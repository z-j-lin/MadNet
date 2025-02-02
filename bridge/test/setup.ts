import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/dist/src/signer-with-address";
import {
  BigNumber,
  BigNumberish,
  BytesLike,
  Contract,
  ContractTransaction,
  Signer,
  Wallet,
} from "ethers";
import { isHexString } from "ethers/lib/utils";
import { ethers, network } from "hardhat";
import { Artifact } from "hardhat/types";
import SnapshotsV1Artifact from "../../legacy/V1/artifacts/contracts/Snapshots.sol/Snapshots.json";
import { INITIALIZER } from "../scripts/lib/constants";
import {
  AliceNetFactory,
  AToken,
  ATokenBurner,
  ATokenMinter,
  BToken,
  ETHDKG,
  Foundation,
  LegacyToken,
  LiquidityProviderStaking,
  PublicStaking,
  Snapshots,
  StakingPositionDescriptor,
  ValidatorPool,
  ValidatorPoolMock,
  ValidatorStaking,
} from "../typechain-types";
import { completeETHDKGRound, ValidatorRawData } from "./ethdkg/setup";
import {
  signedData1,
  validatorsSnapshotsG1,
} from "./sharedConstants/4-validators-snapshots-100-Group1";
import { createValidators, stakeValidators } from "./validatorPool/setup";

export const SNAPSHOT_BUFFER_LENGTH = 6;

export const PLACEHOLDER_ADDRESS = "0x0000000000000000000000000000000000000000";
export { assert, expect } from "./chai-setup";
export interface SignedBClaims {
  BClaims: string;
  GroupSignature: string;
}
export interface Snapshot {
  BClaims: string;
  GroupSignature: string;
  height: BigNumberish;
  validatorIndex: number;
}

export interface BaseFixture {
  factory: AliceNetFactory;
  [key: string]: any;
}

export interface BaseTokensFixture extends BaseFixture {
  aToken: AToken;
  bToken: BToken;
  legacyToken: LegacyToken;
  publicStaking: PublicStaking;
}

export interface Fixture extends BaseTokensFixture {
  validatorStaking: ValidatorStaking;
  validatorPool: ValidatorPool | ValidatorPoolMock;
  snapshots: Snapshots;
  ethdkg: ETHDKG;
  stakingPositionDescriptor: StakingPositionDescriptor;
  namedSigners: SignerWithAddress[];
}

/**
 * Shuffles array in place. ES6 version
 * https://stackoverflow.com/questions/6274339/how-can-i-shuffle-an-array/6274381#6274381
 * @param {Array} a items An array containing the items.
 */
export function shuffle(a: ValidatorRawData[]) {
  for (let i = a.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [a[i], a[j]] = [a[j], a[i]];
  }
  return a;
}

export const mineBlocks = async (nBlocks: bigint) => {
  if (nBlocks > BigInt(0)) {
    await network.provider.send("hardhat_mine", [
      ethers.utils.hexValue(nBlocks),
    ]);
  }
};

export const getBlockByNumber = async () => {
  return await network.provider.send("eth_getBlockByNumber", [
    "pending",
    false,
  ]);
};

export const getPendingTransactions = async () => {
  return await network.provider.send("eth_pendingTransactions");
};

export const getValidatorEthAccount = async (
  validator: ValidatorRawData | string
): Promise<Signer> => {
  if (typeof validator === "string") {
    return ethers.getSigner(validator);
  } else {
    const balance = await ethers.provider.getBalance(validator.address);
    if (balance.eq(0)) {
      await (
        await ethers.getSigners()
      )[0].sendTransaction({
        to: validator.address,
        value: ethers.utils.parseEther("10"),
      });
    }
    if (typeof validator.privateKey !== "undefined") {
      return new Wallet(validator.privateKey, ethers.provider);
    }
    return ethers.getSigner(validator.address);
  }
};

export const createUsers = async (
  numberOfUsers: number,
  createWithNoFunds: boolean = false
): Promise<SignerWithAddress[]> => {
  const users: SignerWithAddress[] = [];
  const admin = (await ethers.getSigners())[0];
  for (let i = 0; i < numberOfUsers; i++) {
    const user = new Wallet(ethers.utils.randomBytes(64), ethers.provider);
    if (!createWithNoFunds) {
      const balance = await ethers.provider.getBalance(user.address);
      if (balance.eq(0)) {
        await admin.sendTransaction({
          to: user.address,
          value: ethers.utils.parseEther("1"),
        });
      }
    }
    users.push(user as Signer as SignerWithAddress);
  }
  return users;
};

