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

// ISignersMetaData contains all meta data concerning the ISigners contract.
var ISignersMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"checkSignatureAndIncrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"getSigComponents\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"validateChangeAddressSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ISignersABI is the input ABI used to generate the binding from.
// Deprecated: Use ISignersMetaData.ABI instead.
var ISignersABI = ISignersMetaData.ABI

// ISigners is an auto generated Go binding around an Ethereum contract.
type ISigners struct {
	ISignersCaller     // Read-only binding to the contract
	ISignersTransactor // Write-only binding to the contract
	ISignersFilterer   // Log filterer for contract events
}

// ISignersCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISignersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISignersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISignersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISignersSession struct {
	Contract     *ISigners         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISignersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISignersCallerSession struct {
	Contract *ISignersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ISignersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISignersTransactorSession struct {
	Contract     *ISignersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ISignersRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISignersRaw struct {
	Contract *ISigners // Generic contract binding to access the raw methods on
}

// ISignersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISignersCallerRaw struct {
	Contract *ISignersCaller // Generic read-only contract binding to access the raw methods on
}

// ISignersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISignersTransactorRaw struct {
	Contract *ISignersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISigners creates a new instance of ISigners, bound to a specific deployed contract.
func NewISigners(address common.Address, backend bind.ContractBackend) (*ISigners, error) {
	contract, err := bindISigners(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISigners{ISignersCaller: ISignersCaller{contract: contract}, ISignersTransactor: ISignersTransactor{contract: contract}, ISignersFilterer: ISignersFilterer{contract: contract}}, nil
}

// NewISignersCaller creates a new read-only instance of ISigners, bound to a specific deployed contract.
func NewISignersCaller(address common.Address, caller bind.ContractCaller) (*ISignersCaller, error) {
	contract, err := bindISigners(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISignersCaller{contract: contract}, nil
}

// NewISignersTransactor creates a new write-only instance of ISigners, bound to a specific deployed contract.
func NewISignersTransactor(address common.Address, transactor bind.ContractTransactor) (*ISignersTransactor, error) {
	contract, err := bindISigners(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISignersTransactor{contract: contract}, nil
}

// NewISignersFilterer creates a new log filterer instance of ISigners, bound to a specific deployed contract.
func NewISignersFilterer(address common.Address, filterer bind.ContractFilterer) (*ISignersFilterer, error) {
	contract, err := bindISigners(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISignersFilterer{contract: contract}, nil
}

// bindISigners binds a generic wrapper to an already deployed contract.
func bindISigners(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISignersMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISigners *ISignersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISigners.Contract.ISignersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISigners *ISignersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISigners.Contract.ISignersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISigners *ISignersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISigners.Contract.ISignersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISigners *ISignersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISigners.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISigners *ISignersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISigners.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISigners *ISignersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISigners.Contract.contract.Transact(opts, method, params...)
}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_ISigners *ISignersCaller) GetSigComponents(opts *bind.CallOpts, methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	var out []interface{}
	err := _ISigners.contract.Call(opts, &out, "getSigComponents", methodId_, contractAddress_)

	outstruct := new(struct {
		ChainName string
		Nonce     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Nonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_ISigners *ISignersSession) GetSigComponents(methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	return _ISigners.Contract.GetSigComponents(&_ISigners.CallOpts, methodId_, contractAddress_)
}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_ISigners *ISignersCallerSession) GetSigComponents(methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	return _ISigners.Contract.GetSigComponents(&_ISigners.CallOpts, methodId_, contractAddress_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_ISigners *ISignersTransactor) CheckSignatureAndIncrementNonce(opts *bind.TransactOpts, methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _ISigners.contract.Transact(opts, "checkSignatureAndIncrementNonce", methodId_, contractAddress_, signHash_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_ISigners *ISignersSession) CheckSignatureAndIncrementNonce(methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _ISigners.Contract.CheckSignatureAndIncrementNonce(&_ISigners.TransactOpts, methodId_, contractAddress_, signHash_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_ISigners *ISignersTransactorSession) CheckSignatureAndIncrementNonce(methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _ISigners.Contract.CheckSignatureAndIncrementNonce(&_ISigners.TransactOpts, methodId_, contractAddress_, signHash_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_ISigners *ISignersTransactor) ValidateChangeAddressSignature(opts *bind.TransactOpts, methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _ISigners.contract.Transact(opts, "validateChangeAddressSignature", methodId_, contractAddress_, newAddress_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_ISigners *ISignersSession) ValidateChangeAddressSignature(methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _ISigners.Contract.ValidateChangeAddressSignature(&_ISigners.TransactOpts, methodId_, contractAddress_, newAddress_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_ISigners *ISignersTransactorSession) ValidateChangeAddressSignature(methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _ISigners.Contract.ValidateChangeAddressSignature(&_ISigners.TransactOpts, methodId_, contractAddress_, newAddress_, signature_)
}
