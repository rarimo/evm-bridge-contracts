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

// IBridgeFacadeDepositFeeERC1155Parameters is an auto generated low-level Go binding around an user-defined struct.
type IBridgeFacadeDepositFeeERC1155Parameters struct {
	FeeToken common.Address
}

// IBridgeFacadeDepositFeeERC20Parameters is an auto generated low-level Go binding around an user-defined struct.
type IBridgeFacadeDepositFeeERC20Parameters struct {
	FeeToken common.Address
}

// IBridgeFacadeDepositFeeERC721Parameters is an auto generated low-level Go binding around an user-defined struct.
type IBridgeFacadeDepositFeeERC721Parameters struct {
	FeeToken common.Address
}

// IBridgeFacadeDepositFeeNativeParameters is an auto generated low-level Go binding around an user-defined struct.
type IBridgeFacadeDepositFeeNativeParameters struct {
	FeeToken common.Address
}

// IBridgeFacadeDepositFeeSBTParameters is an auto generated low-level Go binding around an user-defined struct.
type IBridgeFacadeDepositFeeSBTParameters struct {
	FeeToken common.Address
}

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

// IBridgeFacadeMetaData contains all meta data concerning the IBridgeFacade contract.
var IBridgeFacadeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"AddedFeeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"RemovedFeeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"UpdatedFeeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnFeeToken\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.AddFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"addFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"internalType\":\"structIBridgeFacade.DepositFeeERC1155Parameters\",\"name\":\"feeParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.DepositERC1155Parameters\",\"name\":\"depositParams_\",\"type\":\"tuple\"}],\"name\":\"depositERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"internalType\":\"structIBridgeFacade.DepositFeeERC20Parameters\",\"name\":\"feeParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.DepositERC20Parameters\",\"name\":\"depositParams_\",\"type\":\"tuple\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"internalType\":\"structIBridgeFacade.DepositFeeERC721Parameters\",\"name\":\"feeParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.DepositERC721Parameters\",\"name\":\"depositParams_\",\"type\":\"tuple\"}],\"name\":\"depositERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"internalType\":\"structIBridgeFacade.DepositFeeNativeParameters\",\"name\":\"feeParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"internalType\":\"structINativeHandler.DepositNativeParameters\",\"name\":\"depositParams_\",\"type\":\"tuple\"}],\"name\":\"depositNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"internalType\":\"structIBridgeFacade.DepositFeeSBTParameters\",\"name\":\"feeParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"internalType\":\"structISBTHandler.DepositSBTParameters\",\"name\":\"depositParams_\",\"type\":\"tuple\"}],\"name\":\"depositSBT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken_\",\"type\":\"address\"}],\"name\":\"getCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"commission_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.RemoveFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"removeFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.UpdateFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"updateFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.WithdrawERC1155Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.WithdrawERC20Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.WithdrawERC721Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.WithdrawFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structINativeHandler.WithdrawNativeParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structISBTHandler.WithdrawSBTParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawSBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBridgeFacadeABI is the input ABI used to generate the binding from.
// Deprecated: Use IBridgeFacadeMetaData.ABI instead.
var IBridgeFacadeABI = IBridgeFacadeMetaData.ABI

// IBridgeFacade is an auto generated Go binding around an Ethereum contract.
type IBridgeFacade struct {
	IBridgeFacadeCaller     // Read-only binding to the contract
	IBridgeFacadeTransactor // Write-only binding to the contract
	IBridgeFacadeFilterer   // Log filterer for contract events
}

// IBridgeFacadeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBridgeFacadeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeFacadeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBridgeFacadeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeFacadeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBridgeFacadeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeFacadeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBridgeFacadeSession struct {
	Contract     *IBridgeFacade    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBridgeFacadeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBridgeFacadeCallerSession struct {
	Contract *IBridgeFacadeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IBridgeFacadeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBridgeFacadeTransactorSession struct {
	Contract     *IBridgeFacadeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IBridgeFacadeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBridgeFacadeRaw struct {
	Contract *IBridgeFacade // Generic contract binding to access the raw methods on
}

// IBridgeFacadeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBridgeFacadeCallerRaw struct {
	Contract *IBridgeFacadeCaller // Generic read-only contract binding to access the raw methods on
}

// IBridgeFacadeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBridgeFacadeTransactorRaw struct {
	Contract *IBridgeFacadeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBridgeFacade creates a new instance of IBridgeFacade, bound to a specific deployed contract.
func NewIBridgeFacade(address common.Address, backend bind.ContractBackend) (*IBridgeFacade, error) {
	contract, err := bindIBridgeFacade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBridgeFacade{IBridgeFacadeCaller: IBridgeFacadeCaller{contract: contract}, IBridgeFacadeTransactor: IBridgeFacadeTransactor{contract: contract}, IBridgeFacadeFilterer: IBridgeFacadeFilterer{contract: contract}}, nil
}

// NewIBridgeFacadeCaller creates a new read-only instance of IBridgeFacade, bound to a specific deployed contract.
func NewIBridgeFacadeCaller(address common.Address, caller bind.ContractCaller) (*IBridgeFacadeCaller, error) {
	contract, err := bindIBridgeFacade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeFacadeCaller{contract: contract}, nil
}

// NewIBridgeFacadeTransactor creates a new write-only instance of IBridgeFacade, bound to a specific deployed contract.
func NewIBridgeFacadeTransactor(address common.Address, transactor bind.ContractTransactor) (*IBridgeFacadeTransactor, error) {
	contract, err := bindIBridgeFacade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeFacadeTransactor{contract: contract}, nil
}

// NewIBridgeFacadeFilterer creates a new log filterer instance of IBridgeFacade, bound to a specific deployed contract.
func NewIBridgeFacadeFilterer(address common.Address, filterer bind.ContractFilterer) (*IBridgeFacadeFilterer, error) {
	contract, err := bindIBridgeFacade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBridgeFacadeFilterer{contract: contract}, nil
}

// bindIBridgeFacade binds a generic wrapper to an already deployed contract.
func bindIBridgeFacade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBridgeFacadeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridgeFacade *IBridgeFacadeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridgeFacade.Contract.IBridgeFacadeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridgeFacade *IBridgeFacadeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.IBridgeFacadeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridgeFacade *IBridgeFacadeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.IBridgeFacadeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridgeFacade *IBridgeFacadeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridgeFacade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridgeFacade *IBridgeFacadeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridgeFacade *IBridgeFacadeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.contract.Transact(opts, method, params...)
}

// GetCommission is a free data retrieval call binding the contract method 0xf334dcb2.
//
// Solidity: function getCommission(address feeToken_) view returns(uint256 commission_)
func (_IBridgeFacade *IBridgeFacadeCaller) GetCommission(opts *bind.CallOpts, feeToken_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBridgeFacade.contract.Call(opts, &out, "getCommission", feeToken_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommission is a free data retrieval call binding the contract method 0xf334dcb2.
//
// Solidity: function getCommission(address feeToken_) view returns(uint256 commission_)
func (_IBridgeFacade *IBridgeFacadeSession) GetCommission(feeToken_ common.Address) (*big.Int, error) {
	return _IBridgeFacade.Contract.GetCommission(&_IBridgeFacade.CallOpts, feeToken_)
}

// GetCommission is a free data retrieval call binding the contract method 0xf334dcb2.
//
// Solidity: function getCommission(address feeToken_) view returns(uint256 commission_)
func (_IBridgeFacade *IBridgeFacadeCallerSession) GetCommission(feeToken_ common.Address) (*big.Int, error) {
	return _IBridgeFacade.Contract.GetCommission(&_IBridgeFacade.CallOpts, feeToken_)
}

// AddFeeToken is a paid mutator transaction binding the contract method 0x4915b86e.
//
// Solidity: function addFeeToken((address[],uint256[],bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactor) AddFeeToken(opts *bind.TransactOpts, params_ IFeeManagerAddFeeTokenParameters) (*types.Transaction, error) {
	return _IBridgeFacade.contract.Transact(opts, "addFeeToken", params_)
}

// AddFeeToken is a paid mutator transaction binding the contract method 0x4915b86e.
//
// Solidity: function addFeeToken((address[],uint256[],bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeSession) AddFeeToken(params_ IFeeManagerAddFeeTokenParameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.AddFeeToken(&_IBridgeFacade.TransactOpts, params_)
}

