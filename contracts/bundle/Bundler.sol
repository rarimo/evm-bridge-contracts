// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import {IBundler} from "../interfaces/bundle/IBundler.sol";

import {BundleExecutorProxy} from "./proxy/BundleExecutorProxy.sol";
import {BundleExecutorImplementation} from "./proxy/BundleExecutorImplementation.sol";

abstract contract Bundler is IBundler, Initializable {
    address public bundleExecutorImplementation;

    function __Bundler_init(address bundleExecutorImplementation_) public onlyInitializing {
        bundleExecutorImplementation = bundleExecutorImplementation_;
    }

    function _bundleUp(Bundle memory bundle_) internal {
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

    uint256[49] private _gap;
}
