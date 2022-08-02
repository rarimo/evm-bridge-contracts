// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

interface IERC721MintableBurnable is IERC721 {
    function mintTo(address receiver_, uint256 tokenId_) external;

    function burn(uint256 tokenId_) external;
}