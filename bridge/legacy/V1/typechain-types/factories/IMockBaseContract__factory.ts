/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Provider } from "@ethersproject/providers";
import { Contract, Signer, utils } from "ethers";
import type {
  IMockBaseContract,
  IMockBaseContractInterface,
} from "../IMockBaseContract";

const _abi = [
  {
    inputs: [],
    name: "fail",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "getImut",
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
  {
    inputs: [],
    name: "getVar",
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
  {
    inputs: [],
    name: "lock",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "_v",
        type: "uint256",
      },
    ],
    name: "setV",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "unlock",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
];

export class IMockBaseContract__factory {
  static readonly abi = _abi;
  static createInterface(): IMockBaseContractInterface {
    return new utils.Interface(_abi) as IMockBaseContractInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): IMockBaseContract {
    return new Contract(address, _abi, signerOrProvider) as IMockBaseContract;
  }
}
