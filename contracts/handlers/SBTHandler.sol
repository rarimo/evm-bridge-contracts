// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";

import "../interfaces/handlers/ISBTHandler.sol";
import "../interfaces/tokens/IERC721MintableBurnable.sol";
import "../interfaces/tokens/SBT/ISBT.sol";

import "../libs/Encoder.sol";

import "../bundle/Bundler.sol";

abstract contract SBTHandler is ISBTHandler, Bundler, ERC721Holder {
    using Encoder for bytes32;

    function depositSBT(DepositSBTParameters calldata params_) external override onlyFacade {
        emit DepositedSBT(
            params_.token,
            params_.tokenId,
            params_.bundle.salt.encode(),
            params_.bundle.bundle,
            params_.network,
            params_.receiver
        );
    }

    function withdrawSBT(WithdrawSBTParameters memory params_) public override onlyFacade {
        require(params_.token != address(0), "SBTHandler: zero token");
        require(params_.receiver != address(0), "SBTHandler: zero receiver");

        ISBT(params_.token).attestTo(params_.receiver, params_.tokenId, params_.tokenURI);
    }

    function withdrawSBTBundle(WithdrawSBTParameters memory params_) external override onlyFacade {
        params_.receiver = determineProxyAddress(params_.bundle.salt);

        withdrawSBT(params_);
        _bundleUp(params_.bundle);
    }

    uint256[50] private _gap;
}
