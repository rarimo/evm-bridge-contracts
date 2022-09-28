// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";

import "../interfaces/handlers/IERC721Handler.sol";
import "../interfaces/tokens/IERC721MintableBurnable.sol";

import "../bundle/Bundler.sol";

abstract contract ERC721Handler is IERC721Handler, ERC721Holder, Bundler {
    function depositERC721(
        address token_,
        uint256 tokenId_,
        string calldata receiver_,
        IBundler.Bundle calldata bundle_,
        string calldata network_,
        bool isWrapped_
    ) external override {
        require(token_ != address(0), "ERC721Handler: zero token");

        IERC721MintableBurnable erc721_ = IERC721MintableBurnable(token_);

        if (isWrapped_) {
            erc721_.burnFrom(msg.sender, tokenId_);
        } else {
            erc721_.safeTransferFrom(msg.sender, address(this), tokenId_);
        }

        emit DepositedERC721(
            token_,
            tokenId_,
            receiver_,
            _encodeSalt(bundle_.salt),
            bundle_.bundle,
            network_,
            isWrapped_
        );
    }

    function withdrawERC721Bundle(
        address token_,
        uint256 tokenId_,
        string calldata tokenURI_,
        IBundler.Bundle calldata bundle_,
        bool isWrapped_
    ) external onlyThis {
        address bundleProxy_ = determineProxyAddress(bundle_.salt);

        _withdraw(token_, tokenId_, tokenURI_, bundleProxy_, isWrapped_);
        _bundleUp(bundle_);
    }

    function _withdrawERC721(
        address token_,
        uint256 tokenId_,
        string calldata tokenURI_,
        address receiver_,
        IBundler.Bundle calldata bundle_,
        bool isWrapped_
    ) internal {
        require(token_ != address(0), "ERC721Handler: zero token");

        if (bundle_.bundle.length > 0) {
            try this.withdrawERC721Bundle(token_, tokenId_, tokenURI_, bundle_, isWrapped_) {
                return;
            } catch {}
        }

        _withdraw(token_, tokenId_, tokenURI_, receiver_, isWrapped_);
    }

    function _withdraw(
        address token_,
        uint256 tokenId_,
        string calldata tokenURI_,
        address receiver_,
        bool isWrapped_
    ) private {
        require(receiver_ != address(0), "ERC721Handler: zero receiver");

        IERC721MintableBurnable erc721_ = IERC721MintableBurnable(token_);

        if (isWrapped_) {
            erc721_.mintTo(receiver_, tokenId_, tokenURI_);
        } else {
            erc721_.safeTransferFrom(address(this), receiver_, tokenId_);
        }
    }

    function getERC721MerkleLeaf(
        address token_,
        uint256 tokenId_,
        string calldata tokenURI_,
        address receiver_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        string memory chainName_
    ) public view override returns (bytes32) {
        return
            keccak256(
                abi.encode(
                    token_,
                    tokenId_,
                    1,
                    tokenURI_,
                    receiver_,
                    bundle_,
                    originHash_,
                    chainName_,
                    address(this)
                )
            );
    }
}
