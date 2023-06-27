// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

import "../../interfaces/tokens/IWrappedNative.sol";

contract WrappedNativeMock is IWrappedNative, ERC20 {
    constructor() ERC20("WrappedNativeMock", "WNM") {}

    function deposit() external payable override {
        _mint(msg.sender, msg.value);
    }

    function withdraw(uint256 amount_) external override {
        _burn(msg.sender, amount_);

        payable(msg.sender).transfer(amount_);
    }
}
