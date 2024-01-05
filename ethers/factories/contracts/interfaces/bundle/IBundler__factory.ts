/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  IBundler,
  IBundlerInterface,
} from "../../../../contracts/interfaces/bundle/IBundler";

const _abi = [
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
] as const;

export class IBundler__factory {
  static readonly abi = _abi;
  static createInterface(): IBundlerInterface {
    return new utils.Interface(_abi) as IBundlerInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): IBundler {
    return new Contract(address, _abi, signerOrProvider) as IBundler;
  }
}