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

// ArraysABI is the input ABI used to generate the binding from.
const ArraysABI = "[]"

// ArraysBin is the compiled bytecode used for deploying new contracts.
const ArraysBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582049b634e4ae8f8aa5b5378ac1c73afb29ece1440fb226df74194c2c6c18a9f5100029`

// DeployArrays deploys a new Ethereum contract, binding an instance of Arrays to it.
func DeployArrays(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Arrays, error) {
	parsed, err := abi.JSON(strings.NewReader(ArraysABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArraysBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Arrays{ArraysCaller: ArraysCaller{contract: contract}, ArraysTransactor: ArraysTransactor{contract: contract}, ArraysFilterer: ArraysFilterer{contract: contract}}, nil
}

// Arrays is an auto generated Go binding around an Ethereum contract.
type Arrays struct {
	ArraysCaller     // Read-only binding to the contract
	ArraysTransactor // Write-only binding to the contract
	ArraysFilterer   // Log filterer for contract events
}

// ArraysCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArraysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArraysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArraysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArraysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArraysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArraysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArraysSession struct {
	Contract     *Arrays           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArraysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArraysCallerSession struct {
	Contract *ArraysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArraysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArraysTransactorSession struct {
	Contract     *ArraysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArraysRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArraysRaw struct {
	Contract *Arrays // Generic contract binding to access the raw methods on
}

// ArraysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArraysCallerRaw struct {
	Contract *ArraysCaller // Generic read-only contract binding to access the raw methods on
}

// ArraysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArraysTransactorRaw struct {
	Contract *ArraysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArrays creates a new instance of Arrays, bound to a specific deployed contract.
func NewArrays(address common.Address, backend bind.ContractBackend) (*Arrays, error) {
	contract, err := bindArrays(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Arrays{ArraysCaller: ArraysCaller{contract: contract}, ArraysTransactor: ArraysTransactor{contract: contract}, ArraysFilterer: ArraysFilterer{contract: contract}}, nil
}

// NewArraysCaller creates a new read-only instance of Arrays, bound to a specific deployed contract.
func NewArraysCaller(address common.Address, caller bind.ContractCaller) (*ArraysCaller, error) {
	contract, err := bindArrays(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArraysCaller{contract: contract}, nil
}

// NewArraysTransactor creates a new write-only instance of Arrays, bound to a specific deployed contract.
func NewArraysTransactor(address common.Address, transactor bind.ContractTransactor) (*ArraysTransactor, error) {
	contract, err := bindArrays(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArraysTransactor{contract: contract}, nil
}

// NewArraysFilterer creates a new log filterer instance of Arrays, bound to a specific deployed contract.
func NewArraysFilterer(address common.Address, filterer bind.ContractFilterer) (*ArraysFilterer, error) {
	contract, err := bindArrays(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArraysFilterer{contract: contract}, nil
}

// bindArrays binds a generic wrapper to an already deployed contract.
func bindArrays(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArraysABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arrays *ArraysRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Arrays.Contract.ArraysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arrays *ArraysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arrays.Contract.ArraysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arrays *ArraysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arrays.Contract.ArraysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arrays *ArraysCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Arrays.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arrays *ArraysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arrays.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arrays *ArraysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arrays.Contract.contract.Transact(opts, method, params...)
}

// BatchTransactionABI is the input ABI used to generate the binding from.
const BatchTransactionABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_v\",\"type\":\"uint8[]\"},{\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"name\":\"_from\",\"type\":\"address[]\"},{\"name\":\"_dest\",\"type\":\"address[]\"},{\"name\":\"_value\",\"type\":\"uint256[]\"},{\"name\":\"_ts\",\"type\":\"uint256[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"TransactionStatus\",\"type\":\"event\"}]"

// BatchTransactionBin is the compiled bytecode used for deploying new contracts.
const BatchTransactionBin = `0x60806040526104a9806100136000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663e9861ab18114610042575b005b34801561004e57600080fd5b506040805160206004803580820135838102808601850190965280855261004095369593946024949385019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506101ed9650505050505050565b600080600088518a5114151561020257600080fd5b875189511461021057600080fd5b865188511461021e57600080fd5b855187511461022c57600080fd5b845186511461023a57600080fd5b835185511461024857600080fd5b600092505b895183101561047157868381518110151561026457fe5b9060200190602002015191508173ffffffffffffffffffffffffffffffffffffffff166394489c428b8581518110151561029a57fe5b906020019060200201518b868151811015156102b257fe5b906020019060200201518b878151811015156102ca57fe5b906020019060200201518a888151811015156102e257fe5b906020019060200201518a898151811015156102fa57fe5b906020019060200201518a8a81518110151561031257fe5b6020908102909101810151604080517c010000000000000000000000000000000000000000000000000000000063ffffffff8b1602815260ff90981660048901526024880196909652604487019490945273ffffffffffffffffffffffffffffffffffffffff9092166064860152608485015260a4840191909152905160c4808401938290030181600087803b1580156103ab57600080fd5b505af11580156103bf573d6000803e3d6000fd5b505050506040513d60208110156103d557600080fd5b505190507f41f5d8bdb5631ef260ce5fb88cf0a48c22d261827bebce333beb5963b3fe3edd81888581518110151561040957fe5b90602001906020020151888681518110151561042157fe5b602090810290910181015160408051941515855273ffffffffffffffffffffffffffffffffffffffff938416928501929092529190911682820152519081900360600190a160019092019161024d565b505050505050505050505600a165627a7a72305820fe4a1125dd6b7b7f8ecc2005110e076c9596e64a72b0df4cc330ab98f92382ac0029`

// DeployBatchTransaction deploys a new Ethereum contract, binding an instance of BatchTransaction to it.
func DeployBatchTransaction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BatchTransaction, error) {
	parsed, err := abi.JSON(strings.NewReader(BatchTransactionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BatchTransactionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BatchTransaction{BatchTransactionCaller: BatchTransactionCaller{contract: contract}, BatchTransactionTransactor: BatchTransactionTransactor{contract: contract}, BatchTransactionFilterer: BatchTransactionFilterer{contract: contract}}, nil
}

// BatchTransaction is an auto generated Go binding around an Ethereum contract.
type BatchTransaction struct {
	BatchTransactionCaller     // Read-only binding to the contract
	BatchTransactionTransactor // Write-only binding to the contract
	BatchTransactionFilterer   // Log filterer for contract events
}

// BatchTransactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchTransactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchTransactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchTransactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchTransactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchTransactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchTransactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchTransactionSession struct {
	Contract     *BatchTransaction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchTransactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchTransactionCallerSession struct {
	Contract *BatchTransactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BatchTransactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchTransactionTransactorSession struct {
	Contract     *BatchTransactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BatchTransactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchTransactionRaw struct {
	Contract *BatchTransaction // Generic contract binding to access the raw methods on
}

// BatchTransactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchTransactionCallerRaw struct {
	Contract *BatchTransactionCaller // Generic read-only contract binding to access the raw methods on
}

// BatchTransactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchTransactionTransactorRaw struct {
	Contract *BatchTransactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchTransaction creates a new instance of BatchTransaction, bound to a specific deployed contract.
func NewBatchTransaction(address common.Address, backend bind.ContractBackend) (*BatchTransaction, error) {
	contract, err := bindBatchTransaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchTransaction{BatchTransactionCaller: BatchTransactionCaller{contract: contract}, BatchTransactionTransactor: BatchTransactionTransactor{contract: contract}, BatchTransactionFilterer: BatchTransactionFilterer{contract: contract}}, nil
}

// NewBatchTransactionCaller creates a new read-only instance of BatchTransaction, bound to a specific deployed contract.
func NewBatchTransactionCaller(address common.Address, caller bind.ContractCaller) (*BatchTransactionCaller, error) {
	contract, err := bindBatchTransaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchTransactionCaller{contract: contract}, nil
}

// NewBatchTransactionTransactor creates a new write-only instance of BatchTransaction, bound to a specific deployed contract.
func NewBatchTransactionTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchTransactionTransactor, error) {
	contract, err := bindBatchTransaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchTransactionTransactor{contract: contract}, nil
}

// NewBatchTransactionFilterer creates a new log filterer instance of BatchTransaction, bound to a specific deployed contract.
func NewBatchTransactionFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchTransactionFilterer, error) {
	contract, err := bindBatchTransaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchTransactionFilterer{contract: contract}, nil
}

// bindBatchTransaction binds a generic wrapper to an already deployed contract.
func bindBatchTransaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BatchTransactionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchTransaction *BatchTransactionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BatchTransaction.Contract.BatchTransactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchTransaction *BatchTransactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchTransaction.Contract.BatchTransactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchTransaction *BatchTransactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchTransaction.Contract.BatchTransactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchTransaction *BatchTransactionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BatchTransaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchTransaction *BatchTransactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchTransaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchTransaction *BatchTransactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchTransaction.Contract.contract.Transact(opts, method, params...)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe9861ab1.
//
// Solidity: function executeBatch(uint8[] _v, bytes32[] _r, bytes32[] _s, address[] _from, address[] _dest, uint256[] _value, uint256[] _ts) returns()
func (_BatchTransaction *BatchTransactionTransactor) ExecuteBatch(opts *bind.TransactOpts, _v []uint8, _r [][32]byte, _s [][32]byte, _from []common.Address, _dest []common.Address, _value []*big.Int, _ts []*big.Int) (*types.Transaction, error) {
	return _BatchTransaction.contract.Transact(opts, "executeBatch", _v, _r, _s, _from, _dest, _value, _ts)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe9861ab1.
//
// Solidity: function executeBatch(uint8[] _v, bytes32[] _r, bytes32[] _s, address[] _from, address[] _dest, uint256[] _value, uint256[] _ts) returns()
func (_BatchTransaction *BatchTransactionSession) ExecuteBatch(_v []uint8, _r [][32]byte, _s [][32]byte, _from []common.Address, _dest []common.Address, _value []*big.Int, _ts []*big.Int) (*types.Transaction, error) {
	return _BatchTransaction.Contract.ExecuteBatch(&_BatchTransaction.TransactOpts, _v, _r, _s, _from, _dest, _value, _ts)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe9861ab1.
//
// Solidity: function executeBatch(uint8[] _v, bytes32[] _r, bytes32[] _s, address[] _from, address[] _dest, uint256[] _value, uint256[] _ts) returns()
func (_BatchTransaction *BatchTransactionTransactorSession) ExecuteBatch(_v []uint8, _r [][32]byte, _s [][32]byte, _from []common.Address, _dest []common.Address, _value []*big.Int, _ts []*big.Int) (*types.Transaction, error) {
	return _BatchTransaction.Contract.ExecuteBatch(&_BatchTransaction.TransactOpts, _v, _r, _s, _from, _dest, _value, _ts)
}

// BatchTransactionTransactionStatusIterator is returned from FilterTransactionStatus and is used to iterate over the raw logs and unpacked data for TransactionStatus events raised by the BatchTransaction contract.
type BatchTransactionTransactionStatusIterator struct {
	Event *BatchTransactionTransactionStatus // Event containing the contract specifics and raw log

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
func (it *BatchTransactionTransactionStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchTransactionTransactionStatus)
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
		it.Event = new(BatchTransactionTransactionStatus)
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
func (it *BatchTransactionTransactionStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchTransactionTransactionStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchTransactionTransactionStatus represents a TransactionStatus event raised by the BatchTransaction contract.
type BatchTransactionTransactionStatus struct {
	Success bool
	From    common.Address
	Dest    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransactionStatus is a free log retrieval operation binding the contract event 0x41f5d8bdb5631ef260ce5fb88cf0a48c22d261827bebce333beb5963b3fe3edd.
//
// Solidity: event TransactionStatus(bool success, address from, address dest)
func (_BatchTransaction *BatchTransactionFilterer) FilterTransactionStatus(opts *bind.FilterOpts) (*BatchTransactionTransactionStatusIterator, error) {

	logs, sub, err := _BatchTransaction.contract.FilterLogs(opts, "TransactionStatus")
	if err != nil {
		return nil, err
	}
	return &BatchTransactionTransactionStatusIterator{contract: _BatchTransaction.contract, event: "TransactionStatus", logs: logs, sub: sub}, nil
}

// WatchTransactionStatus is a free log subscription operation binding the contract event 0x41f5d8bdb5631ef260ce5fb88cf0a48c22d261827bebce333beb5963b3fe3edd.
//
// Solidity: event TransactionStatus(bool success, address from, address dest)
func (_BatchTransaction *BatchTransactionFilterer) WatchTransactionStatus(opts *bind.WatchOpts, sink chan<- *BatchTransactionTransactionStatus) (event.Subscription, error) {

	logs, sub, err := _BatchTransaction.contract.WatchLogs(opts, "TransactionStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchTransactionTransactionStatus)
				if err := _BatchTransaction.contract.UnpackLog(event, "TransactionStatus", log); err != nil {
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

// ElvWalletABI is the input ABI used to generate the binding from.
const ElvWalletABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_dest\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_ts\",\"type\":\"uint256\"}],\"name\":\"validateTransaction\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_dest\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_ts\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ts\",\"type\":\"uint256\"}],\"name\":\"validateTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"code\",\"type\":\"int256\"}],\"name\":\"ExecStatus\",\"type\":\"event\"}]"

// ElvWalletBin is the compiled bytecode used for deploying new contracts.
const ElvWalletBin = `0x6080604052604051602080610564833981016040525160018054600160a060020a031916600160a060020a03909216919091179055610521806100436000396000f3006080604052600436106100615763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663763d5ee681146100665780638da5cb5b146100ad57806394489c42146100de578063f50b2efe14610111575b600080fd5b34801561007257600080fd5b5061009960ff60043516602435604435600160a060020a036064351660843560a435610144565b604080519115158252519081900360200190f35b3480156100b957600080fd5b506100c26102ad565b60408051600160a060020a039092168252519081900360200190f35b3480156100ea57600080fd5b5061009960ff60043516602435604435600160a060020a036064351660843560a4356102bc565b34801561011d57600080fd5b506101296004356103c5565b60408051921515835260208301919091528051918290030190f35b604080516c01000000000000000000000000308102602080840191909152600160a060020a038716909102603483015260488201859052606880830185905283518084039091018152608890920192839052815160009384938493909282918401908083835b602083106101c95780518252601f1990920191602091820191016101aa565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902091506001828a8a8a604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af115801561026f573d6000803e3d6000fd5b5050604051601f190151600154909250600160a060020a03808416911614905061029c57600092506102a1565b600192505b50509695505050505050565b600154600160a060020a031681565b60008080303185111561030657604080516001815290517fd35e92c995c781ece41c1626ca10bcc46127ab6f6a443a9d4072a9961a2fecdf9181900360200190a1600092506102a1565b610314898989898989610144565b915081151561035a57604080516002815290517fd35e92c995c781ece41c1626ca10bcc46127ab6f6a443a9d4072a9961a2fecdf9181900360200190a1600092506102a1565b604051600160a060020a0387169086156108fc029087906000818181858888f19350505050905080151561029c57604080516004815290517fd35e92c995c781ece41c1626ca10bcc46127ab6f6a443a9d4072a9961a2fecdf9181900360200190a1600092506102a1565b600080806103d9818563ffffffff61042216565b60005490915081101561041457836000828154811015156103f657fe5b9060005260206000200154101515610414576000925082915061041c565b600181925092505b50915091565b60008060008085805490506000141561043e57600093506104cd565b85546000935091505b8183101561048f5761045983836104d6565b905084868281548110151561046a57fe5b906000526020600020015411156104835780915061048a565b8060010192505b610447565b6000831180156104b957508486600185038154811015156104ac57fe5b9060005260206000200154145b156104c9576001830393506104cd565b8293505b50505092915050565b60006002600184811690841601046002830460028504010193925050505600a165627a7a7230582028835e500ff35a8f5d21955e0c9f659cc2a7edc188ee31f3d4bd31b6b0759c680029`

// DeployElvWallet deploys a new Ethereum contract, binding an instance of ElvWallet to it.
func DeployElvWallet(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *ElvWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(ElvWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ElvWalletBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ElvWallet{ElvWalletCaller: ElvWalletCaller{contract: contract}, ElvWalletTransactor: ElvWalletTransactor{contract: contract}, ElvWalletFilterer: ElvWalletFilterer{contract: contract}}, nil
}

// ElvWallet is an auto generated Go binding around an Ethereum contract.
type ElvWallet struct {
	ElvWalletCaller     // Read-only binding to the contract
	ElvWalletTransactor // Write-only binding to the contract
	ElvWalletFilterer   // Log filterer for contract events
}

// ElvWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type ElvWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElvWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ElvWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElvWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ElvWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElvWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ElvWalletSession struct {
	Contract     *ElvWallet        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ElvWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ElvWalletCallerSession struct {
	Contract *ElvWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ElvWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ElvWalletTransactorSession struct {
	Contract     *ElvWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ElvWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type ElvWalletRaw struct {
	Contract *ElvWallet // Generic contract binding to access the raw methods on
}

// ElvWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ElvWalletCallerRaw struct {
	Contract *ElvWalletCaller // Generic read-only contract binding to access the raw methods on
}

// ElvWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ElvWalletTransactorRaw struct {
	Contract *ElvWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElvWallet creates a new instance of ElvWallet, bound to a specific deployed contract.
func NewElvWallet(address common.Address, backend bind.ContractBackend) (*ElvWallet, error) {
	contract, err := bindElvWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ElvWallet{ElvWalletCaller: ElvWalletCaller{contract: contract}, ElvWalletTransactor: ElvWalletTransactor{contract: contract}, ElvWalletFilterer: ElvWalletFilterer{contract: contract}}, nil
}

// NewElvWalletCaller creates a new read-only instance of ElvWallet, bound to a specific deployed contract.
func NewElvWalletCaller(address common.Address, caller bind.ContractCaller) (*ElvWalletCaller, error) {
	contract, err := bindElvWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ElvWalletCaller{contract: contract}, nil
}

// NewElvWalletTransactor creates a new write-only instance of ElvWallet, bound to a specific deployed contract.
func NewElvWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*ElvWalletTransactor, error) {
	contract, err := bindElvWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ElvWalletTransactor{contract: contract}, nil
}

// NewElvWalletFilterer creates a new log filterer instance of ElvWallet, bound to a specific deployed contract.
func NewElvWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*ElvWalletFilterer, error) {
	contract, err := bindElvWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ElvWalletFilterer{contract: contract}, nil
}

// bindElvWallet binds a generic wrapper to an already deployed contract.
func bindElvWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ElvWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElvWallet *ElvWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ElvWallet.Contract.ElvWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElvWallet *ElvWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElvWallet.Contract.ElvWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElvWallet *ElvWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElvWallet.Contract.ElvWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElvWallet *ElvWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ElvWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElvWallet *ElvWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElvWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElvWallet *ElvWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElvWallet.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ElvWallet *ElvWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ElvWallet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ElvWallet *ElvWalletSession) Owner() (common.Address, error) {
	return _ElvWallet.Contract.Owner(&_ElvWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ElvWallet *ElvWalletCallerSession) Owner() (common.Address, error) {
	return _ElvWallet.Contract.Owner(&_ElvWallet.CallOpts)
}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 _ts) constant returns(bool, uint256)
func (_ElvWallet *ElvWalletCaller) ValidateTimestamp(opts *bind.CallOpts, _ts *big.Int) (bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ElvWallet.contract.Call(opts, out, "validateTimestamp", _ts)
	return *ret0, *ret1, err
}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 _ts) constant returns(bool, uint256)
func (_ElvWallet *ElvWalletSession) ValidateTimestamp(_ts *big.Int) (bool, *big.Int, error) {
	return _ElvWallet.Contract.ValidateTimestamp(&_ElvWallet.CallOpts, _ts)
}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 _ts) constant returns(bool, uint256)
func (_ElvWallet *ElvWalletCallerSession) ValidateTimestamp(_ts *big.Int) (bool, *big.Int, error) {
	return _ElvWallet.Contract.ValidateTimestamp(&_ElvWallet.CallOpts, _ts)
}

// ValidateTransaction is a free data retrieval call binding the contract method 0x763d5ee6.
//
// Solidity: function validateTransaction(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) constant returns(bool)
func (_ElvWallet *ElvWalletCaller) ValidateTransaction(opts *bind.CallOpts, _v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ElvWallet.contract.Call(opts, out, "validateTransaction", _v, _r, _s, _dest, _value, _ts)
	return *ret0, err
}

// ValidateTransaction is a free data retrieval call binding the contract method 0x763d5ee6.
//
// Solidity: function validateTransaction(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) constant returns(bool)
func (_ElvWallet *ElvWalletSession) ValidateTransaction(_v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (bool, error) {
	return _ElvWallet.Contract.ValidateTransaction(&_ElvWallet.CallOpts, _v, _r, _s, _dest, _value, _ts)
}

// ValidateTransaction is a free data retrieval call binding the contract method 0x763d5ee6.
//
// Solidity: function validateTransaction(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) constant returns(bool)
func (_ElvWallet *ElvWalletCallerSession) ValidateTransaction(_v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (bool, error) {
	return _ElvWallet.Contract.ValidateTransaction(&_ElvWallet.CallOpts, _v, _r, _s, _dest, _value, _ts)
}

// Execute is a paid mutator transaction binding the contract method 0x94489c42.
//
// Solidity: function execute(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) returns(bool)
func (_ElvWallet *ElvWalletTransactor) Execute(opts *bind.TransactOpts, _v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (*types.Transaction, error) {
	return _ElvWallet.contract.Transact(opts, "execute", _v, _r, _s, _dest, _value, _ts)
}

// Execute is a paid mutator transaction binding the contract method 0x94489c42.
//
// Solidity: function execute(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) returns(bool)
func (_ElvWallet *ElvWalletSession) Execute(_v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (*types.Transaction, error) {
	return _ElvWallet.Contract.Execute(&_ElvWallet.TransactOpts, _v, _r, _s, _dest, _value, _ts)
}

// Execute is a paid mutator transaction binding the contract method 0x94489c42.
//
// Solidity: function execute(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts) returns(bool)
func (_ElvWallet *ElvWalletTransactorSession) Execute(_v uint8, _r [32]byte, _s [32]byte, _dest common.Address, _value *big.Int, _ts *big.Int) (*types.Transaction, error) {
	return _ElvWallet.Contract.Execute(&_ElvWallet.TransactOpts, _v, _r, _s, _dest, _value, _ts)
}

// ElvWalletExecStatusIterator is returned from FilterExecStatus and is used to iterate over the raw logs and unpacked data for ExecStatus events raised by the ElvWallet contract.
type ElvWalletExecStatusIterator struct {
	Event *ElvWalletExecStatus // Event containing the contract specifics and raw log

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
func (it *ElvWalletExecStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElvWalletExecStatus)
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
		it.Event = new(ElvWalletExecStatus)
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
func (it *ElvWalletExecStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElvWalletExecStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElvWalletExecStatus represents a ExecStatus event raised by the ElvWallet contract.
type ElvWalletExecStatus struct {
	Code *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExecStatus is a free log retrieval operation binding the contract event 0xd35e92c995c781ece41c1626ca10bcc46127ab6f6a443a9d4072a9961a2fecdf.
//
// Solidity: event ExecStatus(int256 code)
func (_ElvWallet *ElvWalletFilterer) FilterExecStatus(opts *bind.FilterOpts) (*ElvWalletExecStatusIterator, error) {

	logs, sub, err := _ElvWallet.contract.FilterLogs(opts, "ExecStatus")
	if err != nil {
		return nil, err
	}
	return &ElvWalletExecStatusIterator{contract: _ElvWallet.contract, event: "ExecStatus", logs: logs, sub: sub}, nil
}

// WatchExecStatus is a free log subscription operation binding the contract event 0xd35e92c995c781ece41c1626ca10bcc46127ab6f6a443a9d4072a9961a2fecdf.
//
// Solidity: event ExecStatus(int256 code)
func (_ElvWallet *ElvWalletFilterer) WatchExecStatus(opts *bind.WatchOpts, sink chan<- *ElvWalletExecStatus) (event.Subscription, error) {

	logs, sub, err := _ElvWallet.contract.WatchLogs(opts, "ExecStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElvWalletExecStatus)
				if err := _ElvWallet.contract.UnpackLog(event, "ExecStatus", log); err != nil {
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

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
const MathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820239ffefe423472e750b4fd5193f3c4fc943128384872e96d094a8d15d5627ddd0029`

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
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
