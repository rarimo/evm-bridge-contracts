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

  describe("depositERC20", () => {
    it("should deposit 100 tokens, isWrapped = true", async () => {
      let expectedAmount = wei("100");

      let tx = await handler.depositERC20(token.address, expectedAmount, "receiver", "kovan", true);

      assert.equal((await token.balanceOf(OWNER)).toFixed(), toBN(baseBalance).minus(toBN(expectedAmount)).toFixed());
      assert.equal(await token.balanceOf(handler.address), "0");
      assert.equal(tx.receipt.logs[0].event, "DepositedERC20");
      assert.equal(tx.receipt.logs[0].args.token, token.address);
      assert.equal(tx.receipt.logs[0].args.amount, expectedAmount);
      assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
      assert.equal(tx.receipt.logs[0].args.network, "kovan");
      assert.isTrue(tx.receipt.logs[0].args.isWrapped);
    });

    it("should not burn tokens if they are not approved", async () => {
      let expectedAmount = wei("100");

      await token.approve(handler.address, 0);

      await truffleAssert.reverts(
        handler.depositERC20(token.address, expectedAmount, "receiver", "kovan", true),
        "ERC20: insufficient allowance"
      );
    });

    it("should deposit 52 tokens, isWrapped = false", async () => {
      let expectedAmount = wei("52");

      let tx = await handler.depositERC20(token.address, expectedAmount, "receiver", "kovan", false);

      assert.equal((await token.balanceOf(OWNER)).toFixed(), toBN(baseBalance).minus(toBN(expectedAmount)).toFixed());
      assert.equal((await token.balanceOf(handler.address)).toFixed(), expectedAmount);
      assert.equal(tx.receipt.logs[0].event, "DepositedERC20");
      assert.equal(tx.receipt.logs[0].args.token, token.address);
      assert.equal(tx.receipt.logs[0].args.amount, expectedAmount);
      assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
      assert.equal(tx.receipt.logs[0].args.network, "kovan");
      assert.isFalse(tx.receipt.logs[0].args.isWrapped);
    });

    it("should revert when try deposit 0 tokens", async () => {
      let expectedAmount = wei("0");
      await truffleAssert.reverts(
        handler.depositERC20(token.address, expectedAmount, "receiver", "kovan", false),
        "ERC20Handler: amount is zero"
      );
    });

    it("should revert when token address is 0", async () => {
      let expectedAmount = wei("1");
      await truffleAssert.reverts(
        handler.depositERC20("0x0000000000000000000000000000000000000000", expectedAmount, "receiver", "kovan", false),
        "ERC20Handler: zero token"
      );
    });
  });

  describe("getERC20MerkleLeaf", () => {
    it("should encode args", async () => {
      let amount = wei("100");
      let originHash = "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d";

      let merkleLeaf0 = await handler.getERC20MerkleLeaf(token.address, amount, OWNER, originHash, chainName);

      assert.equal(
        merkleLeaf0,
        web3.utils.soliditySha3(
          { value: token.address, type: "address" },
          { value: amount, type: "uint256" },
          { value: OWNER, type: "address" },
          { value: originHash, type: "bytes32" },
          { value: chainName, type: "string" },
          { value: handler.address, type: "address" }
        )
      );

      let merkleLeaf1 = await handler.getERC20MerkleLeaf(token.address, amount, OWNER, originHash, "BSC");

      assert.equal(
        merkleLeaf1,
        web3.utils.soliditySha3(
          { value: token.address, type: "address" },
          { value: amount, type: "uint256" },
          { value: OWNER, type: "address" },
          { value: originHash, type: "bytes32" },
          { value: "BSC", type: "string" },
          { value: handler.address, type: "address" }
        )
      );

      assert.notEqual(merkleLeaf0, merkleLeaf1);
    });
  });

  describe("withdrawERC20", () => {
    it("should withdraw 100 tokens, is wrapped = true", async () => {
      let amount = wei("100");

      await handler.depositERC20(token.address, amount, "receiver", "kovan", true);
      await handler.withdrawERC20(token.address, amount, OWNER, true);

      assert.equal((await token.balanceOf(OWNER)).toFixed(), baseBalance);
      assert.equal(await token.balanceOf(handler.address), "0");
    });

    it("should withdraw 52 tokens, is wrapped = false", async () => {
      let amount = wei("52");

      await handler.depositERC20(token.address, amount, "receiver", "kovan", false);
      await handler.withdrawERC20(token.address, amount, OWNER, false);

      assert.equal((await token.balanceOf(OWNER)).toFixed(), baseBalance);
      assert.equal(await token.balanceOf(handler.address), "0");
    });

    it("should revert when try to withdraw 0 tokens", async () => {
      await truffleAssert.reverts(
        handler.withdrawERC20(token.address, "0", OWNER, false),
        "ERC20Handler: amount is zero"
      );
    });

    it("should revert when try token address 0", async () => {
      await truffleAssert.reverts(
        handler.withdrawERC20("0x0000000000000000000000000000000000000000", wei("100"), OWNER, false),
        "ERC20Handler: zero token"
      );
    });

    it("should revert when receiver address is 0", async () => {
      await truffleAssert.reverts(
        handler.withdrawERC20(token.address, wei("100"), "0x0000000000000000000000000000000000000000", false),
        "ERC20Handler: zero receiver"
      );
    });
  });
});
