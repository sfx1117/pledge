// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// BscPledgeOracleMainnetTokenMetaData contains all meta data concerning the BscPledgeOracleMainnetToken contract.
var BscPledgeOracleMainnetTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetsAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"}],\"name\":\"getPrices\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"}],\"name\":\"getUnderlyingAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"}],\"name\":\"getUnderlyingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aggergator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"setAssetsAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDecimals\",\"type\":\"uint256\"}],\"name\":\"setDecimals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"aggergator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"setUnderlyingAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setUnderlyingPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BscPledgeOracleMainnetTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BscPledgeOracleMainnetTokenMetaData.ABI instead.
var BscPledgeOracleMainnetTokenABI = BscPledgeOracleMainnetTokenMetaData.ABI

// BscPledgeOracleMainnetToken is an auto generated Go binding around an Ethereum contract.
type BscPledgeOracleMainnetToken struct {
	BscPledgeOracleMainnetTokenCaller     // Read-only binding to the contract
	BscPledgeOracleMainnetTokenTransactor // Write-only binding to the contract
	BscPledgeOracleMainnetTokenFilterer   // Log filterer for contract events
}

// BscPledgeOracleMainnetTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BscPledgeOracleMainnetTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleMainnetTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BscPledgeOracleMainnetTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleMainnetTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BscPledgeOracleMainnetTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleMainnetTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BscPledgeOracleMainnetTokenSession struct {
	Contract     *BscPledgeOracleMainnetToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BscPledgeOracleMainnetTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BscPledgeOracleMainnetTokenCallerSession struct {
	Contract *BscPledgeOracleMainnetTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// BscPledgeOracleMainnetTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BscPledgeOracleMainnetTokenTransactorSession struct {
	Contract     *BscPledgeOracleMainnetTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// BscPledgeOracleMainnetTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BscPledgeOracleMainnetTokenRaw struct {
	Contract *BscPledgeOracleMainnetToken // Generic contract binding to access the raw methods on
}

// BscPledgeOracleMainnetTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BscPledgeOracleMainnetTokenCallerRaw struct {
	Contract *BscPledgeOracleMainnetTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BscPledgeOracleMainnetTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BscPledgeOracleMainnetTokenTransactorRaw struct {
	Contract *BscPledgeOracleMainnetTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBscPledgeOracleMainnetToken creates a new instance of BscPledgeOracleMainnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleMainnetToken(address common.Address, backend bind.ContractBackend) (*BscPledgeOracleMainnetToken, error) {
	contract, err := bindBscPledgeOracleMainnetToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleMainnetToken{BscPledgeOracleMainnetTokenCaller: BscPledgeOracleMainnetTokenCaller{contract: contract}, BscPledgeOracleMainnetTokenTransactor: BscPledgeOracleMainnetTokenTransactor{contract: contract}, BscPledgeOracleMainnetTokenFilterer: BscPledgeOracleMainnetTokenFilterer{contract: contract}}, nil
}

// NewBscPledgeOracleMainnetTokenCaller creates a new read-only instance of BscPledgeOracleMainnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleMainnetTokenCaller(address common.Address, caller bind.ContractCaller) (*BscPledgeOracleMainnetTokenCaller, error) {
	contract, err := bindBscPledgeOracleMainnetToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleMainnetTokenCaller{contract: contract}, nil
}

// NewBscPledgeOracleMainnetTokenTransactor creates a new write-only instance of BscPledgeOracleMainnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleMainnetTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BscPledgeOracleMainnetTokenTransactor, error) {
	contract, err := bindBscPledgeOracleMainnetToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleMainnetTokenTransactor{contract: contract}, nil
}

// NewBscPledgeOracleMainnetTokenFilterer creates a new log filterer instance of BscPledgeOracleMainnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleMainnetTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BscPledgeOracleMainnetTokenFilterer, error) {
	contract, err := bindBscPledgeOracleMainnetToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleMainnetTokenFilterer{contract: contract}, nil
}

