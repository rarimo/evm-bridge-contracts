// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../interfaces/facade/IBridgeFacade.sol";
import "../interfaces/bridge/IBridge.sol";
import "../interfaces/tokens/IERC20MintableBurnable.sol";
import "../interfaces/tokens/IERC721MintableBurnable.sol";
import "../interfaces/tokens/SBT/ISBT.sol";
import "../interfaces/tokens/IERC1155MintableBurnable.sol";

import "../libs/Encoder.sol";

import "./FeeManager.sol";

contract BridgeFacade is IBridgeFacade, FeeManager {
    using SafeERC20 for *;
    using Encoder for *;

    function __BridgeFacade_init(address bridge_) external initializer {
        __FeeManager_init(bridge_);
    }

    function depositERC20(
        DepositFeeERC20Parameters calldata feeParams_,
        IBridge.DepositERC20Parameters calldata depositParams_
    ) external payable override {
        require(_payCommission(feeParams_.feeToken) == 0, "BridgeFacade: excessive native amount");
        require(depositParams_.token != address(0), "BridgeFacade: zero token");
        require(depositParams_.amount > 0, "BridgeFacade: amount is zero");

        address bridge_ = bridge;

        if (depositParams_.isWrapped) {
            IERC20MintableBurnable(depositParams_.token).burnFrom(
                msg.sender,
                depositParams_.amount
            );
        } else {
            IERC20MintableBurnable(depositParams_.token).safeTransferFrom(
                msg.sender,
                bridge_,
                depositParams_.amount
            );
        }

        IBridge(bridge_).depositERC20(depositParams_);
    }

    function depositERC721(
        DepositFeeERC721Parameters calldata feeParams_,
        IBridge.DepositERC721Parameters calldata depositParams_
    ) external payable override {
        require(_payCommission(feeParams_.feeToken) == 0, "BridgeFacade: excessive native amount");
        require(depositParams_.token != address(0), "BridgeFacade: zero token");

        address bridge_ = bridge;

        if (depositParams_.isWrapped) {
            IERC721MintableBurnable(depositParams_.token).burnFrom(
                msg.sender,
                depositParams_.tokenId
            );
        } else {
            IERC721MintableBurnable(depositParams_.token).safeTransferFrom(
                msg.sender,
                bridge_,
                depositParams_.tokenId
            );
        }

        IBridge(bridge_).depositERC721(depositParams_);
    }

    function depositSBT(
        DepositFeeSBTParameters calldata feeParams_,
        IBridge.DepositSBTParameters calldata depositParams_
    ) external payable override {
        require(_payCommission(feeParams_.feeToken) == 0, "BridgeFacade: excessive native amount");
        require(depositParams_.token != address(0), "BridgeFacade: zero token");
        require(
            ISBT(depositParams_.token).ownerOf(depositParams_.tokenId) == msg.sender,
            "BridgeFacade: invalid token id"
        );

        IBridge(bridge).depositSBT(depositParams_);
    }

    function depositERC1155(
        DepositFeeERC1155Parameters calldata feeParams_,
        IBridge.DepositERC1155Parameters calldata depositParams_
    ) external payable override {
        require(_payCommission(feeParams_.feeToken) == 0, "BridgeFacade: excessive native amount");
        require(depositParams_.token != address(0), "BridgeFacade: zero token");
        require(depositParams_.amount > 0, "BridgeFacade: amount is zero");

        address bridge_ = bridge;

        if (depositParams_.isWrapped) {
            IERC1155MintableBurnable(depositParams_.token).burnFrom(
                msg.sender,
                depositParams_.tokenId,
                depositParams_.amount
            );
        } else {
            IERC1155MintableBurnable(depositParams_.token).safeTransferFrom(
                msg.sender,
                bridge_,
                depositParams_.tokenId,
                depositParams_.amount,
                ""
            );
        }

        IBridge(bridge_).depositERC1155(depositParams_);
    }

    function depositNative(
        DepositFeeNativeParameters calldata feeParams_,
        IBridge.DepositNativeParameters calldata depositParams_
    ) external payable override {
        uint256 amount_ = _payCommission(feeParams_.feeToken);
        require(amount_ > 0, "BridgeFacade: no funds to deposit");

        IBridge(bridge).depositNative{value: amount_}(depositParams_);
    }

    function withdrawERC20(IBridge.WithdrawERC20Parameters calldata params_) external override {
        address bridge_ = bridge;

        IBridge(bridge_).verifyMerkleLeaf(
            params_.getTokenDataLeaf(),
            params_.bundle,
            params_.originHash,
            params_.receiver,
            params_.proof
        );

        if (params_.bundle.bundle.length > 0) {
            try IBridge(bridge_).withdrawERC20Bundle(params_) {
                return;
            } catch {}
        }

        IBridge(bridge_).withdrawERC20(params_);
    }

    function withdrawERC721(IBridge.WithdrawERC721Parameters calldata params_) external override {
        address bridge_ = bridge;

        IBridge(bridge_).verifyMerkleLeaf(
            params_.getTokenDataLeaf(),
            params_.bundle,
            params_.originHash,
            params_.receiver,
            params_.proof
        );

        if (params_.bundle.bundle.length > 0) {
            try IBridge(bridge_).withdrawERC721Bundle(params_) {
                return;
            } catch {}
        }

        IBridge(bridge_).withdrawERC721(params_);
    }

    function withdrawSBT(IBridge.WithdrawSBTParameters calldata params_) external override {
        address bridge_ = bridge;

        IBridge(bridge_).verifyMerkleLeaf(
            params_.getTokenDataLeaf(),
            params_.bundle,
            params_.originHash,
            params_.receiver,
            params_.proof
        );

        if (params_.bundle.bundle.length > 0) {
            try IBridge(bridge_).withdrawSBTBundle(params_) {
                return;
            } catch {}
        }

        IBridge(bridge_).withdrawSBT(params_);
    }

    function withdrawERC1155(
        IBridge.WithdrawERC1155Parameters calldata params_
    ) external override {
        address bridge_ = bridge;

        IBridge(bridge_).verifyMerkleLeaf(
            params_.getTokenDataLeaf(),
            params_.bundle,
            params_.originHash,
            params_.receiver,
            params_.proof
        );

        if (params_.bundle.bundle.length > 0) {
            try IBridge(bridge_).withdrawERC1155Bundle(params_) {
                return;
            } catch {}
        }

        IBridge(bridge_).withdrawERC1155(params_);
    }

    function withdrawNative(IBridge.WithdrawNativeParameters calldata params_) external override {
        address bridge_ = bridge;

        IBridge(bridge_).verifyMerkleLeaf(
            params_.getTokenDataLeaf(),
            params_.bundle,
            params_.originHash,
            params_.receiver,
            params_.proof
        );

        if (params_.bundle.bundle.length > 0) {
            try IBridge(bridge_).withdrawNativeBundle(params_) {
                return;
            } catch {}
        }

        IBridge(bridge_).withdrawNative(params_);
    }
}
