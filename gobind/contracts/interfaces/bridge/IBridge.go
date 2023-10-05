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

// IBridgeMetaData contains all meta data concerning the IBridge contract.
var IBridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"DepositedERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"DepositedERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"DepositedERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"name\":\"DepositedNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"name\":\"DepositedSBT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"WithdrawnERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"WithdrawnERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"WithdrawnERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"WithdrawnNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"WithdrawnSBT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"checkSignatureAndIncrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.DepositERC1155Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.DepositERC20Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.DepositERC721Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"internalType\":\"structINativeHandler.DepositNativeParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"internalType\":\"structISBTHandler.DepositSBTParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"depositSBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt_\",\"type\":\"bytes32\"}],\"name\":\"determineProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"getSigComponents\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"validateChangeAddressSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"tokenDataLeaf_\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle_\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"}],\"name\":\"verifyMerkleLeaf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.WithdrawERC1155Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.WithdrawERC1155Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC1155Bundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.WithdrawERC20Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.WithdrawERC20Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC20Bundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.WithdrawERC721Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.WithdrawERC721Parameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawERC721Bundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structINativeHandler.WithdrawNativeParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structINativeHandler.WithdrawNativeParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawNativeBundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structISBTHandler.WithdrawSBTParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawSBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"originHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structISBTHandler.WithdrawSBTParameters\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"withdrawSBTBundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use IBridgeMetaData.ABI instead.
var IBridgeABI = IBridgeMetaData.ABI

// IBridge is an auto generated Go binding around an Ethereum contract.
type IBridge struct {
	IBridgeCaller     // Read-only binding to the contract
	IBridgeTransactor // Write-only binding to the contract
	IBridgeFilterer   // Log filterer for contract events
}

// IBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBridgeSession struct {
	Contract     *IBridge          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBridgeCallerSession struct {
	Contract *IBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBridgeTransactorSession struct {
	Contract     *IBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBridgeRaw struct {
	Contract *IBridge // Generic contract binding to access the raw methods on
}

// IBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBridgeCallerRaw struct {
	Contract *IBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBridgeTransactorRaw struct {
	Contract *IBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBridge creates a new instance of IBridge, bound to a specific deployed contract.
func NewIBridge(address common.Address, backend bind.ContractBackend) (*IBridge, error) {
	contract, err := bindIBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBridge{IBridgeCaller: IBridgeCaller{contract: contract}, IBridgeTransactor: IBridgeTransactor{contract: contract}, IBridgeFilterer: IBridgeFilterer{contract: contract}}, nil
}

// NewIBridgeCaller creates a new read-only instance of IBridge, bound to a specific deployed contract.
func NewIBridgeCaller(address common.Address, caller bind.ContractCaller) (*IBridgeCaller, error) {
	contract, err := bindIBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeCaller{contract: contract}, nil
}

// NewIBridgeTransactor creates a new write-only instance of IBridge, bound to a specific deployed contract.
func NewIBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IBridgeTransactor, error) {
	contract, err := bindIBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeTransactor{contract: contract}, nil
}

// NewIBridgeFilterer creates a new log filterer instance of IBridge, bound to a specific deployed contract.
func NewIBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IBridgeFilterer, error) {
	contract, err := bindIBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBridgeFilterer{contract: contract}, nil
}

// bindIBridge binds a generic wrapper to an already deployed contract.
func bindIBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridge *IBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridge.Contract.IBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridge *IBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridge.Contract.IBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridge *IBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridge.Contract.IBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridge *IBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridge *IBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridge *IBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridge.Contract.contract.Transact(opts, method, params...)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_IBridge *IBridgeCaller) DetermineProxyAddress(opts *bind.CallOpts, salt_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "determineProxyAddress", salt_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_IBridge *IBridgeSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _IBridge.Contract.DetermineProxyAddress(&_IBridge.CallOpts, salt_)
}

