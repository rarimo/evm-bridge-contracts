// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";

contract BundleReceiverMock is ERC721Holder {
    function doSomething(bool shouldRevert) external pure {
        require(!shouldRevert, "BundleReceiverMock: oops");
    }

    function withdrawNFT(IERC721 nft, uint256 tokenId) external {
        nft.safeTransferFrom(address(this), msg.sender, tokenId);
    }

    function withdrawETH(uint256 amount) external {
        payable(msg.sender).transfer(amount);
    }

    receive() external payable {}
}
