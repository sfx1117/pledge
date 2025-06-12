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

// BscPledgeOracleTestnetTokenMetaData contains all meta data concerning the BscPledgeOracleTestnetToken contract.
var BscPledgeOracleTestnetTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetsAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"}],\"name\":\"getPrices\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"}],\"name\":\"getUnderlyingAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"}],\"name\":\"getUnderlyingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aggergator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"setAssetsAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDecimals\",\"type\":\"uint256\"}],\"name\":\"setDecimals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"aggergator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"setUnderlyingAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setUnderlyingPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BscPledgeOracleTestnetTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BscPledgeOracleTestnetTokenMetaData.ABI instead.
var BscPledgeOracleTestnetTokenABI = BscPledgeOracleTestnetTokenMetaData.ABI

// BscPledgeOracleTestnetToken is an auto generated Go binding around an Ethereum contract.
type BscPledgeOracleTestnetToken struct {
	BscPledgeOracleTestnetTokenCaller     // Read-only binding to the contract
	BscPledgeOracleTestnetTokenTransactor // Write-only binding to the contract
	BscPledgeOracleTestnetTokenFilterer   // Log filterer for contract events
}

// BscPledgeOracleTestnetTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BscPledgeOracleTestnetTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleTestnetTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BscPledgeOracleTestnetTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleTestnetTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BscPledgeOracleTestnetTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleTestnetTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BscPledgeOracleTestnetTokenSession struct {
	Contract     *BscPledgeOracleTestnetToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BscPledgeOracleTestnetTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BscPledgeOracleTestnetTokenCallerSession struct {
	Contract *BscPledgeOracleTestnetTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// BscPledgeOracleTestnetTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BscPledgeOracleTestnetTokenTransactorSession struct {
	Contract     *BscPledgeOracleTestnetTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// BscPledgeOracleTestnetTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BscPledgeOracleTestnetTokenRaw struct {
	Contract *BscPledgeOracleTestnetToken // Generic contract binding to access the raw methods on
}

// BscPledgeOracleTestnetTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BscPledgeOracleTestnetTokenCallerRaw struct {
	Contract *BscPledgeOracleTestnetTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BscPledgeOracleTestnetTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BscPledgeOracleTestnetTokenTransactorRaw struct {
	Contract *BscPledgeOracleTestnetTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBscPledgeOracleTestnetToken creates a new instance of BscPledgeOracleTestnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleTestnetToken(address common.Address, backend bind.ContractBackend) (*BscPledgeOracleTestnetToken, error) {
	contract, err := bindBscPledgeOracleTestnetToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleTestnetToken{BscPledgeOracleTestnetTokenCaller: BscPledgeOracleTestnetTokenCaller{contract: contract}, BscPledgeOracleTestnetTokenTransactor: BscPledgeOracleTestnetTokenTransactor{contract: contract}, BscPledgeOracleTestnetTokenFilterer: BscPledgeOracleTestnetTokenFilterer{contract: contract}}, nil
}

// NewBscPledgeOracleTestnetTokenCaller creates a new read-only instance of BscPledgeOracleTestnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleTestnetTokenCaller(address common.Address, caller bind.ContractCaller) (*BscPledgeOracleTestnetTokenCaller, error) {
	contract, err := bindBscPledgeOracleTestnetToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleTestnetTokenCaller{contract: contract}, nil
}

// NewBscPledgeOracleTestnetTokenTransactor creates a new write-only instance of BscPledgeOracleTestnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleTestnetTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BscPledgeOracleTestnetTokenTransactor, error) {
	contract, err := bindBscPledgeOracleTestnetToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleTestnetTokenTransactor{contract: contract}, nil
}

// NewBscPledgeOracleTestnetTokenFilterer creates a new log filterer instance of BscPledgeOracleTestnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleTestnetTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BscPledgeOracleTestnetTokenFilterer, error) {
	contract, err := bindBscPledgeOracleTestnetToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleTestnetTokenFilterer{contract: contract}, nil
}