async function getContractAddressFromDeployedStaticEvent(
  tx: ContractTransaction
): Promise<string> {
  const eventSignature = "event DeployedStatic(address contractAddr)";
  const eventName = "DeployedStatic";
  return await getContractAddressFromEventLog(tx, eventSignature, eventName);
}

async function getContractAddressFromDeployedProxyEvent(
  tx: ContractTransaction
): Promise<string> {
  const eventSignature = "event DeployedProxy(address contractAddr)";
  const eventName = "DeployedProxy";
  return await getContractAddressFromEventLog(tx, eventSignature, eventName);
}

async function getContractAddressFromDeployedRawEvent(
  tx: ContractTransaction
): Promise<string> {
  const eventSignature = "event DeployedRaw(address contractAddr)";
  const eventName = "DeployedRaw";
  return await getContractAddressFromEventLog(tx, eventSignature, eventName);
}

async function getContractAddressFromEventLog(
  tx: ContractTransaction,
  eventSignature: string,
  eventName: string
): Promise<string> {
  const receipt = await ethers.provider.getTransactionReceipt(tx.hash);
  const intrface = new ethers.utils.Interface([eventSignature]);
  let result = "";
  for (const log of receipt.logs) {
    const topics = log.topics;
    const data = log.data;
    const topicHash = intrface.getEventTopic(intrface.getEvent(eventName));
    if (!isHexString(topics[0], 32) || topics[0].toLowerCase() !== topicHash) {
      continue;
    }
    result = intrface.decodeEventLog(eventName, data, topics).contractAddr;
  }
  if (result === "") {
    throw new Error(
      "Couldn't parse logs in the transaction!\nReceipt:\n" + receipt
    );
  }
  return result;
}

function getBytes32Salt(contractName: string) {
  return ethers.utils.formatBytes32String(contractName);
}

export const deployStaticWithFactory = async (
  factory: AliceNetFactory,
  contractName: string,
  salt?: string,
  initCallData?: any[],
  constructorArgs: any[] = []
): Promise<Contract> => {
  const _Contract = await ethers.getContractFactory(contractName);
  const contractTx = await factory.deployTemplate(
    _Contract.getDeployTransaction(...constructorArgs).data as BytesLike
  );
  let receipt = await ethers.provider.getTransactionReceipt(contractTx.hash);
  if (receipt.gasUsed.gt(10_000_000)) {
    throw new Error(
      `Contract deployment size:${receipt.gasUsed} is greater than 10 million`
    );
  }
  let initCallDataBin;

  try {
    initCallDataBin = _Contract.interface.encodeFunctionData(
      "initialize",
      initCallData
    );
  } catch (error) {
    console.log(
      `Warning couldnt get init call data for contract: ${contractName}`
    );
    console.log(error);
    initCallDataBin = "0x";
  }
  let saltBytes;
  if (salt === undefined) {
    saltBytes = getBytes32Salt(contractName) as string;
  } else {
    saltBytes = getBytes32Salt(salt);
  }

  const tx = await factory.deployStatic(saltBytes, initCallDataBin);
  receipt = await ethers.provider.getTransactionReceipt(tx.hash);
  if (receipt.gasUsed.gt(10_000_000)) {
    throw new Error(
      `Contract deployment size:${receipt.gasUsed} is greater than 10 million`
    );
  }

  return _Contract.attach(await getContractAddressFromDeployedStaticEvent(tx));
};

export const deployUpgradeableWithFactory = async (
  factory: AliceNetFactory,
  contractName: string,
  salt?: string,
  initCallData?: any[],
  constructorArgs: any[] = []
): Promise<Contract> => {
  const _Contract = await ethers.getContractFactory(contractName);
  let deployCode = _Contract.getDeployTransaction(...constructorArgs)
    .data as BytesLike;
  const transaction = await factory.deployCreate(deployCode);
  let receipt = await transaction.wait();
  if (receipt.gasUsed.gt(10_000_000)) {
    throw new Error(
      `Contract deployment size:${receipt.gasUsed} is greater than 10 million`
    );
  }
  const logicAddr = await getContractAddressFromDeployedRawEvent(transaction);
  let saltBytes;
  if (salt === undefined) {
    saltBytes = getBytes32Salt(contractName) as string;
  } else {
    saltBytes = getBytes32Salt(salt);
  }

  const transaction2 = await factory.deployProxy(saltBytes);
  receipt = await ethers.provider.getTransactionReceipt(transaction2.hash);
  if (receipt.gasUsed.gt(10_000_000)) {
    throw new Error(
      `Contract deployment size:${receipt.gasUsed} is greater than 10 million`
    );
  }

  let initCallDataBin = "0x";
  if (initCallData !== undefined) {
    try {
      initCallDataBin = _Contract.interface.encodeFunctionData(
        "initialize",
        initCallData
      );
    } catch (error) {
      console.warn(
        `Error deploying contract ${contractName} couldn't get initialize arguments: ${error}`
      );
    }
  }
  await factory.upgradeProxy(saltBytes, logicAddr, initCallDataBin);
  return _Contract.attach(
    await getContractAddressFromDeployedProxyEvent(transaction2)
  );
};

