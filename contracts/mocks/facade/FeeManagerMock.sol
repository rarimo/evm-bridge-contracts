// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {FeeManager} from "../../facade/FeeManager.sol";

contract FeeManagerMock is FeeManager {
    function __FeeManagerMock_init(address bridge_) public initializer {
        __FeeManager_init(bridge_);
    }
}
