// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package signatureSeries

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

// SignatureSeriesLazyNFTVoucher is an auto generated low-level Go binding around an user-defined struct.
type SignatureSeriesLazyNFTVoucher struct {
	Price     *big.Int
	Uri       string
	Signature []byte
}

// SignatureSeriesMetaData contains all meta data concerning the SignatureSeries contract.
var SignatureSeriesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_nftPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tradeHubAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"flowContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"reciepient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRentable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"renter\",\"type\":\"address\"}],\"name\":\"RentalInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metaDataURI\",\"type\":\"string\"}],\"name\":\"SignatureSeriesAssetCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ownerOrApproved\",\"type\":\"address\"}],\"name\":\"SignatureSeriesAssetDestroyed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expires\",\"type\":\"uint64\"}],\"name\":\"UpdateUser\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Counter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNING_DOMAIN\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessMasterAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"amountRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"uint96\",\"name\":\"royaltyPercentBasisPoint\",\"type\":\"uint96\"}],\"name\":\"createAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"uint96\",\"name\":\"royaltyPercentBasisPoint\",\"type\":\"uint96\"}],\"name\":\"delegateAssetCreation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"destroyAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSignatureSeries.LazyNFTVoucher\",\"name\":\"voucher\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"royaltyPercentBasisPoint\",\"type\":\"uint96\"}],\"name\":\"lazyAssetCreation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSignatureSeries.LazyNFTVoucher\",\"name\":\"voucher\",\"type\":\"tuple\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeInHours\",\"type\":\"uint256\"}],\"name\":\"rent\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rentables\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRentable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expires\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"hourlyRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftPrice\",\"type\":\"uint256\"}],\"name\":\"setNftPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRentable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pricePerHour\",\"type\":\"uint256\"}],\"name\":\"setRentInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expires\",\"type\":\"uint64\"}],\"name\":\"setUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradeHub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"userExpires\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"userOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SignatureSeriesABI is the input ABI used to generate the binding from.
// Deprecated: Use SignatureSeriesMetaData.ABI instead.
var SignatureSeriesABI = SignatureSeriesMetaData.ABI

// SignatureSeries is an auto generated Go binding around an Ethereum contract.
type SignatureSeries struct {
	SignatureSeriesCaller     // Read-only binding to the contract
	SignatureSeriesTransactor // Write-only binding to the contract
	SignatureSeriesFilterer   // Log filterer for contract events
}

// SignatureSeriesCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignatureSeriesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureSeriesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignatureSeriesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureSeriesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignatureSeriesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureSeriesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignatureSeriesSession struct {
	Contract     *SignatureSeries  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignatureSeriesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignatureSeriesCallerSession struct {
	Contract *SignatureSeriesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SignatureSeriesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignatureSeriesTransactorSession struct {
	Contract     *SignatureSeriesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SignatureSeriesRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignatureSeriesRaw struct {
	Contract *SignatureSeries // Generic contract binding to access the raw methods on
}

// SignatureSeriesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignatureSeriesCallerRaw struct {
	Contract *SignatureSeriesCaller // Generic read-only contract binding to access the raw methods on
}

// SignatureSeriesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignatureSeriesTransactorRaw struct {
	Contract *SignatureSeriesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignatureSeries creates a new instance of SignatureSeries, bound to a specific deployed contract.
func NewSignatureSeries(address common.Address, backend bind.ContractBackend) (*SignatureSeries, error) {
	contract, err := bindSignatureSeries(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignatureSeries{SignatureSeriesCaller: SignatureSeriesCaller{contract: contract}, SignatureSeriesTransactor: SignatureSeriesTransactor{contract: contract}, SignatureSeriesFilterer: SignatureSeriesFilterer{contract: contract}}, nil
}

// NewSignatureSeriesCaller creates a new read-only instance of SignatureSeries, bound to a specific deployed contract.
func NewSignatureSeriesCaller(address common.Address, caller bind.ContractCaller) (*SignatureSeriesCaller, error) {
	contract, err := bindSignatureSeries(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureSeriesCaller{contract: contract}, nil
}

// NewSignatureSeriesTransactor creates a new write-only instance of SignatureSeries, bound to a specific deployed contract.
func NewSignatureSeriesTransactor(address common.Address, transactor bind.ContractTransactor) (*SignatureSeriesTransactor, error) {
	contract, err := bindSignatureSeries(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureSeriesTransactor{contract: contract}, nil
}

// NewSignatureSeriesFilterer creates a new log filterer instance of SignatureSeries, bound to a specific deployed contract.
func NewSignatureSeriesFilterer(address common.Address, filterer bind.ContractFilterer) (*SignatureSeriesFilterer, error) {
	contract, err := bindSignatureSeries(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignatureSeriesFilterer{contract: contract}, nil
}

// bindSignatureSeries binds a generic wrapper to an already deployed contract.
func bindSignatureSeries(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SignatureSeriesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureSeries *SignatureSeriesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignatureSeries.Contract.SignatureSeriesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureSeries *SignatureSeriesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureSeries.Contract.SignatureSeriesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureSeries *SignatureSeriesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureSeries.Contract.SignatureSeriesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureSeries *SignatureSeriesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignatureSeries.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureSeries *SignatureSeriesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureSeries.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureSeries *SignatureSeriesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureSeries.Contract.contract.Transact(opts, method, params...)
}

// Counter is a free data retrieval call binding the contract method 0x4eeca076.
//
// Solidity: function Counter() view returns(uint256)
func (_SignatureSeries *SignatureSeriesCaller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "Counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x4eeca076.
//
// Solidity: function Counter() view returns(uint256)
func (_SignatureSeries *SignatureSeriesSession) Counter() (*big.Int, error) {
	return _SignatureSeries.Contract.Counter(&_SignatureSeries.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x4eeca076.
//
// Solidity: function Counter() view returns(uint256)
func (_SignatureSeries *SignatureSeriesCallerSession) Counter() (*big.Int, error) {
	return _SignatureSeries.Contract.Counter(&_SignatureSeries.CallOpts)
}

// SIGNATUREVERSION is a free data retrieval call binding the contract method 0xa4157296.
//
// Solidity: function SIGNATURE_VERSION() view returns(string)
func (_SignatureSeries *SignatureSeriesCaller) SIGNATUREVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "SIGNATURE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SIGNATUREVERSION is a free data retrieval call binding the contract method 0xa4157296.
//
// Solidity: function SIGNATURE_VERSION() view returns(string)
func (_SignatureSeries *SignatureSeriesSession) SIGNATUREVERSION() (string, error) {
	return _SignatureSeries.Contract.SIGNATUREVERSION(&_SignatureSeries.CallOpts)
}

// SIGNATUREVERSION is a free data retrieval call binding the contract method 0xa4157296.
//
// Solidity: function SIGNATURE_VERSION() view returns(string)
func (_SignatureSeries *SignatureSeriesCallerSession) SIGNATUREVERSION() (string, error) {
	return _SignatureSeries.Contract.SIGNATUREVERSION(&_SignatureSeries.CallOpts)
}

// SIGNINGDOMAIN is a free data retrieval call binding the contract method 0xe3faad94.
//
// Solidity: function SIGNING_DOMAIN() view returns(string)
func (_SignatureSeries *SignatureSeriesCaller) SIGNINGDOMAIN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "SIGNING_DOMAIN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SIGNINGDOMAIN is a free data retrieval call binding the contract method 0xe3faad94.
//
// Solidity: function SIGNING_DOMAIN() view returns(string)
func (_SignatureSeries *SignatureSeriesSession) SIGNINGDOMAIN() (string, error) {
	return _SignatureSeries.Contract.SIGNINGDOMAIN(&_SignatureSeries.CallOpts)
}

// SIGNINGDOMAIN is a free data retrieval call binding the contract method 0xe3faad94.
//
// Solidity: function SIGNING_DOMAIN() view returns(string)
func (_SignatureSeries *SignatureSeriesCallerSession) SIGNINGDOMAIN() (string, error) {
	return _SignatureSeries.Contract.SIGNINGDOMAIN(&_SignatureSeries.CallOpts)
}

// AccessMasterAddress is a free data retrieval call binding the contract method 0x0aa7c854.
//
// Solidity: function accessMasterAddress() view returns(address)
func (_SignatureSeries *SignatureSeriesCaller) AccessMasterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "accessMasterAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccessMasterAddress is a free data retrieval call binding the contract method 0x0aa7c854.
//
// Solidity: function accessMasterAddress() view returns(address)
func (_SignatureSeries *SignatureSeriesSession) AccessMasterAddress() (common.Address, error) {
	return _SignatureSeries.Contract.AccessMasterAddress(&_SignatureSeries.CallOpts)
}

// AccessMasterAddress is a free data retrieval call binding the contract method 0x0aa7c854.
//
// Solidity: function accessMasterAddress() view returns(address)
func (_SignatureSeries *SignatureSeriesCallerSession) AccessMasterAddress() (common.Address, error) {
	return _SignatureSeries.Contract.AccessMasterAddress(&_SignatureSeries.CallOpts)
}

// AmountRequired is a free data retrieval call binding the contract method 0xba91caae.
//
// Solidity: function amountRequired(uint256 tokenId, uint256 time) view returns(uint256 amount)
func (_SignatureSeries *SignatureSeriesCaller) AmountRequired(opts *bind.CallOpts, tokenId *big.Int, time *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "amountRequired", tokenId, time)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountRequired is a free data retrieval call binding the contract method 0xba91caae.
//
// Solidity: function amountRequired(uint256 tokenId, uint256 time) view returns(uint256 amount)
func (_SignatureSeries *SignatureSeriesSession) AmountRequired(tokenId *big.Int, time *big.Int) (*big.Int, error) {
	return _SignatureSeries.Contract.AmountRequired(&_SignatureSeries.CallOpts, tokenId, time)
}

// AmountRequired is a free data retrieval call binding the contract method 0xba91caae.
//
// Solidity: function amountRequired(uint256 tokenId, uint256 time) view returns(uint256 amount)
func (_SignatureSeries *SignatureSeriesCallerSession) AmountRequired(tokenId *big.Int, time *big.Int) (*big.Int, error) {
	return _SignatureSeries.Contract.AmountRequired(&_SignatureSeries.CallOpts, tokenId, time)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SignatureSeries *SignatureSeriesCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SignatureSeries *SignatureSeriesSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SignatureSeries.Contract.BalanceOf(&_SignatureSeries.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SignatureSeries *SignatureSeriesCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SignatureSeries.Contract.BalanceOf(&_SignatureSeries.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SignatureSeries *SignatureSeriesCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SignatureSeries *SignatureSeriesSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SignatureSeries.Contract.GetApproved(&_SignatureSeries.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SignatureSeries *SignatureSeriesCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SignatureSeries.Contract.GetApproved(&_SignatureSeries.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SignatureSeries *SignatureSeriesCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SignatureSeries *SignatureSeriesSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SignatureSeries.Contract.IsApprovedForAll(&_SignatureSeries.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SignatureSeries *SignatureSeriesCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SignatureSeries.Contract.IsApprovedForAll(&_SignatureSeries.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SignatureSeries *SignatureSeriesCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SignatureSeries *SignatureSeriesSession) Name() (string, error) {
	return _SignatureSeries.Contract.Name(&_SignatureSeries.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SignatureSeries *SignatureSeriesCallerSession) Name() (string, error) {
	return _SignatureSeries.Contract.Name(&_SignatureSeries.CallOpts)
}

// NftPrice is a free data retrieval call binding the contract method 0x0d39fc81.
//
// Solidity: function nftPrice() view returns(uint256)
func (_SignatureSeries *SignatureSeriesCaller) NftPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "nftPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftPrice is a free data retrieval call binding the contract method 0x0d39fc81.
//
// Solidity: function nftPrice() view returns(uint256)
func (_SignatureSeries *SignatureSeriesSession) NftPrice() (*big.Int, error) {
	return _SignatureSeries.Contract.NftPrice(&_SignatureSeries.CallOpts)
}

// NftPrice is a free data retrieval call binding the contract method 0x0d39fc81.
//
// Solidity: function nftPrice() view returns(uint256)
func (_SignatureSeries *SignatureSeriesCallerSession) NftPrice() (*big.Int, error) {
	return _SignatureSeries.Contract.NftPrice(&_SignatureSeries.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SignatureSeries *SignatureSeriesCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SignatureSeries *SignatureSeriesSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SignatureSeries.Contract.OwnerOf(&_SignatureSeries.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SignatureSeries *SignatureSeriesCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SignatureSeries.Contract.OwnerOf(&_SignatureSeries.CallOpts, tokenId)
}

// Recover is a free data retrieval call binding the contract method 0x3a61d52c.
//
// Solidity: function recover((uint256,string,bytes) voucher) view returns(address)
func (_SignatureSeries *SignatureSeriesCaller) Recover(opts *bind.CallOpts, voucher SignatureSeriesLazyNFTVoucher) (common.Address, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "recover", voucher)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recover is a free data retrieval call binding the contract method 0x3a61d52c.
//
// Solidity: function recover((uint256,string,bytes) voucher) view returns(address)
func (_SignatureSeries *SignatureSeriesSession) Recover(voucher SignatureSeriesLazyNFTVoucher) (common.Address, error) {
	return _SignatureSeries.Contract.Recover(&_SignatureSeries.CallOpts, voucher)
}

// Recover is a free data retrieval call binding the contract method 0x3a61d52c.
//
// Solidity: function recover((uint256,string,bytes) voucher) view returns(address)
func (_SignatureSeries *SignatureSeriesCallerSession) Recover(voucher SignatureSeriesLazyNFTVoucher) (common.Address, error) {
	return _SignatureSeries.Contract.Recover(&_SignatureSeries.CallOpts, voucher)
}

// Rentables is a free data retrieval call binding the contract method 0x9874cd97.
//
// Solidity: function rentables(uint256 ) view returns(bool isRentable, address user, uint64 expires, uint256 hourlyRate)
func (_SignatureSeries *SignatureSeriesCaller) Rentables(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IsRentable bool
	User       common.Address
	Expires    uint64
	HourlyRate *big.Int
}, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "rentables", arg0)

	outstruct := new(struct {
		IsRentable bool
		User       common.Address
		Expires    uint64
		HourlyRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsRentable = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.User = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Expires = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.HourlyRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Rentables is a free data retrieval call binding the contract method 0x9874cd97.
//
// Solidity: function rentables(uint256 ) view returns(bool isRentable, address user, uint64 expires, uint256 hourlyRate)
func (_SignatureSeries *SignatureSeriesSession) Rentables(arg0 *big.Int) (struct {
	IsRentable bool
	User       common.Address
	Expires    uint64
	HourlyRate *big.Int
}, error) {
	return _SignatureSeries.Contract.Rentables(&_SignatureSeries.CallOpts, arg0)
}

// Rentables is a free data retrieval call binding the contract method 0x9874cd97.
//
// Solidity: function rentables(uint256 ) view returns(bool isRentable, address user, uint64 expires, uint256 hourlyRate)
func (_SignatureSeries *SignatureSeriesCallerSession) Rentables(arg0 *big.Int) (struct {
	IsRentable bool
	User       common.Address
	Expires    uint64
	HourlyRate *big.Int
}, error) {
	return _SignatureSeries.Contract.Rentables(&_SignatureSeries.CallOpts, arg0)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_SignatureSeries *SignatureSeriesCaller) RoyaltyInfo(opts *bind.CallOpts, _tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "royaltyInfo", _tokenId, _salePrice)

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
func (_SignatureSeries *SignatureSeriesSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _SignatureSeries.Contract.RoyaltyInfo(&_SignatureSeries.CallOpts, _tokenId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_SignatureSeries *SignatureSeriesCallerSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _SignatureSeries.Contract.RoyaltyInfo(&_SignatureSeries.CallOpts, _tokenId, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SignatureSeries *SignatureSeriesCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SignatureSeries *SignatureSeriesSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SignatureSeries.Contract.SupportsInterface(&_SignatureSeries.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SignatureSeries *SignatureSeriesCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SignatureSeries.Contract.SupportsInterface(&_SignatureSeries.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SignatureSeries *SignatureSeriesCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SignatureSeries *SignatureSeriesSession) Symbol() (string, error) {
	return _SignatureSeries.Contract.Symbol(&_SignatureSeries.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SignatureSeries *SignatureSeriesCallerSession) Symbol() (string, error) {
	return _SignatureSeries.Contract.Symbol(&_SignatureSeries.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_SignatureSeries *SignatureSeriesCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_SignatureSeries *SignatureSeriesSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _SignatureSeries.Contract.TokenByIndex(&_SignatureSeries.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_SignatureSeries *SignatureSeriesCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _SignatureSeries.Contract.TokenByIndex(&_SignatureSeries.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_SignatureSeries *SignatureSeriesCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_SignatureSeries *SignatureSeriesSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _SignatureSeries.Contract.TokenOfOwnerByIndex(&_SignatureSeries.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_SignatureSeries *SignatureSeriesCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _SignatureSeries.Contract.TokenOfOwnerByIndex(&_SignatureSeries.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SignatureSeries *SignatureSeriesCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SignatureSeries *SignatureSeriesSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SignatureSeries.Contract.TokenURI(&_SignatureSeries.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SignatureSeries *SignatureSeriesCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SignatureSeries.Contract.TokenURI(&_SignatureSeries.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SignatureSeries *SignatureSeriesCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SignatureSeries *SignatureSeriesSession) TotalSupply() (*big.Int, error) {
	return _SignatureSeries.Contract.TotalSupply(&_SignatureSeries.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SignatureSeries *SignatureSeriesCallerSession) TotalSupply() (*big.Int, error) {
	return _SignatureSeries.Contract.TotalSupply(&_SignatureSeries.CallOpts)
}

// TradeHub is a free data retrieval call binding the contract method 0xa4b11052.
//
// Solidity: function tradeHub() view returns(address)
func (_SignatureSeries *SignatureSeriesCaller) TradeHub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "tradeHub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TradeHub is a free data retrieval call binding the contract method 0xa4b11052.
//
// Solidity: function tradeHub() view returns(address)
func (_SignatureSeries *SignatureSeriesSession) TradeHub() (common.Address, error) {
	return _SignatureSeries.Contract.TradeHub(&_SignatureSeries.CallOpts)
}

// TradeHub is a free data retrieval call binding the contract method 0xa4b11052.
//
// Solidity: function tradeHub() view returns(address)
func (_SignatureSeries *SignatureSeriesCallerSession) TradeHub() (common.Address, error) {
	return _SignatureSeries.Contract.TradeHub(&_SignatureSeries.CallOpts)
}

// UserExpires is a free data retrieval call binding the contract method 0x8fc88c48.
//
// Solidity: function userExpires(uint256 tokenId) view returns(uint256)
func (_SignatureSeries *SignatureSeriesCaller) UserExpires(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "userExpires", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserExpires is a free data retrieval call binding the contract method 0x8fc88c48.
//
// Solidity: function userExpires(uint256 tokenId) view returns(uint256)
func (_SignatureSeries *SignatureSeriesSession) UserExpires(tokenId *big.Int) (*big.Int, error) {
	return _SignatureSeries.Contract.UserExpires(&_SignatureSeries.CallOpts, tokenId)
}

// UserExpires is a free data retrieval call binding the contract method 0x8fc88c48.
//
// Solidity: function userExpires(uint256 tokenId) view returns(uint256)
func (_SignatureSeries *SignatureSeriesCallerSession) UserExpires(tokenId *big.Int) (*big.Int, error) {
	return _SignatureSeries.Contract.UserExpires(&_SignatureSeries.CallOpts, tokenId)
}

// UserOf is a free data retrieval call binding the contract method 0xc2f1f14a.
//
// Solidity: function userOf(uint256 tokenId) view returns(address)
func (_SignatureSeries *SignatureSeriesCaller) UserOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "userOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserOf is a free data retrieval call binding the contract method 0xc2f1f14a.
//
// Solidity: function userOf(uint256 tokenId) view returns(address)
func (_SignatureSeries *SignatureSeriesSession) UserOf(tokenId *big.Int) (common.Address, error) {
	return _SignatureSeries.Contract.UserOf(&_SignatureSeries.CallOpts, tokenId)
}

// UserOf is a free data retrieval call binding the contract method 0xc2f1f14a.
//
// Solidity: function userOf(uint256 tokenId) view returns(address)
func (_SignatureSeries *SignatureSeriesCallerSession) UserOf(tokenId *big.Int) (common.Address, error) {
	return _SignatureSeries.Contract.UserOf(&_SignatureSeries.CallOpts, tokenId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_SignatureSeries *SignatureSeriesCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SignatureSeries.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_SignatureSeries *SignatureSeriesSession) Version() (uint8, error) {
	return _SignatureSeries.Contract.Version(&_SignatureSeries.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_SignatureSeries *SignatureSeriesCallerSession) Version() (uint8, error) {
	return _SignatureSeries.Contract.Version(&_SignatureSeries.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SignatureSeries *SignatureSeriesTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SignatureSeries *SignatureSeriesSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.Approve(&_SignatureSeries.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SignatureSeries *SignatureSeriesTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.Approve(&_SignatureSeries.TransactOpts, to, tokenId)
}

// CreateAsset is a paid mutator transaction binding the contract method 0x83c6f9a9.
//
// Solidity: function createAsset(string metadataURI, uint96 royaltyPercentBasisPoint) returns(uint256)
func (_SignatureSeries *SignatureSeriesTransactor) CreateAsset(opts *bind.TransactOpts, metadataURI string, royaltyPercentBasisPoint *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.contract.Transact(opts, "createAsset", metadataURI, royaltyPercentBasisPoint)
}

// CreateAsset is a paid mutator transaction binding the contract method 0x83c6f9a9.
//
// Solidity: function createAsset(string metadataURI, uint96 royaltyPercentBasisPoint) returns(uint256)
func (_SignatureSeries *SignatureSeriesSession) CreateAsset(metadataURI string, royaltyPercentBasisPoint *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.CreateAsset(&_SignatureSeries.TransactOpts, metadataURI, royaltyPercentBasisPoint)
}

// CreateAsset is a paid mutator transaction binding the contract method 0x83c6f9a9.
//
// Solidity: function createAsset(string metadataURI, uint96 royaltyPercentBasisPoint) returns(uint256)
func (_SignatureSeries *SignatureSeriesTransactorSession) CreateAsset(metadataURI string, royaltyPercentBasisPoint *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.CreateAsset(&_SignatureSeries.TransactOpts, metadataURI, royaltyPercentBasisPoint)
}

// DelegateAssetCreation is a paid mutator transaction binding the contract method 0x8c275eec.
//
// Solidity: function delegateAssetCreation(address creator, string metadataURI, uint96 royaltyPercentBasisPoint) returns(uint256)
func (_SignatureSeries *SignatureSeriesTransactor) DelegateAssetCreation(opts *bind.TransactOpts, creator common.Address, metadataURI string, royaltyPercentBasisPoint *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.contract.Transact(opts, "delegateAssetCreation", creator, metadataURI, royaltyPercentBasisPoint)
}

// DelegateAssetCreation is a paid mutator transaction binding the contract method 0x8c275eec.
//
// Solidity: function delegateAssetCreation(address creator, string metadataURI, uint96 royaltyPercentBasisPoint) returns(uint256)
func (_SignatureSeries *SignatureSeriesSession) DelegateAssetCreation(creator common.Address, metadataURI string, royaltyPercentBasisPoint *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.DelegateAssetCreation(&_SignatureSeries.TransactOpts, creator, metadataURI, royaltyPercentBasisPoint)
}

// DelegateAssetCreation is a paid mutator transaction binding the contract method 0x8c275eec.
//
// Solidity: function delegateAssetCreation(address creator, string metadataURI, uint96 royaltyPercentBasisPoint) returns(uint256)
func (_SignatureSeries *SignatureSeriesTransactorSession) DelegateAssetCreation(creator common.Address, metadataURI string, royaltyPercentBasisPoint *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.DelegateAssetCreation(&_SignatureSeries.TransactOpts, creator, metadataURI, royaltyPercentBasisPoint)
}

// DestroyAsset is a paid mutator transaction binding the contract method 0x474c1520.
//
// Solidity: function destroyAsset(uint256 tokenId) returns()
func (_SignatureSeries *SignatureSeriesTransactor) DestroyAsset(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.contract.Transact(opts, "destroyAsset", tokenId)
}

// DestroyAsset is a paid mutator transaction binding the contract method 0x474c1520.
//
// Solidity: function destroyAsset(uint256 tokenId) returns()
func (_SignatureSeries *SignatureSeriesSession) DestroyAsset(tokenId *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.DestroyAsset(&_SignatureSeries.TransactOpts, tokenId)
}

// DestroyAsset is a paid mutator transaction binding the contract method 0x474c1520.
//
// Solidity: function destroyAsset(uint256 tokenId) returns()
func (_SignatureSeries *SignatureSeriesTransactorSession) DestroyAsset(tokenId *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.DestroyAsset(&_SignatureSeries.TransactOpts, tokenId)
}

// LazyAssetCreation is a paid mutator transaction binding the contract method 0xdae48e6b.
//
// Solidity: function lazyAssetCreation((uint256,string,bytes) voucher, uint96 royaltyPercentBasisPoint) payable returns(uint256)
func (_SignatureSeries *SignatureSeriesTransactor) LazyAssetCreation(opts *bind.TransactOpts, voucher SignatureSeriesLazyNFTVoucher, royaltyPercentBasisPoint *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.contract.Transact(opts, "lazyAssetCreation", voucher, royaltyPercentBasisPoint)
}

// LazyAssetCreation is a paid mutator transaction binding the contract method 0xdae48e6b.
//
// Solidity: function lazyAssetCreation((uint256,string,bytes) voucher, uint96 royaltyPercentBasisPoint) payable returns(uint256)
func (_SignatureSeries *SignatureSeriesSession) LazyAssetCreation(voucher SignatureSeriesLazyNFTVoucher, royaltyPercentBasisPoint *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.LazyAssetCreation(&_SignatureSeries.TransactOpts, voucher, royaltyPercentBasisPoint)
}

// LazyAssetCreation is a paid mutator transaction binding the contract method 0xdae48e6b.
//
// Solidity: function lazyAssetCreation((uint256,string,bytes) voucher, uint96 royaltyPercentBasisPoint) payable returns(uint256)
func (_SignatureSeries *SignatureSeriesTransactorSession) LazyAssetCreation(voucher SignatureSeriesLazyNFTVoucher, royaltyPercentBasisPoint *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.LazyAssetCreation(&_SignatureSeries.TransactOpts, voucher, royaltyPercentBasisPoint)
}

// Rent is a paid mutator transaction binding the contract method 0x783a112b.
//
// Solidity: function rent(uint256 _tokenId, uint256 _timeInHours) payable returns()
func (_SignatureSeries *SignatureSeriesTransactor) Rent(opts *bind.TransactOpts, _tokenId *big.Int, _timeInHours *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.contract.Transact(opts, "rent", _tokenId, _timeInHours)
}

// Rent is a paid mutator transaction binding the contract method 0x783a112b.
//
// Solidity: function rent(uint256 _tokenId, uint256 _timeInHours) payable returns()
func (_SignatureSeries *SignatureSeriesSession) Rent(_tokenId *big.Int, _timeInHours *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.Rent(&_SignatureSeries.TransactOpts, _tokenId, _timeInHours)
}

// Rent is a paid mutator transaction binding the contract method 0x783a112b.
//
// Solidity: function rent(uint256 _tokenId, uint256 _timeInHours) payable returns()
func (_SignatureSeries *SignatureSeriesTransactorSession) Rent(_tokenId *big.Int, _timeInHours *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.Rent(&_SignatureSeries.TransactOpts, _tokenId, _timeInHours)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SignatureSeries *SignatureSeriesTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SignatureSeries *SignatureSeriesSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.SafeTransferFrom(&_SignatureSeries.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SignatureSeries *SignatureSeriesTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.SafeTransferFrom(&_SignatureSeries.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SignatureSeries *SignatureSeriesTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SignatureSeries.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SignatureSeries *SignatureSeriesSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SignatureSeries.Contract.SafeTransferFrom0(&_SignatureSeries.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SignatureSeries *SignatureSeriesTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SignatureSeries.Contract.SafeTransferFrom0(&_SignatureSeries.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SignatureSeries *SignatureSeriesTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _SignatureSeries.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SignatureSeries *SignatureSeriesSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SignatureSeries.Contract.SetApprovalForAll(&_SignatureSeries.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SignatureSeries *SignatureSeriesTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SignatureSeries.Contract.SetApprovalForAll(&_SignatureSeries.TransactOpts, operator, approved)
}

// SetNftPrice is a paid mutator transaction binding the contract method 0x7d9a7a4c.
//
// Solidity: function setNftPrice(uint256 _nftPrice) returns()
func (_SignatureSeries *SignatureSeriesTransactor) SetNftPrice(opts *bind.TransactOpts, _nftPrice *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.contract.Transact(opts, "setNftPrice", _nftPrice)
}

// SetNftPrice is a paid mutator transaction binding the contract method 0x7d9a7a4c.
//
// Solidity: function setNftPrice(uint256 _nftPrice) returns()
func (_SignatureSeries *SignatureSeriesSession) SetNftPrice(_nftPrice *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.SetNftPrice(&_SignatureSeries.TransactOpts, _nftPrice)
}

// SetNftPrice is a paid mutator transaction binding the contract method 0x7d9a7a4c.
//
// Solidity: function setNftPrice(uint256 _nftPrice) returns()
func (_SignatureSeries *SignatureSeriesTransactorSession) SetNftPrice(_nftPrice *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.SetNftPrice(&_SignatureSeries.TransactOpts, _nftPrice)
}

// SetRentInfo is a paid mutator transaction binding the contract method 0x655bc8db.
//
// Solidity: function setRentInfo(uint256 tokenId, bool isRentable, uint256 pricePerHour) returns()
func (_SignatureSeries *SignatureSeriesTransactor) SetRentInfo(opts *bind.TransactOpts, tokenId *big.Int, isRentable bool, pricePerHour *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.contract.Transact(opts, "setRentInfo", tokenId, isRentable, pricePerHour)
}

// SetRentInfo is a paid mutator transaction binding the contract method 0x655bc8db.
//
// Solidity: function setRentInfo(uint256 tokenId, bool isRentable, uint256 pricePerHour) returns()
func (_SignatureSeries *SignatureSeriesSession) SetRentInfo(tokenId *big.Int, isRentable bool, pricePerHour *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.SetRentInfo(&_SignatureSeries.TransactOpts, tokenId, isRentable, pricePerHour)
}

// SetRentInfo is a paid mutator transaction binding the contract method 0x655bc8db.
//
// Solidity: function setRentInfo(uint256 tokenId, bool isRentable, uint256 pricePerHour) returns()
func (_SignatureSeries *SignatureSeriesTransactorSession) SetRentInfo(tokenId *big.Int, isRentable bool, pricePerHour *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.SetRentInfo(&_SignatureSeries.TransactOpts, tokenId, isRentable, pricePerHour)
}

// SetUser is a paid mutator transaction binding the contract method 0xe030565e.
//
// Solidity: function setUser(uint256 tokenId, address user, uint64 expires) returns()
func (_SignatureSeries *SignatureSeriesTransactor) SetUser(opts *bind.TransactOpts, tokenId *big.Int, user common.Address, expires uint64) (*types.Transaction, error) {
	return _SignatureSeries.contract.Transact(opts, "setUser", tokenId, user, expires)
}

// SetUser is a paid mutator transaction binding the contract method 0xe030565e.
//
// Solidity: function setUser(uint256 tokenId, address user, uint64 expires) returns()
func (_SignatureSeries *SignatureSeriesSession) SetUser(tokenId *big.Int, user common.Address, expires uint64) (*types.Transaction, error) {
	return _SignatureSeries.Contract.SetUser(&_SignatureSeries.TransactOpts, tokenId, user, expires)
}

// SetUser is a paid mutator transaction binding the contract method 0xe030565e.
//
// Solidity: function setUser(uint256 tokenId, address user, uint64 expires) returns()
func (_SignatureSeries *SignatureSeriesTransactorSession) SetUser(tokenId *big.Int, user common.Address, expires uint64) (*types.Transaction, error) {
	return _SignatureSeries.Contract.SetUser(&_SignatureSeries.TransactOpts, tokenId, user, expires)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SignatureSeries *SignatureSeriesTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SignatureSeries *SignatureSeriesSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.TransferFrom(&_SignatureSeries.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SignatureSeries *SignatureSeriesTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SignatureSeries.Contract.TransferFrom(&_SignatureSeries.TransactOpts, from, to, tokenId)
}

// SignatureSeriesApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SignatureSeries contract.
type SignatureSeriesApprovalIterator struct {
	Event *SignatureSeriesApproval // Event containing the contract specifics and raw log

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
func (it *SignatureSeriesApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignatureSeriesApproval)
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
		it.Event = new(SignatureSeriesApproval)
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
func (it *SignatureSeriesApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignatureSeriesApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignatureSeriesApproval represents a Approval event raised by the SignatureSeries contract.
type SignatureSeriesApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SignatureSeries *SignatureSeriesFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SignatureSeriesApprovalIterator, error) {

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

	logs, sub, err := _SignatureSeries.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SignatureSeriesApprovalIterator{contract: _SignatureSeries.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SignatureSeries *SignatureSeriesFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SignatureSeriesApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _SignatureSeries.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignatureSeriesApproval)
				if err := _SignatureSeries.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_SignatureSeries *SignatureSeriesFilterer) ParseApproval(log types.Log) (*SignatureSeriesApproval, error) {
	event := new(SignatureSeriesApproval)
	if err := _SignatureSeries.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SignatureSeriesApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the SignatureSeries contract.
type SignatureSeriesApprovalForAllIterator struct {
	Event *SignatureSeriesApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SignatureSeriesApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignatureSeriesApprovalForAll)
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
		it.Event = new(SignatureSeriesApprovalForAll)
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
func (it *SignatureSeriesApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignatureSeriesApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignatureSeriesApprovalForAll represents a ApprovalForAll event raised by the SignatureSeries contract.
type SignatureSeriesApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SignatureSeries *SignatureSeriesFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SignatureSeriesApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SignatureSeries.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SignatureSeriesApprovalForAllIterator{contract: _SignatureSeries.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SignatureSeries *SignatureSeriesFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SignatureSeriesApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SignatureSeries.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignatureSeriesApprovalForAll)
				if err := _SignatureSeries.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_SignatureSeries *SignatureSeriesFilterer) ParseApprovalForAll(log types.Log) (*SignatureSeriesApprovalForAll, error) {
	event := new(SignatureSeriesApprovalForAll)
	if err := _SignatureSeries.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SignatureSeriesFundTransferredIterator is returned from FilterFundTransferred and is used to iterate over the raw logs and unpacked data for FundTransferred events raised by the SignatureSeries contract.
type SignatureSeriesFundTransferredIterator struct {
	Event *SignatureSeriesFundTransferred // Event containing the contract specifics and raw log

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
func (it *SignatureSeriesFundTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignatureSeriesFundTransferred)
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
		it.Event = new(SignatureSeriesFundTransferred)
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
func (it *SignatureSeriesFundTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignatureSeriesFundTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignatureSeriesFundTransferred represents a FundTransferred event raised by the SignatureSeries contract.
type SignatureSeriesFundTransferred struct {
	Sender     common.Address
	Reciepient common.Address
	TokenId    *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFundTransferred is a free log retrieval operation binding the contract event 0xaeadc2c673108f922a1042ef404b73ae1af992c9ae54249a4a7d13ef21ef2c18.
//
// Solidity: event FundTransferred(address sender, address reciepient, uint256 tokenId, uint256 amount)
func (_SignatureSeries *SignatureSeriesFilterer) FilterFundTransferred(opts *bind.FilterOpts) (*SignatureSeriesFundTransferredIterator, error) {

	logs, sub, err := _SignatureSeries.contract.FilterLogs(opts, "FundTransferred")
	if err != nil {
		return nil, err
	}
	return &SignatureSeriesFundTransferredIterator{contract: _SignatureSeries.contract, event: "FundTransferred", logs: logs, sub: sub}, nil
}

// WatchFundTransferred is a free log subscription operation binding the contract event 0xaeadc2c673108f922a1042ef404b73ae1af992c9ae54249a4a7d13ef21ef2c18.
//
// Solidity: event FundTransferred(address sender, address reciepient, uint256 tokenId, uint256 amount)
func (_SignatureSeries *SignatureSeriesFilterer) WatchFundTransferred(opts *bind.WatchOpts, sink chan<- *SignatureSeriesFundTransferred) (event.Subscription, error) {

	logs, sub, err := _SignatureSeries.contract.WatchLogs(opts, "FundTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignatureSeriesFundTransferred)
				if err := _SignatureSeries.contract.UnpackLog(event, "FundTransferred", log); err != nil {
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

// ParseFundTransferred is a log parse operation binding the contract event 0xaeadc2c673108f922a1042ef404b73ae1af992c9ae54249a4a7d13ef21ef2c18.
//
// Solidity: event FundTransferred(address sender, address reciepient, uint256 tokenId, uint256 amount)
func (_SignatureSeries *SignatureSeriesFilterer) ParseFundTransferred(log types.Log) (*SignatureSeriesFundTransferred, error) {
	event := new(SignatureSeriesFundTransferred)
	if err := _SignatureSeries.contract.UnpackLog(event, "FundTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SignatureSeriesRentalInfoIterator is returned from FilterRentalInfo and is used to iterate over the raw logs and unpacked data for RentalInfo events raised by the SignatureSeries contract.
type SignatureSeriesRentalInfoIterator struct {
	Event *SignatureSeriesRentalInfo // Event containing the contract specifics and raw log

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
func (it *SignatureSeriesRentalInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignatureSeriesRentalInfo)
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
		it.Event = new(SignatureSeriesRentalInfo)
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
func (it *SignatureSeriesRentalInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignatureSeriesRentalInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignatureSeriesRentalInfo represents a RentalInfo event raised by the SignatureSeries contract.
type SignatureSeriesRentalInfo struct {
	TokenId    *big.Int
	IsRentable bool
	Price      *big.Int
	Renter     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRentalInfo is a free log retrieval operation binding the contract event 0xa9dcc06084f66afbdc1841a28e84400fcbf607131220fdcf404c8502f3c03203.
//
// Solidity: event RentalInfo(uint256 tokenId, bool isRentable, uint256 price, address indexed renter)
func (_SignatureSeries *SignatureSeriesFilterer) FilterRentalInfo(opts *bind.FilterOpts, renter []common.Address) (*SignatureSeriesRentalInfoIterator, error) {

	var renterRule []interface{}
	for _, renterItem := range renter {
		renterRule = append(renterRule, renterItem)
	}

	logs, sub, err := _SignatureSeries.contract.FilterLogs(opts, "RentalInfo", renterRule)
	if err != nil {
		return nil, err
	}
	return &SignatureSeriesRentalInfoIterator{contract: _SignatureSeries.contract, event: "RentalInfo", logs: logs, sub: sub}, nil
}

// WatchRentalInfo is a free log subscription operation binding the contract event 0xa9dcc06084f66afbdc1841a28e84400fcbf607131220fdcf404c8502f3c03203.
//
// Solidity: event RentalInfo(uint256 tokenId, bool isRentable, uint256 price, address indexed renter)
func (_SignatureSeries *SignatureSeriesFilterer) WatchRentalInfo(opts *bind.WatchOpts, sink chan<- *SignatureSeriesRentalInfo, renter []common.Address) (event.Subscription, error) {

	var renterRule []interface{}
	for _, renterItem := range renter {
		renterRule = append(renterRule, renterItem)
	}

	logs, sub, err := _SignatureSeries.contract.WatchLogs(opts, "RentalInfo", renterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignatureSeriesRentalInfo)
				if err := _SignatureSeries.contract.UnpackLog(event, "RentalInfo", log); err != nil {
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

// ParseRentalInfo is a log parse operation binding the contract event 0xa9dcc06084f66afbdc1841a28e84400fcbf607131220fdcf404c8502f3c03203.
//
// Solidity: event RentalInfo(uint256 tokenId, bool isRentable, uint256 price, address indexed renter)
func (_SignatureSeries *SignatureSeriesFilterer) ParseRentalInfo(log types.Log) (*SignatureSeriesRentalInfo, error) {
	event := new(SignatureSeriesRentalInfo)
	if err := _SignatureSeries.contract.UnpackLog(event, "RentalInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SignatureSeriesSignatureSeriesAssetCreatedIterator is returned from FilterSignatureSeriesAssetCreated and is used to iterate over the raw logs and unpacked data for SignatureSeriesAssetCreated events raised by the SignatureSeries contract.
type SignatureSeriesSignatureSeriesAssetCreatedIterator struct {
	Event *SignatureSeriesSignatureSeriesAssetCreated // Event containing the contract specifics and raw log

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
func (it *SignatureSeriesSignatureSeriesAssetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignatureSeriesSignatureSeriesAssetCreated)
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
		it.Event = new(SignatureSeriesSignatureSeriesAssetCreated)
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
func (it *SignatureSeriesSignatureSeriesAssetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignatureSeriesSignatureSeriesAssetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignatureSeriesSignatureSeriesAssetCreated represents a SignatureSeriesAssetCreated event raised by the SignatureSeries contract.
type SignatureSeriesSignatureSeriesAssetCreated struct {
	TokenID     *big.Int
	Creator     common.Address
	MetaDataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSignatureSeriesAssetCreated is a free log retrieval operation binding the contract event 0x07ac23775381c7e58d18cef592a12cb5ea8576f2034c691e2500aba9b51f9d7f.
//
// Solidity: event SignatureSeriesAssetCreated(uint256 tokenID, address indexed creator, string metaDataURI)
func (_SignatureSeries *SignatureSeriesFilterer) FilterSignatureSeriesAssetCreated(opts *bind.FilterOpts, creator []common.Address) (*SignatureSeriesSignatureSeriesAssetCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _SignatureSeries.contract.FilterLogs(opts, "SignatureSeriesAssetCreated", creatorRule)
	if err != nil {
		return nil, err
	}
	return &SignatureSeriesSignatureSeriesAssetCreatedIterator{contract: _SignatureSeries.contract, event: "SignatureSeriesAssetCreated", logs: logs, sub: sub}, nil
}

// WatchSignatureSeriesAssetCreated is a free log subscription operation binding the contract event 0x07ac23775381c7e58d18cef592a12cb5ea8576f2034c691e2500aba9b51f9d7f.
//
// Solidity: event SignatureSeriesAssetCreated(uint256 tokenID, address indexed creator, string metaDataURI)
func (_SignatureSeries *SignatureSeriesFilterer) WatchSignatureSeriesAssetCreated(opts *bind.WatchOpts, sink chan<- *SignatureSeriesSignatureSeriesAssetCreated, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _SignatureSeries.contract.WatchLogs(opts, "SignatureSeriesAssetCreated", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignatureSeriesSignatureSeriesAssetCreated)
				if err := _SignatureSeries.contract.UnpackLog(event, "SignatureSeriesAssetCreated", log); err != nil {
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

// ParseSignatureSeriesAssetCreated is a log parse operation binding the contract event 0x07ac23775381c7e58d18cef592a12cb5ea8576f2034c691e2500aba9b51f9d7f.
//
// Solidity: event SignatureSeriesAssetCreated(uint256 tokenID, address indexed creator, string metaDataURI)
func (_SignatureSeries *SignatureSeriesFilterer) ParseSignatureSeriesAssetCreated(log types.Log) (*SignatureSeriesSignatureSeriesAssetCreated, error) {
	event := new(SignatureSeriesSignatureSeriesAssetCreated)
	if err := _SignatureSeries.contract.UnpackLog(event, "SignatureSeriesAssetCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SignatureSeriesSignatureSeriesAssetDestroyedIterator is returned from FilterSignatureSeriesAssetDestroyed and is used to iterate over the raw logs and unpacked data for SignatureSeriesAssetDestroyed events raised by the SignatureSeries contract.
type SignatureSeriesSignatureSeriesAssetDestroyedIterator struct {
	Event *SignatureSeriesSignatureSeriesAssetDestroyed // Event containing the contract specifics and raw log

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
func (it *SignatureSeriesSignatureSeriesAssetDestroyedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignatureSeriesSignatureSeriesAssetDestroyed)
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
		it.Event = new(SignatureSeriesSignatureSeriesAssetDestroyed)
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
func (it *SignatureSeriesSignatureSeriesAssetDestroyedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignatureSeriesSignatureSeriesAssetDestroyedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignatureSeriesSignatureSeriesAssetDestroyed represents a SignatureSeriesAssetDestroyed event raised by the SignatureSeries contract.
type SignatureSeriesSignatureSeriesAssetDestroyed struct {
	TokenId         *big.Int
	OwnerOrApproved common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSignatureSeriesAssetDestroyed is a free log retrieval operation binding the contract event 0x70dfc397541721d0da674bb15688bca9a5fc5dfcf822c7e879e29197a6f73a29.
//
// Solidity: event SignatureSeriesAssetDestroyed(uint256 indexed tokenId, address ownerOrApproved)
func (_SignatureSeries *SignatureSeriesFilterer) FilterSignatureSeriesAssetDestroyed(opts *bind.FilterOpts, tokenId []*big.Int) (*SignatureSeriesSignatureSeriesAssetDestroyedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SignatureSeries.contract.FilterLogs(opts, "SignatureSeriesAssetDestroyed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SignatureSeriesSignatureSeriesAssetDestroyedIterator{contract: _SignatureSeries.contract, event: "SignatureSeriesAssetDestroyed", logs: logs, sub: sub}, nil
}

// WatchSignatureSeriesAssetDestroyed is a free log subscription operation binding the contract event 0x70dfc397541721d0da674bb15688bca9a5fc5dfcf822c7e879e29197a6f73a29.
//
// Solidity: event SignatureSeriesAssetDestroyed(uint256 indexed tokenId, address ownerOrApproved)
func (_SignatureSeries *SignatureSeriesFilterer) WatchSignatureSeriesAssetDestroyed(opts *bind.WatchOpts, sink chan<- *SignatureSeriesSignatureSeriesAssetDestroyed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SignatureSeries.contract.WatchLogs(opts, "SignatureSeriesAssetDestroyed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignatureSeriesSignatureSeriesAssetDestroyed)
				if err := _SignatureSeries.contract.UnpackLog(event, "SignatureSeriesAssetDestroyed", log); err != nil {
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

// ParseSignatureSeriesAssetDestroyed is a log parse operation binding the contract event 0x70dfc397541721d0da674bb15688bca9a5fc5dfcf822c7e879e29197a6f73a29.
//
// Solidity: event SignatureSeriesAssetDestroyed(uint256 indexed tokenId, address ownerOrApproved)
func (_SignatureSeries *SignatureSeriesFilterer) ParseSignatureSeriesAssetDestroyed(log types.Log) (*SignatureSeriesSignatureSeriesAssetDestroyed, error) {
	event := new(SignatureSeriesSignatureSeriesAssetDestroyed)
	if err := _SignatureSeries.contract.UnpackLog(event, "SignatureSeriesAssetDestroyed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SignatureSeriesTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SignatureSeries contract.
type SignatureSeriesTransferIterator struct {
	Event *SignatureSeriesTransfer // Event containing the contract specifics and raw log

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
func (it *SignatureSeriesTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignatureSeriesTransfer)
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
		it.Event = new(SignatureSeriesTransfer)
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
func (it *SignatureSeriesTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignatureSeriesTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignatureSeriesTransfer represents a Transfer event raised by the SignatureSeries contract.
type SignatureSeriesTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SignatureSeries *SignatureSeriesFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SignatureSeriesTransferIterator, error) {

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

	logs, sub, err := _SignatureSeries.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SignatureSeriesTransferIterator{contract: _SignatureSeries.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SignatureSeries *SignatureSeriesFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SignatureSeriesTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _SignatureSeries.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignatureSeriesTransfer)
				if err := _SignatureSeries.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_SignatureSeries *SignatureSeriesFilterer) ParseTransfer(log types.Log) (*SignatureSeriesTransfer, error) {
	event := new(SignatureSeriesTransfer)
	if err := _SignatureSeries.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SignatureSeriesUpdateUserIterator is returned from FilterUpdateUser and is used to iterate over the raw logs and unpacked data for UpdateUser events raised by the SignatureSeries contract.
type SignatureSeriesUpdateUserIterator struct {
	Event *SignatureSeriesUpdateUser // Event containing the contract specifics and raw log

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
func (it *SignatureSeriesUpdateUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignatureSeriesUpdateUser)
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
		it.Event = new(SignatureSeriesUpdateUser)
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
func (it *SignatureSeriesUpdateUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignatureSeriesUpdateUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignatureSeriesUpdateUser represents a UpdateUser event raised by the SignatureSeries contract.
type SignatureSeriesUpdateUser struct {
	TokenId *big.Int
	User    common.Address
	Expires uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateUser is a free log retrieval operation binding the contract event 0x4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe.
//
// Solidity: event UpdateUser(uint256 indexed tokenId, address indexed user, uint64 expires)
func (_SignatureSeries *SignatureSeriesFilterer) FilterUpdateUser(opts *bind.FilterOpts, tokenId []*big.Int, user []common.Address) (*SignatureSeriesUpdateUserIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SignatureSeries.contract.FilterLogs(opts, "UpdateUser", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &SignatureSeriesUpdateUserIterator{contract: _SignatureSeries.contract, event: "UpdateUser", logs: logs, sub: sub}, nil
}

// WatchUpdateUser is a free log subscription operation binding the contract event 0x4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe.
//
// Solidity: event UpdateUser(uint256 indexed tokenId, address indexed user, uint64 expires)
func (_SignatureSeries *SignatureSeriesFilterer) WatchUpdateUser(opts *bind.WatchOpts, sink chan<- *SignatureSeriesUpdateUser, tokenId []*big.Int, user []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SignatureSeries.contract.WatchLogs(opts, "UpdateUser", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignatureSeriesUpdateUser)
				if err := _SignatureSeries.contract.UnpackLog(event, "UpdateUser", log); err != nil {
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

// ParseUpdateUser is a log parse operation binding the contract event 0x4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe.
//
// Solidity: event UpdateUser(uint256 indexed tokenId, address indexed user, uint64 expires)
func (_SignatureSeries *SignatureSeriesFilterer) ParseUpdateUser(log types.Log) (*SignatureSeriesUpdateUser, error) {
	event := new(SignatureSeriesUpdateUser)
	if err := _SignatureSeries.contract.UnpackLog(event, "UpdateUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
