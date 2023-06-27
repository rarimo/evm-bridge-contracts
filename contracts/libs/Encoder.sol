// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../interfaces/bridge/IBridge.sol";
import "../interfaces/bundle/IBundler.sol";

library Encoder {
    function encode(bytes32 salt_) internal view returns (bytes32) {
        return keccak256(abi.encodePacked(salt_, tx.origin));
    }

    function encode(
        bytes calldata tokenDataLeaf_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        string memory chainName_,
        address receiver_
    ) internal view returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    tokenDataLeaf_,
                    _getBundleLeaf(bundle_),
                    _getMetadataLeaf(originHash_, chainName_, receiver_)
                )
            );
    }

    function getTokenDataLeaf(
        IBridge.WithdrawERC20Parameters calldata params_
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(params_.token, params_.amount);
    }

    function getTokenDataLeaf(
        IBridge.WithdrawERC721Parameters calldata params_
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(params_.token, params_.tokenId, params_.tokenURI);
    }

    function getTokenDataLeaf(
        IBridge.WithdrawSBTParameters calldata params_
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(params_.token, params_.tokenId, params_.tokenURI);
    }

    function getTokenDataLeaf(
        IBridge.WithdrawERC1155Parameters calldata params_
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(params_.token, params_.tokenId, params_.tokenURI, params_.amount);
    }

    function getTokenDataLeaf(
        IBridge.WithdrawNativeParameters calldata params_
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(params_.amount);
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
