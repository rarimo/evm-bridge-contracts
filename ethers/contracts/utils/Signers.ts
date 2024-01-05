/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BigNumberish,
  BytesLike,
  CallOverrides,
  ContractTransaction,
  Overrides,
  PopulatedTransaction,
  Signer,
  utils,
} from "ethers";
import type {
  FunctionFragment,
  Result,
  EventFragment,
} from "@ethersproject/abi";
import type { Listener, Provider } from "@ethersproject/providers";
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
  PromiseOrValue,
} from "../../common";

export interface SignersInterface extends utils.Interface {
  functions: {
    "P()": FunctionFragment;
    "__Signers_init(address,address,string)": FunctionFragment;
    "chainName()": FunctionFragment;
    "checkSignatureAndIncrementNonce(uint8,address,bytes32,bytes)": FunctionFragment;
    "facade()": FunctionFragment;
    "getSigComponents(uint8,address)": FunctionFragment;
    "nonces(address,uint8)": FunctionFragment;
    "signer()": FunctionFragment;
    "validateChangeAddressSignature(uint8,address,address,bytes)": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "P"
      | "P()"
      | "__Signers_init"
      | "__Signers_init(address,address,string)"
      | "chainName"
      | "chainName()"
      | "checkSignatureAndIncrementNonce"
      | "checkSignatureAndIncrementNonce(uint8,address,bytes32,bytes)"
      | "facade"
      | "facade()"
      | "getSigComponents"
      | "getSigComponents(uint8,address)"
      | "nonces"
      | "nonces(address,uint8)"
      | "signer"
      | "signer()"
      | "validateChangeAddressSignature"
      | "validateChangeAddressSignature(uint8,address,address,bytes)"
  ): FunctionFragment;

