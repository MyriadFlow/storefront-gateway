// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package flowSubscription

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

// FlowSubscriptionMetaData contains all meta data concerning the FlowSubscription contract.
var FlowSubscriptionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_initialURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_publicSalePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_subscriptionPricePerMonth\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"royaltyBasisPoint\",\"type\":\"uint96\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"Time\",\"type\":\"uint256\"}],\"name\":\"SubscriptionCancelRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SubscriptionIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerOrApproved\",\"type\":\"address\"}],\"name\":\"SubscriptionRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiration\",\"type\":\"uint64\"}],\"name\":\"SubscriptionUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MONTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelSubscription\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cancellationRequested\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"delegateSubscribe\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"expiresAt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isRenewable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicSalePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"renewSubscription\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"revokeSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_publicSalePrice\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenUri\",\"type\":\"string\"}],\"name\":\"setTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subscribe\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subscriptionPricePerMonth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FlowSubscriptionABI is the input ABI used to generate the binding from.
// Deprecated: Use FlowSubscriptionMetaData.ABI instead.
var FlowSubscriptionABI = FlowSubscriptionMetaData.ABI

// FlowSubscription is an auto generated Go binding around an Ethereum contract.
type FlowSubscription struct {
	FlowSubscriptionCaller     // Read-only binding to the contract
	FlowSubscriptionTransactor // Write-only binding to the contract
	FlowSubscriptionFilterer   // Log filterer for contract events
}

// FlowSubscriptionCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlowSubscriptionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlowSubscriptionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlowSubscriptionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlowSubscriptionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlowSubscriptionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlowSubscriptionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlowSubscriptionSession struct {
	Contract     *FlowSubscription // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlowSubscriptionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlowSubscriptionCallerSession struct {
	Contract *FlowSubscriptionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// FlowSubscriptionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlowSubscriptionTransactorSession struct {
	Contract     *FlowSubscriptionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FlowSubscriptionRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlowSubscriptionRaw struct {
	Contract *FlowSubscription // Generic contract binding to access the raw methods on
}

// FlowSubscriptionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlowSubscriptionCallerRaw struct {
	Contract *FlowSubscriptionCaller // Generic read-only contract binding to access the raw methods on
}

// FlowSubscriptionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlowSubscriptionTransactorRaw struct {
	Contract *FlowSubscriptionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlowSubscription creates a new instance of FlowSubscription, bound to a specific deployed contract.
