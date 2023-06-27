// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

interface IFeeManager {
    /**
     * @notice the enum that helps distinguish functions for calling within the signature
     * @param AddFeeToken the type corresponding to the addFeeToken function
     * @param RemoveFeeToken the type corresponding to the removeFeeToken function
     * @param UpdateFeeToken the type corresponding to the updateFeeToken function
     * @param AuthorizeUpgrade the type corresponding to the _authorizeUpgrade function
     */
    enum MethodId {
        AddFeeToken,
        RemoveFeeToken,
        UpdateFeeToken,
        WithdrawFeeToken,
        AuthorizeUpgrade
    }

    /**
     * @notice the event emitted from the addFeeToken function
     */
    event AddedFeeToken(address feeToken, uint256 feeAmount);

    /**
     * @notice the event emitted from the removeFeeToken function
     */
    event RemovedFeeToken(address feeToken, uint256 feeAmount);

    /**
     * @notice the event emitted from the updateFeeToken function
     */
    event UpdatedFeeToken(address feeToken, uint256 feeAmount);

    /**
     * @notice the event emitted from the withdrawFeeToken function
     */
    event WithdrawnFeeToken(address receiver, address feeToken, uint256 amount);

    /**
     * @notice the struct that represents add fee token parameters
     * @param feeTokens the list of fee tokens to be added
     * @param feeAmounts the list of corresponding fee token amounts
     * @param signature the add fee token signature
     */
    struct AddFeeTokenParameters {
        address[] feeTokens;
        uint256[] feeAmounts;
        bytes signature;
    }

    /**
     * @notice the struct that represents remove fee token parameters
     * @param feeTokens the list of fee tokens to be removed
     * @param feeAmounts the list of corresponding fee token amounts
     * @param signature the remove fee token signature
     */
    struct RemoveFeeTokenParameters {
        address[] feeTokens;
        uint256[] feeAmounts;
        bytes signature;
    }

    /**
     * @notice the struct that represents update fee token parameters
     * @param feeTokens the list of fee tokens to be updated
     * @param feeAmounts the list of corresponding fee token amounts
     * @param signature the update fee token signature
     */
    struct UpdateFeeTokenParameters {
        address[] feeTokens;
        uint256[] feeAmounts;
        bytes signature;
    }

    /**
     * @notice the struct that represents withdraw fee token parameters
     * @param receiver the address who will receive tokens
     * @param feeTokens the list of fee tokens to be withdrawn
     * @param feeAmounts the list of corresponding fee token amounts
     * @param signature the withdraw fee token signature
     */
    struct WithdrawFeeTokenParameters {
        address receiver;
        address[] feeTokens;
        uint256[] amounts;
        bytes signature;
    }

    /**
     * @notice the function that adds fee tokens
     * @param params_ the parameters for adding fee tokens
     */
    function addFeeToken(AddFeeTokenParameters calldata params_) external;

    /**
     * @notice the function that removes fee tokens
     * @param params_ the parameters for removing fee tokens
     */
    function removeFeeToken(RemoveFeeTokenParameters calldata params_) external;

    /**
     * @notice the function that updates fee tokens
     * @param params_ the parameters for updating fee tokens
     */
    function updateFeeToken(UpdateFeeTokenParameters calldata params_) external;

    /**
     * @notice the function that withdraws fee tokens
     * @param params_ the parameters for the fee tokens withdrawal
     */
    function withdrawFeeToken(WithdrawFeeTokenParameters calldata params_) external;
}
