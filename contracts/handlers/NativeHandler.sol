// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../interfaces/handlers/INativeHandler.sol";

import "../bundle/Bundler.sol";

abstract contract NativeHandler is INativeHandler, Bundler {
    function depositNative(
        string calldata receiver_,
        bytes calldata bundle_,
        string calldata network_
    ) external payable override {
        require(msg.value > 0, "NativeHandler: zero value");

        emit DepositedNative(msg.value, receiver_, bundle_, network_);
    }

    function withdrawNativeBundle(
        uint256 amount_,
        bytes calldata bundle_,
        bytes32 salt_
    ) external onlyThis {
        address bundleProxy_ = determineProxyAddress(salt_, bundle_);

        _withdraw(amount_, bundleProxy_);
        _bundleUp(salt_, bundle_);
    }

    function _withdrawNative(
        uint256 amount_,
        address receiver_,
        bytes calldata bundle_,
        bytes32 originHash_
    ) internal {
        require(amount_ > 0, "NativeHandler: amount is zero");

        if (bundle_.length > 0) {
            try this.withdrawNativeBundle(amount_, bundle_, originHash_) {
                return;
            } catch {}
        }

        _withdraw(amount_, receiver_);
    }

    function _withdraw(uint256 amount_, address receiver_) private {
        require(receiver_ != address(0), "NativeHandler: receiver is zero");

        (bool sent_, ) = payable(receiver_).call{value: amount_}("");

        require(sent_, "NativeHandler: can't send eth");
    }

    function getNativeMerkleLeaf(
        uint256 amount_,
        address receiver_,
        bytes calldata bundle_,
        bytes32 originHash_,
        string memory chainName_
    ) public view override returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    amount_,
                    receiver_,
                    bundle_,
                    originHash_,
                    chainName_,
                    address(this)
                )
            );
    }
}
