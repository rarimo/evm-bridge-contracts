// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../../handlers/ERC1155Handler.sol";

contract ERC1155HandlerMock is ERC1155Handler {
    function withdrawERC1155(
        address token_,
        uint256 tokenId_,
        uint256 amount_,
        string calldata tokenURI_,
        address receiver_,
        IBundler.Bundle calldata bundle_,
        bool isWrapped_
    ) external {
        _withdrawERC1155(token_, tokenId_, amount_, tokenURI_, receiver_, bundle_, isWrapped_);
    }
}
