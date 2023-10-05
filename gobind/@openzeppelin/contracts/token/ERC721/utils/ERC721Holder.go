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

// ERC721HolderMetaData contains all meta data concerning the ERC721Holder contract.
var ERC721HolderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC721HolderABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721HolderMetaData.ABI instead.
var ERC721HolderABI = ERC721HolderMetaData.ABI

// ERC721Holder is an auto generated Go binding around an Ethereum contract.
type ERC721Holder struct {
	ERC721HolderCaller     // Read-only binding to the contract
	ERC721HolderTransactor // Write-only binding to the contract
	ERC721HolderFilterer   // Log filterer for contract events
}

// ERC721HolderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721HolderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HolderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721HolderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HolderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721HolderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HolderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721HolderSession struct {
	Contract     *ERC721Holder     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721HolderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721HolderCallerSession struct {
	Contract *ERC721HolderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC721HolderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721HolderTransactorSession struct {
	Contract     *ERC721HolderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC721HolderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721HolderRaw struct {
	Contract *ERC721Holder // Generic contract binding to access the raw methods on
}

// ERC721HolderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721HolderCallerRaw struct {
	Contract *ERC721HolderCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721HolderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721HolderTransactorRaw struct {
	Contract *ERC721HolderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Holder creates a new instance of ERC721Holder, bound to a specific deployed contract.
func NewERC721Holder(address common.Address, backend bind.ContractBackend) (*ERC721Holder, error) {
	contract, err := bindERC721Holder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Holder{ERC721HolderCaller: ERC721HolderCaller{contract: contract}, ERC721HolderTransactor: ERC721HolderTransactor{contract: contract}, ERC721HolderFilterer: ERC721HolderFilterer{contract: contract}}, nil
}

// NewERC721HolderCaller creates a new read-only instance of ERC721Holder, bound to a specific deployed contract.
func NewERC721HolderCaller(address common.Address, caller bind.ContractCaller) (*ERC721HolderCaller, error) {
	contract, err := bindERC721Holder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721HolderCaller{contract: contract}, nil
}

// NewERC721HolderTransactor creates a new write-only instance of ERC721Holder, bound to a specific deployed contract.
func NewERC721HolderTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721HolderTransactor, error) {
	contract, err := bindERC721Holder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721HolderTransactor{contract: contract}, nil
}

// NewERC721HolderFilterer creates a new log filterer instance of ERC721Holder, bound to a specific deployed contract.
func NewERC721HolderFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721HolderFilterer, error) {
	contract, err := bindERC721Holder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721HolderFilterer{contract: contract}, nil
}

// bindERC721Holder binds a generic wrapper to an already deployed contract.
func bindERC721Holder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC721HolderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Holder *ERC721HolderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Holder.Contract.ERC721HolderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Holder *ERC721HolderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Holder.Contract.ERC721HolderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Holder *ERC721HolderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Holder.Contract.ERC721HolderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Holder *ERC721HolderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Holder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Holder *ERC721HolderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Holder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Holder *ERC721HolderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Holder.Contract.contract.Transact(opts, method, params...)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_ERC721Holder *ERC721HolderTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _ERC721Holder.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_ERC721Holder *ERC721HolderSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _ERC721Holder.Contract.OnERC721Received(&_ERC721Holder.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_ERC721Holder *ERC721HolderTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _ERC721Holder.Contract.OnERC721Received(&_ERC721Holder.TransactOpts, arg0, arg1, arg2, arg3)
}
