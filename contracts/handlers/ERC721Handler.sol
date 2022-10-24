// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";

import "../interfaces/handlers/IERC721Handler.sol";
import "../interfaces/tokens/IERC721MintableBurnable.sol";

import "../libs/Encoder.sol";

import "../bundle/Bundler.sol";

abstract contract ERC721Handler is IERC721Handler, ERC721Holder, Bundler {
    using Encoder for bytes32;

    function depositERC721(
        address token_,
        uint256 tokenId_,
        IBundler.Bundle calldata bundle_,
        string calldata network_,
        string calldata receiver_,
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
            bundle_.salt.encode(),
            bundle_.bundle,
            network_,
            receiver_,
            isWrapped_
        );
    }

    function withdrawERC721Bundle(
        bytes calldata tokenData_,
        IBundler.Bundle calldata bundle_,
        bool isWrapped_
    ) external onlyThis {
        address bundleProxy_ = determineProxyAddress(bundle_.salt);

        _withdrawERC721(tokenData_, bundleProxy_, isWrapped_);
        _bundleUp(bundle_);
    }

    function _withdrawERC721(bytes calldata tokenData_, address receiver_, bool isWrapped_)
        internal
    {
        (address token_, uint256 tokenId_, string memory tokenURI_) = _decodeERC721TokenData(
            tokenData_
        );

        require(token_ != address(0), "ERC721Handler: zero token");
        require(receiver_ != address(0), "ERC721Handler: zero receiver");

        IERC721MintableBurnable erc721_ = IERC721MintableBurnable(token_);

        if (isWrapped_) {
            erc721_.mintTo(receiver_, tokenId_, tokenURI_);
        } else {
            erc721_.safeTransferFrom(address(this), receiver_, tokenId_);
        }
    }

    function _getERC721TokenDataLeaf(bytes calldata tokenData_)
        internal
        pure
        returns (bytes memory)
    {
        (address token_, uint256 tokenId_, string memory tokenURI_) = _decodeERC721TokenData(
            tokenData_
        );

        return abi.encodePacked(token_, tokenId_, tokenURI_);
    }

    function _decodeERC721TokenData(bytes calldata tokenData_)
        private
        pure
        returns (address, uint256, string memory)
    {
        return abi.decode(tokenData_, (address, uint256, string));
    }
}
