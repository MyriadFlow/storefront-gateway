// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accessMaster

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

// AccessMasterMetaData contains all meta data concerning the AccessMaster contract.
var AccessMasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLOW_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLOW_CREATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLOW_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isCreator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"updateName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"updateSymbol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AccessMasterABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessMasterMetaData.ABI instead.
var AccessMasterABI = AccessMasterMetaData.ABI

// AccessMaster is an auto generated Go binding around an Ethereum contract.
type AccessMaster struct {
	AccessMasterCaller     // Read-only binding to the contract
	AccessMasterTransactor // Write-only binding to the contract
	AccessMasterFilterer   // Log filterer for contract events
}

// AccessMasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessMasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessMasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessMasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessMasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessMasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessMasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessMasterSession struct {
	Contract     *AccessMaster     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessMasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessMasterCallerSession struct {
	Contract *AccessMasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AccessMasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessMasterTransactorSession struct {
	Contract     *AccessMasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AccessMasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessMasterRaw struct {
	Contract *AccessMaster // Generic contract binding to access the raw methods on
}

// AccessMasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessMasterCallerRaw struct {
	Contract *AccessMasterCaller // Generic read-only contract binding to access the raw methods on
}

// AccessMasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessMasterTransactorRaw struct {
	Contract *AccessMasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessMaster creates a new instance of AccessMaster, bound to a specific deployed contract.
func NewAccessMaster(address common.Address, backend bind.ContractBackend) (*AccessMaster, error) {
	contract, err := bindAccessMaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessMaster{AccessMasterCaller: AccessMasterCaller{contract: contract}, AccessMasterTransactor: AccessMasterTransactor{contract: contract}, AccessMasterFilterer: AccessMasterFilterer{contract: contract}}, nil
}

// NewAccessMasterCaller creates a new read-only instance of AccessMaster, bound to a specific deployed contract.
func NewAccessMasterCaller(address common.Address, caller bind.ContractCaller) (*AccessMasterCaller, error) {
	contract, err := bindAccessMaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessMasterCaller{contract: contract}, nil
}

// NewAccessMasterTransactor creates a new write-only instance of AccessMaster, bound to a specific deployed contract.
func NewAccessMasterTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessMasterTransactor, error) {
	contract, err := bindAccessMaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessMasterTransactor{contract: contract}, nil
}

// NewAccessMasterFilterer creates a new log filterer instance of AccessMaster, bound to a specific deployed contract.
func NewAccessMasterFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessMasterFilterer, error) {
	contract, err := bindAccessMaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessMasterFilterer{contract: contract}, nil
}

// bindAccessMaster binds a generic wrapper to an already deployed contract.
func bindAccessMaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccessMasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessMaster *AccessMasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessMaster.Contract.AccessMasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessMaster *AccessMasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessMaster.Contract.AccessMasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessMaster *AccessMasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessMaster.Contract.AccessMasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessMaster *AccessMasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessMaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessMaster *AccessMasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessMaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessMaster *AccessMasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessMaster.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessMaster *AccessMasterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessMaster.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessMaster *AccessMasterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessMaster.Contract.DEFAULTADMINROLE(&_AccessMaster.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessMaster *AccessMasterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessMaster.Contract.DEFAULTADMINROLE(&_AccessMaster.CallOpts)
}

