import { expect } from "chai";
import { BytesLike, ContractReceipt } from "ethers";
import { ethers, run } from "hardhat";
import { Receipt } from "hardhat-deploy/dist/types";
import {
  CONTRACT_ADDR,
  DEPLOYED_RAW,
  DEPLOYED_STATIC,
  DEPLOY_CREATE,
  DEPLOY_FACTORY,
  DEPLOY_METAMORPHIC,
  DEPLOY_PROXY,
  DEPLOY_STATIC,
  DEPLOY_TEMPLATE,
  DEPLOY_UPGRADEABLE_PROXY,
  MADNET_FACTORY,
  MOCK,
  MOCK_INITIALIZABLE,
  MULTI_CALL_DEPLOY_PROXY,
  MULTI_CALL_UPGRADE_PROXY,
  UPGRADE_DEPLOYED_PROXY,
  UTILS,
} from "../../scripts/lib/constants";
import { deployFactory } from "../../scripts/lib/deployment/deploymentUtil";
import {
  DeployCreateData,
  MetaContractData,
  ProxyData,
  TemplateData,
} from "../../scripts/lib/deployment/factoryStateUtil";
import { getSalt } from "../../scripts/lib/MadnetFactory";
import {
  getAccounts,
  getEventVar,
  getMetamorphicAddress,
  getReceiptEventVar,
  predictFactoryAddress,
} from "./Setup";

describe("Cli tasks", async () => {
  let firstOwner: string;
  let firstDelegator: string;
  const accounts: Array<string> = [];

  beforeEach(async () => {
    process.env.test = "true";
    process.env.silencer = "true";
    // set owner and delegator
    firstOwner = accounts[0];
    firstDelegator = accounts[1];
  });

  xit("deploys factory with cli and checks if the default factory is updated in factory state toml file", async () => {
    const accounts = await getAccounts();
    const futureFactoryAddress = await predictFactoryAddress(accounts[0]);
    const factoryAddress = await run(DEPLOY_FACTORY);
    // check if the address is the predicted
    expect(factoryAddress).to.equal(futureFactoryAddress);
  });
  // todo add init call data and check init vars
  xit("deploys MockInitializable contract with deployUpgradeableProxy", async () => {
    // deploys factory using the deployFactory task
    const factory = await deployFactory(run);
    const proxyData: ProxyData = await run(DEPLOY_UPGRADEABLE_PROXY, {
      contractName: MOCK_INITIALIZABLE,
      initCallData: "14",
    });
    const expectedProxyAddress = getMetamorphicAddress(
      factory,
      ethers.utils.formatBytes32String(MOCK_INITIALIZABLE)
    );
    expect(proxyData.proxyAddress).to.equal(expectedProxyAddress);
  });
  // todo check mock logic
  xit("deploys mock contract with deployStatic", async () => {
    // deploys factory using the deployFactory task
    let factoryAddress = await cliDeployFactory();
    const metaData = await cliDeployMetamorphic(MOCK, factoryAddress, undefined, [
      "2",
      "s",
    ]);
    const salt = ethers.utils.formatBytes32String("Mock");
    const expectedMetaAddr = getMetamorphicAddress(factoryAddress, salt);
    expect(metaData.metaAddress).to.equal(expectedMetaAddr);
  });

  xit("deploys MockInitializable contract with deployCreate", async () => {
    let factoryAddress = await cliDeployFactory();
    const deployCreateData = await cliDeployCreate(MOCK_INITIALIZABLE, factoryAddress);
  });

  xit("deploys MockInitializable with deploy create, deploys proxy, then upgrades proxy to point to MockInitializable with initCallData", async () => {
    let factoryAddress = await cliDeployFactory();
    const test = "1";
    const deployCreateData = await cliDeployCreate(MOCK_INITIALIZABLE, factoryAddress);
    const salt = (await getSalt(MOCK_INITIALIZABLE)) as string;
    const proxyData = await cliDeployProxy(salt, factoryAddress);
    const logicFactory = await ethers.getContractFactory(MOCK_INITIALIZABLE);
    const upgradedProxyData = await cliUpgradeDeployedProxy(
      MOCK_INITIALIZABLE,
      deployCreateData.address,
      factoryAddress,
      test
    );
    const mockContract = logicFactory.attach(upgradedProxyData.proxyAddress);
    const i = await mockContract.callStatic.getImut();
    expect(i.toNumber()).to.equal(parseInt(test, 10));
  });

  xit("deploys mock contract with deployTemplate then deploys a metamorphic contract", async () => {
    let factoryAddress = await cliDeployFactory();
    const testVar1 = "14";
    const testVar2 = "s";
    const templateData = await cliDeployTemplate(MOCK, factoryAddress, [
      testVar1,
      testVar2,
    ]);
    const metaData = await cliDeployStatic(MOCK, factoryAddress, undefined);
    const logicFactory = await ethers.getContractFactory(MOCK);
    const mockContract = logicFactory.attach(metaData.metaAddress);
    const i = await mockContract.callStatic.getImut();
    expect(i.toNumber()).to.equal(parseInt(testVar1, 10));
    const pString = await mockContract.callStatic.getpString();
    expect(pString).to.equal(testVar2);
  });

  xit("deploys mockInitializable with deployCreate, then deploy and upgrades a proxy with multiCallDeployProxy", async () => {
    let factoryAddress = await cliDeployFactory();
    const logicData = await cliDeployCreate(MOCK_INITIALIZABLE, factoryAddress);
    const proxyData = await cliMultiCallDeployProxy(
      MOCK_INITIALIZABLE,
      logicData.address,
      factoryAddress,
      "1"
    );
  });

  xit("deploys mock with deployCreate", async () => {
    let factoryAddress = await cliDeployFactory();
    const deployCreateData = await cliDeployCreate(MOCK, factoryAddress, ["2", "s"]);
  });

  it("deploys MadByte with deployMetamorphic", async () => {
    let factoryAddress = await cliDeployFactory();
    let utilsBase = await ethers.getContractFactory(UTILS)
    let utils = await utilsBase.deploy();
    let metaContractData = await cliDeployMetamorphic("MadByte", factoryAddress, undefined, ["0x"])
    let receipt = metaContractData.receipt as ContractReceipt
    console.log(metaContractData.gas)
  });

  it("deploys MadByte with multiCallDeployMetamorphic", async () => {
    let factoryAddress = await cliDeployFactory();
    let receipt = await cliMultiCallDeployMetamorphic("MadByte", factoryAddress)
    //console.log(receipt.events, receipt.logs)
    let templateAddress = await getReceiptEventVar(receipt, DEPLOYED_RAW, CONTRACT_ADDR)
    let utilsBase = await ethers.getContractFactory(UTILS)
    let utils = await utilsBase.deploy();
    let codesize = await utils.callStatic.getCodeSize(templateAddress)
    let metaAddress = await getReceiptEventVar(receipt, DEPLOYED_STATIC, CONTRACT_ADDR)
    let metaCodeSize = await utils.callStatic.getCodeSize(metaAddress)
    console.log(receipt.cumulativeGasUsed)
    console.log(`template Size: ${codesize}`)
    console.log(`metamorphic size: ${metaCodeSize}`)
  });
});
export async function cliMultiCallDeployMetamorphic(contractName: string,
  factoryAddress: string,
  initCallData?: string,
  constructorArgs?: Array<string>
  ): Promise<ContractReceipt>{
    return await run("multiCallDeployMetamorphic", {
      contractName: contractName,
      factoryAddress: factoryAddress,
      initCallData: initCallData,
      constructorArgs: constructorArgs
    })
  }
