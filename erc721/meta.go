// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc721meta

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

// Erc721metaMetaData contains all meta data concerning the Erc721meta contract.
var Erc721metaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Erc721metaABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc721metaMetaData.ABI instead.
var Erc721metaABI = Erc721metaMetaData.ABI

// Erc721meta is an auto generated Go binding around an Ethereum contract.
type Erc721meta struct {
	Erc721metaCaller     // Read-only binding to the contract
	Erc721metaTransactor // Write-only binding to the contract
	Erc721metaFilterer   // Log filterer for contract events
}

// Erc721metaCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc721metaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721metaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc721metaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721metaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc721metaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721metaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc721metaSession struct {
	Contract     *Erc721meta       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc721metaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc721metaCallerSession struct {
	Contract *Erc721metaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Erc721metaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc721metaTransactorSession struct {
	Contract     *Erc721metaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Erc721metaRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc721metaRaw struct {
	Contract *Erc721meta // Generic contract binding to access the raw methods on
}

// Erc721metaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc721metaCallerRaw struct {
	Contract *Erc721metaCaller // Generic read-only contract binding to access the raw methods on
}

// Erc721metaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc721metaTransactorRaw struct {
	Contract *Erc721metaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc721meta creates a new instance of Erc721meta, bound to a specific deployed contract.
func NewErc721meta(address common.Address, backend bind.ContractBackend) (*Erc721meta, error) {
	contract, err := bindErc721meta(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc721meta{Erc721metaCaller: Erc721metaCaller{contract: contract}, Erc721metaTransactor: Erc721metaTransactor{contract: contract}, Erc721metaFilterer: Erc721metaFilterer{contract: contract}}, nil
}

// NewErc721metaCaller creates a new read-only instance of Erc721meta, bound to a specific deployed contract.
func NewErc721metaCaller(address common.Address, caller bind.ContractCaller) (*Erc721metaCaller, error) {
	contract, err := bindErc721meta(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721metaCaller{contract: contract}, nil
}

// NewErc721metaTransactor creates a new write-only instance of Erc721meta, bound to a specific deployed contract.
func NewErc721metaTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc721metaTransactor, error) {
	contract, err := bindErc721meta(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721metaTransactor{contract: contract}, nil
}

// NewErc721metaFilterer creates a new log filterer instance of Erc721meta, bound to a specific deployed contract.
func NewErc721metaFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc721metaFilterer, error) {
	contract, err := bindErc721meta(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc721metaFilterer{contract: contract}, nil
}

// bindErc721meta binds a generic wrapper to an already deployed contract.
func bindErc721meta(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc721metaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721meta *Erc721metaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721meta.Contract.Erc721metaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721meta *Erc721metaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721meta.Contract.Erc721metaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721meta *Erc721metaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721meta.Contract.Erc721metaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721meta *Erc721metaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721meta.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721meta *Erc721metaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721meta.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721meta *Erc721metaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721meta.Contract.contract.Transact(opts, method, params...)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string _name)
func (_Erc721meta *Erc721metaCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc721meta.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string _name)
func (_Erc721meta *Erc721metaSession) Name() (string, error) {
	return _Erc721meta.Contract.Name(&_Erc721meta.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string _name)
func (_Erc721meta *Erc721metaCallerSession) Name() (string, error) {
	return _Erc721meta.Contract.Name(&_Erc721meta.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string _symbol)
func (_Erc721meta *Erc721metaCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc721meta.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string _symbol)
func (_Erc721meta *Erc721metaSession) Symbol() (string, error) {
	return _Erc721meta.Contract.Symbol(&_Erc721meta.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string _symbol)
func (_Erc721meta *Erc721metaCallerSession) Symbol() (string, error) {
	return _Erc721meta.Contract.Symbol(&_Erc721meta.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Erc721meta *Erc721metaCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Erc721meta.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Erc721meta *Erc721metaSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Erc721meta.Contract.TokenURI(&_Erc721meta.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Erc721meta *Erc721metaCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Erc721meta.Contract.TokenURI(&_Erc721meta.CallOpts, _tokenId)
}
