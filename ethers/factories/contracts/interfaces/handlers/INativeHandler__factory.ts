/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  INativeHandler,
  INativeHandlerInterface,
} from "../../../../contracts/interfaces/handlers/INativeHandler";

const _abi = [
  {
    anonymous: false,
    inputs: [
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
    ],
    name: "DepositedNative",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
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
    ],
    name: "WithdrawnNative",
    type: "event",
  },
  {
    inputs: [
      {
        components: [
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
        ],
        internalType: "struct INativeHandler.DepositNativeParameters",
        name: "params_",
        type: "tuple",
      },
    ],
    name: "depositNative",
    outputs: [],
    stateMutability: "payable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "bytes32",
        name: "salt_",
        type: "bytes32",
      },
    ],
    name: "determineProxyAddress",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
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
        ],
        internalType: "struct INativeHandler.WithdrawNativeParameters",
        name: "params_",
        type: "tuple",
      },
    ],
    name: "withdrawNative",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
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
        ],
        internalType: "struct INativeHandler.WithdrawNativeParameters",
        name: "params_",
        type: "tuple",
      },
    ],
    name: "withdrawNativeBundle",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
] as const;

export class INativeHandler__factory {
  static readonly abi = _abi;
  static createInterface(): INativeHandlerInterface {
    return new utils.Interface(_abi) as INativeHandlerInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): INativeHandler {
    return new Contract(address, _abi, signerOrProvider) as INativeHandler;
  }
}
