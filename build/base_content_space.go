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
const AccessibleABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"accessRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AccessRequest\",\"type\":\"event\"}]"

// AccessibleBin is the compiled bytecode used for deploying new contracts.
const AccessibleBin = `0x6080604052348015600f57600080fd5b5060c68061001e6000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663f155188781146043575b600080fd5b348015604e57600080fd5b5060556069565b604080519115158252519081900360200190f35b6040516000907fed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88908290a1506001905600a165627a7a72305820502aa6075371a347b26ea6f9bbd1637bc548b3ec00afd0db61334295f83141a00029`

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
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAccessRequest is a free log retrieval operation binding the contract event 0xed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88.
//
// Solidity: event AccessRequest()
func (_Accessible *AccessibleFilterer) FilterAccessRequest(opts *bind.FilterOpts) (*AccessibleAccessRequestIterator, error) {

	logs, sub, err := _Accessible.contract.FilterLogs(opts, "AccessRequest")
	if err != nil {
		return nil, err
	}
	return &AccessibleAccessRequestIterator{contract: _Accessible.contract, event: "AccessRequest", logs: logs, sub: sub}, nil
}

// WatchAccessRequest is a free log subscription operation binding the contract event 0xed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88.
//
// Solidity: event AccessRequest()
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
const BaseAccessControlGroupABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"members\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"grantAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"hasManagerAccess\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"grantManagerAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"revokeAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"hasAccess\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"revokeManagerAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"managers\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"MemberAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"ManagerAccessGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"MemberRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"ManagerAccessRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"operationCode\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"UnauthorizedOperation\",\"type\":\"event\"}]"

