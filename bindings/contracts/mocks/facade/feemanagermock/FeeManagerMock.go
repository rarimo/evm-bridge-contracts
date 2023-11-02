// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package feemanagermock

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

// FeeManagerMockMetaData contains all meta data concerning the FeeManagerMock contract.
var FeeManagerMockMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"AddedFeeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"RemovedFeeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"UpdatedFeeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnFeeToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"}],\"name\":\"__FeeManagerMock_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"}],\"name\":\"__FeeManager_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.AddFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"addFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken_\",\"type\":\"address\"}],\"name\":\"getCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"commission_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.RemoveFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"removeFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.UpdateFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"updateFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"upgradeToWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.WithdrawFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FeeManagerMockABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeManagerMockMetaData.ABI instead.
var FeeManagerMockABI = FeeManagerMockMetaData.ABI

// FeeManagerMock is an auto generated Go binding around an Ethereum contract.
type FeeManagerMock struct {
	FeeManagerMockCaller     // Read-only binding to the contract
	FeeManagerMockTransactor // Write-only binding to the contract
	FeeManagerMockFilterer   // Log filterer for contract events
}

// FeeManagerMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeManagerMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeManagerMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeManagerMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeManagerMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeManagerMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeManagerMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeManagerMockSession struct {
	Contract     *FeeManagerMock   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeManagerMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeManagerMockCallerSession struct {
	Contract *FeeManagerMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FeeManagerMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeManagerMockTransactorSession struct {
	Contract     *FeeManagerMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FeeManagerMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeManagerMockRaw struct {
	Contract *FeeManagerMock // Generic contract binding to access the raw methods on
}

// FeeManagerMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeManagerMockCallerRaw struct {
	Contract *FeeManagerMockCaller // Generic read-only contract binding to access the raw methods on
}

// FeeManagerMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeManagerMockTransactorRaw struct {
	Contract *FeeManagerMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeManagerMock creates a new instance of FeeManagerMock, bound to a specific deployed contract.
func NewFeeManagerMock(address common.Address, backend bind.ContractBackend) (*FeeManagerMock, error) {
	contract, err := bindFeeManagerMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeManagerMock{FeeManagerMockCaller: FeeManagerMockCaller{contract: contract}, FeeManagerMockTransactor: FeeManagerMockTransactor{contract: contract}, FeeManagerMockFilterer: FeeManagerMockFilterer{contract: contract}}, nil
}

// NewFeeManagerMockCaller creates a new read-only instance of FeeManagerMock, bound to a specific deployed contract.
func NewFeeManagerMockCaller(address common.Address, caller bind.ContractCaller) (*FeeManagerMockCaller, error) {
	contract, err := bindFeeManagerMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeManagerMockCaller{contract: contract}, nil
}

// NewFeeManagerMockTransactor creates a new write-only instance of FeeManagerMock, bound to a specific deployed contract.
func NewFeeManagerMockTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeManagerMockTransactor, error) {
	contract, err := bindFeeManagerMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeManagerMockTransactor{contract: contract}, nil
}

// NewFeeManagerMockFilterer creates a new log filterer instance of FeeManagerMock, bound to a specific deployed contract.
func NewFeeManagerMockFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeManagerMockFilterer, error) {
	contract, err := bindFeeManagerMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeManagerMockFilterer{contract: contract}, nil
}

