// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

interface IERC721Handler {
    /**
     * @notice event emits from depositERC721 function
     */
    event DepositedERC721(
        address token,
        uint256 tokenId,
        string receiver,
        string network,
        bool isWrapped
    );

    /**
     * @notice function for depositing erc721 tokens, emits event DepositedERC721
     * @param token_ the address of deposited token
     * @param tokenId_ the id of deposited token
     * @param receiver_ the receiver address in destination network, information field for event
     * @param network_ the network name of destination network, information field for event
     * @param isWrapped_ the boolean flag, if true - token will burned, false - token will transferred
     */
    function depositERC721(
        address token_,
        uint256 tokenId_,
        string calldata receiver_,
        string calldata network_,
        bool isWrapped_
    ) external;

    /**
     * @notice function for getting the leaf of a merkle tree
     * @param token_ the address of withdrawn token
     * @param tokenId_ the id of deposited token
     * @param amount_ should always equal 1
     * @param receiver_ the receiver address in destination network
     * @param originHash_ the keccak256 hash of abi.encodePacked(origin chain name . origin tx hash . event nonce)
     * @param chainName_ the name of this chain
     * @param verifyingContract_ this contract address
     * @return bytes32 the keccak256 hash of abi.encodePacked concatenation of arguments
     */
    function getERC721MerkleLeaf(
        address token_,
        uint256 tokenId_,
        uint256 amount_,
        address receiver_,
        bytes32 originHash_,
        string memory chainName_,
        address verifyingContract_
    ) external pure returns (bytes32);
}
