const { assert } = require("chai");
const { accounts } = require("../../scripts/helpers/utils");
const truffleAssert = require("truffle-assertions");

const ERC721HandlerMock = artifacts.require("ERC721HandlerMock");
const ERC721MB = artifacts.require("ERC721MintableBurnable");
const SBTMB = artifacts.require("SBTMintableBurnable");

ERC721HandlerMock.numberFormat = "BigNumber";
ERC721MB.numberFormat = "BigNumber";
SBTMB.numberFormat = "BigNumber";

describe("ERC721Handler", () => {
  const baseId = "5000";
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
    handler = await ERC721HandlerMock.new();
  });

  describe("ERC721", () => {
    beforeEach(async () => {
      token = await ERC721MB.new("Mock", "MK", OWNER, "");

      await token.mintTo(OWNER, baseId, "URI");
      await token.approve(handler.address, baseId);

      await token.transferOwnership(handler.address);
    });

    describe("access", () => {
      it("only this should call this method", async () => {
        const tokenData = web3.eth.abi.encodeParameters(["address", "uint256", "string"], [token.address, 0, ""]);

        await truffleAssert.reverts(
          handler.withdrawERC721Bundle(tokenData, { salt: salt, bundle: "0x" }, true),
          "Bundler: not this"
        );
      });
    });

    describe("depositERC721", () => {
      it("should deposit token, isWrapped = true", async () => {
        assert.isFalse(await token.supportsInterface("0x00000000"));

        let tx = await handler.depositERC721(
          token.address,
          baseId,
          { salt: salt, bundle: "0x00" },
          "kovan",
          "receiver",
          true
        );

        const realSalt = web3.utils.soliditySha3({ value: salt, type: "bytes32" }, { value: OWNER, type: "address" });

        assert.equal(tx.receipt.logs[0].event, "DepositedERC721");
        assert.equal(tx.receipt.logs[0].args.token, token.address);
        assert.equal(tx.receipt.logs[0].args.tokenId, baseId);
        assert.equal(tx.receipt.logs[0].args.salt, realSalt);
        assert.equal(tx.receipt.logs[0].args.bundle, "0x00");
        assert.equal(tx.receipt.logs[0].args.network, "kovan");
        assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
        assert.isTrue(tx.receipt.logs[0].args.isWrapped);

        await truffleAssert.reverts(token.ownerOf(baseId), "ERC721: owner query for nonexistent token");
      });

      it("should deposit token, isWrapped = true (2)", async () => {
        await token.approve("0x0000000000000000000000000000000000000000", baseId);
        await token.setApprovalForAll(handler.address, true);

        await truffleAssert.passes(
          handler.depositERC721(token.address, baseId, { salt: salt, bundle: "0x00" }, "kovan", "receiver", true),
          "pass"
        );
      });

      it("should not burn token if it is not approved", async () => {
        await token.approve(token.address, baseId);

        await truffleAssert.reverts(
          handler.depositERC721(token.address, baseId, { salt: salt, bundle: "0x" }, "kovan", "receiver", true),
          "ERC721MintableBurnable: not approved"
        );
      });

      it("should not burn token if it is approved but not owned", async () => {
        await truffleAssert.reverts(
          handler.depositERC721(token.address, baseId, { salt: salt, bundle: "0x" }, "kovan", "receiver", true, {
            from: SECOND,
          }),
          "ERC721MintableBurnable: not approved"
        );
      });

      it("should deposit token, isWrapped = false", async () => {
        let tx = await handler.depositERC721(
          token.address,
          baseId,
          { salt: salt, bundle: "0x0123" },
          "kovan",
          "receiver",
          false
        );

        const realSalt = web3.utils.soliditySha3({ value: salt, type: "bytes32" }, { value: OWNER, type: "address" });

        assert.equal(await token.ownerOf(baseId), handler.address);

        assert.equal(tx.receipt.logs[0].event, "DepositedERC721");
        assert.equal(tx.receipt.logs[0].args.token, token.address);
        assert.equal(tx.receipt.logs[0].args.tokenId, baseId);
        assert.equal(tx.receipt.logs[0].args.salt, realSalt);
        assert.equal(tx.receipt.logs[0].args.bundle, "0x0123");
        assert.equal(tx.receipt.logs[0].args.network, "kovan");
        assert.equal(tx.receipt.logs[0].args.receiver, "receiver");
        assert.isFalse(tx.receipt.logs[0].args.isWrapped);
      });

      it("should revert when token address is 0", async () => {
        await truffleAssert.reverts(
          handler.depositERC721(
            "0x0000000000000000000000000000000000000000",
            baseId,
            { salt: salt, bundle: "0x00" },
            "kovan",
            "receiver",
            false
          ),
          "ERC721Handler: zero token"
        );
      });
    });

    describe("withdrawERC721", async () => {
      it("should withdraw token, wrapped = true", async () => {
        const tokenData = web3.eth.abi.encodeParameters(
          ["address", "uint256", "string"],
          [token.address, baseId, "URI1"]
        );

        assert.equal(await token.tokenURI(baseId), "URI");

        await handler.depositERC721(token.address, baseId, { salt: salt, bundle: "0x" }, "kovan", "receiver", true);
        await handler.withdrawERC721(tokenData, OWNER, true);

        assert.equal(await token.ownerOf(baseId), OWNER);
        assert.equal(await token.tokenURI(baseId), "URI1");
      });

      it("should withdraw token, wrapped = false", async () => {
        const tokenData = web3.eth.abi.encodeParameters(
          ["address", "uint256", "string"],
          [token.address, baseId, "URI1"]
        );

        await handler.depositERC721(token.address, baseId, { salt: salt, bundle: "0x" }, "kovan", "receiver", false);
        await handler.withdrawERC721(tokenData, OWNER, false);

        assert.equal(await token.ownerOf(baseId), OWNER);
        assert.equal(await token.tokenURI(baseId), "URI");
      });

      it("should revert when token address is 0", async () => {
        const tokenData = web3.eth.abi.encodeParameters(
          ["address", "uint256", "string"],
          ["0x0000000000000000000000000000000000000000", baseId, ""]
        );

        await truffleAssert.reverts(handler.withdrawERC721(tokenData, OWNER, false), "ERC721Handler: zero token");
      });

      it("should revert when receiver address is 0", async () => {
        const tokenData = web3.eth.abi.encodeParameters(["address", "uint256", "string"], [token.address, baseId, ""]);

        await truffleAssert.reverts(
          handler.withdrawERC721(tokenData, "0x0000000000000000000000000000000000000000", false),
          "ERC721Handler: zero receiver"
        );
      });
    });
  });

  describe("SBT", () => {
    beforeEach(async () => {
      token = await SBTMB.new("Mock", "MK", OWNER, "");

      await token.attestTo(OWNER, baseId, "URI");
      await token.attestTo(SECOND, 1, "");
    });

    describe("token", () => {
      describe("supportsInterface", () => {
        it("should support correct interfaces", async () => {
          assert.isTrue(await token.supportsInterface("0x80ac58cd"));
          assert.isTrue(await token.supportsInterface("0xf6809be0"));
          assert.isTrue(await token.supportsInterface("0xb45a3c0e"));
          assert.isFalse(await token.supportsInterface("0x00000000"));
        });
      });

      describe("attest", () => {
        it("should not attest if user already has a token", async () => {
          await truffleAssert.reverts(
            token.attestTo(SECOND, 2, ""),
            "SBTMintableBurnable: receiver already has a token"
          );
        });

        it("should not attest if caller is not the owner", async () => {
          await truffleAssert.reverts(
            token.attestTo(SECOND, 2, "", { from: SECOND }),
            "Ownable: caller is not the owner"
          );
        });
      });

      describe("transfer", () => {
        it("should not transfer an sbt token", async () => {
          await truffleAssert.reverts(
            token.safeTransferFrom(OWNER, SECOND, baseId),
            "SBTMintableBurnable: not transferable"
          );
        });
      });
    });

    describe("handler", () => {
      beforeEach(async () => {
        await token.transferOwnership(handler.address);
      });

      describe("access", () => {
        it("only this should call this method", async () => {
          const tokenData = web3.eth.abi.encodeParameters(["address", "uint256", "string"], [token.address, 0, ""]);

          await truffleAssert.reverts(
            handler.withdrawSBTBundle(tokenData, { salt: salt, bundle: "0x" }, false),
            "Bundler: not this"
          );
        });
      });

      describe("depositSBT", () => {
        it("should revert when token address is 0", async () => {
          await truffleAssert.reverts(
            handler.depositSBT(
              "0x0000000000000000000000000000000000000000",
              baseId,
              { salt: salt, bundle: "0x00" },
              "kovan",
              "receiver"
            ),
            "ERC721Handler: zero token"
          );
        });

        it("should revert if not owned token", async () => {
          await truffleAssert.reverts(
            handler.depositSBT(token.address, 1, { salt: salt, bundle: "0x00" }, "kovan", "receiver"),
            "ERC721Handler: invalid token id"
          );
        });

        it("should deposit if all conditions are met", async () => {
          const tx = await handler.depositSBT(
            token.address,
            baseId,
            { salt: salt, bundle: "0x00" },
            "kovan",
            "receiver"
          );

          const realSalt = web3.utils.soliditySha3({ value: salt, type: "bytes32" }, { value: OWNER, type: "address" });

          assert.equal(tx.receipt.logs[0].event, "DepositedSBT");
          assert.equal(tx.receipt.logs[0].args.token, token.address);
          assert.equal(tx.receipt.logs[0].args.tokenId, baseId);
          assert.equal(tx.receipt.logs[0].args.salt, realSalt);
          assert.equal(tx.receipt.logs[0].args.bundle, "0x00");
          assert.equal(tx.receipt.logs[0].args.network, "kovan");
          assert.equal(tx.receipt.logs[0].args.receiver, "receiver");

          assert.equal((await token.tokenIdOf(OWNER)).toFixed(), baseId);
        });
      });

      describe("withdrawSBT", () => {
        it("should revert when token address is 0", async () => {
          const tokenData = web3.eth.abi.encodeParameters(
            ["address", "uint256", "string"],
            ["0x0000000000000000000000000000000000000000", baseId, ""]
          );

          await truffleAssert.reverts(handler.withdrawSBT(tokenData, OWNER), "ERC721Handler: zero token");
        });

        it("should revert when receiver address is 0", async () => {
          const tokenData = web3.eth.abi.encodeParameters(
            ["address", "uint256", "string"],
            [token.address, baseId, ""]
          );

          await truffleAssert.reverts(
            handler.withdrawSBT(tokenData, "0x0000000000000000000000000000000000000000"),
            "ERC721Handler: zero receiver"
          );
        });

        it("should revert when receiver address is 0", async () => {
          await handler.depositSBT(token.address, baseId, { salt: salt, bundle: "0x00" }, "kovan", "receiver");

          await token.burn();

          await truffleAssert.reverts(token.tokenIdOf(OWNER), "SBTMintableBurnable: user doesn't have a token");
          await truffleAssert.reverts(token.locked(baseId), "SBTMintableBurnable: token is assigned to zero address");

          assert.equal((await token.balanceOf(OWNER)).toFixed(), "0");

          const tokenData = web3.eth.abi.encodeParameters(
            ["address", "uint256", "string"],
            [token.address, baseId, "URI"]
          );

          await handler.withdrawSBT(tokenData, OWNER);

          assert.equal((await token.tokenIdOf(OWNER)).toFixed(), baseId);
          assert.isTrue(await token.locked(baseId));
          assert.equal((await token.balanceOf(OWNER)).toFixed(), "1");
          assert.equal(await token.tokenURI(baseId), "URI");
        });
      });
    });
  });
});