// bindBscPledgeOracleTestnetToken binds a generic wrapper to an already deployed contract.
func bindBscPledgeOracleTestnetToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BscPledgeOracleTestnetTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscPledgeOracleTestnetToken.Contract.BscPledgeOracleTestnetTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.BscPledgeOracleTestnetTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.BscPledgeOracleTestnetTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscPledgeOracleTestnetToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.contract.Transact(opts, method, params...)
}

// GetAssetsAggregator is a free data retrieval call binding the contract method 0xb889a989.
//
// Solidity: function getAssetsAggregator(address asset) view returns(address, uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCaller) GetAssetsAggregator(opts *bind.CallOpts, asset common.Address) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleTestnetToken.contract.Call(opts, &out, "getAssetsAggregator", asset)

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
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) GetAssetsAggregator(asset common.Address) (common.Address, *big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetAssetsAggregator(&_BscPledgeOracleTestnetToken.CallOpts, asset)
}

// GetAssetsAggregator is a free data retrieval call binding the contract method 0xb889a989.
//
// Solidity: function getAssetsAggregator(address asset) view returns(address, uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCallerSession) GetAssetsAggregator(asset common.Address) (common.Address, *big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetAssetsAggregator(&_BscPledgeOracleTestnetToken.CallOpts, asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCaller) GetPrice(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleTestnetToken.contract.Call(opts, &out, "getPrice", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) GetPrice(asset common.Address) (*big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetPrice(&_BscPledgeOracleTestnetToken.CallOpts, asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCallerSession) GetPrice(asset common.Address) (*big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetPrice(&_BscPledgeOracleTestnetToken.CallOpts, asset)
}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCaller) GetPrices(opts *bind.CallOpts, assets []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleTestnetToken.contract.Call(opts, &out, "getPrices", assets)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) GetPrices(assets []*big.Int) ([]*big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetPrices(&_BscPledgeOracleTestnetToken.CallOpts, assets)
}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCallerSession) GetPrices(assets []*big.Int) ([]*big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetPrices(&_BscPledgeOracleTestnetToken.CallOpts, assets)
}

// GetUnderlyingAggregator is a free data retrieval call binding the contract method 0x75e443aa.
//
// Solidity: function getUnderlyingAggregator(uint256 underlying) view returns(address, uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCaller) GetUnderlyingAggregator(opts *bind.CallOpts, underlying *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleTestnetToken.contract.Call(opts, &out, "getUnderlyingAggregator", underlying)

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
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) GetUnderlyingAggregator(underlying *big.Int) (common.Address, *big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetUnderlyingAggregator(&_BscPledgeOracleTestnetToken.CallOpts, underlying)
}

