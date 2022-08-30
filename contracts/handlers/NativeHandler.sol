// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../interfaces/handlers/INativeHandler.sol";

abstract contract NativeHandler is INativeHandler {
    function depositNative(string calldata receiver_, string calldata network_)
        external
        payable
        override
    {
        require(msg.value > 0, "NativeHandler: zero value");

        emit DepositedNative(msg.value, receiver_, network_);
    }

    function _withdrawNative(uint256 amount_, address receiver_) internal {
        require(amount_ > 0, "NativeHandler: amount is zero");
        require(receiver_ != address(0), "NativeHandler: receiver is zero");

        (bool sent_, ) = payable(receiver_).call{value: amount_}("");

        require(sent_, "NativeHandler: can't send eth");
    }

    function getNativeMerkleLeaf(
        uint256 amount_,
        address receiver_,
        bytes32 originHash_,
        string memory chainName_,
        address verifyingContract_
    ) public pure override returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(amount_, receiver_, originHash_, chainName_, verifyingContract_)
            );
    }
}
