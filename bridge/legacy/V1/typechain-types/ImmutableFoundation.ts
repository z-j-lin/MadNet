/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { FunctionFragment, Result } from "@ethersproject/abi";
import { Listener, Provider } from "@ethersproject/providers";
import {
  BaseContract,
  BigNumber,
  BytesLike,
  CallOverrides,
  PopulatedTransaction,
  Signer,
  utils,
} from "ethers";
import { OnEvent, TypedEvent, TypedEventFilter, TypedListener } from "./common";

export interface ImmutableFoundationInterface extends utils.Interface {
  contractName: "ImmutableFoundation";
  functions: {
    "getMetamorphicContractAddress(bytes32,address)": FunctionFragment;
  };

  encodeFunctionData(
    functionFragment: "getMetamorphicContractAddress",
    values: [BytesLike, string]
  ): string;

  decodeFunctionResult(
    functionFragment: "getMetamorphicContractAddress",
    data: BytesLike
  ): Result;

  events: {};
}

export interface ImmutableFoundation extends BaseContract {
  contractName: "ImmutableFoundation";
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: ImmutableFoundationInterface;

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
    getMetamorphicContractAddress(
      _salt: BytesLike,
      _factory: string,
      overrides?: CallOverrides
    ): Promise<[string]>;
  };

  getMetamorphicContractAddress(
    _salt: BytesLike,
    _factory: string,
    overrides?: CallOverrides
  ): Promise<string>;

  callStatic: {
    getMetamorphicContractAddress(
      _salt: BytesLike,
      _factory: string,
      overrides?: CallOverrides
    ): Promise<string>;
  };

  filters: {};

  estimateGas: {
    getMetamorphicContractAddress(
      _salt: BytesLike,
      _factory: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    getMetamorphicContractAddress(
      _salt: BytesLike,
      _factory: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;
  };
}