func NewFlowSubscription(address common.Address, backend bind.ContractBackend) (*FlowSubscription, error) {
	contract, err := bindFlowSubscription(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlowSubscription{FlowSubscriptionCaller: FlowSubscriptionCaller{contract: contract}, FlowSubscriptionTransactor: FlowSubscriptionTransactor{contract: contract}, FlowSubscriptionFilterer: FlowSubscriptionFilterer{contract: contract}}, nil
}

// NewFlowSubscriptionCaller creates a new read-only instance of FlowSubscription, bound to a specific deployed contract.
func NewFlowSubscriptionCaller(address common.Address, caller bind.ContractCaller) (*FlowSubscriptionCaller, error) {
	contract, err := bindFlowSubscription(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlowSubscriptionCaller{contract: contract}, nil
}

// NewFlowSubscriptionTransactor creates a new write-only instance of FlowSubscription, bound to a specific deployed contract.
func NewFlowSubscriptionTransactor(address common.Address, transactor bind.ContractTransactor) (*FlowSubscriptionTransactor, error) {
	contract, err := bindFlowSubscription(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlowSubscriptionTransactor{contract: contract}, nil
}

// NewFlowSubscriptionFilterer creates a new log filterer instance of FlowSubscription, bound to a specific deployed contract.
func NewFlowSubscriptionFilterer(address common.Address, filterer bind.ContractFilterer) (*FlowSubscriptionFilterer, error) {
	contract, err := bindFlowSubscription(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlowSubscriptionFilterer{contract: contract}, nil
}

// bindFlowSubscription binds a generic wrapper to an already deployed contract.
func bindFlowSubscription(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FlowSubscriptionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlowSubscription *FlowSubscriptionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlowSubscription.Contract.FlowSubscriptionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlowSubscription *FlowSubscriptionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlowSubscription.Contract.FlowSubscriptionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlowSubscription *FlowSubscriptionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlowSubscription.Contract.FlowSubscriptionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlowSubscription *FlowSubscriptionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlowSubscription.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlowSubscription *FlowSubscriptionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlowSubscription.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlowSubscription *FlowSubscriptionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlowSubscription.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_FlowSubscription *FlowSubscriptionCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_FlowSubscription *FlowSubscriptionSession) ADMINROLE() ([32]byte, error) {
	return _FlowSubscription.Contract.ADMINROLE(&_FlowSubscription.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_FlowSubscription *FlowSubscriptionCallerSession) ADMINROLE() ([32]byte, error) {
	return _FlowSubscription.Contract.ADMINROLE(&_FlowSubscription.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FlowSubscription *FlowSubscriptionCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FlowSubscription *FlowSubscriptionSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FlowSubscription.Contract.DEFAULTADMINROLE(&_FlowSubscription.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FlowSubscription *FlowSubscriptionCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FlowSubscription.Contract.DEFAULTADMINROLE(&_FlowSubscription.CallOpts)
}

// MONTH is a free data retrieval call binding the contract method 0xd5999a5c.
//
// Solidity: function MONTH() view returns(uint256)
func (_FlowSubscription *FlowSubscriptionCaller) MONTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "MONTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MONTH is a free data retrieval call binding the contract method 0xd5999a5c.
//
// Solidity: function MONTH() view returns(uint256)
func (_FlowSubscription *FlowSubscriptionSession) MONTH() (*big.Int, error) {
	return _FlowSubscription.Contract.MONTH(&_FlowSubscription.CallOpts)
}

// MONTH is a free data retrieval call binding the contract method 0xd5999a5c.
//
// Solidity: function MONTH() view returns(uint256)
func (_FlowSubscription *FlowSubscriptionCallerSession) MONTH() (*big.Int, error) {
	return _FlowSubscription.Contract.MONTH(&_FlowSubscription.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_FlowSubscription *FlowSubscriptionCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_FlowSubscription *FlowSubscriptionSession) OPERATORROLE() ([32]byte, error) {
	return _FlowSubscription.Contract.OPERATORROLE(&_FlowSubscription.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_FlowSubscription *FlowSubscriptionCallerSession) OPERATORROLE() ([32]byte, error) {
	return _FlowSubscription.Contract.OPERATORROLE(&_FlowSubscription.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FlowSubscription *FlowSubscriptionCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FlowSubscription *FlowSubscriptionSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _FlowSubscription.Contract.BalanceOf(&_FlowSubscription.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FlowSubscription *FlowSubscriptionCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _FlowSubscription.Contract.BalanceOf(&_FlowSubscription.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_FlowSubscription *FlowSubscriptionCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_FlowSubscription *FlowSubscriptionSession) BaseURI() (string, error) {
	return _FlowSubscription.Contract.BaseURI(&_FlowSubscription.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_FlowSubscription *FlowSubscriptionCallerSession) BaseURI() (string, error) {
	return _FlowSubscription.Contract.BaseURI(&_FlowSubscription.CallOpts)
}

// CancellationRequested is a free data retrieval call binding the contract method 0x210542c7.
//
// Solidity: function cancellationRequested(uint256 ) view returns(bool)
func (_FlowSubscription *FlowSubscriptionCaller) CancellationRequested(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "cancellationRequested", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CancellationRequested is a free data retrieval call binding the contract method 0x210542c7.
//
// Solidity: function cancellationRequested(uint256 ) view returns(bool)
func (_FlowSubscription *FlowSubscriptionSession) CancellationRequested(arg0 *big.Int) (bool, error) {
	return _FlowSubscription.Contract.CancellationRequested(&_FlowSubscription.CallOpts, arg0)
}

// CancellationRequested is a free data retrieval call binding the contract method 0x210542c7.
//
// Solidity: function cancellationRequested(uint256 ) view returns(bool)
func (_FlowSubscription *FlowSubscriptionCallerSession) CancellationRequested(arg0 *big.Int) (bool, error) {
	return _FlowSubscription.Contract.CancellationRequested(&_FlowSubscription.CallOpts, arg0)
}

// ExpiresAt is a free data retrieval call binding the contract method 0x17c95709.
//
// Solidity: function expiresAt(uint256 tokenId) view returns(uint64)
func (_FlowSubscription *FlowSubscriptionCaller) ExpiresAt(opts *bind.CallOpts, tokenId *big.Int) (uint64, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "expiresAt", tokenId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ExpiresAt is a free data retrieval call binding the contract method 0x17c95709.
//
// Solidity: function expiresAt(uint256 tokenId) view returns(uint64)
func (_FlowSubscription *FlowSubscriptionSession) ExpiresAt(tokenId *big.Int) (uint64, error) {
	return _FlowSubscription.Contract.ExpiresAt(&_FlowSubscription.CallOpts, tokenId)
}

// ExpiresAt is a free data retrieval call binding the contract method 0x17c95709.
//
// Solidity: function expiresAt(uint256 tokenId) view returns(uint64)
func (_FlowSubscription *FlowSubscriptionCallerSession) ExpiresAt(tokenId *big.Int) (uint64, error) {
	return _FlowSubscription.Contract.ExpiresAt(&_FlowSubscription.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FlowSubscription *FlowSubscriptionCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FlowSubscription *FlowSubscriptionSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _FlowSubscription.Contract.GetApproved(&_FlowSubscription.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FlowSubscription *FlowSubscriptionCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _FlowSubscription.Contract.GetApproved(&_FlowSubscription.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FlowSubscription *FlowSubscriptionCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FlowSubscription *FlowSubscriptionSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FlowSubscription.Contract.GetRoleAdmin(&_FlowSubscription.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FlowSubscription *FlowSubscriptionCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FlowSubscription.Contract.GetRoleAdmin(&_FlowSubscription.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_FlowSubscription *FlowSubscriptionCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_FlowSubscription *FlowSubscriptionSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _FlowSubscription.Contract.GetRoleMember(&_FlowSubscription.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_FlowSubscription *FlowSubscriptionCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _FlowSubscription.Contract.GetRoleMember(&_FlowSubscription.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_FlowSubscription *FlowSubscriptionCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_FlowSubscription *FlowSubscriptionSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _FlowSubscription.Contract.GetRoleMemberCount(&_FlowSubscription.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_FlowSubscription *FlowSubscriptionCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _FlowSubscription.Contract.GetRoleMemberCount(&_FlowSubscription.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FlowSubscription *FlowSubscriptionCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FlowSubscription *FlowSubscriptionSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FlowSubscription.Contract.HasRole(&_FlowSubscription.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FlowSubscription *FlowSubscriptionCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FlowSubscription.Contract.HasRole(&_FlowSubscription.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FlowSubscription *FlowSubscriptionCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FlowSubscription *FlowSubscriptionSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _FlowSubscription.Contract.IsApprovedForAll(&_FlowSubscription.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FlowSubscription *FlowSubscriptionCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _FlowSubscription.Contract.IsApprovedForAll(&_FlowSubscription.CallOpts, owner, operator)
}

// IsRenewable is a free data retrieval call binding the contract method 0xcde317af.
//
// Solidity: function isRenewable(uint256 tokenId) view returns(bool)
func (_FlowSubscription *FlowSubscriptionCaller) IsRenewable(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "isRenewable", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRenewable is a free data retrieval call binding the contract method 0xcde317af.
//
// Solidity: function isRenewable(uint256 tokenId) view returns(bool)
func (_FlowSubscription *FlowSubscriptionSession) IsRenewable(tokenId *big.Int) (bool, error) {
	return _FlowSubscription.Contract.IsRenewable(&_FlowSubscription.CallOpts, tokenId)
}

// IsRenewable is a free data retrieval call binding the contract method 0xcde317af.
//
// Solidity: function isRenewable(uint256 tokenId) view returns(bool)
func (_FlowSubscription *FlowSubscriptionCallerSession) IsRenewable(tokenId *big.Int) (bool, error) {
	return _FlowSubscription.Contract.IsRenewable(&_FlowSubscription.CallOpts, tokenId)
}

// MintPaused is a free data retrieval call binding the contract method 0x7e4831d3.
//
// Solidity: function mintPaused() view returns(bool)
func (_FlowSubscription *FlowSubscriptionCaller) MintPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "mintPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintPaused is a free data retrieval call binding the contract method 0x7e4831d3.
//
// Solidity: function mintPaused() view returns(bool)
func (_FlowSubscription *FlowSubscriptionSession) MintPaused() (bool, error) {
	return _FlowSubscription.Contract.MintPaused(&_FlowSubscription.CallOpts)
}

// MintPaused is a free data retrieval call binding the contract method 0x7e4831d3.
//
// Solidity: function mintPaused() view returns(bool)
func (_FlowSubscription *FlowSubscriptionCallerSession) MintPaused() (bool, error) {
	return _FlowSubscription.Contract.MintPaused(&_FlowSubscription.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FlowSubscription *FlowSubscriptionCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FlowSubscription *FlowSubscriptionSession) Name() (string, error) {
	return _FlowSubscription.Contract.Name(&_FlowSubscription.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FlowSubscription *FlowSubscriptionCallerSession) Name() (string, error) {
	return _FlowSubscription.Contract.Name(&_FlowSubscription.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FlowSubscription *FlowSubscriptionCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FlowSubscription *FlowSubscriptionSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _FlowSubscription.Contract.OwnerOf(&_FlowSubscription.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FlowSubscription *FlowSubscriptionCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _FlowSubscription.Contract.OwnerOf(&_FlowSubscription.CallOpts, tokenId)
}

// PublicSalePrice is a free data retrieval call binding the contract method 0x9b6860c8.
//
// Solidity: function publicSalePrice() view returns(uint256)
func (_FlowSubscription *FlowSubscriptionCaller) PublicSalePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "publicSalePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicSalePrice is a free data retrieval call binding the contract method 0x9b6860c8.
//
// Solidity: function publicSalePrice() view returns(uint256)
func (_FlowSubscription *FlowSubscriptionSession) PublicSalePrice() (*big.Int, error) {
	return _FlowSubscription.Contract.PublicSalePrice(&_FlowSubscription.CallOpts)
}

// PublicSalePrice is a free data retrieval call binding the contract method 0x9b6860c8.
//
// Solidity: function publicSalePrice() view returns(uint256)
func (_FlowSubscription *FlowSubscriptionCallerSession) PublicSalePrice() (*big.Int, error) {
	return _FlowSubscription.Contract.PublicSalePrice(&_FlowSubscription.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_FlowSubscription *FlowSubscriptionCaller) RoyaltyInfo(opts *bind.CallOpts, _tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "royaltyInfo", _tokenId, _salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_FlowSubscription *FlowSubscriptionSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _FlowSubscription.Contract.RoyaltyInfo(&_FlowSubscription.CallOpts, _tokenId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_FlowSubscription *FlowSubscriptionCallerSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _FlowSubscription.Contract.RoyaltyInfo(&_FlowSubscription.CallOpts, _tokenId, _salePrice)
}

// SubscriptionPricePerMonth is a free data retrieval call binding the contract method 0x022b50c0.
//
// Solidity: function subscriptionPricePerMonth() view returns(uint256)
func (_FlowSubscription *FlowSubscriptionCaller) SubscriptionPricePerMonth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "subscriptionPricePerMonth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubscriptionPricePerMonth is a free data retrieval call binding the contract method 0x022b50c0.
//
// Solidity: function subscriptionPricePerMonth() view returns(uint256)
func (_FlowSubscription *FlowSubscriptionSession) SubscriptionPricePerMonth() (*big.Int, error) {
	return _FlowSubscription.Contract.SubscriptionPricePerMonth(&_FlowSubscription.CallOpts)
}

// SubscriptionPricePerMonth is a free data retrieval call binding the contract method 0x022b50c0.
//
// Solidity: function subscriptionPricePerMonth() view returns(uint256)
func (_FlowSubscription *FlowSubscriptionCallerSession) SubscriptionPricePerMonth() (*big.Int, error) {
	return _FlowSubscription.Contract.SubscriptionPricePerMonth(&_FlowSubscription.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FlowSubscription *FlowSubscriptionCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FlowSubscription *FlowSubscriptionSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FlowSubscription.Contract.SupportsInterface(&_FlowSubscription.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FlowSubscription *FlowSubscriptionCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FlowSubscription.Contract.SupportsInterface(&_FlowSubscription.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FlowSubscription *FlowSubscriptionCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FlowSubscription *FlowSubscriptionSession) Symbol() (string, error) {
	return _FlowSubscription.Contract.Symbol(&_FlowSubscription.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FlowSubscription *FlowSubscriptionCallerSession) Symbol() (string, error) {
	return _FlowSubscription.Contract.Symbol(&_FlowSubscription.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_FlowSubscription *FlowSubscriptionCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_FlowSubscription *FlowSubscriptionSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _FlowSubscription.Contract.TokenByIndex(&_FlowSubscription.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_FlowSubscription *FlowSubscriptionCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _FlowSubscription.Contract.TokenByIndex(&_FlowSubscription.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FlowSubscription *FlowSubscriptionCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FlowSubscription *FlowSubscriptionSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _FlowSubscription.Contract.TokenOfOwnerByIndex(&_FlowSubscription.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FlowSubscription *FlowSubscriptionCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _FlowSubscription.Contract.TokenOfOwnerByIndex(&_FlowSubscription.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FlowSubscription *FlowSubscriptionCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FlowSubscription *FlowSubscriptionSession) TokenURI(tokenId *big.Int) (string, error) {
	return _FlowSubscription.Contract.TokenURI(&_FlowSubscription.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FlowSubscription *FlowSubscriptionCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _FlowSubscription.Contract.TokenURI(&_FlowSubscription.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FlowSubscription *FlowSubscriptionCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FlowSubscription.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FlowSubscription *FlowSubscriptionSession) TotalSupply() (*big.Int, error) {
	return _FlowSubscription.Contract.TotalSupply(&_FlowSubscription.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FlowSubscription *FlowSubscriptionCallerSession) TotalSupply() (*big.Int, error) {
	return _FlowSubscription.Contract.TotalSupply(&_FlowSubscription.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_FlowSubscription *FlowSubscriptionTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_FlowSubscription *FlowSubscriptionSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.Contract.Approve(&_FlowSubscription.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.Contract.Approve(&_FlowSubscription.TransactOpts, to, tokenId)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0x21235083.
//
// Solidity: function cancelSubscription(uint256 tokenId) payable returns()
func (_FlowSubscription *FlowSubscriptionTransactor) CancelSubscription(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "cancelSubscription", tokenId)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0x21235083.
//
// Solidity: function cancelSubscription(uint256 tokenId) payable returns()
func (_FlowSubscription *FlowSubscriptionSession) CancelSubscription(tokenId *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.Contract.CancelSubscription(&_FlowSubscription.TransactOpts, tokenId)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0x21235083.
//
// Solidity: function cancelSubscription(uint256 tokenId) payable returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) CancelSubscription(tokenId *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.Contract.CancelSubscription(&_FlowSubscription.TransactOpts, tokenId)
}

// DelegateSubscribe is a paid mutator transaction binding the contract method 0x3171f5b0.
//
// Solidity: function delegateSubscribe(address creator) returns(uint256 tokenId)
func (_FlowSubscription *FlowSubscriptionTransactor) DelegateSubscribe(opts *bind.TransactOpts, creator common.Address) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "delegateSubscribe", creator)
}

// DelegateSubscribe is a paid mutator transaction binding the contract method 0x3171f5b0.
//
// Solidity: function delegateSubscribe(address creator) returns(uint256 tokenId)
func (_FlowSubscription *FlowSubscriptionSession) DelegateSubscribe(creator common.Address) (*types.Transaction, error) {
	return _FlowSubscription.Contract.DelegateSubscribe(&_FlowSubscription.TransactOpts, creator)
}

// DelegateSubscribe is a paid mutator transaction binding the contract method 0x3171f5b0.
//
// Solidity: function delegateSubscribe(address creator) returns(uint256 tokenId)
func (_FlowSubscription *FlowSubscriptionTransactorSession) DelegateSubscribe(creator common.Address) (*types.Transaction, error) {
	return _FlowSubscription.Contract.DelegateSubscribe(&_FlowSubscription.TransactOpts, creator)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FlowSubscription *FlowSubscriptionTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FlowSubscription *FlowSubscriptionSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FlowSubscription.Contract.GrantRole(&_FlowSubscription.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FlowSubscription.Contract.GrantRole(&_FlowSubscription.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_FlowSubscription *FlowSubscriptionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_FlowSubscription *FlowSubscriptionSession) Pause() (*types.Transaction, error) {
	return _FlowSubscription.Contract.Pause(&_FlowSubscription.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) Pause() (*types.Transaction, error) {
	return _FlowSubscription.Contract.Pause(&_FlowSubscription.TransactOpts)
}

// RenewSubscription is a paid mutator transaction binding the contract method 0x776ce868.
//
// Solidity: function renewSubscription(uint256 tokenId, uint64 duration) payable returns()
func (_FlowSubscription *FlowSubscriptionTransactor) RenewSubscription(opts *bind.TransactOpts, tokenId *big.Int, duration uint64) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "renewSubscription", tokenId, duration)
}

// RenewSubscription is a paid mutator transaction binding the contract method 0x776ce868.
//
// Solidity: function renewSubscription(uint256 tokenId, uint64 duration) payable returns()
func (_FlowSubscription *FlowSubscriptionSession) RenewSubscription(tokenId *big.Int, duration uint64) (*types.Transaction, error) {
	return _FlowSubscription.Contract.RenewSubscription(&_FlowSubscription.TransactOpts, tokenId, duration)
}

// RenewSubscription is a paid mutator transaction binding the contract method 0x776ce868.
//
// Solidity: function renewSubscription(uint256 tokenId, uint64 duration) payable returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) RenewSubscription(tokenId *big.Int, duration uint64) (*types.Transaction, error) {
	return _FlowSubscription.Contract.RenewSubscription(&_FlowSubscription.TransactOpts, tokenId, duration)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FlowSubscription *FlowSubscriptionTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FlowSubscription *FlowSubscriptionSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FlowSubscription.Contract.RenounceRole(&_FlowSubscription.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FlowSubscription.Contract.RenounceRole(&_FlowSubscription.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FlowSubscription *FlowSubscriptionTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FlowSubscription *FlowSubscriptionSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FlowSubscription.Contract.RevokeRole(&_FlowSubscription.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FlowSubscription.Contract.RevokeRole(&_FlowSubscription.TransactOpts, role, account)
}

// RevokeSubscription is a paid mutator transaction binding the contract method 0x659e2aef.
//
// Solidity: function revokeSubscription(uint256 _tokenId) returns()
func (_FlowSubscription *FlowSubscriptionTransactor) RevokeSubscription(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "revokeSubscription", _tokenId)
}

// RevokeSubscription is a paid mutator transaction binding the contract method 0x659e2aef.
//
// Solidity: function revokeSubscription(uint256 _tokenId) returns()
func (_FlowSubscription *FlowSubscriptionSession) RevokeSubscription(_tokenId *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.Contract.RevokeSubscription(&_FlowSubscription.TransactOpts, _tokenId)
}

// RevokeSubscription is a paid mutator transaction binding the contract method 0x659e2aef.
//
// Solidity: function revokeSubscription(uint256 _tokenId) returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) RevokeSubscription(_tokenId *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.Contract.RevokeSubscription(&_FlowSubscription.TransactOpts, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FlowSubscription *FlowSubscriptionTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FlowSubscription *FlowSubscriptionSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.Contract.SafeTransferFrom(&_FlowSubscription.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.Contract.SafeTransferFrom(&_FlowSubscription.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_FlowSubscription *FlowSubscriptionTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_FlowSubscription *FlowSubscriptionSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _FlowSubscription.Contract.SafeTransferFrom0(&_FlowSubscription.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _FlowSubscription.Contract.SafeTransferFrom0(&_FlowSubscription.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FlowSubscription *FlowSubscriptionTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FlowSubscription *FlowSubscriptionSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FlowSubscription.Contract.SetApprovalForAll(&_FlowSubscription.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FlowSubscription.Contract.SetApprovalForAll(&_FlowSubscription.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _tokenBaseURI) returns()
func (_FlowSubscription *FlowSubscriptionTransactor) SetBaseURI(opts *bind.TransactOpts, _tokenBaseURI string) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "setBaseURI", _tokenBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _tokenBaseURI) returns()
func (_FlowSubscription *FlowSubscriptionSession) SetBaseURI(_tokenBaseURI string) (*types.Transaction, error) {
	return _FlowSubscription.Contract.SetBaseURI(&_FlowSubscription.TransactOpts, _tokenBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _tokenBaseURI) returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) SetBaseURI(_tokenBaseURI string) (*types.Transaction, error) {
	return _FlowSubscription.Contract.SetBaseURI(&_FlowSubscription.TransactOpts, _tokenBaseURI)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _publicSalePrice) returns()
func (_FlowSubscription *FlowSubscriptionTransactor) SetPrice(opts *bind.TransactOpts, _publicSalePrice *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "setPrice", _publicSalePrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _publicSalePrice) returns()
func (_FlowSubscription *FlowSubscriptionSession) SetPrice(_publicSalePrice *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.Contract.SetPrice(&_FlowSubscription.TransactOpts, _publicSalePrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _publicSalePrice) returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) SetPrice(_publicSalePrice *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.Contract.SetPrice(&_FlowSubscription.TransactOpts, _publicSalePrice)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string _tokenUri) returns()
func (_FlowSubscription *FlowSubscriptionTransactor) SetTokenURI(opts *bind.TransactOpts, tokenId *big.Int, _tokenUri string) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "setTokenURI", tokenId, _tokenUri)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string _tokenUri) returns()
func (_FlowSubscription *FlowSubscriptionSession) SetTokenURI(tokenId *big.Int, _tokenUri string) (*types.Transaction, error) {
	return _FlowSubscription.Contract.SetTokenURI(&_FlowSubscription.TransactOpts, tokenId, _tokenUri)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string _tokenUri) returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) SetTokenURI(tokenId *big.Int, _tokenUri string) (*types.Transaction, error) {
	return _FlowSubscription.Contract.SetTokenURI(&_FlowSubscription.TransactOpts, tokenId, _tokenUri)
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() payable returns(uint256)
func (_FlowSubscription *FlowSubscriptionTransactor) Subscribe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "subscribe")
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() payable returns(uint256)
func (_FlowSubscription *FlowSubscriptionSession) Subscribe() (*types.Transaction, error) {
	return _FlowSubscription.Contract.Subscribe(&_FlowSubscription.TransactOpts)
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() payable returns(uint256)
func (_FlowSubscription *FlowSubscriptionTransactorSession) Subscribe() (*types.Transaction, error) {
	return _FlowSubscription.Contract.Subscribe(&_FlowSubscription.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FlowSubscription *FlowSubscriptionTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FlowSubscription *FlowSubscriptionSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.Contract.TransferFrom(&_FlowSubscription.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FlowSubscription.Contract.TransferFrom(&_FlowSubscription.TransactOpts, from, to, tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_FlowSubscription *FlowSubscriptionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_FlowSubscription *FlowSubscriptionSession) Unpause() (*types.Transaction, error) {
	return _FlowSubscription.Contract.Unpause(&_FlowSubscription.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) Unpause() (*types.Transaction, error) {
	return _FlowSubscription.Contract.Unpause(&_FlowSubscription.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FlowSubscription *FlowSubscriptionTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlowSubscription.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FlowSubscription *FlowSubscriptionSession) Withdraw() (*types.Transaction, error) {
	return _FlowSubscription.Contract.Withdraw(&_FlowSubscription.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FlowSubscription *FlowSubscriptionTransactorSession) Withdraw() (*types.Transaction, error) {
	return _FlowSubscription.Contract.Withdraw(&_FlowSubscription.TransactOpts)
}

// FlowSubscriptionApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FlowSubscription contract.
type FlowSubscriptionApprovalIterator struct {
	Event *FlowSubscriptionApproval // Event containing the contract specifics and raw log

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
func (it *FlowSubscriptionApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlowSubscriptionApproval)
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
		it.Event = new(FlowSubscriptionApproval)
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
func (it *FlowSubscriptionApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlowSubscriptionApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlowSubscriptionApproval represents a Approval event raised by the FlowSubscription contract.
type FlowSubscriptionApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FlowSubscription *FlowSubscriptionFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*FlowSubscriptionApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FlowSubscription.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FlowSubscriptionApprovalIterator{contract: _FlowSubscription.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FlowSubscription *FlowSubscriptionFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FlowSubscriptionApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FlowSubscription.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlowSubscriptionApproval)
				if err := _FlowSubscription.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FlowSubscription *FlowSubscriptionFilterer) ParseApproval(log types.Log) (*FlowSubscriptionApproval, error) {
	event := new(FlowSubscriptionApproval)
	if err := _FlowSubscription.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlowSubscriptionApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the FlowSubscription contract.
type FlowSubscriptionApprovalForAllIterator struct {
	Event *FlowSubscriptionApprovalForAll // Event containing the contract specifics and raw log

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
func (it *FlowSubscriptionApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlowSubscriptionApprovalForAll)
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
		it.Event = new(FlowSubscriptionApprovalForAll)
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
func (it *FlowSubscriptionApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlowSubscriptionApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlowSubscriptionApprovalForAll represents a ApprovalForAll event raised by the FlowSubscription contract.
type FlowSubscriptionApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FlowSubscription *FlowSubscriptionFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*FlowSubscriptionApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FlowSubscription.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &FlowSubscriptionApprovalForAllIterator{contract: _FlowSubscription.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FlowSubscription *FlowSubscriptionFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *FlowSubscriptionApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FlowSubscription.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlowSubscriptionApprovalForAll)
				if err := _FlowSubscription.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FlowSubscription *FlowSubscriptionFilterer) ParseApprovalForAll(log types.Log) (*FlowSubscriptionApprovalForAll, error) {
	event := new(FlowSubscriptionApprovalForAll)
	if err := _FlowSubscription.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlowSubscriptionRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the FlowSubscription contract.
type FlowSubscriptionRoleAdminChangedIterator struct {
	Event *FlowSubscriptionRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *FlowSubscriptionRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlowSubscriptionRoleAdminChanged)
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
		it.Event = new(FlowSubscriptionRoleAdminChanged)
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
func (it *FlowSubscriptionRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlowSubscriptionRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlowSubscriptionRoleAdminChanged represents a RoleAdminChanged event raised by the FlowSubscription contract.
type FlowSubscriptionRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FlowSubscription *FlowSubscriptionFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FlowSubscriptionRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _FlowSubscription.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FlowSubscriptionRoleAdminChangedIterator{contract: _FlowSubscription.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FlowSubscription *FlowSubscriptionFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FlowSubscriptionRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _FlowSubscription.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlowSubscriptionRoleAdminChanged)
				if err := _FlowSubscription.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FlowSubscription *FlowSubscriptionFilterer) ParseRoleAdminChanged(log types.Log) (*FlowSubscriptionRoleAdminChanged, error) {
	event := new(FlowSubscriptionRoleAdminChanged)
	if err := _FlowSubscription.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlowSubscriptionRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the FlowSubscription contract.
type FlowSubscriptionRoleGrantedIterator struct {
	Event *FlowSubscriptionRoleGranted // Event containing the contract specifics and raw log

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
func (it *FlowSubscriptionRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlowSubscriptionRoleGranted)
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
		it.Event = new(FlowSubscriptionRoleGranted)
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
func (it *FlowSubscriptionRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlowSubscriptionRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlowSubscriptionRoleGranted represents a RoleGranted event raised by the FlowSubscription contract.
type FlowSubscriptionRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FlowSubscription *FlowSubscriptionFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FlowSubscriptionRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FlowSubscription.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FlowSubscriptionRoleGrantedIterator{contract: _FlowSubscription.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FlowSubscription *FlowSubscriptionFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FlowSubscriptionRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FlowSubscription.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlowSubscriptionRoleGranted)
				if err := _FlowSubscription.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FlowSubscription *FlowSubscriptionFilterer) ParseRoleGranted(log types.Log) (*FlowSubscriptionRoleGranted, error) {
	event := new(FlowSubscriptionRoleGranted)
	if err := _FlowSubscription.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlowSubscriptionRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the FlowSubscription contract.
type FlowSubscriptionRoleRevokedIterator struct {
	Event *FlowSubscriptionRoleRevoked // Event containing the contract specifics and raw log

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
func (it *FlowSubscriptionRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlowSubscriptionRoleRevoked)
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
		it.Event = new(FlowSubscriptionRoleRevoked)
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
func (it *FlowSubscriptionRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlowSubscriptionRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlowSubscriptionRoleRevoked represents a RoleRevoked event raised by the FlowSubscription contract.
type FlowSubscriptionRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FlowSubscription *FlowSubscriptionFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FlowSubscriptionRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FlowSubscription.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FlowSubscriptionRoleRevokedIterator{contract: _FlowSubscription.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FlowSubscription *FlowSubscriptionFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FlowSubscriptionRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FlowSubscription.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlowSubscriptionRoleRevoked)
				if err := _FlowSubscription.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FlowSubscription *FlowSubscriptionFilterer) ParseRoleRevoked(log types.Log) (*FlowSubscriptionRoleRevoked, error) {
	event := new(FlowSubscriptionRoleRevoked)
	if err := _FlowSubscription.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlowSubscriptionSubscriptionCancelRequestedIterator is returned from FilterSubscriptionCancelRequested and is used to iterate over the raw logs and unpacked data for SubscriptionCancelRequested events raised by the FlowSubscription contract.
type FlowSubscriptionSubscriptionCancelRequestedIterator struct {
	Event *FlowSubscriptionSubscriptionCancelRequested // Event containing the contract specifics and raw log

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
func (it *FlowSubscriptionSubscriptionCancelRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlowSubscriptionSubscriptionCancelRequested)
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
		it.Event = new(FlowSubscriptionSubscriptionCancelRequested)
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
func (it *FlowSubscriptionSubscriptionCancelRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlowSubscriptionSubscriptionCancelRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlowSubscriptionSubscriptionCancelRequested represents a SubscriptionCancelRequested event raised by the FlowSubscription contract.
type FlowSubscriptionSubscriptionCancelRequested struct {
	TokenId *big.Int
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionCancelRequested is a free log retrieval operation binding the contract event 0x59019954c1cf4fce9b3bea5af48057bbfcfe96716860eb3e2b8ced18cc805f47.
//
// Solidity: event SubscriptionCancelRequested(uint256 indexed tokenId, uint256 indexed Time)
func (_FlowSubscription *FlowSubscriptionFilterer) FilterSubscriptionCancelRequested(opts *bind.FilterOpts, tokenId []*big.Int, Time []*big.Int) (*FlowSubscriptionSubscriptionCancelRequestedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var TimeRule []interface{}
	for _, TimeItem := range Time {
		TimeRule = append(TimeRule, TimeItem)
	}

	logs, sub, err := _FlowSubscription.contract.FilterLogs(opts, "SubscriptionCancelRequested", tokenIdRule, TimeRule)
	if err != nil {
		return nil, err
	}
	return &FlowSubscriptionSubscriptionCancelRequestedIterator{contract: _FlowSubscription.contract, event: "SubscriptionCancelRequested", logs: logs, sub: sub}, nil
}

// WatchSubscriptionCancelRequested is a free log subscription operation binding the contract event 0x59019954c1cf4fce9b3bea5af48057bbfcfe96716860eb3e2b8ced18cc805f47.
//
// Solidity: event SubscriptionCancelRequested(uint256 indexed tokenId, uint256 indexed Time)
func (_FlowSubscription *FlowSubscriptionFilterer) WatchSubscriptionCancelRequested(opts *bind.WatchOpts, sink chan<- *FlowSubscriptionSubscriptionCancelRequested, tokenId []*big.Int, Time []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var TimeRule []interface{}
	for _, TimeItem := range Time {
		TimeRule = append(TimeRule, TimeItem)
	}

	logs, sub, err := _FlowSubscription.contract.WatchLogs(opts, "SubscriptionCancelRequested", tokenIdRule, TimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlowSubscriptionSubscriptionCancelRequested)
				if err := _FlowSubscription.contract.UnpackLog(event, "SubscriptionCancelRequested", log); err != nil {
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

// ParseSubscriptionCancelRequested is a log parse operation binding the contract event 0x59019954c1cf4fce9b3bea5af48057bbfcfe96716860eb3e2b8ced18cc805f47.
//
// Solidity: event SubscriptionCancelRequested(uint256 indexed tokenId, uint256 indexed Time)
func (_FlowSubscription *FlowSubscriptionFilterer) ParseSubscriptionCancelRequested(log types.Log) (*FlowSubscriptionSubscriptionCancelRequested, error) {
	event := new(FlowSubscriptionSubscriptionCancelRequested)
	if err := _FlowSubscription.contract.UnpackLog(event, "SubscriptionCancelRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlowSubscriptionSubscriptionIssuedIterator is returned from FilterSubscriptionIssued and is used to iterate over the raw logs and unpacked data for SubscriptionIssued events raised by the FlowSubscription contract.
type FlowSubscriptionSubscriptionIssuedIterator struct {
	Event *FlowSubscriptionSubscriptionIssued // Event containing the contract specifics and raw log

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
func (it *FlowSubscriptionSubscriptionIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlowSubscriptionSubscriptionIssued)
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
		it.Event = new(FlowSubscriptionSubscriptionIssued)
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
func (it *FlowSubscriptionSubscriptionIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlowSubscriptionSubscriptionIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlowSubscriptionSubscriptionIssued represents a SubscriptionIssued event raised by the FlowSubscription contract.
type FlowSubscriptionSubscriptionIssued struct {
	TokenId *big.Int
	Owner   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionIssued is a free log retrieval operation binding the contract event 0xe0c7a574c27509120ebe3cdaa5048236ebd7d58e553021eef7b9e5090a1b8ef5.
//
// Solidity: event SubscriptionIssued(uint256 tokenId, address indexed owner)
func (_FlowSubscription *FlowSubscriptionFilterer) FilterSubscriptionIssued(opts *bind.FilterOpts, owner []common.Address) (*FlowSubscriptionSubscriptionIssuedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FlowSubscription.contract.FilterLogs(opts, "SubscriptionIssued", ownerRule)
	if err != nil {
		return nil, err
	}
	return &FlowSubscriptionSubscriptionIssuedIterator{contract: _FlowSubscription.contract, event: "SubscriptionIssued", logs: logs, sub: sub}, nil
}

// WatchSubscriptionIssued is a free log subscription operation binding the contract event 0xe0c7a574c27509120ebe3cdaa5048236ebd7d58e553021eef7b9e5090a1b8ef5.
//
// Solidity: event SubscriptionIssued(uint256 tokenId, address indexed owner)
func (_FlowSubscription *FlowSubscriptionFilterer) WatchSubscriptionIssued(opts *bind.WatchOpts, sink chan<- *FlowSubscriptionSubscriptionIssued, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FlowSubscription.contract.WatchLogs(opts, "SubscriptionIssued", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlowSubscriptionSubscriptionIssued)
				if err := _FlowSubscription.contract.UnpackLog(event, "SubscriptionIssued", log); err != nil {
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

// ParseSubscriptionIssued is a log parse operation binding the contract event 0xe0c7a574c27509120ebe3cdaa5048236ebd7d58e553021eef7b9e5090a1b8ef5.
//
// Solidity: event SubscriptionIssued(uint256 tokenId, address indexed owner)
func (_FlowSubscription *FlowSubscriptionFilterer) ParseSubscriptionIssued(log types.Log) (*FlowSubscriptionSubscriptionIssued, error) {
	event := new(FlowSubscriptionSubscriptionIssued)
	if err := _FlowSubscription.contract.UnpackLog(event, "SubscriptionIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlowSubscriptionSubscriptionRevokedIterator is returned from FilterSubscriptionRevoked and is used to iterate over the raw logs and unpacked data for SubscriptionRevoked events raised by the FlowSubscription contract.
type FlowSubscriptionSubscriptionRevokedIterator struct {
	Event *FlowSubscriptionSubscriptionRevoked // Event containing the contract specifics and raw log

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
func (it *FlowSubscriptionSubscriptionRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlowSubscriptionSubscriptionRevoked)
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
		it.Event = new(FlowSubscriptionSubscriptionRevoked)
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
func (it *FlowSubscriptionSubscriptionRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlowSubscriptionSubscriptionRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlowSubscriptionSubscriptionRevoked represents a SubscriptionRevoked event raised by the FlowSubscription contract.
type FlowSubscriptionSubscriptionRevoked struct {
	TokenId         *big.Int
	OwnerOrApproved common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionRevoked is a free log retrieval operation binding the contract event 0x35d33492e57db62bbca25dc09ec8f2becf2079f985d313c9e69957d786145453.
//
// Solidity: event SubscriptionRevoked(uint256 tokenId, address indexed ownerOrApproved)
func (_FlowSubscription *FlowSubscriptionFilterer) FilterSubscriptionRevoked(opts *bind.FilterOpts, ownerOrApproved []common.Address) (*FlowSubscriptionSubscriptionRevokedIterator, error) {

	var ownerOrApprovedRule []interface{}
	for _, ownerOrApprovedItem := range ownerOrApproved {
		ownerOrApprovedRule = append(ownerOrApprovedRule, ownerOrApprovedItem)
	}

	logs, sub, err := _FlowSubscription.contract.FilterLogs(opts, "SubscriptionRevoked", ownerOrApprovedRule)
	if err != nil {
		return nil, err
	}
	return &FlowSubscriptionSubscriptionRevokedIterator{contract: _FlowSubscription.contract, event: "SubscriptionRevoked", logs: logs, sub: sub}, nil
}

// WatchSubscriptionRevoked is a free log subscription operation binding the contract event 0x35d33492e57db62bbca25dc09ec8f2becf2079f985d313c9e69957d786145453.
//
// Solidity: event SubscriptionRevoked(uint256 tokenId, address indexed ownerOrApproved)
func (_FlowSubscription *FlowSubscriptionFilterer) WatchSubscriptionRevoked(opts *bind.WatchOpts, sink chan<- *FlowSubscriptionSubscriptionRevoked, ownerOrApproved []common.Address) (event.Subscription, error) {

	var ownerOrApprovedRule []interface{}
	for _, ownerOrApprovedItem := range ownerOrApproved {
		ownerOrApprovedRule = append(ownerOrApprovedRule, ownerOrApprovedItem)
	}

	logs, sub, err := _FlowSubscription.contract.WatchLogs(opts, "SubscriptionRevoked", ownerOrApprovedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlowSubscriptionSubscriptionRevoked)
				if err := _FlowSubscription.contract.UnpackLog(event, "SubscriptionRevoked", log); err != nil {
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

// ParseSubscriptionRevoked is a log parse operation binding the contract event 0x35d33492e57db62bbca25dc09ec8f2becf2079f985d313c9e69957d786145453.
//
// Solidity: event SubscriptionRevoked(uint256 tokenId, address indexed ownerOrApproved)
func (_FlowSubscription *FlowSubscriptionFilterer) ParseSubscriptionRevoked(log types.Log) (*FlowSubscriptionSubscriptionRevoked, error) {
	event := new(FlowSubscriptionSubscriptionRevoked)
	if err := _FlowSubscription.contract.UnpackLog(event, "SubscriptionRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlowSubscriptionSubscriptionUpdateIterator is returned from FilterSubscriptionUpdate and is used to iterate over the raw logs and unpacked data for SubscriptionUpdate events raised by the FlowSubscription contract.
type FlowSubscriptionSubscriptionUpdateIterator struct {
	Event *FlowSubscriptionSubscriptionUpdate // Event containing the contract specifics and raw log

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
func (it *FlowSubscriptionSubscriptionUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlowSubscriptionSubscriptionUpdate)
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
		it.Event = new(FlowSubscriptionSubscriptionUpdate)
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
func (it *FlowSubscriptionSubscriptionUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlowSubscriptionSubscriptionUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlowSubscriptionSubscriptionUpdate represents a SubscriptionUpdate event raised by the FlowSubscription contract.
type FlowSubscriptionSubscriptionUpdate struct {
	TokenId    *big.Int
	Expiration uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionUpdate is a free log retrieval operation binding the contract event 0x2ec2be2c4b90c2cf13ecb6751a24daed6bb741ae5ed3f7371aabf9402f6d62e8.
//
// Solidity: event SubscriptionUpdate(uint256 indexed tokenId, uint64 expiration)
func (_FlowSubscription *FlowSubscriptionFilterer) FilterSubscriptionUpdate(opts *bind.FilterOpts, tokenId []*big.Int) (*FlowSubscriptionSubscriptionUpdateIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FlowSubscription.contract.FilterLogs(opts, "SubscriptionUpdate", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FlowSubscriptionSubscriptionUpdateIterator{contract: _FlowSubscription.contract, event: "SubscriptionUpdate", logs: logs, sub: sub}, nil
}

// WatchSubscriptionUpdate is a free log subscription operation binding the contract event 0x2ec2be2c4b90c2cf13ecb6751a24daed6bb741ae5ed3f7371aabf9402f6d62e8.
//
// Solidity: event SubscriptionUpdate(uint256 indexed tokenId, uint64 expiration)
func (_FlowSubscription *FlowSubscriptionFilterer) WatchSubscriptionUpdate(opts *bind.WatchOpts, sink chan<- *FlowSubscriptionSubscriptionUpdate, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FlowSubscription.contract.WatchLogs(opts, "SubscriptionUpdate", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlowSubscriptionSubscriptionUpdate)
				if err := _FlowSubscription.contract.UnpackLog(event, "SubscriptionUpdate", log); err != nil {
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

// ParseSubscriptionUpdate is a log parse operation binding the contract event 0x2ec2be2c4b90c2cf13ecb6751a24daed6bb741ae5ed3f7371aabf9402f6d62e8.
//
// Solidity: event SubscriptionUpdate(uint256 indexed tokenId, uint64 expiration)
func (_FlowSubscription *FlowSubscriptionFilterer) ParseSubscriptionUpdate(log types.Log) (*FlowSubscriptionSubscriptionUpdate, error) {
	event := new(FlowSubscriptionSubscriptionUpdate)
	if err := _FlowSubscription.contract.UnpackLog(event, "SubscriptionUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlowSubscriptionTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FlowSubscription contract.
type FlowSubscriptionTransferIterator struct {
	Event *FlowSubscriptionTransfer // Event containing the contract specifics and raw log

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
func (it *FlowSubscriptionTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlowSubscriptionTransfer)
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
		it.Event = new(FlowSubscriptionTransfer)
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
func (it *FlowSubscriptionTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlowSubscriptionTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlowSubscriptionTransfer represents a Transfer event raised by the FlowSubscription contract.
type FlowSubscriptionTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FlowSubscription *FlowSubscriptionFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*FlowSubscriptionTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FlowSubscription.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FlowSubscriptionTransferIterator{contract: _FlowSubscription.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FlowSubscription *FlowSubscriptionFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FlowSubscriptionTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FlowSubscription.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlowSubscriptionTransfer)
				if err := _FlowSubscription.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FlowSubscription *FlowSubscriptionFilterer) ParseTransfer(log types.Log) (*FlowSubscriptionTransfer, error) {
	event := new(FlowSubscriptionTransfer)
	if err := _FlowSubscription.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
