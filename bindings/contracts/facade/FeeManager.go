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

// FeeManagerMetaData contains all meta data concerning the FeeManager contract.
var FeeManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"AddedFeeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"RemovedFeeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"UpdatedFeeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnFeeToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"}],\"name\":\"__FeeManager_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.AddFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"addFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken_\",\"type\":\"address\"}],\"name\":\"getCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"commission_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.RemoveFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"removeFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.UpdateFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"updateFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"upgradeToWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.WithdrawFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FeeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeManagerMetaData.ABI instead.
var FeeManagerABI = FeeManagerMetaData.ABI

// FeeManager is an auto generated Go binding around an Ethereum contract.
type FeeManager struct {
	FeeManagerCaller     // Read-only binding to the contract
	FeeManagerTransactor // Write-only binding to the contract
	FeeManagerFilterer   // Log filterer for contract events
}

// FeeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeManagerSession struct {
	Contract     *FeeManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeManagerCallerSession struct {
	Contract *FeeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FeeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeManagerTransactorSession struct {
	Contract     *FeeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FeeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeManagerRaw struct {
	Contract *FeeManager // Generic contract binding to access the raw methods on
}

// FeeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeManagerCallerRaw struct {
	Contract *FeeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// FeeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeManagerTransactorRaw struct {
	Contract *FeeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeManager creates a new instance of FeeManager, bound to a specific deployed contract.
func NewFeeManager(address common.Address, backend bind.ContractBackend) (*FeeManager, error) {
	contract, err := bindFeeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeManager{FeeManagerCaller: FeeManagerCaller{contract: contract}, FeeManagerTransactor: FeeManagerTransactor{contract: contract}, FeeManagerFilterer: FeeManagerFilterer{contract: contract}}, nil
}

// NewFeeManagerCaller creates a new read-only instance of FeeManager, bound to a specific deployed contract.
func NewFeeManagerCaller(address common.Address, caller bind.ContractCaller) (*FeeManagerCaller, error) {
	contract, err := bindFeeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeManagerCaller{contract: contract}, nil
}

// NewFeeManagerTransactor creates a new write-only instance of FeeManager, bound to a specific deployed contract.
func NewFeeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeManagerTransactor, error) {
	contract, err := bindFeeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeManagerTransactor{contract: contract}, nil
}

// NewFeeManagerFilterer creates a new log filterer instance of FeeManager, bound to a specific deployed contract.
func NewFeeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeManagerFilterer, error) {
	contract, err := bindFeeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeManagerFilterer{contract: contract}, nil
}

// bindFeeManager binds a generic wrapper to an already deployed contract.
func bindFeeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeManager *FeeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeManager.Contract.FeeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeManager *FeeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeManager.Contract.FeeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeManager *FeeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeManager.Contract.FeeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeManager *FeeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeManager *FeeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeManager *FeeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeManager.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_FeeManager *FeeManagerCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_FeeManager *FeeManagerSession) Bridge() (common.Address, error) {
	return _FeeManager.Contract.Bridge(&_FeeManager.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_FeeManager *FeeManagerCallerSession) Bridge() (common.Address, error) {
	return _FeeManager.Contract.Bridge(&_FeeManager.CallOpts)
}

// GetCommission is a free data retrieval call binding the contract method 0xf334dcb2.
//
// Solidity: function getCommission(address feeToken_) view returns(uint256 commission_)
func (_FeeManager *FeeManagerCaller) GetCommission(opts *bind.CallOpts, feeToken_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "getCommission", feeToken_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommission is a free data retrieval call binding the contract method 0xf334dcb2.
//
// Solidity: function getCommission(address feeToken_) view returns(uint256 commission_)
func (_FeeManager *FeeManagerSession) GetCommission(feeToken_ common.Address) (*big.Int, error) {
	return _FeeManager.Contract.GetCommission(&_FeeManager.CallOpts, feeToken_)
}

// GetCommission is a free data retrieval call binding the contract method 0xf334dcb2.
//
// Solidity: function getCommission(address feeToken_) view returns(uint256 commission_)
func (_FeeManager *FeeManagerCallerSession) GetCommission(feeToken_ common.Address) (*big.Int, error) {
	return _FeeManager.Contract.GetCommission(&_FeeManager.CallOpts, feeToken_)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FeeManager *FeeManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FeeManager *FeeManagerSession) ProxiableUUID() ([32]byte, error) {
	return _FeeManager.Contract.ProxiableUUID(&_FeeManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FeeManager *FeeManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _FeeManager.Contract.ProxiableUUID(&_FeeManager.CallOpts)
}

// FeeManagerInit is a paid mutator transaction binding the contract method 0xa77c8c86.
//
// Solidity: function __FeeManager_init(address bridge_) returns()
func (_FeeManager *FeeManagerTransactor) FeeManagerInit(opts *bind.TransactOpts, bridge_ common.Address) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "__FeeManager_init", bridge_)
}

// FeeManagerInit is a paid mutator transaction binding the contract method 0xa77c8c86.
//
// Solidity: function __FeeManager_init(address bridge_) returns()
func (_FeeManager *FeeManagerSession) FeeManagerInit(bridge_ common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.FeeManagerInit(&_FeeManager.TransactOpts, bridge_)
}

// FeeManagerInit is a paid mutator transaction binding the contract method 0xa77c8c86.
//
// Solidity: function __FeeManager_init(address bridge_) returns()
func (_FeeManager *FeeManagerTransactorSession) FeeManagerInit(bridge_ common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.FeeManagerInit(&_FeeManager.TransactOpts, bridge_)
}

// AddFeeToken is a paid mutator transaction binding the contract method 0x4915b86e.
//
// Solidity: function addFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManager *FeeManagerTransactor) AddFeeToken(opts *bind.TransactOpts, params_ IFeeManagerAddFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "addFeeToken", params_)
}

