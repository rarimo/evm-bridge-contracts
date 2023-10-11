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

// ERC721HandlerMetaData contains all meta data concerning the ERC721Handler contract.
var ERC721HandlerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"DepositedERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"WithdrawnERC721\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bundleExecutorImplementation_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"facade_\",\"type\":\"address\"}],\"name\":\"__Bundler_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bundleExecutorImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.DepositERC721Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt_\",\"type\":\"bytes32\"}],\"name\":\"determineProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.WithdrawERC721Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.WithdrawERC721Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC721Bundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC721HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721HandlerMetaData.ABI instead.
var ERC721HandlerABI = ERC721HandlerMetaData.ABI

// ERC721Handler is an auto generated Go binding around an Ethereum contract.
type ERC721Handler struct {
	ERC721HandlerCaller     // Read-only binding to the contract
	ERC721HandlerTransactor // Write-only binding to the contract
	ERC721HandlerFilterer   // Log filterer for contract events
}

// ERC721HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721HandlerSession struct {
	Contract     *ERC721Handler    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721HandlerCallerSession struct {
	Contract *ERC721HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC721HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721HandlerTransactorSession struct {
	Contract     *ERC721HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC721HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721HandlerRaw struct {
	Contract *ERC721Handler // Generic contract binding to access the raw methods on
}

// ERC721HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721HandlerCallerRaw struct {
	Contract *ERC721HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721HandlerTransactorRaw struct {
	Contract *ERC721HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Handler creates a new instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721Handler(address common.Address, backend bind.ContractBackend) (*ERC721Handler, error) {
	contract, err := bindERC721Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Handler{ERC721HandlerCaller: ERC721HandlerCaller{contract: contract}, ERC721HandlerTransactor: ERC721HandlerTransactor{contract: contract}, ERC721HandlerFilterer: ERC721HandlerFilterer{contract: contract}}, nil
}

// NewERC721HandlerCaller creates a new read-only instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC721HandlerCaller, error) {
	contract, err := bindERC721Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerCaller{contract: contract}, nil
}

// NewERC721HandlerTransactor creates a new write-only instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721HandlerTransactor, error) {
	contract, err := bindERC721Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerTransactor{contract: contract}, nil
}

// NewERC721HandlerFilterer creates a new log filterer instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721HandlerFilterer, error) {
	contract, err := bindERC721Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerFilterer{contract: contract}, nil
}

// bindERC721Handler binds a generic wrapper to an already deployed contract.
func bindERC721Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC721HandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Handler *ERC721HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Handler.Contract.ERC721HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Handler *ERC721HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ERC721HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Handler *ERC721HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ERC721HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Handler *ERC721HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Handler *ERC721HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Handler *ERC721HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Handler.Contract.contract.Transact(opts, method, params...)
}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_ERC721Handler *ERC721HandlerCaller) BundleExecutorImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "bundleExecutorImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_ERC721Handler *ERC721HandlerSession) BundleExecutorImplementation() (common.Address, error) {
	return _ERC721Handler.Contract.BundleExecutorImplementation(&_ERC721Handler.CallOpts)
}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_ERC721Handler *ERC721HandlerCallerSession) BundleExecutorImplementation() (common.Address, error) {
	return _ERC721Handler.Contract.BundleExecutorImplementation(&_ERC721Handler.CallOpts)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_ERC721Handler *ERC721HandlerCaller) DetermineProxyAddress(opts *bind.CallOpts, salt_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "determineProxyAddress", salt_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_ERC721Handler *ERC721HandlerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _ERC721Handler.Contract.DetermineProxyAddress(&_ERC721Handler.CallOpts, salt_)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_ERC721Handler *ERC721HandlerCallerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _ERC721Handler.Contract.DetermineProxyAddress(&_ERC721Handler.CallOpts, salt_)
}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_ERC721Handler *ERC721HandlerCaller) Facade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "facade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_ERC721Handler *ERC721HandlerSession) Facade() (common.Address, error) {
	return _ERC721Handler.Contract.Facade(&_ERC721Handler.CallOpts)
}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_ERC721Handler *ERC721HandlerCallerSession) Facade() (common.Address, error) {
	return _ERC721Handler.Contract.Facade(&_ERC721Handler.CallOpts)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x96de44c2.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_, address facade_) returns()
