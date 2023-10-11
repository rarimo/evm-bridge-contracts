// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package handlers

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

// IBundlerBundle is an auto generated low-level Go binding around an user-defined struct.
type IBundlerBundle struct {
	Salt   [32]byte
	Bundle []byte
}

// IERC1155HandlerDepositERC1155Parameters is an auto generated low-level Go binding around an user-defined struct.
type IERC1155HandlerDepositERC1155Parameters struct {
	Token     common.Address
	TokenId   *big.Int
	Amount    *big.Int
	Bundle    IBundlerBundle
	Network   string
	Receiver  string
	IsWrapped bool
}

// IERC1155HandlerWithdrawERC1155Parameters is an auto generated low-level Go binding around an user-defined struct.
type IERC1155HandlerWithdrawERC1155Parameters struct {
	Token      common.Address
	TokenId    *big.Int
	TokenURI   string
	Amount     *big.Int
	Bundle     IBundlerBundle
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
	IsWrapped  bool
}

// ERC1155HandlerMetaData contains all meta data concerning the ERC1155Handler contract.
var ERC1155HandlerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"DepositedERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"WithdrawnERC1155\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bundleExecutorImplementation_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"facade_\",\"type\":\"address\"}],\"name\":\"__Bundler_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bundleExecutorImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.DepositERC1155Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt_\",\"type\":\"bytes32\"}],\"name\":\"determineProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.WithdrawERC1155Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.WithdrawERC1155Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC1155Bundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC1155HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1155HandlerMetaData.ABI instead.
var ERC1155HandlerABI = ERC1155HandlerMetaData.ABI

// ERC1155Handler is an auto generated Go binding around an Ethereum contract.
type ERC1155Handler struct {
	ERC1155HandlerCaller     // Read-only binding to the contract
	ERC1155HandlerTransactor // Write-only binding to the contract
	ERC1155HandlerFilterer   // Log filterer for contract events
}

// ERC1155HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1155HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1155HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1155HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1155HandlerSession struct {
	Contract     *ERC1155Handler   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1155HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1155HandlerCallerSession struct {
	Contract *ERC1155HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC1155HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1155HandlerTransactorSession struct {
	Contract     *ERC1155HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC1155HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1155HandlerRaw struct {
	Contract *ERC1155Handler // Generic contract binding to access the raw methods on
}

// ERC1155HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1155HandlerCallerRaw struct {
	Contract *ERC1155HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1155HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1155HandlerTransactorRaw struct {
	Contract *ERC1155HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155Handler creates a new instance of ERC1155Handler, bound to a specific deployed contract.
func NewERC1155Handler(address common.Address, backend bind.ContractBackend) (*ERC1155Handler, error) {
	contract, err := bindERC1155Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155Handler{ERC1155HandlerCaller: ERC1155HandlerCaller{contract: contract}, ERC1155HandlerTransactor: ERC1155HandlerTransactor{contract: contract}, ERC1155HandlerFilterer: ERC1155HandlerFilterer{contract: contract}}, nil
}

// NewERC1155HandlerCaller creates a new read-only instance of ERC1155Handler, bound to a specific deployed contract.
func NewERC1155HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC1155HandlerCaller, error) {
	contract, err := bindERC1155Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155HandlerCaller{contract: contract}, nil
}

// NewERC1155HandlerTransactor creates a new write-only instance of ERC1155Handler, bound to a specific deployed contract.
func NewERC1155HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1155HandlerTransactor, error) {
	contract, err := bindERC1155Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155HandlerTransactor{contract: contract}, nil
}

// NewERC1155HandlerFilterer creates a new log filterer instance of ERC1155Handler, bound to a specific deployed contract.
func NewERC1155HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1155HandlerFilterer, error) {
	contract, err := bindERC1155Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155HandlerFilterer{contract: contract}, nil
}