// bindFeeManagerMock binds a generic wrapper to an already deployed contract.
func bindFeeManagerMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeManagerMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeManagerMock *FeeManagerMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeManagerMock.Contract.FeeManagerMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeManagerMock *FeeManagerMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.FeeManagerMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeManagerMock *FeeManagerMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.FeeManagerMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeManagerMock *FeeManagerMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeManagerMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeManagerMock *FeeManagerMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeManagerMock *FeeManagerMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_FeeManagerMock *FeeManagerMockCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeManagerMock.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_FeeManagerMock *FeeManagerMockSession) Bridge() (common.Address, error) {
	return _FeeManagerMock.Contract.Bridge(&_FeeManagerMock.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_FeeManagerMock *FeeManagerMockCallerSession) Bridge() (common.Address, error) {
	return _FeeManagerMock.Contract.Bridge(&_FeeManagerMock.CallOpts)
}

// GetCommission is a free data retrieval call binding the contract method 0xf334dcb2.
//
// Solidity: function getCommission(address feeToken_) view returns(uint256 commission_)
func (_FeeManagerMock *FeeManagerMockCaller) GetCommission(opts *bind.CallOpts, feeToken_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeManagerMock.contract.Call(opts, &out, "getCommission", feeToken_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommission is a free data retrieval call binding the contract method 0xf334dcb2.
//
// Solidity: function getCommission(address feeToken_) view returns(uint256 commission_)
func (_FeeManagerMock *FeeManagerMockSession) GetCommission(feeToken_ common.Address) (*big.Int, error) {
	return _FeeManagerMock.Contract.GetCommission(&_FeeManagerMock.CallOpts, feeToken_)
}

// GetCommission is a free data retrieval call binding the contract method 0xf334dcb2.
//
// Solidity: function getCommission(address feeToken_) view returns(uint256 commission_)
func (_FeeManagerMock *FeeManagerMockCallerSession) GetCommission(feeToken_ common.Address) (*big.Int, error) {
	return _FeeManagerMock.Contract.GetCommission(&_FeeManagerMock.CallOpts, feeToken_)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FeeManagerMock *FeeManagerMockCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FeeManagerMock.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FeeManagerMock *FeeManagerMockSession) ProxiableUUID() ([32]byte, error) {
	return _FeeManagerMock.Contract.ProxiableUUID(&_FeeManagerMock.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FeeManagerMock *FeeManagerMockCallerSession) ProxiableUUID() ([32]byte, error) {
	return _FeeManagerMock.Contract.ProxiableUUID(&_FeeManagerMock.CallOpts)
}

// FeeManagerMockInit is a paid mutator transaction binding the contract method 0xf0653a28.
//
// Solidity: function __FeeManagerMock_init(address bridge_) returns()
func (_FeeManagerMock *FeeManagerMockTransactor) FeeManagerMockInit(opts *bind.TransactOpts, bridge_ common.Address) (*types.Transaction, error) {
	return _FeeManagerMock.contract.Transact(opts, "__FeeManagerMock_init", bridge_)
}

// FeeManagerMockInit is a paid mutator transaction binding the contract method 0xf0653a28.
//
// Solidity: function __FeeManagerMock_init(address bridge_) returns()
func (_FeeManagerMock *FeeManagerMockSession) FeeManagerMockInit(bridge_ common.Address) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.FeeManagerMockInit(&_FeeManagerMock.TransactOpts, bridge_)
}

// FeeManagerMockInit is a paid mutator transaction binding the contract method 0xf0653a28.
//
// Solidity: function __FeeManagerMock_init(address bridge_) returns()
func (_FeeManagerMock *FeeManagerMockTransactorSession) FeeManagerMockInit(bridge_ common.Address) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.FeeManagerMockInit(&_FeeManagerMock.TransactOpts, bridge_)
}

// FeeManagerInit is a paid mutator transaction binding the contract method 0xa77c8c86.
//
// Solidity: function __FeeManager_init(address bridge_) returns()
func (_FeeManagerMock *FeeManagerMockTransactor) FeeManagerInit(opts *bind.TransactOpts, bridge_ common.Address) (*types.Transaction, error) {
	return _FeeManagerMock.contract.Transact(opts, "__FeeManager_init", bridge_)
}

// FeeManagerInit is a paid mutator transaction binding the contract method 0xa77c8c86.
//
// Solidity: function __FeeManager_init(address bridge_) returns()
func (_FeeManagerMock *FeeManagerMockSession) FeeManagerInit(bridge_ common.Address) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.FeeManagerInit(&_FeeManagerMock.TransactOpts, bridge_)
}

// FeeManagerInit is a paid mutator transaction binding the contract method 0xa77c8c86.
//
// Solidity: function __FeeManager_init(address bridge_) returns()
func (_FeeManagerMock *FeeManagerMockTransactorSession) FeeManagerInit(bridge_ common.Address) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.FeeManagerInit(&_FeeManagerMock.TransactOpts, bridge_)
}

