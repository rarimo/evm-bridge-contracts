// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package handlers

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

// IBundlerBundle is an auto generated low-level Go binding around an user-defined struct.
type IBundlerBundle struct {
	Salt   [32]byte
	Bundle []byte
}

// INativeHandlerDepositNativeParameters is an auto generated low-level Go binding around an user-defined struct.
type INativeHandlerDepositNativeParameters struct {
	Amount   *big.Int
	Bundle   IBundlerBundle
	Network  string
	Receiver string
}

// INativeHandlerWithdrawNativeParameters is an auto generated low-level Go binding around an user-defined struct.
type INativeHandlerWithdrawNativeParameters struct {
	Amount     *big.Int
	Bundle     IBundlerBundle
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
}

// INativeHandlerMetaData contains all meta data concerning the INativeHandler contract.
var INativeHandlerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"name\":\"DepositedNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"WithdrawnNative\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"internalType\":\"structINativeHandler.DepositNativeParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt_\",\"type\":\"bytes32\"}],\"name\":\"determineProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structINativeHandler.WithdrawNativeParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structINativeHandler.WithdrawNativeParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawNativeBundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// INativeHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use INativeHandlerMetaData.ABI instead.
var INativeHandlerABI = INativeHandlerMetaData.ABI

// INativeHandler is an auto generated Go binding around an Ethereum contract.
type INativeHandler struct {
	INativeHandlerCaller     // Read-only binding to the contract
	INativeHandlerTransactor // Write-only binding to the contract
	INativeHandlerFilterer   // Log filterer for contract events
}

// INativeHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type INativeHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INativeHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INativeHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INativeHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INativeHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INativeHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INativeHandlerSession struct {
	Contract     *INativeHandler   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INativeHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INativeHandlerCallerSession struct {
	Contract *INativeHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// INativeHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INativeHandlerTransactorSession struct {
	Contract     *INativeHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// INativeHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type INativeHandlerRaw struct {
	Contract *INativeHandler // Generic contract binding to access the raw methods on
}

// INativeHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INativeHandlerCallerRaw struct {
	Contract *INativeHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// INativeHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INativeHandlerTransactorRaw struct {
	Contract *INativeHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINativeHandler creates a new instance of INativeHandler, bound to a specific deployed contract.
func NewINativeHandler(address common.Address, backend bind.ContractBackend) (*INativeHandler, error) {
	contract, err := bindINativeHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INativeHandler{INativeHandlerCaller: INativeHandlerCaller{contract: contract}, INativeHandlerTransactor: INativeHandlerTransactor{contract: contract}, INativeHandlerFilterer: INativeHandlerFilterer{contract: contract}}, nil
}

// NewINativeHandlerCaller creates a new read-only instance of INativeHandler, bound to a specific deployed contract.
func NewINativeHandlerCaller(address common.Address, caller bind.ContractCaller) (*INativeHandlerCaller, error) {
	contract, err := bindINativeHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INativeHandlerCaller{contract: contract}, nil
}

// NewINativeHandlerTransactor creates a new write-only instance of INativeHandler, bound to a specific deployed contract.
func NewINativeHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*INativeHandlerTransactor, error) {
	contract, err := bindINativeHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INativeHandlerTransactor{contract: contract}, nil
}

// NewINativeHandlerFilterer creates a new log filterer instance of INativeHandler, bound to a specific deployed contract.
func NewINativeHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*INativeHandlerFilterer, error) {
	contract, err := bindINativeHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INativeHandlerFilterer{contract: contract}, nil
}

