// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "../interfaces/tokens/IERC20MintableBurnable.sol";
import "../interfaces/handlers/IERC20Handler.sol";

import "../bundle/Bundler.sol";

abstract contract ERC20Handler is IERC20Handler, Bundler {
    using SafeERC20 for IERC20MintableBurnable;

    function depositERC20(
        address token_,
        uint256 amount_,
        string calldata receiver_,
        IBundler.Bundle calldata bundle_,
        string calldata network_,
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
            receiver_,
            _encodeSalt(bundle_.salt),
            bundle_.bundle,
            network_,
            isWrapped_
        );
    }

    function withdrawERC20Bundle(
        address token_,
        uint256 amount_,
        IBundler.Bundle calldata bundle_,
        bool isWrapped_
    ) external onlyThis {
        address bundleProxy_ = determineProxyAddress(bundle_.salt);

        _withdraw(token_, amount_, bundleProxy_, isWrapped_);
        _bundleUp(bundle_);
    }

    function _withdrawERC20(
        address token_,
        uint256 amount_,
        address receiver_,
        IBundler.Bundle calldata bundle_,
        bool isWrapped_
    ) internal {
        require(token_ != address(0), "ERC20Handler: zero token");
        require(amount_ > 0, "ERC20Handler: amount is zero");

        if (bundle_.bundle.length > 0) {
            try this.withdrawERC20Bundle(token_, amount_, bundle_, isWrapped_) {
                return;
            } catch {}
        }

        _withdraw(token_, amount_, receiver_, isWrapped_);
    }

    function _withdraw(
        address token_,
        uint256 amount_,
        address receiver_,
        bool isWrapped_
    ) private {
        require(receiver_ != address(0), "ERC20Handler: zero receiver");

        IERC20MintableBurnable erc20_ = IERC20MintableBurnable(token_);

        if (isWrapped_) {
            erc20_.mintTo(receiver_, amount_);
        } else {
            erc20_.safeTransfer(receiver_, amount_);
        }
    }

    function getERC20MerkleLeaf(
        address token_,
        uint256 amount_,
        address receiver_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        string memory chainName_
    ) public view override returns (bytes32) {
        return
            keccak256(
                abi.encode(
                    token_,
                    amount_,
                    receiver_,
                    bundle_,
                    originHash_,
                    chainName_,
                    address(this)
                )
            );
    }
}