// AddFeeToken is a paid mutator transaction binding the contract method 0x4915b86e.
//
// Solidity: function addFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManagerMock *FeeManagerMockTransactor) AddFeeToken(opts *bind.TransactOpts, params_ IFeeManagerAddFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManagerMock.contract.Transact(opts, "addFeeToken", params_)
}

// AddFeeToken is a paid mutator transaction binding the contract method 0x4915b86e.
//
// Solidity: function addFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManagerMock *FeeManagerMockSession) AddFeeToken(params_ IFeeManagerAddFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.AddFeeToken(&_FeeManagerMock.TransactOpts, params_)
}

// AddFeeToken is a paid mutator transaction binding the contract method 0x4915b86e.
//
// Solidity: function addFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManagerMock *FeeManagerMockTransactorSession) AddFeeToken(params_ IFeeManagerAddFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.AddFeeToken(&_FeeManagerMock.TransactOpts, params_)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x4eb5943a.
//
// Solidity: function removeFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManagerMock *FeeManagerMockTransactor) RemoveFeeToken(opts *bind.TransactOpts, params_ IFeeManagerRemoveFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManagerMock.contract.Transact(opts, "removeFeeToken", params_)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x4eb5943a.
//
// Solidity: function removeFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManagerMock *FeeManagerMockSession) RemoveFeeToken(params_ IFeeManagerRemoveFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.RemoveFeeToken(&_FeeManagerMock.TransactOpts, params_)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x4eb5943a.
//
// Solidity: function removeFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManagerMock *FeeManagerMockTransactorSession) RemoveFeeToken(params_ IFeeManagerRemoveFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.RemoveFeeToken(&_FeeManagerMock.TransactOpts, params_)
}

// UpdateFeeToken is a paid mutator transaction binding the contract method 0x63925ea2.
//
// Solidity: function updateFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManagerMock *FeeManagerMockTransactor) UpdateFeeToken(opts *bind.TransactOpts, params_ IFeeManagerUpdateFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManagerMock.contract.Transact(opts, "updateFeeToken", params_)
}

// UpdateFeeToken is a paid mutator transaction binding the contract method 0x63925ea2.
//
// Solidity: function updateFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManagerMock *FeeManagerMockSession) UpdateFeeToken(params_ IFeeManagerUpdateFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.UpdateFeeToken(&_FeeManagerMock.TransactOpts, params_)
}

// UpdateFeeToken is a paid mutator transaction binding the contract method 0x63925ea2.
//
// Solidity: function updateFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManagerMock *FeeManagerMockTransactorSession) UpdateFeeToken(params_ IFeeManagerUpdateFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.UpdateFeeToken(&_FeeManagerMock.TransactOpts, params_)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FeeManagerMock *FeeManagerMockTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _FeeManagerMock.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FeeManagerMock *FeeManagerMockSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.UpgradeTo(&_FeeManagerMock.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FeeManagerMock *FeeManagerMockTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.UpgradeTo(&_FeeManagerMock.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FeeManagerMock *FeeManagerMockTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FeeManagerMock.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FeeManagerMock *FeeManagerMockSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.UpgradeToAndCall(&_FeeManagerMock.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FeeManagerMock *FeeManagerMockTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.UpgradeToAndCall(&_FeeManagerMock.TransactOpts, newImplementation, data)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_FeeManagerMock *FeeManagerMockTransactor) UpgradeToWithSig(opts *bind.TransactOpts, newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _FeeManagerMock.contract.Transact(opts, "upgradeToWithSig", newImplementation_, signature_)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_FeeManagerMock *FeeManagerMockSession) UpgradeToWithSig(newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.UpgradeToWithSig(&_FeeManagerMock.TransactOpts, newImplementation_, signature_)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_FeeManagerMock *FeeManagerMockTransactorSession) UpgradeToWithSig(newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.UpgradeToWithSig(&_FeeManagerMock.TransactOpts, newImplementation_, signature_)
}

