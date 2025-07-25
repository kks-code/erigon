// Code generated by abigen. DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"

	ethereum "github.com/erigontech/erigon"
	"github.com/erigontech/erigon-lib/abi"
	"github.com/erigontech/erigon-lib/common"
	"github.com/erigontech/erigon-lib/types"
	"github.com/erigontech/erigon/p2p/event"
	"github.com/erigontech/erigon/execution/abi/bind"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = fmt.Errorf
	_ = reflect.ValueOf
)

// TestcontractABI is the input ABI used to generate the binding from.
const TestcontractABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"createAndException\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"createAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeAndException\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"updateAndException\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"updateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TestcontractBin is the compiled bytecode used for deploying new contracts.
var TestcontractBin = "0x608060405234801561001057600080fd5b503360009081526020819052604090206064905561023a806100336000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063c2ce0ef711610066578063c2ce0ef714610114578063c53e5ae314610127578063cb946a0714610127578063d592ed1f1461013a578063f64c050d1461011457600080fd5b806327e235e3146100a3578063660cc200146100d5578063780900dc146100df57806382ab890a146100df578063a7f43779146100ff575b600080fd5b6100c36100b13660046101a5565b60006020819052908152604090205481565b60405190815260200160405180910390f35b6100dd610142565b005b6100dd6100ed3660046101d5565b33600090815260208190526040902055565b6100dd33600090815260208190526040812055565b6100dd6101223660046101d5565b61015c565b6100dd6101353660046101d5565b610179565b6100dd610190565b3360009081526020819052604081205561015a6101ee565b565b3360009081526020819052604090208190556101766101ee565b50565b336000908152602081905260408120829055819080fd5b33600090815260208190526040812081905580fd5b6000602082840312156101b757600080fd5b81356001600160a01b03811681146101ce57600080fd5b9392505050565b6000602082840312156101e757600080fd5b5035919050565b634e487b7160e01b600052600160045260246000fdfea2646970667358221220cbc0ec1bd2504a9e5d6a3f3a8b92de36a1af219062eec962f407d937d0acadc664736f6c63430008130033"

// DeployTestcontract deploys a new Ethereum contract, binding an instance of Testcontract to it.
func DeployTestcontract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, types.Transaction, *Testcontract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestcontractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestcontractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Testcontract{TestcontractCaller: TestcontractCaller{contract: contract}, TestcontractTransactor: TestcontractTransactor{contract: contract}, TestcontractFilterer: TestcontractFilterer{contract: contract}}, nil
}

// Testcontract is an auto generated Go binding around an Ethereum contract.
type Testcontract struct {
	TestcontractCaller     // Read-only binding to the contract
	TestcontractTransactor // Write-only binding to the contract
	TestcontractFilterer   // Log filterer for contract events
}

// TestcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestcontractSession struct {
	Contract     *Testcontract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestcontractCallerSession struct {
	Contract *TestcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TestcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestcontractTransactorSession struct {
	Contract     *TestcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TestcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestcontractRaw struct {
	Contract *Testcontract // Generic contract binding to access the raw methods on
}

// TestcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestcontractCallerRaw struct {
	Contract *TestcontractCaller // Generic read-only contract binding to access the raw methods on
}

// TestcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestcontractTransactorRaw struct {
	Contract *TestcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestcontract creates a new instance of Testcontract, bound to a specific deployed contract.
func NewTestcontract(address common.Address, backend bind.ContractBackend) (*Testcontract, error) {
	contract, err := bindTestcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Testcontract{TestcontractCaller: TestcontractCaller{contract: contract}, TestcontractTransactor: TestcontractTransactor{contract: contract}, TestcontractFilterer: TestcontractFilterer{contract: contract}}, nil
}

// NewTestcontractCaller creates a new read-only instance of Testcontract, bound to a specific deployed contract.
func NewTestcontractCaller(address common.Address, caller bind.ContractCaller) (*TestcontractCaller, error) {
	contract, err := bindTestcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestcontractCaller{contract: contract}, nil
}

// NewTestcontractTransactor creates a new write-only instance of Testcontract, bound to a specific deployed contract.
func NewTestcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*TestcontractTransactor, error) {
	contract, err := bindTestcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestcontractTransactor{contract: contract}, nil
}

