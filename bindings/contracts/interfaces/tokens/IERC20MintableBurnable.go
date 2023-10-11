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

// IERC20MintableBurnableMetaData contains all meta data concerning the IERC20MintableBurnable contract.
var IERC20MintableBurnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"mintTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20MintableBurnableABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MintableBurnableMetaData.ABI instead.
var IERC20MintableBurnableABI = IERC20MintableBurnableMetaData.ABI

// IERC20MintableBurnable is an auto generated Go binding around an Ethereum contract.
type IERC20MintableBurnable struct {
	IERC20MintableBurnableCaller     // Read-only binding to the contract
	IERC20MintableBurnableTransactor // Write-only binding to the contract
	IERC20MintableBurnableFilterer   // Log filterer for contract events
}

// IERC20MintableBurnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20MintableBurnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MintableBurnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20MintableBurnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MintableBurnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20MintableBurnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MintableBurnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20MintableBurnableSession struct {
	Contract     *IERC20MintableBurnable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IERC20MintableBurnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20MintableBurnableCallerSession struct {
	Contract *IERC20MintableBurnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IERC20MintableBurnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20MintableBurnableTransactorSession struct {
	Contract     *IERC20MintableBurnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IERC20MintableBurnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20MintableBurnableRaw struct {
	Contract *IERC20MintableBurnable // Generic contract binding to access the raw methods on
}

// IERC20MintableBurnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20MintableBurnableCallerRaw struct {
	Contract *IERC20MintableBurnableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20MintableBurnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20MintableBurnableTransactorRaw struct {
	Contract *IERC20MintableBurnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20MintableBurnable creates a new instance of IERC20MintableBurnable, bound to a specific deployed contract.
func NewIERC20MintableBurnable(address common.Address, backend bind.ContractBackend) (*IERC20MintableBurnable, error) {
	contract, err := bindIERC20MintableBurnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20MintableBurnable{IERC20MintableBurnableCaller: IERC20MintableBurnableCaller{contract: contract}, IERC20MintableBurnableTransactor: IERC20MintableBurnableTransactor{contract: contract}, IERC20MintableBurnableFilterer: IERC20MintableBurnableFilterer{contract: contract}}, nil
}

// NewIERC20MintableBurnableCaller creates a new read-only instance of IERC20MintableBurnable, bound to a specific deployed contract.
func NewIERC20MintableBurnableCaller(address common.Address, caller bind.ContractCaller) (*IERC20MintableBurnableCaller, error) {
	contract, err := bindIERC20MintableBurnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MintableBurnableCaller{contract: contract}, nil
}

// NewIERC20MintableBurnableTransactor creates a new write-only instance of IERC20MintableBurnable, bound to a specific deployed contract.
func NewIERC20MintableBurnableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20MintableBurnableTransactor, error) {
	contract, err := bindIERC20MintableBurnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MintableBurnableTransactor{contract: contract}, nil
}

// NewIERC20MintableBurnableFilterer creates a new log filterer instance of IERC20MintableBurnable, bound to a specific deployed contract.
func NewIERC20MintableBurnableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20MintableBurnableFilterer, error) {
	contract, err := bindIERC20MintableBurnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20MintableBurnableFilterer{contract: contract}, nil
}

// bindIERC20MintableBurnable binds a generic wrapper to an already deployed contract.
func bindIERC20MintableBurnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20MintableBurnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20MintableBurnable *IERC20MintableBurnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20MintableBurnable.Contract.IERC20MintableBurnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20MintableBurnable *IERC20MintableBurnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20MintableBurnable.Contract.IERC20MintableBurnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20MintableBurnable *IERC20MintableBurnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20MintableBurnable.Contract.IERC20MintableBurnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20MintableBurnable *IERC20MintableBurnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20MintableBurnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20MintableBurnable *IERC20MintableBurnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20MintableBurnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20MintableBurnable *IERC20MintableBurnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20MintableBurnable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20MintableBurnable *IERC20MintableBurnableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20MintableBurnable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20MintableBurnable *IERC20MintableBurnableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20MintableBurnable.Contract.Allowance(&_IERC20MintableBurnable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20MintableBurnable *IERC20MintableBurnableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20MintableBurnable.Contract.Allowance(&_IERC20MintableBurnable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20MintableBurnable *IERC20MintableBurnableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20MintableBurnable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20MintableBurnable *IERC20MintableBurnableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20MintableBurnable.Contract.BalanceOf(&_IERC20MintableBurnable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20MintableBurnable *IERC20MintableBurnableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20MintableBurnable.Contract.BalanceOf(&_IERC20MintableBurnable.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20MintableBurnable *IERC20MintableBurnableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20MintableBurnable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20MintableBurnable *IERC20MintableBurnableSession) TotalSupply() (*big.Int, error) {
	return _IERC20MintableBurnable.Contract.TotalSupply(&_IERC20MintableBurnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20MintableBurnable *IERC20MintableBurnableCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20MintableBurnable.Contract.TotalSupply(&_IERC20MintableBurnable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20MintableBurnable *IERC20MintableBurnableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MintableBurnable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20MintableBurnable *IERC20MintableBurnableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MintableBurnable.Contract.Approve(&_IERC20MintableBurnable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20MintableBurnable *IERC20MintableBurnableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MintableBurnable.Contract.Approve(&_IERC20MintableBurnable.TransactOpts, spender, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address payer_, uint256 amount_) returns()
func (_IERC20MintableBurnable *IERC20MintableBurnableTransactor) BurnFrom(opts *bind.TransactOpts, payer_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IERC20MintableBurnable.contract.Transact(opts, "burnFrom", payer_, amount_)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address payer_, uint256 amount_) returns()
func (_IERC20MintableBurnable *IERC20MintableBurnableSession) BurnFrom(payer_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IERC20MintableBurnable.Contract.BurnFrom(&_IERC20MintableBurnable.TransactOpts, payer_, amount_)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address payer_, uint256 amount_) returns()
func (_IERC20MintableBurnable *IERC20MintableBurnableTransactorSession) BurnFrom(payer_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IERC20MintableBurnable.Contract.BurnFrom(&_IERC20MintableBurnable.TransactOpts, payer_, amount_)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address receiver_, uint256 amount_) returns()
func (_IERC20MintableBurnable *IERC20MintableBurnableTransactor) MintTo(opts *bind.TransactOpts, receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IERC20MintableBurnable.contract.Transact(opts, "mintTo", receiver_, amount_)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address receiver_, uint256 amount_) returns()
func (_IERC20MintableBurnable *IERC20MintableBurnableSession) MintTo(receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IERC20MintableBurnable.Contract.MintTo(&_IERC20MintableBurnable.TransactOpts, receiver_, amount_)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address receiver_, uint256 amount_) returns()
func (_IERC20MintableBurnable *IERC20MintableBurnableTransactorSession) MintTo(receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IERC20MintableBurnable.Contract.MintTo(&_IERC20MintableBurnable.TransactOpts, receiver_, amount_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20MintableBurnable *IERC20MintableBurnableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MintableBurnable.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20MintableBurnable *IERC20MintableBurnableSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MintableBurnable.Contract.Transfer(&_IERC20MintableBurnable.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20MintableBurnable *IERC20MintableBurnableTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MintableBurnable.Contract.Transfer(&_IERC20MintableBurnable.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20MintableBurnable *IERC20MintableBurnableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MintableBurnable.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20MintableBurnable *IERC20MintableBurnableSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MintableBurnable.Contract.TransferFrom(&_IERC20MintableBurnable.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20MintableBurnable *IERC20MintableBurnableTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MintableBurnable.Contract.TransferFrom(&_IERC20MintableBurnable.TransactOpts, from, to, amount)
}

// IERC20MintableBurnableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20MintableBurnable contract.
type IERC20MintableBurnableApprovalIterator struct {
	Event *IERC20MintableBurnableApproval // Event containing the contract specifics and raw log

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
func (it *IERC20MintableBurnableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MintableBurnableApproval)
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
		it.Event = new(IERC20MintableBurnableApproval)
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
func (it *IERC20MintableBurnableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MintableBurnableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MintableBurnableApproval represents a Approval event raised by the IERC20MintableBurnable contract.
type IERC20MintableBurnableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20MintableBurnable *IERC20MintableBurnableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20MintableBurnableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20MintableBurnable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MintableBurnableApprovalIterator{contract: _IERC20MintableBurnable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20MintableBurnable *IERC20MintableBurnableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20MintableBurnableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20MintableBurnable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MintableBurnableApproval)
				if err := _IERC20MintableBurnable.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20MintableBurnable *IERC20MintableBurnableFilterer) ParseApproval(log types.Log) (*IERC20MintableBurnableApproval, error) {
	event := new(IERC20MintableBurnableApproval)
	if err := _IERC20MintableBurnable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MintableBurnableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20MintableBurnable contract.
type IERC20MintableBurnableTransferIterator struct {
	Event *IERC20MintableBurnableTransfer // Event containing the contract specifics and raw log

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
func (it *IERC20MintableBurnableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MintableBurnableTransfer)
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
		it.Event = new(IERC20MintableBurnableTransfer)
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
func (it *IERC20MintableBurnableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MintableBurnableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MintableBurnableTransfer represents a Transfer event raised by the IERC20MintableBurnable contract.
type IERC20MintableBurnableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20MintableBurnable *IERC20MintableBurnableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20MintableBurnableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20MintableBurnable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MintableBurnableTransferIterator{contract: _IERC20MintableBurnable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20MintableBurnable *IERC20MintableBurnableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20MintableBurnableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20MintableBurnable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MintableBurnableTransfer)
				if err := _IERC20MintableBurnable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20MintableBurnable *IERC20MintableBurnableFilterer) ParseTransfer(log types.Log) (*IERC20MintableBurnableTransfer, error) {
	event := new(IERC20MintableBurnableTransfer)
	if err := _IERC20MintableBurnable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
