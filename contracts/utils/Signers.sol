// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import "@openzeppelin/contracts/proxy/utils/Initializable.sol";

abstract contract Signers is Initializable {
    using ECDSA for bytes32;
    using MerkleProof for bytes32[];

    address public signer;
    uint256 public nonce;

    string public chainName;

    function __Signers_init(address signer_, string calldata chainName_) public onlyInitializing {
        signer = signer_;
        chainName = chainName_;
    }

    function _checkSignatureAndIncrementNonce(bytes32 signHash_, bytes memory signature_)
        internal
    {
        _checkSignature(signHash_, signature_);
        nonce++;
    }

    function _checkSignature(bytes32 signHash_, bytes memory signature_) internal view {
        address signer_ = signHash_.recover(signature_);

        require(signer == signer_, "Signers: invalid signature");
    }

    function _checkMerkleSignature(bytes32 merkleLeaf_, bytes calldata proof_) internal view {
        (bytes32[] memory merklePath_, bytes memory signature_) = abi.decode(
            proof_,
            (bytes32[], bytes)
        );

        bytes32 merkleRoot_ = merklePath_.processProof(merkleLeaf_);

        _checkSignature(merkleRoot_, signature_);
    }

    function _getAddressChangeHash(address newAddress_) internal view returns (bytes32) {
        require(newAddress_ != address(0), "Signers: new address is 0");

        return keccak256(abi.encodePacked(newAddress_, chainName, nonce, address(this)));
    }
}
