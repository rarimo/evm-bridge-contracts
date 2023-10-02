// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";

import "../interfaces/handlers/IERC721Handler.sol";
import "../interfaces/tokens/IERC721MintableBurnable.sol";
import "../interfaces/tokens/SBT/ISBT.sol";

import "../libs/Encoder.sol";

import "../bundle/Bundler.sol";

abstract contract ERC721Handler is IERC721Handler, Bundler, ERC721Holder {
    using Encoder for bytes32;

    function depositERC721(DepositERC721Parameters calldata params_) external override onlyFacade {
        emit DepositedERC721(
            params_.token,
            params_.tokenId,
            params_.bundle.salt.encode(),
            params_.bundle.bundle,
            params_.network,
            params_.receiver,
            params_.isWrapped
        );
    }

    function withdrawERC721(WithdrawERC721Parameters memory params_) public override onlyFacade {
        require(params_.token != address(0), "ERC721Handler: zero token");
        require(params_.receiver != address(0), "ERC721Handler: zero receiver");

        if (params_.isWrapped) {
            IERC721MintableBurnable(params_.token).mintTo(
                params_.receiver,
                params_.tokenId,
                params_.tokenURI
            );
        } else {
            IERC721MintableBurnable(params_.token).safeTransferFrom(
                address(this),
                params_.receiver,
                params_.tokenId
            );
        }

        emit WithdrawnERC721(
            params_.token,
            params_.tokenId,
            params_.tokenURI,
            params_.bundle.salt,
            params_.bundle.bundle,
            params_.originHash,
            params_.receiver,
            params_.proof,
            params_.isWrapped
        );
    }

    function withdrawERC721Bundle(
        WithdrawERC721Parameters memory params_
    ) external override onlyFacade {
        params_.receiver = determineProxyAddress(params_.bundle.salt);

        withdrawERC721(params_);
        _bundleUp(params_.bundle);
    }

    uint256[50] private _gap;
}
