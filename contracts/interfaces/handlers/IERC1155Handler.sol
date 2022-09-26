// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

interface IERC1155Handler {
    /**
     * @notice event emits from depositERC1155 function
     */
    event DepositedERC1155(
        address token,
        uint256 tokenId,
        uint256 amount,
        string receiver,
        string network,
        bool isWrapped
    );

    /**
     * @notice function for depositing erc1155 tokens, emits event DepositedERC115
     * @param token_ the address of deposited tokens
     * @param tokenId_ the id of deposited tokens
     * @param amount_ the amount of deposited tokens
     * @param receiver_ the receiver address in destination network, information field for event
     * @param network_ the network name of destination network, information field for event
     * @param isWrapped_ the boolean flag, if true - tokens will burned, false - tokens will transferred
     */
    function depositERC1155(
        address token_,
        uint256 tokenId_,
        uint256 amount_,
        string calldata receiver_,
        string calldata network_,
        bool isWrapped_
    ) external;

    /**
     * @notice function for getting the leaf of a merkle tree
     * @param token_ the address of withdrawn token
     * @param tokenId_ the id of deposited token
     * @param amount_ the amount of withdrawn tokens
     * @param tokenURI_ the token metadata URI or token index if base URI is set
     * @param receiver_ the receiver address in destination network
     * @param originHash_ the keccak256 hash of abi.encodePacked(origin chain name . origin tx hash . event nonce)
     * @param chainName_ the name of this chain
     * @return bytes32 the keccak256 hash of abi.encodePacked concatenation of arguments + address(this)
     */
    function getERC1155MerkleLeaf(
        address token_,
        uint256 tokenId_,
        uint256 amount_,
        string memory tokenURI_,
        address receiver_,
        bytes32 originHash_,
        string memory chainName_
    ) external view returns (bytes32);
}
