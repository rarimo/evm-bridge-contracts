// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../bundle/IBundler.sol";

interface IERC20Handler is IBundler {
    /**
     * @notice event emits from depositERC20 function
     */
    event DepositedERC20(
        address token,
        uint256 amount,
        string receiver,
        bytes32 salt,
        bytes bundle,
        string network,
        bool isWrapped
    );

    /**
     * @notice function for depositing erc20 tokens, emits event DepositedERC20
     * @param token_ the address of deposited token
     * @param amount_ the amount of deposited tokens
     * @param receiver_ the receiver address in destination network, information field for event
     * @param bundle_ the encoded transaction bundle with salt
     * @param network_ the network name of destination network, information field for event
     * @param isWrapped_ the boolean flag, if true - tokens will burned, false - tokens will transferred
     */
    function depositERC20(
        address token_,
        uint256 amount_,
        string calldata receiver_,
        IBundler.Bundle calldata bundle_,
        string calldata network_,
        bool isWrapped_
    ) external;

    /**
     * @notice function for getting the leaf of a merkle tree
     * @param token_ the address of withdrawn token
     * @param amount_ the amount of withdrawn tokens
     * @param receiver_ the receiver address in destination network
     * @param bundle_ the encoded transaction bundle with encoded salt
     * @param originHash_ the keccak256 hash of abi.encodePacked(origin chain name . origin tx hash . event nonce)
     * @param chainName_ the name of this chain
     * @return bytes32 the keccak256 hash of abi.encodePacked concatenation of arguments + address(this)
     */
    function getERC20MerkleLeaf(
        address token_,
        uint256 amount_,
        address receiver_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        string memory chainName_
    ) external view returns (bytes32);
}
