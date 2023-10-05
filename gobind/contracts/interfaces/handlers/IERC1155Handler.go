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

// IERC1155HandlerDepositERC1155Parameters is an auto generated low-level Go binding around an user-defined struct.
type IERC1155HandlerDepositERC1155Parameters struct {
	Token     common.Address
	TokenId   *big.Int
	Amount    *big.Int
	Bundle    IBundlerBundle
	Network   string
	Receiver  string
	IsWrapped bool
}

// IERC1155HandlerWithdrawERC1155Parameters is an auto generated low-level Go binding around an user-defined struct.
type IERC1155HandlerWithdrawERC1155Parameters struct {
	Token      common.Address
	TokenId    *big.Int
	TokenURI   string
	Amount     *big.Int
	Bundle     IBundlerBundle
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
	IsWrapped  bool
}

// IERC1155HandlerMetaData contains all meta data concerning the IERC1155Handler contract.
var IERC1155HandlerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"DepositedERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"WithdrawnERC1155\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.DepositERC1155Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt_\",\"type\":\"bytes32\"}],\"name\":\"determineProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.WithdrawERC1155Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.WithdrawERC1155Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC1155Bundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC1155HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC1155HandlerMetaData.ABI instead.
var IERC1155HandlerABI = IERC1155HandlerMetaData.ABI

// IERC1155Handler is an auto generated Go binding around an Ethereum contract.
type IERC1155Handler struct {
	IERC1155HandlerCaller     // Read-only binding to the contract
	IERC1155HandlerTransactor // Write-only binding to the contract
	IERC1155HandlerFilterer   // Log filterer for contract events
}

// IERC1155HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1155HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1155HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1155HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1155HandlerSession struct {
	Contract     *IERC1155Handler  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC1155HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1155HandlerCallerSession struct {
	Contract *IERC1155HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IERC1155HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1155HandlerTransactorSession struct {
	Contract     *IERC1155HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC1155HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1155HandlerRaw struct {
	Contract *IERC1155Handler // Generic contract binding to access the raw methods on
}

// IERC1155HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1155HandlerCallerRaw struct {
	Contract *IERC1155HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IERC1155HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1155HandlerTransactorRaw struct {
	Contract *IERC1155HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1155Handler creates a new instance of IERC1155Handler, bound to a specific deployed contract.
func NewIERC1155Handler(address common.Address, backend bind.ContractBackend) (*IERC1155Handler, error) {
	contract, err := bindIERC1155Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1155Handler{IERC1155HandlerCaller: IERC1155HandlerCaller{contract: contract}, IERC1155HandlerTransactor: IERC1155HandlerTransactor{contract: contract}, IERC1155HandlerFilterer: IERC1155HandlerFilterer{contract: contract}}, nil
}

// NewIERC1155HandlerCaller creates a new read-only instance of IERC1155Handler, bound to a specific deployed contract.
func NewIERC1155HandlerCaller(address common.Address, caller bind.ContractCaller) (*IERC1155HandlerCaller, error) {
	contract, err := bindIERC1155Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155HandlerCaller{contract: contract}, nil
}

// NewIERC1155HandlerTransactor creates a new write-only instance of IERC1155Handler, bound to a specific deployed contract.
func NewIERC1155HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC1155HandlerTransactor, error) {
	contract, err := bindIERC1155Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155HandlerTransactor{contract: contract}, nil
}

// NewIERC1155HandlerFilterer creates a new log filterer instance of IERC1155Handler, bound to a specific deployed contract.
func NewIERC1155HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC1155HandlerFilterer, error) {
	contract, err := bindIERC1155Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1155HandlerFilterer{contract: contract}, nil
}

// bindIERC1155Handler binds a generic wrapper to an already deployed contract.
func bindIERC1155Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC1155HandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155Handler *IERC1155HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155Handler.Contract.IERC1155HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155Handler *IERC1155HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155Handler.Contract.IERC1155HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155Handler *IERC1155HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155Handler.Contract.IERC1155HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155Handler *IERC1155HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155Handler *IERC1155HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155Handler *IERC1155HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155Handler.Contract.contract.Transact(opts, method, params...)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_IERC1155Handler *IERC1155HandlerCaller) DetermineProxyAddress(opts *bind.CallOpts, salt_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IERC1155Handler.contract.Call(opts, &out, "determineProxyAddress", salt_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_IERC1155Handler *IERC1155HandlerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _IERC1155Handler.Contract.DetermineProxyAddress(&_IERC1155Handler.CallOpts, salt_)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_IERC1155Handler *IERC1155HandlerCallerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _IERC1155Handler.Contract.DetermineProxyAddress(&_IERC1155Handler.CallOpts, salt_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x969aef06.
//
// Solidity: function depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IERC1155Handler *IERC1155HandlerTransactor) DepositERC1155(opts *bind.TransactOpts, params_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _IERC1155Handler.contract.Transact(opts, "depositERC1155", params_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x969aef06.
//
// Solidity: function depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IERC1155Handler *IERC1155HandlerSession) DepositERC1155(params_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _IERC1155Handler.Contract.DepositERC1155(&_IERC1155Handler.TransactOpts, params_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x969aef06.
//
// Solidity: function depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IERC1155Handler *IERC1155HandlerTransactorSession) DepositERC1155(params_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _IERC1155Handler.Contract.DepositERC1155(&_IERC1155Handler.TransactOpts, params_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC1155Handler *IERC1155HandlerTransactor) WithdrawERC1155(opts *bind.TransactOpts, params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _IERC1155Handler.contract.Transact(opts, "withdrawERC1155", params_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC1155Handler *IERC1155HandlerSession) WithdrawERC1155(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _IERC1155Handler.Contract.WithdrawERC1155(&_IERC1155Handler.TransactOpts, params_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC1155Handler *IERC1155HandlerTransactorSession) WithdrawERC1155(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _IERC1155Handler.Contract.WithdrawERC1155(&_IERC1155Handler.TransactOpts, params_)
}

// WithdrawERC1155Bundle is a paid mutator transaction binding the contract method 0xb76da974.
//
// Solidity: function withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC1155Handler *IERC1155HandlerTransactor) WithdrawERC1155Bundle(opts *bind.TransactOpts, params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _IERC1155Handler.contract.Transact(opts, "withdrawERC1155Bundle", params_)
}

// WithdrawERC1155Bundle is a paid mutator transaction binding the contract method 0xb76da974.
//
// Solidity: function withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC1155Handler *IERC1155HandlerSession) WithdrawERC1155Bundle(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _IERC1155Handler.Contract.WithdrawERC1155Bundle(&_IERC1155Handler.TransactOpts, params_)
}