export async function cliDeployUpgradeableProxy(
  contractName: string,
  factoryAddress: string,
  initCallData?: string,
  constructorArgs?: Array<string>
): Promise<ProxyData> {
  return await run(DEPLOY_UPGRADEABLE_PROXY, {
    contractName,
    factoryAddress,
    initCallData,
    constructorArgs,
  });
}

export async function cliDeployMetamorphic(
  contractName: string,
  factoryAddress: string,
  initCallData?: string,
  constructorArgs?: Array<string>
): Promise<MetaContractData> {
  return await run(DEPLOY_METAMORPHIC, {
    contractName,
    factoryAddress,
    initCallData,
    constructorArgs,
  });
}

export async function cliDeployCreate(
  contractName: string,
  factoryAddress: string,
  constructorArgs?: Array<string>
): Promise<DeployCreateData> {
  return await run(DEPLOY_CREATE, {
    contractName,
    factoryAddress,
    constructorArgs,
  });
}

export async function cliUpgradeDeployedProxy(
  contractName: string,
  logicAddress: string,
  factoryAddress: string,
  initCallData?: string
): Promise<ProxyData> {
  return await run(UPGRADE_DEPLOYED_PROXY, {
    contractName,
    logicAddress,
    factoryAddress,
    initCallData,
  });
}

export async function cliDeployTemplate(
  contractName: string,
  factoryAddress: string,
  constructorArgs?: Array<string>
): Promise<TemplateData> {
  return await run(DEPLOY_TEMPLATE, {
    contractName,
    factoryAddress,
    constructorArgs,
  });
}

export async function cliDeployStatic(
  contractName: string,
  factoryAddress: string,
  initCallData?: Array<string>
): Promise<MetaContractData> {
  return await run(DEPLOY_STATIC, {
    contractName,
    factoryAddress,
    initCallData,
  });
}

export async function cliMultiCallDeployProxy(
  contractName: string,
  logicAddress: string,
  factoryAddress: string,
  initCallData?: string,
  salt?: string
): Promise<ProxyData> {
  return await run(MULTI_CALL_DEPLOY_PROXY, {
    contractName,
    logicAddress,
    factoryAddress,
    initCallData,
    salt,
  });
}

export async function cliMultiCallUpgradeProxy(
  contractName: string,
  factoryAddress: BytesLike,
  initCallData?: BytesLike,
  salt?: BytesLike,
  constructorArgs?: Array<string>
): Promise<ProxyData> {
  return await run(MULTI_CALL_UPGRADE_PROXY, {
    contractName,
    factoryAddress,
    initCallData,
    salt,
    constructorArgs,
  });
}

export async function cliDeployFactory() {
  return await run(DEPLOY_FACTORY);
}

export async function cliDeployProxy(
  salt: string,
  factoryAddress: string
): Promise<ProxyData> {
  return await run(DEPLOY_PROXY, {
    salt,
    factoryAddress,
  });
}