export const deployUpgradeableV1WithFactory = async (
  factory: AliceNetFactory,
  contractArtifact: Artifact,
  salt?: string,
  initCallData?: any[],
  constructorArgs: any[] = []
): Promise<Contract> => {
  const _Contract = new ethers.ContractFactory(
    contractArtifact.abi,
    contractArtifact.bytecode,
    factory.signer
  );
  const deployCode = (await _Contract.getDeployTransaction(...constructorArgs)
    .data) as BytesLike;
  const transaction = await factory.deployCreate(deployCode);
  let receipt = await transaction.wait();
  if (receipt.gasUsed.gt(10_000_000)) {
    throw new Error(
      `Contract deployment size:${receipt.gasUsed} is greater than 10 million`
    );
  }
  const logicAddr = await getContractAddressFromDeployedRawEvent(transaction);
  let saltBytes;
  if (salt === undefined) {
    saltBytes = getBytes32Salt(contractArtifact.contractName) as string;
  } else {
    saltBytes = getBytes32Salt(salt);
  }

  const transaction2 = await factory.deployProxy(saltBytes);
  receipt = await ethers.provider.getTransactionReceipt(transaction2.hash);
  if (receipt.gasUsed.gt(10_000_000)) {
    throw new Error(
      `Contract deployment size:${receipt.gasUsed} is greater than 10 million`
    );
  }

  let initCallDataBin = "0x";
  if (initCallData !== undefined) {
    try {
      initCallDataBin = _Contract.interface.encodeFunctionData(
        "initialize",
        initCallData
      );
    } catch (error) {
      console.warn(
        `Error deploying contract ${contractArtifact.contractName} couldn't get initialize arguments: ${error}`
      );
    }
  }
  await factory.upgradeProxy(saltBytes, logicAddr, initCallDataBin);
  let address = await getContractAddressFromDeployedProxyEvent(transaction2);
  return _Contract.attach(address);
};

export const deployLogicAndUpgradeWithFactory = async (
  factory: AliceNetFactory,
  contractName: string,
  proxyAddress: string,
  salt?: string,
  initCallData?: any[],
  constructorArgs: any[] = []
): Promise<Contract> => {
  const _Contract = await ethers.getContractFactory(contractName);
  let deployTx = _Contract.getDeployTransaction(...constructorArgs);
  const logicDeploymentCode = deployTx.data as BytesLike;
  const transaction = await factory.deployCreate(logicDeploymentCode, {
    gasLimit: 30000000,
  });
  const receipt = await transaction.wait();
  if (receipt.gasUsed.gt(10_000_000)) {
    throw new Error(
      `Contract deployment size:${receipt.gasUsed} is greater than 10 million`
    );
  }
  const logicAddr = await getContractAddressFromDeployedRawEvent(transaction);
  let saltBytes;
  if (salt === undefined) {
    saltBytes = getBytes32Salt(contractName);
  } else {
    saltBytes = getBytes32Salt(salt);
  }
  let initCallDataBin = "0x";
  if (initCallData !== undefined) {
    try {
      initCallDataBin = _Contract.interface.encodeFunctionData(
        INITIALIZER,
        initCallData
      );
    } catch (error) {
      console.warn(
        `Error deploying contract ${contractName} couldn't get initialize arguments: ${error}`
      );
    }
  }
  await factory.upgradeProxy(saltBytes, logicAddr, initCallDataBin, {
    gasLimit: 30000000,
  });
  return await ethers.getContractAt(contractName, proxyAddress);
};

