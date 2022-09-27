// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../../handlers/NativeHandler.sol";

contract NativeHandlerMock is NativeHandler {
    function withdrawNative(
        uint256 amount_,
        address receiver_,
        bytes calldata bundle_,
        bytes32 originHash_
    ) external {
        _withdrawNative(amount_, receiver_, bundle_, originHash_);
    }
}
