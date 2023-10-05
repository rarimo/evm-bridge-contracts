// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package facade

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

// IFeeManagerAddFeeTokenParameters is an auto generated low-level Go binding around an user-defined struct.
type IFeeManagerAddFeeTokenParameters struct {
	FeeTokens  []common.Address
	FeeAmounts []*big.Int
	Signature  []byte
}

// IFeeManagerRemoveFeeTokenParameters is an auto generated low-level Go binding around an user-defined struct.
type IFeeManagerRemoveFeeTokenParameters struct {
	FeeTokens  []common.Address
	FeeAmounts []*big.Int
	Signature  []byte
}

// IFeeManagerUpdateFeeTokenParameters is an auto generated low-level Go binding around an user-defined struct.
type IFeeManagerUpdateFeeTokenParameters struct {
	FeeTokens  []common.Address
	FeeAmounts []*big.Int
	Signature  []byte
}

// IFeeManagerWithdrawFeeTokenParameters is an auto generated low-level Go binding around an user-defined struct.
type IFeeManagerWithdrawFeeTokenParameters struct {
	Receiver  common.Address
	FeeTokens []common.Address
	Amounts   []*big.Int
	Signature []byte
}

// IFeeManagerMetaData contains all meta data concerning the IFeeManager contract.
var IFeeManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"AddedFeeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"RemovedFeeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"UpdatedFeeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnFeeToken\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.AddFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"addFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken_\",\"type\":\"address\"}],\"name\":\"getCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"commission_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.RemoveFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"removeFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.UpdateFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"updateFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.WithdrawFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IFeeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IFeeManagerMetaData.ABI instead.
var IFeeManagerABI = IFeeManagerMetaData.ABI

// IFeeManager is an auto generated Go binding around an Ethereum contract.
type IFeeManager struct {
	IFeeManagerCaller     // Read-only binding to the contract
	IFeeManagerTransactor // Write-only binding to the contract
	IFeeManagerFilterer   // Log filterer for contract events
}

// IFeeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFeeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFeeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFeeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFeeManagerSession struct {
	Contract     *IFeeManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFeeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFeeManagerCallerSession struct {
	Contract *IFeeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IFeeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFeeManagerTransactorSession struct {
	Contract     *IFeeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IFeeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFeeManagerRaw struct {
	Contract *IFeeManager // Generic contract binding to access the raw methods on
}

// IFeeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFeeManagerCallerRaw struct {
	Contract *IFeeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IFeeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFeeManagerTransactorRaw struct {
	Contract *IFeeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFeeManager creates a new instance of IFeeManager, bound to a specific deployed contract.
func NewIFeeManager(address common.Address, backend bind.ContractBackend) (*IFeeManager, error) {
	contract, err := bindIFeeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFeeManager{IFeeManagerCaller: IFeeManagerCaller{contract: contract}, IFeeManagerTransactor: IFeeManagerTransactor{contract: contract}, IFeeManagerFilterer: IFeeManagerFilterer{contract: contract}}, nil
}

// NewIFeeManagerCaller creates a new read-only instance of IFeeManager, bound to a specific deployed contract.
func NewIFeeManagerCaller(address common.Address, caller bind.ContractCaller) (*IFeeManagerCaller, error) {
	contract, err := bindIFeeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFeeManagerCaller{contract: contract}, nil
}

// NewIFeeManagerTransactor creates a new write-only instance of IFeeManager, bound to a specific deployed contract.
func NewIFeeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IFeeManagerTransactor, error) {
	contract, err := bindIFeeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFeeManagerTransactor{contract: contract}, nil
}

// NewIFeeManagerFilterer creates a new log filterer instance of IFeeManager, bound to a specific deployed contract.
func NewIFeeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IFeeManagerFilterer, error) {
	contract, err := bindIFeeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFeeManagerFilterer{contract: contract}, nil
}

// bindIFeeManager binds a generic wrapper to an already deployed contract.
func bindIFeeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IFeeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFeeManager *IFeeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFeeManager.Contract.IFeeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFeeManager *IFeeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFeeManager.Contract.IFeeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFeeManager *IFeeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFeeManager.Contract.IFeeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFeeManager *IFeeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFeeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFeeManager *IFeeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFeeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFeeManager *IFeeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFeeManager.Contract.contract.Transact(opts, method, params...)
}

