// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

abstract contract Signers is OwnableUpgradeable {
    using ECDSA for bytes32;
    using EnumerableSet for EnumerableSet.AddressSet;

    EnumerableSet.AddressSet internal _signers;

    function __Signers_init(address[] calldata signers_) public onlyInitializing {
        __Ownable_init();

        addSigners(signers_);
    }

    function _checkSignatures(bytes32 signHash_, bytes[] calldata signatures_) internal view {
        for (uint256 i = 0; i < signatures_.length; i++) {
            address recovered_ = signHash_.toEthSignedMessageHash().recover(signatures_[i]);

            require(_signers.contains(recovered_), "Signers: invalid signer");
        }
    }

    function addSigners(address[] calldata signers_) public onlyOwner {
        for (uint256 i = 0; i < signers_.length; i++) {
            require(signers_[i] != address(0), "Signers: zero signer");

            _signers.add(signers_[i]);
        }
    }

    function removeSigners(address[] calldata signers_) public onlyOwner {
        for (uint256 i = 0; i < signers_.length; i++) {
            _signers.remove(signers_[i]);
        }
    }

    function getSigners() external view returns (address[] memory) {
        return _signers.values();
    }
}
