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

// IERC20HandlerDepositERC20Parameters is an auto generated low-level Go binding around an user-defined struct.
type IERC20HandlerDepositERC20Parameters struct {
	Token     common.Address
	Amount    *big.Int
	Bundle    IBundlerBundle
	Network   string
	Receiver  string
	IsWrapped bool
}

// IERC20HandlerWithdrawERC20Parameters is an auto generated low-level Go binding around an user-defined struct.
type IERC20HandlerWithdrawERC20Parameters struct {
	Token      common.Address
	Amount     *big.Int
	Bundle     IBundlerBundle
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
	IsWrapped  bool
}

// IERC20HandlerMetaData contains all meta data concerning the IERC20Handler contract.
var IERC20HandlerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"DepositedERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"WithdrawnERC20\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.DepositERC20Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt_\",\"type\":\"bytes32\"}],\"name\":\"determineProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.WithdrawERC20Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.WithdrawERC20Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC20Bundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20HandlerMetaData.ABI instead.
var IERC20HandlerABI = IERC20HandlerMetaData.ABI

// IERC20Handler is an auto generated Go binding around an Ethereum contract.
type IERC20Handler struct {
	IERC20HandlerCaller     // Read-only binding to the contract
	IERC20HandlerTransactor // Write-only binding to the contract
	IERC20HandlerFilterer   // Log filterer for contract events
}

// IERC20HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20HandlerSession struct {
	Contract     *IERC20Handler    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20HandlerCallerSession struct {
	Contract *IERC20HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IERC20HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20HandlerTransactorSession struct {
	Contract     *IERC20HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IERC20HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20HandlerRaw struct {
	Contract *IERC20Handler // Generic contract binding to access the raw methods on
}

// IERC20HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20HandlerCallerRaw struct {
	Contract *IERC20HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20HandlerTransactorRaw struct {
	Contract *IERC20HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Handler creates a new instance of IERC20Handler, bound to a specific deployed contract.
func NewIERC20Handler(address common.Address, backend bind.ContractBackend) (*IERC20Handler, error) {
	contract, err := bindIERC20Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Handler{IERC20HandlerCaller: IERC20HandlerCaller{contract: contract}, IERC20HandlerTransactor: IERC20HandlerTransactor{contract: contract}, IERC20HandlerFilterer: IERC20HandlerFilterer{contract: contract}}, nil
}

// NewIERC20HandlerCaller creates a new read-only instance of IERC20Handler, bound to a specific deployed contract.
func NewIERC20HandlerCaller(address common.Address, caller bind.ContractCaller) (*IERC20HandlerCaller, error) {
	contract, err := bindIERC20Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20HandlerCaller{contract: contract}, nil
}

// NewIERC20HandlerTransactor creates a new write-only instance of IERC20Handler, bound to a specific deployed contract.
func NewIERC20HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20HandlerTransactor, error) {
	contract, err := bindIERC20Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20HandlerTransactor{contract: contract}, nil
}

// NewIERC20HandlerFilterer creates a new log filterer instance of IERC20Handler, bound to a specific deployed contract.
func NewIERC20HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20HandlerFilterer, error) {
	contract, err := bindIERC20Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20HandlerFilterer{contract: contract}, nil
}

// bindIERC20Handler binds a generic wrapper to an already deployed contract.
func bindIERC20Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20HandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Handler *IERC20HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Handler.Contract.IERC20HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Handler *IERC20HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Handler.Contract.IERC20HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Handler *IERC20HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Handler.Contract.IERC20HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Handler *IERC20HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Handler *IERC20HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Handler *IERC20HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Handler.Contract.contract.Transact(opts, method, params...)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_IERC20Handler *IERC20HandlerCaller) DetermineProxyAddress(opts *bind.CallOpts, salt_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IERC20Handler.contract.Call(opts, &out, "determineProxyAddress", salt_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_IERC20Handler *IERC20HandlerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _IERC20Handler.Contract.DetermineProxyAddress(&_IERC20Handler.CallOpts, salt_)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_IERC20Handler *IERC20HandlerCallerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _IERC20Handler.Contract.DetermineProxyAddress(&_IERC20Handler.CallOpts, salt_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xbebd5a0b.
//
// Solidity: function depositERC20((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IERC20Handler *IERC20HandlerTransactor) DepositERC20(opts *bind.TransactOpts, params_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _IERC20Handler.contract.Transact(opts, "depositERC20", params_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xbebd5a0b.
//
// Solidity: function depositERC20((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IERC20Handler *IERC20HandlerSession) DepositERC20(params_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _IERC20Handler.Contract.DepositERC20(&_IERC20Handler.TransactOpts, params_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xbebd5a0b.
//
// Solidity: function depositERC20((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IERC20Handler *IERC20HandlerTransactorSession) DepositERC20(params_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _IERC20Handler.Contract.DepositERC20(&_IERC20Handler.TransactOpts, params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC20Handler *IERC20HandlerTransactor) WithdrawERC20(opts *bind.TransactOpts, params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _IERC20Handler.contract.Transact(opts, "withdrawERC20", params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC20Handler *IERC20HandlerSession) WithdrawERC20(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _IERC20Handler.Contract.WithdrawERC20(&_IERC20Handler.TransactOpts, params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC20Handler *IERC20HandlerTransactorSession) WithdrawERC20(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _IERC20Handler.Contract.WithdrawERC20(&_IERC20Handler.TransactOpts, params_)
}

// WithdrawERC20Bundle is a paid mutator transaction binding the contract method 0xe53507bd.
//
// Solidity: function withdrawERC20Bundle((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC20Handler *IERC20HandlerTransactor) WithdrawERC20Bundle(opts *bind.TransactOpts, params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _IERC20Handler.contract.Transact(opts, "withdrawERC20Bundle", params_)
}

// WithdrawERC20Bundle is a paid mutator transaction binding the contract method 0xe53507bd.
//
// Solidity: function withdrawERC20Bundle((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC20Handler *IERC20HandlerSession) WithdrawERC20Bundle(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _IERC20Handler.Contract.WithdrawERC20Bundle(&_IERC20Handler.TransactOpts, params_)
}

