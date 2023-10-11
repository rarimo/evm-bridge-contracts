// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxy

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

// UUPSSignableUpgradeableMetaData contains all meta data concerning the UUPSSignableUpgradeable contract.
var UUPSSignableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"upgradeToWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UUPSSignableUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use UUPSSignableUpgradeableMetaData.ABI instead.
var UUPSSignableUpgradeableABI = UUPSSignableUpgradeableMetaData.ABI

// UUPSSignableUpgradeable is an auto generated Go binding around an Ethereum contract.
type UUPSSignableUpgradeable struct {
	UUPSSignableUpgradeableCaller     // Read-only binding to the contract
	UUPSSignableUpgradeableTransactor // Write-only binding to the contract
	UUPSSignableUpgradeableFilterer   // Log filterer for contract events
}

// UUPSSignableUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type UUPSSignableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UUPSSignableUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UUPSSignableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UUPSSignableUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UUPSSignableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UUPSSignableUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UUPSSignableUpgradeableSession struct {
	Contract     *UUPSSignableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UUPSSignableUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UUPSSignableUpgradeableCallerSession struct {
	Contract *UUPSSignableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// UUPSSignableUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UUPSSignableUpgradeableTransactorSession struct {
	Contract     *UUPSSignableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// UUPSSignableUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type UUPSSignableUpgradeableRaw struct {
	Contract *UUPSSignableUpgradeable // Generic contract binding to access the raw methods on
}

// UUPSSignableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UUPSSignableUpgradeableCallerRaw struct {
	Contract *UUPSSignableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// UUPSSignableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UUPSSignableUpgradeableTransactorRaw struct {
	Contract *UUPSSignableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUUPSSignableUpgradeable creates a new instance of UUPSSignableUpgradeable, bound to a specific deployed contract.
func NewUUPSSignableUpgradeable(address common.Address, backend bind.ContractBackend) (*UUPSSignableUpgradeable, error) {
	contract, err := bindUUPSSignableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UUPSSignableUpgradeable{UUPSSignableUpgradeableCaller: UUPSSignableUpgradeableCaller{contract: contract}, UUPSSignableUpgradeableTransactor: UUPSSignableUpgradeableTransactor{contract: contract}, UUPSSignableUpgradeableFilterer: UUPSSignableUpgradeableFilterer{contract: contract}}, nil
}

// NewUUPSSignableUpgradeableCaller creates a new read-only instance of UUPSSignableUpgradeable, bound to a specific deployed contract.
func NewUUPSSignableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*UUPSSignableUpgradeableCaller, error) {
	contract, err := bindUUPSSignableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UUPSSignableUpgradeableCaller{contract: contract}, nil
}

// NewUUPSSignableUpgradeableTransactor creates a new write-only instance of UUPSSignableUpgradeable, bound to a specific deployed contract.
func NewUUPSSignableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*UUPSSignableUpgradeableTransactor, error) {
	contract, err := bindUUPSSignableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UUPSSignableUpgradeableTransactor{contract: contract}, nil
}

// NewUUPSSignableUpgradeableFilterer creates a new log filterer instance of UUPSSignableUpgradeable, bound to a specific deployed contract.
func NewUUPSSignableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*UUPSSignableUpgradeableFilterer, error) {
	contract, err := bindUUPSSignableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UUPSSignableUpgradeableFilterer{contract: contract}, nil
}