// bindERC1155Handler binds a generic wrapper to an already deployed contract.
func bindERC1155Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC1155HandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155Handler *ERC1155HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155Handler.Contract.ERC1155HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155Handler *ERC1155HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.ERC1155HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155Handler *ERC1155HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.ERC1155HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155Handler *ERC1155HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155Handler *ERC1155HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155Handler *ERC1155HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.contract.Transact(opts, method, params...)
}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_ERC1155Handler *ERC1155HandlerCaller) BundleExecutorImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC1155Handler.contract.Call(opts, &out, "bundleExecutorImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_ERC1155Handler *ERC1155HandlerSession) BundleExecutorImplementation() (common.Address, error) {
	return _ERC1155Handler.Contract.BundleExecutorImplementation(&_ERC1155Handler.CallOpts)
}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_ERC1155Handler *ERC1155HandlerCallerSession) BundleExecutorImplementation() (common.Address, error) {
	return _ERC1155Handler.Contract.BundleExecutorImplementation(&_ERC1155Handler.CallOpts)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_ERC1155Handler *ERC1155HandlerCaller) DetermineProxyAddress(opts *bind.CallOpts, salt_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC1155Handler.contract.Call(opts, &out, "determineProxyAddress", salt_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_ERC1155Handler *ERC1155HandlerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _ERC1155Handler.Contract.DetermineProxyAddress(&_ERC1155Handler.CallOpts, salt_)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_ERC1155Handler *ERC1155HandlerCallerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _ERC1155Handler.Contract.DetermineProxyAddress(&_ERC1155Handler.CallOpts, salt_)
}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_ERC1155Handler *ERC1155HandlerCaller) Facade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC1155Handler.contract.Call(opts, &out, "facade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_ERC1155Handler *ERC1155HandlerSession) Facade() (common.Address, error) {
	return _ERC1155Handler.Contract.Facade(&_ERC1155Handler.CallOpts)
}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_ERC1155Handler *ERC1155HandlerCallerSession) Facade() (common.Address, error) {
	return _ERC1155Handler.Contract.Facade(&_ERC1155Handler.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155Handler *ERC1155HandlerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC1155Handler.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155Handler *ERC1155HandlerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155Handler.Contract.SupportsInterface(&_ERC1155Handler.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155Handler *ERC1155HandlerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155Handler.Contract.SupportsInterface(&_ERC1155Handler.CallOpts, interfaceId)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x96de44c2.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_, address facade_) returns()
func (_ERC1155Handler *ERC1155HandlerTransactor) BundlerInit(opts *bind.TransactOpts, bundleExecutorImplementation_ common.Address, facade_ common.Address) (*types.Transaction, error) {
	return _ERC1155Handler.contract.Transact(opts, "__Bundler_init", bundleExecutorImplementation_, facade_)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x96de44c2.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_, address facade_) returns()
func (_ERC1155Handler *ERC1155HandlerSession) BundlerInit(bundleExecutorImplementation_ common.Address, facade_ common.Address) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.BundlerInit(&_ERC1155Handler.TransactOpts, bundleExecutorImplementation_, facade_)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x96de44c2.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_, address facade_) returns()
func (_ERC1155Handler *ERC1155HandlerTransactorSession) BundlerInit(bundleExecutorImplementation_ common.Address, facade_ common.Address) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.BundlerInit(&_ERC1155Handler.TransactOpts, bundleExecutorImplementation_, facade_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x969aef06.
//
// Solidity: function depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_ERC1155Handler *ERC1155HandlerTransactor) DepositERC1155(opts *bind.TransactOpts, params_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _ERC1155Handler.contract.Transact(opts, "depositERC1155", params_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x969aef06.
//
// Solidity: function depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_ERC1155Handler *ERC1155HandlerSession) DepositERC1155(params_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.DepositERC1155(&_ERC1155Handler.TransactOpts, params_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x969aef06.
//
// Solidity: function depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_ERC1155Handler *ERC1155HandlerTransactorSession) DepositERC1155(params_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.DepositERC1155(&_ERC1155Handler.TransactOpts, params_)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ERC1155Handler *ERC1155HandlerTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Handler.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ERC1155Handler *ERC1155HandlerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.OnERC1155BatchReceived(&_ERC1155Handler.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ERC1155Handler *ERC1155HandlerTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.OnERC1155BatchReceived(&_ERC1155Handler.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ERC1155Handler *ERC1155HandlerTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Handler.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ERC1155Handler *ERC1155HandlerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.OnERC1155Received(&_ERC1155Handler.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ERC1155Handler *ERC1155HandlerTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.OnERC1155Received(&_ERC1155Handler.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC1155Handler *ERC1155HandlerTransactor) WithdrawERC1155(opts *bind.TransactOpts, params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _ERC1155Handler.contract.Transact(opts, "withdrawERC1155", params_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC1155Handler *ERC1155HandlerSession) WithdrawERC1155(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.WithdrawERC1155(&_ERC1155Handler.TransactOpts, params_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC1155Handler *ERC1155HandlerTransactorSession) WithdrawERC1155(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.WithdrawERC1155(&_ERC1155Handler.TransactOpts, params_)
}

// WithdrawERC1155Bundle is a paid mutator transaction binding the contract method 0xb76da974.
//
// Solidity: function withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC1155Handler *ERC1155HandlerTransactor) WithdrawERC1155Bundle(opts *bind.TransactOpts, params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _ERC1155Handler.contract.Transact(opts, "withdrawERC1155Bundle", params_)
}

