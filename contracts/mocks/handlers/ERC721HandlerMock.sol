// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../../handlers/ERC721Handler.sol";

contract ERC721HandlerMock is ERC721Handler {
    function withdrawERC721(
        bytes calldata tokenData_,
        address receiver_,
        bool isWrapped_
    ) external {
        _withdrawERC721(tokenData_, receiver_, isWrapped_);
    }
}