// GetCommission is a free data retrieval call binding the contract method 0xf334dcb2.
//
// Solidity: function getCommission(address feeToken_) view returns(uint256 commission_)
func (_IFeeManager *IFeeManagerCaller) GetCommission(opts *bind.CallOpts, feeToken_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IFeeManager.contract.Call(opts, &out, "getCommission", feeToken_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommission is a free data retrieval call binding the contract method 0xf334dcb2.
//
// Solidity: function getCommission(address feeToken_) view returns(uint256 commission_)
func (_IFeeManager *IFeeManagerSession) GetCommission(feeToken_ common.Address) (*big.Int, error) {
	return _IFeeManager.Contract.GetCommission(&_IFeeManager.CallOpts, feeToken_)
}

// GetCommission is a free data retrieval call binding the contract method 0xf334dcb2.
//
// Solidity: function getCommission(address feeToken_) view returns(uint256 commission_)
func (_IFeeManager *IFeeManagerCallerSession) GetCommission(feeToken_ common.Address) (*big.Int, error) {
	return _IFeeManager.Contract.GetCommission(&_IFeeManager.CallOpts, feeToken_)
}

// AddFeeToken is a paid mutator transaction binding the contract method 0x4915b86e.
//
// Solidity: function addFeeToken((address[],uint256[],bytes) params_) returns()
func (_IFeeManager *IFeeManagerTransactor) AddFeeToken(opts *bind.TransactOpts, params_ IFeeManagerAddFeeTokenParameters) (*types.Transaction, error) {
	return _IFeeManager.contract.Transact(opts, "addFeeToken", params_)
}

// AddFeeToken is a paid mutator transaction binding the contract method 0x4915b86e.
//
// Solidity: function addFeeToken((address[],uint256[],bytes) params_) returns()
func (_IFeeManager *IFeeManagerSession) AddFeeToken(params_ IFeeManagerAddFeeTokenParameters) (*types.Transaction, error) {
	return _IFeeManager.Contract.AddFeeToken(&_IFeeManager.TransactOpts, params_)
}

// AddFeeToken is a paid mutator transaction binding the contract method 0x4915b86e.
//
// Solidity: function addFeeToken((address[],uint256[],bytes) params_) returns()
func (_IFeeManager *IFeeManagerTransactorSession) AddFeeToken(params_ IFeeManagerAddFeeTokenParameters) (*types.Transaction, error) {
	return _IFeeManager.Contract.AddFeeToken(&_IFeeManager.TransactOpts, params_)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x4eb5943a.
//
// Solidity: function removeFeeToken((address[],uint256[],bytes) params_) returns()
func (_IFeeManager *IFeeManagerTransactor) RemoveFeeToken(opts *bind.TransactOpts, params_ IFeeManagerRemoveFeeTokenParameters) (*types.Transaction, error) {
	return _IFeeManager.contract.Transact(opts, "removeFeeToken", params_)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x4eb5943a.
//
// Solidity: function removeFeeToken((address[],uint256[],bytes) params_) returns()
func (_IFeeManager *IFeeManagerSession) RemoveFeeToken(params_ IFeeManagerRemoveFeeTokenParameters) (*types.Transaction, error) {
	return _IFeeManager.Contract.RemoveFeeToken(&_IFeeManager.TransactOpts, params_)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x4eb5943a.
//
// Solidity: function removeFeeToken((address[],uint256[],bytes) params_) returns()
func (_IFeeManager *IFeeManagerTransactorSession) RemoveFeeToken(params_ IFeeManagerRemoveFeeTokenParameters) (*types.Transaction, error) {
	return _IFeeManager.Contract.RemoveFeeToken(&_IFeeManager.TransactOpts, params_)
}

// UpdateFeeToken is a paid mutator transaction binding the contract method 0x63925ea2.
//
// Solidity: function updateFeeToken((address[],uint256[],bytes) params_) returns()
func (_IFeeManager *IFeeManagerTransactor) UpdateFeeToken(opts *bind.TransactOpts, params_ IFeeManagerUpdateFeeTokenParameters) (*types.Transaction, error) {
	return _IFeeManager.contract.Transact(opts, "updateFeeToken", params_)
}

// UpdateFeeToken is a paid mutator transaction binding the contract method 0x63925ea2.
//
// Solidity: function updateFeeToken((address[],uint256[],bytes) params_) returns()
func (_IFeeManager *IFeeManagerSession) UpdateFeeToken(params_ IFeeManagerUpdateFeeTokenParameters) (*types.Transaction, error) {
	return _IFeeManager.Contract.UpdateFeeToken(&_IFeeManager.TransactOpts, params_)
}

// UpdateFeeToken is a paid mutator transaction binding the contract method 0x63925ea2.
//
// Solidity: function updateFeeToken((address[],uint256[],bytes) params_) returns()
func (_IFeeManager *IFeeManagerTransactorSession) UpdateFeeToken(params_ IFeeManagerUpdateFeeTokenParameters) (*types.Transaction, error) {
	return _IFeeManager.Contract.UpdateFeeToken(&_IFeeManager.TransactOpts, params_)
}

// WithdrawFeeToken is a paid mutator transaction binding the contract method 0x9f33cf22.
//
// Solidity: function withdrawFeeToken((address,address[],uint256[],bytes) params_) returns()
func (_IFeeManager *IFeeManagerTransactor) WithdrawFeeToken(opts *bind.TransactOpts, params_ IFeeManagerWithdrawFeeTokenParameters) (*types.Transaction, error) {
	return _IFeeManager.contract.Transact(opts, "withdrawFeeToken", params_)
}

// WithdrawFeeToken is a paid mutator transaction binding the contract method 0x9f33cf22.
//
// Solidity: function withdrawFeeToken((address,address[],uint256[],bytes) params_) returns()
func (_IFeeManager *IFeeManagerSession) WithdrawFeeToken(params_ IFeeManagerWithdrawFeeTokenParameters) (*types.Transaction, error) {
	return _IFeeManager.Contract.WithdrawFeeToken(&_IFeeManager.TransactOpts, params_)
}

// WithdrawFeeToken is a paid mutator transaction binding the contract method 0x9f33cf22.
//
// Solidity: function withdrawFeeToken((address,address[],uint256[],bytes) params_) returns()
func (_IFeeManager *IFeeManagerTransactorSession) WithdrawFeeToken(params_ IFeeManagerWithdrawFeeTokenParameters) (*types.Transaction, error) {
	return _IFeeManager.Contract.WithdrawFeeToken(&_IFeeManager.TransactOpts, params_)
}

// IFeeManagerAddedFeeTokenIterator is returned from FilterAddedFeeToken and is used to iterate over the raw logs and unpacked data for AddedFeeToken events raised by the IFeeManager contract.
type IFeeManagerAddedFeeTokenIterator struct {
	Event *IFeeManagerAddedFeeToken // Event containing the contract specifics and raw log

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
func (it *IFeeManagerAddedFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFeeManagerAddedFeeToken)
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
		it.Event = new(IFeeManagerAddedFeeToken)
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
func (it *IFeeManagerAddedFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFeeManagerAddedFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFeeManagerAddedFeeToken represents a AddedFeeToken event raised by the IFeeManager contract.
type IFeeManagerAddedFeeToken struct {
	FeeToken  common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddedFeeToken is a free log retrieval operation binding the contract event 0x5d966ab5c51224ee3c5146e53a0b57969d97d0e1779f52a1017813c13344efc4.
//
// Solidity: event AddedFeeToken(address feeToken, uint256 feeAmount)
func (_IFeeManager *IFeeManagerFilterer) FilterAddedFeeToken(opts *bind.FilterOpts) (*IFeeManagerAddedFeeTokenIterator, error) {

	logs, sub, err := _IFeeManager.contract.FilterLogs(opts, "AddedFeeToken")
	if err != nil {
		return nil, err
	}
	return &IFeeManagerAddedFeeTokenIterator{contract: _IFeeManager.contract, event: "AddedFeeToken", logs: logs, sub: sub}, nil
}

// WatchAddedFeeToken is a free log subscription operation binding the contract event 0x5d966ab5c51224ee3c5146e53a0b57969d97d0e1779f52a1017813c13344efc4.
//
// Solidity: event AddedFeeToken(address feeToken, uint256 feeAmount)
func (_IFeeManager *IFeeManagerFilterer) WatchAddedFeeToken(opts *bind.WatchOpts, sink chan<- *IFeeManagerAddedFeeToken) (event.Subscription, error) {

	logs, sub, err := _IFeeManager.contract.WatchLogs(opts, "AddedFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFeeManagerAddedFeeToken)
				if err := _IFeeManager.contract.UnpackLog(event, "AddedFeeToken", log); err != nil {
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

// ParseAddedFeeToken is a log parse operation binding the contract event 0x5d966ab5c51224ee3c5146e53a0b57969d97d0e1779f52a1017813c13344efc4.
//
// Solidity: event AddedFeeToken(address feeToken, uint256 feeAmount)
func (_IFeeManager *IFeeManagerFilterer) ParseAddedFeeToken(log types.Log) (*IFeeManagerAddedFeeToken, error) {
	event := new(IFeeManagerAddedFeeToken)
	if err := _IFeeManager.contract.UnpackLog(event, "AddedFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFeeManagerRemovedFeeTokenIterator is returned from FilterRemovedFeeToken and is used to iterate over the raw logs and unpacked data for RemovedFeeToken events raised by the IFeeManager contract.
type IFeeManagerRemovedFeeTokenIterator struct {
	Event *IFeeManagerRemovedFeeToken // Event containing the contract specifics and raw log

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
func (it *IFeeManagerRemovedFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFeeManagerRemovedFeeToken)
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
		it.Event = new(IFeeManagerRemovedFeeToken)
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
func (it *IFeeManagerRemovedFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFeeManagerRemovedFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFeeManagerRemovedFeeToken represents a RemovedFeeToken event raised by the IFeeManager contract.
type IFeeManagerRemovedFeeToken struct {
	FeeToken  common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemovedFeeToken is a free log retrieval operation binding the contract event 0xeac220a493a0693e32f1240aae311de65b2d84391a9ba8310fcaf359d7c9bd8f.
//
// Solidity: event RemovedFeeToken(address feeToken, uint256 feeAmount)
func (_IFeeManager *IFeeManagerFilterer) FilterRemovedFeeToken(opts *bind.FilterOpts) (*IFeeManagerRemovedFeeTokenIterator, error) {

	logs, sub, err := _IFeeManager.contract.FilterLogs(opts, "RemovedFeeToken")
	if err != nil {
		return nil, err
	}
	return &IFeeManagerRemovedFeeTokenIterator{contract: _IFeeManager.contract, event: "RemovedFeeToken", logs: logs, sub: sub}, nil
}

// WatchRemovedFeeToken is a free log subscription operation binding the contract event 0xeac220a493a0693e32f1240aae311de65b2d84391a9ba8310fcaf359d7c9bd8f.
//
// Solidity: event RemovedFeeToken(address feeToken, uint256 feeAmount)
func (_IFeeManager *IFeeManagerFilterer) WatchRemovedFeeToken(opts *bind.WatchOpts, sink chan<- *IFeeManagerRemovedFeeToken) (event.Subscription, error) {

	logs, sub, err := _IFeeManager.contract.WatchLogs(opts, "RemovedFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFeeManagerRemovedFeeToken)
				if err := _IFeeManager.contract.UnpackLog(event, "RemovedFeeToken", log); err != nil {
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

// ParseRemovedFeeToken is a log parse operation binding the contract event 0xeac220a493a0693e32f1240aae311de65b2d84391a9ba8310fcaf359d7c9bd8f.
//
// Solidity: event RemovedFeeToken(address feeToken, uint256 feeAmount)
func (_IFeeManager *IFeeManagerFilterer) ParseRemovedFeeToken(log types.Log) (*IFeeManagerRemovedFeeToken, error) {
	event := new(IFeeManagerRemovedFeeToken)
	if err := _IFeeManager.contract.UnpackLog(event, "RemovedFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFeeManagerUpdatedFeeTokenIterator is returned from FilterUpdatedFeeToken and is used to iterate over the raw logs and unpacked data for UpdatedFeeToken events raised by the IFeeManager contract.
type IFeeManagerUpdatedFeeTokenIterator struct {
	Event *IFeeManagerUpdatedFeeToken // Event containing the contract specifics and raw log

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
func (it *IFeeManagerUpdatedFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFeeManagerUpdatedFeeToken)
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
		it.Event = new(IFeeManagerUpdatedFeeToken)
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
func (it *IFeeManagerUpdatedFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFeeManagerUpdatedFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFeeManagerUpdatedFeeToken represents a UpdatedFeeToken event raised by the IFeeManager contract.
type IFeeManagerUpdatedFeeToken struct {
	FeeToken  common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedFeeToken is a free log retrieval operation binding the contract event 0x70698787bb7456af2d4339ba1e7d059c4eda7c473412e7ed2c0d5b361a49d50a.
//
// Solidity: event UpdatedFeeToken(address feeToken, uint256 feeAmount)
func (_IFeeManager *IFeeManagerFilterer) FilterUpdatedFeeToken(opts *bind.FilterOpts) (*IFeeManagerUpdatedFeeTokenIterator, error) {

	logs, sub, err := _IFeeManager.contract.FilterLogs(opts, "UpdatedFeeToken")
	if err != nil {
		return nil, err
	}
	return &IFeeManagerUpdatedFeeTokenIterator{contract: _IFeeManager.contract, event: "UpdatedFeeToken", logs: logs, sub: sub}, nil
}

// WatchUpdatedFeeToken is a free log subscription operation binding the contract event 0x70698787bb7456af2d4339ba1e7d059c4eda7c473412e7ed2c0d5b361a49d50a.
//
// Solidity: event UpdatedFeeToken(address feeToken, uint256 feeAmount)
func (_IFeeManager *IFeeManagerFilterer) WatchUpdatedFeeToken(opts *bind.WatchOpts, sink chan<- *IFeeManagerUpdatedFeeToken) (event.Subscription, error) {

	logs, sub, err := _IFeeManager.contract.WatchLogs(opts, "UpdatedFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFeeManagerUpdatedFeeToken)
				if err := _IFeeManager.contract.UnpackLog(event, "UpdatedFeeToken", log); err != nil {
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

// ParseUpdatedFeeToken is a log parse operation binding the contract event 0x70698787bb7456af2d4339ba1e7d059c4eda7c473412e7ed2c0d5b361a49d50a.
//
// Solidity: event UpdatedFeeToken(address feeToken, uint256 feeAmount)
func (_IFeeManager *IFeeManagerFilterer) ParseUpdatedFeeToken(log types.Log) (*IFeeManagerUpdatedFeeToken, error) {
	event := new(IFeeManagerUpdatedFeeToken)
	if err := _IFeeManager.contract.UnpackLog(event, "UpdatedFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFeeManagerWithdrawnFeeTokenIterator is returned from FilterWithdrawnFeeToken and is used to iterate over the raw logs and unpacked data for WithdrawnFeeToken events raised by the IFeeManager contract.
type IFeeManagerWithdrawnFeeTokenIterator struct {
	Event *IFeeManagerWithdrawnFeeToken // Event containing the contract specifics and raw log

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
func (it *IFeeManagerWithdrawnFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFeeManagerWithdrawnFeeToken)
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
		it.Event = new(IFeeManagerWithdrawnFeeToken)
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
func (it *IFeeManagerWithdrawnFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFeeManagerWithdrawnFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFeeManagerWithdrawnFeeToken represents a WithdrawnFeeToken event raised by the IFeeManager contract.
type IFeeManagerWithdrawnFeeToken struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnFeeToken is a free log retrieval operation binding the contract event 0x0221f5dbeb176269bc9dbce8b10193c570930431bbe525ae79b4599910500bf4.
//
// Solidity: event WithdrawnFeeToken(address receiver, address feeToken, uint256 amount)
func (_IFeeManager *IFeeManagerFilterer) FilterWithdrawnFeeToken(opts *bind.FilterOpts) (*IFeeManagerWithdrawnFeeTokenIterator, error) {

	logs, sub, err := _IFeeManager.contract.FilterLogs(opts, "WithdrawnFeeToken")
	if err != nil {
		return nil, err
	}
	return &IFeeManagerWithdrawnFeeTokenIterator{contract: _IFeeManager.contract, event: "WithdrawnFeeToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawnFeeToken is a free log subscription operation binding the contract event 0x0221f5dbeb176269bc9dbce8b10193c570930431bbe525ae79b4599910500bf4.
//
// Solidity: event WithdrawnFeeToken(address receiver, address feeToken, uint256 amount)
func (_IFeeManager *IFeeManagerFilterer) WatchWithdrawnFeeToken(opts *bind.WatchOpts, sink chan<- *IFeeManagerWithdrawnFeeToken) (event.Subscription, error) {

	logs, sub, err := _IFeeManager.contract.WatchLogs(opts, "WithdrawnFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFeeManagerWithdrawnFeeToken)
				if err := _IFeeManager.contract.UnpackLog(event, "WithdrawnFeeToken", log); err != nil {
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

// ParseWithdrawnFeeToken is a log parse operation binding the contract event 0x0221f5dbeb176269bc9dbce8b10193c570930431bbe525ae79b4599910500bf4.
//
// Solidity: event WithdrawnFeeToken(address receiver, address feeToken, uint256 amount)
func (_IFeeManager *IFeeManagerFilterer) ParseWithdrawnFeeToken(log types.Log) (*IFeeManagerWithdrawnFeeToken, error) {
	event := new(IFeeManagerWithdrawnFeeToken)
	if err := _IFeeManager.contract.UnpackLog(event, "WithdrawnFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
