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

// ISBTHandlerDepositSBTParameters is an auto generated low-level Go binding around an user-defined struct.
type ISBTHandlerDepositSBTParameters struct {
	Token    common.Address
	TokenId  *big.Int
	Bundle   IBundlerBundle
	Network  string
	Receiver string
}

// ISBTHandlerWithdrawSBTParameters is an auto generated low-level Go binding around an user-defined struct.
type ISBTHandlerWithdrawSBTParameters struct {
	Token      common.Address
	TokenId    *big.Int
	TokenURI   string
	Bundle     IBundlerBundle
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
}

// ISBTHandlerMetaData contains all meta data concerning the ISBTHandler contract.
var ISBTHandlerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"name\":\"DepositedSBT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"WithdrawnSBT\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"internalType\":\"structISBTHandler.DepositSBTParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositSBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt_\",\"type\":\"bytes32\"}],\"name\":\"determineProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structISBTHandler.WithdrawSBTParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawSBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structISBTHandler.WithdrawSBTParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawSBTBundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ISBTHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use ISBTHandlerMetaData.ABI instead.
var ISBTHandlerABI = ISBTHandlerMetaData.ABI

// ISBTHandler is an auto generated Go binding around an Ethereum contract.
type ISBTHandler struct {
	ISBTHandlerCaller     // Read-only binding to the contract
	ISBTHandlerTransactor // Write-only binding to the contract
	ISBTHandlerFilterer   // Log filterer for contract events
}

// ISBTHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISBTHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISBTHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISBTHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISBTHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISBTHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISBTHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISBTHandlerSession struct {
	Contract     *ISBTHandler      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISBTHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISBTHandlerCallerSession struct {
	Contract *ISBTHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ISBTHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISBTHandlerTransactorSession struct {
	Contract     *ISBTHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ISBTHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISBTHandlerRaw struct {
	Contract *ISBTHandler // Generic contract binding to access the raw methods on
}

// ISBTHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISBTHandlerCallerRaw struct {
	Contract *ISBTHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ISBTHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISBTHandlerTransactorRaw struct {
	Contract *ISBTHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISBTHandler creates a new instance of ISBTHandler, bound to a specific deployed contract.
func NewISBTHandler(address common.Address, backend bind.ContractBackend) (*ISBTHandler, error) {
	contract, err := bindISBTHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISBTHandler{ISBTHandlerCaller: ISBTHandlerCaller{contract: contract}, ISBTHandlerTransactor: ISBTHandlerTransactor{contract: contract}, ISBTHandlerFilterer: ISBTHandlerFilterer{contract: contract}}, nil
}

// NewISBTHandlerCaller creates a new read-only instance of ISBTHandler, bound to a specific deployed contract.
func NewISBTHandlerCaller(address common.Address, caller bind.ContractCaller) (*ISBTHandlerCaller, error) {
	contract, err := bindISBTHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISBTHandlerCaller{contract: contract}, nil
}

// NewISBTHandlerTransactor creates a new write-only instance of ISBTHandler, bound to a specific deployed contract.
func NewISBTHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ISBTHandlerTransactor, error) {
	contract, err := bindISBTHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISBTHandlerTransactor{contract: contract}, nil
}

// NewISBTHandlerFilterer creates a new log filterer instance of ISBTHandler, bound to a specific deployed contract.
func NewISBTHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ISBTHandlerFilterer, error) {
	contract, err := bindISBTHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISBTHandlerFilterer{contract: contract}, nil
}

// bindISBTHandler binds a generic wrapper to an already deployed contract.
func bindISBTHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISBTHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISBTHandler *ISBTHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISBTHandler.Contract.ISBTHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISBTHandler *ISBTHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISBTHandler.Contract.ISBTHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISBTHandler *ISBTHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISBTHandler.Contract.ISBTHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISBTHandler *ISBTHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISBTHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISBTHandler *ISBTHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISBTHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISBTHandler *ISBTHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISBTHandler.Contract.contract.Transact(opts, method, params...)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_ISBTHandler *ISBTHandlerCaller) DetermineProxyAddress(opts *bind.CallOpts, salt_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ISBTHandler.contract.Call(opts, &out, "determineProxyAddress", salt_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_ISBTHandler *ISBTHandlerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _ISBTHandler.Contract.DetermineProxyAddress(&_ISBTHandler.CallOpts, salt_)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_ISBTHandler *ISBTHandlerCallerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _ISBTHandler.Contract.DetermineProxyAddress(&_ISBTHandler.CallOpts, salt_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0x755f3823.
//
// Solidity: function depositSBT((address,uint256,(bytes32,bytes),string,string) params_) returns()
func (_ISBTHandler *ISBTHandlerTransactor) DepositSBT(opts *bind.TransactOpts, params_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _ISBTHandler.contract.Transact(opts, "depositSBT", params_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0x755f3823.
//
// Solidity: function depositSBT((address,uint256,(bytes32,bytes),string,string) params_) returns()
func (_ISBTHandler *ISBTHandlerSession) DepositSBT(params_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _ISBTHandler.Contract.DepositSBT(&_ISBTHandler.TransactOpts, params_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0x755f3823.
//
// Solidity: function depositSBT((address,uint256,(bytes32,bytes),string,string) params_) returns()
func (_ISBTHandler *ISBTHandlerTransactorSession) DepositSBT(params_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _ISBTHandler.Contract.DepositSBT(&_ISBTHandler.TransactOpts, params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_ISBTHandler *ISBTHandlerTransactor) WithdrawSBT(opts *bind.TransactOpts, params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _ISBTHandler.contract.Transact(opts, "withdrawSBT", params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_ISBTHandler *ISBTHandlerSession) WithdrawSBT(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _ISBTHandler.Contract.WithdrawSBT(&_ISBTHandler.TransactOpts, params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_ISBTHandler *ISBTHandlerTransactorSession) WithdrawSBT(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _ISBTHandler.Contract.WithdrawSBT(&_ISBTHandler.TransactOpts, params_)
}

// WithdrawSBTBundle is a paid mutator transaction binding the contract method 0x9a830d4c.
//
// Solidity: function withdrawSBTBundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_ISBTHandler *ISBTHandlerTransactor) WithdrawSBTBundle(opts *bind.TransactOpts, params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _ISBTHandler.contract.Transact(opts, "withdrawSBTBundle", params_)
}

// WithdrawSBTBundle is a paid mutator transaction binding the contract method 0x9a830d4c.
//
// Solidity: function withdrawSBTBundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_ISBTHandler *ISBTHandlerSession) WithdrawSBTBundle(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _ISBTHandler.Contract.WithdrawSBTBundle(&_ISBTHandler.TransactOpts, params_)
}

// WithdrawSBTBundle is a paid mutator transaction binding the contract method 0x9a830d4c.
//
// Solidity: function withdrawSBTBundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_ISBTHandler *ISBTHandlerTransactorSession) WithdrawSBTBundle(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _ISBTHandler.Contract.WithdrawSBTBundle(&_ISBTHandler.TransactOpts, params_)
}

// ISBTHandlerDepositedSBTIterator is returned from FilterDepositedSBT and is used to iterate over the raw logs and unpacked data for DepositedSBT events raised by the ISBTHandler contract.
type ISBTHandlerDepositedSBTIterator struct {
	Event *ISBTHandlerDepositedSBT // Event containing the contract specifics and raw log

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
func (it *ISBTHandlerDepositedSBTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISBTHandlerDepositedSBT)
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
		it.Event = new(ISBTHandlerDepositedSBT)
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
func (it *ISBTHandlerDepositedSBTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISBTHandlerDepositedSBTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISBTHandlerDepositedSBT represents a DepositedSBT event raised by the ISBTHandler contract.
type ISBTHandlerDepositedSBT struct {
	Token    common.Address
	TokenId  *big.Int
	Salt     [32]byte
	Bundle   []byte
	Network  string
	Receiver string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDepositedSBT is a free log retrieval operation binding the contract event 0x1452835f6476968a263680d4519d5a7d74a60354947a95bad57f7274c339ae86.
//
// Solidity: event DepositedSBT(address token, uint256 tokenId, bytes32 salt, bytes bundle, string network, string receiver)
func (_ISBTHandler *ISBTHandlerFilterer) FilterDepositedSBT(opts *bind.FilterOpts) (*ISBTHandlerDepositedSBTIterator, error) {

	logs, sub, err := _ISBTHandler.contract.FilterLogs(opts, "DepositedSBT")
	if err != nil {
		return nil, err
	}
	return &ISBTHandlerDepositedSBTIterator{contract: _ISBTHandler.contract, event: "DepositedSBT", logs: logs, sub: sub}, nil
}

// WatchDepositedSBT is a free log subscription operation binding the contract event 0x1452835f6476968a263680d4519d5a7d74a60354947a95bad57f7274c339ae86.
//
// Solidity: event DepositedSBT(address token, uint256 tokenId, bytes32 salt, bytes bundle, string network, string receiver)
func (_ISBTHandler *ISBTHandlerFilterer) WatchDepositedSBT(opts *bind.WatchOpts, sink chan<- *ISBTHandlerDepositedSBT) (event.Subscription, error) {

	logs, sub, err := _ISBTHandler.contract.WatchLogs(opts, "DepositedSBT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISBTHandlerDepositedSBT)
				if err := _ISBTHandler.contract.UnpackLog(event, "DepositedSBT", log); err != nil {
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

// ParseDepositedSBT is a log parse operation binding the contract event 0x1452835f6476968a263680d4519d5a7d74a60354947a95bad57f7274c339ae86.
//
// Solidity: event DepositedSBT(address token, uint256 tokenId, bytes32 salt, bytes bundle, string network, string receiver)
func (_ISBTHandler *ISBTHandlerFilterer) ParseDepositedSBT(log types.Log) (*ISBTHandlerDepositedSBT, error) {
	event := new(ISBTHandlerDepositedSBT)
	if err := _ISBTHandler.contract.UnpackLog(event, "DepositedSBT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISBTHandlerWithdrawnSBTIterator is returned from FilterWithdrawnSBT and is used to iterate over the raw logs and unpacked data for WithdrawnSBT events raised by the ISBTHandler contract.
type ISBTHandlerWithdrawnSBTIterator struct {
	Event *ISBTHandlerWithdrawnSBT // Event containing the contract specifics and raw log

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
func (it *ISBTHandlerWithdrawnSBTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISBTHandlerWithdrawnSBT)
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
		it.Event = new(ISBTHandlerWithdrawnSBT)
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
func (it *ISBTHandlerWithdrawnSBTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISBTHandlerWithdrawnSBTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISBTHandlerWithdrawnSBT represents a WithdrawnSBT event raised by the ISBTHandler contract.
type ISBTHandlerWithdrawnSBT struct {
	Token      common.Address
	TokenId    *big.Int
	TokenURI   string
	Salt       [32]byte
	Bundle     []byte
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnSBT is a free log retrieval operation binding the contract event 0xbdb8f203981c7db82a7d157c20285399875b02508b52db10a634dd2a72c1b623.
//
// Solidity: event WithdrawnSBT(address token, uint256 tokenId, string tokenURI, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_ISBTHandler *ISBTHandlerFilterer) FilterWithdrawnSBT(opts *bind.FilterOpts) (*ISBTHandlerWithdrawnSBTIterator, error) {

	logs, sub, err := _ISBTHandler.contract.FilterLogs(opts, "WithdrawnSBT")
	if err != nil {
		return nil, err
	}
	return &ISBTHandlerWithdrawnSBTIterator{contract: _ISBTHandler.contract, event: "WithdrawnSBT", logs: logs, sub: sub}, nil
}

// WatchWithdrawnSBT is a free log subscription operation binding the contract event 0xbdb8f203981c7db82a7d157c20285399875b02508b52db10a634dd2a72c1b623.
//
// Solidity: event WithdrawnSBT(address token, uint256 tokenId, string tokenURI, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_ISBTHandler *ISBTHandlerFilterer) WatchWithdrawnSBT(opts *bind.WatchOpts, sink chan<- *ISBTHandlerWithdrawnSBT) (event.Subscription, error) {

	logs, sub, err := _ISBTHandler.contract.WatchLogs(opts, "WithdrawnSBT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISBTHandlerWithdrawnSBT)
				if err := _ISBTHandler.contract.UnpackLog(event, "WithdrawnSBT", log); err != nil {
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

// ParseWithdrawnSBT is a log parse operation binding the contract event 0xbdb8f203981c7db82a7d157c20285399875b02508b52db10a634dd2a72c1b623.
//
// Solidity: event WithdrawnSBT(address token, uint256 tokenId, string tokenURI, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_ISBTHandler *ISBTHandlerFilterer) ParseWithdrawnSBT(log types.Log) (*ISBTHandlerWithdrawnSBT, error) {
	event := new(ISBTHandlerWithdrawnSBT)
	if err := _ISBTHandler.contract.UnpackLog(event, "WithdrawnSBT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
