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

// TestElvPrecompileABI is the input ABI used to generate the binding from.
const TestElvPrecompileABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"elv_precomp_addr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"testPrecompile\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transOrd\",\"type\":\"int256\"},{\"name\":\"batchId\",\"type\":\"uint256\"}],\"name\":\"getTrans\",\"outputs\":[{\"name\":\"ret\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes32\"}],\"name\":\"GetElvData\",\"type\":\"event\"}]"

// TestElvPrecompileBin is the compiled bytecode used for deploying new contracts.
const TestElvPrecompileBin = `0x60806040819052600080547f662829000000000000000000000000000000000000000000000000000000000090925260ff600160a060020a03199092169190911760a060020a60c060020a0319167726121ff0000000000000000000000000000000000000000017905534801561007557600080fd5b50610275806100856000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166320104aa6811461005b5780639f6f32aa14610099578063bc8b7495146100b0575b600080fd5b34801561006757600080fd5b506100706100dd565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156100a557600080fd5b506100ae6100f9565b005b3480156100bc57600080fd5b506100cb6004356024356101ae565b60408051918252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600080805b8115156101a9576101108360006101ae565b90507fff000000000000000000000000000000000000000000000000000000000000007f0100000000000000000000000000000000000000000000000000000000000000600083901a0216151561016a57600191506101a4565b6040805182815290517fe5a6a3cf3a8e75295099dc4358f1909dff82370e8daa4f8d435726548634da4f9181900360200190a16001909201915b6100fe565b505050565b600080546040517c0100000000000000000000000000000000000000000000000000000000740100000000000000000000000000000000000000008304810280835290860260048301819052600883018690529092909173ffffffffffffffffffffffffffffffffffffffff90911690602081602c81888681f180151561023457600080fd5b508051602090910160405296955050505050505600a165627a7a7230582097d8f97867a5cd0ff9977257c568d1c0110992c94e2d1eaeeec5744ae23cb5f70029`