// WithdrawFeeToken is a paid mutator transaction binding the contract method 0x9f33cf22.
//
// Solidity: function withdrawFeeToken((address,address[],uint256[],bytes) params_) returns()
func (_FeeManagerMock *FeeManagerMockTransactor) WithdrawFeeToken(opts *bind.TransactOpts, params_ IFeeManagerWithdrawFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManagerMock.contract.Transact(opts, "withdrawFeeToken", params_)
}

// WithdrawFeeToken is a paid mutator transaction binding the contract method 0x9f33cf22.
//
// Solidity: function withdrawFeeToken((address,address[],uint256[],bytes) params_) returns()
func (_FeeManagerMock *FeeManagerMockSession) WithdrawFeeToken(params_ IFeeManagerWithdrawFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.WithdrawFeeToken(&_FeeManagerMock.TransactOpts, params_)
}

// WithdrawFeeToken is a paid mutator transaction binding the contract method 0x9f33cf22.
//
// Solidity: function withdrawFeeToken((address,address[],uint256[],bytes) params_) returns()
func (_FeeManagerMock *FeeManagerMockTransactorSession) WithdrawFeeToken(params_ IFeeManagerWithdrawFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManagerMock.Contract.WithdrawFeeToken(&_FeeManagerMock.TransactOpts, params_)
}

