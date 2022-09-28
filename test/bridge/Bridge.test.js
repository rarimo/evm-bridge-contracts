const { assert } = require("chai");
const { accounts, wei, toBN } = require("../../scripts/helpers/utils");
const {
  getBytesTransferERC20,
  getBytesSafeTransferERC721,
  getBytesSafeTransferERC1155,
  getBytesWithdrawNFT,
  getBytesWithdrawETH,
} = require("../bundler/bundler-utils");
const { constructTree, getProof, getRoot } = require("../../scripts/helpers/merkletree");
const ethSigUtil = require("@metamask/eth-sig-util");
const truffleAssert = require("truffle-assertions");

const Bridge = artifacts.require("Bridge");
const BundleImpl = artifacts.require("BundleExecutorImplementation");
const ERC20MB = artifacts.require("ERC20MintableBurnable");
const ERC721MB = artifacts.require("ERC721MintableBurnable");
const ERC1155MB = artifacts.require("ERC1155MintableBurnable");
const BundleReceiver = artifacts.require("BundleReceiverMock");

ERC1155MB.numberFormat = "BigNumber";
ERC721MB.numberFormat = "BigNumber";
ERC20MB.numberFormat = "BigNumber";
Bridge.numberFormat = "BigNumber";
BundleImpl.numberFormat = "BigNumber";
BundleReceiver.numberFormat = "BigNumber";

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
  const salt = "0x0000000000000000000000000000000000000000000000000000000000000001";

  let OWNER;
  let SECOND;

  let bridge;
  let bundleImpl;
  let bundleReceiver;
  let erc20;
  let erc721;
  let erc1155;

  before("setup", async () => {
    OWNER = await accounts(0);
    SECOND = await accounts(1);
  });

  beforeEach("setup", async () => {
    bundleImpl = await BundleImpl.new();
    bundleReceiver = await BundleReceiver.new();

    bridge = await Bridge.new();
    await bridge.__Bridge_init(OWNER, bundleImpl.address, chainName);

    erc20 = await ERC20MB.new("Mock", "MK", OWNER);
    await erc20.mintTo(OWNER, baseBalance);
    await erc20.approve(bridge.address, baseBalance);

    erc721 = await ERC721MB.new("Mock", "MK", OWNER, "");
    await erc721.mintTo(OWNER, baseId, "URI");
    await erc721.approve(bridge.address, baseId);

    erc1155 = await ERC1155MB.new(OWNER, "");
    await erc1155.mintTo(OWNER, baseId, baseBalance, "URI");
    await erc1155.setApprovalForAll(bridge.address, true);

    await erc20.transferOwnership(bridge.address);
    await erc721.transferOwnership(bridge.address);
    await erc1155.transferOwnership(bridge.address);
  });

  describe("ERC20 flow", () => {
    it("should withdrawERC20", async () => {
      const isWrapped = true;
      const privateKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

      let tx = await bridge.depositERC20(
        erc20.address,
        baseBalance,
        OWNER,
        { salt: salt, bundle: "0x" },
        chainName,
        isWrapped
      );

      let leaf = await bridge.getERC20MerkleLeaf(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.amount,
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle == null ? "0x" : tx.receipt.logs[0].args.bundle,
        },
        originHash,
        tx.receipt.logs[0].args.network
      );

      existingLeaves.push(leaf);

      const tree = constructTree(existingLeaves);
      const root = getRoot(tree);
      const path = getProof(leaf, tree);

      const signature = ethSigUtil.personalSign({ privateKey: privateKey, data: root });

      const proof = web3.eth.abi.encodeParameters(["bytes32[]", "bytes"], [path, signature]);

      await bridge.withdrawERC20(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.amount,
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle == null ? "0x" : tx.receipt.logs[0].args.bundle,
        },
        originHash,
        proof,
        isWrapped
      );

      assert.equal((await erc20.balanceOf(OWNER)).toFixed(), baseBalance);
      assert.equal(await erc20.balanceOf(bridge.address), "0");

      assert.isTrue(await bridge.usedHashes(originHash));
    });

    it("should withdraw ERC20 through bundling", async () => {
      const isWrapped = true;
      const privateKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");
      const bundle = web3.eth.abi.encodeParameters(
        ["address[]", "uint256[]", "bytes[]"],
        [[erc20.address], [0], [getBytesTransferERC20(SECOND, wei("100"))]]
      );
      const realSalt = web3.utils.soliditySha3({
        value: web3.eth.abi.encodeParameters(["bytes32", "address"], [salt, OWNER]),
        type: "bytes",
      });

      let tx = await bridge.depositERC20(
        erc20.address,
        baseBalance,
        OWNER,
        { salt: salt, bundle: bundle },
        chainName,
        isWrapped
      );

      let leaf = await bridge.getERC20MerkleLeaf(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.amount,
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle,
        },
        originHash,
        tx.receipt.logs[0].args.network
      );

      existingLeaves.push(leaf);

      const tree = constructTree(existingLeaves);
      const root = getRoot(tree);
      const path = getProof(leaf, tree);

      const signature = ethSigUtil.personalSign({ privateKey: privateKey, data: root });

      const proof = web3.eth.abi.encodeParameters(["bytes32[]", "bytes"], [path, signature]);

      const bundleProxy = await bridge.determineProxyAddress(realSalt);

      await bridge.withdrawERC20(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.amount,
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle,
        },
        originHash,
        proof,
        isWrapped
      );

      assert.equal((await erc20.balanceOf(SECOND)).toFixed(), wei("100"));
      assert.equal(await erc20.balanceOf(bridge.address), "0");
      assert.equal((await erc20.balanceOf(bundleProxy)).toFixed(), wei("900"));
    });

    it("should fallback to a normal withdrawal if bundling reverts", async () => {
      const isWrapped = true;
      const privateKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");
      const bundle = web3.eth.abi.encodeParameters(
        ["address[]", "uint256[]", "bytes[]"],
        [[erc20.address], [0], [getBytesTransferERC20(SECOND, wei("10000"))]]
      );

      let tx = await bridge.depositERC20(
        erc20.address,
        baseBalance,
        OWNER,
        { salt: salt, bundle: bundle },
        chainName,
        isWrapped
      );

      let leaf = await bridge.getERC20MerkleLeaf(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.amount,
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle,
        },
        originHash,
        tx.receipt.logs[0].args.network
      );

      existingLeaves.push(leaf);

      const tree = constructTree(existingLeaves);
      const root = getRoot(tree);
      const path = getProof(leaf, tree);

      const signature = ethSigUtil.personalSign({ privateKey: privateKey, data: root });

      const proof = web3.eth.abi.encodeParameters(["bytes32[]", "bytes"], [path, signature]);

      const bundleProxy = await bridge.determineProxyAddress(tx.receipt.logs[0].args.salt);

      await bridge.withdrawERC20(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.amount,
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle,
        },
        originHash,
        proof,
        isWrapped
      );

      assert.equal((await erc20.balanceOf(OWNER)).toFixed(), baseBalance);
      assert.equal(await erc20.balanceOf(bridge.address), "0");
      assert.equal(await erc20.balanceOf(bundleProxy), "0");
    });
  });

  describe("ERC721 flow", () => {
    it("should withdrawERC721", async () => {
      const isWrapped = true;
      const privateKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

      let tx = await bridge.depositERC721(
        erc721.address,
        baseId,
        OWNER,
        { salt: salt, bundle: "0x" },
        chainName,
        isWrapped
      );

      let leaf = await bridge.getERC721MerkleLeaf(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.tokenId,
        "URI1",
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle == null ? "0x" : tx.receipt.logs[0].args.bundle,
        },
        originHash,
        tx.receipt.logs[0].args.network
      );

      existingLeaves.push(leaf);

      const tree = constructTree(existingLeaves);
      const root = getRoot(tree);
      const path = getProof(leaf, tree);

      const signature = ethSigUtil.personalSign({ privateKey: privateKey, data: root });

      const proof = web3.eth.abi.encodeParameters(["bytes32[]", "bytes"], [path, signature]);

      await bridge.withdrawERC721(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.tokenId,
        "URI1",
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle == null ? "0x" : tx.receipt.logs[0].args.bundle,
        },
        originHash,
        proof,
        isWrapped
      );

      assert.equal(await erc721.ownerOf(baseId), OWNER);
      assert.equal(await erc721.tokenURI(baseId), "URI1");

      assert.isTrue(await bridge.usedHashes(originHash));
    });

    it("should transfer NFT through bundling", async () => {
      const isWrapped = false;
      const privateKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

      const realSalt = web3.utils.soliditySha3({
        value: web3.eth.abi.encodeParameters(["bytes32", "address"], [salt, OWNER]),
        type: "bytes",
      });

      const bundleProxy = await bridge.determineProxyAddress(realSalt);

      const bundle = web3.eth.abi.encodeParameters(
        ["address[]", "uint256[]", "bytes[]"],
        [
          [erc721.address, bundleReceiver.address, erc721.address],
          [0, 0, 0],
          [
            getBytesSafeTransferERC721(bundleProxy, bundleReceiver.address, baseId),
            getBytesWithdrawNFT(erc721.address, baseId),
            getBytesSafeTransferERC721(bundleProxy, SECOND, baseId),
          ],
        ]
      );

      let tx = await bridge.depositERC721(
        erc721.address,
        baseId,
        OWNER,
        { salt: salt, bundle: bundle },
        chainName,
        isWrapped
      );

      let leaf = await bridge.getERC721MerkleLeaf(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.tokenId,
        "URI1",
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle,
        },
        originHash,
        tx.receipt.logs[0].args.network
      );

      existingLeaves.push(leaf);

      const tree = constructTree(existingLeaves);
      const root = getRoot(tree);
      const path = getProof(leaf, tree);

      const signature = ethSigUtil.personalSign({ privateKey: privateKey, data: root });

      const proof = web3.eth.abi.encodeParameters(["bytes32[]", "bytes"], [path, signature]);

      await bridge.withdrawERC721(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.tokenId,
        "URI1",
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle,
        },
        originHash,
        proof,
        isWrapped
      );

      assert.equal(await erc721.ownerOf(baseId), SECOND);
      assert.equal(await erc721.tokenURI(baseId), "URI");
    });
  });

  describe("ERC1155 flow", () => {
    it("should withdrawERC1155", async () => {
      const isWrapped = true;
      const privateKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

      let tx = await bridge.depositERC1155(
        erc1155.address,
        baseId,
        baseBalance,
        OWNER,
        { salt: salt, bundle: "0x" },
        chainName,
        isWrapped
      );

      let leaf = await bridge.getERC1155MerkleLeaf(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.tokenId,
        tx.receipt.logs[0].args.amount,
        "URI2",
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle == null ? "0x" : tx.receipt.logs[0].args.bundle,
        },
        originHash,
        tx.receipt.logs[0].args.network
      );

      existingLeaves.push(leaf);

      const tree = constructTree(existingLeaves);
      const root = getRoot(tree);
      const path = getProof(leaf, tree);

      const signature = ethSigUtil.personalSign({ privateKey: privateKey, data: root });

      const proof = web3.eth.abi.encodeParameters(["bytes32[]", "bytes"], [path, signature]);

      await bridge.withdrawERC1155(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.tokenId,
        tx.receipt.logs[0].args.amount,
        "URI2",
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle == null ? "0x" : tx.receipt.logs[0].args.bundle,
        },
        originHash,
        proof,
        isWrapped
      );

      assert.equal((await erc1155.balanceOf(bridge.address, baseId)).toFixed(), "0");
      assert.equal((await erc1155.balanceOf(OWNER, baseId)).toFixed(), baseBalance);
      assert.equal(await erc1155.uri(baseId), "URI2");

      assert.isTrue(await bridge.usedHashes(originHash));
    });

    it("should withdraw ERC1155 through bundling", async () => {
      const isWrapped = false;
      const privateKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

      const realSalt = web3.utils.soliditySha3({
        value: web3.eth.abi.encodeParameters(["bytes32", "address"], [salt, OWNER]),
        type: "bytes",
      });

      const bundleProxy = await bridge.determineProxyAddress(realSalt);

      const bundle = web3.eth.abi.encodeParameters(
        ["address[]", "uint256[]", "bytes[]"],
        [[erc1155.address], [0], [getBytesSafeTransferERC1155(bundleProxy, SECOND, baseId, wei("100"), "0x")]]
      );

      let tx = await bridge.depositERC1155(
        erc1155.address,
        baseId,
        baseBalance,
        OWNER,
        { salt: salt, bundle: bundle },
        chainName,
        isWrapped
      );

      let leaf = await bridge.getERC1155MerkleLeaf(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.tokenId,
        tx.receipt.logs[0].args.amount,
        "URI2",
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle,
        },
        originHash,
        tx.receipt.logs[0].args.network
      );

      existingLeaves.push(leaf);

      const tree = constructTree(existingLeaves);
      const root = getRoot(tree);
      const path = getProof(leaf, tree);

      const signature = ethSigUtil.personalSign({ privateKey: privateKey, data: root });

      const proof = web3.eth.abi.encodeParameters(["bytes32[]", "bytes"], [path, signature]);

      await bridge.withdrawERC1155(
        tx.receipt.logs[0].args.token,
        tx.receipt.logs[0].args.tokenId,
        tx.receipt.logs[0].args.amount,
        "URI2",
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle,
        },
        originHash,
        proof,
        isWrapped
      );

      assert.equal((await erc1155.balanceOf(bridge.address, baseId)).toFixed(), "0");
      assert.equal((await erc1155.balanceOf(bundleProxy, baseId)).toFixed(), wei("900"));
      assert.equal((await erc1155.balanceOf(SECOND, baseId)).toFixed(), wei("100"));

      assert.equal(await erc1155.uri(baseId), "URI");
    });
  });

  describe("Native flow", () => {
    it("should withdrawNative", async () => {
      const privateKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

      let tx = await bridge.depositNative(OWNER, { salt: salt, bundle: "0x" }, chainName, { value: baseBalance });

      let leaf = await bridge.getNativeMerkleLeaf(
        tx.receipt.logs[0].args.amount,
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle == null ? "0x" : tx.receipt.logs[0].args.bundle,
        },
        originHash,
        tx.receipt.logs[0].args.network
      );

      existingLeaves.push(leaf);

      const tree = constructTree(existingLeaves);
      const root = getRoot(tree);
      const path = getProof(leaf, tree);

      const signature = ethSigUtil.personalSign({ privateKey: privateKey, data: root });

      const proof = web3.eth.abi.encodeParameters(["bytes32[]", "bytes"], [path, signature]);

      const prevBalance = await web3.eth.getBalance(OWNER);

      await bridge.withdrawNative(
        tx.receipt.logs[0].args.amount,
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle == null ? "0x" : tx.receipt.logs[0].args.bundle,
        },
        originHash,
        proof,
        { from: SECOND }
      );

      assert.equal(await web3.eth.getBalance(OWNER), toBN(prevBalance).plus(baseBalance).toFixed());
      assert.equal(await web3.eth.getBalance(bridge.address), "0");

      assert.isTrue(await bridge.usedHashes(originHash));
    });

    it("should withdraw native through bundling", async () => {
      const privateKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

      const bundle = web3.eth.abi.encodeParameters(
        ["address[]", "uint256[]", "bytes[]"],
        [
          [bundleReceiver.address, bundleReceiver.address, SECOND],
          [wei("10"), 0, wei("100")],
          ["0x", getBytesWithdrawETH(wei("1")), "0x"],
        ]
      );

      let tx = await bridge.depositNative(OWNER, { salt: salt, bundle: bundle }, chainName, { value: baseBalance });

      let leaf = await bridge.getNativeMerkleLeaf(
        tx.receipt.logs[0].args.amount,
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle,
        },
        originHash,
        tx.receipt.logs[0].args.network
      );

      existingLeaves.push(leaf);

      const tree = constructTree(existingLeaves);
      const root = getRoot(tree);
      const path = getProof(leaf, tree);

      const signature = ethSigUtil.personalSign({ privateKey: privateKey, data: root });

      const proof = web3.eth.abi.encodeParameters(["bytes32[]", "bytes"], [path, signature]);

      const bundleProxy = await bridge.determineProxyAddress(tx.receipt.logs[0].args.salt);

      const prevBalance = await web3.eth.getBalance(SECOND);

      await bridge.withdrawNative(
        tx.receipt.logs[0].args.amount,
        tx.receipt.logs[0].args.receiver,
        {
          salt: tx.receipt.logs[0].args.salt,
          bundle: tx.receipt.logs[0].args.bundle,
        },
        originHash,
        proof
      );

      assert.equal(await web3.eth.getBalance(SECOND), toBN(prevBalance).plus(wei("100")).toFixed());
      assert.equal(await web3.eth.getBalance(bridge.address), wei("891"));
      assert.equal(await web3.eth.getBalance(bundleProxy), "0");
      assert.equal(await web3.eth.getBalance(bundleReceiver.address), wei("9"));
    });
  });

  describe("changeSigner", () => {
    it("should correctly change signer", async () => {
      const ownerKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

      assert.equal(await bridge.signer(), OWNER);
      assert.equal(await bridge.nonce(), "0");

      const hashToSign = web3.utils.soliditySha3(
        { value: SECOND, type: "address" },
        { value: chainName, type: "string" },
        { value: "0", type: "uint256" },
        { value: bridge.address, type: "address" }
      );

      const signature = ethSigUtil.personalSign({ privateKey: ownerKey, data: hashToSign });

      await bridge.changeSigner(SECOND, signature);

      assert.equal(await bridge.signer(), SECOND);
      assert.equal(await bridge.nonce(), "1");
    });

    it("should not change signer if bad signature is passed", async () => {
      const ownerKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

      const hashToSign = web3.utils.soliditySha3(
        { value: SECOND, type: "address" },
        { value: chainName, type: "string" },
        { value: "0", type: "uint256" },
        { value: bridge.address, type: "address" }
      );
      const signature = ethSigUtil.personalSign({ privateKey: ownerKey, data: hashToSign });

      await bridge.changeSigner(SECOND, signature);

      await truffleAssert.reverts(bridge.changeSigner(SECOND, signature), "Signers: invalid signature");
    });
  });

  describe("changeBundleExecutorImplementation", () => {
    it("should correctly change executor implementation", async () => {
      const ownerKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

      assert.equal(await bridge.bundleExecutorImplementation(), bundleImpl.address);
      assert.equal(await bridge.nonce(), "0");

      const hashToSign = web3.utils.soliditySha3(
        { value: SECOND, type: "address" },
        { value: chainName, type: "string" },
        { value: "0", type: "uint256" },
        { value: bridge.address, type: "address" }
      );

      const signature = ethSigUtil.personalSign({ privateKey: ownerKey, data: hashToSign });

      await bridge.changeBundleExecutorImplementation(SECOND, signature);

      assert.equal(await bridge.bundleExecutorImplementation(), SECOND);
      assert.equal(await bridge.nonce(), "1");
    });
  });
});
