// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../../bundle/Bundler.sol";

contract BundlerMock is Bundler {
    function __BundlerMock_init(address bundleExecutorImplementation_) external initializer {
        __Bundler_init(bundleExecutorImplementation_);
    }

    function bundleUp(IBundler.Bundle calldata bundle_) external {
        _bundleUp(bundle_);
    }
}