func (_ERC721Handler *ERC721HandlerTransactor) BundlerInit(opts *bind.TransactOpts, bundleExecutorImplementation_ common.Address, facade_ common.Address) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "__Bundler_init", bundleExecutorImplementation_, facade_)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x96de44c2.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_, address facade_) returns()
func (_ERC721Handler *ERC721HandlerSession) BundlerInit(bundleExecutorImplementation_ common.Address, facade_ common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.BundlerInit(&_ERC721Handler.TransactOpts, bundleExecutorImplementation_, facade_)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x96de44c2.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_, address facade_) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) BundlerInit(bundleExecutorImplementation_ common.Address, facade_ common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.BundlerInit(&_ERC721Handler.TransactOpts, bundleExecutorImplementation_, facade_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x6a38abbf.
//
// Solidity: function depositERC721((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_ERC721Handler *ERC721HandlerTransactor) DepositERC721(opts *bind.TransactOpts, params_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "depositERC721", params_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x6a38abbf.
//
// Solidity: function depositERC721((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_ERC721Handler *ERC721HandlerSession) DepositERC721(params_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _ERC721Handler.Contract.DepositERC721(&_ERC721Handler.TransactOpts, params_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x6a38abbf.
//
// Solidity: function depositERC721((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) DepositERC721(params_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _ERC721Handler.Contract.DepositERC721(&_ERC721Handler.TransactOpts, params_)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_ERC721Handler *ERC721HandlerTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_ERC721Handler *ERC721HandlerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.OnERC721Received(&_ERC721Handler.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_ERC721Handler *ERC721HandlerTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.OnERC721Received(&_ERC721Handler.TransactOpts, arg0, arg1, arg2, arg3)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC721Handler *ERC721HandlerTransactor) WithdrawERC721(opts *bind.TransactOpts, params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "withdrawERC721", params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC721Handler *ERC721HandlerSession) WithdrawERC721(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _ERC721Handler.Contract.WithdrawERC721(&_ERC721Handler.TransactOpts, params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) WithdrawERC721(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _ERC721Handler.Contract.WithdrawERC721(&_ERC721Handler.TransactOpts, params_)
}

// WithdrawERC721Bundle is a paid mutator transaction binding the contract method 0x9ae00a57.
//
// Solidity: function withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC721Handler *ERC721HandlerTransactor) WithdrawERC721Bundle(opts *bind.TransactOpts, params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "withdrawERC721Bundle", params_)
}

// WithdrawERC721Bundle is a paid mutator transaction binding the contract method 0x9ae00a57.
//
// Solidity: function withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC721Handler *ERC721HandlerSession) WithdrawERC721Bundle(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _ERC721Handler.Contract.WithdrawERC721Bundle(&_ERC721Handler.TransactOpts, params_)
}

// WithdrawERC721Bundle is a paid mutator transaction binding the contract method 0x9ae00a57.
//
// Solidity: function withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) WithdrawERC721Bundle(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _ERC721Handler.Contract.WithdrawERC721Bundle(&_ERC721Handler.TransactOpts, params_)
}

