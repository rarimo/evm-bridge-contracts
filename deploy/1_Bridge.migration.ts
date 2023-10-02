import { Deployer, Logger } from "@solarity/hardhat-migrate";
import { artifacts } from "hardhat";
import config from "./config/config.json";

const BundleExecutorImplementation = artifacts.require("BundleExecutorImplementation");
const ERC1967Proxy = artifacts.require("ERC1967Proxy");
const Bridge = artifacts.require("Bridge");
const BridgeFacade = artifacts.require("BridgeFacade");

export = async (deployer: Deployer, logger: Logger) => {
  const bundleExecutorImplementation = await deployer.deploy(
    BundleExecutorImplementation,
    config.integrations.wrappedNative
  );

  const bridge = await deployer.deploy(Bridge);
  await deployer.deploy(ERC1967Proxy, bridge.address, []);
  const bridgeProxy = await Bridge.at((await ERC1967Proxy.deployed()).address);

  const bridgeFacade = await deployer.deploy(BridgeFacade);
  await deployer.deploy(ERC1967Proxy, bridgeFacade.address, []);
  const bridgeFacadeProxy = await BridgeFacade.at((await ERC1967Proxy.deployed()).address);

  logger.logTransaction(
    await bridgeProxy.__Bridge_init(
      config.accounts.validator,
      bundleExecutorImplementation.address,
      config.constants.chainName,
      bridgeFacadeProxy.address
    ),
    "Init Bridge"
  );
  logger.logTransaction(await bridgeFacadeProxy.__BridgeFacade_init(bridgeProxy.address), "Init BridgeFacade");
  logger.logContracts(
    ["Bridge proxy", bridgeProxy.address],
    ["Bridge implementation", bridge.address],
    ["BridgeFacade proxy", bridgeFacadeProxy.address],
    ["BridgeFacade implementation", bridgeFacade.address]
  );
};
