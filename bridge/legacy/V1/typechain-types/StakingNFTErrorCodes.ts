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

export interface StakingNFTErrorCodesInterface extends utils.Interface {
  contractName: "StakingNFTErrorCodes";
  functions: {
    "STAKENFT_BALANCE_LESS_THAN_RESERVE()": FunctionFragment;
    "STAKENFT_CALLER_NOT_TOKEN_OWNER()": FunctionFragment;
    "STAKENFT_FREE_AFTER_TIME_NOT_REACHED()": FunctionFragment;
    "STAKENFT_INVALID_TOKEN_ID()": FunctionFragment;
    "STAKENFT_LOCK_DURATION_GREATER_THAN_GOVERNANCE_LOCK()": FunctionFragment;
    "STAKENFT_LOCK_DURATION_GREATER_THAN_MINT_LOCK()": FunctionFragment;
    "STAKENFT_LOCK_DURATION_WITHDRAW_TIME_NOT_REACHED()": FunctionFragment;
    "STAKENFT_MINT_AMOUNT_EXCEEDS_MAX_SUPPLY()": FunctionFragment;
    "STAKENFT_SLUSH_TOO_LARGE()": FunctionFragment;
  };

  encodeFunctionData(
    functionFragment: "STAKENFT_BALANCE_LESS_THAN_RESERVE",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "STAKENFT_CALLER_NOT_TOKEN_OWNER",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "STAKENFT_FREE_AFTER_TIME_NOT_REACHED",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "STAKENFT_INVALID_TOKEN_ID",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "STAKENFT_LOCK_DURATION_GREATER_THAN_GOVERNANCE_LOCK",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "STAKENFT_LOCK_DURATION_GREATER_THAN_MINT_LOCK",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "STAKENFT_LOCK_DURATION_WITHDRAW_TIME_NOT_REACHED",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "STAKENFT_MINT_AMOUNT_EXCEEDS_MAX_SUPPLY",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "STAKENFT_SLUSH_TOO_LARGE",
    values?: undefined
  ): string;

  decodeFunctionResult(
    functionFragment: "STAKENFT_BALANCE_LESS_THAN_RESERVE",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "STAKENFT_CALLER_NOT_TOKEN_OWNER",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "STAKENFT_FREE_AFTER_TIME_NOT_REACHED",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "STAKENFT_INVALID_TOKEN_ID",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "STAKENFT_LOCK_DURATION_GREATER_THAN_GOVERNANCE_LOCK",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "STAKENFT_LOCK_DURATION_GREATER_THAN_MINT_LOCK",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "STAKENFT_LOCK_DURATION_WITHDRAW_TIME_NOT_REACHED",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "STAKENFT_MINT_AMOUNT_EXCEEDS_MAX_SUPPLY",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "STAKENFT_SLUSH_TOO_LARGE",
    data: BytesLike
  ): Result;

  events: {};
}

export interface StakingNFTErrorCodes extends BaseContract {
  contractName: "StakingNFTErrorCodes";
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: StakingNFTErrorCodesInterface;

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
    STAKENFT_BALANCE_LESS_THAN_RESERVE(
      overrides?: CallOverrides
    ): Promise<[string]>;

    STAKENFT_CALLER_NOT_TOKEN_OWNER(
      overrides?: CallOverrides
    ): Promise<[string]>;

    STAKENFT_FREE_AFTER_TIME_NOT_REACHED(
      overrides?: CallOverrides
    ): Promise<[string]>;

    STAKENFT_INVALID_TOKEN_ID(overrides?: CallOverrides): Promise<[string]>;

    STAKENFT_LOCK_DURATION_GREATER_THAN_GOVERNANCE_LOCK(
      overrides?: CallOverrides
    ): Promise<[string]>;

    STAKENFT_LOCK_DURATION_GREATER_THAN_MINT_LOCK(
      overrides?: CallOverrides
    ): Promise<[string]>;

    STAKENFT_LOCK_DURATION_WITHDRAW_TIME_NOT_REACHED(
      overrides?: CallOverrides
    ): Promise<[string]>;

    STAKENFT_MINT_AMOUNT_EXCEEDS_MAX_SUPPLY(
      overrides?: CallOverrides
    ): Promise<[string]>;

