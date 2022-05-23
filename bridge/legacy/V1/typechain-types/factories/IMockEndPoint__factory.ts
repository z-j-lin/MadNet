/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Provider } from "@ethersproject/providers";
import { Contract, Signer, utils } from "ethers";
import type { IMockEndPoint, IMockEndPointInterface } from "../IMockEndPoint";

const _abi = [
  {
    inputs: [],
    name: "addOne",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "addTwo",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "factory",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "i",
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

export class IMockEndPoint__factory {
  static readonly abi = _abi;
  static createInterface(): IMockEndPointInterface {
    return new utils.Interface(_abi) as IMockEndPointInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): IMockEndPoint {
    return new Contract(address, _abi, signerOrProvider) as IMockEndPoint;
  }
}
