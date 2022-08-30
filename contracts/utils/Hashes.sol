// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

abstract contract Hashes {
    mapping(bytes32 => bool) public usedHashes; // keccak256(origin chain name . origin tx hash . event nonce) => is used

    function _checkAndUpdateHashes(bytes32 originHash_) internal {
        require(!usedHashes[originHash_], "Hashes: the hash nonce is used");

        usedHashes[originHash_] = true;
    }
}
