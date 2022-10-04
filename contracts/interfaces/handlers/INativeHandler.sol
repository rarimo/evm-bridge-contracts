// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../bundle/IBundler.sol";

interface INativeHandler is IBundler {
    /**
     * @notice event emits from depositNative function
     */
    event DepositedNative(
        uint256 amount,
        bytes32 salt,
        bytes bundle,
        string network,
        string receiver
    );

    /**
     * @notice function for depositing native currency, emits event DepositedNative
     * @param bundle_ the encoded transaction bundle with salt
     * @param network_ the network name of destination network, information field for event
     * @param receiver_ the receiver address in destination network, information field for event
     */
    function depositNative(
        IBundler.Bundle calldata bundle_,
        string calldata network_,
        string calldata receiver_
    ) external payable;
}