// WithdrawERC1155Bundle is a paid mutator transaction binding the contract method 0xb76da974.
//
// Solidity: function withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC1155Handler *ERC1155HandlerSession) WithdrawERC1155Bundle(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.WithdrawERC1155Bundle(&_ERC1155Handler.TransactOpts, params_)
}

// WithdrawERC1155Bundle is a paid mutator transaction binding the contract method 0xb76da974.
//
// Solidity: function withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC1155Handler *ERC1155HandlerTransactorSession) WithdrawERC1155Bundle(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _ERC1155Handler.Contract.WithdrawERC1155Bundle(&_ERC1155Handler.TransactOpts, params_)
}

// ERC1155HandlerDepositedERC1155Iterator is returned from FilterDepositedERC1155 and is used to iterate over the raw logs and unpacked data for DepositedERC1155 events raised by the ERC1155Handler contract.
type ERC1155HandlerDepositedERC1155Iterator struct {
	Event *ERC1155HandlerDepositedERC1155 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155HandlerDepositedERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155HandlerDepositedERC1155)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155HandlerDepositedERC1155)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155HandlerDepositedERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155HandlerDepositedERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155HandlerDepositedERC1155 represents a DepositedERC1155 event raised by the ERC1155Handler contract.
type ERC1155HandlerDepositedERC1155 struct {
	Token     common.Address
	TokenId   *big.Int
	Amount    *big.Int
	Salt      [32]byte
	Bundle    []byte
	Network   string
	Receiver  string
	IsWrapped bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositedERC1155 is a free log retrieval operation binding the contract event 0x103b790f2fa3a8676ff87c3620a55f0853d0e45128a8c7e9fadf29e17c51d07a.
//
// Solidity: event DepositedERC1155(address token, uint256 tokenId, uint256 amount, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_ERC1155Handler *ERC1155HandlerFilterer) FilterDepositedERC1155(opts *bind.FilterOpts) (*ERC1155HandlerDepositedERC1155Iterator, error) {

	logs, sub, err := _ERC1155Handler.contract.FilterLogs(opts, "DepositedERC1155")
	if err != nil {
		return nil, err
	}
	return &ERC1155HandlerDepositedERC1155Iterator{contract: _ERC1155Handler.contract, event: "DepositedERC1155", logs: logs, sub: sub}, nil
}

