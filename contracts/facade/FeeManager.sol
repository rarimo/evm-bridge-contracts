// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "../interfaces/facade/IFeeManager.sol";
import "../interfaces/utils/ISigners.sol";

import "../bridge/proxy/UUPSSignableUpgradeable.sol";

import "../libs/Constants.sol";

abstract contract FeeManager is IFeeManager, UUPSSignableUpgradeable, Initializable {
    using SafeERC20 for IERC20;

    address public bridge;
    mapping(address => uint256) internal _feeTokens;

    function __FeeManager_init(address bridge_) public onlyInitializing {
        bridge = bridge_;
    }

    function addFeeToken(AddFeeTokenParameters calldata params_) external override {
        require(
            params_.feeTokens.length == params_.feeAmounts.length,
            "FeeManager: params lengths mismatch"
        );

        _validateTokenManagementSignature(
            uint8(MethodId.AddFeeToken),
            params_.feeTokens,
            params_.feeAmounts,
            params_.signature
        );

        for (uint256 i = 0; i < params_.feeTokens.length; ++i) {
            _addFeeToken(params_.feeTokens[i], params_.feeAmounts[i]);
        }
    }

    function removeFeeToken(RemoveFeeTokenParameters calldata params_) external override {
        require(
            params_.feeTokens.length == params_.feeAmounts.length,
            "FeeManager: params lengths mismatch"
        );

        _validateTokenManagementSignature(
            uint8(MethodId.RemoveFeeToken),
            params_.feeTokens,
            params_.feeAmounts,
            params_.signature
        );

        for (uint256 i = 0; i < params_.feeTokens.length; ++i) {
            _removeFeeToken(params_.feeTokens[i], params_.feeAmounts[i]);
        }
    }

    function updateFeeToken(UpdateFeeTokenParameters calldata params_) external override {
        require(
            params_.feeTokens.length == params_.feeAmounts.length,
            "FeeManager: params lengths mismatch"
        );

        _validateTokenManagementSignature(
            uint8(MethodId.UpdateFeeToken),
            params_.feeTokens,
            params_.feeAmounts,
            params_.signature
        );

        for (uint256 i = 0; i < params_.feeTokens.length; ++i) {
            _updateFeeToken(params_.feeTokens[i], params_.feeAmounts[i]);
        }
    }

    function withdrawFeeToken(WithdrawFeeTokenParameters calldata params_) external override {
        require(
            params_.feeTokens.length == params_.amounts.length,
            "FeeManager: params lengths mismatch"
        );

        _validateWithdrawSignature(params_);

        for (uint256 i = 0; i < params_.feeTokens.length; ++i) {
            _withdrawFeeToken(params_.receiver, params_.feeTokens[i], params_.amounts[i]);
        }
    }

    function getCommission(
        address feeToken_
    ) external view override returns (uint256 commission_) {
        return _feeTokens[feeToken_];
    }

    function _payCommission(address feeToken_) internal returns (uint256 nativeLeft_) {
        nativeLeft_ = msg.value;

        uint256 commission_ = _feeTokens[feeToken_];
        require(commission_ != 0, "FeeManager: wrong fee token");

        if (feeToken_ == Constants.ETHEREUM_ADDRESS) {
            require(commission_ <= nativeLeft_, "FeeManager: wrong native amount");

            nativeLeft_ -= commission_;
        } else if (feeToken_ != Constants.COMMISSION_ADDRESS) {
            IERC20(feeToken_).safeTransferFrom(msg.sender, address(this), commission_);
        }
    }

    function _authorizeUpgrade(address) internal pure override {
        revert("FeeManager: this upgrade method is off");
    }

    function _authorizeUpgrade(
        address newImplementation_,
        bytes calldata signature_
    ) internal override {
        require(newImplementation_ != address(0), "FeeManager: zero address");

        ISigners(bridge).validateChangeAddressSignature(
            uint8(MethodId.AuthorizeUpgrade),
            address(this),
            newImplementation_,
            signature_
        );
    }

    function _validateTokenManagementSignature(
        uint8 methodId_,
        address[] calldata feeTokens_,
        uint256[] calldata feeAmounts_,
        bytes calldata signature_
    ) private {
        address bridge_ = bridge;

        (string memory chainName_, uint256 nonce_) = ISigners(bridge_).getSigComponents(
            methodId_,
            address(this)
        );

        bytes32 signHash_ = keccak256(
            abi.encodePacked(nonce_, address(this), chainName_, methodId_, feeTokens_, feeAmounts_)
        );

        ISigners(bridge_).checkSignatureAndIncrementNonce(
            methodId_,
            address(this),
            signHash_,
            signature_
        );
    }

    function _validateWithdrawSignature(WithdrawFeeTokenParameters calldata params_) private {
        address bridge_ = bridge;
        uint8 methodId_ = uint8(MethodId.WithdrawFeeToken);

        (string memory chainName_, uint256 nonce_) = ISigners(bridge_).getSigComponents(
            methodId_,
            address(this)
        );

        bytes32 signHash_ = keccak256(
            abi.encodePacked(
                nonce_,
                params_.receiver,
                address(this),
                chainName_,
                methodId_,
                params_.feeTokens,
                params_.amounts
            )
        );

        ISigners(bridge_).checkSignatureAndIncrementNonce(
            methodId_,
            address(this),
            signHash_,
            params_.signature
        );
    }

    function _addFeeToken(address feeToken_, uint256 feeAmount_) private {
        require(_feeTokens[feeToken_] == 0, "FeeManager: token already exists");

        _feeTokens[feeToken_] = feeAmount_;

        emit AddedFeeToken(feeToken_, feeAmount_);
    }

    function _removeFeeToken(address feeToken_, uint256 feeAmount_) private {
        require(_feeTokens[feeToken_] == feeAmount_, "FeeManager: wrong token address or amount");

        delete _feeTokens[feeToken_];

        emit RemovedFeeToken(feeToken_, feeAmount_);
    }

    function _updateFeeToken(address feeToken_, uint256 feeAmount_) private {
        require(_feeTokens[feeToken_] != 0, "FeeManager: token doesn't exist");

        _feeTokens[feeToken_] = feeAmount_;

        emit UpdatedFeeToken(feeToken_, feeAmount_);
    }

    function _withdrawFeeToken(address receiver_, address feeToken_, uint256 amount_) private {
        require(feeToken_ != Constants.COMMISSION_ADDRESS, "FeeManager: commission token");

        if (feeToken_ == Constants.ETHEREUM_ADDRESS) {
            (bool ok_, ) = receiver_.call{value: amount_}("");
            require(ok_, "FeeManager: failed to withdraw native");
        } else {
            IERC20(feeToken_).safeTransfer(receiver_, amount_);
        }

        emit WithdrawnFeeToken(receiver_, feeToken_, amount_);
    }

    uint256[48] private _gap;
}
