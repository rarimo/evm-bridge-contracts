// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativehandler

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

// INativeHandlerDepositNativeParameters is an auto generated low-level Go binding around an user-defined struct.
type INativeHandlerDepositNativeParameters struct {
	Amount   *big.Int
	Bundle   IBundlerBundle
	Network  string
	Receiver string
}

// INativeHandlerWithdrawNativeParameters is an auto generated low-level Go binding around an user-defined struct.
type INativeHandlerWithdrawNativeParameters struct {
	Amount     *big.Int
	Bundle     IBundlerBundle
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
}

// NativeHandlerMetaData contains all meta data concerning the NativeHandler contract.
var NativeHandlerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"name\":\"DepositedNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"WithdrawnNative\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"P\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bundleExecutorImplementation_\",\"type\":\"address\"}],\"name\":\"__Bundler_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"facade_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"}],\"name\":\"__Signers_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bundleExecutorImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"checkSignatureAndIncrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"internalType\":\"structINativeHandler.DepositNativeParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt_\",\"type\":\"bytes32\"}],\"name\":\"determineProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"getSigComponents\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"validateChangeAddressSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structINativeHandler.WithdrawNativeParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structINativeHandler.WithdrawNativeParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawNativeBundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// NativeHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeHandlerMetaData.ABI instead.
var NativeHandlerABI = NativeHandlerMetaData.ABI

// NativeHandler is an auto generated Go binding around an Ethereum contract.
type NativeHandler struct {
	NativeHandlerCaller     // Read-only binding to the contract
	NativeHandlerTransactor // Write-only binding to the contract
	NativeHandlerFilterer   // Log filterer for contract events
}

// NativeHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeHandlerSession struct {
	Contract     *NativeHandler    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NativeHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeHandlerCallerSession struct {
	Contract *NativeHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// NativeHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeHandlerTransactorSession struct {
	Contract     *NativeHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// NativeHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeHandlerRaw struct {
	Contract *NativeHandler // Generic contract binding to access the raw methods on
}

// NativeHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeHandlerCallerRaw struct {
	Contract *NativeHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// NativeHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeHandlerTransactorRaw struct {
	Contract *NativeHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeHandler creates a new instance of NativeHandler, bound to a specific deployed contract.
func NewNativeHandler(address common.Address, backend bind.ContractBackend) (*NativeHandler, error) {
	contract, err := bindNativeHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeHandler{NativeHandlerCaller: NativeHandlerCaller{contract: contract}, NativeHandlerTransactor: NativeHandlerTransactor{contract: contract}, NativeHandlerFilterer: NativeHandlerFilterer{contract: contract}}, nil
}

// NewNativeHandlerCaller creates a new read-only instance of NativeHandler, bound to a specific deployed contract.
func NewNativeHandlerCaller(address common.Address, caller bind.ContractCaller) (*NativeHandlerCaller, error) {
	contract, err := bindNativeHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeHandlerCaller{contract: contract}, nil
}

// NewNativeHandlerTransactor creates a new write-only instance of NativeHandler, bound to a specific deployed contract.
func NewNativeHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeHandlerTransactor, error) {
	contract, err := bindNativeHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeHandlerTransactor{contract: contract}, nil
}

// NewNativeHandlerFilterer creates a new log filterer instance of NativeHandler, bound to a specific deployed contract.
func NewNativeHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeHandlerFilterer, error) {
	contract, err := bindNativeHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeHandlerFilterer{contract: contract}, nil
}

