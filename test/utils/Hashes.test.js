const { assert } = require("chai");
const truffleAssert = require("truffle-assertions");

const Hashes = artifacts.require("HashesMock");

Hashes.numberFormat = "BigNumber";

describe("Hashes", () => {
  let hashes;

  let originHash;

  beforeEach("setup", async () => {
    hashes = await Hashes.new();

    let originChainName = "BSC";
    let originTxHash = "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d";
    let originEventNonce = "1794147";

    originHash = web3.utils.soliditySha3(
      { value: originChainName, type: "string" },
      { value: originTxHash, type: "bytes32" },
      { value: originEventNonce, type: "uint256" }
    );
  });

  describe("checkAndUpdateHashes", () => {
    it("should call checkAndUpdateHashes", async () => {
      await hashes.checkAndUpdateHashes(originHash);

      assert.isTrue(await hashes.usedHashes(originHash));
    });

    it("should revert when try add hash twice", async () => {
      await hashes.checkAndUpdateHashes(originHash);

      await truffleAssert.reverts(hashes.checkAndUpdateHashes(originHash), "Hashes: the hash nonce is used");
    });
  });
});
