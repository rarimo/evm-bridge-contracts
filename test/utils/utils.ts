import { ethers } from "hardhat";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";

export function encodeSalt(salt: string, txOrigin: SignerWithAddress): string {
  return ethers.utils.solidityKeccak256(["bytes32", "address"], [salt, txOrigin.address]);
}