// GetUnderlyingAggregator is a free data retrieval call binding the contract method 0x75e443aa.
//
// Solidity: function getUnderlyingAggregator(uint256 underlying) view returns(address, uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCallerSession) GetUnderlyingAggregator(underlying *big.Int) (common.Address, *big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetUnderlyingAggregator(&_BscPledgeOracleTestnetToken.CallOpts, underlying)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCaller) GetUnderlyingPrice(opts *bind.CallOpts, underlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleTestnetToken.contract.Call(opts, &out, "getUnderlyingPrice", underlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) GetUnderlyingPrice(underlying *big.Int) (*big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetUnderlyingPrice(&_BscPledgeOracleTestnetToken.CallOpts, underlying)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCallerSession) GetUnderlyingPrice(underlying *big.Int) (*big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetUnderlyingPrice(&_BscPledgeOracleTestnetToken.CallOpts, underlying)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscPledgeOracleTestnetToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) Owner() (common.Address, error) {
	return _BscPledgeOracleTestnetToken.Contract.Owner(&_BscPledgeOracleTestnetToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCallerSession) Owner() (common.Address, error) {
	return _BscPledgeOracleTestnetToken.Contract.Owner(&_BscPledgeOracleTestnetToken.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.RenounceOwnership(&_BscPledgeOracleTestnetToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.RenounceOwnership(&_BscPledgeOracleTestnetToken.TransactOpts)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xcd9ffa0b.
//
// Solidity: function setAssetsAggregator(address asset, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactor) SetAssetsAggregator(opts *bind.TransactOpts, asset common.Address, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.contract.Transact(opts, "setAssetsAggregator", asset, aggergator, _decimals)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xcd9ffa0b.
//
// Solidity: function setAssetsAggregator(address asset, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) SetAssetsAggregator(asset common.Address, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetAssetsAggregator(&_BscPledgeOracleTestnetToken.TransactOpts, asset, aggergator, _decimals)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xcd9ffa0b.
//
// Solidity: function setAssetsAggregator(address asset, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorSession) SetAssetsAggregator(asset common.Address, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetAssetsAggregator(&_BscPledgeOracleTestnetToken.TransactOpts, asset, aggergator, _decimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactor) SetDecimals(opts *bind.TransactOpts, newDecimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.contract.Transact(opts, "setDecimals", newDecimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) SetDecimals(newDecimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetDecimals(&_BscPledgeOracleTestnetToken.TransactOpts, newDecimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorSession) SetDecimals(newDecimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetDecimals(&_BscPledgeOracleTestnetToken.TransactOpts, newDecimals)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactor) SetPrice(opts *bind.TransactOpts, asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.contract.Transact(opts, "setPrice", asset, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) SetPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetPrice(&_BscPledgeOracleTestnetToken.TransactOpts, asset, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorSession) SetPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetPrice(&_BscPledgeOracleTestnetToken.TransactOpts, asset, price)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactor) SetPrices(opts *bind.TransactOpts, assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.contract.Transact(opts, "setPrices", assets, prices)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) SetPrices(assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetPrices(&_BscPledgeOracleTestnetToken.TransactOpts, assets, prices)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorSession) SetPrices(assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetPrices(&_BscPledgeOracleTestnetToken.TransactOpts, assets, prices)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x434b33ac.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactor) SetUnderlyingAggregator(opts *bind.TransactOpts, underlying *big.Int, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.contract.Transact(opts, "setUnderlyingAggregator", underlying, aggergator, _decimals)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x434b33ac.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) SetUnderlyingAggregator(underlying *big.Int, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetUnderlyingAggregator(&_BscPledgeOracleTestnetToken.TransactOpts, underlying, aggergator, _decimals)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x434b33ac.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorSession) SetUnderlyingAggregator(underlying *big.Int, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetUnderlyingAggregator(&_BscPledgeOracleTestnetToken.TransactOpts, underlying, aggergator, _decimals)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactor) SetUnderlyingPrice(opts *bind.TransactOpts, underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.contract.Transact(opts, "setUnderlyingPrice", underlying, price)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) SetUnderlyingPrice(underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetUnderlyingPrice(&_BscPledgeOracleTestnetToken.TransactOpts, underlying, price)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorSession) SetUnderlyingPrice(underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetUnderlyingPrice(&_BscPledgeOracleTestnetToken.TransactOpts, underlying, price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.TransferOwnership(&_BscPledgeOracleTestnetToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.TransferOwnership(&_BscPledgeOracleTestnetToken.TransactOpts, newOwner)
}

// BscPledgeOracleTestnetTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BscPledgeOracleTestnetToken contract.
type BscPledgeOracleTestnetTokenOwnershipTransferredIterator struct {
	Event *BscPledgeOracleTestnetTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BscPledgeOracleTestnetTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscPledgeOracleTestnetTokenOwnershipTransferred)
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
		it.Event = new(BscPledgeOracleTestnetTokenOwnershipTransferred)
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
func (it *BscPledgeOracleTestnetTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscPledgeOracleTestnetTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscPledgeOracleTestnetTokenOwnershipTransferred represents a OwnershipTransferred event raised by the BscPledgeOracleTestnetToken contract.
type BscPledgeOracleTestnetTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BscPledgeOracleTestnetTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscPledgeOracleTestnetToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleTestnetTokenOwnershipTransferredIterator{contract: _BscPledgeOracleTestnetToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BscPledgeOracleTestnetTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscPledgeOracleTestnetToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscPledgeOracleTestnetTokenOwnershipTransferred)
				if err := _BscPledgeOracleTestnetToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenFilterer) ParseOwnershipTransferred(log types.Log) (*BscPledgeOracleTestnetTokenOwnershipTransferred, error) {
	event := new(BscPledgeOracleTestnetTokenOwnershipTransferred)
	if err := _BscPledgeOracleTestnetToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
