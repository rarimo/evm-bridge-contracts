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

// IWrappedNativeMetaData contains all meta data concerning the IWrappedNative contract.
var IWrappedNativeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IWrappedNativeABI is the input ABI used to generate the binding from.
// Deprecated: Use IWrappedNativeMetaData.ABI instead.
var IWrappedNativeABI = IWrappedNativeMetaData.ABI

// IWrappedNative is an auto generated Go binding around an Ethereum contract.
type IWrappedNative struct {
	IWrappedNativeCaller     // Read-only binding to the contract
	IWrappedNativeTransactor // Write-only binding to the contract
	IWrappedNativeFilterer   // Log filterer for contract events
}

// IWrappedNativeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWrappedNativeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrappedNativeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWrappedNativeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrappedNativeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWrappedNativeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrappedNativeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWrappedNativeSession struct {
	Contract     *IWrappedNative   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWrappedNativeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWrappedNativeCallerSession struct {
	Contract *IWrappedNativeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IWrappedNativeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWrappedNativeTransactorSession struct {
	Contract     *IWrappedNativeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IWrappedNativeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWrappedNativeRaw struct {
	Contract *IWrappedNative // Generic contract binding to access the raw methods on
}

// IWrappedNativeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWrappedNativeCallerRaw struct {
	Contract *IWrappedNativeCaller // Generic read-only contract binding to access the raw methods on
}

// IWrappedNativeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWrappedNativeTransactorRaw struct {
	Contract *IWrappedNativeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWrappedNative creates a new instance of IWrappedNative, bound to a specific deployed contract.
func NewIWrappedNative(address common.Address, backend bind.ContractBackend) (*IWrappedNative, error) {
	contract, err := bindIWrappedNative(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWrappedNative{IWrappedNativeCaller: IWrappedNativeCaller{contract: contract}, IWrappedNativeTransactor: IWrappedNativeTransactor{contract: contract}, IWrappedNativeFilterer: IWrappedNativeFilterer{contract: contract}}, nil
}

// NewIWrappedNativeCaller creates a new read-only instance of IWrappedNative, bound to a specific deployed contract.
func NewIWrappedNativeCaller(address common.Address, caller bind.ContractCaller) (*IWrappedNativeCaller, error) {
	contract, err := bindIWrappedNative(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWrappedNativeCaller{contract: contract}, nil
}

// NewIWrappedNativeTransactor creates a new write-only instance of IWrappedNative, bound to a specific deployed contract.
func NewIWrappedNativeTransactor(address common.Address, transactor bind.ContractTransactor) (*IWrappedNativeTransactor, error) {
	contract, err := bindIWrappedNative(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWrappedNativeTransactor{contract: contract}, nil
}

// NewIWrappedNativeFilterer creates a new log filterer instance of IWrappedNative, bound to a specific deployed contract.
func NewIWrappedNativeFilterer(address common.Address, filterer bind.ContractFilterer) (*IWrappedNativeFilterer, error) {
	contract, err := bindIWrappedNative(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWrappedNativeFilterer{contract: contract}, nil
}

// bindIWrappedNative binds a generic wrapper to an already deployed contract.
func bindIWrappedNative(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IWrappedNativeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWrappedNative *IWrappedNativeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWrappedNative.Contract.IWrappedNativeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWrappedNative *IWrappedNativeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWrappedNative.Contract.IWrappedNativeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWrappedNative *IWrappedNativeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWrappedNative.Contract.IWrappedNativeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWrappedNative *IWrappedNativeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWrappedNative.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWrappedNative *IWrappedNativeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWrappedNative.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWrappedNative *IWrappedNativeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWrappedNative.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWrappedNative *IWrappedNativeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWrappedNative.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWrappedNative *IWrappedNativeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IWrappedNative.Contract.Allowance(&_IWrappedNative.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWrappedNative *IWrappedNativeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IWrappedNative.Contract.Allowance(&_IWrappedNative.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWrappedNative *IWrappedNativeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWrappedNative.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWrappedNative *IWrappedNativeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IWrappedNative.Contract.BalanceOf(&_IWrappedNative.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWrappedNative *IWrappedNativeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IWrappedNative.Contract.BalanceOf(&_IWrappedNative.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWrappedNative *IWrappedNativeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWrappedNative.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWrappedNative *IWrappedNativeSession) TotalSupply() (*big.Int, error) {
	return _IWrappedNative.Contract.TotalSupply(&_IWrappedNative.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWrappedNative *IWrappedNativeCallerSession) TotalSupply() (*big.Int, error) {
	return _IWrappedNative.Contract.TotalSupply(&_IWrappedNative.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWrappedNative *IWrappedNativeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWrappedNative.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWrappedNative *IWrappedNativeSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWrappedNative.Contract.Approve(&_IWrappedNative.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWrappedNative *IWrappedNativeTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWrappedNative.Contract.Approve(&_IWrappedNative.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWrappedNative *IWrappedNativeTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWrappedNative.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWrappedNative *IWrappedNativeSession) Deposit() (*types.Transaction, error) {
	return _IWrappedNative.Contract.Deposit(&_IWrappedNative.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWrappedNative *IWrappedNativeTransactorSession) Deposit() (*types.Transaction, error) {
	return _IWrappedNative.Contract.Deposit(&_IWrappedNative.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IWrappedNative *IWrappedNativeTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWrappedNative.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IWrappedNative *IWrappedNativeSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWrappedNative.Contract.Transfer(&_IWrappedNative.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IWrappedNative *IWrappedNativeTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWrappedNative.Contract.Transfer(&_IWrappedNative.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IWrappedNative *IWrappedNativeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWrappedNative.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IWrappedNative *IWrappedNativeSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWrappedNative.Contract.TransferFrom(&_IWrappedNative.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IWrappedNative *IWrappedNativeTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWrappedNative.Contract.TransferFrom(&_IWrappedNative.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWrappedNative *IWrappedNativeTransactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IWrappedNative.contract.Transact(opts, "withdraw", arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWrappedNative *IWrappedNativeSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWrappedNative.Contract.Withdraw(&_IWrappedNative.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWrappedNative *IWrappedNativeTransactorSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWrappedNative.Contract.Withdraw(&_IWrappedNative.TransactOpts, arg0)
}

// IWrappedNativeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IWrappedNative contract.
type IWrappedNativeApprovalIterator struct {
	Event *IWrappedNativeApproval // Event containing the contract specifics and raw log

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
func (it *IWrappedNativeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWrappedNativeApproval)
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
		it.Event = new(IWrappedNativeApproval)
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
func (it *IWrappedNativeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWrappedNativeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWrappedNativeApproval represents a Approval event raised by the IWrappedNative contract.
type IWrappedNativeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWrappedNative *IWrappedNativeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IWrappedNativeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IWrappedNative.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IWrappedNativeApprovalIterator{contract: _IWrappedNative.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWrappedNative *IWrappedNativeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IWrappedNativeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IWrappedNative.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWrappedNativeApproval)
				if err := _IWrappedNative.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IWrappedNative *IWrappedNativeFilterer) ParseApproval(log types.Log) (*IWrappedNativeApproval, error) {
	event := new(IWrappedNativeApproval)
	if err := _IWrappedNative.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWrappedNativeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IWrappedNative contract.
type IWrappedNativeTransferIterator struct {
	Event *IWrappedNativeTransfer // Event containing the contract specifics and raw log

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
func (it *IWrappedNativeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWrappedNativeTransfer)
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
		it.Event = new(IWrappedNativeTransfer)
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
func (it *IWrappedNativeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWrappedNativeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWrappedNativeTransfer represents a Transfer event raised by the IWrappedNative contract.
type IWrappedNativeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWrappedNative *IWrappedNativeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IWrappedNativeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IWrappedNative.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IWrappedNativeTransferIterator{contract: _IWrappedNative.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWrappedNative *IWrappedNativeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IWrappedNativeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IWrappedNative.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWrappedNativeTransfer)
				if err := _IWrappedNative.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IWrappedNative *IWrappedNativeFilterer) ParseTransfer(log types.Log) (*IWrappedNativeTransfer, error) {
	event := new(IWrappedNativeTransfer)
	if err := _IWrappedNative.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
