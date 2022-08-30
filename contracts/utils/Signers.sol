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

    function __Signers_init(address signer_) public onlyInitializing {
        signer = signer_;
    }

    function _checkSignature(bytes32 signHash_, bytes memory signature_) internal view {
        address signer_ = signHash_.toEthSignedMessageHash().recover(signature_);

        require(signer == signer_, "Signers: invalid signature");
    }

    function _checkMerkleSignature(
        bytes32 merkleLeaf_,
        bytes32[] calldata merklePath_,
        bytes calldata signature_
    ) internal view {
        bytes32 merkleRoot_ = merklePath_.processProof(merkleLeaf_);

        _checkSignature(merkleRoot_, signature_);
    }

    function changeSigner(address newSigner_, bytes memory signature_) external {
        _checkSignature(keccak256(abi.encodePacked(newSigner_, nonce++)), signature_);

        signer = newSigner_;
    }
}
