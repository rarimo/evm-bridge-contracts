// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import {ERC721URIStorage} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

import {IERC5192} from "../interfaces/tokens/SBT/IERC5192.sol";
import {ISBT} from "../interfaces/tokens/SBT/ISBT.sol";

contract SBTMintableBurnable is IERC5192, ISBT, Ownable, ERC721URIStorage {
    string public baseURI;

    mapping(address => uint256) internal _ownersToTokens;

    constructor(
        string memory name_,
        string memory symbol_,
        address owner_,
        string memory baseURI_
    ) ERC721(name_, symbol_) {
        baseURI = baseURI_;

        transferOwnership(owner_);
    }

    function attestTo(
        address receiver_,
        uint256 tokenId_,
        string calldata tokenURI_
    ) external override onlyOwner {
        _ownersToTokens[receiver_] = tokenId_;

        _mint(receiver_, tokenId_);
        _setTokenURI(tokenId_, tokenURI_);

        emit Locked(tokenId_);
    }

    function burn() external override {
        uint256 tokenId_ = tokenIdOf(msg.sender);

        delete _ownersToTokens[msg.sender];

        _burn(tokenId_);
    }

    function tokenIdOf(address from_) public view override returns (uint256 tokenId_) {
        tokenId_ = _ownersToTokens[from_];

        require(tokenId_ != 0, "SBTMintableBurnable: user doesn't have a token");
    }

    function locked(uint256 tokenId_) external view override returns (bool) {
        require(_exists(tokenId_), "SBTMintableBurnable: token is assigned to zero address");

        return true;
    }

    function supportsInterface(
        bytes4 interfaceId_
    ) public view override(ERC721, IERC165) returns (bool) {
        return
            interfaceId_ == type(IERC5192).interfaceId ||
            interfaceId_ == type(ISBT).interfaceId ||
            super.supportsInterface(interfaceId_);
    }

    function _beforeTokenTransfer(address from_, address to_, uint256 tokenId_) internal override {
        super._beforeTokenTransfer(from_, to_, tokenId_);

        if (from_ == address(0)) {
            require(balanceOf(to_) == 0, "SBTMintableBurnable: receiver already has a token");
        } else {
            require(to_ == address(0), "SBTMintableBurnable: not transferable");
        }
    }

    function _baseURI() internal view override returns (string memory) {
        return baseURI;
    }
}
