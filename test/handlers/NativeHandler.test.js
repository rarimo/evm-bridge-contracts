const { assert } = require("chai");
const { accounts, wei } = require("../../scripts/helpers/utils");
const truffleAssert = require("truffle-assertions");

const NativeHandlerMock = artifacts.require("NativeHandlerMock");

NativeHandlerMock.numberFormat = "BigNumber";

describe("NativeHandler", () => {
  const chainName = "ethereum";
  const baseAmount = wei("10");
  const originHash = "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d";
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
      await truffleAssert.reverts(handler.withdrawNativeBundle(0, { salt: salt, bundle: "0x" }), "Bundler: not this");
    });
  });

  describe("depositNative", () => {
    it("should deposit native", async () => {
      let tx = await handler.depositNative("receiver", { salt: salt, bundle: "0x000001" }, "kovan", {
        value: baseAmount,
      });

      let realSalt = web3.utils.soliditySha3({
        value: web3.eth.abi.encodeParameters(["bytes32", "address"], [salt, OWNER]),
        type: "bytes",
      });

      assert.equal(await web3.eth.getBalance(handler.address), baseAmount);
      assert.equal(tx.receipt.logs[0].event, "DepositedNative");
      assert.equal(tx.receipt.logs[0].args.amount, baseAmount);
      assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
      assert.equal(tx.receipt.logs[0].args.salt, realSalt);
      assert.equal(tx.receipt.logs[0].args.bundle, "0x000001");
      assert.equal(tx.receipt.logs[0].args.network, "kovan");
    });

    it("should revert when depositing 0 tokens", async () => {
      await truffleAssert.reverts(
        handler.depositNative("receiver", { salt: salt, bundle: "0x" }, "kovan"),
        "NativeHandler: zero value"
      );
    });
  });

  describe("getNativeMerkleLeaf", () => {
    it("should encode args", async () => {
      let bundle = "0x";

      let merkleLeaf0 = await handler.getNativeMerkleLeaf(
        baseAmount,
        OWNER,
        { salt: salt, bundle: bundle },
        originHash,
        chainName
      );

      assert.equal(
        merkleLeaf0,
        web3.utils.soliditySha3(
          { value: baseAmount, type: "uint256" },
          { value: OWNER, type: "address" },
          { value: salt, type: "bytes32" },
          { value: bundle, type: "bytes" },
          { value: originHash, type: "bytes32" },
          { value: chainName, type: "string" },
          { value: handler.address, type: "address" }
        )
      );

      let merkleLeaf1 = await handler.getNativeMerkleLeaf(
        wei("1"),
        OWNER,
        { salt: salt, bundle: bundle },
        originHash,
        "BSC"
      );

      assert.equal(
        merkleLeaf1,
        web3.utils.soliditySha3(
          { value: wei("1"), type: "uint256" },
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

  describe("withdrawNative", () => {
    it("should withdraw native", async () => {
      await handler.depositNative("receiver", { salt: salt, bundle: "0x" }, "kovan", { value: baseAmount });
      await handler.withdrawNative(baseAmount, OWNER, { salt: salt, bundle: "0x" });

      assert.equal(await web3.eth.getBalance(handler.address), 0);
    });

    it("should revert when amount is 0", async () => {
      await truffleAssert.reverts(
        handler.withdrawNative(0, OWNER, { salt: salt, bundle: "0x" }),
        "NativeHandler: amount is zero"
      );
    });

    it("should revert when receiver address is 0", async () => {
      await truffleAssert.reverts(
        handler.withdrawNative(baseAmount, "0x0000000000000000000000000000000000000000", { salt: salt, bundle: "0x" }),
        "NativeHandler: receiver is zero"
      );
    });

    it("should revert when amount more than balance", async () => {
      await truffleAssert.reverts(
        handler.withdrawNative(wei("1000000"), OWNER, { salt: salt, bundle: "0x" }),
        "NativeHandler: failed to send eth"
      );
    });
  });
});
