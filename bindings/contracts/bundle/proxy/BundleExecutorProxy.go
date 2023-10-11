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

// BundleExecutorProxyMetaData contains all meta data concerning the BundleExecutorProxy contract.
var BundleExecutorProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BundleExecutorProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use BundleExecutorProxyMetaData.ABI instead.
var BundleExecutorProxyABI = BundleExecutorProxyMetaData.ABI

// BundleExecutorProxy is an auto generated Go binding around an Ethereum contract.
type BundleExecutorProxy struct {
	BundleExecutorProxyCaller     // Read-only binding to the contract
	BundleExecutorProxyTransactor // Write-only binding to the contract
	BundleExecutorProxyFilterer   // Log filterer for contract events
}

// BundleExecutorProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type BundleExecutorProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundleExecutorProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BundleExecutorProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundleExecutorProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BundleExecutorProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundleExecutorProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BundleExecutorProxySession struct {
	Contract     *BundleExecutorProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BundleExecutorProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BundleExecutorProxyCallerSession struct {
	Contract *BundleExecutorProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BundleExecutorProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BundleExecutorProxyTransactorSession struct {
	Contract     *BundleExecutorProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BundleExecutorProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type BundleExecutorProxyRaw struct {
	Contract *BundleExecutorProxy // Generic contract binding to access the raw methods on
}

// BundleExecutorProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BundleExecutorProxyCallerRaw struct {
	Contract *BundleExecutorProxyCaller // Generic read-only contract binding to access the raw methods on
}

// BundleExecutorProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BundleExecutorProxyTransactorRaw struct {
	Contract *BundleExecutorProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBundleExecutorProxy creates a new instance of BundleExecutorProxy, bound to a specific deployed contract.
func NewBundleExecutorProxy(address common.Address, backend bind.ContractBackend) (*BundleExecutorProxy, error) {
	contract, err := bindBundleExecutorProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BundleExecutorProxy{BundleExecutorProxyCaller: BundleExecutorProxyCaller{contract: contract}, BundleExecutorProxyTransactor: BundleExecutorProxyTransactor{contract: contract}, BundleExecutorProxyFilterer: BundleExecutorProxyFilterer{contract: contract}}, nil
}

// NewBundleExecutorProxyCaller creates a new read-only instance of BundleExecutorProxy, bound to a specific deployed contract.
func NewBundleExecutorProxyCaller(address common.Address, caller bind.ContractCaller) (*BundleExecutorProxyCaller, error) {
	contract, err := bindBundleExecutorProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BundleExecutorProxyCaller{contract: contract}, nil
}

// NewBundleExecutorProxyTransactor creates a new write-only instance of BundleExecutorProxy, bound to a specific deployed contract.
func NewBundleExecutorProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*BundleExecutorProxyTransactor, error) {
	contract, err := bindBundleExecutorProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BundleExecutorProxyTransactor{contract: contract}, nil
}

// NewBundleExecutorProxyFilterer creates a new log filterer instance of BundleExecutorProxy, bound to a specific deployed contract.
func NewBundleExecutorProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*BundleExecutorProxyFilterer, error) {
	contract, err := bindBundleExecutorProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BundleExecutorProxyFilterer{contract: contract}, nil
}

// bindBundleExecutorProxy binds a generic wrapper to an already deployed contract.
func bindBundleExecutorProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BundleExecutorProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BundleExecutorProxy *BundleExecutorProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BundleExecutorProxy.Contract.BundleExecutorProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BundleExecutorProxy *BundleExecutorProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BundleExecutorProxy.Contract.BundleExecutorProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BundleExecutorProxy *BundleExecutorProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BundleExecutorProxy.Contract.BundleExecutorProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BundleExecutorProxy *BundleExecutorProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BundleExecutorProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BundleExecutorProxy *BundleExecutorProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BundleExecutorProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BundleExecutorProxy *BundleExecutorProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BundleExecutorProxy.Contract.contract.Transact(opts, method, params...)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_BundleExecutorProxy *BundleExecutorProxyTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BundleExecutorProxy.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_BundleExecutorProxy *BundleExecutorProxySession) Destroy() (*types.Transaction, error) {
	return _BundleExecutorProxy.Contract.Destroy(&_BundleExecutorProxy.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_BundleExecutorProxy *BundleExecutorProxyTransactorSession) Destroy() (*types.Transaction, error) {
	return _BundleExecutorProxy.Contract.Destroy(&_BundleExecutorProxy.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BundleExecutorProxy *BundleExecutorProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BundleExecutorProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BundleExecutorProxy *BundleExecutorProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BundleExecutorProxy.Contract.Fallback(&_BundleExecutorProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BundleExecutorProxy *BundleExecutorProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BundleExecutorProxy.Contract.Fallback(&_BundleExecutorProxy.TransactOpts, calldata)
}
