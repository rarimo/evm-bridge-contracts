const { accounts } = require("../../scripts/helpers/utils");
const { constructTree, getProof, getRoot } = require("../../scripts/helpers/merkletree");
const { rawSign } = require("../helpers/signer");
const { OWNER_PRIVATE_KEY, ANOTHER_PRIVATE_KEY } = require("../helpers/keys");
const truffleAssert = require("truffle-assertions");

const Signers = artifacts.require("SignersMock");

Signers.numberFormat = "BigNumber";

describe("Signers", () => {
  const chainName = "ethereum";

  let OWNER;
  let SECOND;

  let signers;

  before("setup", async () => {
    OWNER = await accounts(0);
    SECOND = await accounts(0);
  });

  beforeEach("setup", async () => {
    signers = await Signers.new();

    await signers.__SignersMock_init(OWNER, chainName);
  });

  describe("access", () => {
    it("should not initialize twice", async () => {
      await truffleAssert.reverts(
        signers.__Signers_init(OWNER, chainName),
        "Initializable: contract is not initializing"
      );
    });
  });

  describe("checkSignatures", () => {
    it("should check signatures", async () => {
      const hashToSign = "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d";

      const signature = rawSign(hashToSign, OWNER_PRIVATE_KEY);

      await truffleAssert.passes(signers.checkSignature(hashToSign, signature));
    });

    it("should revert when signer in not signer", async () => {
      const hashToSign = "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d";

      const signature = rawSign(hashToSign, ANOTHER_PRIVATE_KEY);

      await truffleAssert.reverts(signers.checkSignature(hashToSign, signature), "Signers: invalid signature");
    });

    it("should revert when passing wrong data", async () => {
      const hashToSign = "0xc4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d";
      const fakeHash = "0xd4f46c912cc2a1f30891552ac72871ab0f0e977886852bdd5dccd221a595647d";

      const signature = rawSign(hashToSign, OWNER_PRIVATE_KEY);

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
      const leaf = leaves[0];
      const hashToSign = getRoot(tree);
      const path = getProof(leaf, tree);

      const signature = rawSign(hashToSign, OWNER_PRIVATE_KEY);

      const proof = web3.eth.abi.encodeParameters(["bytes32[]", "bytes"], [path, signature]);

      await truffleAssert.passes(signers.checkMerkleSignature(leaf, proof));
    });

    it("should revert if passed merkle leaf is wrong", async () => {
      const leaf = leaves[0];
      const hashToSign = getRoot(tree);
      const path = getProof(leaf, tree);

      const signature = rawSign(hashToSign, OWNER_PRIVATE_KEY);

      const proof = web3.eth.abi.encodeParameters(["bytes32[]", "bytes"], [path, signature]);

      await truffleAssert.reverts(signers.checkMerkleSignature(leaves[1], proof), "Signers: invalid signature");
    });
  });
});
