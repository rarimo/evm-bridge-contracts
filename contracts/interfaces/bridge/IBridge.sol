// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

interface IBridge {
    function withdrawERC20(
        address token_,
        uint256 amount_,
        bytes32 txHash_,
        uint256 txNonce_,
        bool isWrapped,
        bytes[] calldata signatures_
    ) external;

    function withdrawERC721(
        address token_,
        uint256 tokenId_,
        bytes32 txHash_,
        uint256 txNonce_,
        bool isWrapped,
        bytes[] calldata signatures_
    ) external;

    function withdrawERC1155(
        address token_,
        uint256 tokenId_,
        uint256 amount_,
        bytes32 txHash_,
        uint256 txNonce_,
        bool isWrapped,
        bytes[] calldata signatures_
    ) external;

    function withdrawNative(
        uint256 amount_,
        bytes32 txHash_,
        uint256 txNonce_,
        bool isWrapped,
        bytes[] calldata signatures_
    ) external;
}
