/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  IERC1155Handler,
  IERC1155HandlerInterface,
} from "../../../../contracts/interfaces/handlers/IERC1155Handler";

const _abi = [
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "token",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "bytes32",
        name: "salt",
        type: "bytes32",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "bundle",
        type: "bytes",
      },
      {
        indexed: false,
        internalType: "string",
        name: "network",
        type: "string",
      },
      {
        indexed: false,
        internalType: "string",
        name: "receiver",
        type: "string",
      },
      {
        indexed: false,
        internalType: "bool",
        name: "isWrapped",
        type: "bool",
      },
    ],
    name: "DepositedERC1155",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "token",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "string",
        name: "tokenURI",
        type: "string",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "bytes32",
        name: "salt",
        type: "bytes32",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "bundle",
        type: "bytes",
      },
      {
        indexed: false,
        internalType: "bytes32",
        name: "originHash",
        type: "bytes32",
      },
      {
        indexed: false,
        internalType: "address",
        name: "receiver",
        type: "address",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "proof",
        type: "bytes",
      },
      {
        indexed: false,
        internalType: "bool",
        name: "isWrapped",
        type: "bool",
      },
    ],
    name: "WithdrawnERC1155",
    type: "event",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "address",
            name: "token",
            type: "address",
          },
          {
            internalType: "uint256",
            name: "tokenId",
            type: "uint256",
          },
          {
            internalType: "uint256",
            name: "amount",
            type: "uint256",
          },
          {
            components: [
              {
                internalType: "bytes32",
                name: "salt",
                type: "bytes32",
              },
              {
                internalType: "bytes",
                name: "bundle",
                type: "bytes",
              },
            ],
            internalType: "struct IBundler.Bundle",
            name: "bundle",
            type: "tuple",
          },
          {
            internalType: "string",
            name: "network",
            type: "string",
          },
          {
            internalType: "string",
            name: "receiver",
            type: "string",
          },
          {
            internalType: "bool",
            name: "isWrapped",
            type: "bool",
          },
        ],
        internalType: "struct IERC1155Handler.DepositERC1155Parameters",
        name: "params_",
        type: "tuple",
      },
    ],
    name: "depositERC1155",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "address",
            name: "token",
            type: "address",
          },
          {
            internalType: "uint256",
            name: "tokenId",
            type: "uint256",
          },
          {
            internalType: "string",
            name: "tokenURI",
            type: "string",
          },
          {
            internalType: "uint256",
            name: "amount",
            type: "uint256",
          },
          {
            components: [
              {
                internalType: "bytes32",
                name: "salt",
                type: "bytes32",
              },
              {
                internalType: "bytes",
                name: "bundle",
                type: "bytes",
              },
            ],
            internalType: "struct IBundler.Bundle",
            name: "bundle",
            type: "tuple",
          },
          {
            internalType: "bytes32",
            name: "originHash",
            type: "bytes32",
          },
          {
            internalType: "address",
            name: "receiver",
            type: "address",
          },
          {
            internalType: "bytes",
            name: "proof",
            type: "bytes",
          },
          {
            internalType: "bool",
            name: "isWrapped",
            type: "bool",
          },
        ],
        internalType: "struct IERC1155Handler.WithdrawERC1155Parameters",
        name: "params_",
        type: "tuple",
      },
    ],
    name: "withdrawERC1155",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "address",
            name: "token",
            type: "address",
          },
          {
            internalType: "uint256",
            name: "tokenId",
            type: "uint256",
          },
          {
            internalType: "string",
            name: "tokenURI",
            type: "string",
          },
          {
            internalType: "uint256",
            name: "amount",
            type: "uint256",
          },
          {
            components: [
              {
                internalType: "bytes32",
                name: "salt",
                type: "bytes32",
              },
              {
                internalType: "bytes",
                name: "bundle",
                type: "bytes",
              },
            ],
            internalType: "struct IBundler.Bundle",
            name: "bundle",
            type: "tuple",
          },
          {
            internalType: "bytes32",
            name: "originHash",
            type: "bytes32",
          },
          {
            internalType: "address",
            name: "receiver",
            type: "address",
          },
          {
            internalType: "bytes",
            name: "proof",
            type: "bytes",
          },
          {
            internalType: "bool",
            name: "isWrapped",
            type: "bool",
          },
        ],
        internalType: "struct IERC1155Handler.WithdrawERC1155Parameters",
        name: "params_",
        type: "tuple",
      },
    ],
    name: "withdrawERC1155Bundle",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
] as const;

export class IERC1155Handler__factory {
  static readonly abi = _abi;
  static createInterface(): IERC1155HandlerInterface {
    return new utils.Interface(_abi) as IERC1155HandlerInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): IERC1155Handler {
    return new Contract(address, _abi, signerOrProvider) as IERC1155Handler;
  }
}
