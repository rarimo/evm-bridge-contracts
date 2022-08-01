// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155Receiver.sol";

interface IERC1155MintableBurnable is IERC1155, IERC1155Receiver {
    function mintTo(
        address receiver_,
        uint256 tokenId_,
        uint256 amount_
    ) external returns (bool);

    function burn(
        address payer_,
        uint256 tokenId_,
        uint256 amount_
    ) external;
}