// FeeManagerMockAddedFeeTokenIterator is returned from FilterAddedFeeToken and is used to iterate over the raw logs and unpacked data for AddedFeeToken events raised by the FeeManagerMock contract.
type FeeManagerMockAddedFeeTokenIterator struct {
	Event *FeeManagerMockAddedFeeToken // Event containing the contract specifics and raw log

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
func (it *FeeManagerMockAddedFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerMockAddedFeeToken)
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
		it.Event = new(FeeManagerMockAddedFeeToken)
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
func (it *FeeManagerMockAddedFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerMockAddedFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerMockAddedFeeToken represents a AddedFeeToken event raised by the FeeManagerMock contract.
type FeeManagerMockAddedFeeToken struct {
	FeeToken  common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddedFeeToken is a free log retrieval operation binding the contract event 0x5d966ab5c51224ee3c5146e53a0b57969d97d0e1779f52a1017813c13344efc4.
//
// Solidity: event AddedFeeToken(address feeToken, uint256 feeAmount)
func (_FeeManagerMock *FeeManagerMockFilterer) FilterAddedFeeToken(opts *bind.FilterOpts) (*FeeManagerMockAddedFeeTokenIterator, error) {

	logs, sub, err := _FeeManagerMock.contract.FilterLogs(opts, "AddedFeeToken")
	if err != nil {
		return nil, err
	}
	return &FeeManagerMockAddedFeeTokenIterator{contract: _FeeManagerMock.contract, event: "AddedFeeToken", logs: logs, sub: sub}, nil
}

// WatchAddedFeeToken is a free log subscription operation binding the contract event 0x5d966ab5c51224ee3c5146e53a0b57969d97d0e1779f52a1017813c13344efc4.
//
// Solidity: event AddedFeeToken(address feeToken, uint256 feeAmount)
func (_FeeManagerMock *FeeManagerMockFilterer) WatchAddedFeeToken(opts *bind.WatchOpts, sink chan<- *FeeManagerMockAddedFeeToken) (event.Subscription, error) {

	logs, sub, err := _FeeManagerMock.contract.WatchLogs(opts, "AddedFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerMockAddedFeeToken)
				if err := _FeeManagerMock.contract.UnpackLog(event, "AddedFeeToken", log); err != nil {
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
func (_FeeManagerMock *FeeManagerMockFilterer) ParseAddedFeeToken(log types.Log) (*FeeManagerMockAddedFeeToken, error) {
	event := new(FeeManagerMockAddedFeeToken)
	if err := _FeeManagerMock.contract.UnpackLog(event, "AddedFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerMockAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the FeeManagerMock contract.
type FeeManagerMockAdminChangedIterator struct {
	Event *FeeManagerMockAdminChanged // Event containing the contract specifics and raw log

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
func (it *FeeManagerMockAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerMockAdminChanged)
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
		it.Event = new(FeeManagerMockAdminChanged)
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
func (it *FeeManagerMockAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerMockAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerMockAdminChanged represents a AdminChanged event raised by the FeeManagerMock contract.
type FeeManagerMockAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_FeeManagerMock *FeeManagerMockFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*FeeManagerMockAdminChangedIterator, error) {

	logs, sub, err := _FeeManagerMock.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &FeeManagerMockAdminChangedIterator{contract: _FeeManagerMock.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_FeeManagerMock *FeeManagerMockFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *FeeManagerMockAdminChanged) (event.Subscription, error) {

	logs, sub, err := _FeeManagerMock.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerMockAdminChanged)
				if err := _FeeManagerMock.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_FeeManagerMock *FeeManagerMockFilterer) ParseAdminChanged(log types.Log) (*FeeManagerMockAdminChanged, error) {
	event := new(FeeManagerMockAdminChanged)
	if err := _FeeManagerMock.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerMockBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the FeeManagerMock contract.
type FeeManagerMockBeaconUpgradedIterator struct {
	Event *FeeManagerMockBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *FeeManagerMockBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerMockBeaconUpgraded)
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
		it.Event = new(FeeManagerMockBeaconUpgraded)
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
func (it *FeeManagerMockBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerMockBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerMockBeaconUpgraded represents a BeaconUpgraded event raised by the FeeManagerMock contract.
type FeeManagerMockBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_FeeManagerMock *FeeManagerMockFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*FeeManagerMockBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _FeeManagerMock.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &FeeManagerMockBeaconUpgradedIterator{contract: _FeeManagerMock.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_FeeManagerMock *FeeManagerMockFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *FeeManagerMockBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _FeeManagerMock.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerMockBeaconUpgraded)
				if err := _FeeManagerMock.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_FeeManagerMock *FeeManagerMockFilterer) ParseBeaconUpgraded(log types.Log) (*FeeManagerMockBeaconUpgraded, error) {
	event := new(FeeManagerMockBeaconUpgraded)
	if err := _FeeManagerMock.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerMockInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FeeManagerMock contract.
type FeeManagerMockInitializedIterator struct {
	Event *FeeManagerMockInitialized // Event containing the contract specifics and raw log

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
func (it *FeeManagerMockInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerMockInitialized)
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
		it.Event = new(FeeManagerMockInitialized)
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
func (it *FeeManagerMockInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerMockInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerMockInitialized represents a Initialized event raised by the FeeManagerMock contract.
type FeeManagerMockInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_FeeManagerMock *FeeManagerMockFilterer) FilterInitialized(opts *bind.FilterOpts) (*FeeManagerMockInitializedIterator, error) {

	logs, sub, err := _FeeManagerMock.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FeeManagerMockInitializedIterator{contract: _FeeManagerMock.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_FeeManagerMock *FeeManagerMockFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FeeManagerMockInitialized) (event.Subscription, error) {

	logs, sub, err := _FeeManagerMock.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerMockInitialized)
				if err := _FeeManagerMock.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_FeeManagerMock *FeeManagerMockFilterer) ParseInitialized(log types.Log) (*FeeManagerMockInitialized, error) {
	event := new(FeeManagerMockInitialized)
	if err := _FeeManagerMock.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerMockRemovedFeeTokenIterator is returned from FilterRemovedFeeToken and is used to iterate over the raw logs and unpacked data for RemovedFeeToken events raised by the FeeManagerMock contract.
type FeeManagerMockRemovedFeeTokenIterator struct {
	Event *FeeManagerMockRemovedFeeToken // Event containing the contract specifics and raw log

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
func (it *FeeManagerMockRemovedFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerMockRemovedFeeToken)
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
		it.Event = new(FeeManagerMockRemovedFeeToken)
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
func (it *FeeManagerMockRemovedFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerMockRemovedFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerMockRemovedFeeToken represents a RemovedFeeToken event raised by the FeeManagerMock contract.
type FeeManagerMockRemovedFeeToken struct {
	FeeToken  common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemovedFeeToken is a free log retrieval operation binding the contract event 0xeac220a493a0693e32f1240aae311de65b2d84391a9ba8310fcaf359d7c9bd8f.
//
// Solidity: event RemovedFeeToken(address feeToken, uint256 feeAmount)
func (_FeeManagerMock *FeeManagerMockFilterer) FilterRemovedFeeToken(opts *bind.FilterOpts) (*FeeManagerMockRemovedFeeTokenIterator, error) {

	logs, sub, err := _FeeManagerMock.contract.FilterLogs(opts, "RemovedFeeToken")
	if err != nil {
		return nil, err
	}
	return &FeeManagerMockRemovedFeeTokenIterator{contract: _FeeManagerMock.contract, event: "RemovedFeeToken", logs: logs, sub: sub}, nil
}

// WatchRemovedFeeToken is a free log subscription operation binding the contract event 0xeac220a493a0693e32f1240aae311de65b2d84391a9ba8310fcaf359d7c9bd8f.
//
// Solidity: event RemovedFeeToken(address feeToken, uint256 feeAmount)
func (_FeeManagerMock *FeeManagerMockFilterer) WatchRemovedFeeToken(opts *bind.WatchOpts, sink chan<- *FeeManagerMockRemovedFeeToken) (event.Subscription, error) {

	logs, sub, err := _FeeManagerMock.contract.WatchLogs(opts, "RemovedFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerMockRemovedFeeToken)
				if err := _FeeManagerMock.contract.UnpackLog(event, "RemovedFeeToken", log); err != nil {
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
func (_FeeManagerMock *FeeManagerMockFilterer) ParseRemovedFeeToken(log types.Log) (*FeeManagerMockRemovedFeeToken, error) {
	event := new(FeeManagerMockRemovedFeeToken)
	if err := _FeeManagerMock.contract.UnpackLog(event, "RemovedFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerMockUpdatedFeeTokenIterator is returned from FilterUpdatedFeeToken and is used to iterate over the raw logs and unpacked data for UpdatedFeeToken events raised by the FeeManagerMock contract.
type FeeManagerMockUpdatedFeeTokenIterator struct {
	Event *FeeManagerMockUpdatedFeeToken // Event containing the contract specifics and raw log

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
func (it *FeeManagerMockUpdatedFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerMockUpdatedFeeToken)
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
		it.Event = new(FeeManagerMockUpdatedFeeToken)
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
func (it *FeeManagerMockUpdatedFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerMockUpdatedFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerMockUpdatedFeeToken represents a UpdatedFeeToken event raised by the FeeManagerMock contract.
type FeeManagerMockUpdatedFeeToken struct {
	FeeToken  common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedFeeToken is a free log retrieval operation binding the contract event 0x70698787bb7456af2d4339ba1e7d059c4eda7c473412e7ed2c0d5b361a49d50a.
//
// Solidity: event UpdatedFeeToken(address feeToken, uint256 feeAmount)
func (_FeeManagerMock *FeeManagerMockFilterer) FilterUpdatedFeeToken(opts *bind.FilterOpts) (*FeeManagerMockUpdatedFeeTokenIterator, error) {

	logs, sub, err := _FeeManagerMock.contract.FilterLogs(opts, "UpdatedFeeToken")
	if err != nil {
		return nil, err
	}
	return &FeeManagerMockUpdatedFeeTokenIterator{contract: _FeeManagerMock.contract, event: "UpdatedFeeToken", logs: logs, sub: sub}, nil
}

// WatchUpdatedFeeToken is a free log subscription operation binding the contract event 0x70698787bb7456af2d4339ba1e7d059c4eda7c473412e7ed2c0d5b361a49d50a.
//
// Solidity: event UpdatedFeeToken(address feeToken, uint256 feeAmount)
func (_FeeManagerMock *FeeManagerMockFilterer) WatchUpdatedFeeToken(opts *bind.WatchOpts, sink chan<- *FeeManagerMockUpdatedFeeToken) (event.Subscription, error) {

	logs, sub, err := _FeeManagerMock.contract.WatchLogs(opts, "UpdatedFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerMockUpdatedFeeToken)
				if err := _FeeManagerMock.contract.UnpackLog(event, "UpdatedFeeToken", log); err != nil {
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
func (_FeeManagerMock *FeeManagerMockFilterer) ParseUpdatedFeeToken(log types.Log) (*FeeManagerMockUpdatedFeeToken, error) {
	event := new(FeeManagerMockUpdatedFeeToken)
	if err := _FeeManagerMock.contract.UnpackLog(event, "UpdatedFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerMockUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the FeeManagerMock contract.
type FeeManagerMockUpgradedIterator struct {
	Event *FeeManagerMockUpgraded // Event containing the contract specifics and raw log

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
func (it *FeeManagerMockUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerMockUpgraded)
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
		it.Event = new(FeeManagerMockUpgraded)
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
func (it *FeeManagerMockUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerMockUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerMockUpgraded represents a Upgraded event raised by the FeeManagerMock contract.
type FeeManagerMockUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FeeManagerMock *FeeManagerMockFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*FeeManagerMockUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FeeManagerMock.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &FeeManagerMockUpgradedIterator{contract: _FeeManagerMock.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FeeManagerMock *FeeManagerMockFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *FeeManagerMockUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FeeManagerMock.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerMockUpgraded)
				if err := _FeeManagerMock.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_FeeManagerMock *FeeManagerMockFilterer) ParseUpgraded(log types.Log) (*FeeManagerMockUpgraded, error) {
	event := new(FeeManagerMockUpgraded)
	if err := _FeeManagerMock.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerMockWithdrawnFeeTokenIterator is returned from FilterWithdrawnFeeToken and is used to iterate over the raw logs and unpacked data for WithdrawnFeeToken events raised by the FeeManagerMock contract.
type FeeManagerMockWithdrawnFeeTokenIterator struct {
	Event *FeeManagerMockWithdrawnFeeToken // Event containing the contract specifics and raw log

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
func (it *FeeManagerMockWithdrawnFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerMockWithdrawnFeeToken)
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
		it.Event = new(FeeManagerMockWithdrawnFeeToken)
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
func (it *FeeManagerMockWithdrawnFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerMockWithdrawnFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerMockWithdrawnFeeToken represents a WithdrawnFeeToken event raised by the FeeManagerMock contract.
type FeeManagerMockWithdrawnFeeToken struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnFeeToken is a free log retrieval operation binding the contract event 0x0221f5dbeb176269bc9dbce8b10193c570930431bbe525ae79b4599910500bf4.
//
// Solidity: event WithdrawnFeeToken(address receiver, address feeToken, uint256 amount)
func (_FeeManagerMock *FeeManagerMockFilterer) FilterWithdrawnFeeToken(opts *bind.FilterOpts) (*FeeManagerMockWithdrawnFeeTokenIterator, error) {

	logs, sub, err := _FeeManagerMock.contract.FilterLogs(opts, "WithdrawnFeeToken")
	if err != nil {
		return nil, err
	}
	return &FeeManagerMockWithdrawnFeeTokenIterator{contract: _FeeManagerMock.contract, event: "WithdrawnFeeToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawnFeeToken is a free log subscription operation binding the contract event 0x0221f5dbeb176269bc9dbce8b10193c570930431bbe525ae79b4599910500bf4.
//
// Solidity: event WithdrawnFeeToken(address receiver, address feeToken, uint256 amount)
func (_FeeManagerMock *FeeManagerMockFilterer) WatchWithdrawnFeeToken(opts *bind.WatchOpts, sink chan<- *FeeManagerMockWithdrawnFeeToken) (event.Subscription, error) {

	logs, sub, err := _FeeManagerMock.contract.WatchLogs(opts, "WithdrawnFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerMockWithdrawnFeeToken)
				if err := _FeeManagerMock.contract.UnpackLog(event, "WithdrawnFeeToken", log); err != nil {
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
func (_FeeManagerMock *FeeManagerMockFilterer) ParseWithdrawnFeeToken(log types.Log) (*FeeManagerMockWithdrawnFeeToken, error) {
	event := new(FeeManagerMockWithdrawnFeeToken)
	if err := _FeeManagerMock.contract.UnpackLog(event, "WithdrawnFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
