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

// ISBTHandlerDepositSBTParameters is an auto generated low-level Go binding around an user-defined struct.
type ISBTHandlerDepositSBTParameters struct {
	Token    common.Address
	TokenId  *big.Int
	Bundle   IBundlerBundle
	Network  string
	Receiver string
}

// ISBTHandlerWithdrawSBTParameters is an auto generated low-level Go binding around an user-defined struct.
type ISBTHandlerWithdrawSBTParameters struct {
	Token      common.Address
	TokenId    *big.Int
	TokenURI   string
	Bundle     IBundlerBundle
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
}

// SBTHandlerMetaData contains all meta data concerning the SBTHandler contract.
var SBTHandlerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"name\":\"DepositedSBT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"WithdrawnSBT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bundleExecutorImplementation_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"facade_\",\"type\":\"address\"}],\"name\":\"__Bundler_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bundleExecutorImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"internalType\":\"structISBTHandler.DepositSBTParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositSBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt_\",\"type\":\"bytes32\"}],\"name\":\"determineProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structISBTHandler.WithdrawSBTParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawSBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structISBTHandler.WithdrawSBTParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawSBTBundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SBTHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use SBTHandlerMetaData.ABI instead.
var SBTHandlerABI = SBTHandlerMetaData.ABI

// SBTHandler is an auto generated Go binding around an Ethereum contract.
type SBTHandler struct {
	SBTHandlerCaller     // Read-only binding to the contract
	SBTHandlerTransactor // Write-only binding to the contract
	SBTHandlerFilterer   // Log filterer for contract events
}

// SBTHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SBTHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SBTHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SBTHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SBTHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SBTHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SBTHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SBTHandlerSession struct {
	Contract     *SBTHandler       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SBTHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SBTHandlerCallerSession struct {
	Contract *SBTHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SBTHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SBTHandlerTransactorSession struct {
	Contract     *SBTHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SBTHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SBTHandlerRaw struct {
	Contract *SBTHandler // Generic contract binding to access the raw methods on
}

// SBTHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SBTHandlerCallerRaw struct {
	Contract *SBTHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// SBTHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SBTHandlerTransactorRaw struct {
	Contract *SBTHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSBTHandler creates a new instance of SBTHandler, bound to a specific deployed contract.
func NewSBTHandler(address common.Address, backend bind.ContractBackend) (*SBTHandler, error) {
	contract, err := bindSBTHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SBTHandler{SBTHandlerCaller: SBTHandlerCaller{contract: contract}, SBTHandlerTransactor: SBTHandlerTransactor{contract: contract}, SBTHandlerFilterer: SBTHandlerFilterer{contract: contract}}, nil
}

// NewSBTHandlerCaller creates a new read-only instance of SBTHandler, bound to a specific deployed contract.
func NewSBTHandlerCaller(address common.Address, caller bind.ContractCaller) (*SBTHandlerCaller, error) {
	contract, err := bindSBTHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SBTHandlerCaller{contract: contract}, nil
}

// NewSBTHandlerTransactor creates a new write-only instance of SBTHandler, bound to a specific deployed contract.
func NewSBTHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*SBTHandlerTransactor, error) {
	contract, err := bindSBTHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SBTHandlerTransactor{contract: contract}, nil
}

// NewSBTHandlerFilterer creates a new log filterer instance of SBTHandler, bound to a specific deployed contract.
func NewSBTHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*SBTHandlerFilterer, error) {
	contract, err := bindSBTHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SBTHandlerFilterer{contract: contract}, nil
}

