// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc1155meta

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

// Erc1155metaMetaData contains all meta data concerning the Erc1155meta contract.
var Erc1155metaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Erc1155metaABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc1155metaMetaData.ABI instead.
var Erc1155metaABI = Erc1155metaMetaData.ABI

// Erc1155meta is an auto generated Go binding around an Ethereum contract.
type Erc1155meta struct {
	Erc1155metaCaller     // Read-only binding to the contract
	Erc1155metaTransactor // Write-only binding to the contract
	Erc1155metaFilterer   // Log filterer for contract events
}

// Erc1155metaCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc1155metaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1155metaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc1155metaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1155metaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc1155metaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1155metaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc1155metaSession struct {
	Contract     *Erc1155meta      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc1155metaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc1155metaCallerSession struct {
	Contract *Erc1155metaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Erc1155metaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc1155metaTransactorSession struct {
	Contract     *Erc1155metaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Erc1155metaRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc1155metaRaw struct {
	Contract *Erc1155meta // Generic contract binding to access the raw methods on
}

// Erc1155metaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc1155metaCallerRaw struct {
	Contract *Erc1155metaCaller // Generic read-only contract binding to access the raw methods on
}

// Erc1155metaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc1155metaTransactorRaw struct {
	Contract *Erc1155metaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc1155meta creates a new instance of Erc1155meta, bound to a specific deployed contract.
func NewErc1155meta(address common.Address, backend bind.ContractBackend) (*Erc1155meta, error) {
	contract, err := bindErc1155meta(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc1155meta{Erc1155metaCaller: Erc1155metaCaller{contract: contract}, Erc1155metaTransactor: Erc1155metaTransactor{contract: contract}, Erc1155metaFilterer: Erc1155metaFilterer{contract: contract}}, nil
}

// NewErc1155metaCaller creates a new read-only instance of Erc1155meta, bound to a specific deployed contract.
func NewErc1155metaCaller(address common.Address, caller bind.ContractCaller) (*Erc1155metaCaller, error) {
	contract, err := bindErc1155meta(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc1155metaCaller{contract: contract}, nil
}

// NewErc1155metaTransactor creates a new write-only instance of Erc1155meta, bound to a specific deployed contract.
func NewErc1155metaTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc1155metaTransactor, error) {
	contract, err := bindErc1155meta(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc1155metaTransactor{contract: contract}, nil
}

// NewErc1155metaFilterer creates a new log filterer instance of Erc1155meta, bound to a specific deployed contract.
func NewErc1155metaFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc1155metaFilterer, error) {
	contract, err := bindErc1155meta(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc1155metaFilterer{contract: contract}, nil
}

// bindErc1155meta binds a generic wrapper to an already deployed contract.
func bindErc1155meta(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc1155metaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc1155meta *Erc1155metaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc1155meta.Contract.Erc1155metaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc1155meta *Erc1155metaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1155meta.Contract.Erc1155metaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc1155meta *Erc1155metaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc1155meta.Contract.Erc1155metaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc1155meta *Erc1155metaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc1155meta.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc1155meta *Erc1155metaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1155meta.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc1155meta *Erc1155metaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc1155meta.Contract.contract.Transact(opts, method, params...)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_Erc1155meta *Erc1155metaCaller) Uri(opts *bind.CallOpts, _id *big.Int) (string, error) {
	var out []interface{}
	err := _Erc1155meta.contract.Call(opts, &out, "uri", _id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_Erc1155meta *Erc1155metaSession) Uri(_id *big.Int) (string, error) {
	return _Erc1155meta.Contract.Uri(&_Erc1155meta.CallOpts, _id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_Erc1155meta *Erc1155metaCallerSession) Uri(_id *big.Int) (string, error) {
	return _Erc1155meta.Contract.Uri(&_Erc1155meta.CallOpts, _id)
}