// WithdrawERC20Bundle is a paid mutator transaction binding the contract method 0xe53507bd.
//
// Solidity: function withdrawERC20Bundle((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC20Handler *IERC20HandlerTransactorSession) WithdrawERC20Bundle(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _IERC20Handler.Contract.WithdrawERC20Bundle(&_IERC20Handler.TransactOpts, params_)
}

// IERC20HandlerDepositedERC20Iterator is returned from FilterDepositedERC20 and is used to iterate over the raw logs and unpacked data for DepositedERC20 events raised by the IERC20Handler contract.
type IERC20HandlerDepositedERC20Iterator struct {
	Event *IERC20HandlerDepositedERC20 // Event containing the contract specifics and raw log

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
func (it *IERC20HandlerDepositedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20HandlerDepositedERC20)
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
		it.Event = new(IERC20HandlerDepositedERC20)
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
func (it *IERC20HandlerDepositedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20HandlerDepositedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20HandlerDepositedERC20 represents a DepositedERC20 event raised by the IERC20Handler contract.
type IERC20HandlerDepositedERC20 struct {
	Token     common.Address
	Amount    *big.Int
	Salt      [32]byte
	Bundle    []byte
	Network   string
	Receiver  string
	IsWrapped bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositedERC20 is a free log retrieval operation binding the contract event 0x043d52f9acdd847f0210803c386559db9e09d492143f2072fe30ea62ff0bb639.
//
// Solidity: event DepositedERC20(address token, uint256 amount, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_IERC20Handler *IERC20HandlerFilterer) FilterDepositedERC20(opts *bind.FilterOpts) (*IERC20HandlerDepositedERC20Iterator, error) {

	logs, sub, err := _IERC20Handler.contract.FilterLogs(opts, "DepositedERC20")
	if err != nil {
		return nil, err
	}
	return &IERC20HandlerDepositedERC20Iterator{contract: _IERC20Handler.contract, event: "DepositedERC20", logs: logs, sub: sub}, nil
}

// WatchDepositedERC20 is a free log subscription operation binding the contract event 0x043d52f9acdd847f0210803c386559db9e09d492143f2072fe30ea62ff0bb639.
//
// Solidity: event DepositedERC20(address token, uint256 amount, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_IERC20Handler *IERC20HandlerFilterer) WatchDepositedERC20(opts *bind.WatchOpts, sink chan<- *IERC20HandlerDepositedERC20) (event.Subscription, error) {

	logs, sub, err := _IERC20Handler.contract.WatchLogs(opts, "DepositedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20HandlerDepositedERC20)
				if err := _IERC20Handler.contract.UnpackLog(event, "DepositedERC20", log); err != nil {
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

// ParseDepositedERC20 is a log parse operation binding the contract event 0x043d52f9acdd847f0210803c386559db9e09d492143f2072fe30ea62ff0bb639.
//
// Solidity: event DepositedERC20(address token, uint256 amount, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_IERC20Handler *IERC20HandlerFilterer) ParseDepositedERC20(log types.Log) (*IERC20HandlerDepositedERC20, error) {
	event := new(IERC20HandlerDepositedERC20)
	if err := _IERC20Handler.contract.UnpackLog(event, "DepositedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20HandlerWithdrawnERC20Iterator is returned from FilterWithdrawnERC20 and is used to iterate over the raw logs and unpacked data for WithdrawnERC20 events raised by the IERC20Handler contract.
type IERC20HandlerWithdrawnERC20Iterator struct {
	Event *IERC20HandlerWithdrawnERC20 // Event containing the contract specifics and raw log

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
func (it *IERC20HandlerWithdrawnERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20HandlerWithdrawnERC20)
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
		it.Event = new(IERC20HandlerWithdrawnERC20)
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
func (it *IERC20HandlerWithdrawnERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20HandlerWithdrawnERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20HandlerWithdrawnERC20 represents a WithdrawnERC20 event raised by the IERC20Handler contract.
type IERC20HandlerWithdrawnERC20 struct {
	Token      common.Address
	Amount     *big.Int
	Salt       [32]byte
	Bundle     []byte
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
	IsWrapped  bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnERC20 is a free log retrieval operation binding the contract event 0xff14005e9cd01bf3385cde89fb85a18453a0e3df47a3224d84ebab3fbd174128.
//
// Solidity: event WithdrawnERC20(address token, uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_IERC20Handler *IERC20HandlerFilterer) FilterWithdrawnERC20(opts *bind.FilterOpts) (*IERC20HandlerWithdrawnERC20Iterator, error) {

	logs, sub, err := _IERC20Handler.contract.FilterLogs(opts, "WithdrawnERC20")
	if err != nil {
		return nil, err
	}
	return &IERC20HandlerWithdrawnERC20Iterator{contract: _IERC20Handler.contract, event: "WithdrawnERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawnERC20 is a free log subscription operation binding the contract event 0xff14005e9cd01bf3385cde89fb85a18453a0e3df47a3224d84ebab3fbd174128.
//
// Solidity: event WithdrawnERC20(address token, uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_IERC20Handler *IERC20HandlerFilterer) WatchWithdrawnERC20(opts *bind.WatchOpts, sink chan<- *IERC20HandlerWithdrawnERC20) (event.Subscription, error) {

	logs, sub, err := _IERC20Handler.contract.WatchLogs(opts, "WithdrawnERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20HandlerWithdrawnERC20)
				if err := _IERC20Handler.contract.UnpackLog(event, "WithdrawnERC20", log); err != nil {
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

// ParseWithdrawnERC20 is a log parse operation binding the contract event 0xff14005e9cd01bf3385cde89fb85a18453a0e3df47a3224d84ebab3fbd174128.
//
// Solidity: event WithdrawnERC20(address token, uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_IERC20Handler *IERC20HandlerFilterer) ParseWithdrawnERC20(log types.Log) (*IERC20HandlerWithdrawnERC20, error) {
	event := new(IERC20HandlerWithdrawnERC20)
	if err := _IERC20Handler.contract.UnpackLog(event, "WithdrawnERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
