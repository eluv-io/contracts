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

// BatchTransactionABI is the input ABI used to generate the binding from.
const BatchTransactionABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_v\",\"type\":\"uint8[]\"},{\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"name\":\"_from\",\"type\":\"address[]\"},{\"name\":\"_dest\",\"type\":\"address[]\"},{\"name\":\"_value\",\"type\":\"uint256[]\"},{\"name\":\"_ts\",\"type\":\"uint256[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"TransactionStatus\",\"type\":\"event\"}]"

// BatchTransactionBin is the compiled bytecode used for deploying new contracts.
const BatchTransactionBin = `0x608060405260008054600160a060020a031916331790556104cb806100256000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663e9861ab18114610042575b005b34801561004e57600080fd5b506040805160206004803580820135838102808601850190965280855261004095369593946024949385019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506101ed9650505050505050565b600080548190819073ffffffffffffffffffffffffffffffffffffffff16331461021657600080fd5b88518a511461022457600080fd5b875189511461023257600080fd5b865188511461024057600080fd5b855187511461024e57600080fd5b845186511461025c57600080fd5b835185511461026a57600080fd5b600092505b895183101561049357868381518110151561028657fe5b9060200190602002015191508173ffffffffffffffffffffffffffffffffffffffff166394489c428b858151811015156102bc57fe5b906020019060200201518b868151811015156102d457fe5b906020019060200201518b878151811015156102ec57fe5b906020019060200201518a8881518110151561030457fe5b906020019060200201518a8981518110151561031c57fe5b906020019060200201518a8a81518110151561033457fe5b6020908102909101810151604080517c010000000000000000000000000000000000000000000000000000000063ffffffff8b1602815260ff90981660048901526024880196909652604487019490945273ffffffffffffffffffffffffffffffffffffffff9092166064860152608485015260a4840191909152905160c4808401938290030181600087803b1580156103cd57600080fd5b505af11580156103e1573d6000803e3d6000fd5b505050506040513d60208110156103f757600080fd5b505190507f41f5d8bdb5631ef260ce5fb88cf0a48c22d261827bebce333beb5963b3fe3edd81888581518110151561042b57fe5b90602001906020020151888681518110151561044357fe5b602090810290910181015160408051941515855273ffffffffffffffffffffffffffffffffffffffff938416928501929092529190911682820152519081900360600190a160019092019161026f565b505050505050505050505600a165627a7a723058204ec3d64c516918fc23b6fbf9adfcae0d0c5fff3a3cb7af85c8b673ded64a86f20029`

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

// ChannelWalletABI is the input ABI used to generate the binding from.
const ChannelWalletABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"currentTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"guarantor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_dest\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_ts\",\"type\":\"uint256\"}],\"name\":\"validateTransaction\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_dest\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_ts\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"execStatusOk\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ts\",\"type\":\"uint256\"}],\"name\":\"validateTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_guarantor\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"guarantor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"code\",\"type\":\"int256\"}],\"name\":\"ExecStatus\",\"type\":\"event\"}]"

// ChannelWalletBin is the compiled bytecode used for deploying new contracts.
const ChannelWalletBin = `0x6080604052604051602080610569833981016040525160018054600160a060020a0319908116331790915560028054600160a060020a039093169290911691909117905542600055610513806100566000396000f3006080604052600436106100825763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631e2ff94f811461008757806371cf516b146100ae578063763d5ee6146100df5780638da5cb5b1461012657806394489c421461013b57806395ba60ba1461016e578063f50b2efe14610183575b600080fd5b34801561009357600080fd5b5061009c61019b565b60408051918252519081900360200190f35b3480156100ba57600080fd5b506100c36101a1565b60408051600160a060020a039092168252519081900360200190f35b3480156100eb57600080fd5b5061011260ff60043516602435604435600160a060020a036064351660843560a4356101b0565b604080519115158252519081900360200190f35b34801561013257600080fd5b506100c3610319565b34801561014757600080fd5b5061011260ff60043516602435604435600160a060020a036064351660843560a435610328565b34801561017a57600080fd5b5061009c6104a6565b34801561018f57600080fd5b506101126004356104ab565b60005481565b600254600160a060020a031681565b604080516c01000000000000000000000000308102602080840191909152600160a060020a038716909102603483015260488201859052606880830185905283518084039091018152608890920192839052815160009384938493909282918401908083835b602083106102355780518252601f199092019160209182019101610216565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902091506001828a8a8a604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af11580156102db573d6000803e3d6000fd5b5050604051601f190151600154909250600160a060020a038084169116149050610308576000925061030d565b600192505b50509695505050505050565b600154600160a060020a031681565b60025460009081908190600160a060020a0316331461034657600080fd5b600054841161035457600080fd5b303185111561039b5760025460408051600160a060020a0390921682526001602083015280516000805160206104c88339815191529281900390910190a16000925061030d565b6103a98989898989896101b0565b91508115156103f3576002805460408051600160a060020a039092168252602082019290925281516000805160206104c8833981519152929181900390910190a16000925061030d565b6000848155604051600160a060020a0388169187156108fc02918891818181858888f1935050505090508015156104625760025460408051600160a060020a0390921682526003602083015280516000805160206104c88339815191529281900390910190a16000925061030d565b60025460408051600160a060020a0390921682526000602083015280516000805160206104c88339815191529281900390910190a150600198975050505050505050565b600081565b600080548211156104be575060016104c2565b5060005b9190505600583d8312ef7016406c7ea8ba9796b9e55ac1fdc22455754cbc93869509faefada165627a7a7230582044e6db8b470b05696314e2d7363190a0947115d564416b7578673c882dd3c8580029`

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

// ExecStatusOk is a free data retrieval call binding the contract method 0x95ba60ba.
//
// Solidity: function execStatusOk() constant returns(int256)
func (_ChannelWallet *ChannelWalletCaller) ExecStatusOk(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChannelWallet.contract.Call(opts, out, "execStatusOk")
	return *ret0, err
}

// ExecStatusOk is a free data retrieval call binding the contract method 0x95ba60ba.
//
// Solidity: function execStatusOk() constant returns(int256)
func (_ChannelWallet *ChannelWalletSession) ExecStatusOk() (*big.Int, error) {
	return _ChannelWallet.Contract.ExecStatusOk(&_ChannelWallet.CallOpts)
}

// ExecStatusOk is a free data retrieval call binding the contract method 0x95ba60ba.
//
// Solidity: function execStatusOk() constant returns(int256)
func (_ChannelWallet *ChannelWalletCallerSession) ExecStatusOk() (*big.Int, error) {
	return _ChannelWallet.Contract.ExecStatusOk(&_ChannelWallet.CallOpts)
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
// Solidity: function validateTimestamp(uint256 _ts) constant returns(bool)
func (_ChannelWallet *ChannelWalletCaller) ValidateTimestamp(opts *bind.CallOpts, _ts *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChannelWallet.contract.Call(opts, out, "validateTimestamp", _ts)
	return *ret0, err
}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 _ts) constant returns(bool)
func (_ChannelWallet *ChannelWalletSession) ValidateTimestamp(_ts *big.Int) (bool, error) {
	return _ChannelWallet.Contract.ValidateTimestamp(&_ChannelWallet.CallOpts, _ts)
}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 _ts) constant returns(bool)
func (_ChannelWallet *ChannelWalletCallerSession) ValidateTimestamp(_ts *big.Int) (bool, error) {
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