// WithdrawERC1155Bundle is a paid mutator transaction binding the contract method 0xb76da974.
//
// Solidity: function withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IERC1155Handler *IERC1155HandlerTransactorSession) WithdrawERC1155Bundle(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _IERC1155Handler.Contract.WithdrawERC1155Bundle(&_IERC1155Handler.TransactOpts, params_)
}

// IERC1155HandlerDepositedERC1155Iterator is returned from FilterDepositedERC1155 and is used to iterate over the raw logs and unpacked data for DepositedERC1155 events raised by the IERC1155Handler contract.
type IERC1155HandlerDepositedERC1155Iterator struct {
	Event *IERC1155HandlerDepositedERC1155 // Event containing the contract specifics and raw log

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
func (it *IERC1155HandlerDepositedERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155HandlerDepositedERC1155)
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
		it.Event = new(IERC1155HandlerDepositedERC1155)
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
func (it *IERC1155HandlerDepositedERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155HandlerDepositedERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155HandlerDepositedERC1155 represents a DepositedERC1155 event raised by the IERC1155Handler contract.
type IERC1155HandlerDepositedERC1155 struct {
	Token     common.Address
	TokenId   *big.Int
	Amount    *big.Int
	Salt      [32]byte
	Bundle    []byte
	Network   string
	Receiver  string
	IsWrapped bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositedERC1155 is a free log retrieval operation binding the contract event 0x103b790f2fa3a8676ff87c3620a55f0853d0e45128a8c7e9fadf29e17c51d07a.
//
// Solidity: event DepositedERC1155(address token, uint256 tokenId, uint256 amount, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_IERC1155Handler *IERC1155HandlerFilterer) FilterDepositedERC1155(opts *bind.FilterOpts) (*IERC1155HandlerDepositedERC1155Iterator, error) {

	logs, sub, err := _IERC1155Handler.contract.FilterLogs(opts, "DepositedERC1155")
	if err != nil {
		return nil, err
	}
	return &IERC1155HandlerDepositedERC1155Iterator{contract: _IERC1155Handler.contract, event: "DepositedERC1155", logs: logs, sub: sub}, nil
}

// WatchDepositedERC1155 is a free log subscription operation binding the contract event 0x103b790f2fa3a8676ff87c3620a55f0853d0e45128a8c7e9fadf29e17c51d07a.
//
// Solidity: event DepositedERC1155(address token, uint256 tokenId, uint256 amount, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_IERC1155Handler *IERC1155HandlerFilterer) WatchDepositedERC1155(opts *bind.WatchOpts, sink chan<- *IERC1155HandlerDepositedERC1155) (event.Subscription, error) {

	logs, sub, err := _IERC1155Handler.contract.WatchLogs(opts, "DepositedERC1155")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155HandlerDepositedERC1155)
				if err := _IERC1155Handler.contract.UnpackLog(event, "DepositedERC1155", log); err != nil {
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

// ParseDepositedERC1155 is a log parse operation binding the contract event 0x103b790f2fa3a8676ff87c3620a55f0853d0e45128a8c7e9fadf29e17c51d07a.
//
// Solidity: event DepositedERC1155(address token, uint256 tokenId, uint256 amount, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_IERC1155Handler *IERC1155HandlerFilterer) ParseDepositedERC1155(log types.Log) (*IERC1155HandlerDepositedERC1155, error) {
	event := new(IERC1155HandlerDepositedERC1155)
	if err := _IERC1155Handler.contract.UnpackLog(event, "DepositedERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155HandlerWithdrawnERC1155Iterator is returned from FilterWithdrawnERC1155 and is used to iterate over the raw logs and unpacked data for WithdrawnERC1155 events raised by the IERC1155Handler contract.
type IERC1155HandlerWithdrawnERC1155Iterator struct {
	Event *IERC1155HandlerWithdrawnERC1155 // Event containing the contract specifics and raw log

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
func (it *IERC1155HandlerWithdrawnERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155HandlerWithdrawnERC1155)
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
		it.Event = new(IERC1155HandlerWithdrawnERC1155)
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
func (it *IERC1155HandlerWithdrawnERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155HandlerWithdrawnERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155HandlerWithdrawnERC1155 represents a WithdrawnERC1155 event raised by the IERC1155Handler contract.
type IERC1155HandlerWithdrawnERC1155 struct {
	Token      common.Address
	TokenId    *big.Int
	TokenURI   string
	Amount     *big.Int
	Salt       [32]byte
	Bundle     []byte
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
	IsWrapped  bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnERC1155 is a free log retrieval operation binding the contract event 0x30acf6d8c142717c265a7b8ab5be21a380a1ab05e2cd7ce3ca4fdd0c3260303e.
//
// Solidity: event WithdrawnERC1155(address token, uint256 tokenId, string tokenURI, uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_IERC1155Handler *IERC1155HandlerFilterer) FilterWithdrawnERC1155(opts *bind.FilterOpts) (*IERC1155HandlerWithdrawnERC1155Iterator, error) {

	logs, sub, err := _IERC1155Handler.contract.FilterLogs(opts, "WithdrawnERC1155")
	if err != nil {
		return nil, err
	}
	return &IERC1155HandlerWithdrawnERC1155Iterator{contract: _IERC1155Handler.contract, event: "WithdrawnERC1155", logs: logs, sub: sub}, nil
}

// WatchWithdrawnERC1155 is a free log subscription operation binding the contract event 0x30acf6d8c142717c265a7b8ab5be21a380a1ab05e2cd7ce3ca4fdd0c3260303e.
//
// Solidity: event WithdrawnERC1155(address token, uint256 tokenId, string tokenURI, uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_IERC1155Handler *IERC1155HandlerFilterer) WatchWithdrawnERC1155(opts *bind.WatchOpts, sink chan<- *IERC1155HandlerWithdrawnERC1155) (event.Subscription, error) {

	logs, sub, err := _IERC1155Handler.contract.WatchLogs(opts, "WithdrawnERC1155")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155HandlerWithdrawnERC1155)
				if err := _IERC1155Handler.contract.UnpackLog(event, "WithdrawnERC1155", log); err != nil {
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

// ParseWithdrawnERC1155 is a log parse operation binding the contract event 0x30acf6d8c142717c265a7b8ab5be21a380a1ab05e2cd7ce3ca4fdd0c3260303e.
//
// Solidity: event WithdrawnERC1155(address token, uint256 tokenId, string tokenURI, uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_IERC1155Handler *IERC1155HandlerFilterer) ParseWithdrawnERC1155(log types.Log) (*IERC1155HandlerWithdrawnERC1155, error) {
	event := new(IERC1155HandlerWithdrawnERC1155)
	if err := _IERC1155Handler.contract.UnpackLog(event, "WithdrawnERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
