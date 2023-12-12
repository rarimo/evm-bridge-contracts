// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ERC721Holder} from "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";
import {ERC1155Holder} from "@openzeppelin/contracts/token/ERC1155/utils/ERC1155Holder.sol";

import {IWrappedNative} from "../../interfaces/tokens/IWrappedNative.sol";

contract BundleExecutorImplementation is ERC721Holder, ERC1155Holder {
    address public immutable WRAPPED_NATIVE;

    constructor(address wrappedNative_) {
        WRAPPED_NATIVE = wrappedNative_;
    }

    function execute(bytes calldata bundle_) external payable {
        (address[] memory contracts_, uint256[] memory values_, bytes[] memory data_) = abi.decode(
            bundle_,
            (address[], uint256[], bytes[])
        );

        require(
            contracts_.length == values_.length && values_.length == data_.length,
            "BundleExecutorImplementation: lengths mismatch"
        );
        require(contracts_.length != 0, "BundleExecutorImplementation: zero bundle");

        for (uint256 i = 0; i < contracts_.length; i++) {
            (bool success_, ) = contracts_[i].call{value: values_[i]}(data_[i]);

            require(success_, "BundleExecutorImplementation: call reverted");
        }

        if (address(this).balance > 0) {
            IWrappedNative(WRAPPED_NATIVE).deposit{value: address(this).balance}();
        }
    }

    receive() external payable {}
}