// bindNativeHandler binds a generic wrapper to an already deployed contract.
func bindNativeHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeHandler *NativeHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeHandler.Contract.NativeHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeHandler *NativeHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeHandler.Contract.NativeHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeHandler *NativeHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeHandler.Contract.NativeHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeHandler *NativeHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeHandler *NativeHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeHandler *NativeHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeHandler.Contract.contract.Transact(opts, method, params...)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_NativeHandler *NativeHandlerCaller) P(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeHandler.contract.Call(opts, &out, "P")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_NativeHandler *NativeHandlerSession) P() (*big.Int, error) {
	return _NativeHandler.Contract.P(&_NativeHandler.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_NativeHandler *NativeHandlerCallerSession) P() (*big.Int, error) {
	return _NativeHandler.Contract.P(&_NativeHandler.CallOpts)
}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_NativeHandler *NativeHandlerCaller) BundleExecutorImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeHandler.contract.Call(opts, &out, "bundleExecutorImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_NativeHandler *NativeHandlerSession) BundleExecutorImplementation() (common.Address, error) {
	return _NativeHandler.Contract.BundleExecutorImplementation(&_NativeHandler.CallOpts)
}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_NativeHandler *NativeHandlerCallerSession) BundleExecutorImplementation() (common.Address, error) {
	return _NativeHandler.Contract.BundleExecutorImplementation(&_NativeHandler.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_NativeHandler *NativeHandlerCaller) ChainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeHandler.contract.Call(opts, &out, "chainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_NativeHandler *NativeHandlerSession) ChainName() (string, error) {
	return _NativeHandler.Contract.ChainName(&_NativeHandler.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_NativeHandler *NativeHandlerCallerSession) ChainName() (string, error) {
	return _NativeHandler.Contract.ChainName(&_NativeHandler.CallOpts)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_NativeHandler *NativeHandlerCaller) DetermineProxyAddress(opts *bind.CallOpts, salt_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _NativeHandler.contract.Call(opts, &out, "determineProxyAddress", salt_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_NativeHandler *NativeHandlerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _NativeHandler.Contract.DetermineProxyAddress(&_NativeHandler.CallOpts, salt_)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_NativeHandler *NativeHandlerCallerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _NativeHandler.Contract.DetermineProxyAddress(&_NativeHandler.CallOpts, salt_)
}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_NativeHandler *NativeHandlerCaller) Facade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeHandler.contract.Call(opts, &out, "facade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_NativeHandler *NativeHandlerSession) Facade() (common.Address, error) {
	return _NativeHandler.Contract.Facade(&_NativeHandler.CallOpts)
}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_NativeHandler *NativeHandlerCallerSession) Facade() (common.Address, error) {
	return _NativeHandler.Contract.Facade(&_NativeHandler.CallOpts)
}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_NativeHandler *NativeHandlerCaller) GetSigComponents(opts *bind.CallOpts, methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	var out []interface{}
	err := _NativeHandler.contract.Call(opts, &out, "getSigComponents", methodId_, contractAddress_)

	outstruct := new(struct {
		ChainName string
		Nonce     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Nonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_NativeHandler *NativeHandlerSession) GetSigComponents(methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	return _NativeHandler.Contract.GetSigComponents(&_NativeHandler.CallOpts, methodId_, contractAddress_)
}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_NativeHandler *NativeHandlerCallerSession) GetSigComponents(methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	return _NativeHandler.Contract.GetSigComponents(&_NativeHandler.CallOpts, methodId_, contractAddress_)
}

// Nonces is a free data retrieval call binding the contract method 0xed3218a2.
//
// Solidity: function nonces(address , uint8 ) view returns(uint256)
func (_NativeHandler *NativeHandlerCaller) Nonces(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) (*big.Int, error) {
	var out []interface{}
	err := _NativeHandler.contract.Call(opts, &out, "nonces", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0xed3218a2.
//
// Solidity: function nonces(address , uint8 ) view returns(uint256)
func (_NativeHandler *NativeHandlerSession) Nonces(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _NativeHandler.Contract.Nonces(&_NativeHandler.CallOpts, arg0, arg1)
}

// Nonces is a free data retrieval call binding the contract method 0xed3218a2.
//
// Solidity: function nonces(address , uint8 ) view returns(uint256)
func (_NativeHandler *NativeHandlerCallerSession) Nonces(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _NativeHandler.Contract.Nonces(&_NativeHandler.CallOpts, arg0, arg1)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_NativeHandler *NativeHandlerCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeHandler.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_NativeHandler *NativeHandlerSession) Signer() (common.Address, error) {
	return _NativeHandler.Contract.Signer(&_NativeHandler.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_NativeHandler *NativeHandlerCallerSession) Signer() (common.Address, error) {
	return _NativeHandler.Contract.Signer(&_NativeHandler.CallOpts)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x654d62aa.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_) returns()
func (_NativeHandler *NativeHandlerTransactor) BundlerInit(opts *bind.TransactOpts, bundleExecutorImplementation_ common.Address) (*types.Transaction, error) {
	return _NativeHandler.contract.Transact(opts, "__Bundler_init", bundleExecutorImplementation_)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x654d62aa.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_) returns()
func (_NativeHandler *NativeHandlerSession) BundlerInit(bundleExecutorImplementation_ common.Address) (*types.Transaction, error) {
	return _NativeHandler.Contract.BundlerInit(&_NativeHandler.TransactOpts, bundleExecutorImplementation_)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x654d62aa.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_) returns()
func (_NativeHandler *NativeHandlerTransactorSession) BundlerInit(bundleExecutorImplementation_ common.Address) (*types.Transaction, error) {
	return _NativeHandler.Contract.BundlerInit(&_NativeHandler.TransactOpts, bundleExecutorImplementation_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x3baa7892.
//
// Solidity: function __Signers_init(address signer_, address facade_, string chainName_) returns()
func (_NativeHandler *NativeHandlerTransactor) SignersInit(opts *bind.TransactOpts, signer_ common.Address, facade_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _NativeHandler.contract.Transact(opts, "__Signers_init", signer_, facade_, chainName_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x3baa7892.
//
// Solidity: function __Signers_init(address signer_, address facade_, string chainName_) returns()
func (_NativeHandler *NativeHandlerSession) SignersInit(signer_ common.Address, facade_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _NativeHandler.Contract.SignersInit(&_NativeHandler.TransactOpts, signer_, facade_, chainName_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x3baa7892.
//
// Solidity: function __Signers_init(address signer_, address facade_, string chainName_) returns()
func (_NativeHandler *NativeHandlerTransactorSession) SignersInit(signer_ common.Address, facade_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _NativeHandler.Contract.SignersInit(&_NativeHandler.TransactOpts, signer_, facade_, chainName_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_NativeHandler *NativeHandlerTransactor) CheckSignatureAndIncrementNonce(opts *bind.TransactOpts, methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _NativeHandler.contract.Transact(opts, "checkSignatureAndIncrementNonce", methodId_, contractAddress_, signHash_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_NativeHandler *NativeHandlerSession) CheckSignatureAndIncrementNonce(methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _NativeHandler.Contract.CheckSignatureAndIncrementNonce(&_NativeHandler.TransactOpts, methodId_, contractAddress_, signHash_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_NativeHandler *NativeHandlerTransactorSession) CheckSignatureAndIncrementNonce(methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _NativeHandler.Contract.CheckSignatureAndIncrementNonce(&_NativeHandler.TransactOpts, methodId_, contractAddress_, signHash_, signature_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xb27588e5.
//
// Solidity: function depositNative((uint256,(bytes32,bytes),string,string) params_) payable returns()
func (_NativeHandler *NativeHandlerTransactor) DepositNative(opts *bind.TransactOpts, params_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _NativeHandler.contract.Transact(opts, "depositNative", params_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xb27588e5.
//
// Solidity: function depositNative((uint256,(bytes32,bytes),string,string) params_) payable returns()
func (_NativeHandler *NativeHandlerSession) DepositNative(params_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _NativeHandler.Contract.DepositNative(&_NativeHandler.TransactOpts, params_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xb27588e5.
//
// Solidity: function depositNative((uint256,(bytes32,bytes),string,string) params_) payable returns()
func (_NativeHandler *NativeHandlerTransactorSession) DepositNative(params_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _NativeHandler.Contract.DepositNative(&_NativeHandler.TransactOpts, params_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_NativeHandler *NativeHandlerTransactor) ValidateChangeAddressSignature(opts *bind.TransactOpts, methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _NativeHandler.contract.Transact(opts, "validateChangeAddressSignature", methodId_, contractAddress_, newAddress_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_NativeHandler *NativeHandlerSession) ValidateChangeAddressSignature(methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _NativeHandler.Contract.ValidateChangeAddressSignature(&_NativeHandler.TransactOpts, methodId_, contractAddress_, newAddress_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_NativeHandler *NativeHandlerTransactorSession) ValidateChangeAddressSignature(methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _NativeHandler.Contract.ValidateChangeAddressSignature(&_NativeHandler.TransactOpts, methodId_, contractAddress_, newAddress_, signature_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_NativeHandler *NativeHandlerTransactor) WithdrawNative(opts *bind.TransactOpts, params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _NativeHandler.contract.Transact(opts, "withdrawNative", params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_NativeHandler *NativeHandlerSession) WithdrawNative(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _NativeHandler.Contract.WithdrawNative(&_NativeHandler.TransactOpts, params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_NativeHandler *NativeHandlerTransactorSession) WithdrawNative(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _NativeHandler.Contract.WithdrawNative(&_NativeHandler.TransactOpts, params_)
}

// WithdrawNativeBundle is a paid mutator transaction binding the contract method 0xc4f54e21.
//
// Solidity: function withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_NativeHandler *NativeHandlerTransactor) WithdrawNativeBundle(opts *bind.TransactOpts, params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _NativeHandler.contract.Transact(opts, "withdrawNativeBundle", params_)
}

// WithdrawNativeBundle is a paid mutator transaction binding the contract method 0xc4f54e21.
//
// Solidity: function withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_NativeHandler *NativeHandlerSession) WithdrawNativeBundle(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _NativeHandler.Contract.WithdrawNativeBundle(&_NativeHandler.TransactOpts, params_)
}

// WithdrawNativeBundle is a paid mutator transaction binding the contract method 0xc4f54e21.
//
// Solidity: function withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_NativeHandler *NativeHandlerTransactorSession) WithdrawNativeBundle(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _NativeHandler.Contract.WithdrawNativeBundle(&_NativeHandler.TransactOpts, params_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeHandler *NativeHandlerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeHandler.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeHandler *NativeHandlerSession) Receive() (*types.Transaction, error) {
	return _NativeHandler.Contract.Receive(&_NativeHandler.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeHandler *NativeHandlerTransactorSession) Receive() (*types.Transaction, error) {
	return _NativeHandler.Contract.Receive(&_NativeHandler.TransactOpts)
}

// NativeHandlerDepositedNativeIterator is returned from FilterDepositedNative and is used to iterate over the raw logs and unpacked data for DepositedNative events raised by the NativeHandler contract.
type NativeHandlerDepositedNativeIterator struct {
	Event *NativeHandlerDepositedNative // Event containing the contract specifics and raw log

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
func (it *NativeHandlerDepositedNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeHandlerDepositedNative)
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
		it.Event = new(NativeHandlerDepositedNative)
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
func (it *NativeHandlerDepositedNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeHandlerDepositedNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeHandlerDepositedNative represents a DepositedNative event raised by the NativeHandler contract.
type NativeHandlerDepositedNative struct {
	Amount   *big.Int
	Salt     [32]byte
	Bundle   []byte
	Network  string
	Receiver string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDepositedNative is a free log retrieval operation binding the contract event 0x9a47c8733424880a9e86a368eff95da5e7d36b68474a95eb097be2e43c116f27.
//
// Solidity: event DepositedNative(uint256 amount, bytes32 salt, bytes bundle, string network, string receiver)
func (_NativeHandler *NativeHandlerFilterer) FilterDepositedNative(opts *bind.FilterOpts) (*NativeHandlerDepositedNativeIterator, error) {

	logs, sub, err := _NativeHandler.contract.FilterLogs(opts, "DepositedNative")
	if err != nil {
		return nil, err
	}
	return &NativeHandlerDepositedNativeIterator{contract: _NativeHandler.contract, event: "DepositedNative", logs: logs, sub: sub}, nil
}

// WatchDepositedNative is a free log subscription operation binding the contract event 0x9a47c8733424880a9e86a368eff95da5e7d36b68474a95eb097be2e43c116f27.
//
// Solidity: event DepositedNative(uint256 amount, bytes32 salt, bytes bundle, string network, string receiver)
func (_NativeHandler *NativeHandlerFilterer) WatchDepositedNative(opts *bind.WatchOpts, sink chan<- *NativeHandlerDepositedNative) (event.Subscription, error) {

	logs, sub, err := _NativeHandler.contract.WatchLogs(opts, "DepositedNative")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeHandlerDepositedNative)
				if err := _NativeHandler.contract.UnpackLog(event, "DepositedNative", log); err != nil {
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

// ParseDepositedNative is a log parse operation binding the contract event 0x9a47c8733424880a9e86a368eff95da5e7d36b68474a95eb097be2e43c116f27.
//
// Solidity: event DepositedNative(uint256 amount, bytes32 salt, bytes bundle, string network, string receiver)
func (_NativeHandler *NativeHandlerFilterer) ParseDepositedNative(log types.Log) (*NativeHandlerDepositedNative, error) {
	event := new(NativeHandlerDepositedNative)
	if err := _NativeHandler.contract.UnpackLog(event, "DepositedNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeHandlerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NativeHandler contract.
type NativeHandlerInitializedIterator struct {
	Event *NativeHandlerInitialized // Event containing the contract specifics and raw log

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
func (it *NativeHandlerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeHandlerInitialized)
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
		it.Event = new(NativeHandlerInitialized)
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
func (it *NativeHandlerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeHandlerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeHandlerInitialized represents a Initialized event raised by the NativeHandler contract.
type NativeHandlerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NativeHandler *NativeHandlerFilterer) FilterInitialized(opts *bind.FilterOpts) (*NativeHandlerInitializedIterator, error) {

	logs, sub, err := _NativeHandler.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NativeHandlerInitializedIterator{contract: _NativeHandler.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NativeHandler *NativeHandlerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NativeHandlerInitialized) (event.Subscription, error) {

	logs, sub, err := _NativeHandler.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeHandlerInitialized)
				if err := _NativeHandler.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NativeHandler *NativeHandlerFilterer) ParseInitialized(log types.Log) (*NativeHandlerInitialized, error) {
	event := new(NativeHandlerInitialized)
	if err := _NativeHandler.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeHandlerWithdrawnNativeIterator is returned from FilterWithdrawnNative and is used to iterate over the raw logs and unpacked data for WithdrawnNative events raised by the NativeHandler contract.
type NativeHandlerWithdrawnNativeIterator struct {
	Event *NativeHandlerWithdrawnNative // Event containing the contract specifics and raw log

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
func (it *NativeHandlerWithdrawnNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeHandlerWithdrawnNative)
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
		it.Event = new(NativeHandlerWithdrawnNative)
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
func (it *NativeHandlerWithdrawnNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeHandlerWithdrawnNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeHandlerWithdrawnNative represents a WithdrawnNative event raised by the NativeHandler contract.
type NativeHandlerWithdrawnNative struct {
	Amount     *big.Int
	Salt       [32]byte
	Bundle     []byte
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnNative is a free log retrieval operation binding the contract event 0x319eed2457c60bc7ee23fdc5be5941c7b628c01f6f2c526dedeca75054d66b18.
//
// Solidity: event WithdrawnNative(uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_NativeHandler *NativeHandlerFilterer) FilterWithdrawnNative(opts *bind.FilterOpts) (*NativeHandlerWithdrawnNativeIterator, error) {

	logs, sub, err := _NativeHandler.contract.FilterLogs(opts, "WithdrawnNative")
	if err != nil {
		return nil, err
	}
	return &NativeHandlerWithdrawnNativeIterator{contract: _NativeHandler.contract, event: "WithdrawnNative", logs: logs, sub: sub}, nil
}

// WatchWithdrawnNative is a free log subscription operation binding the contract event 0x319eed2457c60bc7ee23fdc5be5941c7b628c01f6f2c526dedeca75054d66b18.
//
// Solidity: event WithdrawnNative(uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_NativeHandler *NativeHandlerFilterer) WatchWithdrawnNative(opts *bind.WatchOpts, sink chan<- *NativeHandlerWithdrawnNative) (event.Subscription, error) {

	logs, sub, err := _NativeHandler.contract.WatchLogs(opts, "WithdrawnNative")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeHandlerWithdrawnNative)
				if err := _NativeHandler.contract.UnpackLog(event, "WithdrawnNative", log); err != nil {
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

// ParseWithdrawnNative is a log parse operation binding the contract event 0x319eed2457c60bc7ee23fdc5be5941c7b628c01f6f2c526dedeca75054d66b18.
//
// Solidity: event WithdrawnNative(uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_NativeHandler *NativeHandlerFilterer) ParseWithdrawnNative(log types.Log) (*NativeHandlerWithdrawnNative, error) {
	event := new(NativeHandlerWithdrawnNative)
	if err := _NativeHandler.contract.UnpackLog(event, "WithdrawnNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