    STAKENFT_SLUSH_TOO_LARGE(overrides?: CallOverrides): Promise<[string]>;
  };

  STAKENFT_BALANCE_LESS_THAN_RESERVE(
    overrides?: CallOverrides
  ): Promise<string>;

  STAKENFT_CALLER_NOT_TOKEN_OWNER(overrides?: CallOverrides): Promise<string>;

  STAKENFT_FREE_AFTER_TIME_NOT_REACHED(
    overrides?: CallOverrides
  ): Promise<string>;

  STAKENFT_INVALID_TOKEN_ID(overrides?: CallOverrides): Promise<string>;

  STAKENFT_LOCK_DURATION_GREATER_THAN_GOVERNANCE_LOCK(
    overrides?: CallOverrides
  ): Promise<string>;

  STAKENFT_LOCK_DURATION_GREATER_THAN_MINT_LOCK(
    overrides?: CallOverrides
  ): Promise<string>;

  STAKENFT_LOCK_DURATION_WITHDRAW_TIME_NOT_REACHED(
    overrides?: CallOverrides
  ): Promise<string>;

  STAKENFT_MINT_AMOUNT_EXCEEDS_MAX_SUPPLY(
    overrides?: CallOverrides
  ): Promise<string>;

  STAKENFT_SLUSH_TOO_LARGE(overrides?: CallOverrides): Promise<string>;

  callStatic: {
    STAKENFT_BALANCE_LESS_THAN_RESERVE(
      overrides?: CallOverrides
    ): Promise<string>;

    STAKENFT_CALLER_NOT_TOKEN_OWNER(overrides?: CallOverrides): Promise<string>;

    STAKENFT_FREE_AFTER_TIME_NOT_REACHED(
      overrides?: CallOverrides
    ): Promise<string>;

    STAKENFT_INVALID_TOKEN_ID(overrides?: CallOverrides): Promise<string>;

    STAKENFT_LOCK_DURATION_GREATER_THAN_GOVERNANCE_LOCK(
      overrides?: CallOverrides
    ): Promise<string>;

    STAKENFT_LOCK_DURATION_GREATER_THAN_MINT_LOCK(
      overrides?: CallOverrides
    ): Promise<string>;

    STAKENFT_LOCK_DURATION_WITHDRAW_TIME_NOT_REACHED(
      overrides?: CallOverrides
    ): Promise<string>;

    STAKENFT_MINT_AMOUNT_EXCEEDS_MAX_SUPPLY(
      overrides?: CallOverrides
    ): Promise<string>;

    STAKENFT_SLUSH_TOO_LARGE(overrides?: CallOverrides): Promise<string>;
  };

  filters: {};

  estimateGas: {
    STAKENFT_BALANCE_LESS_THAN_RESERVE(
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    STAKENFT_CALLER_NOT_TOKEN_OWNER(
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    STAKENFT_FREE_AFTER_TIME_NOT_REACHED(
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    STAKENFT_INVALID_TOKEN_ID(overrides?: CallOverrides): Promise<BigNumber>;

    STAKENFT_LOCK_DURATION_GREATER_THAN_GOVERNANCE_LOCK(
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    STAKENFT_LOCK_DURATION_GREATER_THAN_MINT_LOCK(
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    STAKENFT_LOCK_DURATION_WITHDRAW_TIME_NOT_REACHED(
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    STAKENFT_MINT_AMOUNT_EXCEEDS_MAX_SUPPLY(
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    STAKENFT_SLUSH_TOO_LARGE(overrides?: CallOverrides): Promise<BigNumber>;
  };

  populateTransaction: {
    STAKENFT_BALANCE_LESS_THAN_RESERVE(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    STAKENFT_CALLER_NOT_TOKEN_OWNER(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    STAKENFT_FREE_AFTER_TIME_NOT_REACHED(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    STAKENFT_INVALID_TOKEN_ID(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    STAKENFT_LOCK_DURATION_GREATER_THAN_GOVERNANCE_LOCK(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    STAKENFT_LOCK_DURATION_GREATER_THAN_MINT_LOCK(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    STAKENFT_LOCK_DURATION_WITHDRAW_TIME_NOT_REACHED(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    STAKENFT_MINT_AMOUNT_EXCEEDS_MAX_SUPPLY(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    STAKENFT_SLUSH_TOO_LARGE(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;
  };
}
