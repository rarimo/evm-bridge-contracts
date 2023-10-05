// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxy

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BundleExecutorImplementationMetaData contains all meta data concerning the BundleExecutorImplementation contract.
var BundleExecutorImplementationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedNative_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WRAPPED_NATIVE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"bundle_\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BundleExecutorImplementationABI is the input ABI used to generate the binding from.
// Deprecated: Use BundleExecutorImplementationMetaData.ABI instead.
var BundleExecutorImplementationABI = BundleExecutorImplementationMetaData.ABI

// BundleExecutorImplementation is an auto generated Go binding around an Ethereum contract.
type BundleExecutorImplementation struct {
	BundleExecutorImplementationCaller     // Read-only binding to the contract
	BundleExecutorImplementationTransactor // Write-only binding to the contract
	BundleExecutorImplementationFilterer   // Log filterer for contract events
}

// BundleExecutorImplementationCaller is an auto generated read-only Go binding around an Ethereum contract.
type BundleExecutorImplementationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundleExecutorImplementationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BundleExecutorImplementationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundleExecutorImplementationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BundleExecutorImplementationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundleExecutorImplementationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BundleExecutorImplementationSession struct {
	Contract     *BundleExecutorImplementation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BundleExecutorImplementationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BundleExecutorImplementationCallerSession struct {
	Contract *BundleExecutorImplementationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// BundleExecutorImplementationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BundleExecutorImplementationTransactorSession struct {
	Contract     *BundleExecutorImplementationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// BundleExecutorImplementationRaw is an auto generated low-level Go binding around an Ethereum contract.
type BundleExecutorImplementationRaw struct {
	Contract *BundleExecutorImplementation // Generic contract binding to access the raw methods on
}

// BundleExecutorImplementationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BundleExecutorImplementationCallerRaw struct {
	Contract *BundleExecutorImplementationCaller // Generic read-only contract binding to access the raw methods on
}

// BundleExecutorImplementationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BundleExecutorImplementationTransactorRaw struct {
	Contract *BundleExecutorImplementationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBundleExecutorImplementation creates a new instance of BundleExecutorImplementation, bound to a specific deployed contract.
func NewBundleExecutorImplementation(address common.Address, backend bind.ContractBackend) (*BundleExecutorImplementation, error) {
	contract, err := bindBundleExecutorImplementation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BundleExecutorImplementation{BundleExecutorImplementationCaller: BundleExecutorImplementationCaller{contract: contract}, BundleExecutorImplementationTransactor: BundleExecutorImplementationTransactor{contract: contract}, BundleExecutorImplementationFilterer: BundleExecutorImplementationFilterer{contract: contract}}, nil
}

// NewBundleExecutorImplementationCaller creates a new read-only instance of BundleExecutorImplementation, bound to a specific deployed contract.
func NewBundleExecutorImplementationCaller(address common.Address, caller bind.ContractCaller) (*BundleExecutorImplementationCaller, error) {
	contract, err := bindBundleExecutorImplementation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BundleExecutorImplementationCaller{contract: contract}, nil
}

// NewBundleExecutorImplementationTransactor creates a new write-only instance of BundleExecutorImplementation, bound to a specific deployed contract.
func NewBundleExecutorImplementationTransactor(address common.Address, transactor bind.ContractTransactor) (*BundleExecutorImplementationTransactor, error) {
	contract, err := bindBundleExecutorImplementation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BundleExecutorImplementationTransactor{contract: contract}, nil
}

// NewBundleExecutorImplementationFilterer creates a new log filterer instance of BundleExecutorImplementation, bound to a specific deployed contract.
func NewBundleExecutorImplementationFilterer(address common.Address, filterer bind.ContractFilterer) (*BundleExecutorImplementationFilterer, error) {
	contract, err := bindBundleExecutorImplementation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BundleExecutorImplementationFilterer{contract: contract}, nil
}

// bindBundleExecutorImplementation binds a generic wrapper to an already deployed contract.
func bindBundleExecutorImplementation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BundleExecutorImplementationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BundleExecutorImplementation *BundleExecutorImplementationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BundleExecutorImplementation.Contract.BundleExecutorImplementationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BundleExecutorImplementation *BundleExecutorImplementationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BundleExecutorImplementation.Contract.BundleExecutorImplementationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BundleExecutorImplementation *BundleExecutorImplementationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BundleExecutorImplementation.Contract.BundleExecutorImplementationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BundleExecutorImplementation *BundleExecutorImplementationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BundleExecutorImplementation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BundleExecutorImplementation *BundleExecutorImplementationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BundleExecutorImplementation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BundleExecutorImplementation *BundleExecutorImplementationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BundleExecutorImplementation.Contract.contract.Transact(opts, method, params...)
}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_BundleExecutorImplementation *BundleExecutorImplementationCaller) WRAPPEDNATIVE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BundleExecutorImplementation.contract.Call(opts, &out, "WRAPPED_NATIVE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_BundleExecutorImplementation *BundleExecutorImplementationSession) WRAPPEDNATIVE() (common.Address, error) {
	return _BundleExecutorImplementation.Contract.WRAPPEDNATIVE(&_BundleExecutorImplementation.CallOpts)
}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_BundleExecutorImplementation *BundleExecutorImplementationCallerSession) WRAPPEDNATIVE() (common.Address, error) {
	return _BundleExecutorImplementation.Contract.WRAPPEDNATIVE(&_BundleExecutorImplementation.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BundleExecutorImplementation *BundleExecutorImplementationCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BundleExecutorImplementation.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BundleExecutorImplementation *BundleExecutorImplementationSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BundleExecutorImplementation.Contract.SupportsInterface(&_BundleExecutorImplementation.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BundleExecutorImplementation *BundleExecutorImplementationCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BundleExecutorImplementation.Contract.SupportsInterface(&_BundleExecutorImplementation.CallOpts, interfaceId)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes bundle_) payable returns()
func (_BundleExecutorImplementation *BundleExecutorImplementationTransactor) Execute(opts *bind.TransactOpts, bundle_ []byte) (*types.Transaction, error) {
	return _BundleExecutorImplementation.contract.Transact(opts, "execute", bundle_)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes bundle_) payable returns()
func (_BundleExecutorImplementation *BundleExecutorImplementationSession) Execute(bundle_ []byte) (*types.Transaction, error) {
	return _BundleExecutorImplementation.Contract.Execute(&_BundleExecutorImplementation.TransactOpts, bundle_)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes bundle_) payable returns()
func (_BundleExecutorImplementation *BundleExecutorImplementationTransactorSession) Execute(bundle_ []byte) (*types.Transaction, error) {
	return _BundleExecutorImplementation.Contract.Execute(&_BundleExecutorImplementation.TransactOpts, bundle_)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_BundleExecutorImplementation *BundleExecutorImplementationTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BundleExecutorImplementation.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_BundleExecutorImplementation *BundleExecutorImplementationSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BundleExecutorImplementation.Contract.OnERC1155BatchReceived(&_BundleExecutorImplementation.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_BundleExecutorImplementation *BundleExecutorImplementationTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BundleExecutorImplementation.Contract.OnERC1155BatchReceived(&_BundleExecutorImplementation.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_BundleExecutorImplementation *BundleExecutorImplementationTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BundleExecutorImplementation.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_BundleExecutorImplementation *BundleExecutorImplementationSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BundleExecutorImplementation.Contract.OnERC1155Received(&_BundleExecutorImplementation.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_BundleExecutorImplementation *BundleExecutorImplementationTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BundleExecutorImplementation.Contract.OnERC1155Received(&_BundleExecutorImplementation.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_BundleExecutorImplementation *BundleExecutorImplementationTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _BundleExecutorImplementation.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_BundleExecutorImplementation *BundleExecutorImplementationSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _BundleExecutorImplementation.Contract.OnERC721Received(&_BundleExecutorImplementation.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_BundleExecutorImplementation *BundleExecutorImplementationTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _BundleExecutorImplementation.Contract.OnERC721Received(&_BundleExecutorImplementation.TransactOpts, arg0, arg1, arg2, arg3)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BundleExecutorImplementation *BundleExecutorImplementationTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BundleExecutorImplementation.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BundleExecutorImplementation *BundleExecutorImplementationSession) Receive() (*types.Transaction, error) {
	return _BundleExecutorImplementation.Contract.Receive(&_BundleExecutorImplementation.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BundleExecutorImplementation *BundleExecutorImplementationTransactorSession) Receive() (*types.Transaction, error) {
	return _BundleExecutorImplementation.Contract.Receive(&_BundleExecutorImplementation.TransactOpts)
}
