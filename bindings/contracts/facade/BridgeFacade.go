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

// BridgeFacadeMetaData contains all meta data concerning the BridgeFacade contract.
var BridgeFacadeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"AddedFeeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"RemovedFeeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"UpdatedFeeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnFeeToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"}],\"name\":\"__BridgeFacade_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"}],\"name\":\"__FeeManager_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.AddFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"addFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"internalType\":\"structIBridgeFacade.DepositFeeERC1155Parameters\",\"name\":\"feeParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.DepositERC1155Parameters\",\"name\":\"depositParams_\",\"type\":\"tuple\"}],\"name\":\"depositERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"internalType\":\"structIBridgeFacade.DepositFeeERC20Parameters\",\"name\":\"feeParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.DepositERC20Parameters\",\"name\":\"depositParams_\",\"type\":\"tuple\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"internalType\":\"structIBridgeFacade.DepositFeeERC721Parameters\",\"name\":\"feeParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.DepositERC721Parameters\",\"name\":\"depositParams_\",\"type\":\"tuple\"}],\"name\":\"depositERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"internalType\":\"structIBridgeFacade.DepositFeeNativeParameters\",\"name\":\"feeParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"internalType\":\"structINativeHandler.DepositNativeParameters\",\"name\":\"depositParams_\",\"type\":\"tuple\"}],\"name\":\"depositNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"internalType\":\"structIBridgeFacade.DepositFeeSBTParameters\",\"name\":\"feeParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"internalType\":\"structISBTHandler.DepositSBTParameters\",\"name\":\"depositParams_\",\"type\":\"tuple\"}],\"name\":\"depositSBT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken_\",\"type\":\"address\"}],\"name\":\"getCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"commission_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.RemoveFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"removeFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.UpdateFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"updateFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"upgradeToWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.WithdrawERC1155Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.WithdrawERC20Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.WithdrawERC721Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFeeManager.WithdrawFeeTokenParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structINativeHandler.WithdrawNativeParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structISBTHandler.WithdrawSBTParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawSBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeFacadeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeFacadeMetaData.ABI instead.
var BridgeFacadeABI = BridgeFacadeMetaData.ABI

// BridgeFacade is an auto generated Go binding around an Ethereum contract.
type BridgeFacade struct {
	BridgeFacadeCaller     // Read-only binding to the contract
	BridgeFacadeTransactor // Write-only binding to the contract
	BridgeFacadeFilterer   // Log filterer for contract events
}

// BridgeFacadeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeFacadeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFacadeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeFacadeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFacadeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFacadeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFacadeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeFacadeSession struct {
	Contract     *BridgeFacade     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeFacadeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeFacadeCallerSession struct {
	Contract *BridgeFacadeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BridgeFacadeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeFacadeTransactorSession struct {
	Contract     *BridgeFacadeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BridgeFacadeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeFacadeRaw struct {
	Contract *BridgeFacade // Generic contract binding to access the raw methods on
}

// BridgeFacadeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeFacadeCallerRaw struct {
	Contract *BridgeFacadeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeFacadeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeFacadeTransactorRaw struct {
	Contract *BridgeFacadeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeFacade creates a new instance of BridgeFacade, bound to a specific deployed contract.
func NewBridgeFacade(address common.Address, backend bind.ContractBackend) (*BridgeFacade, error) {
	contract, err := bindBridgeFacade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeFacade{BridgeFacadeCaller: BridgeFacadeCaller{contract: contract}, BridgeFacadeTransactor: BridgeFacadeTransactor{contract: contract}, BridgeFacadeFilterer: BridgeFacadeFilterer{contract: contract}}, nil
}

// NewBridgeFacadeCaller creates a new read-only instance of BridgeFacade, bound to a specific deployed contract.
func NewBridgeFacadeCaller(address common.Address, caller bind.ContractCaller) (*BridgeFacadeCaller, error) {
	contract, err := bindBridgeFacade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeFacadeCaller{contract: contract}, nil
}

// NewBridgeFacadeTransactor creates a new write-only instance of BridgeFacade, bound to a specific deployed contract.
func NewBridgeFacadeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeFacadeTransactor, error) {
	contract, err := bindBridgeFacade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeFacadeTransactor{contract: contract}, nil
}

// NewBridgeFacadeFilterer creates a new log filterer instance of BridgeFacade, bound to a specific deployed contract.
func NewBridgeFacadeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFacadeFilterer, error) {
	contract, err := bindBridgeFacade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFacadeFilterer{contract: contract}, nil
}

// bindBridgeFacade binds a generic wrapper to an already deployed contract.
func bindBridgeFacade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeFacadeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeFacade *BridgeFacadeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeFacade.Contract.BridgeFacadeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeFacade *BridgeFacadeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeFacade.Contract.BridgeFacadeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeFacade *BridgeFacadeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeFacade.Contract.BridgeFacadeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeFacade *BridgeFacadeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeFacade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeFacade *BridgeFacadeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeFacade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeFacade *BridgeFacadeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeFacade.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BridgeFacade *BridgeFacadeCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeFacade.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BridgeFacade *BridgeFacadeSession) Bridge() (common.Address, error) {
	return _BridgeFacade.Contract.Bridge(&_BridgeFacade.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BridgeFacade *BridgeFacadeCallerSession) Bridge() (common.Address, error) {
	return _BridgeFacade.Contract.Bridge(&_BridgeFacade.CallOpts)
}

// GetCommission is a free data retrieval call binding the contract method 0xf334dcb2.
//
// Solidity: function getCommission(address feeToken_) view returns(uint256 commission_)
func (_BridgeFacade *BridgeFacadeCaller) GetCommission(opts *bind.CallOpts, feeToken_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeFacade.contract.Call(opts, &out, "getCommission", feeToken_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommission is a free data retrieval call binding the contract method 0xf334dcb2.
//
// Solidity: function getCommission(address feeToken_) view returns(uint256 commission_)
func (_BridgeFacade *BridgeFacadeSession) GetCommission(feeToken_ common.Address) (*big.Int, error) {
	return _BridgeFacade.Contract.GetCommission(&_BridgeFacade.CallOpts, feeToken_)
}

// GetCommission is a free data retrieval call binding the contract method 0xf334dcb2.
//
// Solidity: function getCommission(address feeToken_) view returns(uint256 commission_)
func (_BridgeFacade *BridgeFacadeCallerSession) GetCommission(feeToken_ common.Address) (*big.Int, error) {
	return _BridgeFacade.Contract.GetCommission(&_BridgeFacade.CallOpts, feeToken_)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeFacade *BridgeFacadeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeFacade.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeFacade *BridgeFacadeSession) ProxiableUUID() ([32]byte, error) {
	return _BridgeFacade.Contract.ProxiableUUID(&_BridgeFacade.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeFacade *BridgeFacadeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BridgeFacade.Contract.ProxiableUUID(&_BridgeFacade.CallOpts)
}

// BridgeFacadeInit is a paid mutator transaction binding the contract method 0x9c9ae4f8.
//
// Solidity: function __BridgeFacade_init(address bridge_) returns()
func (_BridgeFacade *BridgeFacadeTransactor) BridgeFacadeInit(opts *bind.TransactOpts, bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "__BridgeFacade_init", bridge_)
}

// BridgeFacadeInit is a paid mutator transaction binding the contract method 0x9c9ae4f8.
//
// Solidity: function __BridgeFacade_init(address bridge_) returns()
func (_BridgeFacade *BridgeFacadeSession) BridgeFacadeInit(bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeFacade.Contract.BridgeFacadeInit(&_BridgeFacade.TransactOpts, bridge_)
}

// BridgeFacadeInit is a paid mutator transaction binding the contract method 0x9c9ae4f8.
//
// Solidity: function __BridgeFacade_init(address bridge_) returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) BridgeFacadeInit(bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeFacade.Contract.BridgeFacadeInit(&_BridgeFacade.TransactOpts, bridge_)
}

// FeeManagerInit is a paid mutator transaction binding the contract method 0xa77c8c86.
//
// Solidity: function __FeeManager_init(address bridge_) returns()
func (_BridgeFacade *BridgeFacadeTransactor) FeeManagerInit(opts *bind.TransactOpts, bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "__FeeManager_init", bridge_)
}

// FeeManagerInit is a paid mutator transaction binding the contract method 0xa77c8c86.
//
// Solidity: function __FeeManager_init(address bridge_) returns()
func (_BridgeFacade *BridgeFacadeSession) FeeManagerInit(bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeFacade.Contract.FeeManagerInit(&_BridgeFacade.TransactOpts, bridge_)
}

// FeeManagerInit is a paid mutator transaction binding the contract method 0xa77c8c86.
//
// Solidity: function __FeeManager_init(address bridge_) returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) FeeManagerInit(bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeFacade.Contract.FeeManagerInit(&_BridgeFacade.TransactOpts, bridge_)
}

