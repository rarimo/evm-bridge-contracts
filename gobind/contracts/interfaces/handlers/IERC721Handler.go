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

// IERC721HandlerDepositERC721Parameters is an auto generated low-level Go binding around an user-defined struct.
type IERC721HandlerDepositERC721Parameters struct {
	Token     common.Address
	TokenId   *big.Int
	Bundle    IBundlerBundle
	Network   string
	Receiver  string
	IsWrapped bool
}

// IERC721HandlerWithdrawERC721Parameters is an auto generated low-level Go binding around an user-defined struct.
type IERC721HandlerWithdrawERC721Parameters struct {
	Token      common.Address
	TokenId    *big.Int
	TokenURI   string
	Bundle     IBundlerBundle
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
	IsWrapped  bool
}

// IERC721HandlerMetaData contains all meta data concerning the IERC721Handler contract.
var IERC721HandlerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"DepositedERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"WithdrawnERC721\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.DepositERC721Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt_\",\"type\":\"bytes32\"}],\"name\":\"determineProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.WithdrawERC721Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.WithdrawERC721Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC721Bundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC721HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721HandlerMetaData.ABI instead.
var IERC721HandlerABI = IERC721HandlerMetaData.ABI

// IERC721Handler is an auto generated Go binding around an Ethereum contract.
type IERC721Handler struct {
	IERC721HandlerCaller     // Read-only binding to the contract
	IERC721HandlerTransactor // Write-only binding to the contract
	IERC721HandlerFilterer   // Log filterer for contract events
}

// IERC721HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721HandlerSession struct {
	Contract     *IERC721Handler   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721HandlerCallerSession struct {
	Contract *IERC721HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IERC721HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721HandlerTransactorSession struct {
	Contract     *IERC721HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC721HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721HandlerRaw struct {
	Contract *IERC721Handler // Generic contract binding to access the raw methods on
}

// IERC721HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721HandlerCallerRaw struct {
	Contract *IERC721HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721HandlerTransactorRaw struct {
	Contract *IERC721HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Handler creates a new instance of IERC721Handler, bound to a specific deployed contract.
func NewIERC721Handler(address common.Address, backend bind.ContractBackend) (*IERC721Handler, error) {
	contract, err := bindIERC721Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Handler{IERC721HandlerCaller: IERC721HandlerCaller{contract: contract}, IERC721HandlerTransactor: IERC721HandlerTransactor{contract: contract}, IERC721HandlerFilterer: IERC721HandlerFilterer{contract: contract}}, nil
}

// NewIERC721HandlerCaller creates a new read-only instance of IERC721Handler, bound to a specific deployed contract.
func NewIERC721HandlerCaller(address common.Address, caller bind.ContractCaller) (*IERC721HandlerCaller, error) {
	contract, err := bindIERC721Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721HandlerCaller{contract: contract}, nil
}

// NewIERC721HandlerTransactor creates a new write-only instance of IERC721Handler, bound to a specific deployed contract.
func NewIERC721HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721HandlerTransactor, error) {
	contract, err := bindIERC721Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721HandlerTransactor{contract: contract}, nil
}

// NewIERC721HandlerFilterer creates a new log filterer instance of IERC721Handler, bound to a specific deployed contract.
func NewIERC721HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721HandlerFilterer, error) {
	contract, err := bindIERC721Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721HandlerFilterer{contract: contract}, nil
}

// bindIERC721Handler binds a generic wrapper to an already deployed contract.
func bindIERC721Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC721HandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Handler *IERC721HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Handler.Contract.IERC721HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Handler *IERC721HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Handler.Contract.IERC721HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Handler *IERC721HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Handler.Contract.IERC721HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Handler *IERC721HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Handler *IERC721HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Handler *IERC721HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Handler.Contract.contract.Transact(opts, method, params...)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_IERC721Handler *IERC721HandlerCaller) DetermineProxyAddress(opts *bind.CallOpts, salt_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IERC721Handler.contract.Call(opts, &out, "determineProxyAddress", salt_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_IERC721Handler *IERC721HandlerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _IERC721Handler.Contract.DetermineProxyAddress(&_IERC721Handler.CallOpts, salt_)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_IERC721Handler *IERC721HandlerCallerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _IERC721Handler.Contract.DetermineProxyAddress(&_IERC721Handler.CallOpts, salt_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x6a38abbf.
//
// Solidity: function depositERC721((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IERC721Handler *IERC721HandlerTransactor) DepositERC721(opts *bind.TransactOpts, params_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _IERC721Handler.contract.Transact(opts, "depositERC721", params_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x6a38abbf.
//
// Solidity: function depositERC721((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IERC721Handler *IERC721HandlerSession) DepositERC721(params_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _IERC721Handler.Contract.DepositERC721(&_IERC721Handler.TransactOpts, params_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x6a38abbf.
//
// Solidity: function depositERC721((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IERC721Handler *IERC721HandlerTransactorSession) DepositERC721(params_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _IERC721Handler.Contract.DepositERC721(&_IERC721Handler.TransactOpts, params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC721Handler *IERC721HandlerTransactor) WithdrawERC721(opts *bind.TransactOpts, params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _IERC721Handler.contract.Transact(opts, "withdrawERC721", params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC721Handler *IERC721HandlerSession) WithdrawERC721(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _IERC721Handler.Contract.WithdrawERC721(&_IERC721Handler.TransactOpts, params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC721Handler *IERC721HandlerTransactorSession) WithdrawERC721(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _IERC721Handler.Contract.WithdrawERC721(&_IERC721Handler.TransactOpts, params_)
}

// WithdrawERC721Bundle is a paid mutator transaction binding the contract method 0x9ae00a57.
//
// Solidity: function withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC721Handler *IERC721HandlerTransactor) WithdrawERC721Bundle(opts *bind.TransactOpts, params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _IERC721Handler.contract.Transact(opts, "withdrawERC721Bundle", params_)
}

// WithdrawERC721Bundle is a paid mutator transaction binding the contract method 0x9ae00a57.
//
// Solidity: function withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC721Handler *IERC721HandlerSession) WithdrawERC721Bundle(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _IERC721Handler.Contract.WithdrawERC721Bundle(&_IERC721Handler.TransactOpts, params_)
}

