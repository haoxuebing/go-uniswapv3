// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package routerV2

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

// RouterV2MetaData contains all meta data concerning the RouterV2 contract.
var RouterV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"WooPoolChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumIWooRouterV2.SwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"WooRouterSwap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approveTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minToAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"externalSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"realToAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stuckToken\",\"type\":\"address\"}],\"name\":\"inCaseTokenGotStuck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"}],\"name\":\"querySwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"setPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"setWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minToAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rebateTo\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"realToAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"}],\"name\":\"tryQuerySwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wooPool\",\"outputs\":[{\"internalType\":\"contractIWooPPV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RouterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterV2MetaData.ABI instead.
var RouterV2ABI = RouterV2MetaData.ABI

// RouterV2 is an auto generated Go binding around an Ethereum contract.
type RouterV2 struct {
	RouterV2Caller     // Read-only binding to the contract
	RouterV2Transactor // Write-only binding to the contract
	RouterV2Filterer   // Log filterer for contract events
}

// RouterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type RouterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterV2Session struct {
	Contract     *RouterV2         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterV2CallerSession struct {
	Contract *RouterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RouterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterV2TransactorSession struct {
	Contract     *RouterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RouterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type RouterV2Raw struct {
	Contract *RouterV2 // Generic contract binding to access the raw methods on
}

// RouterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterV2CallerRaw struct {
	Contract *RouterV2Caller // Generic read-only contract binding to access the raw methods on
}

// RouterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterV2TransactorRaw struct {
	Contract *RouterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRouterV2 creates a new instance of RouterV2, bound to a specific deployed contract.
func NewRouterV2(address common.Address, backend bind.ContractBackend) (*RouterV2, error) {
	contract, err := bindRouterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RouterV2{RouterV2Caller: RouterV2Caller{contract: contract}, RouterV2Transactor: RouterV2Transactor{contract: contract}, RouterV2Filterer: RouterV2Filterer{contract: contract}}, nil
}

// NewRouterV2Caller creates a new read-only instance of RouterV2, bound to a specific deployed contract.
func NewRouterV2Caller(address common.Address, caller bind.ContractCaller) (*RouterV2Caller, error) {
	contract, err := bindRouterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterV2Caller{contract: contract}, nil
}

// NewRouterV2Transactor creates a new write-only instance of RouterV2, bound to a specific deployed contract.
func NewRouterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*RouterV2Transactor, error) {
	contract, err := bindRouterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterV2Transactor{contract: contract}, nil
}

// NewRouterV2Filterer creates a new log filterer instance of RouterV2, bound to a specific deployed contract.
func NewRouterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*RouterV2Filterer, error) {
	contract, err := bindRouterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterV2Filterer{contract: contract}, nil
}

