// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"DepositedERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"DepositedERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"DepositedERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"name\":\"DepositedNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"name\":\"DepositedSBT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"WithdrawnERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"WithdrawnERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"WithdrawnERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"WithdrawnNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"WithdrawnSBT\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"P\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bundleImplementation_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"facade_\",\"type\":\"address\"}],\"name\":\"__Bridge_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bundleExecutorImplementation_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"facade_\",\"type\":\"address\"}],\"name\":\"__Bundler_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"}],\"name\":\"__Signers_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bundleExecutorImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"changeBundleExecutorImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFacade_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"changeFacade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"newSignerPubKey_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"changeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"checkSignatureAndIncrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.DepositERC1155Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.DepositERC20Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.DepositERC721Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"internalType\":\"structINativeHandler.DepositNativeParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"internalType\":\"structISBTHandler.DepositSBTParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositSBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt_\",\"type\":\"bytes32\"}],\"name\":\"determineProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"getSigComponents\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"upgradeToWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"validateChangeAddressSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"tokenDataLeaf_\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle_\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"}],\"name\":\"verifyMerkleLeaf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.WithdrawERC1155Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.WithdrawERC1155Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC1155Bundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.WithdrawERC20Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.WithdrawERC20Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC20Bundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.WithdrawERC721Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.WithdrawERC721Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC721Bundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structINativeHandler.WithdrawNativeParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structINativeHandler.WithdrawNativeParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawNativeBundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structISBTHandler.WithdrawSBTParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawSBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structISBTHandler.WithdrawSBTParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawSBTBundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_Bridge *BridgeCaller) P(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "P")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_Bridge *BridgeSession) P() (*big.Int, error) {
	return _Bridge.Contract.P(&_Bridge.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_Bridge *BridgeCallerSession) P() (*big.Int, error) {
	return _Bridge.Contract.P(&_Bridge.CallOpts)
}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_Bridge *BridgeCaller) BundleExecutorImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "bundleExecutorImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_Bridge *BridgeSession) BundleExecutorImplementation() (common.Address, error) {
	return _Bridge.Contract.BundleExecutorImplementation(&_Bridge.CallOpts)
}

// BundleExecutorImplementation is a free data retrieval call binding the contract method 0x59e46336.
//
// Solidity: function bundleExecutorImplementation() view returns(address)
func (_Bridge *BridgeCallerSession) BundleExecutorImplementation() (common.Address, error) {
	return _Bridge.Contract.BundleExecutorImplementation(&_Bridge.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_Bridge *BridgeCaller) ChainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "chainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_Bridge *BridgeSession) ChainName() (string, error) {
	return _Bridge.Contract.ChainName(&_Bridge.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_Bridge *BridgeCallerSession) ChainName() (string, error) {
	return _Bridge.Contract.ChainName(&_Bridge.CallOpts)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_Bridge *BridgeCaller) DetermineProxyAddress(opts *bind.CallOpts, salt_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "determineProxyAddress", salt_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_Bridge *BridgeSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _Bridge.Contract.DetermineProxyAddress(&_Bridge.CallOpts, salt_)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_Bridge *BridgeCallerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _Bridge.Contract.DetermineProxyAddress(&_Bridge.CallOpts, salt_)
}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_Bridge *BridgeCaller) Facade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "facade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_Bridge *BridgeSession) Facade() (common.Address, error) {
	return _Bridge.Contract.Facade(&_Bridge.CallOpts)
}

// Facade is a free data retrieval call binding the contract method 0x5014a0fb.
//
// Solidity: function facade() view returns(address)
func (_Bridge *BridgeCallerSession) Facade() (common.Address, error) {
	return _Bridge.Contract.Facade(&_Bridge.CallOpts)
}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_Bridge *BridgeCaller) GetSigComponents(opts *bind.CallOpts, methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getSigComponents", methodId_, contractAddress_)

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
func (_Bridge *BridgeSession) GetSigComponents(methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	return _Bridge.Contract.GetSigComponents(&_Bridge.CallOpts, methodId_, contractAddress_)
}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_Bridge *BridgeCallerSession) GetSigComponents(methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	return _Bridge.Contract.GetSigComponents(&_Bridge.CallOpts, methodId_, contractAddress_)
}