// AddFeeToken is a paid mutator transaction binding the contract method 0x4915b86e.
//
// Solidity: function addFeeToken((address[],uint256[],bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactor) AddFeeToken(opts *bind.TransactOpts, params_ IFeeManagerAddFeeTokenParameters) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "addFeeToken", params_)
}

// AddFeeToken is a paid mutator transaction binding the contract method 0x4915b86e.
//
// Solidity: function addFeeToken((address[],uint256[],bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeSession) AddFeeToken(params_ IFeeManagerAddFeeTokenParameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.AddFeeToken(&_BridgeFacade.TransactOpts, params_)
}

// AddFeeToken is a paid mutator transaction binding the contract method 0x4915b86e.
//
// Solidity: function addFeeToken((address[],uint256[],bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) AddFeeToken(params_ IFeeManagerAddFeeTokenParameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.AddFeeToken(&_BridgeFacade.TransactOpts, params_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x9950523c.
//
// Solidity: function depositERC1155((address) feeParams_, (address,uint256,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeFacade *BridgeFacadeTransactor) DepositERC1155(opts *bind.TransactOpts, feeParams_ IBridgeFacadeDepositFeeERC1155Parameters, depositParams_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "depositERC1155", feeParams_, depositParams_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x9950523c.
//
// Solidity: function depositERC1155((address) feeParams_, (address,uint256,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeFacade *BridgeFacadeSession) DepositERC1155(feeParams_ IBridgeFacadeDepositFeeERC1155Parameters, depositParams_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.DepositERC1155(&_BridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x9950523c.
//
// Solidity: function depositERC1155((address) feeParams_, (address,uint256,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) DepositERC1155(feeParams_ IBridgeFacadeDepositFeeERC1155Parameters, depositParams_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.DepositERC1155(&_BridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x73fb5a66.
//
// Solidity: function depositERC20((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeFacade *BridgeFacadeTransactor) DepositERC20(opts *bind.TransactOpts, feeParams_ IBridgeFacadeDepositFeeERC20Parameters, depositParams_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "depositERC20", feeParams_, depositParams_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x73fb5a66.
//
// Solidity: function depositERC20((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeFacade *BridgeFacadeSession) DepositERC20(feeParams_ IBridgeFacadeDepositFeeERC20Parameters, depositParams_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.DepositERC20(&_BridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x73fb5a66.
//
// Solidity: function depositERC20((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) DepositERC20(feeParams_ IBridgeFacadeDepositFeeERC20Parameters, depositParams_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.DepositERC20(&_BridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x1087dcd1.
//
// Solidity: function depositERC721((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeFacade *BridgeFacadeTransactor) DepositERC721(opts *bind.TransactOpts, feeParams_ IBridgeFacadeDepositFeeERC721Parameters, depositParams_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "depositERC721", feeParams_, depositParams_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x1087dcd1.
//
// Solidity: function depositERC721((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeFacade *BridgeFacadeSession) DepositERC721(feeParams_ IBridgeFacadeDepositFeeERC721Parameters, depositParams_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.DepositERC721(&_BridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x1087dcd1.
//
// Solidity: function depositERC721((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) DepositERC721(feeParams_ IBridgeFacadeDepositFeeERC721Parameters, depositParams_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.DepositERC721(&_BridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xa7f40f2f.
//
// Solidity: function depositNative((address) feeParams_, (uint256,(bytes32,bytes),string,string) depositParams_) payable returns()
func (_BridgeFacade *BridgeFacadeTransactor) DepositNative(opts *bind.TransactOpts, feeParams_ IBridgeFacadeDepositFeeNativeParameters, depositParams_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "depositNative", feeParams_, depositParams_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xa7f40f2f.
//
// Solidity: function depositNative((address) feeParams_, (uint256,(bytes32,bytes),string,string) depositParams_) payable returns()
func (_BridgeFacade *BridgeFacadeSession) DepositNative(feeParams_ IBridgeFacadeDepositFeeNativeParameters, depositParams_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.DepositNative(&_BridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xa7f40f2f.
//
// Solidity: function depositNative((address) feeParams_, (uint256,(bytes32,bytes),string,string) depositParams_) payable returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) DepositNative(feeParams_ IBridgeFacadeDepositFeeNativeParameters, depositParams_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.DepositNative(&_BridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0xd93ff678.
//
// Solidity: function depositSBT((address) feeParams_, (address,uint256,(bytes32,bytes),string,string) depositParams_) payable returns()
func (_BridgeFacade *BridgeFacadeTransactor) DepositSBT(opts *bind.TransactOpts, feeParams_ IBridgeFacadeDepositFeeSBTParameters, depositParams_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "depositSBT", feeParams_, depositParams_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0xd93ff678.
//
// Solidity: function depositSBT((address) feeParams_, (address,uint256,(bytes32,bytes),string,string) depositParams_) payable returns()
func (_BridgeFacade *BridgeFacadeSession) DepositSBT(feeParams_ IBridgeFacadeDepositFeeSBTParameters, depositParams_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.DepositSBT(&_BridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0xd93ff678.
//
// Solidity: function depositSBT((address) feeParams_, (address,uint256,(bytes32,bytes),string,string) depositParams_) payable returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) DepositSBT(feeParams_ IBridgeFacadeDepositFeeSBTParameters, depositParams_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.DepositSBT(&_BridgeFacade.TransactOpts, feeParams_, depositParams_)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x4eb5943a.
//
// Solidity: function removeFeeToken((address[],uint256[],bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactor) RemoveFeeToken(opts *bind.TransactOpts, params_ IFeeManagerRemoveFeeTokenParameters) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "removeFeeToken", params_)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x4eb5943a.
//
// Solidity: function removeFeeToken((address[],uint256[],bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeSession) RemoveFeeToken(params_ IFeeManagerRemoveFeeTokenParameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.RemoveFeeToken(&_BridgeFacade.TransactOpts, params_)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x4eb5943a.
//
// Solidity: function removeFeeToken((address[],uint256[],bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) RemoveFeeToken(params_ IFeeManagerRemoveFeeTokenParameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.RemoveFeeToken(&_BridgeFacade.TransactOpts, params_)
}

// UpdateFeeToken is a paid mutator transaction binding the contract method 0x63925ea2.
//
// Solidity: function updateFeeToken((address[],uint256[],bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactor) UpdateFeeToken(opts *bind.TransactOpts, params_ IFeeManagerUpdateFeeTokenParameters) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "updateFeeToken", params_)
}

// UpdateFeeToken is a paid mutator transaction binding the contract method 0x63925ea2.
//
// Solidity: function updateFeeToken((address[],uint256[],bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeSession) UpdateFeeToken(params_ IFeeManagerUpdateFeeTokenParameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.UpdateFeeToken(&_BridgeFacade.TransactOpts, params_)
}

// UpdateFeeToken is a paid mutator transaction binding the contract method 0x63925ea2.
//
// Solidity: function updateFeeToken((address[],uint256[],bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) UpdateFeeToken(params_ IFeeManagerUpdateFeeTokenParameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.UpdateFeeToken(&_BridgeFacade.TransactOpts, params_)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BridgeFacade *BridgeFacadeTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BridgeFacade *BridgeFacadeSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _BridgeFacade.Contract.UpgradeTo(&_BridgeFacade.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _BridgeFacade.Contract.UpgradeTo(&_BridgeFacade.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeFacade *BridgeFacadeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeFacade *BridgeFacadeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeFacade.Contract.UpgradeToAndCall(&_BridgeFacade.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeFacade.Contract.UpgradeToAndCall(&_BridgeFacade.TransactOpts, newImplementation, data)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_BridgeFacade *BridgeFacadeTransactor) UpgradeToWithSig(opts *bind.TransactOpts, newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "upgradeToWithSig", newImplementation_, signature_)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_BridgeFacade *BridgeFacadeSession) UpgradeToWithSig(newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _BridgeFacade.Contract.UpgradeToWithSig(&_BridgeFacade.TransactOpts, newImplementation_, signature_)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) UpgradeToWithSig(newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _BridgeFacade.Contract.UpgradeToWithSig(&_BridgeFacade.TransactOpts, newImplementation_, signature_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactor) WithdrawERC1155(opts *bind.TransactOpts, params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "withdrawERC1155", params_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_BridgeFacade *BridgeFacadeSession) WithdrawERC1155(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.WithdrawERC1155(&_BridgeFacade.TransactOpts, params_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) WithdrawERC1155(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.WithdrawERC1155(&_BridgeFacade.TransactOpts, params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactor) WithdrawERC20(opts *bind.TransactOpts, params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "withdrawERC20", params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_BridgeFacade *BridgeFacadeSession) WithdrawERC20(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.WithdrawERC20(&_BridgeFacade.TransactOpts, params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) WithdrawERC20(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.WithdrawERC20(&_BridgeFacade.TransactOpts, params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactor) WithdrawERC721(opts *bind.TransactOpts, params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "withdrawERC721", params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_BridgeFacade *BridgeFacadeSession) WithdrawERC721(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.WithdrawERC721(&_BridgeFacade.TransactOpts, params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) WithdrawERC721(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.WithdrawERC721(&_BridgeFacade.TransactOpts, params_)
}

// WithdrawFeeToken is a paid mutator transaction binding the contract method 0x9f33cf22.
//
// Solidity: function withdrawFeeToken((address,address[],uint256[],bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactor) WithdrawFeeToken(opts *bind.TransactOpts, params_ IFeeManagerWithdrawFeeTokenParameters) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "withdrawFeeToken", params_)
}

// WithdrawFeeToken is a paid mutator transaction binding the contract method 0x9f33cf22.
//
// Solidity: function withdrawFeeToken((address,address[],uint256[],bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeSession) WithdrawFeeToken(params_ IFeeManagerWithdrawFeeTokenParameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.WithdrawFeeToken(&_BridgeFacade.TransactOpts, params_)
}