// FLOWADMINROLE is a free data retrieval call binding the contract method 0x7b75083c.
//
// Solidity: function FLOW_ADMIN_ROLE() view returns(bytes32)
func (_AccessMaster *AccessMasterCaller) FLOWADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessMaster.contract.Call(opts, &out, "FLOW_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FLOWADMINROLE is a free data retrieval call binding the contract method 0x7b75083c.
//
// Solidity: function FLOW_ADMIN_ROLE() view returns(bytes32)
func (_AccessMaster *AccessMasterSession) FLOWADMINROLE() ([32]byte, error) {
	return _AccessMaster.Contract.FLOWADMINROLE(&_AccessMaster.CallOpts)
}

// FLOWADMINROLE is a free data retrieval call binding the contract method 0x7b75083c.
//
// Solidity: function FLOW_ADMIN_ROLE() view returns(bytes32)
func (_AccessMaster *AccessMasterCallerSession) FLOWADMINROLE() ([32]byte, error) {
	return _AccessMaster.Contract.FLOWADMINROLE(&_AccessMaster.CallOpts)
}

// FLOWCREATORROLE is a free data retrieval call binding the contract method 0x9ea5722b.
//
// Solidity: function FLOW_CREATOR_ROLE() view returns(bytes32)
func (_AccessMaster *AccessMasterCaller) FLOWCREATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessMaster.contract.Call(opts, &out, "FLOW_CREATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FLOWCREATORROLE is a free data retrieval call binding the contract method 0x9ea5722b.
//
// Solidity: function FLOW_CREATOR_ROLE() view returns(bytes32)
func (_AccessMaster *AccessMasterSession) FLOWCREATORROLE() ([32]byte, error) {
	return _AccessMaster.Contract.FLOWCREATORROLE(&_AccessMaster.CallOpts)
}

// FLOWCREATORROLE is a free data retrieval call binding the contract method 0x9ea5722b.
//
// Solidity: function FLOW_CREATOR_ROLE() view returns(bytes32)
func (_AccessMaster *AccessMasterCallerSession) FLOWCREATORROLE() ([32]byte, error) {
	return _AccessMaster.Contract.FLOWCREATORROLE(&_AccessMaster.CallOpts)
}

// FLOWOPERATORROLE is a free data retrieval call binding the contract method 0x611ed8cb.
//
// Solidity: function FLOW_OPERATOR_ROLE() view returns(bytes32)
func (_AccessMaster *AccessMasterCaller) FLOWOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessMaster.contract.Call(opts, &out, "FLOW_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FLOWOPERATORROLE is a free data retrieval call binding the contract method 0x611ed8cb.
//
// Solidity: function FLOW_OPERATOR_ROLE() view returns(bytes32)
func (_AccessMaster *AccessMasterSession) FLOWOPERATORROLE() ([32]byte, error) {
	return _AccessMaster.Contract.FLOWOPERATORROLE(&_AccessMaster.CallOpts)
}

// FLOWOPERATORROLE is a free data retrieval call binding the contract method 0x611ed8cb.
//
// Solidity: function FLOW_OPERATOR_ROLE() view returns(bytes32)
func (_AccessMaster *AccessMasterCallerSession) FLOWOPERATORROLE() ([32]byte, error) {
	return _AccessMaster.Contract.FLOWOPERATORROLE(&_AccessMaster.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessMaster *AccessMasterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccessMaster.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessMaster *AccessMasterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessMaster.Contract.GetRoleAdmin(&_AccessMaster.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessMaster *AccessMasterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessMaster.Contract.GetRoleAdmin(&_AccessMaster.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessMaster *AccessMasterCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccessMaster.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessMaster *AccessMasterSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessMaster.Contract.GetRoleMember(&_AccessMaster.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessMaster *AccessMasterCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessMaster.Contract.GetRoleMember(&_AccessMaster.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessMaster *AccessMasterCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AccessMaster.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessMaster *AccessMasterSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessMaster.Contract.GetRoleMemberCount(&_AccessMaster.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessMaster *AccessMasterCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessMaster.Contract.GetRoleMemberCount(&_AccessMaster.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessMaster *AccessMasterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessMaster.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessMaster *AccessMasterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessMaster.Contract.HasRole(&_AccessMaster.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessMaster *AccessMasterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessMaster.Contract.HasRole(&_AccessMaster.CallOpts, role, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address user) view returns(bool)
func (_AccessMaster *AccessMasterCaller) IsAdmin(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _AccessMaster.contract.Call(opts, &out, "isAdmin", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address user) view returns(bool)
func (_AccessMaster *AccessMasterSession) IsAdmin(user common.Address) (bool, error) {
	return _AccessMaster.Contract.IsAdmin(&_AccessMaster.CallOpts, user)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address user) view returns(bool)
func (_AccessMaster *AccessMasterCallerSession) IsAdmin(user common.Address) (bool, error) {
	return _AccessMaster.Contract.IsAdmin(&_AccessMaster.CallOpts, user)
}

// IsCreator is a free data retrieval call binding the contract method 0xefd46065.
//
// Solidity: function isCreator(address user) view returns(bool)
func (_AccessMaster *AccessMasterCaller) IsCreator(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _AccessMaster.contract.Call(opts, &out, "isCreator", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCreator is a free data retrieval call binding the contract method 0xefd46065.
//
// Solidity: function isCreator(address user) view returns(bool)
func (_AccessMaster *AccessMasterSession) IsCreator(user common.Address) (bool, error) {
	return _AccessMaster.Contract.IsCreator(&_AccessMaster.CallOpts, user)
}

// IsCreator is a free data retrieval call binding the contract method 0xefd46065.
//
// Solidity: function isCreator(address user) view returns(bool)
func (_AccessMaster *AccessMasterCallerSession) IsCreator(user common.Address) (bool, error) {
	return _AccessMaster.Contract.IsCreator(&_AccessMaster.CallOpts, user)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address user) view returns(bool)
func (_AccessMaster *AccessMasterCaller) IsOperator(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _AccessMaster.contract.Call(opts, &out, "isOperator", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address user) view returns(bool)
func (_AccessMaster *AccessMasterSession) IsOperator(user common.Address) (bool, error) {
	return _AccessMaster.Contract.IsOperator(&_AccessMaster.CallOpts, user)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address user) view returns(bool)
func (_AccessMaster *AccessMasterCallerSession) IsOperator(user common.Address) (bool, error) {
	return _AccessMaster.Contract.IsOperator(&_AccessMaster.CallOpts, user)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AccessMaster *AccessMasterCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AccessMaster.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AccessMaster *AccessMasterSession) Name() (string, error) {
	return _AccessMaster.Contract.Name(&_AccessMaster.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AccessMaster *AccessMasterCallerSession) Name() (string, error) {
	return _AccessMaster.Contract.Name(&_AccessMaster.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessMaster *AccessMasterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AccessMaster.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessMaster *AccessMasterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessMaster.Contract.SupportsInterface(&_AccessMaster.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessMaster *AccessMasterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessMaster.Contract.SupportsInterface(&_AccessMaster.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AccessMaster *AccessMasterCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AccessMaster.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AccessMaster *AccessMasterSession) Symbol() (string, error) {
	return _AccessMaster.Contract.Symbol(&_AccessMaster.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AccessMaster *AccessMasterCallerSession) Symbol() (string, error) {
	return _AccessMaster.Contract.Symbol(&_AccessMaster.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_AccessMaster *AccessMasterCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessMaster.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_AccessMaster *AccessMasterSession) Version() (uint8, error) {
	return _AccessMaster.Contract.Version(&_AccessMaster.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_AccessMaster *AccessMasterCallerSession) Version() (uint8, error) {
	return _AccessMaster.Contract.Version(&_AccessMaster.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessMaster *AccessMasterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessMaster.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessMaster *AccessMasterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessMaster.Contract.GrantRole(&_AccessMaster.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessMaster *AccessMasterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessMaster.Contract.GrantRole(&_AccessMaster.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessMaster *AccessMasterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessMaster.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessMaster *AccessMasterSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessMaster.Contract.RenounceRole(&_AccessMaster.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessMaster *AccessMasterTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessMaster.Contract.RenounceRole(&_AccessMaster.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessMaster *AccessMasterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessMaster.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessMaster *AccessMasterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessMaster.Contract.RevokeRole(&_AccessMaster.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessMaster *AccessMasterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessMaster.Contract.RevokeRole(&_AccessMaster.TransactOpts, role, account)
}

// UpdateName is a paid mutator transaction binding the contract method 0x84da92a7.
//
// Solidity: function updateName(string _name) returns()
func (_AccessMaster *AccessMasterTransactor) UpdateName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _AccessMaster.contract.Transact(opts, "updateName", _name)
}

// UpdateName is a paid mutator transaction binding the contract method 0x84da92a7.
//
// Solidity: function updateName(string _name) returns()
func (_AccessMaster *AccessMasterSession) UpdateName(_name string) (*types.Transaction, error) {
	return _AccessMaster.Contract.UpdateName(&_AccessMaster.TransactOpts, _name)
}

// UpdateName is a paid mutator transaction binding the contract method 0x84da92a7.
//
// Solidity: function updateName(string _name) returns()
func (_AccessMaster *AccessMasterTransactorSession) UpdateName(_name string) (*types.Transaction, error) {
	return _AccessMaster.Contract.UpdateName(&_AccessMaster.TransactOpts, _name)
}

// UpdateSymbol is a paid mutator transaction binding the contract method 0x537f5312.
//
// Solidity: function updateSymbol(string _symbol) returns()
func (_AccessMaster *AccessMasterTransactor) UpdateSymbol(opts *bind.TransactOpts, _symbol string) (*types.Transaction, error) {
	return _AccessMaster.contract.Transact(opts, "updateSymbol", _symbol)
}

// UpdateSymbol is a paid mutator transaction binding the contract method 0x537f5312.
//
// Solidity: function updateSymbol(string _symbol) returns()
func (_AccessMaster *AccessMasterSession) UpdateSymbol(_symbol string) (*types.Transaction, error) {
	return _AccessMaster.Contract.UpdateSymbol(&_AccessMaster.TransactOpts, _symbol)
}

// UpdateSymbol is a paid mutator transaction binding the contract method 0x537f5312.
//
// Solidity: function updateSymbol(string _symbol) returns()
func (_AccessMaster *AccessMasterTransactorSession) UpdateSymbol(_symbol string) (*types.Transaction, error) {
	return _AccessMaster.Contract.UpdateSymbol(&_AccessMaster.TransactOpts, _symbol)
}

// AccessMasterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AccessMaster contract.
type AccessMasterRoleAdminChangedIterator struct {
	Event *AccessMasterRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccessMasterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessMasterRoleAdminChanged)
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
		it.Event = new(AccessMasterRoleAdminChanged)
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
func (it *AccessMasterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessMasterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessMasterRoleAdminChanged represents a RoleAdminChanged event raised by the AccessMaster contract.
type AccessMasterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessMaster *AccessMasterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AccessMasterRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AccessMaster.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AccessMasterRoleAdminChangedIterator{contract: _AccessMaster.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessMaster *AccessMasterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessMasterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AccessMaster.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessMasterRoleAdminChanged)
				if err := _AccessMaster.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AccessMaster *AccessMasterFilterer) ParseRoleAdminChanged(log types.Log) (*AccessMasterRoleAdminChanged, error) {
	event := new(AccessMasterRoleAdminChanged)
	if err := _AccessMaster.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessMasterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessMaster contract.
type AccessMasterRoleGrantedIterator struct {
	Event *AccessMasterRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccessMasterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessMasterRoleGranted)
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
		it.Event = new(AccessMasterRoleGranted)
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
func (it *AccessMasterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessMasterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessMasterRoleGranted represents a RoleGranted event raised by the AccessMaster contract.
type AccessMasterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessMaster *AccessMasterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessMasterRoleGrantedIterator, error) {

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

	logs, sub, err := _AccessMaster.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessMasterRoleGrantedIterator{contract: _AccessMaster.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessMaster *AccessMasterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessMasterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessMaster.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessMasterRoleGranted)
				if err := _AccessMaster.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AccessMaster *AccessMasterFilterer) ParseRoleGranted(log types.Log) (*AccessMasterRoleGranted, error) {
	event := new(AccessMasterRoleGranted)
	if err := _AccessMaster.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessMasterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessMaster contract.
type AccessMasterRoleRevokedIterator struct {
	Event *AccessMasterRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccessMasterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessMasterRoleRevoked)
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
		it.Event = new(AccessMasterRoleRevoked)
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
func (it *AccessMasterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessMasterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessMasterRoleRevoked represents a RoleRevoked event raised by the AccessMaster contract.
type AccessMasterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessMaster *AccessMasterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessMasterRoleRevokedIterator, error) {

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

	logs, sub, err := _AccessMaster.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessMasterRoleRevokedIterator{contract: _AccessMaster.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessMaster *AccessMasterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessMasterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessMaster.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessMasterRoleRevoked)
				if err := _AccessMaster.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AccessMaster *AccessMasterFilterer) ParseRoleRevoked(log types.Log) (*AccessMasterRoleRevoked, error) {
	event := new(AccessMasterRoleRevoked)
	if err := _AccessMaster.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
