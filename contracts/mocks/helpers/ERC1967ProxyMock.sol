// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract ERC1967ProxyMock is ERC1967Proxy {
    constructor(address logic_, bytes memory data_) ERC1967Proxy(logic_, data_) {}

    function implementation() external view returns (address) {
        return _implementation();
    }
}
