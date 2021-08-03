// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC20DetailedABI is the input ABI used to generate the binding from.
const ERC20DetailedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ERC20DetailedFuncSigs maps the 4-byte function signature to its string representation.
var ERC20DetailedFuncSigs = map[string]string{
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
}

// ERC20DetailedBin is the compiled bytecode used for deploying new contracts.
var ERC20DetailedBin = "0x608060405234801561001057600080fd5b5060405161049f38038061049f8339818101604052606081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825164010000000081118282018810171561008257600080fd5b82525081516020918201929091019080838360005b838110156100af578181015183820152602001610097565b50505050905090810190601f1680156100dc5780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156100ff57600080fd5b90830190602082018581111561011457600080fd5b825164010000000081118282018810171561012e57600080fd5b82525081516020918201929091019080838360005b8381101561015b578181015183820152602001610143565b50505050905090810190601f1680156101885780820380516001836020036101000a031916815260200191505b5060405260209081015185519093506101a792506000918601906101d8565b5081516101bb9060019060208501906101d8565b506002805460ff191660ff92909216919091179055506102739050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061021957805160ff1916838001178555610246565b82800160010185558215610246579182015b8281111561024657825182559160200191906001019061022b565b50610252929150610256565b5090565b61027091905b80821115610252576000815560010161025c565b90565b61021d806102826000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806306fdde0314610046578063313ce567146100c357806395d89b41146100e1575b600080fd5b61004e6100e9565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610088578181015183820152602001610070565b50505050905090810190601f1680156100b55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100cb61017f565b6040805160ff9092168252519081900360200190f35b61004e610188565b60008054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156101755780601f1061014a57610100808354040283529160200191610175565b820191906000526020600020905b81548152906001019060200180831161015857829003601f168201915b5050505050905090565b60025460ff1690565b60018054604080516020601f600260001961010087891615020190951694909404938401819004810282018101909252828152606093909290918301828280156101755780601f1061014a5761010080835404028352916020019161017556fea265627a7a72315820d125a9eba67b478da452e6d828c3f8dbb2f2a4f0bcf72555eec937dc9059706464736f6c634300050b0032"

// DeployERC20Detailed deploys a new Ethereum contract, binding an instance of ERC20Detailed to it.
func DeployERC20Detailed(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals uint8) (common.Address, *types.Transaction, *ERC20Detailed, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20DetailedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20DetailedBin), backend, name, symbol, decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Detailed{ERC20DetailedCaller: ERC20DetailedCaller{contract: contract}, ERC20DetailedTransactor: ERC20DetailedTransactor{contract: contract}, ERC20DetailedFilterer: ERC20DetailedFilterer{contract: contract}}, nil
}

// ERC20Detailed is an auto generated Go binding around an Ethereum contract.
type ERC20Detailed struct {
	ERC20DetailedCaller     // Read-only binding to the contract
	ERC20DetailedTransactor // Write-only binding to the contract
	ERC20DetailedFilterer   // Log filterer for contract events
}

