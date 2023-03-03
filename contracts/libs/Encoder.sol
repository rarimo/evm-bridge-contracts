// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../interfaces/bundle/IBundler.sol";

library Encoder {
    function encode(bytes32 salt_) internal view returns (bytes32) {
        return keccak256(abi.encodePacked(salt_, tx.origin));
    }

    function encode(
        function(bytes calldata) internal pure returns (bytes memory) _getTokenDataLeaf,
        bytes calldata tokenData_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        string memory chainName_,
        address receiver_
    ) internal view returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    _getTokenDataLeaf(tokenData_),
                    _getBundleLeaf(bundle_),
                    _getMetadataLeaf(originHash_, chainName_, receiver_)
                )
            );
    }

    function _getBundleLeaf(IBundler.Bundle calldata bundle_) private pure returns (bytes memory) {
        if (bundle_.bundle.length == 0) {
            return bytes("");
        }

        return abi.encodePacked(bundle_.salt, bundle_.bundle);
    }

    function _getMetadataLeaf(
        bytes32 originHash_,
        string memory chainName_,
        address receiver_
    ) private view returns (bytes memory) {
        return abi.encodePacked(originHash_, chainName_, receiver_, address(this));
    }
}