// Nonces is a free data retrieval call binding the contract method 0xed3218a2.
//
// Solidity: function nonces(address , uint8 ) view returns(uint256)
func (_Bridge *BridgeCaller) Nonces(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "nonces", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0xed3218a2.
//
// Solidity: function nonces(address , uint8 ) view returns(uint256)
func (_Bridge *BridgeSession) Nonces(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _Bridge.Contract.Nonces(&_Bridge.CallOpts, arg0, arg1)
}

// Nonces is a free data retrieval call binding the contract method 0xed3218a2.
//
// Solidity: function nonces(address , uint8 ) view returns(uint256)
func (_Bridge *BridgeCallerSession) Nonces(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _Bridge.Contract.Nonces(&_Bridge.CallOpts, arg0, arg1)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bridge *BridgeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bridge *BridgeSession) ProxiableUUID() ([32]byte, error) {
	return _Bridge.Contract.ProxiableUUID(&_Bridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bridge *BridgeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Bridge.Contract.ProxiableUUID(&_Bridge.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Bridge *BridgeCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Bridge *BridgeSession) Signer() (common.Address, error) {
	return _Bridge.Contract.Signer(&_Bridge.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Bridge *BridgeCallerSession) Signer() (common.Address, error) {
	return _Bridge.Contract.Signer(&_Bridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 ) view returns(bool)
func (_Bridge *BridgeCaller) UsedHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "usedHashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 ) view returns(bool)
func (_Bridge *BridgeSession) UsedHashes(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.UsedHashes(&_Bridge.CallOpts, arg0)
}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) UsedHashes(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.UsedHashes(&_Bridge.CallOpts, arg0)
}

// BridgeInit is a paid mutator transaction binding the contract method 0x1f23d35d.
//
// Solidity: function __Bridge_init(address signer_, address bundleImplementation_, string chainName_, address facade_) returns()
func (_Bridge *BridgeTransactor) BridgeInit(opts *bind.TransactOpts, signer_ common.Address, bundleImplementation_ common.Address, chainName_ string, facade_ common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "__Bridge_init", signer_, bundleImplementation_, chainName_, facade_)
}

// BridgeInit is a paid mutator transaction binding the contract method 0x1f23d35d.
//
// Solidity: function __Bridge_init(address signer_, address bundleImplementation_, string chainName_, address facade_) returns()
func (_Bridge *BridgeSession) BridgeInit(signer_ common.Address, bundleImplementation_ common.Address, chainName_ string, facade_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeInit(&_Bridge.TransactOpts, signer_, bundleImplementation_, chainName_, facade_)
}

// BridgeInit is a paid mutator transaction binding the contract method 0x1f23d35d.
//
// Solidity: function __Bridge_init(address signer_, address bundleImplementation_, string chainName_, address facade_) returns()
func (_Bridge *BridgeTransactorSession) BridgeInit(signer_ common.Address, bundleImplementation_ common.Address, chainName_ string, facade_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeInit(&_Bridge.TransactOpts, signer_, bundleImplementation_, chainName_, facade_)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x96de44c2.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_, address facade_) returns()
func (_Bridge *BridgeTransactor) BundlerInit(opts *bind.TransactOpts, bundleExecutorImplementation_ common.Address, facade_ common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "__Bundler_init", bundleExecutorImplementation_, facade_)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x96de44c2.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_, address facade_) returns()
func (_Bridge *BridgeSession) BundlerInit(bundleExecutorImplementation_ common.Address, facade_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.BundlerInit(&_Bridge.TransactOpts, bundleExecutorImplementation_, facade_)
}

// BundlerInit is a paid mutator transaction binding the contract method 0x96de44c2.
//
// Solidity: function __Bundler_init(address bundleExecutorImplementation_, address facade_) returns()
func (_Bridge *BridgeTransactorSession) BundlerInit(bundleExecutorImplementation_ common.Address, facade_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.BundlerInit(&_Bridge.TransactOpts, bundleExecutorImplementation_, facade_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x509ace95.
//
// Solidity: function __Signers_init(address signer_, string chainName_) returns()
func (_Bridge *BridgeTransactor) SignersInit(opts *bind.TransactOpts, signer_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "__Signers_init", signer_, chainName_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x509ace95.
//
// Solidity: function __Signers_init(address signer_, string chainName_) returns()
func (_Bridge *BridgeSession) SignersInit(signer_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _Bridge.Contract.SignersInit(&_Bridge.TransactOpts, signer_, chainName_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x509ace95.
//
// Solidity: function __Signers_init(address signer_, string chainName_) returns()
func (_Bridge *BridgeTransactorSession) SignersInit(signer_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _Bridge.Contract.SignersInit(&_Bridge.TransactOpts, signer_, chainName_)
}

// ChangeBundleExecutorImplementation is a paid mutator transaction binding the contract method 0xb297c523.
//
// Solidity: function changeBundleExecutorImplementation(address newImplementation_, bytes signature_) returns()
func (_Bridge *BridgeTransactor) ChangeBundleExecutorImplementation(opts *bind.TransactOpts, newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "changeBundleExecutorImplementation", newImplementation_, signature_)
}

// ChangeBundleExecutorImplementation is a paid mutator transaction binding the contract method 0xb297c523.
//
// Solidity: function changeBundleExecutorImplementation(address newImplementation_, bytes signature_) returns()
func (_Bridge *BridgeSession) ChangeBundleExecutorImplementation(newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeBundleExecutorImplementation(&_Bridge.TransactOpts, newImplementation_, signature_)
}

// ChangeBundleExecutorImplementation is a paid mutator transaction binding the contract method 0xb297c523.
//
// Solidity: function changeBundleExecutorImplementation(address newImplementation_, bytes signature_) returns()
func (_Bridge *BridgeTransactorSession) ChangeBundleExecutorImplementation(newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeBundleExecutorImplementation(&_Bridge.TransactOpts, newImplementation_, signature_)
}

// ChangeFacade is a paid mutator transaction binding the contract method 0xa8e108ff.
//
// Solidity: function changeFacade(address newFacade_, bytes signature_) returns()
func (_Bridge *BridgeTransactor) ChangeFacade(opts *bind.TransactOpts, newFacade_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "changeFacade", newFacade_, signature_)
}

// ChangeFacade is a paid mutator transaction binding the contract method 0xa8e108ff.
//
// Solidity: function changeFacade(address newFacade_, bytes signature_) returns()
func (_Bridge *BridgeSession) ChangeFacade(newFacade_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeFacade(&_Bridge.TransactOpts, newFacade_, signature_)
}

// ChangeFacade is a paid mutator transaction binding the contract method 0xa8e108ff.
//
// Solidity: function changeFacade(address newFacade_, bytes signature_) returns()
func (_Bridge *BridgeTransactorSession) ChangeFacade(newFacade_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeFacade(&_Bridge.TransactOpts, newFacade_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_Bridge *BridgeTransactor) ChangeSigner(opts *bind.TransactOpts, newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "changeSigner", newSignerPubKey_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_Bridge *BridgeSession) ChangeSigner(newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeSigner(&_Bridge.TransactOpts, newSignerPubKey_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_Bridge *BridgeTransactorSession) ChangeSigner(newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeSigner(&_Bridge.TransactOpts, newSignerPubKey_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_Bridge *BridgeTransactor) CheckSignatureAndIncrementNonce(opts *bind.TransactOpts, methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "checkSignatureAndIncrementNonce", methodId_, contractAddress_, signHash_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_Bridge *BridgeSession) CheckSignatureAndIncrementNonce(methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.Contract.CheckSignatureAndIncrementNonce(&_Bridge.TransactOpts, methodId_, contractAddress_, signHash_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_Bridge *BridgeTransactorSession) CheckSignatureAndIncrementNonce(methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.Contract.CheckSignatureAndIncrementNonce(&_Bridge.TransactOpts, methodId_, contractAddress_, signHash_, signature_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x969aef06.
//
// Solidity: function depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_Bridge *BridgeTransactor) DepositERC1155(opts *bind.TransactOpts, params_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositERC1155", params_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x969aef06.
//
// Solidity: function depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_Bridge *BridgeSession) DepositERC1155(params_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC1155(&_Bridge.TransactOpts, params_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x969aef06.
//
// Solidity: function depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_Bridge *BridgeTransactorSession) DepositERC1155(params_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC1155(&_Bridge.TransactOpts, params_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xbebd5a0b.
//
// Solidity: function depositERC20((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_Bridge *BridgeTransactor) DepositERC20(opts *bind.TransactOpts, params_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositERC20", params_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xbebd5a0b.
//
// Solidity: function depositERC20((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_Bridge *BridgeSession) DepositERC20(params_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC20(&_Bridge.TransactOpts, params_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xbebd5a0b.
//
// Solidity: function depositERC20((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_Bridge *BridgeTransactorSession) DepositERC20(params_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC20(&_Bridge.TransactOpts, params_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x6a38abbf.
//
// Solidity: function depositERC721((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_Bridge *BridgeTransactor) DepositERC721(opts *bind.TransactOpts, params_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositERC721", params_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x6a38abbf.
//
// Solidity: function depositERC721((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_Bridge *BridgeSession) DepositERC721(params_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC721(&_Bridge.TransactOpts, params_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x6a38abbf.
//
// Solidity: function depositERC721((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_Bridge *BridgeTransactorSession) DepositERC721(params_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC721(&_Bridge.TransactOpts, params_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xb27588e5.
//
// Solidity: function depositNative((uint256,(bytes32,bytes),string,string) params_) payable returns()
func (_Bridge *BridgeTransactor) DepositNative(opts *bind.TransactOpts, params_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositNative", params_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xb27588e5.
//
// Solidity: function depositNative((uint256,(bytes32,bytes),string,string) params_) payable returns()
func (_Bridge *BridgeSession) DepositNative(params_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _Bridge.Contract.DepositNative(&_Bridge.TransactOpts, params_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xb27588e5.
//
// Solidity: function depositNative((uint256,(bytes32,bytes),string,string) params_) payable returns()
func (_Bridge *BridgeTransactorSession) DepositNative(params_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _Bridge.Contract.DepositNative(&_Bridge.TransactOpts, params_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0x755f3823.
//
// Solidity: function depositSBT((address,uint256,(bytes32,bytes),string,string) params_) returns()
func (_Bridge *BridgeTransactor) DepositSBT(opts *bind.TransactOpts, params_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositSBT", params_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0x755f3823.
//
// Solidity: function depositSBT((address,uint256,(bytes32,bytes),string,string) params_) returns()
func (_Bridge *BridgeSession) DepositSBT(params_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _Bridge.Contract.DepositSBT(&_Bridge.TransactOpts, params_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0x755f3823.
//
// Solidity: function depositSBT((address,uint256,(bytes32,bytes),string,string) params_) returns()
func (_Bridge *BridgeTransactorSession) DepositSBT(params_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _Bridge.Contract.DepositSBT(&_Bridge.TransactOpts, params_)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Bridge *BridgeSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC1155BatchReceived(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC1155BatchReceived(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC1155Received(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC1155Received(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC721Received(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC721Received(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Bridge *BridgeTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Bridge *BridgeSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.UpgradeTo(&_Bridge.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Bridge *BridgeTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.UpgradeTo(&_Bridge.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bridge *BridgeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bridge *BridgeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.UpgradeToAndCall(&_Bridge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bridge *BridgeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.UpgradeToAndCall(&_Bridge.TransactOpts, newImplementation, data)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_Bridge *BridgeTransactor) UpgradeToWithSig(opts *bind.TransactOpts, newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "upgradeToWithSig", newImplementation_, signature_)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_Bridge *BridgeSession) UpgradeToWithSig(newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.Contract.UpgradeToWithSig(&_Bridge.TransactOpts, newImplementation_, signature_)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_Bridge *BridgeTransactorSession) UpgradeToWithSig(newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.Contract.UpgradeToWithSig(&_Bridge.TransactOpts, newImplementation_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_Bridge *BridgeTransactor) ValidateChangeAddressSignature(opts *bind.TransactOpts, methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "validateChangeAddressSignature", methodId_, contractAddress_, newAddress_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_Bridge *BridgeSession) ValidateChangeAddressSignature(methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ValidateChangeAddressSignature(&_Bridge.TransactOpts, methodId_, contractAddress_, newAddress_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_Bridge *BridgeTransactorSession) ValidateChangeAddressSignature(methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ValidateChangeAddressSignature(&_Bridge.TransactOpts, methodId_, contractAddress_, newAddress_, signature_)
}

// VerifyMerkleLeaf is a paid mutator transaction binding the contract method 0x02fd4e26.
//
// Solidity: function verifyMerkleLeaf(bytes tokenDataLeaf_, (bytes32,bytes) bundle_, bytes32 originHash_, address receiver_, bytes proof_) returns()
func (_Bridge *BridgeTransactor) VerifyMerkleLeaf(opts *bind.TransactOpts, tokenDataLeaf_ []byte, bundle_ IBundlerBundle, originHash_ [32]byte, receiver_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "verifyMerkleLeaf", tokenDataLeaf_, bundle_, originHash_, receiver_, proof_)
}

// VerifyMerkleLeaf is a paid mutator transaction binding the contract method 0x02fd4e26.
//
// Solidity: function verifyMerkleLeaf(bytes tokenDataLeaf_, (bytes32,bytes) bundle_, bytes32 originHash_, address receiver_, bytes proof_) returns()
func (_Bridge *BridgeSession) VerifyMerkleLeaf(tokenDataLeaf_ []byte, bundle_ IBundlerBundle, originHash_ [32]byte, receiver_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _Bridge.Contract.VerifyMerkleLeaf(&_Bridge.TransactOpts, tokenDataLeaf_, bundle_, originHash_, receiver_, proof_)
}

// VerifyMerkleLeaf is a paid mutator transaction binding the contract method 0x02fd4e26.
//
// Solidity: function verifyMerkleLeaf(bytes tokenDataLeaf_, (bytes32,bytes) bundle_, bytes32 originHash_, address receiver_, bytes proof_) returns()
func (_Bridge *BridgeTransactorSession) VerifyMerkleLeaf(tokenDataLeaf_ []byte, bundle_ IBundlerBundle, originHash_ [32]byte, receiver_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _Bridge.Contract.VerifyMerkleLeaf(&_Bridge.TransactOpts, tokenDataLeaf_, bundle_, originHash_, receiver_, proof_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeTransactor) WithdrawERC1155(opts *bind.TransactOpts, params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC1155", params_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeSession) WithdrawERC1155(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC1155(&_Bridge.TransactOpts, params_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC1155(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC1155(&_Bridge.TransactOpts, params_)
}

// WithdrawERC1155Bundle is a paid mutator transaction binding the contract method 0xb76da974.
//
// Solidity: function withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeTransactor) WithdrawERC1155Bundle(opts *bind.TransactOpts, params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC1155Bundle", params_)
}

// WithdrawERC1155Bundle is a paid mutator transaction binding the contract method 0xb76da974.
//
// Solidity: function withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeSession) WithdrawERC1155Bundle(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC1155Bundle(&_Bridge.TransactOpts, params_)
}

// WithdrawERC1155Bundle is a paid mutator transaction binding the contract method 0xb76da974.
//
// Solidity: function withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC1155Bundle(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC1155Bundle(&_Bridge.TransactOpts, params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeTransactor) WithdrawERC20(opts *bind.TransactOpts, params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC20", params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeSession) WithdrawERC20(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC20(&_Bridge.TransactOpts, params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC20(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC20(&_Bridge.TransactOpts, params_)
}

// WithdrawERC20Bundle is a paid mutator transaction binding the contract method 0xe53507bd.
//
// Solidity: function withdrawERC20Bundle((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeTransactor) WithdrawERC20Bundle(opts *bind.TransactOpts, params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC20Bundle", params_)
}

// WithdrawERC20Bundle is a paid mutator transaction binding the contract method 0xe53507bd.
//
// Solidity: function withdrawERC20Bundle((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeSession) WithdrawERC20Bundle(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC20Bundle(&_Bridge.TransactOpts, params_)
}

// WithdrawERC20Bundle is a paid mutator transaction binding the contract method 0xe53507bd.
//
// Solidity: function withdrawERC20Bundle((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC20Bundle(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC20Bundle(&_Bridge.TransactOpts, params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeTransactor) WithdrawERC721(opts *bind.TransactOpts, params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC721", params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeSession) WithdrawERC721(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC721(&_Bridge.TransactOpts, params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC721(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC721(&_Bridge.TransactOpts, params_)
}

// WithdrawERC721Bundle is a paid mutator transaction binding the contract method 0x9ae00a57.
//
// Solidity: function withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeTransactor) WithdrawERC721Bundle(opts *bind.TransactOpts, params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC721Bundle", params_)
}

// WithdrawERC721Bundle is a paid mutator transaction binding the contract method 0x9ae00a57.
//
// Solidity: function withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeSession) WithdrawERC721Bundle(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC721Bundle(&_Bridge.TransactOpts, params_)
}

// WithdrawERC721Bundle is a paid mutator transaction binding the contract method 0x9ae00a57.
//
// Solidity: function withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC721Bundle(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC721Bundle(&_Bridge.TransactOpts, params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_Bridge *BridgeTransactor) WithdrawNative(opts *bind.TransactOpts, params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawNative", params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_Bridge *BridgeSession) WithdrawNative(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawNative(&_Bridge.TransactOpts, params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawNative(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawNative(&_Bridge.TransactOpts, params_)
}

// WithdrawNativeBundle is a paid mutator transaction binding the contract method 0xc4f54e21.
//
// Solidity: function withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_Bridge *BridgeTransactor) WithdrawNativeBundle(opts *bind.TransactOpts, params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawNativeBundle", params_)
}

// WithdrawNativeBundle is a paid mutator transaction binding the contract method 0xc4f54e21.
//
// Solidity: function withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_Bridge *BridgeSession) WithdrawNativeBundle(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawNativeBundle(&_Bridge.TransactOpts, params_)
}

// WithdrawNativeBundle is a paid mutator transaction binding the contract method 0xc4f54e21.
//
// Solidity: function withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawNativeBundle(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawNativeBundle(&_Bridge.TransactOpts, params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_Bridge *BridgeTransactor) WithdrawSBT(opts *bind.TransactOpts, params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawSBT", params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_Bridge *BridgeSession) WithdrawSBT(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawSBT(&_Bridge.TransactOpts, params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawSBT(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawSBT(&_Bridge.TransactOpts, params_)
}

// WithdrawSBTBundle is a paid mutator transaction binding the contract method 0x9a830d4c.
//
// Solidity: function withdrawSBTBundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_Bridge *BridgeTransactor) WithdrawSBTBundle(opts *bind.TransactOpts, params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawSBTBundle", params_)
}

// WithdrawSBTBundle is a paid mutator transaction binding the contract method 0x9a830d4c.
//
// Solidity: function withdrawSBTBundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_Bridge *BridgeSession) WithdrawSBTBundle(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawSBTBundle(&_Bridge.TransactOpts, params_)
}

// WithdrawSBTBundle is a paid mutator transaction binding the contract method 0x9a830d4c.
//
// Solidity: function withdrawSBTBundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawSBTBundle(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawSBTBundle(&_Bridge.TransactOpts, params_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// BridgeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Bridge contract.
type BridgeAdminChangedIterator struct {
	Event *BridgeAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeAdminChanged)
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
		it.Event = new(BridgeAdminChanged)
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
func (it *BridgeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeAdminChanged represents a AdminChanged event raised by the Bridge contract.
type BridgeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Bridge *BridgeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*BridgeAdminChangedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &BridgeAdminChangedIterator{contract: _Bridge.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Bridge *BridgeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeAdminChanged)
				if err := _Bridge.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseAdminChanged(log types.Log) (*BridgeAdminChanged, error) {
	event := new(BridgeAdminChanged)
	if err := _Bridge.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Bridge contract.
type BridgeBeaconUpgradedIterator struct {
	Event *BridgeBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *BridgeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBeaconUpgraded)
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
		it.Event = new(BridgeBeaconUpgraded)
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
func (it *BridgeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBeaconUpgraded represents a BeaconUpgraded event raised by the Bridge contract.
type BridgeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Bridge *BridgeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*BridgeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBeaconUpgradedIterator{contract: _Bridge.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Bridge *BridgeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *BridgeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBeaconUpgraded)
				if err := _Bridge.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseBeaconUpgraded(log types.Log) (*BridgeBeaconUpgraded, error) {
	event := new(BridgeBeaconUpgraded)
	if err := _Bridge.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositedERC1155Iterator is returned from FilterDepositedERC1155 and is used to iterate over the raw logs and unpacked data for DepositedERC1155 events raised by the Bridge contract.
type BridgeDepositedERC1155Iterator struct {
	Event *BridgeDepositedERC1155 // Event containing the contract specifics and raw log

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
func (it *BridgeDepositedERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositedERC1155)
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
		it.Event = new(BridgeDepositedERC1155)
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
func (it *BridgeDepositedERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositedERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositedERC1155 represents a DepositedERC1155 event raised by the Bridge contract.
type BridgeDepositedERC1155 struct {
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
func (_Bridge *BridgeFilterer) FilterDepositedERC1155(opts *bind.FilterOpts) (*BridgeDepositedERC1155Iterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositedERC1155")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedERC1155Iterator{contract: _Bridge.contract, event: "DepositedERC1155", logs: logs, sub: sub}, nil
}

// WatchDepositedERC1155 is a free log subscription operation binding the contract event 0x103b790f2fa3a8676ff87c3620a55f0853d0e45128a8c7e9fadf29e17c51d07a.
//
// Solidity: event DepositedERC1155(address token, uint256 tokenId, uint256 amount, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_Bridge *BridgeFilterer) WatchDepositedERC1155(opts *bind.WatchOpts, sink chan<- *BridgeDepositedERC1155) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositedERC1155")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositedERC1155)
				if err := _Bridge.contract.UnpackLog(event, "DepositedERC1155", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseDepositedERC1155(log types.Log) (*BridgeDepositedERC1155, error) {
	event := new(BridgeDepositedERC1155)
	if err := _Bridge.contract.UnpackLog(event, "DepositedERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositedERC20Iterator is returned from FilterDepositedERC20 and is used to iterate over the raw logs and unpacked data for DepositedERC20 events raised by the Bridge contract.
type BridgeDepositedERC20Iterator struct {
	Event *BridgeDepositedERC20 // Event containing the contract specifics and raw log

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
func (it *BridgeDepositedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositedERC20)
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
		it.Event = new(BridgeDepositedERC20)
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
func (it *BridgeDepositedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositedERC20 represents a DepositedERC20 event raised by the Bridge contract.
type BridgeDepositedERC20 struct {
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
func (_Bridge *BridgeFilterer) FilterDepositedERC20(opts *bind.FilterOpts) (*BridgeDepositedERC20Iterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositedERC20")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedERC20Iterator{contract: _Bridge.contract, event: "DepositedERC20", logs: logs, sub: sub}, nil
}

// WatchDepositedERC20 is a free log subscription operation binding the contract event 0x043d52f9acdd847f0210803c386559db9e09d492143f2072fe30ea62ff0bb639.
//
// Solidity: event DepositedERC20(address token, uint256 amount, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_Bridge *BridgeFilterer) WatchDepositedERC20(opts *bind.WatchOpts, sink chan<- *BridgeDepositedERC20) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositedERC20)
				if err := _Bridge.contract.UnpackLog(event, "DepositedERC20", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseDepositedERC20(log types.Log) (*BridgeDepositedERC20, error) {
	event := new(BridgeDepositedERC20)
	if err := _Bridge.contract.UnpackLog(event, "DepositedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositedERC721Iterator is returned from FilterDepositedERC721 and is used to iterate over the raw logs and unpacked data for DepositedERC721 events raised by the Bridge contract.
type BridgeDepositedERC721Iterator struct {
	Event *BridgeDepositedERC721 // Event containing the contract specifics and raw log

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
func (it *BridgeDepositedERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositedERC721)
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
		it.Event = new(BridgeDepositedERC721)
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
func (it *BridgeDepositedERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositedERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositedERC721 represents a DepositedERC721 event raised by the Bridge contract.
type BridgeDepositedERC721 struct {
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
func (_Bridge *BridgeFilterer) FilterDepositedERC721(opts *bind.FilterOpts) (*BridgeDepositedERC721Iterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositedERC721")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedERC721Iterator{contract: _Bridge.contract, event: "DepositedERC721", logs: logs, sub: sub}, nil
}

// WatchDepositedERC721 is a free log subscription operation binding the contract event 0x7f787dd0c844dac4f8bfc4044046cdab3be531f7eefa9b740c531e48a99725e1.
//
// Solidity: event DepositedERC721(address token, uint256 tokenId, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_Bridge *BridgeFilterer) WatchDepositedERC721(opts *bind.WatchOpts, sink chan<- *BridgeDepositedERC721) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositedERC721")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositedERC721)
				if err := _Bridge.contract.UnpackLog(event, "DepositedERC721", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseDepositedERC721(log types.Log) (*BridgeDepositedERC721, error) {
	event := new(BridgeDepositedERC721)
	if err := _Bridge.contract.UnpackLog(event, "DepositedERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositedNativeIterator is returned from FilterDepositedNative and is used to iterate over the raw logs and unpacked data for DepositedNative events raised by the Bridge contract.
type BridgeDepositedNativeIterator struct {
	Event *BridgeDepositedNative // Event containing the contract specifics and raw log

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
func (it *BridgeDepositedNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositedNative)
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
		it.Event = new(BridgeDepositedNative)
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
func (it *BridgeDepositedNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositedNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositedNative represents a DepositedNative event raised by the Bridge contract.
type BridgeDepositedNative struct {
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
func (_Bridge *BridgeFilterer) FilterDepositedNative(opts *bind.FilterOpts) (*BridgeDepositedNativeIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositedNative")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedNativeIterator{contract: _Bridge.contract, event: "DepositedNative", logs: logs, sub: sub}, nil
}

// WatchDepositedNative is a free log subscription operation binding the contract event 0x9a47c8733424880a9e86a368eff95da5e7d36b68474a95eb097be2e43c116f27.
//
// Solidity: event DepositedNative(uint256 amount, bytes32 salt, bytes bundle, string network, string receiver)
func (_Bridge *BridgeFilterer) WatchDepositedNative(opts *bind.WatchOpts, sink chan<- *BridgeDepositedNative) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositedNative")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositedNative)
				if err := _Bridge.contract.UnpackLog(event, "DepositedNative", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseDepositedNative(log types.Log) (*BridgeDepositedNative, error) {
	event := new(BridgeDepositedNative)
	if err := _Bridge.contract.UnpackLog(event, "DepositedNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositedSBTIterator is returned from FilterDepositedSBT and is used to iterate over the raw logs and unpacked data for DepositedSBT events raised by the Bridge contract.
type BridgeDepositedSBTIterator struct {
	Event *BridgeDepositedSBT // Event containing the contract specifics and raw log

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
func (it *BridgeDepositedSBTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositedSBT)
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
		it.Event = new(BridgeDepositedSBT)
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
func (it *BridgeDepositedSBTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositedSBTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositedSBT represents a DepositedSBT event raised by the Bridge contract.
type BridgeDepositedSBT struct {
	Token    common.Address
	TokenId  *big.Int
	Salt     [32]byte
	Bundle   []byte
	Network  string
	Receiver string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDepositedSBT is a free log retrieval operation binding the contract event 0x1452835f6476968a263680d4519d5a7d74a60354947a95bad57f7274c339ae86.
//
// Solidity: event DepositedSBT(address token, uint256 tokenId, bytes32 salt, bytes bundle, string network, string receiver)
func (_Bridge *BridgeFilterer) FilterDepositedSBT(opts *bind.FilterOpts) (*BridgeDepositedSBTIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositedSBT")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedSBTIterator{contract: _Bridge.contract, event: "DepositedSBT", logs: logs, sub: sub}, nil
}

// WatchDepositedSBT is a free log subscription operation binding the contract event 0x1452835f6476968a263680d4519d5a7d74a60354947a95bad57f7274c339ae86.
//
// Solidity: event DepositedSBT(address token, uint256 tokenId, bytes32 salt, bytes bundle, string network, string receiver)
func (_Bridge *BridgeFilterer) WatchDepositedSBT(opts *bind.WatchOpts, sink chan<- *BridgeDepositedSBT) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositedSBT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositedSBT)
				if err := _Bridge.contract.UnpackLog(event, "DepositedSBT", log); err != nil {
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

// ParseDepositedSBT is a log parse operation binding the contract event 0x1452835f6476968a263680d4519d5a7d74a60354947a95bad57f7274c339ae86.
//
// Solidity: event DepositedSBT(address token, uint256 tokenId, bytes32 salt, bytes bundle, string network, string receiver)
func (_Bridge *BridgeFilterer) ParseDepositedSBT(log types.Log) (*BridgeDepositedSBT, error) {
	event := new(BridgeDepositedSBT)
	if err := _Bridge.contract.UnpackLog(event, "DepositedSBT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bridge contract.
type BridgeInitializedIterator struct {
	Event *BridgeInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeInitialized)
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
		it.Event = new(BridgeInitialized)
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
func (it *BridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeInitialized represents a Initialized event raised by the Bridge contract.
type BridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridge *BridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeInitializedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeInitializedIterator{contract: _Bridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridge *BridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeInitialized)
				if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseInitialized(log types.Log) (*BridgeInitialized, error) {
	event := new(BridgeInitialized)
	if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Bridge contract.
type BridgeUpgradedIterator struct {
	Event *BridgeUpgraded // Event containing the contract specifics and raw log

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
func (it *BridgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUpgraded)
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
		it.Event = new(BridgeUpgraded)
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
func (it *BridgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUpgraded represents a Upgraded event raised by the Bridge contract.
type BridgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bridge *BridgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BridgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BridgeUpgradedIterator{contract: _Bridge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bridge *BridgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BridgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUpgraded)
				if err := _Bridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseUpgraded(log types.Log) (*BridgeUpgraded, error) {
	event := new(BridgeUpgraded)
	if err := _Bridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeWithdrawnERC1155Iterator is returned from FilterWithdrawnERC1155 and is used to iterate over the raw logs and unpacked data for WithdrawnERC1155 events raised by the Bridge contract.
type BridgeWithdrawnERC1155Iterator struct {
	Event *BridgeWithdrawnERC1155 // Event containing the contract specifics and raw log

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
func (it *BridgeWithdrawnERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeWithdrawnERC1155)
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
		it.Event = new(BridgeWithdrawnERC1155)
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
func (it *BridgeWithdrawnERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeWithdrawnERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeWithdrawnERC1155 represents a WithdrawnERC1155 event raised by the Bridge contract.
type BridgeWithdrawnERC1155 struct {
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
func (_Bridge *BridgeFilterer) FilterWithdrawnERC1155(opts *bind.FilterOpts) (*BridgeWithdrawnERC1155Iterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "WithdrawnERC1155")
	if err != nil {
		return nil, err
	}
	return &BridgeWithdrawnERC1155Iterator{contract: _Bridge.contract, event: "WithdrawnERC1155", logs: logs, sub: sub}, nil
}

// WatchWithdrawnERC1155 is a free log subscription operation binding the contract event 0x30acf6d8c142717c265a7b8ab5be21a380a1ab05e2cd7ce3ca4fdd0c3260303e.
//
// Solidity: event WithdrawnERC1155(address token, uint256 tokenId, string tokenURI, uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_Bridge *BridgeFilterer) WatchWithdrawnERC1155(opts *bind.WatchOpts, sink chan<- *BridgeWithdrawnERC1155) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "WithdrawnERC1155")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeWithdrawnERC1155)
				if err := _Bridge.contract.UnpackLog(event, "WithdrawnERC1155", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseWithdrawnERC1155(log types.Log) (*BridgeWithdrawnERC1155, error) {
	event := new(BridgeWithdrawnERC1155)
	if err := _Bridge.contract.UnpackLog(event, "WithdrawnERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeWithdrawnERC20Iterator is returned from FilterWithdrawnERC20 and is used to iterate over the raw logs and unpacked data for WithdrawnERC20 events raised by the Bridge contract.
type BridgeWithdrawnERC20Iterator struct {
	Event *BridgeWithdrawnERC20 // Event containing the contract specifics and raw log

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
func (it *BridgeWithdrawnERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeWithdrawnERC20)
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
		it.Event = new(BridgeWithdrawnERC20)
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
func (it *BridgeWithdrawnERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeWithdrawnERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeWithdrawnERC20 represents a WithdrawnERC20 event raised by the Bridge contract.
type BridgeWithdrawnERC20 struct {
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
func (_Bridge *BridgeFilterer) FilterWithdrawnERC20(opts *bind.FilterOpts) (*BridgeWithdrawnERC20Iterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "WithdrawnERC20")
	if err != nil {
		return nil, err
	}
	return &BridgeWithdrawnERC20Iterator{contract: _Bridge.contract, event: "WithdrawnERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawnERC20 is a free log subscription operation binding the contract event 0xff14005e9cd01bf3385cde89fb85a18453a0e3df47a3224d84ebab3fbd174128.
//
// Solidity: event WithdrawnERC20(address token, uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_Bridge *BridgeFilterer) WatchWithdrawnERC20(opts *bind.WatchOpts, sink chan<- *BridgeWithdrawnERC20) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "WithdrawnERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeWithdrawnERC20)
				if err := _Bridge.contract.UnpackLog(event, "WithdrawnERC20", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseWithdrawnERC20(log types.Log) (*BridgeWithdrawnERC20, error) {
	event := new(BridgeWithdrawnERC20)
	if err := _Bridge.contract.UnpackLog(event, "WithdrawnERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeWithdrawnERC721Iterator is returned from FilterWithdrawnERC721 and is used to iterate over the raw logs and unpacked data for WithdrawnERC721 events raised by the Bridge contract.
type BridgeWithdrawnERC721Iterator struct {
	Event *BridgeWithdrawnERC721 // Event containing the contract specifics and raw log

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
func (it *BridgeWithdrawnERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeWithdrawnERC721)
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
		it.Event = new(BridgeWithdrawnERC721)
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
func (it *BridgeWithdrawnERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeWithdrawnERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeWithdrawnERC721 represents a WithdrawnERC721 event raised by the Bridge contract.
type BridgeWithdrawnERC721 struct {
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
func (_Bridge *BridgeFilterer) FilterWithdrawnERC721(opts *bind.FilterOpts) (*BridgeWithdrawnERC721Iterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "WithdrawnERC721")
	if err != nil {
		return nil, err
	}
	return &BridgeWithdrawnERC721Iterator{contract: _Bridge.contract, event: "WithdrawnERC721", logs: logs, sub: sub}, nil
}

// WatchWithdrawnERC721 is a free log subscription operation binding the contract event 0x8548956d276e213edf284f96ff17636f64d436743236a9866a7112888c442edd.
//
// Solidity: event WithdrawnERC721(address token, uint256 tokenId, string tokenURI, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_Bridge *BridgeFilterer) WatchWithdrawnERC721(opts *bind.WatchOpts, sink chan<- *BridgeWithdrawnERC721) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "WithdrawnERC721")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeWithdrawnERC721)
				if err := _Bridge.contract.UnpackLog(event, "WithdrawnERC721", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseWithdrawnERC721(log types.Log) (*BridgeWithdrawnERC721, error) {
	event := new(BridgeWithdrawnERC721)
	if err := _Bridge.contract.UnpackLog(event, "WithdrawnERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeWithdrawnNativeIterator is returned from FilterWithdrawnNative and is used to iterate over the raw logs and unpacked data for WithdrawnNative events raised by the Bridge contract.
type BridgeWithdrawnNativeIterator struct {
	Event *BridgeWithdrawnNative // Event containing the contract specifics and raw log

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
func (it *BridgeWithdrawnNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeWithdrawnNative)
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
		it.Event = new(BridgeWithdrawnNative)
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
func (it *BridgeWithdrawnNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeWithdrawnNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeWithdrawnNative represents a WithdrawnNative event raised by the Bridge contract.
type BridgeWithdrawnNative struct {
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
func (_Bridge *BridgeFilterer) FilterWithdrawnNative(opts *bind.FilterOpts) (*BridgeWithdrawnNativeIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "WithdrawnNative")
	if err != nil {
		return nil, err
	}
	return &BridgeWithdrawnNativeIterator{contract: _Bridge.contract, event: "WithdrawnNative", logs: logs, sub: sub}, nil
}

// WatchWithdrawnNative is a free log subscription operation binding the contract event 0x319eed2457c60bc7ee23fdc5be5941c7b628c01f6f2c526dedeca75054d66b18.
//
// Solidity: event WithdrawnNative(uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_Bridge *BridgeFilterer) WatchWithdrawnNative(opts *bind.WatchOpts, sink chan<- *BridgeWithdrawnNative) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "WithdrawnNative")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeWithdrawnNative)
				if err := _Bridge.contract.UnpackLog(event, "WithdrawnNative", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseWithdrawnNative(log types.Log) (*BridgeWithdrawnNative, error) {
	event := new(BridgeWithdrawnNative)
	if err := _Bridge.contract.UnpackLog(event, "WithdrawnNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeWithdrawnSBTIterator is returned from FilterWithdrawnSBT and is used to iterate over the raw logs and unpacked data for WithdrawnSBT events raised by the Bridge contract.
type BridgeWithdrawnSBTIterator struct {
	Event *BridgeWithdrawnSBT // Event containing the contract specifics and raw log

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
func (it *BridgeWithdrawnSBTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeWithdrawnSBT)
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
		it.Event = new(BridgeWithdrawnSBT)
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
func (it *BridgeWithdrawnSBTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeWithdrawnSBTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeWithdrawnSBT represents a WithdrawnSBT event raised by the Bridge contract.
type BridgeWithdrawnSBT struct {
	Token      common.Address
	TokenId    *big.Int
	TokenURI   string
	Salt       [32]byte
	Bundle     []byte
	OriginHash [32]byte
	Receiver   common.Address
	Proof      []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnSBT is a free log retrieval operation binding the contract event 0xbdb8f203981c7db82a7d157c20285399875b02508b52db10a634dd2a72c1b623.
//
// Solidity: event WithdrawnSBT(address token, uint256 tokenId, string tokenURI, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_Bridge *BridgeFilterer) FilterWithdrawnSBT(opts *bind.FilterOpts) (*BridgeWithdrawnSBTIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "WithdrawnSBT")
	if err != nil {
		return nil, err
	}
	return &BridgeWithdrawnSBTIterator{contract: _Bridge.contract, event: "WithdrawnSBT", logs: logs, sub: sub}, nil
}

// WatchWithdrawnSBT is a free log subscription operation binding the contract event 0xbdb8f203981c7db82a7d157c20285399875b02508b52db10a634dd2a72c1b623.
//
// Solidity: event WithdrawnSBT(address token, uint256 tokenId, string tokenURI, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_Bridge *BridgeFilterer) WatchWithdrawnSBT(opts *bind.WatchOpts, sink chan<- *BridgeWithdrawnSBT) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "WithdrawnSBT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeWithdrawnSBT)
				if err := _Bridge.contract.UnpackLog(event, "WithdrawnSBT", log); err != nil {
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

// ParseWithdrawnSBT is a log parse operation binding the contract event 0xbdb8f203981c7db82a7d157c20285399875b02508b52db10a634dd2a72c1b623.
//
// Solidity: event WithdrawnSBT(address token, uint256 tokenId, string tokenURI, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_Bridge *BridgeFilterer) ParseWithdrawnSBT(log types.Log) (*BridgeWithdrawnSBT, error) {
	event := new(BridgeWithdrawnSBT)
	if err := _Bridge.contract.UnpackLog(event, "WithdrawnSBT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
