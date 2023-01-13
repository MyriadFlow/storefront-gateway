// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketplace

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
)

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_platformFee\",\"type\":\"uint96\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metaDataURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"forSale\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"activity\",\"type\":\"string\"}],\"name\":\"MarketplaceItem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MARKETPLACE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"newPlatformFee\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"newPayoutAddress\",\"type\":\"address\"}],\"name\":\"changeFeeAndPayoutAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"createMarketItem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"}],\"name\":\"createMarketSale\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"idToMarketItem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forSale\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"deleted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payoutAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFeeBasisPoint\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"}],\"name\":\"removeFromSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Marketplace *MarketplaceCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Marketplace *MarketplaceSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Marketplace.Contract.DEFAULTADMINROLE(&_Marketplace.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Marketplace.Contract.DEFAULTADMINROLE(&_Marketplace.CallOpts)
}

// MARKETPLACEADMINROLE is a free data retrieval call binding the contract method 0xb2a4eea0.
//
// Solidity: function MARKETPLACE_ADMIN_ROLE() view returns(bytes32)
func (_Marketplace *MarketplaceCaller) MARKETPLACEADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "MARKETPLACE_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MARKETPLACEADMINROLE is a free data retrieval call binding the contract method 0xb2a4eea0.
//
// Solidity: function MARKETPLACE_ADMIN_ROLE() view returns(bytes32)
func (_Marketplace *MarketplaceSession) MARKETPLACEADMINROLE() ([32]byte, error) {
	return _Marketplace.Contract.MARKETPLACEADMINROLE(&_Marketplace.CallOpts)
}

