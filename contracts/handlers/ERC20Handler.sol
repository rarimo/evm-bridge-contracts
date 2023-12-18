// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {IERC20MintableBurnable} from "../interfaces/tokens/IERC20MintableBurnable.sol";
import {IERC20Handler} from "../interfaces/handlers/IERC20Handler.sol";

import {Encoder} from "../libs/Encoder.sol";

import {Bundler} from "../bundle/Bundler.sol";

abstract contract ERC20Handler is IERC20Handler, Bundler {
    using SafeERC20 for IERC20MintableBurnable;
    using Encoder for bytes32;

    function depositERC20(DepositERC20Parameters calldata params_) external override onlyFacade {
        emit DepositedERC20(
            params_.token,
            params_.amount,
            params_.bundle.salt.encode(),
            params_.bundle.bundle,
            params_.network,
            params_.receiver,
            params_.isWrapped
        );
    }

    function withdrawERC20(WithdrawERC20Parameters memory params_) public override onlyFacade {
        require(params_.token != address(0), "ERC20Handler: zero token");
        require(params_.receiver != address(0), "ERC20Handler: zero receiver");

        if (params_.isWrapped) {
            IERC20MintableBurnable(params_.token).mintTo(params_.receiver, params_.amount);
        } else {
            IERC20MintableBurnable(params_.token).safeTransfer(params_.receiver, params_.amount);
        }

        emit WithdrawnERC20(
            params_.token,
            params_.amount,
            params_.bundle.salt,
            params_.bundle.bundle,
            params_.originHash,
            params_.receiver,
            params_.proof,
            params_.isWrapped
        );
    }

    function withdrawERC20Bundle(
        WithdrawERC20Parameters memory params_
    ) external override onlyFacade {
        params_.receiver = determineProxyAddress(params_.bundle.salt);

        withdrawERC20(params_);
        _bundleUp(params_.bundle);
    }

    uint256[50] private _gap;
}
