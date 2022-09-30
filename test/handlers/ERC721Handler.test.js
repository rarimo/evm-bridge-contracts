const { assert } = require("chai");
const { accounts } = require("../../scripts/helpers/utils");
const truffleAssert = require("truffle-assertions");

const ERC721HandlerMock = artifacts.require("ERC721HandlerMock");
const ERC721MB = artifacts.require("ERC721MintableBurnable");

ERC721MB.numberFormat = "BigNumber";
ERC721HandlerMock.numberFormat = "BigNumber";

describe("ERC721Handler", () => {
  const chainName = "ethereum";
  const baseId = "5000";
  const originHash = "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d";
  const salt = "0x0000000000000000000000000000000000000000000000000000000000000002";

  let OWNER;
  let SECOND;
  let handler;
  let token;

  before("setup", async () => {
    OWNER = await accounts(0);
    SECOND = await accounts(1);
  });

  beforeEach("setup", async () => {
    token = await ERC721MB.new("Mock", "MK", OWNER, "");
    handler = await ERC721HandlerMock.new();

    await token.mintTo(OWNER, baseId, "URI");
    await token.approve(handler.address, baseId);

    await token.transferOwnership(handler.address);
  });

  describe("access", () => {
    it("only this should call this method", async () => {
      await truffleAssert.reverts(
        handler.withdrawERC721Bundle(token.address, 0, "", { salt: salt, bundle: "0x" }, true),
        "Bundler: not this"
      );
    });
  });

  describe("depositERC721", () => {
    it("should deposit token, isWrapped = true", async () => {
      let tx = await handler.depositERC721(
        token.address,
        baseId,
        "receiver",
        { salt: salt, bundle: "0x00" },
        "kovan",
        true
      );

      let realSalt = web3.utils.soliditySha3({
        value: web3.eth.abi.encodeParameters(["bytes32", "address"], [salt, OWNER]),
        type: "bytes",
      });

      assert.equal(tx.receipt.logs[0].event, "DepositedERC721");
      assert.equal(tx.receipt.logs[0].args.token, token.address);
      assert.equal(tx.receipt.logs[0].args.tokenId, baseId);
      assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
      assert.equal(tx.receipt.logs[0].args.salt, realSalt);
      assert.equal(tx.receipt.logs[0].args.bundle, "0x00");
      assert.equal(tx.receipt.logs[0].args.network, "kovan");
      assert.isTrue(tx.receipt.logs[0].args.isWrapped);

      await truffleAssert.reverts(token.ownerOf(baseId), "ERC721: owner query for nonexistent token");
    });

    it("should not burn token if it is not approved", async () => {
      await token.approve(token.address, baseId);

      await truffleAssert.reverts(
        handler.depositERC721(token.address, baseId, "receiver", { salt: salt, bundle: "0x" }, "kovan", true),
        "ERC721MintableBurnable: not approved"
      );
    });

    it("should not burn token if it is approved but not owned", async () => {
      await truffleAssert.reverts(
        handler.depositERC721(token.address, baseId, "receiver", { salt: salt, bundle: "0x" }, "kovan", true, {
          from: SECOND,
        }),
        "ERC721MintableBurnable: not approved"
      );
    });

    it("should deposit token, isWrapped = false", async () => {
      let tx = await handler.depositERC721(
        token.address,
        baseId,
        "receiver",
        { salt: salt, bundle: "0x0123" },
        "kovan",
        false
      );

      let realSalt = web3.utils.soliditySha3({
        value: web3.eth.abi.encodeParameters(["bytes32", "address"], [salt, OWNER]),
        type: "bytes",
      });

      assert.equal(await token.ownerOf(baseId), handler.address);
      assert.equal(tx.receipt.logs[0].event, "DepositedERC721");
      assert.equal(tx.receipt.logs[0].args.token, token.address);
      assert.equal(tx.receipt.logs[0].args.tokenId, baseId);
      assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
      assert.equal(tx.receipt.logs[0].args.salt, realSalt);
      assert.equal(tx.receipt.logs[0].args.bundle, "0x0123");
      assert.equal(tx.receipt.logs[0].args.network, "kovan");
      assert.isFalse(tx.receipt.logs[0].args.isWrapped);
    });

    it("should revert when token address is 0", async () => {
      await truffleAssert.reverts(
        handler.depositERC721(
          "0x0000000000000000000000000000000000000000",
          baseId,
          "receiver",
          { salt: salt, bundle: "0x00" },
          "kovan",
          false
        ),
        "ERC721Handler: zero token"
      );
    });
  });

  describe("getERC721MerkleLeaf", () => {
    it("should encode args", async () => {
      let bundle = "0xFFAAFFAA";

      let merkleLeaf0 = await handler.getERC721MerkleLeaf(
        token.address,
        baseId,
        "URI1",
        OWNER,
        { salt: salt, bundle: bundle },
        originHash,
        chainName
      );

      assert.equal(
        merkleLeaf0,
        web3.utils.soliditySha3(
          { value: token.address, type: "address" },
          { value: baseId, type: "uint256" },
          { value: "1", type: "uint256" },
          { value: "URI1", type: "string" },
          { value: OWNER, type: "address" },
          { value: salt, type: "bytes32" },
          { value: bundle, type: "bytes" },
          { value: originHash, type: "bytes32" },
          { value: chainName, type: "string" },
          { value: handler.address, type: "address" }
        )
      );

      let merkleLeaf1 = await handler.getERC721MerkleLeaf(
        token.address,
        baseId,
        "URI2",
        OWNER,
        { salt: salt, bundle: bundle },
        originHash,
        "BSC"
      );

      assert.equal(
        merkleLeaf1,
        web3.utils.soliditySha3(
          { value: token.address, type: "address" },
          { value: baseId, type: "uint256" },
          { value: "1", type: "uint256" },
          { value: "URI2", type: "string" },
          { value: OWNER, type: "address" },
          { value: salt, type: "bytes32" },
          { value: bundle, type: "bytes" },
          { value: originHash, type: "bytes32" },
          { value: "BSC", type: "string" },
          { value: handler.address, type: "address" }
        )
      );

      assert.notEqual(merkleLeaf0, merkleLeaf1);
    });
  });

  describe("withdrawERC721", async () => {
    it("should withdraw token, wrapped = true", async () => {
      await handler.depositERC721(token.address, baseId, "receiver", { salt: salt, bundle: "0x" }, "kovan", true);
      await handler.withdrawERC721(token.address, baseId, "URI1", OWNER, { salt: salt, bundle: "0x" }, true);

      assert.equal(await token.ownerOf(baseId), OWNER);
      assert.equal(await token.tokenURI(baseId), "URI1");
    });

    it("should withdraw token, wrapped = false", async () => {
      await handler.depositERC721(token.address, baseId, "receiver", { salt: salt, bundle: "0x" }, "kovan", false);
      await handler.withdrawERC721(token.address, baseId, "URI1", OWNER, { salt: salt, bundle: "0x" }, false);

      assert.equal(await token.ownerOf(baseId), OWNER);
      assert.equal(await token.tokenURI(baseId), "URI");
    });

    it("should revert when token address is 0", async () => {
      await truffleAssert.reverts(
        handler.withdrawERC721(
          "0x0000000000000000000000000000000000000000",
          baseId,
          "",
          OWNER,
          { salt: salt, bundle: "0x" },
          false
        ),
        "ERC721Handler: zero token"
      );
    });

    it("should revert when receiver address is 0", async () => {
      await truffleAssert.reverts(
        handler.withdrawERC721(
          token.address,
          baseId,
          "",
          "0x0000000000000000000000000000000000000000",
          { salt: salt, bundle: "0x" },
          false
        ),
        "ERC721Handler: zero receiver"
      );
    });
  });
});
