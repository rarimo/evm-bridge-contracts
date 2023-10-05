// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package helpers

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

// ERC1967ProxyMockMetaData contains all meta data concerning the ERC1967ProxyMock contract.
var ERC1967ProxyMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"logic_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ERC1967ProxyMockABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1967ProxyMockMetaData.ABI instead.
var ERC1967ProxyMockABI = ERC1967ProxyMockMetaData.ABI

// ERC1967ProxyMock is an auto generated Go binding around an Ethereum contract.
type ERC1967ProxyMock struct {
	ERC1967ProxyMockCaller     // Read-only binding to the contract
	ERC1967ProxyMockTransactor // Write-only binding to the contract
	ERC1967ProxyMockFilterer   // Log filterer for contract events
}

// ERC1967ProxyMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1967ProxyMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967ProxyMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1967ProxyMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967ProxyMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1967ProxyMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967ProxyMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1967ProxyMockSession struct {
	Contract     *ERC1967ProxyMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1967ProxyMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1967ProxyMockCallerSession struct {
	Contract *ERC1967ProxyMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC1967ProxyMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1967ProxyMockTransactorSession struct {
	Contract     *ERC1967ProxyMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC1967ProxyMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1967ProxyMockRaw struct {
	Contract *ERC1967ProxyMock // Generic contract binding to access the raw methods on
}

// ERC1967ProxyMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1967ProxyMockCallerRaw struct {
	Contract *ERC1967ProxyMockCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1967ProxyMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1967ProxyMockTransactorRaw struct {
	Contract *ERC1967ProxyMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1967ProxyMock creates a new instance of ERC1967ProxyMock, bound to a specific deployed contract.
func NewERC1967ProxyMock(address common.Address, backend bind.ContractBackend) (*ERC1967ProxyMock, error) {
	contract, err := bindERC1967ProxyMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyMock{ERC1967ProxyMockCaller: ERC1967ProxyMockCaller{contract: contract}, ERC1967ProxyMockTransactor: ERC1967ProxyMockTransactor{contract: contract}, ERC1967ProxyMockFilterer: ERC1967ProxyMockFilterer{contract: contract}}, nil
}

// NewERC1967ProxyMockCaller creates a new read-only instance of ERC1967ProxyMock, bound to a specific deployed contract.
func NewERC1967ProxyMockCaller(address common.Address, caller bind.ContractCaller) (*ERC1967ProxyMockCaller, error) {
	contract, err := bindERC1967ProxyMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyMockCaller{contract: contract}, nil
}

// NewERC1967ProxyMockTransactor creates a new write-only instance of ERC1967ProxyMock, bound to a specific deployed contract.
func NewERC1967ProxyMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1967ProxyMockTransactor, error) {
	contract, err := bindERC1967ProxyMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyMockTransactor{contract: contract}, nil
}

// NewERC1967ProxyMockFilterer creates a new log filterer instance of ERC1967ProxyMock, bound to a specific deployed contract.
func NewERC1967ProxyMockFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1967ProxyMockFilterer, error) {
	contract, err := bindERC1967ProxyMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyMockFilterer{contract: contract}, nil
}

// bindERC1967ProxyMock binds a generic wrapper to an already deployed contract.
func bindERC1967ProxyMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC1967ProxyMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967ProxyMock *ERC1967ProxyMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967ProxyMock.Contract.ERC1967ProxyMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967ProxyMock *ERC1967ProxyMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967ProxyMock.Contract.ERC1967ProxyMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967ProxyMock *ERC1967ProxyMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967ProxyMock.Contract.ERC1967ProxyMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967ProxyMock *ERC1967ProxyMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967ProxyMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967ProxyMock *ERC1967ProxyMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967ProxyMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967ProxyMock *ERC1967ProxyMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967ProxyMock.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_ERC1967ProxyMock *ERC1967ProxyMockCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC1967ProxyMock.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_ERC1967ProxyMock *ERC1967ProxyMockSession) Implementation() (common.Address, error) {
	return _ERC1967ProxyMock.Contract.Implementation(&_ERC1967ProxyMock.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_ERC1967ProxyMock *ERC1967ProxyMockCallerSession) Implementation() (common.Address, error) {
	return _ERC1967ProxyMock.Contract.Implementation(&_ERC1967ProxyMock.CallOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ERC1967ProxyMock *ERC1967ProxyMockTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ERC1967ProxyMock.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ERC1967ProxyMock *ERC1967ProxyMockSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ERC1967ProxyMock.Contract.Fallback(&_ERC1967ProxyMock.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ERC1967ProxyMock *ERC1967ProxyMockTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ERC1967ProxyMock.Contract.Fallback(&_ERC1967ProxyMock.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ERC1967ProxyMock *ERC1967ProxyMockTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967ProxyMock.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ERC1967ProxyMock *ERC1967ProxyMockSession) Receive() (*types.Transaction, error) {
	return _ERC1967ProxyMock.Contract.Receive(&_ERC1967ProxyMock.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ERC1967ProxyMock *ERC1967ProxyMockTransactorSession) Receive() (*types.Transaction, error) {
	return _ERC1967ProxyMock.Contract.Receive(&_ERC1967ProxyMock.TransactOpts)
}

// ERC1967ProxyMockAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ERC1967ProxyMock contract.
type ERC1967ProxyMockAdminChangedIterator struct {
	Event *ERC1967ProxyMockAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC1967ProxyMockAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967ProxyMockAdminChanged)
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
		it.Event = new(ERC1967ProxyMockAdminChanged)
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
func (it *ERC1967ProxyMockAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967ProxyMockAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967ProxyMockAdminChanged represents a AdminChanged event raised by the ERC1967ProxyMock contract.
type ERC1967ProxyMockAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967ProxyMock *ERC1967ProxyMockFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ERC1967ProxyMockAdminChangedIterator, error) {

	logs, sub, err := _ERC1967ProxyMock.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyMockAdminChangedIterator{contract: _ERC1967ProxyMock.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967ProxyMock *ERC1967ProxyMockFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC1967ProxyMockAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ERC1967ProxyMock.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967ProxyMockAdminChanged)
				if err := _ERC1967ProxyMock.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967ProxyMock *ERC1967ProxyMockFilterer) ParseAdminChanged(log types.Log) (*ERC1967ProxyMockAdminChanged, error) {
	event := new(ERC1967ProxyMockAdminChanged)
	if err := _ERC1967ProxyMock.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967ProxyMockBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ERC1967ProxyMock contract.
type ERC1967ProxyMockBeaconUpgradedIterator struct {
	Event *ERC1967ProxyMockBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967ProxyMockBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967ProxyMockBeaconUpgraded)
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
		it.Event = new(ERC1967ProxyMockBeaconUpgraded)
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
func (it *ERC1967ProxyMockBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967ProxyMockBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967ProxyMockBeaconUpgraded represents a BeaconUpgraded event raised by the ERC1967ProxyMock contract.
type ERC1967ProxyMockBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967ProxyMock *ERC1967ProxyMockFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ERC1967ProxyMockBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967ProxyMock.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyMockBeaconUpgradedIterator{contract: _ERC1967ProxyMock.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967ProxyMock *ERC1967ProxyMockFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967ProxyMockBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967ProxyMock.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967ProxyMockBeaconUpgraded)
				if err := _ERC1967ProxyMock.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967ProxyMock *ERC1967ProxyMockFilterer) ParseBeaconUpgraded(log types.Log) (*ERC1967ProxyMockBeaconUpgraded, error) {
	event := new(ERC1967ProxyMockBeaconUpgraded)
	if err := _ERC1967ProxyMock.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967ProxyMockUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ERC1967ProxyMock contract.
type ERC1967ProxyMockUpgradedIterator struct {
	Event *ERC1967ProxyMockUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967ProxyMockUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967ProxyMockUpgraded)
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
		it.Event = new(ERC1967ProxyMockUpgraded)
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
func (it *ERC1967ProxyMockUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967ProxyMockUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967ProxyMockUpgraded represents a Upgraded event raised by the ERC1967ProxyMock contract.
type ERC1967ProxyMockUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967ProxyMock *ERC1967ProxyMockFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ERC1967ProxyMockUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967ProxyMock.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyMockUpgradedIterator{contract: _ERC1967ProxyMock.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967ProxyMock *ERC1967ProxyMockFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967ProxyMockUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967ProxyMock.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967ProxyMockUpgraded)
				if err := _ERC1967ProxyMock.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967ProxyMock *ERC1967ProxyMockFilterer) ParseUpgraded(log types.Log) (*ERC1967ProxyMockUpgraded, error) {
	event := new(ERC1967ProxyMockUpgraded)
	if err := _ERC1967ProxyMock.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
