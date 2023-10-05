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

// ERC20HandlerMetaData contains all meta data concerning the ERC20Handler contract.
var ERC20HandlerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"DepositedERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"WithdrawnERC20\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bundleExecutorImplementation_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"facade_\",\"type\":\"address\"}],\"name\":\"__Bundler_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bundleExecutorImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.DepositERC20Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt_\",\"type\":\"bytes32\"}],\"name\":\"determineProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.WithdrawERC20Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.WithdrawERC20Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC20Bundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC20HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20HandlerMetaData.ABI instead.
var ERC20HandlerABI = ERC20HandlerMetaData.ABI

// ERC20Handler is an auto generated Go binding around an Ethereum contract.
type ERC20Handler struct {
	ERC20HandlerCaller     // Read-only binding to the contract
	ERC20HandlerTransactor // Write-only binding to the contract
	ERC20HandlerFilterer   // Log filterer for contract events
}

// ERC20HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20HandlerSession struct {
	Contract     *ERC20Handler     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20HandlerCallerSession struct {
	Contract *ERC20HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC20HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20HandlerTransactorSession struct {
	Contract     *ERC20HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC20HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20HandlerRaw struct {
	Contract *ERC20Handler // Generic contract binding to access the raw methods on
}

// ERC20HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20HandlerCallerRaw struct {
	Contract *ERC20HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactorRaw struct {
	Contract *ERC20HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Handler creates a new instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20Handler(address common.Address, backend bind.ContractBackend) (*ERC20Handler, error) {
	contract, err := bindERC20Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}, ERC20HandlerFilterer: ERC20HandlerFilterer{contract: contract}}, nil
}

// NewERC20HandlerCaller creates a new read-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC20HandlerCaller, error) {
	contract, err := bindERC20Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerCaller{contract: contract}, nil
}

// NewERC20HandlerTransactor creates a new write-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20HandlerTransactor, error) {
	contract, err := bindERC20Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerTransactor{contract: contract}, nil
}

// NewERC20HandlerFilterer creates a new log filterer instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20HandlerFilterer, error) {
	contract, err := bindERC20Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerFilterer{contract: contract}, nil
}

// bindERC20Handler binds a generic wrapper to an already deployed contract.
func bindERC20Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20HandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Handler *ERC20HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Handler.Contract.ERC20HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Handler *ERC20HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ERC20HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Handler *ERC20HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ERC20HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Handler *ERC20HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Handler *ERC20HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Handler *ERC20HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Handler.Contract.contract.Transact(opts, method, params...)
}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_ERC20Handler *ERC20HandlerCaller) BundleExecutorImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "bundleExecutorImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_ERC20Handler *ERC20HandlerSession) BundleExecutorImplementation() (common.Address, error) {
	return _ERC20Handler.Contract.BundleExecutorImplementation(&_ERC20Handler.CallOpts)
}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) BundleExecutorImplementation() (common.Address, error) {
	return _ERC20Handler.Contract.BundleExecutorImplementation(&_ERC20Handler.CallOpts)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_ERC20Handler *ERC20HandlerCaller) DetermineProxyAddress(opts *bind.CallOpts, salt_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "determineProxyAddress", salt_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_ERC20Handler *ERC20HandlerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.DetermineProxyAddress(&_ERC20Handler.CallOpts, salt_)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.DetermineProxyAddress(&_ERC20Handler.CallOpts, salt_)
}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_ERC20Handler *ERC20HandlerCaller) Facade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "facade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_ERC20Handler *ERC20HandlerSession) Facade() (common.Address, error) {
	return _ERC20Handler.Contract.Facade(&_ERC20Handler.CallOpts)
}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) Facade() (common.Address, error) {
	return _ERC20Handler.Contract.Facade(&_ERC20Handler.CallOpts)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x96de44c2.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_, address facade_) returns()
