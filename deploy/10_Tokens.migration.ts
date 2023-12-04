import { Deployer, Logger } from "@solarity/hardhat-migrate";
import { artifacts } from "hardhat";

const ERC1967Proxy = artifacts.require("ERC1967Proxy");
const ERC20MintableBurnable = artifacts.require("ERC20MintableBurnable");
const ERC721MintableBurnable = artifacts.require("ERC721MintableBurnable");
const ERC1155MintableBurnable = artifacts.require("ERC1155MintableBurnable");

module.exports = async (deployer: Deployer, logger: Logger) => {
  const bridgeAddress = "";

  let contracts = [];

  contracts.push([
    "WFUJI",
    (await deployer.deploy(ERC20MintableBurnable, "WFUJI", "WFUJI", 18, bridgeAddress)).address,
  ]);

  contracts.push([
    "WNEAR",
    (await deployer.deploy(ERC20MintableBurnable, "WNEAR", "WNEAR", 24, bridgeAddress)).address,
  ]);

  contracts.push(["WSOL", (await deployer.deploy(ERC20MintableBurnable, "WSOL", "WSOL", 18, bridgeAddress)).address]);

  contracts.push(["USDC", (await deployer.deploy(ERC20MintableBurnable, "USDC", "USDC", 6, bridgeAddress)).address]);

  contracts.push([
    "ERC20Native",
    (await deployer.deploy(ERC20MintableBurnable, "ERC20Native", "ERC20Native", 18, bridgeAddress)).address,
  ]);

  contracts.push([
    "ERC721Native",
    (await deployer.deploy(ERC721MintableBurnable, "ERC721Native", "ERC721Native", bridgeAddress, "")).address,
  ]);

  contracts.push(["ERC1155Native", (await deployer.deploy(ERC1155MintableBurnable, bridgeAddress, "")).address]);

  logger.logContracts(...contracts);
};