// bindRouterV2 binds a generic wrapper to an already deployed contract.
func bindRouterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RouterV2 *RouterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RouterV2.Contract.RouterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RouterV2 *RouterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RouterV2.Contract.RouterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RouterV2 *RouterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RouterV2.Contract.RouterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RouterV2 *RouterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RouterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RouterV2 *RouterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RouterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RouterV2 *RouterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RouterV2.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_RouterV2 *RouterV2Caller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RouterV2.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_RouterV2 *RouterV2Session) WETH() (common.Address, error) {
	return _RouterV2.Contract.WETH(&_RouterV2.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_RouterV2 *RouterV2CallerSession) WETH() (common.Address, error) {
	return _RouterV2.Contract.WETH(&_RouterV2.CallOpts)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_RouterV2 *RouterV2Caller) IsWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _RouterV2.contract.Call(opts, &out, "isWhitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_RouterV2 *RouterV2Session) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _RouterV2.Contract.IsWhitelisted(&_RouterV2.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address ) view returns(bool)
func (_RouterV2 *RouterV2CallerSession) IsWhitelisted(arg0 common.Address) (bool, error) {
	return _RouterV2.Contract.IsWhitelisted(&_RouterV2.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RouterV2 *RouterV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RouterV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RouterV2 *RouterV2Session) Owner() (common.Address, error) {
	return _RouterV2.Contract.Owner(&_RouterV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RouterV2 *RouterV2CallerSession) Owner() (common.Address, error) {
	return _RouterV2.Contract.Owner(&_RouterV2.CallOpts)
}

// QuerySwap is a free data retrieval call binding the contract method 0xe94803f4.
//
// Solidity: function querySwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_RouterV2 *RouterV2Caller) QuerySwap(opts *bind.CallOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RouterV2.contract.Call(opts, &out, "querySwap", fromToken, toToken, fromAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySwap is a free data retrieval call binding the contract method 0xe94803f4.
//
// Solidity: function querySwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_RouterV2 *RouterV2Session) QuerySwap(fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	return _RouterV2.Contract.QuerySwap(&_RouterV2.CallOpts, fromToken, toToken, fromAmount)
}

// QuerySwap is a free data retrieval call binding the contract method 0xe94803f4.
//
// Solidity: function querySwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_RouterV2 *RouterV2CallerSession) QuerySwap(fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	return _RouterV2.Contract.QuerySwap(&_RouterV2.CallOpts, fromToken, toToken, fromAmount)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_RouterV2 *RouterV2Caller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RouterV2.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_RouterV2 *RouterV2Session) QuoteToken() (common.Address, error) {
	return _RouterV2.Contract.QuoteToken(&_RouterV2.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_RouterV2 *RouterV2CallerSession) QuoteToken() (common.Address, error) {
	return _RouterV2.Contract.QuoteToken(&_RouterV2.CallOpts)
}

// TryQuerySwap is a free data retrieval call binding the contract method 0x9d82e123.
//
// Solidity: function tryQuerySwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_RouterV2 *RouterV2Caller) TryQuerySwap(opts *bind.CallOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RouterV2.contract.Call(opts, &out, "tryQuerySwap", fromToken, toToken, fromAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TryQuerySwap is a free data retrieval call binding the contract method 0x9d82e123.
//
// Solidity: function tryQuerySwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_RouterV2 *RouterV2Session) TryQuerySwap(fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	return _RouterV2.Contract.TryQuerySwap(&_RouterV2.CallOpts, fromToken, toToken, fromAmount)
}

// TryQuerySwap is a free data retrieval call binding the contract method 0x9d82e123.
//
// Solidity: function tryQuerySwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 toAmount)
func (_RouterV2 *RouterV2CallerSession) TryQuerySwap(fromToken common.Address, toToken common.Address, fromAmount *big.Int) (*big.Int, error) {
	return _RouterV2.Contract.TryQuerySwap(&_RouterV2.CallOpts, fromToken, toToken, fromAmount)
}

// WooPool is a free data retrieval call binding the contract method 0xa7394603.
//
// Solidity: function wooPool() view returns(address)
func (_RouterV2 *RouterV2Caller) WooPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RouterV2.contract.Call(opts, &out, "wooPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WooPool is a free data retrieval call binding the contract method 0xa7394603.
//
// Solidity: function wooPool() view returns(address)
func (_RouterV2 *RouterV2Session) WooPool() (common.Address, error) {
	return _RouterV2.Contract.WooPool(&_RouterV2.CallOpts)
}

// WooPool is a free data retrieval call binding the contract method 0xa7394603.
//
// Solidity: function wooPool() view returns(address)
func (_RouterV2 *RouterV2CallerSession) WooPool() (common.Address, error) {
	return _RouterV2.Contract.WooPool(&_RouterV2.CallOpts)
}

// ExternalSwap is a paid mutator transaction binding the contract method 0x199b83fa.
//
// Solidity: function externalSwap(address approveTarget, address swapTarget, address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, bytes data) payable returns(uint256 realToAmount)
func (_RouterV2 *RouterV2Transactor) ExternalSwap(opts *bind.TransactOpts, approveTarget common.Address, swapTarget common.Address, fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _RouterV2.contract.Transact(opts, "externalSwap", approveTarget, swapTarget, fromToken, toToken, fromAmount, minToAmount, to, data)
}

// ExternalSwap is a paid mutator transaction binding the contract method 0x199b83fa.
//
// Solidity: function externalSwap(address approveTarget, address swapTarget, address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, bytes data) payable returns(uint256 realToAmount)
func (_RouterV2 *RouterV2Session) ExternalSwap(approveTarget common.Address, swapTarget common.Address, fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _RouterV2.Contract.ExternalSwap(&_RouterV2.TransactOpts, approveTarget, swapTarget, fromToken, toToken, fromAmount, minToAmount, to, data)
}

// ExternalSwap is a paid mutator transaction binding the contract method 0x199b83fa.
//
// Solidity: function externalSwap(address approveTarget, address swapTarget, address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, bytes data) payable returns(uint256 realToAmount)
func (_RouterV2 *RouterV2TransactorSession) ExternalSwap(approveTarget common.Address, swapTarget common.Address, fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _RouterV2.Contract.ExternalSwap(&_RouterV2.TransactOpts, approveTarget, swapTarget, fromToken, toToken, fromAmount, minToAmount, to, data)
}

// InCaseTokenGotStuck is a paid mutator transaction binding the contract method 0xe1a4e72a.
//
// Solidity: function inCaseTokenGotStuck(address stuckToken) returns()
func (_RouterV2 *RouterV2Transactor) InCaseTokenGotStuck(opts *bind.TransactOpts, stuckToken common.Address) (*types.Transaction, error) {
	return _RouterV2.contract.Transact(opts, "inCaseTokenGotStuck", stuckToken)
}

// InCaseTokenGotStuck is a paid mutator transaction binding the contract method 0xe1a4e72a.
//
// Solidity: function inCaseTokenGotStuck(address stuckToken) returns()
func (_RouterV2 *RouterV2Session) InCaseTokenGotStuck(stuckToken common.Address) (*types.Transaction, error) {
	return _RouterV2.Contract.InCaseTokenGotStuck(&_RouterV2.TransactOpts, stuckToken)
}

// InCaseTokenGotStuck is a paid mutator transaction binding the contract method 0xe1a4e72a.
//
// Solidity: function inCaseTokenGotStuck(address stuckToken) returns()
func (_RouterV2 *RouterV2TransactorSession) InCaseTokenGotStuck(stuckToken common.Address) (*types.Transaction, error) {
	return _RouterV2.Contract.InCaseTokenGotStuck(&_RouterV2.TransactOpts, stuckToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RouterV2 *RouterV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RouterV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RouterV2 *RouterV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _RouterV2.Contract.RenounceOwnership(&_RouterV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RouterV2 *RouterV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RouterV2.Contract.RenounceOwnership(&_RouterV2.TransactOpts)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address newPool) returns()
func (_RouterV2 *RouterV2Transactor) SetPool(opts *bind.TransactOpts, newPool common.Address) (*types.Transaction, error) {
	return _RouterV2.contract.Transact(opts, "setPool", newPool)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address newPool) returns()
func (_RouterV2 *RouterV2Session) SetPool(newPool common.Address) (*types.Transaction, error) {
	return _RouterV2.Contract.SetPool(&_RouterV2.TransactOpts, newPool)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address newPool) returns()
func (_RouterV2 *RouterV2TransactorSession) SetPool(newPool common.Address) (*types.Transaction, error) {
	return _RouterV2.Contract.SetPool(&_RouterV2.TransactOpts, newPool)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0x9281aa0b.
//
// Solidity: function setWhitelisted(address target, bool whitelisted) returns()
func (_RouterV2 *RouterV2Transactor) SetWhitelisted(opts *bind.TransactOpts, target common.Address, whitelisted bool) (*types.Transaction, error) {
	return _RouterV2.contract.Transact(opts, "setWhitelisted", target, whitelisted)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0x9281aa0b.
//
// Solidity: function setWhitelisted(address target, bool whitelisted) returns()
func (_RouterV2 *RouterV2Session) SetWhitelisted(target common.Address, whitelisted bool) (*types.Transaction, error) {
	return _RouterV2.Contract.SetWhitelisted(&_RouterV2.TransactOpts, target, whitelisted)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0x9281aa0b.
//
// Solidity: function setWhitelisted(address target, bool whitelisted) returns()
func (_RouterV2 *RouterV2TransactorSession) SetWhitelisted(target common.Address, whitelisted bool) (*types.Transaction, error) {
	return _RouterV2.Contract.SetWhitelisted(&_RouterV2.TransactOpts, target, whitelisted)
}

// Swap is a paid mutator transaction binding the contract method 0x7dc20382.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, address rebateTo) payable returns(uint256 realToAmount)
func (_RouterV2 *RouterV2Transactor) Swap(opts *bind.TransactOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _RouterV2.contract.Transact(opts, "swap", fromToken, toToken, fromAmount, minToAmount, to, rebateTo)
}

// Swap is a paid mutator transaction binding the contract method 0x7dc20382.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, address rebateTo) payable returns(uint256 realToAmount)
func (_RouterV2 *RouterV2Session) Swap(fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _RouterV2.Contract.Swap(&_RouterV2.TransactOpts, fromToken, toToken, fromAmount, minToAmount, to, rebateTo)
}

// Swap is a paid mutator transaction binding the contract method 0x7dc20382.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minToAmount, address to, address rebateTo) payable returns(uint256 realToAmount)
func (_RouterV2 *RouterV2TransactorSession) Swap(fromToken common.Address, toToken common.Address, fromAmount *big.Int, minToAmount *big.Int, to common.Address, rebateTo common.Address) (*types.Transaction, error) {
	return _RouterV2.Contract.Swap(&_RouterV2.TransactOpts, fromToken, toToken, fromAmount, minToAmount, to, rebateTo)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RouterV2 *RouterV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RouterV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RouterV2 *RouterV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RouterV2.Contract.TransferOwnership(&_RouterV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RouterV2 *RouterV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RouterV2.Contract.TransferOwnership(&_RouterV2.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RouterV2 *RouterV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RouterV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RouterV2 *RouterV2Session) Receive() (*types.Transaction, error) {
	return _RouterV2.Contract.Receive(&_RouterV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RouterV2 *RouterV2TransactorSession) Receive() (*types.Transaction, error) {
	return _RouterV2.Contract.Receive(&_RouterV2.TransactOpts)
}

// RouterV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RouterV2 contract.
type RouterV2OwnershipTransferredIterator struct {
	Event *RouterV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RouterV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterV2OwnershipTransferred)
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
		it.Event = new(RouterV2OwnershipTransferred)
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
func (it *RouterV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterV2OwnershipTransferred represents a OwnershipTransferred event raised by the RouterV2 contract.
type RouterV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RouterV2 *RouterV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RouterV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RouterV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RouterV2OwnershipTransferredIterator{contract: _RouterV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RouterV2 *RouterV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RouterV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RouterV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterV2OwnershipTransferred)
				if err := _RouterV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RouterV2 *RouterV2Filterer) ParseOwnershipTransferred(log types.Log) (*RouterV2OwnershipTransferred, error) {
	event := new(RouterV2OwnershipTransferred)
	if err := _RouterV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterV2WooPoolChangedIterator is returned from FilterWooPoolChanged and is used to iterate over the raw logs and unpacked data for WooPoolChanged events raised by the RouterV2 contract.
type RouterV2WooPoolChangedIterator struct {
	Event *RouterV2WooPoolChanged // Event containing the contract specifics and raw log

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
func (it *RouterV2WooPoolChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterV2WooPoolChanged)
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
		it.Event = new(RouterV2WooPoolChanged)
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
func (it *RouterV2WooPoolChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterV2WooPoolChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterV2WooPoolChanged represents a WooPoolChanged event raised by the RouterV2 contract.
type RouterV2WooPoolChanged struct {
	NewPool common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWooPoolChanged is a free log retrieval operation binding the contract event 0x4577a21bd8e55848c574b7582f8e6cc6a7cf1c1958f36a9751eab6329d656b1e.
//
// Solidity: event WooPoolChanged(address newPool)
func (_RouterV2 *RouterV2Filterer) FilterWooPoolChanged(opts *bind.FilterOpts) (*RouterV2WooPoolChangedIterator, error) {

	logs, sub, err := _RouterV2.contract.FilterLogs(opts, "WooPoolChanged")
	if err != nil {
		return nil, err
	}
	return &RouterV2WooPoolChangedIterator{contract: _RouterV2.contract, event: "WooPoolChanged", logs: logs, sub: sub}, nil
}

// WatchWooPoolChanged is a free log subscription operation binding the contract event 0x4577a21bd8e55848c574b7582f8e6cc6a7cf1c1958f36a9751eab6329d656b1e.
//
// Solidity: event WooPoolChanged(address newPool)
func (_RouterV2 *RouterV2Filterer) WatchWooPoolChanged(opts *bind.WatchOpts, sink chan<- *RouterV2WooPoolChanged) (event.Subscription, error) {

	logs, sub, err := _RouterV2.contract.WatchLogs(opts, "WooPoolChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterV2WooPoolChanged)
				if err := _RouterV2.contract.UnpackLog(event, "WooPoolChanged", log); err != nil {
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

// ParseWooPoolChanged is a log parse operation binding the contract event 0x4577a21bd8e55848c574b7582f8e6cc6a7cf1c1958f36a9751eab6329d656b1e.
//
// Solidity: event WooPoolChanged(address newPool)
func (_RouterV2 *RouterV2Filterer) ParseWooPoolChanged(log types.Log) (*RouterV2WooPoolChanged, error) {
	event := new(RouterV2WooPoolChanged)
	if err := _RouterV2.contract.UnpackLog(event, "WooPoolChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterV2WooRouterSwapIterator is returned from FilterWooRouterSwap and is used to iterate over the raw logs and unpacked data for WooRouterSwap events raised by the RouterV2 contract.
type RouterV2WooRouterSwapIterator struct {
	Event *RouterV2WooRouterSwap // Event containing the contract specifics and raw log

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
func (it *RouterV2WooRouterSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterV2WooRouterSwap)
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
		it.Event = new(RouterV2WooRouterSwap)
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
func (it *RouterV2WooRouterSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterV2WooRouterSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterV2WooRouterSwap represents a WooRouterSwap event raised by the RouterV2 contract.
type RouterV2WooRouterSwap struct {
	SwapType   uint8
	FromToken  common.Address
	ToToken    common.Address
	FromAmount *big.Int
	ToAmount   *big.Int
	From       common.Address
	To         common.Address
	RebateTo   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWooRouterSwap is a free log retrieval operation binding the contract event 0x27c98e911efdd224f4002f6cd831c3ad0d2759ee176f9ee8466d95826af22a1c.
//
// Solidity: event WooRouterSwap(uint8 swapType, address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to, address rebateTo)
func (_RouterV2 *RouterV2Filterer) FilterWooRouterSwap(opts *bind.FilterOpts, fromToken []common.Address, toToken []common.Address, to []common.Address) (*RouterV2WooRouterSwapIterator, error) {

	var fromTokenRule []interface{}
	for _, fromTokenItem := range fromToken {
		fromTokenRule = append(fromTokenRule, fromTokenItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RouterV2.contract.FilterLogs(opts, "WooRouterSwap", fromTokenRule, toTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RouterV2WooRouterSwapIterator{contract: _RouterV2.contract, event: "WooRouterSwap", logs: logs, sub: sub}, nil
}

// WatchWooRouterSwap is a free log subscription operation binding the contract event 0x27c98e911efdd224f4002f6cd831c3ad0d2759ee176f9ee8466d95826af22a1c.
//
// Solidity: event WooRouterSwap(uint8 swapType, address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to, address rebateTo)
func (_RouterV2 *RouterV2Filterer) WatchWooRouterSwap(opts *bind.WatchOpts, sink chan<- *RouterV2WooRouterSwap, fromToken []common.Address, toToken []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenRule []interface{}
	for _, fromTokenItem := range fromToken {
		fromTokenRule = append(fromTokenRule, fromTokenItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RouterV2.contract.WatchLogs(opts, "WooRouterSwap", fromTokenRule, toTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterV2WooRouterSwap)
				if err := _RouterV2.contract.UnpackLog(event, "WooRouterSwap", log); err != nil {
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

// ParseWooRouterSwap is a log parse operation binding the contract event 0x27c98e911efdd224f4002f6cd831c3ad0d2759ee176f9ee8466d95826af22a1c.
//
// Solidity: event WooRouterSwap(uint8 swapType, address indexed fromToken, address indexed toToken, uint256 fromAmount, uint256 toAmount, address from, address indexed to, address rebateTo)
func (_RouterV2 *RouterV2Filterer) ParseWooRouterSwap(log types.Log) (*RouterV2WooRouterSwap, error) {
	event := new(RouterV2WooRouterSwap)
	if err := _RouterV2.contract.UnpackLog(event, "WooRouterSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
