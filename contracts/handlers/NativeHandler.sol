// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../interfaces/handlers/INativeHandler.sol";

import "../bundle/Bundler.sol";

abstract contract NativeHandler is INativeHandler, Bundler {
    function depositNative(
        string calldata receiver_,
        IBundler.Bundle calldata bundle_,
        string calldata network_
    ) external payable override {
        require(msg.value > 0, "NativeHandler: zero value");

        emit DepositedNative(
            msg.value,
            receiver_,
            _encodeSalt(bundle_.salt),
            bundle_.bundle,
            network_
        );
    }

    function withdrawNativeBundle(uint256 amount_, IBundler.Bundle calldata bundle_)
        external
        onlyThis
    {
        address bundleProxy_ = determineProxyAddress(bundle_.salt);

        _withdraw(amount_, bundleProxy_);
        _bundleUp(bundle_);
    }

    function _withdrawNative(
        uint256 amount_,
        address receiver_,
        IBundler.Bundle calldata bundle_
    ) internal {
        require(amount_ > 0, "NativeHandler: amount is zero");

        if (bundle_.bundle.length > 0) {
            try this.withdrawNativeBundle(amount_, bundle_) {
                return;
            } catch {}
        }

        _withdraw(amount_, receiver_);
    }

    function _withdraw(uint256 amount_, address receiver_) private {
        require(receiver_ != address(0), "NativeHandler: receiver is zero");

        (bool success_, ) = payable(receiver_).call{value: amount_}("");

        require(success_, "NativeHandler: failed to send eth");
    }

    function getNativeMerkleLeaf(
        uint256 amount_,
        address receiver_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        string memory chainName_
    ) public view override returns (bytes32) {
        return
            keccak256(
                abi.encode(amount_, receiver_, bundle_, originHash_, chainName_, address(this))
            );
    }
}
