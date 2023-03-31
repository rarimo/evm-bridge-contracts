// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

interface IERC5192 {
    event Locked(uint256 tokenId);

    function locked(uint256 tokenId) external view returns (bool);
}