// BaseAccessControlGroupBin is the compiled bytecode used for deploying new contracts.
const BaseAccessControlGroupBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a0319908116329081178084556001805484169092178255600280549093163317909255600160a060020a039182168352600460209081526040808520805460ff199081168517909155855490941685526003909152909220805490911690911790556106888061008d6000396000f3006080604052600436106100cf5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100d157806308ae4b0c146101025780630ae5e7391461013757806341c0e1b51461015857806342e7ba7b1461016d5780636d2e4b1b1461018e57806375861a95146101af57806385e68531146101d05780638da5cb5b146101f157806395a078e814610206578063af570c0414610227578063cdb849b71461023c578063f2fde38b1461025d578063fdff9b4d1461027e575b005b3480156100dd57600080fd5b506100e661029f565b60408051600160a060020a039092168252519081900360200190f35b34801561010e57600080fd5b50610123600160a060020a03600435166102ae565b604080519115158252519081900360200190f35b34801561014357600080fd5b506100cf600160a060020a03600435166102c3565b34801561016457600080fd5b506100cf61033f565b34801561017957600080fd5b50610123600160a060020a036004351661037b565b34801561019a57600080fd5b506100cf600160a060020a036004351661039e565b3480156101bb57600080fd5b506100cf600160a060020a03600435166103f9565b3480156101dc57600080fd5b506100cf600160a060020a0360043516610482565b3480156101fd57600080fd5b506100e6610510565b34801561021257600080fd5b50610123600160a060020a036004351661051f565b34801561023357600080fd5b506100e6610542565b34801561024857600080fd5b506100cf600160a060020a0360043516610551565b34801561026957600080fd5b506100cf600160a060020a03600435166105d5565b34801561028a57600080fd5b50610123600160a060020a0360043516610647565b600054600160a060020a031681565b60036020526000908152604090205460ff1681565b3360009081526004602052604090205460ff1615156001146102e457600080fd5b600160a060020a038116600081815260036020908152604091829020805460ff19166001179055815192835290517fb251eb052afc73ffd02ffe85ad79990a8b3fed60d76dbc2fa2fdd7123dffd9149281900390910190a150565b600154600160a060020a03163214806103625750600154600160a060020a031633145b151561036d57600080fd5b600154600160a060020a0316ff5b600160a060020a031660009081526004602052604090205460ff16151560011490565b600054600160a060020a031632146103b557600080fd5b600160a060020a03811615156103ca57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031632148061041c5750600154600160a060020a031633145b151561042757600080fd5b600160a060020a038116600081815260046020908152604091829020805460ff19166001179055815192835290517f93bcaab179551bde429187645251f8e1fb8ac85801fcb1cf91eb2c9043d611179281900390910190a150565b3360009081526004602052604090205460ff161515600114806104ad575033600160a060020a038216145b15156104b857600080fd5b600160a060020a038116600081815260036020908152604091829020805460ff19169055815192835290517f745cd29407db644ed93e3ceb61cbcab96d1dfb496989ac5d5bf514fc5a9fab9c9281900390910190a150565b600154600160a060020a031681565b600160a060020a031660009081526003602052604090205460ff16151560011490565b600254600160a060020a031681565b600154600160a060020a0316331480610572575033600160a060020a038216145b151561057d57600080fd5b600160a060020a038116600081815260046020908152604091829020805460ff19169055815192835290517f2d6aa1a9629d125e23a0cf692cda7cd6795dff1652eedd4673b38ec31e387b959281900390910190a150565b600154600160a060020a03163214806105f85750600154600160a060020a031633145b151561060357600080fd5b600160a060020a038116151561061857600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60046020526000908152604090205460ff16815600a165627a7a7230582012ef4fbe6555a73830212786a66a6e16a4b3c0c313a0de46d8f4fb95ddf39afa0029`

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

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_BaseAccessControlGroup *BaseAccessControlGroupCaller) ContentSpace(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseAccessControlGroup.contract.Call(opts, out, "contentSpace")
	return *ret0, err
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) ContentSpace() (common.Address, error) {
	return _BaseAccessControlGroup.Contract.ContentSpace(&_BaseAccessControlGroup.CallOpts)
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_BaseAccessControlGroup *BaseAccessControlGroupCallerSession) ContentSpace() (common.Address, error) {
	return _BaseAccessControlGroup.Contract.ContentSpace(&_BaseAccessControlGroup.CallOpts)
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
// Solidity: function hasAccess(address candidate) constant returns(bool)
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
// Solidity: function hasAccess(address candidate) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) HasAccess(candidate common.Address) (bool, error) {
	return _BaseAccessControlGroup.Contract.HasAccess(&_BaseAccessControlGroup.CallOpts, candidate)
}

// HasAccess is a free data retrieval call binding the contract method 0x95a078e8.
//
// Solidity: function hasAccess(address candidate) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupCallerSession) HasAccess(candidate common.Address) (bool, error) {
	return _BaseAccessControlGroup.Contract.HasAccess(&_BaseAccessControlGroup.CallOpts, candidate)
}

// HasManagerAccess is a free data retrieval call binding the contract method 0x42e7ba7b.
//
// Solidity: function hasManagerAccess(address candidate) constant returns(bool)
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
// Solidity: function hasManagerAccess(address candidate) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) HasManagerAccess(candidate common.Address) (bool, error) {
	return _BaseAccessControlGroup.Contract.HasManagerAccess(&_BaseAccessControlGroup.CallOpts, candidate)
}

// HasManagerAccess is a free data retrieval call binding the contract method 0x42e7ba7b.
//
// Solidity: function hasManagerAccess(address candidate) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupCallerSession) HasManagerAccess(candidate common.Address) (bool, error) {
	return _BaseAccessControlGroup.Contract.HasManagerAccess(&_BaseAccessControlGroup.CallOpts, candidate)
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) constant returns(bool)
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
// Solidity: function managers(address ) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) Managers(arg0 common.Address) (bool, error) {
	return _BaseAccessControlGroup.Contract.Managers(&_BaseAccessControlGroup.CallOpts, arg0)
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupCallerSession) Managers(arg0 common.Address) (bool, error) {
	return _BaseAccessControlGroup.Contract.Managers(&_BaseAccessControlGroup.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) constant returns(bool)
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
// Solidity: function members(address ) constant returns(bool)
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) Members(arg0 common.Address) (bool, error) {
	return _BaseAccessControlGroup.Contract.Members(&_BaseAccessControlGroup.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) constant returns(bool)
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
// Solidity: function grantAccess(address candidate) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactor) GrantAccess(opts *bind.TransactOpts, candidate common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.contract.Transact(opts, "grantAccess", candidate)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x0ae5e739.
//
// Solidity: function grantAccess(address candidate) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) GrantAccess(candidate common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.GrantAccess(&_BaseAccessControlGroup.TransactOpts, candidate)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x0ae5e739.
//
// Solidity: function grantAccess(address candidate) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactorSession) GrantAccess(candidate common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.GrantAccess(&_BaseAccessControlGroup.TransactOpts, candidate)
}

// GrantManagerAccess is a paid mutator transaction binding the contract method 0x75861a95.
//
// Solidity: function grantManagerAccess(address manager) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactor) GrantManagerAccess(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.contract.Transact(opts, "grantManagerAccess", manager)
}

// GrantManagerAccess is a paid mutator transaction binding the contract method 0x75861a95.
//
// Solidity: function grantManagerAccess(address manager) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) GrantManagerAccess(manager common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.GrantManagerAccess(&_BaseAccessControlGroup.TransactOpts, manager)
}

// GrantManagerAccess is a paid mutator transaction binding the contract method 0x75861a95.
//
// Solidity: function grantManagerAccess(address manager) returns()
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
// Solidity: function revokeAccess(address candidate) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactor) RevokeAccess(opts *bind.TransactOpts, candidate common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.contract.Transact(opts, "revokeAccess", candidate)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x85e68531.
//
// Solidity: function revokeAccess(address candidate) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) RevokeAccess(candidate common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.RevokeAccess(&_BaseAccessControlGroup.TransactOpts, candidate)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x85e68531.
//
// Solidity: function revokeAccess(address candidate) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactorSession) RevokeAccess(candidate common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.RevokeAccess(&_BaseAccessControlGroup.TransactOpts, candidate)
}

// RevokeManagerAccess is a paid mutator transaction binding the contract method 0xcdb849b7.
//
// Solidity: function revokeManagerAccess(address manager) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactor) RevokeManagerAccess(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.contract.Transact(opts, "revokeManagerAccess", manager)
}

// RevokeManagerAccess is a paid mutator transaction binding the contract method 0xcdb849b7.
//
// Solidity: function revokeManagerAccess(address manager) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) RevokeManagerAccess(manager common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.RevokeManagerAccess(&_BaseAccessControlGroup.TransactOpts, manager)
}

// RevokeManagerAccess is a paid mutator transaction binding the contract method 0xcdb849b7.
//
// Solidity: function revokeManagerAccess(address manager) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactorSession) RevokeManagerAccess(manager common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.RevokeManagerAccess(&_BaseAccessControlGroup.TransactOpts, manager)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.TransferCreatorship(&_BaseAccessControlGroup.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.TransferCreatorship(&_BaseAccessControlGroup.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseAccessControlGroup.Contract.TransferOwnership(&_BaseAccessControlGroup.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
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
// Solidity: event ManagerAccessGranted(address candidate)
func (_BaseAccessControlGroup *BaseAccessControlGroupFilterer) FilterManagerAccessGranted(opts *bind.FilterOpts) (*BaseAccessControlGroupManagerAccessGrantedIterator, error) {

	logs, sub, err := _BaseAccessControlGroup.contract.FilterLogs(opts, "ManagerAccessGranted")
	if err != nil {
		return nil, err
	}
	return &BaseAccessControlGroupManagerAccessGrantedIterator{contract: _BaseAccessControlGroup.contract, event: "ManagerAccessGranted", logs: logs, sub: sub}, nil
}

// WatchManagerAccessGranted is a free log subscription operation binding the contract event 0x93bcaab179551bde429187645251f8e1fb8ac85801fcb1cf91eb2c9043d61117.
//
// Solidity: event ManagerAccessGranted(address candidate)
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
// Solidity: event ManagerAccessRevoked(address candidate)
func (_BaseAccessControlGroup *BaseAccessControlGroupFilterer) FilterManagerAccessRevoked(opts *bind.FilterOpts) (*BaseAccessControlGroupManagerAccessRevokedIterator, error) {

	logs, sub, err := _BaseAccessControlGroup.contract.FilterLogs(opts, "ManagerAccessRevoked")
	if err != nil {
		return nil, err
	}
	return &BaseAccessControlGroupManagerAccessRevokedIterator{contract: _BaseAccessControlGroup.contract, event: "ManagerAccessRevoked", logs: logs, sub: sub}, nil
}

// WatchManagerAccessRevoked is a free log subscription operation binding the contract event 0x2d6aa1a9629d125e23a0cf692cda7cd6795dff1652eedd4673b38ec31e387b95.
//
// Solidity: event ManagerAccessRevoked(address candidate)
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
// Solidity: event MemberAdded(address candidate)
func (_BaseAccessControlGroup *BaseAccessControlGroupFilterer) FilterMemberAdded(opts *bind.FilterOpts) (*BaseAccessControlGroupMemberAddedIterator, error) {

	logs, sub, err := _BaseAccessControlGroup.contract.FilterLogs(opts, "MemberAdded")
	if err != nil {
		return nil, err
	}
	return &BaseAccessControlGroupMemberAddedIterator{contract: _BaseAccessControlGroup.contract, event: "MemberAdded", logs: logs, sub: sub}, nil
}

// WatchMemberAdded is a free log subscription operation binding the contract event 0xb251eb052afc73ffd02ffe85ad79990a8b3fed60d76dbc2fa2fdd7123dffd914.
//
// Solidity: event MemberAdded(address candidate)
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
// Solidity: event MemberRevoked(address candidate)
func (_BaseAccessControlGroup *BaseAccessControlGroupFilterer) FilterMemberRevoked(opts *bind.FilterOpts) (*BaseAccessControlGroupMemberRevokedIterator, error) {

	logs, sub, err := _BaseAccessControlGroup.contract.FilterLogs(opts, "MemberRevoked")
	if err != nil {
		return nil, err
	}
	return &BaseAccessControlGroupMemberRevokedIterator{contract: _BaseAccessControlGroup.contract, event: "MemberRevoked", logs: logs, sub: sub}, nil
}

// WatchMemberRevoked is a free log subscription operation binding the contract event 0x745cd29407db644ed93e3ceb61cbcab96d1dfb496989ac5d5bf514fc5a9fab9c.
//
// Solidity: event MemberRevoked(address candidate)
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
// Solidity: event UnauthorizedOperation(uint256 operationCode, address candidate)
func (_BaseAccessControlGroup *BaseAccessControlGroupFilterer) FilterUnauthorizedOperation(opts *bind.FilterOpts) (*BaseAccessControlGroupUnauthorizedOperationIterator, error) {

	logs, sub, err := _BaseAccessControlGroup.contract.FilterLogs(opts, "UnauthorizedOperation")
	if err != nil {
		return nil, err
	}
	return &BaseAccessControlGroupUnauthorizedOperationIterator{contract: _BaseAccessControlGroup.contract, event: "UnauthorizedOperation", logs: logs, sub: sub}, nil
}

// WatchUnauthorizedOperation is a free log subscription operation binding the contract event 0x23de2adc3e22f171f66b3e5a333e17feb9dc30ba9570933bd259cb6c13ef7ab7.
//
// Solidity: event UnauthorizedOperation(uint256 operationCode, address candidate)
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
const BaseContentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"STATUS_PUBLISHED\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"publish\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestMap\",\"outputs\":[{\"name\":\"originator\",\"type\":\"address\"},{\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"statusCode\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressKMS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentType\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"status_code\",\"type\":\"int256\"}],\"name\":\"statusCodeDescription\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STATUS_DRAFT\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"status_code\",\"type\":\"int256\"}],\"name\":\"setStatusCode\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"request_ID\",\"type\":\"uint256\"},{\"name\":\"score_pct\",\"type\":\"uint256\"},{\"name\":\"ml_out_hash\",\"type\":\"bytes32\"}],\"name\":\"accessComplete\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accessCharge\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"status_code\",\"type\":\"int256\"}],\"name\":\"updateStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"level\",\"type\":\"uint8\"},{\"name\":\"custom_values\",\"type\":\"bytes32[]\"},{\"name\":\"stakeholders\",\"type\":\"address[]\"}],\"name\":\"getAccessCharge\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requestID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"level\",\"type\":\"uint8\"},{\"name\":\"pke_requestor\",\"type\":\"string\"},{\"name\":\"pke_AFGH\",\"type\":\"string\"},{\"name\":\"custom_values\",\"type\":\"bytes32[]\"},{\"name\":\"stakeholders\",\"type\":\"address[]\"}],\"name\":\"accessRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"libraryAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"address_KMS\",\"type\":\"address\"}],\"name\":\"setAddressKMS\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STATUS_REVIEW\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"objectHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setContentContractAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"request_ID\",\"type\":\"uint256\"},{\"name\":\"access_granted\",\"type\":\"bool\"},{\"name\":\"re_key\",\"type\":\"string\"},{\"name\":\"encrypted_AES_key\",\"type\":\"string\"}],\"name\":\"accessGrant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"object_hash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"charge\",\"type\":\"uint256\"}],\"name\":\"setAccessCharge\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"statusDescription\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"content_type\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"containingLibrary\",\"type\":\"address\"}],\"name\":\"ContentObjectCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentType\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"contentContractAddress\",\"type\":\"address\"}],\"name\":\"SetContentType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"contentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"pkeRequestor\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"pkeAFGH\",\"type\":\"string\"}],\"name\":\"AccessRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"access_granted\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"reKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"encryptedAESKey\",\"type\":\"string\"}],\"name\":\"AccessGrant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"customValue\",\"type\":\"bytes32\"}],\"name\":\"AccessRequestValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"stakeholder\",\"type\":\"address\"}],\"name\":\"AccessRequestStakeholder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"scorePct\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"mlOutHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"customContractResult\",\"type\":\"bool\"}],\"name\":\"AccessComplete\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentContractAddress\",\"type\":\"address\"}],\"name\":\"SetContentContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accessCharge\",\"type\":\"uint256\"}],\"name\":\"SetAccessCharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"accessCharge\",\"type\":\"uint256\"}],\"name\":\"GetAccessCharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accessCharge\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amountProvided\",\"type\":\"uint256\"}],\"name\":\"InsufficientFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"statusCode\",\"type\":\"int256\"}],\"name\":\"SetStatusCode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestStatus\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"statusCode\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"Publish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"custom_contract\",\"type\":\"address\"}],\"name\":\"InvokeCustomPreHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"custom_contract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"ReturnCustomHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"custom_contract\",\"type\":\"address\"}],\"name\":\"InvokeCustomPostHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"Commit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"UpdateRequest\",\"type\":\"event\"}]"

// BaseContentBin is the compiled bytecode used for deploying new contracts.
const BaseContentBin = `0x6080604052604051602080611d80833981016040818152915160008054600160a060020a0319908116329081178355600180548316909117905560068054821633179081905560001960085560099290925560038054600160a060020a038086169190931617905516825291517fc3decc188980e855666b70498ca85e8fa284d97d30483d828fa126f7303d7d199181900360200190a150611cda806100a66000396000f3006080604052600436106101705763ffffffff60e060020a60003504166217de98811461017257806302d05d3f14610199578063075d4782146101ca5780631a735f18146101e65780632310167f1461022d57806327c1c21d1461024257806332eaf21b1461025757806336ebffca1461026c578063388642841461028157806341c0e1b5146102995780634dd70788146102ae5780635267db44146102c35780635cc4aa9b146102db57806364ade32b146102ec5780636d2e4b1b146103015780638280dd8f14610322578063879fe48f1461033a5780638da5cb5b146103d05780638f779201146103e5578063a1ff106e146103fa578063b816f513146104fe578063c287e0ed14610513578063c9e8e72d14610528578063d810f8c814610549578063e02dd9c21461055e578063e538530314610573578063ee56d76714610594578063f14fcbc814610634578063f2fde38b1461064c578063f4d9bae81461066d578063f81ab0ae14610685575b005b34801561017e57600080fd5b5061018761069a565b60408051918252519081900360200190f35b3480156101a557600080fd5b506101ae6106be565b60408051600160a060020a039092168252519081900360200190f35b6101d26106cd565b604080519115158252519081900360200190f35b3480156101f257600080fd5b506101fe6004356107fc565b60408051600160a060020a0390941684526020840192909252600090810b900b82820152519081900360600190f35b34801561023957600080fd5b506101ae610827565b34801561024e57600080fd5b50610187610836565b34801561026357600080fd5b506101ae61083c565b34801561027857600080fd5b506101ae61084b565b34801561028d57600080fd5b5061018760043561085a565b3480156102a557600080fd5b506101706109ac565b3480156102ba57600080fd5b50610187610a91565b3480156102cf57600080fd5b50610187600435610ab5565b6101d2600435602435604435610b4a565b3480156102f857600080fd5b50610187610d0b565b34801561030d57600080fd5b50610170600160a060020a0360043516610d11565b34801561032e57600080fd5b50610187600435610d6c565b34801561034657600080fd5b5060408051602060046024803582810135848102808701860190975280865261018796843560ff1696369660449591949091019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750949750610f159650505050505050565b3480156103dc57600080fd5b506101ae6110ae565b3480156103f157600080fd5b506101876110bd565b60408051602060046024803582810135601f81018590048502860185019096528585526101d295833560ff1695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506110c39650505050505050565b34801561050a57600080fd5b506101ae6115a5565b34801561051f57600080fd5b506101706115b4565b34801561053457600080fd5b50610170600160a060020a0360043516611619565b34801561055557600080fd5b50610187611676565b34801561056a57600080fd5b5061018761169a565b34801561057f57600080fd5b50610170600160a060020a03600435166116a0565b3480156105a057600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101d2948235946024803515159536959460649492019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506118649650505050505050565b34801561064057600080fd5b50610170600435611af4565b34801561065857600080fd5b50610170600160a060020a0360043516611b5d565b34801561067957600080fd5b50610187600435611bcf565b34801561069157600080fd5b50610187611c40565b7f5075626c6973686564000000000000000000000000000000000000000000000081565b600054600160a060020a031681565b60015460009081908190600160a060020a03163214806106f75750600154600160a060020a031633145b151561070257600080fd5b61070c6001610d6c565b5060009150600060085413156107ad5750600654604080517f49102e610000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169182916349102e619160048083019260209291908290030181600087803b15801561077e57600080fd5b505af1158015610792573d6000803e3d6000fd5b505050506040513d60208110156107a857600080fd5b505191505b600854600254604080518515158152602081019390935282810191909152517f901e6f3cdc4c61620d5d424116934b9af6e31ba79cdeaa349336d93ecfe846d49181900360600190a150919050565b600a60205260009081526040812080546001820154600290920154600160a060020a03909116920b83565b600554600160a060020a031681565b60085481565b600454600160a060020a031681565b600354600160a060020a031681565b60055460009081908190600160a060020a03161561090a5750600554604080517f45080442000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a0390921691829163450804429160248083019260209291908290030181600087803b1580156108db57600080fd5b505af11580156108ef573d6000803e3d6000fd5b505050506040513d602081101561090557600080fd5b505191505b8115610918578192506109a5565b831515610947577f5075626c6973686564000000000000000000000000000000000000000000000092506109a5565b6000841215610978577f447261667400000000000000000000000000000000000000000000000000000092506109a5565b60008413156109a5577f447261667420696e20726576696577000000000000000000000000000000000092505b5050919050565b600154600090600160a060020a03163214806109d25750600154600160a060020a031633145b15156109dd57600080fd5b600554600160a060020a031615610a865750600554604080517f9e99bbea0000000000000000000000000000000000000000000000000000000081529051600160a060020a03909216918291639e99bbea9160048083019260209291908290030181600087803b158015610a5057600080fd5b505af1158015610a64573d6000803e3d6000fd5b505050506040513d6020811015610a7a57600080fd5b505115610a8657600080fd5b610a8e611c52565b50565b7f447261667400000000000000000000000000000000000000000000000000000081565b600154600090600160a060020a031632148015610ae957506000821280610ae95750600082138015610ae957506000600854125b15610af45760088290555b600654600160a060020a0316331415610b0d5760088290555b60085460408051918252517fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a29181900360200190a1505060085490565b6000838152600a60205260408120805482908190600160a060020a031615801590610b7e57508254600160a060020a031633145b1515610b8957600080fd5b60055460019250600160a060020a031615610c375750600554604080517f6b2d1324000000000000000000000000000000000000000000000000000000008152600481018990529051600160a060020a03909216918291636b2d13249160248083019260209291908290030181600087803b158015610c0757600080fd5b505af1158015610c1b573d6000803e3d6000fd5b505050506040513d6020811015610c3157600080fd5b50511591505b6002830154600090810b900b1515610c7b576001830154604051339180156108fc02916000818181858888f19350505050158015610c79573d6000803e3d6000fd5b505b6000878152600a60209081526040808320805473ffffffffffffffffffffffffffffffffffffffff1916815560018101939093556002909201805460ff191690558151898152908101889052808201879052831515606082015290517f7f1f4b28434ce7beab4983e64a8b5bb96e195a67029fdaff925028aec57fbc6b9181900360800190a15095945050505050565b60075481565b600054600160a060020a03163214610d2857600080fd5b600160a060020a0381161515610d3d57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60015460009081908190600160a060020a0316321480610d965750600154600160a060020a031633145b80610dab5750600654600160a060020a031633145b1515610db657600080fd5b600554600160a060020a03161515610e3a57600154600160a060020a0316321480610deb5750600154600160a060020a031633145b8015610e025750836000191480610e025750836001145b15610e0f57839150610e35565b600654600160a060020a031633148015610e2c5750600060085412155b15610e35578391505b610ed2565b50600554604080517f3513a805000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a03909216918291633513a8059160248083019260209291908290030181600087803b158015610ea357600080fd5b505af1158015610eb7573d6000803e3d6000fd5b505050506040513d6020811015610ecd57600080fd5b505191505b60088290556040805183815290517fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a29181900360200190a1505060085492915050565b6007546005546000919082908190600160a060020a031615611065576005546040517ff8117ca200000000000000000000000000000000000000000000000000000000815260ff891660048201908152606060248301908152895160648401528951600160a060020a039094169550859363f8117ca2938c938c938c93919290916044820191608401906020808801910280838360005b83811015610fc4578181015183820152602001610fac565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015611003578181015183820152602001610feb565b5050505090500195505050505050602060405180830381600087803b15801561102b57600080fd5b505af115801561103f573d6000803e3d6000fd5b505050506040513d602081101561105557600080fd5b5051905060008112611065578092505b6040805160ff891681526020810185905281517fa58326ee5bb617cb8b4f0d0f5f557c469d2d05d7a738f777037deda9c724b370929181900390910190a1509095945050505050565b600154600160a060020a031681565b60095481565b60008060006110d0611c8e565b600980546001019055600654604080517f95a078e80000000000000000000000000000000000000000000000000000000081523260048201529051600160a060020a0390921694506000918291829187916395a078e891602480830192602092919082900301818787803b15801561114757600080fd5b505af115801561115b573d6000803e3d6000fd5b505050506040513d602081101561117157600080fd5b5051151561117e57600080fd5b600154600160a060020a031632146111aa5761119b8c8a8a610f15565b9450348511156111aa57600080fd5b6040805160608101825233815234602080830191825260008385018181526009548252600a90925293842083518154600160a060020a0391821673ffffffffffffffffffffffffffffffffffffffff19909116178255925160018201559051600290910180549190940b60ff1660ff1990911617909255600554909550161561136357600560009054906101000a9004600160a060020a0316925082600160a060020a031663123e0e806009548e8c8c6040518563ffffffff1660e060020a028152600401808581526020018460ff1660ff1681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156112c25781810151838201526020016112aa565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156113015781810151838201526020016112e9565b505050509050019650505050505050602060405180830381600087803b15801561132a57600080fd5b505af115801561133e573d6000803e3d6000fd5b505050506040513d602081101561135457600080fd5b50519150811561136357600080fd5b7f089a6f1788a3c353423e1be4ba12533bdde7d908bb41abeee185af0acb3df5626009548d6002548e8e604051808681526020018560ff1660ff16815260200184600019166000191681526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156113ee5781810151838201526020016113d6565b50505050905090810190601f16801561141b5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561144e578181015183820152602001611436565b50505050905090810190601f16801561147b5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15060005b885181101561150b5788818151811015156114a857fe5b6020908102909101015115611503577f515e0a48b385fce2a8e4d9f169a97c4f6ea669a752358f5e6ab37cc3c2e84c3889828151811015156114e657fe5b602090810290910181015160408051918252519081900390910190a15b600101611491565b5060005b875181101561159357888181518110151561152657fe5b602090810290910101511561158b577fb6e3239e521a6c66920ae634f8e921a37e6991d520ac44d52f8516397f41b684888281518110151561156457fe5b602090810290910181015160408051600160a060020a039092168252519081900390910190a15b60010161150f565b5060019b9a5050505050505050505050565b600654600160a060020a031681565b600154600160a060020a03163214806115d75750600154600160a060020a031633145b15156115e257600080fd5b60025460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600154600160a060020a031632148061163c5750600154600160a060020a031633145b151561164757600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b7f447261667420696e20726576696577000000000000000000000000000000000081565b60025481565b60015460009081908190600160a060020a03163214806116ca5750600154600160a060020a031633145b15156116d557600080fd5b600554600160a060020a03161561177557600560009054906101000a9004600160a060020a0316925082600160a060020a0316639e99bbea6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561173c57600080fd5b505af1158015611750573d6000803e3d6000fd5b505050506040513d602081101561176657600080fd5b50519150811561177557600080fd5b6005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0386169081179091551561181f5783925082600160a060020a0316637b1cdb3e6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156117e657600080fd5b505af11580156117fa573d6000803e3d6000fd5b505050506040513d602081101561181057600080fd5b50519050801561181f57600080fd5b60055460408051600160a060020a039092168252517fa6f2e38f0cfebf27212317fced3ac40bc62e00bd33f38d69603710740c69acb79181900360200190a150505050565b6001546000908190600160a060020a031633148061188c5750600454600160a060020a031633145b151561189757600080fd5b506000858152600a602052604090208054600160a060020a031615156118bc57600080fd5b6002810154600090810b900b1515611ae85784151561197d5780546001820154604051600160a060020a039092169181156108fc0291906000818181858888f19350505050158015611912573d6000803e3d6000fd5b5060028101805460ff191660ff179055604080518781526000602082018190526080828401819052820181905260c06060830181905282015290517f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718918190036101000190a1611ae8565b6001805490820154604051600160a060020a039092169181156108fc0291906000818181858888f193505050501580156119bb573d6000803e3d6000fd5b5060028101805460ff19166001908117909155604080518881526020808201849052608092820183815288519383019390935287517f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718948b9490938a938a93919291606084019160a08501919087019080838360005b83811015611a49578181015183820152602001611a31565b50505050905090810190601f168015611a765780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611aa9578181015183820152602001611a91565b50505050905090810190601f168015611ad65780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a15b50600195945050505050565b600154600160a060020a0316321480611b175750600154600160a060020a031633145b1515611b2257600080fd5b60028190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b600154600160a060020a0316321480611b805750600154600160a060020a031633145b1515611b8b57600080fd5b600160a060020a0381161515611ba057600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600090600160a060020a0316321480611bf55750600154600160a060020a031633145b1515611c0057600080fd5b60078290556040805183815290517f4114f8ef80b6de2161db580cbefa14e1892d15d3ebe2062c9914e4a5773114a39181900360200190a1505060075490565b6000611c4d60085461085a565b905090565b600154600160a060020a0316321480611c755750600154600160a060020a031633145b1515611c8057600080fd5b600154600160a060020a0316ff5b6040805160608101825260008082526020820181905291810191909152905600a165627a7a723058209a663fbea2d3d07ef8d8b67303e675cf66c7f16a6dfe96500c87f5d31e35f0080029`

// DeployBaseContent deploys a new Ethereum contract, binding an instance of BaseContent to it.
func DeployBaseContent(auth *bind.TransactOpts, backend bind.ContractBackend, content_type common.Address) (common.Address, *types.Transaction, *BaseContent, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseContentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaseContentBin), backend, content_type)
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

// STATUSDRAFT is a free data retrieval call binding the contract method 0x4dd70788.
//
// Solidity: function STATUS_DRAFT() constant returns(bytes32)
func (_BaseContent *BaseContentCaller) STATUSDRAFT(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "STATUS_DRAFT")
	return *ret0, err
}

// STATUSDRAFT is a free data retrieval call binding the contract method 0x4dd70788.
//
// Solidity: function STATUS_DRAFT() constant returns(bytes32)
func (_BaseContent *BaseContentSession) STATUSDRAFT() ([32]byte, error) {
	return _BaseContent.Contract.STATUSDRAFT(&_BaseContent.CallOpts)
}

// STATUSDRAFT is a free data retrieval call binding the contract method 0x4dd70788.
//
// Solidity: function STATUS_DRAFT() constant returns(bytes32)
func (_BaseContent *BaseContentCallerSession) STATUSDRAFT() ([32]byte, error) {
	return _BaseContent.Contract.STATUSDRAFT(&_BaseContent.CallOpts)
}

// STATUSPUBLISHED is a free data retrieval call binding the contract method 0x0017de98.
//
// Solidity: function STATUS_PUBLISHED() constant returns(bytes32)
func (_BaseContent *BaseContentCaller) STATUSPUBLISHED(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "STATUS_PUBLISHED")
	return *ret0, err
}

// STATUSPUBLISHED is a free data retrieval call binding the contract method 0x0017de98.
//
// Solidity: function STATUS_PUBLISHED() constant returns(bytes32)
func (_BaseContent *BaseContentSession) STATUSPUBLISHED() ([32]byte, error) {
	return _BaseContent.Contract.STATUSPUBLISHED(&_BaseContent.CallOpts)
}

// STATUSPUBLISHED is a free data retrieval call binding the contract method 0x0017de98.
//
// Solidity: function STATUS_PUBLISHED() constant returns(bytes32)
func (_BaseContent *BaseContentCallerSession) STATUSPUBLISHED() ([32]byte, error) {
	return _BaseContent.Contract.STATUSPUBLISHED(&_BaseContent.CallOpts)
}

// STATUSREVIEW is a free data retrieval call binding the contract method 0xd810f8c8.
//
// Solidity: function STATUS_REVIEW() constant returns(bytes32)
func (_BaseContent *BaseContentCaller) STATUSREVIEW(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "STATUS_REVIEW")
	return *ret0, err
}

// STATUSREVIEW is a free data retrieval call binding the contract method 0xd810f8c8.
//
// Solidity: function STATUS_REVIEW() constant returns(bytes32)
func (_BaseContent *BaseContentSession) STATUSREVIEW() ([32]byte, error) {
	return _BaseContent.Contract.STATUSREVIEW(&_BaseContent.CallOpts)
}

// STATUSREVIEW is a free data retrieval call binding the contract method 0xd810f8c8.
//
// Solidity: function STATUS_REVIEW() constant returns(bytes32)
func (_BaseContent *BaseContentCallerSession) STATUSREVIEW() ([32]byte, error) {
	return _BaseContent.Contract.STATUSREVIEW(&_BaseContent.CallOpts)
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

// GetAccessCharge is a free data retrieval call binding the contract method 0x879fe48f.
//
// Solidity: function getAccessCharge(uint8 level, bytes32[] custom_values, address[] stakeholders) constant returns(uint256)
func (_BaseContent *BaseContentCaller) GetAccessCharge(opts *bind.CallOpts, level uint8, custom_values [][32]byte, stakeholders []common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "getAccessCharge", level, custom_values, stakeholders)
	return *ret0, err
}

// GetAccessCharge is a free data retrieval call binding the contract method 0x879fe48f.
//
// Solidity: function getAccessCharge(uint8 level, bytes32[] custom_values, address[] stakeholders) constant returns(uint256)
func (_BaseContent *BaseContentSession) GetAccessCharge(level uint8, custom_values [][32]byte, stakeholders []common.Address) (*big.Int, error) {
	return _BaseContent.Contract.GetAccessCharge(&_BaseContent.CallOpts, level, custom_values, stakeholders)
}

// GetAccessCharge is a free data retrieval call binding the contract method 0x879fe48f.
//
// Solidity: function getAccessCharge(uint8 level, bytes32[] custom_values, address[] stakeholders) constant returns(uint256)
func (_BaseContent *BaseContentCallerSession) GetAccessCharge(level uint8, custom_values [][32]byte, stakeholders []common.Address) (*big.Int, error) {
	return _BaseContent.Contract.GetAccessCharge(&_BaseContent.CallOpts, level, custom_values, stakeholders)
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
// Solidity: function requestMap(uint256 ) constant returns(address originator, uint256 amountPaid, int8 status)
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
// Solidity: function requestMap(uint256 ) constant returns(address originator, uint256 amountPaid, int8 status)
func (_BaseContent *BaseContentSession) RequestMap(arg0 *big.Int) (struct {
	Originator common.Address
	AmountPaid *big.Int
	Status     int8
}, error) {
	return _BaseContent.Contract.RequestMap(&_BaseContent.CallOpts, arg0)
}

// RequestMap is a free data retrieval call binding the contract method 0x1a735f18.
//
// Solidity: function requestMap(uint256 ) constant returns(address originator, uint256 amountPaid, int8 status)
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

// StatusCodeDescription is a free data retrieval call binding the contract method 0x38864284.
//
// Solidity: function statusCodeDescription(int256 status_code) constant returns(bytes32)
func (_BaseContent *BaseContentCaller) StatusCodeDescription(opts *bind.CallOpts, status_code *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "statusCodeDescription", status_code)
	return *ret0, err
}

// StatusCodeDescription is a free data retrieval call binding the contract method 0x38864284.
//
// Solidity: function statusCodeDescription(int256 status_code) constant returns(bytes32)
func (_BaseContent *BaseContentSession) StatusCodeDescription(status_code *big.Int) ([32]byte, error) {
	return _BaseContent.Contract.StatusCodeDescription(&_BaseContent.CallOpts, status_code)
}

// StatusCodeDescription is a free data retrieval call binding the contract method 0x38864284.
//
// Solidity: function statusCodeDescription(int256 status_code) constant returns(bytes32)
func (_BaseContent *BaseContentCallerSession) StatusCodeDescription(status_code *big.Int) ([32]byte, error) {
	return _BaseContent.Contract.StatusCodeDescription(&_BaseContent.CallOpts, status_code)
}

// StatusDescription is a free data retrieval call binding the contract method 0xf81ab0ae.
//
// Solidity: function statusDescription() constant returns(bytes32)
func (_BaseContent *BaseContentCaller) StatusDescription(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "statusDescription")
	return *ret0, err
}

// StatusDescription is a free data retrieval call binding the contract method 0xf81ab0ae.
//
// Solidity: function statusDescription() constant returns(bytes32)
func (_BaseContent *BaseContentSession) StatusDescription() ([32]byte, error) {
	return _BaseContent.Contract.StatusDescription(&_BaseContent.CallOpts)
}

// StatusDescription is a free data retrieval call binding the contract method 0xf81ab0ae.
//
// Solidity: function statusDescription() constant returns(bytes32)
func (_BaseContent *BaseContentCallerSession) StatusDescription() ([32]byte, error) {
	return _BaseContent.Contract.StatusDescription(&_BaseContent.CallOpts)
}

// AccessComplete is a paid mutator transaction binding the contract method 0x5cc4aa9b.
//
// Solidity: function accessComplete(uint256 request_ID, uint256 score_pct, bytes32 ml_out_hash) returns(bool)
func (_BaseContent *BaseContentTransactor) AccessComplete(opts *bind.TransactOpts, request_ID *big.Int, score_pct *big.Int, ml_out_hash [32]byte) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "accessComplete", request_ID, score_pct, ml_out_hash)
}

// AccessComplete is a paid mutator transaction binding the contract method 0x5cc4aa9b.
//
// Solidity: function accessComplete(uint256 request_ID, uint256 score_pct, bytes32 ml_out_hash) returns(bool)
func (_BaseContent *BaseContentSession) AccessComplete(request_ID *big.Int, score_pct *big.Int, ml_out_hash [32]byte) (*types.Transaction, error) {
	return _BaseContent.Contract.AccessComplete(&_BaseContent.TransactOpts, request_ID, score_pct, ml_out_hash)
}

// AccessComplete is a paid mutator transaction binding the contract method 0x5cc4aa9b.
//
// Solidity: function accessComplete(uint256 request_ID, uint256 score_pct, bytes32 ml_out_hash) returns(bool)
func (_BaseContent *BaseContentTransactorSession) AccessComplete(request_ID *big.Int, score_pct *big.Int, ml_out_hash [32]byte) (*types.Transaction, error) {
	return _BaseContent.Contract.AccessComplete(&_BaseContent.TransactOpts, request_ID, score_pct, ml_out_hash)
}

// AccessGrant is a paid mutator transaction binding the contract method 0xee56d767.
//
// Solidity: function accessGrant(uint256 request_ID, bool access_granted, string re_key, string encrypted_AES_key) returns(bool)
func (_BaseContent *BaseContentTransactor) AccessGrant(opts *bind.TransactOpts, request_ID *big.Int, access_granted bool, re_key string, encrypted_AES_key string) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "accessGrant", request_ID, access_granted, re_key, encrypted_AES_key)
}

// AccessGrant is a paid mutator transaction binding the contract method 0xee56d767.
//
// Solidity: function accessGrant(uint256 request_ID, bool access_granted, string re_key, string encrypted_AES_key) returns(bool)
func (_BaseContent *BaseContentSession) AccessGrant(request_ID *big.Int, access_granted bool, re_key string, encrypted_AES_key string) (*types.Transaction, error) {
	return _BaseContent.Contract.AccessGrant(&_BaseContent.TransactOpts, request_ID, access_granted, re_key, encrypted_AES_key)
}

// AccessGrant is a paid mutator transaction binding the contract method 0xee56d767.
//
// Solidity: function accessGrant(uint256 request_ID, bool access_granted, string re_key, string encrypted_AES_key) returns(bool)
func (_BaseContent *BaseContentTransactorSession) AccessGrant(request_ID *big.Int, access_granted bool, re_key string, encrypted_AES_key string) (*types.Transaction, error) {
	return _BaseContent.Contract.AccessGrant(&_BaseContent.TransactOpts, request_ID, access_granted, re_key, encrypted_AES_key)
}

// AccessRequest is a paid mutator transaction binding the contract method 0xa1ff106e.
//
// Solidity: function accessRequest(uint8 level, string pke_requestor, string pke_AFGH, bytes32[] custom_values, address[] stakeholders) returns(bool)
func (_BaseContent *BaseContentTransactor) AccessRequest(opts *bind.TransactOpts, level uint8, pke_requestor string, pke_AFGH string, custom_values [][32]byte, stakeholders []common.Address) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "accessRequest", level, pke_requestor, pke_AFGH, custom_values, stakeholders)
}

// AccessRequest is a paid mutator transaction binding the contract method 0xa1ff106e.
//
// Solidity: function accessRequest(uint8 level, string pke_requestor, string pke_AFGH, bytes32[] custom_values, address[] stakeholders) returns(bool)
func (_BaseContent *BaseContentSession) AccessRequest(level uint8, pke_requestor string, pke_AFGH string, custom_values [][32]byte, stakeholders []common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.AccessRequest(&_BaseContent.TransactOpts, level, pke_requestor, pke_AFGH, custom_values, stakeholders)
}

// AccessRequest is a paid mutator transaction binding the contract method 0xa1ff106e.
//
// Solidity: function accessRequest(uint8 level, string pke_requestor, string pke_AFGH, bytes32[] custom_values, address[] stakeholders) returns(bool)
func (_BaseContent *BaseContentTransactorSession) AccessRequest(level uint8, pke_requestor string, pke_AFGH string, custom_values [][32]byte, stakeholders []common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.AccessRequest(&_BaseContent.TransactOpts, level, pke_requestor, pke_AFGH, custom_values, stakeholders)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 object_hash) returns()
func (_BaseContent *BaseContentTransactor) Commit(opts *bind.TransactOpts, object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "commit", object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 object_hash) returns()
func (_BaseContent *BaseContentSession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContent.Contract.Commit(&_BaseContent.TransactOpts, object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 object_hash) returns()
func (_BaseContent *BaseContentTransactorSession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContent.Contract.Commit(&_BaseContent.TransactOpts, object_hash)
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

// Publish is a paid mutator transaction binding the contract method 0x075d4782.
//
// Solidity: function publish() returns(bool)
func (_BaseContent *BaseContentTransactor) Publish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "publish")
}

// Publish is a paid mutator transaction binding the contract method 0x075d4782.
//
// Solidity: function publish() returns(bool)
func (_BaseContent *BaseContentSession) Publish() (*types.Transaction, error) {
	return _BaseContent.Contract.Publish(&_BaseContent.TransactOpts)
}

// Publish is a paid mutator transaction binding the contract method 0x075d4782.
//
// Solidity: function publish() returns(bool)
func (_BaseContent *BaseContentTransactorSession) Publish() (*types.Transaction, error) {
	return _BaseContent.Contract.Publish(&_BaseContent.TransactOpts)
}

// SetAccessCharge is a paid mutator transaction binding the contract method 0xf4d9bae8.
//
// Solidity: function setAccessCharge(uint256 charge) returns(uint256)
func (_BaseContent *BaseContentTransactor) SetAccessCharge(opts *bind.TransactOpts, charge *big.Int) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "setAccessCharge", charge)
}

// SetAccessCharge is a paid mutator transaction binding the contract method 0xf4d9bae8.
//
// Solidity: function setAccessCharge(uint256 charge) returns(uint256)
func (_BaseContent *BaseContentSession) SetAccessCharge(charge *big.Int) (*types.Transaction, error) {
	return _BaseContent.Contract.SetAccessCharge(&_BaseContent.TransactOpts, charge)
}

// SetAccessCharge is a paid mutator transaction binding the contract method 0xf4d9bae8.
//
// Solidity: function setAccessCharge(uint256 charge) returns(uint256)
func (_BaseContent *BaseContentTransactorSession) SetAccessCharge(charge *big.Int) (*types.Transaction, error) {
	return _BaseContent.Contract.SetAccessCharge(&_BaseContent.TransactOpts, charge)
}

// SetAddressKMS is a paid mutator transaction binding the contract method 0xc9e8e72d.
//
// Solidity: function setAddressKMS(address address_KMS) returns()
func (_BaseContent *BaseContentTransactor) SetAddressKMS(opts *bind.TransactOpts, address_KMS common.Address) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "setAddressKMS", address_KMS)
}

// SetAddressKMS is a paid mutator transaction binding the contract method 0xc9e8e72d.
//
// Solidity: function setAddressKMS(address address_KMS) returns()
func (_BaseContent *BaseContentSession) SetAddressKMS(address_KMS common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.SetAddressKMS(&_BaseContent.TransactOpts, address_KMS)
}

// SetAddressKMS is a paid mutator transaction binding the contract method 0xc9e8e72d.
//
// Solidity: function setAddressKMS(address address_KMS) returns()
func (_BaseContent *BaseContentTransactorSession) SetAddressKMS(address_KMS common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.SetAddressKMS(&_BaseContent.TransactOpts, address_KMS)
}

// SetContentContractAddress is a paid mutator transaction binding the contract method 0xe5385303.
//
// Solidity: function setContentContractAddress(address addr) returns()
func (_BaseContent *BaseContentTransactor) SetContentContractAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "setContentContractAddress", addr)
}

// SetContentContractAddress is a paid mutator transaction binding the contract method 0xe5385303.
//
// Solidity: function setContentContractAddress(address addr) returns()
func (_BaseContent *BaseContentSession) SetContentContractAddress(addr common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.SetContentContractAddress(&_BaseContent.TransactOpts, addr)
}

// SetContentContractAddress is a paid mutator transaction binding the contract method 0xe5385303.
//
// Solidity: function setContentContractAddress(address addr) returns()
func (_BaseContent *BaseContentTransactorSession) SetContentContractAddress(addr common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.SetContentContractAddress(&_BaseContent.TransactOpts, addr)
}

// SetStatusCode is a paid mutator transaction binding the contract method 0x5267db44.
//
// Solidity: function setStatusCode(int256 status_code) returns(int256)
func (_BaseContent *BaseContentTransactor) SetStatusCode(opts *bind.TransactOpts, status_code *big.Int) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "setStatusCode", status_code)
}

// SetStatusCode is a paid mutator transaction binding the contract method 0x5267db44.
//
// Solidity: function setStatusCode(int256 status_code) returns(int256)
func (_BaseContent *BaseContentSession) SetStatusCode(status_code *big.Int) (*types.Transaction, error) {
	return _BaseContent.Contract.SetStatusCode(&_BaseContent.TransactOpts, status_code)
}

// SetStatusCode is a paid mutator transaction binding the contract method 0x5267db44.
//
// Solidity: function setStatusCode(int256 status_code) returns(int256)
func (_BaseContent *BaseContentTransactorSession) SetStatusCode(status_code *big.Int) (*types.Transaction, error) {
	return _BaseContent.Contract.SetStatusCode(&_BaseContent.TransactOpts, status_code)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseContent *BaseContentTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseContent *BaseContentSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.TransferCreatorship(&_BaseContent.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseContent *BaseContentTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.TransferCreatorship(&_BaseContent.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseContent *BaseContentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseContent *BaseContentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.TransferOwnership(&_BaseContent.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
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

// UpdateStatus is a paid mutator transaction binding the contract method 0x8280dd8f.
//
// Solidity: function updateStatus(int256 status_code) returns(int256)
func (_BaseContent *BaseContentTransactor) UpdateStatus(opts *bind.TransactOpts, status_code *big.Int) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "updateStatus", status_code)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x8280dd8f.
//
// Solidity: function updateStatus(int256 status_code) returns(int256)
func (_BaseContent *BaseContentSession) UpdateStatus(status_code *big.Int) (*types.Transaction, error) {
	return _BaseContent.Contract.UpdateStatus(&_BaseContent.TransactOpts, status_code)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x8280dd8f.
//
// Solidity: function updateStatus(int256 status_code) returns(int256)
func (_BaseContent *BaseContentTransactorSession) UpdateStatus(status_code *big.Int) (*types.Transaction, error) {
	return _BaseContent.Contract.UpdateStatus(&_BaseContent.TransactOpts, status_code)
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
// Solidity: event AccessComplete(uint256 requestID, uint256 scorePct, bytes32 mlOutHash, bool customContractResult)
func (_BaseContent *BaseContentFilterer) FilterAccessComplete(opts *bind.FilterOpts) (*BaseContentAccessCompleteIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "AccessComplete")
	if err != nil {
		return nil, err
	}
	return &BaseContentAccessCompleteIterator{contract: _BaseContent.contract, event: "AccessComplete", logs: logs, sub: sub}, nil
}

// WatchAccessComplete is a free log subscription operation binding the contract event 0x7f1f4b28434ce7beab4983e64a8b5bb96e195a67029fdaff925028aec57fbc6b.
//
// Solidity: event AccessComplete(uint256 requestID, uint256 scorePct, bytes32 mlOutHash, bool customContractResult)
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
// Solidity: event AccessGrant(uint256 requestID, bool access_granted, string reKey, string encryptedAESKey)
func (_BaseContent *BaseContentFilterer) FilterAccessGrant(opts *bind.FilterOpts) (*BaseContentAccessGrantIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "AccessGrant")
	if err != nil {
		return nil, err
	}
	return &BaseContentAccessGrantIterator{contract: _BaseContent.contract, event: "AccessGrant", logs: logs, sub: sub}, nil
}

// WatchAccessGrant is a free log subscription operation binding the contract event 0x475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718.
//
// Solidity: event AccessGrant(uint256 requestID, bool access_granted, string reKey, string encryptedAESKey)
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
	RequestID    *big.Int
	Level        uint8
	ContentHash  [32]byte
	PkeRequestor string
	PkeAFGH      string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAccessRequest is a free log retrieval operation binding the contract event 0x089a6f1788a3c353423e1be4ba12533bdde7d908bb41abeee185af0acb3df562.
//
// Solidity: event AccessRequest(uint256 requestID, uint8 level, bytes32 contentHash, string pkeRequestor, string pkeAFGH)
func (_BaseContent *BaseContentFilterer) FilterAccessRequest(opts *bind.FilterOpts) (*BaseContentAccessRequestIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "AccessRequest")
	if err != nil {
		return nil, err
	}
	return &BaseContentAccessRequestIterator{contract: _BaseContent.contract, event: "AccessRequest", logs: logs, sub: sub}, nil
}

// WatchAccessRequest is a free log subscription operation binding the contract event 0x089a6f1788a3c353423e1be4ba12533bdde7d908bb41abeee185af0acb3df562.
//
// Solidity: event AccessRequest(uint256 requestID, uint8 level, bytes32 contentHash, string pkeRequestor, string pkeAFGH)
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
// Solidity: event AccessRequestStakeholder(address stakeholder)
func (_BaseContent *BaseContentFilterer) FilterAccessRequestStakeholder(opts *bind.FilterOpts) (*BaseContentAccessRequestStakeholderIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "AccessRequestStakeholder")
	if err != nil {
		return nil, err
	}
	return &BaseContentAccessRequestStakeholderIterator{contract: _BaseContent.contract, event: "AccessRequestStakeholder", logs: logs, sub: sub}, nil
}

// WatchAccessRequestStakeholder is a free log subscription operation binding the contract event 0xb6e3239e521a6c66920ae634f8e921a37e6991d520ac44d52f8516397f41b684.
//
// Solidity: event AccessRequestStakeholder(address stakeholder)
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
// Solidity: event AccessRequestValue(bytes32 customValue)
func (_BaseContent *BaseContentFilterer) FilterAccessRequestValue(opts *bind.FilterOpts) (*BaseContentAccessRequestValueIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "AccessRequestValue")
	if err != nil {
		return nil, err
	}
	return &BaseContentAccessRequestValueIterator{contract: _BaseContent.contract, event: "AccessRequestValue", logs: logs, sub: sub}, nil
}

// WatchAccessRequestValue is a free log subscription operation binding the contract event 0x515e0a48b385fce2a8e4d9f169a97c4f6ea669a752358f5e6ab37cc3c2e84c38.
//
// Solidity: event AccessRequestValue(bytes32 customValue)
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
// Solidity: event Commit(bytes32 objectHash)
func (_BaseContent *BaseContentFilterer) FilterCommit(opts *bind.FilterOpts) (*BaseContentCommitIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "Commit")
	if err != nil {
		return nil, err
	}
	return &BaseContentCommitIterator{contract: _BaseContent.contract, event: "Commit", logs: logs, sub: sub}, nil
}

// WatchCommit is a free log subscription operation binding the contract event 0x9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff.
//
// Solidity: event Commit(bytes32 objectHash)
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
// Solidity: event ContentObjectCreate(address containingLibrary)
func (_BaseContent *BaseContentFilterer) FilterContentObjectCreate(opts *bind.FilterOpts) (*BaseContentContentObjectCreateIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "ContentObjectCreate")
	if err != nil {
		return nil, err
	}
	return &BaseContentContentObjectCreateIterator{contract: _BaseContent.contract, event: "ContentObjectCreate", logs: logs, sub: sub}, nil
}

// WatchContentObjectCreate is a free log subscription operation binding the contract event 0xc3decc188980e855666b70498ca85e8fa284d97d30483d828fa126f7303d7d19.
//
// Solidity: event ContentObjectCreate(address containingLibrary)
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
// Solidity: event GetAccessCharge(uint8 level, uint256 accessCharge)
func (_BaseContent *BaseContentFilterer) FilterGetAccessCharge(opts *bind.FilterOpts) (*BaseContentGetAccessChargeIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "GetAccessCharge")
	if err != nil {
		return nil, err
	}
	return &BaseContentGetAccessChargeIterator{contract: _BaseContent.contract, event: "GetAccessCharge", logs: logs, sub: sub}, nil
}

// WatchGetAccessCharge is a free log subscription operation binding the contract event 0xa58326ee5bb617cb8b4f0d0f5f557c469d2d05d7a738f777037deda9c724b370.
//
// Solidity: event GetAccessCharge(uint8 level, uint256 accessCharge)
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
// Solidity: event InsufficientFunds(uint256 accessCharge, uint256 amountProvided)
func (_BaseContent *BaseContentFilterer) FilterInsufficientFunds(opts *bind.FilterOpts) (*BaseContentInsufficientFundsIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "InsufficientFunds")
	if err != nil {
		return nil, err
	}
	return &BaseContentInsufficientFundsIterator{contract: _BaseContent.contract, event: "InsufficientFunds", logs: logs, sub: sub}, nil
}

// WatchInsufficientFunds is a free log subscription operation binding the contract event 0x03eb8b54a949acec2cd08fdb6d6bd4647a1f2c907d75d6900648effa92eb147f.
//
// Solidity: event InsufficientFunds(uint256 accessCharge, uint256 amountProvided)
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
// Solidity: event InvokeCustomPostHook(address custom_contract)
func (_BaseContent *BaseContentFilterer) FilterInvokeCustomPostHook(opts *bind.FilterOpts) (*BaseContentInvokeCustomPostHookIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "InvokeCustomPostHook")
	if err != nil {
		return nil, err
	}
	return &BaseContentInvokeCustomPostHookIterator{contract: _BaseContent.contract, event: "InvokeCustomPostHook", logs: logs, sub: sub}, nil
}

// WatchInvokeCustomPostHook is a free log subscription operation binding the contract event 0x97d9c9779ed3ed8b9a6edfe16d17b1fdec843245747a19abfb621806e37d4a89.
//
// Solidity: event InvokeCustomPostHook(address custom_contract)
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
// Solidity: event InvokeCustomPreHook(address custom_contract)
func (_BaseContent *BaseContentFilterer) FilterInvokeCustomPreHook(opts *bind.FilterOpts) (*BaseContentInvokeCustomPreHookIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "InvokeCustomPreHook")
	if err != nil {
		return nil, err
	}
	return &BaseContentInvokeCustomPreHookIterator{contract: _BaseContent.contract, event: "InvokeCustomPreHook", logs: logs, sub: sub}, nil
}

// WatchInvokeCustomPreHook is a free log subscription operation binding the contract event 0x12b04791b5caab768e2757268992f0c62801e3921d9e310c893f0d5f9caa5f71.
//
// Solidity: event InvokeCustomPreHook(address custom_contract)
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
	RequestStatus bool
	StatusCode    *big.Int
	ObjectHash    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPublish is a free log retrieval operation binding the contract event 0x901e6f3cdc4c61620d5d424116934b9af6e31ba79cdeaa349336d93ecfe846d4.
//
// Solidity: event Publish(bool requestStatus, int256 statusCode, bytes32 objectHash)
func (_BaseContent *BaseContentFilterer) FilterPublish(opts *bind.FilterOpts) (*BaseContentPublishIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "Publish")
	if err != nil {
		return nil, err
	}
	return &BaseContentPublishIterator{contract: _BaseContent.contract, event: "Publish", logs: logs, sub: sub}, nil
}

// WatchPublish is a free log subscription operation binding the contract event 0x901e6f3cdc4c61620d5d424116934b9af6e31ba79cdeaa349336d93ecfe846d4.
//
// Solidity: event Publish(bool requestStatus, int256 statusCode, bytes32 objectHash)
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
// Solidity: event ReturnCustomHook(address custom_contract, uint256 result)
func (_BaseContent *BaseContentFilterer) FilterReturnCustomHook(opts *bind.FilterOpts) (*BaseContentReturnCustomHookIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "ReturnCustomHook")
	if err != nil {
		return nil, err
	}
	return &BaseContentReturnCustomHookIterator{contract: _BaseContent.contract, event: "ReturnCustomHook", logs: logs, sub: sub}, nil
}

// WatchReturnCustomHook is a free log subscription operation binding the contract event 0x8c693e8b27db7caf9b9637b66dcc11444760023a4d53e95407a3acef1b249f50.
//
// Solidity: event ReturnCustomHook(address custom_contract, uint256 result)
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
// Solidity: event SetAccessCharge(uint256 accessCharge)
func (_BaseContent *BaseContentFilterer) FilterSetAccessCharge(opts *bind.FilterOpts) (*BaseContentSetAccessChargeIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "SetAccessCharge")
	if err != nil {
		return nil, err
	}
	return &BaseContentSetAccessChargeIterator{contract: _BaseContent.contract, event: "SetAccessCharge", logs: logs, sub: sub}, nil
}

// WatchSetAccessCharge is a free log subscription operation binding the contract event 0x4114f8ef80b6de2161db580cbefa14e1892d15d3ebe2062c9914e4a5773114a3.
//
// Solidity: event SetAccessCharge(uint256 accessCharge)
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
// Solidity: event SetContentContract(address contentContractAddress)
func (_BaseContent *BaseContentFilterer) FilterSetContentContract(opts *bind.FilterOpts) (*BaseContentSetContentContractIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "SetContentContract")
	if err != nil {
		return nil, err
	}
	return &BaseContentSetContentContractIterator{contract: _BaseContent.contract, event: "SetContentContract", logs: logs, sub: sub}, nil
}

// WatchSetContentContract is a free log subscription operation binding the contract event 0xa6f2e38f0cfebf27212317fced3ac40bc62e00bd33f38d69603710740c69acb7.
//
// Solidity: event SetContentContract(address contentContractAddress)
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
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetContentType is a free log retrieval operation binding the contract event 0x4f692e87baf302f7281e83eec109053efc2ca8e7bddfc6ce88c579cd9767f71f.
//
// Solidity: event SetContentType(address contentType, address contentContractAddress)
func (_BaseContent *BaseContentFilterer) FilterSetContentType(opts *bind.FilterOpts) (*BaseContentSetContentTypeIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "SetContentType")
	if err != nil {
		return nil, err
	}
	return &BaseContentSetContentTypeIterator{contract: _BaseContent.contract, event: "SetContentType", logs: logs, sub: sub}, nil
}

// WatchSetContentType is a free log subscription operation binding the contract event 0x4f692e87baf302f7281e83eec109053efc2ca8e7bddfc6ce88c579cd9767f71f.
//
// Solidity: event SetContentType(address contentType, address contentContractAddress)
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
// Solidity: event SetStatusCode(int256 statusCode)
func (_BaseContent *BaseContentFilterer) FilterSetStatusCode(opts *bind.FilterOpts) (*BaseContentSetStatusCodeIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "SetStatusCode")
	if err != nil {
		return nil, err
	}
	return &BaseContentSetStatusCodeIterator{contract: _BaseContent.contract, event: "SetStatusCode", logs: logs, sub: sub}, nil
}

// WatchSetStatusCode is a free log subscription operation binding the contract event 0xda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a2.
//
// Solidity: event SetStatusCode(int256 statusCode)
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
// Solidity: event UpdateRequest(bytes32 objectHash)
func (_BaseContent *BaseContentFilterer) FilterUpdateRequest(opts *bind.FilterOpts) (*BaseContentUpdateRequestIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return &BaseContentUpdateRequestIterator{contract: _BaseContent.contract, event: "UpdateRequest", logs: logs, sub: sub}, nil
}

// WatchUpdateRequest is a free log subscription operation binding the contract event 0x1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef19.
//
// Solidity: event UpdateRequest(bytes32 objectHash)
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
const BaseContentSpaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"address_KMS\",\"type\":\"address\"}],\"name\":\"createLibrary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content_space_description\",\"type\":\"string\"}],\"name\":\"setDescription\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createContentType\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"engageAccountLibrary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"objectHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"object_hash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accessRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"content_space_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentTypeAddress\",\"type\":\"address\"}],\"name\":\"CreateContentType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"libraryAddress\",\"type\":\"address\"}],\"name\":\"CreateLibrary\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"groupAddress\",\"type\":\"address\"}],\"name\":\"CreateGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountAddress\",\"type\":\"address\"}],\"name\":\"EngageAccountLibrary\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"Commit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"UpdateRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AccessRequest\",\"type\":\"event\"}]"

// BaseContentSpaceBin is the compiled bytecode used for deploying new contracts.
const BaseContentSpaceBin = `0x608060405234801561001057600080fd5b506040516150f63803806150f68339810160405280516000805432600160a060020a0319918216811790925560018054909116909117905501805161005c906003906020840190610063565b50506100fe565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a457805160ff19168380011785556100d1565b828001600101855582156100d1579182015b828111156100d15782518255916020019190600101906100b6565b506100dd9291506100e1565b5090565b6100fb91905b808211156100dd57600081556001016100e7565b90565b614fe98061010d6000396000f3006080604052600436106100e55763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100e757806306fdde031461011857806340b89f06146101a257806341c0e1b5146101c3578063575185ed146101d85780636d2e4b1b146101ed5780637284e4161461020e5780638da5cb5b1461022357806390c3f38f14610238578063b8cfaf0514610291578063c287e0ed146102a6578063c82710c1146102bb578063e02dd9c2146102d0578063f14fcbc8146102f7578063f15518871461030f578063f2fde38b14610338575b005b3480156100f357600080fd5b506100fc610359565b60408051600160a060020a039092168252519081900360200190f35b34801561012457600080fd5b5061012d610368565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561016757818101518382015260200161014f565b50505050905090810190601f1680156101945780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101ae57600080fd5b506100fc600160a060020a03600435166103f6565b3480156101cf57600080fd5b506100e5610476565b3480156101e457600080fd5b506100fc6104b2565b3480156101f957600080fd5b506100e5600160a060020a0360043516610522565b34801561021a57600080fd5b5061012d61057d565b34801561022f57600080fd5b506100fc6105d8565b34801561024457600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100e59436949293602493928401919081908401838280828437509497506105e79650505050505050565b34801561029d57600080fd5b506100fc61062c565b3480156102b257600080fd5b506100e5610699565b3480156102c757600080fd5b506100fc6106fe565b3480156102dc57600080fd5b506102e5610738565b60408051918252519081900360200190f35b34801561030357600080fd5b506100e560043561073e565b34801561031b57600080fd5b506103246107a7565b604080519115158252519081900360200190f35b34801561034457600080fd5b506100e5600160a060020a03600435166107d9565b600054600160a060020a031681565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103ee5780601f106103c3576101008083540402835291602001916103ee565b820191906000526020600020905b8154815290600101906020018083116103d157829003601f168201915b505050505081565b6000808261040261084b565b600160a060020a03909116815260405190819003602001906000f08015801561042f573d6000803e3d6000fd5b5060408051600160a060020a038316815290519192507f473c07a6d0228c4fb8fe2be3b4617c3b5fb7c0f8cd9ba4b67e8631844b9b6571919081900360200190a192915050565b600154600160a060020a03163214806104995750600154600160a060020a031633145b15156104a457600080fd5b600154600160a060020a0316ff5b6000806104bd61085b565b604051809103906000f0801580156104d9573d6000803e3d6000fd5b5060408051600160a060020a038316815290519192507fa3b1fe71ae61bad8cffa485b230e24e518938f76182a30fa0d9979e7237ad159919081900360200190a18091505b5090565b600054600160a060020a0316321461053957600080fd5b600160a060020a038116151561054e57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103ee5780601f106103c3576101008083540402835291602001916103ee565b600154600160a060020a031681565b600154600160a060020a031632148061060a5750600154600160a060020a031633145b151561061557600080fd5b805161062890600490602084019061086b565b5050565b6000806106376108e5565b604051809103906000f080158015610653573d6000803e3d6000fd5b5060408051600160a060020a038316815290519192507f9e69777f30c55126be256664fa7beff4b796ac32ebceab94df5071b0148017f8919081900360200190a1919050565b600154600160a060020a03163214806106bc5750600154600160a060020a031633145b15156106c757600080fd5b60025460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b6040805132815290516000917f53ce35a7383a3ea3f695bdf0f87d7e5485ba816b382673e849bfdd24e7f5e3ca919081900360200190a190565b60025481565b600154600160a060020a03163214806107615750600154600160a060020a031633145b151561076c57600080fd5b60028190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b6040516000907fed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88908290a15060015b90565b600154600160a060020a03163214806107fc5750600154600160a060020a031633145b151561080757600080fd5b600160a060020a038116151561081c57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b604051613b1d8061091083390190565b6040516107158061442d83390190565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106108ac57805160ff19168380011785556108d9565b828001600101855582156108d9579182015b828111156108d95782518255916020019190600101906108be565b5061051e9291506108f5565b60405161047c80614b4283390190565b6107d691905b8082111561051e57600081556001016108fb560060806040526000600f55604051602080613b1d833981016040525160008054600160a060020a031990811632908117835560018054831690911790556003805482163317905560088290556009829055600a91909155600e8054600160a060020a0390931692909116919091179055613aa08061007d6000396000f3006080604052600436106101b35763ffffffff60e060020a60003504166302d05d3f81146101b55780630eaec2c5146101e65780630f58a7861461021b57806316308394146102425780631b969895146102695780631cdbee5a1461028a5780631d0f4351146102ab5780631e35d8fa146102cc57806321770a84146102ed5780632393553b1461030257806329d002191461031a5780632c11f3921461033b57806332eaf21b14610353578063386493e01461036857806341c0e1b514610389578063470750bb1461039e57806349102e61146103b357806363dab9d4146103c8578063679a9a3c146103e05780636d2e4b1b1461040157806387e86b2c146104225780638cb13c2e1461048d5780638da5cb5b146104a5578063952e464b146104ba57806395a078e8146104d2578063991a3a7c146104f3578063af570c041461050b578063c287e0ed14610520578063c65bcbe214610535578063c9e8e72d1461054a578063dc3c29c01461056b578063e02dd9c21461058c578063e5538fd2146105a1578063e8de515f146105b6578063f14fcbc8146105d7578063f1551887146105ef578063f2fde38b14610604575b005b3480156101c157600080fd5b506101ca610625565b60408051600160a060020a039092168252519081900360200190f35b3480156101f257600080fd5b50610207600160a060020a0360043516610634565b604080519115158252519081900360200190f35b34801561022757600080fd5b506101b3600160a060020a036004358116906024351661073f565b34801561024e57600080fd5b50610257610820565b60408051918252519081900360200190f35b34801561027557600080fd5b506101b3600160a060020a0360043516610826565b34801561029657600080fd5b506101ca600160a060020a03600435166109b6565b3480156102b757600080fd5b506101b3600160a060020a03600435166109d1565b3480156102d857600080fd5b506101ca600160a060020a0360043516610ae4565b3480156102f957600080fd5b50610257610cff565b34801561030e57600080fd5b506101ca600435610d05565b34801561032657600080fd5b50610207600160a060020a0360043516610d2d565b34801561034757600080fd5b506101ca600435610e38565b34801561035f57600080fd5b506101ca610e46565b34801561037457600080fd5b506101b3600160a060020a0360043516610e55565b34801561039557600080fd5b506101b3610fe1565b3480156103aa57600080fd5b5061025761101d565b3480156103bf57600080fd5b50610207611023565b3480156103d457600080fd5b506101ca60043561123a565b3480156103ec57600080fd5b506101b3600160a060020a0360043516611287565b34801561040d57600080fd5b506101b3600160a060020a036004351661139a565b34801561042e57600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610207948235600160a060020a031694602480351515953695946064949201919081908401838280828437509497506113e89650505050505050565b34801561049957600080fd5b506101ca600435611708565b3480156104b157600080fd5b506101ca611716565b3480156104c657600080fd5b506101ca600435611725565b3480156104de57600080fd5b50610207600160a060020a0360043516611733565b3480156104ff57600080fd5b506101ca600435611830565b34801561051757600080fd5b506101ca61183e565b34801561052c57600080fd5b506101b361184d565b34801561054157600080fd5b506102576118b2565b34801561055657600080fd5b506101b3600160a060020a03600435166118b8565b34801561057757600080fd5b506101b3600160a060020a0360043516611908565b34801561059857600080fd5b50610257611a1b565b3480156105ad57600080fd5b50610257611a21565b3480156105c257600080fd5b506101b3600160a060020a0360043516611a27565b3480156105e357600080fd5b506101b3600435611bb3565b3480156105fb57600080fd5b50610207611c1c565b34801561061057600080fd5b506101b3600160a060020a0360043516611c7f565b600054600160a060020a031681565b6000806000806000600854600014156106505760019450610736565b5060005b60085481101561073157600480548290811061066c57fe5b600091825260209091200154600160a060020a0316935083156107295783915081600160a060020a03166395a078e8876040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b1580156106e757600080fd5b505af11580156106fb573d6000803e3d6000fd5b505050506040513d602081101561071157600080fd5b50519250600183151514156107295760019450610736565b600101610654565b600094505b50505050919050565b600154600160a060020a03163214806107625750600154600160a060020a031633145b151561076d57600080fd5b6007805460018181019092557fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688018054600160a060020a03808616600160a060020a03199283168117909355600b80549094019093556000828152600c6020908152604091829020805495871695909316851790925580519283529082019290925281517f280016f7418306a55542432120fd1a239ef9fcc1a92694d8d44ca76be0249ea7929181900390910190a15050565b600f5481565b600154600090600160a060020a031632148061084c5750600154600160a060020a031633145b151561085757600080fd5b5060005b6009548110156109b25781600160a060020a031660058281548110151561087e57fe5b600091825260209091200154600160a060020a031614156109aa5760058054829081106108a757fe5b60009182526020909120018054600160a060020a03191690556009546000190181146109635760056001600954038154811015156108e157fe5b60009182526020909120015460058054600160a060020a03909216918390811061090757fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550600560016009540381548110151561094957fe5b60009182526020909120018054600160a060020a03191690555b6009805460001901905560408051600160a060020a038416815290517fdf9d78c5635b72b709c85300a786eb7238acbe5bffe01c60c16464e45c6eb6eb9181900360200190a15b60010161085b565b5050565b600c60205260009081526040902054600160a060020a031681565b600154600160a060020a03163214806109f45750600154600160a060020a031633145b15156109ff57600080fd5b600654600a541015610a5057806006600a54815481101515610a1d57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550610a9c565b600680546001810182556000919091527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f018054600160a060020a031916600160a060020a0383161790555b600a8054600101905560408051600160a060020a038316815290517f3a94857e4393737f73edb175a7d0c195c7f635d9ae995e12740616ec55c9d4119181900360200190a150565b6000806000806000610af532610634565b1515610b0057600080fd5b600b5415610b6b5760009350600092505b600b54831015610b5f5785600160a060020a0316600784815481101515610b3457fe5b600091825260209091200154600160a060020a03161415610b5457600193505b600190920191610b11565b831515610b6b57600080fd5b85610b74611ce4565b600160a060020a03909116815260405190819003602001906000f080158015610ba1573d6000803e3d6000fd5b50600e54604080517fc9e8e72d000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015290519294508493509083169163c9e8e72d9160248082019260009290919082900301818387803b158015610c0f57600080fd5b505af1158015610c23573d6000803e3d6000fd5b50505050600160a060020a038681166000908152600c60205260408082205481517fe5385303000000000000000000000000000000000000000000000000000000008152908416600482015290519284169263e53853039260248084019391929182900301818387803b158015610c9957600080fd5b505af1158015610cad573d6000803e3d6000fd5b505060408051600160a060020a0380871682528a16602082015281517f3981e74ab81857b375ec391a4f7c31ee89462cd927de6d8fbdb98f77da009c569450908190039091019150a150949350505050565b60095481565b6004805482908110610d1357fe5b600091825260209091200154600160a060020a0316905081565b6001546000908190819081908190600160a060020a0387811691161415610d575760019450610736565b5060005b600954811015610731576005805482908110610d7357fe5b600091825260209091200154600160a060020a031693508315610e305783915081600160a060020a03166395a078e8876040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b158015610dee57600080fd5b505af1158015610e02573d6000803e3d6000fd5b505050506040513d6020811015610e1857600080fd5b5051925060018315151415610e305760019450610736565b600101610d5b565b6006805482908110610d1357fe5b600e54600160a060020a031681565b600154600090600160a060020a0316321480610e7b5750600154600160a060020a031633145b1515610e8657600080fd5b5060005b6008548110156109b25781600160a060020a0316600482815481101515610ead57fe5b600091825260209091200154600160a060020a03161415610fd9576004805482908110610ed657fe5b60009182526020909120018054600160a060020a0319169055600854600019018114610f92576004600160085403815481101515610f1057fe5b60009182526020909120015460048054600160a060020a039092169183908110610f3657fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055506004600160085403815481101515610f7857fe5b60009182526020909120018054600160a060020a03191690555b6008805460001901905560408051600160a060020a038416815290517fbbd97daa1862eb12f77ed128a557406737cee07b131b1e2d7140dff2005e197c9181900360200190a15b600101610e8a565b600154600160a060020a03163214806110045750600154600160a060020a031633145b151561100f57600080fd5b600154600160a060020a0316ff5b60085481565b6009546000903390819015156111035780600160a060020a0316638280dd8f60006040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b15801561107d57600080fd5b505af1158015611091573d6000803e3d6000fd5b505050506040513d60208110156110a757600080fd5b505060408051600160a060020a038416815260016020820152606081830181905260009082015290517f70234ce475fee4ab40e5e55cf533f67f12b47ef4c860e62dd7affa84ead4b4429181900360a00190a160019250611235565b600160a060020a0382166000908152601060205260409020541561112a5760009250611235565b600d54600f54101561117b5781600d600f5481548110151561114857fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055506111c7565b600d80546001810182556000919091527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb5018054600160a060020a031916600160a060020a0384161790555b600f8054600160a060020a0384166000818152601060209081526040918290206001948501905584549093019093558251908152329181019190915281517f0588a34cf0de4e025d359c89ca4bacbcbf175440909952d91c814412d9da996a929181900390910190a1600192505b505090565b6000600f546000148061124f575081600f5411155b1561125c57506000611282565b600d80548390811061126a57fe5b600091825260209091200154600160a060020a031690505b919050565b600154600160a060020a03163214806112aa5750600154600160a060020a031633145b15156112b557600080fd5b6004546008541015611306578060046008548154811015156112d357fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550611352565b600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b018054600160a060020a031916600160a060020a0383161790555b60088054600101905560408051600160a060020a038316815290517f218673669018c25b89bfbf1b58d0075e37c8847ef16e707b92355b7833e97d619181900360200190a150565b600054600160a060020a031632146113b157600080fd5b600160a060020a03811615156113c657600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b6000806000806000806113fa32610d2d565b151560011461140857600080fd5b600160a060020a038916600090815260106020526040902054600d805460001990920196508a9550908690811061143b57fe5b600091825260208083209091018054600160a060020a0319169055600f8054600019018155600160a060020a038c168352601090915260408220919091555485101561153057600d600f5481548110151561149257fe5b600091825260209091200154600d8054600160a060020a0390921694508491879081106114bb57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550600d600f548154811015156114fa57fe5b600091825260208083209091018054600160a060020a0319169055600160a060020a038516825260109052604090206001860190555b83600160a060020a03166327c1c21d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561156e57600080fd5b505af1158015611582573d6000803e3d6000fd5b505050506040513d602081101561159857600080fd5b5051915060008213156116f757600188151514156115b8575060006115bd565b506000195b83600160a060020a0316638280dd8f826040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b15801561160657600080fd5b505af115801561161a573d6000803e3d6000fd5b505050506040513d602081101561163057600080fd5b505060408051600160a060020a038b16815289151560208281019190915260609282018381528a519383019390935289517f70234ce475fee4ab40e5e55cf533f67f12b47ef4c860e62dd7affa84ead4b442938d938d938d93919291608084019185019080838360005b838110156116b257818101518382015260200161169a565b50505050905090810190601f1680156116df5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a1600195506116fc565b600095505b50505050509392505050565b600d805482908110610d1357fe5b600154600160a060020a031681565b6005805482908110610d1357fe5b6000806000806000600a546000141561174f5760019450610736565b5060005b600a5481101561073157600680548290811061176b57fe5b600091825260209091200154600160a060020a0316935083156118285783915081600160a060020a03166395a078e8876040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b1580156117e657600080fd5b505af11580156117fa573d6000803e3d6000fd5b505050506040513d602081101561181057600080fd5b50519250600183151514156118285760019450610736565b600101611753565b6007805482908110610d1357fe5b600354600160a060020a031681565b600154600160a060020a03163214806118705750600154600160a060020a031633145b151561187b57600080fd5b60025460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600b5481565b600154600160a060020a03163214806118db5750600154600160a060020a031633145b15156118e657600080fd5b600e8054600160a060020a031916600160a060020a0392909216919091179055565b600154600160a060020a031632148061192b5750600154600160a060020a031633145b151561193657600080fd5b60055460095410156119875780600560095481548110151561195457fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055506119d3565b600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0018054600160a060020a031916600160a060020a0383161790555b60098054600101905560408051600160a060020a038316815290517f1b88a571cc8ac2e87512f05648e79d184f5cc0cbb2889bc487c41f8b9a3202eb9181900360200190a150565b60025481565b600a5481565b600154600090600160a060020a0316321480611a4d5750600154600160a060020a031633145b1515611a5857600080fd5b5060005b600a548110156109b25781600160a060020a0316600682815481101515611a7f57fe5b600091825260209091200154600160a060020a03161415611bab576006805482908110611aa857fe5b60009182526020909120018054600160a060020a0319169055600a54600019018114611b645760066001600a5403815481101515611ae257fe5b60009182526020909120015460068054600160a060020a039092169183908110611b0857fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a0316021790555060066001600a5403815481101515611b4a57fe5b60009182526020909120018054600160a060020a03191690555b600a805460001901905560408051600160a060020a038416815290517fc5224c4118417a068eeac7d714e6d8af6f99ec3fb611bc965185460b0e38f0819181900360200190a15b600101611a5c565b600154600160a060020a0316321480611bd65750600154600160a060020a031633145b1515611be157600080fd5b60028190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b6000611c2732611733565b80611c365750611c3632610634565b80611c455750611c4532610d2d565b1515611c5057600080fd5b6040517fed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e8890600090a150600190565b600154600160a060020a0316321480611ca25750600154600160a060020a031633145b1515611cad57600080fd5b600160a060020a0381161515611cc257600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b604051611d8080611cf58339019056006080604052604051602080611d80833981016040818152915160008054600160a060020a0319908116329081178355600180548316909117905560068054821633179081905560001960085560099290925560038054600160a060020a038086169190931617905516825291517fc3decc188980e855666b70498ca85e8fa284d97d30483d828fa126f7303d7d199181900360200190a150611cda806100a66000396000f3006080604052600436106101705763ffffffff60e060020a60003504166217de98811461017257806302d05d3f14610199578063075d4782146101ca5780631a735f18146101e65780632310167f1461022d57806327c1c21d1461024257806332eaf21b1461025757806336ebffca1461026c578063388642841461028157806341c0e1b5146102995780634dd70788146102ae5780635267db44146102c35780635cc4aa9b146102db57806364ade32b146102ec5780636d2e4b1b146103015780638280dd8f14610322578063879fe48f1461033a5780638da5cb5b146103d05780638f779201146103e5578063a1ff106e146103fa578063b816f513146104fe578063c287e0ed14610513578063c9e8e72d14610528578063d810f8c814610549578063e02dd9c21461055e578063e538530314610573578063ee56d76714610594578063f14fcbc814610634578063f2fde38b1461064c578063f4d9bae81461066d578063f81ab0ae14610685575b005b34801561017e57600080fd5b5061018761069a565b60408051918252519081900360200190f35b3480156101a557600080fd5b506101ae6106be565b60408051600160a060020a039092168252519081900360200190f35b6101d26106cd565b604080519115158252519081900360200190f35b3480156101f257600080fd5b506101fe6004356107fc565b60408051600160a060020a0390941684526020840192909252600090810b900b82820152519081900360600190f35b34801561023957600080fd5b506101ae610827565b34801561024e57600080fd5b50610187610836565b34801561026357600080fd5b506101ae61083c565b34801561027857600080fd5b506101ae61084b565b34801561028d57600080fd5b5061018760043561085a565b3480156102a557600080fd5b506101706109ac565b3480156102ba57600080fd5b50610187610a91565b3480156102cf57600080fd5b50610187600435610ab5565b6101d2600435602435604435610b4a565b3480156102f857600080fd5b50610187610d0b565b34801561030d57600080fd5b50610170600160a060020a0360043516610d11565b34801561032e57600080fd5b50610187600435610d6c565b34801561034657600080fd5b5060408051602060046024803582810135848102808701860190975280865261018796843560ff1696369660449591949091019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750949750610f159650505050505050565b3480156103dc57600080fd5b506101ae6110ae565b3480156103f157600080fd5b506101876110bd565b60408051602060046024803582810135601f81018590048502860185019096528585526101d295833560ff1695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506110c39650505050505050565b34801561050a57600080fd5b506101ae6115a5565b34801561051f57600080fd5b506101706115b4565b34801561053457600080fd5b50610170600160a060020a0360043516611619565b34801561055557600080fd5b50610187611676565b34801561056a57600080fd5b5061018761169a565b34801561057f57600080fd5b50610170600160a060020a03600435166116a0565b3480156105a057600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101d2948235946024803515159536959460649492019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506118649650505050505050565b34801561064057600080fd5b50610170600435611af4565b34801561065857600080fd5b50610170600160a060020a0360043516611b5d565b34801561067957600080fd5b50610187600435611bcf565b34801561069157600080fd5b50610187611c40565b7f5075626c6973686564000000000000000000000000000000000000000000000081565b600054600160a060020a031681565b60015460009081908190600160a060020a03163214806106f75750600154600160a060020a031633145b151561070257600080fd5b61070c6001610d6c565b5060009150600060085413156107ad5750600654604080517f49102e610000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169182916349102e619160048083019260209291908290030181600087803b15801561077e57600080fd5b505af1158015610792573d6000803e3d6000fd5b505050506040513d60208110156107a857600080fd5b505191505b600854600254604080518515158152602081019390935282810191909152517f901e6f3cdc4c61620d5d424116934b9af6e31ba79cdeaa349336d93ecfe846d49181900360600190a150919050565b600a60205260009081526040812080546001820154600290920154600160a060020a03909116920b83565b600554600160a060020a031681565b60085481565b600454600160a060020a031681565b600354600160a060020a031681565b60055460009081908190600160a060020a03161561090a5750600554604080517f45080442000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a0390921691829163450804429160248083019260209291908290030181600087803b1580156108db57600080fd5b505af11580156108ef573d6000803e3d6000fd5b505050506040513d602081101561090557600080fd5b505191505b8115610918578192506109a5565b831515610947577f5075626c6973686564000000000000000000000000000000000000000000000092506109a5565b6000841215610978577f447261667400000000000000000000000000000000000000000000000000000092506109a5565b60008413156109a5577f447261667420696e20726576696577000000000000000000000000000000000092505b5050919050565b600154600090600160a060020a03163214806109d25750600154600160a060020a031633145b15156109dd57600080fd5b600554600160a060020a031615610a865750600554604080517f9e99bbea0000000000000000000000000000000000000000000000000000000081529051600160a060020a03909216918291639e99bbea9160048083019260209291908290030181600087803b158015610a5057600080fd5b505af1158015610a64573d6000803e3d6000fd5b505050506040513d6020811015610a7a57600080fd5b505115610a8657600080fd5b610a8e611c52565b50565b7f447261667400000000000000000000000000000000000000000000000000000081565b600154600090600160a060020a031632148015610ae957506000821280610ae95750600082138015610ae957506000600854125b15610af45760088290555b600654600160a060020a0316331415610b0d5760088290555b60085460408051918252517fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a29181900360200190a1505060085490565b6000838152600a60205260408120805482908190600160a060020a031615801590610b7e57508254600160a060020a031633145b1515610b8957600080fd5b60055460019250600160a060020a031615610c375750600554604080517f6b2d1324000000000000000000000000000000000000000000000000000000008152600481018990529051600160a060020a03909216918291636b2d13249160248083019260209291908290030181600087803b158015610c0757600080fd5b505af1158015610c1b573d6000803e3d6000fd5b505050506040513d6020811015610c3157600080fd5b50511591505b6002830154600090810b900b1515610c7b576001830154604051339180156108fc02916000818181858888f19350505050158015610c79573d6000803e3d6000fd5b505b6000878152600a60209081526040808320805473ffffffffffffffffffffffffffffffffffffffff1916815560018101939093556002909201805460ff191690558151898152908101889052808201879052831515606082015290517f7f1f4b28434ce7beab4983e64a8b5bb96e195a67029fdaff925028aec57fbc6b9181900360800190a15095945050505050565b60075481565b600054600160a060020a03163214610d2857600080fd5b600160a060020a0381161515610d3d57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60015460009081908190600160a060020a0316321480610d965750600154600160a060020a031633145b80610dab5750600654600160a060020a031633145b1515610db657600080fd5b600554600160a060020a03161515610e3a57600154600160a060020a0316321480610deb5750600154600160a060020a031633145b8015610e025750836000191480610e025750836001145b15610e0f57839150610e35565b600654600160a060020a031633148015610e2c5750600060085412155b15610e35578391505b610ed2565b50600554604080517f3513a805000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a03909216918291633513a8059160248083019260209291908290030181600087803b158015610ea357600080fd5b505af1158015610eb7573d6000803e3d6000fd5b505050506040513d6020811015610ecd57600080fd5b505191505b60088290556040805183815290517fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a29181900360200190a1505060085492915050565b6007546005546000919082908190600160a060020a031615611065576005546040517ff8117ca200000000000000000000000000000000000000000000000000000000815260ff891660048201908152606060248301908152895160648401528951600160a060020a039094169550859363f8117ca2938c938c938c93919290916044820191608401906020808801910280838360005b83811015610fc4578181015183820152602001610fac565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015611003578181015183820152602001610feb565b5050505090500195505050505050602060405180830381600087803b15801561102b57600080fd5b505af115801561103f573d6000803e3d6000fd5b505050506040513d602081101561105557600080fd5b5051905060008112611065578092505b6040805160ff891681526020810185905281517fa58326ee5bb617cb8b4f0d0f5f557c469d2d05d7a738f777037deda9c724b370929181900390910190a1509095945050505050565b600154600160a060020a031681565b60095481565b60008060006110d0611c8e565b600980546001019055600654604080517f95a078e80000000000000000000000000000000000000000000000000000000081523260048201529051600160a060020a0390921694506000918291829187916395a078e891602480830192602092919082900301818787803b15801561114757600080fd5b505af115801561115b573d6000803e3d6000fd5b505050506040513d602081101561117157600080fd5b5051151561117e57600080fd5b600154600160a060020a031632146111aa5761119b8c8a8a610f15565b9450348511156111aa57600080fd5b6040805160608101825233815234602080830191825260008385018181526009548252600a90925293842083518154600160a060020a0391821673ffffffffffffffffffffffffffffffffffffffff19909116178255925160018201559051600290910180549190940b60ff1660ff1990911617909255600554909550161561136357600560009054906101000a9004600160a060020a0316925082600160a060020a031663123e0e806009548e8c8c6040518563ffffffff1660e060020a028152600401808581526020018460ff1660ff1681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156112c25781810151838201526020016112aa565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156113015781810151838201526020016112e9565b505050509050019650505050505050602060405180830381600087803b15801561132a57600080fd5b505af115801561133e573d6000803e3d6000fd5b505050506040513d602081101561135457600080fd5b50519150811561136357600080fd5b7f089a6f1788a3c353423e1be4ba12533bdde7d908bb41abeee185af0acb3df5626009548d6002548e8e604051808681526020018560ff1660ff16815260200184600019166000191681526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156113ee5781810151838201526020016113d6565b50505050905090810190601f16801561141b5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561144e578181015183820152602001611436565b50505050905090810190601f16801561147b5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15060005b885181101561150b5788818151811015156114a857fe5b6020908102909101015115611503577f515e0a48b385fce2a8e4d9f169a97c4f6ea669a752358f5e6ab37cc3c2e84c3889828151811015156114e657fe5b602090810290910181015160408051918252519081900390910190a15b600101611491565b5060005b875181101561159357888181518110151561152657fe5b602090810290910101511561158b577fb6e3239e521a6c66920ae634f8e921a37e6991d520ac44d52f8516397f41b684888281518110151561156457fe5b602090810290910181015160408051600160a060020a039092168252519081900390910190a15b60010161150f565b5060019b9a5050505050505050505050565b600654600160a060020a031681565b600154600160a060020a03163214806115d75750600154600160a060020a031633145b15156115e257600080fd5b60025460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600154600160a060020a031632148061163c5750600154600160a060020a031633145b151561164757600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b7f447261667420696e20726576696577000000000000000000000000000000000081565b60025481565b60015460009081908190600160a060020a03163214806116ca5750600154600160a060020a031633145b15156116d557600080fd5b600554600160a060020a03161561177557600560009054906101000a9004600160a060020a0316925082600160a060020a0316639e99bbea6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561173c57600080fd5b505af1158015611750573d6000803e3d6000fd5b505050506040513d602081101561176657600080fd5b50519150811561177557600080fd5b6005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0386169081179091551561181f5783925082600160a060020a0316637b1cdb3e6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156117e657600080fd5b505af11580156117fa573d6000803e3d6000fd5b505050506040513d602081101561181057600080fd5b50519050801561181f57600080fd5b60055460408051600160a060020a039092168252517fa6f2e38f0cfebf27212317fced3ac40bc62e00bd33f38d69603710740c69acb79181900360200190a150505050565b6001546000908190600160a060020a031633148061188c5750600454600160a060020a031633145b151561189757600080fd5b506000858152600a602052604090208054600160a060020a031615156118bc57600080fd5b6002810154600090810b900b1515611ae85784151561197d5780546001820154604051600160a060020a039092169181156108fc0291906000818181858888f19350505050158015611912573d6000803e3d6000fd5b5060028101805460ff191660ff179055604080518781526000602082018190526080828401819052820181905260c06060830181905282015290517f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718918190036101000190a1611ae8565b6001805490820154604051600160a060020a039092169181156108fc0291906000818181858888f193505050501580156119bb573d6000803e3d6000fd5b5060028101805460ff19166001908117909155604080518881526020808201849052608092820183815288519383019390935287517f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718948b9490938a938a93919291606084019160a08501919087019080838360005b83811015611a49578181015183820152602001611a31565b50505050905090810190601f168015611a765780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611aa9578181015183820152602001611a91565b50505050905090810190601f168015611ad65780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a15b50600195945050505050565b600154600160a060020a0316321480611b175750600154600160a060020a031633145b1515611b2257600080fd5b60028190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b600154600160a060020a0316321480611b805750600154600160a060020a031633145b1515611b8b57600080fd5b600160a060020a0381161515611ba057600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600090600160a060020a0316321480611bf55750600154600160a060020a031633145b1515611c0057600080fd5b60078290556040805183815290517f4114f8ef80b6de2161db580cbefa14e1892d15d3ebe2062c9914e4a5773114a39181900360200190a1505060075490565b6000611c4d60085461085a565b905090565b600154600160a060020a0316321480611c755750600154600160a060020a031633145b1515611c8057600080fd5b600154600160a060020a0316ff5b6040805160608101825260008082526020820181905291810191909152905600a165627a7a723058209a663fbea2d3d07ef8d8b67303e675cf66c7f16a6dfe96500c87f5d31e35f0080029a165627a7a72305820cd8e2d6ba5fef7e22adb64708a33d1f4c49486b64e633418780124f7a11882370029608060405234801561001057600080fd5b5060008054600160a060020a0319908116329081178084556001805484169092178255600280549093163317909255600160a060020a039182168352600460209081526040808520805460ff199081168517909155855490941685526003909152909220805490911690911790556106888061008d6000396000f3006080604052600436106100cf5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100d157806308ae4b0c146101025780630ae5e7391461013757806341c0e1b51461015857806342e7ba7b1461016d5780636d2e4b1b1461018e57806375861a95146101af57806385e68531146101d05780638da5cb5b146101f157806395a078e814610206578063af570c0414610227578063cdb849b71461023c578063f2fde38b1461025d578063fdff9b4d1461027e575b005b3480156100dd57600080fd5b506100e661029f565b60408051600160a060020a039092168252519081900360200190f35b34801561010e57600080fd5b50610123600160a060020a03600435166102ae565b604080519115158252519081900360200190f35b34801561014357600080fd5b506100cf600160a060020a03600435166102c3565b34801561016457600080fd5b506100cf61033f565b34801561017957600080fd5b50610123600160a060020a036004351661037b565b34801561019a57600080fd5b506100cf600160a060020a036004351661039e565b3480156101bb57600080fd5b506100cf600160a060020a03600435166103f9565b3480156101dc57600080fd5b506100cf600160a060020a0360043516610482565b3480156101fd57600080fd5b506100e6610510565b34801561021257600080fd5b50610123600160a060020a036004351661051f565b34801561023357600080fd5b506100e6610542565b34801561024857600080fd5b506100cf600160a060020a0360043516610551565b34801561026957600080fd5b506100cf600160a060020a03600435166105d5565b34801561028a57600080fd5b50610123600160a060020a0360043516610647565b600054600160a060020a031681565b60036020526000908152604090205460ff1681565b3360009081526004602052604090205460ff1615156001146102e457600080fd5b600160a060020a038116600081815260036020908152604091829020805460ff19166001179055815192835290517fb251eb052afc73ffd02ffe85ad79990a8b3fed60d76dbc2fa2fdd7123dffd9149281900390910190a150565b600154600160a060020a03163214806103625750600154600160a060020a031633145b151561036d57600080fd5b600154600160a060020a0316ff5b600160a060020a031660009081526004602052604090205460ff16151560011490565b600054600160a060020a031632146103b557600080fd5b600160a060020a03811615156103ca57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031632148061041c5750600154600160a060020a031633145b151561042757600080fd5b600160a060020a038116600081815260046020908152604091829020805460ff19166001179055815192835290517f93bcaab179551bde429187645251f8e1fb8ac85801fcb1cf91eb2c9043d611179281900390910190a150565b3360009081526004602052604090205460ff161515600114806104ad575033600160a060020a038216145b15156104b857600080fd5b600160a060020a038116600081815260036020908152604091829020805460ff19169055815192835290517f745cd29407db644ed93e3ceb61cbcab96d1dfb496989ac5d5bf514fc5a9fab9c9281900390910190a150565b600154600160a060020a031681565b600160a060020a031660009081526003602052604090205460ff16151560011490565b600254600160a060020a031681565b600154600160a060020a0316331480610572575033600160a060020a038216145b151561057d57600080fd5b600160a060020a038116600081815260046020908152604091829020805460ff19169055815192835290517f2d6aa1a9629d125e23a0cf692cda7cd6795dff1652eedd4673b38ec31e387b959281900390910190a150565b600154600160a060020a03163214806105f85750600154600160a060020a031633145b151561060357600080fd5b600160a060020a038116151561061857600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60046020526000908152604090205460ff16815600a165627a7a7230582012ef4fbe6555a73830212786a66a6e16a4b3c0c313a0de46d8f4fb95ddf39afa0029608060405260008054600160a060020a031990811632908117909255600180548216909217909155600380549091163317905561043b806100416000396000f3006080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100a557806341c0e1b5146100d65780636d2e4b1b146100eb5780638da5cb5b1461010c578063af570c0414610121578063c287e0ed14610136578063e02dd9c21461014b578063f14fcbc814610172578063f15518871461018a578063f2fde38b146101b3575b005b3480156100b157600080fd5b506100ba6101d4565b60408051600160a060020a039092168252519081900360200190f35b3480156100e257600080fd5b506100a36101e3565b3480156100f757600080fd5b506100a3600160a060020a036004351661021f565b34801561011857600080fd5b506100ba61027a565b34801561012d57600080fd5b506100ba610289565b34801561014257600080fd5b506100a3610298565b34801561015757600080fd5b506101606102fd565b60408051918252519081900360200190f35b34801561017e57600080fd5b506100a3600435610303565b34801561019657600080fd5b5061019f61036c565b604080519115158252519081900360200190f35b3480156101bf57600080fd5b506100a3600160a060020a036004351661039d565b600054600160a060020a031681565b600154600160a060020a03163214806102065750600154600160a060020a031633145b151561021157600080fd5b600154600160a060020a0316ff5b600054600160a060020a0316321461023657600080fd5b600160a060020a038116151561024b57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031681565b600354600160a060020a031681565b600154600160a060020a03163214806102bb5750600154600160a060020a031633145b15156102c657600080fd5b60025460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b60025481565b600154600160a060020a03163214806103265750600154600160a060020a031633145b151561033157600080fd5b60028190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b6040516000907fed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88908290a150600190565b600154600160a060020a03163214806103c05750600154600160a060020a031633145b15156103cb57600080fd5b600160a060020a03811615156103e057600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820e2015a2830fdae8ea2d7369aa8aee94caa5d0f6b7d8c85dd3e78d2f9a854169b0029a165627a7a723058207ea13089944a4699513953e1250c90d2f4cbc1d2e5877713fc5529e99fc937510029`

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
// Solidity: function commit(bytes32 object_hash) returns()
func (_BaseContentSpace *BaseContentSpaceTransactor) Commit(opts *bind.TransactOpts, object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "commit", object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 object_hash) returns()
func (_BaseContentSpace *BaseContentSpaceSession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.Commit(&_BaseContentSpace.TransactOpts, object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 object_hash) returns()
func (_BaseContentSpace *BaseContentSpaceTransactorSession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.Commit(&_BaseContentSpace.TransactOpts, object_hash)
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
// Solidity: function createLibrary(address address_KMS) returns(address)
func (_BaseContentSpace *BaseContentSpaceTransactor) CreateLibrary(opts *bind.TransactOpts, address_KMS common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "createLibrary", address_KMS)
}

// CreateLibrary is a paid mutator transaction binding the contract method 0x40b89f06.
//
// Solidity: function createLibrary(address address_KMS) returns(address)
func (_BaseContentSpace *BaseContentSpaceSession) CreateLibrary(address_KMS common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.CreateLibrary(&_BaseContentSpace.TransactOpts, address_KMS)
}

// CreateLibrary is a paid mutator transaction binding the contract method 0x40b89f06.
//
// Solidity: function createLibrary(address address_KMS) returns(address)
func (_BaseContentSpace *BaseContentSpaceTransactorSession) CreateLibrary(address_KMS common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.CreateLibrary(&_BaseContentSpace.TransactOpts, address_KMS)
}

// EngageAccountLibrary is a paid mutator transaction binding the contract method 0xc82710c1.
//
// Solidity: function engageAccountLibrary() returns(address)
func (_BaseContentSpace *BaseContentSpaceTransactor) EngageAccountLibrary(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "engageAccountLibrary")
}

// EngageAccountLibrary is a paid mutator transaction binding the contract method 0xc82710c1.
//
// Solidity: function engageAccountLibrary() returns(address)
func (_BaseContentSpace *BaseContentSpaceSession) EngageAccountLibrary() (*types.Transaction, error) {
	return _BaseContentSpace.Contract.EngageAccountLibrary(&_BaseContentSpace.TransactOpts)
}

// EngageAccountLibrary is a paid mutator transaction binding the contract method 0xc82710c1.
//
// Solidity: function engageAccountLibrary() returns(address)
func (_BaseContentSpace *BaseContentSpaceTransactorSession) EngageAccountLibrary() (*types.Transaction, error) {
	return _BaseContentSpace.Contract.EngageAccountLibrary(&_BaseContentSpace.TransactOpts)
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
// Solidity: function setDescription(string content_space_description) returns()
func (_BaseContentSpace *BaseContentSpaceTransactor) SetDescription(opts *bind.TransactOpts, content_space_description string) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "setDescription", content_space_description)
}

// SetDescription is a paid mutator transaction binding the contract method 0x90c3f38f.
//
// Solidity: function setDescription(string content_space_description) returns()
func (_BaseContentSpace *BaseContentSpaceSession) SetDescription(content_space_description string) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.SetDescription(&_BaseContentSpace.TransactOpts, content_space_description)
}

// SetDescription is a paid mutator transaction binding the contract method 0x90c3f38f.
//
// Solidity: function setDescription(string content_space_description) returns()
func (_BaseContentSpace *BaseContentSpaceTransactorSession) SetDescription(content_space_description string) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.SetDescription(&_BaseContentSpace.TransactOpts, content_space_description)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseContentSpace *BaseContentSpaceTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseContentSpace *BaseContentSpaceSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.TransferCreatorship(&_BaseContentSpace.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseContentSpace *BaseContentSpaceTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.TransferCreatorship(&_BaseContentSpace.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseContentSpace *BaseContentSpaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseContentSpace *BaseContentSpaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.TransferOwnership(&_BaseContentSpace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
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
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAccessRequest is a free log retrieval operation binding the contract event 0xed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88.
//
// Solidity: event AccessRequest()
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterAccessRequest(opts *bind.FilterOpts) (*BaseContentSpaceAccessRequestIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "AccessRequest")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceAccessRequestIterator{contract: _BaseContentSpace.contract, event: "AccessRequest", logs: logs, sub: sub}, nil
}

// WatchAccessRequest is a free log subscription operation binding the contract event 0xed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88.
//
// Solidity: event AccessRequest()
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

// BaseContentSpaceCommitIterator is returned from FilterCommit and is used to iterate over the raw UnpackLog and unpacked data for Commit events raised by the BaseContentSpace contract.
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
// Solidity: event Commit(bytes32 objectHash)
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterCommit(opts *bind.FilterOpts) (*BaseContentSpaceCommitIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "Commit")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceCommitIterator{contract: _BaseContentSpace.contract, event: "Commit", logs: logs, sub: sub}, nil
}

// WatchCommit is a free log subscription operation binding the contract event 0x9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff.
//
// Solidity: event Commit(bytes32 objectHash)
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
// Solidity: event CreateContentType(address contentTypeAddress)
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterCreateContentType(opts *bind.FilterOpts) (*BaseContentSpaceCreateContentTypeIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "CreateContentType")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceCreateContentTypeIterator{contract: _BaseContentSpace.contract, event: "CreateContentType", logs: logs, sub: sub}, nil
}

// WatchCreateContentType is a free log subscription operation binding the contract event 0x9e69777f30c55126be256664fa7beff4b796ac32ebceab94df5071b0148017f8.
//
// Solidity: event CreateContentType(address contentTypeAddress)
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
// Solidity: event CreateGroup(address groupAddress)
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterCreateGroup(opts *bind.FilterOpts) (*BaseContentSpaceCreateGroupIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "CreateGroup")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceCreateGroupIterator{contract: _BaseContentSpace.contract, event: "CreateGroup", logs: logs, sub: sub}, nil
}

// WatchCreateGroup is a free log subscription operation binding the contract event 0xa3b1fe71ae61bad8cffa485b230e24e518938f76182a30fa0d9979e7237ad159.
//
// Solidity: event CreateGroup(address groupAddress)
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
// Solidity: event CreateLibrary(address libraryAddress)
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterCreateLibrary(opts *bind.FilterOpts) (*BaseContentSpaceCreateLibraryIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "CreateLibrary")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceCreateLibraryIterator{contract: _BaseContentSpace.contract, event: "CreateLibrary", logs: logs, sub: sub}, nil
}

// WatchCreateLibrary is a free log subscription operation binding the contract event 0x473c07a6d0228c4fb8fe2be3b4617c3b5fb7c0f8cd9ba4b67e8631844b9b6571.
//
// Solidity: event CreateLibrary(address libraryAddress)
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

// BaseContentSpaceEngageAccountLibraryIterator is returned from FilterEngageAccountLibrary and is used to iterate over the raw logs and unpacked data for EngageAccountLibrary events raised by the BaseContentSpace contract.
type BaseContentSpaceEngageAccountLibraryIterator struct {
	Event *BaseContentSpaceEngageAccountLibrary // Event containing the contract specifics and raw log

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
func (it *BaseContentSpaceEngageAccountLibraryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentSpaceEngageAccountLibrary)
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
		it.Event = new(BaseContentSpaceEngageAccountLibrary)
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
func (it *BaseContentSpaceEngageAccountLibraryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentSpaceEngageAccountLibraryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentSpaceEngageAccountLibrary represents a EngageAccountLibrary event raised by the BaseContentSpace contract.
type BaseContentSpaceEngageAccountLibrary struct {
	AccountAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEngageAccountLibrary is a free log retrieval operation binding the contract event 0x53ce35a7383a3ea3f695bdf0f87d7e5485ba816b382673e849bfdd24e7f5e3ca.
//
// Solidity: event EngageAccountLibrary(address accountAddress)
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterEngageAccountLibrary(opts *bind.FilterOpts) (*BaseContentSpaceEngageAccountLibraryIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "EngageAccountLibrary")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceEngageAccountLibraryIterator{contract: _BaseContentSpace.contract, event: "EngageAccountLibrary", logs: logs, sub: sub}, nil
}

// WatchEngageAccountLibrary is a free log subscription operation binding the contract event 0x53ce35a7383a3ea3f695bdf0f87d7e5485ba816b382673e849bfdd24e7f5e3ca.
//
// Solidity: event EngageAccountLibrary(address accountAddress)
func (_BaseContentSpace *BaseContentSpaceFilterer) WatchEngageAccountLibrary(opts *bind.WatchOpts, sink chan<- *BaseContentSpaceEngageAccountLibrary) (event.Subscription, error) {

	logs, sub, err := _BaseContentSpace.contract.WatchLogs(opts, "EngageAccountLibrary")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentSpaceEngageAccountLibrary)
				if err := _BaseContentSpace.contract.UnpackLog(event, "EngageAccountLibrary", log); err != nil {
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
// Solidity: event UpdateRequest(bytes32 objectHash)
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterUpdateRequest(opts *bind.FilterOpts) (*BaseContentSpaceUpdateRequestIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceUpdateRequestIterator{contract: _BaseContentSpace.contract, event: "UpdateRequest", logs: logs, sub: sub}, nil
}

// WatchUpdateRequest is a free log subscription operation binding the contract event 0x1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef19.
//
// Solidity: event UpdateRequest(bytes32 objectHash)
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
const BaseContentTypeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"objectHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"object_hash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accessRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"Commit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"UpdateRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AccessRequest\",\"type\":\"event\"}]"

// BaseContentTypeBin is the compiled bytecode used for deploying new contracts.
const BaseContentTypeBin = `0x608060405260008054600160a060020a031990811632908117909255600180548216909217909155600380549091163317905561043b806100416000396000f3006080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100a557806341c0e1b5146100d65780636d2e4b1b146100eb5780638da5cb5b1461010c578063af570c0414610121578063c287e0ed14610136578063e02dd9c21461014b578063f14fcbc814610172578063f15518871461018a578063f2fde38b146101b3575b005b3480156100b157600080fd5b506100ba6101d4565b60408051600160a060020a039092168252519081900360200190f35b3480156100e257600080fd5b506100a36101e3565b3480156100f757600080fd5b506100a3600160a060020a036004351661021f565b34801561011857600080fd5b506100ba61027a565b34801561012d57600080fd5b506100ba610289565b34801561014257600080fd5b506100a3610298565b34801561015757600080fd5b506101606102fd565b60408051918252519081900360200190f35b34801561017e57600080fd5b506100a3600435610303565b34801561019657600080fd5b5061019f61036c565b604080519115158252519081900360200190f35b3480156101bf57600080fd5b506100a3600160a060020a036004351661039d565b600054600160a060020a031681565b600154600160a060020a03163214806102065750600154600160a060020a031633145b151561021157600080fd5b600154600160a060020a0316ff5b600054600160a060020a0316321461023657600080fd5b600160a060020a038116151561024b57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031681565b600354600160a060020a031681565b600154600160a060020a03163214806102bb5750600154600160a060020a031633145b15156102c657600080fd5b60025460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b60025481565b600154600160a060020a03163214806103265750600154600160a060020a031633145b151561033157600080fd5b60028190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b6040516000907fed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88908290a150600190565b600154600160a060020a03163214806103c05750600154600160a060020a031633145b15156103cb57600080fd5b600160a060020a03811615156103e057600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820e2015a2830fdae8ea2d7369aa8aee94caa5d0f6b7d8c85dd3e78d2f9a854169b0029`

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
// Solidity: function commit(bytes32 object_hash) returns()
func (_BaseContentType *BaseContentTypeTransactor) Commit(opts *bind.TransactOpts, object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContentType.contract.Transact(opts, "commit", object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 object_hash) returns()
func (_BaseContentType *BaseContentTypeSession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _BaseContentType.Contract.Commit(&_BaseContentType.TransactOpts, object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 object_hash) returns()
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
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseContentType *BaseContentTypeTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _BaseContentType.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseContentType *BaseContentTypeSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseContentType.Contract.TransferCreatorship(&_BaseContentType.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseContentType *BaseContentTypeTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseContentType.Contract.TransferCreatorship(&_BaseContentType.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseContentType *BaseContentTypeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseContentType.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseContentType *BaseContentTypeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseContentType.Contract.TransferOwnership(&_BaseContentType.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
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
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAccessRequest is a free log retrieval operation binding the contract event 0xed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88.
//
// Solidity: event AccessRequest()
func (_BaseContentType *BaseContentTypeFilterer) FilterAccessRequest(opts *bind.FilterOpts) (*BaseContentTypeAccessRequestIterator, error) {

	logs, sub, err := _BaseContentType.contract.FilterLogs(opts, "AccessRequest")
	if err != nil {
		return nil, err
	}
	return &BaseContentTypeAccessRequestIterator{contract: _BaseContentType.contract, event: "AccessRequest", logs: logs, sub: sub}, nil
}

// WatchAccessRequest is a free log subscription operation binding the contract event 0xed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88.
//
// Solidity: event AccessRequest()
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
// Solidity: event Commit(bytes32 objectHash)
func (_BaseContentType *BaseContentTypeFilterer) FilterCommit(opts *bind.FilterOpts) (*BaseContentTypeCommitIterator, error) {

	logs, sub, err := _BaseContentType.contract.FilterLogs(opts, "Commit")
	if err != nil {
		return nil, err
	}
	return &BaseContentTypeCommitIterator{contract: _BaseContentType.contract, event: "Commit", logs: logs, sub: sub}, nil
}

// WatchCommit is a free log subscription operation binding the contract event 0x9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff.
//
// Solidity: event Commit(bytes32 objectHash)
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
// Solidity: event UpdateRequest(bytes32 objectHash)
func (_BaseContentType *BaseContentTypeFilterer) FilterUpdateRequest(opts *bind.FilterOpts) (*BaseContentTypeUpdateRequestIterator, error) {

	logs, sub, err := _BaseContentType.contract.FilterLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return &BaseContentTypeUpdateRequestIterator{contract: _BaseContentType.contract, event: "UpdateRequest", logs: logs, sub: sub}, nil
}

// WatchUpdateRequest is a free log subscription operation binding the contract event 0x1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef19.
//
// Solidity: event UpdateRequest(bytes32 objectHash)
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
const BaseLibraryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"canContribute\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content_type\",\"type\":\"address\"},{\"name\":\"content_contract\",\"type\":\"address\"}],\"name\":\"addContentType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"approvalRequestsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"removeReviewerGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contentTypeContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"addAccessorGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content_type\",\"type\":\"address\"}],\"name\":\"createContent\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reviewerGroupsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contributorGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"canReview\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accessorGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressKMS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"removeContributorGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contributorGroupsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"submitApprovalRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingApprovalRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"addContributorGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content_contract\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"},{\"name\":\"note\",\"type\":\"string\"}],\"name\":\"approveContent\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approvalRequests\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reviewerGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"hasAccess\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contentTypes\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentTypesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"address_KMS\",\"type\":\"address\"}],\"name\":\"setAddressKMS\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"addReviewerGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"objectHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accessorGroupsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"removeAccessorGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"object_hash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accessRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"address_KMS\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"content_type\",\"type\":\"address\"}],\"name\":\"ContentObjectCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ContributorGroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ContributorGroupRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ReviewerGroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ReviewerGroupRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"AccessorGroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"AccessorGroupRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentType\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"contentContract\",\"type\":\"address\"}],\"name\":\"ContentTypeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"operationCode\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"UnauthorizedOperation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"ApproveContentRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"note\",\"type\":\"string\"}],\"name\":\"ApproveContent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"Commit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"UpdateRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AccessRequest\",\"type\":\"event\"}]"

// BaseLibraryBin is the compiled bytecode used for deploying new contracts.
const BaseLibraryBin = `0x60806040526000600f55604051602080613b1d833981016040525160008054600160a060020a031990811632908117835560018054831690911790556003805482163317905560088290556009829055600a91909155600e8054600160a060020a0390931692909116919091179055613aa08061007d6000396000f3006080604052600436106101b35763ffffffff60e060020a60003504166302d05d3f81146101b55780630eaec2c5146101e65780630f58a7861461021b57806316308394146102425780631b969895146102695780631cdbee5a1461028a5780631d0f4351146102ab5780631e35d8fa146102cc57806321770a84146102ed5780632393553b1461030257806329d002191461031a5780632c11f3921461033b57806332eaf21b14610353578063386493e01461036857806341c0e1b514610389578063470750bb1461039e57806349102e61146103b357806363dab9d4146103c8578063679a9a3c146103e05780636d2e4b1b1461040157806387e86b2c146104225780638cb13c2e1461048d5780638da5cb5b146104a5578063952e464b146104ba57806395a078e8146104d2578063991a3a7c146104f3578063af570c041461050b578063c287e0ed14610520578063c65bcbe214610535578063c9e8e72d1461054a578063dc3c29c01461056b578063e02dd9c21461058c578063e5538fd2146105a1578063e8de515f146105b6578063f14fcbc8146105d7578063f1551887146105ef578063f2fde38b14610604575b005b3480156101c157600080fd5b506101ca610625565b60408051600160a060020a039092168252519081900360200190f35b3480156101f257600080fd5b50610207600160a060020a0360043516610634565b604080519115158252519081900360200190f35b34801561022757600080fd5b506101b3600160a060020a036004358116906024351661073f565b34801561024e57600080fd5b50610257610820565b60408051918252519081900360200190f35b34801561027557600080fd5b506101b3600160a060020a0360043516610826565b34801561029657600080fd5b506101ca600160a060020a03600435166109b6565b3480156102b757600080fd5b506101b3600160a060020a03600435166109d1565b3480156102d857600080fd5b506101ca600160a060020a0360043516610ae4565b3480156102f957600080fd5b50610257610cff565b34801561030e57600080fd5b506101ca600435610d05565b34801561032657600080fd5b50610207600160a060020a0360043516610d2d565b34801561034757600080fd5b506101ca600435610e38565b34801561035f57600080fd5b506101ca610e46565b34801561037457600080fd5b506101b3600160a060020a0360043516610e55565b34801561039557600080fd5b506101b3610fe1565b3480156103aa57600080fd5b5061025761101d565b3480156103bf57600080fd5b50610207611023565b3480156103d457600080fd5b506101ca60043561123a565b3480156103ec57600080fd5b506101b3600160a060020a0360043516611287565b34801561040d57600080fd5b506101b3600160a060020a036004351661139a565b34801561042e57600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610207948235600160a060020a031694602480351515953695946064949201919081908401838280828437509497506113e89650505050505050565b34801561049957600080fd5b506101ca600435611708565b3480156104b157600080fd5b506101ca611716565b3480156104c657600080fd5b506101ca600435611725565b3480156104de57600080fd5b50610207600160a060020a0360043516611733565b3480156104ff57600080fd5b506101ca600435611830565b34801561051757600080fd5b506101ca61183e565b34801561052c57600080fd5b506101b361184d565b34801561054157600080fd5b506102576118b2565b34801561055657600080fd5b506101b3600160a060020a03600435166118b8565b34801561057757600080fd5b506101b3600160a060020a0360043516611908565b34801561059857600080fd5b50610257611a1b565b3480156105ad57600080fd5b50610257611a21565b3480156105c257600080fd5b506101b3600160a060020a0360043516611a27565b3480156105e357600080fd5b506101b3600435611bb3565b3480156105fb57600080fd5b50610207611c1c565b34801561061057600080fd5b506101b3600160a060020a0360043516611c7f565b600054600160a060020a031681565b6000806000806000600854600014156106505760019450610736565b5060005b60085481101561073157600480548290811061066c57fe5b600091825260209091200154600160a060020a0316935083156107295783915081600160a060020a03166395a078e8876040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b1580156106e757600080fd5b505af11580156106fb573d6000803e3d6000fd5b505050506040513d602081101561071157600080fd5b50519250600183151514156107295760019450610736565b600101610654565b600094505b50505050919050565b600154600160a060020a03163214806107625750600154600160a060020a031633145b151561076d57600080fd5b6007805460018181019092557fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688018054600160a060020a03808616600160a060020a03199283168117909355600b80549094019093556000828152600c6020908152604091829020805495871695909316851790925580519283529082019290925281517f280016f7418306a55542432120fd1a239ef9fcc1a92694d8d44ca76be0249ea7929181900390910190a15050565b600f5481565b600154600090600160a060020a031632148061084c5750600154600160a060020a031633145b151561085757600080fd5b5060005b6009548110156109b25781600160a060020a031660058281548110151561087e57fe5b600091825260209091200154600160a060020a031614156109aa5760058054829081106108a757fe5b60009182526020909120018054600160a060020a03191690556009546000190181146109635760056001600954038154811015156108e157fe5b60009182526020909120015460058054600160a060020a03909216918390811061090757fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550600560016009540381548110151561094957fe5b60009182526020909120018054600160a060020a03191690555b6009805460001901905560408051600160a060020a038416815290517fdf9d78c5635b72b709c85300a786eb7238acbe5bffe01c60c16464e45c6eb6eb9181900360200190a15b60010161085b565b5050565b600c60205260009081526040902054600160a060020a031681565b600154600160a060020a03163214806109f45750600154600160a060020a031633145b15156109ff57600080fd5b600654600a541015610a5057806006600a54815481101515610a1d57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550610a9c565b600680546001810182556000919091527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f018054600160a060020a031916600160a060020a0383161790555b600a8054600101905560408051600160a060020a038316815290517f3a94857e4393737f73edb175a7d0c195c7f635d9ae995e12740616ec55c9d4119181900360200190a150565b6000806000806000610af532610634565b1515610b0057600080fd5b600b5415610b6b5760009350600092505b600b54831015610b5f5785600160a060020a0316600784815481101515610b3457fe5b600091825260209091200154600160a060020a03161415610b5457600193505b600190920191610b11565b831515610b6b57600080fd5b85610b74611ce4565b600160a060020a03909116815260405190819003602001906000f080158015610ba1573d6000803e3d6000fd5b50600e54604080517fc9e8e72d000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015290519294508493509083169163c9e8e72d9160248082019260009290919082900301818387803b158015610c0f57600080fd5b505af1158015610c23573d6000803e3d6000fd5b50505050600160a060020a038681166000908152600c60205260408082205481517fe5385303000000000000000000000000000000000000000000000000000000008152908416600482015290519284169263e53853039260248084019391929182900301818387803b158015610c9957600080fd5b505af1158015610cad573d6000803e3d6000fd5b505060408051600160a060020a0380871682528a16602082015281517f3981e74ab81857b375ec391a4f7c31ee89462cd927de6d8fbdb98f77da009c569450908190039091019150a150949350505050565b60095481565b6004805482908110610d1357fe5b600091825260209091200154600160a060020a0316905081565b6001546000908190819081908190600160a060020a0387811691161415610d575760019450610736565b5060005b600954811015610731576005805482908110610d7357fe5b600091825260209091200154600160a060020a031693508315610e305783915081600160a060020a03166395a078e8876040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b158015610dee57600080fd5b505af1158015610e02573d6000803e3d6000fd5b505050506040513d6020811015610e1857600080fd5b5051925060018315151415610e305760019450610736565b600101610d5b565b6006805482908110610d1357fe5b600e54600160a060020a031681565b600154600090600160a060020a0316321480610e7b5750600154600160a060020a031633145b1515610e8657600080fd5b5060005b6008548110156109b25781600160a060020a0316600482815481101515610ead57fe5b600091825260209091200154600160a060020a03161415610fd9576004805482908110610ed657fe5b60009182526020909120018054600160a060020a0319169055600854600019018114610f92576004600160085403815481101515610f1057fe5b60009182526020909120015460048054600160a060020a039092169183908110610f3657fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055506004600160085403815481101515610f7857fe5b60009182526020909120018054600160a060020a03191690555b6008805460001901905560408051600160a060020a038416815290517fbbd97daa1862eb12f77ed128a557406737cee07b131b1e2d7140dff2005e197c9181900360200190a15b600101610e8a565b600154600160a060020a03163214806110045750600154600160a060020a031633145b151561100f57600080fd5b600154600160a060020a0316ff5b60085481565b6009546000903390819015156111035780600160a060020a0316638280dd8f60006040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b15801561107d57600080fd5b505af1158015611091573d6000803e3d6000fd5b505050506040513d60208110156110a757600080fd5b505060408051600160a060020a038416815260016020820152606081830181905260009082015290517f70234ce475fee4ab40e5e55cf533f67f12b47ef4c860e62dd7affa84ead4b4429181900360a00190a160019250611235565b600160a060020a0382166000908152601060205260409020541561112a5760009250611235565b600d54600f54101561117b5781600d600f5481548110151561114857fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055506111c7565b600d80546001810182556000919091527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb5018054600160a060020a031916600160a060020a0384161790555b600f8054600160a060020a0384166000818152601060209081526040918290206001948501905584549093019093558251908152329181019190915281517f0588a34cf0de4e025d359c89ca4bacbcbf175440909952d91c814412d9da996a929181900390910190a1600192505b505090565b6000600f546000148061124f575081600f5411155b1561125c57506000611282565b600d80548390811061126a57fe5b600091825260209091200154600160a060020a031690505b919050565b600154600160a060020a03163214806112aa5750600154600160a060020a031633145b15156112b557600080fd5b6004546008541015611306578060046008548154811015156112d357fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550611352565b600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b018054600160a060020a031916600160a060020a0383161790555b60088054600101905560408051600160a060020a038316815290517f218673669018c25b89bfbf1b58d0075e37c8847ef16e707b92355b7833e97d619181900360200190a150565b600054600160a060020a031632146113b157600080fd5b600160a060020a03811615156113c657600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b6000806000806000806113fa32610d2d565b151560011461140857600080fd5b600160a060020a038916600090815260106020526040902054600d805460001990920196508a9550908690811061143b57fe5b600091825260208083209091018054600160a060020a0319169055600f8054600019018155600160a060020a038c168352601090915260408220919091555485101561153057600d600f5481548110151561149257fe5b600091825260209091200154600d8054600160a060020a0390921694508491879081106114bb57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550600d600f548154811015156114fa57fe5b600091825260208083209091018054600160a060020a0319169055600160a060020a038516825260109052604090206001860190555b83600160a060020a03166327c1c21d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561156e57600080fd5b505af1158015611582573d6000803e3d6000fd5b505050506040513d602081101561159857600080fd5b5051915060008213156116f757600188151514156115b8575060006115bd565b506000195b83600160a060020a0316638280dd8f826040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b15801561160657600080fd5b505af115801561161a573d6000803e3d6000fd5b505050506040513d602081101561163057600080fd5b505060408051600160a060020a038b16815289151560208281019190915260609282018381528a519383019390935289517f70234ce475fee4ab40e5e55cf533f67f12b47ef4c860e62dd7affa84ead4b442938d938d938d93919291608084019185019080838360005b838110156116b257818101518382015260200161169a565b50505050905090810190601f1680156116df5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a1600195506116fc565b600095505b50505050509392505050565b600d805482908110610d1357fe5b600154600160a060020a031681565b6005805482908110610d1357fe5b6000806000806000600a546000141561174f5760019450610736565b5060005b600a5481101561073157600680548290811061176b57fe5b600091825260209091200154600160a060020a0316935083156118285783915081600160a060020a03166395a078e8876040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b1580156117e657600080fd5b505af11580156117fa573d6000803e3d6000fd5b505050506040513d602081101561181057600080fd5b50519250600183151514156118285760019450610736565b600101611753565b6007805482908110610d1357fe5b600354600160a060020a031681565b600154600160a060020a03163214806118705750600154600160a060020a031633145b151561187b57600080fd5b60025460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600b5481565b600154600160a060020a03163214806118db5750600154600160a060020a031633145b15156118e657600080fd5b600e8054600160a060020a031916600160a060020a0392909216919091179055565b600154600160a060020a031632148061192b5750600154600160a060020a031633145b151561193657600080fd5b60055460095410156119875780600560095481548110151561195457fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055506119d3565b600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0018054600160a060020a031916600160a060020a0383161790555b60098054600101905560408051600160a060020a038316815290517f1b88a571cc8ac2e87512f05648e79d184f5cc0cbb2889bc487c41f8b9a3202eb9181900360200190a150565b60025481565b600a5481565b600154600090600160a060020a0316321480611a4d5750600154600160a060020a031633145b1515611a5857600080fd5b5060005b600a548110156109b25781600160a060020a0316600682815481101515611a7f57fe5b600091825260209091200154600160a060020a03161415611bab576006805482908110611aa857fe5b60009182526020909120018054600160a060020a0319169055600a54600019018114611b645760066001600a5403815481101515611ae257fe5b60009182526020909120015460068054600160a060020a039092169183908110611b0857fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a0316021790555060066001600a5403815481101515611b4a57fe5b60009182526020909120018054600160a060020a03191690555b600a805460001901905560408051600160a060020a038416815290517fc5224c4118417a068eeac7d714e6d8af6f99ec3fb611bc965185460b0e38f0819181900360200190a15b600101611a5c565b600154600160a060020a0316321480611bd65750600154600160a060020a031633145b1515611be157600080fd5b60028190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b6000611c2732611733565b80611c365750611c3632610634565b80611c455750611c4532610d2d565b1515611c5057600080fd5b6040517fed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e8890600090a150600190565b600154600160a060020a0316321480611ca25750600154600160a060020a031633145b1515611cad57600080fd5b600160a060020a0381161515611cc257600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b604051611d8080611cf58339019056006080604052604051602080611d80833981016040818152915160008054600160a060020a0319908116329081178355600180548316909117905560068054821633179081905560001960085560099290925560038054600160a060020a038086169190931617905516825291517fc3decc188980e855666b70498ca85e8fa284d97d30483d828fa126f7303d7d199181900360200190a150611cda806100a66000396000f3006080604052600436106101705763ffffffff60e060020a60003504166217de98811461017257806302d05d3f14610199578063075d4782146101ca5780631a735f18146101e65780632310167f1461022d57806327c1c21d1461024257806332eaf21b1461025757806336ebffca1461026c578063388642841461028157806341c0e1b5146102995780634dd70788146102ae5780635267db44146102c35780635cc4aa9b146102db57806364ade32b146102ec5780636d2e4b1b146103015780638280dd8f14610322578063879fe48f1461033a5780638da5cb5b146103d05780638f779201146103e5578063a1ff106e146103fa578063b816f513146104fe578063c287e0ed14610513578063c9e8e72d14610528578063d810f8c814610549578063e02dd9c21461055e578063e538530314610573578063ee56d76714610594578063f14fcbc814610634578063f2fde38b1461064c578063f4d9bae81461066d578063f81ab0ae14610685575b005b34801561017e57600080fd5b5061018761069a565b60408051918252519081900360200190f35b3480156101a557600080fd5b506101ae6106be565b60408051600160a060020a039092168252519081900360200190f35b6101d26106cd565b604080519115158252519081900360200190f35b3480156101f257600080fd5b506101fe6004356107fc565b60408051600160a060020a0390941684526020840192909252600090810b900b82820152519081900360600190f35b34801561023957600080fd5b506101ae610827565b34801561024e57600080fd5b50610187610836565b34801561026357600080fd5b506101ae61083c565b34801561027857600080fd5b506101ae61084b565b34801561028d57600080fd5b5061018760043561085a565b3480156102a557600080fd5b506101706109ac565b3480156102ba57600080fd5b50610187610a91565b3480156102cf57600080fd5b50610187600435610ab5565b6101d2600435602435604435610b4a565b3480156102f857600080fd5b50610187610d0b565b34801561030d57600080fd5b50610170600160a060020a0360043516610d11565b34801561032e57600080fd5b50610187600435610d6c565b34801561034657600080fd5b5060408051602060046024803582810135848102808701860190975280865261018796843560ff1696369660449591949091019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750949750610f159650505050505050565b3480156103dc57600080fd5b506101ae6110ae565b3480156103f157600080fd5b506101876110bd565b60408051602060046024803582810135601f81018590048502860185019096528585526101d295833560ff1695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506110c39650505050505050565b34801561050a57600080fd5b506101ae6115a5565b34801561051f57600080fd5b506101706115b4565b34801561053457600080fd5b50610170600160a060020a0360043516611619565b34801561055557600080fd5b50610187611676565b34801561056a57600080fd5b5061018761169a565b34801561057f57600080fd5b50610170600160a060020a03600435166116a0565b3480156105a057600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101d2948235946024803515159536959460649492019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506118649650505050505050565b34801561064057600080fd5b50610170600435611af4565b34801561065857600080fd5b50610170600160a060020a0360043516611b5d565b34801561067957600080fd5b50610187600435611bcf565b34801561069157600080fd5b50610187611c40565b7f5075626c6973686564000000000000000000000000000000000000000000000081565b600054600160a060020a031681565b60015460009081908190600160a060020a03163214806106f75750600154600160a060020a031633145b151561070257600080fd5b61070c6001610d6c565b5060009150600060085413156107ad5750600654604080517f49102e610000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169182916349102e619160048083019260209291908290030181600087803b15801561077e57600080fd5b505af1158015610792573d6000803e3d6000fd5b505050506040513d60208110156107a857600080fd5b505191505b600854600254604080518515158152602081019390935282810191909152517f901e6f3cdc4c61620d5d424116934b9af6e31ba79cdeaa349336d93ecfe846d49181900360600190a150919050565b600a60205260009081526040812080546001820154600290920154600160a060020a03909116920b83565b600554600160a060020a031681565b60085481565b600454600160a060020a031681565b600354600160a060020a031681565b60055460009081908190600160a060020a03161561090a5750600554604080517f45080442000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a0390921691829163450804429160248083019260209291908290030181600087803b1580156108db57600080fd5b505af11580156108ef573d6000803e3d6000fd5b505050506040513d602081101561090557600080fd5b505191505b8115610918578192506109a5565b831515610947577f5075626c6973686564000000000000000000000000000000000000000000000092506109a5565b6000841215610978577f447261667400000000000000000000000000000000000000000000000000000092506109a5565b60008413156109a5577f447261667420696e20726576696577000000000000000000000000000000000092505b5050919050565b600154600090600160a060020a03163214806109d25750600154600160a060020a031633145b15156109dd57600080fd5b600554600160a060020a031615610a865750600554604080517f9e99bbea0000000000000000000000000000000000000000000000000000000081529051600160a060020a03909216918291639e99bbea9160048083019260209291908290030181600087803b158015610a5057600080fd5b505af1158015610a64573d6000803e3d6000fd5b505050506040513d6020811015610a7a57600080fd5b505115610a8657600080fd5b610a8e611c52565b50565b7f447261667400000000000000000000000000000000000000000000000000000081565b600154600090600160a060020a031632148015610ae957506000821280610ae95750600082138015610ae957506000600854125b15610af45760088290555b600654600160a060020a0316331415610b0d5760088290555b60085460408051918252517fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a29181900360200190a1505060085490565b6000838152600a60205260408120805482908190600160a060020a031615801590610b7e57508254600160a060020a031633145b1515610b8957600080fd5b60055460019250600160a060020a031615610c375750600554604080517f6b2d1324000000000000000000000000000000000000000000000000000000008152600481018990529051600160a060020a03909216918291636b2d13249160248083019260209291908290030181600087803b158015610c0757600080fd5b505af1158015610c1b573d6000803e3d6000fd5b505050506040513d6020811015610c3157600080fd5b50511591505b6002830154600090810b900b1515610c7b576001830154604051339180156108fc02916000818181858888f19350505050158015610c79573d6000803e3d6000fd5b505b6000878152600a60209081526040808320805473ffffffffffffffffffffffffffffffffffffffff1916815560018101939093556002909201805460ff191690558151898152908101889052808201879052831515606082015290517f7f1f4b28434ce7beab4983e64a8b5bb96e195a67029fdaff925028aec57fbc6b9181900360800190a15095945050505050565b60075481565b600054600160a060020a03163214610d2857600080fd5b600160a060020a0381161515610d3d57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60015460009081908190600160a060020a0316321480610d965750600154600160a060020a031633145b80610dab5750600654600160a060020a031633145b1515610db657600080fd5b600554600160a060020a03161515610e3a57600154600160a060020a0316321480610deb5750600154600160a060020a031633145b8015610e025750836000191480610e025750836001145b15610e0f57839150610e35565b600654600160a060020a031633148015610e2c5750600060085412155b15610e35578391505b610ed2565b50600554604080517f3513a805000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a03909216918291633513a8059160248083019260209291908290030181600087803b158015610ea357600080fd5b505af1158015610eb7573d6000803e3d6000fd5b505050506040513d6020811015610ecd57600080fd5b505191505b60088290556040805183815290517fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a29181900360200190a1505060085492915050565b6007546005546000919082908190600160a060020a031615611065576005546040517ff8117ca200000000000000000000000000000000000000000000000000000000815260ff891660048201908152606060248301908152895160648401528951600160a060020a039094169550859363f8117ca2938c938c938c93919290916044820191608401906020808801910280838360005b83811015610fc4578181015183820152602001610fac565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015611003578181015183820152602001610feb565b5050505090500195505050505050602060405180830381600087803b15801561102b57600080fd5b505af115801561103f573d6000803e3d6000fd5b505050506040513d602081101561105557600080fd5b5051905060008112611065578092505b6040805160ff891681526020810185905281517fa58326ee5bb617cb8b4f0d0f5f557c469d2d05d7a738f777037deda9c724b370929181900390910190a1509095945050505050565b600154600160a060020a031681565b60095481565b60008060006110d0611c8e565b600980546001019055600654604080517f95a078e80000000000000000000000000000000000000000000000000000000081523260048201529051600160a060020a0390921694506000918291829187916395a078e891602480830192602092919082900301818787803b15801561114757600080fd5b505af115801561115b573d6000803e3d6000fd5b505050506040513d602081101561117157600080fd5b5051151561117e57600080fd5b600154600160a060020a031632146111aa5761119b8c8a8a610f15565b9450348511156111aa57600080fd5b6040805160608101825233815234602080830191825260008385018181526009548252600a90925293842083518154600160a060020a0391821673ffffffffffffffffffffffffffffffffffffffff19909116178255925160018201559051600290910180549190940b60ff1660ff1990911617909255600554909550161561136357600560009054906101000a9004600160a060020a0316925082600160a060020a031663123e0e806009548e8c8c6040518563ffffffff1660e060020a028152600401808581526020018460ff1660ff1681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156112c25781810151838201526020016112aa565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156113015781810151838201526020016112e9565b505050509050019650505050505050602060405180830381600087803b15801561132a57600080fd5b505af115801561133e573d6000803e3d6000fd5b505050506040513d602081101561135457600080fd5b50519150811561136357600080fd5b7f089a6f1788a3c353423e1be4ba12533bdde7d908bb41abeee185af0acb3df5626009548d6002548e8e604051808681526020018560ff1660ff16815260200184600019166000191681526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156113ee5781810151838201526020016113d6565b50505050905090810190601f16801561141b5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561144e578181015183820152602001611436565b50505050905090810190601f16801561147b5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15060005b885181101561150b5788818151811015156114a857fe5b6020908102909101015115611503577f515e0a48b385fce2a8e4d9f169a97c4f6ea669a752358f5e6ab37cc3c2e84c3889828151811015156114e657fe5b602090810290910181015160408051918252519081900390910190a15b600101611491565b5060005b875181101561159357888181518110151561152657fe5b602090810290910101511561158b577fb6e3239e521a6c66920ae634f8e921a37e6991d520ac44d52f8516397f41b684888281518110151561156457fe5b602090810290910181015160408051600160a060020a039092168252519081900390910190a15b60010161150f565b5060019b9a5050505050505050505050565b600654600160a060020a031681565b600154600160a060020a03163214806115d75750600154600160a060020a031633145b15156115e257600080fd5b60025460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600154600160a060020a031632148061163c5750600154600160a060020a031633145b151561164757600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b7f447261667420696e20726576696577000000000000000000000000000000000081565b60025481565b60015460009081908190600160a060020a03163214806116ca5750600154600160a060020a031633145b15156116d557600080fd5b600554600160a060020a03161561177557600560009054906101000a9004600160a060020a0316925082600160a060020a0316639e99bbea6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561173c57600080fd5b505af1158015611750573d6000803e3d6000fd5b505050506040513d602081101561176657600080fd5b50519150811561177557600080fd5b6005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0386169081179091551561181f5783925082600160a060020a0316637b1cdb3e6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156117e657600080fd5b505af11580156117fa573d6000803e3d6000fd5b505050506040513d602081101561181057600080fd5b50519050801561181f57600080fd5b60055460408051600160a060020a039092168252517fa6f2e38f0cfebf27212317fced3ac40bc62e00bd33f38d69603710740c69acb79181900360200190a150505050565b6001546000908190600160a060020a031633148061188c5750600454600160a060020a031633145b151561189757600080fd5b506000858152600a602052604090208054600160a060020a031615156118bc57600080fd5b6002810154600090810b900b1515611ae85784151561197d5780546001820154604051600160a060020a039092169181156108fc0291906000818181858888f19350505050158015611912573d6000803e3d6000fd5b5060028101805460ff191660ff179055604080518781526000602082018190526080828401819052820181905260c06060830181905282015290517f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718918190036101000190a1611ae8565b6001805490820154604051600160a060020a039092169181156108fc0291906000818181858888f193505050501580156119bb573d6000803e3d6000fd5b5060028101805460ff19166001908117909155604080518881526020808201849052608092820183815288519383019390935287517f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718948b9490938a938a93919291606084019160a08501919087019080838360005b83811015611a49578181015183820152602001611a31565b50505050905090810190601f168015611a765780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611aa9578181015183820152602001611a91565b50505050905090810190601f168015611ad65780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a15b50600195945050505050565b600154600160a060020a0316321480611b175750600154600160a060020a031633145b1515611b2257600080fd5b60028190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b600154600160a060020a0316321480611b805750600154600160a060020a031633145b1515611b8b57600080fd5b600160a060020a0381161515611ba057600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600090600160a060020a0316321480611bf55750600154600160a060020a031633145b1515611c0057600080fd5b60078290556040805183815290517f4114f8ef80b6de2161db580cbefa14e1892d15d3ebe2062c9914e4a5773114a39181900360200190a1505060075490565b6000611c4d60085461085a565b905090565b600154600160a060020a0316321480611c755750600154600160a060020a031633145b1515611c8057600080fd5b600154600160a060020a0316ff5b6040805160608101825260008082526020820181905291810191909152905600a165627a7a723058209a663fbea2d3d07ef8d8b67303e675cf66c7f16a6dfe96500c87f5d31e35f0080029a165627a7a72305820cd8e2d6ba5fef7e22adb64708a33d1f4c49486b64e633418780124f7a11882370029`

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
// Solidity: function accessorGroups(uint256 ) constant returns(address)
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
// Solidity: function accessorGroups(uint256 ) constant returns(address)
func (_BaseLibrary *BaseLibrarySession) AccessorGroups(arg0 *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.AccessorGroups(&_BaseLibrary.CallOpts, arg0)
}

// AccessorGroups is a free data retrieval call binding the contract method 0x2c11f392.
//
// Solidity: function accessorGroups(uint256 ) constant returns(address)
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
// Solidity: function approvalRequests(uint256 ) constant returns(address)
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
// Solidity: function approvalRequests(uint256 ) constant returns(address)
func (_BaseLibrary *BaseLibrarySession) ApprovalRequests(arg0 *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.ApprovalRequests(&_BaseLibrary.CallOpts, arg0)
}

// ApprovalRequests is a free data retrieval call binding the contract method 0x8cb13c2e.
//
// Solidity: function approvalRequests(uint256 ) constant returns(address)
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
// Solidity: function canContribute(address candidate) constant returns(bool)
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
// Solidity: function canContribute(address candidate) constant returns(bool)
func (_BaseLibrary *BaseLibrarySession) CanContribute(candidate common.Address) (bool, error) {
	return _BaseLibrary.Contract.CanContribute(&_BaseLibrary.CallOpts, candidate)
}

// CanContribute is a free data retrieval call binding the contract method 0x0eaec2c5.
//
// Solidity: function canContribute(address candidate) constant returns(bool)
func (_BaseLibrary *BaseLibraryCallerSession) CanContribute(candidate common.Address) (bool, error) {
	return _BaseLibrary.Contract.CanContribute(&_BaseLibrary.CallOpts, candidate)
}

// CanReview is a free data retrieval call binding the contract method 0x29d00219.
//
// Solidity: function canReview(address candidate) constant returns(bool)
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
// Solidity: function canReview(address candidate) constant returns(bool)
func (_BaseLibrary *BaseLibrarySession) CanReview(candidate common.Address) (bool, error) {
	return _BaseLibrary.Contract.CanReview(&_BaseLibrary.CallOpts, candidate)
}

// CanReview is a free data retrieval call binding the contract method 0x29d00219.
//
// Solidity: function canReview(address candidate) constant returns(bool)
func (_BaseLibrary *BaseLibraryCallerSession) CanReview(candidate common.Address) (bool, error) {
	return _BaseLibrary.Contract.CanReview(&_BaseLibrary.CallOpts, candidate)
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_BaseLibrary *BaseLibraryCaller) ContentSpace(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "contentSpace")
	return *ret0, err
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_BaseLibrary *BaseLibrarySession) ContentSpace() (common.Address, error) {
	return _BaseLibrary.Contract.ContentSpace(&_BaseLibrary.CallOpts)
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_BaseLibrary *BaseLibraryCallerSession) ContentSpace() (common.Address, error) {
	return _BaseLibrary.Contract.ContentSpace(&_BaseLibrary.CallOpts)
}

// ContentTypeContracts is a free data retrieval call binding the contract method 0x1cdbee5a.
//
// Solidity: function contentTypeContracts(address ) constant returns(address)
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
// Solidity: function contentTypeContracts(address ) constant returns(address)
func (_BaseLibrary *BaseLibrarySession) ContentTypeContracts(arg0 common.Address) (common.Address, error) {
	return _BaseLibrary.Contract.ContentTypeContracts(&_BaseLibrary.CallOpts, arg0)
}

// ContentTypeContracts is a free data retrieval call binding the contract method 0x1cdbee5a.
//
// Solidity: function contentTypeContracts(address ) constant returns(address)
func (_BaseLibrary *BaseLibraryCallerSession) ContentTypeContracts(arg0 common.Address) (common.Address, error) {
	return _BaseLibrary.Contract.ContentTypeContracts(&_BaseLibrary.CallOpts, arg0)
}

// ContentTypes is a free data retrieval call binding the contract method 0x991a3a7c.
//
// Solidity: function contentTypes(uint256 ) constant returns(address)
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
// Solidity: function contentTypes(uint256 ) constant returns(address)
func (_BaseLibrary *BaseLibrarySession) ContentTypes(arg0 *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.ContentTypes(&_BaseLibrary.CallOpts, arg0)
}

// ContentTypes is a free data retrieval call binding the contract method 0x991a3a7c.
//
// Solidity: function contentTypes(uint256 ) constant returns(address)
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
// Solidity: function contributorGroups(uint256 ) constant returns(address)
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
// Solidity: function contributorGroups(uint256 ) constant returns(address)
func (_BaseLibrary *BaseLibrarySession) ContributorGroups(arg0 *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.ContributorGroups(&_BaseLibrary.CallOpts, arg0)
}

// ContributorGroups is a free data retrieval call binding the contract method 0x2393553b.
//
// Solidity: function contributorGroups(uint256 ) constant returns(address)
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
// Solidity: function getPendingApprovalRequest(uint256 index) constant returns(address)
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
// Solidity: function getPendingApprovalRequest(uint256 index) constant returns(address)
func (_BaseLibrary *BaseLibrarySession) GetPendingApprovalRequest(index *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.GetPendingApprovalRequest(&_BaseLibrary.CallOpts, index)
}

// GetPendingApprovalRequest is a free data retrieval call binding the contract method 0x63dab9d4.
//
// Solidity: function getPendingApprovalRequest(uint256 index) constant returns(address)
func (_BaseLibrary *BaseLibraryCallerSession) GetPendingApprovalRequest(index *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.GetPendingApprovalRequest(&_BaseLibrary.CallOpts, index)
}

// HasAccess is a free data retrieval call binding the contract method 0x95a078e8.
//
// Solidity: function hasAccess(address candidate) constant returns(bool)
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
// Solidity: function hasAccess(address candidate) constant returns(bool)
func (_BaseLibrary *BaseLibrarySession) HasAccess(candidate common.Address) (bool, error) {
	return _BaseLibrary.Contract.HasAccess(&_BaseLibrary.CallOpts, candidate)
}

// HasAccess is a free data retrieval call binding the contract method 0x95a078e8.
//
// Solidity: function hasAccess(address candidate) constant returns(bool)
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
// Solidity: function reviewerGroups(uint256 ) constant returns(address)
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
// Solidity: function reviewerGroups(uint256 ) constant returns(address)
func (_BaseLibrary *BaseLibrarySession) ReviewerGroups(arg0 *big.Int) (common.Address, error) {
	return _BaseLibrary.Contract.ReviewerGroups(&_BaseLibrary.CallOpts, arg0)
}

// ReviewerGroups is a free data retrieval call binding the contract method 0x952e464b.
//
// Solidity: function reviewerGroups(uint256 ) constant returns(address)
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
// Solidity: function addAccessorGroup(address group) returns()
func (_BaseLibrary *BaseLibraryTransactor) AddAccessorGroup(opts *bind.TransactOpts, group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "addAccessorGroup", group)
}

// AddAccessorGroup is a paid mutator transaction binding the contract method 0x1d0f4351.
//
// Solidity: function addAccessorGroup(address group) returns()
func (_BaseLibrary *BaseLibrarySession) AddAccessorGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.AddAccessorGroup(&_BaseLibrary.TransactOpts, group)
}

// AddAccessorGroup is a paid mutator transaction binding the contract method 0x1d0f4351.
//
// Solidity: function addAccessorGroup(address group) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) AddAccessorGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.AddAccessorGroup(&_BaseLibrary.TransactOpts, group)
}

// AddContentType is a paid mutator transaction binding the contract method 0x0f58a786.
//
// Solidity: function addContentType(address content_type, address content_contract) returns()
func (_BaseLibrary *BaseLibraryTransactor) AddContentType(opts *bind.TransactOpts, content_type common.Address, content_contract common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "addContentType", content_type, content_contract)
}

// AddContentType is a paid mutator transaction binding the contract method 0x0f58a786.
//
// Solidity: function addContentType(address content_type, address content_contract) returns()
func (_BaseLibrary *BaseLibrarySession) AddContentType(content_type common.Address, content_contract common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.AddContentType(&_BaseLibrary.TransactOpts, content_type, content_contract)
}

// AddContentType is a paid mutator transaction binding the contract method 0x0f58a786.
//
// Solidity: function addContentType(address content_type, address content_contract) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) AddContentType(content_type common.Address, content_contract common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.AddContentType(&_BaseLibrary.TransactOpts, content_type, content_contract)
}

// AddContributorGroup is a paid mutator transaction binding the contract method 0x679a9a3c.
//
// Solidity: function addContributorGroup(address group) returns()
func (_BaseLibrary *BaseLibraryTransactor) AddContributorGroup(opts *bind.TransactOpts, group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "addContributorGroup", group)
}

// AddContributorGroup is a paid mutator transaction binding the contract method 0x679a9a3c.
//
// Solidity: function addContributorGroup(address group) returns()
func (_BaseLibrary *BaseLibrarySession) AddContributorGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.AddContributorGroup(&_BaseLibrary.TransactOpts, group)
}

// AddContributorGroup is a paid mutator transaction binding the contract method 0x679a9a3c.
//
// Solidity: function addContributorGroup(address group) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) AddContributorGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.AddContributorGroup(&_BaseLibrary.TransactOpts, group)
}

// AddReviewerGroup is a paid mutator transaction binding the contract method 0xdc3c29c0.
//
// Solidity: function addReviewerGroup(address group) returns()
func (_BaseLibrary *BaseLibraryTransactor) AddReviewerGroup(opts *bind.TransactOpts, group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "addReviewerGroup", group)
}

// AddReviewerGroup is a paid mutator transaction binding the contract method 0xdc3c29c0.
//
// Solidity: function addReviewerGroup(address group) returns()
func (_BaseLibrary *BaseLibrarySession) AddReviewerGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.AddReviewerGroup(&_BaseLibrary.TransactOpts, group)
}

// AddReviewerGroup is a paid mutator transaction binding the contract method 0xdc3c29c0.
//
// Solidity: function addReviewerGroup(address group) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) AddReviewerGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.AddReviewerGroup(&_BaseLibrary.TransactOpts, group)
}

// ApproveContent is a paid mutator transaction binding the contract method 0x87e86b2c.
//
// Solidity: function approveContent(address content_contract, bool approved, string note) returns(bool)
func (_BaseLibrary *BaseLibraryTransactor) ApproveContent(opts *bind.TransactOpts, content_contract common.Address, approved bool, note string) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "approveContent", content_contract, approved, note)
}

// ApproveContent is a paid mutator transaction binding the contract method 0x87e86b2c.
//
// Solidity: function approveContent(address content_contract, bool approved, string note) returns(bool)
func (_BaseLibrary *BaseLibrarySession) ApproveContent(content_contract common.Address, approved bool, note string) (*types.Transaction, error) {
	return _BaseLibrary.Contract.ApproveContent(&_BaseLibrary.TransactOpts, content_contract, approved, note)
}

// ApproveContent is a paid mutator transaction binding the contract method 0x87e86b2c.
//
// Solidity: function approveContent(address content_contract, bool approved, string note) returns(bool)
func (_BaseLibrary *BaseLibraryTransactorSession) ApproveContent(content_contract common.Address, approved bool, note string) (*types.Transaction, error) {
	return _BaseLibrary.Contract.ApproveContent(&_BaseLibrary.TransactOpts, content_contract, approved, note)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 object_hash) returns()
func (_BaseLibrary *BaseLibraryTransactor) Commit(opts *bind.TransactOpts, object_hash [32]byte) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "commit", object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 object_hash) returns()
func (_BaseLibrary *BaseLibrarySession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _BaseLibrary.Contract.Commit(&_BaseLibrary.TransactOpts, object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 object_hash) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _BaseLibrary.Contract.Commit(&_BaseLibrary.TransactOpts, object_hash)
}

// CreateContent is a paid mutator transaction binding the contract method 0x1e35d8fa.
//
// Solidity: function createContent(address content_type) returns(address)
func (_BaseLibrary *BaseLibraryTransactor) CreateContent(opts *bind.TransactOpts, content_type common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "createContent", content_type)
}

// CreateContent is a paid mutator transaction binding the contract method 0x1e35d8fa.
//
// Solidity: function createContent(address content_type) returns(address)
func (_BaseLibrary *BaseLibrarySession) CreateContent(content_type common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.CreateContent(&_BaseLibrary.TransactOpts, content_type)
}

// CreateContent is a paid mutator transaction binding the contract method 0x1e35d8fa.
//
// Solidity: function createContent(address content_type) returns(address)
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

// RemoveAccessorGroup is a paid mutator transaction binding the contract method 0xe8de515f.
//
// Solidity: function removeAccessorGroup(address group) returns()
func (_BaseLibrary *BaseLibraryTransactor) RemoveAccessorGroup(opts *bind.TransactOpts, group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "removeAccessorGroup", group)
}

// RemoveAccessorGroup is a paid mutator transaction binding the contract method 0xe8de515f.
//
// Solidity: function removeAccessorGroup(address group) returns()
func (_BaseLibrary *BaseLibrarySession) RemoveAccessorGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.RemoveAccessorGroup(&_BaseLibrary.TransactOpts, group)
}

// RemoveAccessorGroup is a paid mutator transaction binding the contract method 0xe8de515f.
//
// Solidity: function removeAccessorGroup(address group) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) RemoveAccessorGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.RemoveAccessorGroup(&_BaseLibrary.TransactOpts, group)
}

// RemoveContributorGroup is a paid mutator transaction binding the contract method 0x386493e0.
//
// Solidity: function removeContributorGroup(address group) returns()
func (_BaseLibrary *BaseLibraryTransactor) RemoveContributorGroup(opts *bind.TransactOpts, group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "removeContributorGroup", group)
}

// RemoveContributorGroup is a paid mutator transaction binding the contract method 0x386493e0.
//
// Solidity: function removeContributorGroup(address group) returns()
func (_BaseLibrary *BaseLibrarySession) RemoveContributorGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.RemoveContributorGroup(&_BaseLibrary.TransactOpts, group)
}

// RemoveContributorGroup is a paid mutator transaction binding the contract method 0x386493e0.
//
// Solidity: function removeContributorGroup(address group) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) RemoveContributorGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.RemoveContributorGroup(&_BaseLibrary.TransactOpts, group)
}

// RemoveReviewerGroup is a paid mutator transaction binding the contract method 0x1b969895.
//
// Solidity: function removeReviewerGroup(address group) returns()
func (_BaseLibrary *BaseLibraryTransactor) RemoveReviewerGroup(opts *bind.TransactOpts, group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "removeReviewerGroup", group)
}

// RemoveReviewerGroup is a paid mutator transaction binding the contract method 0x1b969895.
//
// Solidity: function removeReviewerGroup(address group) returns()
func (_BaseLibrary *BaseLibrarySession) RemoveReviewerGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.RemoveReviewerGroup(&_BaseLibrary.TransactOpts, group)
}

// RemoveReviewerGroup is a paid mutator transaction binding the contract method 0x1b969895.
//
// Solidity: function removeReviewerGroup(address group) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) RemoveReviewerGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.RemoveReviewerGroup(&_BaseLibrary.TransactOpts, group)
}

// SetAddressKMS is a paid mutator transaction binding the contract method 0xc9e8e72d.
//
// Solidity: function setAddressKMS(address address_KMS) returns()
func (_BaseLibrary *BaseLibraryTransactor) SetAddressKMS(opts *bind.TransactOpts, address_KMS common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "setAddressKMS", address_KMS)
}

// SetAddressKMS is a paid mutator transaction binding the contract method 0xc9e8e72d.
//
// Solidity: function setAddressKMS(address address_KMS) returns()
func (_BaseLibrary *BaseLibrarySession) SetAddressKMS(address_KMS common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.SetAddressKMS(&_BaseLibrary.TransactOpts, address_KMS)
}

// SetAddressKMS is a paid mutator transaction binding the contract method 0xc9e8e72d.
//
// Solidity: function setAddressKMS(address address_KMS) returns()
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
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseLibrary *BaseLibraryTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseLibrary *BaseLibrarySession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.TransferCreatorship(&_BaseLibrary.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseLibrary *BaseLibraryTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.TransferCreatorship(&_BaseLibrary.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseLibrary *BaseLibraryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseLibrary *BaseLibrarySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.TransferOwnership(&_BaseLibrary.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
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
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAccessRequest is a free log retrieval operation binding the contract event 0xed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88.
//
// Solidity: event AccessRequest()
func (_BaseLibrary *BaseLibraryFilterer) FilterAccessRequest(opts *bind.FilterOpts) (*BaseLibraryAccessRequestIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "AccessRequest")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryAccessRequestIterator{contract: _BaseLibrary.contract, event: "AccessRequest", logs: logs, sub: sub}, nil
}

// WatchAccessRequest is a free log subscription operation binding the contract event 0xed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88.
//
// Solidity: event AccessRequest()
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
// Solidity: event AccessorGroupAdded(address group)
func (_BaseLibrary *BaseLibraryFilterer) FilterAccessorGroupAdded(opts *bind.FilterOpts) (*BaseLibraryAccessorGroupAddedIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "AccessorGroupAdded")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryAccessorGroupAddedIterator{contract: _BaseLibrary.contract, event: "AccessorGroupAdded", logs: logs, sub: sub}, nil
}

// WatchAccessorGroupAdded is a free log subscription operation binding the contract event 0x3a94857e4393737f73edb175a7d0c195c7f635d9ae995e12740616ec55c9d411.
//
// Solidity: event AccessorGroupAdded(address group)
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

// BaseLibraryAccessorGroupRemovedIterator is returned from FilterAccessorGroupRemoved and is used to iterate over the raw logs and unpacked data for AccessorGroupRemoved events raised by the BaseLibrary contract.
type BaseLibraryAccessorGroupRemovedIterator struct {
	Event *BaseLibraryAccessorGroupRemoved // Event containing the contract specifics and raw log

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
func (it *BaseLibraryAccessorGroupRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryAccessorGroupRemoved)
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
		it.Event = new(BaseLibraryAccessorGroupRemoved)
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
func (it *BaseLibraryAccessorGroupRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryAccessorGroupRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryAccessorGroupRemoved represents a AccessorGroupRemoved event raised by the BaseLibrary contract.
type BaseLibraryAccessorGroupRemoved struct {
	Group common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAccessorGroupRemoved is a free log retrieval operation binding the contract event 0xc5224c4118417a068eeac7d714e6d8af6f99ec3fb611bc965185460b0e38f081.
//
// Solidity: event AccessorGroupRemoved(address group)
func (_BaseLibrary *BaseLibraryFilterer) FilterAccessorGroupRemoved(opts *bind.FilterOpts) (*BaseLibraryAccessorGroupRemovedIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "AccessorGroupRemoved")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryAccessorGroupRemovedIterator{contract: _BaseLibrary.contract, event: "AccessorGroupRemoved", logs: logs, sub: sub}, nil
}

// WatchAccessorGroupRemoved is a free log subscription operation binding the contract event 0xc5224c4118417a068eeac7d714e6d8af6f99ec3fb611bc965185460b0e38f081.
//
// Solidity: event AccessorGroupRemoved(address group)
func (_BaseLibrary *BaseLibraryFilterer) WatchAccessorGroupRemoved(opts *bind.WatchOpts, sink chan<- *BaseLibraryAccessorGroupRemoved) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "AccessorGroupRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryAccessorGroupRemoved)
				if err := _BaseLibrary.contract.UnpackLog(event, "AccessorGroupRemoved", log); err != nil {
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

// BaseLibraryApproveContentIterator is returned from FilterApproveContent and is used to iterate over the raw logs and unpacked data for ApproveContent events raised by the BaseLibrary contract.
type BaseLibraryApproveContentIterator struct {
	Event *BaseLibraryApproveContent // Event containing the contract specifics and raw log

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
func (it *BaseLibraryApproveContentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryApproveContent)
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
		it.Event = new(BaseLibraryApproveContent)
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
func (it *BaseLibraryApproveContentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryApproveContentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryApproveContent represents a ApproveContent event raised by the BaseLibrary contract.
type BaseLibraryApproveContent struct {
	ContentAddress common.Address
	Approved       bool
	Note           string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterApproveContent is a free log retrieval operation binding the contract event 0x70234ce475fee4ab40e5e55cf533f67f12b47ef4c860e62dd7affa84ead4b442.
//
// Solidity: event ApproveContent(address contentAddress, bool approved, string note)
func (_BaseLibrary *BaseLibraryFilterer) FilterApproveContent(opts *bind.FilterOpts) (*BaseLibraryApproveContentIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "ApproveContent")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryApproveContentIterator{contract: _BaseLibrary.contract, event: "ApproveContent", logs: logs, sub: sub}, nil
}

// WatchApproveContent is a free log subscription operation binding the contract event 0x70234ce475fee4ab40e5e55cf533f67f12b47ef4c860e62dd7affa84ead4b442.
//
// Solidity: event ApproveContent(address contentAddress, bool approved, string note)
func (_BaseLibrary *BaseLibraryFilterer) WatchApproveContent(opts *bind.WatchOpts, sink chan<- *BaseLibraryApproveContent) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "ApproveContent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryApproveContent)
				if err := _BaseLibrary.contract.UnpackLog(event, "ApproveContent", log); err != nil {
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
// Solidity: event ApproveContentRequest(address contentAddress, address submitter)
func (_BaseLibrary *BaseLibraryFilterer) FilterApproveContentRequest(opts *bind.FilterOpts) (*BaseLibraryApproveContentRequestIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "ApproveContentRequest")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryApproveContentRequestIterator{contract: _BaseLibrary.contract, event: "ApproveContentRequest", logs: logs, sub: sub}, nil
}

// WatchApproveContentRequest is a free log subscription operation binding the contract event 0x0588a34cf0de4e025d359c89ca4bacbcbf175440909952d91c814412d9da996a.
//
// Solidity: event ApproveContentRequest(address contentAddress, address submitter)
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
// Solidity: event Commit(bytes32 objectHash)
func (_BaseLibrary *BaseLibraryFilterer) FilterCommit(opts *bind.FilterOpts) (*BaseLibraryCommitIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "Commit")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryCommitIterator{contract: _BaseLibrary.contract, event: "Commit", logs: logs, sub: sub}, nil
}

// WatchCommit is a free log subscription operation binding the contract event 0x9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff.
//
// Solidity: event Commit(bytes32 objectHash)
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
// Solidity: event ContentObjectCreated(address contentAddress, address content_type)
func (_BaseLibrary *BaseLibraryFilterer) FilterContentObjectCreated(opts *bind.FilterOpts) (*BaseLibraryContentObjectCreatedIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "ContentObjectCreated")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryContentObjectCreatedIterator{contract: _BaseLibrary.contract, event: "ContentObjectCreated", logs: logs, sub: sub}, nil
}

// WatchContentObjectCreated is a free log subscription operation binding the contract event 0x3981e74ab81857b375ec391a4f7c31ee89462cd927de6d8fbdb98f77da009c56.
//
// Solidity: event ContentObjectCreated(address contentAddress, address content_type)
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
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContentTypeAdded is a free log retrieval operation binding the contract event 0x280016f7418306a55542432120fd1a239ef9fcc1a92694d8d44ca76be0249ea7.
//
// Solidity: event ContentTypeAdded(address contentType, address contentContract)
func (_BaseLibrary *BaseLibraryFilterer) FilterContentTypeAdded(opts *bind.FilterOpts) (*BaseLibraryContentTypeAddedIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "ContentTypeAdded")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryContentTypeAddedIterator{contract: _BaseLibrary.contract, event: "ContentTypeAdded", logs: logs, sub: sub}, nil
}

// WatchContentTypeAdded is a free log subscription operation binding the contract event 0x280016f7418306a55542432120fd1a239ef9fcc1a92694d8d44ca76be0249ea7.
//
// Solidity: event ContentTypeAdded(address contentType, address contentContract)
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
// Solidity: event ContributorGroupAdded(address group)
func (_BaseLibrary *BaseLibraryFilterer) FilterContributorGroupAdded(opts *bind.FilterOpts) (*BaseLibraryContributorGroupAddedIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "ContributorGroupAdded")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryContributorGroupAddedIterator{contract: _BaseLibrary.contract, event: "ContributorGroupAdded", logs: logs, sub: sub}, nil
}

// WatchContributorGroupAdded is a free log subscription operation binding the contract event 0x218673669018c25b89bfbf1b58d0075e37c8847ef16e707b92355b7833e97d61.
//
// Solidity: event ContributorGroupAdded(address group)
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

// BaseLibraryContributorGroupRemovedIterator is returned from FilterContributorGroupRemoved and is used to iterate over the raw logs and unpacked data for ContributorGroupRemoved events raised by the BaseLibrary contract.
type BaseLibraryContributorGroupRemovedIterator struct {
	Event *BaseLibraryContributorGroupRemoved // Event containing the contract specifics and raw log

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
func (it *BaseLibraryContributorGroupRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryContributorGroupRemoved)
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
		it.Event = new(BaseLibraryContributorGroupRemoved)
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
func (it *BaseLibraryContributorGroupRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryContributorGroupRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryContributorGroupRemoved represents a ContributorGroupRemoved event raised by the BaseLibrary contract.
type BaseLibraryContributorGroupRemoved struct {
	Group common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterContributorGroupRemoved is a free log retrieval operation binding the contract event 0xbbd97daa1862eb12f77ed128a557406737cee07b131b1e2d7140dff2005e197c.
//
// Solidity: event ContributorGroupRemoved(address group)
func (_BaseLibrary *BaseLibraryFilterer) FilterContributorGroupRemoved(opts *bind.FilterOpts) (*BaseLibraryContributorGroupRemovedIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "ContributorGroupRemoved")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryContributorGroupRemovedIterator{contract: _BaseLibrary.contract, event: "ContributorGroupRemoved", logs: logs, sub: sub}, nil
}

// WatchContributorGroupRemoved is a free log subscription operation binding the contract event 0xbbd97daa1862eb12f77ed128a557406737cee07b131b1e2d7140dff2005e197c.
//
// Solidity: event ContributorGroupRemoved(address group)
func (_BaseLibrary *BaseLibraryFilterer) WatchContributorGroupRemoved(opts *bind.WatchOpts, sink chan<- *BaseLibraryContributorGroupRemoved) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "ContributorGroupRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryContributorGroupRemoved)
				if err := _BaseLibrary.contract.UnpackLog(event, "ContributorGroupRemoved", log); err != nil {
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
// Solidity: event ReviewerGroupAdded(address group)
func (_BaseLibrary *BaseLibraryFilterer) FilterReviewerGroupAdded(opts *bind.FilterOpts) (*BaseLibraryReviewerGroupAddedIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "ReviewerGroupAdded")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryReviewerGroupAddedIterator{contract: _BaseLibrary.contract, event: "ReviewerGroupAdded", logs: logs, sub: sub}, nil
}

// WatchReviewerGroupAdded is a free log subscription operation binding the contract event 0x1b88a571cc8ac2e87512f05648e79d184f5cc0cbb2889bc487c41f8b9a3202eb.
//
// Solidity: event ReviewerGroupAdded(address group)
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

// BaseLibraryReviewerGroupRemovedIterator is returned from FilterReviewerGroupRemoved and is used to iterate over the raw logs and unpacked data for ReviewerGroupRemoved events raised by the BaseLibrary contract.
type BaseLibraryReviewerGroupRemovedIterator struct {
	Event *BaseLibraryReviewerGroupRemoved // Event containing the contract specifics and raw log

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
func (it *BaseLibraryReviewerGroupRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryReviewerGroupRemoved)
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
		it.Event = new(BaseLibraryReviewerGroupRemoved)
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
func (it *BaseLibraryReviewerGroupRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryReviewerGroupRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryReviewerGroupRemoved represents a ReviewerGroupRemoved event raised by the BaseLibrary contract.
type BaseLibraryReviewerGroupRemoved struct {
	Group common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReviewerGroupRemoved is a free log retrieval operation binding the contract event 0xdf9d78c5635b72b709c85300a786eb7238acbe5bffe01c60c16464e45c6eb6eb.
//
// Solidity: event ReviewerGroupRemoved(address group)
func (_BaseLibrary *BaseLibraryFilterer) FilterReviewerGroupRemoved(opts *bind.FilterOpts) (*BaseLibraryReviewerGroupRemovedIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "ReviewerGroupRemoved")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryReviewerGroupRemovedIterator{contract: _BaseLibrary.contract, event: "ReviewerGroupRemoved", logs: logs, sub: sub}, nil
}

// WatchReviewerGroupRemoved is a free log subscription operation binding the contract event 0xdf9d78c5635b72b709c85300a786eb7238acbe5bffe01c60c16464e45c6eb6eb.
//
// Solidity: event ReviewerGroupRemoved(address group)
func (_BaseLibrary *BaseLibraryFilterer) WatchReviewerGroupRemoved(opts *bind.WatchOpts, sink chan<- *BaseLibraryReviewerGroupRemoved) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "ReviewerGroupRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryReviewerGroupRemoved)
				if err := _BaseLibrary.contract.UnpackLog(event, "ReviewerGroupRemoved", log); err != nil {
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
// Solidity: event UnauthorizedOperation(uint256 operationCode, address candidate)
func (_BaseLibrary *BaseLibraryFilterer) FilterUnauthorizedOperation(opts *bind.FilterOpts) (*BaseLibraryUnauthorizedOperationIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "UnauthorizedOperation")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryUnauthorizedOperationIterator{contract: _BaseLibrary.contract, event: "UnauthorizedOperation", logs: logs, sub: sub}, nil
}

// WatchUnauthorizedOperation is a free log subscription operation binding the contract event 0x23de2adc3e22f171f66b3e5a333e17feb9dc30ba9570933bd259cb6c13ef7ab7.
//
// Solidity: event UnauthorizedOperation(uint256 operationCode, address candidate)
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
// Solidity: event UpdateRequest(bytes32 objectHash)
func (_BaseLibrary *BaseLibraryFilterer) FilterUpdateRequest(opts *bind.FilterOpts) (*BaseLibraryUpdateRequestIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryUpdateRequestIterator{contract: _BaseLibrary.contract, event: "UpdateRequest", logs: logs, sub: sub}, nil
}

// WatchUpdateRequest is a free log subscription operation binding the contract event 0x1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef19.
//
// Solidity: event UpdateRequest(bytes32 objectHash)
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
const ContentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"address[]\"}],\"name\":\"runAccess\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proposed_status_code\",\"type\":\"int256\"}],\"name\":\"runStatusChange\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"name\":\"runDescribeStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"runFinalize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"runCreate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"runKill\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"address[]\"}],\"name\":\"runAccessCharge\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"s\",\"type\":\"string\"}],\"name\":\"DbgString\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"a\",\"type\":\"address\"}],\"name\":\"DbgAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"u\",\"type\":\"uint256\"}],\"name\":\"DbgUint256\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"u\",\"type\":\"uint256\"}],\"name\":\"DbgUint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"RunCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"RunKill\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"proposedStatusCode\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"returnStatusCode\",\"type\":\"int256\"}],\"name\":\"RunStatusChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"calculateAccessCharge\",\"type\":\"int256\"}],\"name\":\"RunAccessCharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"RunAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"RunFinalize\",\"type\":\"event\"}]"

// ContentBin is the compiled bytecode used for deploying new contracts.
const ContentBin = `0x60806040526000805432600160a060020a0319918216811783556001805490921617905561042b90819061003390396000f3006080604052600436106100b95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100bb578063123e0e80146100ec5780633513a8051461018b57806341c0e1b51461019657806345080442146101ab5780636b2d1324146101c35780636d2e4b1b146101ce5780637b1cdb3e146101ef5780638da5cb5b146101f75780639e99bbea146101ef578063f2fde38b1461020c578063f8117ca21461022d575b005b3480156100c757600080fd5b506100d06102b6565b60408051600160a060020a039092168252519081900360200190f35b6040805160206004604435818101358381028086018501909652808552610179958335956024803560ff1696369695606495939492019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506102c59650505050505050565b60408051918252519081900360200190f35b6101796004356102cf565b3480156101a257600080fd5b506100b96102d2565b3480156101b757600080fd5b5061017960043561030e565b61017960043561030e565b3480156101da57600080fd5b506100b9600160a060020a0360043516610314565b61017961036f565b34801561020357600080fd5b506100d0610374565b34801561021857600080fd5b506100b9600160a060020a0360043516610383565b60408051602060046024803582810135848102808701860190975280865261017996843560ff1696369660449591949091019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506103f59650505050505050565b600054600160a060020a031681565b6000949350505050565b90565b600154600160a060020a03163214806102f55750600154600160a060020a031633145b151561030057600080fd5b600154600160a060020a0316ff5b50600090565b600054600160a060020a0316321461032b57600080fd5b600160a060020a038116151561034057600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600090565b600154600160a060020a031681565b600154600160a060020a03163214806103a65750600154600160a060020a031633145b15156103b157600080fd5b600160a060020a03811615156103c657600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60001993925050505600a165627a7a723058201d3fa5325658e0c82b76b8baf1321fc6ab9a0b25c4db5683b58e37d1460735bb0029`

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

// RunDescribeStatus is a free data retrieval call binding the contract method 0x45080442.
//
// Solidity: function runDescribeStatus(int256 ) constant returns(bytes32)
func (_Content *ContentCaller) RunDescribeStatus(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Content.contract.Call(opts, out, "runDescribeStatus", arg0)
	return *ret0, err
}

// RunDescribeStatus is a free data retrieval call binding the contract method 0x45080442.
//
// Solidity: function runDescribeStatus(int256 ) constant returns(bytes32)
func (_Content *ContentSession) RunDescribeStatus(arg0 *big.Int) ([32]byte, error) {
	return _Content.Contract.RunDescribeStatus(&_Content.CallOpts, arg0)
}

// RunDescribeStatus is a free data retrieval call binding the contract method 0x45080442.
//
// Solidity: function runDescribeStatus(int256 ) constant returns(bytes32)
func (_Content *ContentCallerSession) RunDescribeStatus(arg0 *big.Int) ([32]byte, error) {
	return _Content.Contract.RunDescribeStatus(&_Content.CallOpts, arg0)
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
// Solidity: function runAccess(uint256 , uint8 , bytes32[] , address[] ) returns(uint256)
func (_Content *ContentTransactor) RunAccess(opts *bind.TransactOpts, arg0 *big.Int, arg1 uint8, arg2 [][32]byte, arg3 []common.Address) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "runAccess", arg0, arg1, arg2, arg3)
}

// RunAccess is a paid mutator transaction binding the contract method 0x123e0e80.
//
// Solidity: function runAccess(uint256 , uint8 , bytes32[] , address[] ) returns(uint256)
func (_Content *ContentSession) RunAccess(arg0 *big.Int, arg1 uint8, arg2 [][32]byte, arg3 []common.Address) (*types.Transaction, error) {
	return _Content.Contract.RunAccess(&_Content.TransactOpts, arg0, arg1, arg2, arg3)
}

// RunAccess is a paid mutator transaction binding the contract method 0x123e0e80.
//
// Solidity: function runAccess(uint256 , uint8 , bytes32[] , address[] ) returns(uint256)
func (_Content *ContentTransactorSession) RunAccess(arg0 *big.Int, arg1 uint8, arg2 [][32]byte, arg3 []common.Address) (*types.Transaction, error) {
	return _Content.Contract.RunAccess(&_Content.TransactOpts, arg0, arg1, arg2, arg3)
}

// RunAccessCharge is a paid mutator transaction binding the contract method 0xf8117ca2.
//
// Solidity: function runAccessCharge(uint8 , bytes32[] , address[] ) returns(int256)
func (_Content *ContentTransactor) RunAccessCharge(opts *bind.TransactOpts, arg0 uint8, arg1 [][32]byte, arg2 []common.Address) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "runAccessCharge", arg0, arg1, arg2)
}

// RunAccessCharge is a paid mutator transaction binding the contract method 0xf8117ca2.
//
// Solidity: function runAccessCharge(uint8 , bytes32[] , address[] ) returns(int256)
func (_Content *ContentSession) RunAccessCharge(arg0 uint8, arg1 [][32]byte, arg2 []common.Address) (*types.Transaction, error) {
	return _Content.Contract.RunAccessCharge(&_Content.TransactOpts, arg0, arg1, arg2)
}

// RunAccessCharge is a paid mutator transaction binding the contract method 0xf8117ca2.
//
// Solidity: function runAccessCharge(uint8 , bytes32[] , address[] ) returns(int256)
func (_Content *ContentTransactorSession) RunAccessCharge(arg0 uint8, arg1 [][32]byte, arg2 []common.Address) (*types.Transaction, error) {
	return _Content.Contract.RunAccessCharge(&_Content.TransactOpts, arg0, arg1, arg2)
}

// RunCreate is a paid mutator transaction binding the contract method 0x7b1cdb3e.
//
// Solidity: function runCreate() returns(uint256)
func (_Content *ContentTransactor) RunCreate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "runCreate")
}

// RunCreate is a paid mutator transaction binding the contract method 0x7b1cdb3e.
//
// Solidity: function runCreate() returns(uint256)
func (_Content *ContentSession) RunCreate() (*types.Transaction, error) {
	return _Content.Contract.RunCreate(&_Content.TransactOpts)
}

// RunCreate is a paid mutator transaction binding the contract method 0x7b1cdb3e.
//
// Solidity: function runCreate() returns(uint256)
func (_Content *ContentTransactorSession) RunCreate() (*types.Transaction, error) {
	return _Content.Contract.RunCreate(&_Content.TransactOpts)
}

// RunFinalize is a paid mutator transaction binding the contract method 0x6b2d1324.
//
// Solidity: function runFinalize(uint256 ) returns(uint256)
func (_Content *ContentTransactor) RunFinalize(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "runFinalize", arg0)
}

// RunFinalize is a paid mutator transaction binding the contract method 0x6b2d1324.
//
// Solidity: function runFinalize(uint256 ) returns(uint256)
func (_Content *ContentSession) RunFinalize(arg0 *big.Int) (*types.Transaction, error) {
	return _Content.Contract.RunFinalize(&_Content.TransactOpts, arg0)
}

// RunFinalize is a paid mutator transaction binding the contract method 0x6b2d1324.
//
// Solidity: function runFinalize(uint256 ) returns(uint256)
func (_Content *ContentTransactorSession) RunFinalize(arg0 *big.Int) (*types.Transaction, error) {
	return _Content.Contract.RunFinalize(&_Content.TransactOpts, arg0)
}

// RunKill is a paid mutator transaction binding the contract method 0x9e99bbea.
//
// Solidity: function runKill() returns(uint256)
func (_Content *ContentTransactor) RunKill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "runKill")
}

// RunKill is a paid mutator transaction binding the contract method 0x9e99bbea.
//
// Solidity: function runKill() returns(uint256)
func (_Content *ContentSession) RunKill() (*types.Transaction, error) {
	return _Content.Contract.RunKill(&_Content.TransactOpts)
}

// RunKill is a paid mutator transaction binding the contract method 0x9e99bbea.
//
// Solidity: function runKill() returns(uint256)
func (_Content *ContentTransactorSession) RunKill() (*types.Transaction, error) {
	return _Content.Contract.RunKill(&_Content.TransactOpts)
}

// RunStatusChange is a paid mutator transaction binding the contract method 0x3513a805.
//
// Solidity: function runStatusChange(int256 proposed_status_code) returns(int256)
func (_Content *ContentTransactor) RunStatusChange(opts *bind.TransactOpts, proposed_status_code *big.Int) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "runStatusChange", proposed_status_code)
}

// RunStatusChange is a paid mutator transaction binding the contract method 0x3513a805.
//
// Solidity: function runStatusChange(int256 proposed_status_code) returns(int256)
func (_Content *ContentSession) RunStatusChange(proposed_status_code *big.Int) (*types.Transaction, error) {
	return _Content.Contract.RunStatusChange(&_Content.TransactOpts, proposed_status_code)
}

// RunStatusChange is a paid mutator transaction binding the contract method 0x3513a805.
//
// Solidity: function runStatusChange(int256 proposed_status_code) returns(int256)
func (_Content *ContentTransactorSession) RunStatusChange(proposed_status_code *big.Int) (*types.Transaction, error) {
	return _Content.Contract.RunStatusChange(&_Content.TransactOpts, proposed_status_code)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Content *ContentTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Content *ContentSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _Content.Contract.TransferCreatorship(&_Content.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Content *ContentTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _Content.Contract.TransferCreatorship(&_Content.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Content *ContentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Content *ContentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Content.Contract.TransferOwnership(&_Content.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
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
// Solidity: event DbgAddress(address a)
func (_Content *ContentFilterer) FilterDbgAddress(opts *bind.FilterOpts) (*ContentDbgAddressIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "DbgAddress")
	if err != nil {
		return nil, err
	}
	return &ContentDbgAddressIterator{contract: _Content.contract, event: "DbgAddress", logs: logs, sub: sub}, nil
}

// WatchDbgAddress is a free log subscription operation binding the contract event 0x474589881ea4b1937d02648256f59b0dec5c28f4d1a8be6efcb50aa51aaf69de.
//
// Solidity: event DbgAddress(address a)
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
// Solidity: event DbgString(string s)
func (_Content *ContentFilterer) FilterDbgString(opts *bind.FilterOpts) (*ContentDbgStringIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "DbgString")
	if err != nil {
		return nil, err
	}
	return &ContentDbgStringIterator{contract: _Content.contract, event: "DbgString", logs: logs, sub: sub}, nil
}

// WatchDbgString is a free log subscription operation binding the contract event 0x415302daeab7bd01c23adb4dee059bc666b313134024f09781aa6d73ff6ec042.
//
// Solidity: event DbgString(string s)
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
// Solidity: event DbgUint(uint256 u)
func (_Content *ContentFilterer) FilterDbgUint(opts *bind.FilterOpts) (*ContentDbgUintIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "DbgUint")
	if err != nil {
		return nil, err
	}
	return &ContentDbgUintIterator{contract: _Content.contract, event: "DbgUint", logs: logs, sub: sub}, nil
}

// WatchDbgUint is a free log subscription operation binding the contract event 0xf209d49352c8b16a38abaf19dd720dfc5a398b0209f3d48a673a52234e3c6aac.
//
// Solidity: event DbgUint(uint256 u)
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
// Solidity: event DbgUint256(uint256 u)
func (_Content *ContentFilterer) FilterDbgUint256(opts *bind.FilterOpts) (*ContentDbgUint256Iterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "DbgUint256")
	if err != nil {
		return nil, err
	}
	return &ContentDbgUint256Iterator{contract: _Content.contract, event: "DbgUint256", logs: logs, sub: sub}, nil
}

// WatchDbgUint256 is a free log subscription operation binding the contract event 0x6b0cf6507e95b7606ee5ef593c4219bc4056b30825b0eecd12e5e7022126f23f.
//
// Solidity: event DbgUint256(uint256 u)
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
// Solidity: event RunAccess(uint256 requestID, uint256 result)
func (_Content *ContentFilterer) FilterRunAccess(opts *bind.FilterOpts) (*ContentRunAccessIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "RunAccess")
	if err != nil {
		return nil, err
	}
	return &ContentRunAccessIterator{contract: _Content.contract, event: "RunAccess", logs: logs, sub: sub}, nil
}

// WatchRunAccess is a free log subscription operation binding the contract event 0x3e68dc35f88d76818f276322c37f5021ee00e232fe0d27a93c02801aec4d9c58.
//
// Solidity: event RunAccess(uint256 requestID, uint256 result)
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
// Solidity: event RunAccessCharge(uint8 level, int256 calculateAccessCharge)
func (_Content *ContentFilterer) FilterRunAccessCharge(opts *bind.FilterOpts) (*ContentRunAccessChargeIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "RunAccessCharge")
	if err != nil {
		return nil, err
	}
	return &ContentRunAccessChargeIterator{contract: _Content.contract, event: "RunAccessCharge", logs: logs, sub: sub}, nil
}

// WatchRunAccessCharge is a free log subscription operation binding the contract event 0xffadad18ab3777a19f664019a6261b011ab9405749e01a45950d44fb9360b385.
//
// Solidity: event RunAccessCharge(uint8 level, int256 calculateAccessCharge)
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

// ContentRunCreateIterator is returned from FilterRunCreate and is used to iterate over the raw logs and unpacked data for RunCreate events raised by the Content contract.
type ContentRunCreateIterator struct {
	Event *ContentRunCreate // Event containing the contract specifics and raw log

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
func (it *ContentRunCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentRunCreate)
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
		it.Event = new(ContentRunCreate)
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
func (it *ContentRunCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentRunCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentRunCreate represents a RunCreate event raised by the Content contract.
type ContentRunCreate struct {
	Result *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRunCreate is a free log retrieval operation binding the contract event 0x9df71221e13c480b974b5d5bd7591b30b7ea3bfff8a56dfa7fde810a14c1c39b.
//
// Solidity: event RunCreate(uint256 result)
func (_Content *ContentFilterer) FilterRunCreate(opts *bind.FilterOpts) (*ContentRunCreateIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "RunCreate")
	if err != nil {
		return nil, err
	}
	return &ContentRunCreateIterator{contract: _Content.contract, event: "RunCreate", logs: logs, sub: sub}, nil
}

// WatchRunCreate is a free log subscription operation binding the contract event 0x9df71221e13c480b974b5d5bd7591b30b7ea3bfff8a56dfa7fde810a14c1c39b.
//
// Solidity: event RunCreate(uint256 result)
func (_Content *ContentFilterer) WatchRunCreate(opts *bind.WatchOpts, sink chan<- *ContentRunCreate) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "RunCreate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentRunCreate)
				if err := _Content.contract.UnpackLog(event, "RunCreate", log); err != nil {
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
	Result    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRunFinalize is a free log retrieval operation binding the contract event 0xbf0f2215c45c5ee802d4c20bdfc915308c4459b0f6a78f23ad350e6408bf2891.
//
// Solidity: event RunFinalize(uint256 requestID, uint256 result)
func (_Content *ContentFilterer) FilterRunFinalize(opts *bind.FilterOpts) (*ContentRunFinalizeIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "RunFinalize")
	if err != nil {
		return nil, err
	}
	return &ContentRunFinalizeIterator{contract: _Content.contract, event: "RunFinalize", logs: logs, sub: sub}, nil
}

// WatchRunFinalize is a free log subscription operation binding the contract event 0xbf0f2215c45c5ee802d4c20bdfc915308c4459b0f6a78f23ad350e6408bf2891.
//
// Solidity: event RunFinalize(uint256 requestID, uint256 result)
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

// ContentRunKillIterator is returned from FilterRunKill and is used to iterate over the raw logs and unpacked data for RunKill events raised by the Content contract.
type ContentRunKillIterator struct {
	Event *ContentRunKill // Event containing the contract specifics and raw log

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
func (it *ContentRunKillIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentRunKill)
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
		it.Event = new(ContentRunKill)
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
func (it *ContentRunKillIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentRunKillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentRunKill represents a RunKill event raised by the Content contract.
type ContentRunKill struct {
	Result *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRunKill is a free log retrieval operation binding the contract event 0x6d0dbfc3805aef247651b04b50fc717599f7e0b66c6b022ae1544406f7bf8f86.
//
// Solidity: event RunKill(uint256 result)
func (_Content *ContentFilterer) FilterRunKill(opts *bind.FilterOpts) (*ContentRunKillIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "RunKill")
	if err != nil {
		return nil, err
	}
	return &ContentRunKillIterator{contract: _Content.contract, event: "RunKill", logs: logs, sub: sub}, nil
}

// WatchRunKill is a free log subscription operation binding the contract event 0x6d0dbfc3805aef247651b04b50fc717599f7e0b66c6b022ae1544406f7bf8f86.
//
// Solidity: event RunKill(uint256 result)
func (_Content *ContentFilterer) WatchRunKill(opts *bind.WatchOpts, sink chan<- *ContentRunKill) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "RunKill")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentRunKill)
				if err := _Content.contract.UnpackLog(event, "RunKill", log); err != nil {
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
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRunStatusChange is a free log retrieval operation binding the contract event 0xb6c1c013bb5004fe8e943c6890e300ccedf9bd73dcd4eb291b31b9f96874feff.
//
// Solidity: event RunStatusChange(int256 proposedStatusCode, int256 returnStatusCode)
func (_Content *ContentFilterer) FilterRunStatusChange(opts *bind.FilterOpts) (*ContentRunStatusChangeIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "RunStatusChange")
	if err != nil {
		return nil, err
	}
	return &ContentRunStatusChangeIterator{contract: _Content.contract, event: "RunStatusChange", logs: logs, sub: sub}, nil
}

// WatchRunStatusChange is a free log subscription operation binding the contract event 0xb6c1c013bb5004fe8e943c6890e300ccedf9bd73dcd4eb291b31b9f96874feff.
//
// Solidity: event RunStatusChange(int256 proposedStatusCode, int256 returnStatusCode)
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
const EditableBin = `0x60806040526000805432600160a060020a031991821681178355600180549092161790556103a790819061003390396000f30060806040526004361061008d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f811461008f57806341c0e1b5146100c05780636d2e4b1b146100d55780638da5cb5b146100f6578063c287e0ed1461010b578063e02dd9c214610120578063f14fcbc814610147578063f2fde38b1461015f575b005b34801561009b57600080fd5b506100a4610180565b60408051600160a060020a039092168252519081900360200190f35b3480156100cc57600080fd5b5061008d61018f565b3480156100e157600080fd5b5061008d600160a060020a03600435166101cb565b34801561010257600080fd5b506100a4610226565b34801561011757600080fd5b5061008d610235565b34801561012c57600080fd5b5061013561029a565b60408051918252519081900360200190f35b34801561015357600080fd5b5061008d6004356102a0565b34801561016b57600080fd5b5061008d600160a060020a0360043516610309565b600054600160a060020a031681565b600154600160a060020a03163214806101b25750600154600160a060020a031633145b15156101bd57600080fd5b600154600160a060020a0316ff5b600054600160a060020a031632146101e257600080fd5b600160a060020a03811615156101f757600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031681565b600154600160a060020a03163214806102585750600154600160a060020a031633145b151561026357600080fd5b60025460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b60025481565b600154600160a060020a03163214806102c35750600154600160a060020a031633145b15156102ce57600080fd5b60028190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b600154600160a060020a031632148061032c5750600154600160a060020a031633145b151561033757600080fd5b600160a060020a038116151561034c57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820420ef50907fd2bbec63ceb5c0b9986d7e826f0ce278e0afaeea3ad30071798fb0029`

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
// Solidity: function commit(bytes32 object_hash) returns()
func (_Editable *EditableTransactor) Commit(opts *bind.TransactOpts, object_hash [32]byte) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "commit", object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 object_hash) returns()
func (_Editable *EditableSession) Commit(object_hash [32]byte) (*types.Transaction, error) {
	return _Editable.Contract.Commit(&_Editable.TransactOpts, object_hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 object_hash) returns()
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
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Editable *EditableTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Editable *EditableSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _Editable.Contract.TransferCreatorship(&_Editable.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Editable *EditableTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _Editable.Contract.TransferCreatorship(&_Editable.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Editable *EditableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Editable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Editable *EditableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Editable.Contract.TransferOwnership(&_Editable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
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
// Solidity: event Commit(bytes32 objectHash)
func (_Editable *EditableFilterer) FilterCommit(opts *bind.FilterOpts) (*EditableCommitIterator, error) {

	logs, sub, err := _Editable.contract.FilterLogs(opts, "Commit")
	if err != nil {
		return nil, err
	}
	return &EditableCommitIterator{contract: _Editable.contract, event: "Commit", logs: logs, sub: sub}, nil
}

// WatchCommit is a free log subscription operation binding the contract event 0x9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff.
//
// Solidity: event Commit(bytes32 objectHash)
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
// Solidity: event UpdateRequest(bytes32 objectHash)
func (_Editable *EditableFilterer) FilterUpdateRequest(opts *bind.FilterOpts) (*EditableUpdateRequestIterator, error) {

	logs, sub, err := _Editable.contract.FilterLogs(opts, "UpdateRequest")
	if err != nil {
		return nil, err
	}
	return &EditableUpdateRequestIterator{contract: _Editable.contract, event: "UpdateRequest", logs: logs, sub: sub}, nil
}

// WatchUpdateRequest is a free log subscription operation binding the contract event 0x1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef19.
//
// Solidity: event UpdateRequest(bytes32 objectHash)
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
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x60806040526000805432600160a060020a0319918216811783556001805490921617905561025e90819061003390396000f30060806040526004361061006c5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f811461006e57806341c0e1b51461009f5780636d2e4b1b146100b45780638da5cb5b146100d5578063f2fde38b146100ea575b005b34801561007a57600080fd5b5061008361010b565b60408051600160a060020a039092168252519081900360200190f35b3480156100ab57600080fd5b5061006c61011a565b3480156100c057600080fd5b5061006c600160a060020a0360043516610156565b3480156100e157600080fd5b506100836101b1565b3480156100f657600080fd5b5061006c600160a060020a03600435166101c0565b600054600160a060020a031681565b600154600160a060020a031632148061013d5750600154600160a060020a031633145b151561014857600080fd5b600154600160a060020a0316ff5b600054600160a060020a0316321461016d57600080fd5b600160a060020a038116151561018257600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031681565b600154600160a060020a03163214806101e35750600154600160a060020a031633145b15156101ee57600080fd5b600160a060020a038116151561020357600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820d63ab1762d0c1232e12eecb9466def742fcb9080d6c024f82a29ba69d63057650029`

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
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Ownable *OwnableTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Ownable *OwnableSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferCreatorship(&_Ownable.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Ownable *OwnableTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferCreatorship(&_Ownable.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}
