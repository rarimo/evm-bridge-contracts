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

    function _bundleUp(bytes32 salt_, bytes calldata bundle_) internal {
        new BundleExecutorProxy{salt: salt_}(
            bundleExecutorImplementation,
            abi.encodeWithSelector(BundleExecutorImplementation.execute.selector, bundle_)
        );
    }

    function determineProxyAddress(bytes32 salt_, bytes calldata bundle_)
        public
        view
        override
        returns (address)
    {
        bytes memory delegateData_ = abi.encodeWithSelector(
            BundleExecutorImplementation.execute.selector,
            bundle_
        );

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
                                        abi.encode(bundleExecutorImplementation, delegateData_)
                                    )
                                )
                            )
                        )
                    )
                )
            );
    }
}
