const { assert } = require("chai");
const { accounts, wei } = require("../../scripts/helpers/utils");
const truffleAssert = require("truffle-assertions");

const ERC1155HandlerMock = artifacts.require("ERC1155HandlerMock");
const ERC1155MB = artifacts.require("ERC1155MintableBurnable");

ERC1155MB.numberFormat = "BigNumber";
ERC1155HandlerMock.numberFormat = "BigNumber";

describe("ERC1155Handler", () => {
  const chainName = "ethereum";
  const baseAmount = "10";
  const baseId = "5000";
  const originHash = "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d";
  const salt = "0x0000000000000000000000000000000000000000000000000000000000000003";

  let OWNER;
  let handler;
  let token;

  before("setup", async () => {
    OWNER = await accounts(0);
  });

  beforeEach("setup", async () => {
    token = await ERC1155MB.new(OWNER, "");
    handler = await ERC1155HandlerMock.new();

    await token.mintTo(OWNER, baseId, baseAmount, "URI");
    await token.setApprovalForAll(handler.address, true);

    await token.transferOwnership(handler.address);
  });

  describe("access", () => {
    it("only this should call this method", async () => {
      await truffleAssert.reverts(
        handler.withdrawERC1155Bundle(token.address, 0, 0, "", { salt: salt, bundle: "0x" }, true),
        "Bundler: not this"
      );
    });
  });

  describe("depositERC1155", () => {
    it("should deposit token, isWrapped = true", async () => {
      let tx = await handler.depositERC1155(
        token.address,
        baseId,
        baseAmount,
        "receiver",
        { salt: salt, bundle: "0xFFAA" },
        "kovan",
        true
      );

      let realSalt = web3.utils.soliditySha3({
        value: web3.eth.abi.encodeParameters(["bytes32", "address"], [salt, OWNER]),
        type: "bytes",
      });

      assert.equal(await token.balanceOf(OWNER, baseId), "0");
      assert.equal(tx.receipt.logs[0].event, "DepositedERC1155");
      assert.equal(tx.receipt.logs[0].args.token, token.address);
      assert.equal(tx.receipt.logs[0].args.tokenId, baseId);
      assert.equal(tx.receipt.logs[0].args.amount, baseAmount);
      assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
      assert.equal(tx.receipt.logs[0].args.salt, realSalt);
      assert.equal(tx.receipt.logs[0].args.bundle, "0xffaa");
      assert.equal(tx.receipt.logs[0].args.network, "kovan");
      assert.isTrue(tx.receipt.logs[0].args.isWrapped);
    });

    it("should not burn tokens if they are not approved", async () => {
      await token.setApprovalForAll(handler.address, false);

      await truffleAssert.reverts(
        handler.depositERC1155(
          token.address,
          baseId,
          baseAmount,
          "receiver",
          { salt: salt, bundle: "0x" },
          "kovan",
          true
        ),
        "ERC1155MintableBurnable: not approved"
      );
    });

    it("should deposit token, isWrapped = false", async () => {
      let tx = await handler.depositERC1155(
        token.address,
        baseId,
        baseAmount,
        "receiver",
        { salt: salt, bundle: "0x000000" },
        "kovan",
        false
      );

      let realSalt = web3.utils.soliditySha3({
        value: web3.eth.abi.encodeParameters(["bytes32", "address"], [salt, OWNER]),
        type: "bytes",
      });

      assert.equal(await token.balanceOf(OWNER, baseId), "0");
      assert.equal(await token.balanceOf(handler.address, baseId), baseAmount);
      assert.equal(tx.receipt.logs[0].event, "DepositedERC1155");
      assert.equal(tx.receipt.logs[0].args.token, token.address);
      assert.equal(tx.receipt.logs[0].args.tokenId, baseId);
      assert.equal(tx.receipt.logs[0].args.amount, baseAmount);
      assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
      assert.equal(tx.receipt.logs[0].args.salt, realSalt);
      assert.equal(tx.receipt.logs[0].args.bundle, "0x000000");
      assert.equal(tx.receipt.logs[0].args.network, "kovan");
      assert.isFalse(tx.receipt.logs[0].args.isWrapped);
    });

    it("should revert when token address is 0", async () => {
      await truffleAssert.reverts(
        handler.depositERC1155(
          "0x0000000000000000000000000000000000000000",
          baseId,
          baseAmount,
          "receiver",
          { salt: salt, bundle: "0x" },
          "kovan",
          false
        ),
        "ERC1155Handler: zero token"
      );
    });

    it("should revert when depositing 0 tokens", async () => {
      await truffleAssert.reverts(
        handler.depositERC1155(
          token.address,
          baseId,
          wei("0"),
          "receiver",
          { salt: salt, bundle: "0x" },
          "kovan",
          false
        ),
        "ERC1155Handler: amount is zero"
      );
    });
  });

  describe("getERC1155MerkleLeaf", () => {
    it("should encode args", async () => {
      let bundle = "0x112233445566778899";

      let merkleLeaf0 = await handler.getERC1155MerkleLeaf(
        token.address,
        baseId,
        baseAmount,
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
          { value: baseAmount, type: "uint256" },
          { value: "URI1", type: "string" },
          { value: OWNER, type: "address" },
          { value: salt, type: "bytes32" },
          { value: bundle, type: "bytes" },
          { value: originHash, type: "bytes32" },
          { value: chainName, type: "string" },
          { value: handler.address, type: "address" }
        )
      );

      let merkleLeaf1 = await handler.getERC1155MerkleLeaf(
        token.address,
        baseId,
        baseAmount,
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
          { value: baseAmount, type: "uint256" },
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

  describe("withdrawERC1155", () => {
    it("should withdraw 100 tokens, wrapped = true", async () => {
      await handler.depositERC1155(
        token.address,
        baseId,
        baseAmount,
        "receiver",
        { salt: salt, bundle: "0x" },
        "kovan",
        true
      );

      assert.equal(await token.uri(baseId), "URI");

      await handler.withdrawERC1155(
        token.address,
        baseId,
        baseAmount,
        "URI1",
        OWNER,
        { salt: salt, bundle: "0x" },
        true
      );

      assert.equal(await token.balanceOf(OWNER, baseId), baseAmount);
      assert.equal(await token.balanceOf(handler.address, baseId), "0");
      assert.equal(await token.uri(baseId), "URI1");
    });

    it("should withdraw 52 tokens, wrapped = false", async () => {
      await handler.depositERC1155(
        token.address,
        baseId,
        baseAmount,
        "receiver",
        { salt: salt, bundle: "0x" },
        "kovan",
        false
      );
      await handler.withdrawERC1155(token.address, baseId, baseAmount, "", OWNER, { salt: salt, bundle: "0x" }, false);

      assert.equal(await token.balanceOf(OWNER, baseId), baseAmount);
      assert.equal(await token.balanceOf(handler.address, baseId), "0");
      assert.equal(await token.uri(baseId), "URI");
    });

    it("should revert when token address is 0", async () => {
      await truffleAssert.reverts(
        handler.withdrawERC1155(
          "0x0000000000000000000000000000000000000000",
          baseId,
          baseAmount,
          "",
          OWNER,
          { salt: salt, bundle: "0x" },
          true
        ),
        "ERC1155Handler: zero token"
      );
    });

    it("should revert when amount is 0", async () => {
      await truffleAssert.reverts(
        handler.withdrawERC1155(token.address, baseId, 0, "", OWNER, { salt: salt, bundle: "0x" }, true),
        "ERC1155Handler: amount is zero"
      );
    });

    it("should revert when receiver address is 0", async () => {
      await truffleAssert.reverts(
        handler.withdrawERC1155(
          token.address,
          baseId,
          baseAmount,
          "",
          "0x0000000000000000000000000000000000000000",
          { salt: salt, bundle: "0x" },
          true
        ),
        "ERC1155Handler: zero receiver"
      );
    });
  });
});
