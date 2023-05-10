// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import "@openzeppelin/contracts/proxy/utils/Initializable.sol";

import "../interfaces/bridge/IBridge.sol";

abstract contract Signers is Initializable {
    using ECDSA for bytes32;
    using MerkleProof for bytes32[];

    uint256 public constant P = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F;

    address public signer;
    string public chainName;

    mapping(IBridge.MethodId => uint256) public nonces;

    function __Signers_init(address signer_, string calldata chainName_) public onlyInitializing {
        signer = signer_;
        chainName = chainName_;
    }

    function _checkSignatureAndIncrementNonce(
        IBridge.MethodId methodId_,
        bytes32 signHash_,
        bytes memory signature_
    ) internal {
        _checkSignature(signHash_, signature_);
        ++nonces[methodId_];
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

    function _getAddressChangeHash(
        IBridge.MethodId methodId_,
        address newAddress_
    ) internal view returns (bytes32) {
        require(newAddress_ != address(0), "Signers: new address is 0");

        return
            keccak256(
                abi.encodePacked(
                    methodId_,
                    newAddress_,
                    chainName,
                    nonces[methodId_],
                    address(this)
                )
            );
    }

    function _convertPubKeyToAddress(bytes memory pubKey_) internal pure returns (address) {
        require(pubKey_.length == 64, "Signers: wrong pubKey length");

        (uint256 x_, uint256 y_) = abi.decode(pubKey_, (uint256, uint256));

        // @dev y^2 = x^3 + 7, x != 0, y != 0 (mod P)
        require(x_ != 0 && y_ != 0 && x_ != P && y_ != P, "Signers: zero pubKey");
        require(
            mulmod(y_, y_, P) == addmod(mulmod(mulmod(x_, x_, P), x_, P), 7, P),
            "Signers: pubKey not on the curve"
        );

        return address(uint160(uint256(keccak256(pubKey_))));
    }
}
