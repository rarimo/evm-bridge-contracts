const { assert } = require("chai");
const { toBN, accounts, wei } = require("../../scripts/helpers/utils");
const truffleAssert = require("truffle-assertions");

const ERC20HandlerMock = artifacts.require("ERC20HandlerMock");
const ERC20MB = artifacts.require("ERC20MintableBurnable");

ERC20MB.numberFormat = "BigNumber";
ERC20HandlerMock.numberFormat = "BigNumber";

describe("ERC20Handler", () => {
  const chainName = "ethereum";
  const baseBalance = wei("1000000");
  const originHash = "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d";
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
      await truffleAssert.reverts(
        handler.withdrawERC20Bundle(token.address, 0, { salt: salt, bundle: "0x" }, true),
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
        "receiver",
        { salt: salt, bundle: "0x" },
        "kovan",
        true
      );

      let realSalt = web3.utils.soliditySha3({
        value: web3.eth.abi.encodeParameters(["bytes32", "address"], [salt, OWNER]),
        type: "bytes",
      });

      assert.equal((await token.balanceOf(OWNER)).toFixed(), toBN(baseBalance).minus(toBN(expectedAmount)).toFixed());
      assert.equal(await token.balanceOf(handler.address), "0");
      assert.equal(tx.receipt.logs[0].event, "DepositedERC20");
      assert.equal(tx.receipt.logs[0].args.token, token.address);
      assert.equal(tx.receipt.logs[0].args.amount, expectedAmount);
      assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
      assert.equal(tx.receipt.logs[0].args.salt, realSalt);
      assert.isNull(tx.receipt.logs[0].args.bundle);
      assert.equal(tx.receipt.logs[0].args.network, "kovan");
      assert.isTrue(tx.receipt.logs[0].args.isWrapped);
    });

    it("should not burn tokens if they are not approved", async () => {
      let expectedAmount = wei("100");

      await token.approve(handler.address, 0);

      await truffleAssert.reverts(
        handler.depositERC20(token.address, expectedAmount, "receiver", { salt: salt, bundle: "0x" }, "kovan", true),
        "ERC20: insufficient allowance"
      );
    });

    it("should deposit 52 tokens, isWrapped = false", async () => {
      let expectedAmount = wei("52");

      let tx = await handler.depositERC20(
        token.address,
        expectedAmount,
        "receiver",
        { salt: salt, bundle: "0x0123" },
        "kovan",
        false
      );

      let realSalt = web3.utils.soliditySha3({
        value: web3.eth.abi.encodeParameters(["bytes32", "address"], [salt, OWNER]),
        type: "bytes",
      });

      assert.equal((await token.balanceOf(OWNER)).toFixed(), toBN(baseBalance).minus(toBN(expectedAmount)).toFixed());
      assert.equal((await token.balanceOf(handler.address)).toFixed(), expectedAmount);
      assert.equal(tx.receipt.logs[0].event, "DepositedERC20");
      assert.equal(tx.receipt.logs[0].args.token, token.address);
      assert.equal(tx.receipt.logs[0].args.amount, expectedAmount);
      assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
      assert.equal(tx.receipt.logs[0].args.salt, realSalt);
      assert.equal(tx.receipt.logs[0].args.bundle, "0x0123");
      assert.equal(tx.receipt.logs[0].args.network, "kovan");
      assert.isFalse(tx.receipt.logs[0].args.isWrapped);
    });

    it("should revert when depositing 0 tokens", async () => {
      await truffleAssert.reverts(
        handler.depositERC20(token.address, wei("0"), "receiver", { salt: salt, bundle: "0x" }, "kovan", false),
        "ERC20Handler: amount is zero"
      );
    });

    it("should revert when token address is 0", async () => {
      await truffleAssert.reverts(
        handler.depositERC20(
          "0x0000000000000000000000000000000000000000",
          wei("1"),
          "receiver",
          { salt: salt, bundle: "0x" },
          "kovan",
          false
        ),
        "ERC20Handler: zero token"
      );
    });
  });

  describe("getERC20MerkleLeaf", () => {
    it("should encode args", async () => {
      let amount = wei("100");
      let bundle = "0x012345678987654321";

      let merkleLeaf0 = await handler.getERC20MerkleLeaf(
        token.address,
        amount,
        OWNER,
        { salt: salt, bundle: bundle },
        originHash,
        chainName
      );

      assert.equal(
        merkleLeaf0,
        web3.utils.soliditySha3({
          value: web3.eth.abi.encodeParameters(
            ["address", "uint256", "address", "tuple(bytes32,bytes)", "bytes32", "string", "address"],
            [token.address, amount, OWNER, [salt, bundle], originHash, chainName, handler.address]
          ),
          type: "bytes",
        })
      );

      let merkleLeaf1 = await handler.getERC20MerkleLeaf(
        token.address,
        amount,
        OWNER,
        { salt: salt, bundle: bundle },
        originHash,
        "BSC"
      );

      assert.equal(
        merkleLeaf1,
        web3.utils.soliditySha3({
          value: web3.eth.abi.encodeParameters(
            ["address", "uint256", "address", "tuple(bytes32,bytes)", "bytes32", "string", "address"],
            [token.address, amount, OWNER, [salt, bundle], originHash, "BSC", handler.address]
          ),
          type: "bytes",
        })
      );

      assert.notEqual(merkleLeaf0, merkleLeaf1);
    });
  });

  describe("withdrawERC20", () => {
    it("should withdraw 100 tokens, is wrapped = true", async () => {
      let amount = wei("100");

      await handler.depositERC20(token.address, amount, "receiver", { salt: salt, bundle: "0x0123" }, "kovan", true);
      await handler.withdrawERC20(token.address, amount, OWNER, { salt: salt, bundle: "0x" }, true);

      assert.equal((await token.balanceOf(OWNER)).toFixed(), baseBalance);
      assert.equal(await token.balanceOf(handler.address), "0");
    });

    it("should withdraw 52 tokens, is wrapped = false", async () => {
      let amount = wei("52");

      await handler.depositERC20(token.address, amount, "receiver", { salt: salt, bundle: "0x0123" }, "kovan", false);
      await handler.withdrawERC20(token.address, amount, OWNER, { salt: salt, bundle: "0x" }, false);

      assert.equal((await token.balanceOf(OWNER)).toFixed(), baseBalance);
      assert.equal(await token.balanceOf(handler.address), "0");
    });

    it("should revert when try to withdraw 0 tokens", async () => {
      await truffleAssert.reverts(
        handler.withdrawERC20(token.address, "0", OWNER, { salt: salt, bundle: "0x" }, false),
        "ERC20Handler: amount is zero"
      );
    });

    it("should revert when try token address 0", async () => {
      await truffleAssert.reverts(
        handler.withdrawERC20(
          "0x0000000000000000000000000000000000000000",
          wei("100"),
          OWNER,
          { salt: salt, bundle: "0x" },
          false
        ),
        "ERC20Handler: zero token"
      );
    });

    it("should revert when receiver address is 0", async () => {
      await truffleAssert.reverts(
        handler.withdrawERC20(
          token.address,
          wei("100"),
          "0x0000000000000000000000000000000000000000",
          { salt: salt, bundle: "0x" },
          false
        ),
        "ERC20Handler: zero receiver"
      );
    });
  });
});
