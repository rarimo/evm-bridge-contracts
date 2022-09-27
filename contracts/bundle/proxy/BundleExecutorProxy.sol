// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract BundleExecutorProxy {
    constructor(address implementation_, bytes memory delegateData_) {
        (bool success_, ) = implementation_.delegatecall(delegateData_);

        if (!success_) {
            assembly {
                let ptr_ := mload(0x40)
                let size_ := returndatasize()
                returndatacopy(ptr_, 0, size_)
                revert(ptr_, size_)
            }
        }

        selfdestruct(payable(msg.sender));
    }

    receive() external payable {}
}
