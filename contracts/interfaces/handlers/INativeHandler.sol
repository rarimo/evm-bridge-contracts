// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

interface INativeHandler {
    event DepositedNative(uint256 tokenAmount, string receiver, string network, address token_);

    function depositNative(
        string calldata receiver_,
        string calldata network_,
        address token_
    ) external payable;
}
