const { assert } = require("chai");
const { accounts, wei } = require("../../scripts/helpers/utils");
const truffleAssert = require("truffle-assertions");

const NativeHandlerMock = artifacts.require("NativeHandlerMock");

NativeHandlerMock.numberFormat = "BigNumber";

describe("NativeHandler", () => {
  const baseAmount = wei("10");

  let OWNER;
  let handler;

  before("setup", async () => {
    OWNER = await accounts(0);
  });

  beforeEach("setup", async () => {
    handler = await NativeHandlerMock.new();
  });

  describe("depositNative", () => {
    it("should deposit native", async () => {
      let tx = await handler.depositNative("receiver", "kovan", { value: baseAmount });

      assert.equal(await web3.eth.getBalance(handler.address), baseAmount);
      assert.equal(tx.receipt.logs[0].event, "DepositedNative");
      assert.equal(tx.receipt.logs[0].args.amount, baseAmount);
      assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
      assert.equal(tx.receipt.logs[0].args.network, "kovan");
    });

    it("should revert when try deposit 0 tokens", async () => {
      await truffleAssert.reverts(handler.depositNative("receiver", "kovan"), "NativeHandler: zero value");
    });
  });

  describe("getNativeMerkleLeaf", () => {
    it("should encode args", async () => {
      let originHash = "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d";

      let merkleLeaf0 = await handler.getNativeMerkleLeaf(baseAmount, OWNER, originHash, "ethereum", handler.address);

      assert.equal(
        merkleLeaf0,
        web3.utils.soliditySha3(
          { value: baseAmount, type: "uint256" },
          { value: OWNER, type: "address" },
          { value: originHash, type: "bytes32" },
          { value: "ethereum", type: "string" },
          { value: handler.address, type: "address" }
        )
      );

      let merkleLeaf1 = await handler.getNativeMerkleLeaf(wei("1"), OWNER, originHash, "BSC", handler.address);

      assert.equal(
        merkleLeaf1,
        web3.utils.soliditySha3(
          { value: wei("1"), type: "uint256" },
          { value: OWNER, type: "address" },
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
      await handler.depositNative("receiver", "kovan", { value: baseAmount });
      await handler.withdrawNative(baseAmount, OWNER);

      assert.equal(await web3.eth.getBalance(handler.address), 0);
    });

    it("should revert when amount is 0", async () => {
      await truffleAssert.reverts(handler.withdrawNative(0, OWNER), "NativeHandler: amount is zero");
    });

    it("should revert when receiver address is 0", async () => {
      await truffleAssert.reverts(
        handler.withdrawNative(baseAmount, "0x0000000000000000000000000000000000000000"),
        "NativeHandler: receiver is zero"
      );
    });

    it("should revert when amount more than balance", async () => {
      await truffleAssert.reverts(handler.withdrawNative(wei("1000000"), OWNER), "NativeHandler: can't send eth");
    });
  });
});