  encodeFunctionData(functionFragment: "P", values?: undefined): string;
  encodeFunctionData(functionFragment: "P()", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "__Signers_init",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<string>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "__Signers_init(address,address,string)",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<string>
    ]
  ): string;
  encodeFunctionData(functionFragment: "chainName", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "chainName()",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "checkSignatureAndIncrementNonce",
    values: [
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<string>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "checkSignatureAndIncrementNonce(uint8,address,bytes32,bytes)",
    values: [
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<string>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(functionFragment: "facade", values?: undefined): string;
  encodeFunctionData(functionFragment: "facade()", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "getSigComponents",
    values: [PromiseOrValue<BigNumberish>, PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "getSigComponents(uint8,address)",
    values: [PromiseOrValue<BigNumberish>, PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "nonces",
    values: [PromiseOrValue<string>, PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "nonces(address,uint8)",
    values: [PromiseOrValue<string>, PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(functionFragment: "signer", values?: undefined): string;
  encodeFunctionData(functionFragment: "signer()", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "validateChangeAddressSignature",
    values: [
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "validateChangeAddressSignature(uint8,address,address,bytes)",
    values: [
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BytesLike>
    ]
  ): string;

  decodeFunctionResult(functionFragment: "P", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "P()", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "__Signers_init",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "__Signers_init(address,address,string)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "chainName", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "chainName()",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "checkSignatureAndIncrementNonce",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "checkSignatureAndIncrementNonce(uint8,address,bytes32,bytes)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "facade", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "facade()", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "getSigComponents",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getSigComponents(uint8,address)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "nonces", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "nonces(address,uint8)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "signer", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "signer()", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "validateChangeAddressSignature",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "validateChangeAddressSignature(uint8,address,address,bytes)",
    data: BytesLike
  ): Result;

  events: {
    "Initialized(uint8)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "Initialized"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Initialized(uint8)"): EventFragment;
}

export interface InitializedEventObject {
  version: number;
}
export type InitializedEvent = TypedEvent<[number], InitializedEventObject>;

export type InitializedEventFilter = TypedEventFilter<InitializedEvent>;

export interface Signers extends BaseContract {
  contractName: "Signers";

  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: SignersInterface;

  queryFilter<TEvent extends TypedEvent>(
    event: TypedEventFilter<TEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TEvent>>;

  listeners<TEvent extends TypedEvent>(
    eventFilter?: TypedEventFilter<TEvent>
  ): Array<TypedListener<TEvent>>;
  listeners(eventName?: string): Array<Listener>;
  removeAllListeners<TEvent extends TypedEvent>(
    eventFilter: TypedEventFilter<TEvent>
  ): this;
  removeAllListeners(eventName?: string): this;
  off: OnEvent<this>;
  on: OnEvent<this>;
  once: OnEvent<this>;
  removeListener: OnEvent<this>;

  functions: {
    P(overrides?: CallOverrides): Promise<[BigNumber]>;

    "P()"(overrides?: CallOverrides): Promise<[BigNumber]>;

    __Signers_init(
      signer_: PromiseOrValue<string>,
      facade_: PromiseOrValue<string>,
      chainName_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "__Signers_init(address,address,string)"(
      signer_: PromiseOrValue<string>,
      facade_: PromiseOrValue<string>,
      chainName_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    chainName(overrides?: CallOverrides): Promise<[string]>;

    "chainName()"(overrides?: CallOverrides): Promise<[string]>;

    checkSignatureAndIncrementNonce(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      signHash_: PromiseOrValue<BytesLike>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "checkSignatureAndIncrementNonce(uint8,address,bytes32,bytes)"(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      signHash_: PromiseOrValue<BytesLike>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    facade(overrides?: CallOverrides): Promise<[string]>;

    "facade()"(overrides?: CallOverrides): Promise<[string]>;

    getSigComponents(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[string, BigNumber] & { chainName_: string; nonce_: BigNumber }>;

    "getSigComponents(uint8,address)"(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[string, BigNumber] & { chainName_: string; nonce_: BigNumber }>;

    nonces(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    "nonces(address,uint8)"(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    signer(overrides?: CallOverrides): Promise<[string]>;

    "signer()"(overrides?: CallOverrides): Promise<[string]>;

    validateChangeAddressSignature(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      newAddress_: PromiseOrValue<string>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "validateChangeAddressSignature(uint8,address,address,bytes)"(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      newAddress_: PromiseOrValue<string>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  P(overrides?: CallOverrides): Promise<BigNumber>;

  "P()"(overrides?: CallOverrides): Promise<BigNumber>;

  __Signers_init(
    signer_: PromiseOrValue<string>,
    facade_: PromiseOrValue<string>,
    chainName_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "__Signers_init(address,address,string)"(
    signer_: PromiseOrValue<string>,
    facade_: PromiseOrValue<string>,
    chainName_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  chainName(overrides?: CallOverrides): Promise<string>;

  "chainName()"(overrides?: CallOverrides): Promise<string>;

  checkSignatureAndIncrementNonce(
    methodId_: PromiseOrValue<BigNumberish>,
    contractAddress_: PromiseOrValue<string>,
    signHash_: PromiseOrValue<BytesLike>,
    signature_: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "checkSignatureAndIncrementNonce(uint8,address,bytes32,bytes)"(
    methodId_: PromiseOrValue<BigNumberish>,
    contractAddress_: PromiseOrValue<string>,
    signHash_: PromiseOrValue<BytesLike>,
    signature_: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  facade(overrides?: CallOverrides): Promise<string>;

  "facade()"(overrides?: CallOverrides): Promise<string>;

  getSigComponents(
    methodId_: PromiseOrValue<BigNumberish>,
    contractAddress_: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<[string, BigNumber] & { chainName_: string; nonce_: BigNumber }>;

  "getSigComponents(uint8,address)"(
    methodId_: PromiseOrValue<BigNumberish>,
    contractAddress_: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<[string, BigNumber] & { chainName_: string; nonce_: BigNumber }>;

  nonces(
    arg0: PromiseOrValue<string>,
    arg1: PromiseOrValue<BigNumberish>,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  "nonces(address,uint8)"(
    arg0: PromiseOrValue<string>,
    arg1: PromiseOrValue<BigNumberish>,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  validateChangeAddressSignature(
    methodId_: PromiseOrValue<BigNumberish>,
    contractAddress_: PromiseOrValue<string>,
    newAddress_: PromiseOrValue<string>,
    signature_: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "validateChangeAddressSignature(uint8,address,address,bytes)"(
    methodId_: PromiseOrValue<BigNumberish>,
    contractAddress_: PromiseOrValue<string>,
    newAddress_: PromiseOrValue<string>,
    signature_: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    P(overrides?: CallOverrides): Promise<BigNumber>;

    "P()"(overrides?: CallOverrides): Promise<BigNumber>;

    __Signers_init(
      signer_: PromiseOrValue<string>,
      facade_: PromiseOrValue<string>,
      chainName_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    "__Signers_init(address,address,string)"(
      signer_: PromiseOrValue<string>,
      facade_: PromiseOrValue<string>,
      chainName_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    chainName(overrides?: CallOverrides): Promise<string>;

    "chainName()"(overrides?: CallOverrides): Promise<string>;

    checkSignatureAndIncrementNonce(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      signHash_: PromiseOrValue<BytesLike>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    "checkSignatureAndIncrementNonce(uint8,address,bytes32,bytes)"(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      signHash_: PromiseOrValue<BytesLike>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    facade(overrides?: CallOverrides): Promise<string>;

    "facade()"(overrides?: CallOverrides): Promise<string>;

    getSigComponents(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[string, BigNumber] & { chainName_: string; nonce_: BigNumber }>;

    "getSigComponents(uint8,address)"(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[string, BigNumber] & { chainName_: string; nonce_: BigNumber }>;

    nonces(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    "nonces(address,uint8)"(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    signer(overrides?: CallOverrides): Promise<string>;

    "signer()"(overrides?: CallOverrides): Promise<string>;

    validateChangeAddressSignature(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      newAddress_: PromiseOrValue<string>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    "validateChangeAddressSignature(uint8,address,address,bytes)"(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      newAddress_: PromiseOrValue<string>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {
    "Initialized(uint8)"(version?: null): InitializedEventFilter;
    Initialized(version?: null): InitializedEventFilter;
  };

  estimateGas: {
    P(overrides?: CallOverrides): Promise<BigNumber>;

    "P()"(overrides?: CallOverrides): Promise<BigNumber>;

    __Signers_init(
      signer_: PromiseOrValue<string>,
      facade_: PromiseOrValue<string>,
      chainName_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "__Signers_init(address,address,string)"(
      signer_: PromiseOrValue<string>,
      facade_: PromiseOrValue<string>,
      chainName_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    chainName(overrides?: CallOverrides): Promise<BigNumber>;

    "chainName()"(overrides?: CallOverrides): Promise<BigNumber>;

    checkSignatureAndIncrementNonce(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      signHash_: PromiseOrValue<BytesLike>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "checkSignatureAndIncrementNonce(uint8,address,bytes32,bytes)"(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      signHash_: PromiseOrValue<BytesLike>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    facade(overrides?: CallOverrides): Promise<BigNumber>;

    "facade()"(overrides?: CallOverrides): Promise<BigNumber>;

    getSigComponents(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    "getSigComponents(uint8,address)"(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    nonces(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    "nonces(address,uint8)"(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    signer(overrides?: CallOverrides): Promise<BigNumber>;

    "signer()"(overrides?: CallOverrides): Promise<BigNumber>;

    validateChangeAddressSignature(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      newAddress_: PromiseOrValue<string>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "validateChangeAddressSignature(uint8,address,address,bytes)"(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      newAddress_: PromiseOrValue<string>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    P(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    "P()"(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    __Signers_init(
      signer_: PromiseOrValue<string>,
      facade_: PromiseOrValue<string>,
      chainName_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "__Signers_init(address,address,string)"(
      signer_: PromiseOrValue<string>,
      facade_: PromiseOrValue<string>,
      chainName_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    chainName(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    "chainName()"(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    checkSignatureAndIncrementNonce(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      signHash_: PromiseOrValue<BytesLike>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "checkSignatureAndIncrementNonce(uint8,address,bytes32,bytes)"(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      signHash_: PromiseOrValue<BytesLike>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    facade(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    "facade()"(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    getSigComponents(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    "getSigComponents(uint8,address)"(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    nonces(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    "nonces(address,uint8)"(
      arg0: PromiseOrValue<string>,
      arg1: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    signer(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    "signer()"(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    validateChangeAddressSignature(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      newAddress_: PromiseOrValue<string>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "validateChangeAddressSignature(uint8,address,address,bytes)"(
      methodId_: PromiseOrValue<BigNumberish>,
      contractAddress_: PromiseOrValue<string>,
      newAddress_: PromiseOrValue<string>,
      signature_: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
