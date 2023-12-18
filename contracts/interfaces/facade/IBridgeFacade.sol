// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {IFeeManager} from "./IFeeManager.sol";
import {IBridge} from "../bridge/IBridge.sol";

/**
 * @notice The Bridge Facade contract.
 *
 * The Bridge Facade contract serves as a facade interface for interacting with the underlying Bridge contract.
 * It delegates all deposit and withdrawal requests to the Bridge contract and collects fees.
 * It also provides an opportunity to manage the list of fee tokens and withdraw commissions by utilizing the signature-based methods.
 */
interface IBridgeFacade is IFeeManager {
    /**
     * @notice the struct that represents fee parameters for the erc20 deposit
     * @param token the address of the fee token
     */
    struct DepositFeeERC20Parameters {
        address feeToken;
    }

    /**
     * @notice the struct that represents fee parameters for the erc721 deposit
     * @param feeToken the address of the fee token
     */
    struct DepositFeeERC721Parameters {
        address feeToken;
    }

    /**
     * @notice the struct that represents fee parameters for the sbt deposit
     * @param feeToken the address of the fee token
     */
    struct DepositFeeSBTParameters {
        address feeToken;
    }

    /**
     * @notice the struct that represents fee parameters for the erc1155 deposit
     * @param feeToken the address of the fee token
     */
    struct DepositFeeERC1155Parameters {
        address feeToken;
    }

    /**
     * @notice the struct that represents fee parameters for the native deposit
     * @param feeToken the address of the fee token
     */
    struct DepositFeeNativeParameters {
        address feeToken;
    }

    /**
     * @notice the function that deposits erc20 tokens with a fee
     * @param feeParams_ the fee parameters for the erc20 deposit
     * @param depositParams_ the parameters for the erc20 deposit
     */
    function depositERC20(
        DepositFeeERC20Parameters calldata feeParams_,
        IBridge.DepositERC20Parameters calldata depositParams_
    ) external payable;

    /**
     * @notice the function that deposits erc721 tokens with a fee
     * @param feeParams_ the fee parameters for the erc721 deposit
     * @param depositParams_ the parameters for the erc721 deposit
     */
    function depositERC721(
        DepositFeeERC721Parameters calldata feeParams_,
        IBridge.DepositERC721Parameters calldata depositParams_
    ) external payable;

    /**
     * @notice the function that deposits sbt tokens with a fee
     * @param feeParams_ the fee parameters for the sbt deposit
     * @param depositParams_ the parameters for the sbt deposit
     */
    function depositSBT(
        DepositFeeSBTParameters calldata feeParams_,
        IBridge.DepositSBTParameters calldata depositParams_
    ) external payable;

    /**
     * @notice the function that deposits erc1155 tokens with a fee
     * @param feeParams_ the fee parameters for the erc1155 deposit
     * @param depositParams_ the parameters for the erc1155 deposit
     */
    function depositERC1155(
        DepositFeeERC1155Parameters calldata feeParams_,
        IBridge.DepositERC1155Parameters calldata depositParams_
    ) external payable;

    /**
     * @notice the function that deposits native tokens with a fee
     * @param feeParams_ the fee parameters for the native deposit
     * @param depositParams_ the parameters for the native deposit
     */
    function depositNative(
        DepositFeeNativeParameters calldata feeParams_,
        IBridge.DepositNativeParameters calldata depositParams_
    ) external payable;

    /**
     * @notice the function to withdraw erc20 tokens
     * @param params_ the parameters for the erc20 withdrawal
     */
    function withdrawERC20(IBridge.WithdrawERC20Parameters calldata params_) external;

    /**
     * @notice the function to withdraw erc721 tokens
     * @param params_ the parameters for the erc721 withdrawal
     */
    function withdrawERC721(IBridge.WithdrawERC721Parameters calldata params_) external;

    /**
     * @notice the function to withdraw sbt tokens
     * @param params_ the parameters for the sbt withdrawal
     */
    function withdrawSBT(IBridge.WithdrawSBTParameters calldata params_) external;

    /**
     * @notice the function to withdraw erc1155 tokens
     * @param params_ the parameters for the erc1155 withdrawal
     */
    function withdrawERC1155(IBridge.WithdrawERC1155Parameters calldata params_) external;

    /**
     * @notice the function to withdraw native tokens
     * @param params_ the parameters for the native withdrawal
     */
    function withdrawNative(IBridge.WithdrawNativeParameters calldata params_) external;
}
