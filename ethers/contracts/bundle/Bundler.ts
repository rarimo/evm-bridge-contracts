/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
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

export interface BundlerInterface extends utils.Interface {
  functions: {
    "__Bundler_init(address)": FunctionFragment;
    "bundleExecutorImplementation()": FunctionFragment;
    "determineProxyAddress(bytes32)": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "__Bundler_init"
      | "__Bundler_init(address)"
      | "bundleExecutorImplementation"
      | "bundleExecutorImplementation()"
      | "determineProxyAddress"
      | "determineProxyAddress(bytes32)"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "__Bundler_init",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "__Bundler_init(address)",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "bundleExecutorImplementation",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "bundleExecutorImplementation()",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "determineProxyAddress",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "determineProxyAddress(bytes32)",
    values: [PromiseOrValue<BytesLike>]
  ): string;

  decodeFunctionResult(
    functionFragment: "__Bundler_init",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "__Bundler_init(address)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "bundleExecutorImplementation",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "bundleExecutorImplementation()",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "determineProxyAddress",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "determineProxyAddress(bytes32)",
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

export interface Bundler extends BaseContract {
  contractName: "Bundler";

  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: BundlerInterface;

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
    __Bundler_init(
      bundleExecutorImplementation_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "__Bundler_init(address)"(
      bundleExecutorImplementation_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    bundleExecutorImplementation(overrides?: CallOverrides): Promise<[string]>;

    "bundleExecutorImplementation()"(
      overrides?: CallOverrides
    ): Promise<[string]>;

    determineProxyAddress(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    "determineProxyAddress(bytes32)"(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[string]>;
  };

  __Bundler_init(
    bundleExecutorImplementation_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "__Bundler_init(address)"(
    bundleExecutorImplementation_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  bundleExecutorImplementation(overrides?: CallOverrides): Promise<string>;

  "bundleExecutorImplementation()"(overrides?: CallOverrides): Promise<string>;

  determineProxyAddress(
    salt_: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<string>;

  "determineProxyAddress(bytes32)"(
    salt_: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<string>;

  callStatic: {
    __Bundler_init(
      bundleExecutorImplementation_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    "__Bundler_init(address)"(
      bundleExecutorImplementation_: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    bundleExecutorImplementation(overrides?: CallOverrides): Promise<string>;

    "bundleExecutorImplementation()"(
      overrides?: CallOverrides
    ): Promise<string>;

    determineProxyAddress(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;

    "determineProxyAddress(bytes32)"(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;
  };

  filters: {
    "Initialized(uint8)"(version?: null): InitializedEventFilter;
    Initialized(version?: null): InitializedEventFilter;
  };

  estimateGas: {
    __Bundler_init(
      bundleExecutorImplementation_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "__Bundler_init(address)"(
      bundleExecutorImplementation_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    bundleExecutorImplementation(overrides?: CallOverrides): Promise<BigNumber>;

    "bundleExecutorImplementation()"(
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    determineProxyAddress(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    "determineProxyAddress(bytes32)"(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    __Bundler_init(
      bundleExecutorImplementation_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "__Bundler_init(address)"(
      bundleExecutorImplementation_: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    bundleExecutorImplementation(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    "bundleExecutorImplementation()"(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    determineProxyAddress(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    "determineProxyAddress(bytes32)"(
      salt_: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;
  };
}
