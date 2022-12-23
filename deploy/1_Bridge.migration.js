const BundleExecutorImpl = artifacts.require("BundleExecutorImplementation");
const Bridge = artifacts.require("Bridge");
const ERC1967Proxy = artifacts.require("ERC1967Proxy");

const VALIDATOR = process.env.VALIDATOR;
const CHAIN_NAME = process.env.CHAIN_NAME;

module.exports = async (deployer, logger) => {
  const bundleExecutorImpl = await deployer.deploy(BundleExecutorImpl);

  const bridge = await deployer.deploy(Bridge);
  const proxy = await deployer.deploy(ERC1967Proxy, bridge.address, []);

  const bridgeProxy = await Bridge.at(proxy.address);

  logger.logTransaction(
    await bridgeProxy.__Bridge_init(VALIDATOR, bundleExecutorImpl.address, CHAIN_NAME),
    "Init Bridge"
  );
};
