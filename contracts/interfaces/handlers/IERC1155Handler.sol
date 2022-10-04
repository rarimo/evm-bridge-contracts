// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../bundle/IBundler.sol";

interface IERC1155Handler is IBundler {
    /**
     * @notice event emits from depositERC1155 function
     */
    event DepositedERC1155(
        address token,
        uint256 tokenId,
        uint256 amount,
        bytes32 salt,
        bytes bundle,
        string network,
        string receiver,
        bool isWrapped
    );

    /**
     * @notice function for depositing erc1155 tokens, emits event DepositedERC115
     * @param token_ the address of deposited tokens
     * @param tokenId_ the id of deposited tokens
     * @param amount_ the amount of deposited tokens
     * @param bundle_ the encoded transaction bundle with salt
     * @param network_ the network name of destination network, information field for event
     * @param receiver_ the receiver address in destination network, information field for event
     * @param isWrapped_ the boolean flag, if true - tokens will burned, false - tokens will transferred
     */
    function depositERC1155(
        address token_,
        uint256 tokenId_,
        uint256 amount_,
        IBundler.Bundle calldata bundle_,
        string calldata network_,
        string calldata receiver_,
        bool isWrapped_
    ) external;
}
