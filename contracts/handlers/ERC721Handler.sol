// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../interfaces/handlers/IERC721Handler.sol";

abstract contract ERC721Handler is IERC721Handler {
    function depositERC721(
        address token_,
        uint256 tokenId_,
        string calldata receiver_,
        string calldata network_,
        bool isWrapped_
    ) external override {
        // TODO
    }
}