// bindSBTHandler binds a generic wrapper to an already deployed contract.
func bindSBTHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SBTHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SBTHandler *SBTHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SBTHandler.Contract.SBTHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SBTHandler *SBTHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SBTHandler.Contract.SBTHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SBTHandler *SBTHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SBTHandler.Contract.SBTHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SBTHandler *SBTHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SBTHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SBTHandler *SBTHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SBTHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SBTHandler *SBTHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SBTHandler.Contract.contract.Transact(opts, method, params...)
}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_SBTHandler *SBTHandlerCaller) BundleExecutorImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SBTHandler.contract.Call(opts, &out, "bundleExecutorImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_SBTHandler *SBTHandlerSession) BundleExecutorImplementation() (common.Address, error) {
	return _SBTHandler.Contract.BundleExecutorImplementation(&_SBTHandler.CallOpts)
}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_SBTHandler *SBTHandlerCallerSession) BundleExecutorImplementation() (common.Address, error) {
	return _SBTHandler.Contract.BundleExecutorImplementation(&_SBTHandler.CallOpts)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_SBTHandler *SBTHandlerCaller) DetermineProxyAddress(opts *bind.CallOpts, salt_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _SBTHandler.contract.Call(opts, &out, "determineProxyAddress", salt_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_SBTHandler *SBTHandlerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _SBTHandler.Contract.DetermineProxyAddress(&_SBTHandler.CallOpts, salt_)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_SBTHandler *SBTHandlerCallerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _SBTHandler.Contract.DetermineProxyAddress(&_SBTHandler.CallOpts, salt_)
}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_SBTHandler *SBTHandlerCaller) Facade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SBTHandler.contract.Call(opts, &out, "facade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_SBTHandler *SBTHandlerSession) Facade() (common.Address, error) {
	return _SBTHandler.Contract.Facade(&_SBTHandler.CallOpts)
}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_SBTHandler *SBTHandlerCallerSession) Facade() (common.Address, error) {
	return _SBTHandler.Contract.Facade(&_SBTHandler.CallOpts)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x96de44c2.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_, address facade_) returns()
func (_SBTHandler *SBTHandlerTransactor) BundlerInit(opts *bind.TransactOpts, bundleExecutorImplementation_ common.Address, facade_ common.Address) (*types.Transaction, error) {
	return _SBTHandler.contract.Transact(opts, "__Bundler_init", bundleExecutorImplementation_, facade_)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x96de44c2.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_, address facade_) returns()
func (_SBTHandler *SBTHandlerSession) BundlerInit(bundleExecutorImplementation_ common.Address, facade_ common.Address) (*types.Transaction, error) {
	return _SBTHandler.Contract.BundlerInit(&_SBTHandler.TransactOpts, bundleExecutorImplementation_, facade_)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x96de44c2.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_, address facade_) returns()
