// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC1155/utils/ERC1155Holder.sol";

import "../interfaces/handlers/IERC1155Handler.sol";
import "../interfaces/tokens/IERC1155MintableBurnable.sol";

import "../bundle/Bundler.sol";

abstract contract ERC1155Handler is IERC1155Handler, ERC1155Holder, Bundler {
    function depositERC1155(
        address token_,
        uint256 tokenId_,
        uint256 amount_,
        string calldata receiver_,
        IBundler.Bundle calldata bundle_,
        string calldata network_,
        bool isWrapped_
    ) external override {
        require(token_ != address(0), "ERC1155Handler: zero token");
        require(amount_ > 0, "ERC1155Handler: amount is zero");

        IERC1155MintableBurnable erc1155_ = IERC1155MintableBurnable(token_);

        if (isWrapped_) {
            erc1155_.burnFrom(msg.sender, tokenId_, amount_);
        } else {
            erc1155_.safeTransferFrom(msg.sender, address(this), tokenId_, amount_, "");
        }

        emit DepositedERC1155(
            token_,
            tokenId_,
            amount_,
            receiver_,
            _encodeSalt(bundle_.salt),
            bundle_.bundle,
            network_,
            isWrapped_
        );
    }

    function withdrawERC1155Bundle(
        address token_,
        uint256 tokenId_,
        uint256 amount_,
        string calldata tokenURI_,
        IBundler.Bundle calldata bundle_,
        bool isWrapped_
    ) external onlyThis {
        address bundleProxy_ = determineProxyAddress(bundle_.salt);

        _withdraw(token_, tokenId_, amount_, tokenURI_, bundleProxy_, isWrapped_);
        _bundleUp(bundle_);
    }

    function _withdrawERC1155(
        address token_,
        uint256 tokenId_,
        uint256 amount_,
        string calldata tokenURI_,
        address receiver_,
        IBundler.Bundle calldata bundle_,
        bool isWrapped_
    ) internal {
        require(token_ != address(0), "ERC1155Handler: zero token");
        require(amount_ > 0, "ERC1155Handler: amount is zero");

        if (bundle_.bundle.length > 0) {
            try
                this.withdrawERC1155Bundle(
                    token_,
                    tokenId_,
                    amount_,
                    tokenURI_,
                    bundle_,
                    isWrapped_
                )
            {
                return;
            } catch {}
        }

        _withdraw(token_, tokenId_, amount_, tokenURI_, receiver_, isWrapped_);
    }

    function _withdraw(
        address token_,
        uint256 tokenId_,
        uint256 amount_,
        string calldata tokenURI_,
        address receiver_,
        bool isWrapped_
    ) private {
        require(receiver_ != address(0), "ERC1155Handler: zero receiver");

        IERC1155MintableBurnable erc1155_ = IERC1155MintableBurnable(token_);

        if (isWrapped_) {
            erc1155_.mintTo(receiver_, tokenId_, amount_, tokenURI_);
        } else {
            erc1155_.safeTransferFrom(address(this), receiver_, tokenId_, amount_, "");
        }
    }

    function getERC1155MerkleLeaf(
        address token_,
        uint256 tokenId_,
        uint256 amount_,
        string calldata tokenURI_,
        address receiver_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        string memory chainName_
    ) public view override returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    token_,
                    tokenId_,
                    amount_,
                    tokenURI_,
                    receiver_,
                    bundle_.salt,
                    bundle_.bundle,
                    originHash_,
                    chainName_,
                    address(this)
                )
            );
    }
}
