// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ChannelWalletABI is the input ABI used to generate the binding from.
const ChannelWalletABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"currentTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ts\",\"type\":\"uint64\"}],\"name\":\"validateTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"guarantor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_dest\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_ts\",\"type\":\"uint256\"}],\"name\":\"validateTransaction\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_dest\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_ts\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ts\",\"type\":\"uint256\"}],\"name\":\"validateTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_guarantor\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"guarantor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"code\",\"type\":\"int256\"}],\"name\":\"ExecStatus\",\"type\":\"event\"}]"

// ChannelWalletBin is the compiled bytecode used for deploying new contracts.
const ChannelWalletBin = `0x`

// DeployChannelWallet deploys a new Ethereum contract, binding an instance of ChannelWallet to it.
func DeployChannelWallet(auth *bind.TransactOpts, backend bind.ContractBackend, _guarantor common.Address) (common.Address, *types.Transaction, *ChannelWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChannelWalletBin), backend, _guarantor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChannelWallet{ChannelWalletCaller: ChannelWalletCaller{contract: contract}, ChannelWalletTransactor: ChannelWalletTransactor{contract: contract}, ChannelWalletFilterer: ChannelWalletFilterer{contract: contract}}, nil
}

// ChannelWallet is an auto generated Go binding around an Ethereum contract.
type ChannelWallet struct {
	ChannelWalletCaller     // Read-only binding to the contract
	ChannelWalletTransactor // Write-only binding to the contract
	ChannelWalletFilterer   // Log filterer for contract events
}

// ChannelWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChannelWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChannelWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChannelWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChannelWalletSession struct {
	Contract     *ChannelWallet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChannelWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChannelWalletCallerSession struct {
	Contract *ChannelWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ChannelWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChannelWalletTransactorSession struct {
	Contract     *ChannelWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ChannelWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChannelWalletRaw struct {
	Contract *ChannelWallet // Generic contract binding to access the raw methods on
}

// ChannelWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChannelWalletCallerRaw struct {
	Contract *ChannelWalletCaller // Generic read-only contract binding to access the raw methods on
}

// ChannelWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChannelWalletTransactorRaw struct {
	Contract *ChannelWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChannelWallet creates a new instance of ChannelWallet, bound to a specific deployed contract.
func NewChannelWallet(address common.Address, backend bind.ContractBackend) (*ChannelWallet, error) {
	contract, err := bindChannelWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChannelWallet{ChannelWalletCaller: ChannelWalletCaller{contract: contract}, ChannelWalletTransactor: ChannelWalletTransactor{contract: contract}, ChannelWalletFilterer: ChannelWalletFilterer{contract: contract}}, nil
}

// NewChannelWalletCaller creates a new read-only instance of ChannelWallet, bound to a specific deployed contract.
func NewChannelWalletCaller(address common.Address, caller bind.ContractCaller) (*ChannelWalletCaller, error) {
	contract, err := bindChannelWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelWalletCaller{contract: contract}, nil
}

// NewChannelWalletTransactor creates a new write-only instance of ChannelWallet, bound to a specific deployed contract.
func NewChannelWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*ChannelWalletTransactor, error) {
	contract, err := bindChannelWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelWalletTransactor{contract: contract}, nil
}

// NewChannelWalletFilterer creates a new log filterer instance of ChannelWallet, bound to a specific deployed contract.
func NewChannelWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*ChannelWalletFilterer, error) {
	contract, err := bindChannelWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChannelWalletFilterer{contract: contract}, nil
}

// bindChannelWallet binds a generic wrapper to an already deployed contract.
func bindChannelWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChannelWallet *ChannelWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChannelWallet.Contract.ChannelWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChannelWallet *ChannelWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelWallet.Contract.ChannelWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChannelWallet *ChannelWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChannelWallet.Contract.ChannelWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChannelWallet *ChannelWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChannelWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChannelWallet *ChannelWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChannelWallet *ChannelWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChannelWallet.Contract.contract.Transact(opts, method, params...)
}

// CurrentTimestamp is a free data retrieval call binding the contract method 0x1e2ff94f.
//
// Solidity: function currentTimestamp() constant returns(uint256)
func (_ChannelWallet *ChannelWalletCaller) CurrentTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChannelWallet.contract.Call(opts, out, "currentTimestamp")
	return *ret0, err
}