// NewTestcontractFilterer creates a new log filterer instance of Testcontract, bound to a specific deployed contract.
func NewTestcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*TestcontractFilterer, error) {
	contract, err := bindTestcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestcontractFilterer{contract: contract}, nil
}

// bindTestcontract binds a generic wrapper to an already deployed contract.
func bindTestcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testcontract *TestcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Testcontract.Contract.TestcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testcontract *TestcontractRaw) Transfer(opts *bind.TransactOpts) (types.Transaction, error) {
	return _Testcontract.Contract.TestcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testcontract *TestcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (types.Transaction, error) {
	return _Testcontract.Contract.TestcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testcontract *TestcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Testcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testcontract *TestcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (types.Transaction, error) {
	return _Testcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testcontract *TestcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (types.Transaction, error) {
	return _Testcontract.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Testcontract *TestcontractCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Testcontract.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Testcontract *TestcontractSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Testcontract.Contract.Balances(&_Testcontract.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Testcontract *TestcontractCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Testcontract.Contract.Balances(&_Testcontract.CallOpts, arg0)
}

// Create is a paid mutator transaction binding the contract method 0x780900dc.
//
// Solidity: function create(uint256 newBalance) returns()
func (_Testcontract *TestcontractTransactor) Create(opts *bind.TransactOpts, newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.contract.Transact(opts, "create", newBalance)
}

// Create is a paid mutator transaction binding the contract method 0x780900dc.
//
// Solidity: function create(uint256 newBalance) returns()
func (_Testcontract *TestcontractSession) Create(newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.Contract.Create(&_Testcontract.TransactOpts, newBalance)
}

// Create is a paid mutator transaction binding the contract method 0x780900dc.
//
// Solidity: function create(uint256 newBalance) returns()
func (_Testcontract *TestcontractTransactorSession) Create(newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.Contract.Create(&_Testcontract.TransactOpts, newBalance)
}

// CreateAndException is a paid mutator transaction binding the contract method 0xc2ce0ef7.
//
// Solidity: function createAndException(uint256 newBalance) returns()
func (_Testcontract *TestcontractTransactor) CreateAndException(opts *bind.TransactOpts, newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.contract.Transact(opts, "createAndException", newBalance)
}

// CreateAndException is a paid mutator transaction binding the contract method 0xc2ce0ef7.
//
// Solidity: function createAndException(uint256 newBalance) returns()
func (_Testcontract *TestcontractSession) CreateAndException(newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.Contract.CreateAndException(&_Testcontract.TransactOpts, newBalance)
}

// CreateAndException is a paid mutator transaction binding the contract method 0xc2ce0ef7.
//
// Solidity: function createAndException(uint256 newBalance) returns()
func (_Testcontract *TestcontractTransactorSession) CreateAndException(newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.Contract.CreateAndException(&_Testcontract.TransactOpts, newBalance)
}

// CreateAndRevert is a paid mutator transaction binding the contract method 0xc53e5ae3.
//
// Solidity: function createAndRevert(uint256 newBalance) returns()
func (_Testcontract *TestcontractTransactor) CreateAndRevert(opts *bind.TransactOpts, newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.contract.Transact(opts, "createAndRevert", newBalance)
}

// CreateAndRevert is a paid mutator transaction binding the contract method 0xc53e5ae3.
//
// Solidity: function createAndRevert(uint256 newBalance) returns()
func (_Testcontract *TestcontractSession) CreateAndRevert(newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.Contract.CreateAndRevert(&_Testcontract.TransactOpts, newBalance)
}

// CreateAndRevert is a paid mutator transaction binding the contract method 0xc53e5ae3.
//
// Solidity: function createAndRevert(uint256 newBalance) returns()
func (_Testcontract *TestcontractTransactorSession) CreateAndRevert(newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.Contract.CreateAndRevert(&_Testcontract.TransactOpts, newBalance)
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Testcontract *TestcontractTransactor) Remove(opts *bind.TransactOpts) (types.Transaction, error) {
	return _Testcontract.contract.Transact(opts, "remove")
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Testcontract *TestcontractSession) Remove() (types.Transaction, error) {
	return _Testcontract.Contract.Remove(&_Testcontract.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Testcontract *TestcontractTransactorSession) Remove() (types.Transaction, error) {
	return _Testcontract.Contract.Remove(&_Testcontract.TransactOpts)
}

// RemoveAndException is a paid mutator transaction binding the contract method 0x660cc200.
//
// Solidity: function removeAndException() returns()
func (_Testcontract *TestcontractTransactor) RemoveAndException(opts *bind.TransactOpts) (types.Transaction, error) {
	return _Testcontract.contract.Transact(opts, "removeAndException")
}

// RemoveAndException is a paid mutator transaction binding the contract method 0x660cc200.
//
// Solidity: function removeAndException() returns()
func (_Testcontract *TestcontractSession) RemoveAndException() (types.Transaction, error) {
	return _Testcontract.Contract.RemoveAndException(&_Testcontract.TransactOpts)
}

// RemoveAndException is a paid mutator transaction binding the contract method 0x660cc200.
//
// Solidity: function removeAndException() returns()
func (_Testcontract *TestcontractTransactorSession) RemoveAndException() (types.Transaction, error) {
	return _Testcontract.Contract.RemoveAndException(&_Testcontract.TransactOpts)
}

// RemoveAndRevert is a paid mutator transaction binding the contract method 0xd592ed1f.
//
// Solidity: function removeAndRevert() returns()
func (_Testcontract *TestcontractTransactor) RemoveAndRevert(opts *bind.TransactOpts) (types.Transaction, error) {
	return _Testcontract.contract.Transact(opts, "removeAndRevert")
}

// RemoveAndRevert is a paid mutator transaction binding the contract method 0xd592ed1f.
//
// Solidity: function removeAndRevert() returns()
func (_Testcontract *TestcontractSession) RemoveAndRevert() (types.Transaction, error) {
	return _Testcontract.Contract.RemoveAndRevert(&_Testcontract.TransactOpts)
}

// RemoveAndRevert is a paid mutator transaction binding the contract method 0xd592ed1f.
//
// Solidity: function removeAndRevert() returns()
func (_Testcontract *TestcontractTransactorSession) RemoveAndRevert() (types.Transaction, error) {
	return _Testcontract.Contract.RemoveAndRevert(&_Testcontract.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(uint256 newBalance) returns()
func (_Testcontract *TestcontractTransactor) Update(opts *bind.TransactOpts, newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.contract.Transact(opts, "update", newBalance)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(uint256 newBalance) returns()
func (_Testcontract *TestcontractSession) Update(newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.Contract.Update(&_Testcontract.TransactOpts, newBalance)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(uint256 newBalance) returns()
func (_Testcontract *TestcontractTransactorSession) Update(newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.Contract.Update(&_Testcontract.TransactOpts, newBalance)
}

// UpdateAndException is a paid mutator transaction binding the contract method 0xf64c050d.
//
// Solidity: function updateAndException(uint256 newBalance) returns()
func (_Testcontract *TestcontractTransactor) UpdateAndException(opts *bind.TransactOpts, newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.contract.Transact(opts, "updateAndException", newBalance)
}

// UpdateAndException is a paid mutator transaction binding the contract method 0xf64c050d.
//
// Solidity: function updateAndException(uint256 newBalance) returns()
func (_Testcontract *TestcontractSession) UpdateAndException(newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.Contract.UpdateAndException(&_Testcontract.TransactOpts, newBalance)
}

// UpdateAndException is a paid mutator transaction binding the contract method 0xf64c050d.
//
// Solidity: function updateAndException(uint256 newBalance) returns()
func (_Testcontract *TestcontractTransactorSession) UpdateAndException(newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.Contract.UpdateAndException(&_Testcontract.TransactOpts, newBalance)
}

// UpdateAndRevert is a paid mutator transaction binding the contract method 0xcb946a07.
//
// Solidity: function updateAndRevert(uint256 newBalance) returns()
func (_Testcontract *TestcontractTransactor) UpdateAndRevert(opts *bind.TransactOpts, newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.contract.Transact(opts, "updateAndRevert", newBalance)
}

// UpdateAndRevert is a paid mutator transaction binding the contract method 0xcb946a07.
//
// Solidity: function updateAndRevert(uint256 newBalance) returns()
func (_Testcontract *TestcontractSession) UpdateAndRevert(newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.Contract.UpdateAndRevert(&_Testcontract.TransactOpts, newBalance)
}

// UpdateAndRevert is a paid mutator transaction binding the contract method 0xcb946a07.
//
// Solidity: function updateAndRevert(uint256 newBalance) returns()
func (_Testcontract *TestcontractTransactorSession) UpdateAndRevert(newBalance *big.Int) (types.Transaction, error) {
	return _Testcontract.Contract.UpdateAndRevert(&_Testcontract.TransactOpts, newBalance)
}

// TestcontractCreateParams is an auto generated read-only Go binding of transcaction calldata params
type TestcontractCreateParams struct {
	Param_newBalance *big.Int
}

// Parse Create method from calldata of a transaction
//
// Solidity: function create(uint256 newBalance) returns()
func ParseTestcontractCreateParams(calldata []byte) (*TestcontractCreateParams, error) {
	if len(calldata) <= 4 {
		return nil, fmt.Errorf("invalid calldata input")
	}

	_abi, err := abi.JSON(strings.NewReader(TestcontractABI))
	if err != nil {
		return nil, fmt.Errorf("failed to get abi of registry metadata: %w", err)
	}

	out, err := _abi.Methods["create"].Inputs.Unpack(calldata[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack create params data: %w", err)
	}

	var paramsResult = new(TestcontractCreateParams)
	value := reflect.ValueOf(paramsResult).Elem()

	if value.NumField() != len(out) {
		return nil, fmt.Errorf("failed to match calldata with param field number")
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &TestcontractCreateParams{
		Param_newBalance: out0,
	}, nil
}

// TestcontractCreateAndExceptionParams is an auto generated read-only Go binding of transcaction calldata params
type TestcontractCreateAndExceptionParams struct {
	Param_newBalance *big.Int
}

// Parse CreateAndException method from calldata of a transaction
//
// Solidity: function createAndException(uint256 newBalance) returns()
func ParseTestcontractCreateAndExceptionParams(calldata []byte) (*TestcontractCreateAndExceptionParams, error) {
	if len(calldata) <= 4 {
		return nil, fmt.Errorf("invalid calldata input")
	}

	_abi, err := abi.JSON(strings.NewReader(TestcontractABI))
	if err != nil {
		return nil, fmt.Errorf("failed to get abi of registry metadata: %w", err)
	}

	out, err := _abi.Methods["createAndException"].Inputs.Unpack(calldata[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack createAndException params data: %w", err)
	}

	var paramsResult = new(TestcontractCreateAndExceptionParams)
	value := reflect.ValueOf(paramsResult).Elem()

	if value.NumField() != len(out) {
		return nil, fmt.Errorf("failed to match calldata with param field number")
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &TestcontractCreateAndExceptionParams{
		Param_newBalance: out0,
	}, nil
}

// TestcontractCreateAndRevertParams is an auto generated read-only Go binding of transcaction calldata params
type TestcontractCreateAndRevertParams struct {
	Param_newBalance *big.Int
}

// Parse CreateAndRevert method from calldata of a transaction
//
// Solidity: function createAndRevert(uint256 newBalance) returns()
func ParseTestcontractCreateAndRevertParams(calldata []byte) (*TestcontractCreateAndRevertParams, error) {
	if len(calldata) <= 4 {
		return nil, fmt.Errorf("invalid calldata input")
	}

	_abi, err := abi.JSON(strings.NewReader(TestcontractABI))
	if err != nil {
		return nil, fmt.Errorf("failed to get abi of registry metadata: %w", err)
	}

	out, err := _abi.Methods["createAndRevert"].Inputs.Unpack(calldata[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack createAndRevert params data: %w", err)
	}

	var paramsResult = new(TestcontractCreateAndRevertParams)
	value := reflect.ValueOf(paramsResult).Elem()

	if value.NumField() != len(out) {
		return nil, fmt.Errorf("failed to match calldata with param field number")
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &TestcontractCreateAndRevertParams{
		Param_newBalance: out0,
	}, nil
}

// TestcontractUpdateParams is an auto generated read-only Go binding of transcaction calldata params
type TestcontractUpdateParams struct {
	Param_newBalance *big.Int
}

// Parse Update method from calldata of a transaction
//
// Solidity: function update(uint256 newBalance) returns()
func ParseTestcontractUpdateParams(calldata []byte) (*TestcontractUpdateParams, error) {
	if len(calldata) <= 4 {
		return nil, fmt.Errorf("invalid calldata input")
	}

	_abi, err := abi.JSON(strings.NewReader(TestcontractABI))
	if err != nil {
		return nil, fmt.Errorf("failed to get abi of registry metadata: %w", err)
	}

	out, err := _abi.Methods["update"].Inputs.Unpack(calldata[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack update params data: %w", err)
	}

	var paramsResult = new(TestcontractUpdateParams)
	value := reflect.ValueOf(paramsResult).Elem()

	if value.NumField() != len(out) {
		return nil, fmt.Errorf("failed to match calldata with param field number")
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &TestcontractUpdateParams{
		Param_newBalance: out0,
	}, nil
}

// TestcontractUpdateAndExceptionParams is an auto generated read-only Go binding of transcaction calldata params
type TestcontractUpdateAndExceptionParams struct {
	Param_newBalance *big.Int
}

// Parse UpdateAndException method from calldata of a transaction
//
// Solidity: function updateAndException(uint256 newBalance) returns()
func ParseTestcontractUpdateAndExceptionParams(calldata []byte) (*TestcontractUpdateAndExceptionParams, error) {
	if len(calldata) <= 4 {
		return nil, fmt.Errorf("invalid calldata input")
	}

	_abi, err := abi.JSON(strings.NewReader(TestcontractABI))
	if err != nil {
		return nil, fmt.Errorf("failed to get abi of registry metadata: %w", err)
	}

	out, err := _abi.Methods["updateAndException"].Inputs.Unpack(calldata[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack updateAndException params data: %w", err)
	}

	var paramsResult = new(TestcontractUpdateAndExceptionParams)
	value := reflect.ValueOf(paramsResult).Elem()

	if value.NumField() != len(out) {
		return nil, fmt.Errorf("failed to match calldata with param field number")
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &TestcontractUpdateAndExceptionParams{
		Param_newBalance: out0,
	}, nil
}

// TestcontractUpdateAndRevertParams is an auto generated read-only Go binding of transcaction calldata params
type TestcontractUpdateAndRevertParams struct {
	Param_newBalance *big.Int
}

// Parse UpdateAndRevert method from calldata of a transaction
//
// Solidity: function updateAndRevert(uint256 newBalance) returns()
func ParseTestcontractUpdateAndRevertParams(calldata []byte) (*TestcontractUpdateAndRevertParams, error) {
	if len(calldata) <= 4 {
		return nil, fmt.Errorf("invalid calldata input")
	}

	_abi, err := abi.JSON(strings.NewReader(TestcontractABI))
	if err != nil {
		return nil, fmt.Errorf("failed to get abi of registry metadata: %w", err)
	}

	out, err := _abi.Methods["updateAndRevert"].Inputs.Unpack(calldata[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack updateAndRevert params data: %w", err)
	}

	var paramsResult = new(TestcontractUpdateAndRevertParams)
	value := reflect.ValueOf(paramsResult).Elem()

	if value.NumField() != len(out) {
		return nil, fmt.Errorf("failed to match calldata with param field number")
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &TestcontractUpdateAndRevertParams{
		Param_newBalance: out0,
	}, nil
}