// DetermineProxyAddress is a free data retrieval call binding the contract method 0x0492e493.
//
// Solidity: function determineProxyAddress(bytes32 salt_) view returns(address)
func (_IBridge *IBridgeCallerSession) DetermineProxyAddress(salt_ [32]byte) (common.Address, error) {
	return _IBridge.Contract.DetermineProxyAddress(&_IBridge.CallOpts, salt_)
}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_IBridge *IBridgeCaller) GetSigComponents(opts *bind.CallOpts, methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "getSigComponents", methodId_, contractAddress_)

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
func (_IBridge *IBridgeSession) GetSigComponents(methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	return _IBridge.Contract.GetSigComponents(&_IBridge.CallOpts, methodId_, contractAddress_)
}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_IBridge *IBridgeCallerSession) GetSigComponents(methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	return _IBridge.Contract.GetSigComponents(&_IBridge.CallOpts, methodId_, contractAddress_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_IBridge *IBridgeTransactor) CheckSignatureAndIncrementNonce(opts *bind.TransactOpts, methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "checkSignatureAndIncrementNonce", methodId_, contractAddress_, signHash_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_IBridge *IBridgeSession) CheckSignatureAndIncrementNonce(methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _IBridge.Contract.CheckSignatureAndIncrementNonce(&_IBridge.TransactOpts, methodId_, contractAddress_, signHash_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_IBridge *IBridgeTransactorSession) CheckSignatureAndIncrementNonce(methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _IBridge.Contract.CheckSignatureAndIncrementNonce(&_IBridge.TransactOpts, methodId_, contractAddress_, signHash_, signature_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x969aef06.
//
// Solidity: function depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IBridge *IBridgeTransactor) DepositERC1155(opts *bind.TransactOpts, params_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "depositERC1155", params_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x969aef06.
//
// Solidity: function depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IBridge *IBridgeSession) DepositERC1155(params_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.DepositERC1155(&_IBridge.TransactOpts, params_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x969aef06.
//
// Solidity: function depositERC1155((address,uint256,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IBridge *IBridgeTransactorSession) DepositERC1155(params_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.DepositERC1155(&_IBridge.TransactOpts, params_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xbebd5a0b.
//
// Solidity: function depositERC20((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IBridge *IBridgeTransactor) DepositERC20(opts *bind.TransactOpts, params_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "depositERC20", params_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xbebd5a0b.
//
// Solidity: function depositERC20((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IBridge *IBridgeSession) DepositERC20(params_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.DepositERC20(&_IBridge.TransactOpts, params_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xbebd5a0b.
//
// Solidity: function depositERC20((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IBridge *IBridgeTransactorSession) DepositERC20(params_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.DepositERC20(&_IBridge.TransactOpts, params_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x6a38abbf.
//
// Solidity: function depositERC721((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IBridge *IBridgeTransactor) DepositERC721(opts *bind.TransactOpts, params_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "depositERC721", params_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x6a38abbf.
//
// Solidity: function depositERC721((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IBridge *IBridgeSession) DepositERC721(params_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.DepositERC721(&_IBridge.TransactOpts, params_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x6a38abbf.
//
// Solidity: function depositERC721((address,uint256,(bytes32,bytes),string,string,bool) params_) returns()
func (_IBridge *IBridgeTransactorSession) DepositERC721(params_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.DepositERC721(&_IBridge.TransactOpts, params_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xb27588e5.
//
// Solidity: function depositNative((uint256,(bytes32,bytes),string,string) params_) payable returns()
func (_IBridge *IBridgeTransactor) DepositNative(opts *bind.TransactOpts, params_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "depositNative", params_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xb27588e5.
//
// Solidity: function depositNative((uint256,(bytes32,bytes),string,string) params_) payable returns()
func (_IBridge *IBridgeSession) DepositNative(params_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _IBridge.Contract.DepositNative(&_IBridge.TransactOpts, params_)
}

// DepositNative is a paid mutator transaction binding the contract method 0xb27588e5.
//
// Solidity: function depositNative((uint256,(bytes32,bytes),string,string) params_) payable returns()
func (_IBridge *IBridgeTransactorSession) DepositNative(params_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _IBridge.Contract.DepositNative(&_IBridge.TransactOpts, params_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0x755f3823.
//
// Solidity: function depositSBT((address,uint256,(bytes32,bytes),string,string) params_) returns()
func (_IBridge *IBridgeTransactor) DepositSBT(opts *bind.TransactOpts, params_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "depositSBT", params_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0x755f3823.
//
// Solidity: function depositSBT((address,uint256,(bytes32,bytes),string,string) params_) returns()
func (_IBridge *IBridgeSession) DepositSBT(params_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _IBridge.Contract.DepositSBT(&_IBridge.TransactOpts, params_)
}

// DepositSBT is a paid mutator transaction binding the contract method 0x755f3823.
//
// Solidity: function depositSBT((address,uint256,(bytes32,bytes),string,string) params_) returns()
func (_IBridge *IBridgeTransactorSession) DepositSBT(params_ ISBTHandlerDepositSBTParameters) (*types.Transaction, error) {
	return _IBridge.Contract.DepositSBT(&_IBridge.TransactOpts, params_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_IBridge *IBridgeTransactor) ValidateChangeAddressSignature(opts *bind.TransactOpts, methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "validateChangeAddressSignature", methodId_, contractAddress_, newAddress_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_IBridge *IBridgeSession) ValidateChangeAddressSignature(methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _IBridge.Contract.ValidateChangeAddressSignature(&_IBridge.TransactOpts, methodId_, contractAddress_, newAddress_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_IBridge *IBridgeTransactorSession) ValidateChangeAddressSignature(methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _IBridge.Contract.ValidateChangeAddressSignature(&_IBridge.TransactOpts, methodId_, contractAddress_, newAddress_, signature_)
}

// VerifyMerkleLeaf is a paid mutator transaction binding the contract method 0x02fd4e26.
//
// Solidity: function verifyMerkleLeaf(bytes tokenDataLeaf_, (bytes32,bytes) bundle_, bytes32 originHash_, address receiver_, bytes proof_) returns()
func (_IBridge *IBridgeTransactor) VerifyMerkleLeaf(opts *bind.TransactOpts, tokenDataLeaf_ []byte, bundle_ IBundlerBundle, originHash_ [32]byte, receiver_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "verifyMerkleLeaf", tokenDataLeaf_, bundle_, originHash_, receiver_, proof_)
}

// VerifyMerkleLeaf is a paid mutator transaction binding the contract method 0x02fd4e26.
//
// Solidity: function verifyMerkleLeaf(bytes tokenDataLeaf_, (bytes32,bytes) bundle_, bytes32 originHash_, address receiver_, bytes proof_) returns()
func (_IBridge *IBridgeSession) VerifyMerkleLeaf(tokenDataLeaf_ []byte, bundle_ IBundlerBundle, originHash_ [32]byte, receiver_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _IBridge.Contract.VerifyMerkleLeaf(&_IBridge.TransactOpts, tokenDataLeaf_, bundle_, originHash_, receiver_, proof_)
}

// VerifyMerkleLeaf is a paid mutator transaction binding the contract method 0x02fd4e26.
//
// Solidity: function verifyMerkleLeaf(bytes tokenDataLeaf_, (bytes32,bytes) bundle_, bytes32 originHash_, address receiver_, bytes proof_) returns()
func (_IBridge *IBridgeTransactorSession) VerifyMerkleLeaf(tokenDataLeaf_ []byte, bundle_ IBundlerBundle, originHash_ [32]byte, receiver_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _IBridge.Contract.VerifyMerkleLeaf(&_IBridge.TransactOpts, tokenDataLeaf_, bundle_, originHash_, receiver_, proof_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeTransactor) WithdrawERC1155(opts *bind.TransactOpts, params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "withdrawERC1155", params_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeSession) WithdrawERC1155(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawERC1155(&_IBridge.TransactOpts, params_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x00903e5d.
//
// Solidity: function withdrawERC1155((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeTransactorSession) WithdrawERC1155(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawERC1155(&_IBridge.TransactOpts, params_)
}

// WithdrawERC1155Bundle is a paid mutator transaction binding the contract method 0xb76da974.
//
// Solidity: function withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeTransactor) WithdrawERC1155Bundle(opts *bind.TransactOpts, params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "withdrawERC1155Bundle", params_)
}

// WithdrawERC1155Bundle is a paid mutator transaction binding the contract method 0xb76da974.
//
// Solidity: function withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeSession) WithdrawERC1155Bundle(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawERC1155Bundle(&_IBridge.TransactOpts, params_)
}

// WithdrawERC1155Bundle is a paid mutator transaction binding the contract method 0xb76da974.
//
// Solidity: function withdrawERC1155Bundle((address,uint256,string,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeTransactorSession) WithdrawERC1155Bundle(params_ IERC1155HandlerWithdrawERC1155Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawERC1155Bundle(&_IBridge.TransactOpts, params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeTransactor) WithdrawERC20(opts *bind.TransactOpts, params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "withdrawERC20", params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeSession) WithdrawERC20(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawERC20(&_IBridge.TransactOpts, params_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x942b0e31.
//
// Solidity: function withdrawERC20((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeTransactorSession) WithdrawERC20(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawERC20(&_IBridge.TransactOpts, params_)
}

// WithdrawERC20Bundle is a paid mutator transaction binding the contract method 0xe53507bd.
//
// Solidity: function withdrawERC20Bundle((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeTransactor) WithdrawERC20Bundle(opts *bind.TransactOpts, params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "withdrawERC20Bundle", params_)
}

// WithdrawERC20Bundle is a paid mutator transaction binding the contract method 0xe53507bd.
//
// Solidity: function withdrawERC20Bundle((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeSession) WithdrawERC20Bundle(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawERC20Bundle(&_IBridge.TransactOpts, params_)
}

// WithdrawERC20Bundle is a paid mutator transaction binding the contract method 0xe53507bd.
//
// Solidity: function withdrawERC20Bundle((address,uint256,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeTransactorSession) WithdrawERC20Bundle(params_ IERC20HandlerWithdrawERC20Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawERC20Bundle(&_IBridge.TransactOpts, params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeTransactor) WithdrawERC721(opts *bind.TransactOpts, params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "withdrawERC721", params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeSession) WithdrawERC721(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawERC721(&_IBridge.TransactOpts, params_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x1c250708.
//
// Solidity: function withdrawERC721((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeTransactorSession) WithdrawERC721(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawERC721(&_IBridge.TransactOpts, params_)
}

// WithdrawERC721Bundle is a paid mutator transaction binding the contract method 0x9ae00a57.
//
// Solidity: function withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeTransactor) WithdrawERC721Bundle(opts *bind.TransactOpts, params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "withdrawERC721Bundle", params_)
}

// WithdrawERC721Bundle is a paid mutator transaction binding the contract method 0x9ae00a57.
//
// Solidity: function withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeSession) WithdrawERC721Bundle(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawERC721Bundle(&_IBridge.TransactOpts, params_)
}

// WithdrawERC721Bundle is a paid mutator transaction binding the contract method 0x9ae00a57.
//
// Solidity: function withdrawERC721Bundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes,bool) params_) returns()
func (_IBridge *IBridgeTransactorSession) WithdrawERC721Bundle(params_ IERC721HandlerWithdrawERC721Parameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawERC721Bundle(&_IBridge.TransactOpts, params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridge *IBridgeTransactor) WithdrawNative(opts *bind.TransactOpts, params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "withdrawNative", params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridge *IBridgeSession) WithdrawNative(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawNative(&_IBridge.TransactOpts, params_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x922e37c0.
//
// Solidity: function withdrawNative((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridge *IBridgeTransactorSession) WithdrawNative(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawNative(&_IBridge.TransactOpts, params_)
}

// WithdrawNativeBundle is a paid mutator transaction binding the contract method 0xc4f54e21.
//
// Solidity: function withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridge *IBridgeTransactor) WithdrawNativeBundle(opts *bind.TransactOpts, params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "withdrawNativeBundle", params_)
}

// WithdrawNativeBundle is a paid mutator transaction binding the contract method 0xc4f54e21.
//
// Solidity: function withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridge *IBridgeSession) WithdrawNativeBundle(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawNativeBundle(&_IBridge.TransactOpts, params_)
}

// WithdrawNativeBundle is a paid mutator transaction binding the contract method 0xc4f54e21.
//
// Solidity: function withdrawNativeBundle((uint256,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridge *IBridgeTransactorSession) WithdrawNativeBundle(params_ INativeHandlerWithdrawNativeParameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawNativeBundle(&_IBridge.TransactOpts, params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridge *IBridgeTransactor) WithdrawSBT(opts *bind.TransactOpts, params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "withdrawSBT", params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridge *IBridgeSession) WithdrawSBT(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawSBT(&_IBridge.TransactOpts, params_)
}

// WithdrawSBT is a paid mutator transaction binding the contract method 0x73710258.
//
// Solidity: function withdrawSBT((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridge *IBridgeTransactorSession) WithdrawSBT(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawSBT(&_IBridge.TransactOpts, params_)
}

// WithdrawSBTBundle is a paid mutator transaction binding the contract method 0x9a830d4c.
//
// Solidity: function withdrawSBTBundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridge *IBridgeTransactor) WithdrawSBTBundle(opts *bind.TransactOpts, params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "withdrawSBTBundle", params_)
}

// WithdrawSBTBundle is a paid mutator transaction binding the contract method 0x9a830d4c.
//
// Solidity: function withdrawSBTBundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridge *IBridgeSession) WithdrawSBTBundle(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawSBTBundle(&_IBridge.TransactOpts, params_)
}

// WithdrawSBTBundle is a paid mutator transaction binding the contract method 0x9a830d4c.
//
// Solidity: function withdrawSBTBundle((address,uint256,string,(bytes32,bytes),bytes32,address,bytes) params_) returns()
func (_IBridge *IBridgeTransactorSession) WithdrawSBTBundle(params_ ISBTHandlerWithdrawSBTParameters) (*types.Transaction, error) {
	return _IBridge.Contract.WithdrawSBTBundle(&_IBridge.TransactOpts, params_)
}

// IBridgeDepositedERC1155Iterator is returned from FilterDepositedERC1155 and is used to iterate over the raw logs and unpacked data for DepositedERC1155 events raised by the IBridge contract.
type IBridgeDepositedERC1155Iterator struct {
	Event *IBridgeDepositedERC1155 // Event containing the contract specifics and raw log

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
func (it *IBridgeDepositedERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeDepositedERC1155)
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
		it.Event = new(IBridgeDepositedERC1155)
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
func (it *IBridgeDepositedERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeDepositedERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeDepositedERC1155 represents a DepositedERC1155 event raised by the IBridge contract.
type IBridgeDepositedERC1155 struct {
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
func (_IBridge *IBridgeFilterer) FilterDepositedERC1155(opts *bind.FilterOpts) (*IBridgeDepositedERC1155Iterator, error) {

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "DepositedERC1155")
	if err != nil {
		return nil, err
	}
	return &IBridgeDepositedERC1155Iterator{contract: _IBridge.contract, event: "DepositedERC1155", logs: logs, sub: sub}, nil
}

// WatchDepositedERC1155 is a free log subscription operation binding the contract event 0x103b790f2fa3a8676ff87c3620a55f0853d0e45128a8c7e9fadf29e17c51d07a.
//
// Solidity: event DepositedERC1155(address token, uint256 tokenId, uint256 amount, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_IBridge *IBridgeFilterer) WatchDepositedERC1155(opts *bind.WatchOpts, sink chan<- *IBridgeDepositedERC1155) (event.Subscription, error) {

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "DepositedERC1155")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeDepositedERC1155)
				if err := _IBridge.contract.UnpackLog(event, "DepositedERC1155", log); err != nil {
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
func (_IBridge *IBridgeFilterer) ParseDepositedERC1155(log types.Log) (*IBridgeDepositedERC1155, error) {
	event := new(IBridgeDepositedERC1155)
	if err := _IBridge.contract.UnpackLog(event, "DepositedERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeDepositedERC20Iterator is returned from FilterDepositedERC20 and is used to iterate over the raw logs and unpacked data for DepositedERC20 events raised by the IBridge contract.
type IBridgeDepositedERC20Iterator struct {
	Event *IBridgeDepositedERC20 // Event containing the contract specifics and raw log

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
func (it *IBridgeDepositedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeDepositedERC20)
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
		it.Event = new(IBridgeDepositedERC20)
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
func (it *IBridgeDepositedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeDepositedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeDepositedERC20 represents a DepositedERC20 event raised by the IBridge contract.
type IBridgeDepositedERC20 struct {
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
func (_IBridge *IBridgeFilterer) FilterDepositedERC20(opts *bind.FilterOpts) (*IBridgeDepositedERC20Iterator, error) {

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "DepositedERC20")
	if err != nil {
		return nil, err
	}
	return &IBridgeDepositedERC20Iterator{contract: _IBridge.contract, event: "DepositedERC20", logs: logs, sub: sub}, nil
}

// WatchDepositedERC20 is a free log subscription operation binding the contract event 0x043d52f9acdd847f0210803c386559db9e09d492143f2072fe30ea62ff0bb639.
//
// Solidity: event DepositedERC20(address token, uint256 amount, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_IBridge *IBridgeFilterer) WatchDepositedERC20(opts *bind.WatchOpts, sink chan<- *IBridgeDepositedERC20) (event.Subscription, error) {

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "DepositedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeDepositedERC20)
				if err := _IBridge.contract.UnpackLog(event, "DepositedERC20", log); err != nil {
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
func (_IBridge *IBridgeFilterer) ParseDepositedERC20(log types.Log) (*IBridgeDepositedERC20, error) {
	event := new(IBridgeDepositedERC20)
	if err := _IBridge.contract.UnpackLog(event, "DepositedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeDepositedERC721Iterator is returned from FilterDepositedERC721 and is used to iterate over the raw logs and unpacked data for DepositedERC721 events raised by the IBridge contract.
type IBridgeDepositedERC721Iterator struct {
	Event *IBridgeDepositedERC721 // Event containing the contract specifics and raw log

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
func (it *IBridgeDepositedERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeDepositedERC721)
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
		it.Event = new(IBridgeDepositedERC721)
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
func (it *IBridgeDepositedERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeDepositedERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeDepositedERC721 represents a DepositedERC721 event raised by the IBridge contract.
type IBridgeDepositedERC721 struct {
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
func (_IBridge *IBridgeFilterer) FilterDepositedERC721(opts *bind.FilterOpts) (*IBridgeDepositedERC721Iterator, error) {

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "DepositedERC721")
	if err != nil {
		return nil, err
	}
	return &IBridgeDepositedERC721Iterator{contract: _IBridge.contract, event: "DepositedERC721", logs: logs, sub: sub}, nil
}

// WatchDepositedERC721 is a free log subscription operation binding the contract event 0x7f787dd0c844dac4f8bfc4044046cdab3be531f7eefa9b740c531e48a99725e1.
//
// Solidity: event DepositedERC721(address token, uint256 tokenId, bytes32 salt, bytes bundle, string network, string receiver, bool isWrapped)
func (_IBridge *IBridgeFilterer) WatchDepositedERC721(opts *bind.WatchOpts, sink chan<- *IBridgeDepositedERC721) (event.Subscription, error) {

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "DepositedERC721")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeDepositedERC721)
				if err := _IBridge.contract.UnpackLog(event, "DepositedERC721", log); err != nil {
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
func (_IBridge *IBridgeFilterer) ParseDepositedERC721(log types.Log) (*IBridgeDepositedERC721, error) {
	event := new(IBridgeDepositedERC721)
	if err := _IBridge.contract.UnpackLog(event, "DepositedERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeDepositedNativeIterator is returned from FilterDepositedNative and is used to iterate over the raw logs and unpacked data for DepositedNative events raised by the IBridge contract.
type IBridgeDepositedNativeIterator struct {
	Event *IBridgeDepositedNative // Event containing the contract specifics and raw log

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
func (it *IBridgeDepositedNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeDepositedNative)
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
		it.Event = new(IBridgeDepositedNative)
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
func (it *IBridgeDepositedNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeDepositedNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeDepositedNative represents a DepositedNative event raised by the IBridge contract.
type IBridgeDepositedNative struct {
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
func (_IBridge *IBridgeFilterer) FilterDepositedNative(opts *bind.FilterOpts) (*IBridgeDepositedNativeIterator, error) {

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "DepositedNative")
	if err != nil {
		return nil, err
	}
	return &IBridgeDepositedNativeIterator{contract: _IBridge.contract, event: "DepositedNative", logs: logs, sub: sub}, nil
}

// WatchDepositedNative is a free log subscription operation binding the contract event 0x9a47c8733424880a9e86a368eff95da5e7d36b68474a95eb097be2e43c116f27.
//
// Solidity: event DepositedNative(uint256 amount, bytes32 salt, bytes bundle, string network, string receiver)
func (_IBridge *IBridgeFilterer) WatchDepositedNative(opts *bind.WatchOpts, sink chan<- *IBridgeDepositedNative) (event.Subscription, error) {

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "DepositedNative")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeDepositedNative)
				if err := _IBridge.contract.UnpackLog(event, "DepositedNative", log); err != nil {
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
func (_IBridge *IBridgeFilterer) ParseDepositedNative(log types.Log) (*IBridgeDepositedNative, error) {
	event := new(IBridgeDepositedNative)
	if err := _IBridge.contract.UnpackLog(event, "DepositedNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeDepositedSBTIterator is returned from FilterDepositedSBT and is used to iterate over the raw logs and unpacked data for DepositedSBT events raised by the IBridge contract.
type IBridgeDepositedSBTIterator struct {
	Event *IBridgeDepositedSBT // Event containing the contract specifics and raw log

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
func (it *IBridgeDepositedSBTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeDepositedSBT)
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
		it.Event = new(IBridgeDepositedSBT)
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
func (it *IBridgeDepositedSBTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeDepositedSBTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeDepositedSBT represents a DepositedSBT event raised by the IBridge contract.
type IBridgeDepositedSBT struct {
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
func (_IBridge *IBridgeFilterer) FilterDepositedSBT(opts *bind.FilterOpts) (*IBridgeDepositedSBTIterator, error) {

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "DepositedSBT")
	if err != nil {
		return nil, err
	}
	return &IBridgeDepositedSBTIterator{contract: _IBridge.contract, event: "DepositedSBT", logs: logs, sub: sub}, nil
}

// WatchDepositedSBT is a free log subscription operation binding the contract event 0x1452835f6476968a263680d4519d5a7d74a60354947a95bad57f7274c339ae86.
//
// Solidity: event DepositedSBT(address token, uint256 tokenId, bytes32 salt, bytes bundle, string network, string receiver)
func (_IBridge *IBridgeFilterer) WatchDepositedSBT(opts *bind.WatchOpts, sink chan<- *IBridgeDepositedSBT) (event.Subscription, error) {

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "DepositedSBT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeDepositedSBT)
				if err := _IBridge.contract.UnpackLog(event, "DepositedSBT", log); err != nil {
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
func (_IBridge *IBridgeFilterer) ParseDepositedSBT(log types.Log) (*IBridgeDepositedSBT, error) {
	event := new(IBridgeDepositedSBT)
	if err := _IBridge.contract.UnpackLog(event, "DepositedSBT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeWithdrawnERC1155Iterator is returned from FilterWithdrawnERC1155 and is used to iterate over the raw logs and unpacked data for WithdrawnERC1155 events raised by the IBridge contract.
type IBridgeWithdrawnERC1155Iterator struct {
	Event *IBridgeWithdrawnERC1155 // Event containing the contract specifics and raw log

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
func (it *IBridgeWithdrawnERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeWithdrawnERC1155)
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
		it.Event = new(IBridgeWithdrawnERC1155)
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
func (it *IBridgeWithdrawnERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeWithdrawnERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeWithdrawnERC1155 represents a WithdrawnERC1155 event raised by the IBridge contract.
type IBridgeWithdrawnERC1155 struct {
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
func (_IBridge *IBridgeFilterer) FilterWithdrawnERC1155(opts *bind.FilterOpts) (*IBridgeWithdrawnERC1155Iterator, error) {

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "WithdrawnERC1155")
	if err != nil {
		return nil, err
	}
	return &IBridgeWithdrawnERC1155Iterator{contract: _IBridge.contract, event: "WithdrawnERC1155", logs: logs, sub: sub}, nil
}

// WatchWithdrawnERC1155 is a free log subscription operation binding the contract event 0x30acf6d8c142717c265a7b8ab5be21a380a1ab05e2cd7ce3ca4fdd0c3260303e.
//
// Solidity: event WithdrawnERC1155(address token, uint256 tokenId, string tokenURI, uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_IBridge *IBridgeFilterer) WatchWithdrawnERC1155(opts *bind.WatchOpts, sink chan<- *IBridgeWithdrawnERC1155) (event.Subscription, error) {

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "WithdrawnERC1155")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeWithdrawnERC1155)
				if err := _IBridge.contract.UnpackLog(event, "WithdrawnERC1155", log); err != nil {
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
func (_IBridge *IBridgeFilterer) ParseWithdrawnERC1155(log types.Log) (*IBridgeWithdrawnERC1155, error) {
	event := new(IBridgeWithdrawnERC1155)
	if err := _IBridge.contract.UnpackLog(event, "WithdrawnERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeWithdrawnERC20Iterator is returned from FilterWithdrawnERC20 and is used to iterate over the raw logs and unpacked data for WithdrawnERC20 events raised by the IBridge contract.
type IBridgeWithdrawnERC20Iterator struct {
	Event *IBridgeWithdrawnERC20 // Event containing the contract specifics and raw log

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
func (it *IBridgeWithdrawnERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeWithdrawnERC20)
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
		it.Event = new(IBridgeWithdrawnERC20)
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
func (it *IBridgeWithdrawnERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeWithdrawnERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeWithdrawnERC20 represents a WithdrawnERC20 event raised by the IBridge contract.
type IBridgeWithdrawnERC20 struct {
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
func (_IBridge *IBridgeFilterer) FilterWithdrawnERC20(opts *bind.FilterOpts) (*IBridgeWithdrawnERC20Iterator, error) {

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "WithdrawnERC20")
	if err != nil {
		return nil, err
	}
	return &IBridgeWithdrawnERC20Iterator{contract: _IBridge.contract, event: "WithdrawnERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawnERC20 is a free log subscription operation binding the contract event 0xff14005e9cd01bf3385cde89fb85a18453a0e3df47a3224d84ebab3fbd174128.
//
// Solidity: event WithdrawnERC20(address token, uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_IBridge *IBridgeFilterer) WatchWithdrawnERC20(opts *bind.WatchOpts, sink chan<- *IBridgeWithdrawnERC20) (event.Subscription, error) {

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "WithdrawnERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeWithdrawnERC20)
				if err := _IBridge.contract.UnpackLog(event, "WithdrawnERC20", log); err != nil {
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
func (_IBridge *IBridgeFilterer) ParseWithdrawnERC20(log types.Log) (*IBridgeWithdrawnERC20, error) {
	event := new(IBridgeWithdrawnERC20)
	if err := _IBridge.contract.UnpackLog(event, "WithdrawnERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeWithdrawnERC721Iterator is returned from FilterWithdrawnERC721 and is used to iterate over the raw logs and unpacked data for WithdrawnERC721 events raised by the IBridge contract.
type IBridgeWithdrawnERC721Iterator struct {
	Event *IBridgeWithdrawnERC721 // Event containing the contract specifics and raw log

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
func (it *IBridgeWithdrawnERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeWithdrawnERC721)
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
		it.Event = new(IBridgeWithdrawnERC721)
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
func (it *IBridgeWithdrawnERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeWithdrawnERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeWithdrawnERC721 represents a WithdrawnERC721 event raised by the IBridge contract.
type IBridgeWithdrawnERC721 struct {
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
func (_IBridge *IBridgeFilterer) FilterWithdrawnERC721(opts *bind.FilterOpts) (*IBridgeWithdrawnERC721Iterator, error) {

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "WithdrawnERC721")
	if err != nil {
		return nil, err
	}
	return &IBridgeWithdrawnERC721Iterator{contract: _IBridge.contract, event: "WithdrawnERC721", logs: logs, sub: sub}, nil
}

// WatchWithdrawnERC721 is a free log subscription operation binding the contract event 0x8548956d276e213edf284f96ff17636f64d436743236a9866a7112888c442edd.
//
// Solidity: event WithdrawnERC721(address token, uint256 tokenId, string tokenURI, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof, bool isWrapped)
func (_IBridge *IBridgeFilterer) WatchWithdrawnERC721(opts *bind.WatchOpts, sink chan<- *IBridgeWithdrawnERC721) (event.Subscription, error) {

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "WithdrawnERC721")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeWithdrawnERC721)
				if err := _IBridge.contract.UnpackLog(event, "WithdrawnERC721", log); err != nil {
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
func (_IBridge *IBridgeFilterer) ParseWithdrawnERC721(log types.Log) (*IBridgeWithdrawnERC721, error) {
	event := new(IBridgeWithdrawnERC721)
	if err := _IBridge.contract.UnpackLog(event, "WithdrawnERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeWithdrawnNativeIterator is returned from FilterWithdrawnNative and is used to iterate over the raw logs and unpacked data for WithdrawnNative events raised by the IBridge contract.
type IBridgeWithdrawnNativeIterator struct {
	Event *IBridgeWithdrawnNative // Event containing the contract specifics and raw log

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
func (it *IBridgeWithdrawnNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeWithdrawnNative)
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
		it.Event = new(IBridgeWithdrawnNative)
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
func (it *IBridgeWithdrawnNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeWithdrawnNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeWithdrawnNative represents a WithdrawnNative event raised by the IBridge contract.
type IBridgeWithdrawnNative struct {
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
func (_IBridge *IBridgeFilterer) FilterWithdrawnNative(opts *bind.FilterOpts) (*IBridgeWithdrawnNativeIterator, error) {

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "WithdrawnNative")
	if err != nil {
		return nil, err
	}
	return &IBridgeWithdrawnNativeIterator{contract: _IBridge.contract, event: "WithdrawnNative", logs: logs, sub: sub}, nil
}

// WatchWithdrawnNative is a free log subscription operation binding the contract event 0x319eed2457c60bc7ee23fdc5be5941c7b628c01f6f2c526dedeca75054d66b18.
//
// Solidity: event WithdrawnNative(uint256 amount, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_IBridge *IBridgeFilterer) WatchWithdrawnNative(opts *bind.WatchOpts, sink chan<- *IBridgeWithdrawnNative) (event.Subscription, error) {

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "WithdrawnNative")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeWithdrawnNative)
				if err := _IBridge.contract.UnpackLog(event, "WithdrawnNative", log); err != nil {
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
func (_IBridge *IBridgeFilterer) ParseWithdrawnNative(log types.Log) (*IBridgeWithdrawnNative, error) {
	event := new(IBridgeWithdrawnNative)
	if err := _IBridge.contract.UnpackLog(event, "WithdrawnNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeWithdrawnSBTIterator is returned from FilterWithdrawnSBT and is used to iterate over the raw logs and unpacked data for WithdrawnSBT events raised by the IBridge contract.
type IBridgeWithdrawnSBTIterator struct {
	Event *IBridgeWithdrawnSBT // Event containing the contract specifics and raw log

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
func (it *IBridgeWithdrawnSBTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeWithdrawnSBT)
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
		it.Event = new(IBridgeWithdrawnSBT)
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
func (it *IBridgeWithdrawnSBTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeWithdrawnSBTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeWithdrawnSBT represents a WithdrawnSBT event raised by the IBridge contract.
type IBridgeWithdrawnSBT struct {
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
func (_IBridge *IBridgeFilterer) FilterWithdrawnSBT(opts *bind.FilterOpts) (*IBridgeWithdrawnSBTIterator, error) {

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "WithdrawnSBT")
	if err != nil {
		return nil, err
	}
	return &IBridgeWithdrawnSBTIterator{contract: _IBridge.contract, event: "WithdrawnSBT", logs: logs, sub: sub}, nil
}

// WatchWithdrawnSBT is a free log subscription operation binding the contract event 0xbdb8f203981c7db82a7d157c20285399875b02508b52db10a634dd2a72c1b623.
//
// Solidity: event WithdrawnSBT(address token, uint256 tokenId, string tokenURI, bytes32 salt, bytes bundle, bytes32 originHash, address receiver, bytes proof)
func (_IBridge *IBridgeFilterer) WatchWithdrawnSBT(opts *bind.WatchOpts, sink chan<- *IBridgeWithdrawnSBT) (event.Subscription, error) {

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "WithdrawnSBT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeWithdrawnSBT)
				if err := _IBridge.contract.UnpackLog(event, "WithdrawnSBT", log); err != nil {
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
func (_IBridge *IBridgeFilterer) ParseWithdrawnSBT(log types.Log) (*IBridgeWithdrawnSBT, error) {
	event := new(IBridgeWithdrawnSBT)
	if err := _IBridge.contract.UnpackLog(event, "WithdrawnSBT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