// WithdrawERC721Bundle is a paid mutator transaction binding the contract method 0x9ae00a57.
//
// Solidity: function withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC721Handler *IERC721HandlerTransactorSession) WithdrawERC721Bundle(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _IERC721Handler.Contract.WithdrawERC721Bundle(&_IERC721Handler.TransactOpts, params_)
}

// IERC721HandlerDepositedERC721Iterator is returned from FilterDepositedERC721 and is used to iterate over the raw logs and unpacked data for DepositedERC721 events raised by the IERC721Handler contract.
type IERC721HandlerDepositedERC721Iterator struct {
	Event *IERC721HandlerDepositedERC721 // Event containing the contract specifics and raw log

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
func (it *IERC721HandlerDepositedERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721HandlerDepositedERC721)
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
		it.Event = new(IERC721HandlerDepositedERC721)
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
func (it *IERC721HandlerDepositedERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721HandlerDepositedERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721HandlerDepositedERC721 represents a DepositedERC721 event raised by the IERC721Handler contract.
type IERC721HandlerDepositedERC721 struct {
	Token     common.Address
	TokenId   *big.Int
	Salt      [32]byte
	Bundle    []byte
	Network   string
	Receiver  string
	IsWrapped bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositedERC721 is a free log retrieval operation binding the contract event 0x7f787dd0c844dac4f8bfc4044046cdab3be531f7eefa9b740c531e48a99725e1.
//
// Solidity: event DepositedERC721(address token, uint256 tokenId, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_IERC721Handler *IERC721HandlerFilterer) FilterDepositedERC721(opts *bind.FilterOpts) (*IERC721HandlerDepositedERC721Iterator, error) {

	logs, sub, err := _IERC721Handler.contract.FilterLogs(opts, "DepositedERC721")
	if err != nil {
		return nil, err
	}
	return &IERC721HandlerDepositedERC721Iterator{contract: _IERC721Handler.contract, event: "DepositedERC721", logs: logs, sub: sub}, nil
}

// WatchDepositedERC721 is a free log subscription operation binding the contract event 0x7f787dd0c844dac4f8bfc4044046cdab3be531f7eefa9b740c531e48a99725e1.
//
// Solidity: event DepositedERC721(address token, uint256 tokenId, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_IERC721Handler *IERC721HandlerFilterer) WatchDepositedERC721(opts *bind.WatchOpts, sink chan<- *IERC721HandlerDepositedERC721) (event.Subscription, error) {

	logs, sub, err := _IERC721Handler.contract.WatchLogs(opts, "DepositedERC721")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721HandlerDepositedERC721)
				if err := _IERC721Handler.contract.UnpackLog(event, "DepositedERC721", log); err != nil {
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

// ParseDepositedERC721 is a log parse operation binding the contract event 0x7f787dd0c844dac4f8bfc4044046cdab3be531f7eefa9b740c531e48a99725e1.
//
// Solidity: event DepositedERC721(address token, uint256 tokenId, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_IERC721Handler *IERC721HandlerFilterer) ParseDepositedERC721(log types.Log) (*IERC721HandlerDepositedERC721, error) {
	event := new(IERC721HandlerDepositedERC721)
	if err := _IERC721Handler.contract.UnpackLog(event, "DepositedERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721HandlerWithdrawnERC721Iterator is returned from FilterWithdrawnERC721 and is used to iterate over the raw logs and unpacked data for WithdrawnERC721 events raised by the IERC721Handler contract.
type IERC721HandlerWithdrawnERC721Iterator struct {
	Event *IERC721HandlerWithdrawnERC721 // Event containing the contract specifics and raw log

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
func (it *IERC721HandlerWithdrawnERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721HandlerWithdrawnERC721)
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
		it.Event = new(IERC721HandlerWithdrawnERC721)
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
func (it *IERC721HandlerWithdrawnERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721HandlerWithdrawnERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721HandlerWithdrawnERC721 represents a WithdrawnERC721 event raised by the IERC721Handler contract.
type IERC721HandlerWithdrawnERC721 struct {
	Token      common.Address
	TokenId    *big.Int
	TokenURI   string
	Salt       [32]byte
	Bundle     []byte
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
	IsWrapped  bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnERC721 is a free log retrieval operation binding the contract event 0x8548956d276e213edf284f96ff17636f64d436743236a9866a7112888c442edd.
//
// Solidity: event WithdrawnERC721(address token, uint256 tokenId, string tokenURI, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_IERC721Handler *IERC721HandlerFilterer) FilterWithdrawnERC721(opts *bind.FilterOpts) (*IERC721HandlerWithdrawnERC721Iterator, error) {

	logs, sub, err := _IERC721Handler.contract.FilterLogs(opts, "WithdrawnERC721")
	if err != nil {
		return nil, err
	}
	return &IERC721HandlerWithdrawnERC721Iterator{contract: _IERC721Handler.contract, event: "WithdrawnERC721", logs: logs, sub: sub}, nil
}

// WatchWithdrawnERC721 is a free log subscription operation binding the contract event 0x8548956d276e213edf284f96ff17636f64d436743236a9866a7112888c442edd.
//
// Solidity: event WithdrawnERC721(address token, uint256 tokenId, string tokenURI, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_IERC721Handler *IERC721HandlerFilterer) WatchWithdrawnERC721(opts *bind.WatchOpts, sink chan<- *IERC721HandlerWithdrawnERC721) (event.Subscription, error) {

	logs, sub, err := _IERC721Handler.contract.WatchLogs(opts, "WithdrawnERC721")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721HandlerWithdrawnERC721)
				if err := _IERC721Handler.contract.UnpackLog(event, "WithdrawnERC721", log); err != nil {
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

// ParseWithdrawnERC721 is a log parse operation binding the contract event 0x8548956d276e213edf284f96ff17636f64d436743236a9866a7112888c442edd.
//
// Solidity: event WithdrawnERC721(address token, uint256 tokenId, string tokenURI, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_IERC721Handler *IERC721HandlerFilterer) ParseWithdrawnERC721(log types.Log) (*IERC721HandlerWithdrawnERC721, error) {
	event := new(IERC721HandlerWithdrawnERC721)
	if err := _IERC721Handler.contract.UnpackLog(event, "WithdrawnERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
