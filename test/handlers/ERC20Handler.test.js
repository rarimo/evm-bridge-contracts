const { assert } = require("chai");
const { toBN, accounts, wei } = require("../../scripts/helpers/utils");
const truffleAssert = require("truffle-assertions");

const ERC20HandlerMock = artifacts.require("ERC20HandlerMock");
const ERC20MB = artifacts.require("ERC20MintableBurnable");

ERC20MB.numberFormat = "BigNumber";
ERC20HandlerMock.numberFormat = "BigNumber";

describe("ERC20Handler", () => {
  const baseBalance = wei("1000000");
  const salt = "0x0000000000000000000000000000000000000000000000000000000000000001";

  let OWNER;
  let handler;
  let token;

  before("setup", async () => {
    OWNER = await accounts(0);
  });

  beforeEach("setup", async () => {
    token = await ERC20MB.new("Mock", "MK", OWNER);
    handler = await ERC20HandlerMock.new();

    await token.mintTo(OWNER, baseBalance);
    await token.approve(handler.address, baseBalance);

    await token.transferOwnership(handler.address);
  });

  describe("access", () => {
    it("only this should call this method", async () => {
      const tokenData = web3.eth.abi.encodeParameters(["address", "uint256"], [token.address, 0]);

      await truffleAssert.reverts(
        handler.withdrawERC20Bundle(tokenData, { salt: salt, bundle: "0x" }, true),
        "Bundler: not this"
      );
    });
  });

  describe("depositERC20", () => {
    it("should deposit 100 tokens, isWrapped = true", async () => {
      let expectedAmount = wei("100");

      let tx = await handler.depositERC20(
        token.address,
        expectedAmount,
        { salt: salt, bundle: "0x" },
        "kovan",
        "receiver",
        true
      );

      const realSalt = web3.utils.soliditySha3({ value: salt, type: "bytes32" }, { value: OWNER, type: "address" });

      assert.equal((await token.balanceOf(OWNER)).toFixed(), toBN(baseBalance).minus(toBN(expectedAmount)).toFixed());
      assert.equal(await token.balanceOf(handler.address), "0");

      assert.equal(tx.receipt.logs[0].event, "DepositedERC20");
      assert.equal(tx.receipt.logs[0].args.token, token.address);
      assert.equal(tx.receipt.logs[0].args.amount, expectedAmount);
      assert.equal(tx.receipt.logs[0].args.salt, realSalt);
      assert.isNull(tx.receipt.logs[0].args.bundle);
      assert.equal(tx.receipt.logs[0].args.network, "kovan");
      assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
      assert.isTrue(tx.receipt.logs[0].args.isWrapped);
    });

    it("should not burn tokens if they are not approved", async () => {
      await token.approve(handler.address, 0);

      await truffleAssert.reverts(
        handler.depositERC20(token.address, wei("100"), { salt: salt, bundle: "0x" }, "kovan", "receiver", true),
        "ERC20: insufficient allowance"
      );
    });

    it("should deposit 52 tokens, isWrapped = false", async () => {
      let expectedAmount = wei("52");

      let tx = await handler.depositERC20(
        token.address,
        expectedAmount,
        { salt: salt, bundle: "0x0123" },
        "kovan",
        "receiver",
        false
      );

      const realSalt = web3.utils.soliditySha3({ value: salt, type: "bytes32" }, { value: OWNER, type: "address" });

      assert.equal((await token.balanceOf(OWNER)).toFixed(), toBN(baseBalance).minus(toBN(expectedAmount)).toFixed());
      assert.equal((await token.balanceOf(handler.address)).toFixed(), expectedAmount);

      assert.equal(tx.receipt.logs[0].event, "DepositedERC20");
      assert.equal(tx.receipt.logs[0].args.token, token.address);
      assert.equal(tx.receipt.logs[0].args.amount, expectedAmount);
      assert.equal(tx.receipt.logs[0].args.salt, realSalt);
      assert.equal(tx.receipt.logs[0].args.bundle, "0x0123");
      assert.equal(tx.receipt.logs[0].args.network, "kovan");
      assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
      assert.isFalse(tx.receipt.logs[0].args.isWrapped);
    });

    it("should revert when depositing 0 tokens", async () => {
      await truffleAssert.reverts(
        handler.depositERC20(token.address, 0, { salt: salt, bundle: "0x" }, "kovan", "receiver", false),
        "ERC20Handler: amount is zero"
      );
    });

    it("should revert when token address is 0", async () => {
      await truffleAssert.reverts(
        handler.depositERC20(
          "0x0000000000000000000000000000000000000000",
          wei("1"),
          { salt: salt, bundle: "0x" },
          "kovan",
          "receiver",
          false
        ),
        "ERC20Handler: zero token"
      );
    });
  });

  describe("withdrawERC20", () => {
    it("should withdraw 100 tokens, is wrapped = true", async () => {
      const tokenData = web3.eth.abi.encodeParameters(["address", "uint256"], [token.address, wei("100")]);

      await handler.depositERC20(
        token.address,
        wei("100"),
        { salt: salt, bundle: "0x0123" },
        "kovan",
        "receiver",
        true
      );
      await handler.withdrawERC20(tokenData, OWNER, true);

      assert.equal((await token.balanceOf(OWNER)).toFixed(), baseBalance);
      assert.equal(await token.balanceOf(handler.address), "0");
    });

    it("should withdraw 52 tokens, is wrapped = false", async () => {
      const tokenData = web3.eth.abi.encodeParameters(["address", "uint256"], [token.address, wei("52")]);

      await handler.depositERC20(
        token.address,
        wei("52"),
        { salt: salt, bundle: "0x0123" },
        "kovan",
        "receiver",
        false
      );
      await handler.withdrawERC20(tokenData, OWNER, false);

      assert.equal((await token.balanceOf(OWNER)).toFixed(), baseBalance);
      assert.equal(await token.balanceOf(handler.address), "0");
    });

    it("should withdraw 0 tokens, isWrapped = false", async () => {
      const tokenData = web3.eth.abi.encodeParameters(["address", "uint256"], [token.address, 0]);

      await truffleAssert.passes(handler.withdrawERC20(tokenData, OWNER, false), "pass");
    });

    it("should withdraw 0 tokens, isWrapped = true", async () => {
      const tokenData = web3.eth.abi.encodeParameters(["address", "uint256"], [token.address, 0]);

      await truffleAssert.passes(handler.withdrawERC20(tokenData, OWNER, true), "pass");
    });

    it("should revert when withdrawal token address is 0", async () => {
      const tokenData = web3.eth.abi.encodeParameters(
        ["address", "uint256"],
        ["0x0000000000000000000000000000000000000000", wei("100")]
      );

      await truffleAssert.reverts(handler.withdrawERC20(tokenData, OWNER, false), "ERC20Handler: zero token");
    });

    it("should revert when receiver address is 0", async () => {
      const tokenData = web3.eth.abi.encodeParameters(["address", "uint256"], [token.address, wei("100")]);

      await truffleAssert.reverts(
        handler.withdrawERC20(tokenData, "0x0000000000000000000000000000000000000000", false),
        "ERC20Handler: zero receiver"
      );
    });
  });
});