// AddFeeToken is a paid mutator transaction binding the contract method 0x4915b86e.
//
// Solidity: function addFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManager *FeeManagerSession) AddFeeToken(params_ IFeeManagerAddFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManager.Contract.AddFeeToken(&_FeeManager.TransactOpts, params_)
}

// AddFeeToken is a paid mutator transaction binding the contract method 0x4915b86e.
//
// Solidity: function addFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManager *FeeManagerTransactorSession) AddFeeToken(params_ IFeeManagerAddFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManager.Contract.AddFeeToken(&_FeeManager.TransactOpts, params_)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x4eb5943a.
//
// Solidity: function removeFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManager *FeeManagerTransactor) RemoveFeeToken(opts *bind.TransactOpts, params_ IFeeManagerRemoveFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "removeFeeToken", params_)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x4eb5943a.
//
// Solidity: function removeFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManager *FeeManagerSession) RemoveFeeToken(params_ IFeeManagerRemoveFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManager.Contract.RemoveFeeToken(&_FeeManager.TransactOpts, params_)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x4eb5943a.
//
// Solidity: function removeFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManager *FeeManagerTransactorSession) RemoveFeeToken(params_ IFeeManagerRemoveFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManager.Contract.RemoveFeeToken(&_FeeManager.TransactOpts, params_)
}

// UpdateFeeToken is a paid mutator transaction binding the contract method 0x63925ea2.
//
// Solidity: function updateFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManager *FeeManagerTransactor) UpdateFeeToken(opts *bind.TransactOpts, params_ IFeeManagerUpdateFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "updateFeeToken", params_)
}

// UpdateFeeToken is a paid mutator transaction binding the contract method 0x63925ea2.
//
// Solidity: function updateFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManager *FeeManagerSession) UpdateFeeToken(params_ IFeeManagerUpdateFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManager.Contract.UpdateFeeToken(&_FeeManager.TransactOpts, params_)
}

