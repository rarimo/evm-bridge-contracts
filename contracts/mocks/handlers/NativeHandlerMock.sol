// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../../handlers/NativeHandler.sol";

contract NativeHandlerMock is NativeHandler {
    function withdrawNative(bytes calldata tokenData_, address receiver_) external {
        _withdrawNative(tokenData_, receiver_, false);
    }
}
