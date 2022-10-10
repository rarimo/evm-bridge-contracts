const getBytesDoSomething = (shouldRevert) => {
  return web3.eth.abi.encodeFunctionCall(
    {
      name: "doSomething",
      type: "function",
      inputs: [
        {
          type: "bool",
          name: "shouldRevert",
        },
      ],
    },
    [shouldRevert]
  );
};

const getBytesWithdrawNFT = (nft, tokenId) => {
  return web3.eth.abi.encodeFunctionCall(
    {
      name: "withdrawNFT",
      type: "function",
      inputs: [
        {
          type: "address",
          name: "nft",
        },
        {
          type: "uint256",
          name: "tokenId",
        },
      ],
    },
    [nft, tokenId]
  );
};

const getBytesWithdrawETH = (amount) => {
  return web3.eth.abi.encodeFunctionCall(
    {
      name: "withdrawETH",
      type: "function",
      inputs: [
        {
          type: "uint256",
          name: "amount",
        },
      ],
    },
    [amount]
  );
};

const getBytesTransferERC20 = (to, amount) => {
  return web3.eth.abi.encodeFunctionCall(
    {
      name: "transfer",
      type: "function",
      inputs: [
        {
          type: "address",
          name: "to",
        },
        {
          type: "uint256",
          name: "amount",
        },
      ],
    },
    [to, amount]
  );
};

const getBytesSafeTransferERC721 = (from, to, tokenId) => {
  return web3.eth.abi.encodeFunctionCall(
    {
      name: "safeTransferFrom",
      type: "function",
      inputs: [
        {
          type: "address",
          name: "from",
        },
        {
          type: "address",
          name: "to",
        },
        {
          type: "uint256",
          name: "tokenId",
        },
      ],
    },
    [from, to, tokenId]
  );
};

const getBytesSafeTransferERC1155 = (from, to, id, amount, data) => {
  return web3.eth.abi.encodeFunctionCall(
    {
      name: "safeTransferFrom",
      type: "function",
      inputs: [
        {
          type: "address",
          name: "from",
        },
        {
          type: "address",
          name: "to",
        },
        {
          type: "uint256",
          name: "id",
        },
        {
          type: "uint256",
          name: "amount",
        },
        {
          type: "bytes",
          name: "data",
        },
      ],
    },
    [from, to, id, amount, data]
  );
};

module.exports = {
  getBytesDoSomething,
  getBytesWithdrawNFT,
  getBytesWithdrawETH,
  getBytesTransferERC20,
  getBytesSafeTransferERC721,
  getBytesSafeTransferERC1155,
};
