// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";

import "../interfaces/handlers/IERC721Handler.sol";
import "../interfaces/tokens/IERC721MintableBurnable.sol";

abstract contract ERC721Handler is IERC721Handler, ERC721Holder {
    function depositERC721(
        address token_,
        uint256 tokenId_,
        string calldata receiver_,
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

        emit DepositedERC721(token_, tokenId_, receiver_, network_, isWrapped_);
    }

    function _withdrawERC721(
        address token_,
        uint256 tokenId_,
        address receiver_,
        bool isWrapped_
    ) internal {
        require(token_ != address(0), "ERC721Handler: zero token");
        require(receiver_ != address(0), "ERC721Handler: zero receiver");

        IERC721MintableBurnable erc721_ = IERC721MintableBurnable(token_);

        if (isWrapped_) {
            erc721_.mintTo(receiver_, tokenId_);
        } else {
            erc721_.safeTransferFrom(address(this), receiver_, tokenId_);
        }
    }

    function getERC721MerkleLeaf(
        address token_,
        uint256 tokenId_,
        uint256 amount_,
        address receiver_,
        bytes32 originHash_,
        string memory chainName_,
        address verifyingContract_
    ) public pure override returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    token_,
                    tokenId_,
                    amount_,
                    receiver_,
                    originHash_,
                    chainName_,
                    verifyingContract_
                )
            );
    }
}