// bindBscPledgeOracleMainnetToken binds a generic wrapper to an already deployed contract.
func bindBscPledgeOracleMainnetToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BscPledgeOracleMainnetTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscPledgeOracleMainnetToken.Contract.BscPledgeOracleMainnetTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.BscPledgeOracleMainnetTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.BscPledgeOracleMainnetTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscPledgeOracleMainnetToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.contract.Transact(opts, method, params...)
}

// GetAssetsAggregator is a free data retrieval call binding the contract method 0xb889a989.
//
// Solidity: function getAssetsAggregator(address asset) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) GetAssetsAggregator(opts *bind.CallOpts, asset common.Address) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "getAssetsAggregator", asset)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAssetsAggregator is a free data retrieval call binding the contract method 0xb889a989.
//
// Solidity: function getAssetsAggregator(address asset) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) GetAssetsAggregator(asset common.Address) (common.Address, *big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetAssetsAggregator(&_BscPledgeOracleMainnetToken.CallOpts, asset)
}

// GetAssetsAggregator is a free data retrieval call binding the contract method 0xb889a989.
//
// Solidity: function getAssetsAggregator(address asset) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) GetAssetsAggregator(asset common.Address) (common.Address, *big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetAssetsAggregator(&_BscPledgeOracleMainnetToken.CallOpts, asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) GetPrice(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "getPrice", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) GetPrice(asset common.Address) (*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetPrice(&_BscPledgeOracleMainnetToken.CallOpts, asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) GetPrice(asset common.Address) (*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetPrice(&_BscPledgeOracleMainnetToken.CallOpts, asset)
}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) GetPrices(opts *bind.CallOpts, assets []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "getPrices", assets)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) GetPrices(assets []*big.Int) ([]*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetPrices(&_BscPledgeOracleMainnetToken.CallOpts, assets)
}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) GetPrices(assets []*big.Int) ([]*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetPrices(&_BscPledgeOracleMainnetToken.CallOpts, assets)
}

// GetUnderlyingAggregator is a free data retrieval call binding the contract method 0x75e443aa.
//
// Solidity: function getUnderlyingAggregator(uint256 underlying) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) GetUnderlyingAggregator(opts *bind.CallOpts, underlying *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "getUnderlyingAggregator", underlying)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetUnderlyingAggregator is a free data retrieval call binding the contract method 0x75e443aa.
//
// Solidity: function getUnderlyingAggregator(uint256 underlying) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) GetUnderlyingAggregator(underlying *big.Int) (common.Address, *big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetUnderlyingAggregator(&_BscPledgeOracleMainnetToken.CallOpts, underlying)
}

