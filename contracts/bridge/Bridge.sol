// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@solarity/solidity-lib/access-control/MultiOwnable.sol";

import "../interfaces/bridge/IBridge.sol";

import "../handlers/ERC20Handler.sol";
import "../handlers/ERC721Handler.sol";
import "../handlers/SBTHandler.sol";
import "../handlers/ERC1155Handler.sol";
import "../handlers/NativeHandler.sol";

import "../bundle/Bundler.sol";

import "./proxy/UUPSSignableUpgradeable.sol";

import "../utils/Signers.sol";
import "../utils/Hashes.sol";

import "../libs/Encoder.sol";

contract Bridge is
    IBridge,
    MultiOwnable,
    UUPSSignableUpgradeable,
    Bundler,
    Signers,
    Hashes,
    ERC20Handler,
    ERC721Handler,
    SBTHandler,
    ERC1155Handler,
    NativeHandler
{
    using Encoder for *;

    function __Bridge_init(
        address signer_,
        address bundleImplementation_,
        string calldata chainName_,
        address facade_
    ) external initializer {
        __Signers_init(signer_, chainName_);
        __Bundler_init(bundleImplementation_, facade_);
        __MultiOwnable_init();
    }

    function verifyMerkleLeaf(
        bytes calldata tokenDataLeaf_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        address receiver_,
        bytes calldata proof_
    ) external override onlyFacade {
        bytes32 merkleLeaf_ = tokenDataLeaf_.encode(bundle_, originHash_, chainName, receiver_);

        _checkAndUpdateHashes(originHash_);
        _checkMerkleSignature(merkleLeaf_, proof_);
    }

    function changeSignerOwner(bytes calldata newSignerPubKey_) external onlyOwner {
        signer = _convertPubKeyToAddress(newSignerPubKey_);
    }

    function changeBundleExecutorImplementationOwner(
        address newImplementation_
    ) external onlyOwner {
        require(newImplementation_ != address(0), "Bridge: zero address");

        bundleExecutorImplementation = newImplementation_;
    }

    function changeFacadeOwner(address newFacade_) external onlyOwner {
        require(newFacade_ != address(0), "Bridge: zero address");

        facade = newFacade_;
    }

    function changeSigner(bytes calldata newSignerPubKey_, bytes calldata signature_) external {
        _checkSignature(keccak256(newSignerPubKey_), signature_);

        signer = _convertPubKeyToAddress(newSignerPubKey_);
    }

    function changeBundleExecutorImplementation(
        address newImplementation_,
        bytes calldata signature_
    ) external {
        require(newImplementation_ != address(0), "Bridge: zero address");

        validateChangeAddressSignature(
            uint8(MethodId.ChangeBundleExecutorImplementation),
            address(this),
            newImplementation_,
            signature_
        );

        bundleExecutorImplementation = newImplementation_;
    }

    function changeFacade(address newFacade_, bytes calldata signature_) external {
        require(newFacade_ != address(0), "Bridge: zero address");

        validateChangeAddressSignature(
            uint8(MethodId.ChangeFacade),
            address(this),
            newFacade_,
            signature_
        );

        facade = newFacade_;
    }

    function _authorizeUpgrade(address) internal view override onlyOwner {}

    function _authorizeUpgrade(
        address newImplementation_,
        bytes calldata signature_
    ) internal override {
        require(newImplementation_ != address(0), "Bridge: zero address");

        validateChangeAddressSignature(
            uint8(MethodId.AuthorizeUpgrade),
            address(this),
            newImplementation_,
            signature_
        );
    }
}