// DeployTestElvPrecompile deploys a new Ethereum contract, binding an instance of TestElvPrecompile to it.
func DeployTestElvPrecompile(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestElvPrecompile, error) {
	parsed, err := abi.JSON(strings.NewReader(TestElvPrecompileABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestElvPrecompileBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestElvPrecompile{TestElvPrecompileCaller: TestElvPrecompileCaller{contract: contract}, TestElvPrecompileTransactor: TestElvPrecompileTransactor{contract: contract}, TestElvPrecompileFilterer: TestElvPrecompileFilterer{contract: contract}}, nil
}

// TestElvPrecompile is an auto generated Go binding around an Ethereum contract.
type TestElvPrecompile struct {
	TestElvPrecompileCaller     // Read-only binding to the contract
	TestElvPrecompileTransactor // Write-only binding to the contract
	TestElvPrecompileFilterer   // Log filterer for contract events
}

// TestElvPrecompileCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestElvPrecompileCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestElvPrecompileTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestElvPrecompileTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestElvPrecompileFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestElvPrecompileFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestElvPrecompileSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestElvPrecompileSession struct {
	Contract     *TestElvPrecompile // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TestElvPrecompileCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestElvPrecompileCallerSession struct {
	Contract *TestElvPrecompileCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TestElvPrecompileTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestElvPrecompileTransactorSession struct {
	Contract     *TestElvPrecompileTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TestElvPrecompileRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestElvPrecompileRaw struct {
	Contract *TestElvPrecompile // Generic contract binding to access the raw methods on
}

// TestElvPrecompileCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestElvPrecompileCallerRaw struct {
	Contract *TestElvPrecompileCaller // Generic read-only contract binding to access the raw methods on
}

// TestElvPrecompileTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestElvPrecompileTransactorRaw struct {
	Contract *TestElvPrecompileTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestElvPrecompile creates a new instance of TestElvPrecompile, bound to a specific deployed contract.
func NewTestElvPrecompile(address common.Address, backend bind.ContractBackend) (*TestElvPrecompile, error) {
	contract, err := bindTestElvPrecompile(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestElvPrecompile{TestElvPrecompileCaller: TestElvPrecompileCaller{contract: contract}, TestElvPrecompileTransactor: TestElvPrecompileTransactor{contract: contract}, TestElvPrecompileFilterer: TestElvPrecompileFilterer{contract: contract}}, nil
}

// NewTestElvPrecompileCaller creates a new read-only instance of TestElvPrecompile, bound to a specific deployed contract.
func NewTestElvPrecompileCaller(address common.Address, caller bind.ContractCaller) (*TestElvPrecompileCaller, error) {
	contract, err := bindTestElvPrecompile(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestElvPrecompileCaller{contract: contract}, nil
}

// NewTestElvPrecompileTransactor creates a new write-only instance of TestElvPrecompile, bound to a specific deployed contract.
func NewTestElvPrecompileTransactor(address common.Address, transactor bind.ContractTransactor) (*TestElvPrecompileTransactor, error) {
	contract, err := bindTestElvPrecompile(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestElvPrecompileTransactor{contract: contract}, nil
}

// NewTestElvPrecompileFilterer creates a new log filterer instance of TestElvPrecompile, bound to a specific deployed contract.
func NewTestElvPrecompileFilterer(address common.Address, filterer bind.ContractFilterer) (*TestElvPrecompileFilterer, error) {
	contract, err := bindTestElvPrecompile(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestElvPrecompileFilterer{contract: contract}, nil
}

// bindTestElvPrecompile binds a generic wrapper to an already deployed contract.
func bindTestElvPrecompile(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestElvPrecompileABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestElvPrecompile *TestElvPrecompileRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestElvPrecompile.Contract.TestElvPrecompileCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestElvPrecompile *TestElvPrecompileRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestElvPrecompile.Contract.TestElvPrecompileTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestElvPrecompile *TestElvPrecompileRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestElvPrecompile.Contract.TestElvPrecompileTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestElvPrecompile *TestElvPrecompileCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestElvPrecompile.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestElvPrecompile *TestElvPrecompileTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestElvPrecompile.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestElvPrecompile *TestElvPrecompileTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestElvPrecompile.Contract.contract.Transact(opts, method, params...)
}

// ElvPrecompAddr is a free data retrieval call binding the contract method 0x20104aa6.
//
// Solidity: function elv_precomp_addr() constant returns(address)
func (_TestElvPrecompile *TestElvPrecompileCaller) ElvPrecompAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TestElvPrecompile.contract.Call(opts, out, "elv_precomp_addr")
	return *ret0, err
}

// ElvPrecompAddr is a free data retrieval call binding the contract method 0x20104aa6.
//
// Solidity: function elv_precomp_addr() constant returns(address)
func (_TestElvPrecompile *TestElvPrecompileSession) ElvPrecompAddr() (common.Address, error) {
	return _TestElvPrecompile.Contract.ElvPrecompAddr(&_TestElvPrecompile.CallOpts)
}

// ElvPrecompAddr is a free data retrieval call binding the contract method 0x20104aa6.
//
// Solidity: function elv_precomp_addr() constant returns(address)
func (_TestElvPrecompile *TestElvPrecompileCallerSession) ElvPrecompAddr() (common.Address, error) {
	return _TestElvPrecompile.Contract.ElvPrecompAddr(&_TestElvPrecompile.CallOpts)
}

// GetTrans is a free data retrieval call binding the contract method 0xbc8b7495.
//
// Solidity: function getTrans(int256 transOrd, uint256 batchId) constant returns(bytes32 ret)
func (_TestElvPrecompile *TestElvPrecompileCaller) GetTrans(opts *bind.CallOpts, transOrd *big.Int, batchId *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TestElvPrecompile.contract.Call(opts, out, "getTrans", transOrd, batchId)
	return *ret0, err
}

// GetTrans is a free data retrieval call binding the contract method 0xbc8b7495.
//
// Solidity: function getTrans(int256 transOrd, uint256 batchId) constant returns(bytes32 ret)
func (_TestElvPrecompile *TestElvPrecompileSession) GetTrans(transOrd *big.Int, batchId *big.Int) ([32]byte, error) {
	return _TestElvPrecompile.Contract.GetTrans(&_TestElvPrecompile.CallOpts, transOrd, batchId)
}

// GetTrans is a free data retrieval call binding the contract method 0xbc8b7495.
//
// Solidity: function getTrans(int256 transOrd, uint256 batchId) constant returns(bytes32 ret)
func (_TestElvPrecompile *TestElvPrecompileCallerSession) GetTrans(transOrd *big.Int, batchId *big.Int) ([32]byte, error) {
	return _TestElvPrecompile.Contract.GetTrans(&_TestElvPrecompile.CallOpts, transOrd, batchId)
}

// TestPrecompile is a paid mutator transaction binding the contract method 0x9f6f32aa.
//
// Solidity: function testPrecompile() returns()
func (_TestElvPrecompile *TestElvPrecompileTransactor) TestPrecompile(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestElvPrecompile.contract.Transact(opts, "testPrecompile")
}

// TestPrecompile is a paid mutator transaction binding the contract method 0x9f6f32aa.
//
// Solidity: function testPrecompile() returns()
func (_TestElvPrecompile *TestElvPrecompileSession) TestPrecompile() (*types.Transaction, error) {
	return _TestElvPrecompile.Contract.TestPrecompile(&_TestElvPrecompile.TransactOpts)
}

// TestPrecompile is a paid mutator transaction binding the contract method 0x9f6f32aa.
//
// Solidity: function testPrecompile() returns()
func (_TestElvPrecompile *TestElvPrecompileTransactorSession) TestPrecompile() (*types.Transaction, error) {
	return _TestElvPrecompile.Contract.TestPrecompile(&_TestElvPrecompile.TransactOpts)
}

// TestElvPrecompileGetElvDataIterator is returned from FilterGetElvData and is used to iterate over the raw logs and unpacked data for GetElvData events raised by the TestElvPrecompile contract.
type TestElvPrecompileGetElvDataIterator struct {
	Event *TestElvPrecompileGetElvData // Event containing the contract specifics and raw log

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
func (it *TestElvPrecompileGetElvDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestElvPrecompileGetElvData)
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
		it.Event = new(TestElvPrecompileGetElvData)
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
func (it *TestElvPrecompileGetElvDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestElvPrecompileGetElvDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestElvPrecompileGetElvData represents a GetElvData event raised by the TestElvPrecompile contract.
type TestElvPrecompileGetElvData struct {
	Data [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGetElvData is a free log retrieval operation binding the contract event 0xe5a6a3cf3a8e75295099dc4358f1909dff82370e8daa4f8d435726548634da4f.
//
// Solidity: event GetElvData(bytes32 data)
func (_TestElvPrecompile *TestElvPrecompileFilterer) FilterGetElvData(opts *bind.FilterOpts) (*TestElvPrecompileGetElvDataIterator, error) {

	logs, sub, err := _TestElvPrecompile.contract.FilterLogs(opts, "GetElvData")
	if err != nil {
		return nil, err
	}
	return &TestElvPrecompileGetElvDataIterator{contract: _TestElvPrecompile.contract, event: "GetElvData", logs: logs, sub: sub}, nil
}

// WatchGetElvData is a free log subscription operation binding the contract event 0xe5a6a3cf3a8e75295099dc4358f1909dff82370e8daa4f8d435726548634da4f.
//
// Solidity: event GetElvData(bytes32 data)
func (_TestElvPrecompile *TestElvPrecompileFilterer) WatchGetElvData(opts *bind.WatchOpts, sink chan<- *TestElvPrecompileGetElvData) (event.Subscription, error) {

	logs, sub, err := _TestElvPrecompile.contract.WatchLogs(opts, "GetElvData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestElvPrecompileGetElvData)
				if err := _TestElvPrecompile.contract.UnpackLog(event, "GetElvData", log); err != nil {
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
