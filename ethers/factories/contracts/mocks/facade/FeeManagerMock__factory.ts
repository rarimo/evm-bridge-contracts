/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../common";
import type {
  FeeManagerMock,
  FeeManagerMockInterface,
} from "../../../../contracts/mocks/facade/FeeManagerMock";

const _abi = [
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "feeToken",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "feeAmount",
        type: "uint256",
      },
    ],
    name: "AddedFeeToken",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "previousAdmin",
        type: "address",
      },
      {
        indexed: false,
        internalType: "address",
        name: "newAdmin",
        type: "address",
      },
    ],
    name: "AdminChanged",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "beacon",
        type: "address",
      },
    ],
    name: "BeaconUpgraded",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint8",
        name: "version",
        type: "uint8",
      },
    ],
    name: "Initialized",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "feeToken",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "feeAmount",
        type: "uint256",
      },
    ],
    name: "RemovedFeeToken",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "feeToken",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "feeAmount",
        type: "uint256",
      },
    ],
    name: "UpdatedFeeToken",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "implementation",
        type: "address",
      },
    ],
    name: "Upgraded",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "receiver",
        type: "address",
      },
      {
        indexed: false,
        internalType: "address",
        name: "feeToken",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
    ],
    name: "WithdrawnFeeToken",
    type: "event",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "bridge_",
        type: "address",
      },
    ],
    name: "__FeeManagerMock_init",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "bridge_",
        type: "address",
      },
    ],
    name: "__FeeManager_init",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "address[]",
            name: "feeTokens",
            type: "address[]",
          },
          {
            internalType: "uint256[]",
            name: "feeAmounts",
            type: "uint256[]",
          },
          {
            internalType: "bytes",
            name: "signature",
            type: "bytes",
          },
        ],
        internalType: "struct IFeeManager.AddFeeTokenParameters",
        name: "params_",
        type: "tuple",
      },
    ],
    name: "addFeeToken",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "bridge",
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
        name: "feeToken_",
        type: "address",
      },
    ],
    name: "getCommission",
    outputs: [
      {
        internalType: "uint256",
        name: "commission_",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "proxiableUUID",
    outputs: [
      {
        internalType: "bytes32",
        name: "",
        type: "bytes32",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "address[]",
            name: "feeTokens",
            type: "address[]",
          },
          {
            internalType: "uint256[]",
            name: "feeAmounts",
            type: "uint256[]",
          },
          {
            internalType: "bytes",
            name: "signature",
            type: "bytes",
          },
        ],
        internalType: "struct IFeeManager.RemoveFeeTokenParameters",
        name: "params_",
        type: "tuple",
      },
    ],
    name: "removeFeeToken",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "address[]",
            name: "feeTokens",
            type: "address[]",
          },
          {
            internalType: "uint256[]",
            name: "feeAmounts",
            type: "uint256[]",
          },
          {
            internalType: "bytes",
            name: "signature",
            type: "bytes",
          },
        ],
        internalType: "struct IFeeManager.UpdateFeeTokenParameters",
        name: "params_",
        type: "tuple",
      },
    ],
    name: "updateFeeToken",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "newImplementation",
        type: "address",
      },
    ],
    name: "upgradeTo",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "newImplementation",
        type: "address",
      },
      {
        internalType: "bytes",
        name: "data",
        type: "bytes",
      },
    ],
    name: "upgradeToAndCall",
    outputs: [],
    stateMutability: "payable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "newImplementation_",
        type: "address",
      },
      {
        internalType: "bytes",
        name: "signature_",
        type: "bytes",
      },
    ],
    name: "upgradeToWithSig",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "address",
            name: "receiver",
            type: "address",
          },
          {
            internalType: "address[]",
            name: "feeTokens",
            type: "address[]",
          },
          {
            internalType: "uint256[]",
            name: "amounts",
            type: "uint256[]",
          },
          {
            internalType: "bytes",
            name: "signature",
            type: "bytes",
          },
        ],
        internalType: "struct IFeeManager.WithdrawFeeTokenParameters",
        name: "params_",
        type: "tuple",
      },
    ],
    name: "withdrawFeeToken",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
] as const;

