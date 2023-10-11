// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interfaces

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

// IERC1822ProxiableMetaData contains all meta data concerning the IERC1822Proxiable contract.
var IERC1822ProxiableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IERC1822ProxiableABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC1822ProxiableMetaData.ABI instead.
var IERC1822ProxiableABI = IERC1822ProxiableMetaData.ABI

// IERC1822Proxiable is an auto generated Go binding around an Ethereum contract.
type IERC1822Proxiable struct {
	IERC1822ProxiableCaller     // Read-only binding to the contract
	IERC1822ProxiableTransactor // Write-only binding to the contract
	IERC1822ProxiableFilterer   // Log filterer for contract events
}

// IERC1822ProxiableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1822ProxiableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1822ProxiableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1822ProxiableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1822ProxiableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1822ProxiableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1822ProxiableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1822ProxiableSession struct {
	Contract     *IERC1822Proxiable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC1822ProxiableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1822ProxiableCallerSession struct {
	Contract *IERC1822ProxiableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IERC1822ProxiableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1822ProxiableTransactorSession struct {
	Contract     *IERC1822ProxiableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IERC1822ProxiableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1822ProxiableRaw struct {
	Contract *IERC1822Proxiable // Generic contract binding to access the raw methods on
}

// IERC1822ProxiableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1822ProxiableCallerRaw struct {
	Contract *IERC1822ProxiableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC1822ProxiableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1822ProxiableTransactorRaw struct {
	Contract *IERC1822ProxiableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1822Proxiable creates a new instance of IERC1822Proxiable, bound to a specific deployed contract.
func NewIERC1822Proxiable(address common.Address, backend bind.ContractBackend) (*IERC1822Proxiable, error) {
	contract, err := bindIERC1822Proxiable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1822Proxiable{IERC1822ProxiableCaller: IERC1822ProxiableCaller{contract: contract}, IERC1822ProxiableTransactor: IERC1822ProxiableTransactor{contract: contract}, IERC1822ProxiableFilterer: IERC1822ProxiableFilterer{contract: contract}}, nil
}

// NewIERC1822ProxiableCaller creates a new read-only instance of IERC1822Proxiable, bound to a specific deployed contract.
func NewIERC1822ProxiableCaller(address common.Address, caller bind.ContractCaller) (*IERC1822ProxiableCaller, error) {
	contract, err := bindIERC1822Proxiable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableCaller{contract: contract}, nil
}

// NewIERC1822ProxiableTransactor creates a new write-only instance of IERC1822Proxiable, bound to a specific deployed contract.
func NewIERC1822ProxiableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC1822ProxiableTransactor, error) {
	contract, err := bindIERC1822Proxiable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableTransactor{contract: contract}, nil
}

// NewIERC1822ProxiableFilterer creates a new log filterer instance of IERC1822Proxiable, bound to a specific deployed contract.
func NewIERC1822ProxiableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC1822ProxiableFilterer, error) {
	contract, err := bindIERC1822Proxiable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableFilterer{contract: contract}, nil
}

// bindIERC1822Proxiable binds a generic wrapper to an already deployed contract.
func bindIERC1822Proxiable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC1822ProxiableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1822Proxiable *IERC1822ProxiableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1822Proxiable.Contract.IERC1822ProxiableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1822Proxiable *IERC1822ProxiableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1822Proxiable.Contract.IERC1822ProxiableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1822Proxiable *IERC1822ProxiableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1822Proxiable.Contract.IERC1822ProxiableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1822Proxiable *IERC1822ProxiableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1822Proxiable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1822Proxiable *IERC1822ProxiableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1822Proxiable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1822Proxiable *IERC1822ProxiableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1822Proxiable.Contract.contract.Transact(opts, method, params...)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IERC1822Proxiable *IERC1822ProxiableCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IERC1822Proxiable.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IERC1822Proxiable *IERC1822ProxiableSession) ProxiableUUID() ([32]byte, error) {
	return _IERC1822Proxiable.Contract.ProxiableUUID(&_IERC1822Proxiable.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IERC1822Proxiable *IERC1822ProxiableCallerSession) ProxiableUUID() ([32]byte, error) {
	return _IERC1822Proxiable.Contract.ProxiableUUID(&_IERC1822Proxiable.CallOpts)
}