export const deployFactoryAndBaseTokens = async (
  admin: SignerWithAddress
): Promise<BaseTokensFixture> => {
  const factory = await deployAliceNetFactory(admin);
  // LegacyToken
  const legacyToken = (await deployStaticWithFactory(
    factory,
    "LegacyToken"
  )) as LegacyToken;

  const aToken = (await deployStaticWithFactory(
    factory,
    "AToken",
    "AToken",
    undefined,
    [legacyToken.address]
  )) as AToken;

  // BToken
  const bToken = (await deployStaticWithFactory(factory, "BToken")) as BToken;

  // PublicStaking
  const publicStaking = (await deployUpgradeableWithFactory(
    factory,
    "PublicStaking",
    "PublicStaking",
    []
  )) as PublicStaking;

  return {
    factory,
    aToken,
    bToken,
    legacyToken,
    publicStaking,
  };
};

export const deployAliceNetFactory = async (
  admin: SignerWithAddress
): Promise<AliceNetFactory> => {
  const txCount = await ethers.provider.getTransactionCount(admin.address);
  // calculate the factory address for the constructor arg
  const futureFactoryAddress = ethers.utils.getContractAddress({
    from: admin.address,
    nonce: txCount,
  });

  const Factory = await ethers.getContractFactory("AliceNetFactory");
  const factory = await Factory.deploy(futureFactoryAddress);
  await factory.deployed();
  return factory;
};

export const preFixtureSetup = async () => {
  await network.provider.send("evm_setAutomine", [true]);
  // hardhat is not being able to estimate correctly the tx gas due to the massive bytes array
  // being sent as input to the function (the contract bytecode), so we need to increase the block
  // gas limit temporally in order to deploy the template
  await network.provider.send("evm_setBlockGasLimit", ["0x3000000000000000"]);
};

export const posFixtureSetup = async (
  factory: AliceNetFactory,
  aToken: AToken,
  legacyToken: LegacyToken
) => {
  // finish workaround, putting the blockgas limit to the previous value 30_000_000
  await network.provider.send("evm_setBlockGasLimit", ["0x1C9C380"]);
  const [admin] = await ethers.getSigners();
  // transferring some part of the legacy token from the factory to the admin
  await factory.callAny(
    legacyToken.address,
    0,
    aToken.interface.encodeFunctionData("transfer", [
      admin.address,
      ethers.utils.parseEther("100000000"),
    ])
  );
  // migrating the rest of the legacy tokens to fresh new Atokens
  await factory.callAny(
    legacyToken.address,
    0,
    aToken.interface.encodeFunctionData("approve", [
      aToken.address,
      ethers.utils.parseEther("100000000"),
    ])
  );
  await factory.callAny(
    aToken.address,
    0,
    aToken.interface.encodeFunctionData("allowMigration")
  );
  await factory.callAny(
    aToken.address,
    0,
    aToken.interface.encodeFunctionData("migrate", [
      ethers.utils.parseEther("100000000"),
    ])
  );
  // transferring those Atokens to the admin
  await factory.callAny(
    aToken.address,
    0,
    aToken.interface.encodeFunctionData("transfer", [
      admin.address,
      ethers.utils.parseEther("100000000"),
    ])
  );
};

export const getBaseTokensFixture = async (): Promise<BaseTokensFixture> => {
  await preFixtureSetup();
  const [admin] = await ethers.getSigners();
  // AToken
  const fixture = await deployFactoryAndBaseTokens(admin);
  await posFixtureSetup(fixture.factory, fixture.aToken, fixture.legacyToken);
  return fixture;
};
/**
 *
 * @param mockValidatorPool deploys a mock validator pool
 * @param mockSnapshots deploys the mock snapshots contract
 * @param mockETHDKG deploys mock ethdkg
 * @param runEthDKG runs ethdkg and performs 6 snapshots
 * @param upgradeSnapshots deploys current snapshot contract and upgrades snapshots proxy
 * @param snapshotsV1 deploys snapshot V1 initialy
 * @returns
 */
