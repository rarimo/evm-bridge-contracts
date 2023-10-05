// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bundle

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

// IBundlerMetaData contains all meta data concerning the IBundler contract.
var IBundlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt_\",\"type\":\"bytes32\"}],\"name\":\"determineProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IBundlerABI is the input ABI used to generate the binding from.
// Deprecated: Use IBundlerMetaData.ABI instead.
var IBundlerABI = IBundlerMetaData.ABI

// IBundler is an auto generated Go binding around an Ethereum contract.
type IBundler struct {
	IBundlerCaller     // Read-only binding to the contract
	IBundlerTransactor // Write-only binding to the contract
	IBundlerFilterer   // Log filterer for contract events
}

// IBundlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBundlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBundlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBundlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBundlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBundlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBundlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBundlerSession struct {
	Contract     *IBundler         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBundlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBundlerCallerSession struct {
	Contract *IBundlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IBundlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBundlerTransactorSession struct {
	Contract     *IBundlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IBundlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBundlerRaw struct {
	Contract *IBundler // Generic contract binding to access the raw methods on
}

// IBundlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBundlerCallerRaw struct {
	Contract *IBundlerCaller // Generic read-only contract binding to access the raw methods on
}

// IBundlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBundlerTransactorRaw struct {
	Contract *IBundlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBundler creates a new instance of IBundler, bound to a specific deployed contract.
func NewIBundler(address common.Address, backend bind.ContractBackend) (*IBundler, error) {
	contract, err := bindIBundler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBundler{IBundlerCaller: IBundlerCaller{contract: contract}, IBundlerTransactor: IBundlerTransactor{contract: contract}, IBundlerFilterer: IBundlerFilterer{contract: contract}}, nil
}

// NewIBundlerCaller creates a new read-only instance of IBundler, bound to a specific deployed contract.
func NewIBundlerCaller(address common.Address, caller bind.ContractCaller) (*IBundlerCaller, error) {
	contract, err := bindIBundler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBundlerCaller{contract: contract}, nil
}

// NewIBundlerTransactor creates a new write-only instance of IBundler, bound to a specific deployed contract.
func NewIBundlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IBundlerTransactor, error) {
	contract, err := bindIBundler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBundlerTransactor{contract: contract}, nil
}

// NewIBundlerFilterer creates a new log filterer instance of IBundler, bound to a specific deployed contract.
func NewIBundlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IBundlerFilterer, error) {
	contract, err := bindIBundler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBundlerFilterer{contract: contract}, nil
}

// bindIBundler binds a generic wrapper to an already deployed contract.
func bindIBundler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBundlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBundler *IBundlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBundler.Contract.IBundlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBundler *IBundlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBundler.Contract.IBundlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBundler *IBundlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBundler.Contract.IBundlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBundler *IBundlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBundler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBundler *IBundlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBundler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBundler *IBundlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBundler.Contract.contract.Transact(opts, method, params...)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_IBundler *IBundlerCaller) DetermineProxyAddress(opts *bind.CallOpts, salt_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IBundler.contract.Call(opts, &out, "determineProxyAddress", salt_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_IBundler *IBundlerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _IBundler.Contract.DetermineProxyAddress(&_IBundler.CallOpts, salt_)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_IBundler *IBundlerCallerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _IBundler.Contract.DetermineProxyAddress(&_IBundler.CallOpts, salt_)
}