// UpdateFeeToken is a paid mutator transaction binding the contract method 0x63925ea2.
//
// Solidity: function updateFeeToken((address[],uint256[],bytes) params_) returns()
func (_FeeManager *FeeManagerTransactorSession) UpdateFeeToken(params_ IFeeManagerUpdateFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManager.Contract.UpdateFeeToken(&_FeeManager.TransactOpts, params_)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FeeManager *FeeManagerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FeeManager *FeeManagerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.UpgradeTo(&_FeeManager.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FeeManager *FeeManagerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.UpgradeTo(&_FeeManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FeeManager *FeeManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FeeManager *FeeManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FeeManager.Contract.UpgradeToAndCall(&_FeeManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FeeManager *FeeManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FeeManager.Contract.UpgradeToAndCall(&_FeeManager.TransactOpts, newImplementation, data)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_FeeManager *FeeManagerTransactor) UpgradeToWithSig(opts *bind.TransactOpts, newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "upgradeToWithSig", newImplementation_, signature_)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_FeeManager *FeeManagerSession) UpgradeToWithSig(newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _FeeManager.Contract.UpgradeToWithSig(&_FeeManager.TransactOpts, newImplementation_, signature_)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_FeeManager *FeeManagerTransactorSession) UpgradeToWithSig(newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _FeeManager.Contract.UpgradeToWithSig(&_FeeManager.TransactOpts, newImplementation_, signature_)
}

// WithdrawFeeToken is a paid mutator transaction binding the contract method 0x9f33cf22.
//
// Solidity: function withdrawFeeToken((address,address[],uint256[],bytes) params_) returns()
func (_FeeManager *FeeManagerTransactor) WithdrawFeeToken(opts *bind.TransactOpts, params_ IFeeManagerWithdrawFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "withdrawFeeToken", params_)
}

// WithdrawFeeToken is a paid mutator transaction binding the contract method 0x9f33cf22.
//
// Solidity: function withdrawFeeToken((address,address[],uint256[],bytes) params_) returns()
func (_FeeManager *FeeManagerSession) WithdrawFeeToken(params_ IFeeManagerWithdrawFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManager.Contract.WithdrawFeeToken(&_FeeManager.TransactOpts, params_)
}

// WithdrawFeeToken is a paid mutator transaction binding the contract method 0x9f33cf22.
//
// Solidity: function withdrawFeeToken((address,address[],uint256[],bytes) params_) returns()
func (_FeeManager *FeeManagerTransactorSession) WithdrawFeeToken(params_ IFeeManagerWithdrawFeeTokenParameters) (*types.Transaction, error) {
	return _FeeManager.Contract.WithdrawFeeToken(&_FeeManager.TransactOpts, params_)
}

// FeeManagerAddedFeeTokenIterator is returned from FilterAddedFeeToken and is used to iterate over the raw logs and unpacked data for AddedFeeToken events raised by the FeeManager contract.
type FeeManagerAddedFeeTokenIterator struct {
	Event *FeeManagerAddedFeeToken // Event containing the contract specifics and raw log

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
func (it *FeeManagerAddedFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerAddedFeeToken)
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
		it.Event = new(FeeManagerAddedFeeToken)
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
func (it *FeeManagerAddedFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerAddedFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerAddedFeeToken represents a AddedFeeToken event raised by the FeeManager contract.
type FeeManagerAddedFeeToken struct {
	FeeToken  common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddedFeeToken is a free log retrieval operation binding the contract event 0x5d966ab5c51224ee3c5146e53a0b57969d97d0e1779f52a1017813c13344efc4.
//
// Solidity: event AddedFeeToken(address feeToken, uint256 feeAmount)
func (_FeeManager *FeeManagerFilterer) FilterAddedFeeToken(opts *bind.FilterOpts) (*FeeManagerAddedFeeTokenIterator, error) {

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "AddedFeeToken")
	if err != nil {
		return nil, err
	}
	return &FeeManagerAddedFeeTokenIterator{contract: _FeeManager.contract, event: "AddedFeeToken", logs: logs, sub: sub}, nil
}

// WatchAddedFeeToken is a free log subscription operation binding the contract event 0x5d966ab5c51224ee3c5146e53a0b57969d97d0e1779f52a1017813c13344efc4.
//
// Solidity: event AddedFeeToken(address feeToken, uint256 feeAmount)
func (_FeeManager *FeeManagerFilterer) WatchAddedFeeToken(opts *bind.WatchOpts, sink chan<- *FeeManagerAddedFeeToken) (event.Subscription, error) {

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "AddedFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerAddedFeeToken)
				if err := _FeeManager.contract.UnpackLog(event, "AddedFeeToken", log); err != nil {
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
func (_FeeManager *FeeManagerFilterer) ParseAddedFeeToken(log types.Log) (*FeeManagerAddedFeeToken, error) {
	event := new(FeeManagerAddedFeeToken)
	if err := _FeeManager.contract.UnpackLog(event, "AddedFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the FeeManager contract.
type FeeManagerAdminChangedIterator struct {
	Event *FeeManagerAdminChanged // Event containing the contract specifics and raw log

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
func (it *FeeManagerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerAdminChanged)
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
		it.Event = new(FeeManagerAdminChanged)
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
func (it *FeeManagerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerAdminChanged represents a AdminChanged event raised by the FeeManager contract.
type FeeManagerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_FeeManager *FeeManagerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*FeeManagerAdminChangedIterator, error) {

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &FeeManagerAdminChangedIterator{contract: _FeeManager.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_FeeManager *FeeManagerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *FeeManagerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerAdminChanged)
				if err := _FeeManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_FeeManager *FeeManagerFilterer) ParseAdminChanged(log types.Log) (*FeeManagerAdminChanged, error) {
	event := new(FeeManagerAdminChanged)
	if err := _FeeManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the FeeManager contract.
type FeeManagerBeaconUpgradedIterator struct {
	Event *FeeManagerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *FeeManagerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerBeaconUpgraded)
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
		it.Event = new(FeeManagerBeaconUpgraded)
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
func (it *FeeManagerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerBeaconUpgraded represents a BeaconUpgraded event raised by the FeeManager contract.
type FeeManagerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_FeeManager *FeeManagerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*FeeManagerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &FeeManagerBeaconUpgradedIterator{contract: _FeeManager.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_FeeManager *FeeManagerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *FeeManagerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerBeaconUpgraded)
				if err := _FeeManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_FeeManager *FeeManagerFilterer) ParseBeaconUpgraded(log types.Log) (*FeeManagerBeaconUpgraded, error) {
	event := new(FeeManagerBeaconUpgraded)
	if err := _FeeManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FeeManager contract.
type FeeManagerInitializedIterator struct {
	Event *FeeManagerInitialized // Event containing the contract specifics and raw log

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
func (it *FeeManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerInitialized)
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
		it.Event = new(FeeManagerInitialized)
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
func (it *FeeManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerInitialized represents a Initialized event raised by the FeeManager contract.
type FeeManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_FeeManager *FeeManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*FeeManagerInitializedIterator, error) {

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FeeManagerInitializedIterator{contract: _FeeManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_FeeManager *FeeManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FeeManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerInitialized)
				if err := _FeeManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_FeeManager *FeeManagerFilterer) ParseInitialized(log types.Log) (*FeeManagerInitialized, error) {
	event := new(FeeManagerInitialized)
	if err := _FeeManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerRemovedFeeTokenIterator is returned from FilterRemovedFeeToken and is used to iterate over the raw logs and unpacked data for RemovedFeeToken events raised by the FeeManager contract.
type FeeManagerRemovedFeeTokenIterator struct {
	Event *FeeManagerRemovedFeeToken // Event containing the contract specifics and raw log

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
func (it *FeeManagerRemovedFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerRemovedFeeToken)
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
		it.Event = new(FeeManagerRemovedFeeToken)
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
func (it *FeeManagerRemovedFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerRemovedFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerRemovedFeeToken represents a RemovedFeeToken event raised by the FeeManager contract.
type FeeManagerRemovedFeeToken struct {
	FeeToken  common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemovedFeeToken is a free log retrieval operation binding the contract event 0xeac220a493a0693e32f1240aae311de65b2d84391a9ba8310fcaf359d7c9bd8f.
//
// Solidity: event RemovedFeeToken(address feeToken, uint256 feeAmount)
func (_FeeManager *FeeManagerFilterer) FilterRemovedFeeToken(opts *bind.FilterOpts) (*FeeManagerRemovedFeeTokenIterator, error) {

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "RemovedFeeToken")
	if err != nil {
		return nil, err
	}
	return &FeeManagerRemovedFeeTokenIterator{contract: _FeeManager.contract, event: "RemovedFeeToken", logs: logs, sub: sub}, nil
}

// WatchRemovedFeeToken is a free log subscription operation binding the contract event 0xeac220a493a0693e32f1240aae311de65b2d84391a9ba8310fcaf359d7c9bd8f.
//
// Solidity: event RemovedFeeToken(address feeToken, uint256 feeAmount)
func (_FeeManager *FeeManagerFilterer) WatchRemovedFeeToken(opts *bind.WatchOpts, sink chan<- *FeeManagerRemovedFeeToken) (event.Subscription, error) {

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "RemovedFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerRemovedFeeToken)
				if err := _FeeManager.contract.UnpackLog(event, "RemovedFeeToken", log); err != nil {
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
func (_FeeManager *FeeManagerFilterer) ParseRemovedFeeToken(log types.Log) (*FeeManagerRemovedFeeToken, error) {
	event := new(FeeManagerRemovedFeeToken)
	if err := _FeeManager.contract.UnpackLog(event, "RemovedFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerUpdatedFeeTokenIterator is returned from FilterUpdatedFeeToken and is used to iterate over the raw logs and unpacked data for UpdatedFeeToken events raised by the FeeManager contract.
type FeeManagerUpdatedFeeTokenIterator struct {
	Event *FeeManagerUpdatedFeeToken // Event containing the contract specifics and raw log

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
func (it *FeeManagerUpdatedFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerUpdatedFeeToken)
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
		it.Event = new(FeeManagerUpdatedFeeToken)
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
func (it *FeeManagerUpdatedFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerUpdatedFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerUpdatedFeeToken represents a UpdatedFeeToken event raised by the FeeManager contract.
type FeeManagerUpdatedFeeToken struct {
	FeeToken  common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedFeeToken is a free log retrieval operation binding the contract event 0x70698787bb7456af2d4339ba1e7d059c4eda7c473412e7ed2c0d5b361a49d50a.
//
// Solidity: event UpdatedFeeToken(address feeToken, uint256 feeAmount)
func (_FeeManager *FeeManagerFilterer) FilterUpdatedFeeToken(opts *bind.FilterOpts) (*FeeManagerUpdatedFeeTokenIterator, error) {

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "UpdatedFeeToken")
	if err != nil {
		return nil, err
	}
	return &FeeManagerUpdatedFeeTokenIterator{contract: _FeeManager.contract, event: "UpdatedFeeToken", logs: logs, sub: sub}, nil
}

// WatchUpdatedFeeToken is a free log subscription operation binding the contract event 0x70698787bb7456af2d4339ba1e7d059c4eda7c473412e7ed2c0d5b361a49d50a.
//
// Solidity: event UpdatedFeeToken(address feeToken, uint256 feeAmount)
func (_FeeManager *FeeManagerFilterer) WatchUpdatedFeeToken(opts *bind.WatchOpts, sink chan<- *FeeManagerUpdatedFeeToken) (event.Subscription, error) {

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "UpdatedFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerUpdatedFeeToken)
				if err := _FeeManager.contract.UnpackLog(event, "UpdatedFeeToken", log); err != nil {
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
func (_FeeManager *FeeManagerFilterer) ParseUpdatedFeeToken(log types.Log) (*FeeManagerUpdatedFeeToken, error) {
	event := new(FeeManagerUpdatedFeeToken)
	if err := _FeeManager.contract.UnpackLog(event, "UpdatedFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the FeeManager contract.
type FeeManagerUpgradedIterator struct {
	Event *FeeManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *FeeManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerUpgraded)
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
		it.Event = new(FeeManagerUpgraded)
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
func (it *FeeManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerUpgraded represents a Upgraded event raised by the FeeManager contract.
type FeeManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FeeManager *FeeManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*FeeManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &FeeManagerUpgradedIterator{contract: _FeeManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FeeManager *FeeManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *FeeManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerUpgraded)
				if err := _FeeManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_FeeManager *FeeManagerFilterer) ParseUpgraded(log types.Log) (*FeeManagerUpgraded, error) {
	event := new(FeeManagerUpgraded)
	if err := _FeeManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerWithdrawnFeeTokenIterator is returned from FilterWithdrawnFeeToken and is used to iterate over the raw logs and unpacked data for WithdrawnFeeToken events raised by the FeeManager contract.
type FeeManagerWithdrawnFeeTokenIterator struct {
	Event *FeeManagerWithdrawnFeeToken // Event containing the contract specifics and raw log

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
func (it *FeeManagerWithdrawnFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerWithdrawnFeeToken)
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
		it.Event = new(FeeManagerWithdrawnFeeToken)
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
func (it *FeeManagerWithdrawnFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerWithdrawnFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerWithdrawnFeeToken represents a WithdrawnFeeToken event raised by the FeeManager contract.
type FeeManagerWithdrawnFeeToken struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnFeeToken is a free log retrieval operation binding the contract event 0x0221f5dbeb176269bc9dbce8b10193c570930431bbe525ae79b4599910500bf4.
//
// Solidity: event WithdrawnFeeToken(address receiver, address feeToken, uint256 amount)
func (_FeeManager *FeeManagerFilterer) FilterWithdrawnFeeToken(opts *bind.FilterOpts) (*FeeManagerWithdrawnFeeTokenIterator, error) {

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "WithdrawnFeeToken")
	if err != nil {
		return nil, err
	}
	return &FeeManagerWithdrawnFeeTokenIterator{contract: _FeeManager.contract, event: "WithdrawnFeeToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawnFeeToken is a free log subscription operation binding the contract event 0x0221f5dbeb176269bc9dbce8b10193c570930431bbe525ae79b4599910500bf4.
//
// Solidity: event WithdrawnFeeToken(address receiver, address feeToken, uint256 amount)
func (_FeeManager *FeeManagerFilterer) WatchWithdrawnFeeToken(opts *bind.WatchOpts, sink chan<- *FeeManagerWithdrawnFeeToken) (event.Subscription, error) {

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "WithdrawnFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerWithdrawnFeeToken)
				if err := _FeeManager.contract.UnpackLog(event, "WithdrawnFeeToken", log); err != nil {
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
func (_FeeManager *FeeManagerFilterer) ParseWithdrawnFeeToken(log types.Log) (*FeeManagerWithdrawnFeeToken, error) {
	event := new(FeeManagerWithdrawnFeeToken)
	if err := _FeeManager.contract.UnpackLog(event, "WithdrawnFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