export const getFixture = async (
  mockValidatorPool?: boolean,
  mockSnapshots?: boolean,
  mockETHDKG?: boolean,
  runEthDKG?: boolean,
  upgradeSnapshots?: boolean,
  snapshotsV1?: boolean,
  doSnapshots?: boolean
): Promise<Fixture> => {
  await preFixtureSetup();
  const namedSigners = await ethers.getSigners();
  const [admin] = namedSigners;
  // Deploy the base tokens
  const { factory, aToken, bToken, legacyToken, publicStaking } =
    await deployFactoryAndBaseTokens(admin);

  // ValidatorStaking is not considered a base token since is only used by validators
  const validatorStaking = (await deployUpgradeableWithFactory(
    factory,
    "ValidatorStaking",
    "ValidatorStaking",
    []
  )) as ValidatorStaking;

  // LiquidityProviderStaking
  const liquidityProviderStaking = (await deployUpgradeableWithFactory(
    factory,
    "LiquidityProviderStaking",
    "LiquidityProviderStaking",
    []
  )) as LiquidityProviderStaking;

  // Foundation
  const foundation = (await deployUpgradeableWithFactory(
    factory,
    "Foundation",
    undefined
  )) as Foundation;
  let validatorPool;
  if (typeof mockValidatorPool !== "undefined" && mockValidatorPool) {
    // ValidatorPoolMock
    validatorPool = (await deployUpgradeableWithFactory(
      factory,
      "ValidatorPoolMock",
      "ValidatorPool"
    )) as ValidatorPoolMock;
  } else {
    // ValidatorPool
    validatorPool = (await deployUpgradeableWithFactory(
      factory,
      "ValidatorPool",
      "ValidatorPool",
      [
        ethers.utils.parseUnits("20000", 18),
        10,
        ethers.utils.parseUnits("3", 18),
      ]
    )) as ValidatorPool;
  }

  // ETHDKG Accusations
  await deployUpgradeableWithFactory(factory, "ETHDKGAccusations");

  // StakingPositionDescriptor
  const stakingPositionDescriptor = (await deployUpgradeableWithFactory(
    factory,
    "StakingPositionDescriptor"
  )) as StakingPositionDescriptor;

  // ETHDKG Phases
  await deployUpgradeableWithFactory(factory, "ETHDKGPhases");
  // ETHDKG
  let ethdkg;
  if (typeof mockETHDKG !== "undefined" && mockETHDKG) {
    // ValidatorPoolMock
    ethdkg = (await deployUpgradeableWithFactory(
      factory,
      "ETHDKGMock",
      "ETHDKG",
      [BigNumber.from(40), BigNumber.from(6)]
    )) as ETHDKG;
  } else {
    // ValidatorPool
    ethdkg = (await deployUpgradeableWithFactory(factory, "ETHDKG", "ETHDKG", [
      BigNumber.from(40),
      BigNumber.from(6),
    ])) as ETHDKG;
  }
  let snapshots: any;
  if (typeof mockSnapshots !== "undefined" && mockSnapshots) {
    // Snapshots Mock
    snapshots = (await deployUpgradeableWithFactory(
      factory,
      "SnapshotsMock",
      "Snapshots",
      [10, 40],
      [1, 1]
    )) as Snapshots;
  } else if (upgradeSnapshots || snapshotsV1) {
    // Snapshots
    snapshots = await deployUpgradeableV1WithFactory(
      factory,
      SnapshotsV1Artifact,
      "Snapshots",
      [10, 40],
      [1, 1024]
    );
  } else {
    snapshots = (await deployUpgradeableWithFactory(
      factory,
      "Snapshots",
      "Snapshots",
      [10, 40],
      [1, 1024]
    )) as Snapshots;
  }

  const aTokenMinter = (await deployUpgradeableWithFactory(
    factory,
    "ATokenMinter",
    "ATokenMinter"
  )) as ATokenMinter;
  const aTokenBurner = (await deployUpgradeableWithFactory(
    factory,
    "ATokenBurner",
    "ATokenBurner"
  )) as ATokenBurner;

  await posFixtureSetup(factory, aToken, legacyToken);

  const blockNumber = BigInt(await ethers.provider.getBlockNumber());
  const phaseLength = (await ethdkg.getPhaseLength()).toBigInt();
  if (phaseLength >= blockNumber) {
    await mineBlocks(phaseLength);
  }
  const signedSnapshots = signedData1;
  if (runEthDKG) {
    if (mockETHDKG !== true && mockValidatorPool !== true) {
      const validators = await createValidators(
        factory,
        aToken,
        validatorPool,
        publicStaking,
        validatorStaking,
        validatorsSnapshotsG1
      );
      const stakingTokenIds = await stakeValidators(
        factory,
        aToken,
        validatorPool,
        publicStaking,
        validatorStaking,
        validators
      );
      await factoryCallAny(factory, validatorPool, "registerValidators", [
        validators,
        stakingTokenIds,
      ]);
      await factoryCallAny(factory, validatorPool, "initializeETHDKG");
    }

    await completeETHDKGRound(validatorsSnapshotsG1, {
      ethdkg: ethdkg,
      validatorPool: validatorPool,
    });
    await mineBlocks(
      (await snapshots.getMinimumIntervalBetweenSnapshots()).toBigInt()
    );

    snapshots = snapshots.connect(
      await getValidatorEthAccount(validatorsSnapshotsG1[0])
    );
  }
  if (doSnapshots) {
    //take 6 snapshots
    for (let i = 0; i < SNAPSHOT_BUFFER_LENGTH; i++) {
      await mineBlocks(
        (await snapshots.getMinimumIntervalBetweenSnapshots()).toBigInt()
      );
      const contractTx = await snapshots.snapshot(
        signedSnapshots[i].GroupSignature,
        signedSnapshots[i].BClaims,
        { gasLimit: 30000000 }
      );
      await contractTx.wait();
    }
  }
  if (upgradeSnapshots === true) {
    //upgrade the snapshot contract
    snapshots = (await deployLogicAndUpgradeWithFactory(
      factory,
      "Snapshots",
      snapshots.address,
      undefined,
      [0, 0],
      [1, 1024]
    )) as Snapshots;
  }

  return {
    aToken,
    bToken,
    legacyToken,
    publicStaking,
    validatorStaking,
    validatorPool,
    snapshots,
    ethdkg,
    factory,
    namedSigners,
    aTokenMinter,
    aTokenBurner,
    liquidityProviderStaking,
    foundation,
    stakingPositionDescriptor,
  };
};