// CurrentTimestamp is a free data retrieval call binding the contract method 0x1e2ff94f.
//
// Solidity: function currentTimestamp() constant returns(uint256)
func (_ChannelWallet *ChannelWalletSession) CurrentTimestamp() (*big.Int, error) {
	return _ChannelWallet.Contract.CurrentTimestamp(&_ChannelWallet.CallOpts)
}

// CurrentTimestamp is a free data retrieval call binding the contract method 0x1e2ff94f.
//
// Solidity: function currentTimestamp() constant returns(uint256)
func (_ChannelWallet *ChannelWalletCallerSession) CurrentTimestamp() (*big.Int, error) {
	return _ChannelWallet.Contract.CurrentTimestamp(&_ChannelWallet.CallOpts)
}

// Guarantor is a free data retrieval call binding the contract method 0x71cf516b.
//
// Solidity: function guarantor() constant returns(address)
func (_ChannelWallet *ChannelWalletCaller) Guarantor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChannelWallet.contract.Call(opts, out, "guarantor")
	return *ret0, err
}

// Guarantor is a free data retrieval call binding the contract method 0x71cf516b.
//
// Solidity: function guarantor() constant returns(address)
func (_ChannelWallet *ChannelWalletSession) Guarantor() (common.Address, error) {
	return _ChannelWallet.Contract.Guarantor(&_ChannelWallet.CallOpts)
}

// Guarantor is a free data retrieval call binding the contract method 0x71cf516b.
//
// Solidity: function guarantor() constant returns(address)
func (_ChannelWallet *ChannelWalletCallerSession) Guarantor() (common.Address, error) {
	return _ChannelWallet.Contract.Guarantor(&_ChannelWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChannelWallet *ChannelWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChannelWallet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChannelWallet *ChannelWalletSession) Owner() (common.Address, error) {
	return _ChannelWallet.Contract.Owner(&_ChannelWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChannelWallet *ChannelWalletCallerSession) Owner() (common.Address, error) {
	return _ChannelWallet.Contract.Owner(&_ChannelWallet.CallOpts)
}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 _ts) constant returns(bool, uint256)
func (_ChannelWallet *ChannelWalletCaller) ValidateTimestamp(opts *bind.CallOpts, _ts *big.Int) (bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ChannelWallet.contract.Call(opts, out, "validateTimestamp", _ts)
	return *ret0, *ret1, err
}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 _ts) constant returns(bool, uint256)
func (_ChannelWallet *ChannelWalletSession) ValidateTimestamp(_ts *big.Int) (bool, *big.Int, error) {
	return _ChannelWallet.Contract.ValidateTimestamp(&_ChannelWallet.CallOpts, _ts)
}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 _ts) constant returns(bool, uint256)
func (_ChannelWallet *ChannelWalletCallerSession) ValidateTimestamp(_ts *big.Int) (bool, *big.Int, error) {
	return _ChannelWallet.Contract.ValidateTimestamp(&_ChannelWallet.CallOpts, _ts)
}

// ValidateTransaction is a free data retrieval call binding the contract method 0x763d5ee6.
//
// Solidity: function validateTransaction(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) constant returns(bool)
func (_ChannelWallet *ChannelWalletCaller) ValidateTransaction(opts *bind.CallOpts, _v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChannelWallet.contract.Call(opts, out, "validateTransaction", _v, _r, _s, _dest, _value, _ts)
	return *ret0, err
}

