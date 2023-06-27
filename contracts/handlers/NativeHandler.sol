// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../interfaces/handlers/INativeHandler.sol";

import "../libs/Encoder.sol";

import "../bundle/Bundler.sol";

abstract contract NativeHandler is INativeHandler, Bundler {
    using Encoder for bytes32;

    function depositNative(
        DepositNativeParameters calldata params_
    ) external payable override onlyFacade {
        emit DepositedNative(
            msg.value,
            params_.bundle.salt.encode(),
            params_.bundle.bundle,
            params_.network,
            params_.receiver
        );
    }

    receive() external payable {}

    function withdrawNative(WithdrawNativeParameters memory params_) public override onlyFacade {
        require(params_.receiver != address(0), "NativeHandler: receiver is zero");

        (bool success_, ) = params_.receiver.call{value: params_.amount}("");
        require(success_, "NativeHandler: failed to send eth");
    }

    function withdrawNativeBundle(
        WithdrawNativeParameters memory params_
    ) external override onlyFacade {
        params_.receiver = determineProxyAddress(params_.bundle.salt);

        withdrawNative(params_);
        _bundleUp(params_.bundle);
    }

    uint256[50] private _gap;
}
