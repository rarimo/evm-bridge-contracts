// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../interfaces/handlers/IERC1155Handler.sol";

abstract contract ERC1155Handler is IERC1155Handler {
    function depositERC1155(
        address token_,
        uint256 tokenId_,
        uint256 amount_,
        string calldata receiver_,
        string calldata network_,
        bool isWrapped_
    ) external override {
        // TODO
    }
}
