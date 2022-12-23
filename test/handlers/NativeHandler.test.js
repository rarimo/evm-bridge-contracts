const { assert } = require("chai");
const { accounts, wei } = require("../../scripts/helpers/utils");
const truffleAssert = require("truffle-assertions");

const NativeHandlerMock = artifacts.require("NativeHandlerMock");

NativeHandlerMock.numberFormat = "BigNumber";

describe("NativeHandler", () => {
  const baseAmount = wei("10");
  const salt = "0x0000000000000000000000000000000000000000000000000000000000000004";

  let OWNER;
  let handler;

  before("setup", async () => {
    OWNER = await accounts(0);
  });

  beforeEach("setup", async () => {
    handler = await NativeHandlerMock.new();
  });

  describe("access", () => {
    it("only this should call this method", async () => {
      const tokenData = web3.eth.abi.encodeParameters(["uint256"], [0]);

      await truffleAssert.reverts(
        handler.withdrawNativeBundle(tokenData, { salt: salt, bundle: "0x" }, false),
        "Bundler: not this"
      );
    });
  });

  describe("depositNative", () => {
    it("should deposit native", async () => {
      let tx = await handler.depositNative({ salt: salt, bundle: "0x000001" }, "kovan", "receiver", {
        value: baseAmount,
      });

      const realSalt = web3.utils.soliditySha3({ value: salt, type: "bytes32" }, { value: OWNER, type: "address" });

      assert.equal(await web3.eth.getBalance(handler.address), baseAmount);

      assert.equal(tx.receipt.logs[0].event, "DepositedNative");
      assert.equal(tx.receipt.logs[0].args.amount, baseAmount);
      assert.equal(tx.receipt.logs[0].args.salt, realSalt);
      assert.equal(tx.receipt.logs[0].args.bundle, "0x000001");
      assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
      assert.equal(tx.receipt.logs[0].args.network, "kovan");
    });

    it("should revert when depositing 0 tokens", async () => {
      await truffleAssert.reverts(
        handler.depositNative({ salt: salt, bundle: "0x" }, "kovan", "receiver"),
        "NativeHandler: zero value"
      );
    });
  });

  describe("withdrawNative", () => {
    it("should withdraw native", async () => {
      const tokenData = web3.eth.abi.encodeParameters(["uint256"], [baseAmount]);

      await handler.depositNative({ salt: salt, bundle: "0x" }, "kovan", "receiver", { value: baseAmount });
      await handler.withdrawNative(tokenData, OWNER);

      assert.equal(await web3.eth.getBalance(handler.address), 0);
    });

    it("should withdraw 0 tokens", async () => {
      const tokenData = web3.eth.abi.encodeParameters(["uint256"], [0]);

      await truffleAssert.passes(handler.withdrawNative(tokenData, OWNER), "pass");
    });

    it("should withdraw 0 tokens to the contract with receive function", async () => {
      const tokenData = web3.eth.abi.encodeParameters(["uint256"], [0]);

      await truffleAssert.passes(handler.withdrawNative(tokenData, handler.address), "pass");
    });

    it("should revert when receiver address is 0", async () => {
      const tokenData = web3.eth.abi.encodeParameters(["uint256"], [baseAmount]);

      await truffleAssert.reverts(
        handler.withdrawNative(tokenData, "0x0000000000000000000000000000000000000000"),
        "NativeHandler: receiver is zero"
      );
    });

    it("should revert when amount more than balance", async () => {
      const tokenData = web3.eth.abi.encodeParameters(["uint256"], [wei("1000000")]);

      await truffleAssert.reverts(handler.withdrawNative(tokenData, OWNER), "NativeHandler: failed to send eth");
    });
  });
});
