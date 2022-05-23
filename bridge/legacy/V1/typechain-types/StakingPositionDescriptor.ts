/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { FunctionFragment, Result } from "@ethersproject/abi";
import { Listener, Provider } from "@ethersproject/providers";
import {
  BaseContract,
  BigNumber,
  BigNumberish,
  BytesLike,
  CallOverrides,
  PopulatedTransaction,
  Signer,
  utils,
} from "ethers";
import { OnEvent, TypedEvent, TypedEventFilter, TypedListener } from "./common";

export interface StakingPositionDescriptorInterface extends utils.Interface {
  contractName: "StakingPositionDescriptor";
  functions: {
    "tokenURI(address,uint256)": FunctionFragment;
  };

  encodeFunctionData(
    functionFragment: "tokenURI",
    values: [string, BigNumberish]
  ): string;

  decodeFunctionResult(functionFragment: "tokenURI", data: BytesLike): Result;

  events: {};
}

export interface StakingPositionDescriptor extends BaseContract {
  contractName: "StakingPositionDescriptor";
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: StakingPositionDescriptorInterface;

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
    tokenURI(
      _stakingNFT: string,
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[string]>;
  };

  tokenURI(
    _stakingNFT: string,
    tokenId: BigNumberish,
    overrides?: CallOverrides
  ): Promise<string>;

  callStatic: {
    tokenURI(
      _stakingNFT: string,
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<string>;
  };

  filters: {};

  estimateGas: {
    tokenURI(
      _stakingNFT: string,
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    tokenURI(
      _stakingNFT: string,
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;
  };
}
