// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

interface ISBT is IERC721 {
    function attestTo(address receiver_, uint256 tokenId_, string calldata tokenURI_) external;

    function burn() external;

    function tokenIdOf(address from_) external view returns (uint256 tokenId_);
}
