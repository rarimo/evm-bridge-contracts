// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sbt

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

// IERC5192MetaData contains all meta data concerning the IERC5192 contract.
var IERC5192MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IERC5192ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC5192MetaData.ABI instead.
var IERC5192ABI = IERC5192MetaData.ABI

// IERC5192 is an auto generated Go binding around an Ethereum contract.
type IERC5192 struct {
	IERC5192Caller     // Read-only binding to the contract
	IERC5192Transactor // Write-only binding to the contract
	IERC5192Filterer   // Log filterer for contract events
}

// IERC5192Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC5192Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC5192Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC5192Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC5192Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC5192Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC5192Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC5192Session struct {
	Contract     *IERC5192         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC5192CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC5192CallerSession struct {
	Contract *IERC5192Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IERC5192TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC5192TransactorSession struct {
	Contract     *IERC5192Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IERC5192Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC5192Raw struct {
	Contract *IERC5192 // Generic contract binding to access the raw methods on
}

// IERC5192CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC5192CallerRaw struct {
	Contract *IERC5192Caller // Generic read-only contract binding to access the raw methods on
}

// IERC5192TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC5192TransactorRaw struct {
	Contract *IERC5192Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC5192 creates a new instance of IERC5192, bound to a specific deployed contract.
func NewIERC5192(address common.Address, backend bind.ContractBackend) (*IERC5192, error) {
	contract, err := bindIERC5192(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC5192{IERC5192Caller: IERC5192Caller{contract: contract}, IERC5192Transactor: IERC5192Transactor{contract: contract}, IERC5192Filterer: IERC5192Filterer{contract: contract}}, nil
}

// NewIERC5192Caller creates a new read-only instance of IERC5192, bound to a specific deployed contract.
func NewIERC5192Caller(address common.Address, caller bind.ContractCaller) (*IERC5192Caller, error) {
	contract, err := bindIERC5192(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC5192Caller{contract: contract}, nil
}

// NewIERC5192Transactor creates a new write-only instance of IERC5192, bound to a specific deployed contract.
func NewIERC5192Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC5192Transactor, error) {
	contract, err := bindIERC5192(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC5192Transactor{contract: contract}, nil
}

// NewIERC5192Filterer creates a new log filterer instance of IERC5192, bound to a specific deployed contract.
func NewIERC5192Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC5192Filterer, error) {
	contract, err := bindIERC5192(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC5192Filterer{contract: contract}, nil
}

// bindIERC5192 binds a generic wrapper to an already deployed contract.
func bindIERC5192(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC5192MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC5192 *IERC5192Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC5192.Contract.IERC5192Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC5192 *IERC5192Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC5192.Contract.IERC5192Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC5192 *IERC5192Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC5192.Contract.IERC5192Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC5192 *IERC5192CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC5192.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC5192 *IERC5192TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC5192.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC5192 *IERC5192TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC5192.Contract.contract.Transact(opts, method, params...)
}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool)
func (_IERC5192 *IERC5192Caller) Locked(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _IERC5192.contract.Call(opts, &out, "locked", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool)
func (_IERC5192 *IERC5192Session) Locked(tokenId *big.Int) (bool, error) {
	return _IERC5192.Contract.Locked(&_IERC5192.CallOpts, tokenId)
}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool)
func (_IERC5192 *IERC5192CallerSession) Locked(tokenId *big.Int) (bool, error) {
	return _IERC5192.Contract.Locked(&_IERC5192.CallOpts, tokenId)
}

// IERC5192LockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the IERC5192 contract.
type IERC5192LockedIterator struct {
	Event *IERC5192Locked // Event containing the contract specifics and raw log

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
func (it *IERC5192LockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC5192Locked)
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
		it.Event = new(IERC5192Locked)
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
func (it *IERC5192LockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC5192LockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC5192Locked represents a Locked event raised by the IERC5192 contract.
type IERC5192Locked struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x032bc66be43dbccb7487781d168eb7bda224628a3b2c3388bdf69b532a3a1611.
//
// Solidity: event Locked(uint256 tokenId)
func (_IERC5192 *IERC5192Filterer) FilterLocked(opts *bind.FilterOpts) (*IERC5192LockedIterator, error) {

	logs, sub, err := _IERC5192.contract.FilterLogs(opts, "Locked")
	if err != nil {
		return nil, err
	}
	return &IERC5192LockedIterator{contract: _IERC5192.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x032bc66be43dbccb7487781d168eb7bda224628a3b2c3388bdf69b532a3a1611.
//
// Solidity: event Locked(uint256 tokenId)
func (_IERC5192 *IERC5192Filterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *IERC5192Locked) (event.Subscription, error) {

	logs, sub, err := _IERC5192.contract.WatchLogs(opts, "Locked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC5192Locked)
				if err := _IERC5192.contract.UnpackLog(event, "Locked", log); err != nil {
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

// ParseLocked is a log parse operation binding the contract event 0x032bc66be43dbccb7487781d168eb7bda224628a3b2c3388bdf69b532a3a1611.
//
// Solidity: event Locked(uint256 tokenId)
func (_IERC5192 *IERC5192Filterer) ParseLocked(log types.Log) (*IERC5192Locked, error) {
	event := new(IERC5192Locked)
	if err := _IERC5192.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