// ERC721HandlerDepositedERC721Iterator is returned from FilterDepositedERC721 and is used to iterate over the raw logs and unpacked data for DepositedERC721 events raised by the ERC721Handler contract.
type ERC721HandlerDepositedERC721Iterator struct {
	Event *ERC721HandlerDepositedERC721 // Event containing the contract specifics and raw log

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
func (it *ERC721HandlerDepositedERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721HandlerDepositedERC721)
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
		it.Event = new(ERC721HandlerDepositedERC721)
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
func (it *ERC721HandlerDepositedERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721HandlerDepositedERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721HandlerDepositedERC721 represents a DepositedERC721 event raised by the ERC721Handler contract.
type ERC721HandlerDepositedERC721 struct {
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
func (_ERC721Handler *ERC721HandlerFilterer) FilterDepositedERC721(opts *bind.FilterOpts) (*ERC721HandlerDepositedERC721Iterator, error) {

	logs, sub, err := _ERC721Handler.contract.FilterLogs(opts, "DepositedERC721")
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerDepositedERC721Iterator{contract: _ERC721Handler.contract, event: "DepositedERC721", logs: logs, sub: sub}, nil
}

// WatchDepositedERC721 is a free log subscription operation binding the contract event 0x7f787dd0c844dac4f8bfc4044046cdab3be531f7eefa9b740c531e48a99725e1.
//
// Solidity: event DepositedERC721(address token, uint256 tokenId, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_ERC721Handler *ERC721HandlerFilterer) WatchDepositedERC721(opts *bind.WatchOpts, sink chan<- *ERC721HandlerDepositedERC721) (event.Subscription, error) {

	logs, sub, err := _ERC721Handler.contract.WatchLogs(opts, "DepositedERC721")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721HandlerDepositedERC721)
				if err := _ERC721Handler.contract.UnpackLog(event, "DepositedERC721", log); err != nil {
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
func (_ERC721Handler *ERC721HandlerFilterer) ParseDepositedERC721(log types.Log) (*ERC721HandlerDepositedERC721, error) {
	event := new(ERC721HandlerDepositedERC721)
	if err := _ERC721Handler.contract.UnpackLog(event, "DepositedERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721HandlerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC721Handler contract.
type ERC721HandlerInitializedIterator struct {
	Event *ERC721HandlerInitialized // Event containing the contract specifics and raw log

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
func (it *ERC721HandlerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721HandlerInitialized)
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
		it.Event = new(ERC721HandlerInitialized)
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
func (it *ERC721HandlerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721HandlerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721HandlerInitialized represents a Initialized event raised by the ERC721Handler contract.
type ERC721HandlerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC721Handler *ERC721HandlerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC721HandlerInitializedIterator, error) {

	logs, sub, err := _ERC721Handler.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerInitializedIterator{contract: _ERC721Handler.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC721Handler *ERC721HandlerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC721HandlerInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC721Handler.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721HandlerInitialized)
				if err := _ERC721Handler.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ERC721Handler *ERC721HandlerFilterer) ParseInitialized(log types.Log) (*ERC721HandlerInitialized, error) {
	event := new(ERC721HandlerInitialized)
	if err := _ERC721Handler.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721HandlerWithdrawnERC721Iterator is returned from FilterWithdrawnERC721 and is used to iterate over the raw logs and unpacked data for WithdrawnERC721 events raised by the ERC721Handler contract.
type ERC721HandlerWithdrawnERC721Iterator struct {
	Event *ERC721HandlerWithdrawnERC721 // Event containing the contract specifics and raw log

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
func (it *ERC721HandlerWithdrawnERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721HandlerWithdrawnERC721)
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
		it.Event = new(ERC721HandlerWithdrawnERC721)
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
func (it *ERC721HandlerWithdrawnERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721HandlerWithdrawnERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721HandlerWithdrawnERC721 represents a WithdrawnERC721 event raised by the ERC721Handler contract.
type ERC721HandlerWithdrawnERC721 struct {
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
func (_ERC721Handler *ERC721HandlerFilterer) FilterWithdrawnERC721(opts *bind.FilterOpts) (*ERC721HandlerWithdrawnERC721Iterator, error) {

	logs, sub, err := _ERC721Handler.contract.FilterLogs(opts, "WithdrawnERC721")
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerWithdrawnERC721Iterator{contract: _ERC721Handler.contract, event: "WithdrawnERC721", logs: logs, sub: sub}, nil
}

// WatchWithdrawnERC721 is a free log subscription operation binding the contract event 0x8548956d276e213edf284f96ff17636f64d436743236a9866a7112888c442edd.
//
// Solidity: event WithdrawnERC721(address token, uint256 tokenId, string tokenURI, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_ERC721Handler *ERC721HandlerFilterer) WatchWithdrawnERC721(opts *bind.WatchOpts, sink chan<- *ERC721HandlerWithdrawnERC721) (event.Subscription, error) {

	logs, sub, err := _ERC721Handler.contract.WatchLogs(opts, "WithdrawnERC721")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721HandlerWithdrawnERC721)
				if err := _ERC721Handler.contract.UnpackLog(event, "WithdrawnERC721", log); err != nil {
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
func (_ERC721Handler *ERC721HandlerFilterer) ParseWithdrawnERC721(log types.Log) (*ERC721HandlerWithdrawnERC721, error) {
	event := new(ERC721HandlerWithdrawnERC721)
	if err := _ERC721Handler.contract.UnpackLog(event, "WithdrawnERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