// WatchDepositedERC1155 is a free log subscription operation binding the contract event 0x103b790f2fa3a8676ff87c3620a55f0853d0e45128a8c7e9fadf29e17c51d07a.
//
// Solidity: event DepositedERC1155(address token, uint256 tokenId, uint256 amount, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_ERC1155Handler *ERC1155HandlerFilterer) WatchDepositedERC1155(opts *bind.WatchOpts, sink chan<- *ERC1155HandlerDepositedERC1155) (event.Subscription, error) {

	logs, sub, err := _ERC1155Handler.contract.WatchLogs(opts, "DepositedERC1155")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155HandlerDepositedERC1155)
				if err := _ERC1155Handler.contract.UnpackLog(event, "DepositedERC1155", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositedERC1155 is a log parse operation binding the contract event 0x103b790f2fa3a8676ff87c3620a55f0853d0e45128a8c7e9fadf29e17c51d07a.
//
// Solidity: event DepositedERC1155(address token, uint256 tokenId, uint256 amount, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_ERC1155Handler *ERC1155HandlerFilterer) ParseDepositedERC1155(log types.Log) (*ERC1155HandlerDepositedERC1155, error) {
	event := new(ERC1155HandlerDepositedERC1155)
	if err := _ERC1155Handler.contract.UnpackLog(event, "DepositedERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155HandlerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC1155Handler contract.
type ERC1155HandlerInitializedIterator struct {
	Event *ERC1155HandlerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155HandlerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155HandlerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155HandlerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155HandlerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155HandlerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155HandlerInitialized represents a Initialized event raised by the ERC1155Handler contract.
type ERC1155HandlerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1155Handler *ERC1155HandlerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC1155HandlerInitializedIterator, error) {

	logs, sub, err := _ERC1155Handler.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC1155HandlerInitializedIterator{contract: _ERC1155Handler.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1155Handler *ERC1155HandlerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC1155HandlerInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC1155Handler.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155HandlerInitialized)
				if err := _ERC1155Handler.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1155Handler *ERC1155HandlerFilterer) ParseInitialized(log types.Log) (*ERC1155HandlerInitialized, error) {
	event := new(ERC1155HandlerInitialized)
	if err := _ERC1155Handler.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155HandlerWithdrawnERC1155Iterator is returned from FilterWithdrawnERC1155 and is used to iterate over the raw logs and unpacked data for WithdrawnERC1155 events raised by the ERC1155Handler contract.
type ERC1155HandlerWithdrawnERC1155Iterator struct {
	Event *ERC1155HandlerWithdrawnERC1155 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155HandlerWithdrawnERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155HandlerWithdrawnERC1155)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155HandlerWithdrawnERC1155)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155HandlerWithdrawnERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155HandlerWithdrawnERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155HandlerWithdrawnERC1155 represents a WithdrawnERC1155 event raised by the ERC1155Handler contract.
type ERC1155HandlerWithdrawnERC1155 struct {
	Token      common.Address
	TokenId    *big.Int
	TokenURI   string
	Amount     *big.Int
	Salt       [32]byte
	Bundle     []byte
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
	IsWrapped  bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnERC1155 is a free log retrieval operation binding the contract event 0x30acf6d8c142717c265a7b8ab5be21a380a1ab05e2cd7ce3ca4fdd0c3260303e.
//
// Solidity: event WithdrawnERC1155(address token, uint256 tokenId, string tokenURI, uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_ERC1155Handler *ERC1155HandlerFilterer) FilterWithdrawnERC1155(opts *bind.FilterOpts) (*ERC1155HandlerWithdrawnERC1155Iterator, error) {

	logs, sub, err := _ERC1155Handler.contract.FilterLogs(opts, "WithdrawnERC1155")
	if err != nil {
		return nil, err
	}
	return &ERC1155HandlerWithdrawnERC1155Iterator{contract: _ERC1155Handler.contract, event: "WithdrawnERC1155", logs: logs, sub: sub}, nil
}

// WatchWithdrawnERC1155 is a free log subscription operation binding the contract event 0x30acf6d8c142717c265a7b8ab5be21a380a1ab05e2cd7ce3ca4fdd0c3260303e.
//
// Solidity: event WithdrawnERC1155(address token, uint256 tokenId, string tokenURI, uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_ERC1155Handler *ERC1155HandlerFilterer) WatchWithdrawnERC1155(opts *bind.WatchOpts, sink chan<- *ERC1155HandlerWithdrawnERC1155) (event.Subscription, error) {

	logs, sub, err := _ERC1155Handler.contract.WatchLogs(opts, "WithdrawnERC1155")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155HandlerWithdrawnERC1155)
				if err := _ERC1155Handler.contract.UnpackLog(event, "WithdrawnERC1155", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawnERC1155 is a log parse operation binding the contract event 0x30acf6d8c142717c265a7b8ab5be21a380a1ab05e2cd7ce3ca4fdd0c3260303e.
//
// Solidity: event WithdrawnERC1155(address token, uint256 tokenId, string tokenURI, uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_ERC1155Handler *ERC1155HandlerFilterer) ParseWithdrawnERC1155(log types.Log) (*ERC1155HandlerWithdrawnERC1155, error) {
	event := new(ERC1155HandlerWithdrawnERC1155)
	if err := _ERC1155Handler.contract.UnpackLog(event, "WithdrawnERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
