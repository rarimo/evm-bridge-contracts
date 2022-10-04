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
    using Encoder for *;

    function __Bridge_init(
        address signer_,
        address bundleImplementation_,
        string calldata chainName_
    ) external initializer {
        __Signers_init(signer_, chainName_);
        __Bundler_init(bundleImplementation_);
    }

    function withdrawERC20(
        bytes calldata tokenData_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        address receiver_,
        bytes calldata proof_,
        bool isWrapped_
    ) external override {
        bytes32 merkleLeaf_ = _getERC20TokenDataLeaf.encode(
            tokenData_,
            bundle_,
            originHash_,
            chainName,
            receiver_
        );

        _checkAndUpdateHashes(originHash_);
        _checkMerkleSignature(merkleLeaf_, proof_);

        _withdraw(
            _withdrawERC20,
            this.withdrawERC20Bundle,
            tokenData_,
            bundle_,
            receiver_,
            isWrapped_
        );
    }

    function withdrawERC721(
        bytes calldata tokenData_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        address receiver_,
        bytes calldata proof_,
        bool isWrapped_
    ) external override {
        bytes32 merkleLeaf_ = _getERC721TokenDataLeaf.encode(
            tokenData_,
            bundle_,
            originHash_,
            chainName,
            receiver_
        );

        _checkAndUpdateHashes(originHash_);
        _checkMerkleSignature(merkleLeaf_, proof_);

        _withdraw(
            _withdrawERC721,
            this.withdrawERC721Bundle,
            tokenData_,
            bundle_,
            receiver_,
            isWrapped_
        );
    }

    function withdrawERC1155(
        bytes calldata tokenData_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        address receiver_,
        bytes calldata proof_,
        bool isWrapped_
    ) external override {
        bytes32 merkleLeaf_ = _getERC1155TokenDataLeaf.encode(
            tokenData_,
            bundle_,
            originHash_,
            chainName,
            receiver_
        );

        _checkAndUpdateHashes(originHash_);
        _checkMerkleSignature(merkleLeaf_, proof_);

        _withdraw(
            _withdrawERC1155,
            this.withdrawERC1155Bundle,
            tokenData_,
            bundle_,
            receiver_,
            isWrapped_
        );
    }

    function withdrawNative(
        bytes calldata tokenData_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        address receiver_,
        bytes calldata proof_
    ) external override {
        bytes32 merkleLeaf_ = _getNativeTokenDataLeaf.encode(
            tokenData_,
            bundle_,
            originHash_,
            chainName,
            receiver_
        );

        _checkAndUpdateHashes(originHash_);
        _checkMerkleSignature(merkleLeaf_, proof_);

        _withdraw(
            _withdrawNative,
            this.withdrawNativeBundle,
            tokenData_,
            bundle_,
            receiver_,
            false
        );
    }

    function _withdraw(
        function(bytes calldata, address, bool) internal _withdrawFunc,
        function(bytes memory, IBundler.Bundle memory, bool) external bundleFunc,
        bytes calldata tokenData_,
        IBundler.Bundle calldata bundle_,
        address receiver_,
        bool isWrapped_
    ) internal {
        if (bundle_.bundle.length > 0) {
            try bundleFunc(tokenData_, bundle_, isWrapped_) {
                return;
            } catch {}
        }

        _withdrawFunc(tokenData_, receiver_, isWrapped_);
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
