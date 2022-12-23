// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../interfaces/handlers/INativeHandler.sol";

import "../libs/Encoder.sol";

import "../bundle/Bundler.sol";

abstract contract NativeHandler is INativeHandler, Bundler {
    using Encoder for bytes32;

    function depositNative(
        IBundler.Bundle calldata bundle_,
        string calldata network_,
        string calldata receiver_
    ) external payable override {
        require(msg.value > 0, "NativeHandler: zero value");

        emit DepositedNative(
            msg.value,
            bundle_.salt.encode(),
            bundle_.bundle,
            network_,
            receiver_
        );
    }

    receive() external payable {}

    function withdrawNativeBundle(
        bytes calldata tokenData_,
        IBundler.Bundle calldata bundle_,
        bool
    ) external onlyThis {
        address bundleProxy_ = determineProxyAddress(bundle_.salt);

        _withdrawNative(tokenData_, bundleProxy_, false);
        _bundleUp(bundle_);
    }

    function _withdrawNative(bytes calldata tokenData_, address receiver_, bool) internal {
        uint256 amount_ = _decodeNativeTokenData(tokenData_);

        require(receiver_ != address(0), "NativeHandler: receiver is zero");

        (bool success_, ) = payable(receiver_).call{value: amount_}("");

        require(success_, "NativeHandler: failed to send eth");
    }

    function _getNativeTokenDataLeaf(bytes calldata tokenData_)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(_decodeNativeTokenData(tokenData_));
    }

    function _decodeNativeTokenData(bytes calldata tokenData_) private pure returns (uint256) {
        return abi.decode(tokenData_, (uint256));
    }
}
