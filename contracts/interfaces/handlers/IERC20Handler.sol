// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

interface IERC20Handler {
    /**
     * @notice event emits from depositERC20 function
     */
    event DepositedERC20(
        address token,
        uint256 amount,
        string receiver,
        string network,
        bool isWrapped
    );

    /**
     * @notice function for depositing erc20 tokens, emits event DepositedERC20
     * @param token_ the address of deposited token
     * @param amount_ the amount of deposited tokens
     * @param receiver_ the receiver address in destination network, information field for event
     * @param network_ the network name of destination network, information field for event
     * @param isWrapped_ the boolean flag, if true - tokens will burned, false - tokens will transferred
     */
    function depositERC20(
        address token_,
        uint256 amount_,
        string calldata receiver_,
        string calldata network_,
        bool isWrapped_
    ) external;

    /**
     * @notice function for getting the leaf of a merkle tree
     * @param token_ the address of withdrawn token
     * @param amount_ the amount of withdrawn tokens
     * @param receiver_ the receiver address in destination network
     * @param originHash_ the keccak256 hash of abi.encodePacked(origin chain name . origin tx hash . event nonce)
     * @param chainName_ the name of this chain
     * @param verifyingContract_ this contract address
     * @return bytes32 the keccak256 hash of abi.encodePacked concatenation of arguments
     */
    function getERC20MerkleLeaf(
        address token_,
        uint256 amount_,
        address receiver_,
        bytes32 originHash_,
        string memory chainName_,
        address verifyingContract_
    ) external pure returns (bytes32);
}
