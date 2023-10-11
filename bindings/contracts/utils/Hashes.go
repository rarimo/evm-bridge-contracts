// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package utils

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

// HashesMetaData contains all meta data concerning the Hashes contract.
var HashesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// HashesABI is the input ABI used to generate the binding from.
// Deprecated: Use HashesMetaData.ABI instead.
var HashesABI = HashesMetaData.ABI

// Hashes is an auto generated Go binding around an Ethereum contract.
type Hashes struct {
	HashesCaller     // Read-only binding to the contract
	HashesTransactor // Write-only binding to the contract
	HashesFilterer   // Log filterer for contract events
}

// HashesCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashesSession struct {
	Contract     *Hashes           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashesCallerSession struct {
	Contract *HashesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HashesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashesTransactorSession struct {
	Contract     *HashesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashesRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashesRaw struct {
	Contract *Hashes // Generic contract binding to access the raw methods on
}

// HashesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashesCallerRaw struct {
	Contract *HashesCaller // Generic read-only contract binding to access the raw methods on
}

// HashesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashesTransactorRaw struct {
	Contract *HashesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashes creates a new instance of Hashes, bound to a specific deployed contract.
func NewHashes(address common.Address, backend bind.ContractBackend) (*Hashes, error) {
	contract, err := bindHashes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hashes{HashesCaller: HashesCaller{contract: contract}, HashesTransactor: HashesTransactor{contract: contract}, HashesFilterer: HashesFilterer{contract: contract}}, nil
}

// NewHashesCaller creates a new read-only instance of Hashes, bound to a specific deployed contract.
func NewHashesCaller(address common.Address, caller bind.ContractCaller) (*HashesCaller, error) {
	contract, err := bindHashes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashesCaller{contract: contract}, nil
}

// NewHashesTransactor creates a new write-only instance of Hashes, bound to a specific deployed contract.
func NewHashesTransactor(address common.Address, transactor bind.ContractTransactor) (*HashesTransactor, error) {
	contract, err := bindHashes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashesTransactor{contract: contract}, nil
}

// NewHashesFilterer creates a new log filterer instance of Hashes, bound to a specific deployed contract.
func NewHashesFilterer(address common.Address, filterer bind.ContractFilterer) (*HashesFilterer, error) {
	contract, err := bindHashes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashesFilterer{contract: contract}, nil
}

// bindHashes binds a generic wrapper to an already deployed contract.
func bindHashes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HashesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hashes *HashesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hashes.Contract.HashesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hashes *HashesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hashes.Contract.HashesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hashes *HashesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hashes.Contract.HashesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hashes *HashesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hashes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hashes *HashesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hashes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hashes *HashesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hashes.Contract.contract.Transact(opts, method, params...)
}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 ) view returns(bool)
func (_Hashes *HashesCaller) UsedHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Hashes.contract.Call(opts, &out, "usedHashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 ) view returns(bool)
func (_Hashes *HashesSession) UsedHashes(arg0 [32]byte) (bool, error) {
	return _Hashes.Contract.UsedHashes(&_Hashes.CallOpts, arg0)
}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 ) view returns(bool)
func (_Hashes *HashesCallerSession) UsedHashes(arg0 [32]byte) (bool, error) {
	return _Hashes.Contract.UsedHashes(&_Hashes.CallOpts, arg0)
}