// GetUnderlyingAggregator is a free data retrieval call binding the contract method 0x75e443aa.
//
// Solidity: function getUnderlyingAggregator(uint256 underlying) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) GetUnderlyingAggregator(underlying *big.Int) (common.Address, *big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetUnderlyingAggregator(&_BscPledgeOracleMainnetToken.CallOpts, underlying)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) GetUnderlyingPrice(opts *bind.CallOpts, underlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "getUnderlyingPrice", underlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) GetUnderlyingPrice(underlying *big.Int) (*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetUnderlyingPrice(&_BscPledgeOracleMainnetToken.CallOpts, underlying)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) GetUnderlyingPrice(underlying *big.Int) (*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetUnderlyingPrice(&_BscPledgeOracleMainnetToken.CallOpts, underlying)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) Owner() (common.Address, error) {
	return _BscPledgeOracleMainnetToken.Contract.Owner(&_BscPledgeOracleMainnetToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) Owner() (common.Address, error) {
	return _BscPledgeOracleMainnetToken.Contract.Owner(&_BscPledgeOracleMainnetToken.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.RenounceOwnership(&_BscPledgeOracleMainnetToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.RenounceOwnership(&_BscPledgeOracleMainnetToken.TransactOpts)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xcd9ffa0b.
//
// Solidity: function setAssetsAggregator(address asset, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetAssetsAggregator(opts *bind.TransactOpts, asset common.Address, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setAssetsAggregator", asset, aggergator, _decimals)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xcd9ffa0b.
//
// Solidity: function setAssetsAggregator(address asset, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetAssetsAggregator(asset common.Address, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetAssetsAggregator(&_BscPledgeOracleMainnetToken.TransactOpts, asset, aggergator, _decimals)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xcd9ffa0b.
//
// Solidity: function setAssetsAggregator(address asset, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetAssetsAggregator(asset common.Address, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetAssetsAggregator(&_BscPledgeOracleMainnetToken.TransactOpts, asset, aggergator, _decimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetDecimals(opts *bind.TransactOpts, newDecimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setDecimals", newDecimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetDecimals(newDecimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetDecimals(&_BscPledgeOracleMainnetToken.TransactOpts, newDecimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetDecimals(newDecimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetDecimals(&_BscPledgeOracleMainnetToken.TransactOpts, newDecimals)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetPrice(opts *bind.TransactOpts, asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setPrice", asset, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetPrice(&_BscPledgeOracleMainnetToken.TransactOpts, asset, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetPrice(&_BscPledgeOracleMainnetToken.TransactOpts, asset, price)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetPrices(opts *bind.TransactOpts, assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setPrices", assets, prices)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetPrices(assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetPrices(&_BscPledgeOracleMainnetToken.TransactOpts, assets, prices)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetPrices(assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetPrices(&_BscPledgeOracleMainnetToken.TransactOpts, assets, prices)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x434b33ac.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetUnderlyingAggregator(opts *bind.TransactOpts, underlying *big.Int, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setUnderlyingAggregator", underlying, aggergator, _decimals)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x434b33ac.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetUnderlyingAggregator(underlying *big.Int, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetUnderlyingAggregator(&_BscPledgeOracleMainnetToken.TransactOpts, underlying, aggergator, _decimals)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x434b33ac.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetUnderlyingAggregator(underlying *big.Int, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetUnderlyingAggregator(&_BscPledgeOracleMainnetToken.TransactOpts, underlying, aggergator, _decimals)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetUnderlyingPrice(opts *bind.TransactOpts, underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setUnderlyingPrice", underlying, price)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetUnderlyingPrice(underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetUnderlyingPrice(&_BscPledgeOracleMainnetToken.TransactOpts, underlying, price)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetUnderlyingPrice(underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetUnderlyingPrice(&_BscPledgeOracleMainnetToken.TransactOpts, underlying, price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.TransferOwnership(&_BscPledgeOracleMainnetToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.TransferOwnership(&_BscPledgeOracleMainnetToken.TransactOpts, newOwner)
}

// BscPledgeOracleMainnetTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BscPledgeOracleMainnetToken contract.
type BscPledgeOracleMainnetTokenOwnershipTransferredIterator struct {
	Event *BscPledgeOracleMainnetTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BscPledgeOracleMainnetTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscPledgeOracleMainnetTokenOwnershipTransferred)
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
		it.Event = new(BscPledgeOracleMainnetTokenOwnershipTransferred)
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
func (it *BscPledgeOracleMainnetTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscPledgeOracleMainnetTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscPledgeOracleMainnetTokenOwnershipTransferred represents a OwnershipTransferred event raised by the BscPledgeOracleMainnetToken contract.
type BscPledgeOracleMainnetTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BscPledgeOracleMainnetTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscPledgeOracleMainnetToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleMainnetTokenOwnershipTransferredIterator{contract: _BscPledgeOracleMainnetToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BscPledgeOracleMainnetTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscPledgeOracleMainnetToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscPledgeOracleMainnetTokenOwnershipTransferred)
				if err := _BscPledgeOracleMainnetToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenFilterer) ParseOwnershipTransferred(log types.Log) (*BscPledgeOracleMainnetTokenOwnershipTransferred, error) {
	event := new(BscPledgeOracleMainnetTokenOwnershipTransferred)
	if err := _BscPledgeOracleMainnetToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
