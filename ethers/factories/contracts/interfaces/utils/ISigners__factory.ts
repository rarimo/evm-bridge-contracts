/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  ISigners,
  ISignersInterface,
} from "../../../../contracts/interfaces/utils/ISigners";

const _abi = [
  {
    inputs: [
      {
        internalType: "uint8",
        name: "methodId_",
        type: "uint8",
      },
      {
        internalType: "address",
        name: "contractAddress_",
        type: "address",
      },
      {
        internalType: "bytes32",
        name: "signHash_",
        type: "bytes32",
      },
      {
        internalType: "bytes",
        name: "signature_",
        type: "bytes",
      },
    ],
    name: "checkSignatureAndIncrementNonce",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint8",
        name: "methodId_",
        type: "uint8",
      },
      {
        internalType: "address",
        name: "contractAddress_",
        type: "address",
      },
    ],
    name: "getSigComponents",
    outputs: [
      {
        internalType: "string",
        name: "chainName_",
        type: "string",
      },
      {
        internalType: "uint256",
        name: "nonce_",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint8",
        name: "methodId_",
        type: "uint8",
      },
      {
        internalType: "address",
        name: "contractAddress_",
        type: "address",
      },
      {
        internalType: "address",
        name: "newAddress_",
        type: "address",
      },
      {
        internalType: "bytes",
        name: "signature_",
        type: "bytes",
      },
    ],
    name: "validateChangeAddressSignature",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
] as const;

export class ISigners__factory {
  static readonly abi = _abi;
  static createInterface(): ISignersInterface {
    return new utils.Interface(_abi) as ISignersInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): ISigners {
    return new Contract(address, _abi, signerOrProvider) as ISigners;
  }
}