// MARKETPLACEADMINROLE is a free data retrieval call binding the contract method 0xb2a4eea0.
//
// Solidity: function MARKETPLACE_ADMIN_ROLE() view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) MARKETPLACEADMINROLE() ([32]byte, error) {
	return _Marketplace.Contract.MARKETPLACEADMINROLE(&_Marketplace.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Marketplace *MarketplaceCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Marketplace *MarketplaceSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Marketplace.Contract.GetRoleAdmin(&_Marketplace.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Marketplace.Contract.GetRoleAdmin(&_Marketplace.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Marketplace *MarketplaceCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Marketplace *MarketplaceSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Marketplace.Contract.GetRoleMember(&_Marketplace.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Marketplace *MarketplaceCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Marketplace.Contract.GetRoleMember(&_Marketplace.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Marketplace *MarketplaceSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Marketplace.Contract.GetRoleMemberCount(&_Marketplace.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Marketplace.Contract.GetRoleMemberCount(&_Marketplace.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Marketplace *MarketplaceCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Marketplace *MarketplaceSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Marketplace.Contract.HasRole(&_Marketplace.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Marketplace.Contract.HasRole(&_Marketplace.CallOpts, role, account)
}

// IdToMarketItem is a free data retrieval call binding the contract method 0xe61a70c0.
//
// Solidity: function idToMarketItem(uint256 ) view returns(uint256 itemId, address nftContract, uint256 tokenId, address seller, address owner, uint256 price, bool forSale, bool deleted)
func (_Marketplace *MarketplaceCaller) IdToMarketItem(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	ForSale     bool
	Deleted     bool
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "idToMarketItem", arg0)

	outstruct := new(struct {
		ItemId      *big.Int
		NftContract common.Address
		TokenId     *big.Int
		Seller      common.Address
		Owner       common.Address
		Price       *big.Int
		ForSale     bool
		Deleted     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ItemId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NftContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Owner = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ForSale = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Deleted = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// IdToMarketItem is a free data retrieval call binding the contract method 0xe61a70c0.
//
// Solidity: function idToMarketItem(uint256 ) view returns(uint256 itemId, address nftContract, uint256 tokenId, address seller, address owner, uint256 price, bool forSale, bool deleted)
func (_Marketplace *MarketplaceSession) IdToMarketItem(arg0 *big.Int) (struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	ForSale     bool
	Deleted     bool
}, error) {
	return _Marketplace.Contract.IdToMarketItem(&_Marketplace.CallOpts, arg0)
}

// IdToMarketItem is a free data retrieval call binding the contract method 0xe61a70c0.
//
// Solidity: function idToMarketItem(uint256 ) view returns(uint256 itemId, address nftContract, uint256 tokenId, address seller, address owner, uint256 price, bool forSale, bool deleted)
func (_Marketplace *MarketplaceCallerSession) IdToMarketItem(arg0 *big.Int) (struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	ForSale     bool
	Deleted     bool
}, error) {
	return _Marketplace.Contract.IdToMarketItem(&_Marketplace.CallOpts, arg0)
}

// PayoutAddress is a free data retrieval call binding the contract method 0x5b8d02d7.
//
// Solidity: function payoutAddress() view returns(address)
func (_Marketplace *MarketplaceCaller) PayoutAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "payoutAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PayoutAddress is a free data retrieval call binding the contract method 0x5b8d02d7.
//
// Solidity: function payoutAddress() view returns(address)
func (_Marketplace *MarketplaceSession) PayoutAddress() (common.Address, error) {
	return _Marketplace.Contract.PayoutAddress(&_Marketplace.CallOpts)
}

// PayoutAddress is a free data retrieval call binding the contract method 0x5b8d02d7.
//
// Solidity: function payoutAddress() view returns(address)
func (_Marketplace *MarketplaceCallerSession) PayoutAddress() (common.Address, error) {
	return _Marketplace.Contract.PayoutAddress(&_Marketplace.CallOpts)
}

// PlatformFeeBasisPoint is a free data retrieval call binding the contract method 0x5a9cd033.
//
// Solidity: function platformFeeBasisPoint() view returns(uint96)
func (_Marketplace *MarketplaceCaller) PlatformFeeBasisPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "platformFeeBasisPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFeeBasisPoint is a free data retrieval call binding the contract method 0x5a9cd033.
//
// Solidity: function platformFeeBasisPoint() view returns(uint96)
func (_Marketplace *MarketplaceSession) PlatformFeeBasisPoint() (*big.Int, error) {
	return _Marketplace.Contract.PlatformFeeBasisPoint(&_Marketplace.CallOpts)
}

// PlatformFeeBasisPoint is a free data retrieval call binding the contract method 0x5a9cd033.
//
// Solidity: function platformFeeBasisPoint() view returns(uint96)
func (_Marketplace *MarketplaceCallerSession) PlatformFeeBasisPoint() (*big.Int, error) {
	return _Marketplace.Contract.PlatformFeeBasisPoint(&_Marketplace.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Marketplace *MarketplaceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Marketplace *MarketplaceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Marketplace.Contract.SupportsInterface(&_Marketplace.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Marketplace.Contract.SupportsInterface(&_Marketplace.CallOpts, interfaceId)
}

// ChangeFeeAndPayoutAddress is a paid mutator transaction binding the contract method 0x3738b55a.
//
// Solidity: function changeFeeAndPayoutAddress(uint96 newPlatformFee, address newPayoutAddress) returns()
func (_Marketplace *MarketplaceTransactor) ChangeFeeAndPayoutAddress(opts *bind.TransactOpts, newPlatformFee *big.Int, newPayoutAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "changeFeeAndPayoutAddress", newPlatformFee, newPayoutAddress)
}

// ChangeFeeAndPayoutAddress is a paid mutator transaction binding the contract method 0x3738b55a.
//
// Solidity: function changeFeeAndPayoutAddress(uint96 newPlatformFee, address newPayoutAddress) returns()
func (_Marketplace *MarketplaceSession) ChangeFeeAndPayoutAddress(newPlatformFee *big.Int, newPayoutAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.ChangeFeeAndPayoutAddress(&_Marketplace.TransactOpts, newPlatformFee, newPayoutAddress)
}

// ChangeFeeAndPayoutAddress is a paid mutator transaction binding the contract method 0x3738b55a.
//
// Solidity: function changeFeeAndPayoutAddress(uint96 newPlatformFee, address newPayoutAddress) returns()
func (_Marketplace *MarketplaceTransactorSession) ChangeFeeAndPayoutAddress(newPlatformFee *big.Int, newPayoutAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.ChangeFeeAndPayoutAddress(&_Marketplace.TransactOpts, newPlatformFee, newPayoutAddress)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) returns(uint256)
func (_Marketplace *MarketplaceTransactor) CreateMarketItem(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "createMarketItem", nftContract, tokenId, price)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) returns(uint256)
func (_Marketplace *MarketplaceSession) CreateMarketItem(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateMarketItem(&_Marketplace.TransactOpts, nftContract, tokenId, price)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) returns(uint256)
func (_Marketplace *MarketplaceTransactorSession) CreateMarketItem(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateMarketItem(&_Marketplace.TransactOpts, nftContract, tokenId, price)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xbe9af536.
//
// Solidity: function createMarketSale(uint256 itemId) payable returns()
func (_Marketplace *MarketplaceTransactor) CreateMarketSale(opts *bind.TransactOpts, itemId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "createMarketSale", itemId)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xbe9af536.
//
// Solidity: function createMarketSale(uint256 itemId) payable returns()
func (_Marketplace *MarketplaceSession) CreateMarketSale(itemId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateMarketSale(&_Marketplace.TransactOpts, itemId)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xbe9af536.
//
// Solidity: function createMarketSale(uint256 itemId) payable returns()
func (_Marketplace *MarketplaceTransactorSession) CreateMarketSale(itemId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateMarketSale(&_Marketplace.TransactOpts, itemId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.GrantRole(&_Marketplace.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.GrantRole(&_Marketplace.TransactOpts, role, account)
}

// RemoveFromSale is a paid mutator transaction binding the contract method 0x1361a3b6.
//
// Solidity: function removeFromSale(uint256 itemId) returns()
func (_Marketplace *MarketplaceTransactor) RemoveFromSale(opts *bind.TransactOpts, itemId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "removeFromSale", itemId)
}

// RemoveFromSale is a paid mutator transaction binding the contract method 0x1361a3b6.
//
// Solidity: function removeFromSale(uint256 itemId) returns()
func (_Marketplace *MarketplaceSession) RemoveFromSale(itemId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.RemoveFromSale(&_Marketplace.TransactOpts, itemId)
}

// RemoveFromSale is a paid mutator transaction binding the contract method 0x1361a3b6.
//
// Solidity: function removeFromSale(uint256 itemId) returns()
func (_Marketplace *MarketplaceTransactorSession) RemoveFromSale(itemId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.RemoveFromSale(&_Marketplace.TransactOpts, itemId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.RenounceRole(&_Marketplace.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.RenounceRole(&_Marketplace.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.RevokeRole(&_Marketplace.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.RevokeRole(&_Marketplace.TransactOpts, role, account)
}

// MarketplaceMarketplaceItemIterator is returned from FilterMarketplaceItem and is used to iterate over the raw logs and unpacked data for MarketplaceItem events raised by the Marketplace contract.
type MarketplaceMarketplaceItemIterator struct {
	Event *MarketplaceMarketplaceItem // Event containing the contract specifics and raw log

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
func (it *MarketplaceMarketplaceItemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMarketplaceItem)
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
		it.Event = new(MarketplaceMarketplaceItem)
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
func (it *MarketplaceMarketplaceItemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMarketplaceItemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMarketplaceItem represents a MarketplaceItem event raised by the Marketplace contract.
type MarketplaceMarketplaceItem struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	MetaDataURI string
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	ForSale     bool
	Activity    string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketplaceItem is a free log retrieval operation binding the contract event 0x441ce8810cb23039fd1298c2fe1e7b82e2886a738cfc3281515eedfb7f0e60af.
//
// Solidity: event MarketplaceItem(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, string metaDataURI, address seller, address owner, uint256 price, bool forSale, string activity)
func (_Marketplace *MarketplaceFilterer) FilterMarketplaceItem(opts *bind.FilterOpts, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (*MarketplaceMarketplaceItemIterator, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "MarketplaceItem", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMarketplaceItemIterator{contract: _Marketplace.contract, event: "MarketplaceItem", logs: logs, sub: sub}, nil
}

// WatchMarketplaceItem is a free log subscription operation binding the contract event 0x441ce8810cb23039fd1298c2fe1e7b82e2886a738cfc3281515eedfb7f0e60af.
//
// Solidity: event MarketplaceItem(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, string metaDataURI, address seller, address owner, uint256 price, bool forSale, string activity)
func (_Marketplace *MarketplaceFilterer) WatchMarketplaceItem(opts *bind.WatchOpts, sink chan<- *MarketplaceMarketplaceItem, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "MarketplaceItem", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMarketplaceItem)
				if err := _Marketplace.contract.UnpackLog(event, "MarketplaceItem", log); err != nil {
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

// ParseMarketplaceItem is a log parse operation binding the contract event 0x441ce8810cb23039fd1298c2fe1e7b82e2886a738cfc3281515eedfb7f0e60af.
//
// Solidity: event MarketplaceItem(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, string metaDataURI, address seller, address owner, uint256 price, bool forSale, string activity)
func (_Marketplace *MarketplaceFilterer) ParseMarketplaceItem(log types.Log) (*MarketplaceMarketplaceItem, error) {
	event := new(MarketplaceMarketplaceItem)
	if err := _Marketplace.contract.UnpackLog(event, "MarketplaceItem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Marketplace contract.
type MarketplaceRoleAdminChangedIterator struct {
	Event *MarketplaceRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MarketplaceRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceRoleAdminChanged)
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
		it.Event = new(MarketplaceRoleAdminChanged)
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
func (it *MarketplaceRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceRoleAdminChanged represents a RoleAdminChanged event raised by the Marketplace contract.
type MarketplaceRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Marketplace *MarketplaceFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MarketplaceRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceRoleAdminChangedIterator{contract: _Marketplace.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Marketplace *MarketplaceFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MarketplaceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceRoleAdminChanged)
				if err := _Marketplace.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseRoleAdminChanged(log types.Log) (*MarketplaceRoleAdminChanged, error) {
	event := new(MarketplaceRoleAdminChanged)
	if err := _Marketplace.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Marketplace contract.
type MarketplaceRoleGrantedIterator struct {
	Event *MarketplaceRoleGranted // Event containing the contract specifics and raw log

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
func (it *MarketplaceRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceRoleGranted)
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
		it.Event = new(MarketplaceRoleGranted)
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
func (it *MarketplaceRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceRoleGranted represents a RoleGranted event raised by the Marketplace contract.
type MarketplaceRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marketplace *MarketplaceFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MarketplaceRoleGrantedIterator, error) {

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

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceRoleGrantedIterator{contract: _Marketplace.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marketplace *MarketplaceFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MarketplaceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceRoleGranted)
				if err := _Marketplace.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseRoleGranted(log types.Log) (*MarketplaceRoleGranted, error) {
	event := new(MarketplaceRoleGranted)
	if err := _Marketplace.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Marketplace contract.
type MarketplaceRoleRevokedIterator struct {
	Event *MarketplaceRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MarketplaceRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceRoleRevoked)
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
		it.Event = new(MarketplaceRoleRevoked)
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
func (it *MarketplaceRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceRoleRevoked represents a RoleRevoked event raised by the Marketplace contract.
type MarketplaceRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marketplace *MarketplaceFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MarketplaceRoleRevokedIterator, error) {

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

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceRoleRevokedIterator{contract: _Marketplace.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marketplace *MarketplaceFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MarketplaceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceRoleRevoked)
				if err := _Marketplace.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseRoleRevoked(log types.Log) (*MarketplaceRoleRevoked, error) {
	event := new(MarketplaceRoleRevoked)
	if err := _Marketplace.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
