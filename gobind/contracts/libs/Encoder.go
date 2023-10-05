// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package libs

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

// EncoderMetaData contains all meta data concerning the Encoder contract.
var EncoderMetaData = &bind.MetaData{
	ABI: "[]",
}

// EncoderABI is the input ABI used to generate the binding from.
// Deprecated: Use EncoderMetaData.ABI instead.
var EncoderABI = EncoderMetaData.ABI

// Encoder is an auto generated Go binding around an Ethereum contract.
type Encoder struct {
	EncoderCaller     // Read-only binding to the contract
	EncoderTransactor // Write-only binding to the contract
	EncoderFilterer   // Log filterer for contract events
}

// EncoderCaller is an auto generated read-only Go binding around an Ethereum contract.
type EncoderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EncoderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EncoderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EncoderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EncoderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EncoderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EncoderSession struct {
	Contract     *Encoder          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EncoderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EncoderCallerSession struct {
	Contract *EncoderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EncoderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EncoderTransactorSession struct {
	Contract     *EncoderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EncoderRaw is an auto generated low-level Go binding around an Ethereum contract.
type EncoderRaw struct {
	Contract *Encoder // Generic contract binding to access the raw methods on
}

// EncoderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EncoderCallerRaw struct {
	Contract *EncoderCaller // Generic read-only contract binding to access the raw methods on
}

// EncoderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EncoderTransactorRaw struct {
	Contract *EncoderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEncoder creates a new instance of Encoder, bound to a specific deployed contract.
func NewEncoder(address common.Address, backend bind.ContractBackend) (*Encoder, error) {
	contract, err := bindEncoder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Encoder{EncoderCaller: EncoderCaller{contract: contract}, EncoderTransactor: EncoderTransactor{contract: contract}, EncoderFilterer: EncoderFilterer{contract: contract}}, nil
}

// NewEncoderCaller creates a new read-only instance of Encoder, bound to a specific deployed contract.
func NewEncoderCaller(address common.Address, caller bind.ContractCaller) (*EncoderCaller, error) {
	contract, err := bindEncoder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EncoderCaller{contract: contract}, nil
}

// NewEncoderTransactor creates a new write-only instance of Encoder, bound to a specific deployed contract.
func NewEncoderTransactor(address common.Address, transactor bind.ContractTransactor) (*EncoderTransactor, error) {
	contract, err := bindEncoder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EncoderTransactor{contract: contract}, nil
}

// NewEncoderFilterer creates a new log filterer instance of Encoder, bound to a specific deployed contract.
func NewEncoderFilterer(address common.Address, filterer bind.ContractFilterer) (*EncoderFilterer, error) {
	contract, err := bindEncoder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EncoderFilterer{contract: contract}, nil
}

// bindEncoder binds a generic wrapper to an already deployed contract.
func bindEncoder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EncoderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Encoder *EncoderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Encoder.Contract.EncoderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Encoder *EncoderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Encoder.Contract.EncoderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Encoder *EncoderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Encoder.Contract.EncoderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Encoder *EncoderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Encoder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Encoder *EncoderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Encoder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Encoder *EncoderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Encoder.Contract.contract.Transact(opts, method, params...)
}
