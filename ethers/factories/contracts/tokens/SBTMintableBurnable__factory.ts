/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../common";
import type {
  SBTMintableBurnable,
  SBTMintableBurnableInterface,
} from "../../../contracts/tokens/SBTMintableBurnable";

const _abi = [
  {
    inputs: [
      {
        internalType: "string",
        name: "name_",
        type: "string",
      },
      {
        internalType: "string",
        name: "symbol_",
        type: "string",
      },
      {
        internalType: "address",
        name: "owner_",
        type: "address",
      },
      {
        internalType: "string",
        name: "baseURI_",
        type: "string",
      },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "owner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "approved",
        type: "address",
      },
      {
        indexed: true,
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "Approval",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "owner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "operator",
        type: "address",
      },
      {
        indexed: false,
        internalType: "bool",
        name: "approved",
        type: "bool",
      },
    ],
    name: "ApprovalForAll",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "Locked",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "previousOwner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "OwnershipTransferred",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "from",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "to",
        type: "address",
      },
      {
        indexed: true,
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "Transfer",
    type: "event",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "to",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "approve",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "receiver_",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "tokenId_",
        type: "uint256",
      },
      {
        internalType: "string",
        name: "tokenURI_",
        type: "string",
      },
    ],
    name: "attestTo",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "owner",
        type: "address",
      },
    ],
    name: "balanceOf",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "baseURI",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "burn",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "getApproved",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "owner",
        type: "address",
      },
      {
        internalType: "address",
        name: "operator",
        type: "address",
      },
    ],
    name: "isApprovedForAll",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "tokenId_",
        type: "uint256",
      },
    ],
    name: "locked",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "name",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "owner",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "ownerOf",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "renounceOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "from",
        type: "address",
      },
      {
        internalType: "address",
        name: "to",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "safeTransferFrom",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "from",
        type: "address",
      },
      {
        internalType: "address",
        name: "to",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
      {
        internalType: "bytes",
        name: "_data",
        type: "bytes",
      },
    ],
    name: "safeTransferFrom",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "operator",
        type: "address",
      },
      {
        internalType: "bool",
        name: "approved",
        type: "bool",
      },
    ],
    name: "setApprovalForAll",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "bytes4",
        name: "interfaceId_",
        type: "bytes4",
      },
    ],
    name: "supportsInterface",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "symbol",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "from_",
        type: "address",
      },
    ],
    name: "tokenIdOf",
    outputs: [
      {
        internalType: "uint256",
        name: "tokenId_",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "tokenURI",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "from",
        type: "address",
      },
      {
        internalType: "address",
        name: "to",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "transferFrom",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "transferOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
] as const;

const _bytecode =
  "0x60806040523480156200001157600080fd5b50604051620021cf380380620021cf833981016040819052620000349162000332565b838362000041336200009a565b815162000056906001906020850190620001bf565b5080516200006c906002906020840190620001bf565b505081516200008491506008906020840190620001bf565b506200009082620000ea565b5050505062000422565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000546001600160a01b031633146200014a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6001600160a01b038116620001b15760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840162000141565b620001bc816200009a565b50565b828054620001cd90620003e5565b90600052602060002090601f016020900481019282620001f157600085556200023c565b82601f106200020c57805160ff19168380011785556200023c565b828001600101855582156200023c579182015b828111156200023c5782518255916020019190600101906200021f565b506200024a9291506200024e565b5090565b5b808211156200024a57600081556001016200024f565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200028d57600080fd5b81516001600160401b0380821115620002aa57620002aa62000265565b604051601f8301601f19908116603f01168101908282118183101715620002d557620002d562000265565b81604052838152602092508683858801011115620002f257600080fd5b600091505b83821015620003165785820183015181830184015290820190620002f7565b83821115620003285760008385830101525b9695505050505050565b600080600080608085870312156200034957600080fd5b84516001600160401b03808211156200036157600080fd5b6200036f888389016200027b565b955060208701519150808211156200038657600080fd5b62000394888389016200027b565b604088015190955091506001600160a01b0382168214620003b457600080fd5b606087015191935080821115620003ca57600080fd5b50620003d9878288016200027b565b91505092959194509250565b600181811c90821680620003fa57607f821691505b602082108114156200041c57634e487b7160e01b600052602260045260246000fd5b50919050565b611d9d80620004326000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c8063715018a6116100b8578063b45a3c0e1161007c578063b45a3c0e1461026a578063b88d4fde1461027d578063c563174414610290578063c87b56dd146102a3578063e985e9c5146102b6578063f2fde38b146102f257600080fd5b8063715018a614610223578063773c02d41461022b5780638da5cb5b1461023e57806395d89b411461024f578063a22cb4651461025757600080fd5b806342842e0e116100ff57806342842e0e146101cc57806344df8e70146101df5780636352211e146101e75780636c0360eb146101fa57806370a082311461020257600080fd5b806301ffc9a71461013c57806306fdde0314610164578063081812fc14610179578063095ea7b3146101a457806323b872dd146101b9575b600080fd5b61014f61014a3660046117f6565b610305565b60405190151581526020015b60405180910390f35b61016c61034b565b60405161015b919061186b565b61018c61018736600461187e565b6103dd565b6040516001600160a01b03909116815260200161015b565b6101b76101b23660046118ae565b61046a565b005b6101b76101c73660046118d8565b610580565b6101b76101da3660046118d8565b6105b1565b6101b76105cc565b61018c6101f536600461187e565b6105f5565b61016c61066c565b610215610210366004611914565b6106fa565b60405190815260200161015b565b6101b7610781565b610215610239366004611914565b6107b7565b6000546001600160a01b031661018c565b61016c610839565b6101b761026536600461192f565b610848565b61014f61027836600461187e565b610857565b6101b761028b366004611981565b6108d5565b6101b761029e366004611a5d565b61090d565b61016c6102b136600461187e565b6109d5565b61014f6102c4366004611ae4565b6001600160a01b03918216600090815260066020908152604080832093909416825291909152205460ff1690565b6101b7610300366004611914565b610b3f565b60006001600160e01b03198216635a2d1e0760e11b148061033657506001600160e01b031982166307b404df60e51b145b80610345575061034582610bd7565b92915050565b60606001805461035a90611b17565b80601f016020809104026020016040519081016040528092919081815260200182805461038690611b17565b80156103d35780601f106103a8576101008083540402835291602001916103d3565b820191906000526020600020905b8154815290600101906020018083116103b657829003601f168201915b5050505050905090565b60006103e882610c27565b61044e5760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b60648201526084015b60405180910390fd5b506000908152600560205260409020546001600160a01b031690565b6000610475826105f5565b9050806001600160a01b0316836001600160a01b031614156104e35760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b6064820152608401610445565b336001600160a01b03821614806104ff57506104ff81336102c4565b6105715760405162461bcd60e51b815260206004820152603860248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760448201527f6e6572206e6f7220617070726f76656420666f7220616c6c00000000000000006064820152608401610445565b61057b8383610c44565b505050565b61058a3382610cb2565b6105a65760405162461bcd60e51b815260040161044590611b52565b61057b838383610d9b565b61057b838383604051806020016040528060008152506108d5565b60006105d7336107b7565b3360009081526009602052604081205590506105f281610f42565b50565b6000818152600360205260408120546001600160a01b0316806103455760405162461bcd60e51b815260206004820152602960248201527f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460448201526832b73a103a37b5b2b760b91b6064820152608401610445565b6008805461067990611b17565b80601f01602080910402602001604051908101604052809291908181526020018280546106a590611b17565b80156106f25780601f106106c7576101008083540402835291602001916106f2565b820191906000526020600020905b8154815290600101906020018083116106d557829003601f168201915b505050505081565b60006001600160a01b0382166107655760405162461bcd60e51b815260206004820152602a60248201527f4552433732313a2062616c616e636520717565727920666f7220746865207a65604482015269726f206164647265737360b01b6064820152608401610445565b506001600160a01b031660009081526004602052604090205490565b6000546001600160a01b031633146107ab5760405162461bcd60e51b815260040161044590611ba3565b6107b56000610f82565b565b6001600160a01b038116600090815260096020526040902054806108345760405162461bcd60e51b815260206004820152602e60248201527f5342544d696e7461626c654275726e61626c653a207573657220646f65736e2760448201526d3a103430bb329030903a37b5b2b760911b6064820152608401610445565b919050565b60606002805461035a90611b17565b610853338383610fd2565b5050565b600061086282610c27565b6108cd5760405162461bcd60e51b815260206004820152603660248201527f5342544d696e7461626c654275726e61626c653a20746f6b656e2069732061736044820152757369676e656420746f207a65726f206164647265737360501b6064820152608401610445565b506001919050565b6108df3383610cb2565b6108fb5760405162461bcd60e51b815260040161044590611b52565b610907848484846110a1565b50505050565b6000546001600160a01b031633146109375760405162461bcd60e51b815260040161044590611ba3565b6001600160a01b038416600090815260096020526040902083905561095c84846110d4565b61099c8383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061121392505050565b6040518381527f032bc66be43dbccb7487781d168eb7bda224628a3b2c3388bdf69b532a3a16119060200160405180910390a150505050565b60606109e082610c27565b610a465760405162461bcd60e51b815260206004820152603160248201527f45524337323155524953746f726167653a2055524920717565727920666f72206044820152703737b732bc34b9ba32b73a103a37b5b2b760791b6064820152608401610445565b60008281526007602052604081208054610a5f90611b17565b80601f0160208091040260200160405190810160405280929190818152602001828054610a8b90611b17565b8015610ad85780601f10610aad57610100808354040283529160200191610ad8565b820191906000526020600020905b815481529060010190602001808311610abb57829003601f168201915b505050505090506000610ae961129e565b9050805160001415610afc575092915050565b815115610b2e578082604051602001610b16929190611bd8565b60405160208183030381529060405292505050919050565b610b37846112ad565b949350505050565b6000546001600160a01b03163314610b695760405162461bcd60e51b815260040161044590611ba3565b6001600160a01b038116610bce5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610445565b6105f281610f82565b60006001600160e01b031982166380ac58cd60e01b1480610c0857506001600160e01b03198216635b5e139f60e01b145b8061034557506301ffc9a760e01b6001600160e01b0319831614610345565b6000908152600360205260409020546001600160a01b0316151590565b600081815260056020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610c79826105f5565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000610cbd82610c27565b610d1e5760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b6064820152608401610445565b6000610d29836105f5565b9050806001600160a01b0316846001600160a01b03161480610d7057506001600160a01b0380821660009081526006602090815260408083209388168352929052205460ff165b80610b375750836001600160a01b0316610d89846103dd565b6001600160a01b031614949350505050565b826001600160a01b0316610dae826105f5565b6001600160a01b031614610e125760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b6064820152608401610445565b6001600160a01b038216610e745760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610445565b610e7f838383611378565b610e8a600082610c44565b6001600160a01b0383166000908152600460205260408120805460019290610eb3908490611c1d565b90915550506001600160a01b0382166000908152600460205260408120805460019290610ee1908490611c34565b909155505060008181526003602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b610f4b8161145b565b60008181526007602052604090208054610f6490611b17565b1590506105f25760008181526007602052604081206105f29161170d565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b816001600160a01b0316836001600160a01b031614156110345760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610445565b6001600160a01b03838116600081815260066020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6110ac848484610d9b565b6110b884848484611502565b6109075760405162461bcd60e51b815260040161044590611c4c565b6001600160a01b03821661112a5760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610445565b61113381610c27565b156111805760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610445565b61118c60008383611378565b6001600160a01b03821660009081526004602052604081208054600192906111b5908490611c34565b909155505060008181526003602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b61121c82610c27565b61127f5760405162461bcd60e51b815260206004820152602e60248201527f45524337323155524953746f726167653a2055524920736574206f66206e6f6e60448201526d32bc34b9ba32b73a103a37b5b2b760911b6064820152608401610445565b6000828152600760209081526040909120825161057b92840190611747565b60606008805461035a90611b17565b60606112b882610c27565b61131c5760405162461bcd60e51b815260206004820152602f60248201527f4552433732314d657461646174613a2055524920717565727920666f72206e6f60448201526e3732bc34b9ba32b73a103a37b5b2b760891b6064820152608401610445565b600061132661129e565b905060008151116113465760405180602001604052806000815250611371565b806113508461160f565b604051602001611361929190611bd8565b6040516020818303038152906040525b9392505050565b6001600160a01b0383166113f65761138f826106fa565b1561057b5760405162461bcd60e51b815260206004820152603160248201527f5342544d696e7461626c654275726e61626c653a20726563656976657220616c6044820152703932b0b23c903430b99030903a37b5b2b760791b6064820152608401610445565b6001600160a01b0382161561057b5760405162461bcd60e51b815260206004820152602560248201527f5342544d696e7461626c654275726e61626c653a206e6f74207472616e7366656044820152647261626c6560d81b6064820152608401610445565b6000611466826105f5565b905061147481600084611378565b61147f600083610c44565b6001600160a01b03811660009081526004602052604081208054600192906114a8908490611c1d565b909155505060008281526003602052604080822080546001600160a01b0319169055518391906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b60006001600160a01b0384163b1561160457604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290611546903390899088908890600401611c9e565b602060405180830381600087803b15801561156057600080fd5b505af1925050508015611590575060408051601f3d908101601f1916820190925261158d91810190611cdb565b60015b6115ea573d8080156115be576040519150601f19603f3d011682016040523d82523d6000602084013e6115c3565b606091505b5080516115e25760405162461bcd60e51b815260040161044590611c4c565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610b37565b506001949350505050565b6060816116335750506040805180820190915260018152600360fc1b602082015290565b8160005b811561165d578061164781611cf8565b91506116569050600a83611d29565b9150611637565b60008167ffffffffffffffff8111156116785761167861196b565b6040519080825280601f01601f1916602001820160405280156116a2576020820181803683370190505b5090505b8415610b37576116b7600183611c1d565b91506116c4600a86611d3d565b6116cf906030611c34565b60f81b8183815181106116e4576116e4611d51565b60200101906001600160f81b031916908160001a905350611706600a86611d29565b94506116a6565b50805461171990611b17565b6000825580601f10611729575050565b601f0160209004906000526020600020908101906105f291906117cb565b82805461175390611b17565b90600052602060002090601f01602090048101928261177557600085556117bb565b82601f1061178e57805160ff19168380011785556117bb565b828001600101855582156117bb579182015b828111156117bb5782518255916020019190600101906117a0565b506117c79291506117cb565b5090565b5b808211156117c757600081556001016117cc565b6001600160e01b0319811681146105f257600080fd5b60006020828403121561180857600080fd5b8135611371816117e0565b60005b8381101561182e578181015183820152602001611816565b838111156109075750506000910152565b60008151808452611857816020860160208601611813565b601f01601f19169290920160200192915050565b602081526000611371602083018461183f565b60006020828403121561189057600080fd5b5035919050565b80356001600160a01b038116811461083457600080fd5b600080604083850312156118c157600080fd5b6118ca83611897565b946020939093013593505050565b6000806000606084860312156118ed57600080fd5b6118f684611897565b925061190460208501611897565b9150604084013590509250925092565b60006020828403121561192657600080fd5b61137182611897565b6000806040838503121561194257600080fd5b61194b83611897565b91506020830135801515811461196057600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b6000806000806080858703121561199757600080fd5b6119a085611897565b93506119ae60208601611897565b925060408501359150606085013567ffffffffffffffff808211156119d257600080fd5b818701915087601f8301126119e657600080fd5b8135818111156119f8576119f861196b565b604051601f8201601f19908116603f01168101908382118183101715611a2057611a2061196b565b816040528281528a6020848701011115611a3957600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b60008060008060608587031215611a7357600080fd5b611a7c85611897565b935060208501359250604085013567ffffffffffffffff80821115611aa057600080fd5b818701915087601f830112611ab457600080fd5b813581811115611ac357600080fd5b886020828501011115611ad557600080fd5b95989497505060200194505050565b60008060408385031215611af757600080fd5b611b0083611897565b9150611b0e60208401611897565b90509250929050565b600181811c90821680611b2b57607f821691505b60208210811415611b4c57634e487b7160e01b600052602260045260246000fd5b50919050565b60208082526031908201527f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f6040820152701ddb995c881b9bdc88185c1c1c9bdd9959607a1b606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60008351611bea818460208801611813565b835190830190611bfe818360208801611813565b01949350505050565b634e487b7160e01b600052601160045260246000fd5b600082821015611c2f57611c2f611c07565b500390565b60008219821115611c4757611c47611c07565b500190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090611cd19083018461183f565b9695505050505050565b600060208284031215611ced57600080fd5b8151611371816117e0565b6000600019821415611d0c57611d0c611c07565b5060010190565b634e487b7160e01b600052601260045260246000fd5b600082611d3857611d38611d13565b500490565b600082611d4c57611d4c611d13565b500690565b634e487b7160e01b600052603260045260246000fdfea2646970667358221220a020ff9dc5c4d40021c0347146fb43672d0d200726a3bb78524391dc8d5e9c3b64736f6c63430008090033";

type SBTMintableBurnableConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: SBTMintableBurnableConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class SBTMintableBurnable__factory extends ContractFactory {
  constructor(...args: SBTMintableBurnableConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
    this.contractName = "SBTMintableBurnable";
  }

  override deploy(
    name_: PromiseOrValue<string>,
    symbol_: PromiseOrValue<string>,
    owner_: PromiseOrValue<string>,
    baseURI_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<SBTMintableBurnable> {
    return super.deploy(
      name_,
      symbol_,
      owner_,
      baseURI_,
      overrides || {}
    ) as Promise<SBTMintableBurnable>;
  }
  override getDeployTransaction(
    name_: PromiseOrValue<string>,
    symbol_: PromiseOrValue<string>,
    owner_: PromiseOrValue<string>,
    baseURI_: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(
      name_,
      symbol_,
      owner_,
      baseURI_,
      overrides || {}
    );
  }
  override attach(address: string): SBTMintableBurnable {
    return super.attach(address) as SBTMintableBurnable;
  }
  override connect(signer: Signer): SBTMintableBurnable__factory {
    return super.connect(signer) as SBTMintableBurnable__factory;
  }
  static readonly contractName: "SBTMintableBurnable";

  public readonly contractName: "SBTMintableBurnable";

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): SBTMintableBurnableInterface {
    return new utils.Interface(_abi) as SBTMintableBurnableInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): SBTMintableBurnable {
    return new Contract(address, _abi, signerOrProvider) as SBTMintableBurnable;
  }
}
