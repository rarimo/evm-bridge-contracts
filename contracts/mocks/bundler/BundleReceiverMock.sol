// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";

contract BundleReceiverMock is ERC721Holder {
    uint256 public count;

    function doSomething(bool shouldRevert) external {
        require(!shouldRevert, "BundleReceiverMock: oops");
        ++count;
    }

    function withdrawNFT(IERC721 nft, uint256 tokenId) external {
        nft.safeTransferFrom(address(this), msg.sender, tokenId);
    }

    function withdrawETH(uint256 amount) external {
        payable(msg.sender).transfer(amount);
    }

    receive() external payable {}
}
