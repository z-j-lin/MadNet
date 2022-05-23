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

export interface TXInPreImageParserLibraryErrorCodesInterface
  extends utils.Interface {
  contractName: "TXInPreImageParserLibraryErrorCodes";
  functions: {
    "TXINPREIMAGEPARSERLIB_CHAINID_ZERO()": FunctionFragment;
    "TXINPREIMAGEPARSERLIB_DATA_OFFSET_OVERFLOW()": FunctionFragment;
    "TXINPREIMAGEPARSERLIB_INSUFFICIENT_BYTES()": FunctionFragment;
  };

  encodeFunctionData(
    functionFragment: "TXINPREIMAGEPARSERLIB_CHAINID_ZERO",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "TXINPREIMAGEPARSERLIB_DATA_OFFSET_OVERFLOW",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "TXINPREIMAGEPARSERLIB_INSUFFICIENT_BYTES",
    values?: undefined
  ): string;

  decodeFunctionResult(
    functionFragment: "TXINPREIMAGEPARSERLIB_CHAINID_ZERO",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "TXINPREIMAGEPARSERLIB_DATA_OFFSET_OVERFLOW",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "TXINPREIMAGEPARSERLIB_INSUFFICIENT_BYTES",
    data: BytesLike
  ): Result;

  events: {};
}

export interface TXInPreImageParserLibraryErrorCodes extends BaseContract {
  contractName: "TXInPreImageParserLibraryErrorCodes";
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: TXInPreImageParserLibraryErrorCodesInterface;

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
    TXINPREIMAGEPARSERLIB_CHAINID_ZERO(
      overrides?: CallOverrides
    ): Promise<[string]>;

    TXINPREIMAGEPARSERLIB_DATA_OFFSET_OVERFLOW(
      overrides?: CallOverrides
    ): Promise<[string]>;

    TXINPREIMAGEPARSERLIB_INSUFFICIENT_BYTES(
      overrides?: CallOverrides
    ): Promise<[string]>;
  };

  TXINPREIMAGEPARSERLIB_CHAINID_ZERO(
    overrides?: CallOverrides
  ): Promise<string>;

  TXINPREIMAGEPARSERLIB_DATA_OFFSET_OVERFLOW(
    overrides?: CallOverrides
  ): Promise<string>;

  TXINPREIMAGEPARSERLIB_INSUFFICIENT_BYTES(
    overrides?: CallOverrides
  ): Promise<string>;

  callStatic: {
    TXINPREIMAGEPARSERLIB_CHAINID_ZERO(
      overrides?: CallOverrides
    ): Promise<string>;

    TXINPREIMAGEPARSERLIB_DATA_OFFSET_OVERFLOW(
      overrides?: CallOverrides
    ): Promise<string>;

    TXINPREIMAGEPARSERLIB_INSUFFICIENT_BYTES(
      overrides?: CallOverrides
    ): Promise<string>;
  };

  filters: {};

  estimateGas: {
    TXINPREIMAGEPARSERLIB_CHAINID_ZERO(
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    TXINPREIMAGEPARSERLIB_DATA_OFFSET_OVERFLOW(
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    TXINPREIMAGEPARSERLIB_INSUFFICIENT_BYTES(
      overrides?: CallOverrides
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    TXINPREIMAGEPARSERLIB_CHAINID_ZERO(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    TXINPREIMAGEPARSERLIB_DATA_OFFSET_OVERFLOW(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    TXINPREIMAGEPARSERLIB_INSUFFICIENT_BYTES(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;
  };
}
