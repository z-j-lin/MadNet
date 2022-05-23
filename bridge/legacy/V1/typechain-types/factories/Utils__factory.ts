/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Provider, TransactionRequest } from "@ethersproject/providers";
import { Contract, ContractFactory, Overrides, Signer, utils } from "ethers";
import type { Utils, UtilsInterface } from "../Utils";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "addr_",
        type: "address",
      },
    ],
    name: "getCode",
    outputs: [
      {
        internalType: "bytes",
        name: "outputCode",
        type: "bytes",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "target",
        type: "address",
      },
    ],
    name: "getCodeSize",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
];

const _bytecode =
  "0x608060405234801561001057600080fd5b506101af806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80637e105ce21461003b578063b51c4f9614610064575b600080fd5b61004e6100493660046100c9565b610084565b60405161005b9190610106565b60405180910390f35b6100766100723660046100c9565b3b90565b60405190815260200161005b565b60408051603f833b9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092528181529080600060208401853c50919050565b6000602082840312156100db57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146100ff57600080fd5b9392505050565b600060208083528351808285015260005b8181101561013357858101830151858201604001528201610117565b81811115610145576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01692909201604001939250505056fea26469706673582212203eb8193e80178e9cb9d27d4a60297620e10b01f5dbedd628193a5c4075c449de64736f6c634300080d0033";

type UtilsConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: UtilsConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class Utils__factory extends ContractFactory {
  constructor(...args: UtilsConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
    this.contractName = "Utils";
  }

  deploy(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<Utils> {
    return super.deploy(overrides || {}) as Promise<Utils>;
  }
  getDeployTransaction(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  attach(address: string): Utils {
    return super.attach(address) as Utils;
  }
  connect(signer: Signer): Utils__factory {
    return super.connect(signer) as Utils__factory;
  }
  static readonly contractName: "Utils";
  public readonly contractName: "Utils";
  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): UtilsInterface {
    return new utils.Interface(_abi) as UtilsInterface;
  }
  static connect(address: string, signerOrProvider: Signer | Provider): Utils {
    return new Contract(address, _abi, signerOrProvider) as Utils;
  }
}
