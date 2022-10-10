// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/utils/Initializable.sol";

import "../interfaces/bundle/IBundler.sol";

import "./proxy/BundleExecutorProxy.sol";
import "./proxy/BundleExecutorImplementation.sol";

abstract contract Bundler is IBundler, Initializable {
    address public bundleExecutorImplementation;

    modifier onlyThis() {
        require(msg.sender == address(this), "Bundler: not this");
        _;
    }

    function __Bundler_init(address bundleExecutorImplementation_) public onlyInitializing {
        bundleExecutorImplementation = bundleExecutorImplementation_;
    }

    function _bundleUp(Bundle calldata bundle_) internal {
        address payable executor = payable(
            new BundleExecutorProxy{salt: bundle_.salt}(
                bundleExecutorImplementation,
                address(this)
            )
        );

        BundleExecutorImplementation(executor).execute(bundle_.bundle);
        BundleExecutorProxy(executor).destroy();
    }

    function determineProxyAddress(bytes32 salt_) public view override returns (address) {
        return
            address(
                uint160(
                    uint256(
                        keccak256(
                            abi.encodePacked(
                                bytes1(0xff),
                                address(this),
                                salt_,
                                keccak256(
                                    abi.encodePacked(
                                        type(BundleExecutorProxy).creationCode,
                                        abi.encode(bundleExecutorImplementation, address(this))
                                    )
                                )
                            )
                        )
                    )
                )
            );
    }
}
