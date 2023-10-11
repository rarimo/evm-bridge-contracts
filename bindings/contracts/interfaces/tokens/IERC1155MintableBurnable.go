// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokens

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

// IERC1155MintableBurnableMetaData contains all meta data concerning the IERC1155MintableBurnable contract.
var IERC1155MintableBurnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"}],\"name\":\"mintTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IERC1155MintableBurnableABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC1155MintableBurnableMetaData.ABI instead.
var IERC1155MintableBurnableABI = IERC1155MintableBurnableMetaData.ABI

// IERC1155MintableBurnable is an auto generated Go binding around an Ethereum contract.
type IERC1155MintableBurnable struct {
	IERC1155MintableBurnableCaller     // Read-only binding to the contract
	IERC1155MintableBurnableTransactor // Write-only binding to the contract
	IERC1155MintableBurnableFilterer   // Log filterer for contract events
}

// IERC1155MintableBurnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1155MintableBurnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155MintableBurnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1155MintableBurnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155MintableBurnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1155MintableBurnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155MintableBurnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1155MintableBurnableSession struct {
	Contract     *IERC1155MintableBurnable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC1155MintableBurnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1155MintableBurnableCallerSession struct {
	Contract *IERC1155MintableBurnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IERC1155MintableBurnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1155MintableBurnableTransactorSession struct {
	Contract     *IERC1155MintableBurnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IERC1155MintableBurnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1155MintableBurnableRaw struct {
	Contract *IERC1155MintableBurnable // Generic contract binding to access the raw methods on
}

// IERC1155MintableBurnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1155MintableBurnableCallerRaw struct {
	Contract *IERC1155MintableBurnableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC1155MintableBurnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1155MintableBurnableTransactorRaw struct {
	Contract *IERC1155MintableBurnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1155MintableBurnable creates a new instance of IERC1155MintableBurnable, bound to a specific deployed contract.
func NewIERC1155MintableBurnable(address common.Address, backend bind.ContractBackend) (*IERC1155MintableBurnable, error) {
	contract, err := bindIERC1155MintableBurnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1155MintableBurnable{IERC1155MintableBurnableCaller: IERC1155MintableBurnableCaller{contract: contract}, IERC1155MintableBurnableTransactor: IERC1155MintableBurnableTransactor{contract: contract}, IERC1155MintableBurnableFilterer: IERC1155MintableBurnableFilterer{contract: contract}}, nil
}

// NewIERC1155MintableBurnableCaller creates a new read-only instance of IERC1155MintableBurnable, bound to a specific deployed contract.
func NewIERC1155MintableBurnableCaller(address common.Address, caller bind.ContractCaller) (*IERC1155MintableBurnableCaller, error) {
	contract, err := bindIERC1155MintableBurnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155MintableBurnableCaller{contract: contract}, nil
}

// NewIERC1155MintableBurnableTransactor creates a new write-only instance of IERC1155MintableBurnable, bound to a specific deployed contract.
func NewIERC1155MintableBurnableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC1155MintableBurnableTransactor, error) {
	contract, err := bindIERC1155MintableBurnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155MintableBurnableTransactor{contract: contract}, nil
}

// NewIERC1155MintableBurnableFilterer creates a new log filterer instance of IERC1155MintableBurnable, bound to a specific deployed contract.
func NewIERC1155MintableBurnableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC1155MintableBurnableFilterer, error) {
	contract, err := bindIERC1155MintableBurnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1155MintableBurnableFilterer{contract: contract}, nil
}

// bindIERC1155MintableBurnable binds a generic wrapper to an already deployed contract.
func bindIERC1155MintableBurnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC1155MintableBurnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155MintableBurnable *IERC1155MintableBurnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155MintableBurnable.Contract.IERC1155MintableBurnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155MintableBurnable *IERC1155MintableBurnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.Contract.IERC1155MintableBurnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155MintableBurnable *IERC1155MintableBurnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.Contract.IERC1155MintableBurnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155MintableBurnable *IERC1155MintableBurnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155MintableBurnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155MintableBurnable *IERC1155MintableBurnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155MintableBurnable *IERC1155MintableBurnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC1155MintableBurnable.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _IERC1155MintableBurnable.Contract.BalanceOf(&_IERC1155MintableBurnable.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _IERC1155MintableBurnable.Contract.BalanceOf(&_IERC1155MintableBurnable.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_IERC1155MintableBurnable *IERC1155MintableBurnableCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IERC1155MintableBurnable.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_IERC1155MintableBurnable *IERC1155MintableBurnableSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _IERC1155MintableBurnable.Contract.BalanceOfBatch(&_IERC1155MintableBurnable.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_IERC1155MintableBurnable *IERC1155MintableBurnableCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _IERC1155MintableBurnable.Contract.BalanceOfBatch(&_IERC1155MintableBurnable.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC1155MintableBurnable.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _IERC1155MintableBurnable.Contract.IsApprovedForAll(&_IERC1155MintableBurnable.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _IERC1155MintableBurnable.Contract.IsApprovedForAll(&_IERC1155MintableBurnable.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC1155MintableBurnable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC1155MintableBurnable.Contract.SupportsInterface(&_IERC1155MintableBurnable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC1155MintableBurnable.Contract.SupportsInterface(&_IERC1155MintableBurnable.CallOpts, interfaceId)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x124d91e5.
//
// Solidity: function burnFrom(address payer_, uint256 tokenId_, uint256 amount_) returns()
func (_IERC1155MintableBurnable *IERC1155MintableBurnableTransactor) BurnFrom(opts *bind.TransactOpts, payer_ common.Address, tokenId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.contract.Transact(opts, "burnFrom", payer_, tokenId_, amount_)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x124d91e5.
//
// Solidity: function burnFrom(address payer_, uint256 tokenId_, uint256 amount_) returns()
func (_IERC1155MintableBurnable *IERC1155MintableBurnableSession) BurnFrom(payer_ common.Address, tokenId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.Contract.BurnFrom(&_IERC1155MintableBurnable.TransactOpts, payer_, tokenId_, amount_)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x124d91e5.
//
// Solidity: function burnFrom(address payer_, uint256 tokenId_, uint256 amount_) returns()
func (_IERC1155MintableBurnable *IERC1155MintableBurnableTransactorSession) BurnFrom(payer_ common.Address, tokenId_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.Contract.BurnFrom(&_IERC1155MintableBurnable.TransactOpts, payer_, tokenId_, amount_)
}

// MintTo is a paid mutator transaction binding the contract method 0x3dbd5b25.
//
// Solidity: function mintTo(address receiver_, uint256 tokenId_, uint256 amount_, string tokenURI_) returns()
func (_IERC1155MintableBurnable *IERC1155MintableBurnableTransactor) MintTo(opts *bind.TransactOpts, receiver_ common.Address, tokenId_ *big.Int, amount_ *big.Int, tokenURI_ string) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.contract.Transact(opts, "mintTo", receiver_, tokenId_, amount_, tokenURI_)
}

// MintTo is a paid mutator transaction binding the contract method 0x3dbd5b25.
//
// Solidity: function mintTo(address receiver_, uint256 tokenId_, uint256 amount_, string tokenURI_) returns()
func (_IERC1155MintableBurnable *IERC1155MintableBurnableSession) MintTo(receiver_ common.Address, tokenId_ *big.Int, amount_ *big.Int, tokenURI_ string) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.Contract.MintTo(&_IERC1155MintableBurnable.TransactOpts, receiver_, tokenId_, amount_, tokenURI_)
}

// MintTo is a paid mutator transaction binding the contract method 0x3dbd5b25.
//
// Solidity: function mintTo(address receiver_, uint256 tokenId_, uint256 amount_, string tokenURI_) returns()
func (_IERC1155MintableBurnable *IERC1155MintableBurnableTransactorSession) MintTo(receiver_ common.Address, tokenId_ *big.Int, amount_ *big.Int, tokenURI_ string) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.Contract.MintTo(&_IERC1155MintableBurnable.TransactOpts, receiver_, tokenId_, amount_, tokenURI_)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_IERC1155MintableBurnable *IERC1155MintableBurnableTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_IERC1155MintableBurnable *IERC1155MintableBurnableSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.Contract.SafeBatchTransferFrom(&_IERC1155MintableBurnable.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_IERC1155MintableBurnable *IERC1155MintableBurnableTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.Contract.SafeBatchTransferFrom(&_IERC1155MintableBurnable.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_IERC1155MintableBurnable *IERC1155MintableBurnableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_IERC1155MintableBurnable *IERC1155MintableBurnableSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.Contract.SafeTransferFrom(&_IERC1155MintableBurnable.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_IERC1155MintableBurnable *IERC1155MintableBurnableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.Contract.SafeTransferFrom(&_IERC1155MintableBurnable.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IERC1155MintableBurnable *IERC1155MintableBurnableTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IERC1155MintableBurnable *IERC1155MintableBurnableSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.Contract.SetApprovalForAll(&_IERC1155MintableBurnable.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IERC1155MintableBurnable *IERC1155MintableBurnableTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _IERC1155MintableBurnable.Contract.SetApprovalForAll(&_IERC1155MintableBurnable.TransactOpts, operator, approved)
}

// IERC1155MintableBurnableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC1155MintableBurnable contract.
type IERC1155MintableBurnableApprovalForAllIterator struct {
	Event *IERC1155MintableBurnableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC1155MintableBurnableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155MintableBurnableApprovalForAll)
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
		it.Event = new(IERC1155MintableBurnableApprovalForAll)
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
func (it *IERC1155MintableBurnableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155MintableBurnableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155MintableBurnableApprovalForAll represents a ApprovalForAll event raised by the IERC1155MintableBurnable contract.
type IERC1155MintableBurnableApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*IERC1155MintableBurnableApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC1155MintableBurnable.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155MintableBurnableApprovalForAllIterator{contract: _IERC1155MintableBurnable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC1155MintableBurnableApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC1155MintableBurnable.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155MintableBurnableApprovalForAll)
				if err := _IERC1155MintableBurnable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableFilterer) ParseApprovalForAll(log types.Log) (*IERC1155MintableBurnableApprovalForAll, error) {
	event := new(IERC1155MintableBurnableApprovalForAll)
	if err := _IERC1155MintableBurnable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155MintableBurnableTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the IERC1155MintableBurnable contract.
type IERC1155MintableBurnableTransferBatchIterator struct {
	Event *IERC1155MintableBurnableTransferBatch // Event containing the contract specifics and raw log

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
func (it *IERC1155MintableBurnableTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155MintableBurnableTransferBatch)
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
		it.Event = new(IERC1155MintableBurnableTransferBatch)
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
func (it *IERC1155MintableBurnableTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155MintableBurnableTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155MintableBurnableTransferBatch represents a TransferBatch event raised by the IERC1155MintableBurnable contract.
type IERC1155MintableBurnableTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*IERC1155MintableBurnableTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC1155MintableBurnable.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155MintableBurnableTransferBatchIterator{contract: _IERC1155MintableBurnable.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *IERC1155MintableBurnableTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC1155MintableBurnable.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155MintableBurnableTransferBatch)
				if err := _IERC1155MintableBurnable.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableFilterer) ParseTransferBatch(log types.Log) (*IERC1155MintableBurnableTransferBatch, error) {
	event := new(IERC1155MintableBurnableTransferBatch)
	if err := _IERC1155MintableBurnable.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155MintableBurnableTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the IERC1155MintableBurnable contract.
type IERC1155MintableBurnableTransferSingleIterator struct {
	Event *IERC1155MintableBurnableTransferSingle // Event containing the contract specifics and raw log

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
func (it *IERC1155MintableBurnableTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155MintableBurnableTransferSingle)
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
		it.Event = new(IERC1155MintableBurnableTransferSingle)
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
func (it *IERC1155MintableBurnableTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155MintableBurnableTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155MintableBurnableTransferSingle represents a TransferSingle event raised by the IERC1155MintableBurnable contract.
type IERC1155MintableBurnableTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*IERC1155MintableBurnableTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC1155MintableBurnable.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155MintableBurnableTransferSingleIterator{contract: _IERC1155MintableBurnable.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *IERC1155MintableBurnableTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC1155MintableBurnable.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155MintableBurnableTransferSingle)
				if err := _IERC1155MintableBurnable.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableFilterer) ParseTransferSingle(log types.Log) (*IERC1155MintableBurnableTransferSingle, error) {
	event := new(IERC1155MintableBurnableTransferSingle)
	if err := _IERC1155MintableBurnable.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155MintableBurnableURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the IERC1155MintableBurnable contract.
type IERC1155MintableBurnableURIIterator struct {
	Event *IERC1155MintableBurnableURI // Event containing the contract specifics and raw log

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
func (it *IERC1155MintableBurnableURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155MintableBurnableURI)
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
		it.Event = new(IERC1155MintableBurnableURI)
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
func (it *IERC1155MintableBurnableURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155MintableBurnableURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155MintableBurnableURI represents a URI event raised by the IERC1155MintableBurnable contract.
type IERC1155MintableBurnableURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*IERC1155MintableBurnableURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC1155MintableBurnable.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155MintableBurnableURIIterator{contract: _IERC1155MintableBurnable.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *IERC1155MintableBurnableURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC1155MintableBurnable.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155MintableBurnableURI)
				if err := _IERC1155MintableBurnable.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_IERC1155MintableBurnable *IERC1155MintableBurnableFilterer) ParseURI(log types.Log) (*IERC1155MintableBurnableURI, error) {
	event := new(IERC1155MintableBurnableURI)
	if err := _IERC1155MintableBurnable.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
