// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

interface INativeHandler {
    /**
     * @notice event emits from depositNative function
     */
    event DepositedNative(uint256 amount, string receiver, bytes bundle, string network);

    /**
     * @notice function for depositing native currency, emits event DepositedNative
     * @param receiver_ the receiver address in destination network, information field for event
     * @param bundle_ the encoded transaction bundle
     * @param network_ the network name of destination network, information field for event
     */
    function depositNative(
        string calldata receiver_,
        bytes calldata bundle_,
        string calldata network_
    ) external payable;

    /**
     * @notice function for getting the leaf of a merkle tree
     * @param amount_ the amount of withdrawn tokens
     * @param receiver_ the receiver address in destination network
     * @param bundle_ the encoded transaction bundle
     * @param originHash_ the keccak256 hash of abi.encodePacked(origin chain name . origin tx hash . event nonce)
     * @param chainName_ the name of this chain
     * @return bytes32 the keccak256 hash of abi.encodePacked concatenation of arguments + address(this)
     */
    function getNativeMerkleLeaf(
        uint256 amount_,
        address receiver_,
        bytes calldata bundle_,
        bytes32 originHash_,
        string memory chainName_
    ) external view returns (bytes32);
}
