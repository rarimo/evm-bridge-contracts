// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

interface IBundler {
    /**
     * @notice function to get the bundle executor proxy address for the given salt and bundle
     * @param salt_ the salt for create2 (origin hash)
     * @param bundle_ the tx bundle to execute
     * @return the bundle executor proxy address
     */
    function determineProxyAddress(bytes32 salt_, bytes calldata bundle_)
        external
        view
        returns (address);
}
