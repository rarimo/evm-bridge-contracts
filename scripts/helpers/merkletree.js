const { MerkleTree } = require("merkletreejs");

const constructTree = (leaves) => {
  const tree = new MerkleTree(leaves, (el) => Buffer.from(web3.utils.soliditySha3(el).replace("0x", ""), "hex"), {
    sortPairs: true,
  });

  return tree;
};

const getProof = (leaf, tree) => {
  const proof = tree.getProof(leaf).map((el) => "0x" + el.data.toString("hex"));

  return proof;
};

const getRoot = (tree) => {
  return "0x" + tree.getRoot().toString("hex");
};

module.exports = {
  constructTree,
  getProof,
  getRoot,
};
