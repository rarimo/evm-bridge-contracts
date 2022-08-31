const { logTransaction } = require("../runners/logger/logger.js");

const Bridge = artifacts.require("Bridge");
const ERC1967Proxy = artifacts.require("ERC1967Proxy");

// TODO change parameters
const VALIDATOR = "0x53638975BC11de3029E46DF193d64879EAeA94eB";
const CHAIN_NAME = "ethereum";

module.exports = async (deployer) => {
  const bridge = await deployer.deploy(Bridge);
  const proxy = await deployer.deploy(ERC1967Proxy, bridge.address, []);

  const bridgeProxy = await Bridge.at(proxy.address);

  logTransaction(await bridgeProxy.__Bridge_init(VALIDATOR, CHAIN_NAME), "Init Bridge");
};
