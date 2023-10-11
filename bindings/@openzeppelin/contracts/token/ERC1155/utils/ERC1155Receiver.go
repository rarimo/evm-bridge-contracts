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

// ERC1155ReceiverMetaData contains all meta data concerning the ERC1155Receiver contract.
var ERC1155ReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC1155ReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1155ReceiverMetaData.ABI instead.
var ERC1155ReceiverABI = ERC1155ReceiverMetaData.ABI

// ERC1155Receiver is an auto generated Go binding around an Ethereum contract.
type ERC1155Receiver struct {
	ERC1155ReceiverCaller     // Read-only binding to the contract
	ERC1155ReceiverTransactor // Write-only binding to the contract
	ERC1155ReceiverFilterer   // Log filterer for contract events
}

// ERC1155ReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1155ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155ReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1155ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155ReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1155ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155ReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1155ReceiverSession struct {
	Contract     *ERC1155Receiver  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1155ReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1155ReceiverCallerSession struct {
	Contract *ERC1155ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ERC1155ReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1155ReceiverTransactorSession struct {
	Contract     *ERC1155ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ERC1155ReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1155ReceiverRaw struct {
	Contract *ERC1155Receiver // Generic contract binding to access the raw methods on
}

// ERC1155ReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1155ReceiverCallerRaw struct {
	Contract *ERC1155ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1155ReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1155ReceiverTransactorRaw struct {
	Contract *ERC1155ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155Receiver creates a new instance of ERC1155Receiver, bound to a specific deployed contract.
func NewERC1155Receiver(address common.Address, backend bind.ContractBackend) (*ERC1155Receiver, error) {
	contract, err := bindERC1155Receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155Receiver{ERC1155ReceiverCaller: ERC1155ReceiverCaller{contract: contract}, ERC1155ReceiverTransactor: ERC1155ReceiverTransactor{contract: contract}, ERC1155ReceiverFilterer: ERC1155ReceiverFilterer{contract: contract}}, nil
}

// NewERC1155ReceiverCaller creates a new read-only instance of ERC1155Receiver, bound to a specific deployed contract.
func NewERC1155ReceiverCaller(address common.Address, caller bind.ContractCaller) (*ERC1155ReceiverCaller, error) {
	contract, err := bindERC1155Receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155ReceiverCaller{contract: contract}, nil
}

// NewERC1155ReceiverTransactor creates a new write-only instance of ERC1155Receiver, bound to a specific deployed contract.
func NewERC1155ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1155ReceiverTransactor, error) {
	contract, err := bindERC1155Receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155ReceiverTransactor{contract: contract}, nil
}

// NewERC1155ReceiverFilterer creates a new log filterer instance of ERC1155Receiver, bound to a specific deployed contract.
func NewERC1155ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1155ReceiverFilterer, error) {
	contract, err := bindERC1155Receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155ReceiverFilterer{contract: contract}, nil
}

// bindERC1155Receiver binds a generic wrapper to an already deployed contract.
func bindERC1155Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC1155ReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155Receiver *ERC1155ReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155Receiver.Contract.ERC1155ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155Receiver *ERC1155ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Receiver.Contract.ERC1155ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155Receiver *ERC1155ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155Receiver.Contract.ERC1155ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155Receiver *ERC1155ReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155Receiver *ERC1155ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155Receiver *ERC1155ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155Receiver.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155Receiver *ERC1155ReceiverCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC1155Receiver.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155Receiver *ERC1155ReceiverSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155Receiver.Contract.SupportsInterface(&_ERC1155Receiver.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155Receiver *ERC1155ReceiverCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155Receiver.Contract.SupportsInterface(&_ERC1155Receiver.CallOpts, interfaceId)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_ERC1155Receiver *ERC1155ReceiverTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Receiver.contract.Transact(opts, "onERC1155BatchReceived", operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_ERC1155Receiver *ERC1155ReceiverSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Receiver.Contract.OnERC1155BatchReceived(&_ERC1155Receiver.TransactOpts, operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_ERC1155Receiver *ERC1155ReceiverTransactorSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Receiver.Contract.OnERC1155BatchReceived(&_ERC1155Receiver.TransactOpts, operator, from, ids, values, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_ERC1155Receiver *ERC1155ReceiverTransactor) OnERC1155Received(opts *bind.TransactOpts, operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Receiver.contract.Transact(opts, "onERC1155Received", operator, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_ERC1155Receiver *ERC1155ReceiverSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Receiver.Contract.OnERC1155Received(&_ERC1155Receiver.TransactOpts, operator, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_ERC1155Receiver *ERC1155ReceiverTransactorSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Receiver.Contract.OnERC1155Received(&_ERC1155Receiver.TransactOpts, operator, from, id, value, data)
}
