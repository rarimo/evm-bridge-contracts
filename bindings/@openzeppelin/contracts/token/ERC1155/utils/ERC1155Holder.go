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

// ERC1155HolderMetaData contains all meta data concerning the ERC1155Holder contract.
var ERC1155HolderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC1155HolderABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1155HolderMetaData.ABI instead.
var ERC1155HolderABI = ERC1155HolderMetaData.ABI

// ERC1155Holder is an auto generated Go binding around an Ethereum contract.
type ERC1155Holder struct {
	ERC1155HolderCaller     // Read-only binding to the contract
	ERC1155HolderTransactor // Write-only binding to the contract
	ERC1155HolderFilterer   // Log filterer for contract events
}

// ERC1155HolderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1155HolderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155HolderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1155HolderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155HolderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1155HolderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155HolderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1155HolderSession struct {
	Contract     *ERC1155Holder    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1155HolderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1155HolderCallerSession struct {
	Contract *ERC1155HolderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC1155HolderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1155HolderTransactorSession struct {
	Contract     *ERC1155HolderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC1155HolderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1155HolderRaw struct {
	Contract *ERC1155Holder // Generic contract binding to access the raw methods on
}

// ERC1155HolderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1155HolderCallerRaw struct {
	Contract *ERC1155HolderCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1155HolderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1155HolderTransactorRaw struct {
	Contract *ERC1155HolderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155Holder creates a new instance of ERC1155Holder, bound to a specific deployed contract.
func NewERC1155Holder(address common.Address, backend bind.ContractBackend) (*ERC1155Holder, error) {
	contract, err := bindERC1155Holder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155Holder{ERC1155HolderCaller: ERC1155HolderCaller{contract: contract}, ERC1155HolderTransactor: ERC1155HolderTransactor{contract: contract}, ERC1155HolderFilterer: ERC1155HolderFilterer{contract: contract}}, nil
}

// NewERC1155HolderCaller creates a new read-only instance of ERC1155Holder, bound to a specific deployed contract.
func NewERC1155HolderCaller(address common.Address, caller bind.ContractCaller) (*ERC1155HolderCaller, error) {
	contract, err := bindERC1155Holder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155HolderCaller{contract: contract}, nil
}

// NewERC1155HolderTransactor creates a new write-only instance of ERC1155Holder, bound to a specific deployed contract.
func NewERC1155HolderTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1155HolderTransactor, error) {
	contract, err := bindERC1155Holder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155HolderTransactor{contract: contract}, nil
}

// NewERC1155HolderFilterer creates a new log filterer instance of ERC1155Holder, bound to a specific deployed contract.
func NewERC1155HolderFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1155HolderFilterer, error) {
	contract, err := bindERC1155Holder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155HolderFilterer{contract: contract}, nil
}

// bindERC1155Holder binds a generic wrapper to an already deployed contract.
func bindERC1155Holder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC1155HolderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155Holder *ERC1155HolderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155Holder.Contract.ERC1155HolderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155Holder *ERC1155HolderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Holder.Contract.ERC1155HolderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155Holder *ERC1155HolderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155Holder.Contract.ERC1155HolderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155Holder *ERC1155HolderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155Holder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155Holder *ERC1155HolderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Holder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155Holder *ERC1155HolderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155Holder.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155Holder *ERC1155HolderCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC1155Holder.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155Holder *ERC1155HolderSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155Holder.Contract.SupportsInterface(&_ERC1155Holder.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155Holder *ERC1155HolderCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155Holder.Contract.SupportsInterface(&_ERC1155Holder.CallOpts, interfaceId)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ERC1155Holder *ERC1155HolderTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Holder.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ERC1155Holder *ERC1155HolderSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Holder.Contract.OnERC1155BatchReceived(&_ERC1155Holder.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ERC1155Holder *ERC1155HolderTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Holder.Contract.OnERC1155BatchReceived(&_ERC1155Holder.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ERC1155Holder *ERC1155HolderTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Holder.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ERC1155Holder *ERC1155HolderSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Holder.Contract.OnERC1155Received(&_ERC1155Holder.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ERC1155Holder *ERC1155HolderTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155Holder.Contract.OnERC1155Received(&_ERC1155Holder.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}
