// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract BundleExecutorProxy {
    address internal immutable _IMPLEMENTATION;
    address internal immutable _BRIDGE;

    constructor(address implementation_, address bridge_) {
        _IMPLEMENTATION = implementation_;
        _BRIDGE = bridge_;
    }

    fallback() external payable {
        _delegate(_IMPLEMENTATION);
    }

    function destroy() external {
        address bridge_ = _BRIDGE;

        assembly {
            if iszero(eq(caller(), bridge_)) {
                revert(0, 0)
            }

            selfdestruct(caller())
        }
    }

    function _delegate(address implementation_) internal {
        assembly {
            calldatacopy(0, 0, calldatasize())

            let result_ := delegatecall(gas(), implementation_, 0, calldatasize(), 0, 0)

            returndatacopy(0, 0, returndatasize())

            switch result_
            case 0 {
                revert(0, returndatasize())
            }
            default {
                return(0, returndatasize())
            }
        }
    }
}
