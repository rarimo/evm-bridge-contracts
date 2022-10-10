// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";
import "@openzeppelin/contracts/token/ERC1155/utils/ERC1155Holder.sol";

contract BundleExecutorImplementation is ERC721Holder, ERC1155Holder {
    function execute(bytes calldata bundle_) external payable {
        (address[] memory contracts_, uint256[] memory values_, bytes[] memory data_) = abi.decode(
            bundle_,
            (address[], uint256[], bytes[])
        );

        for (uint256 i = 0; i < contracts_.length; i++) {
            (bool success_, ) = payable(contracts_[i]).call{value: values_[i]}(data_[i]);

            require(success_, "BundleExecutorImplementation: call reverted");
        }
    }

    receive() external payable {}
}