func (_ERC20Handler *ERC20HandlerTransactor) BundlerInit(opts *bind.TransactOpts, bundleExecutorImplementation_ common.Address, facade_ common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "__Bundler_init", bundleExecutorImplementation_, facade_)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x96de44c2.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_, address facade_) returns()
func (_ERC20Handler *ERC20HandlerSession) BundlerInit(bundleExecutorImplementation_ common.Address, facade_ common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.BundlerInit(&_ERC20Handler.TransactOpts, bundleExecutorImplementation_, facade_)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x96de44c2.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_, address facade_) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) BundlerInit(bundleExecutorImplementation_ common.Address, facade_ common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.BundlerInit(&_ERC20Handler.TransactOpts, bundleExecutorImplementation_, facade_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xbebd5a0b.
//
// Solidity: function depositERC20((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_ERC20Handler *ERC20HandlerTransactor) DepositERC20(opts *bind.TransactOpts, params_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "depositERC20", params_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xbebd5a0b.
//
// Solidity: function depositERC20((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_ERC20Handler *ERC20HandlerSession) DepositERC20(params_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _ERC20Handler.Contract.DepositERC20(&_ERC20Handler.TransactOpts, params_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xbebd5a0b.
//
// Solidity: function depositERC20((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) DepositERC20(params_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _ERC20Handler.Contract.DepositERC20(&_ERC20Handler.TransactOpts, params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC20Handler *ERC20HandlerTransactor) WithdrawERC20(opts *bind.TransactOpts, params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "withdrawERC20", params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC20Handler *ERC20HandlerSession) WithdrawERC20(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _ERC20Handler.Contract.WithdrawERC20(&_ERC20Handler.TransactOpts, params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) WithdrawERC20(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _ERC20Handler.Contract.WithdrawERC20(&_ERC20Handler.TransactOpts, params_)
}

// WithdrawERC20Bundle is a paid mutator transaction binding the contract method 0xe53507bd.
//
// Solidity: function withdrawERC20Bundle((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC20Handler *ERC20HandlerTransactor) WithdrawERC20Bundle(opts *bind.TransactOpts, params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "withdrawERC20Bundle", params_)
}

// WithdrawERC20Bundle is a paid mutator transaction binding the contract method 0xe53507bd.
//
// Solidity: function withdrawERC20Bundle((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC20Handler *ERC20HandlerSession) WithdrawERC20Bundle(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _ERC20Handler.Contract.WithdrawERC20Bundle(&_ERC20Handler.TransactOpts, params_)
}

// WithdrawERC20Bundle is a paid mutator transaction binding the contract method 0xe53507bd.
//
// Solidity: function withdrawERC20Bundle((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) WithdrawERC20Bundle(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _ERC20Handler.Contract.WithdrawERC20Bundle(&_ERC20Handler.TransactOpts, params_)
}

