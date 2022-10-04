// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC1155/utils/ERC1155Holder.sol";

import "../interfaces/handlers/IERC1155Handler.sol";
import "../interfaces/tokens/IERC1155MintableBurnable.sol";

import "../libs/Encoder.sol";

import "../bundle/Bundler.sol";

abstract contract ERC1155Handler is IERC1155Handler, ERC1155Holder, Bundler {
    using Encoder for bytes32;

    function depositERC1155(
        address token_,
        uint256 tokenId_,
        uint256 amount_,
        IBundler.Bundle calldata bundle_,
        string calldata network_,
        string calldata receiver_,
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
            bundle_.salt.encode(),
            bundle_.bundle,
            network_,
            receiver_,
            isWrapped_
        );
    }

    function withdrawERC1155Bundle(
        bytes calldata tokenData_,
        IBundler.Bundle calldata bundle_,
        bool isWrapped_
    ) external onlyThis {
        address bundleProxy_ = determineProxyAddress(bundle_.salt);

        _withdrawERC1155(tokenData_, bundleProxy_, isWrapped_);
        _bundleUp(bundle_);
    }

    function _withdrawERC1155(
        bytes calldata tokenData_,
        address receiver_,
        bool isWrapped_
    ) internal {
        (
            address token_,
            uint256 tokenId_,
            string memory tokenURI_,
            uint256 amount_
        ) = _decodeERC1155TokenData(tokenData_);

        require(token_ != address(0), "ERC1155Handler: zero token");
        require(amount_ > 0, "ERC1155Handler: amount is zero");
        require(receiver_ != address(0), "ERC1155Handler: zero receiver");

        IERC1155MintableBurnable erc1155_ = IERC1155MintableBurnable(token_);

        if (isWrapped_) {
            erc1155_.mintTo(receiver_, tokenId_, amount_, tokenURI_);
        } else {
            erc1155_.safeTransferFrom(address(this), receiver_, tokenId_, amount_, "");
        }
    }

    function _getERC1155TokenDataLeaf(bytes calldata tokenData_)
        internal
        pure
        returns (bytes memory)
    {
        (
            address token_,
            uint256 tokenId_,
            string memory tokenURI_,
            uint256 amount_
        ) = _decodeERC1155TokenData(tokenData_);

        return abi.encodePacked(token_, tokenId_, tokenURI_, amount_);
    }

    function _decodeERC1155TokenData(bytes calldata tokenData_)
        private
        pure
        returns (
            address,
            uint256,
            string memory,
            uint256
        )
    {
        return abi.decode(tokenData_, (address, uint256, string, uint256));
    }
}
