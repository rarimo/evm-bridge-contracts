const { assert } = require("chai");
const { accounts } = require("../../scripts/helpers/utils");
const { constructTree, getProof, getRoot } = require("../../scripts/helpers/merkletree");
const truffleAssert = require("truffle-assertions");
const ethSigUtil = require("@metamask/eth-sig-util");

const Signers = artifacts.require("SignersMock");

Signers.numberFormat = "BigNumber";

const OWNER_PRIVATE_KEY = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80";
const ANOTHER_PRIVATE_KEY = "df57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e";

describe("Signers", () => {
  let OWNER;
  let SECOND;

  let signers;

  before("setup", async () => {
    OWNER = await accounts(0);
    SECOND = await accounts(0);
  });

  beforeEach("setup", async () => {
    signers = await Signers.new();

    await signers.__SignersMock_init(OWNER);
  });

  describe("checkSignatures", () => {
    it("should check signatures", async () => {
      const ownerKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");
      const hashToSign = "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d";

      const signature = ethSigUtil.personalSign({ privateKey: ownerKey, data: hashToSign });

      await truffleAssert.passes(signers.checkSignature(hashToSign, signature));
    });

    it("should revert when signer in not signer", async () => {
      const anotherKey = Buffer.from(ANOTHER_PRIVATE_KEY, "hex");
      const hashToSign = "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d";

      const signature = ethSigUtil.personalSign({ privateKey: anotherKey, data: hashToSign });

      await truffleAssert.reverts(signers.checkSignature(hashToSign, signature), "Signers: invalid signature");
    });

    it("should revert when passing wrong data", async () => {
      const ownerKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");
      const hashToSign = "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d";
      const fakeHash = "0xd4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d";

      const signature = ethSigUtil.personalSign({ privateKey: ownerKey, data: hashToSign });

      await truffleAssert.reverts(signers.checkSignature(fakeHash, signature), "Signers: invalid signature");
    });
  });

  describe("checkMerkleSignature", () => {
    let leaves;
    let tree;

    beforeEach("setup", async () => {
      leaves = [
        "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a5956471",
        "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a5956472",
        "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a5956473",
        "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a5956474",
        "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a5956475",
      ];

      tree = constructTree(leaves);
    });

    it("should verify signed merkle root", async () => {
      const ownerKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");
      const leaf = leaves[0];
      const hashToSign = getRoot(tree);
      const proof = getProof(leaf, tree);

      const signature = ethSigUtil.personalSign({ privateKey: ownerKey, data: hashToSign });

      await truffleAssert.passes(signers.checkMerkleSignature(leaf, proof, signature));
    });

    it("should revert if passed merkle leaf is wrong", async () => {
      const ownerKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");
      const leaf = leaves[0];
      const hashToSign = getRoot(tree);
      const proof = getProof(leaf, tree);

      const signature = ethSigUtil.personalSign({ privateKey: ownerKey, data: hashToSign });

      await truffleAssert.reverts(
        signers.checkMerkleSignature(leaves[1], proof, signature),
        "Signers: invalid signature"
      );
    });
  });

  describe("changeSigner", () => {
    it("should correctly change signer", async () => {
      const ownerKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

      assert.equal(await signers.signer(), OWNER);
      assert.equal(await signers.nonce(), "0");

      const hashToSign = web3.utils.soliditySha3(
        { value: SECOND, type: "address" },
        { value: "31337", type: "uint256" },
        { value: "0", type: "uint256" },
        { value: signers.address, type: "address" }
      );

      const signature = ethSigUtil.personalSign({ privateKey: ownerKey, data: hashToSign });

      await signers.changeSigner(SECOND, signature);

      assert.equal(await signers.signer(), SECOND);
      assert.equal(await signers.nonce(), "1");
    });

    it("should not change signer if bad signature is passed", async () => {
      const ownerKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

      const hashToSign = web3.utils.soliditySha3(
        { value: SECOND, type: "address" },
        { value: "31337", type: "uint256" },
        { value: "0", type: "uint256" },
        { value: signers.address, type: "address" }
      );
      const signature = ethSigUtil.personalSign({ privateKey: ownerKey, data: hashToSign });

      await signers.changeSigner(SECOND, signature);

      await truffleAssert.reverts(signers.changeSigner(SECOND, signature), "Signers: invalid signature");
    });
  });
});