// bindINativeHandler binds a generic wrapper to an already deployed contract.
func bindINativeHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := INativeHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INativeHandler *INativeHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INativeHandler.Contract.INativeHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INativeHandler *INativeHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INativeHandler.Contract.INativeHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INativeHandler *INativeHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INativeHandler.Contract.INativeHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INativeHandler *INativeHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INativeHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INativeHandler *INativeHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INativeHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INativeHandler *INativeHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INativeHandler.Contract.contract.Transact(opts, method, params...)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_INativeHandler *INativeHandlerCaller) DetermineProxyAddress(opts *bind.CallOpts, salt_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _INativeHandler.contract.Call(opts, &out, "determineProxyAddress", salt_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_INativeHandler *INativeHandlerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _INativeHandler.Contract.DetermineProxyAddress(&_INativeHandler.CallOpts, salt_)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_INativeHandler *INativeHandlerCallerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _INativeHandler.Contract.DetermineProxyAddress(&_INativeHandler.CallOpts, salt_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xb27588e5.
//
// Solidity: function depositNative((uint256,(bytes32,bytes),string,string) params_) payable returns()
func (_INativeHandler *INativeHandlerTransactor) DepositNative(opts *bind.TransactOpts, params_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _INativeHandler.contract.Transact(opts, "depositNative", params_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xb27588e5.
//
// Solidity: function depositNative((uint256,(bytes32,bytes),string,string) params_) payable returns()
func (_INativeHandler *INativeHandlerSession) DepositNative(params_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _INativeHandler.Contract.DepositNative(&_INativeHandler.TransactOpts, params_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xb27588e5.
//
// Solidity: function depositNative((uint256,(bytes32,bytes),string,string) params_) payable returns()
func (_INativeHandler *INativeHandlerTransactorSession) DepositNative(params_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _INativeHandler.Contract.DepositNative(&_INativeHandler.TransactOpts, params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_INativeHandler *INativeHandlerTransactor) WithdrawNative(opts *bind.TransactOpts, params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _INativeHandler.contract.Transact(opts, "withdrawNative", params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_INativeHandler *INativeHandlerSession) WithdrawNative(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _INativeHandler.Contract.WithdrawNative(&_INativeHandler.TransactOpts, params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_INativeHandler *INativeHandlerTransactorSession) WithdrawNative(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _INativeHandler.Contract.WithdrawNative(&_INativeHandler.TransactOpts, params_)
}

// WithdrawNativeBundle is a paid mutator transaction binding the contract method 0xc4f54e21.
//
// Solidity: function withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_INativeHandler *INativeHandlerTransactor) WithdrawNativeBundle(opts *bind.TransactOpts, params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _INativeHandler.contract.Transact(opts, "withdrawNativeBundle", params_)
}

// WithdrawNativeBundle is a paid mutator transaction binding the contract method 0xc4f54e21.
//
// Solidity: function withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_INativeHandler *INativeHandlerSession) WithdrawNativeBundle(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _INativeHandler.Contract.WithdrawNativeBundle(&_INativeHandler.TransactOpts, params_)
}

// WithdrawNativeBundle is a paid mutator transaction binding the contract method 0xc4f54e21.
//
// Solidity: function withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_INativeHandler *INativeHandlerTransactorSession) WithdrawNativeBundle(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _INativeHandler.Contract.WithdrawNativeBundle(&_INativeHandler.TransactOpts, params_)
}

// INativeHandlerDepositedNativeIterator is returned from FilterDepositedNative and is used to iterate over the raw logs and unpacked data for DepositedNative events raised by the INativeHandler contract.
type INativeHandlerDepositedNativeIterator struct {
	Event *INativeHandlerDepositedNative // Event containing the contract specifics and raw log

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
func (it *INativeHandlerDepositedNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeHandlerDepositedNative)
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
		it.Event = new(INativeHandlerDepositedNative)
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
func (it *INativeHandlerDepositedNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeHandlerDepositedNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeHandlerDepositedNative represents a DepositedNative event raised by the INativeHandler contract.
type INativeHandlerDepositedNative struct {
	Amount   *big.Int
	Salt     [32]byte
	Bundle   []byte
	Network  string
	Receiver string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDepositedNative is a free log retrieval operation binding the contract event 0x9a47c8733424880a9e86a368eff95da5e7d36b68474a95eb097be2e43c116f27.
//
// Solidity: event DepositedNative(uint256 amount, bytes32 salt, bytes bundle, string network, string receiver)
func (_INativeHandler *INativeHandlerFilterer) FilterDepositedNative(opts *bind.FilterOpts) (*INativeHandlerDepositedNativeIterator, error) {

	logs, sub, err := _INativeHandler.contract.FilterLogs(opts, "DepositedNative")
	if err != nil {
		return nil, err
	}
	return &INativeHandlerDepositedNativeIterator{contract: _INativeHandler.contract, event: "DepositedNative", logs: logs, sub: sub}, nil
}

// WatchDepositedNative is a free log subscription operation binding the contract event 0x9a47c8733424880a9e86a368eff95da5e7d36b68474a95eb097be2e43c116f27.
//
// Solidity: event DepositedNative(uint256 amount, bytes32 salt, bytes bundle, string network, string receiver)
func (_INativeHandler *INativeHandlerFilterer) WatchDepositedNative(opts *bind.WatchOpts, sink chan<- *INativeHandlerDepositedNative) (event.Subscription, error) {

	logs, sub, err := _INativeHandler.contract.WatchLogs(opts, "DepositedNative")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeHandlerDepositedNative)
				if err := _INativeHandler.contract.UnpackLog(event, "DepositedNative", log); err != nil {
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

// ParseDepositedNative is a log parse operation binding the contract event 0x9a47c8733424880a9e86a368eff95da5e7d36b68474a95eb097be2e43c116f27.
//
// Solidity: event DepositedNative(uint256 amount, bytes32 salt, bytes bundle, string network, string receiver)
func (_INativeHandler *INativeHandlerFilterer) ParseDepositedNative(log types.Log) (*INativeHandlerDepositedNative, error) {
	event := new(INativeHandlerDepositedNative)
	if err := _INativeHandler.contract.UnpackLog(event, "DepositedNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeHandlerWithdrawnNativeIterator is returned from FilterWithdrawnNative and is used to iterate over the raw logs and unpacked data for WithdrawnNative events raised by the INativeHandler contract.
type INativeHandlerWithdrawnNativeIterator struct {
	Event *INativeHandlerWithdrawnNative // Event containing the contract specifics and raw log

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
func (it *INativeHandlerWithdrawnNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeHandlerWithdrawnNative)
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
		it.Event = new(INativeHandlerWithdrawnNative)
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
func (it *INativeHandlerWithdrawnNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeHandlerWithdrawnNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeHandlerWithdrawnNative represents a WithdrawnNative event raised by the INativeHandler contract.
type INativeHandlerWithdrawnNative struct {
	Amount     *big.Int
	Salt       [32]byte
	Bundle     []byte
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnNative is a free log retrieval operation binding the contract event 0x319eed2457c60bc7ee23fdc5be5941c7b628c01f6f2c526dedeca75054d66b18.
//
// Solidity: event WithdrawnNative(uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_INativeHandler *INativeHandlerFilterer) FilterWithdrawnNative(opts *bind.FilterOpts) (*INativeHandlerWithdrawnNativeIterator, error) {

	logs, sub, err := _INativeHandler.contract.FilterLogs(opts, "WithdrawnNative")
	if err != nil {
		return nil, err
	}
	return &INativeHandlerWithdrawnNativeIterator{contract: _INativeHandler.contract, event: "WithdrawnNative", logs: logs, sub: sub}, nil
}

// WatchWithdrawnNative is a free log subscription operation binding the contract event 0x319eed2457c60bc7ee23fdc5be5941c7b628c01f6f2c526dedeca75054d66b18.
//
// Solidity: event WithdrawnNative(uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_INativeHandler *INativeHandlerFilterer) WatchWithdrawnNative(opts *bind.WatchOpts, sink chan<- *INativeHandlerWithdrawnNative) (event.Subscription, error) {

	logs, sub, err := _INativeHandler.contract.WatchLogs(opts, "WithdrawnNative")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeHandlerWithdrawnNative)
				if err := _INativeHandler.contract.UnpackLog(event, "WithdrawnNative", log); err != nil {
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

// ParseWithdrawnNative is a log parse operation binding the contract event 0x319eed2457c60bc7ee23fdc5be5941c7b628c01f6f2c526dedeca75054d66b18.
//
// Solidity: event WithdrawnNative(uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_INativeHandler *INativeHandlerFilterer) ParseWithdrawnNative(log types.Log) (*INativeHandlerWithdrawnNative, error) {
	event := new(INativeHandlerWithdrawnNative)
	if err := _INativeHandler.contract.UnpackLog(event, "WithdrawnNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