// ERC20HandlerDepositedERC20Iterator is returned from FilterDepositedERC20 and is used to iterate over the raw logs and unpacked data for DepositedERC20 events raised by the ERC20Handler contract.
type ERC20HandlerDepositedERC20Iterator struct {
	Event *ERC20HandlerDepositedERC20 // Event containing the contract specifics and raw log

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
func (it *ERC20HandlerDepositedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20HandlerDepositedERC20)
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
		it.Event = new(ERC20HandlerDepositedERC20)
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
func (it *ERC20HandlerDepositedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20HandlerDepositedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20HandlerDepositedERC20 represents a DepositedERC20 event raised by the ERC20Handler contract.
type ERC20HandlerDepositedERC20 struct {
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
func (_ERC20Handler *ERC20HandlerFilterer) FilterDepositedERC20(opts *bind.FilterOpts) (*ERC20HandlerDepositedERC20Iterator, error) {

	logs, sub, err := _ERC20Handler.contract.FilterLogs(opts, "DepositedERC20")
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerDepositedERC20Iterator{contract: _ERC20Handler.contract, event: "DepositedERC20", logs: logs, sub: sub}, nil
}

// WatchDepositedERC20 is a free log subscription operation binding the contract event 0x043d52f9acdd847f0210803c386559db9e09d492143f2072fe30ea62ff0bb639.
//
// Solidity: event DepositedERC20(address token, uint256 amount, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_ERC20Handler *ERC20HandlerFilterer) WatchDepositedERC20(opts *bind.WatchOpts, sink chan<- *ERC20HandlerDepositedERC20) (event.Subscription, error) {

	logs, sub, err := _ERC20Handler.contract.WatchLogs(opts, "DepositedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20HandlerDepositedERC20)
				if err := _ERC20Handler.contract.UnpackLog(event, "DepositedERC20", log); err != nil {
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
func (_ERC20Handler *ERC20HandlerFilterer) ParseDepositedERC20(log types.Log) (*ERC20HandlerDepositedERC20, error) {
	event := new(ERC20HandlerDepositedERC20)
	if err := _ERC20Handler.contract.UnpackLog(event, "DepositedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20HandlerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC20Handler contract.
type ERC20HandlerInitializedIterator struct {
	Event *ERC20HandlerInitialized // Event containing the contract specifics and raw log

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
func (it *ERC20HandlerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20HandlerInitialized)
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
		it.Event = new(ERC20HandlerInitialized)
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
func (it *ERC20HandlerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20HandlerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20HandlerInitialized represents a Initialized event raised by the ERC20Handler contract.
type ERC20HandlerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC20Handler *ERC20HandlerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC20HandlerInitializedIterator, error) {

	logs, sub, err := _ERC20Handler.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerInitializedIterator{contract: _ERC20Handler.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC20Handler *ERC20HandlerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC20HandlerInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC20Handler.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20HandlerInitialized)
				if err := _ERC20Handler.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC20Handler *ERC20HandlerFilterer) ParseInitialized(log types.Log) (*ERC20HandlerInitialized, error) {
	event := new(ERC20HandlerInitialized)
	if err := _ERC20Handler.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20HandlerWithdrawnERC20Iterator is returned from FilterWithdrawnERC20 and is used to iterate over the raw logs and unpacked data for WithdrawnERC20 events raised by the ERC20Handler contract.
type ERC20HandlerWithdrawnERC20Iterator struct {
	Event *ERC20HandlerWithdrawnERC20 // Event containing the contract specifics and raw log

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
func (it *ERC20HandlerWithdrawnERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20HandlerWithdrawnERC20)
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
		it.Event = new(ERC20HandlerWithdrawnERC20)
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
func (it *ERC20HandlerWithdrawnERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20HandlerWithdrawnERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20HandlerWithdrawnERC20 represents a WithdrawnERC20 event raised by the ERC20Handler contract.
type ERC20HandlerWithdrawnERC20 struct {
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
func (_ERC20Handler *ERC20HandlerFilterer) FilterWithdrawnERC20(opts *bind.FilterOpts) (*ERC20HandlerWithdrawnERC20Iterator, error) {

	logs, sub, err := _ERC20Handler.contract.FilterLogs(opts, "WithdrawnERC20")
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerWithdrawnERC20Iterator{contract: _ERC20Handler.contract, event: "WithdrawnERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawnERC20 is a free log subscription operation binding the contract event 0xff14005e9cd01bf3385cde89fb85a18453a0e3df47a3224d84ebab3fbd174128.
//
// Solidity: event WithdrawnERC20(address token, uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_ERC20Handler *ERC20HandlerFilterer) WatchWithdrawnERC20(opts *bind.WatchOpts, sink chan<- *ERC20HandlerWithdrawnERC20) (event.Subscription, error) {

	logs, sub, err := _ERC20Handler.contract.WatchLogs(opts, "WithdrawnERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20HandlerWithdrawnERC20)
				if err := _ERC20Handler.contract.UnpackLog(event, "WithdrawnERC20", log); err != nil {
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
func (_ERC20Handler *ERC20HandlerFilterer) ParseWithdrawnERC20(log types.Log) (*ERC20HandlerWithdrawnERC20, error) {
	event := new(ERC20HandlerWithdrawnERC20)
	if err := _ERC20Handler.contract.UnpackLog(event, "WithdrawnERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
