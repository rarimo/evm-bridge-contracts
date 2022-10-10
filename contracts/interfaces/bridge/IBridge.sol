// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "../handlers/IERC20Handler.sol";
import "../handlers/IERC721Handler.sol";
import "../handlers/IERC1155Handler.sol";
import "../handlers/INativeHandler.sol";

/**
 * @notice The Bridge contract
 *
 * The Bridge contract acts as a permissioned way of transfering assets (ERC20, ERC721, ERC1155, Native) between
 * 2 different blockchains.
 *
 * In order to correctly use the Bridge, one has to deploy both instances of the contract on the base chain and the
 * destination chain, as well as setup a trusted backend that will act as a `signer`.
 *
 * Each Bridge contract can either give or take the user assets when they want to transfer tokens. Both liquidity pool
 * and mint-and-burn way of transferring assets are supported.
 *
 * The bridge enables the transaction bundling feature as well.
 */
interface IBridge is IBundler, IERC20Handler, IERC721Handler, IERC1155Handler, INativeHandler {
    /**
     * @notice function for withdrawing erc20 tokens
     * @param tokenData_ the encoded token address and amount
     * @param bundle_ the encoded transaction bundle with encoded salt
     * @param receiver_ the address who will receive tokens
     * @param originHash_ the keccak256 hash of abi.encodePacked(origin chain name . origin tx hash . event nonce)
     * @param proof_ the abi encoded merkle path with the signature of a merkle root the signer signed
     * @param isWrapped_ the boolean flag, if true - tokens will minted, false - tokens will transferred
     */
    function withdrawERC20(
        bytes calldata tokenData_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        address receiver_,
        bytes calldata proof_,
        bool isWrapped_
    ) external;

    /**
     * @notice function for withdrawing erc721 tokens
     * @param tokenData_ the encoded token address, token id, amount (1), and token URI
     * @param bundle_ the encoded transaction bundle with encoded salt
     * @param originHash_ the keccak256 hash of abi.encodePacked(origin chain name . origin tx hash . event nonce)
     * @param receiver_ the address who will receive tokens
     * @param proof_ the abi encoded merkle path with the signature of a merkle root the signer signed
     * @param isWrapped_ the boolean flag, if true - tokens will minted, false - tokens will transferred
     */
    function withdrawERC721(
        bytes calldata tokenData_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        address receiver_,
        bytes calldata proof_,
        bool isWrapped_
    ) external;

    /**
     * @notice function for withdrawing erc1155 tokens
     * @param tokenData_ the encoded token address, token id, amount, and token URI
     * @param bundle_ the encoded transaction bundle with encoded salt
     * @param originHash_ the keccak256 hash of abi.encodePacked(origin chain name . origin tx hash . event nonce)
     * @param receiver_ the address who will receive tokens
     * @param proof_ the abi encoded merkle path with the signature of a merkle root the signer signed
     * @param isWrapped_ the boolean flag, if true - tokens will minted, false - tokens will transferred
     */
    function withdrawERC1155(
        bytes calldata tokenData_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        address receiver_,
        bytes calldata proof_,
        bool isWrapped_
    ) external;

    /**
     * @notice function for withdrawing native currency
     * @param tokenData_ the encoded native amount
     * @param bundle_ the encoded transaction bundle
     * @param originHash_ the keccak256 hash of abi.encodePacked(origin chain name . origin tx hash . event nonce)
     * @param receiver_ the address who will receive tokens
     * @param proof_ the abi encoded merkle path with the signature of a merkle root the signer signed
     */
    function withdrawNative(
        bytes calldata tokenData_,
        IBundler.Bundle calldata bundle_,
        bytes32 originHash_,
        address receiver_,
        bytes calldata proof_
    ) external;
}