// WithdrawFeeToken is a paid mutator transaction binding the contract method 0x9f33cf22.
//
// Solidity: function withdrawFeeToken((address,address[],uint256[],bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) WithdrawFeeToken(params_ IFeeManagerWithdrawFeeTokenParameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.WithdrawFeeToken(&_BridgeFacade.TransactOpts, params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactor) WithdrawNative(opts *bind.TransactOpts, params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "withdrawNative", params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeSession) WithdrawNative(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.WithdrawNative(&_BridgeFacade.TransactOpts, params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) WithdrawNative(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.WithdrawNative(&_BridgeFacade.TransactOpts, params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactor) WithdrawSBT(opts *bind.TransactOpts, params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _BridgeFacade.contract.Transact(opts, "withdrawSBT", params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeSession) WithdrawSBT(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.WithdrawSBT(&_BridgeFacade.TransactOpts, params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_BridgeFacade *BridgeFacadeTransactorSession) WithdrawSBT(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _BridgeFacade.Contract.WithdrawSBT(&_BridgeFacade.TransactOpts, params_)
}

// BridgeFacadeAddedFeeTokenIterator is returned from FilterAddedFeeToken and is used to iterate over the raw logs and unpacked data for AddedFeeToken events raised by the BridgeFacade contract.
type BridgeFacadeAddedFeeTokenIterator struct {
	Event *BridgeFacadeAddedFeeToken // Event containing the contract specifics and raw log

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
func (it *BridgeFacadeAddedFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFacadeAddedFeeToken)
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
		it.Event = new(BridgeFacadeAddedFeeToken)
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
func (it *BridgeFacadeAddedFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFacadeAddedFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFacadeAddedFeeToken represents a AddedFeeToken event raised by the BridgeFacade contract.
type BridgeFacadeAddedFeeToken struct {
	FeeToken  common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddedFeeToken is a free log retrieval operation binding the contract event 0x5d966ab5c51224ee3c5146e53a0b57969d97d0e1779f52a1017813c13344efc4.
//
// Solidity: event AddedFeeToken(address feeToken, uint256 feeAmount)
func (_BridgeFacade *BridgeFacadeFilterer) FilterAddedFeeToken(opts *bind.FilterOpts) (*BridgeFacadeAddedFeeTokenIterator, error) {

	logs, sub, err := _BridgeFacade.contract.FilterLogs(opts, "AddedFeeToken")
	if err != nil {
		return nil, err
	}
	return &BridgeFacadeAddedFeeTokenIterator{contract: _BridgeFacade.contract, event: "AddedFeeToken", logs: logs, sub: sub}, nil
}

// WatchAddedFeeToken is a free log subscription operation binding the contract event 0x5d966ab5c51224ee3c5146e53a0b57969d97d0e1779f52a1017813c13344efc4.
//
// Solidity: event AddedFeeToken(address feeToken, uint256 feeAmount)
func (_BridgeFacade *BridgeFacadeFilterer) WatchAddedFeeToken(opts *bind.WatchOpts, sink chan<- *BridgeFacadeAddedFeeToken) (event.Subscription, error) {

	logs, sub, err := _BridgeFacade.contract.WatchLogs(opts, "AddedFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFacadeAddedFeeToken)
				if err := _BridgeFacade.contract.UnpackLog(event, "AddedFeeToken", log); err != nil {
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
func (_BridgeFacade *BridgeFacadeFilterer) ParseAddedFeeToken(log types.Log) (*BridgeFacadeAddedFeeToken, error) {
	event := new(BridgeFacadeAddedFeeToken)
	if err := _BridgeFacade.contract.UnpackLog(event, "AddedFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFacadeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the BridgeFacade contract.
type BridgeFacadeAdminChangedIterator struct {
	Event *BridgeFacadeAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeFacadeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFacadeAdminChanged)
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
		it.Event = new(BridgeFacadeAdminChanged)
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
func (it *BridgeFacadeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFacadeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFacadeAdminChanged represents a AdminChanged event raised by the BridgeFacade contract.
type BridgeFacadeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BridgeFacade *BridgeFacadeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*BridgeFacadeAdminChangedIterator, error) {

	logs, sub, err := _BridgeFacade.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &BridgeFacadeAdminChangedIterator{contract: _BridgeFacade.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BridgeFacade *BridgeFacadeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeFacadeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _BridgeFacade.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFacadeAdminChanged)
				if err := _BridgeFacade.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_BridgeFacade *BridgeFacadeFilterer) ParseAdminChanged(log types.Log) (*BridgeFacadeAdminChanged, error) {
	event := new(BridgeFacadeAdminChanged)
	if err := _BridgeFacade.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFacadeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the BridgeFacade contract.
type BridgeFacadeBeaconUpgradedIterator struct {
	Event *BridgeFacadeBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *BridgeFacadeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFacadeBeaconUpgraded)
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
		it.Event = new(BridgeFacadeBeaconUpgraded)
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
func (it *BridgeFacadeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFacadeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFacadeBeaconUpgraded represents a BeaconUpgraded event raised by the BridgeFacade contract.
type BridgeFacadeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BridgeFacade *BridgeFacadeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*BridgeFacadeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BridgeFacade.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFacadeBeaconUpgradedIterator{contract: _BridgeFacade.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BridgeFacade *BridgeFacadeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *BridgeFacadeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BridgeFacade.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFacadeBeaconUpgraded)
				if err := _BridgeFacade.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_BridgeFacade *BridgeFacadeFilterer) ParseBeaconUpgraded(log types.Log) (*BridgeFacadeBeaconUpgraded, error) {
	event := new(BridgeFacadeBeaconUpgraded)
	if err := _BridgeFacade.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFacadeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BridgeFacade contract.
type BridgeFacadeInitializedIterator struct {
	Event *BridgeFacadeInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeFacadeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFacadeInitialized)
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
		it.Event = new(BridgeFacadeInitialized)
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
func (it *BridgeFacadeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFacadeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFacadeInitialized represents a Initialized event raised by the BridgeFacade contract.
type BridgeFacadeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeFacade *BridgeFacadeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeFacadeInitializedIterator, error) {

	logs, sub, err := _BridgeFacade.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeFacadeInitializedIterator{contract: _BridgeFacade.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeFacade *BridgeFacadeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeFacadeInitialized) (event.Subscription, error) {

	logs, sub, err := _BridgeFacade.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFacadeInitialized)
				if err := _BridgeFacade.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BridgeFacade *BridgeFacadeFilterer) ParseInitialized(log types.Log) (*BridgeFacadeInitialized, error) {
	event := new(BridgeFacadeInitialized)
	if err := _BridgeFacade.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFacadeRemovedFeeTokenIterator is returned from FilterRemovedFeeToken and is used to iterate over the raw logs and unpacked data for RemovedFeeToken events raised by the BridgeFacade contract.
type BridgeFacadeRemovedFeeTokenIterator struct {
	Event *BridgeFacadeRemovedFeeToken // Event containing the contract specifics and raw log

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
func (it *BridgeFacadeRemovedFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFacadeRemovedFeeToken)
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
		it.Event = new(BridgeFacadeRemovedFeeToken)
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
func (it *BridgeFacadeRemovedFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFacadeRemovedFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFacadeRemovedFeeToken represents a RemovedFeeToken event raised by the BridgeFacade contract.
type BridgeFacadeRemovedFeeToken struct {
	FeeToken  common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemovedFeeToken is a free log retrieval operation binding the contract event 0xeac220a493a0693e32f1240aae311de65b2d84391a9ba8310fcaf359d7c9bd8f.
//
// Solidity: event RemovedFeeToken(address feeToken, uint256 feeAmount)
func (_BridgeFacade *BridgeFacadeFilterer) FilterRemovedFeeToken(opts *bind.FilterOpts) (*BridgeFacadeRemovedFeeTokenIterator, error) {

	logs, sub, err := _BridgeFacade.contract.FilterLogs(opts, "RemovedFeeToken")
	if err != nil {
		return nil, err
	}
	return &BridgeFacadeRemovedFeeTokenIterator{contract: _BridgeFacade.contract, event: "RemovedFeeToken", logs: logs, sub: sub}, nil
}

// WatchRemovedFeeToken is a free log subscription operation binding the contract event 0xeac220a493a0693e32f1240aae311de65b2d84391a9ba8310fcaf359d7c9bd8f.
//
// Solidity: event RemovedFeeToken(address feeToken, uint256 feeAmount)
func (_BridgeFacade *BridgeFacadeFilterer) WatchRemovedFeeToken(opts *bind.WatchOpts, sink chan<- *BridgeFacadeRemovedFeeToken) (event.Subscription, error) {

	logs, sub, err := _BridgeFacade.contract.WatchLogs(opts, "RemovedFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFacadeRemovedFeeToken)
				if err := _BridgeFacade.contract.UnpackLog(event, "RemovedFeeToken", log); err != nil {
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
func (_BridgeFacade *BridgeFacadeFilterer) ParseRemovedFeeToken(log types.Log) (*BridgeFacadeRemovedFeeToken, error) {
	event := new(BridgeFacadeRemovedFeeToken)
	if err := _BridgeFacade.contract.UnpackLog(event, "RemovedFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFacadeUpdatedFeeTokenIterator is returned from FilterUpdatedFeeToken and is used to iterate over the raw logs and unpacked data for UpdatedFeeToken events raised by the BridgeFacade contract.
type BridgeFacadeUpdatedFeeTokenIterator struct {
	Event *BridgeFacadeUpdatedFeeToken // Event containing the contract specifics and raw log

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
func (it *BridgeFacadeUpdatedFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFacadeUpdatedFeeToken)
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
		it.Event = new(BridgeFacadeUpdatedFeeToken)
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
func (it *BridgeFacadeUpdatedFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFacadeUpdatedFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFacadeUpdatedFeeToken represents a UpdatedFeeToken event raised by the BridgeFacade contract.
type BridgeFacadeUpdatedFeeToken struct {
	FeeToken  common.Address
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedFeeToken is a free log retrieval operation binding the contract event 0x70698787bb7456af2d4339ba1e7d059c4eda7c473412e7ed2c0d5b361a49d50a.
//
// Solidity: event UpdatedFeeToken(address feeToken, uint256 feeAmount)
func (_BridgeFacade *BridgeFacadeFilterer) FilterUpdatedFeeToken(opts *bind.FilterOpts) (*BridgeFacadeUpdatedFeeTokenIterator, error) {

	logs, sub, err := _BridgeFacade.contract.FilterLogs(opts, "UpdatedFeeToken")
	if err != nil {
		return nil, err
	}
	return &BridgeFacadeUpdatedFeeTokenIterator{contract: _BridgeFacade.contract, event: "UpdatedFeeToken", logs: logs, sub: sub}, nil
}

// WatchUpdatedFeeToken is a free log subscription operation binding the contract event 0x70698787bb7456af2d4339ba1e7d059c4eda7c473412e7ed2c0d5b361a49d50a.
//
// Solidity: event UpdatedFeeToken(address feeToken, uint256 feeAmount)
func (_BridgeFacade *BridgeFacadeFilterer) WatchUpdatedFeeToken(opts *bind.WatchOpts, sink chan<- *BridgeFacadeUpdatedFeeToken) (event.Subscription, error) {

	logs, sub, err := _BridgeFacade.contract.WatchLogs(opts, "UpdatedFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFacadeUpdatedFeeToken)
				if err := _BridgeFacade.contract.UnpackLog(event, "UpdatedFeeToken", log); err != nil {
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
func (_BridgeFacade *BridgeFacadeFilterer) ParseUpdatedFeeToken(log types.Log) (*BridgeFacadeUpdatedFeeToken, error) {
	event := new(BridgeFacadeUpdatedFeeToken)
	if err := _BridgeFacade.contract.UnpackLog(event, "UpdatedFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFacadeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BridgeFacade contract.
type BridgeFacadeUpgradedIterator struct {
	Event *BridgeFacadeUpgraded // Event containing the contract specifics and raw log

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
func (it *BridgeFacadeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFacadeUpgraded)
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
		it.Event = new(BridgeFacadeUpgraded)
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
func (it *BridgeFacadeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFacadeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFacadeUpgraded represents a Upgraded event raised by the BridgeFacade contract.
type BridgeFacadeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BridgeFacade *BridgeFacadeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BridgeFacadeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BridgeFacade.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFacadeUpgradedIterator{contract: _BridgeFacade.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BridgeFacade *BridgeFacadeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BridgeFacadeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BridgeFacade.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFacadeUpgraded)
				if err := _BridgeFacade.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BridgeFacade *BridgeFacadeFilterer) ParseUpgraded(log types.Log) (*BridgeFacadeUpgraded, error) {
	event := new(BridgeFacadeUpgraded)
	if err := _BridgeFacade.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFacadeWithdrawnFeeTokenIterator is returned from FilterWithdrawnFeeToken and is used to iterate over the raw logs and unpacked data for WithdrawnFeeToken events raised by the BridgeFacade contract.
type BridgeFacadeWithdrawnFeeTokenIterator struct {
	Event *BridgeFacadeWithdrawnFeeToken // Event containing the contract specifics and raw log

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
func (it *BridgeFacadeWithdrawnFeeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFacadeWithdrawnFeeToken)
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
		it.Event = new(BridgeFacadeWithdrawnFeeToken)
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
func (it *BridgeFacadeWithdrawnFeeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFacadeWithdrawnFeeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFacadeWithdrawnFeeToken represents a WithdrawnFeeToken event raised by the BridgeFacade contract.
type BridgeFacadeWithdrawnFeeToken struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnFeeToken is a free log retrieval operation binding the contract event 0x0221f5dbeb176269bc9dbce8b10193c570930431bbe525ae79b4599910500bf4.
//
// Solidity: event WithdrawnFeeToken(address receiver, address feeToken, uint256 amount)
func (_BridgeFacade *BridgeFacadeFilterer) FilterWithdrawnFeeToken(opts *bind.FilterOpts) (*BridgeFacadeWithdrawnFeeTokenIterator, error) {

	logs, sub, err := _BridgeFacade.contract.FilterLogs(opts, "WithdrawnFeeToken")
	if err != nil {
		return nil, err
	}
	return &BridgeFacadeWithdrawnFeeTokenIterator{contract: _BridgeFacade.contract, event: "WithdrawnFeeToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawnFeeToken is a free log subscription operation binding the contract event 0x0221f5dbeb176269bc9dbce8b10193c570930431bbe525ae79b4599910500bf4.
//
// Solidity: event WithdrawnFeeToken(address receiver, address feeToken, uint256 amount)
func (_BridgeFacade *BridgeFacadeFilterer) WatchWithdrawnFeeToken(opts *bind.WatchOpts, sink chan<- *BridgeFacadeWithdrawnFeeToken) (event.Subscription, error) {

	logs, sub, err := _BridgeFacade.contract.WatchLogs(opts, "WithdrawnFeeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFacadeWithdrawnFeeToken)
				if err := _BridgeFacade.contract.UnpackLog(event, "WithdrawnFeeToken", log); err != nil {
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
func (_BridgeFacade *BridgeFacadeFilterer) ParseWithdrawnFeeToken(log types.Log) (*BridgeFacadeWithdrawnFeeToken, error) {
	event := new(BridgeFacadeWithdrawnFeeToken)
	if err := _BridgeFacade.contract.UnpackLog(event, "WithdrawnFeeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