// ERC20DetailedCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20DetailedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20DetailedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20DetailedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20DetailedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20DetailedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20DetailedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20DetailedSession struct {
	Contract     *ERC20Detailed    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20DetailedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20DetailedCallerSession struct {
	Contract *ERC20DetailedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC20DetailedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20DetailedTransactorSession struct {
	Contract     *ERC20DetailedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20DetailedRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20DetailedRaw struct {
	Contract *ERC20Detailed // Generic contract binding to access the raw methods on
}

// ERC20DetailedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20DetailedCallerRaw struct {
	Contract *ERC20DetailedCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20DetailedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20DetailedTransactorRaw struct {
	Contract *ERC20DetailedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Detailed creates a new instance of ERC20Detailed, bound to a specific deployed contract.
func NewERC20Detailed(address common.Address, backend bind.ContractBackend) (*ERC20Detailed, error) {
	contract, err := bindERC20Detailed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Detailed{ERC20DetailedCaller: ERC20DetailedCaller{contract: contract}, ERC20DetailedTransactor: ERC20DetailedTransactor{contract: contract}, ERC20DetailedFilterer: ERC20DetailedFilterer{contract: contract}}, nil
}

// NewERC20DetailedCaller creates a new read-only instance of ERC20Detailed, bound to a specific deployed contract.
func NewERC20DetailedCaller(address common.Address, caller bind.ContractCaller) (*ERC20DetailedCaller, error) {
	contract, err := bindERC20Detailed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20DetailedCaller{contract: contract}, nil
}

// NewERC20DetailedTransactor creates a new write-only instance of ERC20Detailed, bound to a specific deployed contract.
func NewERC20DetailedTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20DetailedTransactor, error) {
	contract, err := bindERC20Detailed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20DetailedTransactor{contract: contract}, nil
}

// NewERC20DetailedFilterer creates a new log filterer instance of ERC20Detailed, bound to a specific deployed contract.
func NewERC20DetailedFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20DetailedFilterer, error) {
	contract, err := bindERC20Detailed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20DetailedFilterer{contract: contract}, nil
}

// bindERC20Detailed binds a generic wrapper to an already deployed contract.
func bindERC20Detailed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20DetailedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Detailed *ERC20DetailedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Detailed.Contract.ERC20DetailedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Detailed *ERC20DetailedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.ERC20DetailedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Detailed *ERC20DetailedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.ERC20DetailedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Detailed *ERC20DetailedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Detailed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Detailed *ERC20DetailedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Detailed *ERC20DetailedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Detailed *ERC20DetailedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC20Detailed.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Detailed *ERC20DetailedSession) Decimals() (uint8, error) {
	return _ERC20Detailed.Contract.Decimals(&_ERC20Detailed.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Detailed *ERC20DetailedCallerSession) Decimals() (uint8, error) {
	return _ERC20Detailed.Contract.Decimals(&_ERC20Detailed.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20Detailed *ERC20DetailedCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20Detailed.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20Detailed *ERC20DetailedSession) Name() (string, error) {
	return _ERC20Detailed.Contract.Name(&_ERC20Detailed.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20Detailed *ERC20DetailedCallerSession) Name() (string, error) {
	return _ERC20Detailed.Contract.Name(&_ERC20Detailed.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20Detailed *ERC20DetailedCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20Detailed.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20Detailed *ERC20DetailedSession) Symbol() (string, error) {
	return _ERC20Detailed.Contract.Symbol(&_ERC20Detailed.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20Detailed *ERC20DetailedCallerSession) Symbol() (string, error) {
	return _ERC20Detailed.Contract.Symbol(&_ERC20Detailed.CallOpts)
}

// FrozenABI is the input ABI used to generate the binding from.
const FrozenABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getFrozenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_frozenBalance\",\"type\":\"uint256\"}],\"name\":\"FrozenIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_frozenBalance\",\"type\":\"uint256\"}],\"name\":\"FrozenDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_frozenBalance\",\"type\":\"uint256\"}],\"name\":\"Cashout\",\"type\":\"event\"}]"

// FrozenFuncSigs maps the 4-byte function signature to its string representation.
var FrozenFuncSigs = map[string]string{
	"9f2cfaf1": "getFrozenBalance(address)",
}

// FrozenBin is the compiled bytecode used for deploying new contracts.
var FrozenBin = "0x6080604052348015600f57600080fd5b5060b28061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80639f2cfaf114602d575b600080fd5b605060048036036020811015604157600080fd5b50356001600160a01b03166062565b60408051918252519081900360200190f35b6001600160a01b03166000908152602081905260409020549056fea265627a7a72315820f73c50b26852ef2df3da896c95c5e999d1de07e79a2ce2bf78739829f1955cf864736f6c634300050b0032"

// DeployFrozen deploys a new Ethereum contract, binding an instance of Frozen to it.
func DeployFrozen(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Frozen, error) {
	parsed, err := abi.JSON(strings.NewReader(FrozenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FrozenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Frozen{FrozenCaller: FrozenCaller{contract: contract}, FrozenTransactor: FrozenTransactor{contract: contract}, FrozenFilterer: FrozenFilterer{contract: contract}}, nil
}

// Frozen is an auto generated Go binding around an Ethereum contract.
type Frozen struct {
	FrozenCaller     // Read-only binding to the contract
	FrozenTransactor // Write-only binding to the contract
	FrozenFilterer   // Log filterer for contract events
}

// FrozenCaller is an auto generated read-only Go binding around an Ethereum contract.
type FrozenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FrozenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FrozenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FrozenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FrozenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FrozenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FrozenSession struct {
	Contract     *Frozen           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FrozenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FrozenCallerSession struct {
	Contract *FrozenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FrozenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FrozenTransactorSession struct {
	Contract     *FrozenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FrozenRaw is an auto generated low-level Go binding around an Ethereum contract.
type FrozenRaw struct {
	Contract *Frozen // Generic contract binding to access the raw methods on
}

// FrozenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FrozenCallerRaw struct {
	Contract *FrozenCaller // Generic read-only contract binding to access the raw methods on
}

// FrozenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FrozenTransactorRaw struct {
	Contract *FrozenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFrozen creates a new instance of Frozen, bound to a specific deployed contract.
func NewFrozen(address common.Address, backend bind.ContractBackend) (*Frozen, error) {
	contract, err := bindFrozen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Frozen{FrozenCaller: FrozenCaller{contract: contract}, FrozenTransactor: FrozenTransactor{contract: contract}, FrozenFilterer: FrozenFilterer{contract: contract}}, nil
}

// NewFrozenCaller creates a new read-only instance of Frozen, bound to a specific deployed contract.
func NewFrozenCaller(address common.Address, caller bind.ContractCaller) (*FrozenCaller, error) {
	contract, err := bindFrozen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FrozenCaller{contract: contract}, nil
}

// NewFrozenTransactor creates a new write-only instance of Frozen, bound to a specific deployed contract.
func NewFrozenTransactor(address common.Address, transactor bind.ContractTransactor) (*FrozenTransactor, error) {
	contract, err := bindFrozen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FrozenTransactor{contract: contract}, nil
}

// NewFrozenFilterer creates a new log filterer instance of Frozen, bound to a specific deployed contract.
func NewFrozenFilterer(address common.Address, filterer bind.ContractFilterer) (*FrozenFilterer, error) {
	contract, err := bindFrozen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FrozenFilterer{contract: contract}, nil
}

// bindFrozen binds a generic wrapper to an already deployed contract.
func bindFrozen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FrozenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Frozen *FrozenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Frozen.Contract.FrozenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Frozen *FrozenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Frozen.Contract.FrozenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Frozen *FrozenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Frozen.Contract.FrozenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Frozen *FrozenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Frozen.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Frozen *FrozenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Frozen.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Frozen *FrozenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Frozen.Contract.contract.Transact(opts, method, params...)
}

// GetFrozenBalance is a free data retrieval call binding the contract method 0x9f2cfaf1.
//
// Solidity: function getFrozenBalance(address _validator) constant returns(uint256)
func (_Frozen *FrozenCaller) GetFrozenBalance(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Frozen.contract.Call(opts, out, "getFrozenBalance", _validator)
	return *ret0, err
}

// GetFrozenBalance is a free data retrieval call binding the contract method 0x9f2cfaf1.
//
// Solidity: function getFrozenBalance(address _validator) constant returns(uint256)
func (_Frozen *FrozenSession) GetFrozenBalance(_validator common.Address) (*big.Int, error) {
	return _Frozen.Contract.GetFrozenBalance(&_Frozen.CallOpts, _validator)
}

// GetFrozenBalance is a free data retrieval call binding the contract method 0x9f2cfaf1.
//
// Solidity: function getFrozenBalance(address _validator) constant returns(uint256)
func (_Frozen *FrozenCallerSession) GetFrozenBalance(_validator common.Address) (*big.Int, error) {
	return _Frozen.Contract.GetFrozenBalance(&_Frozen.CallOpts, _validator)
}

// FrozenCashoutIterator is returned from FilterCashout and is used to iterate over the raw logs and unpacked data for Cashout events raised by the Frozen contract.
type FrozenCashoutIterator struct {
	Event *FrozenCashout // Event containing the contract specifics and raw log

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
func (it *FrozenCashoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FrozenCashout)
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
		it.Event = new(FrozenCashout)
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
func (it *FrozenCashoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FrozenCashoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FrozenCashout represents a Cashout event raised by the Frozen contract.
type FrozenCashout struct {
	Validator     common.Address
	Amount        *big.Int
	FrozenBalance *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCashout is a free log retrieval operation binding the contract event 0x4f32a1005fcf45a64a68bc4258b4e0f5522b7aa6a06506781aa4e6c6395c48d0.
//
// Solidity: event Cashout(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Frozen *FrozenFilterer) FilterCashout(opts *bind.FilterOpts, _validator []common.Address) (*FrozenCashoutIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Frozen.contract.FilterLogs(opts, "Cashout", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &FrozenCashoutIterator{contract: _Frozen.contract, event: "Cashout", logs: logs, sub: sub}, nil
}

// WatchCashout is a free log subscription operation binding the contract event 0x4f32a1005fcf45a64a68bc4258b4e0f5522b7aa6a06506781aa4e6c6395c48d0.
//
// Solidity: event Cashout(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Frozen *FrozenFilterer) WatchCashout(opts *bind.WatchOpts, sink chan<- *FrozenCashout, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Frozen.contract.WatchLogs(opts, "Cashout", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FrozenCashout)
				if err := _Frozen.contract.UnpackLog(event, "Cashout", log); err != nil {
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

// ParseCashout is a log parse operation binding the contract event 0x4f32a1005fcf45a64a68bc4258b4e0f5522b7aa6a06506781aa4e6c6395c48d0.
//
// Solidity: event Cashout(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Frozen *FrozenFilterer) ParseCashout(log types.Log) (*FrozenCashout, error) {
	event := new(FrozenCashout)
	if err := _Frozen.contract.UnpackLog(event, "Cashout", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FrozenFrozenDecreasedIterator is returned from FilterFrozenDecreased and is used to iterate over the raw logs and unpacked data for FrozenDecreased events raised by the Frozen contract.
type FrozenFrozenDecreasedIterator struct {
	Event *FrozenFrozenDecreased // Event containing the contract specifics and raw log

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
func (it *FrozenFrozenDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FrozenFrozenDecreased)
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
		it.Event = new(FrozenFrozenDecreased)
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
func (it *FrozenFrozenDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FrozenFrozenDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FrozenFrozenDecreased represents a FrozenDecreased event raised by the Frozen contract.
type FrozenFrozenDecreased struct {
	Validator     common.Address
	Amount        *big.Int
	FrozenBalance *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFrozenDecreased is a free log retrieval operation binding the contract event 0x366703a486081d66a8ef7d8a544aa1f2a91b5f4906e9d5807c3f3ce27ee2d590.
//
// Solidity: event FrozenDecreased(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Frozen *FrozenFilterer) FilterFrozenDecreased(opts *bind.FilterOpts, _validator []common.Address) (*FrozenFrozenDecreasedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Frozen.contract.FilterLogs(opts, "FrozenDecreased", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &FrozenFrozenDecreasedIterator{contract: _Frozen.contract, event: "FrozenDecreased", logs: logs, sub: sub}, nil
}

// WatchFrozenDecreased is a free log subscription operation binding the contract event 0x366703a486081d66a8ef7d8a544aa1f2a91b5f4906e9d5807c3f3ce27ee2d590.
//
// Solidity: event FrozenDecreased(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Frozen *FrozenFilterer) WatchFrozenDecreased(opts *bind.WatchOpts, sink chan<- *FrozenFrozenDecreased, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Frozen.contract.WatchLogs(opts, "FrozenDecreased", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FrozenFrozenDecreased)
				if err := _Frozen.contract.UnpackLog(event, "FrozenDecreased", log); err != nil {
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

// ParseFrozenDecreased is a log parse operation binding the contract event 0x366703a486081d66a8ef7d8a544aa1f2a91b5f4906e9d5807c3f3ce27ee2d590.
//
// Solidity: event FrozenDecreased(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Frozen *FrozenFilterer) ParseFrozenDecreased(log types.Log) (*FrozenFrozenDecreased, error) {
	event := new(FrozenFrozenDecreased)
	if err := _Frozen.contract.UnpackLog(event, "FrozenDecreased", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FrozenFrozenIncreasedIterator is returned from FilterFrozenIncreased and is used to iterate over the raw logs and unpacked data for FrozenIncreased events raised by the Frozen contract.
type FrozenFrozenIncreasedIterator struct {
	Event *FrozenFrozenIncreased // Event containing the contract specifics and raw log

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
func (it *FrozenFrozenIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FrozenFrozenIncreased)
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
		it.Event = new(FrozenFrozenIncreased)
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
func (it *FrozenFrozenIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FrozenFrozenIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FrozenFrozenIncreased represents a FrozenIncreased event raised by the Frozen contract.
type FrozenFrozenIncreased struct {
	Validator     common.Address
	Amount        *big.Int
	FrozenBalance *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFrozenIncreased is a free log retrieval operation binding the contract event 0xaf1eabfa68d345ba5007d2f3fae9c6f58e9a9d3fe16c21e651bd00f9bac4ad4d.
//
// Solidity: event FrozenIncreased(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Frozen *FrozenFilterer) FilterFrozenIncreased(opts *bind.FilterOpts, _validator []common.Address) (*FrozenFrozenIncreasedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Frozen.contract.FilterLogs(opts, "FrozenIncreased", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &FrozenFrozenIncreasedIterator{contract: _Frozen.contract, event: "FrozenIncreased", logs: logs, sub: sub}, nil
}

// WatchFrozenIncreased is a free log subscription operation binding the contract event 0xaf1eabfa68d345ba5007d2f3fae9c6f58e9a9d3fe16c21e651bd00f9bac4ad4d.
//
// Solidity: event FrozenIncreased(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Frozen *FrozenFilterer) WatchFrozenIncreased(opts *bind.WatchOpts, sink chan<- *FrozenFrozenIncreased, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Frozen.contract.WatchLogs(opts, "FrozenIncreased", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FrozenFrozenIncreased)
				if err := _Frozen.contract.UnpackLog(event, "FrozenIncreased", log); err != nil {
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

// ParseFrozenIncreased is a log parse operation binding the contract event 0xaf1eabfa68d345ba5007d2f3fae9c6f58e9a9d3fe16c21e651bd00f9bac4ad4d.
//
// Solidity: event FrozenIncreased(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Frozen *FrozenFilterer) ParseFrozenIncreased(log types.Log) (*FrozenFrozenIncreased, error) {
	event := new(FrozenFrozenIncreased)
	if err := _Frozen.contract.UnpackLog(event, "FrozenIncreased", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockABI is the input ABI used to generate the binding from.
const LockABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unlockedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"statusOf\",\"outputs\":[{\"internalType\":\"enumLock.ValidatorStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SLASHED_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_unlockedAt\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"}]"

// LockFuncSigs maps the 4-byte function signature to its string representation.
var LockFuncSigs = map[string]string{
	"e7065963": "SLASHED_BLOCK()",
	"97a5d5b5": "statusOf(address)",
	"2cf5db13": "unlockedAt(address)",
}

// LockBin is the compiled bytecode used for deploying new contracts.
var LockBin = "0x608060405234801561001057600080fd5b5061016a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632cf5db131461004657806397a5d5b51461007e578063e7065963146100c8575b600080fd5b61006c6004803603602081101561005c57600080fd5b50356001600160a01b03166100d0565b60408051918252519081900360200190f35b6100a46004803603602081101561009457600080fd5b50356001600160a01b03166100e2565b604051808260028111156100b457fe5b60ff16815260200191505060405180910390f35b61006c61012d565b60006020819052908152604090205481565b6001600160a01b0381166000908152602081905260408120544381101561010d576000915050610128565b600160ff1b8114610122576001915050610128565b60029150505b919050565b600160ff1b8156fea265627a7a7231582096623cbc6736ee54562b51e12fa328b1e3703497b6cd641675564751e96f0c5b64736f6c634300050b0032"

// DeployLock deploys a new Ethereum contract, binding an instance of Lock to it.
func DeployLock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Lock, error) {
	parsed, err := abi.JSON(strings.NewReader(LockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lock{LockCaller: LockCaller{contract: contract}, LockTransactor: LockTransactor{contract: contract}, LockFilterer: LockFilterer{contract: contract}}, nil
}

// Lock is an auto generated Go binding around an Ethereum contract.
type Lock struct {
	LockCaller     // Read-only binding to the contract
	LockTransactor // Write-only binding to the contract
	LockFilterer   // Log filterer for contract events
}

// LockCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockSession struct {
	Contract     *Lock             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockCallerSession struct {
	Contract *LockCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockTransactorSession struct {
	Contract     *LockTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockRaw struct {
	Contract *Lock // Generic contract binding to access the raw methods on
}

// LockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockCallerRaw struct {
	Contract *LockCaller // Generic read-only contract binding to access the raw methods on
}

// LockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockTransactorRaw struct {
	Contract *LockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLock creates a new instance of Lock, bound to a specific deployed contract.
func NewLock(address common.Address, backend bind.ContractBackend) (*Lock, error) {
	contract, err := bindLock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lock{LockCaller: LockCaller{contract: contract}, LockTransactor: LockTransactor{contract: contract}, LockFilterer: LockFilterer{contract: contract}}, nil
}

// NewLockCaller creates a new read-only instance of Lock, bound to a specific deployed contract.
func NewLockCaller(address common.Address, caller bind.ContractCaller) (*LockCaller, error) {
	contract, err := bindLock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockCaller{contract: contract}, nil
}

// NewLockTransactor creates a new write-only instance of Lock, bound to a specific deployed contract.
func NewLockTransactor(address common.Address, transactor bind.ContractTransactor) (*LockTransactor, error) {
	contract, err := bindLock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockTransactor{contract: contract}, nil
}

// NewLockFilterer creates a new log filterer instance of Lock, bound to a specific deployed contract.
func NewLockFilterer(address common.Address, filterer bind.ContractFilterer) (*LockFilterer, error) {
	contract, err := bindLock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockFilterer{contract: contract}, nil
}

// bindLock binds a generic wrapper to an already deployed contract.
func bindLock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lock *LockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lock.Contract.LockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lock *LockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lock.Contract.LockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lock *LockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lock.Contract.LockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lock *LockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lock *LockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lock *LockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lock.Contract.contract.Transact(opts, method, params...)
}

// SLASHEDBLOCK is a free data retrieval call binding the contract method 0xe7065963.
//
// Solidity: function SLASHED_BLOCK() constant returns(uint256)
func (_Lock *LockCaller) SLASHEDBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lock.contract.Call(opts, out, "SLASHED_BLOCK")
	return *ret0, err
}

// SLASHEDBLOCK is a free data retrieval call binding the contract method 0xe7065963.
//
// Solidity: function SLASHED_BLOCK() constant returns(uint256)
func (_Lock *LockSession) SLASHEDBLOCK() (*big.Int, error) {
	return _Lock.Contract.SLASHEDBLOCK(&_Lock.CallOpts)
}

// SLASHEDBLOCK is a free data retrieval call binding the contract method 0xe7065963.
//
// Solidity: function SLASHED_BLOCK() constant returns(uint256)
func (_Lock *LockCallerSession) SLASHEDBLOCK() (*big.Int, error) {
	return _Lock.Contract.SLASHEDBLOCK(&_Lock.CallOpts)
}

// StatusOf is a free data retrieval call binding the contract method 0x97a5d5b5.
//
// Solidity: function statusOf(address _validator) constant returns(uint8)
func (_Lock *LockCaller) StatusOf(opts *bind.CallOpts, _validator common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Lock.contract.Call(opts, out, "statusOf", _validator)
	return *ret0, err
}

// StatusOf is a free data retrieval call binding the contract method 0x97a5d5b5.
//
// Solidity: function statusOf(address _validator) constant returns(uint8)
func (_Lock *LockSession) StatusOf(_validator common.Address) (uint8, error) {
	return _Lock.Contract.StatusOf(&_Lock.CallOpts, _validator)
}

// StatusOf is a free data retrieval call binding the contract method 0x97a5d5b5.
//
// Solidity: function statusOf(address _validator) constant returns(uint8)
func (_Lock *LockCallerSession) StatusOf(_validator common.Address) (uint8, error) {
	return _Lock.Contract.StatusOf(&_Lock.CallOpts, _validator)
}

// UnlockedAt is a free data retrieval call binding the contract method 0x2cf5db13.
//
// Solidity: function unlockedAt(address ) constant returns(uint256)
func (_Lock *LockCaller) UnlockedAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lock.contract.Call(opts, out, "unlockedAt", arg0)
	return *ret0, err
}

// UnlockedAt is a free data retrieval call binding the contract method 0x2cf5db13.
//
// Solidity: function unlockedAt(address ) constant returns(uint256)
func (_Lock *LockSession) UnlockedAt(arg0 common.Address) (*big.Int, error) {
	return _Lock.Contract.UnlockedAt(&_Lock.CallOpts, arg0)
}

// UnlockedAt is a free data retrieval call binding the contract method 0x2cf5db13.
//
// Solidity: function unlockedAt(address ) constant returns(uint256)
func (_Lock *LockCallerSession) UnlockedAt(arg0 common.Address) (*big.Int, error) {
	return _Lock.Contract.UnlockedAt(&_Lock.CallOpts, arg0)
}

// LockLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the Lock contract.
type LockLockedIterator struct {
	Event *LockLocked // Event containing the contract specifics and raw log

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
func (it *LockLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockLocked)
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
		it.Event = new(LockLocked)
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
func (it *LockLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockLocked represents a Locked event raised by the Lock contract.
type LockLocked struct {
	Validator  common.Address
	UnlockedAt *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed _validator, uint256 indexed _unlockedAt)
func (_Lock *LockFilterer) FilterLocked(opts *bind.FilterOpts, _validator []common.Address, _unlockedAt []*big.Int) (*LockLockedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}
	var _unlockedAtRule []interface{}
	for _, _unlockedAtItem := range _unlockedAt {
		_unlockedAtRule = append(_unlockedAtRule, _unlockedAtItem)
	}

	logs, sub, err := _Lock.contract.FilterLogs(opts, "Locked", _validatorRule, _unlockedAtRule)
	if err != nil {
		return nil, err
	}
	return &LockLockedIterator{contract: _Lock.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed _validator, uint256 indexed _unlockedAt)
func (_Lock *LockFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *LockLocked, _validator []common.Address, _unlockedAt []*big.Int) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}
	var _unlockedAtRule []interface{}
	for _, _unlockedAtItem := range _unlockedAt {
		_unlockedAtRule = append(_unlockedAtRule, _unlockedAtItem)
	}

	logs, sub, err := _Lock.contract.WatchLogs(opts, "Locked", _validatorRule, _unlockedAtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockLocked)
				if err := _Lock.contract.UnpackLog(event, "Locked", log); err != nil {
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

// ParseLocked is a log parse operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed _validator, uint256 indexed _unlockedAt)
func (_Lock *LockFilterer) ParseLocked(log types.Log) (*LockLocked, error) {
	event := new(LockLocked)
	if err := _Lock.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the Lock contract.
type LockSlashedIterator struct {
	Event *LockSlashed // Event containing the contract specifics and raw log

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
func (it *LockSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockSlashed)
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
		it.Event = new(LockSlashed)
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
func (it *LockSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockSlashed represents a Slashed event raised by the Lock contract.
type LockSlashed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x975ad74f3e1160b20a620ec57c636ffa56c6b6679a0e2fb5689b4f95e65a6946.
//
// Solidity: event Slashed(address indexed _validator)
func (_Lock *LockFilterer) FilterSlashed(opts *bind.FilterOpts, _validator []common.Address) (*LockSlashedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Lock.contract.FilterLogs(opts, "Slashed", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &LockSlashedIterator{contract: _Lock.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x975ad74f3e1160b20a620ec57c636ffa56c6b6679a0e2fb5689b4f95e65a6946.
//
// Solidity: event Slashed(address indexed _validator)
func (_Lock *LockFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *LockSlashed, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Lock.contract.WatchLogs(opts, "Slashed", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockSlashed)
				if err := _Lock.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x975ad74f3e1160b20a620ec57c636ffa56c6b6679a0e2fb5689b4f95e65a6946.
//
// Solidity: event Slashed(address indexed _validator)
func (_Lock *LockFilterer) ParseSlashed(log types.Log) (*LockSlashed, error) {
	event := new(LockSlashed)
	if err := _Lock.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PenalizeABI is the input ABI used to generate the binding from.
const PenalizeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ROYALTY_STACK_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BLOCKS_PER_ROYALTY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"ProofSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_point\",\"type\":\"uint256\"}],\"name\":\"PointUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"StartTimer\",\"type\":\"event\"}]"

// PenalizeFuncSigs maps the 4-byte function signature to its string representation.
var PenalizeFuncSigs = map[string]string{
	"dc49f0dc": "BLOCKS_PER_ROYALTY()",
	"3494e267": "ROYALTY_STACK_LIMIT()",
}

// PenalizeBin is the compiled bytecode used for deploying new contracts.
var PenalizeBin = "0x6080604052600a6000556003600155348015601957600080fd5b506096806100286000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80633494e267146037578063dc49f0dc14604f575b600080fd5b603d6055565b60408051918252519081900360200190f35b603d605b565b60015481565b6000548156fea265627a7a7231582000ce615b777ddcf814768917960b04535ab2849928cf9e908712f14505d9c55064736f6c634300050b0032"

// DeployPenalize deploys a new Ethereum contract, binding an instance of Penalize to it.
func DeployPenalize(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Penalize, error) {
	parsed, err := abi.JSON(strings.NewReader(PenalizeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PenalizeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Penalize{PenalizeCaller: PenalizeCaller{contract: contract}, PenalizeTransactor: PenalizeTransactor{contract: contract}, PenalizeFilterer: PenalizeFilterer{contract: contract}}, nil
}

// Penalize is an auto generated Go binding around an Ethereum contract.
type Penalize struct {
	PenalizeCaller     // Read-only binding to the contract
	PenalizeTransactor // Write-only binding to the contract
	PenalizeFilterer   // Log filterer for contract events
}

// PenalizeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PenalizeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PenalizeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PenalizeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PenalizeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PenalizeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PenalizeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PenalizeSession struct {
	Contract     *Penalize         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PenalizeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PenalizeCallerSession struct {
	Contract *PenalizeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PenalizeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PenalizeTransactorSession struct {
	Contract     *PenalizeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PenalizeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PenalizeRaw struct {
	Contract *Penalize // Generic contract binding to access the raw methods on
}

// PenalizeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PenalizeCallerRaw struct {
	Contract *PenalizeCaller // Generic read-only contract binding to access the raw methods on
}

// PenalizeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PenalizeTransactorRaw struct {
	Contract *PenalizeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPenalize creates a new instance of Penalize, bound to a specific deployed contract.
func NewPenalize(address common.Address, backend bind.ContractBackend) (*Penalize, error) {
	contract, err := bindPenalize(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Penalize{PenalizeCaller: PenalizeCaller{contract: contract}, PenalizeTransactor: PenalizeTransactor{contract: contract}, PenalizeFilterer: PenalizeFilterer{contract: contract}}, nil
}

// NewPenalizeCaller creates a new read-only instance of Penalize, bound to a specific deployed contract.
func NewPenalizeCaller(address common.Address, caller bind.ContractCaller) (*PenalizeCaller, error) {
	contract, err := bindPenalize(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PenalizeCaller{contract: contract}, nil
}

// NewPenalizeTransactor creates a new write-only instance of Penalize, bound to a specific deployed contract.
func NewPenalizeTransactor(address common.Address, transactor bind.ContractTransactor) (*PenalizeTransactor, error) {
	contract, err := bindPenalize(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PenalizeTransactor{contract: contract}, nil
}

// NewPenalizeFilterer creates a new log filterer instance of Penalize, bound to a specific deployed contract.
func NewPenalizeFilterer(address common.Address, filterer bind.ContractFilterer) (*PenalizeFilterer, error) {
	contract, err := bindPenalize(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PenalizeFilterer{contract: contract}, nil
}

// bindPenalize binds a generic wrapper to an already deployed contract.
func bindPenalize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PenalizeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Penalize *PenalizeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Penalize.Contract.PenalizeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Penalize *PenalizeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Penalize.Contract.PenalizeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Penalize *PenalizeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Penalize.Contract.PenalizeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Penalize *PenalizeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Penalize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Penalize *PenalizeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Penalize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Penalize *PenalizeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Penalize.Contract.contract.Transact(opts, method, params...)
}

// BLOCKSPERROYALTY is a free data retrieval call binding the contract method 0xdc49f0dc.
//
// Solidity: function BLOCKS_PER_ROYALTY() constant returns(uint256)
func (_Penalize *PenalizeCaller) BLOCKSPERROYALTY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Penalize.contract.Call(opts, out, "BLOCKS_PER_ROYALTY")
	return *ret0, err
}

// BLOCKSPERROYALTY is a free data retrieval call binding the contract method 0xdc49f0dc.
//
// Solidity: function BLOCKS_PER_ROYALTY() constant returns(uint256)
func (_Penalize *PenalizeSession) BLOCKSPERROYALTY() (*big.Int, error) {
	return _Penalize.Contract.BLOCKSPERROYALTY(&_Penalize.CallOpts)
}

// BLOCKSPERROYALTY is a free data retrieval call binding the contract method 0xdc49f0dc.
//
// Solidity: function BLOCKS_PER_ROYALTY() constant returns(uint256)
func (_Penalize *PenalizeCallerSession) BLOCKSPERROYALTY() (*big.Int, error) {
	return _Penalize.Contract.BLOCKSPERROYALTY(&_Penalize.CallOpts)
}

// ROYALTYSTACKLIMIT is a free data retrieval call binding the contract method 0x3494e267.
//
// Solidity: function ROYALTY_STACK_LIMIT() constant returns(uint256)
func (_Penalize *PenalizeCaller) ROYALTYSTACKLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Penalize.contract.Call(opts, out, "ROYALTY_STACK_LIMIT")
	return *ret0, err
}

// ROYALTYSTACKLIMIT is a free data retrieval call binding the contract method 0x3494e267.
//
// Solidity: function ROYALTY_STACK_LIMIT() constant returns(uint256)
func (_Penalize *PenalizeSession) ROYALTYSTACKLIMIT() (*big.Int, error) {
	return _Penalize.Contract.ROYALTYSTACKLIMIT(&_Penalize.CallOpts)
}

// ROYALTYSTACKLIMIT is a free data retrieval call binding the contract method 0x3494e267.
//
// Solidity: function ROYALTY_STACK_LIMIT() constant returns(uint256)
func (_Penalize *PenalizeCallerSession) ROYALTYSTACKLIMIT() (*big.Int, error) {
	return _Penalize.Contract.ROYALTYSTACKLIMIT(&_Penalize.CallOpts)
}

// PenalizePointUpdatedIterator is returned from FilterPointUpdated and is used to iterate over the raw logs and unpacked data for PointUpdated events raised by the Penalize contract.
type PenalizePointUpdatedIterator struct {
	Event *PenalizePointUpdated // Event containing the contract specifics and raw log

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
func (it *PenalizePointUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PenalizePointUpdated)
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
		it.Event = new(PenalizePointUpdated)
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
func (it *PenalizePointUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PenalizePointUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PenalizePointUpdated represents a PointUpdated event raised by the Penalize contract.
type PenalizePointUpdated struct {
	Validator common.Address
	Point     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPointUpdated is a free log retrieval operation binding the contract event 0x88a501824285cc7436ca02dbfc2953e94f577c28429dec36bac9c7782699a9d8.
//
// Solidity: event PointUpdated(address indexed _validator, uint256 _point)
func (_Penalize *PenalizeFilterer) FilterPointUpdated(opts *bind.FilterOpts, _validator []common.Address) (*PenalizePointUpdatedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Penalize.contract.FilterLogs(opts, "PointUpdated", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &PenalizePointUpdatedIterator{contract: _Penalize.contract, event: "PointUpdated", logs: logs, sub: sub}, nil
}

// WatchPointUpdated is a free log subscription operation binding the contract event 0x88a501824285cc7436ca02dbfc2953e94f577c28429dec36bac9c7782699a9d8.
//
// Solidity: event PointUpdated(address indexed _validator, uint256 _point)
func (_Penalize *PenalizeFilterer) WatchPointUpdated(opts *bind.WatchOpts, sink chan<- *PenalizePointUpdated, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Penalize.contract.WatchLogs(opts, "PointUpdated", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PenalizePointUpdated)
				if err := _Penalize.contract.UnpackLog(event, "PointUpdated", log); err != nil {
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

// ParsePointUpdated is a log parse operation binding the contract event 0x88a501824285cc7436ca02dbfc2953e94f577c28429dec36bac9c7782699a9d8.
//
// Solidity: event PointUpdated(address indexed _validator, uint256 _point)
func (_Penalize *PenalizeFilterer) ParsePointUpdated(log types.Log) (*PenalizePointUpdated, error) {
	event := new(PenalizePointUpdated)
	if err := _Penalize.contract.UnpackLog(event, "PointUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PenalizeProofSubmittedIterator is returned from FilterProofSubmitted and is used to iterate over the raw logs and unpacked data for ProofSubmitted events raised by the Penalize contract.
type PenalizeProofSubmittedIterator struct {
	Event *PenalizeProofSubmitted // Event containing the contract specifics and raw log

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
func (it *PenalizeProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PenalizeProofSubmitted)
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
		it.Event = new(PenalizeProofSubmitted)
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
func (it *PenalizeProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PenalizeProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PenalizeProofSubmitted represents a ProofSubmitted event raised by the Penalize contract.
type PenalizeProofSubmitted struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProofSubmitted is a free log retrieval operation binding the contract event 0x87dc06121a3dc1ab46683cb20422ce8b605fd9a4369ad5128f0dc2a49283e84c.
//
// Solidity: event ProofSubmitted(address indexed _validator)
func (_Penalize *PenalizeFilterer) FilterProofSubmitted(opts *bind.FilterOpts, _validator []common.Address) (*PenalizeProofSubmittedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Penalize.contract.FilterLogs(opts, "ProofSubmitted", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &PenalizeProofSubmittedIterator{contract: _Penalize.contract, event: "ProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchProofSubmitted is a free log subscription operation binding the contract event 0x87dc06121a3dc1ab46683cb20422ce8b605fd9a4369ad5128f0dc2a49283e84c.
//
// Solidity: event ProofSubmitted(address indexed _validator)
func (_Penalize *PenalizeFilterer) WatchProofSubmitted(opts *bind.WatchOpts, sink chan<- *PenalizeProofSubmitted, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Penalize.contract.WatchLogs(opts, "ProofSubmitted", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PenalizeProofSubmitted)
				if err := _Penalize.contract.UnpackLog(event, "ProofSubmitted", log); err != nil {
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

// ParseProofSubmitted is a log parse operation binding the contract event 0x87dc06121a3dc1ab46683cb20422ce8b605fd9a4369ad5128f0dc2a49283e84c.
//
// Solidity: event ProofSubmitted(address indexed _validator)
func (_Penalize *PenalizeFilterer) ParseProofSubmitted(log types.Log) (*PenalizeProofSubmitted, error) {
	event := new(PenalizeProofSubmitted)
	if err := _Penalize.contract.UnpackLog(event, "ProofSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PenalizeStartTimerIterator is returned from FilterStartTimer and is used to iterate over the raw logs and unpacked data for StartTimer events raised by the Penalize contract.
type PenalizeStartTimerIterator struct {
	Event *PenalizeStartTimer // Event containing the contract specifics and raw log

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
func (it *PenalizeStartTimerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PenalizeStartTimer)
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
		it.Event = new(PenalizeStartTimer)
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
func (it *PenalizeStartTimerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PenalizeStartTimerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PenalizeStartTimer represents a StartTimer event raised by the Penalize contract.
type PenalizeStartTimer struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStartTimer is a free log retrieval operation binding the contract event 0x3d364b4503e0f292c41ea2af3eea2464e1e89d01f1a7d6519063a7342ab86522.
//
// Solidity: event StartTimer(address indexed _validator)
func (_Penalize *PenalizeFilterer) FilterStartTimer(opts *bind.FilterOpts, _validator []common.Address) (*PenalizeStartTimerIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Penalize.contract.FilterLogs(opts, "StartTimer", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &PenalizeStartTimerIterator{contract: _Penalize.contract, event: "StartTimer", logs: logs, sub: sub}, nil
}

// WatchStartTimer is a free log subscription operation binding the contract event 0x3d364b4503e0f292c41ea2af3eea2464e1e89d01f1a7d6519063a7342ab86522.
//
// Solidity: event StartTimer(address indexed _validator)
func (_Penalize *PenalizeFilterer) WatchStartTimer(opts *bind.WatchOpts, sink chan<- *PenalizeStartTimer, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Penalize.contract.WatchLogs(opts, "StartTimer", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PenalizeStartTimer)
				if err := _Penalize.contract.UnpackLog(event, "StartTimer", log); err != nil {
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

// ParseStartTimer is a log parse operation binding the contract event 0x3d364b4503e0f292c41ea2af3eea2464e1e89d01f1a7d6519063a7342ab86522.
//
// Solidity: event StartTimer(address indexed _validator)
func (_Penalize *PenalizeFilterer) ParseStartTimer(log types.Log) (*PenalizeStartTimer, error) {
	event := new(PenalizeStartTimer)
	if err := _Penalize.contract.UnpackLog(event, "StartTimer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PremineABI is the input ABI used to generate the binding from.
const PremineABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"premineAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPremineSum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"premineBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPremineAddressesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getPremineAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PremineFuncSigs maps the 4-byte function signature to its string representation.
var PremineFuncSigs = map[string]string{
	"d091afdc": "getPremineAccount(uint256)",
	"93337f3a": "getPremineAddressesCount()",
	"28b8f386": "getPremineSum()",
	"1afa8d29": "premineAddresses(uint256)",
	"86d7b2bb": "premineBalances(uint256)",
}

// PremineBin is the compiled bytecode used for deploying new contracts.
var PremineBin = "0x608060405234801561001057600080fd5b506101e4806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631afa8d291461005c57806328b8f3861461009557806386d7b2bb146100af57806393337f3a146100cc578063d091afdc146100d4575b600080fd5b6100796004803603602081101561007257600080fd5b5035610114565b604080516001600160a01b039092168252519081900360200190f35b61009d61013b565b60408051918252519081900360200190f35b61009d600480360360208110156100c557600080fd5b5035610141565b61009d61015f565b6100f1600480360360208110156100ea57600080fd5b5035610165565b604080516001600160a01b03909316835260208301919091528051918290030190f35b6001818154811061012157fe5b6000918252602090912001546001600160a01b0316905081565b60005490565b6002818154811061014e57fe5b600091825260209091200154905081565b60015490565b6000806001838154811061017557fe5b600091825260209091200154600280546001600160a01b03909216918590811061019b57fe5b90600052602060002001549150915091509156fea265627a7a72315820e4842c67cbcfd589f4b92fa68b0c5d6ccfcf1cf275ddf3ecdc8307d8a7af1aa764736f6c634300050b0032"

// DeployPremine deploys a new Ethereum contract, binding an instance of Premine to it.
func DeployPremine(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Premine, error) {
	parsed, err := abi.JSON(strings.NewReader(PremineABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PremineBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Premine{PremineCaller: PremineCaller{contract: contract}, PremineTransactor: PremineTransactor{contract: contract}, PremineFilterer: PremineFilterer{contract: contract}}, nil
}

// Premine is an auto generated Go binding around an Ethereum contract.
type Premine struct {
	PremineCaller     // Read-only binding to the contract
	PremineTransactor // Write-only binding to the contract
	PremineFilterer   // Log filterer for contract events
}

// PremineCaller is an auto generated read-only Go binding around an Ethereum contract.
type PremineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PremineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PremineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PremineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PremineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PremineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PremineSession struct {
	Contract     *Premine          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PremineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PremineCallerSession struct {
	Contract *PremineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PremineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PremineTransactorSession struct {
	Contract     *PremineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PremineRaw is an auto generated low-level Go binding around an Ethereum contract.
type PremineRaw struct {
	Contract *Premine // Generic contract binding to access the raw methods on
}

// PremineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PremineCallerRaw struct {
	Contract *PremineCaller // Generic read-only contract binding to access the raw methods on
}

// PremineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PremineTransactorRaw struct {
	Contract *PremineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPremine creates a new instance of Premine, bound to a specific deployed contract.
func NewPremine(address common.Address, backend bind.ContractBackend) (*Premine, error) {
	contract, err := bindPremine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Premine{PremineCaller: PremineCaller{contract: contract}, PremineTransactor: PremineTransactor{contract: contract}, PremineFilterer: PremineFilterer{contract: contract}}, nil
}

// NewPremineCaller creates a new read-only instance of Premine, bound to a specific deployed contract.
func NewPremineCaller(address common.Address, caller bind.ContractCaller) (*PremineCaller, error) {
	contract, err := bindPremine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PremineCaller{contract: contract}, nil
}

// NewPremineTransactor creates a new write-only instance of Premine, bound to a specific deployed contract.
func NewPremineTransactor(address common.Address, transactor bind.ContractTransactor) (*PremineTransactor, error) {
	contract, err := bindPremine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PremineTransactor{contract: contract}, nil
}

// NewPremineFilterer creates a new log filterer instance of Premine, bound to a specific deployed contract.
func NewPremineFilterer(address common.Address, filterer bind.ContractFilterer) (*PremineFilterer, error) {
	contract, err := bindPremine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PremineFilterer{contract: contract}, nil
}

// bindPremine binds a generic wrapper to an already deployed contract.
func bindPremine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PremineABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Premine *PremineRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Premine.Contract.PremineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Premine *PremineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Premine.Contract.PremineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Premine *PremineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Premine.Contract.PremineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Premine *PremineCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Premine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Premine *PremineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Premine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Premine *PremineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Premine.Contract.contract.Transact(opts, method, params...)
}

// GetPremineAccount is a free data retrieval call binding the contract method 0xd091afdc.
//
// Solidity: function getPremineAccount(uint256 i) constant returns(address, uint256)
func (_Premine *PremineCaller) GetPremineAccount(opts *bind.CallOpts, i *big.Int) (common.Address, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Premine.contract.Call(opts, out, "getPremineAccount", i)
	return *ret0, *ret1, err
}

// GetPremineAccount is a free data retrieval call binding the contract method 0xd091afdc.
//
// Solidity: function getPremineAccount(uint256 i) constant returns(address, uint256)
func (_Premine *PremineSession) GetPremineAccount(i *big.Int) (common.Address, *big.Int, error) {
	return _Premine.Contract.GetPremineAccount(&_Premine.CallOpts, i)
}

// GetPremineAccount is a free data retrieval call binding the contract method 0xd091afdc.
//
// Solidity: function getPremineAccount(uint256 i) constant returns(address, uint256)
func (_Premine *PremineCallerSession) GetPremineAccount(i *big.Int) (common.Address, *big.Int, error) {
	return _Premine.Contract.GetPremineAccount(&_Premine.CallOpts, i)
}

// GetPremineAddressesCount is a free data retrieval call binding the contract method 0x93337f3a.
//
// Solidity: function getPremineAddressesCount() constant returns(uint256)
func (_Premine *PremineCaller) GetPremineAddressesCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Premine.contract.Call(opts, out, "getPremineAddressesCount")
	return *ret0, err
}

// GetPremineAddressesCount is a free data retrieval call binding the contract method 0x93337f3a.
//
// Solidity: function getPremineAddressesCount() constant returns(uint256)
func (_Premine *PremineSession) GetPremineAddressesCount() (*big.Int, error) {
	return _Premine.Contract.GetPremineAddressesCount(&_Premine.CallOpts)
}

// GetPremineAddressesCount is a free data retrieval call binding the contract method 0x93337f3a.
//
// Solidity: function getPremineAddressesCount() constant returns(uint256)
func (_Premine *PremineCallerSession) GetPremineAddressesCount() (*big.Int, error) {
	return _Premine.Contract.GetPremineAddressesCount(&_Premine.CallOpts)
}

// GetPremineSum is a free data retrieval call binding the contract method 0x28b8f386.
//
// Solidity: function getPremineSum() constant returns(uint256)
func (_Premine *PremineCaller) GetPremineSum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Premine.contract.Call(opts, out, "getPremineSum")
	return *ret0, err
}

// GetPremineSum is a free data retrieval call binding the contract method 0x28b8f386.
//
// Solidity: function getPremineSum() constant returns(uint256)
func (_Premine *PremineSession) GetPremineSum() (*big.Int, error) {
	return _Premine.Contract.GetPremineSum(&_Premine.CallOpts)
}

// GetPremineSum is a free data retrieval call binding the contract method 0x28b8f386.
//
// Solidity: function getPremineSum() constant returns(uint256)
func (_Premine *PremineCallerSession) GetPremineSum() (*big.Int, error) {
	return _Premine.Contract.GetPremineSum(&_Premine.CallOpts)
}

// PremineAddresses is a free data retrieval call binding the contract method 0x1afa8d29.
//
// Solidity: function premineAddresses(uint256 ) constant returns(address)
func (_Premine *PremineCaller) PremineAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Premine.contract.Call(opts, out, "premineAddresses", arg0)
	return *ret0, err
}

// PremineAddresses is a free data retrieval call binding the contract method 0x1afa8d29.
//
// Solidity: function premineAddresses(uint256 ) constant returns(address)
func (_Premine *PremineSession) PremineAddresses(arg0 *big.Int) (common.Address, error) {
	return _Premine.Contract.PremineAddresses(&_Premine.CallOpts, arg0)
}

// PremineAddresses is a free data retrieval call binding the contract method 0x1afa8d29.
//
// Solidity: function premineAddresses(uint256 ) constant returns(address)
func (_Premine *PremineCallerSession) PremineAddresses(arg0 *big.Int) (common.Address, error) {
	return _Premine.Contract.PremineAddresses(&_Premine.CallOpts, arg0)
}

// PremineBalances is a free data retrieval call binding the contract method 0x86d7b2bb.
//
// Solidity: function premineBalances(uint256 ) constant returns(uint256)
func (_Premine *PremineCaller) PremineBalances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Premine.contract.Call(opts, out, "premineBalances", arg0)
	return *ret0, err
}

// PremineBalances is a free data retrieval call binding the contract method 0x86d7b2bb.
//
// Solidity: function premineBalances(uint256 ) constant returns(uint256)
func (_Premine *PremineSession) PremineBalances(arg0 *big.Int) (*big.Int, error) {
	return _Premine.Contract.PremineBalances(&_Premine.CallOpts, arg0)
}

// PremineBalances is a free data retrieval call binding the contract method 0x86d7b2bb.
//
// Solidity: function premineBalances(uint256 ) constant returns(uint256)
func (_Premine *PremineCallerSession) PremineBalances(arg0 *big.Int) (*big.Int, error) {
	return _Premine.Contract.PremineBalances(&_Premine.CallOpts, arg0)
}

// RVSTokenABI is the input ABI used to generate the binding from.
const RVSTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// RVSTokenFuncSigs maps the 4-byte function signature to its string representation.
var RVSTokenFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
}

// RVSTokenBin is the compiled bytecode used for deploying new contracts.
var RVSTokenBin = "0x608060405234801561001057600080fd5b506040516105203803806105208339818101604052606081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825164010000000081118282018810171561008257600080fd5b82525081516020918201929091019080838360005b838110156100af578181015183820152602001610097565b50505050905090810190601f1680156100dc5780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156100ff57600080fd5b90830190602082018581111561011457600080fd5b825164010000000081118282018810171561012e57600080fd5b82525081516020918201929091019080838360005b8381101561015b578181015183820152602001610143565b50505050905090810190601f1680156101885780820380516001836020036101000a031916815260200191505b506040526020908101518551909350859250849184916101ae91600091908601906101e2565b5081516101c29060019060208501906101e2565b506002805460ff191660ff929092169190911790555061027d9350505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061022357805160ff1916838001178555610250565b82800160010185558215610250579182015b82811115610250578251825591602001919060010190610235565b5061025c929150610260565b5090565b61027a91905b8082111561025c5760008155600101610266565b90565b6102948061028c6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806306fdde031461005c57806318160ddd146100d9578063313ce567146100f357806370a082311461011157806395d89b4114610137575b600080fd5b61006461013f565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561009e578181015183820152602001610086565b50505050905090810190601f1680156100cb5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100e16101d5565b60408051918252519081900360200190f35b6100fb6101db565b6040805160ff9092168252519081900360200190f35b6100e16004803603602081101561012757600080fd5b50356001600160a01b03166101e4565b6100646101ff565b60008054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156101cb5780601f106101a0576101008083540402835291602001916101cb565b820191906000526020600020905b8154815290600101906020018083116101ae57829003601f168201915b5050505050905090565b60045490565b60025460ff1690565b6001600160a01b031660009081526003602052604090205490565b60018054604080516020601f600260001961010087891615020190951694909404938401819004810282018101909252828152606093909290918301828280156101cb5780601f106101a0576101008083540402835291602001916101cb56fea265627a7a723158209d427ff5cae52940dab7360bdfaf73e5827f9deefeb8a814b97158662d69630f64736f6c634300050b0032"

// DeployRVSToken deploys a new Ethereum contract, binding an instance of RVSToken to it.
func DeployRVSToken(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals uint8) (common.Address, *types.Transaction, *RVSToken, error) {
	parsed, err := abi.JSON(strings.NewReader(RVSTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RVSTokenBin), backend, name, symbol, decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RVSToken{RVSTokenCaller: RVSTokenCaller{contract: contract}, RVSTokenTransactor: RVSTokenTransactor{contract: contract}, RVSTokenFilterer: RVSTokenFilterer{contract: contract}}, nil
}

// RVSToken is an auto generated Go binding around an Ethereum contract.
type RVSToken struct {
	RVSTokenCaller     // Read-only binding to the contract
	RVSTokenTransactor // Write-only binding to the contract
	RVSTokenFilterer   // Log filterer for contract events
}

// RVSTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type RVSTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RVSTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RVSTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RVSTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RVSTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RVSTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RVSTokenSession struct {
	Contract     *RVSToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RVSTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RVSTokenCallerSession struct {
	Contract *RVSTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RVSTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RVSTokenTransactorSession struct {
	Contract     *RVSTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RVSTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type RVSTokenRaw struct {
	Contract *RVSToken // Generic contract binding to access the raw methods on
}

// RVSTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RVSTokenCallerRaw struct {
	Contract *RVSTokenCaller // Generic read-only contract binding to access the raw methods on
}

// RVSTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RVSTokenTransactorRaw struct {
	Contract *RVSTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRVSToken creates a new instance of RVSToken, bound to a specific deployed contract.
func NewRVSToken(address common.Address, backend bind.ContractBackend) (*RVSToken, error) {
	contract, err := bindRVSToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RVSToken{RVSTokenCaller: RVSTokenCaller{contract: contract}, RVSTokenTransactor: RVSTokenTransactor{contract: contract}, RVSTokenFilterer: RVSTokenFilterer{contract: contract}}, nil
}

// NewRVSTokenCaller creates a new read-only instance of RVSToken, bound to a specific deployed contract.
func NewRVSTokenCaller(address common.Address, caller bind.ContractCaller) (*RVSTokenCaller, error) {
	contract, err := bindRVSToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RVSTokenCaller{contract: contract}, nil
}

// NewRVSTokenTransactor creates a new write-only instance of RVSToken, bound to a specific deployed contract.
func NewRVSTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*RVSTokenTransactor, error) {
	contract, err := bindRVSToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RVSTokenTransactor{contract: contract}, nil
}

// NewRVSTokenFilterer creates a new log filterer instance of RVSToken, bound to a specific deployed contract.
func NewRVSTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*RVSTokenFilterer, error) {
	contract, err := bindRVSToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RVSTokenFilterer{contract: contract}, nil
}

// bindRVSToken binds a generic wrapper to an already deployed contract.
func bindRVSToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RVSTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RVSToken *RVSTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RVSToken.Contract.RVSTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RVSToken *RVSTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RVSToken.Contract.RVSTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RVSToken *RVSTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RVSToken.Contract.RVSTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RVSToken *RVSTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RVSToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RVSToken *RVSTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RVSToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RVSToken *RVSTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RVSToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_RVSToken *RVSTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RVSToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_RVSToken *RVSTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _RVSToken.Contract.BalanceOf(&_RVSToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_RVSToken *RVSTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _RVSToken.Contract.BalanceOf(&_RVSToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RVSToken *RVSTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _RVSToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RVSToken *RVSTokenSession) Decimals() (uint8, error) {
	return _RVSToken.Contract.Decimals(&_RVSToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RVSToken *RVSTokenCallerSession) Decimals() (uint8, error) {
	return _RVSToken.Contract.Decimals(&_RVSToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RVSToken *RVSTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RVSToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RVSToken *RVSTokenSession) Name() (string, error) {
	return _RVSToken.Contract.Name(&_RVSToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RVSToken *RVSTokenCallerSession) Name() (string, error) {
	return _RVSToken.Contract.Name(&_RVSToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RVSToken *RVSTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RVSToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RVSToken *RVSTokenSession) Symbol() (string, error) {
	return _RVSToken.Contract.Symbol(&_RVSToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RVSToken *RVSTokenCallerSession) Symbol() (string, error) {
	return _RVSToken.Contract.Symbol(&_RVSToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RVSToken *RVSTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RVSToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RVSToken *RVSTokenSession) TotalSupply() (*big.Int, error) {
	return _RVSToken.Contract.TotalSupply(&_RVSToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RVSToken *RVSTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _RVSToken.Contract.TotalSupply(&_RVSToken.CallOpts)
}

// RVSTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RVSToken contract.
type RVSTokenTransferIterator struct {
	Event *RVSTokenTransfer // Event containing the contract specifics and raw log

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
func (it *RVSTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RVSTokenTransfer)
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
		it.Event = new(RVSTokenTransfer)
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
func (it *RVSTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RVSTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RVSTokenTransfer represents a Transfer event raised by the RVSToken contract.
type RVSTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RVSToken *RVSTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RVSTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RVSToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RVSTokenTransferIterator{contract: _RVSToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RVSToken *RVSTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RVSTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RVSToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RVSTokenTransfer)
				if err := _RVSToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RVSToken *RVSTokenFilterer) ParseTransfer(log types.Log) (*RVSTokenTransfer, error) {
	event := new(RVSTokenTransfer)
	if err := _RVSToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820235c3db3c3a048d021b92394494f17730fef79e9e8b8624db51d8a6c132af6c364736f6c634300050b0032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StakingABI is the input ABI used to generate the binding from.
const StakingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"rewardOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCpt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CPT_ZOOM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_cpt\",\"type\":\"uint256\"}],\"name\":\"CptUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_credit\",\"type\":\"uint256\"}],\"name\":\"CreditUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_lastBalance\",\"type\":\"uint256\"}],\"name\":\"LastBalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// StakingFuncSigs maps the 4-byte function signature to its string representation.
var StakingFuncSigs = map[string]string{
	"cadba0d1": "CPT_ZOOM()",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"3dda0347": "getCpt()",
	"06fdde03": "name()",
	"1d62ebd9": "rewardOf(address)",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
}

// StakingBin is the compiled bytecode used for deploying new contracts.
var StakingBin = "0x608060405234801561001057600080fd5b506040516106dc3803806106dc8339818101604052606081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825164010000000081118282018810171561008257600080fd5b82525081516020918201929091019080838360005b838110156100af578181015183820152602001610097565b50505050905090810190601f1680156100dc5780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156100ff57600080fd5b90830190602082018581111561011457600080fd5b825164010000000081118282018810171561012e57600080fd5b82525081516020918201929091019080838360005b8381101561015b578181015183820152602001610143565b50505050905090810190601f1680156101885780820380516001836020036101000a031916815260200191505b506040526020908101518551909350859250849184918491849184916101b3916000918601906101ea565b5081516101c79060019060208501906101ea565b506002805460ff191660ff92909216919091179055506102859650505050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061022b57805160ff1916838001178555610258565b82800160010185558215610258579182015b8281111561025857825182559160200191906001019061023d565b50610264929150610268565b5090565b61028291905b80821115610264576000815560010161026e565b90565b610448806102946000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80633dda03471161005b5780633dda03471461016857806370a082311461017057806395d89b4114610196578063cadba0d11461019e57610088565b806306fdde031461008d57806318160ddd1461010a5780631d62ebd914610124578063313ce5671461014a575b600080fd5b6100956101a6565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100cf5781810151838201526020016100b7565b50505050905090810190601f1680156100fc5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61011261023c565b60408051918252519081900360200190f35b6101126004803603602081101561013a57600080fd5b50356001600160a01b0316610242565b6101526102c5565b6040805160ff9092168252519081900360200190f35b6101126102ce565b6101126004803603602081101561018657600080fd5b50356001600160a01b031661030b565b610095610326565b610112610386565b60008054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156102325780601f1061020757610100808354040283529160200191610232565b820191906000526020600020905b81548152906001019060200180831161021557829003601f168201915b5050505050905090565b60045490565b6000806501000000000061026c6102576102ce565b6102608661030b565b9063ffffffff61039016565b8161027357fe5b6001600160a01b03851660009081526005602052604090205491900491508110156102a25760009150506102c0565b6001600160a01b038316600090815260056020526040902054900390505b919050565b60025460ff1690565b600754600090303103816102e061023c565b6102f6836501000000000063ffffffff61039016565b816102fd57fe5b600654919004019250505090565b6001600160a01b031660009081526003602052604090205490565b60018054604080516020601f600260001961010087891615020190951694909404938401819004810282018101909252828152606093909290918301828280156102325780601f1061020757610100808354040283529160200191610232565b6501000000000081565b60008261039f575060006103ec565b828202828482816103ac57fe5b04146103e95760405162461bcd60e51b81526004018080602001828103825260218152602001806103f36021913960400191505060405180910390fd5b90505b9291505056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a265627a7a72315820ff587da7f4cddd7c2f5eaa88773057cded798ddba60e4045b90d42ef41a6342d64736f6c634300050b0032"

// DeployStaking deploys a new Ethereum contract, binding an instance of Staking to it.
func DeployStaking(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals uint8) (common.Address, *types.Transaction, *Staking, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StakingBin), backend, name, symbol, decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// CPTZOOM is a free data retrieval call binding the contract method 0xcadba0d1.
//
// Solidity: function CPT_ZOOM() constant returns(uint256)
func (_Staking *StakingCaller) CPTZOOM(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "CPT_ZOOM")
	return *ret0, err
}

// CPTZOOM is a free data retrieval call binding the contract method 0xcadba0d1.
//
// Solidity: function CPT_ZOOM() constant returns(uint256)
func (_Staking *StakingSession) CPTZOOM() (*big.Int, error) {
	return _Staking.Contract.CPTZOOM(&_Staking.CallOpts)
}

// CPTZOOM is a free data retrieval call binding the contract method 0xcadba0d1.
//
// Solidity: function CPT_ZOOM() constant returns(uint256)
func (_Staking *StakingCallerSession) CPTZOOM() (*big.Int, error) {
	return _Staking.Contract.CPTZOOM(&_Staking.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Staking *StakingCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Staking *StakingSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Staking.Contract.BalanceOf(&_Staking.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Staking *StakingCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Staking.Contract.BalanceOf(&_Staking.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Staking *StakingCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Staking *StakingSession) Decimals() (uint8, error) {
	return _Staking.Contract.Decimals(&_Staking.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Staking *StakingCallerSession) Decimals() (uint8, error) {
	return _Staking.Contract.Decimals(&_Staking.CallOpts)
}

// GetCpt is a free data retrieval call binding the contract method 0x3dda0347.
//
// Solidity: function getCpt() constant returns(uint256)
func (_Staking *StakingCaller) GetCpt(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "getCpt")
	return *ret0, err
}

// GetCpt is a free data retrieval call binding the contract method 0x3dda0347.
//
// Solidity: function getCpt() constant returns(uint256)
func (_Staking *StakingSession) GetCpt() (*big.Int, error) {
	return _Staking.Contract.GetCpt(&_Staking.CallOpts)
}

// GetCpt is a free data retrieval call binding the contract method 0x3dda0347.
//
// Solidity: function getCpt() constant returns(uint256)
func (_Staking *StakingCallerSession) GetCpt() (*big.Int, error) {
	return _Staking.Contract.GetCpt(&_Staking.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Staking *StakingCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Staking *StakingSession) Name() (string, error) {
	return _Staking.Contract.Name(&_Staking.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Staking *StakingCallerSession) Name() (string, error) {
	return _Staking.Contract.Name(&_Staking.CallOpts)
}

// RewardOf is a free data retrieval call binding the contract method 0x1d62ebd9.
//
// Solidity: function rewardOf(address _validator) constant returns(uint256)
func (_Staking *StakingCaller) RewardOf(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "rewardOf", _validator)
	return *ret0, err
}

// RewardOf is a free data retrieval call binding the contract method 0x1d62ebd9.
//
// Solidity: function rewardOf(address _validator) constant returns(uint256)
func (_Staking *StakingSession) RewardOf(_validator common.Address) (*big.Int, error) {
	return _Staking.Contract.RewardOf(&_Staking.CallOpts, _validator)
}

// RewardOf is a free data retrieval call binding the contract method 0x1d62ebd9.
//
// Solidity: function rewardOf(address _validator) constant returns(uint256)
func (_Staking *StakingCallerSession) RewardOf(_validator common.Address) (*big.Int, error) {
	return _Staking.Contract.RewardOf(&_Staking.CallOpts, _validator)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Staking *StakingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Staking *StakingSession) Symbol() (string, error) {
	return _Staking.Contract.Symbol(&_Staking.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Staking *StakingCallerSession) Symbol() (string, error) {
	return _Staking.Contract.Symbol(&_Staking.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Staking *StakingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Staking.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Staking *StakingSession) TotalSupply() (*big.Int, error) {
	return _Staking.Contract.TotalSupply(&_Staking.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Staking *StakingCallerSession) TotalSupply() (*big.Int, error) {
	return _Staking.Contract.TotalSupply(&_Staking.CallOpts)
}

// StakingClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Staking contract.
type StakingClaimedIterator struct {
	Event *StakingClaimed // Event containing the contract specifics and raw log

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
func (it *StakingClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingClaimed)
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
		it.Event = new(StakingClaimed)
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
func (it *StakingClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingClaimed represents a Claimed event raised by the Staking contract.
type StakingClaimed struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed _validator, uint256 _amount)
func (_Staking *StakingFilterer) FilterClaimed(opts *bind.FilterOpts, _validator []common.Address) (*StakingClaimedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Claimed", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingClaimedIterator{contract: _Staking.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed _validator, uint256 _amount)
func (_Staking *StakingFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *StakingClaimed, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Claimed", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingClaimed)
				if err := _Staking.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed _validator, uint256 _amount)
func (_Staking *StakingFilterer) ParseClaimed(log types.Log) (*StakingClaimed, error) {
	event := new(StakingClaimed)
	if err := _Staking.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingCptUpdatedIterator is returned from FilterCptUpdated and is used to iterate over the raw logs and unpacked data for CptUpdated events raised by the Staking contract.
type StakingCptUpdatedIterator struct {
	Event *StakingCptUpdated // Event containing the contract specifics and raw log

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
func (it *StakingCptUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCptUpdated)
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
		it.Event = new(StakingCptUpdated)
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
func (it *StakingCptUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCptUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCptUpdated represents a CptUpdated event raised by the Staking contract.
type StakingCptUpdated struct {
	Cpt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCptUpdated is a free log retrieval operation binding the contract event 0x48f8bdfe8a3b1a47754e147e370713fa5cee7be4f2971581513349a1964cf7a1.
//
// Solidity: event CptUpdated(uint256 indexed _cpt)
func (_Staking *StakingFilterer) FilterCptUpdated(opts *bind.FilterOpts, _cpt []*big.Int) (*StakingCptUpdatedIterator, error) {

	var _cptRule []interface{}
	for _, _cptItem := range _cpt {
		_cptRule = append(_cptRule, _cptItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "CptUpdated", _cptRule)
	if err != nil {
		return nil, err
	}
	return &StakingCptUpdatedIterator{contract: _Staking.contract, event: "CptUpdated", logs: logs, sub: sub}, nil
}

// WatchCptUpdated is a free log subscription operation binding the contract event 0x48f8bdfe8a3b1a47754e147e370713fa5cee7be4f2971581513349a1964cf7a1.
//
// Solidity: event CptUpdated(uint256 indexed _cpt)
func (_Staking *StakingFilterer) WatchCptUpdated(opts *bind.WatchOpts, sink chan<- *StakingCptUpdated, _cpt []*big.Int) (event.Subscription, error) {

	var _cptRule []interface{}
	for _, _cptItem := range _cpt {
		_cptRule = append(_cptRule, _cptItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "CptUpdated", _cptRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCptUpdated)
				if err := _Staking.contract.UnpackLog(event, "CptUpdated", log); err != nil {
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

// ParseCptUpdated is a log parse operation binding the contract event 0x48f8bdfe8a3b1a47754e147e370713fa5cee7be4f2971581513349a1964cf7a1.
//
// Solidity: event CptUpdated(uint256 indexed _cpt)
func (_Staking *StakingFilterer) ParseCptUpdated(log types.Log) (*StakingCptUpdated, error) {
	event := new(StakingCptUpdated)
	if err := _Staking.contract.UnpackLog(event, "CptUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingCreditUpdatedIterator is returned from FilterCreditUpdated and is used to iterate over the raw logs and unpacked data for CreditUpdated events raised by the Staking contract.
type StakingCreditUpdatedIterator struct {
	Event *StakingCreditUpdated // Event containing the contract specifics and raw log

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
func (it *StakingCreditUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCreditUpdated)
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
		it.Event = new(StakingCreditUpdated)
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
func (it *StakingCreditUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCreditUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCreditUpdated represents a CreditUpdated event raised by the Staking contract.
type StakingCreditUpdated struct {
	Validator common.Address
	Credit    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreditUpdated is a free log retrieval operation binding the contract event 0xb8b17c68dbe6cd120dd338473bf86b84c67e766feac2b583d29be17f3259599a.
//
// Solidity: event CreditUpdated(address indexed _validator, uint256 indexed _credit)
func (_Staking *StakingFilterer) FilterCreditUpdated(opts *bind.FilterOpts, _validator []common.Address, _credit []*big.Int) (*StakingCreditUpdatedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}
	var _creditRule []interface{}
	for _, _creditItem := range _credit {
		_creditRule = append(_creditRule, _creditItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "CreditUpdated", _validatorRule, _creditRule)
	if err != nil {
		return nil, err
	}
	return &StakingCreditUpdatedIterator{contract: _Staking.contract, event: "CreditUpdated", logs: logs, sub: sub}, nil
}

// WatchCreditUpdated is a free log subscription operation binding the contract event 0xb8b17c68dbe6cd120dd338473bf86b84c67e766feac2b583d29be17f3259599a.
//
// Solidity: event CreditUpdated(address indexed _validator, uint256 indexed _credit)
func (_Staking *StakingFilterer) WatchCreditUpdated(opts *bind.WatchOpts, sink chan<- *StakingCreditUpdated, _validator []common.Address, _credit []*big.Int) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}
	var _creditRule []interface{}
	for _, _creditItem := range _credit {
		_creditRule = append(_creditRule, _creditItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "CreditUpdated", _validatorRule, _creditRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCreditUpdated)
				if err := _Staking.contract.UnpackLog(event, "CreditUpdated", log); err != nil {
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

// ParseCreditUpdated is a log parse operation binding the contract event 0xb8b17c68dbe6cd120dd338473bf86b84c67e766feac2b583d29be17f3259599a.
//
// Solidity: event CreditUpdated(address indexed _validator, uint256 indexed _credit)
func (_Staking *StakingFilterer) ParseCreditUpdated(log types.Log) (*StakingCreditUpdated, error) {
	event := new(StakingCreditUpdated)
	if err := _Staking.contract.UnpackLog(event, "CreditUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingLastBalanceUpdatedIterator is returned from FilterLastBalanceUpdated and is used to iterate over the raw logs and unpacked data for LastBalanceUpdated events raised by the Staking contract.
type StakingLastBalanceUpdatedIterator struct {
	Event *StakingLastBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *StakingLastBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingLastBalanceUpdated)
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
		it.Event = new(StakingLastBalanceUpdated)
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
func (it *StakingLastBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingLastBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingLastBalanceUpdated represents a LastBalanceUpdated event raised by the Staking contract.
type StakingLastBalanceUpdated struct {
	LastBalance *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLastBalanceUpdated is a free log retrieval operation binding the contract event 0x555b5773a6e2f40d63f33ea8ac35afd7cabd75d29c852e7362faa31676fed7a2.
//
// Solidity: event LastBalanceUpdated(uint256 indexed _lastBalance)
func (_Staking *StakingFilterer) FilterLastBalanceUpdated(opts *bind.FilterOpts, _lastBalance []*big.Int) (*StakingLastBalanceUpdatedIterator, error) {

	var _lastBalanceRule []interface{}
	for _, _lastBalanceItem := range _lastBalance {
		_lastBalanceRule = append(_lastBalanceRule, _lastBalanceItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "LastBalanceUpdated", _lastBalanceRule)
	if err != nil {
		return nil, err
	}
	return &StakingLastBalanceUpdatedIterator{contract: _Staking.contract, event: "LastBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchLastBalanceUpdated is a free log subscription operation binding the contract event 0x555b5773a6e2f40d63f33ea8ac35afd7cabd75d29c852e7362faa31676fed7a2.
//
// Solidity: event LastBalanceUpdated(uint256 indexed _lastBalance)
func (_Staking *StakingFilterer) WatchLastBalanceUpdated(opts *bind.WatchOpts, sink chan<- *StakingLastBalanceUpdated, _lastBalance []*big.Int) (event.Subscription, error) {

	var _lastBalanceRule []interface{}
	for _, _lastBalanceItem := range _lastBalance {
		_lastBalanceRule = append(_lastBalanceRule, _lastBalanceItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "LastBalanceUpdated", _lastBalanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingLastBalanceUpdated)
				if err := _Staking.contract.UnpackLog(event, "LastBalanceUpdated", log); err != nil {
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

// ParseLastBalanceUpdated is a log parse operation binding the contract event 0x555b5773a6e2f40d63f33ea8ac35afd7cabd75d29c852e7362faa31676fed7a2.
//
// Solidity: event LastBalanceUpdated(uint256 indexed _lastBalance)
func (_Staking *StakingFilterer) ParseLastBalanceUpdated(log types.Log) (*StakingLastBalanceUpdated, error) {
	event := new(StakingLastBalanceUpdated)
	if err := _Staking.contract.UnpackLog(event, "LastBalanceUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Staking contract.
type StakingStakedIterator struct {
	Event *StakingStaked // Event containing the contract specifics and raw log

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
func (it *StakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStaked)
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
		it.Event = new(StakingStaked)
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
func (it *StakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStaked represents a Staked event raised by the Staking contract.
type StakingStaked struct {
	Validator common.Address
	Amount    *big.Int
	Stake     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed _validator, uint256 _amount, uint256 _stake)
func (_Staking *StakingFilterer) FilterStaked(opts *bind.FilterOpts, _validator []common.Address) (*StakingStakedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Staked", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakedIterator{contract: _Staking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed _validator, uint256 _amount, uint256 _stake)
func (_Staking *StakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingStaked, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Staked", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStaked)
				if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed _validator, uint256 _amount, uint256 _stake)
func (_Staking *StakingFilterer) ParseStaked(log types.Log) (*StakingStaked, error) {
	event := new(StakingStaked)
	if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Staking contract.
type StakingTransferIterator struct {
	Event *StakingTransfer // Event containing the contract specifics and raw log

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
func (it *StakingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingTransfer)
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
		it.Event = new(StakingTransfer)
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
func (it *StakingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingTransfer represents a Transfer event raised by the Staking contract.
type StakingTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Staking *StakingFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StakingTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StakingTransferIterator{contract: _Staking.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Staking *StakingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StakingTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingTransfer)
				if err := _Staking.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Staking *StakingFilterer) ParseTransfer(log types.Log) (*StakingTransfer, error) {
	event := new(StakingTransfer)
	if err := _Staking.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Staking contract.
type StakingUnstakedIterator struct {
	Event *StakingUnstaked // Event containing the contract specifics and raw log

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
func (it *StakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnstaked)
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
		it.Event = new(StakingUnstaked)
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
func (it *StakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnstaked represents a Unstaked event raised by the Staking contract.
type StakingUnstaked struct {
	Validator common.Address
	Amount    *big.Int
	Stake     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed _validator, uint256 _amount, uint256 _stake)
func (_Staking *StakingFilterer) FilterUnstaked(opts *bind.FilterOpts, _validator []common.Address) (*StakingUnstakedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unstaked", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingUnstakedIterator{contract: _Staking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed _validator, uint256 _amount, uint256 _stake)
func (_Staking *StakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *StakingUnstaked, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unstaked", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnstaked)
				if err := _Staking.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed _validator, uint256 _amount, uint256 _stake)
func (_Staking *StakingFilterer) ParseUnstaked(log types.Log) (*StakingUnstaked, error) {
	event := new(StakingUnstaked)
	if err := _Staking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorABI is the input ABI used to generate the binding from.
const ValidatorABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"stakeRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_coinbase\",\"type\":\"address\"}],\"name\":\"isCoinbase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"coinbases\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"premineAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"rewardOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"penalizedBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPremineSum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_coinbase\",\"type\":\"address\"}],\"name\":\"join\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unlockedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"coinSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROYALTY_STACK_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCpt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royalValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reportedValidator\",\"type\":\"address\"}],\"name\":\"report\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"premineBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"coinbaseValidators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPremineAddressesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"statusOf\",\"outputs\":[{\"internalType\":\"enumLock.ValidatorStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getFrozenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"premineSum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"cashout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VALIDATOR_STAKE_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_firstName\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_lastName\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_licenseId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_fullAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_state\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_zipcode\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_expirationDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_createdDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_updatedDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_contactEmail\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_isCompany\",\"type\":\"bool\"}],\"name\":\"updateMetadata\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_coinbase\",\"type\":\"address\"}],\"name\":\"getValidatorName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"firstName\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastName\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastJoinLeaveBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"submitProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CPT_ZOOM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getPremineAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"coinbasePosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"leave\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BLOCKS_PER_ROYALTY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMinThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SLASHED_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorCoinbases\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miningKey\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voterMiningKey\",\"type\":\"address\"}],\"name\":\"isValidatorAlreadyVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_coinbase\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"firstName\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastName\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"licenseId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fullAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"state\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"zipcode\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contactEmail\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isCompany\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"safeRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"JOINING_PARAM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_genesisValidators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_genesisCoinbases\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_genesisRVSBalances\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"_premineAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_premineBalances\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_coinbase\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"Joined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_coinbase\",\"type\":\"address\"}],\"name\":\"Left\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Penalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"ProofSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_point\",\"type\":\"uint256\"}],\"name\":\"PointUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"StartTimer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_frozenBalance\",\"type\":\"uint256\"}],\"name\":\"FrozenIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_frozenBalance\",\"type\":\"uint256\"}],\"name\":\"FrozenDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_frozenBalance\",\"type\":\"uint256\"}],\"name\":\"Cashout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_cpt\",\"type\":\"uint256\"}],\"name\":\"CptUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_credit\",\"type\":\"uint256\"}],\"name\":\"CreditUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_lastBalance\",\"type\":\"uint256\"}],\"name\":\"LastBalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_unlockedAt\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"MetadataUpdated\",\"type\":\"event\"}]"

// ValidatorFuncSigs maps the 4-byte function signature to its string representation.
var ValidatorFuncSigs = map[string]string{
	"dc49f0dc": "BLOCKS_PER_ROYALTY()",
	"cadba0d1": "CPT_ZOOM()",
	"fd519fc7": "JOINING_PARAM()",
	"3494e267": "ROYALTY_STACK_LIMIT()",
	"e7065963": "SLASHED_BLOCK()",
	"aae505e6": "VALIDATOR_STAKE_RATE()",
	"538ba4f9": "ZERO_ADDRESS()",
	"70a08231": "balanceOf(address)",
	"73f42561": "burned()",
	"a9e732bb": "cashout(uint256)",
	"4e71d92d": "claim()",
	"31f170c2": "coinSupply()",
	"d12616b8": "coinbasePosition(address)",
	"8a97dc97": "coinbaseValidators(address)",
	"0ebf0de7": "coinbases(uint256)",
	"06661abd": "count()",
	"313ce567": "decimals()",
	"3dda0347": "getCpt()",
	"9f2cfaf1": "getFrozenBalance(address)",
	"e6bbe9dd": "getMinThreshold()",
	"d091afdc": "getPremineAccount(uint256)",
	"93337f3a": "getPremineAddressesCount()",
	"28b8f386": "getPremineSum()",
	"557ed1ba": "getTime()",
	"af2e2da9": "getValidatorName(address)",
	"b7ab4db5": "getValidators()",
	"0d8754ac": "isCoinbase(address)",
	"facd743b": "isValidator(address)",
	"f73294b8": "isValidatorAlreadyVoted(address,address)",
	"28ffe6c8": "join(address)",
	"bbc09519": "lastJoinLeaveBlock()",
	"d66d9e19": "leave()",
	"eacc5210": "lockRange()",
	"06fdde03": "name()",
	"25de9cc6": "penalizedBase()",
	"1afa8d29": "premineAddresses(uint256)",
	"86d7b2bb": "premineBalances(uint256)",
	"a6231cf0": "premineSum()",
	"6a00f42e": "report(address,address)",
	"1d62ebd9": "rewardOf(address)",
	"fc20cc90": "safeRange()",
	"3a4b66f1": "stake()",
	"26476204": "stake(address)",
	"05010105": "stakeRequired()",
	"97a5d5b5": "statusOf(address)",
	"c201be23": "submitProof()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"2cf5db13": "unlockedAt(address)",
	"2e17de78": "unstake(uint256)",
	"ad8490f6": "updateMetadata(bytes32,bytes32,bytes32,string,bytes32,bytes32,uint256,uint256,uint256,uint256,bytes32,bool)",
	"f0fe5ac3": "validatorCoinbases(address)",
	"fa52c7d8": "validators(address)",
}

// ValidatorBin is the compiled bytecode used for deploying new contracts.
var ValidatorBin = "0x6080604052600a600c556003600d5560016015556103e86016553480156200002657600080fd5b5060405162003f7f38038062003f7f83398181016040526101008110156200004d57600080fd5b81019080805160405193929190846401000000008211156200006e57600080fd5b9083019060208201858111156200008457600080fd5b8251866020820283011164010000000082111715620000a257600080fd5b82525081516020918201928201910280838360005b83811015620000d1578181015183820152602001620000b7565b5050505090500160405260200180516040519392919084640100000000821115620000fb57600080fd5b9083019060208201858111156200011157600080fd5b82518660208202830111640100000000821117156200012f57600080fd5b82525081516020918201928201910280838360005b838110156200015e57818101518382015260200162000144565b50505050905001604052602001805160405193929190846401000000008211156200018857600080fd5b9083019060208201858111156200019e57600080fd5b8251866020820283011164010000000082111715620001bc57600080fd5b82525081516020918201928201910280838360005b83811015620001eb578181015183820152602001620001d1565b50505050905001604052602001805160405193929190846401000000008211156200021557600080fd5b9083019060208201858111156200022b57600080fd5b82516401000000008111828201881017156200024657600080fd5b82525081516020918201929091019080838360005b83811015620002755781810151838201526020016200025b565b50505050905090810190601f168015620002a35780820380516001836020036101000a031916815260200191505b5060405260200180516040519392919084640100000000821115620002c757600080fd5b908301906020820185811115620002dd57600080fd5b8251640100000000811182820188101715620002f857600080fd5b82525081516020918201929091019080838360005b83811015620003275781810151838201526020016200030d565b50505050905090810190601f168015620003555780820380516001836020036101000a031916815260200191505b506040818152602083015192018051929491939192846401000000008211156200037e57600080fd5b9083019060208201858111156200039457600080fd5b8251866020820283011164010000000082111715620003b257600080fd5b82525081516020918201928201910280838360005b83811015620003e1578181015183820152602001620003c7565b50505050905001604052602001805160405193929190846401000000008211156200040b57600080fd5b9083019060208201858111156200042157600080fd5b82518660208202830111640100000000821117156200043f57600080fd5b82525081516020918201928201910280838360005b838110156200046e57818101518382015260200162000454565b5050505090500160405250505084848482828282828282600290805190602001906200049c92919062001000565b508151620004b290600390602085019062001000565b506004805460ff191660ff9290921691909117905550508c51600096508695508594509250839150505b818110156200055a578b8181518110620004f257fe5b602002602001015194508c81815181106200050957fe5b602002602001015193508a81815181106200052057fe5b602002602001015192506200053c8484620005c060201b60201c565b6200055184866001600160e01b036200068b16565b600101620004dc565b5050845160005b81811015620005ad57620005a48782815181106200057b57fe5b60200260200101518783815181106200059057fe5b6020026020010151620008b660201b60201c565b60010162000561565b50505050505050505050505050620010a2565b600954620005d99082016001600160e01b036200094116565b620005ec6001600160e01b036200097416565b600062000602836001600160e01b0362000a3a16565b90506200061983836001600160e01b0362000acc16565b6200062e83826001600160e01b0362000beb16565b6001600160a01b0383167f1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90836200066e866001600160e01b0362000c7716565b6040805192835260208301919091528051918290030190a2505050565b81620006a0816001600160e01b0362000c9216565b156200070d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f53656e64657220697320616c726561647920612076616c696461746f72000000604482015290519081900360640190fd5b8162000722816001600160e01b0362000cb216565b156200078f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f436f696e6261736520616c726561647920757365640000000000000000000000604482015290519081900360640190fd5b620007a26001600160e01b0362000cd216565b6001600160a01b0380851660008181526018602090815260408083209590955560178054600181019091557fc624b66cc0138b8fabc209247f72d758e1cf3343756d543badbf24212bed8c150180546001600160a01b03199081168517909155938916808352601a82528583208054861685179055838352601990915293902080549092168317909155907fd12045cbb824dccd3bf8edffc06c251eca57b57aa8af4b78ab2f2653007c808c62000862876001600160e01b0362000c7716565b60408051918252519081900360200190a36200089c84436200088c6001600160e01b0362000cd916565b016001600160e01b0362000cf616565b620008b0846001600160e01b0362000dac16565b50505050565b6013805460018082019092557f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a0900180546001600160a01b0319166001600160a01b0394909416939093179092556014805492830181556000527fce6d7b5282bd9a3661ae061feed1dbda4e52ab073b1f9285be6e155d9c38d4ec909101819055601280549091019055565b600981905560405181907f555b5773a6e2f40d63f33ea8ac35afd7cabd75d29c852e7362faa31676fed7a290600090a250565b600954303111620009855762000a38565b600954600a805430319290920391820190556000620009ac6001600160e01b0362000df016565b620009cc650100000000008462000df660201b620019091790919060201c565b81620009d457fe5b049050620009f38160085462000e7760201b62002aaf1790919060201c565b60085562000a096001600160e01b0362000eec16565b6008546040517f48f8bdfe8a3b1a47754e147e370713fa5cee7be4f2971581513349a1964cf7a190600090a250505b565b6000806501000000000062000a7060085462000a5c8662000c7760201b60201c565b62000df660201b620019091790919060201c565b8162000a7857fe5b6001600160a01b038516600090815260076020526040902054919004915081101562000aa957600091505062000ac7565b6001600160a01b038316600090815260076020526040902054900390505b919050565b6001600160a01b03821662000b4257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b62000b5e8160065462000e7760201b62002aaf1790919060201c565b6006556001600160a01b03821660009081526005602090815260409091205462000b9391839062002aaf62000e77821b17901c565b6001600160a01b03831660008181526005602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b60006501000000000062000c0c60085462000a5c8662000c7760201b60201c565b8162000c1457fe5b04905062000c31828262000f0160201b62002a6d1790919060201c565b6001600160a01b038416600081815260076020526040808220849055517fb8b17c68dbe6cd120dd338473bf86b84c67e766feac2b583d29be17f3259599a9190a3505050565b6001600160a01b031660009081526005602052604090205490565b6001600160a01b039081166000908152601a602052604090205416151590565b6001600160a01b0390811660009081526019602052604090205416151590565b6017545b90565b600062000cee6001600160e01b0362000cd216565b600a02905090565b43811162000d6557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f696e76616c696420756e6c6f636b656420626c6f636b00000000000000000000604482015290519081900360640190fd5b6001600160a01b038216600081815260016020526040808220849055518392917f9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd6000891a35050565b6001600160a01b0381166000818152600e6020526040808220439055517f3d364b4503e0f292c41ea2af3eea2464e1e89d01f1a7d6519063a7342ab865229190a250565b60065490565b60008262000e075750600062000e71565b8282028284828162000e1557fe5b041462000e6e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602181526020018062003f5e6021913960400191505060405180910390fd5b90505b92915050565b60008282018381101562000e6e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b62000a3830316001600160e01b036200094116565b600062000e6e83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525062000f4b60201b60201c565b6000818484111562000ff8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101562000fbc57818101518382015260200162000fa2565b50505050905090810190601f16801562000fea5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200104357805160ff191683800117855562001073565b8280016001018555821562001073579182015b828111156200107357825182559160200191906001019062001056565b506200108192915062001085565b5090565b62000cd691905b808211156200108157600081556001016200108c565b612eac80620010b26000396000f3fe60806040526004361061031a5760003560e01c80638a97dc97116101ab578063cadba0d1116100f7578063eacc521011610095578063fa52c7d81161006f578063fa52c7d814610b8b578063facd743b14610c87578063fc20cc9014610cba578063fd519fc714610ccf5761031a565b8063eacc521014610b08578063f0fe5ac314610b1d578063f73294b814610b505761031a565b8063d66d9e19116100d1578063d66d9e1914610ab4578063dc49f0dc14610ac9578063e6bbe9dd14610ade578063e706596314610af35761031a565b8063cadba0d114610a1f578063d091afdc14610a34578063d12616b814610a815761031a565b8063a9e732bb11610164578063af2e2da91161013e578063af2e2da914610944578063b7ab4db514610990578063bbc09519146109f5578063c201be2314610a0a5761031a565b8063a9e732bb14610813578063aae505e61461083d578063ad8490f6146108525761031a565b80638a97dc971461071757806393337f3a1461074a57806395d89b411461075f57806397a5d5b5146107745780639f2cfaf1146107cb578063a6231cf0146107fe5761031a565b80632e17de781161026a5780634e71d92d116102235780636a00f42e116101fd5780636a00f42e1461066a57806370a08231146106a557806373f42561146106d857806386d7b2bb146106ed5761031a565b80634e71d92d1461062b578063538ba4f914610640578063557ed1ba146106555761031a565b80632e17de781461058f578063313ce567146105b957806331f170c2146105e45780633494e267146105f95780633a4b66f11461060e5780633dda0347146106165761031a565b80631afa8d29116102d757806326476204116102b157806326476204146104f957806328b8f3861461052157806328ffe6c8146105365780632cf5db131461055c5761031a565b80631afa8d29146104875780631d62ebd9146104b157806325de9cc6146104e45761031a565b8063050101051461031f57806306661abd1461034657806306fdde031461035b5780630d8754ac146103e55780630ebf0de71461042c57806318160ddd14610472575b600080fd5b34801561032b57600080fd5b50610334610ce4565b60408051918252519081900360200190f35b34801561035257600080fd5b50610334610cff565b34801561036757600080fd5b50610370610d05565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103aa578181015183820152602001610392565b50505050905090810190601f1680156103d75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103f157600080fd5b506104186004803603602081101561040857600080fd5b50356001600160a01b0316610d98565b604080519115158252519081900360200190f35b34801561043857600080fd5b506104566004803603602081101561044f57600080fd5b5035610dbb565b604080516001600160a01b039092168252519081900360200190f35b34801561047e57600080fd5b50610334610de2565b34801561049357600080fd5b50610456600480360360208110156104aa57600080fd5b5035610de8565b3480156104bd57600080fd5b50610334600480360360208110156104d457600080fd5b50356001600160a01b0316610df5565b3480156104f057600080fd5b50610334610e75565b61051f6004803603602081101561050f57600080fd5b50356001600160a01b0316610e98565b005b34801561052d57600080fd5b50610334610f4f565b61051f6004803603602081101561054c57600080fd5b50356001600160a01b0316610f55565b34801561056857600080fd5b506103346004803603602081101561057f57600080fd5b50356001600160a01b03166110b9565b34801561059b57600080fd5b5061051f600480360360208110156105b257600080fd5b50356110cb565b3480156105c557600080fd5b506105ce611131565b6040805160ff9092168252519081900360200190f35b3480156105f057600080fd5b5061033461113a565b34801561060557600080fd5b5061033461115b565b61051f611161565b34801561062257600080fd5b50610334611214565b34801561063757600080fd5b5061051f61124f565b34801561064c57600080fd5b506104566112af565b34801561066157600080fd5b506103346112b4565b34801561067657600080fd5b5061051f6004803603604081101561068d57600080fd5b506001600160a01b03813581169160200135166112b8565b3480156106b157600080fd5b50610334600480360360208110156106c857600080fd5b50356001600160a01b03166113c8565b3480156106e457600080fd5b506103346113e3565b3480156106f957600080fd5b506103346004803603602081101561071057600080fd5b50356113e9565b34801561072357600080fd5b506104566004803603602081101561073a57600080fd5b50356001600160a01b0316611407565b34801561075657600080fd5b50610334611422565b34801561076b57600080fd5b50610370611428565b34801561078057600080fd5b506107a76004803603602081101561079757600080fd5b50356001600160a01b0316611489565b604051808260028111156107b757fe5b60ff16815260200191505060405180910390f35b3480156107d757600080fd5b50610334600480360360208110156107ee57600080fd5b50356001600160a01b03166114d3565b34801561080a57600080fd5b506103346114ee565b34801561081f57600080fd5b5061051f6004803603602081101561083657600080fd5b50356114f4565b34801561084957600080fd5b50610334611577565b34801561085e57600080fd5b5061051f600480360361018081101561087657600080fd5b813591602081013591604082013591908101906080810160608201356401000000008111156108a457600080fd5b8201836020820111156108b657600080fd5b803590602001918460018302840111640100000000831117156108d857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060208101359060408101359060608101359060808101359060a08101359060c08101359060e00135151561157d565b34801561095057600080fd5b506109776004803603602081101561096757600080fd5b50356001600160a01b03166115ea565b6040805192835260208301919091528051918290030190f35b34801561099c57600080fd5b506109a561161c565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156109e15781810151838201526020016109c9565b505050509050019250505060405180910390f35b348015610a0157600080fd5b5061033461167d565b348015610a1657600080fd5b5061051f611683565b348015610a2b57600080fd5b506103346116a5565b348015610a4057600080fd5b50610a5e60048036036020811015610a5757600080fd5b50356116ad565b604080516001600160a01b03909316835260208301919091528051918290030190f35b348015610a8d57600080fd5b5061033460048036036020811015610aa457600080fd5b50356001600160a01b03166116f7565b348015610ac057600080fd5b5061051f611709565b348015610ad557600080fd5b50610334611823565b348015610aea57600080fd5b50610334611829565b348015610aff57600080fd5b5061033461182e565b348015610b1457600080fd5b50610334611836565b348015610b2957600080fd5b5061045660048036036020811015610b4057600080fd5b50356001600160a01b0316611848565b348015610b5c57600080fd5b5061041860048036036040811015610b7357600080fd5b506001600160a01b0381358116916020013516611863565b348015610b9757600080fd5b50610bbe60048036036020811015610bae57600080fd5b50356001600160a01b031661186c565b604051808d81526020018c81526020018b8152602001806020018a81526020018981526020018881526020018781526020018681526020018581526020018481526020018315151515815260200182810382528b818151815260200191508051906020019080838360005b83811015610c41578181015183820152602001610c29565b50505050905090810190601f168015610c6e5780820380516001836020036101000a031916815260200191505b509d505050505050505050505050505060405180910390f35b348015610c9357600080fd5b5061041860048036036020811015610caa57600080fd5b50356001600160a01b03166118d1565b348015610cc657600080fd5b506103346118f1565b348015610cdb57600080fd5b50610334611903565b6000601654610cf161113a565b81610cf857fe5b0490505b90565b60175490565b60028054604080516020601f6000196101006001871615020190941685900493840181900481028201810190925282815260609390929091830182828015610d8e5780601f10610d6357610100808354040283529160200191610d8e565b820191906000526020600020905b815481529060010190602001808311610d7157829003601f168201915b5050505050905090565b6001600160a01b038181166000908152601960205260409020541615155b919050565b60178181548110610dc857fe5b6000918252602090912001546001600160a01b0316905081565b60065490565b60138181548110610dc857fe5b600080600160281b610e1d610e08611214565b610e11866113c8565b9063ffffffff61190916565b81610e2457fe5b6001600160a01b0385166000908152600760205260409020549190049150811015610e53576000915050610db6565b6001600160a01b03831660009081526007602052604090205490039050919050565b6000600a610e81610cff565b610e89610de2565b81610e9057fe5b0481610cf857fe5b610ea1336118d1565b610ee0576040805162461bcd60e51b81526020600482015260176024820152600080516020612e2e833981519152604482015290519081900360640190fd5b6002610eeb33611489565b6002811115610ef657fe5b1415610f40576040805162461bcd60e51b81526020600482015260146024820152731d985b1a59185d1bdc881a5cc81cdb185cda195960621b604482015290519081900360640190fd5b34610f4b8282611969565b5050565b60125490565b6002610f6033611489565b6002811115610f6b57fe5b1415610fb5576040805162461bcd60e51b81526020600482015260146024820152731d985b1a59185d1bdc881a5cc81cdb185cda195960621b604482015290519081900360640190fd5b43610fbe610cff565b60155402601b5401106110025760405162461bcd60e51b8152600401808060200182810382526029815260200180612d7c6029913960400191505060405180910390fd5b43601b5414156110435760405162461bcd60e51b8152600401808060200182810382526024815260200180612da56024913960400191505060405180910390fd5b333461104f8282611969565b611057610ce4565b611060836113c8565b10156110a6576040805162461bcd60e51b815260206004820152601060248201526f0e6e8c2d6ca40dcdee840cadcdeeaced60831b604482015290519081900360640190fd5b6110b082846119f1565b505043601b5550565b60016020526000908152604090205481565b6110d4336118d1565b611113576040805162461bcd60e51b81526020600482015260176024820152600080516020612e2e833981519152604482015290519081900360640190fd5b3360006111208284611b91565b905061112c8282611c0c565b505050565b60045460ff1690565b60006111446113e3565b61114c611c23565b611154610f4f565b0103905090565b600d5481565b61116a336118d1565b6111a9576040805162461bcd60e51b81526020600482015260176024820152600080516020612e2e833981519152604482015290519081900360640190fd5b60026111b433611489565b60028111156111bf57fe5b1415611209576040805162461bcd60e51b81526020600482015260146024820152731d985b1a59185d1bdc881a5cc81cdb185cda195960621b604482015290519081900360640190fd5b61121233610e98565b565b60095460009030310381611226610de2565b61123a83600160281b63ffffffff61190916565b8161124157fe5b600854919004019250505090565b611258336118d1565b611297576040805162461bcd60e51b81526020600482015260176024820152600080516020612e2e833981519152604482015290519081900360640190fd5b3360006112a382611c51565b9050610f4b8282611c0c565b600081565b4290565b806112c2816118d1565b6112fd5760405162461bcd60e51b8152600401808060200182810382526023815260200180612dc96023913960400191505060405180910390fd5b43601b54141561133e5760405162461bcd60e51b8152600401808060200182810382526024815260200180612da56024913960400191505060405180910390fd5b600061135a848461134d6118f1565b611355610e75565b611cb8565b905060006113688483611e0d565b90506113748482611e91565b61137e8482611f18565b6040805182815290516001600160a01b038616917fce6324a0bfd5e264b9b4039aeeab1620f460c269fa080cbd889681a37aaadc65919081900360200190a2505043601b55505050565b6001600160a01b031660009081526005602052604090205490565b60003190565b601481815481106113f657fe5b600091825260209091200154905081565b6019602052600090815260409020546001600160a01b031681565b60135490565b60038054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610d8e5780601f10610d6357610100808354040283529160200191610d8e565b6001600160a01b038116600090815260016020526040812054438110156114b4576000915050610db6565b600160ff1b81146114c9576001915050610db6565b6002915050610db6565b6001600160a01b03166000908152600b602052604090205490565b601c5481565b60006114ff33611489565b600281111561150a57fe5b1461155c576040805162461bcd60e51b815260206004820152601e60248201527f76616c696461746f72206973206c6f636b6564206f7220736c61736865640000604482015290519081900360640190fd5b611564612027565b3361156f81836120c0565b610f4b612158565b60165481565b611586336118d1565b6115c5576040805162461bcd60e51b81526020600482015260176024820152600080516020612e2e833981519152604482015290519081900360640190fd5b336115db818e8e8e8e8e8e8e8e8e8e8e8e612162565b50505050505050505050505050565b6001600160a01b038082166000908152601960205260408120549091829116611612816122bc565b9250925050915091565b60606017805480602002602001604051908101604052809291908181526020018280548015610d8e57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611656575050505050905090565b601b5481565b416000818152601960205260409020546001600160a01b0316610f4b816122df565b600160281b81565b600080601383815481106116bd57fe5b600091825260209091200154601480546001600160a01b0390921691859081106116e357fe5b906000526020600020015491509150915091565b60186020526000908152604090205481565b611712336118d1565b611751576040805162461bcd60e51b81526020600482015260176024820152600080516020612e2e833981519152604482015290519081900360640190fd5b600061175c33611489565b600281111561176757fe5b146117b9576040805162461bcd60e51b815260206004820152601e60248201527f76616c696461746f72206973206c6f636b6564206f7220736c61736865640000604482015290519081900360640190fd5b43601b5414156117fa5760405162461bcd60e51b8152600401808060200182810382526024815260200180612da56024913960400191505060405180910390fd5b336000611806826123b8565b90506118128282611c0c565b61181b826123da565b505043601b55565b600c5481565b600890565b600160ff1b81565b6000611840610cff565b600a02905090565b601a602052600090815260409020546001600160a01b031681565b60015b92915050565b6001600160a01b0380821660009081526019602052604081205490918291829160609183918291829182918291829182918291166118a9816124b4565b9c509c509c509c509c509c509c509c509c509c509c509c505091939597999b5091939597999b565b6001600160a01b039081166000908152601a602052604090205416151590565b60006118fb610cff565b600202905090565b60155481565b60008261191857506000611866565b8282028284828161192557fe5b04146119625760405162461bcd60e51b8152600401808060200182810382526021815260200180612dec6021913960400191505060405180910390fd5b9392505050565b61197681600954016125d2565b61197e612027565b600061198983612605565b9050611995838361261c565b61199f838261270e565b826001600160a01b03167f1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90836119d4866113c8565b6040805192835260208301919091528051918290030190a2505050565b816119fb816118d1565b15611a4d576040805162461bcd60e51b815260206004820152601d60248201527f53656e64657220697320616c726561647920612076616c696461746f72000000604482015290519081900360640190fd5b81611a5781610d98565b15611aa1576040805162461bcd60e51b815260206004820152601560248201527410dbda5b98985cd948185b1c9958591e481d5cd959605a1b604482015290519081900360640190fd5b611aa9610cff565b6001600160a01b0380851660008181526018602090815260408083209590955560178054600181019091557fc624b66cc0138b8fabc209247f72d758e1cf3343756d543badbf24212bed8c150180546001600160a01b03199081168517909155938916808352601a82528583208054861685179055838352601990915293902080549092168317909155907fd12045cbb824dccd3bf8edffc06c251eca57b57aa8af4b78ab2f2653007c808c611b5e876113c8565b60408051918252519081900360200190a3611b828443611b7c611836565b01612784565b611b8b84612818565b50505050565b600080611b9d84611c51565b9050611ba9848461285c565b611bb484600061270e565b8083016001600160a01b0385167f7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e85611bec886113c8565b6040805192835260208301919091528051918290030190a2949350505050565b611c168282612958565b610f4b8243611b7c611836565b600080600954306001600160a01b03163111611c40576000611c47565b6009543031035b600a540191505090565b6000611c5b612027565b6000611c6683612605565b9050611c7383600061270e565b6040805182815290516001600160a01b038516917fd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a919081900360200190a292915050565b6001600160a01b0383166000908152600e602052604081205484908490611d1a576040805162461bcd60e51b81526020600482015260116024820152701d1a5b595c881b9bdd081cdd185c9d1959607a1b604482015290519081900360640190fd5b6001600160a01b0382166000908152600e60205260409020544390820110611d89576040805162461bcd60e51b815260206004820152601a60248201527f746f6f206561726c7920746f6f2062652070656e616c697a6564000000000000604482015290519081900360640190fd5b611d92876129df565b611d9b86612818565b6001600160a01b0386166000908152601160209081526040808320805460010190819055600f9092529091205410611dd557839250611e03565b6001600160a01b0386166000908152600f60209081526040808320546011909252909120540360020a840292505b5050949350505050565b60008181611e1a856114d3565b905081811015611e48576000611e2f86611c51565b9050611e3b8682611c0c565b611e44866114d3565b9150505b81811015611e7d576000611e5b866123b8565b9050611e678682611c0c565b611e70866123da565b611e79866114d3565b9150505b81811015611e89578091505b509392505050565b6001600160a01b0382166000908152600b6020526040902054611eba908263ffffffff612a6d16565b6001600160a01b0383166000818152600b6020908152604091829020849055815185815290810193909352805191927f366703a486081d66a8ef7d8a544aa1f2a91b5f4906e9d5807c3f3ce27ee2d590929081900390910190a25050565b6000611f23836113c8565b611f2b610de2565b03905080611fa5576000611f3d610de2565b611f5184600160281b63ffffffff61190916565b81611f5857fe5b049050611f7081600854612aaf90919063ffffffff16565b60088190556040517f48f8bdfe8a3b1a47754e147e370713fa5cee7be4f2971581513349a1964cf7a190600090a25050610f4b565b600081611fbc84600160281b63ffffffff61190916565b81611fc357fe5b0490506000611fd185612605565b600854909150611fe7908363ffffffff612aaf16565b600855611ff4858261270e565b6008546040517f48f8bdfe8a3b1a47754e147e370713fa5cee7be4f2971581513349a1964cf7a190600090a25050505050565b60095430311161203657611212565b600954600a805430319290920391820190556000612052610de2565b61206683600160281b63ffffffff61190916565b8161206d57fe5b04905061208581600854612aaf90919063ffffffff16565b600855612090612158565b6008546040517f48f8bdfe8a3b1a47754e147e370713fa5cee7be4f2971581513349a1964cf7a190600090a25050565b6120ca8282611e91565b6040516001600160a01b0383169082156108fc029083906000818181858888f19350505050158015612100573d6000803e3d6000fd5b506001600160a01b0382166000818152600b60209081526040918290205482518581529182015281517f4f32a1005fcf45a64a68bc4258b4e0f5522b7aa6a06506781aa4e6c6395c48d0929181900390910190a25050565b61121230316125d2565b6040518061018001604052808d81526020018c81526020018b81526020018a81526020018981526020018881526020018781526020018681526020018581526020018481526020018381526020018215158152506000808f6001600160a01b03166001600160a01b031681526020019081526020016000206000820151816000015560208201518160010155604082015181600201556060820151816003019080519060200190612214929190612ca1565b506080820151600482015560a0820151600582015560c0820151600682015560e0820151600782015561010082015160088201556101208201516009820155610140820151600a82015561016090910151600b909101805460ff19169115159190911790556040516001600160a01b038e16907fe08224dc11618a1e23016c28d3fb80e30630f4df34d9f95890fe9ce89a85d07e90600090a250505050505050505050505050565b6001600160a01b0316600090815260208190526040902080546001909101549091565b6001600160a01b0381166000908152600e6020526040902054431161233e576040805162461bcd60e51b815260206004820152601060248201526f185b1c9958591e481cdd589b5a5d195960821b604482015290519081900360640190fd5b61234781612818565b6001600160a01b0381166000908152601060205260409020805460010190819055600c54908161237357fe5b0661238157612381816129df565b6040516001600160a01b038216907f87dc06121a3dc1ab46683cb20422ce8b605fd9a4369ad5128f0dc2a49283e84c90600090a250565b6000806123c4836113c8565b905060006123d28483611b91565b949350505050565b6001600160a01b038082166000908152601a6020526040902054166123fe81612b09565b6001600160a01b0382166000908152601a6020526040902080546001600160a01b031916905561242d82612bda565b806001600160a01b0316826001600160a01b03167f4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c560405160405180910390a36003612477610cff565b1015610f4b5760405162461bcd60e51b815260040180806020018281038252602a815260200180612e4e602a913960400191505060405180910390fd5b6001600160a01b038116600090815260208181526040808320805460018083015460028085015460038601805488519581161561010002600019011692909204601f81018990048902850189019097528684529397919693956060959294859485948594859485948594859493909290918301828280156125765780601f1061254b57610100808354040283529160200191612576565b820191906000526020600020905b81548152906001019060200180831161255957829003601f168201915b5050505050995080600401549850806005015497508060060154965080600701549550806008015494508060090154935080600a0154925080600b0160009054906101000a900460ff1691505091939597999b5091939597999b565b600981905560405181907f555b5773a6e2f40d63f33ea8ac35afd7cabd75d29c852e7362faa31676fed7a290600090a250565b600080600160281b610e1d600854610e11866113c8565b6001600160a01b038216612677576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b60065461268a908263ffffffff612aaf16565b6006556001600160a01b0382166000908152600560205260409020546126b6908263ffffffff612aaf16565b6001600160a01b03831660008181526005602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b6000600160281b612724600854610e11866113c8565b8161272b57fe5b04905061273e818363ffffffff612a6d16565b6001600160a01b038416600081815260076020526040808220849055517fb8b17c68dbe6cd120dd338473bf86b84c67e766feac2b583d29be17f3259599a9190a3505050565b4381116127d1576040805162461bcd60e51b8152602060048201526016602482015275696e76616c696420756e6c6f636b656420626c6f636b60501b604482015290519081900360640190fd5b6001600160a01b038216600081815260016020526040808220849055518392917f9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd6000891a35050565b6001600160a01b0381166000818152600e6020526040808220439055517f3d364b4503e0f292c41ea2af3eea2464e1e89d01f1a7d6519063a7342ab865229190a250565b6001600160a01b0382166128a15760405162461bcd60e51b8152600401808060200182810382526021815260200180612e0d6021913960400191505060405180910390fd5b6128e481604051806060016040528060228152602001612d5a602291396001600160a01b038516600090815260056020526040902054919063ffffffff612c0a16565b6001600160a01b038316600090815260056020526040902055600654612910908263ffffffff612a6d16565b6006556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b6001600160a01b0382166000908152600b6020526040902054612981908263ffffffff612aaf16565b6001600160a01b0383166000818152600b6020908152604091829020849055815185815290810193909352805191927faf1eabfa68d345ba5007d2f3fae9c6f58e9a9d3fe16c21e651bd00f9bac4ad4d929081900390910190a25050565b600d546001600160a01b038216600090815260116020908152604080832054600f90925290912054910111612a1357612a6a565b6001600160a01b0381166000818152600f6020908152604091829020805460010190819055825190815291517f88a501824285cc7436ca02dbfc2953e94f577c28429dec36bac9c7782699a9d89281900390910190a25b50565b600061196283836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250612c0a565b600082820183811015611962576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6001600160a01b0381166000908152601860205260408120549060176001612b2f610cff565b0381548110612b3a57fe5b60009182526020808320909101546001600160a01b03168083526018909152604090912083905560178054919250829184908110612b7457fe5b600091825260209091200180546001600160a01b0319166001600160a01b03929092169190911790556017805490612bb0906000198301612d1f565b5050506001600160a01b0316600090815260196020526040902080546001600160a01b0319169055565b6001600160a01b03166000908152600e602090815260408083208390556011825280832054600f90925290912055565b60008184841115612c995760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612c5e578181015183820152602001612c46565b50505050905090810190601f168015612c8b5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10612ce257805160ff1916838001178555612d0f565b82800160010185558215612d0f579182015b82811115612d0f578251825591602001919060010190612cf4565b50612d1b929150612d3f565b5090565b81548183558181111561112c5760008381526020902061112c9181019083015b610cfc91905b80821115612d1b5760008155600101612d4556fe45524332303a206275726e20616d6f756e7420657863656564732062616c616e63656a6f696e696e6720746f6f20717569636b20636f756c642064657374726f7920636f6e63656e7375736f6e6c79206f6e6520616374696f6e206a6f696e2f6c656176652070657220626c6f636b7265706f727465642061646472657373206973206e6f7420612076616c696461746f72536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7745524332303a206275726e2066726f6d20746865207a65726f206164647265737353656e646572206973206e6f742076616c696461746f72000000000000000000636f6e63656e73757320776f726b732077697468206174206c6561737420332076616c696461746f7273a265627a7a723158201038763ad024ba3e03a08bf38a1798aacaf8bf893dc53573699124dfb5ac467d64736f6c634300050b0032536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77"

// DeployValidator deploys a new Ethereum contract, binding an instance of Validator to it.
func DeployValidator(auth *bind.TransactOpts, backend bind.ContractBackend, _genesisValidators []common.Address, _genesisCoinbases []common.Address, _genesisRVSBalances []*big.Int, name string, symbol string, decimals uint8, _premineAddresses []common.Address, _premineBalances []*big.Int) (common.Address, *types.Transaction, *Validator, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorBin), backend, _genesisValidators, _genesisCoinbases, _genesisRVSBalances, name, symbol, decimals, _premineAddresses, _premineBalances)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// Validator is an auto generated Go binding around an Ethereum contract.
type Validator struct {
	ValidatorCaller     // Read-only binding to the contract
	ValidatorTransactor // Write-only binding to the contract
	ValidatorFilterer   // Log filterer for contract events
}

// ValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSession struct {
	Contract     *Validator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorCallerSession struct {
	Contract *ValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorTransactorSession struct {
	Contract     *ValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRaw struct {
	Contract *Validator // Generic contract binding to access the raw methods on
}

// ValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorCallerRaw struct {
	Contract *ValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorTransactorRaw struct {
	Contract *ValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidator creates a new instance of Validator, bound to a specific deployed contract.
func NewValidator(address common.Address, backend bind.ContractBackend) (*Validator, error) {
	contract, err := bindValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// NewValidatorCaller creates a new read-only instance of Validator, bound to a specific deployed contract.
func NewValidatorCaller(address common.Address, caller bind.ContractCaller) (*ValidatorCaller, error) {
	contract, err := bindValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorCaller{contract: contract}, nil
}

// NewValidatorTransactor creates a new write-only instance of Validator, bound to a specific deployed contract.
func NewValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorTransactor, error) {
	contract, err := bindValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorTransactor{contract: contract}, nil
}

// NewValidatorFilterer creates a new log filterer instance of Validator, bound to a specific deployed contract.
func NewValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorFilterer, error) {
	contract, err := bindValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorFilterer{contract: contract}, nil
}

// bindValidator binds a generic wrapper to an already deployed contract.
func bindValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.ValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transact(opts, method, params...)
}

// BLOCKSPERROYALTY is a free data retrieval call binding the contract method 0xdc49f0dc.
//
// Solidity: function BLOCKS_PER_ROYALTY() constant returns(uint256)
func (_Validator *ValidatorCaller) BLOCKSPERROYALTY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "BLOCKS_PER_ROYALTY")
	return *ret0, err
}

// BLOCKSPERROYALTY is a free data retrieval call binding the contract method 0xdc49f0dc.
//
// Solidity: function BLOCKS_PER_ROYALTY() constant returns(uint256)
func (_Validator *ValidatorSession) BLOCKSPERROYALTY() (*big.Int, error) {
	return _Validator.Contract.BLOCKSPERROYALTY(&_Validator.CallOpts)
}

// BLOCKSPERROYALTY is a free data retrieval call binding the contract method 0xdc49f0dc.
//
// Solidity: function BLOCKS_PER_ROYALTY() constant returns(uint256)
func (_Validator *ValidatorCallerSession) BLOCKSPERROYALTY() (*big.Int, error) {
	return _Validator.Contract.BLOCKSPERROYALTY(&_Validator.CallOpts)
}

// CPTZOOM is a free data retrieval call binding the contract method 0xcadba0d1.
//
// Solidity: function CPT_ZOOM() constant returns(uint256)
func (_Validator *ValidatorCaller) CPTZOOM(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "CPT_ZOOM")
	return *ret0, err
}

// CPTZOOM is a free data retrieval call binding the contract method 0xcadba0d1.
//
// Solidity: function CPT_ZOOM() constant returns(uint256)
func (_Validator *ValidatorSession) CPTZOOM() (*big.Int, error) {
	return _Validator.Contract.CPTZOOM(&_Validator.CallOpts)
}

// CPTZOOM is a free data retrieval call binding the contract method 0xcadba0d1.
//
// Solidity: function CPT_ZOOM() constant returns(uint256)
func (_Validator *ValidatorCallerSession) CPTZOOM() (*big.Int, error) {
	return _Validator.Contract.CPTZOOM(&_Validator.CallOpts)
}

// JOININGPARAM is a free data retrieval call binding the contract method 0xfd519fc7.
//
// Solidity: function JOINING_PARAM() constant returns(uint256)
func (_Validator *ValidatorCaller) JOININGPARAM(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "JOINING_PARAM")
	return *ret0, err
}

// JOININGPARAM is a free data retrieval call binding the contract method 0xfd519fc7.
//
// Solidity: function JOINING_PARAM() constant returns(uint256)
func (_Validator *ValidatorSession) JOININGPARAM() (*big.Int, error) {
	return _Validator.Contract.JOININGPARAM(&_Validator.CallOpts)
}

// JOININGPARAM is a free data retrieval call binding the contract method 0xfd519fc7.
//
// Solidity: function JOINING_PARAM() constant returns(uint256)
func (_Validator *ValidatorCallerSession) JOININGPARAM() (*big.Int, error) {
	return _Validator.Contract.JOININGPARAM(&_Validator.CallOpts)
}

// ROYALTYSTACKLIMIT is a free data retrieval call binding the contract method 0x3494e267.
//
// Solidity: function ROYALTY_STACK_LIMIT() constant returns(uint256)
func (_Validator *ValidatorCaller) ROYALTYSTACKLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "ROYALTY_STACK_LIMIT")
	return *ret0, err
}

// ROYALTYSTACKLIMIT is a free data retrieval call binding the contract method 0x3494e267.
//
// Solidity: function ROYALTY_STACK_LIMIT() constant returns(uint256)
func (_Validator *ValidatorSession) ROYALTYSTACKLIMIT() (*big.Int, error) {
	return _Validator.Contract.ROYALTYSTACKLIMIT(&_Validator.CallOpts)
}

// ROYALTYSTACKLIMIT is a free data retrieval call binding the contract method 0x3494e267.
//
// Solidity: function ROYALTY_STACK_LIMIT() constant returns(uint256)
func (_Validator *ValidatorCallerSession) ROYALTYSTACKLIMIT() (*big.Int, error) {
	return _Validator.Contract.ROYALTYSTACKLIMIT(&_Validator.CallOpts)
}

// SLASHEDBLOCK is a free data retrieval call binding the contract method 0xe7065963.
//
// Solidity: function SLASHED_BLOCK() constant returns(uint256)
func (_Validator *ValidatorCaller) SLASHEDBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "SLASHED_BLOCK")
	return *ret0, err
}

// SLASHEDBLOCK is a free data retrieval call binding the contract method 0xe7065963.
//
// Solidity: function SLASHED_BLOCK() constant returns(uint256)
func (_Validator *ValidatorSession) SLASHEDBLOCK() (*big.Int, error) {
	return _Validator.Contract.SLASHEDBLOCK(&_Validator.CallOpts)
}

// SLASHEDBLOCK is a free data retrieval call binding the contract method 0xe7065963.
//
// Solidity: function SLASHED_BLOCK() constant returns(uint256)
func (_Validator *ValidatorCallerSession) SLASHEDBLOCK() (*big.Int, error) {
	return _Validator.Contract.SLASHEDBLOCK(&_Validator.CallOpts)
}

// VALIDATORSTAKERATE is a free data retrieval call binding the contract method 0xaae505e6.
//
// Solidity: function VALIDATOR_STAKE_RATE() constant returns(uint256)
func (_Validator *ValidatorCaller) VALIDATORSTAKERATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "VALIDATOR_STAKE_RATE")
	return *ret0, err
}

// VALIDATORSTAKERATE is a free data retrieval call binding the contract method 0xaae505e6.
//
// Solidity: function VALIDATOR_STAKE_RATE() constant returns(uint256)
func (_Validator *ValidatorSession) VALIDATORSTAKERATE() (*big.Int, error) {
	return _Validator.Contract.VALIDATORSTAKERATE(&_Validator.CallOpts)
}

// VALIDATORSTAKERATE is a free data retrieval call binding the contract method 0xaae505e6.
//
// Solidity: function VALIDATOR_STAKE_RATE() constant returns(uint256)
func (_Validator *ValidatorCallerSession) VALIDATORSTAKERATE() (*big.Int, error) {
	return _Validator.Contract.VALIDATORSTAKERATE(&_Validator.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Validator *ValidatorCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Validator *ValidatorSession) ZEROADDRESS() (common.Address, error) {
	return _Validator.Contract.ZEROADDRESS(&_Validator.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Validator *ValidatorCallerSession) ZEROADDRESS() (common.Address, error) {
	return _Validator.Contract.ZEROADDRESS(&_Validator.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Validator *ValidatorCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Validator *ValidatorSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Validator.Contract.BalanceOf(&_Validator.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Validator *ValidatorCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Validator.Contract.BalanceOf(&_Validator.CallOpts, account)
}

// Burned is a free data retrieval call binding the contract method 0x73f42561.
//
// Solidity: function burned() constant returns(uint256)
func (_Validator *ValidatorCaller) Burned(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "burned")
	return *ret0, err
}

// Burned is a free data retrieval call binding the contract method 0x73f42561.
//
// Solidity: function burned() constant returns(uint256)
func (_Validator *ValidatorSession) Burned() (*big.Int, error) {
	return _Validator.Contract.Burned(&_Validator.CallOpts)
}

// Burned is a free data retrieval call binding the contract method 0x73f42561.
//
// Solidity: function burned() constant returns(uint256)
func (_Validator *ValidatorCallerSession) Burned() (*big.Int, error) {
	return _Validator.Contract.Burned(&_Validator.CallOpts)
}

// CoinSupply is a free data retrieval call binding the contract method 0x31f170c2.
//
// Solidity: function coinSupply() constant returns(uint256)
func (_Validator *ValidatorCaller) CoinSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "coinSupply")
	return *ret0, err
}

// CoinSupply is a free data retrieval call binding the contract method 0x31f170c2.
//
// Solidity: function coinSupply() constant returns(uint256)
func (_Validator *ValidatorSession) CoinSupply() (*big.Int, error) {
	return _Validator.Contract.CoinSupply(&_Validator.CallOpts)
}

// CoinSupply is a free data retrieval call binding the contract method 0x31f170c2.
//
// Solidity: function coinSupply() constant returns(uint256)
func (_Validator *ValidatorCallerSession) CoinSupply() (*big.Int, error) {
	return _Validator.Contract.CoinSupply(&_Validator.CallOpts)
}

// CoinbasePosition is a free data retrieval call binding the contract method 0xd12616b8.
//
// Solidity: function coinbasePosition(address ) constant returns(uint256)
func (_Validator *ValidatorCaller) CoinbasePosition(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "coinbasePosition", arg0)
	return *ret0, err
}

// CoinbasePosition is a free data retrieval call binding the contract method 0xd12616b8.
//
// Solidity: function coinbasePosition(address ) constant returns(uint256)
func (_Validator *ValidatorSession) CoinbasePosition(arg0 common.Address) (*big.Int, error) {
	return _Validator.Contract.CoinbasePosition(&_Validator.CallOpts, arg0)
}

// CoinbasePosition is a free data retrieval call binding the contract method 0xd12616b8.
//
// Solidity: function coinbasePosition(address ) constant returns(uint256)
func (_Validator *ValidatorCallerSession) CoinbasePosition(arg0 common.Address) (*big.Int, error) {
	return _Validator.Contract.CoinbasePosition(&_Validator.CallOpts, arg0)
}

// CoinbaseValidators is a free data retrieval call binding the contract method 0x8a97dc97.
//
// Solidity: function coinbaseValidators(address ) constant returns(address)
func (_Validator *ValidatorCaller) CoinbaseValidators(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "coinbaseValidators", arg0)
	return *ret0, err
}

// CoinbaseValidators is a free data retrieval call binding the contract method 0x8a97dc97.
//
// Solidity: function coinbaseValidators(address ) constant returns(address)
func (_Validator *ValidatorSession) CoinbaseValidators(arg0 common.Address) (common.Address, error) {
	return _Validator.Contract.CoinbaseValidators(&_Validator.CallOpts, arg0)
}

// CoinbaseValidators is a free data retrieval call binding the contract method 0x8a97dc97.
//
// Solidity: function coinbaseValidators(address ) constant returns(address)
func (_Validator *ValidatorCallerSession) CoinbaseValidators(arg0 common.Address) (common.Address, error) {
	return _Validator.Contract.CoinbaseValidators(&_Validator.CallOpts, arg0)
}

// Coinbases is a free data retrieval call binding the contract method 0x0ebf0de7.
//
// Solidity: function coinbases(uint256 ) constant returns(address)
func (_Validator *ValidatorCaller) Coinbases(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "coinbases", arg0)
	return *ret0, err
}

// Coinbases is a free data retrieval call binding the contract method 0x0ebf0de7.
//
// Solidity: function coinbases(uint256 ) constant returns(address)
func (_Validator *ValidatorSession) Coinbases(arg0 *big.Int) (common.Address, error) {
	return _Validator.Contract.Coinbases(&_Validator.CallOpts, arg0)
}

// Coinbases is a free data retrieval call binding the contract method 0x0ebf0de7.
//
// Solidity: function coinbases(uint256 ) constant returns(address)
func (_Validator *ValidatorCallerSession) Coinbases(arg0 *big.Int) (common.Address, error) {
	return _Validator.Contract.Coinbases(&_Validator.CallOpts, arg0)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256)
func (_Validator *ValidatorCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "count")
	return *ret0, err
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256)
func (_Validator *ValidatorSession) Count() (*big.Int, error) {
	return _Validator.Contract.Count(&_Validator.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256)
func (_Validator *ValidatorCallerSession) Count() (*big.Int, error) {
	return _Validator.Contract.Count(&_Validator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Validator *ValidatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Validator *ValidatorSession) Decimals() (uint8, error) {
	return _Validator.Contract.Decimals(&_Validator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Validator *ValidatorCallerSession) Decimals() (uint8, error) {
	return _Validator.Contract.Decimals(&_Validator.CallOpts)
}

// GetCpt is a free data retrieval call binding the contract method 0x3dda0347.
//
// Solidity: function getCpt() constant returns(uint256)
func (_Validator *ValidatorCaller) GetCpt(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "getCpt")
	return *ret0, err
}

// GetCpt is a free data retrieval call binding the contract method 0x3dda0347.
//
// Solidity: function getCpt() constant returns(uint256)
func (_Validator *ValidatorSession) GetCpt() (*big.Int, error) {
	return _Validator.Contract.GetCpt(&_Validator.CallOpts)
}

// GetCpt is a free data retrieval call binding the contract method 0x3dda0347.
//
// Solidity: function getCpt() constant returns(uint256)
func (_Validator *ValidatorCallerSession) GetCpt() (*big.Int, error) {
	return _Validator.Contract.GetCpt(&_Validator.CallOpts)
}

// GetFrozenBalance is a free data retrieval call binding the contract method 0x9f2cfaf1.
//
// Solidity: function getFrozenBalance(address _validator) constant returns(uint256)
func (_Validator *ValidatorCaller) GetFrozenBalance(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "getFrozenBalance", _validator)
	return *ret0, err
}

// GetFrozenBalance is a free data retrieval call binding the contract method 0x9f2cfaf1.
//
// Solidity: function getFrozenBalance(address _validator) constant returns(uint256)
func (_Validator *ValidatorSession) GetFrozenBalance(_validator common.Address) (*big.Int, error) {
	return _Validator.Contract.GetFrozenBalance(&_Validator.CallOpts, _validator)
}

// GetFrozenBalance is a free data retrieval call binding the contract method 0x9f2cfaf1.
//
// Solidity: function getFrozenBalance(address _validator) constant returns(uint256)
func (_Validator *ValidatorCallerSession) GetFrozenBalance(_validator common.Address) (*big.Int, error) {
	return _Validator.Contract.GetFrozenBalance(&_Validator.CallOpts, _validator)
}

// GetMinThreshold is a free data retrieval call binding the contract method 0xe6bbe9dd.
//
// Solidity: function getMinThreshold() constant returns(uint256)
func (_Validator *ValidatorCaller) GetMinThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "getMinThreshold")
	return *ret0, err
}

// GetMinThreshold is a free data retrieval call binding the contract method 0xe6bbe9dd.
//
// Solidity: function getMinThreshold() constant returns(uint256)
func (_Validator *ValidatorSession) GetMinThreshold() (*big.Int, error) {
	return _Validator.Contract.GetMinThreshold(&_Validator.CallOpts)
}

// GetMinThreshold is a free data retrieval call binding the contract method 0xe6bbe9dd.
//
// Solidity: function getMinThreshold() constant returns(uint256)
func (_Validator *ValidatorCallerSession) GetMinThreshold() (*big.Int, error) {
	return _Validator.Contract.GetMinThreshold(&_Validator.CallOpts)
}

// GetPremineAccount is a free data retrieval call binding the contract method 0xd091afdc.
//
// Solidity: function getPremineAccount(uint256 i) constant returns(address, uint256)
func (_Validator *ValidatorCaller) GetPremineAccount(opts *bind.CallOpts, i *big.Int) (common.Address, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Validator.contract.Call(opts, out, "getPremineAccount", i)
	return *ret0, *ret1, err
}

// GetPremineAccount is a free data retrieval call binding the contract method 0xd091afdc.
//
// Solidity: function getPremineAccount(uint256 i) constant returns(address, uint256)
func (_Validator *ValidatorSession) GetPremineAccount(i *big.Int) (common.Address, *big.Int, error) {
	return _Validator.Contract.GetPremineAccount(&_Validator.CallOpts, i)
}

// GetPremineAccount is a free data retrieval call binding the contract method 0xd091afdc.
//
// Solidity: function getPremineAccount(uint256 i) constant returns(address, uint256)
func (_Validator *ValidatorCallerSession) GetPremineAccount(i *big.Int) (common.Address, *big.Int, error) {
	return _Validator.Contract.GetPremineAccount(&_Validator.CallOpts, i)
}

// GetPremineAddressesCount is a free data retrieval call binding the contract method 0x93337f3a.
//
// Solidity: function getPremineAddressesCount() constant returns(uint256)
func (_Validator *ValidatorCaller) GetPremineAddressesCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "getPremineAddressesCount")
	return *ret0, err
}

// GetPremineAddressesCount is a free data retrieval call binding the contract method 0x93337f3a.
//
// Solidity: function getPremineAddressesCount() constant returns(uint256)
func (_Validator *ValidatorSession) GetPremineAddressesCount() (*big.Int, error) {
	return _Validator.Contract.GetPremineAddressesCount(&_Validator.CallOpts)
}

// GetPremineAddressesCount is a free data retrieval call binding the contract method 0x93337f3a.
//
// Solidity: function getPremineAddressesCount() constant returns(uint256)
func (_Validator *ValidatorCallerSession) GetPremineAddressesCount() (*big.Int, error) {
	return _Validator.Contract.GetPremineAddressesCount(&_Validator.CallOpts)
}

// GetPremineSum is a free data retrieval call binding the contract method 0x28b8f386.
//
// Solidity: function getPremineSum() constant returns(uint256)
func (_Validator *ValidatorCaller) GetPremineSum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "getPremineSum")
	return *ret0, err
}

// GetPremineSum is a free data retrieval call binding the contract method 0x28b8f386.
//
// Solidity: function getPremineSum() constant returns(uint256)
func (_Validator *ValidatorSession) GetPremineSum() (*big.Int, error) {
	return _Validator.Contract.GetPremineSum(&_Validator.CallOpts)
}

// GetPremineSum is a free data retrieval call binding the contract method 0x28b8f386.
//
// Solidity: function getPremineSum() constant returns(uint256)
func (_Validator *ValidatorCallerSession) GetPremineSum() (*big.Int, error) {
	return _Validator.Contract.GetPremineSum(&_Validator.CallOpts)
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() constant returns(uint256)
func (_Validator *ValidatorCaller) GetTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "getTime")
	return *ret0, err
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() constant returns(uint256)
func (_Validator *ValidatorSession) GetTime() (*big.Int, error) {
	return _Validator.Contract.GetTime(&_Validator.CallOpts)
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() constant returns(uint256)
func (_Validator *ValidatorCallerSession) GetTime() (*big.Int, error) {
	return _Validator.Contract.GetTime(&_Validator.CallOpts)
}

// GetValidatorName is a free data retrieval call binding the contract method 0xaf2e2da9.
//
// Solidity: function getValidatorName(address _coinbase) constant returns(bytes32 firstName, bytes32 lastName)
func (_Validator *ValidatorCaller) GetValidatorName(opts *bind.CallOpts, _coinbase common.Address) (struct {
	FirstName [32]byte
	LastName  [32]byte
}, error) {
	ret := new(struct {
		FirstName [32]byte
		LastName  [32]byte
	})
	out := ret
	err := _Validator.contract.Call(opts, out, "getValidatorName", _coinbase)
	return *ret, err
}

// GetValidatorName is a free data retrieval call binding the contract method 0xaf2e2da9.
//
// Solidity: function getValidatorName(address _coinbase) constant returns(bytes32 firstName, bytes32 lastName)
func (_Validator *ValidatorSession) GetValidatorName(_coinbase common.Address) (struct {
	FirstName [32]byte
	LastName  [32]byte
}, error) {
	return _Validator.Contract.GetValidatorName(&_Validator.CallOpts, _coinbase)
}

// GetValidatorName is a free data retrieval call binding the contract method 0xaf2e2da9.
//
// Solidity: function getValidatorName(address _coinbase) constant returns(bytes32 firstName, bytes32 lastName)
func (_Validator *ValidatorCallerSession) GetValidatorName(_coinbase common.Address) (struct {
	FirstName [32]byte
	LastName  [32]byte
}, error) {
	return _Validator.Contract.GetValidatorName(&_Validator.CallOpts, _coinbase)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() constant returns(address[])
func (_Validator *ValidatorCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "getValidators")
	return *ret0, err
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() constant returns(address[])
func (_Validator *ValidatorSession) GetValidators() ([]common.Address, error) {
	return _Validator.Contract.GetValidators(&_Validator.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() constant returns(address[])
func (_Validator *ValidatorCallerSession) GetValidators() ([]common.Address, error) {
	return _Validator.Contract.GetValidators(&_Validator.CallOpts)
}

// IsCoinbase is a free data retrieval call binding the contract method 0x0d8754ac.
//
// Solidity: function isCoinbase(address _coinbase) constant returns(bool)
func (_Validator *ValidatorCaller) IsCoinbase(opts *bind.CallOpts, _coinbase common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "isCoinbase", _coinbase)
	return *ret0, err
}

// IsCoinbase is a free data retrieval call binding the contract method 0x0d8754ac.
//
// Solidity: function isCoinbase(address _coinbase) constant returns(bool)
func (_Validator *ValidatorSession) IsCoinbase(_coinbase common.Address) (bool, error) {
	return _Validator.Contract.IsCoinbase(&_Validator.CallOpts, _coinbase)
}

// IsCoinbase is a free data retrieval call binding the contract method 0x0d8754ac.
//
// Solidity: function isCoinbase(address _coinbase) constant returns(bool)
func (_Validator *ValidatorCallerSession) IsCoinbase(_coinbase common.Address) (bool, error) {
	return _Validator.Contract.IsCoinbase(&_Validator.CallOpts, _coinbase)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _validator) constant returns(bool)
func (_Validator *ValidatorCaller) IsValidator(opts *bind.CallOpts, _validator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "isValidator", _validator)
	return *ret0, err
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _validator) constant returns(bool)
func (_Validator *ValidatorSession) IsValidator(_validator common.Address) (bool, error) {
	return _Validator.Contract.IsValidator(&_Validator.CallOpts, _validator)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _validator) constant returns(bool)
func (_Validator *ValidatorCallerSession) IsValidator(_validator common.Address) (bool, error) {
	return _Validator.Contract.IsValidator(&_Validator.CallOpts, _validator)
}

// IsValidatorAlreadyVoted is a free data retrieval call binding the contract method 0xf73294b8.
//
// Solidity: function isValidatorAlreadyVoted(address _miningKey, address _voterMiningKey) constant returns(bool)
func (_Validator *ValidatorCaller) IsValidatorAlreadyVoted(opts *bind.CallOpts, _miningKey common.Address, _voterMiningKey common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "isValidatorAlreadyVoted", _miningKey, _voterMiningKey)
	return *ret0, err
}

// IsValidatorAlreadyVoted is a free data retrieval call binding the contract method 0xf73294b8.
//
// Solidity: function isValidatorAlreadyVoted(address _miningKey, address _voterMiningKey) constant returns(bool)
func (_Validator *ValidatorSession) IsValidatorAlreadyVoted(_miningKey common.Address, _voterMiningKey common.Address) (bool, error) {
	return _Validator.Contract.IsValidatorAlreadyVoted(&_Validator.CallOpts, _miningKey, _voterMiningKey)
}

// IsValidatorAlreadyVoted is a free data retrieval call binding the contract method 0xf73294b8.
//
// Solidity: function isValidatorAlreadyVoted(address _miningKey, address _voterMiningKey) constant returns(bool)
func (_Validator *ValidatorCallerSession) IsValidatorAlreadyVoted(_miningKey common.Address, _voterMiningKey common.Address) (bool, error) {
	return _Validator.Contract.IsValidatorAlreadyVoted(&_Validator.CallOpts, _miningKey, _voterMiningKey)
}

// LastJoinLeaveBlock is a free data retrieval call binding the contract method 0xbbc09519.
//
// Solidity: function lastJoinLeaveBlock() constant returns(uint256)
func (_Validator *ValidatorCaller) LastJoinLeaveBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "lastJoinLeaveBlock")
	return *ret0, err
}

// LastJoinLeaveBlock is a free data retrieval call binding the contract method 0xbbc09519.
//
// Solidity: function lastJoinLeaveBlock() constant returns(uint256)
func (_Validator *ValidatorSession) LastJoinLeaveBlock() (*big.Int, error) {
	return _Validator.Contract.LastJoinLeaveBlock(&_Validator.CallOpts)
}

// LastJoinLeaveBlock is a free data retrieval call binding the contract method 0xbbc09519.
//
// Solidity: function lastJoinLeaveBlock() constant returns(uint256)
func (_Validator *ValidatorCallerSession) LastJoinLeaveBlock() (*big.Int, error) {
	return _Validator.Contract.LastJoinLeaveBlock(&_Validator.CallOpts)
}

// LockRange is a free data retrieval call binding the contract method 0xeacc5210.
//
// Solidity: function lockRange() constant returns(uint256)
func (_Validator *ValidatorCaller) LockRange(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "lockRange")
	return *ret0, err
}

// LockRange is a free data retrieval call binding the contract method 0xeacc5210.
//
// Solidity: function lockRange() constant returns(uint256)
func (_Validator *ValidatorSession) LockRange() (*big.Int, error) {
	return _Validator.Contract.LockRange(&_Validator.CallOpts)
}

// LockRange is a free data retrieval call binding the contract method 0xeacc5210.
//
// Solidity: function lockRange() constant returns(uint256)
func (_Validator *ValidatorCallerSession) LockRange() (*big.Int, error) {
	return _Validator.Contract.LockRange(&_Validator.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Validator *ValidatorCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Validator *ValidatorSession) Name() (string, error) {
	return _Validator.Contract.Name(&_Validator.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Validator *ValidatorCallerSession) Name() (string, error) {
	return _Validator.Contract.Name(&_Validator.CallOpts)
}

// PenalizedBase is a free data retrieval call binding the contract method 0x25de9cc6.
//
// Solidity: function penalizedBase() constant returns(uint256)
func (_Validator *ValidatorCaller) PenalizedBase(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "penalizedBase")
	return *ret0, err
}

// PenalizedBase is a free data retrieval call binding the contract method 0x25de9cc6.
//
// Solidity: function penalizedBase() constant returns(uint256)
func (_Validator *ValidatorSession) PenalizedBase() (*big.Int, error) {
	return _Validator.Contract.PenalizedBase(&_Validator.CallOpts)
}

// PenalizedBase is a free data retrieval call binding the contract method 0x25de9cc6.
//
// Solidity: function penalizedBase() constant returns(uint256)
func (_Validator *ValidatorCallerSession) PenalizedBase() (*big.Int, error) {
	return _Validator.Contract.PenalizedBase(&_Validator.CallOpts)
}

// PremineAddresses is a free data retrieval call binding the contract method 0x1afa8d29.
//
// Solidity: function premineAddresses(uint256 ) constant returns(address)
func (_Validator *ValidatorCaller) PremineAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "premineAddresses", arg0)
	return *ret0, err
}

// PremineAddresses is a free data retrieval call binding the contract method 0x1afa8d29.
//
// Solidity: function premineAddresses(uint256 ) constant returns(address)
func (_Validator *ValidatorSession) PremineAddresses(arg0 *big.Int) (common.Address, error) {
	return _Validator.Contract.PremineAddresses(&_Validator.CallOpts, arg0)
}

// PremineAddresses is a free data retrieval call binding the contract method 0x1afa8d29.
//
// Solidity: function premineAddresses(uint256 ) constant returns(address)
func (_Validator *ValidatorCallerSession) PremineAddresses(arg0 *big.Int) (common.Address, error) {
	return _Validator.Contract.PremineAddresses(&_Validator.CallOpts, arg0)
}

// PremineBalances is a free data retrieval call binding the contract method 0x86d7b2bb.
//
// Solidity: function premineBalances(uint256 ) constant returns(uint256)
func (_Validator *ValidatorCaller) PremineBalances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "premineBalances", arg0)
	return *ret0, err
}

// PremineBalances is a free data retrieval call binding the contract method 0x86d7b2bb.
//
// Solidity: function premineBalances(uint256 ) constant returns(uint256)
func (_Validator *ValidatorSession) PremineBalances(arg0 *big.Int) (*big.Int, error) {
	return _Validator.Contract.PremineBalances(&_Validator.CallOpts, arg0)
}

// PremineBalances is a free data retrieval call binding the contract method 0x86d7b2bb.
//
// Solidity: function premineBalances(uint256 ) constant returns(uint256)
func (_Validator *ValidatorCallerSession) PremineBalances(arg0 *big.Int) (*big.Int, error) {
	return _Validator.Contract.PremineBalances(&_Validator.CallOpts, arg0)
}

// PremineSum is a free data retrieval call binding the contract method 0xa6231cf0.
//
// Solidity: function premineSum() constant returns(uint256)
func (_Validator *ValidatorCaller) PremineSum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "premineSum")
	return *ret0, err
}

// PremineSum is a free data retrieval call binding the contract method 0xa6231cf0.
//
// Solidity: function premineSum() constant returns(uint256)
func (_Validator *ValidatorSession) PremineSum() (*big.Int, error) {
	return _Validator.Contract.PremineSum(&_Validator.CallOpts)
}

// PremineSum is a free data retrieval call binding the contract method 0xa6231cf0.
//
// Solidity: function premineSum() constant returns(uint256)
func (_Validator *ValidatorCallerSession) PremineSum() (*big.Int, error) {
	return _Validator.Contract.PremineSum(&_Validator.CallOpts)
}

// RewardOf is a free data retrieval call binding the contract method 0x1d62ebd9.
//
// Solidity: function rewardOf(address _validator) constant returns(uint256)
func (_Validator *ValidatorCaller) RewardOf(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "rewardOf", _validator)
	return *ret0, err
}

// RewardOf is a free data retrieval call binding the contract method 0x1d62ebd9.
//
// Solidity: function rewardOf(address _validator) constant returns(uint256)
func (_Validator *ValidatorSession) RewardOf(_validator common.Address) (*big.Int, error) {
	return _Validator.Contract.RewardOf(&_Validator.CallOpts, _validator)
}

// RewardOf is a free data retrieval call binding the contract method 0x1d62ebd9.
//
// Solidity: function rewardOf(address _validator) constant returns(uint256)
func (_Validator *ValidatorCallerSession) RewardOf(_validator common.Address) (*big.Int, error) {
	return _Validator.Contract.RewardOf(&_Validator.CallOpts, _validator)
}

// SafeRange is a free data retrieval call binding the contract method 0xfc20cc90.
//
// Solidity: function safeRange() constant returns(uint256)
func (_Validator *ValidatorCaller) SafeRange(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "safeRange")
	return *ret0, err
}

// SafeRange is a free data retrieval call binding the contract method 0xfc20cc90.
//
// Solidity: function safeRange() constant returns(uint256)
func (_Validator *ValidatorSession) SafeRange() (*big.Int, error) {
	return _Validator.Contract.SafeRange(&_Validator.CallOpts)
}

// SafeRange is a free data retrieval call binding the contract method 0xfc20cc90.
//
// Solidity: function safeRange() constant returns(uint256)
func (_Validator *ValidatorCallerSession) SafeRange() (*big.Int, error) {
	return _Validator.Contract.SafeRange(&_Validator.CallOpts)
}

// StakeRequired is a free data retrieval call binding the contract method 0x05010105.
//
// Solidity: function stakeRequired() constant returns(uint256)
func (_Validator *ValidatorCaller) StakeRequired(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "stakeRequired")
	return *ret0, err
}

// StakeRequired is a free data retrieval call binding the contract method 0x05010105.
//
// Solidity: function stakeRequired() constant returns(uint256)
func (_Validator *ValidatorSession) StakeRequired() (*big.Int, error) {
	return _Validator.Contract.StakeRequired(&_Validator.CallOpts)
}

// StakeRequired is a free data retrieval call binding the contract method 0x05010105.
//
// Solidity: function stakeRequired() constant returns(uint256)
func (_Validator *ValidatorCallerSession) StakeRequired() (*big.Int, error) {
	return _Validator.Contract.StakeRequired(&_Validator.CallOpts)
}

// StatusOf is a free data retrieval call binding the contract method 0x97a5d5b5.
//
// Solidity: function statusOf(address _validator) constant returns(uint8)
func (_Validator *ValidatorCaller) StatusOf(opts *bind.CallOpts, _validator common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "statusOf", _validator)
	return *ret0, err
}

// StatusOf is a free data retrieval call binding the contract method 0x97a5d5b5.
//
// Solidity: function statusOf(address _validator) constant returns(uint8)
func (_Validator *ValidatorSession) StatusOf(_validator common.Address) (uint8, error) {
	return _Validator.Contract.StatusOf(&_Validator.CallOpts, _validator)
}

// StatusOf is a free data retrieval call binding the contract method 0x97a5d5b5.
//
// Solidity: function statusOf(address _validator) constant returns(uint8)
func (_Validator *ValidatorCallerSession) StatusOf(_validator common.Address) (uint8, error) {
	return _Validator.Contract.StatusOf(&_Validator.CallOpts, _validator)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Validator *ValidatorCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Validator *ValidatorSession) Symbol() (string, error) {
	return _Validator.Contract.Symbol(&_Validator.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Validator *ValidatorCallerSession) Symbol() (string, error) {
	return _Validator.Contract.Symbol(&_Validator.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Validator *ValidatorCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Validator *ValidatorSession) TotalSupply() (*big.Int, error) {
	return _Validator.Contract.TotalSupply(&_Validator.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Validator *ValidatorCallerSession) TotalSupply() (*big.Int, error) {
	return _Validator.Contract.TotalSupply(&_Validator.CallOpts)
}

// UnlockedAt is a free data retrieval call binding the contract method 0x2cf5db13.
//
// Solidity: function unlockedAt(address ) constant returns(uint256)
func (_Validator *ValidatorCaller) UnlockedAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "unlockedAt", arg0)
	return *ret0, err
}

// UnlockedAt is a free data retrieval call binding the contract method 0x2cf5db13.
//
// Solidity: function unlockedAt(address ) constant returns(uint256)
func (_Validator *ValidatorSession) UnlockedAt(arg0 common.Address) (*big.Int, error) {
	return _Validator.Contract.UnlockedAt(&_Validator.CallOpts, arg0)
}

// UnlockedAt is a free data retrieval call binding the contract method 0x2cf5db13.
//
// Solidity: function unlockedAt(address ) constant returns(uint256)
func (_Validator *ValidatorCallerSession) UnlockedAt(arg0 common.Address) (*big.Int, error) {
	return _Validator.Contract.UnlockedAt(&_Validator.CallOpts, arg0)
}

// ValidatorCoinbases is a free data retrieval call binding the contract method 0xf0fe5ac3.
//
// Solidity: function validatorCoinbases(address ) constant returns(address)
func (_Validator *ValidatorCaller) ValidatorCoinbases(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "validatorCoinbases", arg0)
	return *ret0, err
}

// ValidatorCoinbases is a free data retrieval call binding the contract method 0xf0fe5ac3.
//
// Solidity: function validatorCoinbases(address ) constant returns(address)
func (_Validator *ValidatorSession) ValidatorCoinbases(arg0 common.Address) (common.Address, error) {
	return _Validator.Contract.ValidatorCoinbases(&_Validator.CallOpts, arg0)
}

// ValidatorCoinbases is a free data retrieval call binding the contract method 0xf0fe5ac3.
//
// Solidity: function validatorCoinbases(address ) constant returns(address)
func (_Validator *ValidatorCallerSession) ValidatorCoinbases(arg0 common.Address) (common.Address, error) {
	return _Validator.Contract.ValidatorCoinbases(&_Validator.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address _coinbase) constant returns(bytes32 firstName, bytes32 lastName, bytes32 licenseId, string fullAddress, bytes32 state, bytes32 zipcode, uint256 expirationDate, uint256 createdDate, uint256 updatedDate, uint256 minThreshold, bytes32 contactEmail, bool isCompany)
func (_Validator *ValidatorCaller) Validators(opts *bind.CallOpts, _coinbase common.Address) (struct {
	FirstName      [32]byte
	LastName       [32]byte
	LicenseId      [32]byte
	FullAddress    string
	State          [32]byte
	Zipcode        [32]byte
	ExpirationDate *big.Int
	CreatedDate    *big.Int
	UpdatedDate    *big.Int
	MinThreshold   *big.Int
	ContactEmail   [32]byte
	IsCompany      bool
}, error) {
	ret := new(struct {
		FirstName      [32]byte
		LastName       [32]byte
		LicenseId      [32]byte
		FullAddress    string
		State          [32]byte
		Zipcode        [32]byte
		ExpirationDate *big.Int
		CreatedDate    *big.Int
		UpdatedDate    *big.Int
		MinThreshold   *big.Int
		ContactEmail   [32]byte
		IsCompany      bool
	})
	out := ret
	err := _Validator.contract.Call(opts, out, "validators", _coinbase)
	return *ret, err
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address _coinbase) constant returns(bytes32 firstName, bytes32 lastName, bytes32 licenseId, string fullAddress, bytes32 state, bytes32 zipcode, uint256 expirationDate, uint256 createdDate, uint256 updatedDate, uint256 minThreshold, bytes32 contactEmail, bool isCompany)
func (_Validator *ValidatorSession) Validators(_coinbase common.Address) (struct {
	FirstName      [32]byte
	LastName       [32]byte
	LicenseId      [32]byte
	FullAddress    string
	State          [32]byte
	Zipcode        [32]byte
	ExpirationDate *big.Int
	CreatedDate    *big.Int
	UpdatedDate    *big.Int
	MinThreshold   *big.Int
	ContactEmail   [32]byte
	IsCompany      bool
}, error) {
	return _Validator.Contract.Validators(&_Validator.CallOpts, _coinbase)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address _coinbase) constant returns(bytes32 firstName, bytes32 lastName, bytes32 licenseId, string fullAddress, bytes32 state, bytes32 zipcode, uint256 expirationDate, uint256 createdDate, uint256 updatedDate, uint256 minThreshold, bytes32 contactEmail, bool isCompany)
func (_Validator *ValidatorCallerSession) Validators(_coinbase common.Address) (struct {
	FirstName      [32]byte
	LastName       [32]byte
	LicenseId      [32]byte
	FullAddress    string
	State          [32]byte
	Zipcode        [32]byte
	ExpirationDate *big.Int
	CreatedDate    *big.Int
	UpdatedDate    *big.Int
	MinThreshold   *big.Int
	ContactEmail   [32]byte
	IsCompany      bool
}, error) {
	return _Validator.Contract.Validators(&_Validator.CallOpts, _coinbase)
}

// Cashout is a paid mutator transaction binding the contract method 0xa9e732bb.
//
// Solidity: function cashout(uint256 _amount) returns()
func (_Validator *ValidatorTransactor) Cashout(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "cashout", _amount)
}

// Cashout is a paid mutator transaction binding the contract method 0xa9e732bb.
//
// Solidity: function cashout(uint256 _amount) returns()
func (_Validator *ValidatorSession) Cashout(_amount *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.Cashout(&_Validator.TransactOpts, _amount)
}

// Cashout is a paid mutator transaction binding the contract method 0xa9e732bb.
//
// Solidity: function cashout(uint256 _amount) returns()
func (_Validator *ValidatorTransactorSession) Cashout(_amount *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.Cashout(&_Validator.TransactOpts, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Validator *ValidatorTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Validator *ValidatorSession) Claim() (*types.Transaction, error) {
	return _Validator.Contract.Claim(&_Validator.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Validator *ValidatorTransactorSession) Claim() (*types.Transaction, error) {
	return _Validator.Contract.Claim(&_Validator.TransactOpts)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(address _coinbase) returns()
func (_Validator *ValidatorTransactor) Join(opts *bind.TransactOpts, _coinbase common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "join", _coinbase)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(address _coinbase) returns()
func (_Validator *ValidatorSession) Join(_coinbase common.Address) (*types.Transaction, error) {
	return _Validator.Contract.Join(&_Validator.TransactOpts, _coinbase)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(address _coinbase) returns()
func (_Validator *ValidatorTransactorSession) Join(_coinbase common.Address) (*types.Transaction, error) {
	return _Validator.Contract.Join(&_Validator.TransactOpts, _coinbase)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_Validator *ValidatorTransactor) Leave(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "leave")
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_Validator *ValidatorSession) Leave() (*types.Transaction, error) {
	return _Validator.Contract.Leave(&_Validator.TransactOpts)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_Validator *ValidatorTransactorSession) Leave() (*types.Transaction, error) {
	return _Validator.Contract.Leave(&_Validator.TransactOpts)
}

// Report is a paid mutator transaction binding the contract method 0x6a00f42e.
//
// Solidity: function report(address _royalValidator, address _reportedValidator) returns()
func (_Validator *ValidatorTransactor) Report(opts *bind.TransactOpts, _royalValidator common.Address, _reportedValidator common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "report", _royalValidator, _reportedValidator)
}

// Report is a paid mutator transaction binding the contract method 0x6a00f42e.
//
// Solidity: function report(address _royalValidator, address _reportedValidator) returns()
func (_Validator *ValidatorSession) Report(_royalValidator common.Address, _reportedValidator common.Address) (*types.Transaction, error) {
	return _Validator.Contract.Report(&_Validator.TransactOpts, _royalValidator, _reportedValidator)
}

// Report is a paid mutator transaction binding the contract method 0x6a00f42e.
//
// Solidity: function report(address _royalValidator, address _reportedValidator) returns()
func (_Validator *ValidatorTransactorSession) Report(_royalValidator common.Address, _reportedValidator common.Address) (*types.Transaction, error) {
	return _Validator.Contract.Report(&_Validator.TransactOpts, _royalValidator, _reportedValidator)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address _validator) returns()
func (_Validator *ValidatorTransactor) Stake(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "stake", _validator)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address _validator) returns()
func (_Validator *ValidatorSession) Stake(_validator common.Address) (*types.Transaction, error) {
	return _Validator.Contract.Stake(&_Validator.TransactOpts, _validator)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address _validator) returns()
func (_Validator *ValidatorTransactorSession) Stake(_validator common.Address) (*types.Transaction, error) {
	return _Validator.Contract.Stake(&_Validator.TransactOpts, _validator)
}

// Stake0 is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_Validator *ValidatorTransactor) Stake0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "stake0")
}

// Stake0 is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_Validator *ValidatorSession) Stake0() (*types.Transaction, error) {
	return _Validator.Contract.Stake0(&_Validator.TransactOpts)
}

// Stake0 is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_Validator *ValidatorTransactorSession) Stake0() (*types.Transaction, error) {
	return _Validator.Contract.Stake0(&_Validator.TransactOpts)
}

// SubmitProof is a paid mutator transaction binding the contract method 0xc201be23.
//
// Solidity: function submitProof() returns()
func (_Validator *ValidatorTransactor) SubmitProof(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "submitProof")
}

// SubmitProof is a paid mutator transaction binding the contract method 0xc201be23.
//
// Solidity: function submitProof() returns()
func (_Validator *ValidatorSession) SubmitProof() (*types.Transaction, error) {
	return _Validator.Contract.SubmitProof(&_Validator.TransactOpts)
}

// SubmitProof is a paid mutator transaction binding the contract method 0xc201be23.
//
// Solidity: function submitProof() returns()
func (_Validator *ValidatorTransactorSession) SubmitProof() (*types.Transaction, error) {
	return _Validator.Contract.SubmitProof(&_Validator.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_Validator *ValidatorTransactor) Unstake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "unstake", _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_Validator *ValidatorSession) Unstake(_amount *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.Unstake(&_Validator.TransactOpts, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_Validator *ValidatorTransactorSession) Unstake(_amount *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.Unstake(&_Validator.TransactOpts, _amount)
}

// UpdateMetadata is a paid mutator transaction binding the contract method 0xad8490f6.
//
// Solidity: function updateMetadata(bytes32 _firstName, bytes32 _lastName, bytes32 _licenseId, string _fullAddress, bytes32 _state, bytes32 _zipcode, uint256 _expirationDate, uint256 _createdDate, uint256 _updatedDate, uint256 _minThreshold, bytes32 _contactEmail, bool _isCompany) returns()
func (_Validator *ValidatorTransactor) UpdateMetadata(opts *bind.TransactOpts, _firstName [32]byte, _lastName [32]byte, _licenseId [32]byte, _fullAddress string, _state [32]byte, _zipcode [32]byte, _expirationDate *big.Int, _createdDate *big.Int, _updatedDate *big.Int, _minThreshold *big.Int, _contactEmail [32]byte, _isCompany bool) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "updateMetadata", _firstName, _lastName, _licenseId, _fullAddress, _state, _zipcode, _expirationDate, _createdDate, _updatedDate, _minThreshold, _contactEmail, _isCompany)
}

// UpdateMetadata is a paid mutator transaction binding the contract method 0xad8490f6.
//
// Solidity: function updateMetadata(bytes32 _firstName, bytes32 _lastName, bytes32 _licenseId, string _fullAddress, bytes32 _state, bytes32 _zipcode, uint256 _expirationDate, uint256 _createdDate, uint256 _updatedDate, uint256 _minThreshold, bytes32 _contactEmail, bool _isCompany) returns()
func (_Validator *ValidatorSession) UpdateMetadata(_firstName [32]byte, _lastName [32]byte, _licenseId [32]byte, _fullAddress string, _state [32]byte, _zipcode [32]byte, _expirationDate *big.Int, _createdDate *big.Int, _updatedDate *big.Int, _minThreshold *big.Int, _contactEmail [32]byte, _isCompany bool) (*types.Transaction, error) {
	return _Validator.Contract.UpdateMetadata(&_Validator.TransactOpts, _firstName, _lastName, _licenseId, _fullAddress, _state, _zipcode, _expirationDate, _createdDate, _updatedDate, _minThreshold, _contactEmail, _isCompany)
}

// UpdateMetadata is a paid mutator transaction binding the contract method 0xad8490f6.
//
// Solidity: function updateMetadata(bytes32 _firstName, bytes32 _lastName, bytes32 _licenseId, string _fullAddress, bytes32 _state, bytes32 _zipcode, uint256 _expirationDate, uint256 _createdDate, uint256 _updatedDate, uint256 _minThreshold, bytes32 _contactEmail, bool _isCompany) returns()
func (_Validator *ValidatorTransactorSession) UpdateMetadata(_firstName [32]byte, _lastName [32]byte, _licenseId [32]byte, _fullAddress string, _state [32]byte, _zipcode [32]byte, _expirationDate *big.Int, _createdDate *big.Int, _updatedDate *big.Int, _minThreshold *big.Int, _contactEmail [32]byte, _isCompany bool) (*types.Transaction, error) {
	return _Validator.Contract.UpdateMetadata(&_Validator.TransactOpts, _firstName, _lastName, _licenseId, _fullAddress, _state, _zipcode, _expirationDate, _createdDate, _updatedDate, _minThreshold, _contactEmail, _isCompany)
}

// ValidatorCashoutIterator is returned from FilterCashout and is used to iterate over the raw logs and unpacked data for Cashout events raised by the Validator contract.
type ValidatorCashoutIterator struct {
	Event *ValidatorCashout // Event containing the contract specifics and raw log

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
func (it *ValidatorCashoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorCashout)
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
		it.Event = new(ValidatorCashout)
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
func (it *ValidatorCashoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorCashoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorCashout represents a Cashout event raised by the Validator contract.
type ValidatorCashout struct {
	Validator     common.Address
	Amount        *big.Int
	FrozenBalance *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCashout is a free log retrieval operation binding the contract event 0x4f32a1005fcf45a64a68bc4258b4e0f5522b7aa6a06506781aa4e6c6395c48d0.
//
// Solidity: event Cashout(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Validator *ValidatorFilterer) FilterCashout(opts *bind.FilterOpts, _validator []common.Address) (*ValidatorCashoutIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "Cashout", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorCashoutIterator{contract: _Validator.contract, event: "Cashout", logs: logs, sub: sub}, nil
}

// WatchCashout is a free log subscription operation binding the contract event 0x4f32a1005fcf45a64a68bc4258b4e0f5522b7aa6a06506781aa4e6c6395c48d0.
//
// Solidity: event Cashout(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Validator *ValidatorFilterer) WatchCashout(opts *bind.WatchOpts, sink chan<- *ValidatorCashout, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "Cashout", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorCashout)
				if err := _Validator.contract.UnpackLog(event, "Cashout", log); err != nil {
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

// ParseCashout is a log parse operation binding the contract event 0x4f32a1005fcf45a64a68bc4258b4e0f5522b7aa6a06506781aa4e6c6395c48d0.
//
// Solidity: event Cashout(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Validator *ValidatorFilterer) ParseCashout(log types.Log) (*ValidatorCashout, error) {
	event := new(ValidatorCashout)
	if err := _Validator.contract.UnpackLog(event, "Cashout", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Validator contract.
type ValidatorClaimedIterator struct {
	Event *ValidatorClaimed // Event containing the contract specifics and raw log

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
func (it *ValidatorClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorClaimed)
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
		it.Event = new(ValidatorClaimed)
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
func (it *ValidatorClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorClaimed represents a Claimed event raised by the Validator contract.
type ValidatorClaimed struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed _validator, uint256 _amount)
func (_Validator *ValidatorFilterer) FilterClaimed(opts *bind.FilterOpts, _validator []common.Address) (*ValidatorClaimedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "Claimed", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorClaimedIterator{contract: _Validator.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed _validator, uint256 _amount)
func (_Validator *ValidatorFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *ValidatorClaimed, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "Claimed", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorClaimed)
				if err := _Validator.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed _validator, uint256 _amount)
func (_Validator *ValidatorFilterer) ParseClaimed(log types.Log) (*ValidatorClaimed, error) {
	event := new(ValidatorClaimed)
	if err := _Validator.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorCptUpdatedIterator is returned from FilterCptUpdated and is used to iterate over the raw logs and unpacked data for CptUpdated events raised by the Validator contract.
type ValidatorCptUpdatedIterator struct {
	Event *ValidatorCptUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorCptUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorCptUpdated)
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
		it.Event = new(ValidatorCptUpdated)
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
func (it *ValidatorCptUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorCptUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorCptUpdated represents a CptUpdated event raised by the Validator contract.
type ValidatorCptUpdated struct {
	Cpt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCptUpdated is a free log retrieval operation binding the contract event 0x48f8bdfe8a3b1a47754e147e370713fa5cee7be4f2971581513349a1964cf7a1.
//
// Solidity: event CptUpdated(uint256 indexed _cpt)
func (_Validator *ValidatorFilterer) FilterCptUpdated(opts *bind.FilterOpts, _cpt []*big.Int) (*ValidatorCptUpdatedIterator, error) {

	var _cptRule []interface{}
	for _, _cptItem := range _cpt {
		_cptRule = append(_cptRule, _cptItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "CptUpdated", _cptRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorCptUpdatedIterator{contract: _Validator.contract, event: "CptUpdated", logs: logs, sub: sub}, nil
}

// WatchCptUpdated is a free log subscription operation binding the contract event 0x48f8bdfe8a3b1a47754e147e370713fa5cee7be4f2971581513349a1964cf7a1.
//
// Solidity: event CptUpdated(uint256 indexed _cpt)
func (_Validator *ValidatorFilterer) WatchCptUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorCptUpdated, _cpt []*big.Int) (event.Subscription, error) {

	var _cptRule []interface{}
	for _, _cptItem := range _cpt {
		_cptRule = append(_cptRule, _cptItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "CptUpdated", _cptRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorCptUpdated)
				if err := _Validator.contract.UnpackLog(event, "CptUpdated", log); err != nil {
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

// ParseCptUpdated is a log parse operation binding the contract event 0x48f8bdfe8a3b1a47754e147e370713fa5cee7be4f2971581513349a1964cf7a1.
//
// Solidity: event CptUpdated(uint256 indexed _cpt)
func (_Validator *ValidatorFilterer) ParseCptUpdated(log types.Log) (*ValidatorCptUpdated, error) {
	event := new(ValidatorCptUpdated)
	if err := _Validator.contract.UnpackLog(event, "CptUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorCreditUpdatedIterator is returned from FilterCreditUpdated and is used to iterate over the raw logs and unpacked data for CreditUpdated events raised by the Validator contract.
type ValidatorCreditUpdatedIterator struct {
	Event *ValidatorCreditUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorCreditUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorCreditUpdated)
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
		it.Event = new(ValidatorCreditUpdated)
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
func (it *ValidatorCreditUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorCreditUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorCreditUpdated represents a CreditUpdated event raised by the Validator contract.
type ValidatorCreditUpdated struct {
	Validator common.Address
	Credit    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreditUpdated is a free log retrieval operation binding the contract event 0xb8b17c68dbe6cd120dd338473bf86b84c67e766feac2b583d29be17f3259599a.
//
// Solidity: event CreditUpdated(address indexed _validator, uint256 indexed _credit)
func (_Validator *ValidatorFilterer) FilterCreditUpdated(opts *bind.FilterOpts, _validator []common.Address, _credit []*big.Int) (*ValidatorCreditUpdatedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}
	var _creditRule []interface{}
	for _, _creditItem := range _credit {
		_creditRule = append(_creditRule, _creditItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "CreditUpdated", _validatorRule, _creditRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorCreditUpdatedIterator{contract: _Validator.contract, event: "CreditUpdated", logs: logs, sub: sub}, nil
}

// WatchCreditUpdated is a free log subscription operation binding the contract event 0xb8b17c68dbe6cd120dd338473bf86b84c67e766feac2b583d29be17f3259599a.
//
// Solidity: event CreditUpdated(address indexed _validator, uint256 indexed _credit)
func (_Validator *ValidatorFilterer) WatchCreditUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorCreditUpdated, _validator []common.Address, _credit []*big.Int) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}
	var _creditRule []interface{}
	for _, _creditItem := range _credit {
		_creditRule = append(_creditRule, _creditItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "CreditUpdated", _validatorRule, _creditRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorCreditUpdated)
				if err := _Validator.contract.UnpackLog(event, "CreditUpdated", log); err != nil {
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

// ParseCreditUpdated is a log parse operation binding the contract event 0xb8b17c68dbe6cd120dd338473bf86b84c67e766feac2b583d29be17f3259599a.
//
// Solidity: event CreditUpdated(address indexed _validator, uint256 indexed _credit)
func (_Validator *ValidatorFilterer) ParseCreditUpdated(log types.Log) (*ValidatorCreditUpdated, error) {
	event := new(ValidatorCreditUpdated)
	if err := _Validator.contract.UnpackLog(event, "CreditUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorFrozenDecreasedIterator is returned from FilterFrozenDecreased and is used to iterate over the raw logs and unpacked data for FrozenDecreased events raised by the Validator contract.
type ValidatorFrozenDecreasedIterator struct {
	Event *ValidatorFrozenDecreased // Event containing the contract specifics and raw log

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
func (it *ValidatorFrozenDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorFrozenDecreased)
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
		it.Event = new(ValidatorFrozenDecreased)
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
func (it *ValidatorFrozenDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorFrozenDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorFrozenDecreased represents a FrozenDecreased event raised by the Validator contract.
type ValidatorFrozenDecreased struct {
	Validator     common.Address
	Amount        *big.Int
	FrozenBalance *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFrozenDecreased is a free log retrieval operation binding the contract event 0x366703a486081d66a8ef7d8a544aa1f2a91b5f4906e9d5807c3f3ce27ee2d590.
//
// Solidity: event FrozenDecreased(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Validator *ValidatorFilterer) FilterFrozenDecreased(opts *bind.FilterOpts, _validator []common.Address) (*ValidatorFrozenDecreasedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "FrozenDecreased", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorFrozenDecreasedIterator{contract: _Validator.contract, event: "FrozenDecreased", logs: logs, sub: sub}, nil
}

// WatchFrozenDecreased is a free log subscription operation binding the contract event 0x366703a486081d66a8ef7d8a544aa1f2a91b5f4906e9d5807c3f3ce27ee2d590.
//
// Solidity: event FrozenDecreased(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Validator *ValidatorFilterer) WatchFrozenDecreased(opts *bind.WatchOpts, sink chan<- *ValidatorFrozenDecreased, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "FrozenDecreased", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorFrozenDecreased)
				if err := _Validator.contract.UnpackLog(event, "FrozenDecreased", log); err != nil {
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

// ParseFrozenDecreased is a log parse operation binding the contract event 0x366703a486081d66a8ef7d8a544aa1f2a91b5f4906e9d5807c3f3ce27ee2d590.
//
// Solidity: event FrozenDecreased(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Validator *ValidatorFilterer) ParseFrozenDecreased(log types.Log) (*ValidatorFrozenDecreased, error) {
	event := new(ValidatorFrozenDecreased)
	if err := _Validator.contract.UnpackLog(event, "FrozenDecreased", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorFrozenIncreasedIterator is returned from FilterFrozenIncreased and is used to iterate over the raw logs and unpacked data for FrozenIncreased events raised by the Validator contract.
type ValidatorFrozenIncreasedIterator struct {
	Event *ValidatorFrozenIncreased // Event containing the contract specifics and raw log

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
func (it *ValidatorFrozenIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorFrozenIncreased)
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
		it.Event = new(ValidatorFrozenIncreased)
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
func (it *ValidatorFrozenIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorFrozenIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorFrozenIncreased represents a FrozenIncreased event raised by the Validator contract.
type ValidatorFrozenIncreased struct {
	Validator     common.Address
	Amount        *big.Int
	FrozenBalance *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFrozenIncreased is a free log retrieval operation binding the contract event 0xaf1eabfa68d345ba5007d2f3fae9c6f58e9a9d3fe16c21e651bd00f9bac4ad4d.
//
// Solidity: event FrozenIncreased(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Validator *ValidatorFilterer) FilterFrozenIncreased(opts *bind.FilterOpts, _validator []common.Address) (*ValidatorFrozenIncreasedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "FrozenIncreased", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorFrozenIncreasedIterator{contract: _Validator.contract, event: "FrozenIncreased", logs: logs, sub: sub}, nil
}

// WatchFrozenIncreased is a free log subscription operation binding the contract event 0xaf1eabfa68d345ba5007d2f3fae9c6f58e9a9d3fe16c21e651bd00f9bac4ad4d.
//
// Solidity: event FrozenIncreased(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Validator *ValidatorFilterer) WatchFrozenIncreased(opts *bind.WatchOpts, sink chan<- *ValidatorFrozenIncreased, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "FrozenIncreased", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorFrozenIncreased)
				if err := _Validator.contract.UnpackLog(event, "FrozenIncreased", log); err != nil {
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

// ParseFrozenIncreased is a log parse operation binding the contract event 0xaf1eabfa68d345ba5007d2f3fae9c6f58e9a9d3fe16c21e651bd00f9bac4ad4d.
//
// Solidity: event FrozenIncreased(address indexed _validator, uint256 _amount, uint256 _frozenBalance)
func (_Validator *ValidatorFilterer) ParseFrozenIncreased(log types.Log) (*ValidatorFrozenIncreased, error) {
	event := new(ValidatorFrozenIncreased)
	if err := _Validator.contract.UnpackLog(event, "FrozenIncreased", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorJoinedIterator is returned from FilterJoined and is used to iterate over the raw logs and unpacked data for Joined events raised by the Validator contract.
type ValidatorJoinedIterator struct {
	Event *ValidatorJoined // Event containing the contract specifics and raw log

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
func (it *ValidatorJoinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorJoined)
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
		it.Event = new(ValidatorJoined)
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
func (it *ValidatorJoinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorJoinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorJoined represents a Joined event raised by the Validator contract.
type ValidatorJoined struct {
	Validator common.Address
	Coinbase  common.Address
	Stake     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterJoined is a free log retrieval operation binding the contract event 0xd12045cbb824dccd3bf8edffc06c251eca57b57aa8af4b78ab2f2653007c808c.
//
// Solidity: event Joined(address indexed _validator, address indexed _coinbase, uint256 _stake)
func (_Validator *ValidatorFilterer) FilterJoined(opts *bind.FilterOpts, _validator []common.Address, _coinbase []common.Address) (*ValidatorJoinedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}
	var _coinbaseRule []interface{}
	for _, _coinbaseItem := range _coinbase {
		_coinbaseRule = append(_coinbaseRule, _coinbaseItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "Joined", _validatorRule, _coinbaseRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorJoinedIterator{contract: _Validator.contract, event: "Joined", logs: logs, sub: sub}, nil
}

// WatchJoined is a free log subscription operation binding the contract event 0xd12045cbb824dccd3bf8edffc06c251eca57b57aa8af4b78ab2f2653007c808c.
//
// Solidity: event Joined(address indexed _validator, address indexed _coinbase, uint256 _stake)
func (_Validator *ValidatorFilterer) WatchJoined(opts *bind.WatchOpts, sink chan<- *ValidatorJoined, _validator []common.Address, _coinbase []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}
	var _coinbaseRule []interface{}
	for _, _coinbaseItem := range _coinbase {
		_coinbaseRule = append(_coinbaseRule, _coinbaseItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "Joined", _validatorRule, _coinbaseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorJoined)
				if err := _Validator.contract.UnpackLog(event, "Joined", log); err != nil {
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

// ParseJoined is a log parse operation binding the contract event 0xd12045cbb824dccd3bf8edffc06c251eca57b57aa8af4b78ab2f2653007c808c.
//
// Solidity: event Joined(address indexed _validator, address indexed _coinbase, uint256 _stake)
func (_Validator *ValidatorFilterer) ParseJoined(log types.Log) (*ValidatorJoined, error) {
	event := new(ValidatorJoined)
	if err := _Validator.contract.UnpackLog(event, "Joined", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorLastBalanceUpdatedIterator is returned from FilterLastBalanceUpdated and is used to iterate over the raw logs and unpacked data for LastBalanceUpdated events raised by the Validator contract.
type ValidatorLastBalanceUpdatedIterator struct {
	Event *ValidatorLastBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorLastBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorLastBalanceUpdated)
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
		it.Event = new(ValidatorLastBalanceUpdated)
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
func (it *ValidatorLastBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorLastBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorLastBalanceUpdated represents a LastBalanceUpdated event raised by the Validator contract.
type ValidatorLastBalanceUpdated struct {
	LastBalance *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLastBalanceUpdated is a free log retrieval operation binding the contract event 0x555b5773a6e2f40d63f33ea8ac35afd7cabd75d29c852e7362faa31676fed7a2.
//
// Solidity: event LastBalanceUpdated(uint256 indexed _lastBalance)
func (_Validator *ValidatorFilterer) FilterLastBalanceUpdated(opts *bind.FilterOpts, _lastBalance []*big.Int) (*ValidatorLastBalanceUpdatedIterator, error) {

	var _lastBalanceRule []interface{}
	for _, _lastBalanceItem := range _lastBalance {
		_lastBalanceRule = append(_lastBalanceRule, _lastBalanceItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "LastBalanceUpdated", _lastBalanceRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorLastBalanceUpdatedIterator{contract: _Validator.contract, event: "LastBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchLastBalanceUpdated is a free log subscription operation binding the contract event 0x555b5773a6e2f40d63f33ea8ac35afd7cabd75d29c852e7362faa31676fed7a2.
//
// Solidity: event LastBalanceUpdated(uint256 indexed _lastBalance)
func (_Validator *ValidatorFilterer) WatchLastBalanceUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorLastBalanceUpdated, _lastBalance []*big.Int) (event.Subscription, error) {

	var _lastBalanceRule []interface{}
	for _, _lastBalanceItem := range _lastBalance {
		_lastBalanceRule = append(_lastBalanceRule, _lastBalanceItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "LastBalanceUpdated", _lastBalanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorLastBalanceUpdated)
				if err := _Validator.contract.UnpackLog(event, "LastBalanceUpdated", log); err != nil {
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

// ParseLastBalanceUpdated is a log parse operation binding the contract event 0x555b5773a6e2f40d63f33ea8ac35afd7cabd75d29c852e7362faa31676fed7a2.
//
// Solidity: event LastBalanceUpdated(uint256 indexed _lastBalance)
func (_Validator *ValidatorFilterer) ParseLastBalanceUpdated(log types.Log) (*ValidatorLastBalanceUpdated, error) {
	event := new(ValidatorLastBalanceUpdated)
	if err := _Validator.contract.UnpackLog(event, "LastBalanceUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorLeftIterator is returned from FilterLeft and is used to iterate over the raw logs and unpacked data for Left events raised by the Validator contract.
type ValidatorLeftIterator struct {
	Event *ValidatorLeft // Event containing the contract specifics and raw log

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
func (it *ValidatorLeftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorLeft)
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
		it.Event = new(ValidatorLeft)
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
func (it *ValidatorLeftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorLeftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorLeft represents a Left event raised by the Validator contract.
type ValidatorLeft struct {
	Validator common.Address
	Coinbase  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLeft is a free log retrieval operation binding the contract event 0x4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5.
//
// Solidity: event Left(address indexed _validator, address indexed _coinbase)
func (_Validator *ValidatorFilterer) FilterLeft(opts *bind.FilterOpts, _validator []common.Address, _coinbase []common.Address) (*ValidatorLeftIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}
	var _coinbaseRule []interface{}
	for _, _coinbaseItem := range _coinbase {
		_coinbaseRule = append(_coinbaseRule, _coinbaseItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "Left", _validatorRule, _coinbaseRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorLeftIterator{contract: _Validator.contract, event: "Left", logs: logs, sub: sub}, nil
}

// WatchLeft is a free log subscription operation binding the contract event 0x4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5.
//
// Solidity: event Left(address indexed _validator, address indexed _coinbase)
func (_Validator *ValidatorFilterer) WatchLeft(opts *bind.WatchOpts, sink chan<- *ValidatorLeft, _validator []common.Address, _coinbase []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}
	var _coinbaseRule []interface{}
	for _, _coinbaseItem := range _coinbase {
		_coinbaseRule = append(_coinbaseRule, _coinbaseItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "Left", _validatorRule, _coinbaseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorLeft)
				if err := _Validator.contract.UnpackLog(event, "Left", log); err != nil {
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

// ParseLeft is a log parse operation binding the contract event 0x4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5.
//
// Solidity: event Left(address indexed _validator, address indexed _coinbase)
func (_Validator *ValidatorFilterer) ParseLeft(log types.Log) (*ValidatorLeft, error) {
	event := new(ValidatorLeft)
	if err := _Validator.contract.UnpackLog(event, "Left", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the Validator contract.
type ValidatorLockedIterator struct {
	Event *ValidatorLocked // Event containing the contract specifics and raw log

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
func (it *ValidatorLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorLocked)
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
		it.Event = new(ValidatorLocked)
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
func (it *ValidatorLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorLocked represents a Locked event raised by the Validator contract.
type ValidatorLocked struct {
	Validator  common.Address
	UnlockedAt *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed _validator, uint256 indexed _unlockedAt)
func (_Validator *ValidatorFilterer) FilterLocked(opts *bind.FilterOpts, _validator []common.Address, _unlockedAt []*big.Int) (*ValidatorLockedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}
	var _unlockedAtRule []interface{}
	for _, _unlockedAtItem := range _unlockedAt {
		_unlockedAtRule = append(_unlockedAtRule, _unlockedAtItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "Locked", _validatorRule, _unlockedAtRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorLockedIterator{contract: _Validator.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed _validator, uint256 indexed _unlockedAt)
func (_Validator *ValidatorFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *ValidatorLocked, _validator []common.Address, _unlockedAt []*big.Int) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}
	var _unlockedAtRule []interface{}
	for _, _unlockedAtItem := range _unlockedAt {
		_unlockedAtRule = append(_unlockedAtRule, _unlockedAtItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "Locked", _validatorRule, _unlockedAtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorLocked)
				if err := _Validator.contract.UnpackLog(event, "Locked", log); err != nil {
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

// ParseLocked is a log parse operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed _validator, uint256 indexed _unlockedAt)
func (_Validator *ValidatorFilterer) ParseLocked(log types.Log) (*ValidatorLocked, error) {
	event := new(ValidatorLocked)
	if err := _Validator.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorMetadataUpdatedIterator is returned from FilterMetadataUpdated and is used to iterate over the raw logs and unpacked data for MetadataUpdated events raised by the Validator contract.
type ValidatorMetadataUpdatedIterator struct {
	Event *ValidatorMetadataUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorMetadataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMetadataUpdated)
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
		it.Event = new(ValidatorMetadataUpdated)
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
func (it *ValidatorMetadataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMetadataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMetadataUpdated represents a MetadataUpdated event raised by the Validator contract.
type ValidatorMetadataUpdated struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdated is a free log retrieval operation binding the contract event 0xe08224dc11618a1e23016c28d3fb80e30630f4df34d9f95890fe9ce89a85d07e.
//
// Solidity: event MetadataUpdated(address indexed _validator)
func (_Validator *ValidatorFilterer) FilterMetadataUpdated(opts *bind.FilterOpts, _validator []common.Address) (*ValidatorMetadataUpdatedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "MetadataUpdated", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorMetadataUpdatedIterator{contract: _Validator.contract, event: "MetadataUpdated", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdated is a free log subscription operation binding the contract event 0xe08224dc11618a1e23016c28d3fb80e30630f4df34d9f95890fe9ce89a85d07e.
//
// Solidity: event MetadataUpdated(address indexed _validator)
func (_Validator *ValidatorFilterer) WatchMetadataUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorMetadataUpdated, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "MetadataUpdated", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMetadataUpdated)
				if err := _Validator.contract.UnpackLog(event, "MetadataUpdated", log); err != nil {
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

// ParseMetadataUpdated is a log parse operation binding the contract event 0xe08224dc11618a1e23016c28d3fb80e30630f4df34d9f95890fe9ce89a85d07e.
//
// Solidity: event MetadataUpdated(address indexed _validator)
func (_Validator *ValidatorFilterer) ParseMetadataUpdated(log types.Log) (*ValidatorMetadataUpdated, error) {
	event := new(ValidatorMetadataUpdated)
	if err := _Validator.contract.UnpackLog(event, "MetadataUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorPenalizedIterator is returned from FilterPenalized and is used to iterate over the raw logs and unpacked data for Penalized events raised by the Validator contract.
type ValidatorPenalizedIterator struct {
	Event *ValidatorPenalized // Event containing the contract specifics and raw log

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
func (it *ValidatorPenalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorPenalized)
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
		it.Event = new(ValidatorPenalized)
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
func (it *ValidatorPenalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorPenalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorPenalized represents a Penalized event raised by the Validator contract.
type ValidatorPenalized struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPenalized is a free log retrieval operation binding the contract event 0xce6324a0bfd5e264b9b4039aeeab1620f460c269fa080cbd889681a37aaadc65.
//
// Solidity: event Penalized(address indexed _validator, uint256 _amount)
func (_Validator *ValidatorFilterer) FilterPenalized(opts *bind.FilterOpts, _validator []common.Address) (*ValidatorPenalizedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "Penalized", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorPenalizedIterator{contract: _Validator.contract, event: "Penalized", logs: logs, sub: sub}, nil
}

// WatchPenalized is a free log subscription operation binding the contract event 0xce6324a0bfd5e264b9b4039aeeab1620f460c269fa080cbd889681a37aaadc65.
//
// Solidity: event Penalized(address indexed _validator, uint256 _amount)
func (_Validator *ValidatorFilterer) WatchPenalized(opts *bind.WatchOpts, sink chan<- *ValidatorPenalized, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "Penalized", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorPenalized)
				if err := _Validator.contract.UnpackLog(event, "Penalized", log); err != nil {
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

// ParsePenalized is a log parse operation binding the contract event 0xce6324a0bfd5e264b9b4039aeeab1620f460c269fa080cbd889681a37aaadc65.
//
// Solidity: event Penalized(address indexed _validator, uint256 _amount)
func (_Validator *ValidatorFilterer) ParsePenalized(log types.Log) (*ValidatorPenalized, error) {
	event := new(ValidatorPenalized)
	if err := _Validator.contract.UnpackLog(event, "Penalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorPointUpdatedIterator is returned from FilterPointUpdated and is used to iterate over the raw logs and unpacked data for PointUpdated events raised by the Validator contract.
type ValidatorPointUpdatedIterator struct {
	Event *ValidatorPointUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorPointUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorPointUpdated)
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
		it.Event = new(ValidatorPointUpdated)
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
func (it *ValidatorPointUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorPointUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorPointUpdated represents a PointUpdated event raised by the Validator contract.
type ValidatorPointUpdated struct {
	Validator common.Address
	Point     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPointUpdated is a free log retrieval operation binding the contract event 0x88a501824285cc7436ca02dbfc2953e94f577c28429dec36bac9c7782699a9d8.
//
// Solidity: event PointUpdated(address indexed _validator, uint256 _point)
func (_Validator *ValidatorFilterer) FilterPointUpdated(opts *bind.FilterOpts, _validator []common.Address) (*ValidatorPointUpdatedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "PointUpdated", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorPointUpdatedIterator{contract: _Validator.contract, event: "PointUpdated", logs: logs, sub: sub}, nil
}

// WatchPointUpdated is a free log subscription operation binding the contract event 0x88a501824285cc7436ca02dbfc2953e94f577c28429dec36bac9c7782699a9d8.
//
// Solidity: event PointUpdated(address indexed _validator, uint256 _point)
func (_Validator *ValidatorFilterer) WatchPointUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorPointUpdated, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "PointUpdated", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorPointUpdated)
				if err := _Validator.contract.UnpackLog(event, "PointUpdated", log); err != nil {
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

// ParsePointUpdated is a log parse operation binding the contract event 0x88a501824285cc7436ca02dbfc2953e94f577c28429dec36bac9c7782699a9d8.
//
// Solidity: event PointUpdated(address indexed _validator, uint256 _point)
func (_Validator *ValidatorFilterer) ParsePointUpdated(log types.Log) (*ValidatorPointUpdated, error) {
	event := new(ValidatorPointUpdated)
	if err := _Validator.contract.UnpackLog(event, "PointUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorProofSubmittedIterator is returned from FilterProofSubmitted and is used to iterate over the raw logs and unpacked data for ProofSubmitted events raised by the Validator contract.
type ValidatorProofSubmittedIterator struct {
	Event *ValidatorProofSubmitted // Event containing the contract specifics and raw log

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
func (it *ValidatorProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorProofSubmitted)
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
		it.Event = new(ValidatorProofSubmitted)
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
func (it *ValidatorProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorProofSubmitted represents a ProofSubmitted event raised by the Validator contract.
type ValidatorProofSubmitted struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProofSubmitted is a free log retrieval operation binding the contract event 0x87dc06121a3dc1ab46683cb20422ce8b605fd9a4369ad5128f0dc2a49283e84c.
//
// Solidity: event ProofSubmitted(address indexed _validator)
func (_Validator *ValidatorFilterer) FilterProofSubmitted(opts *bind.FilterOpts, _validator []common.Address) (*ValidatorProofSubmittedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ProofSubmitted", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorProofSubmittedIterator{contract: _Validator.contract, event: "ProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchProofSubmitted is a free log subscription operation binding the contract event 0x87dc06121a3dc1ab46683cb20422ce8b605fd9a4369ad5128f0dc2a49283e84c.
//
// Solidity: event ProofSubmitted(address indexed _validator)
func (_Validator *ValidatorFilterer) WatchProofSubmitted(opts *bind.WatchOpts, sink chan<- *ValidatorProofSubmitted, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ProofSubmitted", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorProofSubmitted)
				if err := _Validator.contract.UnpackLog(event, "ProofSubmitted", log); err != nil {
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

// ParseProofSubmitted is a log parse operation binding the contract event 0x87dc06121a3dc1ab46683cb20422ce8b605fd9a4369ad5128f0dc2a49283e84c.
//
// Solidity: event ProofSubmitted(address indexed _validator)
func (_Validator *ValidatorFilterer) ParseProofSubmitted(log types.Log) (*ValidatorProofSubmitted, error) {
	event := new(ValidatorProofSubmitted)
	if err := _Validator.contract.UnpackLog(event, "ProofSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the Validator contract.
type ValidatorSlashedIterator struct {
	Event *ValidatorSlashed // Event containing the contract specifics and raw log

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
func (it *ValidatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSlashed)
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
		it.Event = new(ValidatorSlashed)
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
func (it *ValidatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSlashed represents a Slashed event raised by the Validator contract.
type ValidatorSlashed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x975ad74f3e1160b20a620ec57c636ffa56c6b6679a0e2fb5689b4f95e65a6946.
//
// Solidity: event Slashed(address indexed _validator)
func (_Validator *ValidatorFilterer) FilterSlashed(opts *bind.FilterOpts, _validator []common.Address) (*ValidatorSlashedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "Slashed", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorSlashedIterator{contract: _Validator.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x975ad74f3e1160b20a620ec57c636ffa56c6b6679a0e2fb5689b4f95e65a6946.
//
// Solidity: event Slashed(address indexed _validator)
func (_Validator *ValidatorFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *ValidatorSlashed, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "Slashed", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSlashed)
				if err := _Validator.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x975ad74f3e1160b20a620ec57c636ffa56c6b6679a0e2fb5689b4f95e65a6946.
//
// Solidity: event Slashed(address indexed _validator)
func (_Validator *ValidatorFilterer) ParseSlashed(log types.Log) (*ValidatorSlashed, error) {
	event := new(ValidatorSlashed)
	if err := _Validator.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Validator contract.
type ValidatorStakedIterator struct {
	Event *ValidatorStaked // Event containing the contract specifics and raw log

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
func (it *ValidatorStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorStaked)
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
		it.Event = new(ValidatorStaked)
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
func (it *ValidatorStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorStaked represents a Staked event raised by the Validator contract.
type ValidatorStaked struct {
	Validator common.Address
	Amount    *big.Int
	Stake     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed _validator, uint256 _amount, uint256 _stake)
func (_Validator *ValidatorFilterer) FilterStaked(opts *bind.FilterOpts, _validator []common.Address) (*ValidatorStakedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "Staked", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorStakedIterator{contract: _Validator.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed _validator, uint256 _amount, uint256 _stake)
func (_Validator *ValidatorFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *ValidatorStaked, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "Staked", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorStaked)
				if err := _Validator.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed _validator, uint256 _amount, uint256 _stake)
func (_Validator *ValidatorFilterer) ParseStaked(log types.Log) (*ValidatorStaked, error) {
	event := new(ValidatorStaked)
	if err := _Validator.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorStartTimerIterator is returned from FilterStartTimer and is used to iterate over the raw logs and unpacked data for StartTimer events raised by the Validator contract.
type ValidatorStartTimerIterator struct {
	Event *ValidatorStartTimer // Event containing the contract specifics and raw log

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
func (it *ValidatorStartTimerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorStartTimer)
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
		it.Event = new(ValidatorStartTimer)
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
func (it *ValidatorStartTimerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorStartTimerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorStartTimer represents a StartTimer event raised by the Validator contract.
type ValidatorStartTimer struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStartTimer is a free log retrieval operation binding the contract event 0x3d364b4503e0f292c41ea2af3eea2464e1e89d01f1a7d6519063a7342ab86522.
//
// Solidity: event StartTimer(address indexed _validator)
func (_Validator *ValidatorFilterer) FilterStartTimer(opts *bind.FilterOpts, _validator []common.Address) (*ValidatorStartTimerIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "StartTimer", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorStartTimerIterator{contract: _Validator.contract, event: "StartTimer", logs: logs, sub: sub}, nil
}

// WatchStartTimer is a free log subscription operation binding the contract event 0x3d364b4503e0f292c41ea2af3eea2464e1e89d01f1a7d6519063a7342ab86522.
//
// Solidity: event StartTimer(address indexed _validator)
func (_Validator *ValidatorFilterer) WatchStartTimer(opts *bind.WatchOpts, sink chan<- *ValidatorStartTimer, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "StartTimer", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorStartTimer)
				if err := _Validator.contract.UnpackLog(event, "StartTimer", log); err != nil {
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

// ParseStartTimer is a log parse operation binding the contract event 0x3d364b4503e0f292c41ea2af3eea2464e1e89d01f1a7d6519063a7342ab86522.
//
// Solidity: event StartTimer(address indexed _validator)
func (_Validator *ValidatorFilterer) ParseStartTimer(log types.Log) (*ValidatorStartTimer, error) {
	event := new(ValidatorStartTimer)
	if err := _Validator.contract.UnpackLog(event, "StartTimer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Validator contract.
type ValidatorTransferIterator struct {
	Event *ValidatorTransfer // Event containing the contract specifics and raw log

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
func (it *ValidatorTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorTransfer)
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
		it.Event = new(ValidatorTransfer)
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
func (it *ValidatorTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorTransfer represents a Transfer event raised by the Validator contract.
type ValidatorTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Validator *ValidatorFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ValidatorTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorTransferIterator{contract: _Validator.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Validator *ValidatorFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ValidatorTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorTransfer)
				if err := _Validator.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Validator *ValidatorFilterer) ParseTransfer(log types.Log) (*ValidatorTransfer, error) {
	event := new(ValidatorTransfer)
	if err := _Validator.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Validator contract.
type ValidatorUnstakedIterator struct {
	Event *ValidatorUnstaked // Event containing the contract specifics and raw log

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
func (it *ValidatorUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorUnstaked)
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
		it.Event = new(ValidatorUnstaked)
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
func (it *ValidatorUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorUnstaked represents a Unstaked event raised by the Validator contract.
type ValidatorUnstaked struct {
	Validator common.Address
	Amount    *big.Int
	Stake     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed _validator, uint256 _amount, uint256 _stake)
func (_Validator *ValidatorFilterer) FilterUnstaked(opts *bind.FilterOpts, _validator []common.Address) (*ValidatorUnstakedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "Unstaked", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorUnstakedIterator{contract: _Validator.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed _validator, uint256 _amount, uint256 _stake)
func (_Validator *ValidatorFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *ValidatorUnstaked, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "Unstaked", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorUnstaked)
				if err := _Validator.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed _validator, uint256 _amount, uint256 _stake)
func (_Validator *ValidatorFilterer) ParseUnstaked(log types.Log) (*ValidatorUnstaked, error) {
	event := new(ValidatorUnstaked)
	if err := _Validator.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorMetadataABI is the input ABI used to generate the binding from.
const ValidatorMetadataABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMinThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miningKey\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voterMiningKey\",\"type\":\"address\"}],\"name\":\"isValidatorAlreadyVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"MetadataUpdated\",\"type\":\"event\"}]"

// ValidatorMetadataFuncSigs maps the 4-byte function signature to its string representation.
var ValidatorMetadataFuncSigs = map[string]string{
	"e6bbe9dd": "getMinThreshold()",
	"557ed1ba": "getTime()",
	"f73294b8": "isValidatorAlreadyVoted(address,address)",
}

// ValidatorMetadataBin is the compiled bytecode used for deploying new contracts.
var ValidatorMetadataBin = "0x608060405234801561001057600080fd5b5060e48061001f6000396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c8063557ed1ba146041578063e6bbe9dd146059578063f73294b814605f575b600080fd5b6047609e565b60408051918252519081900360200190f35b604760a2565b608a60048036036040811015607357600080fd5b506001600160a01b038135811691602001351660a7565b604080519115158252519081900360200190f35b4290565b600890565b60019291505056fea265627a7a72315820a3b4e49878a592b14478247f7793c831060b104a021c0288c76ff11ecc346e6e64736f6c634300050b0032"

// DeployValidatorMetadata deploys a new Ethereum contract, binding an instance of ValidatorMetadata to it.
func DeployValidatorMetadata(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidatorMetadata, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorMetadataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorMetadataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorMetadata{ValidatorMetadataCaller: ValidatorMetadataCaller{contract: contract}, ValidatorMetadataTransactor: ValidatorMetadataTransactor{contract: contract}, ValidatorMetadataFilterer: ValidatorMetadataFilterer{contract: contract}}, nil
}

// ValidatorMetadata is an auto generated Go binding around an Ethereum contract.
type ValidatorMetadata struct {
	ValidatorMetadataCaller     // Read-only binding to the contract
	ValidatorMetadataTransactor // Write-only binding to the contract
	ValidatorMetadataFilterer   // Log filterer for contract events
}

// ValidatorMetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorMetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorMetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorMetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorMetadataSession struct {
	Contract     *ValidatorMetadata // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ValidatorMetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorMetadataCallerSession struct {
	Contract *ValidatorMetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ValidatorMetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorMetadataTransactorSession struct {
	Contract     *ValidatorMetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ValidatorMetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorMetadataRaw struct {
	Contract *ValidatorMetadata // Generic contract binding to access the raw methods on
}

// ValidatorMetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorMetadataCallerRaw struct {
	Contract *ValidatorMetadataCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorMetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorMetadataTransactorRaw struct {
	Contract *ValidatorMetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorMetadata creates a new instance of ValidatorMetadata, bound to a specific deployed contract.
func NewValidatorMetadata(address common.Address, backend bind.ContractBackend) (*ValidatorMetadata, error) {
	contract, err := bindValidatorMetadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorMetadata{ValidatorMetadataCaller: ValidatorMetadataCaller{contract: contract}, ValidatorMetadataTransactor: ValidatorMetadataTransactor{contract: contract}, ValidatorMetadataFilterer: ValidatorMetadataFilterer{contract: contract}}, nil
}

// NewValidatorMetadataCaller creates a new read-only instance of ValidatorMetadata, bound to a specific deployed contract.
func NewValidatorMetadataCaller(address common.Address, caller bind.ContractCaller) (*ValidatorMetadataCaller, error) {
	contract, err := bindValidatorMetadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorMetadataCaller{contract: contract}, nil
}

// NewValidatorMetadataTransactor creates a new write-only instance of ValidatorMetadata, bound to a specific deployed contract.
func NewValidatorMetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorMetadataTransactor, error) {
	contract, err := bindValidatorMetadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorMetadataTransactor{contract: contract}, nil
}

// NewValidatorMetadataFilterer creates a new log filterer instance of ValidatorMetadata, bound to a specific deployed contract.
func NewValidatorMetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorMetadataFilterer, error) {
	contract, err := bindValidatorMetadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorMetadataFilterer{contract: contract}, nil
}

// bindValidatorMetadata binds a generic wrapper to an already deployed contract.
func bindValidatorMetadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorMetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorMetadata *ValidatorMetadataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorMetadata.Contract.ValidatorMetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorMetadata *ValidatorMetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMetadata.Contract.ValidatorMetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorMetadata *ValidatorMetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorMetadata.Contract.ValidatorMetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorMetadata *ValidatorMetadataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorMetadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorMetadata *ValidatorMetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMetadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorMetadata *ValidatorMetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorMetadata.Contract.contract.Transact(opts, method, params...)
}

// GetMinThreshold is a free data retrieval call binding the contract method 0xe6bbe9dd.
//
// Solidity: function getMinThreshold() constant returns(uint256)
func (_ValidatorMetadata *ValidatorMetadataCaller) GetMinThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorMetadata.contract.Call(opts, out, "getMinThreshold")
	return *ret0, err
}

// GetMinThreshold is a free data retrieval call binding the contract method 0xe6bbe9dd.
//
// Solidity: function getMinThreshold() constant returns(uint256)
func (_ValidatorMetadata *ValidatorMetadataSession) GetMinThreshold() (*big.Int, error) {
	return _ValidatorMetadata.Contract.GetMinThreshold(&_ValidatorMetadata.CallOpts)
}

// GetMinThreshold is a free data retrieval call binding the contract method 0xe6bbe9dd.
//
// Solidity: function getMinThreshold() constant returns(uint256)
func (_ValidatorMetadata *ValidatorMetadataCallerSession) GetMinThreshold() (*big.Int, error) {
	return _ValidatorMetadata.Contract.GetMinThreshold(&_ValidatorMetadata.CallOpts)
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() constant returns(uint256)
func (_ValidatorMetadata *ValidatorMetadataCaller) GetTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorMetadata.contract.Call(opts, out, "getTime")
	return *ret0, err
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() constant returns(uint256)
func (_ValidatorMetadata *ValidatorMetadataSession) GetTime() (*big.Int, error) {
	return _ValidatorMetadata.Contract.GetTime(&_ValidatorMetadata.CallOpts)
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() constant returns(uint256)
func (_ValidatorMetadata *ValidatorMetadataCallerSession) GetTime() (*big.Int, error) {
	return _ValidatorMetadata.Contract.GetTime(&_ValidatorMetadata.CallOpts)
}

// IsValidatorAlreadyVoted is a free data retrieval call binding the contract method 0xf73294b8.
//
// Solidity: function isValidatorAlreadyVoted(address _miningKey, address _voterMiningKey) constant returns(bool)
func (_ValidatorMetadata *ValidatorMetadataCaller) IsValidatorAlreadyVoted(opts *bind.CallOpts, _miningKey common.Address, _voterMiningKey common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidatorMetadata.contract.Call(opts, out, "isValidatorAlreadyVoted", _miningKey, _voterMiningKey)
	return *ret0, err
}

// IsValidatorAlreadyVoted is a free data retrieval call binding the contract method 0xf73294b8.
//
// Solidity: function isValidatorAlreadyVoted(address _miningKey, address _voterMiningKey) constant returns(bool)
func (_ValidatorMetadata *ValidatorMetadataSession) IsValidatorAlreadyVoted(_miningKey common.Address, _voterMiningKey common.Address) (bool, error) {
	return _ValidatorMetadata.Contract.IsValidatorAlreadyVoted(&_ValidatorMetadata.CallOpts, _miningKey, _voterMiningKey)
}

// IsValidatorAlreadyVoted is a free data retrieval call binding the contract method 0xf73294b8.
//
// Solidity: function isValidatorAlreadyVoted(address _miningKey, address _voterMiningKey) constant returns(bool)
func (_ValidatorMetadata *ValidatorMetadataCallerSession) IsValidatorAlreadyVoted(_miningKey common.Address, _voterMiningKey common.Address) (bool, error) {
	return _ValidatorMetadata.Contract.IsValidatorAlreadyVoted(&_ValidatorMetadata.CallOpts, _miningKey, _voterMiningKey)
}

// ValidatorMetadataMetadataUpdatedIterator is returned from FilterMetadataUpdated and is used to iterate over the raw logs and unpacked data for MetadataUpdated events raised by the ValidatorMetadata contract.
type ValidatorMetadataMetadataUpdatedIterator struct {
	Event *ValidatorMetadataMetadataUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorMetadataMetadataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMetadataMetadataUpdated)
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
		it.Event = new(ValidatorMetadataMetadataUpdated)
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
func (it *ValidatorMetadataMetadataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMetadataMetadataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMetadataMetadataUpdated represents a MetadataUpdated event raised by the ValidatorMetadata contract.
type ValidatorMetadataMetadataUpdated struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdated is a free log retrieval operation binding the contract event 0xe08224dc11618a1e23016c28d3fb80e30630f4df34d9f95890fe9ce89a85d07e.
//
// Solidity: event MetadataUpdated(address indexed _validator)
func (_ValidatorMetadata *ValidatorMetadataFilterer) FilterMetadataUpdated(opts *bind.FilterOpts, _validator []common.Address) (*ValidatorMetadataMetadataUpdatedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _ValidatorMetadata.contract.FilterLogs(opts, "MetadataUpdated", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorMetadataMetadataUpdatedIterator{contract: _ValidatorMetadata.contract, event: "MetadataUpdated", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdated is a free log subscription operation binding the contract event 0xe08224dc11618a1e23016c28d3fb80e30630f4df34d9f95890fe9ce89a85d07e.
//
// Solidity: event MetadataUpdated(address indexed _validator)
func (_ValidatorMetadata *ValidatorMetadataFilterer) WatchMetadataUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorMetadataMetadataUpdated, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _ValidatorMetadata.contract.WatchLogs(opts, "MetadataUpdated", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMetadataMetadataUpdated)
				if err := _ValidatorMetadata.contract.UnpackLog(event, "MetadataUpdated", log); err != nil {
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

// ParseMetadataUpdated is a log parse operation binding the contract event 0xe08224dc11618a1e23016c28d3fb80e30630f4df34d9f95890fe9ce89a85d07e.
//
// Solidity: event MetadataUpdated(address indexed _validator)
func (_ValidatorMetadata *ValidatorMetadataFilterer) ParseMetadataUpdated(log types.Log) (*ValidatorMetadataMetadataUpdated, error) {
	event := new(ValidatorMetadataMetadataUpdated)
	if err := _ValidatorMetadata.contract.UnpackLog(event, "MetadataUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
