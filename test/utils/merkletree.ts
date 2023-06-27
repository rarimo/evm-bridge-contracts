import { MerkleTree } from "merkletreejs";
import { ethers } from "hardhat";
import { SignHelper } from "@/test/utils/signature";
import {
  WithdrawERC20Parameters,
  WithdrawERC721Parameters,
  WithdrawSBTParameters,
  WithdrawERC1155Parameters,
  WithdrawNativeParameters,
} from "@/test/utils/types";

export class MerkleTreeHelper {
  public tree: MerkleTree;

  constructor(public signHelper: SignHelper, public chainName: string, public bridge: string) {
    const leaves = Array.from({ length: 10 }, () => ethers.utils.randomBytes(32));

    this.tree = new MerkleTree(
      leaves,
      (e: Buffer) => {
        const hash = ethers.utils.solidityKeccak256(["bytes"], [e]);

        return Buffer.from(hash.slice(2), "hex");
      },
      { sortPairs: true }
    );
  }

  public encodeLeaf(
    params: Partial<
      WithdrawERC20Parameters &
        WithdrawERC721Parameters &
        WithdrawSBTParameters &
        WithdrawERC1155Parameters &
        WithdrawNativeParameters
    >
  ) {
    const tokenData = (() => {
      let types = new Array<string>();
      let values = new Array<any>();

      if (params.token !== undefined) {
        types.push("address");
        values.push(params.token);
      }

      if (params.tokenId !== undefined) {
        types.push("uint256");
        values.push(params.tokenId);
      }

      if (params.tokenURI !== undefined) {
        types.push("string");
        values.push(params.tokenURI);
      }

      if (params.amount !== undefined) {
        types.push("uint256");
        values.push(params.amount);
      }

      return ethers.utils.solidityPack(types, values);
    })();

    const bundleData = (() => {
      if (params.bundle === undefined) {
        return "0x";
      }

      const salt = params.bundle.salt as string;
      const bundle = params.bundle.bundle as string;

      if (bundle.length > 0 && bundle !== "0x") {
        return ethers.utils.solidityPack(["bytes32", "bytes"], [salt, bundle]);
      }

      return "0x";
    })();

    return ethers.utils.solidityKeccak256(
      ["bytes", "bytes", "bytes32", "string", "address", "address"],
      [tokenData, bundleData, params.originHash, this.chainName, params.receiver, this.bridge]
    );
  }

  public addLeaf(leaf: string) {
    this.tree.addLeaf(Buffer.from(leaf.slice(2), "hex"));
  }

  public getPath(leaf: string): Array<string> {
    return this.tree.getProof(leaf).map((el) => "0x" + el.data.toString("hex"));
  }

  public getProof(leaf: string, addLeaf: boolean = true): string {
    if (addLeaf) {
      this.addLeaf(leaf);
    }

    const root = this.getRoot();
    const path = this.getPath(leaf);

    const signature = this.signHelper.sign(root);

    return ethers.utils.defaultAbiCoder.encode(["bytes32[]", "bytes"], [path, signature]);
  }

  public getRoot(): string {
    return "0x" + this.tree.getRoot().toString("hex");
  }
}
