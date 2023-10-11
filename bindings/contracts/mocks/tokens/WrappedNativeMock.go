// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokens

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

// WrappedNativeMockMetaData contains all meta data concerning the WrappedNativeMock contract.
var WrappedNativeMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WrappedNativeMockABI is the input ABI used to generate the binding from.
// Deprecated: Use WrappedNativeMockMetaData.ABI instead.
var WrappedNativeMockABI = WrappedNativeMockMetaData.ABI

// WrappedNativeMock is an auto generated Go binding around an Ethereum contract.
type WrappedNativeMock struct {
	WrappedNativeMockCaller     // Read-only binding to the contract
	WrappedNativeMockTransactor // Write-only binding to the contract
	WrappedNativeMockFilterer   // Log filterer for contract events
}

// WrappedNativeMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrappedNativeMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedNativeMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrappedNativeMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedNativeMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WrappedNativeMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedNativeMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrappedNativeMockSession struct {
	Contract     *WrappedNativeMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WrappedNativeMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrappedNativeMockCallerSession struct {
	Contract *WrappedNativeMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// WrappedNativeMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrappedNativeMockTransactorSession struct {
	Contract     *WrappedNativeMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// WrappedNativeMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrappedNativeMockRaw struct {
	Contract *WrappedNativeMock // Generic contract binding to access the raw methods on
}

// WrappedNativeMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrappedNativeMockCallerRaw struct {
	Contract *WrappedNativeMockCaller // Generic read-only contract binding to access the raw methods on
}

// WrappedNativeMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrappedNativeMockTransactorRaw struct {
	Contract *WrappedNativeMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrappedNativeMock creates a new instance of WrappedNativeMock, bound to a specific deployed contract.
func NewWrappedNativeMock(address common.Address, backend bind.ContractBackend) (*WrappedNativeMock, error) {
	contract, err := bindWrappedNativeMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeMock{WrappedNativeMockCaller: WrappedNativeMockCaller{contract: contract}, WrappedNativeMockTransactor: WrappedNativeMockTransactor{contract: contract}, WrappedNativeMockFilterer: WrappedNativeMockFilterer{contract: contract}}, nil
}

// NewWrappedNativeMockCaller creates a new read-only instance of WrappedNativeMock, bound to a specific deployed contract.
func NewWrappedNativeMockCaller(address common.Address, caller bind.ContractCaller) (*WrappedNativeMockCaller, error) {
	contract, err := bindWrappedNativeMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeMockCaller{contract: contract}, nil
}

// NewWrappedNativeMockTransactor creates a new write-only instance of WrappedNativeMock, bound to a specific deployed contract.
func NewWrappedNativeMockTransactor(address common.Address, transactor bind.ContractTransactor) (*WrappedNativeMockTransactor, error) {
	contract, err := bindWrappedNativeMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeMockTransactor{contract: contract}, nil
}

// NewWrappedNativeMockFilterer creates a new log filterer instance of WrappedNativeMock, bound to a specific deployed contract.
func NewWrappedNativeMockFilterer(address common.Address, filterer bind.ContractFilterer) (*WrappedNativeMockFilterer, error) {
	contract, err := bindWrappedNativeMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeMockFilterer{contract: contract}, nil
}

