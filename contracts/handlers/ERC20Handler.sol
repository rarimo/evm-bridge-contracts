// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "../interfaces/tokens/IERC20MintableBurnable.sol";
import "../interfaces/handlers/IERC20Handler.sol";

import "../libs/Encoder.sol";

import "../bundle/Bundler.sol";

abstract contract ERC20Handler is IERC20Handler, Bundler {
    using SafeERC20 for IERC20MintableBurnable;
    using Encoder for bytes32;

    function depositERC20(
        address token_,
        uint256 amount_,
        IBundler.Bundle calldata bundle_,
        string calldata network_,
        string calldata receiver_,
        bool isWrapped_
    ) external override {
        require(token_ != address(0), "ERC20Handler: zero token");
        require(amount_ > 0, "ERC20Handler: amount is zero");

        IERC20MintableBurnable erc20_ = IERC20MintableBurnable(token_);

        if (isWrapped_) {
            erc20_.burnFrom(msg.sender, amount_);
        } else {
            erc20_.safeTransferFrom(msg.sender, address(this), amount_);
        }

        emit DepositedERC20(
            token_,
            amount_,
            bundle_.salt.encode(),
            bundle_.bundle,
            network_,
            receiver_,
            isWrapped_
        );
    }

    function withdrawERC20Bundle(
        bytes calldata tokenData_,
        IBundler.Bundle calldata bundle_,
        bool isWrapped_
    ) external onlyThis {
        address bundleProxy_ = determineProxyAddress(bundle_.salt);

        _withdrawERC20(tokenData_, bundleProxy_, isWrapped_);
        _bundleUp(bundle_);
    }

    function _withdrawERC20(
        bytes calldata tokenData_,
        address receiver_,
        bool isWrapped_
    ) internal {
        (address token_, uint256 amount_) = _decodeERC20TokenData(tokenData_);

        require(token_ != address(0), "ERC20Handler: zero token");
        require(receiver_ != address(0), "ERC20Handler: zero receiver");

        IERC20MintableBurnable erc20_ = IERC20MintableBurnable(token_);

        if (isWrapped_) {
            erc20_.mintTo(receiver_, amount_);
        } else {
            erc20_.safeTransfer(receiver_, amount_);
        }
    }

    function _getERC20TokenDataLeaf(
        bytes calldata tokenData_
    ) internal pure returns (bytes memory) {
        (address token_, uint256 amount_) = _decodeERC20TokenData(tokenData_);

        return abi.encodePacked(token_, amount_);
    }

    function _decodeERC20TokenData(
        bytes calldata tokenData_
    ) private pure returns (address, uint256) {
        return abi.decode(tokenData_, (address, uint256));
    }
}
