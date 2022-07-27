// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../interfaces/handlers/INativeHandler.sol";

abstract contract NativeHandler is INativeHandler {
    function depositNative(string calldata receiver_, string calldata network_)
        external
        payable
        override
    {
        // TODO
    }
}
