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

// AccessibleABI is the input ABI used to generate the binding from.
const AccessibleABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"accessRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestValidity\",\"type\":\"uint256\"}],\"name\":\"AccessRequest\",\"type\":\"event\"}]"

// AccessibleBin is the compiled bytecode used for deploying new contracts.
const AccessibleBin = `0x60606040523415600e57600080fd5b60cf8061001c6000396000f300606060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663f155188781146043575b600080fd5b3415604d57600080fd5b60536067565b604051901515815260200160405180910390f35b60007fca8116fd70ecea0cf35f7c08eb31491cce5ebd6bafef5bd6426862df29840463600060405190815260200160405180910390a1506001905600a165627a7a72305820735028607cdc4c73e6525c571ff21d24003b3e565166136fa155a3e7a56cb4250029`

// DeployAccessible deploys a new Ethereum contract, binding an instance of Accessible to it.
func DeployAccessible(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Accessible, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessibleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccessibleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Accessible{AccessibleCaller: AccessibleCaller{contract: contract}, AccessibleTransactor: AccessibleTransactor{contract: contract}, AccessibleFilterer: AccessibleFilterer{contract: contract}}, nil
}

// Accessible is an auto generated Go binding around an Ethereum contract.
type Accessible struct {
	AccessibleCaller     // Read-only binding to the contract
	AccessibleTransactor // Write-only binding to the contract
	AccessibleFilterer   // Log filterer for contract events
}

// AccessibleCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessibleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessibleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessibleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessibleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessibleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessibleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessibleSession struct {
	Contract     *Accessible       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessibleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessibleCallerSession struct {
	Contract *AccessibleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AccessibleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessibleTransactorSession struct {
	Contract     *AccessibleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AccessibleRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessibleRaw struct {
	Contract *Accessible // Generic contract binding to access the raw methods on
}

// AccessibleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessibleCallerRaw struct {
	Contract *AccessibleCaller // Generic read-only contract binding to access the raw methods on
}

// AccessibleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessibleTransactorRaw struct {
	Contract *AccessibleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessible creates a new instance of Accessible, bound to a specific deployed contract.
func NewAccessible(address common.Address, backend bind.ContractBackend) (*Accessible, error) {
	contract, err := bindAccessible(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Accessible{AccessibleCaller: AccessibleCaller{contract: contract}, AccessibleTransactor: AccessibleTransactor{contract: contract}, AccessibleFilterer: AccessibleFilterer{contract: contract}}, nil
}

// NewAccessibleCaller creates a new read-only instance of Accessible, bound to a specific deployed contract.
func NewAccessibleCaller(address common.Address, caller bind.ContractCaller) (*AccessibleCaller, error) {
	contract, err := bindAccessible(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessibleCaller{contract: contract}, nil
}

// NewAccessibleTransactor creates a new write-only instance of Accessible, bound to a specific deployed contract.
func NewAccessibleTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessibleTransactor, error) {
	contract, err := bindAccessible(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessibleTransactor{contract: contract}, nil
}

// NewAccessibleFilterer creates a new log filterer instance of Accessible, bound to a specific deployed contract.
func NewAccessibleFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessibleFilterer, error) {
	contract, err := bindAccessible(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessibleFilterer{contract: contract}, nil
}

// bindAccessible binds a generic wrapper to an already deployed contract.
func bindAccessible(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessibleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accessible *AccessibleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Accessible.Contract.AccessibleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accessible *AccessibleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accessible.Contract.AccessibleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accessible *AccessibleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accessible.Contract.AccessibleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accessible *AccessibleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Accessible.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accessible *AccessibleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accessible.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accessible *AccessibleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accessible.Contract.contract.Transact(opts, method, params...)
}

// AccessRequest is a paid mutator transaction binding the contract method 0xf1551887.
//
// Solidity: function accessRequest() returns(bool)
func (_Accessible *AccessibleTransactor) AccessRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accessible.contract.Transact(opts, "accessRequest")
}

// AccessRequest is a paid mutator transaction binding the contract method 0xf1551887.
//
// Solidity: function accessRequest() returns(bool)
func (_Accessible *AccessibleSession) AccessRequest() (*types.Transaction, error) {
	return _Accessible.Contract.AccessRequest(&_Accessible.TransactOpts)
}

// AccessRequest is a paid mutator transaction binding the contract method 0xf1551887.
//
// Solidity: function accessRequest() returns(bool)
func (_Accessible *AccessibleTransactorSession) AccessRequest() (*types.Transaction, error) {
	return _Accessible.Contract.AccessRequest(&_Accessible.TransactOpts)
}

// AccessibleAccessRequestIterator is returned from FilterAccessRequest and is used to iterate over the raw logs and unpacked data for AccessRequest events raised by the Accessible contract.
type AccessibleAccessRequestIterator struct {
	Event *AccessibleAccessRequest // Event containing the contract specifics and raw log

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
func (it *AccessibleAccessRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessibleAccessRequest)
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
		it.Event = new(AccessibleAccessRequest)
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
func (it *AccessibleAccessRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessibleAccessRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessibleAccessRequest represents a AccessRequest event raised by the Accessible contract.
type AccessibleAccessRequest struct {
	RequestValidity *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccessRequest is a free log retrieval operation binding the contract event 0xca8116fd70ecea0cf35f7c08eb31491cce5ebd6bafef5bd6426862df29840463.
//
// Solidity: e AccessRequest(requestValidity uint256)
func (_Accessible *AccessibleFilterer) FilterAccessRequest(opts *bind.FilterOpts) (*AccessibleAccessRequestIterator, error) {

	logs, sub, err := _Accessible.contract.FilterLogs(opts, "AccessRequest")
	if err != nil {
		return nil, err
	}
	return &AccessibleAccessRequestIterator{contract: _Accessible.contract, event: "AccessRequest", logs: logs, sub: sub}, nil
}

// WatchAccessRequest is a free log subscription operation binding the contract event 0xca8116fd70ecea0cf35f7c08eb31491cce5ebd6bafef5bd6426862df29840463.
//
// Solidity: e AccessRequest(requestValidity uint256)
func (_Accessible *AccessibleFilterer) WatchAccessRequest(opts *bind.WatchOpts, sink chan<- *AccessibleAccessRequest) (event.Subscription, error) {

	logs, sub, err := _Accessible.contract.WatchLogs(opts, "AccessRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessibleAccessRequest)
				if err := _Accessible.contract.UnpackLog(event, "AccessRequest", log); err != nil {
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

// BaseAccessControlGroupABI is the input ABI used to generate the binding from.
const BaseAccessControlGroupABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"members\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"grantAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"hasManagerAccess\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"grantManagerAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"revokeAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"hasAccess\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"revokeManagerAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"managers\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"MemberAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"ManagerAccessGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"MemberRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"ManagerAccessRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"operationCode\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"UnauthorizedOperation\",\"type\":\"event\"}]"

// BaseAccessControlGroupBin is the compiled bytecode used for deploying new contracts.
const BaseAccessControlGroupBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a03328116600160a060020a0319928316811780855560018054909416909117835581168352600360209081526040808520805460ff19908116861790915585549093168552600290915290922080549092161790556107688061007e6000396000f3006060604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100c657806308ae4b0c146100f55780630ae5e7391461012857806341c0e1b51461014757806342e7ba7b1461015a5780636d2e4b1b1461017957806375861a951461019857806385e68531146101b75780638da5cb5b146101d657806395a078e8146101e9578063cdb849b714610208578063f2fde38b14610227578063fdff9b4d14610246575b005b34156100d157600080fd5b6100d9610265565b604051600160a060020a03909116815260200160405180910390f35b341561010057600080fd5b610114600160a060020a0360043516610274565b604051901515815260200160405180910390f35b341561013357600080fd5b6100c4600160a060020a0360043516610289565b341561015257600080fd5b6100c461035e565b341561016557600080fd5b610114600160a060020a0360043516610387565b341561018457600080fd5b6100c4600160a060020a03600435166103aa565b34156101a357600080fd5b6100c4600160a060020a036004351661044b565b34156101c257600080fd5b6100c4600160a060020a03600435166104cb565b34156101e157600080fd5b6100d96105b8565b34156101f457600080fd5b610114600160a060020a03600435166105c7565b341561021357600080fd5b6100c4600160a060020a03600435166105ea565b341561023257600080fd5b6100c4600160a060020a03600435166106c8565b341561025157600080fd5b610114600160a060020a0360043516610727565b600054600160a060020a031681565b60026020526000908152604090205460ff1681565b600160a060020a03331660009081526003602052604090205460ff1615156001141561031657600160a060020a03811660009081526002602052604090819020805460ff191660011790557fb251eb052afc73ffd02ffe85ad79990a8b3fed60d76dbc2fa2fdd7123dffd91490829051600160a060020a03909116815260200160405180910390a161035b565b7f23de2adc3e22f171f66b3e5a333e17feb9dc30ba9570933bd259cb6c13ef7ab7601482604051918252600160a060020a031660208201526040908101905180910390a15b50565b60015432600160a060020a0390811691161461037957600080fd5b600154600160a060020a0316ff5b600160a060020a031660009081526003602052604090205460ff16151560011490565b60005432600160a060020a039081169116146103c557600080fd5b600160a060020a03811615156103da57600080fd5b600054600154600160a060020a039081169116141561041c576001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60015432600160a060020a0390811691161461046657600080fd5b600160a060020a03811660009081526003602052604090819020805460ff191660011790557f93bcaab179551bde429187645251f8e1fb8ac85801fcb1cf91eb2c9043d6111790829051600160a060020a03909116815260200160405180910390a150565b600160a060020a03331660009081526003602052604090205460ff16151560011480610508575080600160a060020a031633600160a060020a0316145b1561057157600160a060020a03811660009081526002602052604090819020805460ff191690557f745cd29407db644ed93e3ceb61cbcab96d1dfb496989ac5d5bf514fc5a9fab9c90829051600160a060020a03909116815260200160405180910390a161035b565b7f23de2adc3e22f171f66b3e5a333e17feb9dc30ba9570933bd259cb6c13ef7ab7601582604051918252600160a060020a031660208201526040908101905180910390a150565b600154600160a060020a031681565b600160a060020a031660009081526002602052604090205460ff16151560011490565b60005433600160a060020a0390811691161480610618575080600160a060020a031633600160a060020a0316145b1561068157600160a060020a03811660009081526003602052604090819020805460ff191690557f2d6aa1a9629d125e23a0cf692cda7cd6795dff1652eedd4673b38ec31e387b9590829051600160a060020a03909116815260200160405180910390a161035b565b7f23de2adc3e22f171f66b3e5a333e17feb9dc30ba9570933bd259cb6c13ef7ab7600b82604051918252600160a060020a031660208201526040908101905180910390a150565b60015432600160a060020a039081169116146106e357600080fd5b600160a060020a03811615156106f857600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60036020526000908152604090205460ff16815600a165627a7a72305820140e959db4d9147a3e337f1842cf9a7aba7c25a281771dae84fa7dcc8ed5d81f0029`

// DeployBaseAccessControlGroup deploys a new Ethereum contract, binding an instance of BaseAccessControlGroup to it.
func DeployBaseAccessControlGroup(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BaseAccessControlGroup, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseAccessControlGroupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaseAccessControlGroupBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseAccessControlGroup{BaseAccessControlGroupCaller: BaseAccessControlGroupCaller{contract: contract}, BaseAccessControlGroupTransactor: BaseAccessControlGroupTransactor{contract: contract}, BaseAccessControlGroupFilterer: BaseAccessControlGroupFilterer{contract: contract}}, nil
}

// BaseAccessControlGroup is an auto generated Go binding around an Ethereum contract.
type BaseAccessControlGroup struct {
	BaseAccessControlGroupCaller     // Read-only binding to the contract
	BaseAccessControlGroupTransactor // Write-only binding to the contract
	BaseAccessControlGroupFilterer   // Log filterer for contract events
}

// BaseAccessControlGroupCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseAccessControlGroupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseAccessControlGroupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseAccessControlGroupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseAccessControlGroupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseAccessControlGroupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseAccessControlGroupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseAccessControlGroupSession struct {
	Contract     *BaseAccessControlGroup // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BaseAccessControlGroupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseAccessControlGroupCallerSession struct {
	Contract *BaseAccessControlGroupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// BaseAccessControlGroupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseAccessControlGroupTransactorSession struct {
	Contract     *BaseAccessControlGroupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// BaseAccessControlGroupRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseAccessControlGroupRaw struct {
	Contract *BaseAccessControlGroup // Generic contract binding to access the raw methods on
}

// BaseAccessControlGroupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseAccessControlGroupCallerRaw struct {
	Contract *BaseAccessControlGroupCaller // Generic read-only contract binding to access the raw methods on
}

// BaseAccessControlGroupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseAccessControlGroupTransactorRaw struct {
	Contract *BaseAccessControlGroupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseAccessControlGroup creates a new instance of BaseAccessControlGroup, bound to a specific deployed contract.
func NewBaseAccessControlGroup(address common.Address, backend bind.ContractBackend) (*BaseAccessControlGroup, error) {
	contract, err := bindBaseAccessControlGroup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseAccessControlGroup{BaseAccessControlGroupCaller: BaseAccessControlGroupCaller{contract: contract}, BaseAccessControlGroupTransactor: BaseAccessControlGroupTransactor{contract: contract}, BaseAccessControlGroupFilterer: BaseAccessControlGroupFilterer{contract: contract}}, nil
}

// NewBaseAccessControlGroupCaller creates a new read-only instance of BaseAccessControlGroup, bound to a specific deployed contract.
func NewBaseAccessControlGroupCaller(address common.Address, caller bind.ContractCaller) (*BaseAccessControlGroupCaller, error) {
	contract, err := bindBaseAccessControlGroup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseAccessControlGroupCaller{contract: contract}, nil
}

// NewBaseAccessControlGroupTransactor creates a new write-only instance of BaseAccessControlGroup, bound to a specific deployed contract.
func NewBaseAccessControlGroupTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseAccessControlGroupTransactor, error) {
	contract, err := bindBaseAccessControlGroup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseAccessControlGroupTransactor{contract: contract}, nil
}

// NewBaseAccessControlGroupFilterer creates a new log filterer instance of BaseAccessControlGroup, bound to a specific deployed contract.
func NewBaseAccessControlGroupFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseAccessControlGroupFilterer, error) {
	contract, err := bindBaseAccessControlGroup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseAccessControlGroupFilterer{contract: contract}, nil
}

// bindBaseAccessControlGroup binds a generic wrapper to an already deployed contract.
func bindBaseAccessControlGroup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseAccessControlGroupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseAccessControlGroup *BaseAccessControlGroupRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseAccessControlGroup.Contract.BaseAccessControlGroupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseAccessControlGroup *BaseAccessControlGroupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.BaseAccessControlGroupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseAccessControlGroup *BaseAccessControlGroupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.BaseAccessControlGroupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseAccessControlGroup *BaseAccessControlGroupCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseAccessControlGroup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.contract.Transact(opts, method, params...)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseAccessControlGroup *BaseAccessControlGroupCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseAccessControlGroup.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) Creator() (common.Address, error) {
	return _BaseAccessControlGroup.Contract.Creator(&_BaseAccessControlGroup.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseAccessControlGroup *BaseAccessControlGroupCallerSession) Creator() (common.Address, error) {
	return _BaseAccessControlGroup.Contract.Creator(&_BaseAccessControlGroup.CallOpts)
}

// HasAccess is a free data retrieval call binding the contract method 0x95a078e8.
//
// Solidity: function hasAccess(candidate address) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupCaller) HasAccess(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseAccessControlGroup.contract.Call(opts, out, "hasAccess", candidate)
	return *ret0, err
}

// HasAccess is a free data retrieval call binding the contract method 0x95a078e8.
//
// Solidity: function hasAccess(candidate address) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) HasAccess(candidate common.Address) (bool, error) {
	return _BaseAccessControlGroup.Contract.HasAccess(&_BaseAccessControlGroup.CallOpts, candidate)
}

// HasAccess is a free data retrieval call binding the contract method 0x95a078e8.
//
// Solidity: function hasAccess(candidate address) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupCallerSession) HasAccess(candidate common.Address) (bool, error) {
	return _BaseAccessControlGroup.Contract.HasAccess(&_BaseAccessControlGroup.CallOpts, candidate)
}

// HasManagerAccess is a free data retrieval call binding the contract method 0x42e7ba7b.
//
// Solidity: function hasManagerAccess(candidate address) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupCaller) HasManagerAccess(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseAccessControlGroup.contract.Call(opts, out, "hasManagerAccess", candidate)
	return *ret0, err
}

// HasManagerAccess is a free data retrieval call binding the contract method 0x42e7ba7b.
//
// Solidity: function hasManagerAccess(candidate address) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) HasManagerAccess(candidate common.Address) (bool, error) {
	return _BaseAccessControlGroup.Contract.HasManagerAccess(&_BaseAccessControlGroup.CallOpts, candidate)
}

// HasManagerAccess is a free data retrieval call binding the contract method 0x42e7ba7b.
//
// Solidity: function hasManagerAccess(candidate address) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupCallerSession) HasManagerAccess(candidate common.Address) (bool, error) {
	return _BaseAccessControlGroup.Contract.HasManagerAccess(&_BaseAccessControlGroup.CallOpts, candidate)
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers( address) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupCaller) Managers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseAccessControlGroup.contract.Call(opts, out, "managers", arg0)
	return *ret0, err
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers( address) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) Managers(arg0 common.Address) (bool, error) {
	return _BaseAccessControlGroup.Contract.Managers(&_BaseAccessControlGroup.CallOpts, arg0)
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers( address) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupCallerSession) Managers(arg0 common.Address) (bool, error) {
	return _BaseAccessControlGroup.Contract.Managers(&_BaseAccessControlGroup.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members( address) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupCaller) Members(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseAccessControlGroup.contract.Call(opts, out, "members", arg0)
	return *ret0, err
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members( address) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) Members(arg0 common.Address) (bool, error) {
	return _BaseAccessControlGroup.Contract.Members(&_BaseAccessControlGroup.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members( address) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupCallerSession) Members(arg0 common.Address) (bool, error) {
	return _BaseAccessControlGroup.Contract.Members(&_BaseAccessControlGroup.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseAccessControlGroup *BaseAccessControlGroupCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseAccessControlGroup.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) Owner() (common.Address, error) {
	return _BaseAccessControlGroup.Contract.Owner(&_BaseAccessControlGroup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseAccessControlGroup *BaseAccessControlGroupCallerSession) Owner() (common.Address, error) {
	return _BaseAccessControlGroup.Contract.Owner(&_BaseAccessControlGroup.CallOpts)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x0ae5e739.
//
// Solidity: function grantAccess(candidate address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactor) GrantAccess(opts *bind.TransactOpts, candidate common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.contract.Transact(opts, "grantAccess", candidate)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x0ae5e739.
//
// Solidity: function grantAccess(candidate address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) GrantAccess(candidate common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.GrantAccess(&_BaseAccessControlGroup.TransactOpts, candidate)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x0ae5e739.
//
// Solidity: function grantAccess(candidate address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactorSession) GrantAccess(candidate common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.GrantAccess(&_BaseAccessControlGroup.TransactOpts, candidate)
}

// GrantManagerAccess is a paid mutator transaction binding the contract method 0x75861a95.
//
// Solidity: function grantManagerAccess(manager address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactor) GrantManagerAccess(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.contract.Transact(opts, "grantManagerAccess", manager)
}

// GrantManagerAccess is a paid mutator transaction binding the contract method 0x75861a95.
//
// Solidity: function grantManagerAccess(manager address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) GrantManagerAccess(manager common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.GrantManagerAccess(&_BaseAccessControlGroup.TransactOpts, manager)
}

// GrantManagerAccess is a paid mutator transaction binding the contract method 0x75861a95.
//
// Solidity: function grantManagerAccess(manager address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactorSession) GrantManagerAccess(manager common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.GrantManagerAccess(&_BaseAccessControlGroup.TransactOpts, manager)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseAccessControlGroup.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) Kill() (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.Kill(&_BaseAccessControlGroup.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactorSession) Kill() (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.Kill(&_BaseAccessControlGroup.TransactOpts)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x85e68531.
//
// Solidity: function revokeAccess(candidate address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactor) RevokeAccess(opts *bind.TransactOpts, candidate common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.contract.Transact(opts, "revokeAccess", candidate)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x85e68531.
//
// Solidity: function revokeAccess(candidate address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) RevokeAccess(candidate common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.RevokeAccess(&_BaseAccessControlGroup.TransactOpts, candidate)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x85e68531.
//
// Solidity: function revokeAccess(candidate address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactorSession) RevokeAccess(candidate common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.RevokeAccess(&_BaseAccessControlGroup.TransactOpts, candidate)
}

// RevokeManagerAccess is a paid mutator transaction binding the contract method 0xcdb849b7.
//
// Solidity: function revokeManagerAccess(manager address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactor) RevokeManagerAccess(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.contract.Transact(opts, "revokeManagerAccess", manager)
}

// RevokeManagerAccess is a paid mutator transaction binding the contract method 0xcdb849b7.
//
// Solidity: function revokeManagerAccess(manager address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) RevokeManagerAccess(manager common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.RevokeManagerAccess(&_BaseAccessControlGroup.TransactOpts, manager)
}

// RevokeManagerAccess is a paid mutator transaction binding the contract method 0xcdb849b7.
//
// Solidity: function revokeManagerAccess(manager address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactorSession) RevokeManagerAccess(manager common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.RevokeManagerAccess(&_BaseAccessControlGroup.TransactOpts, manager)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.TransferCreatorship(&_BaseAccessControlGroup.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.TransferCreatorship(&_BaseAccessControlGroup.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.TransferOwnership(&_BaseAccessControlGroup.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.TransferOwnership(&_BaseAccessControlGroup.TransactOpts, newOwner)
}

// BaseAccessControlGroupManagerAccessGrantedIterator is returned from FilterManagerAccessGranted and is used to iterate over the raw logs and unpacked data for ManagerAccessGranted events raised by the BaseAccessControlGroup contract.
type BaseAccessControlGroupManagerAccessGrantedIterator struct {
	Event *BaseAccessControlGroupManagerAccessGranted // Event containing the contract specifics and raw log

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
func (it *BaseAccessControlGroupManagerAccessGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseAccessControlGroupManagerAccessGranted)
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
		it.Event = new(BaseAccessControlGroupManagerAccessGranted)
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
func (it *BaseAccessControlGroupManagerAccessGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseAccessControlGroupManagerAccessGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseAccessControlGroupManagerAccessGranted represents a ManagerAccessGranted event raised by the BaseAccessControlGroup contract.
type BaseAccessControlGroupManagerAccessGranted struct {
	Candidate common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterManagerAccessGranted is a free log retrieval operation binding the contract event 0x93bcaab179551bde429187645251f8e1fb8ac85801fcb1cf91eb2c9043d61117.
//
// Solidity: e ManagerAccessGranted(candidate address)
func (_BaseAccessControlGroup *BaseAccessControlGroupFilterer) FilterManagerAccessGranted(opts *bind.FilterOpts) (*BaseAccessControlGroupManagerAccessGrantedIterator, error) {

	logs, sub, err := _BaseAccessControlGroup.contract.FilterLogs(opts, "ManagerAccessGranted")
	if err != nil {
		return nil, err
	}
	return &BaseAccessControlGroupManagerAccessGrantedIterator{contract: _BaseAccessControlGroup.contract, event: "ManagerAccessGranted", logs: logs, sub: sub}, nil
}

// WatchManagerAccessGranted is a free log subscription operation binding the contract event 0x93bcaab179551bde429187645251f8e1fb8ac85801fcb1cf91eb2c9043d61117.
//
// Solidity: e ManagerAccessGranted(candidate address)
func (_BaseAccessControlGroup *BaseAccessControlGroupFilterer) WatchManagerAccessGranted(opts *bind.WatchOpts, sink chan<- *BaseAccessControlGroupManagerAccessGranted) (event.Subscription, error) {

	logs, sub, err := _BaseAccessControlGroup.contract.WatchLogs(opts, "ManagerAccessGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseAccessControlGroupManagerAccessGranted)
				if err := _BaseAccessControlGroup.contract.UnpackLog(event, "ManagerAccessGranted", log); err != nil {
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

// BaseAccessControlGroupManagerAccessRevokedIterator is returned from FilterManagerAccessRevoked and is used to iterate over the raw logs and unpacked data for ManagerAccessRevoked events raised by the BaseAccessControlGroup contract.
type BaseAccessControlGroupManagerAccessRevokedIterator struct {
	Event *BaseAccessControlGroupManagerAccessRevoked // Event containing the contract specifics and raw log

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
func (it *BaseAccessControlGroupManagerAccessRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseAccessControlGroupManagerAccessRevoked)
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
		it.Event = new(BaseAccessControlGroupManagerAccessRevoked)
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
func (it *BaseAccessControlGroupManagerAccessRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseAccessControlGroupManagerAccessRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseAccessControlGroupManagerAccessRevoked represents a ManagerAccessRevoked event raised by the BaseAccessControlGroup contract.
type BaseAccessControlGroupManagerAccessRevoked struct {
	Candidate common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterManagerAccessRevoked is a free log retrieval operation binding the contract event 0x2d6aa1a9629d125e23a0cf692cda7cd6795dff1652eedd4673b38ec31e387b95.
//
// Solidity: e ManagerAccessRevoked(candidate address)
func (_BaseAccessControlGroup *BaseAccessControlGroupFilterer) FilterManagerAccessRevoked(opts *bind.FilterOpts) (*BaseAccessControlGroupManagerAccessRevokedIterator, error) {

	logs, sub, err := _BaseAccessControlGroup.contract.FilterLogs(opts, "ManagerAccessRevoked")
	if err != nil {
		return nil, err
	}
	return &BaseAccessControlGroupManagerAccessRevokedIterator{contract: _BaseAccessControlGroup.contract, event: "ManagerAccessRevoked", logs: logs, sub: sub}, nil
}

// WatchManagerAccessRevoked is a free log subscription operation binding the contract event 0x2d6aa1a9629d125e23a0cf692cda7cd6795dff1652eedd4673b38ec31e387b95.
//
// Solidity: e ManagerAccessRevoked(candidate address)
func (_BaseAccessControlGroup *BaseAccessControlGroupFilterer) WatchManagerAccessRevoked(opts *bind.WatchOpts, sink chan<- *BaseAccessControlGroupManagerAccessRevoked) (event.Subscription, error) {

	logs, sub, err := _BaseAccessControlGroup.contract.WatchLogs(opts, "ManagerAccessRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseAccessControlGroupManagerAccessRevoked)
				if err := _BaseAccessControlGroup.contract.UnpackLog(event, "ManagerAccessRevoked", log); err != nil {
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

// BaseAccessControlGroupMemberAddedIterator is returned from FilterMemberAdded and is used to iterate over the raw logs and unpacked data for MemberAdded events raised by the BaseAccessControlGroup contract.
type BaseAccessControlGroupMemberAddedIterator struct {
	Event *BaseAccessControlGroupMemberAdded // Event containing the contract specifics and raw log

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
func (it *BaseAccessControlGroupMemberAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseAccessControlGroupMemberAdded)
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
		it.Event = new(BaseAccessControlGroupMemberAdded)
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
func (it *BaseAccessControlGroupMemberAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseAccessControlGroupMemberAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseAccessControlGroupMemberAdded represents a MemberAdded event raised by the BaseAccessControlGroup contract.
type BaseAccessControlGroupMemberAdded struct {
	Candidate common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMemberAdded is a free log retrieval operation binding the contract event 0xb251eb052afc73ffd02ffe85ad79990a8b3fed60d76dbc2fa2fdd7123dffd914.
//
// Solidity: e MemberAdded(candidate address)
func (_BaseAccessControlGroup *BaseAccessControlGroupFilterer) FilterMemberAdded(opts *bind.FilterOpts) (*BaseAccessControlGroupMemberAddedIterator, error) {

	logs, sub, err := _BaseAccessControlGroup.contract.FilterLogs(opts, "MemberAdded")
	if err != nil {
		return nil, err
	}
	return &BaseAccessControlGroupMemberAddedIterator{contract: _BaseAccessControlGroup.contract, event: "MemberAdded", logs: logs, sub: sub}, nil
}

// WatchMemberAdded is a free log subscription operation binding the contract event 0xb251eb052afc73ffd02ffe85ad79990a8b3fed60d76dbc2fa2fdd7123dffd914.
//
// Solidity: e MemberAdded(candidate address)
func (_BaseAccessControlGroup *BaseAccessControlGroupFilterer) WatchMemberAdded(opts *bind.WatchOpts, sink chan<- *BaseAccessControlGroupMemberAdded) (event.Subscription, error) {

	logs, sub, err := _BaseAccessControlGroup.contract.WatchLogs(opts, "MemberAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseAccessControlGroupMemberAdded)
				if err := _BaseAccessControlGroup.contract.UnpackLog(event, "MemberAdded", log); err != nil {
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

// BaseAccessControlGroupMemberRevokedIterator is returned from FilterMemberRevoked and is used to iterate over the raw logs and unpacked data for MemberRevoked events raised by the BaseAccessControlGroup contract.
type BaseAccessControlGroupMemberRevokedIterator struct {
	Event *BaseAccessControlGroupMemberRevoked // Event containing the contract specifics and raw log

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
func (it *BaseAccessControlGroupMemberRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseAccessControlGroupMemberRevoked)
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
		it.Event = new(BaseAccessControlGroupMemberRevoked)
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
func (it *BaseAccessControlGroupMemberRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseAccessControlGroupMemberRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseAccessControlGroupMemberRevoked represents a MemberRevoked event raised by the BaseAccessControlGroup contract.
type BaseAccessControlGroupMemberRevoked struct {
	Candidate common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMemberRevoked is a free log retrieval operation binding the contract event 0x745cd29407db644ed93e3ceb61cbcab96d1dfb496989ac5d5bf514fc5a9fab9c.
//
// Solidity: e MemberRevoked(candidate address)
func (_BaseAccessControlGroup *BaseAccessControlGroupFilterer) FilterMemberRevoked(opts *bind.FilterOpts) (*BaseAccessControlGroupMemberRevokedIterator, error) {

	logs, sub, err := _BaseAccessControlGroup.contract.FilterLogs(opts, "MemberRevoked")
	if err != nil {
		return nil, err
	}
	return &BaseAccessControlGroupMemberRevokedIterator{contract: _BaseAccessControlGroup.contract, event: "MemberRevoked", logs: logs, sub: sub}, nil
}

// WatchMemberRevoked is a free log subscription operation binding the contract event 0x745cd29407db644ed93e3ceb61cbcab96d1dfb496989ac5d5bf514fc5a9fab9c.
//
// Solidity: e MemberRevoked(candidate address)
func (_BaseAccessControlGroup *BaseAccessControlGroupFilterer) WatchMemberRevoked(opts *bind.WatchOpts, sink chan<- *BaseAccessControlGroupMemberRevoked) (event.Subscription, error) {

	logs, sub, err := _BaseAccessControlGroup.contract.WatchLogs(opts, "MemberRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseAccessControlGroupMemberRevoked)
				if err := _BaseAccessControlGroup.contract.UnpackLog(event, "MemberRevoked", log); err != nil {
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

// BaseAccessControlGroupUnauthorizedOperationIterator is returned from FilterUnauthorizedOperation and is used to iterate over the raw logs and unpacked data for UnauthorizedOperation events raised by the BaseAccessControlGroup contract.
type BaseAccessControlGroupUnauthorizedOperationIterator struct {
	Event *BaseAccessControlGroupUnauthorizedOperation // Event containing the contract specifics and raw log

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
func (it *BaseAccessControlGroupUnauthorizedOperationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseAccessControlGroupUnauthorizedOperation)
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
		it.Event = new(BaseAccessControlGroupUnauthorizedOperation)
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
func (it *BaseAccessControlGroupUnauthorizedOperationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseAccessControlGroupUnauthorizedOperationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseAccessControlGroupUnauthorizedOperation represents a UnauthorizedOperation event raised by the BaseAccessControlGroup contract.
type BaseAccessControlGroupUnauthorizedOperation struct {
	OperationCode *big.Int
	Candidate     common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnauthorizedOperation is a free log retrieval operation binding the contract event 0x23de2adc3e22f171f66b3e5a333e17feb9dc30ba9570933bd259cb6c13ef7ab7.
//
// Solidity: e UnauthorizedOperation(operationCode uint256, candidate address)
func (_BaseAccessControlGroup *BaseAccessControlGroupFilterer) FilterUnauthorizedOperation(opts *bind.FilterOpts) (*BaseAccessControlGroupUnauthorizedOperationIterator, error) {

	logs, sub, err := _BaseAccessControlGroup.contract.FilterLogs(opts, "UnauthorizedOperation")
	if err != nil {
		return nil, err
	}
	return &BaseAccessControlGroupUnauthorizedOperationIterator{contract: _BaseAccessControlGroup.contract, event: "UnauthorizedOperation", logs: logs, sub: sub}, nil
}

// WatchUnauthorizedOperation is a free log subscription operation binding the contract event 0x23de2adc3e22f171f66b3e5a333e17feb9dc30ba9570933bd259cb6c13ef7ab7.
//
// Solidity: e UnauthorizedOperation(operationCode uint256, candidate address)
func (_BaseAccessControlGroup *BaseAccessControlGroupFilterer) WatchUnauthorizedOperation(opts *bind.WatchOpts, sink chan<- *BaseAccessControlGroupUnauthorizedOperation) (event.Subscription, error) {

	logs, sub, err := _BaseAccessControlGroup.contract.WatchLogs(opts, "UnauthorizedOperation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseAccessControlGroupUnauthorizedOperation)
				if err := _BaseAccessControlGroup.contract.UnpackLog(event, "UnauthorizedOperation", log); err != nil {
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

// BaseContentABI is the input ABI used to generate the binding from.
const BaseContentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"licensingFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUnpaidLicensingFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestMap\",\"outputs\":[{\"name\":\"originator\",\"type\":\"address\"},{\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"statusCode\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressKMS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentType\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"status_code\",\"type\":\"int256\"}],\"name\":\"setStatusCode\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"percentComplete\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"request_ID\",\"type\":\"uint256\"},{\"name\":\"score_pct\",\"type\":\"uint256\"},{\"name\":\"ml_out_hash\",\"type\":\"bytes32\"}],\"name\":\"accessComplete\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accessCharge\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content_type\",\"type\":\"address\"}],\"name\":\"setContentType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"status_code\",\"type\":\"int256\"},{\"name\":\"percent_complete\",\"type\":\"uint8\"}],\"name\":\"updateStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"level\",\"type\":\"uint8\"},{\"name\":\"custom_values\",\"type\":\"bytes32[]\"},{\"name\":\"stakeholders\",\"type\":\"address[]\"}],\"name\":\"getAccessCharge\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requestID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"level\",\"type\":\"uint8\"},{\"name\":\"pke_requestor\",\"type\":\"string\"},{\"name\":\"pke_AFGH\",\"type\":\"string\"},{\"name\":\"custom_values\",\"type\":\"bytes32[]\"},{\"name\":\"stakeholders\",\"type\":\"address[]\"}],\"name\":\"accessRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"licensing_fee\",\"type\":\"uint256\"}],\"name\":\"setLicensingFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"libraryAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pct_complete\",\"type\":\"uint8\"},{\"name\":\"signed_verification\",\"type\":\"string\"}],\"name\":\"publish\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"address_KMS\",\"type\":\"address\"}],\"name\":\"setAddressKMS\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"objectHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setContentContractAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"request_ID\",\"type\":\"uint256\"},{\"name\":\"access_granted\",\"type\":\"bool\"},{\"name\":\"re_key\",\"type\":\"string\"},{\"name\":\"encrypted_AES_key\",\"type\":\"string\"}],\"name\":\"accessGrant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"object_hash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"charge\",\"type\":\"uint256\"}],\"name\":\"setAccessCharge\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount_to_be_paid\",\"type\":\"uint256\"}],\"name\":\"addLicensingFeeReceived\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"licensingFeeReceived\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"containingLibrary\",\"type\":\"address\"}],\"name\":\"ContentObjectCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentType\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"contentContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"licensingFee\",\"type\":\"uint256\"}],\"name\":\"SetContentType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestValidity\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"contentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"pkeRequestor\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"pkeAFGH\",\"type\":\"string\"}],\"name\":\"AccessRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"access_granted\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"reKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"encryptedAESKey\",\"type\":\"string\"}],\"name\":\"AccessGrant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"customValue\",\"type\":\"bytes32\"}],\"name\":\"AccessRequestValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"stakeholder\",\"type\":\"address\"}],\"name\":\"AccessRequestStakeholder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"scorePct\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"mlOutHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"customContractResult\",\"type\":\"bool\"}],\"name\":\"AccessComplete\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentContractAddress\",\"type\":\"address\"}],\"name\":\"SetContentContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accessCharge\",\"type\":\"uint256\"}],\"name\":\"SetAccessCharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"accessCharge\",\"type\":\"uint256\"}],\"name\":\"GetAccessCharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accessCharge\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amountProvided\",\"type\":\"uint256\"}],\"name\":\"InsufficientFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"statusCode\",\"type\":\"int256\"}],\"name\":\"SetStatusCode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pctComplete\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"requestStatus\",\"type\":\"bool\"}],\"name\":\"Publish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"custom_contract\",\"type\":\"address\"}],\"name\":\"InvokeCustomPreHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"custom_contract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"ReturnCustomHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"custom_contract\",\"type\":\"address\"}],\"name\":\"InvokeCustomPostHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"Commit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"UpdateRequest\",\"type\":\"event\"}]"

// BaseContentBin is the compiled bytecode used for deploying new contracts.
const BaseContentBin = `0x6060604090815260008054600160a060020a03328116600160a060020a03199283168117845560018054841690911790556006805433831693169290921791829055600019600855600b839055600c929092557fc3decc188980e855666b70498ca85e8fa284d97d30483d828fa126f7303d7d199291169051600160a060020a03909116815260200160405180910390a1611be08061009f6000396000f3006060604052600436106101875763ffffffff60e060020a60003504166302d05d3f81146101895780630779564b146101b857806307b27aca146101dd5780631a735f18146101f05780632310167f1461024257806327c1c21d1461025557806332eaf21b1461026857806336ebffca1461027b57806341c0e1b51461028e5780635267db44146102a1578063545b3539146102b75780635cc4aa9b146102e057806364ade32b1461030557806369348967146103185780636d2e4b1b14610337578063708f53ab14610356578063879fe48f146103725780638da5cb5b146104095780638f7792011461041c578063a1ff106e1461042f578063b65240891461053f578063b816f51314610555578063c287e0ed14610568578063c68c1d1f1461057b578063c9e8e72d146105c9578063e02dd9c2146105e8578063e5385303146105fb578063ee56d7671461061a578063f14fcbc8146106ba578063f2fde38b146106d0578063f4d9bae8146106ef578063f794447114610705578063f94f00141461071b575b005b341561019457600080fd5b61019c61072e565b604051600160a060020a03909116815260200160405180910390f35b34156101c357600080fd5b6101cb61073d565b60405190815260200160405180910390f35b34156101e857600080fd5b6101cb610743565b34156101fb57600080fd5b61020660043561074d565b6040518084600160a060020a0316600160a060020a031681526020018381526020018260000b60000b8152602001935050505060405180910390f35b341561024d57600080fd5b61019c610778565b341561026057600080fd5b6101cb610787565b341561027357600080fd5b61019c61078d565b341561028657600080fd5b61019c61079c565b341561029957600080fd5b6101876107ab565b34156102ac57600080fd5b6101cb6004356107d4565b34156102c257600080fd5b6102ca610873565b60405160ff909116815260200160405180910390f35b6102f160043560243560443561087c565b604051901515815260200160405180910390f35b341561031057600080fd5b6101cb610a12565b341561032357600080fd5b610187600160a060020a0360043516610a18565b341561034257600080fd5b610187600160a060020a0360043516610c00565b341561036157600080fd5b6101cb60043560ff60243516610c87565b341561037d57600080fd5b6101cb6004803560ff1690604460248035908101908301358060208082020160405190810160405280939291908181526020018383602002808284378201915050505050509190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843750949650610dc095505050505050565b341561041457600080fd5b61019c610f28565b341561042757600080fd5b6101cb610f37565b6102f16004803560ff169060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843750949650610f3d95505050505050565b341561054a57600080fd5b610187600435611503565b341561056057600080fd5b61019c611531565b341561057357600080fd5b610187611540565b6102f16004803560ff169060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061159295505050505050565b34156105d457600080fd5b610187600160a060020a03600435166116d5565b34156105f357600080fd5b6101cb611712565b341561060657600080fd5b610187600160a060020a0360043516611718565b341561062557600080fd5b6102f16004803590602480351515919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061179595505050505050565b34156106c557600080fd5b610187600435611a37565b34156106db57600080fd5b610187600160a060020a0360043516611a8d565b34156106fa57600080fd5b6101cb600435611adf565b341561071057600080fd5b6101cb600435611b3d565b341561072657600080fd5b6101cb611b6e565b600054600160a060020a031681565b600a5481565b600b54600a540390565b600d60205260009081526040812080546001820154600290920154600160a060020a03909116920b83565b600554600160a060020a031681565b60085481565b600454600160a060020a031681565b600354600160a060020a031681565b60015432600160a060020a039081169116146107c657600080fd5b600154600160a060020a0316ff5b60015460009032600160a060020a03908116911614801561080c5750600082128061080c575060008213801561080c57506000600854125b156108175760088290555b60065433600160a060020a03908116911614156108345760088290555b7fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a260085460405190815260200160405180910390a1506008545b919050565b60095460ff1681565b60055460009060019082908190600160a060020a0316156108fb57600554600160a060020a0316915081636b2d13248860405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156108e257600080fd5b5af115156108ef57600080fd5b50505060405180519350505b506000868152600d602052604090208054600160a060020a0316158015906109305750805433600160a060020a039081169116145b156109b2576002810154600090810b900b151561097e5733600160a060020a03166108fc82600101549081150290604051600060405180830381858888f19350505050151561097e57600080fd5b6000878152600d602052604081208054600160a060020a03191681556001810191909155600201805460ff191690556109b7565b600092505b7f7f1f4b28434ce7beab4983e64a8b5bb96e195a67029fdaff925028aec57fbc6b87878786604051938452602084019290925260408084019190915290151560608301526080909101905180910390a1509095945050505050565b60075481565b60015460009032600160a060020a03908116911614610a3657600080fd5b5060038054600160a060020a031916600160a060020a03838116919091179091556006541680631cdbee5a8360405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610aa057600080fd5b5af11515610aad57600080fd5b505050604051805160058054600160a060020a031916600160a060020a0392831617905582169050633fbb5d718360405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610b1957600080fd5b5af11515610b2657600080fd5b5050506040518051600a5550600160a060020a0381166332eaf21b6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610b6f57600080fd5b5af11515610b7c57600080fd5b505050604051805160048054600160a060020a031916600160a060020a03928316179055600554600a547f1c158b764edca9bf03a401a388355220a7395febb3c011729d8060572f065627935085929190911690604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15050565b60005432600160a060020a03908116911614610c1b57600080fd5b600160a060020a0381161515610c3057600080fd5b600054600154600160a060020a0390811691161415610c655760018054600160a060020a031916600160a060020a0383161790555b60008054600160a060020a031916600160a060020a0392909216919091179055565b600154600090819081908190819032600160a060020a0390811691161480610cbd575060065433600160a060020a039081169116145b15610db1576009805460ff191660ff881617905560055460009450600160a060020a03161515610d115760095487935060ff166064148015610cfd575086155b15610d0c57600b54600a540393505b610d9f565b600554600160a060020a0316915081633513a8058860405160e060020a63ffffffff841602815260048101919091526024016040805180830381600087803b1515610d5b57600080fd5b5af11515610d6857600080fd5b5050506040518051906020018051919450909150506000811215610d9b57861515610d0c57600b54600a54039350610d9f565b8093505b610da8836107d4565b50839450610db6565b600094505b5050505092915050565b6007546005546000919082908190600160a060020a031615610edf57600554600160a060020a031691508163f8117ca28888886040518463ffffffff1660e060020a028152600401808460ff1660ff1681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015610e55578082015183820152602001610e3d565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610e94578082015183820152602001610e7c565b5050505090500195505050505050602060405180830381600087803b1515610ebb57600080fd5b5af11515610ec857600080fd5b505050604051805191505060008112610edf578092505b7fa58326ee5bb617cb8b4f0d0f5f557c469d2d05d7a738f777037deda9c724b370878460405160ff909216825260208201526040908101905180910390a1509095945050505050565b600154600160a060020a031681565b600c5481565b6000806000610f4a611b74565b600c80546001019055600654600160a060020a0316925060008080856395a078e83260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610faa57600080fd5b5af11515610fb757600080fd5b50505060405180511515905061103457600080516020611b958339815191526069600c548e6000604051938452602084019290925260ff16604080840191909152606083019190915260c060808301819052600090830181905261010060a08401819052830152610140909101905180910390a1600096506114f4565b60015432600160a060020a039081169116146110cc576110558c8a8a610dc0565b9450843410156110cc57600080516020611b958339815191526067600c548e6000604051938452602084019290925260ff16604080840191909152606083019190915260c060808301819052600090830181905261010060a08401819052830152610140909101905180910390a1600096506114f4565b60606040519081016040908152600160a060020a0333168252346020808401919091526000828401819052600c548152600d90915220909450849081518154600160a060020a031916600160a060020a03919091161781556020820151816001015560408201516002909101805460009290920b60ff1660ff1990921691909117905550600554600160a060020a0316156112d857600554600c54600160a060020a039091169350839063123e0e80908e8c8c6040518563ffffffff1660e060020a028152600401808581526020018460ff1660ff1681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156111e75780820151838201526020016111cf565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561122657808201518382015260200161120e565b505050509050019650505050505050602060405180830381600087803b151561124e57600080fd5b5af1151561125b57600080fd5b505050604051805192505081156112d857600080516020611b9583398151915282600c548e6000604051938452602084019290925260ff16604080840191909152606083019190915260c060808301819052600090830181905261010060a08401819052830152610140909101905180910390a1600096506114f4565b600080516020611b958339815191526000600c548e6002548f8f6040518681526020810186905260ff851660408201526060810184905260c0608082018181529060a0830190830185818151815260200191508051906020019080838360005b83811015611350578082015183820152602001611338565b50505050905090810190601f16801561137d5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156113b357808201518382015260200161139b565b50505050905090810190601f1680156113e05780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390a15060005b885181101561146c5788818151811061140c57fe5b9060200190602002015115611464577f515e0a48b385fce2a8e4d9f169a97c4f6ea669a752358f5e6ab37cc3c2e84c3889828151811061144857fe5b9060200190602002015160405190815260200160405180910390a15b6001016113f7565b5060005b87518110156114ef5788818151811061148557fe5b90602001906020020151156114e7577fb6e3239e521a6c66920ae634f8e921a37e6991d520ac44d52f8516397f41b6848882815181106114c157fe5b90602001906020020151604051600160a060020a03909116815260200160405180910390a15b600101611470565b600196505b50505050505095945050505050565b60015432600160a060020a03908116911614801561152357506000600854125b1561152e57600a8190555b50565b600654600160a060020a031681565b60015432600160a060020a0390811691161461155b57600080fd5b7f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef1960025460405190815260200160405180910390a1565b6001546000908190819032600160a060020a039081169116146115b457600080fd5b836040518082805190602001908083835b602083106115e45780518252601f1990920191602091820191016115c5565b6001836020036101000a0380198251168184511617909252505050919091019250604091505051506116199050600186610c87565b50600060085413156116c857600654600160a060020a03169150816349102e616040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561166757600080fd5b5af1151561167457600080fd5b5050506040518051905090507f9625d8ae00bbc11bf4dbac17532594047d634ea66e1baa95d4974d086f261c30858260405160ff9092168252151560208201526040908101905180910390a18092506116cd565b600092505b505092915050565b60015432600160a060020a039081169116146116f057600080fd5b60048054600160a060020a031916600160a060020a0392909216919091179055565b60025481565b60015432600160a060020a0390811691161461173357600080fd5b60058054600160a060020a031916600160a060020a0383811691909117918290557fa6f2e38f0cfebf27212317fced3ac40bc62e00bd33f38d69603710740c69acb79116604051600160a060020a03909116815260200160405180910390a150565b600154600090819033600160a060020a039081169116148015906117c8575060045433600160a060020a03908116911614155b156117d65760009150611a2e565b506000858152600d602052604090208054600160a060020a031615156117ff5760009150611a2e565b6002810154600090810b900b1515611a29578415156118c55780546001820154600160a060020a039091169080156108fc0290604051600060405180830381858888f19350505050151561185257600080fd5b60028101805460ff191660ff1790557f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b9127188660006040519182521515602082015260806040808301829052600091830182905260c060608401819052830191909152610100909101905180910390a1611a29565b6001805490820154600160a060020a039091169080156108fc0290604051600060405180830381858888f19350505050151561190057600080fd5b60028101805460ff191660019081179091557f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b91271890879086866040518481528315156020820152608060408201818152906060830190830185818151815260200191508051906020019080838360005b8381101561198757808201518382015260200161196f565b50505050905090810190601f1680156119b45780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156119ea5780820151838201526020016119d2565b50505050905090810190601f168015611a175780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a15b600191505b50949350505050565b60015432600160a060020a03908116911614611a5257600080fd5b60028190557f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff8160405190815260200160405180910390a150565b60015432600160a060020a03908116911614611aa857600080fd5b600160a060020a0381161515611abd57600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b60015460009032600160a060020a03908116911614611afd57600080fd5b60078290557f4114f8ef80b6de2161db580cbefa14e1892d15d3ebe2062c9914e4a5773114a38260405190815260200160405180910390a1505060075490565b60065460009033600160a060020a03908116911614611b5e5750600061086e565b50600b8054820190819055919050565b600b5481565b6060604051908101604090815260008083526020830181905290820152905600272aa3d0efd05809b06d40e95792681d6b844a29fe97cd703cfd5eee27510740a165627a7a72305820f010d72b4ef3212cb91349cd2c5958f806c34e964fa1b0e1325044ebfbeef0500029`

// DeployBaseContent deploys a new Ethereum contract, binding an instance of BaseContent to it.
func DeployBaseContent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BaseContent, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseContentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaseContentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseContent{BaseContentCaller: BaseContentCaller{contract: contract}, BaseContentTransactor: BaseContentTransactor{contract: contract}, BaseContentFilterer: BaseContentFilterer{contract: contract}}, nil
}

// BaseContent is an auto generated Go binding around an Ethereum contract.
type BaseContent struct {
	BaseContentCaller     // Read-only binding to the contract
	BaseContentTransactor // Write-only binding to the contract
	BaseContentFilterer   // Log filterer for contract events
}

// BaseContentCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseContentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseContentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseContentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseContentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseContentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseContentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseContentSession struct {
	Contract     *BaseContent      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseContentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseContentCallerSession struct {
	Contract *BaseContentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BaseContentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseContentTransactorSession struct {
	Contract     *BaseContentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BaseContentRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseContentRaw struct {
	Contract *BaseContent // Generic contract binding to access the raw methods on
}

// BaseContentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseContentCallerRaw struct {
	Contract *BaseContentCaller // Generic read-only contract binding to access the raw methods on
}

// BaseContentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseContentTransactorRaw struct {
	Contract *BaseContentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseContent creates a new instance of BaseContent, bound to a specific deployed contract.
func NewBaseContent(address common.Address, backend bind.ContractBackend) (*BaseContent, error) {
	contract, err := bindBaseContent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseContent{BaseContentCaller: BaseContentCaller{contract: contract}, BaseContentTransactor: BaseContentTransactor{contract: contract}, BaseContentFilterer: BaseContentFilterer{contract: contract}}, nil
}

// NewBaseContentCaller creates a new read-only instance of BaseContent, bound to a specific deployed contract.
func NewBaseContentCaller(address common.Address, caller bind.ContractCaller) (*BaseContentCaller, error) {
	contract, err := bindBaseContent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseContentCaller{contract: contract}, nil
}

// NewBaseContentTransactor creates a new write-only instance of BaseContent, bound to a specific deployed contract.
func NewBaseContentTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseContentTransactor, error) {
	contract, err := bindBaseContent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseContentTransactor{contract: contract}, nil
}

// NewBaseContentFilterer creates a new log filterer instance of BaseContent, bound to a specific deployed contract.
func NewBaseContentFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseContentFilterer, error) {
	contract, err := bindBaseContent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseContentFilterer{contract: contract}, nil
}

// bindBaseContent binds a generic wrapper to an already deployed contract.
func bindBaseContent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseContentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseContent *BaseContentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseContent.Contract.BaseContentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseContent *BaseContentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContent.Contract.BaseContentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseContent *BaseContentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseContent.Contract.BaseContentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseContent *BaseContentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseContent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseContent *BaseContentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseContent *BaseContentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseContent.Contract.contract.Transact(opts, method, params...)
}

// AccessCharge is a free data retrieval call binding the contract method 0x64ade32b.
//
// Solidity: function accessCharge() constant returns(uint256)
func (_BaseContent *BaseContentCaller) AccessCharge(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "accessCharge")
	return *ret0, err
}

// AccessCharge is a free data retrieval call binding the contract method 0x64ade32b.
//
// Solidity: function accessCharge() constant returns(uint256)
func (_BaseContent *BaseContentSession) AccessCharge() (*big.Int, error) {
	return _BaseContent.Contract.AccessCharge(&_BaseContent.CallOpts)
}

// AccessCharge is a free data retrieval call binding the contract method 0x64ade32b.
//
// Solidity: function accessCharge() constant returns(uint256)
func (_BaseContent *BaseContentCallerSession) AccessCharge() (*big.Int, error) {
	return _BaseContent.Contract.AccessCharge(&_BaseContent.CallOpts)
}

// AddressKMS is a free data retrieval call binding the contract method 0x32eaf21b.
//
// Solidity: function addressKMS() constant returns(address)
func (_BaseContent *BaseContentCaller) AddressKMS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "addressKMS")
	return *ret0, err
}

// AddressKMS is a free data retrieval call binding the contract method 0x32eaf21b.
//
// Solidity: function addressKMS() constant returns(address)
func (_BaseContent *BaseContentSession) AddressKMS() (common.Address, error) {
	return _BaseContent.Contract.AddressKMS(&_BaseContent.CallOpts)
}

// AddressKMS is a free data retrieval call binding the contract method 0x32eaf21b.
//
// Solidity: function addressKMS() constant returns(address)
func (_BaseContent *BaseContentCallerSession) AddressKMS() (common.Address, error) {
	return _BaseContent.Contract.AddressKMS(&_BaseContent.CallOpts)
}

// ContentContractAddress is a free data retrieval call binding the contract method 0x2310167f.
//
// Solidity: function contentContractAddress() constant returns(address)
func (_BaseContent *BaseContentCaller) ContentContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "contentContractAddress")
	return *ret0, err
}

// ContentContractAddress is a free data retrieval call binding the contract method 0x2310167f.
//
// Solidity: function contentContractAddress() constant returns(address)
func (_BaseContent *BaseContentSession) ContentContractAddress() (common.Address, error) {
	return _BaseContent.Contract.ContentContractAddress(&_BaseContent.CallOpts)
}

// ContentContractAddress is a free data retrieval call binding the contract method 0x2310167f.
//
// Solidity: function contentContractAddress() constant returns(address)
func (_BaseContent *BaseContentCallerSession) ContentContractAddress() (common.Address, error) {
	return _BaseContent.Contract.ContentContractAddress(&_BaseContent.CallOpts)
}

// ContentType is a free data retrieval call binding the contract method 0x36ebffca.
//
// Solidity: function contentType() constant returns(address)
func (_BaseContent *BaseContentCaller) ContentType(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "contentType")
	return *ret0, err
}

// ContentType is a free data retrieval call binding the contract method 0x36ebffca.
//
// Solidity: function contentType() constant returns(address)
func (_BaseContent *BaseContentSession) ContentType() (common.Address, error) {
	return _BaseContent.Contract.ContentType(&_BaseContent.CallOpts)
}

// ContentType is a free data retrieval call binding the contract method 0x36ebffca.
//
// Solidity: function contentType() constant returns(address)
func (_BaseContent *BaseContentCallerSession) ContentType() (common.Address, error) {
	return _BaseContent.Contract.ContentType(&_BaseContent.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseContent *BaseContentCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseContent *BaseContentSession) Creator() (common.Address, error) {
	return _BaseContent.Contract.Creator(&_BaseContent.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseContent *BaseContentCallerSession) Creator() (common.Address, error) {
	return _BaseContent.Contract.Creator(&_BaseContent.CallOpts)
}

// GetUnpaidLicensingFee is a free data retrieval call binding the contract method 0x07b27aca.
//
// Solidity: function getUnpaidLicensingFee() constant returns(uint256)
func (_BaseContent *BaseContentCaller) GetUnpaidLicensingFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "getUnpaidLicensingFee")
	return *ret0, err
}

// GetUnpaidLicensingFee is a free data retrieval call binding the contract method 0x07b27aca.
//
// Solidity: function getUnpaidLicensingFee() constant returns(uint256)
func (_BaseContent *BaseContentSession) GetUnpaidLicensingFee() (*big.Int, error) {
	return _BaseContent.Contract.GetUnpaidLicensingFee(&_BaseContent.CallOpts)
}

// GetUnpaidLicensingFee is a free data retrieval call binding the contract method 0x07b27aca.
//
// Solidity: function getUnpaidLicensingFee() constant returns(uint256)
func (_BaseContent *BaseContentCallerSession) GetUnpaidLicensingFee() (*big.Int, error) {
	return _BaseContent.Contract.GetUnpaidLicensingFee(&_BaseContent.CallOpts)
}

// LibraryAddress is a free data retrieval call binding the contract method 0xb816f513.
//
// Solidity: function libraryAddress() constant returns(address)
func (_BaseContent *BaseContentCaller) LibraryAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "libraryAddress")
	return *ret0, err
}

// LibraryAddress is a free data retrieval call binding the contract method 0xb816f513.
//
// Solidity: function libraryAddress() constant returns(address)
func (_BaseContent *BaseContentSession) LibraryAddress() (common.Address, error) {
	return _BaseContent.Contract.LibraryAddress(&_BaseContent.CallOpts)
}

// LibraryAddress is a free data retrieval call binding the contract method 0xb816f513.
//
// Solidity: function libraryAddress() constant returns(address)
func (_BaseContent *BaseContentCallerSession) LibraryAddress() (common.Address, error) {
	return _BaseContent.Contract.LibraryAddress(&_BaseContent.CallOpts)
}

// LicensingFee is a free data retrieval call binding the contract method 0x0779564b.
//
// Solidity: function licensingFee() constant returns(uint256)
func (_BaseContent *BaseContentCaller) LicensingFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "licensingFee")
	return *ret0, err
}

// LicensingFee is a free data retrieval call binding the contract method 0x0779564b.
//
// Solidity: function licensingFee() constant returns(uint256)
func (_BaseContent *BaseContentSession) LicensingFee() (*big.Int, error) {
	return _BaseContent.Contract.LicensingFee(&_BaseContent.CallOpts)
}

// LicensingFee is a free data retrieval call binding the contract method 0x0779564b.
//
// Solidity: function licensingFee() constant returns(uint256)
func (_BaseContent *BaseContentCallerSession) LicensingFee() (*big.Int, error) {
	return _BaseContent.Contract.LicensingFee(&_BaseContent.CallOpts)
}

// LicensingFeeReceived is a free data retrieval call binding the contract method 0xf94f0014.
//
// Solidity: function licensingFeeReceived() constant returns(uint256)
func (_BaseContent *BaseContentCaller) LicensingFeeReceived(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "licensingFeeReceived")
	return *ret0, err
}

// LicensingFeeReceived is a free data retrieval call binding the contract method 0xf94f0014.
//
// Solidity: function licensingFeeReceived() constant returns(uint256)
func (_BaseContent *BaseContentSession) LicensingFeeReceived() (*big.Int, error) {
	return _BaseContent.Contract.LicensingFeeReceived(&_BaseContent.CallOpts)
}

// LicensingFeeReceived is a free data retrieval call binding the contract method 0xf94f0014.
//
// Solidity: function licensingFeeReceived() constant returns(uint256)
func (_BaseContent *BaseContentCallerSession) LicensingFeeReceived() (*big.Int, error) {
	return _BaseContent.Contract.LicensingFeeReceived(&_BaseContent.CallOpts)
}

// ObjectHash is a free data retrieval call binding the contract method 0xe02dd9c2.
//
// Solidity: function objectHash() constant returns(bytes32)
func (_BaseContent *BaseContentCaller) ObjectHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "objectHash")
	return *ret0, err
}

// ObjectHash is a free data retrieval call binding the contract method 0xe02dd9c2.
//
// Solidity: function objectHash() constant returns(bytes32)
func (_BaseContent *BaseContentSession) ObjectHash() ([32]byte, error) {
	return _BaseContent.Contract.ObjectHash(&_BaseContent.CallOpts)
}

// ObjectHash is a free data retrieval call binding the contract method 0xe02dd9c2.
//
// Solidity: function objectHash() constant returns(bytes32)
func (_BaseContent *BaseContentCallerSession) ObjectHash() ([32]byte, error) {
	return _BaseContent.Contract.ObjectHash(&_BaseContent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseContent *BaseContentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseContent *BaseContentSession) Owner() (common.Address, error) {
	return _BaseContent.Contract.Owner(&_BaseContent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseContent *BaseContentCallerSession) Owner() (common.Address, error) {
	return _BaseContent.Contract.Owner(&_BaseContent.CallOpts)
}

// PercentComplete is a free data retrieval call binding the contract method 0x545b3539.
//
// Solidity: function percentComplete() constant returns(uint8)
func (_BaseContent *BaseContentCaller) PercentComplete(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "percentComplete")
	return *ret0, err
}

// PercentComplete is a free data retrieval call binding the contract method 0x545b3539.
//
// Solidity: function percentComplete() constant returns(uint8)
func (_BaseContent *BaseContentSession) PercentComplete() (uint8, error) {
	return _BaseContent.Contract.PercentComplete(&_BaseContent.CallOpts)
}

// PercentComplete is a free data retrieval call binding the contract method 0x545b3539.
//
// Solidity: function percentComplete() constant returns(uint8)
func (_BaseContent *BaseContentCallerSession) PercentComplete() (uint8, error) {
	return _BaseContent.Contract.PercentComplete(&_BaseContent.CallOpts)
}

// RequestID is a free data retrieval call binding the contract method 0x8f779201.
//
// Solidity: function requestID() constant returns(uint256)
func (_BaseContent *BaseContentCaller) RequestID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "requestID")
	return *ret0, err
}

// RequestID is a free data retrieval call binding the contract method 0x8f779201.
//
// Solidity: function requestID() constant returns(uint256)
func (_BaseContent *BaseContentSession) RequestID() (*big.Int, error) {
	return _BaseContent.Contract.RequestID(&_BaseContent.CallOpts)
}

// RequestID is a free data retrieval call binding the contract method 0x8f779201.
//
// Solidity: function requestID() constant returns(uint256)
func (_BaseContent *BaseContentCallerSession) RequestID() (*big.Int, error) {
	return _BaseContent.Contract.RequestID(&_BaseContent.CallOpts)
}

// RequestMap is a free data retrieval call binding the contract method 0x1a735f18.
//
// Solidity: function requestMap( uint256) constant returns(originator address, amountPaid uint256, status int8)
func (_BaseContent *BaseContentCaller) RequestMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Originator common.Address
	AmountPaid *big.Int
	Status     int8
}, error) {
	ret := new(struct {
		Originator common.Address
		AmountPaid *big.Int
		Status     int8
	})
	out := ret
	err := _BaseContent.contract.Call(opts, out, "requestMap", arg0)
	return *ret, err
}

// RequestMap is a free data retrieval call binding the contract method 0x1a735f18.
//
// Solidity: function requestMap( uint256) constant returns(originator address, amountPaid uint256, status int8)
func (_BaseContent *BaseContentSession) RequestMap(arg0 *big.Int) (struct {
	Originator common.Address
	AmountPaid *big.Int
	Status     int8
}, error) {
	return _BaseContent.Contract.RequestMap(&_BaseContent.CallOpts, arg0)
}

// RequestMap is a free data retrieval call binding the contract method 0x1a735f18.
//
// Solidity: function requestMap( uint256) constant returns(originator address, amountPaid uint256, status int8)
func (_BaseContent *BaseContentCallerSession) RequestMap(arg0 *big.Int) (struct {
	Originator common.Address
	AmountPaid *big.Int
	Status     int8
}, error) {
	return _BaseContent.Contract.RequestMap(&_BaseContent.CallOpts, arg0)
}

// StatusCode is a free data retrieval call binding the contract method 0x27c1c21d.
//
// Solidity: function statusCode() constant returns(int256)
func (_BaseContent *BaseContentCaller) StatusCode(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "statusCode")
	return *ret0, err
}

// StatusCode is a free data retrieval call binding the contract method 0x27c1c21d.
//
// Solidity: function statusCode() constant returns(int256)
func (_BaseContent *BaseContentSession) StatusCode() (*big.Int, error) {
	return _BaseContent.Contract.StatusCode(&_BaseContent.CallOpts)
}

// StatusCode is a free data retrieval call binding the contract method 0x27c1c21d.
//
// Solidity: function statusCode() constant returns(int256)
func (_BaseContent *BaseContentCallerSession) StatusCode() (*big.Int, error) {
	return _BaseContent.Contract.StatusCode(&_BaseContent.CallOpts)
}

// AccessComplete is a paid mutator transaction binding the contract method 0x5cc4aa9b.
//
// Solidity: function accessComplete(request_ID uint256, score_pct uint256, ml_out_hash bytes32) returns(bool)
func (_BaseContent *BaseContentTransactor) AccessComplete(opts *bind.TransactOpts, request_ID *big.Int, score_pct *big.Int, ml_out_hash [32]byte) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "accessComplete", request_ID, score_pct, ml_out_hash)
}

// AccessComplete is a paid mutator transaction binding the contract method 0x5cc4aa9b.
//
// Solidity: function accessComplete(request_ID uint256, score_pct uint256, ml_out_hash bytes32) returns(bool)
func (_BaseContent *BaseContentSession) AccessComplete(request_ID *big.Int, score_pct *big.Int, ml_out_hash [32]byte) (*types.Transaction, error) {
	return _BaseContent.Contract.AccessComplete(&_BaseContent.TransactOpts, request_ID, score_pct, ml_out_hash)
}

// AccessComplete is a paid mutator transaction binding the contract method 0x5cc4aa9b.
//
// Solidity: function accessComplete(request_ID uint256, score_pct uint256, ml_out_hash bytes32) returns(bool)
func (_BaseContent *BaseContentTransactorSession) AccessComplete(request_ID *big.Int, score_pct *big.Int, ml_out_hash [32]byte) (*types.Transaction, error) {
	return _BaseContent.Contract.AccessComplete(&_BaseContent.TransactOpts, request_ID, score_pct, ml_out_hash)
}

// AccessGrant is a paid mutator transaction binding the contract method 0xee56d767.
//
// Solidity: function accessGrant(request_ID uint256, access_granted bool, re_key string, encrypted_AES_key string) returns(bool)
func (_BaseContent *BaseContentTransactor) AccessGrant(opts *bind.TransactOpts, request_ID *big.Int, access_granted bool, re_key string, encrypted_AES_key string) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "accessGrant", request_ID, access_granted, re_key, encrypted_AES_key)
}

// AccessGrant is a paid mutator transaction binding the contract method 0xee56d767.
//
// Solidity: function accessGrant(request_ID uint256, access_granted bool, re_key string, encrypted_AES_key string) returns(bool)
func (_BaseContent *BaseContentSession) AccessGrant(request_ID *big.Int, access_granted bool, re_key string, encrypted_AES_key string) (*types.Transaction, error) {
	return _BaseContent.Contract.AccessGrant(&_BaseContent.TransactOpts, request_ID, access_granted, re_key, encrypted_AES_key)
}

// AccessGrant is a paid mutator transaction binding the contract method 0xee56d767.
//
// Solidity: function accessGrant(request_ID uint256, access_granted bool, re_key string, encrypted_AES_key string) returns(bool)
func (_BaseContent *BaseContentTransactorSession) AccessGrant(request_ID *big.Int, access_granted bool, re_key string, encrypted_AES_key string) (*types.Transaction, error) {
	return _BaseContent.Contract.AccessGrant(&_BaseContent.TransactOpts, request_ID, access_granted, re_key, encrypted_AES_key)
}

// AccessRequest is a paid mutator transaction binding the contract method 0xa1ff106e.
//
// Solidity: function accessRequest(level uint8, pke_requestor string, pke_AFGH string, custom_values bytes32[], stakeholders address[]) returns(bool)
func (_BaseContent *BaseContentTransactor) AccessRequest(opts *bind.TransactOpts, level uint8, pke_requestor string, pke_AFGH string, custom_values [][32]byte, stakeholders []common.Address) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "accessRequest", level, pke_requestor, pke_AFGH, custom_values, stakeholders)
}

// AccessRequest is a paid mutator transaction binding the contract method 0xa1ff106e.
//
// Solidity: function accessRequest(level uint8, pke_requestor string, pke_AFGH string, custom_values bytes32[], stakeholders address[]) returns(bool)
func (_BaseContent *BaseContentSession) AccessRequest(level uint8, pke_requestor string, pke_AFGH string, custom_values [][32]byte, stakeholders []common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.AccessRequest(&_BaseContent.TransactOpts, level, pke_requestor, pke_AFGH, custom_values, stakeholders)
}

// AccessRequest is a paid mutator transaction binding the contract method 0xa1ff106e.
//
// Solidity: function accessRequest(level uint8, pke_requestor string, pke_AFGH string, custom_values bytes32[], stakeholders address[]) returns(bool)
func (_BaseContent *BaseContentTransactorSession) AccessRequest(level uint8, pke_requestor string, pke_AFGH string, custom_values [][32]byte, stakeholders []common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.AccessRequest(&_BaseContent.TransactOpts, level, pke_requestor, pke_AFGH, custom_values, stakeholders)
}

// AddLicensingFeeReceived is a paid mutator transaction binding the contract method 0xf7944471.
//
// Solidity: function addLicensingFeeReceived(amount_to_be_paid uint256) returns(uint256)
func (_BaseContent *BaseContentTransactor) AddLicensingFeeReceived(opts *bind.TransactOpts, amount_to_be_paid *big.Int) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "addLicensingFeeReceived", amount_to_be_paid)
}

// AddLicensingFeeReceived is a paid mutator transaction binding the contract method 0xf7944471.
//
// Solidity: function addLicensingFeeReceived(amount_to_be_paid uint256) returns(uint256)
func (_BaseContent *BaseContentSession) AddLicensingFeeReceived(amount_to_be_paid *big.Int) (*types.Transaction, error) {
	return _BaseContent.Contract.AddLicensingFeeReceived(&_BaseContent.TransactOpts, amount_to_be_paid)
}

// AddLicensingFeeReceived is a paid mutator transaction binding the contract method 0xf7944471.
//
// Solidity: function addLicensingFeeReceived(amount_to_be_paid uint256) returns(uint256)
func (_BaseContent *BaseContentTransactorSession) AddLicensingFeeReceived(amount_to_be_paid *big.Int) (*types.Transaction, error) {
	return _BaseContent.Contract.AddLicensingFeeReceived(&_BaseContent.TransactOpts, amount_to_be_paid)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(object_hash bytes32) returns()
func (_BaseContent *BaseContentTransactor) Commit(opts *bind.TransactOpts, object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "commit", object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(object_hash bytes32) returns()
func (_BaseContent *BaseContentSession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContent.Contract.Commit(&_BaseContent.TransactOpts, object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(object_hash bytes32) returns()
func (_BaseContent *BaseContentTransactorSession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContent.Contract.Commit(&_BaseContent.TransactOpts, object_hash)
}

// GetAccessCharge is a paid mutator transaction binding the contract method 0x879fe48f.
//
// Solidity: function getAccessCharge(level uint8, custom_values bytes32[], stakeholders address[]) returns(uint256)
func (_BaseContent *BaseContentTransactor) GetAccessCharge(opts *bind.TransactOpts, level uint8, custom_values [][32]byte, stakeholders []common.Address) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "getAccessCharge", level, custom_values, stakeholders)
}

// GetAccessCharge is a paid mutator transaction binding the contract method 0x879fe48f.
//
// Solidity: function getAccessCharge(level uint8, custom_values bytes32[], stakeholders address[]) returns(uint256)
func (_BaseContent *BaseContentSession) GetAccessCharge(level uint8, custom_values [][32]byte, stakeholders []common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.GetAccessCharge(&_BaseContent.TransactOpts, level, custom_values, stakeholders)
}

// GetAccessCharge is a paid mutator transaction binding the contract method 0x879fe48f.
//
// Solidity: function getAccessCharge(level uint8, custom_values bytes32[], stakeholders address[]) returns(uint256)
func (_BaseContent *BaseContentTransactorSession) GetAccessCharge(level uint8, custom_values [][32]byte, stakeholders []common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.GetAccessCharge(&_BaseContent.TransactOpts, level, custom_values, stakeholders)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseContent *BaseContentTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseContent *BaseContentSession) Kill() (*types.Transaction, error) {
	return _BaseContent.Contract.Kill(&_BaseContent.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseContent *BaseContentTransactorSession) Kill() (*types.Transaction, error) {
	return _BaseContent.Contract.Kill(&_BaseContent.TransactOpts)
}

// Publish is a paid mutator transaction binding the contract method 0xc68c1d1f.
//
// Solidity: function publish(pct_complete uint8, signed_verification string) returns(bool)
func (_BaseContent *BaseContentTransactor) Publish(opts *bind.TransactOpts, pct_complete uint8, signed_verification string) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "publish", pct_complete, signed_verification)
}

// Publish is a paid mutator transaction binding the contract method 0xc68c1d1f.
//
// Solidity: function publish(pct_complete uint8, signed_verification string) returns(bool)
func (_BaseContent *BaseContentSession) Publish(pct_complete uint8, signed_verification string) (*types.Transaction, error) {
	return _BaseContent.Contract.Publish(&_BaseContent.TransactOpts, pct_complete, signed_verification)
}

// Publish is a paid mutator transaction binding the contract method 0xc68c1d1f.
//
// Solidity: function publish(pct_complete uint8, signed_verification string) returns(bool)
func (_BaseContent *BaseContentTransactorSession) Publish(pct_complete uint8, signed_verification string) (*types.Transaction, error) {
	return _BaseContent.Contract.Publish(&_BaseContent.TransactOpts, pct_complete, signed_verification)
}

// SetAccessCharge is a paid mutator transaction binding the contract method 0xf4d9bae8.
//
// Solidity: function setAccessCharge(charge uint256) returns(uint256)
func (_BaseContent *BaseContentTransactor) SetAccessCharge(opts *bind.TransactOpts, charge *big.Int) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "setAccessCharge", charge)
}

// SetAccessCharge is a paid mutator transaction binding the contract method 0xf4d9bae8.
//
// Solidity: function setAccessCharge(charge uint256) returns(uint256)
func (_BaseContent *BaseContentSession) SetAccessCharge(charge *big.Int) (*types.Transaction, error) {
	return _BaseContent.Contract.SetAccessCharge(&_BaseContent.TransactOpts, charge)
}

// SetAccessCharge is a paid mutator transaction binding the contract method 0xf4d9bae8.
//
// Solidity: function setAccessCharge(charge uint256) returns(uint256)
func (_BaseContent *BaseContentTransactorSession) SetAccessCharge(charge *big.Int) (*types.Transaction, error) {
	return _BaseContent.Contract.SetAccessCharge(&_BaseContent.TransactOpts, charge)
}

// SetAddressKMS is a paid mutator transaction binding the contract method 0xc9e8e72d.
//
// Solidity: function setAddressKMS(address_KMS address) returns()
func (_BaseContent *BaseContentTransactor) SetAddressKMS(opts *bind.TransactOpts, address_KMS common.Address) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "setAddressKMS", address_KMS)
}

// SetAddressKMS is a paid mutator transaction binding the contract method 0xc9e8e72d.
//
// Solidity: function setAddressKMS(address_KMS address) returns()
func (_BaseContent *BaseContentSession) SetAddressKMS(address_KMS common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.SetAddressKMS(&_BaseContent.TransactOpts, address_KMS)
}

// SetAddressKMS is a paid mutator transaction binding the contract method 0xc9e8e72d.
//
// Solidity: function setAddressKMS(address_KMS address) returns()
func (_BaseContent *BaseContentTransactorSession) SetAddressKMS(address_KMS common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.SetAddressKMS(&_BaseContent.TransactOpts, address_KMS)
}

// SetContentContractAddress is a paid mutator transaction binding the contract method 0xe5385303.
//
// Solidity: function setContentContractAddress(addr address) returns()
func (_BaseContent *BaseContentTransactor) SetContentContractAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "setContentContractAddress", addr)
}

// SetContentContractAddress is a paid mutator transaction binding the contract method 0xe5385303.
//
// Solidity: function setContentContractAddress(addr address) returns()
func (_BaseContent *BaseContentSession) SetContentContractAddress(addr common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.SetContentContractAddress(&_BaseContent.TransactOpts, addr)
}

// SetContentContractAddress is a paid mutator transaction binding the contract method 0xe5385303.
//
// Solidity: function setContentContractAddress(addr address) returns()
func (_BaseContent *BaseContentTransactorSession) SetContentContractAddress(addr common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.SetContentContractAddress(&_BaseContent.TransactOpts, addr)
}

// SetContentType is a paid mutator transaction binding the contract method 0x69348967.
//
// Solidity: function setContentType(content_type address) returns()
func (_BaseContent *BaseContentTransactor) SetContentType(opts *bind.TransactOpts, content_type common.Address) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "setContentType", content_type)
}

// SetContentType is a paid mutator transaction binding the contract method 0x69348967.
//
// Solidity: function setContentType(content_type address) returns()
func (_BaseContent *BaseContentSession) SetContentType(content_type common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.SetContentType(&_BaseContent.TransactOpts, content_type)
}

// SetContentType is a paid mutator transaction binding the contract method 0x69348967.
//
// Solidity: function setContentType(content_type address) returns()
func (_BaseContent *BaseContentTransactorSession) SetContentType(content_type common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.SetContentType(&_BaseContent.TransactOpts, content_type)
}

// SetLicensingFee is a paid mutator transaction binding the contract method 0xb6524089.
//
// Solidity: function setLicensingFee(licensing_fee uint256) returns()
func (_BaseContent *BaseContentTransactor) SetLicensingFee(opts *bind.TransactOpts, licensing_fee *big.Int) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "setLicensingFee", licensing_fee)
}

// SetLicensingFee is a paid mutator transaction binding the contract method 0xb6524089.
//
// Solidity: function setLicensingFee(licensing_fee uint256) returns()
func (_BaseContent *BaseContentSession) SetLicensingFee(licensing_fee *big.Int) (*types.Transaction, error) {
	return _BaseContent.Contract.SetLicensingFee(&_BaseContent.TransactOpts, licensing_fee)
}

// SetLicensingFee is a paid mutator transaction binding the contract method 0xb6524089.
//
// Solidity: function setLicensingFee(licensing_fee uint256) returns()
func (_BaseContent *BaseContentTransactorSession) SetLicensingFee(licensing_fee *big.Int) (*types.Transaction, error) {
	return _BaseContent.Contract.SetLicensingFee(&_BaseContent.TransactOpts, licensing_fee)
}

// SetStatusCode is a paid mutator transaction binding the contract method 0x5267db44.
//
// Solidity: function setStatusCode(status_code int256) returns(int256)
func (_BaseContent *BaseContentTransactor) SetStatusCode(opts *bind.TransactOpts, status_code *big.Int) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "setStatusCode", status_code)
}

// SetStatusCode is a paid mutator transaction binding the contract method 0x5267db44.
//
// Solidity: function setStatusCode(status_code int256) returns(int256)
func (_BaseContent *BaseContentSession) SetStatusCode(status_code *big.Int) (*types.Transaction, error) {
	return _BaseContent.Contract.SetStatusCode(&_BaseContent.TransactOpts, status_code)
}

// SetStatusCode is a paid mutator transaction binding the contract method 0x5267db44.
//
// Solidity: function setStatusCode(status_code int256) returns(int256)
func (_BaseContent *BaseContentTransactorSession) SetStatusCode(status_code *big.Int) (*types.Transaction, error) {
	return _BaseContent.Contract.SetStatusCode(&_BaseContent.TransactOpts, status_code)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_BaseContent *BaseContentTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_BaseContent *BaseContentSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.TransferCreatorship(&_BaseContent.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_BaseContent *BaseContentTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.TransferCreatorship(&_BaseContent.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseContent *BaseContentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseContent *BaseContentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.TransferOwnership(&_BaseContent.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseContent *BaseContentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.TransferOwnership(&_BaseContent.TransactOpts, newOwner)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xc287e0ed.
//
// Solidity: function updateRequest() returns()
func (_BaseContent *BaseContentTransactor) UpdateRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "updateRequest")
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xc287e0ed.
//
// Solidity: function updateRequest() returns()
func (_BaseContent *BaseContentSession) UpdateRequest() (*types.Transaction, error) {
	return _BaseContent.Contract.UpdateRequest(&_BaseContent.TransactOpts)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xc287e0ed.
//
// Solidity: function updateRequest() returns()
func (_BaseContent *BaseContentTransactorSession) UpdateRequest() (*types.Transaction, error) {
	return _BaseContent.Contract.UpdateRequest(&_BaseContent.TransactOpts)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x708f53ab.
//
// Solidity: function updateStatus(status_code int256, percent_complete uint8) returns(uint256)
func (_BaseContent *BaseContentTransactor) UpdateStatus(opts *bind.TransactOpts, status_code *big.Int, percent_complete uint8) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "updateStatus", status_code, percent_complete)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x708f53ab.
//
// Solidity: function updateStatus(status_code int256, percent_complete uint8) returns(uint256)
func (_BaseContent *BaseContentSession) UpdateStatus(status_code *big.Int, percent_complete uint8) (*types.Transaction, error) {
	return _BaseContent.Contract.UpdateStatus(&_BaseContent.TransactOpts, status_code, percent_complete)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x708f53ab.
//
// Solidity: function updateStatus(status_code int256, percent_complete uint8) returns(uint256)
func (_BaseContent *BaseContentTransactorSession) UpdateStatus(status_code *big.Int, percent_complete uint8) (*types.Transaction, error) {
	return _BaseContent.Contract.UpdateStatus(&_BaseContent.TransactOpts, status_code, percent_complete)
}

// BaseContentAccessCompleteIterator is returned from FilterAccessComplete and is used to iterate over the raw logs and unpacked data for AccessComplete events raised by the BaseContent contract.
type BaseContentAccessCompleteIterator struct {
	Event *BaseContentAccessComplete // Event containing the contract specifics and raw log

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
func (it *BaseContentAccessCompleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentAccessComplete)
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
		it.Event = new(BaseContentAccessComplete)
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
func (it *BaseContentAccessCompleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentAccessCompleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentAccessComplete represents a AccessComplete event raised by the BaseContent contract.
type BaseContentAccessComplete struct {
	RequestID            *big.Int
	ScorePct             *big.Int
	MlOutHash            [32]byte
	CustomContractResult bool
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterAccessComplete is a free log retrieval operation binding the contract event 0x7f1f4b28434ce7beab4983e64a8b5bb96e195a67029fdaff925028aec57fbc6b.
//
// Solidity: e AccessComplete(requestID uint256, scorePct uint256, mlOutHash bytes32, customContractResult bool)
func (_BaseContent *BaseContentFilterer) FilterAccessComplete(opts *bind.FilterOpts) (*BaseContentAccessCompleteIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "AccessComplete")
	if err != nil {
		return nil, err
	}
	return &BaseContentAccessCompleteIterator{contract: _BaseContent.contract, event: "AccessComplete", logs: logs, sub: sub}, nil
}

// WatchAccessComplete is a free log subscription operation binding the contract event 0x7f1f4b28434ce7beab4983e64a8b5bb96e195a67029fdaff925028aec57fbc6b.
//
// Solidity: e AccessComplete(requestID uint256, scorePct uint256, mlOutHash bytes32, customContractResult bool)
func (_BaseContent *BaseContentFilterer) WatchAccessComplete(opts *bind.WatchOpts, sink chan<- *BaseContentAccessComplete) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "AccessComplete")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentAccessComplete)
				if err := _BaseContent.contract.UnpackLog(event, "AccessComplete", log); err != nil {
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

// BaseContentAccessGrantIterator is returned from FilterAccessGrant and is used to iterate over the raw logs and unpacked data for AccessGrant events raised by the BaseContent contract.
type BaseContentAccessGrantIterator struct {
	Event *BaseContentAccessGrant // Event containing the contract specifics and raw log

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
func (it *BaseContentAccessGrantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentAccessGrant)
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
		it.Event = new(BaseContentAccessGrant)
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
func (it *BaseContentAccessGrantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentAccessGrantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentAccessGrant represents a AccessGrant event raised by the BaseContent contract.
type BaseContentAccessGrant struct {
	RequestID       *big.Int
	AccessGranted   bool
	ReKey           string
	EncryptedAESKey string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccessGrant is a free log retrieval operation binding the contract event 0x475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718.
//
// Solidity: e AccessGrant(requestID uint256, access_granted bool, reKey string, encryptedAESKey string)
func (_BaseContent *BaseContentFilterer) FilterAccessGrant(opts *bind.FilterOpts) (*BaseContentAccessGrantIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "AccessGrant")
	if err != nil {
		return nil, err
	}
	return &BaseContentAccessGrantIterator{contract: _BaseContent.contract, event: "AccessGrant", logs: logs, sub: sub}, nil
}

// WatchAccessGrant is a free log subscription operation binding the contract event 0x475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718.
//
// Solidity: e AccessGrant(requestID uint256, access_granted bool, reKey string, encryptedAESKey string)
func (_BaseContent *BaseContentFilterer) WatchAccessGrant(opts *bind.WatchOpts, sink chan<- *BaseContentAccessGrant) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "AccessGrant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentAccessGrant)
				if err := _BaseContent.contract.UnpackLog(event, "AccessGrant", log); err != nil {
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

// BaseContentAccessRequestIterator is returned from FilterAccessRequest and is used to iterate over the raw logs and unpacked data for AccessRequest events raised by the BaseContent contract.
type BaseContentAccessRequestIterator struct {
	Event *BaseContentAccessRequest // Event containing the contract specifics and raw log

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
func (it *BaseContentAccessRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentAccessRequest)
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
		it.Event = new(BaseContentAccessRequest)
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
func (it *BaseContentAccessRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentAccessRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentAccessRequest represents a AccessRequest event raised by the BaseContent contract.
type BaseContentAccessRequest struct {
	RequestValidity *big.Int
	RequestID       *big.Int
	Level           uint8
	ContentHash     [32]byte
	PkeRequestor    string
	PkeAFGH         string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccessRequest is a free log retrieval operation binding the contract event 0x272aa3d0efd05809b06d40e95792681d6b844a29fe97cd703cfd5eee27510740.
//
// Solidity: e AccessRequest(requestValidity uint256, requestID uint256, level uint8, contentHash bytes32, pkeRequestor string, pkeAFGH string)
func (_BaseContent *BaseContentFilterer) FilterAccessRequest(opts *bind.FilterOpts) (*BaseContentAccessRequestIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "AccessRequest")
	if err != nil {
		return nil, err
	}
	return &BaseContentAccessRequestIterator{contract: _BaseContent.contract, event: "AccessRequest", logs: logs, sub: sub}, nil
}

// WatchAccessRequest is a free log subscription operation binding the contract event 0x272aa3d0efd05809b06d40e95792681d6b844a29fe97cd703cfd5eee27510740.
//
// Solidity: e AccessRequest(requestValidity uint256, requestID uint256, level uint8, contentHash bytes32, pkeRequestor string, pkeAFGH string)
func (_BaseContent *BaseContentFilterer) WatchAccessRequest(opts *bind.WatchOpts, sink chan<- *BaseContentAccessRequest) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "AccessRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentAccessRequest)
				if err := _BaseContent.contract.UnpackLog(event, "AccessRequest", log); err != nil {
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

// BaseContentAccessRequestStakeholderIterator is returned from FilterAccessRequestStakeholder and is used to iterate over the raw logs and unpacked data for AccessRequestStakeholder events raised by the BaseContent contract.
type BaseContentAccessRequestStakeholderIterator struct {
	Event *BaseContentAccessRequestStakeholder // Event containing the contract specifics and raw log

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
func (it *BaseContentAccessRequestStakeholderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentAccessRequestStakeholder)
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
		it.Event = new(BaseContentAccessRequestStakeholder)
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
func (it *BaseContentAccessRequestStakeholderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentAccessRequestStakeholderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentAccessRequestStakeholder represents a AccessRequestStakeholder event raised by the BaseContent contract.
type BaseContentAccessRequestStakeholder struct {
	Stakeholder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAccessRequestStakeholder is a free log retrieval operation binding the contract event 0xb6e3239e521a6c66920ae634f8e921a37e6991d520ac44d52f8516397f41b684.
//
// Solidity: e AccessRequestStakeholder(stakeholder address)
func (_BaseContent *BaseContentFilterer) FilterAccessRequestStakeholder(opts *bind.FilterOpts) (*BaseContentAccessRequestStakeholderIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "AccessRequestStakeholder")
	if err != nil {
		return nil, err
	}
	return &BaseContentAccessRequestStakeholderIterator{contract: _BaseContent.contract, event: "AccessRequestStakeholder", logs: logs, sub: sub}, nil
}

// WatchAccessRequestStakeholder is a free log subscription operation binding the contract event 0xb6e3239e521a6c66920ae634f8e921a37e6991d520ac44d52f8516397f41b684.
//
// Solidity: e AccessRequestStakeholder(stakeholder address)
func (_BaseContent *BaseContentFilterer) WatchAccessRequestStakeholder(opts *bind.WatchOpts, sink chan<- *BaseContentAccessRequestStakeholder) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "AccessRequestStakeholder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentAccessRequestStakeholder)
				if err := _BaseContent.contract.UnpackLog(event, "AccessRequestStakeholder", log); err != nil {
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

// BaseContentAccessRequestValueIterator is returned from FilterAccessRequestValue and is used to iterate over the raw logs and unpacked data for AccessRequestValue events raised by the BaseContent contract.
type BaseContentAccessRequestValueIterator struct {
	Event *BaseContentAccessRequestValue // Event containing the contract specifics and raw log

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
func (it *BaseContentAccessRequestValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentAccessRequestValue)
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
		it.Event = new(BaseContentAccessRequestValue)
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
func (it *BaseContentAccessRequestValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentAccessRequestValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentAccessRequestValue represents a AccessRequestValue event raised by the BaseContent contract.
type BaseContentAccessRequestValue struct {
	CustomValue [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAccessRequestValue is a free log retrieval operation binding the contract event 0x515e0a48b385fce2a8e4d9f169a97c4f6ea669a752358f5e6ab37cc3c2e84c38.
//
// Solidity: e AccessRequestValue(customValue bytes32)
func (_BaseContent *BaseContentFilterer) FilterAccessRequestValue(opts *bind.FilterOpts) (*BaseContentAccessRequestValueIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "AccessRequestValue")
	if err != nil {
		return nil, err
	}
	return &BaseContentAccessRequestValueIterator{contract: _BaseContent.contract, event: "AccessRequestValue", logs: logs, sub: sub}, nil
}

// WatchAccessRequestValue is a free log subscription operation binding the contract event 0x515e0a48b385fce2a8e4d9f169a97c4f6ea669a752358f5e6ab37cc3c2e84c38.
//
// Solidity: e AccessRequestValue(customValue bytes32)
func (_BaseContent *BaseContentFilterer) WatchAccessRequestValue(opts *bind.WatchOpts, sink chan<- *BaseContentAccessRequestValue) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "AccessRequestValue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentAccessRequestValue)
				if err := _BaseContent.contract.UnpackLog(event, "AccessRequestValue", log); err != nil {
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

// BaseContentCommitIterator is returned from FilterCommit and is used to iterate over the raw logs and unpacked data for Commit events raised by the BaseContent contract.
type BaseContentCommitIterator struct {
	Event *BaseContentCommit // Event containing the contract specifics and raw log

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
func (it *BaseContentCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentCommit)
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
		it.Event = new(BaseContentCommit)
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
func (it *BaseContentCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentCommit represents a Commit event raised by the BaseContent contract.
type BaseContentCommit struct {
	ObjectHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommit is a free log retrieval operation binding the contract event 0x9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff.
//
// Solidity: e Commit(objectHash bytes32)
func (_BaseContent *BaseContentFilterer) FilterCommit(opts *bind.FilterOpts) (*BaseContentCommitIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "Commit")
	if err != nil {
		return nil, err
	}
	return &BaseContentCommitIterator{contract: _BaseContent.contract, event: "Commit", logs: logs, sub: sub}, nil
}

// WatchCommit is a free log subscription operation binding the contract event 0x9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff.
//
// Solidity: e Commit(objectHash bytes32)
func (_BaseContent *BaseContentFilterer) WatchCommit(opts *bind.WatchOpts, sink chan<- *BaseContentCommit) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "Commit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentCommit)
				if err := _BaseContent.contract.UnpackLog(event, "Commit", log); err != nil {
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

// BaseContentContentObjectCreateIterator is returned from FilterContentObjectCreate and is used to iterate over the raw logs and unpacked data for ContentObjectCreate events raised by the BaseContent contract.
type BaseContentContentObjectCreateIterator struct {
	Event *BaseContentContentObjectCreate // Event containing the contract specifics and raw log

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
func (it *BaseContentContentObjectCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentContentObjectCreate)
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
		it.Event = new(BaseContentContentObjectCreate)
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
func (it *BaseContentContentObjectCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentContentObjectCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentContentObjectCreate represents a ContentObjectCreate event raised by the BaseContent contract.
type BaseContentContentObjectCreate struct {
	ContainingLibrary common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterContentObjectCreate is a free log retrieval operation binding the contract event 0xc3decc188980e855666b70498ca85e8fa284d97d30483d828fa126f7303d7d19.
//
// Solidity: e ContentObjectCreate(containingLibrary address)
func (_BaseContent *BaseContentFilterer) FilterContentObjectCreate(opts *bind.FilterOpts) (*BaseContentContentObjectCreateIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "ContentObjectCreate")
	if err != nil {
		return nil, err
	}
	return &BaseContentContentObjectCreateIterator{contract: _BaseContent.contract, event: "ContentObjectCreate", logs: logs, sub: sub}, nil
}

// WatchContentObjectCreate is a free log subscription operation binding the contract event 0xc3decc188980e855666b70498ca85e8fa284d97d30483d828fa126f7303d7d19.
//
// Solidity: e ContentObjectCreate(containingLibrary address)
func (_BaseContent *BaseContentFilterer) WatchContentObjectCreate(opts *bind.WatchOpts, sink chan<- *BaseContentContentObjectCreate) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "ContentObjectCreate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentContentObjectCreate)
				if err := _BaseContent.contract.UnpackLog(event, "ContentObjectCreate", log); err != nil {
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

// BaseContentGetAccessChargeIterator is returned from FilterGetAccessCharge and is used to iterate over the raw logs and unpacked data for GetAccessCharge events raised by the BaseContent contract.
type BaseContentGetAccessChargeIterator struct {
	Event *BaseContentGetAccessCharge // Event containing the contract specifics and raw log

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
func (it *BaseContentGetAccessChargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentGetAccessCharge)
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
		it.Event = new(BaseContentGetAccessCharge)
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
func (it *BaseContentGetAccessChargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentGetAccessChargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentGetAccessCharge represents a GetAccessCharge event raised by the BaseContent contract.
type BaseContentGetAccessCharge struct {
	Level        uint8
	AccessCharge *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGetAccessCharge is a free log retrieval operation binding the contract event 0xa58326ee5bb617cb8b4f0d0f5f557c469d2d05d7a738f777037deda9c724b370.
//
// Solidity: e GetAccessCharge(level uint8, accessCharge uint256)
func (_BaseContent *BaseContentFilterer) FilterGetAccessCharge(opts *bind.FilterOpts) (*BaseContentGetAccessChargeIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "GetAccessCharge")
	if err != nil {
		return nil, err
	}
	return &BaseContentGetAccessChargeIterator{contract: _BaseContent.contract, event: "GetAccessCharge", logs: logs, sub: sub}, nil
}

// WatchGetAccessCharge is a free log subscription operation binding the contract event 0xa58326ee5bb617cb8b4f0d0f5f557c469d2d05d7a738f777037deda9c724b370.
//
// Solidity: e GetAccessCharge(level uint8, accessCharge uint256)
func (_BaseContent *BaseContentFilterer) WatchGetAccessCharge(opts *bind.WatchOpts, sink chan<- *BaseContentGetAccessCharge) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "GetAccessCharge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentGetAccessCharge)
				if err := _BaseContent.contract.UnpackLog(event, "GetAccessCharge", log); err != nil {
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

// BaseContentInsufficientFundsIterator is returned from FilterInsufficientFunds and is used to iterate over the raw logs and unpacked data for InsufficientFunds events raised by the BaseContent contract.
type BaseContentInsufficientFundsIterator struct {
	Event *BaseContentInsufficientFunds // Event containing the contract specifics and raw log

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
func (it *BaseContentInsufficientFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentInsufficientFunds)
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
		it.Event = new(BaseContentInsufficientFunds)
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
func (it *BaseContentInsufficientFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentInsufficientFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentInsufficientFunds represents a InsufficientFunds event raised by the BaseContent contract.
type BaseContentInsufficientFunds struct {
	AccessCharge   *big.Int
	AmountProvided *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInsufficientFunds is a free log retrieval operation binding the contract event 0x03eb8b54a949acec2cd08fdb6d6bd4647a1f2c907d75d6900648effa92eb147f.
//
// Solidity: e InsufficientFunds(accessCharge uint256, amountProvided uint256)
func (_BaseContent *BaseContentFilterer) FilterInsufficientFunds(opts *bind.FilterOpts) (*BaseContentInsufficientFundsIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "InsufficientFunds")
	if err != nil {
		return nil, err
	}
	return &BaseContentInsufficientFundsIterator{contract: _BaseContent.contract, event: "InsufficientFunds", logs: logs, sub: sub}, nil
}

// WatchInsufficientFunds is a free log subscription operation binding the contract event 0x03eb8b54a949acec2cd08fdb6d6bd4647a1f2c907d75d6900648effa92eb147f.
//
// Solidity: e InsufficientFunds(accessCharge uint256, amountProvided uint256)
func (_BaseContent *BaseContentFilterer) WatchInsufficientFunds(opts *bind.WatchOpts, sink chan<- *BaseContentInsufficientFunds) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "InsufficientFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentInsufficientFunds)
				if err := _BaseContent.contract.UnpackLog(event, "InsufficientFunds", log); err != nil {
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

// BaseContentInvokeCustomPostHookIterator is returned from FilterInvokeCustomPostHook and is used to iterate over the raw logs and unpacked data for InvokeCustomPostHook events raised by the BaseContent contract.
type BaseContentInvokeCustomPostHookIterator struct {
	Event *BaseContentInvokeCustomPostHook // Event containing the contract specifics and raw log

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
func (it *BaseContentInvokeCustomPostHookIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentInvokeCustomPostHook)
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
		it.Event = new(BaseContentInvokeCustomPostHook)
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
func (it *BaseContentInvokeCustomPostHookIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentInvokeCustomPostHookIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentInvokeCustomPostHook represents a InvokeCustomPostHook event raised by the BaseContent contract.
type BaseContentInvokeCustomPostHook struct {
	CustomContract common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInvokeCustomPostHook is a free log retrieval operation binding the contract event 0x97d9c9779ed3ed8b9a6edfe16d17b1fdec843245747a19abfb621806e37d4a89.
//
// Solidity: e InvokeCustomPostHook(custom_contract address)
func (_BaseContent *BaseContentFilterer) FilterInvokeCustomPostHook(opts *bind.FilterOpts) (*BaseContentInvokeCustomPostHookIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "InvokeCustomPostHook")
	if err != nil {
		return nil, err
	}
	return &BaseContentInvokeCustomPostHookIterator{contract: _BaseContent.contract, event: "InvokeCustomPostHook", logs: logs, sub: sub}, nil
}

// WatchInvokeCustomPostHook is a free log subscription operation binding the contract event 0x97d9c9779ed3ed8b9a6edfe16d17b1fdec843245747a19abfb621806e37d4a89.
//
// Solidity: e InvokeCustomPostHook(custom_contract address)
func (_BaseContent *BaseContentFilterer) WatchInvokeCustomPostHook(opts *bind.WatchOpts, sink chan<- *BaseContentInvokeCustomPostHook) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "InvokeCustomPostHook")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentInvokeCustomPostHook)
				if err := _BaseContent.contract.UnpackLog(event, "InvokeCustomPostHook", log); err != nil {
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

// BaseContentInvokeCustomPreHookIterator is returned from FilterInvokeCustomPreHook and is used to iterate over the raw logs and unpacked data for InvokeCustomPreHook events raised by the BaseContent contract.
type BaseContentInvokeCustomPreHookIterator struct {
	Event *BaseContentInvokeCustomPreHook // Event containing the contract specifics and raw log

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
func (it *BaseContentInvokeCustomPreHookIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentInvokeCustomPreHook)
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
		it.Event = new(BaseContentInvokeCustomPreHook)
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
func (it *BaseContentInvokeCustomPreHookIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentInvokeCustomPreHookIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentInvokeCustomPreHook represents a InvokeCustomPreHook event raised by the BaseContent contract.
type BaseContentInvokeCustomPreHook struct {
	CustomContract common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInvokeCustomPreHook is a free log retrieval operation binding the contract event 0x12b04791b5caab768e2757268992f0c62801e3921d9e310c893f0d5f9caa5f71.
//
// Solidity: e InvokeCustomPreHook(custom_contract address)
func (_BaseContent *BaseContentFilterer) FilterInvokeCustomPreHook(opts *bind.FilterOpts) (*BaseContentInvokeCustomPreHookIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "InvokeCustomPreHook")
	if err != nil {
		return nil, err
	}
	return &BaseContentInvokeCustomPreHookIterator{contract: _BaseContent.contract, event: "InvokeCustomPreHook", logs: logs, sub: sub}, nil
}

// WatchInvokeCustomPreHook is a free log subscription operation binding the contract event 0x12b04791b5caab768e2757268992f0c62801e3921d9e310c893f0d5f9caa5f71.
//
// Solidity: e InvokeCustomPreHook(custom_contract address)
func (_BaseContent *BaseContentFilterer) WatchInvokeCustomPreHook(opts *bind.WatchOpts, sink chan<- *BaseContentInvokeCustomPreHook) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "InvokeCustomPreHook")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentInvokeCustomPreHook)
				if err := _BaseContent.contract.UnpackLog(event, "InvokeCustomPreHook", log); err != nil {
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

// BaseContentPublishIterator is returned from FilterPublish and is used to iterate over the raw logs and unpacked data for Publish events raised by the BaseContent contract.
type BaseContentPublishIterator struct {
	Event *BaseContentPublish // Event containing the contract specifics and raw log

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
func (it *BaseContentPublishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentPublish)
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
		it.Event = new(BaseContentPublish)
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
func (it *BaseContentPublishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentPublishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentPublish represents a Publish event raised by the BaseContent contract.
type BaseContentPublish struct {
	PctComplete   uint8
	RequestStatus bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPublish is a free log retrieval operation binding the contract event 0x9625d8ae00bbc11bf4dbac17532594047d634ea66e1baa95d4974d086f261c30.
//
// Solidity: e Publish(pctComplete uint8, requestStatus bool)
func (_BaseContent *BaseContentFilterer) FilterPublish(opts *bind.FilterOpts) (*BaseContentPublishIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "Publish")
	if err != nil {
		return nil, err
	}
	return &BaseContentPublishIterator{contract: _BaseContent.contract, event: "Publish", logs: logs, sub: sub}, nil
}

// WatchPublish is a free log subscription operation binding the contract event 0x9625d8ae00bbc11bf4dbac17532594047d634ea66e1baa95d4974d086f261c30.
//
// Solidity: e Publish(pctComplete uint8, requestStatus bool)
func (_BaseContent *BaseContentFilterer) WatchPublish(opts *bind.WatchOpts, sink chan<- *BaseContentPublish) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "Publish")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentPublish)
				if err := _BaseContent.contract.UnpackLog(event, "Publish", log); err != nil {
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

// BaseContentReturnCustomHookIterator is returned from FilterReturnCustomHook and is used to iterate over the raw logs and unpacked data for ReturnCustomHook events raised by the BaseContent contract.
type BaseContentReturnCustomHookIterator struct {
	Event *BaseContentReturnCustomHook // Event containing the contract specifics and raw log

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
func (it *BaseContentReturnCustomHookIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentReturnCustomHook)
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
		it.Event = new(BaseContentReturnCustomHook)
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
func (it *BaseContentReturnCustomHookIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentReturnCustomHookIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentReturnCustomHook represents a ReturnCustomHook event raised by the BaseContent contract.
type BaseContentReturnCustomHook struct {
	CustomContract common.Address
	Result         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReturnCustomHook is a free log retrieval operation binding the contract event 0x8c693e8b27db7caf9b9637b66dcc11444760023a4d53e95407a3acef1b249f50.
//
// Solidity: e ReturnCustomHook(custom_contract address, result uint256)
func (_BaseContent *BaseContentFilterer) FilterReturnCustomHook(opts *bind.FilterOpts) (*BaseContentReturnCustomHookIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "ReturnCustomHook")
	if err != nil {
		return nil, err
	}
	return &BaseContentReturnCustomHookIterator{contract: _BaseContent.contract, event: "ReturnCustomHook", logs: logs, sub: sub}, nil
}

// WatchReturnCustomHook is a free log subscription operation binding the contract event 0x8c693e8b27db7caf9b9637b66dcc11444760023a4d53e95407a3acef1b249f50.
//
// Solidity: e ReturnCustomHook(custom_contract address, result uint256)
func (_BaseContent *BaseContentFilterer) WatchReturnCustomHook(opts *bind.WatchOpts, sink chan<- *BaseContentReturnCustomHook) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "ReturnCustomHook")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentReturnCustomHook)
				if err := _BaseContent.contract.UnpackLog(event, "ReturnCustomHook", log); err != nil {
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

// BaseContentSetAccessChargeIterator is returned from FilterSetAccessCharge and is used to iterate over the raw logs and unpacked data for SetAccessCharge events raised by the BaseContent contract.
type BaseContentSetAccessChargeIterator struct {
	Event *BaseContentSetAccessCharge // Event containing the contract specifics and raw log

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
func (it *BaseContentSetAccessChargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentSetAccessCharge)
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
		it.Event = new(BaseContentSetAccessCharge)
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
func (it *BaseContentSetAccessChargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentSetAccessChargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentSetAccessCharge represents a SetAccessCharge event raised by the BaseContent contract.
type BaseContentSetAccessCharge struct {
	AccessCharge *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetAccessCharge is a free log retrieval operation binding the contract event 0x4114f8ef80b6de2161db580cbefa14e1892d15d3ebe2062c9914e4a5773114a3.
//
// Solidity: e SetAccessCharge(accessCharge uint256)
func (_BaseContent *BaseContentFilterer) FilterSetAccessCharge(opts *bind.FilterOpts) (*BaseContentSetAccessChargeIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "SetAccessCharge")
	if err != nil {
		return nil, err
	}
	return &BaseContentSetAccessChargeIterator{contract: _BaseContent.contract, event: "SetAccessCharge", logs: logs, sub: sub}, nil
}

// WatchSetAccessCharge is a free log subscription operation binding the contract event 0x4114f8ef80b6de2161db580cbefa14e1892d15d3ebe2062c9914e4a5773114a3.
//
// Solidity: e SetAccessCharge(accessCharge uint256)
func (_BaseContent *BaseContentFilterer) WatchSetAccessCharge(opts *bind.WatchOpts, sink chan<- *BaseContentSetAccessCharge) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "SetAccessCharge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentSetAccessCharge)
				if err := _BaseContent.contract.UnpackLog(event, "SetAccessCharge", log); err != nil {
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

// BaseContentSetContentContractIterator is returned from FilterSetContentContract and is used to iterate over the raw logs and unpacked data for SetContentContract events raised by the BaseContent contract.
type BaseContentSetContentContractIterator struct {
	Event *BaseContentSetContentContract // Event containing the contract specifics and raw log

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
func (it *BaseContentSetContentContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentSetContentContract)
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
		it.Event = new(BaseContentSetContentContract)
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
func (it *BaseContentSetContentContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentSetContentContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentSetContentContract represents a SetContentContract event raised by the BaseContent contract.
type BaseContentSetContentContract struct {
	ContentContractAddress common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetContentContract is a free log retrieval operation binding the contract event 0xa6f2e38f0cfebf27212317fced3ac40bc62e00bd33f38d69603710740c69acb7.
//
// Solidity: e SetContentContract(contentContractAddress address)
func (_BaseContent *BaseContentFilterer) FilterSetContentContract(opts *bind.FilterOpts) (*BaseContentSetContentContractIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "SetContentContract")
	if err != nil {
		return nil, err
	}
	return &BaseContentSetContentContractIterator{contract: _BaseContent.contract, event: "SetContentContract", logs: logs, sub: sub}, nil
}

// WatchSetContentContract is a free log subscription operation binding the contract event 0xa6f2e38f0cfebf27212317fced3ac40bc62e00bd33f38d69603710740c69acb7.
//
// Solidity: e SetContentContract(contentContractAddress address)
func (_BaseContent *BaseContentFilterer) WatchSetContentContract(opts *bind.WatchOpts, sink chan<- *BaseContentSetContentContract) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "SetContentContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentSetContentContract)
				if err := _BaseContent.contract.UnpackLog(event, "SetContentContract", log); err != nil {
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

// BaseContentSetContentTypeIterator is returned from FilterSetContentType and is used to iterate over the raw logs and unpacked data for SetContentType events raised by the BaseContent contract.
type BaseContentSetContentTypeIterator struct {
	Event *BaseContentSetContentType // Event containing the contract specifics and raw log

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
func (it *BaseContentSetContentTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentSetContentType)
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
		it.Event = new(BaseContentSetContentType)
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
func (it *BaseContentSetContentTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentSetContentTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentSetContentType represents a SetContentType event raised by the BaseContent contract.
type BaseContentSetContentType struct {
	ContentType            common.Address
	ContentContractAddress common.Address
	LicensingFee           *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetContentType is a free log retrieval operation binding the contract event 0x1c158b764edca9bf03a401a388355220a7395febb3c011729d8060572f065627.
//
// Solidity: e SetContentType(contentType address, contentContractAddress address, licensingFee uint256)
func (_BaseContent *BaseContentFilterer) FilterSetContentType(opts *bind.FilterOpts) (*BaseContentSetContentTypeIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "SetContentType")
	if err != nil {
		return nil, err
	}
	return &BaseContentSetContentTypeIterator{contract: _BaseContent.contract, event: "SetContentType", logs: logs, sub: sub}, nil
}

// WatchSetContentType is a free log subscription operation binding the contract event 0x1c158b764edca9bf03a401a388355220a7395febb3c011729d8060572f065627.
//
// Solidity: e SetContentType(contentType address, contentContractAddress address, licensingFee uint256)
func (_BaseContent *BaseContentFilterer) WatchSetContentType(opts *bind.WatchOpts, sink chan<- *BaseContentSetContentType) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "SetContentType")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentSetContentType)
				if err := _BaseContent.contract.UnpackLog(event, "SetContentType", log); err != nil {
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

// BaseContentSetStatusCodeIterator is returned from FilterSetStatusCode and is used to iterate over the raw logs and unpacked data for SetStatusCode events raised by the BaseContent contract.
type BaseContentSetStatusCodeIterator struct {
	Event *BaseContentSetStatusCode // Event containing the contract specifics and raw log

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
func (it *BaseContentSetStatusCodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentSetStatusCode)
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
		it.Event = new(BaseContentSetStatusCode)
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
func (it *BaseContentSetStatusCodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentSetStatusCodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentSetStatusCode represents a SetStatusCode event raised by the BaseContent contract.
type BaseContentSetStatusCode struct {
	StatusCode *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetStatusCode is a free log retrieval operation binding the contract event 0xda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a2.
//
// Solidity: e SetStatusCode(statusCode int256)
func (_BaseContent *BaseContentFilterer) FilterSetStatusCode(opts *bind.FilterOpts) (*BaseContentSetStatusCodeIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "SetStatusCode")
	if err != nil {
		return nil, err
	}
	return &BaseContentSetStatusCodeIterator{contract: _BaseContent.contract, event: "SetStatusCode", logs: logs, sub: sub}, nil
}

// WatchSetStatusCode is a free log subscription operation binding the contract event 0xda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a2.
//
// Solidity: e SetStatusCode(statusCode int256)
func (_BaseContent *BaseContentFilterer) WatchSetStatusCode(opts *bind.WatchOpts, sink chan<- *BaseContentSetStatusCode) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "SetStatusCode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentSetStatusCode)
				if err := _BaseContent.contract.UnpackLog(event, "SetStatusCode", log); err != nil {
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

// BaseContentUpdateRequestIterator is returned from FilterUpdateRequest and is used to iterate over the raw logs and unpacked data for UpdateRequest events raised by the BaseContent contract.
type BaseContentUpdateRequestIterator struct {
	Event *BaseContentUpdateRequest // Event containing the contract specifics and raw log

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
func (it *BaseContentUpdateRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentUpdateRequest)
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
		it.Event = new(BaseContentUpdateRequest)
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
func (it *BaseContentUpdateRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentUpdateRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentUpdateRequest represents a UpdateRequest event raised by the BaseContent contract.
type BaseContentUpdateRequest struct {
	ObjectHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateRequest is a free log retrieval operation binding the contract event 0x1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef19.
//
// Solidity: e UpdateRequest(objectHash bytes32)
func (_BaseContent *BaseContentFilterer) FilterUpdateRequest(opts *bind.FilterOpts) (*BaseContentUpdateRequestIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return &BaseContentUpdateRequestIterator{contract: _BaseContent.contract, event: "UpdateRequest", logs: logs, sub: sub}, nil
}

// WatchUpdateRequest is a free log subscription operation binding the contract event 0x1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef19.
//
// Solidity: e UpdateRequest(objectHash bytes32)
func (_BaseContent *BaseContentFilterer) WatchUpdateRequest(opts *bind.WatchOpts, sink chan<- *BaseContentUpdateRequest) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentUpdateRequest)
				if err := _BaseContent.contract.UnpackLog(event, "UpdateRequest", log); err != nil {
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

// BaseContentSpaceABI is the input ABI used to generate the binding from.
const BaseContentSpaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createAccountLibrary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"address_KMS\",\"type\":\"address\"}],\"name\":\"createLibrary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content_space_description\",\"type\":\"string\"}],\"name\":\"setDescription\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createContentType\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"objectHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"object_hash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accessRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"content_space_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentTypeAddress\",\"type\":\"address\"}],\"name\":\"CreateContentType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"libraryAddress\",\"type\":\"address\"}],\"name\":\"CreateLibrary\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"groupAddress\",\"type\":\"address\"}],\"name\":\"CreateGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountAddress\",\"type\":\"address\"}],\"name\":\"CreateAccountLibrary\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"Commit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"UpdateRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestValidity\",\"type\":\"uint256\"}],\"name\":\"AccessRequest\",\"type\":\"event\"}]"

// BaseContentSpaceBin is the compiled bytecode used for deploying new contracts.
const BaseContentSpaceBin = `0x6060604052341561000f57600080fd5b604051614c47380380614c478339810160405280805160008054600160a060020a033216600160a060020a031991821681179092556001805490911690911790559091019050600381805161006892916020019061006f565b505061010a565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100b057805160ff19168380011785556100dd565b828001600101855582156100dd579182015b828111156100dd5782518255916020019190600101906100c2565b506100e99291506100ed565b5090565b61010791905b808211156100e957600081556001016100f3565b90565b614b2e806101196000396000f3006060604052600436106100e55763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100e757806306fdde03146101165780631291dad2146101a057806340b89f06146101b357806341c0e1b5146101d2578063575185ed146101e55780636d2e4b1b146101f85780637284e416146102175780638da5cb5b1461022a57806390c3f38f1461023d578063b8cfaf051461028e578063c287e0ed146102a1578063e02dd9c2146102b4578063f14fcbc8146102d9578063f1551887146102ef578063f2fde38b14610316575b005b34156100f257600080fd5b6100fa610335565b604051600160a060020a03909116815260200160405180910390f35b341561012157600080fd5b610129610344565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561016557808201518382015260200161014d565b50505050905090810190601f1680156101925780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101ab57600080fd5b6100fa6103e2565b34156101be57600080fd5b6100fa600160a060020a0360043516610424565b34156101dd57600080fd5b6100e561049b565b34156101f057600080fd5b6100fa6104c4565b341561020357600080fd5b6100e5600160a060020a036004351661052c565b341561022257600080fd5b6101296105cd565b341561023557600080fd5b6100fa610638565b341561024857600080fd5b6100e560046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061064795505050505050565b341561029957600080fd5b6100fa610679565b34156102ac57600080fd5b6100e56106de565b34156102bf57600080fd5b6102c7610730565b60405190815260200160405180910390f35b34156102e457600080fd5b6100e5600435610736565b34156102fa57600080fd5b61030261078c565b604051901515815260200160405180910390f35b341561032157600080fd5b6100e5600160a060020a03600435166107c9565b600054600160a060020a031681565b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103da5780601f106103af576101008083540402835291602001916103da565b820191906000526020600020905b8154815290600101906020018083116103bd57829003601f168201915b505050505081565b60007feeebced55b1b3d8371f41e9358ca416596184f5c40746f1c10b5224c1c154e0832604051600160a060020a03909116815260200160405180910390a190565b60008082610430610828565b600160a060020a039091168152602001604051809103906000f080151561045657600080fd5b90507f473c07a6d0228c4fb8fe2be3b4617c3b5fb7c0f8cd9ba4b67e8631844b9b657181604051600160a060020a03909116815260200160405180910390a192915050565b60015432600160a060020a039081169116146104b657600080fd5b600154600160a060020a0316ff5b6000806104cf610838565b604051809103906000f08015156104e557600080fd5b90507fa3b1fe71ae61bad8cffa485b230e24e518938f76182a30fa0d9979e7237ad15981604051600160a060020a03909116815260200160405180910390a18091505b5090565b60005432600160a060020a0390811691161461054757600080fd5b600160a060020a038116151561055c57600080fd5b600054600154600160a060020a039081169116141561059e576001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103da5780601f106103af576101008083540402835291602001916103da565b600154600160a060020a031681565b60015432600160a060020a0390811691161461066257600080fd5b6004818051610675929160200190610848565b5050565b6000806106846108c2565b604051809103906000f080151561069a57600080fd5b90507f9e69777f30c55126be256664fa7beff4b796ac32ebceab94df5071b0148017f881604051600160a060020a03909116815260200160405180910390a1919050565b60015432600160a060020a039081169116146106f957600080fd5b7f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef1960025460405190815260200160405180910390a1565b60025481565b60015432600160a060020a0390811691161461075157600080fd5b60028190557f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff8160405190815260200160405180910390a150565b60007fca8116fd70ecea0cf35f7c08eb31491cce5ebd6bafef5bd6426862df29840463600060405190815260200160405180910390a15060015b90565b60015432600160a060020a039081169116146107e457600080fd5b600160a060020a03811615156107f957600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6040516135b7806108ed83390190565b6040516107e680613ea483390190565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061088957805160ff19168380011785556108b6565b828001600101855582156108b6579182015b828111156108b657825182559160200191906001019061089b565b506105289291506108d2565b6040516104798061468a83390190565b6107c691905b8082111561052857600081556001016108d85600606060405260006010556040516020806135b78339810160405280805160008054600160a060020a03328116600160a060020a03199283168117845560018054841690911790556003805433831690841617905560088390556009839055600a92909255600f80549290931691161790555050613536806100816000396000f3006060604052600436106101a85763ffffffff60e060020a60003504166302d05d3f81146101aa5780630eaec2c5146101d9578063163083941461020c5780631cdbee5a146102315780631d0f4351146102505780631e35d8fa1461026f57806321770a841461028e5780632393553b146102a157806327af9e38146102b757806329d00219146102ca5780632c11f392146102e957806332eaf21b146102ff5780633fbb5d711461031257806341c0e1b514610331578063470750bb1461034457806349102e611461035757806363dab9d41461036a578063679a9a3c146103805780636934d45f1461039f5780636d2e4b1b146103c15780638cb13c2e146103e05780638da5cb5b146103f6578063952e464b1461040957806395a078e81461041f578063991a3a7c1461043e578063a90229d514610454578063bf81a0b61461047c578063c287e0ed146104e3578063c65bcbe2146104f6578063c9e8e72d14610509578063dc3c29c014610528578063e02dd9c214610547578063e5538fd21461055a578063f14fcbc81461056d578063f155188714610583578063f2fde38b14610596575b005b34156101b557600080fd5b6101bd6105b5565b604051600160a060020a03909116815260200160405180910390f35b34156101e457600080fd5b6101f8600160a060020a03600435166105c4565b604051901515815260200160405180910390f35b341561021757600080fd5b61021f6106af565b60405190815260200160405180910390f35b341561023c57600080fd5b6101bd600160a060020a03600435166106b5565b341561025b57600080fd5b6101a8600160a060020a03600435166106d0565b341561027a57600080fd5b6101bd600160a060020a036004351661076b565b341561029957600080fd5b61021f61089e565b34156102ac57600080fd5b6101bd6004356108a4565b34156102c257600080fd5b6101bd6108cc565b34156102d557600080fd5b6101f8600160a060020a03600435166108db565b34156102f457600080fd5b6101bd6004356109c3565b341561030a57600080fd5b6101bd6109d1565b341561031d57600080fd5b61021f600160a060020a03600435166109e0565b341561033c57600080fd5b6101a86109f2565b341561034f57600080fd5b61021f610a1b565b341561036257600080fd5b6101f8610a21565b341561037557600080fd5b6101bd600435610cf8565b341561038b57600080fd5b6101a8600160a060020a0360043516610d45565b34156103aa57600080fd5b61021f600160a060020a0360043516602435610de0565b34156103cc57600080fd5b6101a8600160a060020a0360043516610fea565b34156103eb57600080fd5b6101bd600435611071565b341561040157600080fd5b6101bd61107f565b341561041457600080fd5b6101bd60043561108e565b341561042a57600080fd5b6101f8600160a060020a036004351661109c565b341561044957600080fd5b6101bd600435611179565b341561045f57600080fd5b6101a8600160a060020a0360043581169060243516604435611187565b341561048757600080fd5b6101f860048035600160a060020a031690602480351515919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061126795505050505050565b34156104ee57600080fd5b6101a86115c9565b341561050157600080fd5b61021f61161b565b341561051457600080fd5b6101a8600160a060020a0360043516611621565b341561053357600080fd5b6101a8600160a060020a036004351661165e565b341561055257600080fd5b61021f6116f9565b341561056557600080fd5b61021f6116ff565b341561057857600080fd5b6101a8600435611705565b341561058e57600080fd5b6101f861175b565b34156105a157600080fd5b6101a8600160a060020a03600435166117e2565b600054600160a060020a031681565b6000806000806000600480549050600014156105e357600194506106a6565b5060005b6004548110156106a15760048054829081106105ff57fe5b600091825260209091200154600160a060020a0316935083156106995783915081600160a060020a03166395a078e88760405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561066d57600080fd5b5af1151561067a57600080fd5b50505060405180519350506001831515141561069957600194506106a6565b6001016105e7565b600094505b50505050919050565b60105481565b600c60205260009081526040902054600160a060020a031681565b60015432600160a060020a039081169116146106eb57600080fd5b60068054600181016106fd8382611834565b5060009182526020909120018054600160a060020a031916600160a060020a038316179055600a805460010190557f3a94857e4393737f73edb175a7d0c195c7f635d9ae995e12740616ec55c9d41181604051600160a060020a03909116815260200160405180910390a150565b6000806000610779326105c4565b15156107cc577f23de2adc3e22f171f66b3e5a333e17feb9dc30ba9570933bd259cb6c13ef7ab7606532604051918252600160a060020a031660208201526040908101905180910390a160009250610897565b6107d461185d565b604051809103906000f08015156107ea57600080fd5b915081905080600160a060020a031663693489678560405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b151561083d57600080fd5b5af1151561084a57600080fd5b5050507f3981e74ab81857b375ec391a4f7c31ee89462cd927de6d8fbdb98f77da009c568285604051600160a060020a039283168152911660208201526040908101905180910390a18192505b5050919050565b60095481565b60048054829081106108b257fe5b600091825260209091200154600160a060020a0316905081565b600354600160a060020a031681565b6001546000908190819081908190600160a060020a038781169116141561090557600194506106a6565b5060005b6005548110156106a157600580548290811061092157fe5b600091825260209091200154600160a060020a0316935083156109bb5783915081600160a060020a03166395a078e88760405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561098f57600080fd5b5af1151561099c57600080fd5b5050506040518051935050600183151514156109bb57600194506106a6565b600101610909565b60068054829081106108b257fe5b600f54600160a060020a031681565b600d6020526000908152604090205481565b60015432600160a060020a03908116911614610a0d57600080fd5b600154600160a060020a0316ff5b60085481565b600033808280600160a060020a03328116908416638da5cb5b6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610a6857600080fd5b5af11515610a7557600080fd5b50505060405180519050600160a060020a0316141515610a985760009450610cf1565b6009541515610bcd5782600160a060020a031663545b35396040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610ade57600080fd5b5af11515610aeb57600080fd5b5050506040518051925050600160a060020a03831663708f53ab60008460405160e060020a63ffffffff8516028152600481019290925260ff166024820152604401602060405180830381600087803b1515610b4657600080fd5b5af11515610b5357600080fd5b505050604051805190509050610b698382610de0565b507f8a6b83b9f4a16d1fd3ec333d40e20591cfb3ad7c24ecbfbc9a5c5b1a6445cab4846001604051600160a060020a039092168252151560208201526060604080830182905260009183019190915260a0909101905180910390a160019450610cf1565b600160a060020a03841660009081526011602052604090205415610bf45760009450610cf1565b600e546010541015610c405783600e601054815481101515610c1257fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055610c78565b600e805460018101610c528382611834565b5060009182526020909120018054600160a060020a031916600160a060020a0386161790555b60108054600160a060020a038616600090815260116020526040908190206001928301905582549091019091557f0588a34cf0de4e025d359c89ca4bacbcbf175440909952d91c814412d9da996a908590329051600160a060020a039283168152911660208201526040908101905180910390a1600194505b5050505090565b600060105460001480610d0d57508160105411155b15610d1a57506000610d40565b600e805483908110610d2857fe5b600091825260209091200154600160a060020a031690505b919050565b60015432600160a060020a03908116911614610d6057600080fd5b6004805460018101610d728382611834565b5060009182526020909120018054600160a060020a031916600160a060020a0383161790556008805460010190557f218673669018c25b89bfbf1b58d0075e37c8847ef16e707b92355b7833e97d6181604051600160a060020a03909116815260200160405180910390a150565b6000806000806000851115610fdc5785925082600160a060020a03166307b27aca6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610e2f57600080fd5b5af11515610e3c57600080fd5b505050604051805191505080851115610e5757809150610e5b565b8491505b82600160a060020a0316638da5cb5b6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610e9857600080fd5b5af11515610ea557600080fd5b5050506040518051600160a060020a0316905082156108fc0283604051600060405180830381858888f193505050501515610edf57600080fd5b82600160a060020a031663f79444718360405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610f2557600080fd5b5af11515610f3257600080fd5b50505060405180519050507fd8ad32fc5c165753d9da5f44f42ef7daf8002336cab4bdd13b7e58f48b451e3d83600160a060020a0316638da5cb5b6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610f9b57600080fd5b5af11515610fa857600080fd5b5050506040518051905083604051600160a060020a03909216825260208201526040908101905180910390a1819350610fe1565b600093505b50505092915050565b60005432600160a060020a0390811691161461100557600080fd5b600160a060020a038116151561101a57600080fd5b600054600154600160a060020a039081169116141561104f5760018054600160a060020a031916600160a060020a0383161790555b60008054600160a060020a031916600160a060020a0392909216919091179055565b600e8054829081106108b257fe5b600154600160a060020a031681565b60058054829081106108b257fe5b6000806000806000600680549050600014156110bb57600194506106a6565b5060005b6006548110156106a15760068054829081106110d757fe5b600091825260209091200154600160a060020a0316935083156111715783915081600160a060020a03166395a078e88760405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561114557600080fd5b5af1151561115257600080fd5b50505060405180519350506001831515141561117157600194506106a6565b6001016110bf565b60078054829081106108b257fe5b60015432600160a060020a039081169116146111a257600080fd5b60078054600181016111b48382611834565b5060009182526020808320919091018054600160a060020a0319908116600160a060020a03888116918217909355600b805460010190558452600c835260408085208054909216928716929092179055600d909152908190208290557f62c1f342f2a52ee6ac3eeef54df8c836fb31869e9be7c5e5394c042d131c9ecc9084908490849051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60008060008060008060008061127c326108db565b151561128b57600097506115bb565b600160a060020a038b16600090815260116020526040902054600e805460001990920198508c975090889081106112be57fe5b600091825260208083209091018054600160a060020a031916905560108054600019018155600160a060020a038e1683526011909152604082209190915554879011156113b557600e60105481548110151561131657fe5b600091825260209091200154600e8054600160a060020a03909216965086918990811061133f57fe5b906000526020600020900160006101000a815481600160a060020a030219169083600160a060020a03160217905550600e60105481548110151561137f57fe5b600091825260208083209091018054600160a060020a0319169055600160a060020a038716825260119052604090206001880190555b85600160a060020a03166327c1c21d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156113f257600080fd5b5af115156113ff57600080fd5b505050604051805194505060008413156115b65785600160a060020a031663545b35396040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561145057600080fd5b5af1151561145d57600080fd5b505050604051805193505060018a1515141561147c5760009150611487565b836001016000190291505b85600160a060020a031663708f53ab838560405160e060020a63ffffffff8516028152600481019290925260ff166024820152604401602060405180830381600087803b15156114d657600080fd5b5af115156114e357600080fd5b5050506040518051905090506114f98682610de0565b507f8a6b83b9f4a16d1fd3ec333d40e20591cfb3ad7c24ecbfbc9a5c5b1a6445cab48b8b8b604051600160a060020a0384168152821515602082015260606040820181815290820183818151815260200191508051906020019080838360005b83811015611571578082015183820152602001611559565b50505050905090810190601f16801561159e5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a1600197506115bb565b600097505b505050505050509392505050565b60015432600160a060020a039081169116146115e457600080fd5b7f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef1960025460405190815260200160405180910390a1565b600b5481565b60015432600160a060020a0390811691161461163c57600080fd5b600f8054600160a060020a031916600160a060020a0392909216919091179055565b60015432600160a060020a0390811691161461167957600080fd5b600580546001810161168b8382611834565b5060009182526020909120018054600160a060020a031916600160a060020a0383161790556009805460010190557f1b88a571cc8ac2e87512f05648e79d184f5cc0cbb2889bc487c41f8b9a3202eb81604051600160a060020a03909116815260200160405180910390a150565b60025481565b600a5481565b60015432600160a060020a0390811691161461172057600080fd5b60028190557f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff8160405190815260200160405180910390a150565b60006117663261109c565b156117a7577fca8116fd70ecea0cf35f7c08eb31491cce5ebd6bafef5bd6426862df29840463600060405190815260200160405180910390a15060016117df565b7fca8116fd70ecea0cf35f7c08eb31491cce5ebd6bafef5bd6426862df29840463606960405190815260200160405180910390a15060005b90565b60015432600160a060020a039081169116146117fd57600080fd5b600160a060020a038116151561181257600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b8154818355818115116118585760008381526020902061185891810190830161186d565b505050565b604051611c7f8061188c83390190565b6117df91905b808211156118875760008155600101611873565b509056006060604090815260008054600160a060020a03328116600160a060020a03199283168117845560018054841690911790556006805433831693169290921791829055600019600855600b839055600c929092557fc3decc188980e855666b70498ca85e8fa284d97d30483d828fa126f7303d7d199291169051600160a060020a03909116815260200160405180910390a1611be08061009f6000396000f3006060604052600436106101875763ffffffff60e060020a60003504166302d05d3f81146101895780630779564b146101b857806307b27aca146101dd5780631a735f18146101f05780632310167f1461024257806327c1c21d1461025557806332eaf21b1461026857806336ebffca1461027b57806341c0e1b51461028e5780635267db44146102a1578063545b3539146102b75780635cc4aa9b146102e057806364ade32b1461030557806369348967146103185780636d2e4b1b14610337578063708f53ab14610356578063879fe48f146103725780638da5cb5b146104095780638f7792011461041c578063a1ff106e1461042f578063b65240891461053f578063b816f51314610555578063c287e0ed14610568578063c68c1d1f1461057b578063c9e8e72d146105c9578063e02dd9c2146105e8578063e5385303146105fb578063ee56d7671461061a578063f14fcbc8146106ba578063f2fde38b146106d0578063f4d9bae8146106ef578063f794447114610705578063f94f00141461071b575b005b341561019457600080fd5b61019c61072e565b604051600160a060020a03909116815260200160405180910390f35b34156101c357600080fd5b6101cb61073d565b60405190815260200160405180910390f35b34156101e857600080fd5b6101cb610743565b34156101fb57600080fd5b61020660043561074d565b6040518084600160a060020a0316600160a060020a031681526020018381526020018260000b60000b8152602001935050505060405180910390f35b341561024d57600080fd5b61019c610778565b341561026057600080fd5b6101cb610787565b341561027357600080fd5b61019c61078d565b341561028657600080fd5b61019c61079c565b341561029957600080fd5b6101876107ab565b34156102ac57600080fd5b6101cb6004356107d4565b34156102c257600080fd5b6102ca610873565b60405160ff909116815260200160405180910390f35b6102f160043560243560443561087c565b604051901515815260200160405180910390f35b341561031057600080fd5b6101cb610a12565b341561032357600080fd5b610187600160a060020a0360043516610a18565b341561034257600080fd5b610187600160a060020a0360043516610c00565b341561036157600080fd5b6101cb60043560ff60243516610c87565b341561037d57600080fd5b6101cb6004803560ff1690604460248035908101908301358060208082020160405190810160405280939291908181526020018383602002808284378201915050505050509190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843750949650610dc095505050505050565b341561041457600080fd5b61019c610f28565b341561042757600080fd5b6101cb610f37565b6102f16004803560ff169060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843750949650610f3d95505050505050565b341561054a57600080fd5b610187600435611503565b341561056057600080fd5b61019c611531565b341561057357600080fd5b610187611540565b6102f16004803560ff169060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061159295505050505050565b34156105d457600080fd5b610187600160a060020a03600435166116d5565b34156105f357600080fd5b6101cb611712565b341561060657600080fd5b610187600160a060020a0360043516611718565b341561062557600080fd5b6102f16004803590602480351515919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061179595505050505050565b34156106c557600080fd5b610187600435611a37565b34156106db57600080fd5b610187600160a060020a0360043516611a8d565b34156106fa57600080fd5b6101cb600435611adf565b341561071057600080fd5b6101cb600435611b3d565b341561072657600080fd5b6101cb611b6e565b600054600160a060020a031681565b600a5481565b600b54600a540390565b600d60205260009081526040812080546001820154600290920154600160a060020a03909116920b83565b600554600160a060020a031681565b60085481565b600454600160a060020a031681565b600354600160a060020a031681565b60015432600160a060020a039081169116146107c657600080fd5b600154600160a060020a0316ff5b60015460009032600160a060020a03908116911614801561080c5750600082128061080c575060008213801561080c57506000600854125b156108175760088290555b60065433600160a060020a03908116911614156108345760088290555b7fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a260085460405190815260200160405180910390a1506008545b919050565b60095460ff1681565b60055460009060019082908190600160a060020a0316156108fb57600554600160a060020a0316915081636b2d13248860405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156108e257600080fd5b5af115156108ef57600080fd5b50505060405180519350505b506000868152600d602052604090208054600160a060020a0316158015906109305750805433600160a060020a039081169116145b156109b2576002810154600090810b900b151561097e5733600160a060020a03166108fc82600101549081150290604051600060405180830381858888f19350505050151561097e57600080fd5b6000878152600d602052604081208054600160a060020a03191681556001810191909155600201805460ff191690556109b7565b600092505b7f7f1f4b28434ce7beab4983e64a8b5bb96e195a67029fdaff925028aec57fbc6b87878786604051938452602084019290925260408084019190915290151560608301526080909101905180910390a1509095945050505050565b60075481565b60015460009032600160a060020a03908116911614610a3657600080fd5b5060038054600160a060020a031916600160a060020a03838116919091179091556006541680631cdbee5a8360405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610aa057600080fd5b5af11515610aad57600080fd5b505050604051805160058054600160a060020a031916600160a060020a0392831617905582169050633fbb5d718360405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610b1957600080fd5b5af11515610b2657600080fd5b5050506040518051600a5550600160a060020a0381166332eaf21b6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610b6f57600080fd5b5af11515610b7c57600080fd5b505050604051805160048054600160a060020a031916600160a060020a03928316179055600554600a547f1c158b764edca9bf03a401a388355220a7395febb3c011729d8060572f065627935085929190911690604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15050565b60005432600160a060020a03908116911614610c1b57600080fd5b600160a060020a0381161515610c3057600080fd5b600054600154600160a060020a0390811691161415610c655760018054600160a060020a031916600160a060020a0383161790555b60008054600160a060020a031916600160a060020a0392909216919091179055565b600154600090819081908190819032600160a060020a0390811691161480610cbd575060065433600160a060020a039081169116145b15610db1576009805460ff191660ff881617905560055460009450600160a060020a03161515610d115760095487935060ff166064148015610cfd575086155b15610d0c57600b54600a540393505b610d9f565b600554600160a060020a0316915081633513a8058860405160e060020a63ffffffff841602815260048101919091526024016040805180830381600087803b1515610d5b57600080fd5b5af11515610d6857600080fd5b5050506040518051906020018051919450909150506000811215610d9b57861515610d0c57600b54600a54039350610d9f565b8093505b610da8836107d4565b50839450610db6565b600094505b5050505092915050565b6007546005546000919082908190600160a060020a031615610edf57600554600160a060020a031691508163f8117ca28888886040518463ffffffff1660e060020a028152600401808460ff1660ff1681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015610e55578082015183820152602001610e3d565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610e94578082015183820152602001610e7c565b5050505090500195505050505050602060405180830381600087803b1515610ebb57600080fd5b5af11515610ec857600080fd5b505050604051805191505060008112610edf578092505b7fa58326ee5bb617cb8b4f0d0f5f557c469d2d05d7a738f777037deda9c724b370878460405160ff909216825260208201526040908101905180910390a1509095945050505050565b600154600160a060020a031681565b600c5481565b6000806000610f4a611b74565b600c80546001019055600654600160a060020a0316925060008080856395a078e83260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610faa57600080fd5b5af11515610fb757600080fd5b50505060405180511515905061103457600080516020611b958339815191526069600c548e6000604051938452602084019290925260ff16604080840191909152606083019190915260c060808301819052600090830181905261010060a08401819052830152610140909101905180910390a1600096506114f4565b60015432600160a060020a039081169116146110cc576110558c8a8a610dc0565b9450843410156110cc57600080516020611b958339815191526067600c548e6000604051938452602084019290925260ff16604080840191909152606083019190915260c060808301819052600090830181905261010060a08401819052830152610140909101905180910390a1600096506114f4565b60606040519081016040908152600160a060020a0333168252346020808401919091526000828401819052600c548152600d90915220909450849081518154600160a060020a031916600160a060020a03919091161781556020820151816001015560408201516002909101805460009290920b60ff1660ff1990921691909117905550600554600160a060020a0316156112d857600554600c54600160a060020a039091169350839063123e0e80908e8c8c6040518563ffffffff1660e060020a028152600401808581526020018460ff1660ff1681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156111e75780820151838201526020016111cf565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561122657808201518382015260200161120e565b505050509050019650505050505050602060405180830381600087803b151561124e57600080fd5b5af1151561125b57600080fd5b505050604051805192505081156112d857600080516020611b9583398151915282600c548e6000604051938452602084019290925260ff16604080840191909152606083019190915260c060808301819052600090830181905261010060a08401819052830152610140909101905180910390a1600096506114f4565b600080516020611b958339815191526000600c548e6002548f8f6040518681526020810186905260ff851660408201526060810184905260c0608082018181529060a0830190830185818151815260200191508051906020019080838360005b83811015611350578082015183820152602001611338565b50505050905090810190601f16801561137d5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156113b357808201518382015260200161139b565b50505050905090810190601f1680156113e05780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390a15060005b885181101561146c5788818151811061140c57fe5b9060200190602002015115611464577f515e0a48b385fce2a8e4d9f169a97c4f6ea669a752358f5e6ab37cc3c2e84c3889828151811061144857fe5b9060200190602002015160405190815260200160405180910390a15b6001016113f7565b5060005b87518110156114ef5788818151811061148557fe5b90602001906020020151156114e7577fb6e3239e521a6c66920ae634f8e921a37e6991d520ac44d52f8516397f41b6848882815181106114c157fe5b90602001906020020151604051600160a060020a03909116815260200160405180910390a15b600101611470565b600196505b50505050505095945050505050565b60015432600160a060020a03908116911614801561152357506000600854125b1561152e57600a8190555b50565b600654600160a060020a031681565b60015432600160a060020a0390811691161461155b57600080fd5b7f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef1960025460405190815260200160405180910390a1565b6001546000908190819032600160a060020a039081169116146115b457600080fd5b836040518082805190602001908083835b602083106115e45780518252601f1990920191602091820191016115c5565b6001836020036101000a0380198251168184511617909252505050919091019250604091505051506116199050600186610c87565b50600060085413156116c857600654600160a060020a03169150816349102e616040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561166757600080fd5b5af1151561167457600080fd5b5050506040518051905090507f9625d8ae00bbc11bf4dbac17532594047d634ea66e1baa95d4974d086f261c30858260405160ff9092168252151560208201526040908101905180910390a18092506116cd565b600092505b505092915050565b60015432600160a060020a039081169116146116f057600080fd5b60048054600160a060020a031916600160a060020a0392909216919091179055565b60025481565b60015432600160a060020a0390811691161461173357600080fd5b60058054600160a060020a031916600160a060020a0383811691909117918290557fa6f2e38f0cfebf27212317fced3ac40bc62e00bd33f38d69603710740c69acb79116604051600160a060020a03909116815260200160405180910390a150565b600154600090819033600160a060020a039081169116148015906117c8575060045433600160a060020a03908116911614155b156117d65760009150611a2e565b506000858152600d602052604090208054600160a060020a031615156117ff5760009150611a2e565b6002810154600090810b900b1515611a29578415156118c55780546001820154600160a060020a039091169080156108fc0290604051600060405180830381858888f19350505050151561185257600080fd5b60028101805460ff191660ff1790557f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b9127188660006040519182521515602082015260806040808301829052600091830182905260c060608401819052830191909152610100909101905180910390a1611a29565b6001805490820154600160a060020a039091169080156108fc0290604051600060405180830381858888f19350505050151561190057600080fd5b60028101805460ff191660019081179091557f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b91271890879086866040518481528315156020820152608060408201818152906060830190830185818151815260200191508051906020019080838360005b8381101561198757808201518382015260200161196f565b50505050905090810190601f1680156119b45780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156119ea5780820151838201526020016119d2565b50505050905090810190601f168015611a175780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a15b600191505b50949350505050565b60015432600160a060020a03908116911614611a5257600080fd5b60028190557f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff8160405190815260200160405180910390a150565b60015432600160a060020a03908116911614611aa857600080fd5b600160a060020a0381161515611abd57600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b60015460009032600160a060020a03908116911614611afd57600080fd5b60078290557f4114f8ef80b6de2161db580cbefa14e1892d15d3ebe2062c9914e4a5773114a38260405190815260200160405180910390a1505060075490565b60065460009033600160a060020a03908116911614611b5e5750600061086e565b50600b8054820190819055919050565b600b5481565b6060604051908101604090815260008083526020830181905290820152905600272aa3d0efd05809b06d40e95792681d6b844a29fe97cd703cfd5eee27510740a165627a7a72305820f010d72b4ef3212cb91349cd2c5958f806c34e964fa1b0e1325044ebfbeef0500029a165627a7a72305820f540d994d3ac2641583ad9a285b176adee88783dac3139271f046465cdcaa89100296060604052341561000f57600080fd5b60008054600160a060020a03328116600160a060020a0319928316811780855560018054909416909117835581168352600360209081526040808520805460ff19908116861790915585549093168552600290915290922080549092161790556107688061007e6000396000f3006060604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100c657806308ae4b0c146100f55780630ae5e7391461012857806341c0e1b51461014757806342e7ba7b1461015a5780636d2e4b1b1461017957806375861a951461019857806385e68531146101b75780638da5cb5b146101d657806395a078e8146101e9578063cdb849b714610208578063f2fde38b14610227578063fdff9b4d14610246575b005b34156100d157600080fd5b6100d9610265565b604051600160a060020a03909116815260200160405180910390f35b341561010057600080fd5b610114600160a060020a0360043516610274565b604051901515815260200160405180910390f35b341561013357600080fd5b6100c4600160a060020a0360043516610289565b341561015257600080fd5b6100c461035e565b341561016557600080fd5b610114600160a060020a0360043516610387565b341561018457600080fd5b6100c4600160a060020a03600435166103aa565b34156101a357600080fd5b6100c4600160a060020a036004351661044b565b34156101c257600080fd5b6100c4600160a060020a03600435166104cb565b34156101e157600080fd5b6100d96105b8565b34156101f457600080fd5b610114600160a060020a03600435166105c7565b341561021357600080fd5b6100c4600160a060020a03600435166105ea565b341561023257600080fd5b6100c4600160a060020a03600435166106c8565b341561025157600080fd5b610114600160a060020a0360043516610727565b600054600160a060020a031681565b60026020526000908152604090205460ff1681565b600160a060020a03331660009081526003602052604090205460ff1615156001141561031657600160a060020a03811660009081526002602052604090819020805460ff191660011790557fb251eb052afc73ffd02ffe85ad79990a8b3fed60d76dbc2fa2fdd7123dffd91490829051600160a060020a03909116815260200160405180910390a161035b565b7f23de2adc3e22f171f66b3e5a333e17feb9dc30ba9570933bd259cb6c13ef7ab7601482604051918252600160a060020a031660208201526040908101905180910390a15b50565b60015432600160a060020a0390811691161461037957600080fd5b600154600160a060020a0316ff5b600160a060020a031660009081526003602052604090205460ff16151560011490565b60005432600160a060020a039081169116146103c557600080fd5b600160a060020a03811615156103da57600080fd5b600054600154600160a060020a039081169116141561041c576001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60015432600160a060020a0390811691161461046657600080fd5b600160a060020a03811660009081526003602052604090819020805460ff191660011790557f93bcaab179551bde429187645251f8e1fb8ac85801fcb1cf91eb2c9043d6111790829051600160a060020a03909116815260200160405180910390a150565b600160a060020a03331660009081526003602052604090205460ff16151560011480610508575080600160a060020a031633600160a060020a0316145b1561057157600160a060020a03811660009081526002602052604090819020805460ff191690557f745cd29407db644ed93e3ceb61cbcab96d1dfb496989ac5d5bf514fc5a9fab9c90829051600160a060020a03909116815260200160405180910390a161035b565b7f23de2adc3e22f171f66b3e5a333e17feb9dc30ba9570933bd259cb6c13ef7ab7601582604051918252600160a060020a031660208201526040908101905180910390a150565b600154600160a060020a031681565b600160a060020a031660009081526002602052604090205460ff16151560011490565b60005433600160a060020a0390811691161480610618575080600160a060020a031633600160a060020a0316145b1561068157600160a060020a03811660009081526003602052604090819020805460ff191690557f2d6aa1a9629d125e23a0cf692cda7cd6795dff1652eedd4673b38ec31e387b9590829051600160a060020a03909116815260200160405180910390a161035b565b7f23de2adc3e22f171f66b3e5a333e17feb9dc30ba9570933bd259cb6c13ef7ab7600b82604051918252600160a060020a031660208201526040908101905180910390a150565b60015432600160a060020a039081169116146106e357600080fd5b600160a060020a03811615156106f857600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60036020526000908152604090205460ff16815600a165627a7a72305820140e959db4d9147a3e337f1842cf9a7aba7c25a281771dae84fa7dcc8ed5d81f0029606060405260008054600160a060020a03328116600160a060020a031992831681178455600180548416909117905560038054339092169190921617905561042c90819061004d90396000f3006060604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100a557806341c0e1b5146100d45780636d2e4b1b146100e75780638da5cb5b14610106578063af570c0414610119578063c287e0ed1461012c578063e02dd9c21461013f578063f14fcbc814610164578063f15518871461017a578063f2fde38b146101a1575b005b34156100b057600080fd5b6100b86101c0565b604051600160a060020a03909116815260200160405180910390f35b34156100df57600080fd5b6100a36101cf565b34156100f257600080fd5b6100a3600160a060020a03600435166101f8565b341561011157600080fd5b6100b8610299565b341561012457600080fd5b6100b86102a8565b341561013757600080fd5b6100a36102b7565b341561014a57600080fd5b610152610309565b60405190815260200160405180910390f35b341561016f57600080fd5b6100a360043561030f565b341561018557600080fd5b61018d610365565b604051901515815260200160405180910390f35b34156101ac57600080fd5b6100a3600160a060020a03600435166103a1565b600054600160a060020a031681565b60015432600160a060020a039081169116146101ea57600080fd5b600154600160a060020a0316ff5b60005432600160a060020a0390811691161461021357600080fd5b600160a060020a038116151561022857600080fd5b600054600154600160a060020a039081169116141561026a576001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031681565b600354600160a060020a031681565b60015432600160a060020a039081169116146102d257600080fd5b7f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef1960025460405190815260200160405180910390a1565b60025481565b60015432600160a060020a0390811691161461032a57600080fd5b60028190557f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff8160405190815260200160405180910390a150565b60007fca8116fd70ecea0cf35f7c08eb31491cce5ebd6bafef5bd6426862df29840463600060405190815260200160405180910390a150600190565b60015432600160a060020a039081169116146103bc57600080fd5b600160a060020a03811615156103d157600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582002461abc581e32b9ded636540e838ffae65caf9c71a4215780265a749f0276410029a165627a7a723058205ae55b7b3d2497a8e0db51881542374bbd1a23ca5f958da2e862d09c42c7b85f0029`

// DeployBaseContentSpace deploys a new Ethereum contract, binding an instance of BaseContentSpace to it.
func DeployBaseContentSpace(auth *bind.TransactOpts, backend bind.ContractBackend, content_space_name string) (common.Address, *types.Transaction, *BaseContentSpace, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseContentSpaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaseContentSpaceBin), backend, content_space_name)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseContentSpace{BaseContentSpaceCaller: BaseContentSpaceCaller{contract: contract}, BaseContentSpaceTransactor: BaseContentSpaceTransactor{contract: contract}, BaseContentSpaceFilterer: BaseContentSpaceFilterer{contract: contract}}, nil
}

// BaseContentSpace is an auto generated Go binding around an Ethereum contract.
type BaseContentSpace struct {
	BaseContentSpaceCaller     // Read-only binding to the contract
	BaseContentSpaceTransactor // Write-only binding to the contract
	BaseContentSpaceFilterer   // Log filterer for contract events
}

// BaseContentSpaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseContentSpaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseContentSpaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseContentSpaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseContentSpaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseContentSpaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseContentSpaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseContentSpaceSession struct {
	Contract     *BaseContentSpace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseContentSpaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseContentSpaceCallerSession struct {
	Contract *BaseContentSpaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BaseContentSpaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseContentSpaceTransactorSession struct {
	Contract     *BaseContentSpaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BaseContentSpaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseContentSpaceRaw struct {
	Contract *BaseContentSpace // Generic contract binding to access the raw methods on
}

// BaseContentSpaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseContentSpaceCallerRaw struct {
	Contract *BaseContentSpaceCaller // Generic read-only contract binding to access the raw methods on
}

// BaseContentSpaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseContentSpaceTransactorRaw struct {
	Contract *BaseContentSpaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseContentSpace creates a new instance of BaseContentSpace, bound to a specific deployed contract.
func NewBaseContentSpace(address common.Address, backend bind.ContractBackend) (*BaseContentSpace, error) {
	contract, err := bindBaseContentSpace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseContentSpace{BaseContentSpaceCaller: BaseContentSpaceCaller{contract: contract}, BaseContentSpaceTransactor: BaseContentSpaceTransactor{contract: contract}, BaseContentSpaceFilterer: BaseContentSpaceFilterer{contract: contract}}, nil
}

// NewBaseContentSpaceCaller creates a new read-only instance of BaseContentSpace, bound to a specific deployed contract.
func NewBaseContentSpaceCaller(address common.Address, caller bind.ContractCaller) (*BaseContentSpaceCaller, error) {
	contract, err := bindBaseContentSpace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceCaller{contract: contract}, nil
}

// NewBaseContentSpaceTransactor creates a new write-only instance of BaseContentSpace, bound to a specific deployed contract.
func NewBaseContentSpaceTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseContentSpaceTransactor, error) {
	contract, err := bindBaseContentSpace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceTransactor{contract: contract}, nil
}

// NewBaseContentSpaceFilterer creates a new log filterer instance of BaseContentSpace, bound to a specific deployed contract.
func NewBaseContentSpaceFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseContentSpaceFilterer, error) {
	contract, err := bindBaseContentSpace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceFilterer{contract: contract}, nil
}

// bindBaseContentSpace binds a generic wrapper to an already deployed contract.
func bindBaseContentSpace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseContentSpaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseContentSpace *BaseContentSpaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseContentSpace.Contract.BaseContentSpaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseContentSpace *BaseContentSpaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.BaseContentSpaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseContentSpace *BaseContentSpaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.BaseContentSpaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseContentSpace *BaseContentSpaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseContentSpace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseContentSpace *BaseContentSpaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseContentSpace *BaseContentSpaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.contract.Transact(opts, method, params...)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseContentSpace *BaseContentSpaceCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContentSpace.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseContentSpace *BaseContentSpaceSession) Creator() (common.Address, error) {
	return _BaseContentSpace.Contract.Creator(&_BaseContentSpace.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseContentSpace *BaseContentSpaceCallerSession) Creator() (common.Address, error) {
	return _BaseContentSpace.Contract.Creator(&_BaseContentSpace.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() constant returns(string)
func (_BaseContentSpace *BaseContentSpaceCaller) Description(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BaseContentSpace.contract.Call(opts, out, "description")
	return *ret0, err
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() constant returns(string)
func (_BaseContentSpace *BaseContentSpaceSession) Description() (string, error) {
	return _BaseContentSpace.Contract.Description(&_BaseContentSpace.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() constant returns(string)
func (_BaseContentSpace *BaseContentSpaceCallerSession) Description() (string, error) {
	return _BaseContentSpace.Contract.Description(&_BaseContentSpace.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaseContentSpace *BaseContentSpaceCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BaseContentSpace.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaseContentSpace *BaseContentSpaceSession) Name() (string, error) {
	return _BaseContentSpace.Contract.Name(&_BaseContentSpace.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaseContentSpace *BaseContentSpaceCallerSession) Name() (string, error) {
	return _BaseContentSpace.Contract.Name(&_BaseContentSpace.CallOpts)
}

// ObjectHash is a free data retrieval call binding the contract method 0xe02dd9c2.
//
// Solidity: function objectHash() constant returns(bytes32)
func (_BaseContentSpace *BaseContentSpaceCaller) ObjectHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseContentSpace.contract.Call(opts, out, "objectHash")
	return *ret0, err
}

// ObjectHash is a free data retrieval call binding the contract method 0xe02dd9c2.
//
// Solidity: function objectHash() constant returns(bytes32)
func (_BaseContentSpace *BaseContentSpaceSession) ObjectHash() ([32]byte, error) {
	return _BaseContentSpace.Contract.ObjectHash(&_BaseContentSpace.CallOpts)
}

// ObjectHash is a free data retrieval call binding the contract method 0xe02dd9c2.
//
// Solidity: function objectHash() constant returns(bytes32)
func (_BaseContentSpace *BaseContentSpaceCallerSession) ObjectHash() ([32]byte, error) {
	return _BaseContentSpace.Contract.ObjectHash(&_BaseContentSpace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseContentSpace *BaseContentSpaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContentSpace.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseContentSpace *BaseContentSpaceSession) Owner() (common.Address, error) {
	return _BaseContentSpace.Contract.Owner(&_BaseContentSpace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseContentSpace *BaseContentSpaceCallerSession) Owner() (common.Address, error) {
	return _BaseContentSpace.Contract.Owner(&_BaseContentSpace.CallOpts)
}

// AccessRequest is a paid mutator transaction binding the contract method 0xf1551887.
//
// Solidity: function accessRequest() returns(bool)
func (_BaseContentSpace *BaseContentSpaceTransactor) AccessRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "accessRequest")
}

// AccessRequest is a paid mutator transaction binding the contract method 0xf1551887.
//
// Solidity: function accessRequest() returns(bool)
func (_BaseContentSpace *BaseContentSpaceSession) AccessRequest() (*types.Transaction, error) {
	return _BaseContentSpace.Contract.AccessRequest(&_BaseContentSpace.TransactOpts)
}

// AccessRequest is a paid mutator transaction binding the contract method 0xf1551887.
//
// Solidity: function accessRequest() returns(bool)
func (_BaseContentSpace *BaseContentSpaceTransactorSession) AccessRequest() (*types.Transaction, error) {
	return _BaseContentSpace.Contract.AccessRequest(&_BaseContentSpace.TransactOpts)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(object_hash bytes32) returns()
func (_BaseContentSpace *BaseContentSpaceTransactor) Commit(opts *bind.TransactOpts, object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "commit", object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(object_hash bytes32) returns()
func (_BaseContentSpace *BaseContentSpaceSession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.Commit(&_BaseContentSpace.TransactOpts, object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(object_hash bytes32) returns()
func (_BaseContentSpace *BaseContentSpaceTransactorSession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.Commit(&_BaseContentSpace.TransactOpts, object_hash)
}

// CreateAccountLibrary is a paid mutator transaction binding the contract method 0x1291dad2.
//
// Solidity: function createAccountLibrary() returns(address)
func (_BaseContentSpace *BaseContentSpaceTransactor) CreateAccountLibrary(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "createAccountLibrary")
}

// CreateAccountLibrary is a paid mutator transaction binding the contract method 0x1291dad2.
//
// Solidity: function createAccountLibrary() returns(address)
func (_BaseContentSpace *BaseContentSpaceSession) CreateAccountLibrary() (*types.Transaction, error) {
	return _BaseContentSpace.Contract.CreateAccountLibrary(&_BaseContentSpace.TransactOpts)
}

// CreateAccountLibrary is a paid mutator transaction binding the contract method 0x1291dad2.
//
// Solidity: function createAccountLibrary() returns(address)
func (_BaseContentSpace *BaseContentSpaceTransactorSession) CreateAccountLibrary() (*types.Transaction, error) {
	return _BaseContentSpace.Contract.CreateAccountLibrary(&_BaseContentSpace.TransactOpts)
}

// CreateContentType is a paid mutator transaction binding the contract method 0xb8cfaf05.
//
// Solidity: function createContentType() returns(address)
func (_BaseContentSpace *BaseContentSpaceTransactor) CreateContentType(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "createContentType")
}

// CreateContentType is a paid mutator transaction binding the contract method 0xb8cfaf05.
//
// Solidity: function createContentType() returns(address)
func (_BaseContentSpace *BaseContentSpaceSession) CreateContentType() (*types.Transaction, error) {
	return _BaseContentSpace.Contract.CreateContentType(&_BaseContentSpace.TransactOpts)
}

// CreateContentType is a paid mutator transaction binding the contract method 0xb8cfaf05.
//
// Solidity: function createContentType() returns(address)
func (_BaseContentSpace *BaseContentSpaceTransactorSession) CreateContentType() (*types.Transaction, error) {
	return _BaseContentSpace.Contract.CreateContentType(&_BaseContentSpace.TransactOpts)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x575185ed.
//
// Solidity: function createGroup() returns(address)
func (_BaseContentSpace *BaseContentSpaceTransactor) CreateGroup(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "createGroup")
}

// CreateGroup is a paid mutator transaction binding the contract method 0x575185ed.
//
// Solidity: function createGroup() returns(address)
func (_BaseContentSpace *BaseContentSpaceSession) CreateGroup() (*types.Transaction, error) {
	return _BaseContentSpace.Contract.CreateGroup(&_BaseContentSpace.TransactOpts)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x575185ed.
//
// Solidity: function createGroup() returns(address)
func (_BaseContentSpace *BaseContentSpaceTransactorSession) CreateGroup() (*types.Transaction, error) {
	return _BaseContentSpace.Contract.CreateGroup(&_BaseContentSpace.TransactOpts)
}

// CreateLibrary is a paid mutator transaction binding the contract method 0x40b89f06.
//
// Solidity: function createLibrary(address_KMS address) returns(address)
func (_BaseContentSpace *BaseContentSpaceTransactor) CreateLibrary(opts *bind.TransactOpts, address_KMS common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "createLibrary", address_KMS)
}

// CreateLibrary is a paid mutator transaction binding the contract method 0x40b89f06.
//
// Solidity: function createLibrary(address_KMS address) returns(address)
func (_BaseContentSpace *BaseContentSpaceSession) CreateLibrary(address_KMS common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.CreateLibrary(&_BaseContentSpace.TransactOpts, address_KMS)
}

// CreateLibrary is a paid mutator transaction binding the contract method 0x40b89f06.
//
// Solidity: function createLibrary(address_KMS address) returns(address)
func (_BaseContentSpace *BaseContentSpaceTransactorSession) CreateLibrary(address_KMS common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.CreateLibrary(&_BaseContentSpace.TransactOpts, address_KMS)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseContentSpace *BaseContentSpaceTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseContentSpace *BaseContentSpaceSession) Kill() (*types.Transaction, error) {
	return _BaseContentSpace.Contract.Kill(&_BaseContentSpace.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseContentSpace *BaseContentSpaceTransactorSession) Kill() (*types.Transaction, error) {
	return _BaseContentSpace.Contract.Kill(&_BaseContentSpace.TransactOpts)
}

// SetDescription is a paid mutator transaction binding the contract method 0x90c3f38f.
//
// Solidity: function setDescription(content_space_description string) returns()
func (_BaseContentSpace *BaseContentSpaceTransactor) SetDescription(opts *bind.TransactOpts, content_space_description string) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "setDescription", content_space_description)
}

// SetDescription is a paid mutator transaction binding the contract method 0x90c3f38f.
//
// Solidity: function setDescription(content_space_description string) returns()
func (_BaseContentSpace *BaseContentSpaceSession) SetDescription(content_space_description string) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.SetDescription(&_BaseContentSpace.TransactOpts, content_space_description)
}

// SetDescription is a paid mutator transaction binding the contract method 0x90c3f38f.
//
// Solidity: function setDescription(content_space_description string) returns()
func (_BaseContentSpace *BaseContentSpaceTransactorSession) SetDescription(content_space_description string) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.SetDescription(&_BaseContentSpace.TransactOpts, content_space_description)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_BaseContentSpace *BaseContentSpaceTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_BaseContentSpace *BaseContentSpaceSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.TransferCreatorship(&_BaseContentSpace.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_BaseContentSpace *BaseContentSpaceTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.TransferCreatorship(&_BaseContentSpace.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseContentSpace *BaseContentSpaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseContentSpace *BaseContentSpaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.TransferOwnership(&_BaseContentSpace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseContentSpace *BaseContentSpaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.TransferOwnership(&_BaseContentSpace.TransactOpts, newOwner)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xc287e0ed.
//
// Solidity: function updateRequest() returns()
func (_BaseContentSpace *BaseContentSpaceTransactor) UpdateRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "updateRequest")
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xc287e0ed.
//
// Solidity: function updateRequest() returns()
func (_BaseContentSpace *BaseContentSpaceSession) UpdateRequest() (*types.Transaction, error) {
	return _BaseContentSpace.Contract.UpdateRequest(&_BaseContentSpace.TransactOpts)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xc287e0ed.
//
// Solidity: function updateRequest() returns()
func (_BaseContentSpace *BaseContentSpaceTransactorSession) UpdateRequest() (*types.Transaction, error) {
	return _BaseContentSpace.Contract.UpdateRequest(&_BaseContentSpace.TransactOpts)
}

// BaseContentSpaceAccessRequestIterator is returned from FilterAccessRequest and is used to iterate over the raw logs and unpacked data for AccessRequest events raised by the BaseContentSpace contract.
type BaseContentSpaceAccessRequestIterator struct {
	Event *BaseContentSpaceAccessRequest // Event containing the contract specifics and raw log

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
func (it *BaseContentSpaceAccessRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentSpaceAccessRequest)
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
		it.Event = new(BaseContentSpaceAccessRequest)
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
func (it *BaseContentSpaceAccessRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentSpaceAccessRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentSpaceAccessRequest represents a AccessRequest event raised by the BaseContentSpace contract.
type BaseContentSpaceAccessRequest struct {
	RequestValidity *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccessRequest is a free log retrieval operation binding the contract event 0xca8116fd70ecea0cf35f7c08eb31491cce5ebd6bafef5bd6426862df29840463.
//
// Solidity: e AccessRequest(requestValidity uint256)
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterAccessRequest(opts *bind.FilterOpts) (*BaseContentSpaceAccessRequestIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "AccessRequest")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceAccessRequestIterator{contract: _BaseContentSpace.contract, event: "AccessRequest", logs: logs, sub: sub}, nil
}

// WatchAccessRequest is a free log subscription operation binding the contract event 0xca8116fd70ecea0cf35f7c08eb31491cce5ebd6bafef5bd6426862df29840463.
//
// Solidity: e AccessRequest(requestValidity uint256)
func (_BaseContentSpace *BaseContentSpaceFilterer) WatchAccessRequest(opts *bind.WatchOpts, sink chan<- *BaseContentSpaceAccessRequest) (event.Subscription, error) {

	logs, sub, err := _BaseContentSpace.contract.WatchLogs(opts, "AccessRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentSpaceAccessRequest)
				if err := _BaseContentSpace.contract.UnpackLog(event, "AccessRequest", log); err != nil {
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

// BaseContentSpaceCommitIterator is returned from FilterCommit and is used to iterate over the raw logs and unpacked data for Commit events raised by the BaseContentSpace contract.
type BaseContentSpaceCommitIterator struct {
	Event *BaseContentSpaceCommit // Event containing the contract specifics and raw log

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
func (it *BaseContentSpaceCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentSpaceCommit)
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
		it.Event = new(BaseContentSpaceCommit)
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
func (it *BaseContentSpaceCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentSpaceCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentSpaceCommit represents a Commit event raised by the BaseContentSpace contract.
type BaseContentSpaceCommit struct {
	ObjectHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommit is a free log retrieval operation binding the contract event 0x9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff.
//
// Solidity: e Commit(objectHash bytes32)
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterCommit(opts *bind.FilterOpts) (*BaseContentSpaceCommitIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "Commit")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceCommitIterator{contract: _BaseContentSpace.contract, event: "Commit", logs: logs, sub: sub}, nil
}

// WatchCommit is a free log subscription operation binding the contract event 0x9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff.
//
// Solidity: e Commit(objectHash bytes32)
func (_BaseContentSpace *BaseContentSpaceFilterer) WatchCommit(opts *bind.WatchOpts, sink chan<- *BaseContentSpaceCommit) (event.Subscription, error) {

	logs, sub, err := _BaseContentSpace.contract.WatchLogs(opts, "Commit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentSpaceCommit)
				if err := _BaseContentSpace.contract.UnpackLog(event, "Commit", log); err != nil {
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

// BaseContentSpaceCreateAccountLibraryIterator is returned from FilterCreateAccountLibrary and is used to iterate over the raw logs and unpacked data for CreateAccountLibrary events raised by the BaseContentSpace contract.
type BaseContentSpaceCreateAccountLibraryIterator struct {
	Event *BaseContentSpaceCreateAccountLibrary // Event containing the contract specifics and raw log

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
func (it *BaseContentSpaceCreateAccountLibraryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentSpaceCreateAccountLibrary)
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
		it.Event = new(BaseContentSpaceCreateAccountLibrary)
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
func (it *BaseContentSpaceCreateAccountLibraryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentSpaceCreateAccountLibraryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentSpaceCreateAccountLibrary represents a CreateAccountLibrary event raised by the BaseContentSpace contract.
type BaseContentSpaceCreateAccountLibrary struct {
	AccountAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCreateAccountLibrary is a free log retrieval operation binding the contract event 0xeeebced55b1b3d8371f41e9358ca416596184f5c40746f1c10b5224c1c154e08.
//
// Solidity: e CreateAccountLibrary(accountAddress address)
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterCreateAccountLibrary(opts *bind.FilterOpts) (*BaseContentSpaceCreateAccountLibraryIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "CreateAccountLibrary")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceCreateAccountLibraryIterator{contract: _BaseContentSpace.contract, event: "CreateAccountLibrary", logs: logs, sub: sub}, nil
}

// WatchCreateAccountLibrary is a free log subscription operation binding the contract event 0xeeebced55b1b3d8371f41e9358ca416596184f5c40746f1c10b5224c1c154e08.
//
// Solidity: e CreateAccountLibrary(accountAddress address)
func (_BaseContentSpace *BaseContentSpaceFilterer) WatchCreateAccountLibrary(opts *bind.WatchOpts, sink chan<- *BaseContentSpaceCreateAccountLibrary) (event.Subscription, error) {

	logs, sub, err := _BaseContentSpace.contract.WatchLogs(opts, "CreateAccountLibrary")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentSpaceCreateAccountLibrary)
				if err := _BaseContentSpace.contract.UnpackLog(event, "CreateAccountLibrary", log); err != nil {
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

// BaseContentSpaceCreateContentTypeIterator is returned from FilterCreateContentType and is used to iterate over the raw logs and unpacked data for CreateContentType events raised by the BaseContentSpace contract.
type BaseContentSpaceCreateContentTypeIterator struct {
	Event *BaseContentSpaceCreateContentType // Event containing the contract specifics and raw log

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
func (it *BaseContentSpaceCreateContentTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentSpaceCreateContentType)
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
		it.Event = new(BaseContentSpaceCreateContentType)
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
func (it *BaseContentSpaceCreateContentTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentSpaceCreateContentTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentSpaceCreateContentType represents a CreateContentType event raised by the BaseContentSpace contract.
type BaseContentSpaceCreateContentType struct {
	ContentTypeAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCreateContentType is a free log retrieval operation binding the contract event 0x9e69777f30c55126be256664fa7beff4b796ac32ebceab94df5071b0148017f8.
//
// Solidity: e CreateContentType(contentTypeAddress address)
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterCreateContentType(opts *bind.FilterOpts) (*BaseContentSpaceCreateContentTypeIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "CreateContentType")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceCreateContentTypeIterator{contract: _BaseContentSpace.contract, event: "CreateContentType", logs: logs, sub: sub}, nil
}

// WatchCreateContentType is a free log subscription operation binding the contract event 0x9e69777f30c55126be256664fa7beff4b796ac32ebceab94df5071b0148017f8.
//
// Solidity: e CreateContentType(contentTypeAddress address)
func (_BaseContentSpace *BaseContentSpaceFilterer) WatchCreateContentType(opts *bind.WatchOpts, sink chan<- *BaseContentSpaceCreateContentType) (event.Subscription, error) {

	logs, sub, err := _BaseContentSpace.contract.WatchLogs(opts, "CreateContentType")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentSpaceCreateContentType)
				if err := _BaseContentSpace.contract.UnpackLog(event, "CreateContentType", log); err != nil {
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

// BaseContentSpaceCreateGroupIterator is returned from FilterCreateGroup and is used to iterate over the raw logs and unpacked data for CreateGroup events raised by the BaseContentSpace contract.
type BaseContentSpaceCreateGroupIterator struct {
	Event *BaseContentSpaceCreateGroup // Event containing the contract specifics and raw log

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
func (it *BaseContentSpaceCreateGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentSpaceCreateGroup)
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
		it.Event = new(BaseContentSpaceCreateGroup)
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
func (it *BaseContentSpaceCreateGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentSpaceCreateGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentSpaceCreateGroup represents a CreateGroup event raised by the BaseContentSpace contract.
type BaseContentSpaceCreateGroup struct {
	GroupAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateGroup is a free log retrieval operation binding the contract event 0xa3b1fe71ae61bad8cffa485b230e24e518938f76182a30fa0d9979e7237ad159.
//
// Solidity: e CreateGroup(groupAddress address)
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterCreateGroup(opts *bind.FilterOpts) (*BaseContentSpaceCreateGroupIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "CreateGroup")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceCreateGroupIterator{contract: _BaseContentSpace.contract, event: "CreateGroup", logs: logs, sub: sub}, nil
}

// WatchCreateGroup is a free log subscription operation binding the contract event 0xa3b1fe71ae61bad8cffa485b230e24e518938f76182a30fa0d9979e7237ad159.
//
// Solidity: e CreateGroup(groupAddress address)
func (_BaseContentSpace *BaseContentSpaceFilterer) WatchCreateGroup(opts *bind.WatchOpts, sink chan<- *BaseContentSpaceCreateGroup) (event.Subscription, error) {

	logs, sub, err := _BaseContentSpace.contract.WatchLogs(opts, "CreateGroup")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentSpaceCreateGroup)
				if err := _BaseContentSpace.contract.UnpackLog(event, "CreateGroup", log); err != nil {
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

// BaseContentSpaceCreateLibraryIterator is returned from FilterCreateLibrary and is used to iterate over the raw logs and unpacked data for CreateLibrary events raised by the BaseContentSpace contract.
type BaseContentSpaceCreateLibraryIterator struct {
	Event *BaseContentSpaceCreateLibrary // Event containing the contract specifics and raw log

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
func (it *BaseContentSpaceCreateLibraryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentSpaceCreateLibrary)
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
		it.Event = new(BaseContentSpaceCreateLibrary)
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
func (it *BaseContentSpaceCreateLibraryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentSpaceCreateLibraryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentSpaceCreateLibrary represents a CreateLibrary event raised by the BaseContentSpace contract.
type BaseContentSpaceCreateLibrary struct {
	LibraryAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCreateLibrary is a free log retrieval operation binding the contract event 0x473c07a6d0228c4fb8fe2be3b4617c3b5fb7c0f8cd9ba4b67e8631844b9b6571.
//
// Solidity: e CreateLibrary(libraryAddress address)
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterCreateLibrary(opts *bind.FilterOpts) (*BaseContentSpaceCreateLibraryIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "CreateLibrary")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceCreateLibraryIterator{contract: _BaseContentSpace.contract, event: "CreateLibrary", logs: logs, sub: sub}, nil
}

// WatchCreateLibrary is a free log subscription operation binding the contract event 0x473c07a6d0228c4fb8fe2be3b4617c3b5fb7c0f8cd9ba4b67e8631844b9b6571.
//
// Solidity: e CreateLibrary(libraryAddress address)
func (_BaseContentSpace *BaseContentSpaceFilterer) WatchCreateLibrary(opts *bind.WatchOpts, sink chan<- *BaseContentSpaceCreateLibrary) (event.Subscription, error) {

	logs, sub, err := _BaseContentSpace.contract.WatchLogs(opts, "CreateLibrary")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentSpaceCreateLibrary)
				if err := _BaseContentSpace.contract.UnpackLog(event, "CreateLibrary", log); err != nil {
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

// BaseContentSpaceUpdateRequestIterator is returned from FilterUpdateRequest and is used to iterate over the raw logs and unpacked data for UpdateRequest events raised by the BaseContentSpace contract.
type BaseContentSpaceUpdateRequestIterator struct {
	Event *BaseContentSpaceUpdateRequest // Event containing the contract specifics and raw log

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
func (it *BaseContentSpaceUpdateRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentSpaceUpdateRequest)
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
		it.Event = new(BaseContentSpaceUpdateRequest)
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
func (it *BaseContentSpaceUpdateRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentSpaceUpdateRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentSpaceUpdateRequest represents a UpdateRequest event raised by the BaseContentSpace contract.
type BaseContentSpaceUpdateRequest struct {
	ObjectHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateRequest is a free log retrieval operation binding the contract event 0x1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef19.
//
// Solidity: e UpdateRequest(objectHash bytes32)
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterUpdateRequest(opts *bind.FilterOpts) (*BaseContentSpaceUpdateRequestIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceUpdateRequestIterator{contract: _BaseContentSpace.contract, event: "UpdateRequest", logs: logs, sub: sub}, nil
}

// WatchUpdateRequest is a free log subscription operation binding the contract event 0x1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef19.
//
// Solidity: e UpdateRequest(objectHash bytes32)
func (_BaseContentSpace *BaseContentSpaceFilterer) WatchUpdateRequest(opts *bind.WatchOpts, sink chan<- *BaseContentSpaceUpdateRequest) (event.Subscription, error) {

	logs, sub, err := _BaseContentSpace.contract.WatchLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentSpaceUpdateRequest)
				if err := _BaseContentSpace.contract.UnpackLog(event, "UpdateRequest", log); err != nil {
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

// BaseContentTypeABI is the input ABI used to generate the binding from.
const BaseContentTypeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"objectHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"object_hash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accessRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"Commit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"UpdateRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestValidity\",\"type\":\"uint256\"}],\"name\":\"AccessRequest\",\"type\":\"event\"}]"

// BaseContentTypeBin is the compiled bytecode used for deploying new contracts.
const BaseContentTypeBin = `0x606060405260008054600160a060020a03328116600160a060020a031992831681178455600180548416909117905560038054339092169190921617905561042c90819061004d90396000f3006060604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100a557806341c0e1b5146100d45780636d2e4b1b146100e75780638da5cb5b14610106578063af570c0414610119578063c287e0ed1461012c578063e02dd9c21461013f578063f14fcbc814610164578063f15518871461017a578063f2fde38b146101a1575b005b34156100b057600080fd5b6100b86101c0565b604051600160a060020a03909116815260200160405180910390f35b34156100df57600080fd5b6100a36101cf565b34156100f257600080fd5b6100a3600160a060020a03600435166101f8565b341561011157600080fd5b6100b8610299565b341561012457600080fd5b6100b86102a8565b341561013757600080fd5b6100a36102b7565b341561014a57600080fd5b610152610309565b60405190815260200160405180910390f35b341561016f57600080fd5b6100a360043561030f565b341561018557600080fd5b61018d610365565b604051901515815260200160405180910390f35b34156101ac57600080fd5b6100a3600160a060020a03600435166103a1565b600054600160a060020a031681565b60015432600160a060020a039081169116146101ea57600080fd5b600154600160a060020a0316ff5b60005432600160a060020a0390811691161461021357600080fd5b600160a060020a038116151561022857600080fd5b600054600154600160a060020a039081169116141561026a576001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031681565b600354600160a060020a031681565b60015432600160a060020a039081169116146102d257600080fd5b7f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef1960025460405190815260200160405180910390a1565b60025481565b60015432600160a060020a0390811691161461032a57600080fd5b60028190557f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff8160405190815260200160405180910390a150565b60007fca8116fd70ecea0cf35f7c08eb31491cce5ebd6bafef5bd6426862df29840463600060405190815260200160405180910390a150600190565b60015432600160a060020a039081169116146103bc57600080fd5b600160a060020a03811615156103d157600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582002461abc581e32b9ded636540e838ffae65caf9c71a4215780265a749f0276410029`

// DeployBaseContentType deploys a new Ethereum contract, binding an instance of BaseContentType to it.
func DeployBaseContentType(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BaseContentType, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseContentTypeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaseContentTypeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseContentType{BaseContentTypeCaller: BaseContentTypeCaller{contract: contract}, BaseContentTypeTransactor: BaseContentTypeTransactor{contract: contract}, BaseContentTypeFilterer: BaseContentTypeFilterer{contract: contract}}, nil
}

// BaseContentType is an auto generated Go binding around an Ethereum contract.
type BaseContentType struct {
	BaseContentTypeCaller     // Read-only binding to the contract
	BaseContentTypeTransactor // Write-only binding to the contract
	BaseContentTypeFilterer   // Log filterer for contract events
}

// BaseContentTypeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseContentTypeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseContentTypeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseContentTypeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseContentTypeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseContentTypeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseContentTypeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseContentTypeSession struct {
	Contract     *BaseContentType  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseContentTypeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseContentTypeCallerSession struct {
	Contract *BaseContentTypeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BaseContentTypeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseContentTypeTransactorSession struct {
	Contract     *BaseContentTypeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BaseContentTypeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseContentTypeRaw struct {
	Contract *BaseContentType // Generic contract binding to access the raw methods on
}

// BaseContentTypeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseContentTypeCallerRaw struct {
	Contract *BaseContentTypeCaller // Generic read-only contract binding to access the raw methods on
}

// BaseContentTypeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseContentTypeTransactorRaw struct {
	Contract *BaseContentTypeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseContentType creates a new instance of BaseContentType, bound to a specific deployed contract.
func NewBaseContentType(address common.Address, backend bind.ContractBackend) (*BaseContentType, error) {
	contract, err := bindBaseContentType(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseContentType{BaseContentTypeCaller: BaseContentTypeCaller{contract: contract}, BaseContentTypeTransactor: BaseContentTypeTransactor{contract: contract}, BaseContentTypeFilterer: BaseContentTypeFilterer{contract: contract}}, nil
}

// NewBaseContentTypeCaller creates a new read-only instance of BaseContentType, bound to a specific deployed contract.
func NewBaseContentTypeCaller(address common.Address, caller bind.ContractCaller) (*BaseContentTypeCaller, error) {
	contract, err := bindBaseContentType(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseContentTypeCaller{contract: contract}, nil
}

// NewBaseContentTypeTransactor creates a new write-only instance of BaseContentType, bound to a specific deployed contract.
func NewBaseContentTypeTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseContentTypeTransactor, error) {
	contract, err := bindBaseContentType(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseContentTypeTransactor{contract: contract}, nil
}

// NewBaseContentTypeFilterer creates a new log filterer instance of BaseContentType, bound to a specific deployed contract.
func NewBaseContentTypeFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseContentTypeFilterer, error) {
	contract, err := bindBaseContentType(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseContentTypeFilterer{contract: contract}, nil
}

// bindBaseContentType binds a generic wrapper to an already deployed contract.
func bindBaseContentType(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseContentTypeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseContentType *BaseContentTypeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseContentType.Contract.BaseContentTypeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseContentType *BaseContentTypeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContentType.Contract.BaseContentTypeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseContentType *BaseContentTypeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseContentType.Contract.BaseContentTypeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseContentType *BaseContentTypeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseContentType.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseContentType *BaseContentTypeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContentType.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseContentType *BaseContentTypeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseContentType.Contract.contract.Transact(opts, method, params...)
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_BaseContentType *BaseContentTypeCaller) ContentSpace(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContentType.contract.Call(opts, out, "contentSpace")
	return *ret0, err
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_BaseContentType *BaseContentTypeSession) ContentSpace() (common.Address, error) {
	return _BaseContentType.Contract.ContentSpace(&_BaseContentType.CallOpts)
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_BaseContentType *BaseContentTypeCallerSession) ContentSpace() (common.Address, error) {
	return _BaseContentType.Contract.ContentSpace(&_BaseContentType.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseContentType *BaseContentTypeCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContentType.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseContentType *BaseContentTypeSession) Creator() (common.Address, error) {
	return _BaseContentType.Contract.Creator(&_BaseContentType.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseContentType *BaseContentTypeCallerSession) Creator() (common.Address, error) {
	return _BaseContentType.Contract.Creator(&_BaseContentType.CallOpts)
}

// ObjectHash is a free data retrieval call binding the contract method 0xe02dd9c2.
//
// Solidity: function objectHash() constant returns(bytes32)
func (_BaseContentType *BaseContentTypeCaller) ObjectHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseContentType.contract.Call(opts, out, "objectHash")
	return *ret0, err
}

// ObjectHash is a free data retrieval call binding the contract method 0xe02dd9c2.
//
// Solidity: function objectHash() constant returns(bytes32)
func (_BaseContentType *BaseContentTypeSession) ObjectHash() ([32]byte, error) {
	return _BaseContentType.Contract.ObjectHash(&_BaseContentType.CallOpts)
}

// ObjectHash is a free data retrieval call binding the contract method 0xe02dd9c2.
//
// Solidity: function objectHash() constant returns(bytes32)
func (_BaseContentType *BaseContentTypeCallerSession) ObjectHash() ([32]byte, error) {
	return _BaseContentType.Contract.ObjectHash(&_BaseContentType.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseContentType *BaseContentTypeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContentType.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseContentType *BaseContentTypeSession) Owner() (common.Address, error) {
	return _BaseContentType.Contract.Owner(&_BaseContentType.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseContentType *BaseContentTypeCallerSession) Owner() (common.Address, error) {
	return _BaseContentType.Contract.Owner(&_BaseContentType.CallOpts)
}

// AccessRequest is a paid mutator transaction binding the contract method 0xf1551887.
//
// Solidity: function accessRequest() returns(bool)
func (_BaseContentType *BaseContentTypeTransactor) AccessRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContentType.contract.Transact(opts, "accessRequest")
}

// AccessRequest is a paid mutator transaction binding the contract method 0xf1551887.
//
// Solidity: function accessRequest() returns(bool)
func (_BaseContentType *BaseContentTypeSession) AccessRequest() (*types.Transaction, error) {
	return _BaseContentType.Contract.AccessRequest(&_BaseContentType.TransactOpts)
}

// AccessRequest is a paid mutator transaction binding the contract method 0xf1551887.
//
// Solidity: function accessRequest() returns(bool)
func (_BaseContentType *BaseContentTypeTransactorSession) AccessRequest() (*types.Transaction, error) {
	return _BaseContentType.Contract.AccessRequest(&_BaseContentType.TransactOpts)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(object_hash bytes32) returns()
func (_BaseContentType *BaseContentTypeTransactor) Commit(opts *bind.TransactOpts, object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContentType.contract.Transact(opts, "commit", object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(object_hash bytes32) returns()
func (_BaseContentType *BaseContentTypeSession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContentType.Contract.Commit(&_BaseContentType.TransactOpts, object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(object_hash bytes32) returns()
func (_BaseContentType *BaseContentTypeTransactorSession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContentType.Contract.Commit(&_BaseContentType.TransactOpts, object_hash)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseContentType *BaseContentTypeTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContentType.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseContentType *BaseContentTypeSession) Kill() (*types.Transaction, error) {
	return _BaseContentType.Contract.Kill(&_BaseContentType.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseContentType *BaseContentTypeTransactorSession) Kill() (*types.Transaction, error) {
	return _BaseContentType.Contract.Kill(&_BaseContentType.TransactOpts)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_BaseContentType *BaseContentTypeTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _BaseContentType.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_BaseContentType *BaseContentTypeSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseContentType.Contract.TransferCreatorship(&_BaseContentType.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_BaseContentType *BaseContentTypeTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseContentType.Contract.TransferCreatorship(&_BaseContentType.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseContentType *BaseContentTypeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseContentType.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseContentType *BaseContentTypeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseContentType.Contract.TransferOwnership(&_BaseContentType.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseContentType *BaseContentTypeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseContentType.Contract.TransferOwnership(&_BaseContentType.TransactOpts, newOwner)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xc287e0ed.
//
// Solidity: function updateRequest() returns()
func (_BaseContentType *BaseContentTypeTransactor) UpdateRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContentType.contract.Transact(opts, "updateRequest")
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xc287e0ed.
//
// Solidity: function updateRequest() returns()
func (_BaseContentType *BaseContentTypeSession) UpdateRequest() (*types.Transaction, error) {
	return _BaseContentType.Contract.UpdateRequest(&_BaseContentType.TransactOpts)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xc287e0ed.
//
// Solidity: function updateRequest() returns()
func (_BaseContentType *BaseContentTypeTransactorSession) UpdateRequest() (*types.Transaction, error) {
	return _BaseContentType.Contract.UpdateRequest(&_BaseContentType.TransactOpts)
}

// BaseContentTypeAccessRequestIterator is returned from FilterAccessRequest and is used to iterate over the raw logs and unpacked data for AccessRequest events raised by the BaseContentType contract.
type BaseContentTypeAccessRequestIterator struct {
	Event *BaseContentTypeAccessRequest // Event containing the contract specifics and raw log

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
func (it *BaseContentTypeAccessRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentTypeAccessRequest)
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
		it.Event = new(BaseContentTypeAccessRequest)
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
func (it *BaseContentTypeAccessRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentTypeAccessRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentTypeAccessRequest represents a AccessRequest event raised by the BaseContentType contract.
type BaseContentTypeAccessRequest struct {
	RequestValidity *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccessRequest is a free log retrieval operation binding the contract event 0xca8116fd70ecea0cf35f7c08eb31491cce5ebd6bafef5bd6426862df29840463.
//
// Solidity: e AccessRequest(requestValidity uint256)
func (_BaseContentType *BaseContentTypeFilterer) FilterAccessRequest(opts *bind.FilterOpts) (*BaseContentTypeAccessRequestIterator, error) {

	logs, sub, err := _BaseContentType.contract.FilterLogs(opts, "AccessRequest")
	if err != nil {
		return nil, err
	}
	return &BaseContentTypeAccessRequestIterator{contract: _BaseContentType.contract, event: "AccessRequest", logs: logs, sub: sub}, nil
}

// WatchAccessRequest is a free log subscription operation binding the contract event 0xca8116fd70ecea0cf35f7c08eb31491cce5ebd6bafef5bd6426862df29840463.
//
// Solidity: e AccessRequest(requestValidity uint256)
func (_BaseContentType *BaseContentTypeFilterer) WatchAccessRequest(opts *bind.WatchOpts, sink chan<- *BaseContentTypeAccessRequest) (event.Subscription, error) {

	logs, sub, err := _BaseContentType.contract.WatchLogs(opts, "AccessRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentTypeAccessRequest)
				if err := _BaseContentType.contract.UnpackLog(event, "AccessRequest", log); err != nil {
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

// BaseContentTypeCommitIterator is returned from FilterCommit and is used to iterate over the raw logs and unpacked data for Commit events raised by the BaseContentType contract.
type BaseContentTypeCommitIterator struct {
	Event *BaseContentTypeCommit // Event containing the contract specifics and raw log

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
func (it *BaseContentTypeCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentTypeCommit)
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
		it.Event = new(BaseContentTypeCommit)
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
func (it *BaseContentTypeCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentTypeCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentTypeCommit represents a Commit event raised by the BaseContentType contract.
type BaseContentTypeCommit struct {
	ObjectHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommit is a free log retrieval operation binding the contract event 0x9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff.
//
// Solidity: e Commit(objectHash bytes32)
func (_BaseContentType *BaseContentTypeFilterer) FilterCommit(opts *bind.FilterOpts) (*BaseContentTypeCommitIterator, error) {

	logs, sub, err := _BaseContentType.contract.FilterLogs(opts, "Commit")
	if err != nil {
		return nil, err
	}
	return &BaseContentTypeCommitIterator{contract: _BaseContentType.contract, event: "Commit", logs: logs, sub: sub}, nil
}

// WatchCommit is a free log subscription operation binding the contract event 0x9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff.
//
// Solidity: e Commit(objectHash bytes32)
func (_BaseContentType *BaseContentTypeFilterer) WatchCommit(opts *bind.WatchOpts, sink chan<- *BaseContentTypeCommit) (event.Subscription, error) {

	logs, sub, err := _BaseContentType.contract.WatchLogs(opts, "Commit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentTypeCommit)
				if err := _BaseContentType.contract.UnpackLog(event, "Commit", log); err != nil {
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

// BaseContentTypeUpdateRequestIterator is returned from FilterUpdateRequest and is used to iterate over the raw logs and unpacked data for UpdateRequest events raised by the BaseContentType contract.
type BaseContentTypeUpdateRequestIterator struct {
	Event *BaseContentTypeUpdateRequest // Event containing the contract specifics and raw log

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
func (it *BaseContentTypeUpdateRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentTypeUpdateRequest)
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
		it.Event = new(BaseContentTypeUpdateRequest)
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
func (it *BaseContentTypeUpdateRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentTypeUpdateRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentTypeUpdateRequest represents a UpdateRequest event raised by the BaseContentType contract.
type BaseContentTypeUpdateRequest struct {
	ObjectHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateRequest is a free log retrieval operation binding the contract event 0x1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef19.
//
// Solidity: e UpdateRequest(objectHash bytes32)
func (_BaseContentType *BaseContentTypeFilterer) FilterUpdateRequest(opts *bind.FilterOpts) (*BaseContentTypeUpdateRequestIterator, error) {

	logs, sub, err := _BaseContentType.contract.FilterLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return &BaseContentTypeUpdateRequestIterator{contract: _BaseContentType.contract, event: "UpdateRequest", logs: logs, sub: sub}, nil
}

// WatchUpdateRequest is a free log subscription operation binding the contract event 0x1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef19.
//
// Solidity: e UpdateRequest(objectHash bytes32)
func (_BaseContentType *BaseContentTypeFilterer) WatchUpdateRequest(opts *bind.WatchOpts, sink chan<- *BaseContentTypeUpdateRequest) (event.Subscription, error) {

	logs, sub, err := _BaseContentType.contract.WatchLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentTypeUpdateRequest)
				if err := _BaseContentType.contract.UnpackLog(event, "UpdateRequest", log); err != nil {
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

// BaseLibraryABI is the input ABI used to generate the binding from.
const BaseLibraryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"canContribute\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"approvalRequestsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contentTypeContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"addAccessorGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content_type\",\"type\":\"address\"}],\"name\":\"createContent\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reviewerGroupsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contributorGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"space\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"canReview\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accessorGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressKMS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contentTypeLicensingFees\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contributorGroupsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"submitApprovalRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingApprovalRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"addContributorGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content_contract\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payCredit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approvalRequests\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reviewerGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"hasAccess\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contentTypes\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content_type\",\"type\":\"address\"},{\"name\":\"content_contract\",\"type\":\"address\"},{\"name\":\"licensing_fee\",\"type\":\"uint256\"}],\"name\":\"addContentType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content_contract\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"},{\"name\":\"note\",\"type\":\"string\"}],\"name\":\"approveContentExecuted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentTypesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"address_KMS\",\"type\":\"address\"}],\"name\":\"setAddressKMS\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"addReviewerGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"objectHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accessorGroupsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"object_hash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accessRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"address_KMS\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"content_type\",\"type\":\"address\"}],\"name\":\"ContentObjectCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ContributorGroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ReviewerGroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"AccessorGroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentType\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"contentContract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"licensingFee\",\"type\":\"uint256\"}],\"name\":\"ContentTypeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"operationCode\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"UnauthorizedOperation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"ApproveContentRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"note\",\"type\":\"string\"}],\"name\":\"ApproveContentExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PayCredit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"Commit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"UpdateRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestValidity\",\"type\":\"uint256\"}],\"name\":\"AccessRequest\",\"type\":\"event\"}]"

// BaseLibraryBin is the compiled bytecode used for deploying new contracts.
const BaseLibraryBin = `0x606060405260006010556040516020806135b78339810160405280805160008054600160a060020a03328116600160a060020a03199283168117845560018054841690911790556003805433831690841617905560088390556009839055600a92909255600f80549290931691161790555050613536806100816000396000f3006060604052600436106101a85763ffffffff60e060020a60003504166302d05d3f81146101aa5780630eaec2c5146101d9578063163083941461020c5780631cdbee5a146102315780631d0f4351146102505780631e35d8fa1461026f57806321770a841461028e5780632393553b146102a157806327af9e38146102b757806329d00219146102ca5780632c11f392146102e957806332eaf21b146102ff5780633fbb5d711461031257806341c0e1b514610331578063470750bb1461034457806349102e611461035757806363dab9d41461036a578063679a9a3c146103805780636934d45f1461039f5780636d2e4b1b146103c15780638cb13c2e146103e05780638da5cb5b146103f6578063952e464b1461040957806395a078e81461041f578063991a3a7c1461043e578063a90229d514610454578063bf81a0b61461047c578063c287e0ed146104e3578063c65bcbe2146104f6578063c9e8e72d14610509578063dc3c29c014610528578063e02dd9c214610547578063e5538fd21461055a578063f14fcbc81461056d578063f155188714610583578063f2fde38b14610596575b005b34156101b557600080fd5b6101bd6105b5565b604051600160a060020a03909116815260200160405180910390f35b34156101e457600080fd5b6101f8600160a060020a03600435166105c4565b604051901515815260200160405180910390f35b341561021757600080fd5b61021f6106af565b60405190815260200160405180910390f35b341561023c57600080fd5b6101bd600160a060020a03600435166106b5565b341561025b57600080fd5b6101a8600160a060020a03600435166106d0565b341561027a57600080fd5b6101bd600160a060020a036004351661076b565b341561029957600080fd5b61021f61089e565b34156102ac57600080fd5b6101bd6004356108a4565b34156102c257600080fd5b6101bd6108cc565b34156102d557600080fd5b6101f8600160a060020a03600435166108db565b34156102f457600080fd5b6101bd6004356109c3565b341561030a57600080fd5b6101bd6109d1565b341561031d57600080fd5b61021f600160a060020a03600435166109e0565b341561033c57600080fd5b6101a86109f2565b341561034f57600080fd5b61021f610a1b565b341561036257600080fd5b6101f8610a21565b341561037557600080fd5b6101bd600435610cf8565b341561038b57600080fd5b6101a8600160a060020a0360043516610d45565b34156103aa57600080fd5b61021f600160a060020a0360043516602435610de0565b34156103cc57600080fd5b6101a8600160a060020a0360043516610fea565b34156103eb57600080fd5b6101bd600435611071565b341561040157600080fd5b6101bd61107f565b341561041457600080fd5b6101bd60043561108e565b341561042a57600080fd5b6101f8600160a060020a036004351661109c565b341561044957600080fd5b6101bd600435611179565b341561045f57600080fd5b6101a8600160a060020a0360043581169060243516604435611187565b341561048757600080fd5b6101f860048035600160a060020a031690602480351515919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061126795505050505050565b34156104ee57600080fd5b6101a86115c9565b341561050157600080fd5b61021f61161b565b341561051457600080fd5b6101a8600160a060020a0360043516611621565b341561053357600080fd5b6101a8600160a060020a036004351661165e565b341561055257600080fd5b61021f6116f9565b341561056557600080fd5b61021f6116ff565b341561057857600080fd5b6101a8600435611705565b341561058e57600080fd5b6101f861175b565b34156105a157600080fd5b6101a8600160a060020a03600435166117e2565b600054600160a060020a031681565b6000806000806000600480549050600014156105e357600194506106a6565b5060005b6004548110156106a15760048054829081106105ff57fe5b600091825260209091200154600160a060020a0316935083156106995783915081600160a060020a03166395a078e88760405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561066d57600080fd5b5af1151561067a57600080fd5b50505060405180519350506001831515141561069957600194506106a6565b6001016105e7565b600094505b50505050919050565b60105481565b600c60205260009081526040902054600160a060020a031681565b60015432600160a060020a039081169116146106eb57600080fd5b60068054600181016106fd8382611834565b5060009182526020909120018054600160a060020a031916600160a060020a038316179055600a805460010190557f3a94857e4393737f73edb175a7d0c195c7f635d9ae995e12740616ec55c9d41181604051600160a060020a03909116815260200160405180910390a150565b6000806000610779326105c4565b15156107cc577f23de2adc3e22f171f66b3e5a333e17feb9dc30ba9570933bd259cb6c13ef7ab7606532604051918252600160a060020a031660208201526040908101905180910390a160009250610897565b6107d461185d565b604051809103906000f08015156107ea57600080fd5b915081905080600160a060020a031663693489678560405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b151561083d57600080fd5b5af1151561084a57600080fd5b5050507f3981e74ab81857b375ec391a4f7c31ee89462cd927de6d8fbdb98f77da009c568285604051600160a060020a039283168152911660208201526040908101905180910390a18192505b5050919050565b60095481565b60048054829081106108b257fe5b600091825260209091200154600160a060020a0316905081565b600354600160a060020a031681565b6001546000908190819081908190600160a060020a038781169116141561090557600194506106a6565b5060005b6005548110156106a157600580548290811061092157fe5b600091825260209091200154600160a060020a0316935083156109bb5783915081600160a060020a03166395a078e88760405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561098f57600080fd5b5af1151561099c57600080fd5b5050506040518051935050600183151514156109bb57600194506106a6565b600101610909565b60068054829081106108b257fe5b600f54600160a060020a031681565b600d6020526000908152604090205481565b60015432600160a060020a03908116911614610a0d57600080fd5b600154600160a060020a0316ff5b60085481565b600033808280600160a060020a03328116908416638da5cb5b6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610a6857600080fd5b5af11515610a7557600080fd5b50505060405180519050600160a060020a0316141515610a985760009450610cf1565b6009541515610bcd5782600160a060020a031663545b35396040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610ade57600080fd5b5af11515610aeb57600080fd5b5050506040518051925050600160a060020a03831663708f53ab60008460405160e060020a63ffffffff8516028152600481019290925260ff166024820152604401602060405180830381600087803b1515610b4657600080fd5b5af11515610b5357600080fd5b505050604051805190509050610b698382610de0565b507f8a6b83b9f4a16d1fd3ec333d40e20591cfb3ad7c24ecbfbc9a5c5b1a6445cab4846001604051600160a060020a039092168252151560208201526060604080830182905260009183019190915260a0909101905180910390a160019450610cf1565b600160a060020a03841660009081526011602052604090205415610bf45760009450610cf1565b600e546010541015610c405783600e601054815481101515610c1257fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055610c78565b600e805460018101610c528382611834565b5060009182526020909120018054600160a060020a031916600160a060020a0386161790555b60108054600160a060020a038616600090815260116020526040908190206001928301905582549091019091557f0588a34cf0de4e025d359c89ca4bacbcbf175440909952d91c814412d9da996a908590329051600160a060020a039283168152911660208201526040908101905180910390a1600194505b5050505090565b600060105460001480610d0d57508160105411155b15610d1a57506000610d40565b600e805483908110610d2857fe5b600091825260209091200154600160a060020a031690505b919050565b60015432600160a060020a03908116911614610d6057600080fd5b6004805460018101610d728382611834565b5060009182526020909120018054600160a060020a031916600160a060020a0383161790556008805460010190557f218673669018c25b89bfbf1b58d0075e37c8847ef16e707b92355b7833e97d6181604051600160a060020a03909116815260200160405180910390a150565b6000806000806000851115610fdc5785925082600160a060020a03166307b27aca6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610e2f57600080fd5b5af11515610e3c57600080fd5b505050604051805191505080851115610e5757809150610e5b565b8491505b82600160a060020a0316638da5cb5b6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610e9857600080fd5b5af11515610ea557600080fd5b5050506040518051600160a060020a0316905082156108fc0283604051600060405180830381858888f193505050501515610edf57600080fd5b82600160a060020a031663f79444718360405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610f2557600080fd5b5af11515610f3257600080fd5b50505060405180519050507fd8ad32fc5c165753d9da5f44f42ef7daf8002336cab4bdd13b7e58f48b451e3d83600160a060020a0316638da5cb5b6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610f9b57600080fd5b5af11515610fa857600080fd5b5050506040518051905083604051600160a060020a03909216825260208201526040908101905180910390a1819350610fe1565b600093505b50505092915050565b60005432600160a060020a0390811691161461100557600080fd5b600160a060020a038116151561101a57600080fd5b600054600154600160a060020a039081169116141561104f5760018054600160a060020a031916600160a060020a0383161790555b60008054600160a060020a031916600160a060020a0392909216919091179055565b600e8054829081106108b257fe5b600154600160a060020a031681565b60058054829081106108b257fe5b6000806000806000600680549050600014156110bb57600194506106a6565b5060005b6006548110156106a15760068054829081106110d757fe5b600091825260209091200154600160a060020a0316935083156111715783915081600160a060020a03166395a078e88760405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561114557600080fd5b5af1151561115257600080fd5b50505060405180519350506001831515141561117157600194506106a6565b6001016110bf565b60078054829081106108b257fe5b60015432600160a060020a039081169116146111a257600080fd5b60078054600181016111b48382611834565b5060009182526020808320919091018054600160a060020a0319908116600160a060020a03888116918217909355600b805460010190558452600c835260408085208054909216928716929092179055600d909152908190208290557f62c1f342f2a52ee6ac3eeef54df8c836fb31869e9be7c5e5394c042d131c9ecc9084908490849051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60008060008060008060008061127c326108db565b151561128b57600097506115bb565b600160a060020a038b16600090815260116020526040902054600e805460001990920198508c975090889081106112be57fe5b600091825260208083209091018054600160a060020a031916905560108054600019018155600160a060020a038e1683526011909152604082209190915554879011156113b557600e60105481548110151561131657fe5b600091825260209091200154600e8054600160a060020a03909216965086918990811061133f57fe5b906000526020600020900160006101000a815481600160a060020a030219169083600160a060020a03160217905550600e60105481548110151561137f57fe5b600091825260208083209091018054600160a060020a0319169055600160a060020a038716825260119052604090206001880190555b85600160a060020a03166327c1c21d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156113f257600080fd5b5af115156113ff57600080fd5b505050604051805194505060008413156115b65785600160a060020a031663545b35396040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561145057600080fd5b5af1151561145d57600080fd5b505050604051805193505060018a1515141561147c5760009150611487565b836001016000190291505b85600160a060020a031663708f53ab838560405160e060020a63ffffffff8516028152600481019290925260ff166024820152604401602060405180830381600087803b15156114d657600080fd5b5af115156114e357600080fd5b5050506040518051905090506114f98682610de0565b507f8a6b83b9f4a16d1fd3ec333d40e20591cfb3ad7c24ecbfbc9a5c5b1a6445cab48b8b8b604051600160a060020a0384168152821515602082015260606040820181815290820183818151815260200191508051906020019080838360005b83811015611571578082015183820152602001611559565b50505050905090810190601f16801561159e5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a1600197506115bb565b600097505b505050505050509392505050565b60015432600160a060020a039081169116146115e457600080fd5b7f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef1960025460405190815260200160405180910390a1565b600b5481565b60015432600160a060020a0390811691161461163c57600080fd5b600f8054600160a060020a031916600160a060020a0392909216919091179055565b60015432600160a060020a0390811691161461167957600080fd5b600580546001810161168b8382611834565b5060009182526020909120018054600160a060020a031916600160a060020a0383161790556009805460010190557f1b88a571cc8ac2e87512f05648e79d184f5cc0cbb2889bc487c41f8b9a3202eb81604051600160a060020a03909116815260200160405180910390a150565b60025481565b600a5481565b60015432600160a060020a0390811691161461172057600080fd5b60028190557f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff8160405190815260200160405180910390a150565b60006117663261109c565b156117a7577fca8116fd70ecea0cf35f7c08eb31491cce5ebd6bafef5bd6426862df29840463600060405190815260200160405180910390a15060016117df565b7fca8116fd70ecea0cf35f7c08eb31491cce5ebd6bafef5bd6426862df29840463606960405190815260200160405180910390a15060005b90565b60015432600160a060020a039081169116146117fd57600080fd5b600160a060020a038116151561181257600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b8154818355818115116118585760008381526020902061185891810190830161186d565b505050565b604051611c7f8061188c83390190565b6117df91905b808211156118875760008155600101611873565b509056006060604090815260008054600160a060020a03328116600160a060020a03199283168117845560018054841690911790556006805433831693169290921791829055600019600855600b839055600c929092557fc3decc188980e855666b70498ca85e8fa284d97d30483d828fa126f7303d7d199291169051600160a060020a03909116815260200160405180910390a1611be08061009f6000396000f3006060604052600436106101875763ffffffff60e060020a60003504166302d05d3f81146101895780630779564b146101b857806307b27aca146101dd5780631a735f18146101f05780632310167f1461024257806327c1c21d1461025557806332eaf21b1461026857806336ebffca1461027b57806341c0e1b51461028e5780635267db44146102a1578063545b3539146102b75780635cc4aa9b146102e057806364ade32b1461030557806369348967146103185780636d2e4b1b14610337578063708f53ab14610356578063879fe48f146103725780638da5cb5b146104095780638f7792011461041c578063a1ff106e1461042f578063b65240891461053f578063b816f51314610555578063c287e0ed14610568578063c68c1d1f1461057b578063c9e8e72d146105c9578063e02dd9c2146105e8578063e5385303146105fb578063ee56d7671461061a578063f14fcbc8146106ba578063f2fde38b146106d0578063f4d9bae8146106ef578063f794447114610705578063f94f00141461071b575b005b341561019457600080fd5b61019c61072e565b604051600160a060020a03909116815260200160405180910390f35b34156101c357600080fd5b6101cb61073d565b60405190815260200160405180910390f35b34156101e857600080fd5b6101cb610743565b34156101fb57600080fd5b61020660043561074d565b6040518084600160a060020a0316600160a060020a031681526020018381526020018260000b60000b8152602001935050505060405180910390f35b341561024d57600080fd5b61019c610778565b341561026057600080fd5b6101cb610787565b341561027357600080fd5b61019c61078d565b341561028657600080fd5b61019c61079c565b341561029957600080fd5b6101876107ab565b34156102ac57600080fd5b6101cb6004356107d4565b34156102c257600080fd5b6102ca610873565b60405160ff909116815260200160405180910390f35b6102f160043560243560443561087c565b604051901515815260200160405180910390f35b341561031057600080fd5b6101cb610a12565b341561032357600080fd5b610187600160a060020a0360043516610a18565b341561034257600080fd5b610187600160a060020a0360043516610c00565b341561036157600080fd5b6101cb60043560ff60243516610c87565b341561037d57600080fd5b6101cb6004803560ff1690604460248035908101908301358060208082020160405190810160405280939291908181526020018383602002808284378201915050505050509190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843750949650610dc095505050505050565b341561041457600080fd5b61019c610f28565b341561042757600080fd5b6101cb610f37565b6102f16004803560ff169060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843750949650610f3d95505050505050565b341561054a57600080fd5b610187600435611503565b341561056057600080fd5b61019c611531565b341561057357600080fd5b610187611540565b6102f16004803560ff169060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061159295505050505050565b34156105d457600080fd5b610187600160a060020a03600435166116d5565b34156105f357600080fd5b6101cb611712565b341561060657600080fd5b610187600160a060020a0360043516611718565b341561062557600080fd5b6102f16004803590602480351515919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061179595505050505050565b34156106c557600080fd5b610187600435611a37565b34156106db57600080fd5b610187600160a060020a0360043516611a8d565b34156106fa57600080fd5b6101cb600435611adf565b341561071057600080fd5b6101cb600435611b3d565b341561072657600080fd5b6101cb611b6e565b600054600160a060020a031681565b600a5481565b600b54600a540390565b600d60205260009081526040812080546001820154600290920154600160a060020a03909116920b83565b600554600160a060020a031681565b60085481565b600454600160a060020a031681565b600354600160a060020a031681565b60015432600160a060020a039081169116146107c657600080fd5b600154600160a060020a0316ff5b60015460009032600160a060020a03908116911614801561080c5750600082128061080c575060008213801561080c57506000600854125b156108175760088290555b60065433600160a060020a03908116911614156108345760088290555b7fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a260085460405190815260200160405180910390a1506008545b919050565b60095460ff1681565b60055460009060019082908190600160a060020a0316156108fb57600554600160a060020a0316915081636b2d13248860405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156108e257600080fd5b5af115156108ef57600080fd5b50505060405180519350505b506000868152600d602052604090208054600160a060020a0316158015906109305750805433600160a060020a039081169116145b156109b2576002810154600090810b900b151561097e5733600160a060020a03166108fc82600101549081150290604051600060405180830381858888f19350505050151561097e57600080fd5b6000878152600d602052604081208054600160a060020a03191681556001810191909155600201805460ff191690556109b7565b600092505b7f7f1f4b28434ce7beab4983e64a8b5bb96e195a67029fdaff925028aec57fbc6b87878786604051938452602084019290925260408084019190915290151560608301526080909101905180910390a1509095945050505050565b60075481565b60015460009032600160a060020a03908116911614610a3657600080fd5b5060038054600160a060020a031916600160a060020a03838116919091179091556006541680631cdbee5a8360405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610aa057600080fd5b5af11515610aad57600080fd5b505050604051805160058054600160a060020a031916600160a060020a0392831617905582169050633fbb5d718360405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610b1957600080fd5b5af11515610b2657600080fd5b5050506040518051600a5550600160a060020a0381166332eaf21b6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610b6f57600080fd5b5af11515610b7c57600080fd5b505050604051805160048054600160a060020a031916600160a060020a03928316179055600554600a547f1c158b764edca9bf03a401a388355220a7395febb3c011729d8060572f065627935085929190911690604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15050565b60005432600160a060020a03908116911614610c1b57600080fd5b600160a060020a0381161515610c3057600080fd5b600054600154600160a060020a0390811691161415610c655760018054600160a060020a031916600160a060020a0383161790555b60008054600160a060020a031916600160a060020a0392909216919091179055565b600154600090819081908190819032600160a060020a0390811691161480610cbd575060065433600160a060020a039081169116145b15610db1576009805460ff191660ff881617905560055460009450600160a060020a03161515610d115760095487935060ff166064148015610cfd575086155b15610d0c57600b54600a540393505b610d9f565b600554600160a060020a0316915081633513a8058860405160e060020a63ffffffff841602815260048101919091526024016040805180830381600087803b1515610d5b57600080fd5b5af11515610d6857600080fd5b5050506040518051906020018051919450909150506000811215610d9b57861515610d0c57600b54600a54039350610d9f565b8093505b610da8836107d4565b50839450610db6565b600094505b5050505092915050565b6007546005546000919082908190600160a060020a031615610edf57600554600160a060020a031691508163f8117ca28888886040518463ffffffff1660e060020a028152600401808460ff1660ff1681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015610e55578082015183820152602001610e3d565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610e94578082015183820152602001610e7c565b5050505090500195505050505050602060405180830381600087803b1515610ebb57600080fd5b5af11515610ec857600080fd5b505050604051805191505060008112610edf578092505b7fa58326ee5bb617cb8b4f0d0f5f557c469d2d05d7a738f777037deda9c724b370878460405160ff909216825260208201526040908101905180910390a1509095945050505050565b600154600160a060020a031681565b600c5481565b6000806000610f4a611b74565b600c80546001019055600654600160a060020a0316925060008080856395a078e83260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610faa57600080fd5b5af11515610fb757600080fd5b50505060405180511515905061103457600080516020611b958339815191526069600c548e6000604051938452602084019290925260ff16604080840191909152606083019190915260c060808301819052600090830181905261010060a08401819052830152610140909101905180910390a1600096506114f4565b60015432600160a060020a039081169116146110cc576110558c8a8a610dc0565b9450843410156110cc57600080516020611b958339815191526067600c548e6000604051938452602084019290925260ff16604080840191909152606083019190915260c060808301819052600090830181905261010060a08401819052830152610140909101905180910390a1600096506114f4565b60606040519081016040908152600160a060020a0333168252346020808401919091526000828401819052600c548152600d90915220909450849081518154600160a060020a031916600160a060020a03919091161781556020820151816001015560408201516002909101805460009290920b60ff1660ff1990921691909117905550600554600160a060020a0316156112d857600554600c54600160a060020a039091169350839063123e0e80908e8c8c6040518563ffffffff1660e060020a028152600401808581526020018460ff1660ff1681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156111e75780820151838201526020016111cf565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561122657808201518382015260200161120e565b505050509050019650505050505050602060405180830381600087803b151561124e57600080fd5b5af1151561125b57600080fd5b505050604051805192505081156112d857600080516020611b9583398151915282600c548e6000604051938452602084019290925260ff16604080840191909152606083019190915260c060808301819052600090830181905261010060a08401819052830152610140909101905180910390a1600096506114f4565b600080516020611b958339815191526000600c548e6002548f8f6040518681526020810186905260ff851660408201526060810184905260c0608082018181529060a0830190830185818151815260200191508051906020019080838360005b83811015611350578082015183820152602001611338565b50505050905090810190601f16801561137d5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156113b357808201518382015260200161139b565b50505050905090810190601f1680156113e05780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390a15060005b885181101561146c5788818151811061140c57fe5b9060200190602002015115611464577f515e0a48b385fce2a8e4d9f169a97c4f6ea669a752358f5e6ab37cc3c2e84c3889828151811061144857fe5b9060200190602002015160405190815260200160405180910390a15b6001016113f7565b5060005b87518110156114ef5788818151811061148557fe5b90602001906020020151156114e7577fb6e3239e521a6c66920ae634f8e921a37e6991d520ac44d52f8516397f41b6848882815181106114c157fe5b90602001906020020151604051600160a060020a03909116815260200160405180910390a15b600101611470565b600196505b50505050505095945050505050565b60015432600160a060020a03908116911614801561152357506000600854125b1561152e57600a8190555b50565b600654600160a060020a031681565b60015432600160a060020a0390811691161461155b57600080fd5b7f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef1960025460405190815260200160405180910390a1565b6001546000908190819032600160a060020a039081169116146115b457600080fd5b836040518082805190602001908083835b602083106115e45780518252601f1990920191602091820191016115c5565b6001836020036101000a0380198251168184511617909252505050919091019250604091505051506116199050600186610c87565b50600060085413156116c857600654600160a060020a03169150816349102e616040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561166757600080fd5b5af1151561167457600080fd5b5050506040518051905090507f9625d8ae00bbc11bf4dbac17532594047d634ea66e1baa95d4974d086f261c30858260405160ff9092168252151560208201526040908101905180910390a18092506116cd565b600092505b505092915050565b60015432600160a060020a039081169116146116f057600080fd5b60048054600160a060020a031916600160a060020a0392909216919091179055565b60025481565b60015432600160a060020a0390811691161461173357600080fd5b60058054600160a060020a031916600160a060020a0383811691909117918290557fa6f2e38f0cfebf27212317fced3ac40bc62e00bd33f38d69603710740c69acb79116604051600160a060020a03909116815260200160405180910390a150565b600154600090819033600160a060020a039081169116148015906117c8575060045433600160a060020a03908116911614155b156117d65760009150611a2e565b506000858152600d602052604090208054600160a060020a031615156117ff5760009150611a2e565b6002810154600090810b900b1515611a29578415156118c55780546001820154600160a060020a039091169080156108fc0290604051600060405180830381858888f19350505050151561185257600080fd5b60028101805460ff191660ff1790557f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b9127188660006040519182521515602082015260806040808301829052600091830182905260c060608401819052830191909152610100909101905180910390a1611a29565b6001805490820154600160a060020a039091169080156108fc0290604051600060405180830381858888f19350505050151561190057600080fd5b60028101805460ff191660019081179091557f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b91271890879086866040518481528315156020820152608060408201818152906060830190830185818151815260200191508051906020019080838360005b8381101561198757808201518382015260200161196f565b50505050905090810190601f1680156119b45780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156119ea5780820151838201526020016119d2565b50505050905090810190601f168015611a175780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a15b600191505b50949350505050565b60015432600160a060020a03908116911614611a5257600080fd5b60028190557f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff8160405190815260200160405180910390a150565b60015432600160a060020a03908116911614611aa857600080fd5b600160a060020a0381161515611abd57600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b60015460009032600160a060020a03908116911614611afd57600080fd5b60078290557f4114f8ef80b6de2161db580cbefa14e1892d15d3ebe2062c9914e4a5773114a38260405190815260200160405180910390a1505060075490565b60065460009033600160a060020a03908116911614611b5e5750600061086e565b50600b8054820190819055919050565b600b5481565b6060604051908101604090815260008083526020830181905290820152905600272aa3d0efd05809b06d40e95792681d6b844a29fe97cd703cfd5eee27510740a165627a7a72305820f010d72b4ef3212cb91349cd2c5958f806c34e964fa1b0e1325044ebfbeef0500029a165627a7a72305820f540d994d3ac2641583ad9a285b176adee88783dac3139271f046465cdcaa8910029`

// DeployBaseLibrary deploys a new Ethereum contract, binding an instance of BaseLibrary to it.
func DeployBaseLibrary(auth *bind.TransactOpts, backend bind.ContractBackend, address_KMS common.Address) (common.Address, *types.Transaction, *BaseLibrary, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseLibraryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaseLibraryBin), backend, address_KMS)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseLibrary{BaseLibraryCaller: BaseLibraryCaller{contract: contract}, BaseLibraryTransactor: BaseLibraryTransactor{contract: contract}, BaseLibraryFilterer: BaseLibraryFilterer{contract: contract}}, nil
}

// BaseLibrary is an auto generated Go binding around an Ethereum contract.
type BaseLibrary struct {
	BaseLibraryCaller     // Read-only binding to the contract
	BaseLibraryTransactor // Write-only binding to the contract
	BaseLibraryFilterer   // Log filterer for contract events
}

// BaseLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseLibrarySession struct {
	Contract     *BaseLibrary      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseLibraryCallerSession struct {
	Contract *BaseLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BaseLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseLibraryTransactorSession struct {
	Contract     *BaseLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BaseLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseLibraryRaw struct {
	Contract *BaseLibrary // Generic contract binding to access the raw methods on
}

// BaseLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseLibraryCallerRaw struct {
	Contract *BaseLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// BaseLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseLibraryTransactorRaw struct {
	Contract *BaseLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseLibrary creates a new instance of BaseLibrary, bound to a specific deployed contract.
func NewBaseLibrary(address common.Address, backend bind.ContractBackend) (*BaseLibrary, error) {
	contract, err := bindBaseLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseLibrary{BaseLibraryCaller: BaseLibraryCaller{contract: contract}, BaseLibraryTransactor: BaseLibraryTransactor{contract: contract}, BaseLibraryFilterer: BaseLibraryFilterer{contract: contract}}, nil
}

// NewBaseLibraryCaller creates a new read-only instance of BaseLibrary, bound to a specific deployed contract.
func NewBaseLibraryCaller(address common.Address, caller bind.ContractCaller) (*BaseLibraryCaller, error) {
	contract, err := bindBaseLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseLibraryCaller{contract: contract}, nil
}

// NewBaseLibraryTransactor creates a new write-only instance of BaseLibrary, bound to a specific deployed contract.
func NewBaseLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseLibraryTransactor, error) {
	contract, err := bindBaseLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseLibraryTransactor{contract: contract}, nil
}

// NewBaseLibraryFilterer creates a new log filterer instance of BaseLibrary, bound to a specific deployed contract.
func NewBaseLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseLibraryFilterer, error) {
	contract, err := bindBaseLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseLibraryFilterer{contract: contract}, nil
}

// bindBaseLibrary binds a generic wrapper to an already deployed contract.
func bindBaseLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseLibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseLibrary *BaseLibraryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseLibrary.Contract.BaseLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseLibrary *BaseLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseLibrary.Contract.BaseLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseLibrary *BaseLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseLibrary.Contract.BaseLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseLibrary *BaseLibraryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseLibrary *BaseLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseLibrary *BaseLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseLibrary.Contract.contract.Transact(opts, method, params...)
}

// AccessorGroups is a free data retrieval call binding the contract method 0x2c11f392.
//
// Solidity: function accessorGroups( uint256) constant returns(address)
func (_BaseLibrary *BaseLibraryCaller) AccessorGroups(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "accessorGroups", arg0)
	return *ret0, err
}

// AccessorGroups is a free data retrieval call binding the contract method 0x2c11f392.
//
// Solidity: function accessorGroups( uint256) constant returns(address)
func (_BaseLibrary *BaseLibrarySession) AccessorGroups(arg0 *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.AccessorGroups(&_BaseLibrary.CallOpts, arg0)
}

// AccessorGroups is a free data retrieval call binding the contract method 0x2c11f392.
//
// Solidity: function accessorGroups( uint256) constant returns(address)
func (_BaseLibrary *BaseLibraryCallerSession) AccessorGroups(arg0 *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.AccessorGroups(&_BaseLibrary.CallOpts, arg0)
}

// AccessorGroupsLength is a free data retrieval call binding the contract method 0xe5538fd2.
//
// Solidity: function accessorGroupsLength() constant returns(uint256)
func (_BaseLibrary *BaseLibraryCaller) AccessorGroupsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "accessorGroupsLength")
	return *ret0, err
}

// AccessorGroupsLength is a free data retrieval call binding the contract method 0xe5538fd2.
//
// Solidity: function accessorGroupsLength() constant returns(uint256)
func (_BaseLibrary *BaseLibrarySession) AccessorGroupsLength() (*big.Int, error) {
	return _BaseLibrary.Contract.AccessorGroupsLength(&_BaseLibrary.CallOpts)
}

// AccessorGroupsLength is a free data retrieval call binding the contract method 0xe5538fd2.
//
// Solidity: function accessorGroupsLength() constant returns(uint256)
func (_BaseLibrary *BaseLibraryCallerSession) AccessorGroupsLength() (*big.Int, error) {
	return _BaseLibrary.Contract.AccessorGroupsLength(&_BaseLibrary.CallOpts)
}

// AddressKMS is a free data retrieval call binding the contract method 0x32eaf21b.
//
// Solidity: function addressKMS() constant returns(address)
func (_BaseLibrary *BaseLibraryCaller) AddressKMS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "addressKMS")
	return *ret0, err
}

// AddressKMS is a free data retrieval call binding the contract method 0x32eaf21b.
//
// Solidity: function addressKMS() constant returns(address)
func (_BaseLibrary *BaseLibrarySession) AddressKMS() (common.Address, error) {
	return _BaseLibrary.Contract.AddressKMS(&_BaseLibrary.CallOpts)
}

// AddressKMS is a free data retrieval call binding the contract method 0x32eaf21b.
//
// Solidity: function addressKMS() constant returns(address)
func (_BaseLibrary *BaseLibraryCallerSession) AddressKMS() (common.Address, error) {
	return _BaseLibrary.Contract.AddressKMS(&_BaseLibrary.CallOpts)
}

// ApprovalRequests is a free data retrieval call binding the contract method 0x8cb13c2e.
//
// Solidity: function approvalRequests( uint256) constant returns(address)
func (_BaseLibrary *BaseLibraryCaller) ApprovalRequests(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "approvalRequests", arg0)
	return *ret0, err
}

// ApprovalRequests is a free data retrieval call binding the contract method 0x8cb13c2e.
//
// Solidity: function approvalRequests( uint256) constant returns(address)
func (_BaseLibrary *BaseLibrarySession) ApprovalRequests(arg0 *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.ApprovalRequests(&_BaseLibrary.CallOpts, arg0)
}

// ApprovalRequests is a free data retrieval call binding the contract method 0x8cb13c2e.
//
// Solidity: function approvalRequests( uint256) constant returns(address)
func (_BaseLibrary *BaseLibraryCallerSession) ApprovalRequests(arg0 *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.ApprovalRequests(&_BaseLibrary.CallOpts, arg0)
}

// ApprovalRequestsLength is a free data retrieval call binding the contract method 0x16308394.
//
// Solidity: function approvalRequestsLength() constant returns(uint256)
func (_BaseLibrary *BaseLibraryCaller) ApprovalRequestsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "approvalRequestsLength")
	return *ret0, err
}

// ApprovalRequestsLength is a free data retrieval call binding the contract method 0x16308394.
//
// Solidity: function approvalRequestsLength() constant returns(uint256)
func (_BaseLibrary *BaseLibrarySession) ApprovalRequestsLength() (*big.Int, error) {
	return _BaseLibrary.Contract.ApprovalRequestsLength(&_BaseLibrary.CallOpts)
}

// ApprovalRequestsLength is a free data retrieval call binding the contract method 0x16308394.
//
// Solidity: function approvalRequestsLength() constant returns(uint256)
func (_BaseLibrary *BaseLibraryCallerSession) ApprovalRequestsLength() (*big.Int, error) {
	return _BaseLibrary.Contract.ApprovalRequestsLength(&_BaseLibrary.CallOpts)
}

// CanContribute is a free data retrieval call binding the contract method 0x0eaec2c5.
//
// Solidity: function canContribute(candidate address) constant returns(bool)
func (_BaseLibrary *BaseLibraryCaller) CanContribute(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "canContribute", candidate)
	return *ret0, err
}

// CanContribute is a free data retrieval call binding the contract method 0x0eaec2c5.
//
// Solidity: function canContribute(candidate address) constant returns(bool)
func (_BaseLibrary *BaseLibrarySession) CanContribute(candidate common.Address) (bool, error) {
	return _BaseLibrary.Contract.CanContribute(&_BaseLibrary.CallOpts, candidate)
}

// CanContribute is a free data retrieval call binding the contract method 0x0eaec2c5.
//
// Solidity: function canContribute(candidate address) constant returns(bool)
func (_BaseLibrary *BaseLibraryCallerSession) CanContribute(candidate common.Address) (bool, error) {
	return _BaseLibrary.Contract.CanContribute(&_BaseLibrary.CallOpts, candidate)
}

// CanReview is a free data retrieval call binding the contract method 0x29d00219.
//
// Solidity: function canReview(candidate address) constant returns(bool)
func (_BaseLibrary *BaseLibraryCaller) CanReview(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "canReview", candidate)
	return *ret0, err
}

// CanReview is a free data retrieval call binding the contract method 0x29d00219.
//
// Solidity: function canReview(candidate address) constant returns(bool)
func (_BaseLibrary *BaseLibrarySession) CanReview(candidate common.Address) (bool, error) {
	return _BaseLibrary.Contract.CanReview(&_BaseLibrary.CallOpts, candidate)
}

// CanReview is a free data retrieval call binding the contract method 0x29d00219.
//
// Solidity: function canReview(candidate address) constant returns(bool)
func (_BaseLibrary *BaseLibraryCallerSession) CanReview(candidate common.Address) (bool, error) {
	return _BaseLibrary.Contract.CanReview(&_BaseLibrary.CallOpts, candidate)
}

// ContentTypeContracts is a free data retrieval call binding the contract method 0x1cdbee5a.
//
// Solidity: function contentTypeContracts( address) constant returns(address)
func (_BaseLibrary *BaseLibraryCaller) ContentTypeContracts(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "contentTypeContracts", arg0)
	return *ret0, err
}

// ContentTypeContracts is a free data retrieval call binding the contract method 0x1cdbee5a.
//
// Solidity: function contentTypeContracts( address) constant returns(address)
func (_BaseLibrary *BaseLibrarySession) ContentTypeContracts(arg0 common.Address) (common.Address, error) {
	return _BaseLibrary.Contract.ContentTypeContracts(&_BaseLibrary.CallOpts, arg0)
}

// ContentTypeContracts is a free data retrieval call binding the contract method 0x1cdbee5a.
//
// Solidity: function contentTypeContracts( address) constant returns(address)
func (_BaseLibrary *BaseLibraryCallerSession) ContentTypeContracts(arg0 common.Address) (common.Address, error) {
	return _BaseLibrary.Contract.ContentTypeContracts(&_BaseLibrary.CallOpts, arg0)
}

// ContentTypeLicensingFees is a free data retrieval call binding the contract method 0x3fbb5d71.
//
// Solidity: function contentTypeLicensingFees( address) constant returns(uint256)
func (_BaseLibrary *BaseLibraryCaller) ContentTypeLicensingFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "contentTypeLicensingFees", arg0)
	return *ret0, err
}

// ContentTypeLicensingFees is a free data retrieval call binding the contract method 0x3fbb5d71.
//
// Solidity: function contentTypeLicensingFees( address) constant returns(uint256)
func (_BaseLibrary *BaseLibrarySession) ContentTypeLicensingFees(arg0 common.Address) (*big.Int, error) {
	return _BaseLibrary.Contract.ContentTypeLicensingFees(&_BaseLibrary.CallOpts, arg0)
}

// ContentTypeLicensingFees is a free data retrieval call binding the contract method 0x3fbb5d71.
//
// Solidity: function contentTypeLicensingFees( address) constant returns(uint256)
func (_BaseLibrary *BaseLibraryCallerSession) ContentTypeLicensingFees(arg0 common.Address) (*big.Int, error) {
	return _BaseLibrary.Contract.ContentTypeLicensingFees(&_BaseLibrary.CallOpts, arg0)
}

// ContentTypes is a free data retrieval call binding the contract method 0x991a3a7c.
//
// Solidity: function contentTypes( uint256) constant returns(address)
func (_BaseLibrary *BaseLibraryCaller) ContentTypes(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "contentTypes", arg0)
	return *ret0, err
}

// ContentTypes is a free data retrieval call binding the contract method 0x991a3a7c.
//
// Solidity: function contentTypes( uint256) constant returns(address)
func (_BaseLibrary *BaseLibrarySession) ContentTypes(arg0 *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.ContentTypes(&_BaseLibrary.CallOpts, arg0)
}

// ContentTypes is a free data retrieval call binding the contract method 0x991a3a7c.
//
// Solidity: function contentTypes( uint256) constant returns(address)
func (_BaseLibrary *BaseLibraryCallerSession) ContentTypes(arg0 *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.ContentTypes(&_BaseLibrary.CallOpts, arg0)
}

// ContentTypesLength is a free data retrieval call binding the contract method 0xc65bcbe2.
//
// Solidity: function contentTypesLength() constant returns(uint256)
func (_BaseLibrary *BaseLibraryCaller) ContentTypesLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "contentTypesLength")
	return *ret0, err
}

// ContentTypesLength is a free data retrieval call binding the contract method 0xc65bcbe2.
//
// Solidity: function contentTypesLength() constant returns(uint256)
func (_BaseLibrary *BaseLibrarySession) ContentTypesLength() (*big.Int, error) {
	return _BaseLibrary.Contract.ContentTypesLength(&_BaseLibrary.CallOpts)
}

// ContentTypesLength is a free data retrieval call binding the contract method 0xc65bcbe2.
//
// Solidity: function contentTypesLength() constant returns(uint256)
func (_BaseLibrary *BaseLibraryCallerSession) ContentTypesLength() (*big.Int, error) {
	return _BaseLibrary.Contract.ContentTypesLength(&_BaseLibrary.CallOpts)
}

// ContributorGroups is a free data retrieval call binding the contract method 0x2393553b.
//
// Solidity: function contributorGroups( uint256) constant returns(address)
func (_BaseLibrary *BaseLibraryCaller) ContributorGroups(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "contributorGroups", arg0)
	return *ret0, err
}

// ContributorGroups is a free data retrieval call binding the contract method 0x2393553b.
//
// Solidity: function contributorGroups( uint256) constant returns(address)
func (_BaseLibrary *BaseLibrarySession) ContributorGroups(arg0 *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.ContributorGroups(&_BaseLibrary.CallOpts, arg0)
}

// ContributorGroups is a free data retrieval call binding the contract method 0x2393553b.
//
// Solidity: function contributorGroups( uint256) constant returns(address)
func (_BaseLibrary *BaseLibraryCallerSession) ContributorGroups(arg0 *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.ContributorGroups(&_BaseLibrary.CallOpts, arg0)
}

// ContributorGroupsLength is a free data retrieval call binding the contract method 0x470750bb.
//
// Solidity: function contributorGroupsLength() constant returns(uint256)
func (_BaseLibrary *BaseLibraryCaller) ContributorGroupsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "contributorGroupsLength")
	return *ret0, err
}

// ContributorGroupsLength is a free data retrieval call binding the contract method 0x470750bb.
//
// Solidity: function contributorGroupsLength() constant returns(uint256)
func (_BaseLibrary *BaseLibrarySession) ContributorGroupsLength() (*big.Int, error) {
	return _BaseLibrary.Contract.ContributorGroupsLength(&_BaseLibrary.CallOpts)
}

// ContributorGroupsLength is a free data retrieval call binding the contract method 0x470750bb.
//
// Solidity: function contributorGroupsLength() constant returns(uint256)
func (_BaseLibrary *BaseLibraryCallerSession) ContributorGroupsLength() (*big.Int, error) {
	return _BaseLibrary.Contract.ContributorGroupsLength(&_BaseLibrary.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseLibrary *BaseLibraryCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseLibrary *BaseLibrarySession) Creator() (common.Address, error) {
	return _BaseLibrary.Contract.Creator(&_BaseLibrary.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseLibrary *BaseLibraryCallerSession) Creator() (common.Address, error) {
	return _BaseLibrary.Contract.Creator(&_BaseLibrary.CallOpts)
}

// GetPendingApprovalRequest is a free data retrieval call binding the contract method 0x63dab9d4.
//
// Solidity: function getPendingApprovalRequest(index uint256) constant returns(address)
func (_BaseLibrary *BaseLibraryCaller) GetPendingApprovalRequest(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "getPendingApprovalRequest", index)
	return *ret0, err
}

// GetPendingApprovalRequest is a free data retrieval call binding the contract method 0x63dab9d4.
//
// Solidity: function getPendingApprovalRequest(index uint256) constant returns(address)
func (_BaseLibrary *BaseLibrarySession) GetPendingApprovalRequest(index *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.GetPendingApprovalRequest(&_BaseLibrary.CallOpts, index)
}

// GetPendingApprovalRequest is a free data retrieval call binding the contract method 0x63dab9d4.
//
// Solidity: function getPendingApprovalRequest(index uint256) constant returns(address)
func (_BaseLibrary *BaseLibraryCallerSession) GetPendingApprovalRequest(index *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.GetPendingApprovalRequest(&_BaseLibrary.CallOpts, index)
}

// HasAccess is a free data retrieval call binding the contract method 0x95a078e8.
//
// Solidity: function hasAccess(candidate address) constant returns(bool)
func (_BaseLibrary *BaseLibraryCaller) HasAccess(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "hasAccess", candidate)
	return *ret0, err
}

// HasAccess is a free data retrieval call binding the contract method 0x95a078e8.
//
// Solidity: function hasAccess(candidate address) constant returns(bool)
func (_BaseLibrary *BaseLibrarySession) HasAccess(candidate common.Address) (bool, error) {
	return _BaseLibrary.Contract.HasAccess(&_BaseLibrary.CallOpts, candidate)
}

// HasAccess is a free data retrieval call binding the contract method 0x95a078e8.
//
// Solidity: function hasAccess(candidate address) constant returns(bool)
func (_BaseLibrary *BaseLibraryCallerSession) HasAccess(candidate common.Address) (bool, error) {
	return _BaseLibrary.Contract.HasAccess(&_BaseLibrary.CallOpts, candidate)
}

// ObjectHash is a free data retrieval call binding the contract method 0xe02dd9c2.
//
// Solidity: function objectHash() constant returns(bytes32)
func (_BaseLibrary *BaseLibraryCaller) ObjectHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "objectHash")
	return *ret0, err
}

// ObjectHash is a free data retrieval call binding the contract method 0xe02dd9c2.
//
// Solidity: function objectHash() constant returns(bytes32)
func (_BaseLibrary *BaseLibrarySession) ObjectHash() ([32]byte, error) {
	return _BaseLibrary.Contract.ObjectHash(&_BaseLibrary.CallOpts)
}

// ObjectHash is a free data retrieval call binding the contract method 0xe02dd9c2.
//
// Solidity: function objectHash() constant returns(bytes32)
func (_BaseLibrary *BaseLibraryCallerSession) ObjectHash() ([32]byte, error) {
	return _BaseLibrary.Contract.ObjectHash(&_BaseLibrary.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseLibrary *BaseLibraryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseLibrary *BaseLibrarySession) Owner() (common.Address, error) {
	return _BaseLibrary.Contract.Owner(&_BaseLibrary.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseLibrary *BaseLibraryCallerSession) Owner() (common.Address, error) {
	return _BaseLibrary.Contract.Owner(&_BaseLibrary.CallOpts)
}

// ReviewerGroups is a free data retrieval call binding the contract method 0x952e464b.
//
// Solidity: function reviewerGroups( uint256) constant returns(address)
func (_BaseLibrary *BaseLibraryCaller) ReviewerGroups(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "reviewerGroups", arg0)
	return *ret0, err
}

// ReviewerGroups is a free data retrieval call binding the contract method 0x952e464b.
//
// Solidity: function reviewerGroups( uint256) constant returns(address)
func (_BaseLibrary *BaseLibrarySession) ReviewerGroups(arg0 *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.ReviewerGroups(&_BaseLibrary.CallOpts, arg0)
}

// ReviewerGroups is a free data retrieval call binding the contract method 0x952e464b.
//
// Solidity: function reviewerGroups( uint256) constant returns(address)
func (_BaseLibrary *BaseLibraryCallerSession) ReviewerGroups(arg0 *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.ReviewerGroups(&_BaseLibrary.CallOpts, arg0)
}

// ReviewerGroupsLength is a free data retrieval call binding the contract method 0x21770a84.
//
// Solidity: function reviewerGroupsLength() constant returns(uint256)
func (_BaseLibrary *BaseLibraryCaller) ReviewerGroupsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "reviewerGroupsLength")
	return *ret0, err
}

// ReviewerGroupsLength is a free data retrieval call binding the contract method 0x21770a84.
//
// Solidity: function reviewerGroupsLength() constant returns(uint256)
func (_BaseLibrary *BaseLibrarySession) ReviewerGroupsLength() (*big.Int, error) {
	return _BaseLibrary.Contract.ReviewerGroupsLength(&_BaseLibrary.CallOpts)
}

// ReviewerGroupsLength is a free data retrieval call binding the contract method 0x21770a84.
//
// Solidity: function reviewerGroupsLength() constant returns(uint256)
func (_BaseLibrary *BaseLibraryCallerSession) ReviewerGroupsLength() (*big.Int, error) {
	return _BaseLibrary.Contract.ReviewerGroupsLength(&_BaseLibrary.CallOpts)
}

// Space is a free data retrieval call binding the contract method 0x27af9e38.
//
// Solidity: function space() constant returns(address)
func (_BaseLibrary *BaseLibraryCaller) Space(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "space")
	return *ret0, err
}

// Space is a free data retrieval call binding the contract method 0x27af9e38.
//
// Solidity: function space() constant returns(address)
func (_BaseLibrary *BaseLibrarySession) Space() (common.Address, error) {
	return _BaseLibrary.Contract.Space(&_BaseLibrary.CallOpts)
}

// Space is a free data retrieval call binding the contract method 0x27af9e38.
//
// Solidity: function space() constant returns(address)
func (_BaseLibrary *BaseLibraryCallerSession) Space() (common.Address, error) {
	return _BaseLibrary.Contract.Space(&_BaseLibrary.CallOpts)
}

// AccessRequest is a paid mutator transaction binding the contract method 0xf1551887.
//
// Solidity: function accessRequest() returns(bool)
func (_BaseLibrary *BaseLibraryTransactor) AccessRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "accessRequest")
}

// AccessRequest is a paid mutator transaction binding the contract method 0xf1551887.
//
// Solidity: function accessRequest() returns(bool)
func (_BaseLibrary *BaseLibrarySession) AccessRequest() (*types.Transaction, error) {
	return _BaseLibrary.Contract.AccessRequest(&_BaseLibrary.TransactOpts)
}

// AccessRequest is a paid mutator transaction binding the contract method 0xf1551887.
//
// Solidity: function accessRequest() returns(bool)
func (_BaseLibrary *BaseLibraryTransactorSession) AccessRequest() (*types.Transaction, error) {
	return _BaseLibrary.Contract.AccessRequest(&_BaseLibrary.TransactOpts)
}

// AddAccessorGroup is a paid mutator transaction binding the contract method 0x1d0f4351.
//
// Solidity: function addAccessorGroup(group address) returns()
func (_BaseLibrary *BaseLibraryTransactor) AddAccessorGroup(opts *bind.TransactOpts, group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "addAccessorGroup", group)
}

// AddAccessorGroup is a paid mutator transaction binding the contract method 0x1d0f4351.
//
// Solidity: function addAccessorGroup(group address) returns()
func (_BaseLibrary *BaseLibrarySession) AddAccessorGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.AddAccessorGroup(&_BaseLibrary.TransactOpts, group)
}

// AddAccessorGroup is a paid mutator transaction binding the contract method 0x1d0f4351.
//
// Solidity: function addAccessorGroup(group address) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) AddAccessorGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.AddAccessorGroup(&_BaseLibrary.TransactOpts, group)
}

// AddContentType is a paid mutator transaction binding the contract method 0xa90229d5.
//
// Solidity: function addContentType(content_type address, content_contract address, licensing_fee uint256) returns()
func (_BaseLibrary *BaseLibraryTransactor) AddContentType(opts *bind.TransactOpts, content_type common.Address, content_contract common.Address, licensing_fee *big.Int) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "addContentType", content_type, content_contract, licensing_fee)
}

// AddContentType is a paid mutator transaction binding the contract method 0xa90229d5.
//
// Solidity: function addContentType(content_type address, content_contract address, licensing_fee uint256) returns()
func (_BaseLibrary *BaseLibrarySession) AddContentType(content_type common.Address, content_contract common.Address, licensing_fee *big.Int) (*types.Transaction, error) {
	return _BaseLibrary.Contract.AddContentType(&_BaseLibrary.TransactOpts, content_type, content_contract, licensing_fee)
}

// AddContentType is a paid mutator transaction binding the contract method 0xa90229d5.
//
// Solidity: function addContentType(content_type address, content_contract address, licensing_fee uint256) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) AddContentType(content_type common.Address, content_contract common.Address, licensing_fee *big.Int) (*types.Transaction, error) {
	return _BaseLibrary.Contract.AddContentType(&_BaseLibrary.TransactOpts, content_type, content_contract, licensing_fee)
}

// AddContributorGroup is a paid mutator transaction binding the contract method 0x679a9a3c.
//
// Solidity: function addContributorGroup(group address) returns()
func (_BaseLibrary *BaseLibraryTransactor) AddContributorGroup(opts *bind.TransactOpts, group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "addContributorGroup", group)
}

// AddContributorGroup is a paid mutator transaction binding the contract method 0x679a9a3c.
//
// Solidity: function addContributorGroup(group address) returns()
func (_BaseLibrary *BaseLibrarySession) AddContributorGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.AddContributorGroup(&_BaseLibrary.TransactOpts, group)
}

// AddContributorGroup is a paid mutator transaction binding the contract method 0x679a9a3c.
//
// Solidity: function addContributorGroup(group address) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) AddContributorGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.AddContributorGroup(&_BaseLibrary.TransactOpts, group)
}

// AddReviewerGroup is a paid mutator transaction binding the contract method 0xdc3c29c0.
//
// Solidity: function addReviewerGroup(group address) returns()
func (_BaseLibrary *BaseLibraryTransactor) AddReviewerGroup(opts *bind.TransactOpts, group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "addReviewerGroup", group)
}

// AddReviewerGroup is a paid mutator transaction binding the contract method 0xdc3c29c0.
//
// Solidity: function addReviewerGroup(group address) returns()
func (_BaseLibrary *BaseLibrarySession) AddReviewerGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.AddReviewerGroup(&_BaseLibrary.TransactOpts, group)
}

// AddReviewerGroup is a paid mutator transaction binding the contract method 0xdc3c29c0.
//
// Solidity: function addReviewerGroup(group address) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) AddReviewerGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.AddReviewerGroup(&_BaseLibrary.TransactOpts, group)
}

// ApproveContentExecuted is a paid mutator transaction binding the contract method 0xbf81a0b6.
//
// Solidity: function approveContentExecuted(content_contract address, approved bool, note string) returns(bool)
func (_BaseLibrary *BaseLibraryTransactor) ApproveContentExecuted(opts *bind.TransactOpts, content_contract common.Address, approved bool, note string) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "approveContentExecuted", content_contract, approved, note)
}

// ApproveContentExecuted is a paid mutator transaction binding the contract method 0xbf81a0b6.
//
// Solidity: function approveContentExecuted(content_contract address, approved bool, note string) returns(bool)
func (_BaseLibrary *BaseLibrarySession) ApproveContentExecuted(content_contract common.Address, approved bool, note string) (*types.Transaction, error) {
	return _BaseLibrary.Contract.ApproveContentExecuted(&_BaseLibrary.TransactOpts, content_contract, approved, note)
}

// ApproveContentExecuted is a paid mutator transaction binding the contract method 0xbf81a0b6.
//
// Solidity: function approveContentExecuted(content_contract address, approved bool, note string) returns(bool)
func (_BaseLibrary *BaseLibraryTransactorSession) ApproveContentExecuted(content_contract common.Address, approved bool, note string) (*types.Transaction, error) {
	return _BaseLibrary.Contract.ApproveContentExecuted(&_BaseLibrary.TransactOpts, content_contract, approved, note)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(object_hash bytes32) returns()
func (_BaseLibrary *BaseLibraryTransactor) Commit(opts *bind.TransactOpts, object_hash [32]byte) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "commit", object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(object_hash bytes32) returns()
func (_BaseLibrary *BaseLibrarySession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _BaseLibrary.Contract.Commit(&_BaseLibrary.TransactOpts, object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(object_hash bytes32) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _BaseLibrary.Contract.Commit(&_BaseLibrary.TransactOpts, object_hash)
}

// CreateContent is a paid mutator transaction binding the contract method 0x1e35d8fa.
//
// Solidity: function createContent(content_type address) returns(address)
func (_BaseLibrary *BaseLibraryTransactor) CreateContent(opts *bind.TransactOpts, content_type common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "createContent", content_type)
}

// CreateContent is a paid mutator transaction binding the contract method 0x1e35d8fa.
//
// Solidity: function createContent(content_type address) returns(address)
func (_BaseLibrary *BaseLibrarySession) CreateContent(content_type common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.CreateContent(&_BaseLibrary.TransactOpts, content_type)
}

// CreateContent is a paid mutator transaction binding the contract method 0x1e35d8fa.
//
// Solidity: function createContent(content_type address) returns(address)
func (_BaseLibrary *BaseLibraryTransactorSession) CreateContent(content_type common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.CreateContent(&_BaseLibrary.TransactOpts, content_type)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseLibrary *BaseLibraryTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseLibrary *BaseLibrarySession) Kill() (*types.Transaction, error) {
	return _BaseLibrary.Contract.Kill(&_BaseLibrary.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseLibrary *BaseLibraryTransactorSession) Kill() (*types.Transaction, error) {
	return _BaseLibrary.Contract.Kill(&_BaseLibrary.TransactOpts)
}

// PayCredit is a paid mutator transaction binding the contract method 0x6934d45f.
//
// Solidity: function payCredit(content_contract address, amount uint256) returns(uint256)
func (_BaseLibrary *BaseLibraryTransactor) PayCredit(opts *bind.TransactOpts, content_contract common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "payCredit", content_contract, amount)
}

// PayCredit is a paid mutator transaction binding the contract method 0x6934d45f.
//
// Solidity: function payCredit(content_contract address, amount uint256) returns(uint256)
func (_BaseLibrary *BaseLibrarySession) PayCredit(content_contract common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseLibrary.Contract.PayCredit(&_BaseLibrary.TransactOpts, content_contract, amount)
}

// PayCredit is a paid mutator transaction binding the contract method 0x6934d45f.
//
// Solidity: function payCredit(content_contract address, amount uint256) returns(uint256)
func (_BaseLibrary *BaseLibraryTransactorSession) PayCredit(content_contract common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaseLibrary.Contract.PayCredit(&_BaseLibrary.TransactOpts, content_contract, amount)
}

// SetAddressKMS is a paid mutator transaction binding the contract method 0xc9e8e72d.
//
// Solidity: function setAddressKMS(address_KMS address) returns()
func (_BaseLibrary *BaseLibraryTransactor) SetAddressKMS(opts *bind.TransactOpts, address_KMS common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "setAddressKMS", address_KMS)
}

// SetAddressKMS is a paid mutator transaction binding the contract method 0xc9e8e72d.
//
// Solidity: function setAddressKMS(address_KMS address) returns()
func (_BaseLibrary *BaseLibrarySession) SetAddressKMS(address_KMS common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.SetAddressKMS(&_BaseLibrary.TransactOpts, address_KMS)
}

// SetAddressKMS is a paid mutator transaction binding the contract method 0xc9e8e72d.
//
// Solidity: function setAddressKMS(address_KMS address) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) SetAddressKMS(address_KMS common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.SetAddressKMS(&_BaseLibrary.TransactOpts, address_KMS)
}

// SubmitApprovalRequest is a paid mutator transaction binding the contract method 0x49102e61.
//
// Solidity: function submitApprovalRequest() returns(bool)
func (_BaseLibrary *BaseLibraryTransactor) SubmitApprovalRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "submitApprovalRequest")
}

// SubmitApprovalRequest is a paid mutator transaction binding the contract method 0x49102e61.
//
// Solidity: function submitApprovalRequest() returns(bool)
func (_BaseLibrary *BaseLibrarySession) SubmitApprovalRequest() (*types.Transaction, error) {
	return _BaseLibrary.Contract.SubmitApprovalRequest(&_BaseLibrary.TransactOpts)
}

// SubmitApprovalRequest is a paid mutator transaction binding the contract method 0x49102e61.
//
// Solidity: function submitApprovalRequest() returns(bool)
func (_BaseLibrary *BaseLibraryTransactorSession) SubmitApprovalRequest() (*types.Transaction, error) {
	return _BaseLibrary.Contract.SubmitApprovalRequest(&_BaseLibrary.TransactOpts)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_BaseLibrary *BaseLibraryTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_BaseLibrary *BaseLibrarySession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.TransferCreatorship(&_BaseLibrary.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.TransferCreatorship(&_BaseLibrary.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseLibrary *BaseLibraryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseLibrary *BaseLibrarySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.TransferOwnership(&_BaseLibrary.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.TransferOwnership(&_BaseLibrary.TransactOpts, newOwner)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xc287e0ed.
//
// Solidity: function updateRequest() returns()
func (_BaseLibrary *BaseLibraryTransactor) UpdateRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "updateRequest")
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xc287e0ed.
//
// Solidity: function updateRequest() returns()
func (_BaseLibrary *BaseLibrarySession) UpdateRequest() (*types.Transaction, error) {
	return _BaseLibrary.Contract.UpdateRequest(&_BaseLibrary.TransactOpts)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xc287e0ed.
//
// Solidity: function updateRequest() returns()
func (_BaseLibrary *BaseLibraryTransactorSession) UpdateRequest() (*types.Transaction, error) {
	return _BaseLibrary.Contract.UpdateRequest(&_BaseLibrary.TransactOpts)
}

// BaseLibraryAccessRequestIterator is returned from FilterAccessRequest and is used to iterate over the raw logs and unpacked data for AccessRequest events raised by the BaseLibrary contract.
type BaseLibraryAccessRequestIterator struct {
	Event *BaseLibraryAccessRequest // Event containing the contract specifics and raw log

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
func (it *BaseLibraryAccessRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryAccessRequest)
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
		it.Event = new(BaseLibraryAccessRequest)
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
func (it *BaseLibraryAccessRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryAccessRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryAccessRequest represents a AccessRequest event raised by the BaseLibrary contract.
type BaseLibraryAccessRequest struct {
	RequestValidity *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccessRequest is a free log retrieval operation binding the contract event 0xca8116fd70ecea0cf35f7c08eb31491cce5ebd6bafef5bd6426862df29840463.
//
// Solidity: e AccessRequest(requestValidity uint256)
func (_BaseLibrary *BaseLibraryFilterer) FilterAccessRequest(opts *bind.FilterOpts) (*BaseLibraryAccessRequestIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "AccessRequest")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryAccessRequestIterator{contract: _BaseLibrary.contract, event: "AccessRequest", logs: logs, sub: sub}, nil
}

// WatchAccessRequest is a free log subscription operation binding the contract event 0xca8116fd70ecea0cf35f7c08eb31491cce5ebd6bafef5bd6426862df29840463.
//
// Solidity: e AccessRequest(requestValidity uint256)
func (_BaseLibrary *BaseLibraryFilterer) WatchAccessRequest(opts *bind.WatchOpts, sink chan<- *BaseLibraryAccessRequest) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "AccessRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryAccessRequest)
				if err := _BaseLibrary.contract.UnpackLog(event, "AccessRequest", log); err != nil {
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

// BaseLibraryAccessorGroupAddedIterator is returned from FilterAccessorGroupAdded and is used to iterate over the raw logs and unpacked data for AccessorGroupAdded events raised by the BaseLibrary contract.
type BaseLibraryAccessorGroupAddedIterator struct {
	Event *BaseLibraryAccessorGroupAdded // Event containing the contract specifics and raw log

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
func (it *BaseLibraryAccessorGroupAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryAccessorGroupAdded)
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
		it.Event = new(BaseLibraryAccessorGroupAdded)
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
func (it *BaseLibraryAccessorGroupAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryAccessorGroupAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryAccessorGroupAdded represents a AccessorGroupAdded event raised by the BaseLibrary contract.
type BaseLibraryAccessorGroupAdded struct {
	Group common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAccessorGroupAdded is a free log retrieval operation binding the contract event 0x3a94857e4393737f73edb175a7d0c195c7f635d9ae995e12740616ec55c9d411.
//
// Solidity: e AccessorGroupAdded(group address)
func (_BaseLibrary *BaseLibraryFilterer) FilterAccessorGroupAdded(opts *bind.FilterOpts) (*BaseLibraryAccessorGroupAddedIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "AccessorGroupAdded")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryAccessorGroupAddedIterator{contract: _BaseLibrary.contract, event: "AccessorGroupAdded", logs: logs, sub: sub}, nil
}

// WatchAccessorGroupAdded is a free log subscription operation binding the contract event 0x3a94857e4393737f73edb175a7d0c195c7f635d9ae995e12740616ec55c9d411.
//
// Solidity: e AccessorGroupAdded(group address)
func (_BaseLibrary *BaseLibraryFilterer) WatchAccessorGroupAdded(opts *bind.WatchOpts, sink chan<- *BaseLibraryAccessorGroupAdded) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "AccessorGroupAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryAccessorGroupAdded)
				if err := _BaseLibrary.contract.UnpackLog(event, "AccessorGroupAdded", log); err != nil {
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

// BaseLibraryApproveContentExecutedIterator is returned from FilterApproveContentExecuted and is used to iterate over the raw logs and unpacked data for ApproveContentExecuted events raised by the BaseLibrary contract.
type BaseLibraryApproveContentExecutedIterator struct {
	Event *BaseLibraryApproveContentExecuted // Event containing the contract specifics and raw log

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
func (it *BaseLibraryApproveContentExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryApproveContentExecuted)
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
		it.Event = new(BaseLibraryApproveContentExecuted)
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
func (it *BaseLibraryApproveContentExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryApproveContentExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryApproveContentExecuted represents a ApproveContentExecuted event raised by the BaseLibrary contract.
type BaseLibraryApproveContentExecuted struct {
	ContentAddress common.Address
	Approved       bool
	Note           string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterApproveContentExecuted is a free log retrieval operation binding the contract event 0x8a6b83b9f4a16d1fd3ec333d40e20591cfb3ad7c24ecbfbc9a5c5b1a6445cab4.
//
// Solidity: e ApproveContentExecuted(contentAddress address, approved bool, note string)
func (_BaseLibrary *BaseLibraryFilterer) FilterApproveContentExecuted(opts *bind.FilterOpts) (*BaseLibraryApproveContentExecutedIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "ApproveContentExecuted")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryApproveContentExecutedIterator{contract: _BaseLibrary.contract, event: "ApproveContentExecuted", logs: logs, sub: sub}, nil
}

// WatchApproveContentExecuted is a free log subscription operation binding the contract event 0x8a6b83b9f4a16d1fd3ec333d40e20591cfb3ad7c24ecbfbc9a5c5b1a6445cab4.
//
// Solidity: e ApproveContentExecuted(contentAddress address, approved bool, note string)
func (_BaseLibrary *BaseLibraryFilterer) WatchApproveContentExecuted(opts *bind.WatchOpts, sink chan<- *BaseLibraryApproveContentExecuted) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "ApproveContentExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryApproveContentExecuted)
				if err := _BaseLibrary.contract.UnpackLog(event, "ApproveContentExecuted", log); err != nil {
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

// BaseLibraryApproveContentRequestIterator is returned from FilterApproveContentRequest and is used to iterate over the raw logs and unpacked data for ApproveContentRequest events raised by the BaseLibrary contract.
type BaseLibraryApproveContentRequestIterator struct {
	Event *BaseLibraryApproveContentRequest // Event containing the contract specifics and raw log

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
func (it *BaseLibraryApproveContentRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryApproveContentRequest)
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
		it.Event = new(BaseLibraryApproveContentRequest)
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
func (it *BaseLibraryApproveContentRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryApproveContentRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryApproveContentRequest represents a ApproveContentRequest event raised by the BaseLibrary contract.
type BaseLibraryApproveContentRequest struct {
	ContentAddress common.Address
	Submitter      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterApproveContentRequest is a free log retrieval operation binding the contract event 0x0588a34cf0de4e025d359c89ca4bacbcbf175440909952d91c814412d9da996a.
//
// Solidity: e ApproveContentRequest(contentAddress address, submitter address)
func (_BaseLibrary *BaseLibraryFilterer) FilterApproveContentRequest(opts *bind.FilterOpts) (*BaseLibraryApproveContentRequestIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "ApproveContentRequest")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryApproveContentRequestIterator{contract: _BaseLibrary.contract, event: "ApproveContentRequest", logs: logs, sub: sub}, nil
}

// WatchApproveContentRequest is a free log subscription operation binding the contract event 0x0588a34cf0de4e025d359c89ca4bacbcbf175440909952d91c814412d9da996a.
//
// Solidity: e ApproveContentRequest(contentAddress address, submitter address)
func (_BaseLibrary *BaseLibraryFilterer) WatchApproveContentRequest(opts *bind.WatchOpts, sink chan<- *BaseLibraryApproveContentRequest) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "ApproveContentRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryApproveContentRequest)
				if err := _BaseLibrary.contract.UnpackLog(event, "ApproveContentRequest", log); err != nil {
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

// BaseLibraryCommitIterator is returned from FilterCommit and is used to iterate over the raw logs and unpacked data for Commit events raised by the BaseLibrary contract.
type BaseLibraryCommitIterator struct {
	Event *BaseLibraryCommit // Event containing the contract specifics and raw log

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
func (it *BaseLibraryCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryCommit)
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
		it.Event = new(BaseLibraryCommit)
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
func (it *BaseLibraryCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryCommit represents a Commit event raised by the BaseLibrary contract.
type BaseLibraryCommit struct {
	ObjectHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommit is a free log retrieval operation binding the contract event 0x9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff.
//
// Solidity: e Commit(objectHash bytes32)
func (_BaseLibrary *BaseLibraryFilterer) FilterCommit(opts *bind.FilterOpts) (*BaseLibraryCommitIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "Commit")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryCommitIterator{contract: _BaseLibrary.contract, event: "Commit", logs: logs, sub: sub}, nil
}

// WatchCommit is a free log subscription operation binding the contract event 0x9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff.
//
// Solidity: e Commit(objectHash bytes32)
func (_BaseLibrary *BaseLibraryFilterer) WatchCommit(opts *bind.WatchOpts, sink chan<- *BaseLibraryCommit) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "Commit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryCommit)
				if err := _BaseLibrary.contract.UnpackLog(event, "Commit", log); err != nil {
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

// BaseLibraryContentObjectCreatedIterator is returned from FilterContentObjectCreated and is used to iterate over the raw logs and unpacked data for ContentObjectCreated events raised by the BaseLibrary contract.
type BaseLibraryContentObjectCreatedIterator struct {
	Event *BaseLibraryContentObjectCreated // Event containing the contract specifics and raw log

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
func (it *BaseLibraryContentObjectCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryContentObjectCreated)
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
		it.Event = new(BaseLibraryContentObjectCreated)
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
func (it *BaseLibraryContentObjectCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryContentObjectCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryContentObjectCreated represents a ContentObjectCreated event raised by the BaseLibrary contract.
type BaseLibraryContentObjectCreated struct {
	ContentAddress common.Address
	ContentType    common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContentObjectCreated is a free log retrieval operation binding the contract event 0x3981e74ab81857b375ec391a4f7c31ee89462cd927de6d8fbdb98f77da009c56.
//
// Solidity: e ContentObjectCreated(contentAddress address, content_type address)
func (_BaseLibrary *BaseLibraryFilterer) FilterContentObjectCreated(opts *bind.FilterOpts) (*BaseLibraryContentObjectCreatedIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "ContentObjectCreated")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryContentObjectCreatedIterator{contract: _BaseLibrary.contract, event: "ContentObjectCreated", logs: logs, sub: sub}, nil
}

// WatchContentObjectCreated is a free log subscription operation binding the contract event 0x3981e74ab81857b375ec391a4f7c31ee89462cd927de6d8fbdb98f77da009c56.
//
// Solidity: e ContentObjectCreated(contentAddress address, content_type address)
func (_BaseLibrary *BaseLibraryFilterer) WatchContentObjectCreated(opts *bind.WatchOpts, sink chan<- *BaseLibraryContentObjectCreated) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "ContentObjectCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryContentObjectCreated)
				if err := _BaseLibrary.contract.UnpackLog(event, "ContentObjectCreated", log); err != nil {
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

// BaseLibraryContentTypeAddedIterator is returned from FilterContentTypeAdded and is used to iterate over the raw logs and unpacked data for ContentTypeAdded events raised by the BaseLibrary contract.
type BaseLibraryContentTypeAddedIterator struct {
	Event *BaseLibraryContentTypeAdded // Event containing the contract specifics and raw log

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
func (it *BaseLibraryContentTypeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryContentTypeAdded)
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
		it.Event = new(BaseLibraryContentTypeAdded)
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
func (it *BaseLibraryContentTypeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryContentTypeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryContentTypeAdded represents a ContentTypeAdded event raised by the BaseLibrary contract.
type BaseLibraryContentTypeAdded struct {
	ContentType     common.Address
	ContentContract common.Address
	LicensingFee    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContentTypeAdded is a free log retrieval operation binding the contract event 0x62c1f342f2a52ee6ac3eeef54df8c836fb31869e9be7c5e5394c042d131c9ecc.
//
// Solidity: e ContentTypeAdded(contentType address, contentContract address, licensingFee uint256)
func (_BaseLibrary *BaseLibraryFilterer) FilterContentTypeAdded(opts *bind.FilterOpts) (*BaseLibraryContentTypeAddedIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "ContentTypeAdded")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryContentTypeAddedIterator{contract: _BaseLibrary.contract, event: "ContentTypeAdded", logs: logs, sub: sub}, nil
}

// WatchContentTypeAdded is a free log subscription operation binding the contract event 0x62c1f342f2a52ee6ac3eeef54df8c836fb31869e9be7c5e5394c042d131c9ecc.
//
// Solidity: e ContentTypeAdded(contentType address, contentContract address, licensingFee uint256)
func (_BaseLibrary *BaseLibraryFilterer) WatchContentTypeAdded(opts *bind.WatchOpts, sink chan<- *BaseLibraryContentTypeAdded) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "ContentTypeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryContentTypeAdded)
				if err := _BaseLibrary.contract.UnpackLog(event, "ContentTypeAdded", log); err != nil {
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

// BaseLibraryContributorGroupAddedIterator is returned from FilterContributorGroupAdded and is used to iterate over the raw logs and unpacked data for ContributorGroupAdded events raised by the BaseLibrary contract.
type BaseLibraryContributorGroupAddedIterator struct {
	Event *BaseLibraryContributorGroupAdded // Event containing the contract specifics and raw log

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
func (it *BaseLibraryContributorGroupAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryContributorGroupAdded)
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
		it.Event = new(BaseLibraryContributorGroupAdded)
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
func (it *BaseLibraryContributorGroupAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryContributorGroupAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryContributorGroupAdded represents a ContributorGroupAdded event raised by the BaseLibrary contract.
type BaseLibraryContributorGroupAdded struct {
	Group common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterContributorGroupAdded is a free log retrieval operation binding the contract event 0x218673669018c25b89bfbf1b58d0075e37c8847ef16e707b92355b7833e97d61.
//
// Solidity: e ContributorGroupAdded(group address)
func (_BaseLibrary *BaseLibraryFilterer) FilterContributorGroupAdded(opts *bind.FilterOpts) (*BaseLibraryContributorGroupAddedIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "ContributorGroupAdded")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryContributorGroupAddedIterator{contract: _BaseLibrary.contract, event: "ContributorGroupAdded", logs: logs, sub: sub}, nil
}

// WatchContributorGroupAdded is a free log subscription operation binding the contract event 0x218673669018c25b89bfbf1b58d0075e37c8847ef16e707b92355b7833e97d61.
//
// Solidity: e ContributorGroupAdded(group address)
func (_BaseLibrary *BaseLibraryFilterer) WatchContributorGroupAdded(opts *bind.WatchOpts, sink chan<- *BaseLibraryContributorGroupAdded) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "ContributorGroupAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryContributorGroupAdded)
				if err := _BaseLibrary.contract.UnpackLog(event, "ContributorGroupAdded", log); err != nil {
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

// BaseLibraryPayCreditIterator is returned from FilterPayCredit and is used to iterate over the raw logs and unpacked data for PayCredit events raised by the BaseLibrary contract.
type BaseLibraryPayCreditIterator struct {
	Event *BaseLibraryPayCredit // Event containing the contract specifics and raw log

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
func (it *BaseLibraryPayCreditIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryPayCredit)
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
		it.Event = new(BaseLibraryPayCredit)
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
func (it *BaseLibraryPayCreditIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryPayCreditIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryPayCredit represents a PayCredit event raised by the BaseLibrary contract.
type BaseLibraryPayCredit struct {
	Payee  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPayCredit is a free log retrieval operation binding the contract event 0xd8ad32fc5c165753d9da5f44f42ef7daf8002336cab4bdd13b7e58f48b451e3d.
//
// Solidity: e PayCredit(payee address, amount uint256)
func (_BaseLibrary *BaseLibraryFilterer) FilterPayCredit(opts *bind.FilterOpts) (*BaseLibraryPayCreditIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "PayCredit")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryPayCreditIterator{contract: _BaseLibrary.contract, event: "PayCredit", logs: logs, sub: sub}, nil
}

// WatchPayCredit is a free log subscription operation binding the contract event 0xd8ad32fc5c165753d9da5f44f42ef7daf8002336cab4bdd13b7e58f48b451e3d.
//
// Solidity: e PayCredit(payee address, amount uint256)
func (_BaseLibrary *BaseLibraryFilterer) WatchPayCredit(opts *bind.WatchOpts, sink chan<- *BaseLibraryPayCredit) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "PayCredit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryPayCredit)
				if err := _BaseLibrary.contract.UnpackLog(event, "PayCredit", log); err != nil {
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

// BaseLibraryReviewerGroupAddedIterator is returned from FilterReviewerGroupAdded and is used to iterate over the raw logs and unpacked data for ReviewerGroupAdded events raised by the BaseLibrary contract.
type BaseLibraryReviewerGroupAddedIterator struct {
	Event *BaseLibraryReviewerGroupAdded // Event containing the contract specifics and raw log

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
func (it *BaseLibraryReviewerGroupAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryReviewerGroupAdded)
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
		it.Event = new(BaseLibraryReviewerGroupAdded)
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
func (it *BaseLibraryReviewerGroupAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryReviewerGroupAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryReviewerGroupAdded represents a ReviewerGroupAdded event raised by the BaseLibrary contract.
type BaseLibraryReviewerGroupAdded struct {
	Group common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReviewerGroupAdded is a free log retrieval operation binding the contract event 0x1b88a571cc8ac2e87512f05648e79d184f5cc0cbb2889bc487c41f8b9a3202eb.
//
// Solidity: e ReviewerGroupAdded(group address)
func (_BaseLibrary *BaseLibraryFilterer) FilterReviewerGroupAdded(opts *bind.FilterOpts) (*BaseLibraryReviewerGroupAddedIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "ReviewerGroupAdded")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryReviewerGroupAddedIterator{contract: _BaseLibrary.contract, event: "ReviewerGroupAdded", logs: logs, sub: sub}, nil
}

// WatchReviewerGroupAdded is a free log subscription operation binding the contract event 0x1b88a571cc8ac2e87512f05648e79d184f5cc0cbb2889bc487c41f8b9a3202eb.
//
// Solidity: e ReviewerGroupAdded(group address)
func (_BaseLibrary *BaseLibraryFilterer) WatchReviewerGroupAdded(opts *bind.WatchOpts, sink chan<- *BaseLibraryReviewerGroupAdded) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "ReviewerGroupAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryReviewerGroupAdded)
				if err := _BaseLibrary.contract.UnpackLog(event, "ReviewerGroupAdded", log); err != nil {
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

// BaseLibraryUnauthorizedOperationIterator is returned from FilterUnauthorizedOperation and is used to iterate over the raw logs and unpacked data for UnauthorizedOperation events raised by the BaseLibrary contract.
type BaseLibraryUnauthorizedOperationIterator struct {
	Event *BaseLibraryUnauthorizedOperation // Event containing the contract specifics and raw log

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
func (it *BaseLibraryUnauthorizedOperationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryUnauthorizedOperation)
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
		it.Event = new(BaseLibraryUnauthorizedOperation)
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
func (it *BaseLibraryUnauthorizedOperationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryUnauthorizedOperationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryUnauthorizedOperation represents a UnauthorizedOperation event raised by the BaseLibrary contract.
type BaseLibraryUnauthorizedOperation struct {
	OperationCode *big.Int
	Candidate     common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnauthorizedOperation is a free log retrieval operation binding the contract event 0x23de2adc3e22f171f66b3e5a333e17feb9dc30ba9570933bd259cb6c13ef7ab7.
//
// Solidity: e UnauthorizedOperation(operationCode uint256, candidate address)
func (_BaseLibrary *BaseLibraryFilterer) FilterUnauthorizedOperation(opts *bind.FilterOpts) (*BaseLibraryUnauthorizedOperationIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "UnauthorizedOperation")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryUnauthorizedOperationIterator{contract: _BaseLibrary.contract, event: "UnauthorizedOperation", logs: logs, sub: sub}, nil
}

// WatchUnauthorizedOperation is a free log subscription operation binding the contract event 0x23de2adc3e22f171f66b3e5a333e17feb9dc30ba9570933bd259cb6c13ef7ab7.
//
// Solidity: e UnauthorizedOperation(operationCode uint256, candidate address)
func (_BaseLibrary *BaseLibraryFilterer) WatchUnauthorizedOperation(opts *bind.WatchOpts, sink chan<- *BaseLibraryUnauthorizedOperation) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "UnauthorizedOperation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryUnauthorizedOperation)
				if err := _BaseLibrary.contract.UnpackLog(event, "UnauthorizedOperation", log); err != nil {
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

// BaseLibraryUpdateRequestIterator is returned from FilterUpdateRequest and is used to iterate over the raw logs and unpacked data for UpdateRequest events raised by the BaseLibrary contract.
type BaseLibraryUpdateRequestIterator struct {
	Event *BaseLibraryUpdateRequest // Event containing the contract specifics and raw log

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
func (it *BaseLibraryUpdateRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryUpdateRequest)
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
		it.Event = new(BaseLibraryUpdateRequest)
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
func (it *BaseLibraryUpdateRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryUpdateRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryUpdateRequest represents a UpdateRequest event raised by the BaseLibrary contract.
type BaseLibraryUpdateRequest struct {
	ObjectHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateRequest is a free log retrieval operation binding the contract event 0x1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef19.
//
// Solidity: e UpdateRequest(objectHash bytes32)
func (_BaseLibrary *BaseLibraryFilterer) FilterUpdateRequest(opts *bind.FilterOpts) (*BaseLibraryUpdateRequestIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryUpdateRequestIterator{contract: _BaseLibrary.contract, event: "UpdateRequest", logs: logs, sub: sub}, nil
}

// WatchUpdateRequest is a free log subscription operation binding the contract event 0x1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef19.
//
// Solidity: e UpdateRequest(objectHash bytes32)
func (_BaseLibrary *BaseLibraryFilterer) WatchUpdateRequest(opts *bind.WatchOpts, sink chan<- *BaseLibraryUpdateRequest) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryUpdateRequest)
				if err := _BaseLibrary.contract.UnpackLog(event, "UpdateRequest", log); err != nil {
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

// ContentABI is the input ABI used to generate the binding from.
const ContentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"address[]\"}],\"name\":\"runAccess\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proposed_status_code\",\"type\":\"int256\"}],\"name\":\"runStatusChange\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"int256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"runFinalize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"address[]\"}],\"name\":\"runAccessCharge\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"s\",\"type\":\"string\"}],\"name\":\"DbgString\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"a\",\"type\":\"address\"}],\"name\":\"DbgAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"u\",\"type\":\"uint256\"}],\"name\":\"DbgUint256\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"u\",\"type\":\"uint256\"}],\"name\":\"DbgUint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"calculateAccessCharge\",\"type\":\"int256\"}],\"name\":\"RunAccessCharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"RunAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"RunFinalize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"proposedStatusCode\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"returnStatusCode\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"licenseFeeToBePaid\",\"type\":\"int256\"}],\"name\":\"RunStatusChange\",\"type\":\"event\"}]"

// ContentBin is the compiled bytecode used for deploying new contracts.
const ContentBin = `0x606060405260008054600160a060020a033216600160a060020a031991821681179092556001805490911690911790556104338061003e6000396000f3006060604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f811461009a578063123e0e80146100c95780633513a8051461016d57806341c0e1b5146101905780636b2d1324146101a35780636d2e4b1b146101c25780638da5cb5b146101e1578063f2fde38b146101f4578063f8117ca214610213575b005b34156100a557600080fd5b6100ad61029f565b604051600160a060020a03909116815260200160405180910390f35b61015b60048035906024803560ff1691906064906044359081019083013580602080820201604051908101604052809392919081815260200183836020028082843782019150505050505091908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437509496506102ae95505050505050565b60405190815260200160405180910390f35b6101786004356102b8565b60405191825260208201526040908101905180910390f35b341561019b57600080fd5b6100986102bf565b6101ae6004356102e8565b604051901515815260200160405180910390f35b34156101cd57600080fd5b610098600160a060020a03600435166102ee565b34156101ec57600080fd5b6100ad61038f565b34156101ff57600080fd5b610098600160a060020a036004351661039e565b61015b6004803560ff16906044602480359081019083013580602080820201604051908101604052809392919081815260200183836020028082843782019150505050505091908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437509496506103fd95505050505050565b600054600160a060020a031681565b6000949350505050565b9060001990565b60015432600160a060020a039081169116146102da57600080fd5b600154600160a060020a0316ff5b50600190565b60005432600160a060020a0390811691161461030957600080fd5b600160a060020a038116151561031e57600080fd5b600054600154600160a060020a0390811691161415610360576001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031681565b60015432600160a060020a039081169116146103b957600080fd5b600160a060020a03811615156103ce57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60001993925050505600a165627a7a7230582010c77ba01a3fd499336a9d70a968d85f822ba4667bb8fb7bd5a8a91891a39dd00029`

// DeployContent deploys a new Ethereum contract, binding an instance of Content to it.
func DeployContent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Content, error) {
	parsed, err := abi.JSON(strings.NewReader(ContentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Content{ContentCaller: ContentCaller{contract: contract}, ContentTransactor: ContentTransactor{contract: contract}, ContentFilterer: ContentFilterer{contract: contract}}, nil
}

// Content is an auto generated Go binding around an Ethereum contract.
type Content struct {
	ContentCaller     // Read-only binding to the contract
	ContentTransactor // Write-only binding to the contract
	ContentFilterer   // Log filterer for contract events
}

// ContentCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContentSession struct {
	Contract     *Content          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContentCallerSession struct {
	Contract *ContentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContentTransactorSession struct {
	Contract     *ContentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContentRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContentRaw struct {
	Contract *Content // Generic contract binding to access the raw methods on
}

// ContentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContentCallerRaw struct {
	Contract *ContentCaller // Generic read-only contract binding to access the raw methods on
}

// ContentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContentTransactorRaw struct {
	Contract *ContentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContent creates a new instance of Content, bound to a specific deployed contract.
func NewContent(address common.Address, backend bind.ContractBackend) (*Content, error) {
	contract, err := bindContent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Content{ContentCaller: ContentCaller{contract: contract}, ContentTransactor: ContentTransactor{contract: contract}, ContentFilterer: ContentFilterer{contract: contract}}, nil
}

// NewContentCaller creates a new read-only instance of Content, bound to a specific deployed contract.
func NewContentCaller(address common.Address, caller bind.ContractCaller) (*ContentCaller, error) {
	contract, err := bindContent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContentCaller{contract: contract}, nil
}

// NewContentTransactor creates a new write-only instance of Content, bound to a specific deployed contract.
func NewContentTransactor(address common.Address, transactor bind.ContractTransactor) (*ContentTransactor, error) {
	contract, err := bindContent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContentTransactor{contract: contract}, nil
}

// NewContentFilterer creates a new log filterer instance of Content, bound to a specific deployed contract.
func NewContentFilterer(address common.Address, filterer bind.ContractFilterer) (*ContentFilterer, error) {
	contract, err := bindContent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContentFilterer{contract: contract}, nil
}

// bindContent binds a generic wrapper to an already deployed contract.
func bindContent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Content *ContentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Content.Contract.ContentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Content *ContentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Content.Contract.ContentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Content *ContentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Content.Contract.ContentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Content *ContentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Content.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Content *ContentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Content.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Content *ContentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Content.Contract.contract.Transact(opts, method, params...)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Content *ContentCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Content.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Content *ContentSession) Creator() (common.Address, error) {
	return _Content.Contract.Creator(&_Content.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Content *ContentCallerSession) Creator() (common.Address, error) {
	return _Content.Contract.Creator(&_Content.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Content *ContentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Content.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Content *ContentSession) Owner() (common.Address, error) {
	return _Content.Contract.Owner(&_Content.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Content *ContentCallerSession) Owner() (common.Address, error) {
	return _Content.Contract.Owner(&_Content.CallOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Content *ContentTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Content *ContentSession) Kill() (*types.Transaction, error) {
	return _Content.Contract.Kill(&_Content.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Content *ContentTransactorSession) Kill() (*types.Transaction, error) {
	return _Content.Contract.Kill(&_Content.TransactOpts)
}

// RunAccess is a paid mutator transaction binding the contract method 0x123e0e80.
//
// Solidity: function runAccess( uint256,  uint8,  bytes32[],  address[]) returns(uint256)
func (_Content *ContentTransactor) RunAccess(opts *bind.TransactOpts, arg0 *big.Int, arg1 uint8, arg2 [][32]byte, arg3 []common.Address) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "runAccess", arg0, arg1, arg2, arg3)
}

// RunAccess is a paid mutator transaction binding the contract method 0x123e0e80.
//
// Solidity: function runAccess( uint256,  uint8,  bytes32[],  address[]) returns(uint256)
func (_Content *ContentSession) RunAccess(arg0 *big.Int, arg1 uint8, arg2 [][32]byte, arg3 []common.Address) (*types.Transaction, error) {
	return _Content.Contract.RunAccess(&_Content.TransactOpts, arg0, arg1, arg2, arg3)
}

// RunAccess is a paid mutator transaction binding the contract method 0x123e0e80.
//
// Solidity: function runAccess( uint256,  uint8,  bytes32[],  address[]) returns(uint256)
func (_Content *ContentTransactorSession) RunAccess(arg0 *big.Int, arg1 uint8, arg2 [][32]byte, arg3 []common.Address) (*types.Transaction, error) {
	return _Content.Contract.RunAccess(&_Content.TransactOpts, arg0, arg1, arg2, arg3)
}

// RunAccessCharge is a paid mutator transaction binding the contract method 0xf8117ca2.
//
// Solidity: function runAccessCharge( uint8,  bytes32[],  address[]) returns(int256)
func (_Content *ContentTransactor) RunAccessCharge(opts *bind.TransactOpts, arg0 uint8, arg1 [][32]byte, arg2 []common.Address) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "runAccessCharge", arg0, arg1, arg2)
}

// RunAccessCharge is a paid mutator transaction binding the contract method 0xf8117ca2.
//
// Solidity: function runAccessCharge( uint8,  bytes32[],  address[]) returns(int256)
func (_Content *ContentSession) RunAccessCharge(arg0 uint8, arg1 [][32]byte, arg2 []common.Address) (*types.Transaction, error) {
	return _Content.Contract.RunAccessCharge(&_Content.TransactOpts, arg0, arg1, arg2)
}

// RunAccessCharge is a paid mutator transaction binding the contract method 0xf8117ca2.
//
// Solidity: function runAccessCharge( uint8,  bytes32[],  address[]) returns(int256)
func (_Content *ContentTransactorSession) RunAccessCharge(arg0 uint8, arg1 [][32]byte, arg2 []common.Address) (*types.Transaction, error) {
	return _Content.Contract.RunAccessCharge(&_Content.TransactOpts, arg0, arg1, arg2)
}

// RunFinalize is a paid mutator transaction binding the contract method 0x6b2d1324.
//
// Solidity: function runFinalize( uint256) returns(bool)
func (_Content *ContentTransactor) RunFinalize(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "runFinalize", arg0)
}

// RunFinalize is a paid mutator transaction binding the contract method 0x6b2d1324.
//
// Solidity: function runFinalize( uint256) returns(bool)
func (_Content *ContentSession) RunFinalize(arg0 *big.Int) (*types.Transaction, error) {
	return _Content.Contract.RunFinalize(&_Content.TransactOpts, arg0)
}

// RunFinalize is a paid mutator transaction binding the contract method 0x6b2d1324.
//
// Solidity: function runFinalize( uint256) returns(bool)
func (_Content *ContentTransactorSession) RunFinalize(arg0 *big.Int) (*types.Transaction, error) {
	return _Content.Contract.RunFinalize(&_Content.TransactOpts, arg0)
}

// RunStatusChange is a paid mutator transaction binding the contract method 0x3513a805.
//
// Solidity: function runStatusChange(proposed_status_code int256) returns(int256, int256)
func (_Content *ContentTransactor) RunStatusChange(opts *bind.TransactOpts, proposed_status_code *big.Int) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "runStatusChange", proposed_status_code)
}

// RunStatusChange is a paid mutator transaction binding the contract method 0x3513a805.
//
// Solidity: function runStatusChange(proposed_status_code int256) returns(int256, int256)
func (_Content *ContentSession) RunStatusChange(proposed_status_code *big.Int) (*types.Transaction, error) {
	return _Content.Contract.RunStatusChange(&_Content.TransactOpts, proposed_status_code)
}

// RunStatusChange is a paid mutator transaction binding the contract method 0x3513a805.
//
// Solidity: function runStatusChange(proposed_status_code int256) returns(int256, int256)
func (_Content *ContentTransactorSession) RunStatusChange(proposed_status_code *big.Int) (*types.Transaction, error) {
	return _Content.Contract.RunStatusChange(&_Content.TransactOpts, proposed_status_code)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_Content *ContentTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_Content *ContentSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _Content.Contract.TransferCreatorship(&_Content.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_Content *ContentTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _Content.Contract.TransferCreatorship(&_Content.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Content *ContentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Content *ContentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Content.Contract.TransferOwnership(&_Content.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Content *ContentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Content.Contract.TransferOwnership(&_Content.TransactOpts, newOwner)
}

// ContentDbgAddressIterator is returned from FilterDbgAddress and is used to iterate over the raw logs and unpacked data for DbgAddress events raised by the Content contract.
type ContentDbgAddressIterator struct {
	Event *ContentDbgAddress // Event containing the contract specifics and raw log

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
func (it *ContentDbgAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentDbgAddress)
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
		it.Event = new(ContentDbgAddress)
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
func (it *ContentDbgAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentDbgAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentDbgAddress represents a DbgAddress event raised by the Content contract.
type ContentDbgAddress struct {
	A   common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDbgAddress is a free log retrieval operation binding the contract event 0x474589881ea4b1937d02648256f59b0dec5c28f4d1a8be6efcb50aa51aaf69de.
//
// Solidity: e DbgAddress(a address)
func (_Content *ContentFilterer) FilterDbgAddress(opts *bind.FilterOpts) (*ContentDbgAddressIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "DbgAddress")
	if err != nil {
		return nil, err
	}
	return &ContentDbgAddressIterator{contract: _Content.contract, event: "DbgAddress", logs: logs, sub: sub}, nil
}

// WatchDbgAddress is a free log subscription operation binding the contract event 0x474589881ea4b1937d02648256f59b0dec5c28f4d1a8be6efcb50aa51aaf69de.
//
// Solidity: e DbgAddress(a address)
func (_Content *ContentFilterer) WatchDbgAddress(opts *bind.WatchOpts, sink chan<- *ContentDbgAddress) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "DbgAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentDbgAddress)
				if err := _Content.contract.UnpackLog(event, "DbgAddress", log); err != nil {
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

// ContentDbgStringIterator is returned from FilterDbgString and is used to iterate over the raw logs and unpacked data for DbgString events raised by the Content contract.
type ContentDbgStringIterator struct {
	Event *ContentDbgString // Event containing the contract specifics and raw log

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
func (it *ContentDbgStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentDbgString)
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
		it.Event = new(ContentDbgString)
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
func (it *ContentDbgStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentDbgStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentDbgString represents a DbgString event raised by the Content contract.
type ContentDbgString struct {
	S   string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDbgString is a free log retrieval operation binding the contract event 0x415302daeab7bd01c23adb4dee059bc666b313134024f09781aa6d73ff6ec042.
//
// Solidity: e DbgString(s string)
func (_Content *ContentFilterer) FilterDbgString(opts *bind.FilterOpts) (*ContentDbgStringIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "DbgString")
	if err != nil {
		return nil, err
	}
	return &ContentDbgStringIterator{contract: _Content.contract, event: "DbgString", logs: logs, sub: sub}, nil
}

// WatchDbgString is a free log subscription operation binding the contract event 0x415302daeab7bd01c23adb4dee059bc666b313134024f09781aa6d73ff6ec042.
//
// Solidity: e DbgString(s string)
func (_Content *ContentFilterer) WatchDbgString(opts *bind.WatchOpts, sink chan<- *ContentDbgString) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "DbgString")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentDbgString)
				if err := _Content.contract.UnpackLog(event, "DbgString", log); err != nil {
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

// ContentDbgUintIterator is returned from FilterDbgUint and is used to iterate over the raw logs and unpacked data for DbgUint events raised by the Content contract.
type ContentDbgUintIterator struct {
	Event *ContentDbgUint // Event containing the contract specifics and raw log

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
func (it *ContentDbgUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentDbgUint)
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
		it.Event = new(ContentDbgUint)
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
func (it *ContentDbgUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentDbgUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentDbgUint represents a DbgUint event raised by the Content contract.
type ContentDbgUint struct {
	U   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDbgUint is a free log retrieval operation binding the contract event 0xf209d49352c8b16a38abaf19dd720dfc5a398b0209f3d48a673a52234e3c6aac.
//
// Solidity: e DbgUint(u uint256)
func (_Content *ContentFilterer) FilterDbgUint(opts *bind.FilterOpts) (*ContentDbgUintIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "DbgUint")
	if err != nil {
		return nil, err
	}
	return &ContentDbgUintIterator{contract: _Content.contract, event: "DbgUint", logs: logs, sub: sub}, nil
}

// WatchDbgUint is a free log subscription operation binding the contract event 0xf209d49352c8b16a38abaf19dd720dfc5a398b0209f3d48a673a52234e3c6aac.
//
// Solidity: e DbgUint(u uint256)
func (_Content *ContentFilterer) WatchDbgUint(opts *bind.WatchOpts, sink chan<- *ContentDbgUint) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "DbgUint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentDbgUint)
				if err := _Content.contract.UnpackLog(event, "DbgUint", log); err != nil {
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

// ContentDbgUint256Iterator is returned from FilterDbgUint256 and is used to iterate over the raw logs and unpacked data for DbgUint256 events raised by the Content contract.
type ContentDbgUint256Iterator struct {
	Event *ContentDbgUint256 // Event containing the contract specifics and raw log

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
func (it *ContentDbgUint256Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentDbgUint256)
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
		it.Event = new(ContentDbgUint256)
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
func (it *ContentDbgUint256Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentDbgUint256Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentDbgUint256 represents a DbgUint256 event raised by the Content contract.
type ContentDbgUint256 struct {
	U   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDbgUint256 is a free log retrieval operation binding the contract event 0x6b0cf6507e95b7606ee5ef593c4219bc4056b30825b0eecd12e5e7022126f23f.
//
// Solidity: e DbgUint256(u uint256)
func (_Content *ContentFilterer) FilterDbgUint256(opts *bind.FilterOpts) (*ContentDbgUint256Iterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "DbgUint256")
	if err != nil {
		return nil, err
	}
	return &ContentDbgUint256Iterator{contract: _Content.contract, event: "DbgUint256", logs: logs, sub: sub}, nil
}

// WatchDbgUint256 is a free log subscription operation binding the contract event 0x6b0cf6507e95b7606ee5ef593c4219bc4056b30825b0eecd12e5e7022126f23f.
//
// Solidity: e DbgUint256(u uint256)
func (_Content *ContentFilterer) WatchDbgUint256(opts *bind.WatchOpts, sink chan<- *ContentDbgUint256) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "DbgUint256")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentDbgUint256)
				if err := _Content.contract.UnpackLog(event, "DbgUint256", log); err != nil {
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

// ContentRunAccessIterator is returned from FilterRunAccess and is used to iterate over the raw logs and unpacked data for RunAccess events raised by the Content contract.
type ContentRunAccessIterator struct {
	Event *ContentRunAccess // Event containing the contract specifics and raw log

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
func (it *ContentRunAccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentRunAccess)
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
		it.Event = new(ContentRunAccess)
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
func (it *ContentRunAccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentRunAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentRunAccess represents a RunAccess event raised by the Content contract.
type ContentRunAccess struct {
	RequestID *big.Int
	Result    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRunAccess is a free log retrieval operation binding the contract event 0x3e68dc35f88d76818f276322c37f5021ee00e232fe0d27a93c02801aec4d9c58.
//
// Solidity: e RunAccess(requestID uint256, result uint256)
func (_Content *ContentFilterer) FilterRunAccess(opts *bind.FilterOpts) (*ContentRunAccessIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "RunAccess")
	if err != nil {
		return nil, err
	}
	return &ContentRunAccessIterator{contract: _Content.contract, event: "RunAccess", logs: logs, sub: sub}, nil
}

// WatchRunAccess is a free log subscription operation binding the contract event 0x3e68dc35f88d76818f276322c37f5021ee00e232fe0d27a93c02801aec4d9c58.
//
// Solidity: e RunAccess(requestID uint256, result uint256)
func (_Content *ContentFilterer) WatchRunAccess(opts *bind.WatchOpts, sink chan<- *ContentRunAccess) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "RunAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentRunAccess)
				if err := _Content.contract.UnpackLog(event, "RunAccess", log); err != nil {
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

// ContentRunAccessChargeIterator is returned from FilterRunAccessCharge and is used to iterate over the raw logs and unpacked data for RunAccessCharge events raised by the Content contract.
type ContentRunAccessChargeIterator struct {
	Event *ContentRunAccessCharge // Event containing the contract specifics and raw log

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
func (it *ContentRunAccessChargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentRunAccessCharge)
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
		it.Event = new(ContentRunAccessCharge)
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
func (it *ContentRunAccessChargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentRunAccessChargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentRunAccessCharge represents a RunAccessCharge event raised by the Content contract.
type ContentRunAccessCharge struct {
	Level                 uint8
	CalculateAccessCharge *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRunAccessCharge is a free log retrieval operation binding the contract event 0xffadad18ab3777a19f664019a6261b011ab9405749e01a45950d44fb9360b385.
//
// Solidity: e RunAccessCharge(level uint8, calculateAccessCharge int256)
func (_Content *ContentFilterer) FilterRunAccessCharge(opts *bind.FilterOpts) (*ContentRunAccessChargeIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "RunAccessCharge")
	if err != nil {
		return nil, err
	}
	return &ContentRunAccessChargeIterator{contract: _Content.contract, event: "RunAccessCharge", logs: logs, sub: sub}, nil
}

// WatchRunAccessCharge is a free log subscription operation binding the contract event 0xffadad18ab3777a19f664019a6261b011ab9405749e01a45950d44fb9360b385.
//
// Solidity: e RunAccessCharge(level uint8, calculateAccessCharge int256)
func (_Content *ContentFilterer) WatchRunAccessCharge(opts *bind.WatchOpts, sink chan<- *ContentRunAccessCharge) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "RunAccessCharge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentRunAccessCharge)
				if err := _Content.contract.UnpackLog(event, "RunAccessCharge", log); err != nil {
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

// ContentRunFinalizeIterator is returned from FilterRunFinalize and is used to iterate over the raw logs and unpacked data for RunFinalize events raised by the Content contract.
type ContentRunFinalizeIterator struct {
	Event *ContentRunFinalize // Event containing the contract specifics and raw log

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
func (it *ContentRunFinalizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentRunFinalize)
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
		it.Event = new(ContentRunFinalize)
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
func (it *ContentRunFinalizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentRunFinalizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentRunFinalize represents a RunFinalize event raised by the Content contract.
type ContentRunFinalize struct {
	RequestID *big.Int
	Result    bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRunFinalize is a free log retrieval operation binding the contract event 0xab4a2c9cfb6dbb6e74cefc636aceee603aeeb2d2a502c99eee0386d4d9969120.
//
// Solidity: e RunFinalize(requestID uint256, result bool)
func (_Content *ContentFilterer) FilterRunFinalize(opts *bind.FilterOpts) (*ContentRunFinalizeIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "RunFinalize")
	if err != nil {
		return nil, err
	}
	return &ContentRunFinalizeIterator{contract: _Content.contract, event: "RunFinalize", logs: logs, sub: sub}, nil
}

// WatchRunFinalize is a free log subscription operation binding the contract event 0xab4a2c9cfb6dbb6e74cefc636aceee603aeeb2d2a502c99eee0386d4d9969120.
//
// Solidity: e RunFinalize(requestID uint256, result bool)
func (_Content *ContentFilterer) WatchRunFinalize(opts *bind.WatchOpts, sink chan<- *ContentRunFinalize) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "RunFinalize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentRunFinalize)
				if err := _Content.contract.UnpackLog(event, "RunFinalize", log); err != nil {
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

// ContentRunStatusChangeIterator is returned from FilterRunStatusChange and is used to iterate over the raw logs and unpacked data for RunStatusChange events raised by the Content contract.
type ContentRunStatusChangeIterator struct {
	Event *ContentRunStatusChange // Event containing the contract specifics and raw log

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
func (it *ContentRunStatusChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentRunStatusChange)
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
		it.Event = new(ContentRunStatusChange)
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
func (it *ContentRunStatusChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentRunStatusChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentRunStatusChange represents a RunStatusChange event raised by the Content contract.
type ContentRunStatusChange struct {
	ProposedStatusCode *big.Int
	ReturnStatusCode   *big.Int
	LicenseFeeToBePaid *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRunStatusChange is a free log retrieval operation binding the contract event 0x33a32333546026b7e6cc92c04a98029e9fcc452b9796d85be9178531adfb7126.
//
// Solidity: e RunStatusChange(proposedStatusCode int256, returnStatusCode int256, licenseFeeToBePaid int256)
func (_Content *ContentFilterer) FilterRunStatusChange(opts *bind.FilterOpts) (*ContentRunStatusChangeIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "RunStatusChange")
	if err != nil {
		return nil, err
	}
	return &ContentRunStatusChangeIterator{contract: _Content.contract, event: "RunStatusChange", logs: logs, sub: sub}, nil
}

// WatchRunStatusChange is a free log subscription operation binding the contract event 0x33a32333546026b7e6cc92c04a98029e9fcc452b9796d85be9178531adfb7126.
//
// Solidity: e RunStatusChange(proposedStatusCode int256, returnStatusCode int256, licenseFeeToBePaid int256)
func (_Content *ContentFilterer) WatchRunStatusChange(opts *bind.WatchOpts, sink chan<- *ContentRunStatusChange) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "RunStatusChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentRunStatusChange)
				if err := _Content.contract.UnpackLog(event, "RunStatusChange", log); err != nil {
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

// EditableABI is the input ABI used to generate the binding from.
const EditableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"objectHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"object_hash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"Commit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"UpdateRequest\",\"type\":\"event\"}]"

// EditableBin is the compiled bytecode used for deploying new contracts.
const EditableBin = `0x606060405260008054600160a060020a033216600160a060020a031991821681179092556001805490911690911790556103918061003e6000396000f30060606040526004361061008d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f811461008f57806341c0e1b5146100be5780636d2e4b1b146100d15780638da5cb5b146100f0578063c287e0ed14610103578063e02dd9c214610116578063f14fcbc81461013b578063f2fde38b14610151575b005b341561009a57600080fd5b6100a2610170565b604051600160a060020a03909116815260200160405180910390f35b34156100c957600080fd5b61008d61017f565b34156100dc57600080fd5b61008d600160a060020a03600435166101a8565b34156100fb57600080fd5b6100a2610249565b341561010e57600080fd5b61008d610258565b341561012157600080fd5b6101296102aa565b60405190815260200160405180910390f35b341561014657600080fd5b61008d6004356102b0565b341561015c57600080fd5b61008d600160a060020a0360043516610306565b600054600160a060020a031681565b60015432600160a060020a0390811691161461019a57600080fd5b600154600160a060020a0316ff5b60005432600160a060020a039081169116146101c357600080fd5b600160a060020a03811615156101d857600080fd5b600054600154600160a060020a039081169116141561021a576001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031681565b60015432600160a060020a0390811691161461027357600080fd5b7f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef1960025460405190815260200160405180910390a1565b60025481565b60015432600160a060020a039081169116146102cb57600080fd5b60028190557f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff8160405190815260200160405180910390a150565b60015432600160a060020a0390811691161461032157600080fd5b600160a060020a038116151561033657600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058200f88fb35723517385ab38a72990e4c72d0aaaf08334858fca27fa1e2daee00ed0029`

// DeployEditable deploys a new Ethereum contract, binding an instance of Editable to it.
func DeployEditable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Editable, error) {
	parsed, err := abi.JSON(strings.NewReader(EditableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EditableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Editable{EditableCaller: EditableCaller{contract: contract}, EditableTransactor: EditableTransactor{contract: contract}, EditableFilterer: EditableFilterer{contract: contract}}, nil
}

// Editable is an auto generated Go binding around an Ethereum contract.
type Editable struct {
	EditableCaller     // Read-only binding to the contract
	EditableTransactor // Write-only binding to the contract
	EditableFilterer   // Log filterer for contract events
}

// EditableCaller is an auto generated read-only Go binding around an Ethereum contract.
type EditableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EditableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EditableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EditableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EditableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EditableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EditableSession struct {
	Contract     *Editable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EditableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EditableCallerSession struct {
	Contract *EditableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EditableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EditableTransactorSession struct {
	Contract     *EditableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EditableRaw is an auto generated low-level Go binding around an Ethereum contract.
type EditableRaw struct {
	Contract *Editable // Generic contract binding to access the raw methods on
}

// EditableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EditableCallerRaw struct {
	Contract *EditableCaller // Generic read-only contract binding to access the raw methods on
}

// EditableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EditableTransactorRaw struct {
	Contract *EditableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEditable creates a new instance of Editable, bound to a specific deployed contract.
func NewEditable(address common.Address, backend bind.ContractBackend) (*Editable, error) {
	contract, err := bindEditable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Editable{EditableCaller: EditableCaller{contract: contract}, EditableTransactor: EditableTransactor{contract: contract}, EditableFilterer: EditableFilterer{contract: contract}}, nil
}

// NewEditableCaller creates a new read-only instance of Editable, bound to a specific deployed contract.
func NewEditableCaller(address common.Address, caller bind.ContractCaller) (*EditableCaller, error) {
	contract, err := bindEditable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EditableCaller{contract: contract}, nil
}

// NewEditableTransactor creates a new write-only instance of Editable, bound to a specific deployed contract.
func NewEditableTransactor(address common.Address, transactor bind.ContractTransactor) (*EditableTransactor, error) {
	contract, err := bindEditable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EditableTransactor{contract: contract}, nil
}

// NewEditableFilterer creates a new log filterer instance of Editable, bound to a specific deployed contract.
func NewEditableFilterer(address common.Address, filterer bind.ContractFilterer) (*EditableFilterer, error) {
	contract, err := bindEditable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EditableFilterer{contract: contract}, nil
}

// bindEditable binds a generic wrapper to an already deployed contract.
func bindEditable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EditableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Editable *EditableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Editable.Contract.EditableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Editable *EditableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Editable.Contract.EditableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Editable *EditableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Editable.Contract.EditableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Editable *EditableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Editable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Editable *EditableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Editable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Editable *EditableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Editable.Contract.contract.Transact(opts, method, params...)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Editable *EditableCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Editable.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Editable *EditableSession) Creator() (common.Address, error) {
	return _Editable.Contract.Creator(&_Editable.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Editable *EditableCallerSession) Creator() (common.Address, error) {
	return _Editable.Contract.Creator(&_Editable.CallOpts)
}

// ObjectHash is a free data retrieval call binding the contract method 0xe02dd9c2.
//
// Solidity: function objectHash() constant returns(bytes32)
func (_Editable *EditableCaller) ObjectHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Editable.contract.Call(opts, out, "objectHash")
	return *ret0, err
}

// ObjectHash is a free data retrieval call binding the contract method 0xe02dd9c2.
//
// Solidity: function objectHash() constant returns(bytes32)
func (_Editable *EditableSession) ObjectHash() ([32]byte, error) {
	return _Editable.Contract.ObjectHash(&_Editable.CallOpts)
}

// ObjectHash is a free data retrieval call binding the contract method 0xe02dd9c2.
//
// Solidity: function objectHash() constant returns(bytes32)
func (_Editable *EditableCallerSession) ObjectHash() ([32]byte, error) {
	return _Editable.Contract.ObjectHash(&_Editable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Editable *EditableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Editable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Editable *EditableSession) Owner() (common.Address, error) {
	return _Editable.Contract.Owner(&_Editable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Editable *EditableCallerSession) Owner() (common.Address, error) {
	return _Editable.Contract.Owner(&_Editable.CallOpts)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(object_hash bytes32) returns()
func (_Editable *EditableTransactor) Commit(opts *bind.TransactOpts, object_hash [32]byte) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "commit", object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(object_hash bytes32) returns()
func (_Editable *EditableSession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _Editable.Contract.Commit(&_Editable.TransactOpts, object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(object_hash bytes32) returns()
func (_Editable *EditableTransactorSession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _Editable.Contract.Commit(&_Editable.TransactOpts, object_hash)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Editable *EditableTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Editable *EditableSession) Kill() (*types.Transaction, error) {
	return _Editable.Contract.Kill(&_Editable.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Editable *EditableTransactorSession) Kill() (*types.Transaction, error) {
	return _Editable.Contract.Kill(&_Editable.TransactOpts)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_Editable *EditableTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_Editable *EditableSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _Editable.Contract.TransferCreatorship(&_Editable.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_Editable *EditableTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _Editable.Contract.TransferCreatorship(&_Editable.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Editable *EditableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Editable *EditableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Editable.Contract.TransferOwnership(&_Editable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Editable *EditableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Editable.Contract.TransferOwnership(&_Editable.TransactOpts, newOwner)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xc287e0ed.
//
// Solidity: function updateRequest() returns()
func (_Editable *EditableTransactor) UpdateRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "updateRequest")
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xc287e0ed.
//
// Solidity: function updateRequest() returns()
func (_Editable *EditableSession) UpdateRequest() (*types.Transaction, error) {
	return _Editable.Contract.UpdateRequest(&_Editable.TransactOpts)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xc287e0ed.
//
// Solidity: function updateRequest() returns()
func (_Editable *EditableTransactorSession) UpdateRequest() (*types.Transaction, error) {
	return _Editable.Contract.UpdateRequest(&_Editable.TransactOpts)
}

// EditableCommitIterator is returned from FilterCommit and is used to iterate over the raw logs and unpacked data for Commit events raised by the Editable contract.
type EditableCommitIterator struct {
	Event *EditableCommit // Event containing the contract specifics and raw log

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
func (it *EditableCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EditableCommit)
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
		it.Event = new(EditableCommit)
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
func (it *EditableCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EditableCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EditableCommit represents a Commit event raised by the Editable contract.
type EditableCommit struct {
	ObjectHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommit is a free log retrieval operation binding the contract event 0x9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff.
//
// Solidity: e Commit(objectHash bytes32)
func (_Editable *EditableFilterer) FilterCommit(opts *bind.FilterOpts) (*EditableCommitIterator, error) {

	logs, sub, err := _Editable.contract.FilterLogs(opts, "Commit")
	if err != nil {
		return nil, err
	}
	return &EditableCommitIterator{contract: _Editable.contract, event: "Commit", logs: logs, sub: sub}, nil
}

// WatchCommit is a free log subscription operation binding the contract event 0x9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff.
//
// Solidity: e Commit(objectHash bytes32)
func (_Editable *EditableFilterer) WatchCommit(opts *bind.WatchOpts, sink chan<- *EditableCommit) (event.Subscription, error) {

	logs, sub, err := _Editable.contract.WatchLogs(opts, "Commit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EditableCommit)
				if err := _Editable.contract.UnpackLog(event, "Commit", log); err != nil {
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

// EditableUpdateRequestIterator is returned from FilterUpdateRequest and is used to iterate over the raw logs and unpacked data for UpdateRequest events raised by the Editable contract.
type EditableUpdateRequestIterator struct {
	Event *EditableUpdateRequest // Event containing the contract specifics and raw log

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
func (it *EditableUpdateRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EditableUpdateRequest)
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
		it.Event = new(EditableUpdateRequest)
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
func (it *EditableUpdateRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EditableUpdateRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EditableUpdateRequest represents a UpdateRequest event raised by the Editable contract.
type EditableUpdateRequest struct {
	ObjectHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateRequest is a free log retrieval operation binding the contract event 0x1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef19.
//
// Solidity: e UpdateRequest(objectHash bytes32)
func (_Editable *EditableFilterer) FilterUpdateRequest(opts *bind.FilterOpts) (*EditableUpdateRequestIterator, error) {

	logs, sub, err := _Editable.contract.FilterLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return &EditableUpdateRequestIterator{contract: _Editable.contract, event: "UpdateRequest", logs: logs, sub: sub}, nil
}

// WatchUpdateRequest is a free log subscription operation binding the contract event 0x1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef19.
//
// Solidity: e UpdateRequest(objectHash bytes32)
func (_Editable *EditableFilterer) WatchUpdateRequest(opts *bind.WatchOpts, sink chan<- *EditableUpdateRequest) (event.Subscription, error) {

	logs, sub, err := _Editable.contract.WatchLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EditableUpdateRequest)
				if err := _Editable.contract.UnpackLog(event, "UpdateRequest", log); err != nil {
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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033216600160a060020a03199182168117909255600180549091169091179055610274806100496000396000f30060606040526004361061006c5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f811461006e57806341c0e1b51461009d5780636d2e4b1b146100b05780638da5cb5b146100cf578063f2fde38b146100e2575b005b341561007957600080fd5b610081610101565b604051600160a060020a03909116815260200160405180910390f35b34156100a857600080fd5b61006c610110565b34156100bb57600080fd5b61006c600160a060020a0360043516610139565b34156100da57600080fd5b6100816101da565b34156100ed57600080fd5b61006c600160a060020a03600435166101e9565b600054600160a060020a031681565b60015432600160a060020a0390811691161461012b57600080fd5b600154600160a060020a0316ff5b60005432600160a060020a0390811691161461015457600080fd5b600160a060020a038116151561016957600080fd5b600054600154600160a060020a03908116911614156101ab576001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031681565b60015432600160a060020a0390811691161461020457600080fd5b600160a060020a038116151561021957600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058209013bfd3f25cc7f9eeee437fef883eb8e586fbc49f02b4b7ea63bd533dbf071a0029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Ownable *OwnableCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Ownable *OwnableSession) Creator() (common.Address, error) {
	return _Ownable.Contract.Creator(&_Ownable.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Ownable *OwnableCallerSession) Creator() (common.Address, error) {
	return _Ownable.Contract.Creator(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Ownable *OwnableTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Ownable *OwnableSession) Kill() (*types.Transaction, error) {
	return _Ownable.Contract.Kill(&_Ownable.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Ownable *OwnableTransactorSession) Kill() (*types.Transaction, error) {
	return _Ownable.Contract.Kill(&_Ownable.TransactOpts)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_Ownable *OwnableTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_Ownable *OwnableSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferCreatorship(&_Ownable.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(newCreator address) returns()
func (_Ownable *OwnableTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferCreatorship(&_Ownable.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}