export async function getTokenIdFromTx(tx: any) {
  const abi = [
    "event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)",
  ];
  const iface = new ethers.utils.Interface(abi);
  const receipt = await ethers.provider.getTransactionReceipt(tx.hash);
  const logs =
    typeof receipt.logs[2] !== "undefined" ? receipt.logs[2] : receipt.logs[0];
  const log = iface.parseLog(logs);
  return log.args[2];
}

export async function factoryCallAnyFixture(
  fixture: BaseFixture,
  contractName: string,
  functionName: string,
  args?: Array<any>
) {
  const factory = fixture.factory;
  const contract: Contract = fixture[contractName];
  return await factoryCallAny(factory, contract, functionName, args);
}

export async function factoryCallAny(
  factory: AliceNetFactory,
  contract: Contract,
  functionName: string,
  args?: Array<any>
) {
  if (args === undefined) {
    args = [];
  }
  const txResponse = await factory.callAny(
    contract.address,
    0,
    contract.interface.encodeFunctionData(functionName, args)
  );
  const receipt = await txResponse.wait();
  return receipt;
}

export async function delegateFactoryCallAny(
  factory: AliceNetFactory,
  contract: Contract,
  functionName: string,
  args?: Array<any>
) {
  if (args === undefined) {
    args = [];
  }
  const txResponse = await factory.delegateCallAny(
    contract.address,
    contract.interface.encodeFunctionData(functionName, args)
  );
  const receipt = await txResponse.wait();
  return receipt;
}

export async function callFunctionAndGetReturnValues(
  contract: Contract,
  functionName: any,
  account: SignerWithAddress,
  inputParameters: any[],
  messageValue?: BigNumber
): Promise<[any, ContractTransaction]> {
  try {
    let returnValues;
    let tx;
    if (messageValue !== undefined) {
      returnValues = await contract
        .connect(account)
        .callStatic[functionName](...inputParameters, { value: messageValue });
      tx = await contract
        .connect(account)
        [functionName](...inputParameters, { value: messageValue });
    } else {
      returnValues = await contract
        .connect(account)
        .callStatic[functionName](...inputParameters);
      tx = await contract.connect(account)[functionName](...inputParameters);
    }
    return [returnValues, tx];
  } catch (error) {
    throw new Error(
      `Couldn't call function '${functionName}' with account '${account.address}' and input parameters '${inputParameters}'\n${error}`
    );
  }
}

export const getMetamorphicAddress = (
  factoryAddress: string,
  salt: string
): string => {
  const initCode = "0x6020363636335afa1536363636515af43d36363e3d36f3";
  return ethers.utils.getCreate2Address(
    factoryAddress,
    ethers.utils.formatBytes32String(salt),
    ethers.utils.keccak256(initCode)
  );
};