func (_SBTHandler *SBTHandlerTransactorSession) BundlerInit(bundleExecutorImplementation_ common.Address, facade_ common.Address) (*types.Transaction, error) {
	return _SBTHandler.Contract.BundlerInit(&_SBTHandler.TransactOpts, bundleExecutorImplementation_, facade_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0x755f3823.
//
// Solidity: function depositSBT((address,uint256,(bytes32,bytes),string,string) params_) returns()
func (_SBTHandler *SBTHandlerTransactor) DepositSBT(opts *bind.TransactOpts, params_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _SBTHandler.contract.Transact(opts, "depositSBT", params_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0x755f3823.
//
// Solidity: function depositSBT((address,uint256,(bytes32,bytes),string,string) params_) returns()
func (_SBTHandler *SBTHandlerSession) DepositSBT(params_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _SBTHandler.Contract.DepositSBT(&_SBTHandler.TransactOpts, params_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0x755f3823.
//
// Solidity: function depositSBT((address,uint256,(bytes32,bytes),string,string) params_) returns()
func (_SBTHandler *SBTHandlerTransactorSession) DepositSBT(params_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _SBTHandler.Contract.DepositSBT(&_SBTHandler.TransactOpts, params_)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_SBTHandler *SBTHandlerTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _SBTHandler.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_SBTHandler *SBTHandlerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _SBTHandler.Contract.OnERC721Received(&_SBTHandler.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_SBTHandler *SBTHandlerTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _SBTHandler.Contract.OnERC721Received(&_SBTHandler.TransactOpts, arg0, arg1, arg2, arg3)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_SBTHandler *SBTHandlerTransactor) WithdrawSBT(opts *bind.TransactOpts, params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _SBTHandler.contract.Transact(opts, "withdrawSBT", params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_SBTHandler *SBTHandlerSession) WithdrawSBT(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _SBTHandler.Contract.WithdrawSBT(&_SBTHandler.TransactOpts, params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_SBTHandler *SBTHandlerTransactorSession) WithdrawSBT(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _SBTHandler.Contract.WithdrawSBT(&_SBTHandler.TransactOpts, params_)
}

// WithdrawSBTBundle is a paid mutator transaction binding the contract method 0x9a830d4c.
//
// Solidity: function withdrawSBTBundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_SBTHandler *SBTHandlerTransactor) WithdrawSBTBundle(opts *bind.TransactOpts, params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _SBTHandler.contract.Transact(opts, "withdrawSBTBundle", params_)
}

// WithdrawSBTBundle is a paid mutator transaction binding the contract method 0x9a830d4c.
//
// Solidity: function withdrawSBTBundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_SBTHandler *SBTHandlerSession) WithdrawSBTBundle(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _SBTHandler.Contract.WithdrawSBTBundle(&_SBTHandler.TransactOpts, params_)
}

// WithdrawSBTBundle is a paid mutator transaction binding the contract method 0x9a830d4c.
//
// Solidity: function withdrawSBTBundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_SBTHandler *SBTHandlerTransactorSession) WithdrawSBTBundle(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _SBTHandler.Contract.WithdrawSBTBundle(&_SBTHandler.TransactOpts, params_)
}

// SBTHandlerDepositedSBTIterator is returned from FilterDepositedSBT and is used to iterate over the raw logs and unpacked data for DepositedSBT events raised by the SBTHandler contract.
type SBTHandlerDepositedSBTIterator struct {
	Event *SBTHandlerDepositedSBT // Event containing the contract specifics and raw log

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
func (it *SBTHandlerDepositedSBTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SBTHandlerDepositedSBT)
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
		it.Event = new(SBTHandlerDepositedSBT)
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
func (it *SBTHandlerDepositedSBTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SBTHandlerDepositedSBTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SBTHandlerDepositedSBT represents a DepositedSBT event raised by the SBTHandler contract.
type SBTHandlerDepositedSBT struct {
	Token    common.Address
	TokenId  *big.Int
	Salt     [32]byte
	Bundle   []byte
	Network  string
	Receiver string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDepositedSBT is a free log retrieval operation binding the contract event 0x1452835f6476968a263680d4519d5a7d74a60354947a95bad57f7274c339ae86.
//
// Solidity: event DepositedSBT(address token, uint256 tokenId, bytes32 salt, bytes bundle, string network, string receiver)
func (_SBTHandler *SBTHandlerFilterer) FilterDepositedSBT(opts *bind.FilterOpts) (*SBTHandlerDepositedSBTIterator, error) {

	logs, sub, err := _SBTHandler.contract.FilterLogs(opts, "DepositedSBT")
	if err != nil {
		return nil, err
	}
	return &SBTHandlerDepositedSBTIterator{contract: _SBTHandler.contract, event: "DepositedSBT", logs: logs, sub: sub}, nil
}

// WatchDepositedSBT is a free log subscription operation binding the contract event 0x1452835f6476968a263680d4519d5a7d74a60354947a95bad57f7274c339ae86.
//
// Solidity: event DepositedSBT(address token, uint256 tokenId, bytes32 salt, bytes bundle, string network, string receiver)
func (_SBTHandler *SBTHandlerFilterer) WatchDepositedSBT(opts *bind.WatchOpts, sink chan<- *SBTHandlerDepositedSBT) (event.Subscription, error) {

	logs, sub, err := _SBTHandler.contract.WatchLogs(opts, "DepositedSBT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SBTHandlerDepositedSBT)
				if err := _SBTHandler.contract.UnpackLog(event, "DepositedSBT", log); err != nil {
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

// ParseDepositedSBT is a log parse operation binding the contract event 0x1452835f6476968a263680d4519d5a7d74a60354947a95bad57f7274c339ae86.
//
// Solidity: event DepositedSBT(address token, uint256 tokenId, bytes32 salt, bytes bundle, string network, string receiver)
func (_SBTHandler *SBTHandlerFilterer) ParseDepositedSBT(log types.Log) (*SBTHandlerDepositedSBT, error) {
	event := new(SBTHandlerDepositedSBT)
	if err := _SBTHandler.contract.UnpackLog(event, "DepositedSBT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SBTHandlerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SBTHandler contract.
type SBTHandlerInitializedIterator struct {
	Event *SBTHandlerInitialized // Event containing the contract specifics and raw log

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
func (it *SBTHandlerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SBTHandlerInitialized)
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
		it.Event = new(SBTHandlerInitialized)
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
func (it *SBTHandlerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SBTHandlerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SBTHandlerInitialized represents a Initialized event raised by the SBTHandler contract.
type SBTHandlerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SBTHandler *SBTHandlerFilterer) FilterInitialized(opts *bind.FilterOpts) (*SBTHandlerInitializedIterator, error) {

	logs, sub, err := _SBTHandler.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SBTHandlerInitializedIterator{contract: _SBTHandler.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SBTHandler *SBTHandlerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SBTHandlerInitialized) (event.Subscription, error) {

	logs, sub, err := _SBTHandler.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SBTHandlerInitialized)
				if err := _SBTHandler.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SBTHandler *SBTHandlerFilterer) ParseInitialized(log types.Log) (*SBTHandlerInitialized, error) {
	event := new(SBTHandlerInitialized)
	if err := _SBTHandler.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SBTHandlerWithdrawnSBTIterator is returned from FilterWithdrawnSBT and is used to iterate over the raw logs and unpacked data for WithdrawnSBT events raised by the SBTHandler contract.
type SBTHandlerWithdrawnSBTIterator struct {
	Event *SBTHandlerWithdrawnSBT // Event containing the contract specifics and raw log

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
func (it *SBTHandlerWithdrawnSBTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SBTHandlerWithdrawnSBT)
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
		it.Event = new(SBTHandlerWithdrawnSBT)
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
func (it *SBTHandlerWithdrawnSBTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SBTHandlerWithdrawnSBTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SBTHandlerWithdrawnSBT represents a WithdrawnSBT event raised by the SBTHandler contract.
type SBTHandlerWithdrawnSBT struct {
	Token      common.Address
	TokenId    *big.Int
	TokenURI   string
	Salt       [32]byte
	Bundle     []byte
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnSBT is a free log retrieval operation binding the contract event 0xbdb8f203981c7db82a7d157c20285399875b02508b52db10a634dd2a72c1b623.
//
// Solidity: event WithdrawnSBT(address token, uint256 tokenId, string tokenURI, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_SBTHandler *SBTHandlerFilterer) FilterWithdrawnSBT(opts *bind.FilterOpts) (*SBTHandlerWithdrawnSBTIterator, error) {

	logs, sub, err := _SBTHandler.contract.FilterLogs(opts, "WithdrawnSBT")
	if err != nil {
		return nil, err
	}
	return &SBTHandlerWithdrawnSBTIterator{contract: _SBTHandler.contract, event: "WithdrawnSBT", logs: logs, sub: sub}, nil
}

// WatchWithdrawnSBT is a free log subscription operation binding the contract event 0xbdb8f203981c7db82a7d157c20285399875b02508b52db10a634dd2a72c1b623.
//
// Solidity: event WithdrawnSBT(address token, uint256 tokenId, string tokenURI, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_SBTHandler *SBTHandlerFilterer) WatchWithdrawnSBT(opts *bind.WatchOpts, sink chan<- *SBTHandlerWithdrawnSBT) (event.Subscription, error) {

	logs, sub, err := _SBTHandler.contract.WatchLogs(opts, "WithdrawnSBT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SBTHandlerWithdrawnSBT)
				if err := _SBTHandler.contract.UnpackLog(event, "WithdrawnSBT", log); err != nil {
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

// ParseWithdrawnSBT is a log parse operation binding the contract event 0xbdb8f203981c7db82a7d157c20285399875b02508b52db10a634dd2a72c1b623.
//
// Solidity: event WithdrawnSBT(address token, uint256 tokenId, string tokenURI, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_SBTHandler *SBTHandlerFilterer) ParseWithdrawnSBT(log types.Log) (*SBTHandlerWithdrawnSBT, error) {
	event := new(SBTHandlerWithdrawnSBT)
	if err := _SBTHandler.contract.UnpackLog(event, "WithdrawnSBT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