const _bytecode =
  "0x60a06040523060805234801561001457600080fd5b50608051611fd561005a60003960008181610262015281816102ab0152818161050001528181610540015281816105b6015281816105f601526106740152611fd56000f3fe6080604052600436106100a75760003560e01c806363925ea21161006457806363925ea2146101695780639f33cf2214610189578063a77c8c86146101a9578063e78cea92146101c9578063f0653a2814610201578063f334dcb21461022157600080fd5b80633659cfe6146100ac5780634915b86e146100ce5780634eb5943a146100ee5780634f1ef2861461010e57806352d046611461012157806352d1902d14610141575b600080fd5b3480156100b857600080fd5b506100cc6100c736600461181a565b610257565b005b3480156100da57600080fd5b506100cc6100e936600461184d565b610346565b3480156100fa57600080fd5b506100cc61010936600461184d565b610432565b6100cc61011c3660046118f1565b6104f5565b34801561012d57600080fd5b506100cc61013c366004611982565b6105ab565b34801561014d57600080fd5b50610156610667565b6040519081526020015b60405180910390f35b34801561017557600080fd5b506100cc61018436600461184d565b61071a565b34801561019557600080fd5b506100cc6101a4366004611a05565b6107dd565b3480156101b557600080fd5b506100cc6101c436600461181a565b6108b5565b3480156101d557600080fd5b506065546101e9906001600160a01b031681565b6040516001600160a01b039091168152602001610160565b34801561020d57600080fd5b506100cc61021c36600461181a565b610942565b34801561022d57600080fd5b5061015661023c36600461181a565b6001600160a01b031660009081526066602052604090205490565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156102a95760405162461bcd60e51b81526004016102a090611a40565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166102db6109b9565b6001600160a01b0316146103015760405162461bcd60e51b81526004016102a090611a8c565b61030a816109d5565b6103438160005b6040519080825280601f01601f19166020018201604052801561033b576020820181803683370190505b506000610a2c565b50565b6103536020820182611ad8565b905061035f8280611ad8565b90501461037e5760405162461bcd60e51b81526004016102a090611b29565b6103ad60005b61038e8380611ad8565b61039b6020860186611ad8565b6103a86040880188611b6c565b610ba6565b60005b6103ba8280611ad8565b905081101561042e5761041e6103d08380611ad8565b838181106103e0576103e0611bb3565b90506020020160208101906103f5919061181a565b6104026020850185611ad8565b8481811061041257610412611bb3565b90506020020135610ce5565b61042781611bc9565b90506103b0565b5050565b61043f6020820182611ad8565b905061044b8280611ad8565b90501461046a5760405162461bcd60e51b81526004016102a090611b29565b6104746001610384565b60005b6104818280611ad8565b905081101561042e576104e56104978380611ad8565b838181106104a7576104a7611bb3565b90506020020160208101906104bc919061181a565b6104c96020850185611ad8565b848181106104d9576104d9611bb3565b90506020020135610d9c565b6104ee81611bc9565b9050610477565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016141561053e5760405162461bcd60e51b81526004016102a090611a40565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166105706109b9565b6001600160a01b0316146105965760405162461bcd60e51b81526004016102a090611a8c565b61059f826109d5565b61042e82826001610a2c565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156105f45760405162461bcd60e51b81526004016102a090611a40565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166106266109b9565b6001600160a01b03161461064c5760405162461bcd60e51b81526004016102a090611a8c565b610657838383610e66565b610662836000610311565b505050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107075760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c000000000000000060648201526084016102a0565b50600080516020611f5983398151915290565b6107276020820182611ad8565b90506107338280611ad8565b9050146107525760405162461bcd60e51b81526004016102a090611b29565b61075c6002610384565b60005b6107698280611ad8565b905081101561042e576107cd61077f8380611ad8565b8381811061078f5761078f611bb3565b90506020020160208101906107a4919061181a565b6107b16020850185611ad8565b848181106107c1576107c1611bb3565b90506020020135610f2a565b6107d681611bc9565b905061075f565b6107ea6040820182611ad8565b90506107f96020830183611ad8565b9050146108185760405162461bcd60e51b81526004016102a090611b29565b61082181610fe0565b60005b6108316020830183611ad8565b905081101561042e576108a561084a602084018461181a565b6108576020850185611ad8565b8481811061086757610867611bb3565b905060200201602081019061087c919061181a565b6108896040860186611ad8565b8581811061089957610899611bb3565b9050602002013561114d565b6108ae81611bc9565b9050610824565b600054610100900460ff166109205760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016102a0565b606580546001600160a01b0319166001600160a01b0392909216919091179055565b600061094e60016112cf565b90508015610966576000805461ff0019166101001790555b61096f826108b5565b801561042e576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15050565b600080516020611f59833981519152546001600160a01b031690565b60405162461bcd60e51b815260206004820152602660248201527f4665654d616e616765723a20746869732075706772616465206d6574686f642060448201526534b99037b33360d11b60648201526084016102a0565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610a5f576106628361135c565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610a9857600080fd5b505afa925050508015610ac8575060408051601f3d908101601f19168201909252610ac591810190611bf2565b60015b610b2b5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b60648201526084016102a0565b600080516020611f598339815191528114610b9a5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b60648201526084016102a0565b506106628383836113f8565b60655460405163413f04cf60e11b815260ff891660048201523060248201526001600160a01b03909116906000908190839063827e099e9060440160006040518083038186803b158015610bf957600080fd5b505afa158015610c0d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c359190810190611c37565b9150915060008130848d8d8d8d8d604051602001610c5a989796959493929190611d24565b60408051601f19818403018152908290528051602090910120630e3754f960e41b825291506001600160a01b0385169063e3754f9090610ca6908e90309086908c908c90600401611dbc565b600060405180830381600087803b158015610cc057600080fd5b505af1158015610cd4573d6000803e3d6000fd5b505050505050505050505050505050565b6001600160a01b03821660009081526066602052604090205415610d4b5760405162461bcd60e51b815260206004820181905260248201527f4665654d616e616765723a20746f6b656e20616c72656164792065786973747360448201526064016102a0565b6001600160a01b038216600081815260666020908152604091829020849055815192835282018390527f5d966ab5c51224ee3c5146e53a0b57969d97d0e1779f52a1017813c13344efc491016109ad565b6001600160a01b0382166000908152606660205260409020548114610e155760405162461bcd60e51b815260206004820152602960248201527f4665654d616e616765723a2077726f6e6720746f6b656e2061646472657373206044820152681bdc88185b5bdd5b9d60ba1b60648201526084016102a0565b6001600160a01b038216600081815260666020908152604080832092909255815192835282018390527feac220a493a0693e32f1240aae311de65b2d84391a9ba8310fcaf359d7c9bd8f91016109ad565b6001600160a01b038316610ebc5760405162461bcd60e51b815260206004820152601860248201527f4665654d616e616765723a207a65726f2061646472657373000000000000000060448201526064016102a0565b6065546001600160a01b0316637d1e764b6004308686866040518663ffffffff1660e01b8152600401610ef3959493929190611ded565b600060405180830381600087803b158015610f0d57600080fd5b505af1158015610f21573d6000803e3d6000fd5b50505050505050565b6001600160a01b038216600090815260666020526040902054610f8f5760405162461bcd60e51b815260206004820152601f60248201527f4665654d616e616765723a20746f6b656e20646f65736e27742065786973740060448201526064016102a0565b6001600160a01b038216600081815260666020908152604091829020849055815192835282018390527f70698787bb7456af2d4339ba1e7d059c4eda7c473412e7ed2c0d5b361a49d50a91016109ad565b60655460405163413f04cf60e11b81526003600482018190523060248301526001600160a01b0390921691906000908190849063827e099e9060440160006040518083038186803b15801561103457600080fd5b505afa158015611048573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526110709190810190611c37565b9092509050600081611085602088018861181a565b30858761109560208c018c611ad8565b6110a260408e018e611ad8565b6040516020016110ba99989796959493929190611e23565b60408051601f19818403018152919052805160209091012090506001600160a01b03851663e3754f908530846110f360608c018c611b6c565b6040518663ffffffff1660e01b8152600401611113959493929190611dbc565b600060405180830381600087803b15801561112d57600080fd5b505af1158015611141573d6000803e3d6000fd5b50505050505050505050565b6001600160a01b0382811614156111a65760405162461bcd60e51b815260206004820152601c60248201527f4665654d616e616765723a20636f6d6d697373696f6e20746f6b656e0000000060448201526064016102a0565b6001600160a01b03821661126b576000836001600160a01b03168260405160006040518083038185875af1925050503d8060008114611201576040519150601f19603f3d011682016040523d82523d6000602084013e611206565b606091505b50509050806112655760405162461bcd60e51b815260206004820152602560248201527f4665654d616e616765723a206661696c656420746f207769746864726177206e604482015264617469766560d81b60648201526084016102a0565b5061127f565b61127f6001600160a01b0383168483611423565b604080516001600160a01b038086168252841660208201529081018290527f0221f5dbeb176269bc9dbce8b10193c570930431bbe525ae79b4599910500bf49060600160405180910390a1505050565b60008054610100900460ff1615611316578160ff1660011480156112f25750303b155b61130e5760405162461bcd60e51b81526004016102a090611e99565b506000919050565b60005460ff80841691161061133d5760405162461bcd60e51b81526004016102a090611e99565b506000805460ff191660ff92909216919091179055600190565b919050565b6001600160a01b0381163b6113c95760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016102a0565b600080516020611f5983398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b61140183611475565b60008251118061140e5750805b156106625761141d83836114b5565b50505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526106629084906115a9565b61147e8161135c565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606001600160a01b0383163b61151d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016102a0565b600080846001600160a01b0316846040516115389190611ee7565b600060405180830381855af49150503d8060008114611573576040519150601f19603f3d011682016040523d82523d6000602084013e611578565b606091505b50915091506115a08282604051806060016040528060278152602001611f796027913961167b565b95945050505050565b60006115fe826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166116bb9092919063ffffffff16565b805190915015610662578080602001905181019061161c9190611f03565b6106625760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016102a0565b6060831561168a5750816116b4565b82511561169a5782518084602001fd5b8160405162461bcd60e51b81526004016102a09190611f25565b9392505050565b60606116ca84846000856116d2565b949350505050565b6060824710156117335760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016102a0565b6001600160a01b0385163b61178a5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102a0565b600080866001600160a01b031685876040516117a69190611ee7565b60006040518083038185875af1925050503d80600081146117e3576040519150601f19603f3d011682016040523d82523d6000602084013e6117e8565b606091505b50915091506117f882828661167b565b979650505050505050565b80356001600160a01b038116811461135757600080fd5b60006020828403121561182c57600080fd5b6116b482611803565b60006060828403121561184757600080fd5b50919050565b60006020828403121561185f57600080fd5b813567ffffffffffffffff81111561187657600080fd5b6116ca84828501611835565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156118c1576118c1611882565b604052919050565b600067ffffffffffffffff8211156118e3576118e3611882565b50601f01601f191660200190565b6000806040838503121561190457600080fd5b61190d83611803565b9150602083013567ffffffffffffffff81111561192957600080fd5b8301601f8101851361193a57600080fd5b803561194d611948826118c9565b611898565b81815286602083850101111561196257600080fd5b816020840160208301376000602083830101528093505050509250929050565b60008060006040848603121561199757600080fd5b6119a084611803565b9250602084013567ffffffffffffffff808211156119bd57600080fd5b818601915086601f8301126119d157600080fd5b8135818111156119e057600080fd5b8760208285010111156119f257600080fd5b6020830194508093505050509250925092565b600060208284031215611a1757600080fd5b813567ffffffffffffffff811115611a2e57600080fd5b8201608081850312156116b457600080fd5b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b6000808335601e19843603018112611aef57600080fd5b83018035915067ffffffffffffffff821115611b0a57600080fd5b6020019150600581901b3603821315611b2257600080fd5b9250929050565b60208082526023908201527f4665654d616e616765723a20706172616d73206c656e67746873206d69736d616040820152620e8c6d60eb1b606082015260800190565b6000808335601e19843603018112611b8357600080fd5b83018035915067ffffffffffffffff821115611b9e57600080fd5b602001915036819003821315611b2257600080fd5b634e487b7160e01b600052603260045260246000fd5b6000600019821415611beb57634e487b7160e01b600052601160045260246000fd5b5060010190565b600060208284031215611c0457600080fd5b5051919050565b60005b83811015611c26578181015183820152602001611c0e565b8381111561141d5750506000910152565b60008060408385031215611c4a57600080fd5b825167ffffffffffffffff811115611c6157600080fd5b8301601f81018513611c7257600080fd5b8051611c80611948826118c9565b818152866020838501011115611c9557600080fd5b611ca6826020830160208601611c0b565b60209590950151949694955050505050565b60008160005b84811015611ced576001600160a01b03611cd783611803565b1686526020958601959190910190600101611cbe565b5093949350505050565b60006001600160fb1b03831115611d0d57600080fd5b8260051b8083863760009401938452509192915050565b8881526bffffffffffffffffffffffff198860601b16602082015260008751611d54816034850160208c01611c0b565b60f888901b6001600160f81b031916603491840191820152611d84611d7d60358301888a611cb8565b8587611cf7565b9b9a5050505050505050505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60ff8616815260018060a01b03851660208201528360408201526080606082015260006117f8608083018486611d93565b60ff861681526001600160a01b038581166020830152841660408201526080606082018190526000906117f89083018486611d93565b89815260006bffffffffffffffffffffffff19808b60601b166020840152808a60601b166034840152508751611e60816048850160208c01611c0b565b60f888901b6001600160f81b031916604891840191820152611e89611d7d60498301888a611cb8565b9c9b505050505050505050505050565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b60008251611ef9818460208701611c0b565b9190910192915050565b600060208284031215611f1557600080fd5b815180151581146116b457600080fd5b6020815260008251806020840152611f44816040850160208701611c0b565b601f01601f1916919091016040019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212203acb89777ba672ec451a90267cb2794949936dba40754076d10e6353e1df6c4564736f6c63430008090033";

type FeeManagerMockConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: FeeManagerMockConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class FeeManagerMock__factory extends ContractFactory {
  constructor(...args: FeeManagerMockConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
    this.contractName = "FeeManagerMock";
  }

  override deploy(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<FeeManagerMock> {
    return super.deploy(overrides || {}) as Promise<FeeManagerMock>;
  }
  override getDeployTransaction(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  override attach(address: string): FeeManagerMock {
    return super.attach(address) as FeeManagerMock;
  }
  override connect(signer: Signer): FeeManagerMock__factory {
    return super.connect(signer) as FeeManagerMock__factory;
  }
  static readonly contractName: "FeeManagerMock";

  public readonly contractName: "FeeManagerMock";

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): FeeManagerMockInterface {
    return new utils.Interface(_abi) as FeeManagerMockInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): FeeManagerMock {
    return new Contract(address, _abi, signerOrProvider) as FeeManagerMock;
  }
}
