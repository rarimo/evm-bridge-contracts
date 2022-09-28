// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../interfaces/bundle/IBundler.sol";

import "./proxy/BundleExecutorProxy.sol";
import "./proxy/BundleExecutorImplementation.sol";

import "@openzeppelin/contracts/proxy/utils/Initializable.sol";

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

    function _encodeSalt(bytes32 salt_) internal view returns (bytes32) {
        return keccak256(abi.encode(salt_, msg.sender));
    }
}
