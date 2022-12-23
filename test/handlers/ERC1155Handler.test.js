const { assert } = require("chai");
const { accounts, wei } = require("../../scripts/helpers/utils");
const truffleAssert = require("truffle-assertions");

const ERC1155HandlerMock = artifacts.require("ERC1155HandlerMock");
const ERC1155MB = artifacts.require("ERC1155MintableBurnable");

ERC1155MB.numberFormat = "BigNumber";
ERC1155HandlerMock.numberFormat = "BigNumber";

describe("ERC1155Handler", () => {
  const baseAmount = "10";
  const baseId = "5000";
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
      const tokenData = web3.eth.abi.encodeParameters(
        ["address", "uint256", "string", "uint256"],
        [token.address, 0, "", "0"]
      );

      await truffleAssert.reverts(
        handler.withdrawERC1155Bundle(tokenData, { salt: salt, bundle: "0x" }, true),
        "Bundler: not this"
      );
    });
  });

  describe("metadata", () => {
    it("should check token metadata", async () => {
      token = await ERC1155MB.new(OWNER, "URI");

      await token.mintTo(OWNER, baseId, baseAmount, "");

      assert.equal(await token.uri(baseId), "URI");
    });
  });

  describe("depositERC1155", () => {
    it("should deposit token, isWrapped = true", async () => {
      let tx = await handler.depositERC1155(
        token.address,
        baseId,
        baseAmount,
        { salt: salt, bundle: "0xFFAA" },
        "kovan",
        "receiver",
        true
      );

      const realSalt = web3.utils.soliditySha3({ value: salt, type: "bytes32" }, { value: OWNER, type: "address" });

      assert.equal(await token.balanceOf(OWNER, baseId), "0");

      assert.equal(tx.receipt.logs[0].event, "DepositedERC1155");
      assert.equal(tx.receipt.logs[0].args.token, token.address);
      assert.equal(tx.receipt.logs[0].args.tokenId, baseId);
      assert.equal(tx.receipt.logs[0].args.amount, baseAmount);
      assert.equal(tx.receipt.logs[0].args.salt, realSalt);
      assert.equal(tx.receipt.logs[0].args.bundle, "0xffaa");
      assert.equal(tx.receipt.logs[0].args.network, "kovan");
      assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
      assert.isTrue(tx.receipt.logs[0].args.isWrapped);
    });

    it("should not burn tokens if they are not approved", async () => {
      await token.setApprovalForAll(handler.address, false);

      await truffleAssert.reverts(
        handler.depositERC1155(
          token.address,
          baseId,
          baseAmount,
          { salt: salt, bundle: "0x" },
          "kovan",
          "receiver",
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
        { salt: salt, bundle: "0x000000" },
        "kovan",
        "receiver",
        false
      );

      const realSalt = web3.utils.soliditySha3({ value: salt, type: "bytes32" }, { value: OWNER, type: "address" });

      assert.equal(await token.balanceOf(OWNER, baseId), "0");
      assert.equal(await token.balanceOf(handler.address, baseId), baseAmount);

      assert.equal(tx.receipt.logs[0].event, "DepositedERC1155");
      assert.equal(tx.receipt.logs[0].args.token, token.address);
      assert.equal(tx.receipt.logs[0].args.tokenId, baseId);
      assert.equal(tx.receipt.logs[0].args.amount, baseAmount);
      assert.equal(tx.receipt.logs[0].args.salt, realSalt);
      assert.equal(tx.receipt.logs[0].args.bundle, "0x000000");
      assert.equal(tx.receipt.logs[0].args.network, "kovan");
      assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
      assert.isFalse(tx.receipt.logs[0].args.isWrapped);
    });

    it("should revert when token address is 0", async () => {
      await truffleAssert.reverts(
        handler.depositERC1155(
          "0x0000000000000000000000000000000000000000",
          baseId,
          baseAmount,
          { salt: salt, bundle: "0x" },
          "kovan",
          "receiver",
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
          { salt: salt, bundle: "0x" },
          "kovan",
          "receiver",
          false
        ),
        "ERC1155Handler: amount is zero"
      );
    });
  });

  describe("withdrawERC1155", () => {
    it("should withdraw 100 tokens, wrapped = true", async () => {
      const tokenData = web3.eth.abi.encodeParameters(
        ["address", "uint256", "string", "uint256"],
        [token.address, baseId, "URI1", baseAmount]
      );

      await handler.depositERC1155(
        token.address,
        baseId,
        baseAmount,
        { salt: salt, bundle: "0x" },
        "kovan",
        "receiver",
        true
      );

      assert.equal(await token.uri(baseId), "URI");

      await handler.withdrawERC1155(tokenData, OWNER, true);

      assert.equal(await token.balanceOf(OWNER, baseId), baseAmount);
      assert.equal(await token.balanceOf(handler.address, baseId), "0");
      assert.equal(await token.uri(baseId), "URI1");
    });

    it("should withdraw 52 tokens, wrapped = false", async () => {
      const tokenData = web3.eth.abi.encodeParameters(
        ["address", "uint256", "string", "uint256"],
        [token.address, baseId, "", baseAmount]
      );

      await handler.depositERC1155(
        token.address,
        baseId,
        baseAmount,
        { salt: salt, bundle: "0x" },
        "kovan",
        "receiver",
        false
      );
      await handler.withdrawERC1155(tokenData, OWNER, false);

      assert.equal(await token.balanceOf(OWNER, baseId), baseAmount);
      assert.equal(await token.balanceOf(handler.address, baseId), "0");
      assert.equal(await token.uri(baseId), "URI");
    });

    it("should withdraw 0 tokens, isWrapped = false", async () => {
      const tokenData = web3.eth.abi.encodeParameters(
        ["address", "uint256", "string", "uint256"],
        [token.address, baseId, "", "0"]
      );

      await truffleAssert.passes(handler.withdrawERC1155(tokenData, OWNER, false), "pass");
    });

    it("should withdraw 0 tokens, isWrapped = true", async () => {
      const tokenData = web3.eth.abi.encodeParameters(
        ["address", "uint256", "string", "uint256"],
        [token.address, baseId, "", "0"]
      );

      await truffleAssert.passes(handler.withdrawERC1155(tokenData, OWNER, true), "pass");
    });

    it("should revert when token address is 0", async () => {
      const tokenData = web3.eth.abi.encodeParameters(
        ["address", "uint256", "string", "uint256"],
        ["0x0000000000000000000000000000000000000000", baseId, "", baseAmount]
      );

      await truffleAssert.reverts(handler.withdrawERC1155(tokenData, OWNER, true), "ERC1155Handler: zero token");
    });

    it("should revert when receiver address is 0", async () => {
      const tokenData = web3.eth.abi.encodeParameters(
        ["address", "uint256", "string", "uint256"],
        [token.address, baseId, "", baseAmount]
      );

      await truffleAssert.reverts(
        handler.withdrawERC1155(tokenData, "0x0000000000000000000000000000000000000000", true),
        "ERC1155Handler: zero receiver"
      );
    });
  });
});
