const { assert } = require("chai");
const { accounts, wei, toBN } = require("../../scripts/helpers/utils");
const { constructTree, getProof, getRoot } = require("../../scripts/helpers/merkletree");
const ethSigUtil = require("@metamask/eth-sig-util");

const Bridge = artifacts.require("Bridge");
const ERC20MB = artifacts.require("ERC20MintableBurnable");
const ERC721MB = artifacts.require("ERC721MintableBurnable");
const ERC1155MB = artifacts.require("ERC1155MintableBurnable");

ERC1155MB.numberFormat = "BigNumber";
ERC721MB.numberFormat = "BigNumber";
ERC20MB.numberFormat = "BigNumber";
Bridge.numberFormat = "BigNumber";

const OWNER_PRIVATE_KEY = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80";

describe("Bridge", () => {
  const existingLeaves = [
    "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a5956471",
    "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a5956472",
    "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a5956473",
    "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a5956474",
    "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a5956475",
  ];

  const chainName = "ethereum";

  const baseBalance = wei("1000");
  const baseId = "5000";
  const originHash = "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d";

  let OWNER;
  let SECOND;

  let bridge;
  let erc20;
  let erc721;
  let erc1155;

  before("setup", async () => {
    OWNER = await accounts(0);
    SECOND = await accounts(1);
  });

  beforeEach("setup", async () => {
    bridge = await Bridge.new();
    await bridge.__Bridge_init(OWNER, chainName);

    erc20 = await ERC20MB.new("Mock", "MK", OWNER);
    await erc20.mintTo(OWNER, baseBalance);
    await erc20.approve(bridge.address, baseBalance);

    erc721 = await ERC721MB.new("Mock", "MK", OWNER);
    await erc721.mintTo(OWNER, baseId);
    await erc721.approve(bridge.address, baseId);

    erc1155 = await ERC1155MB.new("URI", OWNER);
    await erc1155.mintTo(OWNER, baseId, baseBalance);
    await erc1155.setApprovalForAll(bridge.address, true);

    await erc20.transferOwnership(bridge.address);
    await erc721.transferOwnership(bridge.address);
    await erc1155.transferOwnership(bridge.address);
  });

  describe("ERC20 flow", () => {
    it("should withdrawERC20", async () => {
      const isWrapped = true;
      const privateKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

      let tx = await bridge.depositERC20(erc20.address, baseBalance, OWNER, chainName, isWrapped);

      let leaf = await bridge.getERC20MerkleLeaf(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.amount,
        tx.receipt.logs[0].args.receiver,
        originHash,
        tx.receipt.logs[0].args.network,
        bridge.address
      );

      existingLeaves.push(leaf);

      const tree = constructTree(existingLeaves);
      const root = getRoot(tree);
      const proof = getProof(leaf, tree);

      const signature = ethSigUtil.personalSign({ privateKey: privateKey, data: root });

      await bridge.withdrawERC20(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.amount,
        tx.receipt.logs[0].args.receiver,
        originHash,
        proof,
        signature,
        isWrapped
      );

      assert.equal((await erc20.balanceOf(OWNER)).toFixed(), baseBalance);
      assert.equal(await erc20.balanceOf(bridge.address), "0");

      assert.isTrue(await bridge.usedHashes(originHash));
    });
  });

  describe("ERC721 flow", () => {
    it("should withdrawERC721", async () => {
      const isWrapped = true;
      const privateKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

      let tx = await bridge.depositERC721(erc721.address, baseId, OWNER, chainName, isWrapped);

      let leaf = await bridge.getERC721MerkleLeaf(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.tokenId,
        "1",
        tx.receipt.logs[0].args.receiver,
        originHash,
        tx.receipt.logs[0].args.network,
        bridge.address
      );

      existingLeaves.push(leaf);

      const tree = constructTree(existingLeaves);
      const root = getRoot(tree);
      const proof = getProof(leaf, tree);

      const signature = ethSigUtil.personalSign({ privateKey: privateKey, data: root });

      await bridge.withdrawERC721(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.tokenId,
        tx.receipt.logs[0].args.receiver,
        originHash,
        proof,
        signature,
        isWrapped
      );

      assert.equal(await erc721.ownerOf(baseId), OWNER);

      assert.isTrue(await bridge.usedHashes(originHash));
    });
  });

  describe("ERC1155 flow", () => {
    it("should withdrawERC1155", async () => {
      const isWrapped = true;
      const privateKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

      let tx = await bridge.depositERC1155(erc1155.address, baseId, baseBalance, OWNER, chainName, isWrapped);

      let leaf = await bridge.getERC721MerkleLeaf(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.tokenId,
        tx.receipt.logs[0].args.amount,
        tx.receipt.logs[0].args.receiver,
        originHash,
        tx.receipt.logs[0].args.network,
        bridge.address
      );

      existingLeaves.push(leaf);

      const tree = constructTree(existingLeaves);
      const root = getRoot(tree);
      const proof = getProof(leaf, tree);

      const signature = ethSigUtil.personalSign({ privateKey: privateKey, data: root });

      await bridge.withdrawERC1155(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.tokenId,
        tx.receipt.logs[0].args.amount,
        tx.receipt.logs[0].args.receiver,
        originHash,
        proof,
        signature,
        isWrapped
      );

      assert.equal((await erc1155.balanceOf(bridge.address, baseId)).toFixed(), "0");
      assert.equal((await erc1155.balanceOf(OWNER, baseId)).toFixed(), baseBalance);

      assert.isTrue(await bridge.usedHashes(originHash));
    });
  });

  describe("Native flow", () => {
    it("should withdrawNative", async () => {
      const privateKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

      let tx = await bridge.depositNative(OWNER, chainName, { value: baseBalance });

      let leaf = await bridge.getNativeMerkleLeaf(
        tx.receipt.logs[0].args.amount,
        tx.receipt.logs[0].args.receiver,
        originHash,
        tx.receipt.logs[0].args.network,
        bridge.address
      );

      existingLeaves.push(leaf);

      const tree = constructTree(existingLeaves);
      const root = getRoot(tree);
      const proof = getProof(leaf, tree);

      const signature = ethSigUtil.personalSign({ privateKey: privateKey, data: root });

      const prevBalance = await web3.eth.getBalance(OWNER);

      await bridge.withdrawNative(
        tx.receipt.logs[0].args.amount,
        tx.receipt.logs[0].args.receiver,
        originHash,
        proof,
        signature,
        { from: SECOND }
      );

      assert.equal(await web3.eth.getBalance(OWNER), toBN(prevBalance).plus(baseBalance).toFixed());
      assert.equal(await web3.eth.getBalance(bridge.address), "0");

      assert.isTrue(await bridge.usedHashes(originHash));
    });
  });
});
