// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC1155/utils/ERC1155Holder.sol";

import "../interfaces/handlers/IERC1155Handler.sol";
import "../interfaces/tokens/IERC1155MintableBurnable.sol";

import "../libs/Encoder.sol";

import "../bundle/Bundler.sol";

abstract contract ERC1155Handler is IERC1155Handler, Bundler, ERC1155Holder {
    using Encoder for bytes32;

    function depositERC1155(
        DepositERC1155Parameters calldata params_
    ) external override onlyFacade {
        emit DepositedERC1155(
            params_.token,
            params_.tokenId,
            params_.amount,
            params_.bundle.salt.encode(),
            params_.bundle.bundle,
            params_.network,
            params_.receiver,
            params_.isWrapped
        );
    }

    function withdrawERC1155(WithdrawERC1155Parameters memory params_) public override onlyFacade {
        require(params_.token != address(0), "ERC1155Handler: zero token");
        require(params_.receiver != address(0), "ERC1155Handler: zero receiver");

        if (params_.isWrapped) {
            IERC1155MintableBurnable(params_.token).mintTo(
                params_.receiver,
                params_.tokenId,
                params_.amount,
                params_.tokenURI
            );
        } else {
            IERC1155MintableBurnable(params_.token).safeTransferFrom(
                address(this),
                params_.receiver,
                params_.tokenId,
                params_.amount,
                ""
            );
        }

        emit WithdrawnERC1155(
            params_.token,
            params_.tokenId,
            params_.tokenURI,
            params_.amount,
            params_.bundle.salt,
            params_.bundle.bundle,
            params_.originHash,
            params_.receiver,
            params_.proof,
            params_.isWrapped
        );
    }

    function withdrawERC1155Bundle(
        WithdrawERC1155Parameters memory params_
    ) external override onlyFacade {
        params_.receiver = determineProxyAddress(params_.bundle.salt);

        withdrawERC1155(params_);
        _bundleUp(params_.bundle);
    }

    uint256[50] private _gap;
}
