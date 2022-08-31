// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../../utils/Signers.sol";

contract SignersMock is Signers {
    function __SignersMock_init(address signer_, string calldata chainName_) public initializer {
        __Signers_init(signer_, chainName_);
    }

    function checkSignature(bytes32 signHash_, bytes calldata signature_) external view {
        _checkSignature(signHash_, signature_);
    }

    function checkMerkleSignature(
        bytes32 merkleLeaf_,
        bytes32[] calldata merklePath_,
        bytes calldata signature_
    ) external view {
        _checkMerkleSignature(merkleLeaf_, merklePath_, signature_);
    }
}
