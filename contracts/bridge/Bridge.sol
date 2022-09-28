// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../interfaces/bridge/IBridge.sol";

import "../handlers/ERC20Handler.sol";
import "../handlers/ERC721Handler.sol";
import "../handlers/ERC1155Handler.sol";
import "../handlers/NativeHandler.sol";

import "../bundle/Bundler.sol";

import "./proxy/UUPSSignableUpgradeable.sol";

import "../utils/Signers.sol";
import "../utils/Hashes.sol";

contract Bridge is
    IBridge,
    UUPSSignableUpgradeable,
    Bundler,
    Signers,
    Hashes,
    ERC20Handler,
    ERC721Handler,
    ERC1155Handler,
    NativeHandler
{
    function __Bridge_init(
        address signer_,
        address bundleImplementation_,
        string calldata chainName_
    ) external initializer {
        __Signers_init(signer_, chainName_);
        __Bundler_init(bundleImplementation_);
    }

    function withdrawERC20(
        address token_,
        uint256 amount_,
        address receiver_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        bytes calldata proof_,
        bool isWrapped_
    ) external override {
        bytes32 merkleLeaf_ = getERC20MerkleLeaf(
            token_,
            amount_,
            receiver_,
            bundle_,
            originHash_,
            chainName
        );

        _checkAndUpdateHashes(originHash_);
        _checkMerkleSignature(merkleLeaf_, proof_);

        _withdrawERC20(token_, amount_, receiver_, bundle_, isWrapped_);
    }

    function withdrawERC721(
        address token_,
        uint256 tokenId_,
        string calldata tokenURI_,
        address receiver_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        bytes calldata proof_,
        bool isWrapped_
    ) external override {
        bytes32 merkleLeaf_ = getERC721MerkleLeaf(
            token_,
            tokenId_,
            tokenURI_,
            receiver_,
            bundle_,
            originHash_,
            chainName
        );

        _checkAndUpdateHashes(originHash_);
        _checkMerkleSignature(merkleLeaf_, proof_);

        _withdrawERC721(token_, tokenId_, tokenURI_, receiver_, bundle_, isWrapped_);
    }

    function withdrawERC1155(
        address token_,
        uint256 tokenId_,
        uint256 amount_,
        string calldata tokenURI_,
        address receiver_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        bytes calldata proof_,
        bool isWrapped_
    ) external override {
        bytes32 merkleLeaf_ = getERC1155MerkleLeaf(
            token_,
            tokenId_,
            amount_,
            tokenURI_,
            receiver_,
            bundle_,
            originHash_,
            chainName
        );

        _checkAndUpdateHashes(originHash_);
        _checkMerkleSignature(merkleLeaf_, proof_);

        _withdrawERC1155(token_, tokenId_, amount_, tokenURI_, receiver_, bundle_, isWrapped_);
    }

    function withdrawNative(
        uint256 amount_,
        address receiver_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        bytes calldata proof_
    ) external override {
        bytes32 merkleLeaf_ = getNativeMerkleLeaf(
            amount_,
            receiver_,
            bundle_,
            originHash_,
            chainName
        );

        _checkAndUpdateHashes(originHash_);
        _checkMerkleSignature(merkleLeaf_, proof_);

        _withdrawNative(amount_, receiver_, bundle_);
    }

    function _authorizeUpgrade(address) internal pure override {
        revert("Bridge: this upgrade method is off");
    }

    function _authorizeUpgrade(address newImplementation_, bytes memory signature_)
        internal
        override
        nonceIncrementer
    {
        _checkSignature(_getAddressChangeHash(newImplementation_), signature_);
    }

    function changeSigner(address newSigner_, bytes memory signature_) external nonceIncrementer {
        _checkSignature(_getAddressChangeHash(newSigner_), signature_);

        signer = newSigner_;
    }

    function changeBundleExecutorImplementation(
        address newImplementation_,
        bytes memory signature_
    ) external nonceIncrementer {
        _checkSignature(_getAddressChangeHash(newImplementation_), signature_);

        bundleExecutorImplementation = newImplementation_;
    }
}
