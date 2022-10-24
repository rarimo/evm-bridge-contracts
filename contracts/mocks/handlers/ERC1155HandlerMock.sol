// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../../handlers/ERC1155Handler.sol";

contract ERC1155HandlerMock is ERC1155Handler {
    function withdrawERC1155(bytes calldata tokenData_, address receiver_, bool isWrapped_)
        external
    {
        _withdrawERC1155(tokenData_, receiver_, isWrapped_);
    }
}