// ValidateTransaction is a free data retrieval call binding the contract method 0x763d5ee6.
//
// Solidity: function validateTransaction(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) constant returns(bool)
func (_ChannelWallet *ChannelWalletSession) ValidateTransaction(_v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (bool, error) {
	return _ChannelWallet.Contract.ValidateTransaction(&_ChannelWallet.CallOpts, _v, _r, _s, _dest, _value, _ts)
}

// ValidateTransaction is a free data retrieval call binding the contract method 0x763d5ee6.
//
// Solidity: function validateTransaction(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) constant returns(bool)
func (_ChannelWallet *ChannelWalletCallerSession) ValidateTransaction(_v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (bool, error) {
	return _ChannelWallet.Contract.ValidateTransaction(&_ChannelWallet.CallOpts, _v, _r, _s, _dest, _value, _ts)
}

// Execute is a paid mutator transaction binding the contract method 0x94489c42.
//
// Solidity: function execute(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) returns(bool)
func (_ChannelWallet *ChannelWalletTransactor) Execute(opts *bind.TransactOpts, _v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (*types.Transaction, error) {
	return _ChannelWallet.contract.Transact(opts, "execute", _v, _r, _s, _dest, _value, _ts)
}

// Execute is a paid mutator transaction binding the contract method 0x94489c42.
//
// Solidity: function execute(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) returns(bool)
func (_ChannelWallet *ChannelWalletSession) Execute(_v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (*types.Transaction, error) {
	return _ChannelWallet.Contract.Execute(&_ChannelWallet.TransactOpts, _v, _r, _s, _dest, _value, _ts)
}

// Execute is a paid mutator transaction binding the contract method 0x94489c42.
//
// Solidity: function execute(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) returns(bool)
func (_ChannelWallet *ChannelWalletTransactorSession) Execute(_v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (*types.Transaction, error) {
	return _ChannelWallet.Contract.Execute(&_ChannelWallet.TransactOpts, _v, _r, _s, _dest, _value, _ts)
}

// ChannelWalletExecStatusIterator is returned from FilterExecStatus and is used to iterate over the raw logs and unpacked data for ExecStatus events raised by the ChannelWallet contract.
type ChannelWalletExecStatusIterator struct {
	Event *ChannelWalletExecStatus // Event containing the contract specifics and raw log

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
func (it *ChannelWalletExecStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelWalletExecStatus)
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
		it.Event = new(ChannelWalletExecStatus)
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
func (it *ChannelWalletExecStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelWalletExecStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelWalletExecStatus represents a ExecStatus event raised by the ChannelWallet contract.
type ChannelWalletExecStatus struct {
	Guarantor common.Address
	Code      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExecStatus is a free log retrieval operation binding the contract event 0x583d8312ef7016406c7ea8ba9796b9e55ac1fdc22455754cbc93869509faefad.
//
// Solidity: event ExecStatus(address guarantor, int256 code)
func (_ChannelWallet *ChannelWalletFilterer) FilterExecStatus(opts *bind.FilterOpts) (*ChannelWalletExecStatusIterator, error) {

	logs, sub, err := _ChannelWallet.contract.FilterLogs(opts, "ExecStatus")
	if err != nil {
		return nil, err
	}
	return &ChannelWalletExecStatusIterator{contract: _ChannelWallet.contract, event: "ExecStatus", logs: logs, sub: sub}, nil
}

// WatchExecStatus is a free log subscription operation binding the contract event 0x583d8312ef7016406c7ea8ba9796b9e55ac1fdc22455754cbc93869509faefad.
//
// Solidity: event ExecStatus(address guarantor, int256 code)
func (_ChannelWallet *ChannelWalletFilterer) WatchExecStatus(opts *bind.WatchOpts, sink chan<- *ChannelWalletExecStatus) (event.Subscription, error) {

	logs, sub, err := _ChannelWallet.contract.WatchLogs(opts, "ExecStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelWalletExecStatus)
				if err := _ChannelWallet.contract.UnpackLog(event, "ExecStatus", log); err != nil {
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

// TransactableABI is the input ABI used to generate the binding from.
const TransactableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_dest\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_ts\",\"type\":\"uint256\"}],\"name\":\"validateTransaction\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_dest\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_ts\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ts\",\"type\":\"uint256\"}],\"name\":\"validateTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TransactableBin is the compiled bytecode used for deploying new contracts.
const TransactableBin = `0x`

// DeployTransactable deploys a new Ethereum contract, binding an instance of Transactable to it.
func DeployTransactable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Transactable, error) {
	parsed, err := abi.JSON(strings.NewReader(TransactableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TransactableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Transactable{TransactableCaller: TransactableCaller{contract: contract}, TransactableTransactor: TransactableTransactor{contract: contract}, TransactableFilterer: TransactableFilterer{contract: contract}}, nil
}

// Transactable is an auto generated Go binding around an Ethereum contract.
type Transactable struct {
	TransactableCaller     // Read-only binding to the contract
	TransactableTransactor // Write-only binding to the contract
	TransactableFilterer   // Log filterer for contract events
}

// TransactableCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransactableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransactableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransactableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransactableSession struct {
	Contract     *Transactable     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransactableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransactableCallerSession struct {
	Contract *TransactableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TransactableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransactableTransactorSession struct {
	Contract     *TransactableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TransactableRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransactableRaw struct {
	Contract *Transactable // Generic contract binding to access the raw methods on
}

// TransactableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransactableCallerRaw struct {
	Contract *TransactableCaller // Generic read-only contract binding to access the raw methods on
}

// TransactableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransactableTransactorRaw struct {
	Contract *TransactableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransactable creates a new instance of Transactable, bound to a specific deployed contract.
func NewTransactable(address common.Address, backend bind.ContractBackend) (*Transactable, error) {
	contract, err := bindTransactable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transactable{TransactableCaller: TransactableCaller{contract: contract}, TransactableTransactor: TransactableTransactor{contract: contract}, TransactableFilterer: TransactableFilterer{contract: contract}}, nil
}

// NewTransactableCaller creates a new read-only instance of Transactable, bound to a specific deployed contract.
func NewTransactableCaller(address common.Address, caller bind.ContractCaller) (*TransactableCaller, error) {
	contract, err := bindTransactable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransactableCaller{contract: contract}, nil
}

// NewTransactableTransactor creates a new write-only instance of Transactable, bound to a specific deployed contract.
func NewTransactableTransactor(address common.Address, transactor bind.ContractTransactor) (*TransactableTransactor, error) {
	contract, err := bindTransactable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransactableTransactor{contract: contract}, nil
}

// NewTransactableFilterer creates a new log filterer instance of Transactable, bound to a specific deployed contract.
func NewTransactableFilterer(address common.Address, filterer bind.ContractFilterer) (*TransactableFilterer, error) {
	contract, err := bindTransactable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransactableFilterer{contract: contract}, nil
}

// bindTransactable binds a generic wrapper to an already deployed contract.
func bindTransactable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransactableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transactable *TransactableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transactable.Contract.TransactableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transactable *TransactableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transactable.Contract.TransactableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transactable *TransactableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transactable.Contract.TransactableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transactable *TransactableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transactable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transactable *TransactableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transactable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transactable *TransactableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transactable.Contract.contract.Transact(opts, method, params...)
}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 _ts) constant returns(bool, uint256)
func (_Transactable *TransactableCaller) ValidateTimestamp(opts *bind.CallOpts, _ts *big.Int) (bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Transactable.contract.Call(opts, out, "validateTimestamp", _ts)
	return *ret0, *ret1, err
}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 _ts) constant returns(bool, uint256)
func (_Transactable *TransactableSession) ValidateTimestamp(_ts *big.Int) (bool, *big.Int, error) {
	return _Transactable.Contract.ValidateTimestamp(&_Transactable.CallOpts, _ts)
}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 _ts) constant returns(bool, uint256)
func (_Transactable *TransactableCallerSession) ValidateTimestamp(_ts *big.Int) (bool, *big.Int, error) {
	return _Transactable.Contract.ValidateTimestamp(&_Transactable.CallOpts, _ts)
}

// ValidateTransaction is a free data retrieval call binding the contract method 0x763d5ee6.
//
// Solidity: function validateTransaction(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) constant returns(bool)
func (_Transactable *TransactableCaller) ValidateTransaction(opts *bind.CallOpts, _v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Transactable.contract.Call(opts, out, "validateTransaction", _v, _r, _s, _dest, _value, _ts)
	return *ret0, err
}

// ValidateTransaction is a free data retrieval call binding the contract method 0x763d5ee6.
//
// Solidity: function validateTransaction(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) constant returns(bool)
func (_Transactable *TransactableSession) ValidateTransaction(_v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (bool, error) {
	return _Transactable.Contract.ValidateTransaction(&_Transactable.CallOpts, _v, _r, _s, _dest, _value, _ts)
}

// ValidateTransaction is a free data retrieval call binding the contract method 0x763d5ee6.
//
// Solidity: function validateTransaction(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) constant returns(bool)
func (_Transactable *TransactableCallerSession) ValidateTransaction(_v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (bool, error) {
	return _Transactable.Contract.ValidateTransaction(&_Transactable.CallOpts, _v, _r, _s, _dest, _value, _ts)
}

// Execute is a paid mutator transaction binding the contract method 0x94489c42.
//
// Solidity: function execute(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) returns(bool)
func (_Transactable *TransactableTransactor) Execute(opts *bind.TransactOpts, _v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (*types.Transaction, error) {
	return _Transactable.contract.Transact(opts, "execute", _v, _r, _s, _dest, _value, _ts)
}

// Execute is a paid mutator transaction binding the contract method 0x94489c42.
//
// Solidity: function execute(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) returns(bool)
func (_Transactable *TransactableSession) Execute(_v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (*types.Transaction, error) {
	return _Transactable.Contract.Execute(&_Transactable.TransactOpts, _v, _r, _s, _dest, _value, _ts)
}

// Execute is a paid mutator transaction binding the contract method 0x94489c42.
//
// Solidity: function execute(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) returns(bool)
func (_Transactable *TransactableTransactorSession) Execute(_v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (*types.Transaction, error) {
	return _Transactable.Contract.Execute(&_Transactable.TransactOpts, _v, _r, _s, _dest, _value, _ts)
}