// AddFeeToken is a paid mutator transaction binding the contract method 0x4915b86e.
//
// Solidity: function addFeeToken((address[],uint256[],bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactorSession) AddFeeToken(params_ IFeeManagerAddFeeTokenParameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.AddFeeToken(&_IBridgeFacade.TransactOpts, params_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x9950523c.
//
// Solidity: function depositERC1155((address) feeParams_, (address,uint256,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_IBridgeFacade *IBridgeFacadeTransactor) DepositERC1155(opts *bind.TransactOpts, feeParams_ IBridgeFacadeDepositFeeERC1155Parameters, depositParams_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.contract.Transact(opts, "depositERC1155", feeParams_, depositParams_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x9950523c.
//
// Solidity: function depositERC1155((address) feeParams_, (address,uint256,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_IBridgeFacade *IBridgeFacadeSession) DepositERC1155(feeParams_ IBridgeFacadeDepositFeeERC1155Parameters, depositParams_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.DepositERC1155(&_IBridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x9950523c.
//
// Solidity: function depositERC1155((address) feeParams_, (address,uint256,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_IBridgeFacade *IBridgeFacadeTransactorSession) DepositERC1155(feeParams_ IBridgeFacadeDepositFeeERC1155Parameters, depositParams_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.DepositERC1155(&_IBridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x73fb5a66.
//
// Solidity: function depositERC20((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_IBridgeFacade *IBridgeFacadeTransactor) DepositERC20(opts *bind.TransactOpts, feeParams_ IBridgeFacadeDepositFeeERC20Parameters, depositParams_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.contract.Transact(opts, "depositERC20", feeParams_, depositParams_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x73fb5a66.
//
// Solidity: function depositERC20((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_IBridgeFacade *IBridgeFacadeSession) DepositERC20(feeParams_ IBridgeFacadeDepositFeeERC20Parameters, depositParams_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.DepositERC20(&_IBridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x73fb5a66.
//
// Solidity: function depositERC20((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_IBridgeFacade *IBridgeFacadeTransactorSession) DepositERC20(feeParams_ IBridgeFacadeDepositFeeERC20Parameters, depositParams_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.DepositERC20(&_IBridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x1087dcd1.
//
// Solidity: function depositERC721((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_IBridgeFacade *IBridgeFacadeTransactor) DepositERC721(opts *bind.TransactOpts, feeParams_ IBridgeFacadeDepositFeeERC721Parameters, depositParams_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.contract.Transact(opts, "depositERC721", feeParams_, depositParams_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x1087dcd1.
//
// Solidity: function depositERC721((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_IBridgeFacade *IBridgeFacadeSession) DepositERC721(feeParams_ IBridgeFacadeDepositFeeERC721Parameters, depositParams_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.DepositERC721(&_IBridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x1087dcd1.
//
// Solidity: function depositERC721((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_IBridgeFacade *IBridgeFacadeTransactorSession) DepositERC721(feeParams_ IBridgeFacadeDepositFeeERC721Parameters, depositParams_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.DepositERC721(&_IBridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xa7f40f2f.
//
// Solidity: function depositNative((address) feeParams_, (uint256,(bytes32,bytes),string,string) depositParams_) payable returns()
func (_IBridgeFacade *IBridgeFacadeTransactor) DepositNative(opts *bind.TransactOpts, feeParams_ IBridgeFacadeDepositFeeNativeParameters, depositParams_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _IBridgeFacade.contract.Transact(opts, "depositNative", feeParams_, depositParams_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xa7f40f2f.
//
// Solidity: function depositNative((address) feeParams_, (uint256,(bytes32,bytes),string,string) depositParams_) payable returns()
func (_IBridgeFacade *IBridgeFacadeSession) DepositNative(feeParams_ IBridgeFacadeDepositFeeNativeParameters, depositParams_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.DepositNative(&_IBridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xa7f40f2f.
//
// Solidity: function depositNative((address) feeParams_, (uint256,(bytes32,bytes),string,string) depositParams_) payable returns()
func (_IBridgeFacade *IBridgeFacadeTransactorSession) DepositNative(feeParams_ IBridgeFacadeDepositFeeNativeParameters, depositParams_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.DepositNative(&_IBridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0xd93ff678.
//
// Solidity: function depositSBT((address) feeParams_, (address,uint256,(bytes32,bytes),string,string) depositParams_) payable returns()
func (_IBridgeFacade *IBridgeFacadeTransactor) DepositSBT(opts *bind.TransactOpts, feeParams_ IBridgeFacadeDepositFeeSBTParameters, depositParams_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _IBridgeFacade.contract.Transact(opts, "depositSBT", feeParams_, depositParams_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0xd93ff678.
//
// Solidity: function depositSBT((address) feeParams_, (address,uint256,(bytes32,bytes),string,string) depositParams_) payable returns()
func (_IBridgeFacade *IBridgeFacadeSession) DepositSBT(feeParams_ IBridgeFacadeDepositFeeSBTParameters, depositParams_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.DepositSBT(&_IBridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0xd93ff678.
//
// Solidity: function depositSBT((address) feeParams_, (address,uint256,(bytes32,bytes),string,string) depositParams_) payable returns()
func (_IBridgeFacade *IBridgeFacadeTransactorSession) DepositSBT(feeParams_ IBridgeFacadeDepositFeeSBTParameters, depositParams_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.DepositSBT(&_IBridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x4eb5943a.
//
// Solidity: function removeFeeToken((address[],uint256[],bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactor) RemoveFeeToken(opts *bind.TransactOpts, params_ IFeeManagerRemoveFeeTokenParameters) (*types.Transaction, error) {
	return _IBridgeFacade.contract.Transact(opts, "removeFeeToken", params_)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x4eb5943a.
//
// Solidity: function removeFeeToken((address[],uint256[],bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeSession) RemoveFeeToken(params_ IFeeManagerRemoveFeeTokenParameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.RemoveFeeToken(&_IBridgeFacade.TransactOpts, params_)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x4eb5943a.
//
// Solidity: function removeFeeToken((address[],uint256[],bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactorSession) RemoveFeeToken(params_ IFeeManagerRemoveFeeTokenParameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.RemoveFeeToken(&_IBridgeFacade.TransactOpts, params_)
}

// UpdateFeeToken is a paid mutator transaction binding the contract method 0x63925ea2.
//
// Solidity: function updateFeeToken((address[],uint256[],bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactor) UpdateFeeToken(opts *bind.TransactOpts, params_ IFeeManagerUpdateFeeTokenParameters) (*types.Transaction, error) {
	return _IBridgeFacade.contract.Transact(opts, "updateFeeToken", params_)
}

// UpdateFeeToken is a paid mutator transaction binding the contract method 0x63925ea2.
//
// Solidity: function updateFeeToken((address[],uint256[],bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeSession) UpdateFeeToken(params_ IFeeManagerUpdateFeeTokenParameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.UpdateFeeToken(&_IBridgeFacade.TransactOpts, params_)
}

// UpdateFeeToken is a paid mutator transaction binding the contract method 0x63925ea2.
//
// Solidity: function updateFeeToken((address[],uint256[],bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactorSession) UpdateFeeToken(params_ IFeeManagerUpdateFeeTokenParameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.UpdateFeeToken(&_IBridgeFacade.TransactOpts, params_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactor) WithdrawERC1155(opts *bind.TransactOpts, params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.contract.Transact(opts, "withdrawERC1155", params_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridgeFacade *IBridgeFacadeSession) WithdrawERC1155(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.WithdrawERC1155(&_IBridgeFacade.TransactOpts, params_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactorSession) WithdrawERC1155(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.WithdrawERC1155(&_IBridgeFacade.TransactOpts, params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactor) WithdrawERC20(opts *bind.TransactOpts, params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.contract.Transact(opts, "withdrawERC20", params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridgeFacade *IBridgeFacadeSession) WithdrawERC20(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.WithdrawERC20(&_IBridgeFacade.TransactOpts, params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactorSession) WithdrawERC20(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.WithdrawERC20(&_IBridgeFacade.TransactOpts, params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactor) WithdrawERC721(opts *bind.TransactOpts, params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.contract.Transact(opts, "withdrawERC721", params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridgeFacade *IBridgeFacadeSession) WithdrawERC721(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.WithdrawERC721(&_IBridgeFacade.TransactOpts, params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactorSession) WithdrawERC721(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.WithdrawERC721(&_IBridgeFacade.TransactOpts, params_)
}

// WithdrawFeeToken is a paid mutator transaction binding the contract method 0x9f33cf22.
//
// Solidity: function withdrawFeeToken((address,address[],uint256[],bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactor) WithdrawFeeToken(opts *bind.TransactOpts, params_ IFeeManagerWithdrawFeeTokenParameters) (*types.Transaction, error) {
	return _IBridgeFacade.contract.Transact(opts, "withdrawFeeToken", params_)
}

// WithdrawFeeToken is a paid mutator transaction binding the contract method 0x9f33cf22.
//
// Solidity: function withdrawFeeToken((address,address[],uint256[],bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeSession) WithdrawFeeToken(params_ IFeeManagerWithdrawFeeTokenParameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.WithdrawFeeToken(&_IBridgeFacade.TransactOpts, params_)
}

// WithdrawFeeToken is a paid mutator transaction binding the contract method 0x9f33cf22.
//
// Solidity: function withdrawFeeToken((address,address[],uint256[],bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactorSession) WithdrawFeeToken(params_ IFeeManagerWithdrawFeeTokenParameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.WithdrawFeeToken(&_IBridgeFacade.TransactOpts, params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactor) WithdrawNative(opts *bind.TransactOpts, params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _IBridgeFacade.contract.Transact(opts, "withdrawNative", params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeSession) WithdrawNative(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.WithdrawNative(&_IBridgeFacade.TransactOpts, params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactorSession) WithdrawNative(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.WithdrawNative(&_IBridgeFacade.TransactOpts, params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactor) WithdrawSBT(opts *bind.TransactOpts, params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _IBridgeFacade.contract.Transact(opts, "withdrawSBT", params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeSession) WithdrawSBT(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.WithdrawSBT(&_IBridgeFacade.TransactOpts, params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridgeFacade *IBridgeFacadeTransactorSession) WithdrawSBT(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _IBridgeFacade.Contract.WithdrawSBT(&_IBridgeFacade.TransactOpts, params_)
}

// IBridgeFacadeAddedFeeTokenIterator is returned from FilterAddedFeeToken and is used to iterate over the raw logs and unpacked data for AddedFeeToken events raised by the IBridgeFacade contract.
type IBridgeFacadeAddedFeeTokenIterator struct {
	Event *IBridgeFacadeAddedFeeToken // Event containing the contract specifics and raw log

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
func (it *IBridgeFacadeAddedFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeFacadeAddedFeeToken)
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
		it.Event = new(IBridgeFacadeAddedFeeToken)
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
func (it *IBridgeFacadeAddedFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeFacadeAddedFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeFacadeAddedFeeToken represents a AddedFeeToken event raised by the IBridgeFacade contract.
type IBridgeFacadeAddedFeeToken struct {
	FeeToken  common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddedFeeToken is a free log retrieval operation binding the contract event 0x5d966ab5c51224ee3c5146e53a0b57969d97d0e1779f52a1017813c13344efc4.
//
// Solidity: event AddedFeeToken(address feeToken, uint256 feeAmount)
func (_IBridgeFacade *IBridgeFacadeFilterer) FilterAddedFeeToken(opts *bind.FilterOpts) (*IBridgeFacadeAddedFeeTokenIterator, error) {

	logs, sub, err := _IBridgeFacade.contract.FilterLogs(opts, "AddedFeeToken")
	if err != nil {
		return nil, err
	}
	return &IBridgeFacadeAddedFeeTokenIterator{contract: _IBridgeFacade.contract, event: "AddedFeeToken", logs: logs, sub: sub}, nil
}

// WatchAddedFeeToken is a free log subscription operation binding the contract event 0x5d966ab5c51224ee3c5146e53a0b57969d97d0e1779f52a1017813c13344efc4.
//
// Solidity: event AddedFeeToken(address feeToken, uint256 feeAmount)
func (_IBridgeFacade *IBridgeFacadeFilterer) WatchAddedFeeToken(opts *bind.WatchOpts, sink chan<- *IBridgeFacadeAddedFeeToken) (event.Subscription, error) {

	logs, sub, err := _IBridgeFacade.contract.WatchLogs(opts, "AddedFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeFacadeAddedFeeToken)
				if err := _IBridgeFacade.contract.UnpackLog(event, "AddedFeeToken", log); err != nil {
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
func (_IBridgeFacade *IBridgeFacadeFilterer) ParseAddedFeeToken(log types.Log) (*IBridgeFacadeAddedFeeToken, error) {
	event := new(IBridgeFacadeAddedFeeToken)
	if err := _IBridgeFacade.contract.UnpackLog(event, "AddedFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeFacadeRemovedFeeTokenIterator is returned from FilterRemovedFeeToken and is used to iterate over the raw logs and unpacked data for RemovedFeeToken events raised by the IBridgeFacade contract.
type IBridgeFacadeRemovedFeeTokenIterator struct {
	Event *IBridgeFacadeRemovedFeeToken // Event containing the contract specifics and raw log

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
func (it *IBridgeFacadeRemovedFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeFacadeRemovedFeeToken)
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
		it.Event = new(IBridgeFacadeRemovedFeeToken)
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
func (it *IBridgeFacadeRemovedFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeFacadeRemovedFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeFacadeRemovedFeeToken represents a RemovedFeeToken event raised by the IBridgeFacade contract.
type IBridgeFacadeRemovedFeeToken struct {
	FeeToken  common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemovedFeeToken is a free log retrieval operation binding the contract event 0xeac220a493a0693e32f1240aae311de65b2d84391a9ba8310fcaf359d7c9bd8f.
//
// Solidity: event RemovedFeeToken(address feeToken, uint256 feeAmount)
func (_IBridgeFacade *IBridgeFacadeFilterer) FilterRemovedFeeToken(opts *bind.FilterOpts) (*IBridgeFacadeRemovedFeeTokenIterator, error) {

	logs, sub, err := _IBridgeFacade.contract.FilterLogs(opts, "RemovedFeeToken")
	if err != nil {
		return nil, err
	}
	return &IBridgeFacadeRemovedFeeTokenIterator{contract: _IBridgeFacade.contract, event: "RemovedFeeToken", logs: logs, sub: sub}, nil
}

// WatchRemovedFeeToken is a free log subscription operation binding the contract event 0xeac220a493a0693e32f1240aae311de65b2d84391a9ba8310fcaf359d7c9bd8f.
//
// Solidity: event RemovedFeeToken(address feeToken, uint256 feeAmount)
func (_IBridgeFacade *IBridgeFacadeFilterer) WatchRemovedFeeToken(opts *bind.WatchOpts, sink chan<- *IBridgeFacadeRemovedFeeToken) (event.Subscription, error) {

	logs, sub, err := _IBridgeFacade.contract.WatchLogs(opts, "RemovedFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeFacadeRemovedFeeToken)
				if err := _IBridgeFacade.contract.UnpackLog(event, "RemovedFeeToken", log); err != nil {
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
func (_IBridgeFacade *IBridgeFacadeFilterer) ParseRemovedFeeToken(log types.Log) (*IBridgeFacadeRemovedFeeToken, error) {
	event := new(IBridgeFacadeRemovedFeeToken)
	if err := _IBridgeFacade.contract.UnpackLog(event, "RemovedFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeFacadeUpdatedFeeTokenIterator is returned from FilterUpdatedFeeToken and is used to iterate over the raw logs and unpacked data for UpdatedFeeToken events raised by the IBridgeFacade contract.
type IBridgeFacadeUpdatedFeeTokenIterator struct {
	Event *IBridgeFacadeUpdatedFeeToken // Event containing the contract specifics and raw log

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
func (it *IBridgeFacadeUpdatedFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeFacadeUpdatedFeeToken)
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
		it.Event = new(IBridgeFacadeUpdatedFeeToken)
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
func (it *IBridgeFacadeUpdatedFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeFacadeUpdatedFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeFacadeUpdatedFeeToken represents a UpdatedFeeToken event raised by the IBridgeFacade contract.
type IBridgeFacadeUpdatedFeeToken struct {
	FeeToken  common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedFeeToken is a free log retrieval operation binding the contract event 0x70698787bb7456af2d4339ba1e7d059c4eda7c473412e7ed2c0d5b361a49d50a.
//
// Solidity: event UpdatedFeeToken(address feeToken, uint256 feeAmount)
func (_IBridgeFacade *IBridgeFacadeFilterer) FilterUpdatedFeeToken(opts *bind.FilterOpts) (*IBridgeFacadeUpdatedFeeTokenIterator, error) {

	logs, sub, err := _IBridgeFacade.contract.FilterLogs(opts, "UpdatedFeeToken")
	if err != nil {
		return nil, err
	}
	return &IBridgeFacadeUpdatedFeeTokenIterator{contract: _IBridgeFacade.contract, event: "UpdatedFeeToken", logs: logs, sub: sub}, nil
}

// WatchUpdatedFeeToken is a free log subscription operation binding the contract event 0x70698787bb7456af2d4339ba1e7d059c4eda7c473412e7ed2c0d5b361a49d50a.
//
// Solidity: event UpdatedFeeToken(address feeToken, uint256 feeAmount)
func (_IBridgeFacade *IBridgeFacadeFilterer) WatchUpdatedFeeToken(opts *bind.WatchOpts, sink chan<- *IBridgeFacadeUpdatedFeeToken) (event.Subscription, error) {

	logs, sub, err := _IBridgeFacade.contract.WatchLogs(opts, "UpdatedFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeFacadeUpdatedFeeToken)
				if err := _IBridgeFacade.contract.UnpackLog(event, "UpdatedFeeToken", log); err != nil {
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
func (_IBridgeFacade *IBridgeFacadeFilterer) ParseUpdatedFeeToken(log types.Log) (*IBridgeFacadeUpdatedFeeToken, error) {
	event := new(IBridgeFacadeUpdatedFeeToken)
	if err := _IBridgeFacade.contract.UnpackLog(event, "UpdatedFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeFacadeWithdrawnFeeTokenIterator is returned from FilterWithdrawnFeeToken and is used to iterate over the raw logs and unpacked data for WithdrawnFeeToken events raised by the IBridgeFacade contract.
type IBridgeFacadeWithdrawnFeeTokenIterator struct {
	Event *IBridgeFacadeWithdrawnFeeToken // Event containing the contract specifics and raw log

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
func (it *IBridgeFacadeWithdrawnFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeFacadeWithdrawnFeeToken)
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
		it.Event = new(IBridgeFacadeWithdrawnFeeToken)
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
func (it *IBridgeFacadeWithdrawnFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeFacadeWithdrawnFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeFacadeWithdrawnFeeToken represents a WithdrawnFeeToken event raised by the IBridgeFacade contract.
type IBridgeFacadeWithdrawnFeeToken struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnFeeToken is a free log retrieval operation binding the contract event 0x0221f5dbeb176269bc9dbce8b10193c570930431bbe525ae79b4599910500bf4.
//
// Solidity: event WithdrawnFeeToken(address receiver, address feeToken, uint256 amount)
func (_IBridgeFacade *IBridgeFacadeFilterer) FilterWithdrawnFeeToken(opts *bind.FilterOpts) (*IBridgeFacadeWithdrawnFeeTokenIterator, error) {

	logs, sub, err := _IBridgeFacade.contract.FilterLogs(opts, "WithdrawnFeeToken")
	if err != nil {
		return nil, err
	}
	return &IBridgeFacadeWithdrawnFeeTokenIterator{contract: _IBridgeFacade.contract, event: "WithdrawnFeeToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawnFeeToken is a free log subscription operation binding the contract event 0x0221f5dbeb176269bc9dbce8b10193c570930431bbe525ae79b4599910500bf4.
//
// Solidity: event WithdrawnFeeToken(address receiver, address feeToken, uint256 amount)
func (_IBridgeFacade *IBridgeFacadeFilterer) WatchWithdrawnFeeToken(opts *bind.WatchOpts, sink chan<- *IBridgeFacadeWithdrawnFeeToken) (event.Subscription, error) {

	logs, sub, err := _IBridgeFacade.contract.WatchLogs(opts, "WithdrawnFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeFacadeWithdrawnFeeToken)
				if err := _IBridgeFacade.contract.UnpackLog(event, "WithdrawnFeeToken", log); err != nil {
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
func (_IBridgeFacade *IBridgeFacadeFilterer) ParseWithdrawnFeeToken(log types.Log) (*IBridgeFacadeWithdrawnFeeToken, error) {
	event := new(IBridgeFacadeWithdrawnFeeToken)
	if err := _IBridgeFacade.contract.UnpackLog(event, "WithdrawnFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