// bindUUPSSignableUpgradeable binds a generic wrapper to an already deployed contract.
func bindUUPSSignableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UUPSSignableUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UUPSSignableUpgradeable.Contract.UUPSSignableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UUPSSignableUpgradeable.Contract.UUPSSignableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UUPSSignableUpgradeable.Contract.UUPSSignableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UUPSSignableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UUPSSignableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UUPSSignableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UUPSSignableUpgradeable.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableSession) ProxiableUUID() ([32]byte, error) {
	return _UUPSSignableUpgradeable.Contract.ProxiableUUID(&_UUPSSignableUpgradeable.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableCallerSession) ProxiableUUID() ([32]byte, error) {
	return _UUPSSignableUpgradeable.Contract.ProxiableUUID(&_UUPSSignableUpgradeable.CallOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _UUPSSignableUpgradeable.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UUPSSignableUpgradeable.Contract.UpgradeTo(&_UUPSSignableUpgradeable.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UUPSSignableUpgradeable.Contract.UpgradeTo(&_UUPSSignableUpgradeable.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UUPSSignableUpgradeable.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UUPSSignableUpgradeable.Contract.UpgradeToAndCall(&_UUPSSignableUpgradeable.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UUPSSignableUpgradeable.Contract.UpgradeToAndCall(&_UUPSSignableUpgradeable.TransactOpts, newImplementation, data)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableTransactor) UpgradeToWithSig(opts *bind.TransactOpts, newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _UUPSSignableUpgradeable.contract.Transact(opts, "upgradeToWithSig", newImplementation_, signature_)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableSession) UpgradeToWithSig(newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _UUPSSignableUpgradeable.Contract.UpgradeToWithSig(&_UUPSSignableUpgradeable.TransactOpts, newImplementation_, signature_)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableTransactorSession) UpgradeToWithSig(newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _UUPSSignableUpgradeable.Contract.UpgradeToWithSig(&_UUPSSignableUpgradeable.TransactOpts, newImplementation_, signature_)
}

// UUPSSignableUpgradeableAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the UUPSSignableUpgradeable contract.
type UUPSSignableUpgradeableAdminChangedIterator struct {
	Event *UUPSSignableUpgradeableAdminChanged // Event containing the contract specifics and raw log

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
func (it *UUPSSignableUpgradeableAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSSignableUpgradeableAdminChanged)
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
		it.Event = new(UUPSSignableUpgradeableAdminChanged)
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
func (it *UUPSSignableUpgradeableAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSSignableUpgradeableAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSSignableUpgradeableAdminChanged represents a AdminChanged event raised by the UUPSSignableUpgradeable contract.
type UUPSSignableUpgradeableAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*UUPSSignableUpgradeableAdminChangedIterator, error) {

	logs, sub, err := _UUPSSignableUpgradeable.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &UUPSSignableUpgradeableAdminChangedIterator{contract: _UUPSSignableUpgradeable.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *UUPSSignableUpgradeableAdminChanged) (event.Subscription, error) {

	logs, sub, err := _UUPSSignableUpgradeable.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSSignableUpgradeableAdminChanged)
				if err := _UUPSSignableUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableFilterer) ParseAdminChanged(log types.Log) (*UUPSSignableUpgradeableAdminChanged, error) {
	event := new(UUPSSignableUpgradeableAdminChanged)
	if err := _UUPSSignableUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UUPSSignableUpgradeableBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the UUPSSignableUpgradeable contract.
type UUPSSignableUpgradeableBeaconUpgradedIterator struct {
	Event *UUPSSignableUpgradeableBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *UUPSSignableUpgradeableBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSSignableUpgradeableBeaconUpgraded)
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
		it.Event = new(UUPSSignableUpgradeableBeaconUpgraded)
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
func (it *UUPSSignableUpgradeableBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSSignableUpgradeableBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSSignableUpgradeableBeaconUpgraded represents a BeaconUpgraded event raised by the UUPSSignableUpgradeable contract.
type UUPSSignableUpgradeableBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*UUPSSignableUpgradeableBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _UUPSSignableUpgradeable.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &UUPSSignableUpgradeableBeaconUpgradedIterator{contract: _UUPSSignableUpgradeable.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *UUPSSignableUpgradeableBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _UUPSSignableUpgradeable.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSSignableUpgradeableBeaconUpgraded)
				if err := _UUPSSignableUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableFilterer) ParseBeaconUpgraded(log types.Log) (*UUPSSignableUpgradeableBeaconUpgraded, error) {
	event := new(UUPSSignableUpgradeableBeaconUpgraded)
	if err := _UUPSSignableUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UUPSSignableUpgradeableUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the UUPSSignableUpgradeable contract.
type UUPSSignableUpgradeableUpgradedIterator struct {
	Event *UUPSSignableUpgradeableUpgraded // Event containing the contract specifics and raw log

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
func (it *UUPSSignableUpgradeableUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSSignableUpgradeableUpgraded)
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
		it.Event = new(UUPSSignableUpgradeableUpgraded)
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
func (it *UUPSSignableUpgradeableUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSSignableUpgradeableUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSSignableUpgradeableUpgraded represents a Upgraded event raised by the UUPSSignableUpgradeable contract.
type UUPSSignableUpgradeableUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*UUPSSignableUpgradeableUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UUPSSignableUpgradeable.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &UUPSSignableUpgradeableUpgradedIterator{contract: _UUPSSignableUpgradeable.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *UUPSSignableUpgradeableUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UUPSSignableUpgradeable.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSSignableUpgradeableUpgraded)
				if err := _UUPSSignableUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_UUPSSignableUpgradeable *UUPSSignableUpgradeableFilterer) ParseUpgraded(log types.Log) (*UUPSSignableUpgradeableUpgraded, error) {
	event := new(UUPSSignableUpgradeableUpgraded)
	if err := _UUPSSignableUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
