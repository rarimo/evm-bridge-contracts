// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {IBridge} from "../interfaces/bridge/IBridge.sol";
import {IBundler} from "../interfaces/bundle/IBundler.sol";

import {ERC20Handler} from "../handlers/ERC20Handler.sol";
import {ERC721Handler} from "../handlers/ERC721Handler.sol";
import {SBTHandler} from "../handlers/SBTHandler.sol";
import {ERC1155Handler} from "../handlers/ERC1155Handler.sol";
import {NativeHandler} from "../handlers/NativeHandler.sol";

import {Bundler} from "../bundle/Bundler.sol";

import {UUPSSignableUpgradeable} from "./proxy/UUPSSignableUpgradeable.sol";

import {Signers} from "../utils/Signers.sol";
import {Hashes} from "../utils/Hashes.sol";

import {Encoder} from "../libs/Encoder.sol";

contract Bridge is
    IBridge,
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

    function _authorizeUpgrade(address) internal pure override {
        revert("Bridge: this upgrade method is off");
    }

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