// bindWrappedNativeMock binds a generic wrapper to an already deployed contract.
func bindWrappedNativeMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WrappedNativeMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedNativeMock *WrappedNativeMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedNativeMock.Contract.WrappedNativeMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedNativeMock *WrappedNativeMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.WrappedNativeMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedNativeMock *WrappedNativeMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.WrappedNativeMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedNativeMock *WrappedNativeMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedNativeMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedNativeMock *WrappedNativeMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedNativeMock *WrappedNativeMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedNativeMock *WrappedNativeMockCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedNativeMock.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedNativeMock *WrappedNativeMockSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedNativeMock.Contract.Allowance(&_WrappedNativeMock.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedNativeMock *WrappedNativeMockCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedNativeMock.Contract.Allowance(&_WrappedNativeMock.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedNativeMock *WrappedNativeMockCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedNativeMock.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedNativeMock *WrappedNativeMockSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedNativeMock.Contract.BalanceOf(&_WrappedNativeMock.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedNativeMock *WrappedNativeMockCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedNativeMock.Contract.BalanceOf(&_WrappedNativeMock.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedNativeMock *WrappedNativeMockCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WrappedNativeMock.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedNativeMock *WrappedNativeMockSession) Decimals() (uint8, error) {
	return _WrappedNativeMock.Contract.Decimals(&_WrappedNativeMock.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedNativeMock *WrappedNativeMockCallerSession) Decimals() (uint8, error) {
	return _WrappedNativeMock.Contract.Decimals(&_WrappedNativeMock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedNativeMock *WrappedNativeMockCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedNativeMock.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedNativeMock *WrappedNativeMockSession) Name() (string, error) {
	return _WrappedNativeMock.Contract.Name(&_WrappedNativeMock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedNativeMock *WrappedNativeMockCallerSession) Name() (string, error) {
	return _WrappedNativeMock.Contract.Name(&_WrappedNativeMock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedNativeMock *WrappedNativeMockCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedNativeMock.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedNativeMock *WrappedNativeMockSession) Symbol() (string, error) {
	return _WrappedNativeMock.Contract.Symbol(&_WrappedNativeMock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedNativeMock *WrappedNativeMockCallerSession) Symbol() (string, error) {
	return _WrappedNativeMock.Contract.Symbol(&_WrappedNativeMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedNativeMock *WrappedNativeMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WrappedNativeMock.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedNativeMock *WrappedNativeMockSession) TotalSupply() (*big.Int, error) {
	return _WrappedNativeMock.Contract.TotalSupply(&_WrappedNativeMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedNativeMock *WrappedNativeMockCallerSession) TotalSupply() (*big.Int, error) {
	return _WrappedNativeMock.Contract.TotalSupply(&_WrappedNativeMock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedNativeMock *WrappedNativeMockTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedNativeMock *WrappedNativeMockSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.Approve(&_WrappedNativeMock.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedNativeMock *WrappedNativeMockTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.Approve(&_WrappedNativeMock.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedNativeMock *WrappedNativeMockTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedNativeMock *WrappedNativeMockSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.DecreaseAllowance(&_WrappedNativeMock.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedNativeMock *WrappedNativeMockTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.DecreaseAllowance(&_WrappedNativeMock.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedNativeMock *WrappedNativeMockTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedNativeMock.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedNativeMock *WrappedNativeMockSession) Deposit() (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.Deposit(&_WrappedNativeMock.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedNativeMock *WrappedNativeMockTransactorSession) Deposit() (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.Deposit(&_WrappedNativeMock.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedNativeMock *WrappedNativeMockTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedNativeMock *WrappedNativeMockSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.IncreaseAllowance(&_WrappedNativeMock.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedNativeMock *WrappedNativeMockTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.IncreaseAllowance(&_WrappedNativeMock.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WrappedNativeMock *WrappedNativeMockTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WrappedNativeMock *WrappedNativeMockSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.Transfer(&_WrappedNativeMock.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WrappedNativeMock *WrappedNativeMockTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.Transfer(&_WrappedNativeMock.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WrappedNativeMock *WrappedNativeMockTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WrappedNativeMock *WrappedNativeMockSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.TransferFrom(&_WrappedNativeMock.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WrappedNativeMock *WrappedNativeMockTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.TransferFrom(&_WrappedNativeMock.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount_) returns()
func (_WrappedNativeMock *WrappedNativeMockTransactor) Withdraw(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.contract.Transact(opts, "withdraw", amount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount_) returns()
func (_WrappedNativeMock *WrappedNativeMockSession) Withdraw(amount_ *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.Withdraw(&_WrappedNativeMock.TransactOpts, amount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount_) returns()
func (_WrappedNativeMock *WrappedNativeMockTransactorSession) Withdraw(amount_ *big.Int) (*types.Transaction, error) {
	return _WrappedNativeMock.Contract.Withdraw(&_WrappedNativeMock.TransactOpts, amount_)
}

// WrappedNativeMockApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WrappedNativeMock contract.
type WrappedNativeMockApprovalIterator struct {
	Event *WrappedNativeMockApproval // Event containing the contract specifics and raw log

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
func (it *WrappedNativeMockApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedNativeMockApproval)
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
		it.Event = new(WrappedNativeMockApproval)
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
func (it *WrappedNativeMockApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedNativeMockApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedNativeMockApproval represents a Approval event raised by the WrappedNativeMock contract.
type WrappedNativeMockApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedNativeMock *WrappedNativeMockFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WrappedNativeMockApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedNativeMock.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeMockApprovalIterator{contract: _WrappedNativeMock.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedNativeMock *WrappedNativeMockFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WrappedNativeMockApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedNativeMock.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedNativeMockApproval)
				if err := _WrappedNativeMock.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedNativeMock *WrappedNativeMockFilterer) ParseApproval(log types.Log) (*WrappedNativeMockApproval, error) {
	event := new(WrappedNativeMockApproval)
	if err := _WrappedNativeMock.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedNativeMockTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WrappedNativeMock contract.
type WrappedNativeMockTransferIterator struct {
	Event *WrappedNativeMockTransfer // Event containing the contract specifics and raw log

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
func (it *WrappedNativeMockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedNativeMockTransfer)
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
		it.Event = new(WrappedNativeMockTransfer)
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
func (it *WrappedNativeMockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedNativeMockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedNativeMockTransfer represents a Transfer event raised by the WrappedNativeMock contract.
type WrappedNativeMockTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedNativeMock *WrappedNativeMockFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrappedNativeMockTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedNativeMock.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeMockTransferIterator{contract: _WrappedNativeMock.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedNativeMock *WrappedNativeMockFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WrappedNativeMockTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedNativeMock.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedNativeMockTransfer)
				if err := _WrappedNativeMock.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedNativeMock *WrappedNativeMockFilterer) ParseTransfer(log types.Log) (*WrappedNativeMockTransfer, error) {
	event := new(WrappedNativeMockTransfer)
	if err := _WrappedNativeMock.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
