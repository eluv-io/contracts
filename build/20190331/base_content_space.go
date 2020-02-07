// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts_20190331

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
const AccessibleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accessRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AccessRequest\",\"type\":\"event\"}]"

// AccessibleBin is the compiled bytecode used for deploying new contracts.
const AccessibleBin = `0x60806040527f41636365737369626c6532303139303232323133353930304d4c00000000000060005534801561003457600080fd5b5060fa806100436000396000f30060806040526004361060485763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166354fd4d508114604d578063f1551887146071575b600080fd5b348015605857600080fd5b50605f6097565b60408051918252519081900360200190f35b348015607c57600080fd5b506083609d565b604080519115158252519081900360200190f35b60005481565b6040516000907fed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88908290a1506001905600a165627a7a72305820083ab49a02513a3251123e285170d3419efe10c9c78b93ae45aa32d43d6c6ccc0029`

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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Accessible *AccessibleCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Accessible.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Accessible *AccessibleSession) Version() ([32]byte, error) {
	return _Accessible.Contract.Version(&_Accessible.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Accessible *AccessibleCallerSession) Version() ([32]byte, error) {
	return _Accessible.Contract.Version(&_Accessible.CallOpts)
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
const BaseAccessControlGroupABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"members\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"grantAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"hasManagerAccess\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"grantManagerAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"revokeAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"hasAccess\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"revokeManagerAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"managers\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"content_space\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"MemberAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"ManagerAccessGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"MemberRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"ManagerAccessRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"operationCode\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"UnauthorizedOperation\",\"type\":\"event\"}]"

// BaseAccessControlGroupBin is the compiled bytecode used for deploying new contracts.
const BaseAccessControlGroupBin = `0x60806040527f4f776e61626c6532303139303331353134313530304d4c0000000000000000006000557f42734163636573734374726c47727032303139303331353137323930304d4c0060035534801561005857600080fd5b506040516020806107a983398101604090815290516001805432600160a060020a0319918216811780845560028054841690921790915560048054909216600160a060020a03948516179091558216600090815260066020908152848220805460ff199081168517909155835490941682526005905292832080549092161790556106c09081906100e990396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100dc57806308ae4b0c1461010d5780630ae5e7391461014257806341c0e1b51461016357806342e7ba7b1461017857806354fd4d50146101995780636d2e4b1b146101c057806375861a95146101e157806385e68531146102025780638da5cb5b1461022357806395a078e814610238578063af570c0414610259578063cdb849b71461026e578063f2fde38b1461028f578063fdff9b4d146102b0575b005b3480156100e857600080fd5b506100f16102d1565b60408051600160a060020a039092168252519081900360200190f35b34801561011957600080fd5b5061012e600160a060020a03600435166102e0565b604080519115158252519081900360200190f35b34801561014e57600080fd5b506100da600160a060020a03600435166102f5565b34801561016f57600080fd5b506100da610371565b34801561018457600080fd5b5061012e600160a060020a03600435166103ad565b3480156101a557600080fd5b506101ae6103d0565b60408051918252519081900360200190f35b3480156101cc57600080fd5b506100da600160a060020a03600435166103d6565b3480156101ed57600080fd5b506100da600160a060020a0360043516610431565b34801561020e57600080fd5b506100da600160a060020a03600435166104ba565b34801561022f57600080fd5b506100f1610548565b34801561024457600080fd5b5061012e600160a060020a0360043516610557565b34801561026557600080fd5b506100f161057a565b34801561027a57600080fd5b506100da600160a060020a0360043516610589565b34801561029b57600080fd5b506100da600160a060020a036004351661060d565b3480156102bc57600080fd5b5061012e600160a060020a036004351661067f565b600154600160a060020a031681565b60056020526000908152604090205460ff1681565b3360009081526006602052604090205460ff16151560011461031657600080fd5b600160a060020a038116600081815260056020908152604091829020805460ff19166001179055815192835290517fb251eb052afc73ffd02ffe85ad79990a8b3fed60d76dbc2fa2fdd7123dffd9149281900390910190a150565b600254600160a060020a03163214806103945750600254600160a060020a031633145b151561039f57600080fd5b600254600160a060020a0316ff5b600160a060020a031660009081526006602052604090205460ff16151560011490565b60035481565b600154600160a060020a031632146103ed57600080fd5b600160a060020a038116151561040257600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a03163214806104545750600254600160a060020a031633145b151561045f57600080fd5b600160a060020a038116600081815260066020908152604091829020805460ff19166001179055815192835290517f93bcaab179551bde429187645251f8e1fb8ac85801fcb1cf91eb2c9043d611179281900390910190a150565b3360009081526006602052604090205460ff161515600114806104e5575033600160a060020a038216145b15156104f057600080fd5b600160a060020a038116600081815260056020908152604091829020805460ff19169055815192835290517f745cd29407db644ed93e3ceb61cbcab96d1dfb496989ac5d5bf514fc5a9fab9c9281900390910190a150565b600254600160a060020a031681565b600160a060020a031660009081526005602052604090205460ff16151560011490565b600454600160a060020a031681565b600254600160a060020a03163314806105aa575033600160a060020a038216145b15156105b557600080fd5b600160a060020a038116600081815260066020908152604091829020805460ff19169055815192835290517f2d6aa1a9629d125e23a0cf692cda7cd6795dff1652eedd4673b38ec31e387b959281900390910190a150565b600254600160a060020a03163214806106305750600254600160a060020a031633145b151561063b57600080fd5b600160a060020a038116151561065057600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60066020526000908152604090205460ff16815600a165627a7a72305820a3d364eca755c88b24e28f3029b5ff9767cb2358a596e6de29fab07635e9de520029`

// DeployBaseAccessControlGroup deploys a new Ethereum contract, binding an instance of BaseAccessControlGroup to it.
func DeployBaseAccessControlGroup(auth *bind.TransactOpts, backend bind.ContractBackend, content_space common.Address) (common.Address, *types.Transaction, *BaseAccessControlGroup, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseAccessControlGroupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaseAccessControlGroupBin), backend, content_space)
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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseAccessControlGroup *BaseAccessControlGroupCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseAccessControlGroup.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseAccessControlGroup *BaseAccessControlGroupSession) Version() ([32]byte, error) {
	return _BaseAccessControlGroup.Contract.Version(&_BaseAccessControlGroup.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseAccessControlGroup *BaseAccessControlGroupCallerSession) Version() ([32]byte, error) {
	return _BaseAccessControlGroup.Contract.Version(&_BaseAccessControlGroup.CallOpts)
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
const BaseContentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"STATUS_PUBLISHED\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"publish\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"request_ID\",\"type\":\"uint256\"},{\"name\":\"payee\",\"type\":\"address\"},{\"name\":\"label\",\"type\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"processRequestPayment\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestMap\",\"outputs\":[{\"name\":\"originator\",\"type\":\"address\"},{\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"int8\"},{\"name\":\"settled\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"statusCode\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressKMS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentType\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"status_code\",\"type\":\"int256\"}],\"name\":\"statusCodeDescription\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"level\",\"type\":\"uint8\"},{\"name\":\"custom_values\",\"type\":\"bytes32[]\"},{\"name\":\"stakeholders\",\"type\":\"address[]\"}],\"name\":\"getAccessInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STATUS_DRAFT\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"status_code\",\"type\":\"int256\"}],\"name\":\"setStatusCode\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"request_ID\",\"type\":\"uint256\"},{\"name\":\"score_pct\",\"type\":\"uint256\"},{\"name\":\"ml_out_hash\",\"type\":\"bytes32\"}],\"name\":\"accessComplete\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accessCharge\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"versionHashes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"status_code\",\"type\":\"int256\"}],\"name\":\"updateStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"level\",\"type\":\"uint8\"},{\"name\":\"custom_values\",\"type\":\"bytes32[]\"},{\"name\":\"stakeholders\",\"type\":\"address[]\"}],\"name\":\"getAccessCharge\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requestID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"level\",\"type\":\"uint8\"},{\"name\":\"pke_requestor\",\"type\":\"string\"},{\"name\":\"pke_AFGH\",\"type\":\"string\"},{\"name\":\"custom_values\",\"type\":\"bytes32[]\"},{\"name\":\"stakeholders\",\"type\":\"address[]\"}],\"name\":\"accessRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"libraryAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"address_KMS\",\"type\":\"address\"}],\"name\":\"setAddressKMS\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canPublish\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STATUS_REVIEW\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"objectHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setContentContractAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"request_ID\",\"type\":\"uint256\"},{\"name\":\"access_granted\",\"type\":\"bool\"},{\"name\":\"re_key\",\"type\":\"string\"},{\"name\":\"encrypted_AES_key\",\"type\":\"string\"}],\"name\":\"accessGrant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_objectHash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"charge\",\"type\":\"uint256\"}],\"name\":\"setAccessCharge\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"statusDescription\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"content_type\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"containingLibrary\",\"type\":\"address\"}],\"name\":\"ContentObjectCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentType\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"contentContractAddress\",\"type\":\"address\"}],\"name\":\"SetContentType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"contentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"pkeRequestor\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"pkeAFGH\",\"type\":\"string\"}],\"name\":\"AccessRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"access_granted\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"reKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"encryptedAESKey\",\"type\":\"string\"}],\"name\":\"AccessGrant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"customValue\",\"type\":\"bytes32\"}],\"name\":\"AccessRequestValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"stakeholder\",\"type\":\"address\"}],\"name\":\"AccessRequestStakeholder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"scorePct\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"mlOutHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"customContractResult\",\"type\":\"bool\"}],\"name\":\"AccessComplete\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentContractAddress\",\"type\":\"address\"}],\"name\":\"SetContentContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accessCharge\",\"type\":\"uint256\"}],\"name\":\"SetAccessCharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"accessCharge\",\"type\":\"uint256\"}],\"name\":\"GetAccessCharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accessCharge\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amountProvided\",\"type\":\"uint256\"}],\"name\":\"InsufficientFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"statusCode\",\"type\":\"int256\"}],\"name\":\"SetStatusCode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestStatus\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"statusCode\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"Publish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"custom_contract\",\"type\":\"address\"}],\"name\":\"InvokeCustomPreHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"custom_contract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"ReturnCustomHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"custom_contract\",\"type\":\"address\"}],\"name\":\"InvokeCustomPostHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"CommitPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"Commit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"UpdateRequest\",\"type\":\"event\"}]"

// BaseContentBin is the compiled bytecode used for deploying new contracts.
const BaseContentBin = `0x608060408190527f4f776e61626c6532303139303331353134313530304d4c00000000000000000060009081557f4564697461626c6532303139303331353134313830304d4c00000000000000006003557f42617365436f6e74656e7432303139303331353137353130304d4c0000000000600555600c556020806124e1833981016040818152915160018054600160a060020a031990811632908117909255600280548216909217909155600980548216331790819055600019600b5560068054600160a060020a038086169190941617905516825291517fc3decc188980e855666b70498ca85e8fa284d97d30483d828fa126f7303d7d199181900360200190a1506123cf806101126000396000f3006080604052600436106101a75763ffffffff60e060020a60003504166217de9881146101a957806302d05d3f146101d0578063075d4782146102015780630c6d3f931461021d5780631a735f18146102885780632310167f146102d557806327c1c21d146102ea57806332eaf21b146102ff57806336ebffca14610314578063388642841461032957806338d0f5041461034157806341c0e1b5146103fa5780634dd707881461040f5780635267db441461042457806354fd4d501461043c5780635cc4aa9b1461045157806364ade32b146104625780636d2e4b1b146104775780637ca8f618146104985780638280dd8f146104b0578063879fe48f146104c85780638da5cb5b1461055e5780638f77920114610573578063a1ff106e14610588578063b816f5131461068c578063c287e0ed146106a1578063c9e8e72d146106b6578063cbcd4461146106d7578063d810f8c8146106ec578063e02dd9c214610701578063e538530314610716578063ee56d76714610737578063f14fcbc8146107d7578063f2fde38b146107ef578063f4d9bae814610810578063f81ab0ae14610828575b005b3480156101b557600080fd5b506101be61083d565b60408051918252519081900360200190f35b3480156101dc57600080fd5b506101e5610861565b60408051600160a060020a039092168252519081900360200190f35b610209610870565b604080519115158252519081900360200190f35b34801561022957600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526102099482359460248035600160a060020a03169536959460649492019190819084018382808284375094975050933594506109d19350505050565b34801561029457600080fd5b506102a0600435610a18565b60408051600160a060020a0390951685526020850193909352600091820b90910b838301526060830152519081900360800190f35b3480156102e157600080fd5b506101e5610a4c565b3480156102f657600080fd5b506101be610a5b565b34801561030b57600080fd5b506101e5610a61565b34801561032057600080fd5b506101e5610a70565b34801561033557600080fd5b506101be600435610a7f565b34801561034d57600080fd5b506040805160206004602480358281013584810280870186019097528086526103d796843560ff1696369660449591949091019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750949750610bd19650505050505050565b604051808360000b60000b81526020018281526020019250505060405180910390f35b34801561040657600080fd5b506101a7610d5b565b34801561041b57600080fd5b506101be610e40565b34801561043057600080fd5b506101be600435610e64565b34801561044857600080fd5b506101be610ef9565b610209600435602435604435610eff565b34801561046e57600080fd5b506101be6111c0565b34801561048357600080fd5b506101a7600160a060020a03600435166111c6565b3480156104a457600080fd5b506101be600435611221565b3480156104bc57600080fd5b506101be600435611240565b3480156104d457600080fd5b506040805160206004602480358281013584810280870186019097528086526101be96843560ff1696369660449591949091019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506113b79650505050505050565b34801561056a57600080fd5b506101e5611421565b34801561057f57600080fd5b506101be611430565b60408051602060046024803582810135601f810185900485028601850190965285855261020995833560ff1695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506114369650505050505050565b34801561069857600080fd5b506101e5611927565b3480156106ad57600080fd5b506101a7611936565b3480156106c257600080fd5b506101a7600160a060020a036004351661199b565b3480156106e357600080fd5b506102096119f8565b3480156106f857600080fd5b506101be611ac9565b34801561070d57600080fd5b506101be611aed565b34801561072257600080fd5b506101a7600160a060020a0360043516611af3565b34801561074357600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610209948235946024803515159536959460649492019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750611cb79650505050505050565b3480156107e357600080fd5b506101a760043561204e565b3480156107fb57600080fd5b506101a7600160a060020a03600435166120b7565b34801561081c57600080fd5b506101be600435612129565b34801561083457600080fd5b506101be61219a565b7f5075626c6973686564000000000000000000000000000000000000000000000081565b600154600160a060020a031681565b600080600061087d6119f8565b151561088857600080fd5b600454156108c757600454600e80546001810182556000919091527fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fd01555b6108d2600f546121ac565b6000600f556108e16001611240565b50600091506000600b5413156109825750600954604080517f49102e610000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169182916349102e619160048083019260209291908290030181600087803b15801561095357600080fd5b505af1158015610967573d6000803e3d6000fd5b505050506040513d602081101561097d57600080fd5b505191505b600b54600454604080518515158152602081019390935282810191909152517f901e6f3cdc4c61620d5d424116934b9af6e31ba79cdeaa349336d93ecfe846d49181900360600190a150919050565b600854600090600160a060020a0316158015906109f85750600854600160a060020a031633145b1515610a0357600080fd5b610a0f8585858561220e565b95945050505050565b600d602052600090815260408120805460018201546002830154600390930154600160a060020a03909216939092900b9084565b600854600160a060020a031681565b600b5481565b600754600160a060020a031681565b600654600160a060020a031681565b60085460009081908190600160a060020a031615610b2f5750600854604080517f45080442000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a0390921691829163450804429160248083019260209291908290030181600087803b158015610b0057600080fd5b505af1158015610b14573d6000803e3d6000fd5b505050506040513d6020811015610b2a57600080fd5b505191505b8115610b3d57819250610bca565b831515610b6c577f5075626c697368656400000000000000000000000000000000000000000000009250610bca565b6000841215610b9d577f44726166740000000000000000000000000000000000000000000000000000009250610bca565b6000841315610bca577f447261667420696e20726576696577000000000000000000000000000000000092505b5050919050565b60085460009081908190600019908290600160a060020a031615610d2e57506008546040517f0f82c16f00000000000000000000000000000000000000000000000000000000815260ff891660048201908152606060248301908152895160648401528951600160a060020a03909416938493630f82c16f938d938d938d9360448101916084909101906020808801910280838360005b83811015610c80578181015183820152602001610c68565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610cbf578181015183820152602001610ca7565b50505050905001955050505050506040805180830381600087803b158015610ce657600080fd5b505af1158015610cfa573d6000803e3d6000fd5b505050506040513d6040811015610d1057600080fd5b508051602090910151600a54909450909250831115610d2e57606491505b8160000b6000191415610d4957600a54600095509350610d50565b8183945094505b505050935093915050565b600254600090600160a060020a0316321480610d815750600254600160a060020a031633145b1515610d8c57600080fd5b600854600160a060020a031615610e355750600854604080517f9e99bbea0000000000000000000000000000000000000000000000000000000081529051600160a060020a03909216918291639e99bbea9160048083019260209291908290030181600087803b158015610dff57600080fd5b505af1158015610e13573d6000803e3d6000fd5b505050506040513d6020811015610e2957600080fd5b505115610e3557600080fd5b610e3d612340565b50565b7f447261667400000000000000000000000000000000000000000000000000000081565b600254600090600160a060020a031632148015610e9857506000821280610e985750600082138015610e9857506000600b54125b15610ea357600b8290555b600954600160a060020a0316331415610ebc57600b8290555b600b5460408051918252517fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a29181900360200190a15050600b5490565b60055481565b6000838152600d602052604081208054829081908190600160a060020a031615801590610f4957508354600160a060020a0316331480610f495750600254600160a060020a031633145b1515610f5457600080fd5b6008548715159350600160a060020a03161561100e57600854604080517f17685953000000000000000000000000000000000000000000000000000000008152600481018b9052602481018a90529051600160a060020a03909216935083916317685953916044808201926020929091908290030181600087803b158015610fdb57600080fd5b505af1158015610fef573d6000803e3d6000fd5b505050506040513d602081101561100557600080fd5b50518015935090505b8354600160a060020a031633141561104b57821561103b576002848101805460ff1916909117905561104b565b60028401805460ff191660fe1790555b836001015484600301541015611129576002840154600090810b810b136110cc57835460408051808201909152600681527f726566756e6400000000000000000000000000000000000000000000000000006020820152600386015460018701546110c6938c93600160a060020a039091169290910361220e565b50611129565b60025460408051808201909152600e81527f72656c6561736520657363726f77000000000000000000000000000000000000602082015260038601546001870154611127938c93600160a060020a039091169290910361220e565b505b6000888152600d60209081526040808320805473ffffffffffffffffffffffffffffffffffffffff191681556001810184905560028101805460ff191690556003019290925581518a8152908101899052808201889052841515606082015290517f7f1f4b28434ce7beab4983e64a8b5bb96e195a67029fdaff925028aec57fbc6b9181900360800190a150909695505050505050565b600a5481565b600154600160a060020a031632146111dd57600080fd5b600160a060020a03811615156111f257600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600e80548290811061122f57fe5b600091825260209091200154905081565b600080600061124d6119f8565b151561125857600080fd5b600854600160a060020a031615156112dc57600254600160a060020a031632148061128d5750600254600160a060020a031633145b80156112a457508360001914806112a45750836001145b156112b1578391506112d7565b600954600160a060020a0316331480156112ce57506000600b5412155b156112d7578391505b611374565b50600854604080517f3513a805000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a03909216918291633513a8059160248083019260209291908290030181600087803b15801561134557600080fd5b505af1158015611359573d6000803e3d6000fd5b505050506040513d602081101561136f57600080fd5b505191505b600b8290556040805183815290517fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a29181900360200190a15050600b5492915050565b60008060006113c7868686610bd1565b92509050600081900b156113da57600080fd5b6040805160ff881681526020810184905281517fa58326ee5bb617cb8b4f0d0f5f557c469d2d05d7a738f777037deda9c724b370929181900390910190a150949350505050565b600254600160a060020a031681565b600c5481565b600080600061144361237c565b600c80546001019055600954604080517f95a078e80000000000000000000000000000000000000000000000000000000081523260048201529051600160a060020a0390921694506000918291829187916395a078e891602480830192602092919082900301818787803b1580156114ba57600080fd5b505af11580156114ce573d6000803e3d6000fd5b505050506040513d60208110156114e457600080fd5b505115156114f157600080fd5b600254600160a060020a0316321461151d5761150e8c8a8a6113b7565b94503485111561151d57600080fd5b60408051608081018252338152346020808301918252600083850181815260608501828152600c548352600d90935294812084518154600160a060020a0391821673ffffffffffffffffffffffffffffffffffffffff199091161782559351600182015594516002860180549190920b60ff1660ff199091161790555160039093019290925560085490955016156116e557600860009054906101000a9004600160a060020a0316925082600160a060020a031663123e0e80600c548e8c8c6040518563ffffffff1660e060020a028152600401808581526020018460ff1660ff1681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561164457818101518382015260200161162c565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561168357818101518382015260200161166b565b505050509050019650505050505050602060405180830381600087803b1580156116ac57600080fd5b505af11580156116c0573d6000803e3d6000fd5b505050506040513d60208110156116d657600080fd5b5051915081156116e557600080fd5b7f089a6f1788a3c353423e1be4ba12533bdde7d908bb41abeee185af0acb3df562600c548d6004548e8e604051808681526020018560ff1660ff16815260200184600019166000191681526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611770578181015183820152602001611758565b50505050905090810190601f16801561179d5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156117d05781810151838201526020016117b8565b50505050905090810190601f1680156117fd5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15060005b885181101561188d57888181518110151561182a57fe5b6020908102909101015115611885577f515e0a48b385fce2a8e4d9f169a97c4f6ea669a752358f5e6ab37cc3c2e84c38898281518110151561186857fe5b602090810290910181015160408051918252519081900390910190a15b600101611813565b5060005b87518110156119155788818151811015156118a857fe5b602090810290910101511561190d577fb6e3239e521a6c66920ae634f8e921a37e6991d520ac44d52f8516397f41b68488828151811015156118e657fe5b602090810290910181015160408051600160a060020a039092168252519081900390910190a15b600101611891565b5060019b9a5050505050505050505050565b600954600160a060020a031681565b600254600160a060020a03163214806119595750600254600160a060020a031633145b151561196457600080fd5b60045460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600254600160a060020a03163214806119be5750600254600160a060020a031633145b15156119c957600080fd5b6007805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6002546000908190600160a060020a0316331480611a205750600954600160a060020a031633145b15611a2e5760019150611ac5565b50600954604080517f26683e140000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169182916326683e149160248083019260209291908290030181600087803b158015611a9657600080fd5b505af1158015611aaa573d6000803e3d6000fd5b505050506040513d6020811015611ac057600080fd5b505191505b5090565b7f447261667420696e20726576696577000000000000000000000000000000000081565b60045481565b60025460009081908190600160a060020a0316321480611b1d5750600254600160a060020a031633145b1515611b2857600080fd5b600854600160a060020a031615611bc857600860009054906101000a9004600160a060020a0316925082600160a060020a0316639e99bbea6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611b8f57600080fd5b505af1158015611ba3573d6000803e3d6000fd5b505050506040513d6020811015611bb957600080fd5b505191508115611bc857600080fd5b6008805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03861690811790915515611c725783925082600160a060020a0316637b1cdb3e6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611c3957600080fd5b505af1158015611c4d573d6000803e3d6000fd5b505050506040513d6020811015611c6357600080fd5b505190508015611c7257600080fd5b60085460408051600160a060020a039092168252517fa6f2e38f0cfebf27212317fced3ac40bc62e00bd33f38d69603710740c69acb79181900360200190a150505050565b600254600090819081908190600160a060020a0316331480611ce35750600754600160a060020a031633145b1515611cee57600080fd5b6000888152600d602052604090208054909350600160a060020a03161515611d1557600080fd5b600854879250600160a060020a031615611dce5750600854604080517fe870ed91000000000000000000000000000000000000000000000000000000008152600481018a905288151560248201529051600160a060020a0390921691829163e870ed919160448083019260209291908290030181600087803b158015611d9a57600080fd5b505af1158015611dae573d6000803e3d6000fd5b505050506040513d6020811015611dc457600080fd5b5051159150611ea2565b826001015483600301541015611ea257861515611e4557825460408051808201909152600f81527f616363657373206465636c696e65640000000000000000000000000000000000602082015260038501546001860154611e3f938c93600160a060020a039091169290910361220e565b50611ea2565b60025460408051808201909152600d81527f6f776e6572207061796d656e7400000000000000000000000000000000000000602082015260038501546001860154611ea0938c93600160a060020a039091169290910361220e565b505b60018215151415611fdd5760028301805460ff19166001908117909155604080518a8152602080820184905260809282018381528a519383019390935289517f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718948d9490938c938c93919291606084019160a08501919087019080838360005b83811015611f3a578181015183820152602001611f22565b50505050905090810190601f168015611f675780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611f9a578181015183820152602001611f82565b50505050905090810190601f168015611fc75780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a1612043565b60028301805460ff191660ff179055604080518981526000602082018190526080828401819052820181905260c06060830181905282015290517f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718918190036101000190a15b509695505050505050565b600254600160a060020a03163214806120715750600254600160a060020a031633145b151561207c57600080fd5b600f8190556040805182815290517fa288d3a2aba7d5dec44e93ff46d2f1129e251695be2046de105f9a9c6feefcaa9181900360200190a150565b600254600160a060020a03163214806120da5750600254600160a060020a031633145b15156120e557600080fd5b600160a060020a03811615156120fa57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600090600160a060020a031632148061214f5750600254600160a060020a031633145b151561215a57600080fd5b600a8290556040805183815290517f4114f8ef80b6de2161db580cbefa14e1892d15d3ebe2062c9914e4a5773114a39181900360200190a15050600a5490565b60006121a7600b54610a7f565b905090565b600254600160a060020a03163314806121c857506121c86119f8565b15156121d357600080fd5b60048190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b6000848152600d602052604081206001810154600382015484011161233757604051600160a060020a0386169084156108fc029085906000818181858888f19350505050158015612263573d6000803e3d6000fd5b508281600301540181600301819055507fad58d18ea7292f887da6f15bb4f0badddaa33d169713d09cf49710acc7c3a5b986858786604051808581526020018060200184600160a060020a0316600160a060020a03168152602001838152602001828103825285818151815260200191508051906020019080838360005b838110156122f95781810151838201526020016122e1565b50505050905090810190601f1680156123265780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a15b50949350505050565b600254600160a060020a03163214806123635750600254600160a060020a031633145b151561236e57600080fd5b600254600160a060020a0316ff5b604080516080810182526000808252602082018190529181018290526060810191909152905600a165627a7a723058202fb499b0671f272703d23a44987b8c3eea18538337d5e28f0d8dc8fca16109ae0029`

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

// CanPublish is a free data retrieval call binding the contract method 0xcbcd4461.
//
// Solidity: function canPublish() constant returns(bool)
func (_BaseContent *BaseContentCaller) CanPublish(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "canPublish")
	return *ret0, err
}

// CanPublish is a free data retrieval call binding the contract method 0xcbcd4461.
//
// Solidity: function canPublish() constant returns(bool)
func (_BaseContent *BaseContentSession) CanPublish() (bool, error) {
	return _BaseContent.Contract.CanPublish(&_BaseContent.CallOpts)
}

// CanPublish is a free data retrieval call binding the contract method 0xcbcd4461.
//
// Solidity: function canPublish() constant returns(bool)
func (_BaseContent *BaseContentCallerSession) CanPublish() (bool, error) {
	return _BaseContent.Contract.CanPublish(&_BaseContent.CallOpts)
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

// GetAccessInfo is a free data retrieval call binding the contract method 0x38d0f504.
//
// Solidity: function getAccessInfo(uint8 level, bytes32[] custom_values, address[] stakeholders) constant returns(int8, uint256)
func (_BaseContent *BaseContentCaller) GetAccessInfo(opts *bind.CallOpts, level uint8, custom_values [][32]byte, stakeholders []common.Address) (int8, *big.Int, error) {
	var (
		ret0 = new(int8)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BaseContent.contract.Call(opts, out, "getAccessInfo", level, custom_values, stakeholders)
	return *ret0, *ret1, err
}

// GetAccessInfo is a free data retrieval call binding the contract method 0x38d0f504.
//
// Solidity: function getAccessInfo(uint8 level, bytes32[] custom_values, address[] stakeholders) constant returns(int8, uint256)
func (_BaseContent *BaseContentSession) GetAccessInfo(level uint8, custom_values [][32]byte, stakeholders []common.Address) (int8, *big.Int, error) {
	return _BaseContent.Contract.GetAccessInfo(&_BaseContent.CallOpts, level, custom_values, stakeholders)
}

// GetAccessInfo is a free data retrieval call binding the contract method 0x38d0f504.
//
// Solidity: function getAccessInfo(uint8 level, bytes32[] custom_values, address[] stakeholders) constant returns(int8, uint256)
func (_BaseContent *BaseContentCallerSession) GetAccessInfo(level uint8, custom_values [][32]byte, stakeholders []common.Address) (int8, *big.Int, error) {
	return _BaseContent.Contract.GetAccessInfo(&_BaseContent.CallOpts, level, custom_values, stakeholders)
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
// Solidity: function requestMap(uint256 ) constant returns(address originator, uint256 amountPaid, int8 status, uint256 settled)
func (_BaseContent *BaseContentCaller) RequestMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Originator common.Address
	AmountPaid *big.Int
	Status     int8
	Settled    *big.Int
}, error) {
	ret := new(struct {
		Originator common.Address
		AmountPaid *big.Int
		Status     int8
		Settled    *big.Int
	})
	out := ret
	err := _BaseContent.contract.Call(opts, out, "requestMap", arg0)
	return *ret, err
}

// RequestMap is a free data retrieval call binding the contract method 0x1a735f18.
//
// Solidity: function requestMap(uint256 ) constant returns(address originator, uint256 amountPaid, int8 status, uint256 settled)
func (_BaseContent *BaseContentSession) RequestMap(arg0 *big.Int) (struct {
	Originator common.Address
	AmountPaid *big.Int
	Status     int8
	Settled    *big.Int
}, error) {
	return _BaseContent.Contract.RequestMap(&_BaseContent.CallOpts, arg0)
}

// RequestMap is a free data retrieval call binding the contract method 0x1a735f18.
//
// Solidity: function requestMap(uint256 ) constant returns(address originator, uint256 amountPaid, int8 status, uint256 settled)
func (_BaseContent *BaseContentCallerSession) RequestMap(arg0 *big.Int) (struct {
	Originator common.Address
	AmountPaid *big.Int
	Status     int8
	Settled    *big.Int
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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseContent *BaseContentCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseContent *BaseContentSession) Version() ([32]byte, error) {
	return _BaseContent.Contract.Version(&_BaseContent.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseContent *BaseContentCallerSession) Version() ([32]byte, error) {
	return _BaseContent.Contract.Version(&_BaseContent.CallOpts)
}

// VersionHashes is a free data retrieval call binding the contract method 0x7ca8f618.
//
// Solidity: function versionHashes(uint256 ) constant returns(bytes32)
func (_BaseContent *BaseContentCaller) VersionHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseContent.contract.Call(opts, out, "versionHashes", arg0)
	return *ret0, err
}

// VersionHashes is a free data retrieval call binding the contract method 0x7ca8f618.
//
// Solidity: function versionHashes(uint256 ) constant returns(bytes32)
func (_BaseContent *BaseContentSession) VersionHashes(arg0 *big.Int) ([32]byte, error) {
	return _BaseContent.Contract.VersionHashes(&_BaseContent.CallOpts, arg0)
}

// VersionHashes is a free data retrieval call binding the contract method 0x7ca8f618.
//
// Solidity: function versionHashes(uint256 ) constant returns(bytes32)
func (_BaseContent *BaseContentCallerSession) VersionHashes(arg0 *big.Int) ([32]byte, error) {
	return _BaseContent.Contract.VersionHashes(&_BaseContent.CallOpts, arg0)
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
// Solidity: function commit(bytes32 _objectHash) returns()
func (_BaseContent *BaseContentTransactor) Commit(opts *bind.TransactOpts, _objectHash [32]byte) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "commit", _objectHash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 _objectHash) returns()
func (_BaseContent *BaseContentSession) Commit(_objectHash [32]byte) (*types.Transaction, error) {
	return _BaseContent.Contract.Commit(&_BaseContent.TransactOpts, _objectHash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 _objectHash) returns()
func (_BaseContent *BaseContentTransactorSession) Commit(_objectHash [32]byte) (*types.Transaction, error) {
	return _BaseContent.Contract.Commit(&_BaseContent.TransactOpts, _objectHash)
}

// GetAccessCharge is a paid mutator transaction binding the contract method 0x879fe48f.
//
// Solidity: function getAccessCharge(uint8 level, bytes32[] custom_values, address[] stakeholders) returns(uint256)
func (_BaseContent *BaseContentTransactor) GetAccessCharge(opts *bind.TransactOpts, level uint8, custom_values [][32]byte, stakeholders []common.Address) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "getAccessCharge", level, custom_values, stakeholders)
}

// GetAccessCharge is a paid mutator transaction binding the contract method 0x879fe48f.
//
// Solidity: function getAccessCharge(uint8 level, bytes32[] custom_values, address[] stakeholders) returns(uint256)
func (_BaseContent *BaseContentSession) GetAccessCharge(level uint8, custom_values [][32]byte, stakeholders []common.Address) (*types.Transaction, error) {
	return _BaseContent.Contract.GetAccessCharge(&_BaseContent.TransactOpts, level, custom_values, stakeholders)
}

// GetAccessCharge is a paid mutator transaction binding the contract method 0x879fe48f.
//
// Solidity: function getAccessCharge(uint8 level, bytes32[] custom_values, address[] stakeholders) returns(uint256)
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

// ProcessRequestPayment is a paid mutator transaction binding the contract method 0x0c6d3f93.
//
// Solidity: function processRequestPayment(uint256 request_ID, address payee, string label, uint256 amount) returns(bool)
func (_BaseContent *BaseContentTransactor) ProcessRequestPayment(opts *bind.TransactOpts, request_ID *big.Int, payee common.Address, label string, amount *big.Int) (*types.Transaction, error) {
	return _BaseContent.contract.Transact(opts, "processRequestPayment", request_ID, payee, label, amount)
}

// ProcessRequestPayment is a paid mutator transaction binding the contract method 0x0c6d3f93.
//
// Solidity: function processRequestPayment(uint256 request_ID, address payee, string label, uint256 amount) returns(bool)
func (_BaseContent *BaseContentSession) ProcessRequestPayment(request_ID *big.Int, payee common.Address, label string, amount *big.Int) (*types.Transaction, error) {
	return _BaseContent.Contract.ProcessRequestPayment(&_BaseContent.TransactOpts, request_ID, payee, label, amount)
}

// ProcessRequestPayment is a paid mutator transaction binding the contract method 0x0c6d3f93.
//
// Solidity: function processRequestPayment(uint256 request_ID, address payee, string label, uint256 amount) returns(bool)
func (_BaseContent *BaseContentTransactorSession) ProcessRequestPayment(request_ID *big.Int, payee common.Address, label string, amount *big.Int) (*types.Transaction, error) {
	return _BaseContent.Contract.ProcessRequestPayment(&_BaseContent.TransactOpts, request_ID, payee, label, amount)
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

// BaseContentCommitPendingIterator is returned from FilterCommitPending and is used to iterate over the raw logs and unpacked data for CommitPending events raised by the BaseContent contract.
type BaseContentCommitPendingIterator struct {
	Event *BaseContentCommitPending // Event containing the contract specifics and raw log

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
func (it *BaseContentCommitPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentCommitPending)
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
		it.Event = new(BaseContentCommitPending)
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
func (it *BaseContentCommitPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentCommitPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentCommitPending represents a CommitPending event raised by the BaseContent contract.
type BaseContentCommitPending struct {
	ObjectHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommitPending is a free log retrieval operation binding the contract event 0xa288d3a2aba7d5dec44e93ff46d2f1129e251695be2046de105f9a9c6feefcaa.
//
// Solidity: event CommitPending(bytes32 objectHash)
func (_BaseContent *BaseContentFilterer) FilterCommitPending(opts *bind.FilterOpts) (*BaseContentCommitPendingIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "CommitPending")
	if err != nil {
		return nil, err
	}
	return &BaseContentCommitPendingIterator{contract: _BaseContent.contract, event: "CommitPending", logs: logs, sub: sub}, nil
}

// WatchCommitPending is a free log subscription operation binding the contract event 0xa288d3a2aba7d5dec44e93ff46d2f1129e251695be2046de105f9a9c6feefcaa.
//
// Solidity: event CommitPending(bytes32 objectHash)
func (_BaseContent *BaseContentFilterer) WatchCommitPending(opts *bind.WatchOpts, sink chan<- *BaseContentCommitPending) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "CommitPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentCommitPending)
				if err := _BaseContent.contract.UnpackLog(event, "CommitPending", log); err != nil {
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

// BaseContentLogPaymentIterator is returned from FilterLogPayment and is used to iterate over the raw logs and unpacked data for LogPayment events raised by the BaseContent contract.
type BaseContentLogPaymentIterator struct {
	Event *BaseContentLogPayment // Event containing the contract specifics and raw log

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
func (it *BaseContentLogPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentLogPayment)
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
		it.Event = new(BaseContentLogPayment)
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
func (it *BaseContentLogPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentLogPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentLogPayment represents a LogPayment event raised by the BaseContent contract.
type BaseContentLogPayment struct {
	RequestID *big.Int
	Label     string
	Payee     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogPayment is a free log retrieval operation binding the contract event 0xad58d18ea7292f887da6f15bb4f0badddaa33d169713d09cf49710acc7c3a5b9.
//
// Solidity: event LogPayment(uint256 requestID, string label, address payee, uint256 amount)
func (_BaseContent *BaseContentFilterer) FilterLogPayment(opts *bind.FilterOpts) (*BaseContentLogPaymentIterator, error) {

	logs, sub, err := _BaseContent.contract.FilterLogs(opts, "LogPayment")
	if err != nil {
		return nil, err
	}
	return &BaseContentLogPaymentIterator{contract: _BaseContent.contract, event: "LogPayment", logs: logs, sub: sub}, nil
}

// WatchLogPayment is a free log subscription operation binding the contract event 0xad58d18ea7292f887da6f15bb4f0badddaa33d169713d09cf49710acc7c3a5b9.
//
// Solidity: event LogPayment(uint256 requestID, string label, address payee, uint256 amount)
func (_BaseContent *BaseContentFilterer) WatchLogPayment(opts *bind.WatchOpts, sink chan<- *BaseContentLogPayment) (event.Subscription, error) {

	logs, sub, err := _BaseContent.contract.WatchLogs(opts, "LogPayment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentLogPayment)
				if err := _BaseContent.contract.UnpackLog(event, "LogPayment", log); err != nil {
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
const BaseContentSpaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_locator\",\"type\":\"bytes\"}],\"name\":\"submitNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"canNodePublish\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"address_KMS\",\"type\":\"address\"}],\"name\":\"createLibrary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numActiveNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeNodeLocators\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeNodeAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_factory\",\"type\":\"address\"}],\"name\":\"setFactory\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeAddr\",\"type\":\"address\"},{\"name\":\"_locator\",\"type\":\"bytes\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingNodeLocators\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingNodeAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content_space_description\",\"type\":\"string\"}],\"name\":\"setDescription\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeAddr\",\"type\":\"address\"}],\"name\":\"removeNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createContentType\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"engageAccountLibrary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canPublish\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeAddr\",\"type\":\"address\"}],\"name\":\"approveNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"objectHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"object_hash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accessRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numPendingNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"content_space_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"locator\",\"type\":\"bytes\"}],\"name\":\"NodeSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"locator\",\"type\":\"bytes\"}],\"name\":\"NodeApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentTypeAddress\",\"type\":\"address\"}],\"name\":\"CreateContentType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"libraryAddress\",\"type\":\"address\"}],\"name\":\"CreateLibrary\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"groupAddress\",\"type\":\"address\"}],\"name\":\"CreateGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountAddress\",\"type\":\"address\"}],\"name\":\"EngageAccountLibrary\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"SetFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"Commit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"UpdateRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AccessRequest\",\"type\":\"event\"}]"

// BaseContentSpaceBin is the compiled bytecode used for deploying new contracts.
const BaseContentSpaceBin = `0x60806040527f41636365737369626c6532303139303232323133353930304d4c0000000000006000557f4f776e61626c6532303139303331353134313530304d4c0000000000000000006001557f4564697461626c6532303139303331353134313830304d4c00000000000000006004557f42617365436f6e74656e74537061636532303139303232313131343130304d4c600655348015620000a157600080fd5b50604051620074b3380380620074b38339810160405280516002805432600160a060020a03199182168117909255600380549091169091179055018051620000f190600790602084019062000141565b50620000fc620001c6565b604051809103906000f08015801562000119573d6000803e3d6000fd5b5060098054600160a060020a031916600160a060020a039290921691909117905550620001f7565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200018457805160ff1916838001178555620001b4565b82800160010185558215620001b4579182015b82811115620001b457825182559160200191906001019062000197565b50620001c2929150620001d7565b5090565b6040516158448062001c6f83390190565b620001f491905b80821115620001c25760008155600101620001de565b90565b611a6880620002076000396000f3006080604052600436106101715763ffffffff60e060020a60003504166302d05d3f811461017357806306fdde03146101a4578063160eee741461022e57806326683e141461028757806340b89f06146102bc57806341c0e1b5146102dd57806343f59ec7146102f25780635272ae171461031957806352f82dd81461033157806354fd4d5014610349578063575185ed1461035e5780635bb478081461037357806364f0f0501461039457806369e30ff8146103fb5780636be9514c146104135780636d2e4b1b1461042b5780637284e4161461044c5780638da5cb5b1461046157806390c3f38f14610476578063b2b99ec9146104cf578063b8cfaf05146104f0578063c287e0ed14610505578063c45a01551461051a578063c82710c11461052f578063cbcd446114610544578063dd4c97a014610559578063e02dd9c21461057a578063f14fcbc81461058f578063f1551887146105a7578063f2fde38b146105bc578063f41a1587146105dd575b005b34801561017f57600080fd5b506101886105f2565b60408051600160a060020a039092168252519081900360200190f35b3480156101b057600080fd5b506101b9610601565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101f35781810151838201526020016101db565b50505050905090810190601f1680156102205780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561023a57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261017194369492936024939284019190819084018382808284375094975061068f9650505050505050565b34801561029357600080fd5b506102a8600160a060020a0360043516610a42565b604080519115158252519081900360200190f35b3480156102c857600080fd5b50610188600160a060020a0360043516610aa0565b3480156102e957600080fd5b50610171610b7b565b3480156102fe57600080fd5b50610307610bb7565b60408051918252519081900360200190f35b34801561032557600080fd5b506101b9600435610bbe565b34801561033d57600080fd5b50610188600435610c32565b34801561035557600080fd5b50610307610c5a565b34801561036a57600080fd5b50610188610c60565b34801561037f57600080fd5b50610171600160a060020a0360043516610d2a565b3480156103a057600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610171958335600160a060020a0316953695604494919390910191908190840183828082843750949750610d7a9650505050505050565b34801561040757600080fd5b506101b9600435610f7c565b34801561041f57600080fd5b50610188600435610f8a565b34801561043757600080fd5b50610171600160a060020a0360043516610f98565b34801561045857600080fd5b506101b9610fe6565b34801561046d57600080fd5b50610188611041565b34801561048257600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101719436949293602493928401919081908401838280828437509497506110509650505050505050565b3480156104db57600080fd5b50610171600160a060020a0360043516611095565b3480156104fc57600080fd5b5061018861111d565b34801561051157600080fd5b506101716111e4565b34801561052657600080fd5b50610188611249565b34801561053b57600080fd5b50610188611258565b34801561055057600080fd5b506102a8611292565b34801561056557600080fd5b50610171600160a060020a0360043516611297565b34801561058657600080fd5b506103076114dd565b34801561059b57600080fd5b506101716004356114e3565b3480156105b357600080fd5b506102a8611545565b3480156105c857600080fd5b50610171600160a060020a0360043516611576565b3480156105e957600080fd5b506103076115db565b600254600160a060020a031681565b6007805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106875780601f1061065c57610100808354040283529160200191610687565b820191906000526020600020905b81548152906001019060200180831161066a57829003601f168201915b505050505081565b6107c5600c8054806020026020016040519081016040528092919081815260200182805480156106e857602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116106ca575b5050505050600d805480602002602001604051908101604052809291908181526020016000905b828210156107ba5760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156107a65780601f1061077b576101008083540402835291602001916107a6565b820191906000526020600020905b81548152906001019060200180831161078957829003601f168201915b50505050508152602001906001019061070f565b5050505033846115e1565b156107cf57600080fd5b6108fa600a80548060200260200160405190810160405280929190818152602001828054801561082857602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161080a575b5050505050600b805480602002602001604051908101604052809291908181526020016000905b828210156107ba5760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156108e65780601f106108bb576101008083540402835291602001916108e6565b820191906000526020600020905b8154815290600101906020018083116108c957829003601f168201915b50505050508152602001906001019061084f565b1561090457600080fd5b600c5460641161091357600080fd5b600d8054600181018083556000929092528251610957917fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb501906020850190611885565b5050600c805460018101825560009182527fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c7018054600160a060020a0319163390811790915560408051828152602081810183815286519383019390935285517fae5645569f32b946f7a747113c64094a29a6b84c5ddf55816ef4381ce8a3a46d958794926060850192908601918190849084905b83811015610a045781810151838201526020016109ec565b50505050905090810190601f168015610a315780820380516001836020036101000a031916815260200191505b50935050505060405180910390a150565b6000805b600a54811015610a955782600160a060020a0316600a82815481101515610a6957fe5b600091825260209091200154600160a060020a03161415610a8d5760019150610a9a565b600101610a46565b600091505b50919050565b600954604080517f40b89f06000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151600093849316916340b89f0691602480830192602092919082900301818787803b158015610b0957600080fd5b505af1158015610b1d573d6000803e3d6000fd5b505050506040513d6020811015610b3357600080fd5b505160408051600160a060020a038316815290519192507f473c07a6d0228c4fb8fe2be3b4617c3b5fb7c0f8cd9ba4b67e8631844b9b6571919081900360200190a192915050565b600354600160a060020a0316321480610b9e5750600354600160a060020a031633145b1515610ba957600080fd5b600354600160a060020a0316ff5b600b545b90565b600b805482908110610bcc57fe5b600091825260209182902001805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152935090918301828280156106875780601f1061065c57610100808354040283529160200191610687565b600a805482908110610c4057fe5b600091825260209091200154600160a060020a0316905081565b60065481565b600080600960009054906101000a9004600160a060020a0316600160a060020a031663575185ed6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610cb657600080fd5b505af1158015610cca573d6000803e3d6000fd5b505050506040513d6020811015610ce057600080fd5b505160408051600160a060020a038316815290519192507fa3b1fe71ae61bad8cffa485b230e24e518938f76182a30fa0d9979e7237ad159919081900360200190a18091505b5090565b600354600160a060020a0316321480610d4d5750600354600160a060020a031633145b1515610d5857600080fd5b60098054600160a060020a031916600160a060020a0392909216919091179055565b600354600160a060020a0316321480610d9d5750600354600160a060020a031633145b1515610da857600080fd5b610ede600a805480602002602001604051908101604052809291908181526020018280548015610e0157602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610de3575b5050505050600b805480602002602001604051908101604052809291908181526020016000905b82821015610ed35760008481526020908190208301805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815292830182828015610ebf5780601f10610e9457610100808354040283529160200191610ebf565b820191906000526020600020905b815481529060010190602001808311610ea257829003601f168201915b505050505081526020019060010190610e28565b5050505084846115e1565b15610ee857600080fd5b600a805460018082019092557fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8018054600160a060020a031916600160a060020a038516179055600b805491820180825560009190915282519091610f76917f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db9909101906020850190611885565b50505050565b600d805482908110610bcc57fe5b600c805482908110610c4057fe5b600254600160a060020a03163214610faf57600080fd5b600160a060020a0381161515610fc457600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b6008805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106875780601f1061065c57610100808354040283529160200191610687565b600354600160a060020a031681565b600354600160a060020a03163214806110735750600354600160a060020a031633145b151561107e57600080fd5b8051611091906008906020840190611885565b5050565b600354600090600160a060020a03163214806110bb5750600354600160a060020a031633145b15156110c657600080fd5b5060005b600a548110156110915781600160a060020a0316600a828154811015156110ed57fe5b600091825260209091200154600160a060020a031614156111155761111581600a600b61172f565b6001016110ca565b600080600960009054906101000a9004600160a060020a0316600160a060020a031663b8cfaf056040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561117357600080fd5b505af1158015611187573d6000803e3d6000fd5b505050506040513d602081101561119d57600080fd5b505160408051600160a060020a038316815290519192507f9e69777f30c55126be256664fa7beff4b796ac32ebceab94df5071b0148017f8919081900360200190a1919050565b600354600160a060020a03163214806112075750600354600160a060020a031633145b151561121257600080fd5b60055460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600954600160a060020a031681565b6040805132815290516000917f53ce35a7383a3ea3f695bdf0f87d7e5485ba816b382673e849bfdd24e7f5e3ca919081900360200190a190565b600090565b6003546000908190600160a060020a03163214806112bf5750600354600160a060020a031633145b15156112ca57600080fd5b5060009050805b600c548110156114cc5782600160a060020a0316600c828154811015156112f457fe5b600091825260209091200154600160a060020a031614156114c457600a600c8281548110151561132057fe5b6000918252602080832090910154835460018101855593835291209091018054600160a060020a031916600160a060020a03909216919091179055600d8054600b91908390811061136d57fe5b600091825260208083208454600181810180885596865292909420920180546113b194939093019290916002610100918316159190910260001901909116046118ff565b50507fd644c8164f225d3b7fdbcc404f279bb1e823ef0d93f88dd4b24e85d0e7bc6a54600c828154811015156113e357fe5b600091825260209091200154600d8054600160a060020a03909216918490811061140957fe5b600091825260209182902060408051600160a060020a038616815293840181815291909201805460026000196101006001841615020190911604928401839052929160608301908490801561149f5780601f106114745761010080835404028352916020019161149f565b820191906000526020600020905b81548152906001019060200180831161148257829003601f168201915b5050935050505060405180910390a16114bb81600c600d61172f565b600191506114cc565b6001016112d1565b8115156114d857600080fd5b505050565b60055481565b600354600160a060020a03163314806114ff57506114ff611292565b151561150a57600080fd5b60058190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b6040516000907fed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88908290a150600190565b600354600160a060020a03163214806115995750600354600160a060020a031633145b15156115a457600080fd5b600160a060020a03811615156115b957600080fd5b60038054600160a060020a031916600160a060020a0392909216919091179055565b600d5490565b600080845186511415156115f457600080fd5b5060005b855181101561172157826040518082805190602001908083835b602083106116315780518252601f199092019160209182019101611612565b5181516020939093036101000a600019018019909116921691909117905260405192018290039091208851909350889250849150811061166d57fe5b906020019060200201516040518082805190602001908083835b602083106116a65780518252601f199092019160209182019101611687565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916148061170b575083600160a060020a031686828151811015156116f657fe5b90602001906020020151600160a060020a0316145b156117195760019150611726565b6001016115f8565b600091505b50949350505050565b81548310801561173f5750805483105b151561174a57600080fd5b81546000190183146118115780548190600019810190811061176857fe5b90600052602060002001818481548110151561178057fe5b9060005260206000200190805460018160011615610100020316600290046117a99291906118ff565b508154829060001981019081106117bc57fe5b6000918252602090912001548254600160a060020a03909116908390859081106117e257fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055505b80548190600019810190811061182357fe5b9060005260206000200160006118399190611974565b80546118498260001983016119bb565b5081548290600019810190811061185c57fe5b60009182526020909120018054600160a060020a03191690558154610f768360001983016119df565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106118c657805160ff19168380011785556118f3565b828001600101855582156118f3579182015b828111156118f35782518255916020019190600101906118d8565b50610d269291506119ff565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061193857805485556118f3565b828001600101855582156118f357600052602060002091601f016020900482015b828111156118f3578254825591600101919060010190611959565b50805460018160011615610100020316600290046000825580601f1061199a57506119b8565b601f0160209004906000526020600020908101906119b891906119ff565b50565b8154818355818111156114d8576000838152602090206114d8918101908301611a19565b8154818355818111156114d8576000838152602090206114d89181019083015b610bbb91905b80821115610d265760008155600101611a05565b610bbb91905b80821115610d26576000611a338282611974565b50600101611a1f5600a165627a7a72305820ab33d3b2c465d8b2df8b91d91ee703086783042816c9ffff4bce4c8d60080679002960806040527f4f776e61626c6532303139303331353134313530304d4c0000000000000000006000557f42617365466163746f727932303139303331393139353030304d4c000000000060035560018054600160a060020a0319908116329081179092556002805490911690911790556157c68061007e6000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f811461009a57806340b89f06146100cb57806341c0e1b5146100ec57806354fd4d5014610101578063575185ed146101285780636d2e4b1b1461013d5780638da5cb5b1461015e578063b8cfaf0514610173578063f2fde38b14610188575b005b3480156100a657600080fd5b506100af6101a9565b60408051600160a060020a039092168252519081900360200190f35b3480156100d757600080fd5b506100af600160a060020a03600435166101b8565b3480156100f857600080fd5b506100986101fe565b34801561010d57600080fd5b5061011661023a565b60408051918252519081900360200190f35b34801561013457600080fd5b506100af610240565b34801561014957600080fd5b50610098600160a060020a036004351661027e565b34801561016a57600080fd5b506100af6102d9565b34801561017f57600080fd5b506100af6102e8565b34801561019457600080fd5b50610098600160a060020a03600435166102f3565b600154600160a060020a031681565b600081336101c4610365565b600160a060020a03928316815291166020820152604080519182900301906000f0801580156101f7573d6000803e3d6000fd5b5092915050565b600254600160a060020a03163214806102215750600254600160a060020a031633145b151561022c57600080fd5b600254600160a060020a0316ff5b60035481565b60003361024b610375565b600160a060020a03909116815260405190819003602001906000f080158015610278573d6000803e3d6000fd5b50905090565b600154600160a060020a0316321461029557600080fd5b600160a060020a03811615156102aa57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a031681565b60003361024b610385565b600254600160a060020a03163214806103165750600254600160a060020a031633145b151561032157600080fd5b600160a060020a038116151561033657600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6040516146e48061039683390190565b6040516107a980614a7a83390190565b60405161057880615223833901905600608060408181527f41636365737369626c6532303139303232323133353930304d4c00000000000060009081557f4f776e61626c6532303139303331353134313530304d4c0000000000000000006001557f4564697461626c6532303139303331353134313830304d4c00000000000000006004557f426173654c69627261727932303139303331383130313330304d4c0000000000600655600c819055600d819055600e819055600f819055601355806146e48339810160405280516020909101516002805432600160a060020a03199182168117909255600380548216909217909155600780548216600160a060020a0393841617905560128054909116919092161790556145cf806101156000396000f3006080604052600436106101ea5763ffffffff60e060020a60003504166302d05d3f81146101ec5780630eaec2c51461021d5780630f58a7861461025257806316308394146102795780631b969895146102a05780631cdbee5a146102c15780631d0f4351146102e25780631e35d8fa1461030357806321770a84146103245780632393553b1461033957806326683e141461035157806329d002191461037257806329dedde5146103935780632c11f392146103b457806332eaf21b146103cc578063386493e0146103e157806341c0e1b514610402578063470750bb1461041757806349102e611461042c57806354fd4d501461044157806363dab9d414610456578063679a9a3c1461046e5780636d2e4b1b1461048f57806387e86b2c146104b05780638cb13c2e1461051b5780638da5cb5b14610533578063952e464b1461054857806395a078e814610560578063991a3a7c14610581578063af570c0414610599578063c287e0ed146105ae578063c65bcbe2146105c3578063c9e8e72d146105d8578063cbcd4461146105f9578063dc3c29c01461060e578063e02dd9c21461062f578063e5538fd214610644578063e8de515f14610659578063f14fcbc81461067a578063f155188714610692578063f2fde38b146106a7578063fd089196146106c8575b005b3480156101f857600080fd5b506102016106e9565b60408051600160a060020a039092168252519081900360200190f35b34801561022957600080fd5b5061023e600160a060020a03600435166106f8565b604080519115158252519081900360200190f35b34801561025e57600080fd5b506101ea600160a060020a0360043581169060243516610803565b34801561028557600080fd5b5061028e610978565b60408051918252519081900360200190f35b3480156102ac57600080fd5b5061023e600160a060020a036004351661097e565b3480156102cd57600080fd5b50610201600160a060020a0360043516610b1f565b3480156102ee57600080fd5b506101ea600160a060020a0360043516610b3a565b34801561030f57600080fd5b50610201600160a060020a0360043516610c4d565b34801561033057600080fd5b5061028e610e0e565b34801561034557600080fd5b50610201600435610e14565b34801561035d57600080fd5b5061023e600160a060020a0360043516610e3c565b34801561037e57600080fd5b5061023e600160a060020a0360043516610edb565b34801561039f57600080fd5b5061023e600160a060020a0360043516610fe6565b3480156103c057600080fd5b50610201600435611041565b3480156103d857600080fd5b5061020161104f565b3480156103ed57600080fd5b5061023e600160a060020a036004351661105e565b34801561040e57600080fd5b506101ea6111f4565b34801561042357600080fd5b5061028e611230565b34801561043857600080fd5b5061023e611236565b34801561044d57600080fd5b5061028e61144d565b34801561046257600080fd5b50610201600435611453565b34801561047a57600080fd5b506101ea600160a060020a03600435166114a0565b34801561049b57600080fd5b506101ea600160a060020a03600435166115b3565b3480156104bc57600080fd5b50604080516020600460443581810135601f810184900484028501840190955284845261023e948235600160a060020a031694602480351515953695946064949201919081908401838280828437509497506116019650505050505050565b34801561052757600080fd5b50610201600435611921565b34801561053f57600080fd5b5061020161192f565b34801561055457600080fd5b5061020160043561193e565b34801561056c57600080fd5b5061023e600160a060020a036004351661194c565b34801561058d57600080fd5b50610201600435611a49565b3480156105a557600080fd5b50610201611a57565b3480156105ba57600080fd5b506101ea611a66565b3480156105cf57600080fd5b5061028e611acb565b3480156105e457600080fd5b506101ea600160a060020a0360043516611ad1565b34801561060557600080fd5b5061023e611b21565b34801561061a57600080fd5b506101ea600160a060020a0360043516611b26565b34801561063b57600080fd5b5061028e611c39565b34801561065057600080fd5b5061028e611c3f565b34801561066557600080fd5b5061023e600160a060020a0360043516611c45565b34801561068657600080fd5b506101ea600435611ddb565b34801561069e57600080fd5b5061023e611e3d565b3480156106b357600080fd5b506101ea600160a060020a0360043516611ea0565b3480156106d457600080fd5b5061023e600160a060020a0360043516611f05565b600254600160a060020a031681565b6000806000806000600c546000141561071457600194506107fa565b5060005b600c548110156107f557600880548290811061073057fe5b600091825260209091200154600160a060020a0316935083156107ed5783915081600160a060020a03166395a078e8876040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b1580156107ab57600080fd5b505af11580156107bf573d6000803e3d6000fd5b505050506040513d60208110156107d557600080fd5b50519250600183151514156107ed57600194506107fa565b600101610718565b600094505b50505050919050565b600354600160a060020a03163214806108265750600354600160a060020a031633145b151561083157600080fd5b600160a060020a038083166000908152601060205260409020541615801561085f575061085d82610fe6565b155b1561090b57600b54600f5410156108b55781600b600f5481548110151561088257fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550610901565b600b80546001810182556000919091527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db9018054600160a060020a031916600160a060020a0384161790555b600f805460010190555b600160a060020a038281166000818152601060209081526040918290208054600160a060020a0319169486169485179055815192835282019290925281517f280016f7418306a55542432120fd1a239ef9fcc1a92694d8d44ca76be0249ea7929181900390910190a15050565b60135481565b6003546000908190600160a060020a03163214806109a65750600354600160a060020a031633145b15156109b157600080fd5b5060005b600d54811015610b145782600160a060020a03166009828154811015156109d857fe5b600091825260209091200154600160a060020a03161415610b0c576009805482908110610a0157fe5b60009182526020909120018054600160a060020a0319169055600d54600019018114610abd5760096001600d5403815481101515610a3b57fe5b60009182526020909120015460098054600160a060020a039092169183908110610a6157fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a0316021790555060096001600d5403815481101515610aa357fe5b60009182526020909120018054600160a060020a03191690555b600d805460001901905560408051600160a060020a038516815290517fdf9d78c5635b72b709c85300a786eb7238acbe5bffe01c60c16464e45c6eb6eb9181900360200190a160019150610b19565b6001016109b5565b600091505b50919050565b601060205260009081526040902054600160a060020a031681565b600354600160a060020a0316321480610b5d5750600354600160a060020a031633145b1515610b6857600080fd5b600a54600e541015610bb95780600a600e54815481101515610b8657fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550610c05565b600a80546001810182556000919091527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8018054600160a060020a031916600160a060020a0383161790555b600e8054600101905560408051600160a060020a038316815290517f3a94857e4393737f73edb175a7d0c195c7f635d9ae995e12740616ec55c9d4119181900360200190a150565b600080610c59326106f8565b1515610c6457600080fd5b600f5415610c8057610c7583610fe6565b1515610c8057600080fd5b82610c896120b2565b600160a060020a03909116815260405190819003602001906000f080158015610cb6573d6000803e3d6000fd5b50601254604080517fc9e8e72d000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015290519293509083169163c9e8e72d9160248082019260009290919082900301818387803b158015610d2157600080fd5b505af1158015610d35573d6000803e3d6000fd5b50505050600160a060020a038381166000908152601060205260408082205481517fe5385303000000000000000000000000000000000000000000000000000000008152908416600482015290519284169263e53853039260248084019391929182900301818387803b158015610dab57600080fd5b505af1158015610dbf573d6000803e3d6000fd5b505060408051600160a060020a0380861682528716602082015281517f3981e74ab81857b375ec391a4f7c31ee89462cd927de6d8fbdb98f77da009c569450908190039091019150a192915050565b600d5481565b6008805482908110610e2257fe5b600091825260209091200154600160a060020a0316905081565b600754604080517f26683e14000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151600093929092169182916326683e1491602480830192602092919082900301818887803b158015610ea857600080fd5b505af1158015610ebc573d6000803e3d6000fd5b505050506040513d6020811015610ed257600080fd5b50519392505050565b6003546000908190819081908190600160a060020a0387811691161415610f0557600194506107fa565b5060005b600d548110156107f5576009805482908110610f2157fe5b600091825260209091200154600160a060020a031693508315610fde5783915081600160a060020a03166395a078e8876040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b158015610f9c57600080fd5b505af1158015610fb0573d6000803e3d6000fd5b505050506040513d6020811015610fc657600080fd5b5051925060018315151415610fde57600194506107fa565b600101610f09565b600080805b600f548110156110365783600160a060020a0316600b8281548110151561100e57fe5b600091825260209091200154600160a060020a0316141561102e57600191505b600101610feb565b8192505b5050919050565b600a805482908110610e2257fe5b601254600160a060020a031681565b6003546000908190600160a060020a03163214806110865750600354600160a060020a031633145b151561109157600080fd5b5060005b600c54811015610b145782600160a060020a03166008828154811015156110b857fe5b600091825260209091200154600160a060020a031614156111ec5760088054829081106110e157fe5b60009182526020909120018054600160a060020a0319169055600c5460001901811461119d5760086001600c540381548110151561111b57fe5b60009182526020909120015460088054600160a060020a03909216918390811061114157fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a0316021790555060086001600c540381548110151561118357fe5b60009182526020909120018054600160a060020a03191690555b600c805460001901905560408051600160a060020a038516815290517fbbd97daa1862eb12f77ed128a557406737cee07b131b1e2d7140dff2005e197c9181900360200190a160019150610b19565b600101611095565b600354600160a060020a03163214806112175750600354600160a060020a031633145b151561122257600080fd5b600354600160a060020a0316ff5b600c5481565b600d546000903390819015156113165780600160a060020a0316638280dd8f60006040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b15801561129057600080fd5b505af11580156112a4573d6000803e3d6000fd5b505050506040513d60208110156112ba57600080fd5b505060408051600160a060020a038416815260016020820152606081830181905260009082015290517f70234ce475fee4ab40e5e55cf533f67f12b47ef4c860e62dd7affa84ead4b4429181900360a00190a160019250611448565b600160a060020a0382166000908152601460205260409020541561133d5760009250611448565b601154601354101561138e5781601160135481548110151561135b57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055506113da565b601180546001810182556000919091527f31ecc21a745e3968a04e9570e4425bc18fa8019c68028196b546d1669c200c68018054600160a060020a031916600160a060020a0384161790555b60138054600160a060020a0384166000818152601460209081526040918290206001948501905584549093019093558251908152329181019190915281517f0588a34cf0de4e025d359c89ca4bacbcbf175440909952d91c814412d9da996a929181900390910190a1600192505b505090565b60065481565b60006013546000148061146857508160135411155b156114755750600061149b565b601180548390811061148357fe5b600091825260209091200154600160a060020a031690505b919050565b600354600160a060020a03163214806114c35750600354600160a060020a031633145b15156114ce57600080fd5b600854600c54101561151f57806008600c548154811015156114ec57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a0316021790555061156b565b600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3018054600160a060020a031916600160a060020a0383161790555b600c8054600101905560408051600160a060020a038316815290517f218673669018c25b89bfbf1b58d0075e37c8847ef16e707b92355b7833e97d619181900360200190a150565b600254600160a060020a031632146115ca57600080fd5b600160a060020a03811615156115df57600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b60008060008060008061161332610edb565b151560011461162157600080fd5b600160a060020a0389166000908152601460205260409020546011805460001990920196508a9550908690811061165457fe5b600091825260208083209091018054600160a060020a031916905560138054600019018155600160a060020a038c16835260149091526040822091909155548510156117495760116013548154811015156116ab57fe5b60009182526020909120015460118054600160a060020a0390921694508491879081106116d457fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550601160135481548110151561171357fe5b600091825260208083209091018054600160a060020a0319169055600160a060020a038516825260149052604090206001860190555b83600160a060020a03166327c1c21d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561178757600080fd5b505af115801561179b573d6000803e3d6000fd5b505050506040513d60208110156117b157600080fd5b50519150600082131561191057600188151514156117d1575060006117d6565b506000195b83600160a060020a0316638280dd8f826040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b15801561181f57600080fd5b505af1158015611833573d6000803e3d6000fd5b505050506040513d602081101561184957600080fd5b505060408051600160a060020a038b16815289151560208281019190915260609282018381528a519383019390935289517f70234ce475fee4ab40e5e55cf533f67f12b47ef4c860e62dd7affa84ead4b442938d938d938d93919291608084019185019080838360005b838110156118cb5781810151838201526020016118b3565b50505050905090810190601f1680156118f85780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a160019550611915565b600095505b50505050509392505050565b6011805482908110610e2257fe5b600354600160a060020a031681565b6009805482908110610e2257fe5b6000806000806000600e546000141561196857600194506107fa565b5060005b600e548110156107f557600a80548290811061198457fe5b600091825260209091200154600160a060020a031693508315611a415783915081600160a060020a03166395a078e8876040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b1580156119ff57600080fd5b505af1158015611a13573d6000803e3d6000fd5b505050506040513d6020811015611a2957600080fd5b5051925060018315151415611a4157600194506107fa565b60010161196c565b600b805482908110610e2257fe5b600754600160a060020a031681565b600354600160a060020a0316321480611a895750600354600160a060020a031633145b1515611a9457600080fd5b60055460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600f5481565b600354600160a060020a0316321480611af45750600354600160a060020a031633145b1515611aff57600080fd5b60128054600160a060020a031916600160a060020a0392909216919091179055565b600090565b600354600160a060020a0316321480611b495750600354600160a060020a031633145b1515611b5457600080fd5b600954600d541015611ba557806009600d54815481101515611b7257fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550611bf1565b600980546001810182556000919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af018054600160a060020a031916600160a060020a0383161790555b600d8054600101905560408051600160a060020a038316815290517f1b88a571cc8ac2e87512f05648e79d184f5cc0cbb2889bc487c41f8b9a3202eb9181900360200190a150565b60055481565b600e5481565b6003546000908190600160a060020a0316321480611c6d5750600354600160a060020a031633145b1515611c7857600080fd5b5060005b600e54811015610b145782600160a060020a0316600a82815481101515611c9f57fe5b600091825260209091200154600160a060020a03161415611dd357600a805482908110611cc857fe5b60009182526020909120018054600160a060020a0319169055600e54600019018114611d8457600a6001600e5403815481101515611d0257fe5b600091825260209091200154600a8054600160a060020a039092169183908110611d2857fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550600a6001600e5403815481101515611d6a57fe5b60009182526020909120018054600160a060020a03191690555b600e805460001901905560408051600160a060020a038516815290517fc5224c4118417a068eeac7d714e6d8af6f99ec3fb611bc965185460b0e38f0819181900360200190a160019150610b19565b600101611c7c565b600354600160a060020a0316331480611df75750611df7611b21565b1515611e0257600080fd5b60058190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b6000611e483261194c565b80611e575750611e57326106f8565b80611e665750611e6632610edb565b1515611e7157600080fd5b6040517fed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e8890600090a150600190565b600354600160a060020a0316321480611ec35750600354600160a060020a031633145b1515611ece57600080fd5b600160a060020a0381161515611ee357600080fd5b60038054600160a060020a031916600160a060020a0392909216919091179055565b60035460009081908190600160a060020a0316321480611f2f5750600354600160a060020a031633145b1515611f3a57600080fd5b5050600f546000190160005b600f548110156120a85783600160a060020a0316600b82815481101515611f6957fe5b600091825260209091200154600160a060020a031614156120a057600b805482908110611f9257fe5b60009182526020909120018054600160a060020a031916905580821461203757600b805483908110611fc057fe5b600091825260209091200154600b8054600160a060020a039092169183908110611fe657fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600b80548390811061201d57fe5b60009182526020909120018054600160a060020a03191690555b600f829055600160a060020a0384166000818152601060209081526040918290208054600160a060020a0319169055815192835290517fd41375b9d347dfe722f90a780731abd23b7855f9cf14ea7063c4cab5f9ae58e29281900390910190a16001925061103a565b600101611f46565b5060009392505050565b6040516124e1806120c3833901905600608060408190527f4f776e61626c6532303139303331353134313530304d4c00000000000000000060009081557f4564697461626c6532303139303331353134313830304d4c00000000000000006003557f42617365436f6e74656e7432303139303331353137353130304d4c0000000000600555600c556020806124e1833981016040818152915160018054600160a060020a031990811632908117909255600280548216909217909155600980548216331790819055600019600b5560068054600160a060020a038086169190941617905516825291517fc3decc188980e855666b70498ca85e8fa284d97d30483d828fa126f7303d7d199181900360200190a1506123cf806101126000396000f3006080604052600436106101a75763ffffffff60e060020a60003504166217de9881146101a957806302d05d3f146101d0578063075d4782146102015780630c6d3f931461021d5780631a735f18146102885780632310167f146102d557806327c1c21d146102ea57806332eaf21b146102ff57806336ebffca14610314578063388642841461032957806338d0f5041461034157806341c0e1b5146103fa5780634dd707881461040f5780635267db441461042457806354fd4d501461043c5780635cc4aa9b1461045157806364ade32b146104625780636d2e4b1b146104775780637ca8f618146104985780638280dd8f146104b0578063879fe48f146104c85780638da5cb5b1461055e5780638f77920114610573578063a1ff106e14610588578063b816f5131461068c578063c287e0ed146106a1578063c9e8e72d146106b6578063cbcd4461146106d7578063d810f8c8146106ec578063e02dd9c214610701578063e538530314610716578063ee56d76714610737578063f14fcbc8146107d7578063f2fde38b146107ef578063f4d9bae814610810578063f81ab0ae14610828575b005b3480156101b557600080fd5b506101be61083d565b60408051918252519081900360200190f35b3480156101dc57600080fd5b506101e5610861565b60408051600160a060020a039092168252519081900360200190f35b610209610870565b604080519115158252519081900360200190f35b34801561022957600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526102099482359460248035600160a060020a03169536959460649492019190819084018382808284375094975050933594506109d19350505050565b34801561029457600080fd5b506102a0600435610a18565b60408051600160a060020a0390951685526020850193909352600091820b90910b838301526060830152519081900360800190f35b3480156102e157600080fd5b506101e5610a4c565b3480156102f657600080fd5b506101be610a5b565b34801561030b57600080fd5b506101e5610a61565b34801561032057600080fd5b506101e5610a70565b34801561033557600080fd5b506101be600435610a7f565b34801561034d57600080fd5b506040805160206004602480358281013584810280870186019097528086526103d796843560ff1696369660449591949091019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750949750610bd19650505050505050565b604051808360000b60000b81526020018281526020019250505060405180910390f35b34801561040657600080fd5b506101a7610d5b565b34801561041b57600080fd5b506101be610e40565b34801561043057600080fd5b506101be600435610e64565b34801561044857600080fd5b506101be610ef9565b610209600435602435604435610eff565b34801561046e57600080fd5b506101be6111c0565b34801561048357600080fd5b506101a7600160a060020a03600435166111c6565b3480156104a457600080fd5b506101be600435611221565b3480156104bc57600080fd5b506101be600435611240565b3480156104d457600080fd5b506040805160206004602480358281013584810280870186019097528086526101be96843560ff1696369660449591949091019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506113b79650505050505050565b34801561056a57600080fd5b506101e5611421565b34801561057f57600080fd5b506101be611430565b60408051602060046024803582810135601f810185900485028601850190965285855261020995833560ff1695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506114369650505050505050565b34801561069857600080fd5b506101e5611927565b3480156106ad57600080fd5b506101a7611936565b3480156106c257600080fd5b506101a7600160a060020a036004351661199b565b3480156106e357600080fd5b506102096119f8565b3480156106f857600080fd5b506101be611ac9565b34801561070d57600080fd5b506101be611aed565b34801561072257600080fd5b506101a7600160a060020a0360043516611af3565b34801561074357600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610209948235946024803515159536959460649492019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750611cb79650505050505050565b3480156107e357600080fd5b506101a760043561204e565b3480156107fb57600080fd5b506101a7600160a060020a03600435166120b7565b34801561081c57600080fd5b506101be600435612129565b34801561083457600080fd5b506101be61219a565b7f5075626c6973686564000000000000000000000000000000000000000000000081565b600154600160a060020a031681565b600080600061087d6119f8565b151561088857600080fd5b600454156108c757600454600e80546001810182556000919091527fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fd01555b6108d2600f546121ac565b6000600f556108e16001611240565b50600091506000600b5413156109825750600954604080517f49102e610000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169182916349102e619160048083019260209291908290030181600087803b15801561095357600080fd5b505af1158015610967573d6000803e3d6000fd5b505050506040513d602081101561097d57600080fd5b505191505b600b54600454604080518515158152602081019390935282810191909152517f901e6f3cdc4c61620d5d424116934b9af6e31ba79cdeaa349336d93ecfe846d49181900360600190a150919050565b600854600090600160a060020a0316158015906109f85750600854600160a060020a031633145b1515610a0357600080fd5b610a0f8585858561220e565b95945050505050565b600d602052600090815260408120805460018201546002830154600390930154600160a060020a03909216939092900b9084565b600854600160a060020a031681565b600b5481565b600754600160a060020a031681565b600654600160a060020a031681565b60085460009081908190600160a060020a031615610b2f5750600854604080517f45080442000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a0390921691829163450804429160248083019260209291908290030181600087803b158015610b0057600080fd5b505af1158015610b14573d6000803e3d6000fd5b505050506040513d6020811015610b2a57600080fd5b505191505b8115610b3d57819250610bca565b831515610b6c577f5075626c697368656400000000000000000000000000000000000000000000009250610bca565b6000841215610b9d577f44726166740000000000000000000000000000000000000000000000000000009250610bca565b6000841315610bca577f447261667420696e20726576696577000000000000000000000000000000000092505b5050919050565b60085460009081908190600019908290600160a060020a031615610d2e57506008546040517f0f82c16f00000000000000000000000000000000000000000000000000000000815260ff891660048201908152606060248301908152895160648401528951600160a060020a03909416938493630f82c16f938d938d938d9360448101916084909101906020808801910280838360005b83811015610c80578181015183820152602001610c68565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610cbf578181015183820152602001610ca7565b50505050905001955050505050506040805180830381600087803b158015610ce657600080fd5b505af1158015610cfa573d6000803e3d6000fd5b505050506040513d6040811015610d1057600080fd5b508051602090910151600a54909450909250831115610d2e57606491505b8160000b6000191415610d4957600a54600095509350610d50565b8183945094505b505050935093915050565b600254600090600160a060020a0316321480610d815750600254600160a060020a031633145b1515610d8c57600080fd5b600854600160a060020a031615610e355750600854604080517f9e99bbea0000000000000000000000000000000000000000000000000000000081529051600160a060020a03909216918291639e99bbea9160048083019260209291908290030181600087803b158015610dff57600080fd5b505af1158015610e13573d6000803e3d6000fd5b505050506040513d6020811015610e2957600080fd5b505115610e3557600080fd5b610e3d612340565b50565b7f447261667400000000000000000000000000000000000000000000000000000081565b600254600090600160a060020a031632148015610e9857506000821280610e985750600082138015610e9857506000600b54125b15610ea357600b8290555b600954600160a060020a0316331415610ebc57600b8290555b600b5460408051918252517fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a29181900360200190a15050600b5490565b60055481565b6000838152600d602052604081208054829081908190600160a060020a031615801590610f4957508354600160a060020a0316331480610f495750600254600160a060020a031633145b1515610f5457600080fd5b6008548715159350600160a060020a03161561100e57600854604080517f17685953000000000000000000000000000000000000000000000000000000008152600481018b9052602481018a90529051600160a060020a03909216935083916317685953916044808201926020929091908290030181600087803b158015610fdb57600080fd5b505af1158015610fef573d6000803e3d6000fd5b505050506040513d602081101561100557600080fd5b50518015935090505b8354600160a060020a031633141561104b57821561103b576002848101805460ff1916909117905561104b565b60028401805460ff191660fe1790555b836001015484600301541015611129576002840154600090810b810b136110cc57835460408051808201909152600681527f726566756e6400000000000000000000000000000000000000000000000000006020820152600386015460018701546110c6938c93600160a060020a039091169290910361220e565b50611129565b60025460408051808201909152600e81527f72656c6561736520657363726f77000000000000000000000000000000000000602082015260038601546001870154611127938c93600160a060020a039091169290910361220e565b505b6000888152600d60209081526040808320805473ffffffffffffffffffffffffffffffffffffffff191681556001810184905560028101805460ff191690556003019290925581518a8152908101899052808201889052841515606082015290517f7f1f4b28434ce7beab4983e64a8b5bb96e195a67029fdaff925028aec57fbc6b9181900360800190a150909695505050505050565b600a5481565b600154600160a060020a031632146111dd57600080fd5b600160a060020a03811615156111f257600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600e80548290811061122f57fe5b600091825260209091200154905081565b600080600061124d6119f8565b151561125857600080fd5b600854600160a060020a031615156112dc57600254600160a060020a031632148061128d5750600254600160a060020a031633145b80156112a457508360001914806112a45750836001145b156112b1578391506112d7565b600954600160a060020a0316331480156112ce57506000600b5412155b156112d7578391505b611374565b50600854604080517f3513a805000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a03909216918291633513a8059160248083019260209291908290030181600087803b15801561134557600080fd5b505af1158015611359573d6000803e3d6000fd5b505050506040513d602081101561136f57600080fd5b505191505b600b8290556040805183815290517fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a29181900360200190a15050600b5492915050565b60008060006113c7868686610bd1565b92509050600081900b156113da57600080fd5b6040805160ff881681526020810184905281517fa58326ee5bb617cb8b4f0d0f5f557c469d2d05d7a738f777037deda9c724b370929181900390910190a150949350505050565b600254600160a060020a031681565b600c5481565b600080600061144361237c565b600c80546001019055600954604080517f95a078e80000000000000000000000000000000000000000000000000000000081523260048201529051600160a060020a0390921694506000918291829187916395a078e891602480830192602092919082900301818787803b1580156114ba57600080fd5b505af11580156114ce573d6000803e3d6000fd5b505050506040513d60208110156114e457600080fd5b505115156114f157600080fd5b600254600160a060020a0316321461151d5761150e8c8a8a6113b7565b94503485111561151d57600080fd5b60408051608081018252338152346020808301918252600083850181815260608501828152600c548352600d90935294812084518154600160a060020a0391821673ffffffffffffffffffffffffffffffffffffffff199091161782559351600182015594516002860180549190920b60ff1660ff199091161790555160039093019290925560085490955016156116e557600860009054906101000a9004600160a060020a0316925082600160a060020a031663123e0e80600c548e8c8c6040518563ffffffff1660e060020a028152600401808581526020018460ff1660ff1681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561164457818101518382015260200161162c565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561168357818101518382015260200161166b565b505050509050019650505050505050602060405180830381600087803b1580156116ac57600080fd5b505af11580156116c0573d6000803e3d6000fd5b505050506040513d60208110156116d657600080fd5b5051915081156116e557600080fd5b7f089a6f1788a3c353423e1be4ba12533bdde7d908bb41abeee185af0acb3df562600c548d6004548e8e604051808681526020018560ff1660ff16815260200184600019166000191681526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611770578181015183820152602001611758565b50505050905090810190601f16801561179d5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156117d05781810151838201526020016117b8565b50505050905090810190601f1680156117fd5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15060005b885181101561188d57888181518110151561182a57fe5b6020908102909101015115611885577f515e0a48b385fce2a8e4d9f169a97c4f6ea669a752358f5e6ab37cc3c2e84c38898281518110151561186857fe5b602090810290910181015160408051918252519081900390910190a15b600101611813565b5060005b87518110156119155788818151811015156118a857fe5b602090810290910101511561190d577fb6e3239e521a6c66920ae634f8e921a37e6991d520ac44d52f8516397f41b68488828151811015156118e657fe5b602090810290910181015160408051600160a060020a039092168252519081900390910190a15b600101611891565b5060019b9a5050505050505050505050565b600954600160a060020a031681565b600254600160a060020a03163214806119595750600254600160a060020a031633145b151561196457600080fd5b60045460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600254600160a060020a03163214806119be5750600254600160a060020a031633145b15156119c957600080fd5b6007805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6002546000908190600160a060020a0316331480611a205750600954600160a060020a031633145b15611a2e5760019150611ac5565b50600954604080517f26683e140000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169182916326683e149160248083019260209291908290030181600087803b158015611a9657600080fd5b505af1158015611aaa573d6000803e3d6000fd5b505050506040513d6020811015611ac057600080fd5b505191505b5090565b7f447261667420696e20726576696577000000000000000000000000000000000081565b60045481565b60025460009081908190600160a060020a0316321480611b1d5750600254600160a060020a031633145b1515611b2857600080fd5b600854600160a060020a031615611bc857600860009054906101000a9004600160a060020a0316925082600160a060020a0316639e99bbea6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611b8f57600080fd5b505af1158015611ba3573d6000803e3d6000fd5b505050506040513d6020811015611bb957600080fd5b505191508115611bc857600080fd5b6008805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03861690811790915515611c725783925082600160a060020a0316637b1cdb3e6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611c3957600080fd5b505af1158015611c4d573d6000803e3d6000fd5b505050506040513d6020811015611c6357600080fd5b505190508015611c7257600080fd5b60085460408051600160a060020a039092168252517fa6f2e38f0cfebf27212317fced3ac40bc62e00bd33f38d69603710740c69acb79181900360200190a150505050565b600254600090819081908190600160a060020a0316331480611ce35750600754600160a060020a031633145b1515611cee57600080fd5b6000888152600d602052604090208054909350600160a060020a03161515611d1557600080fd5b600854879250600160a060020a031615611dce5750600854604080517fe870ed91000000000000000000000000000000000000000000000000000000008152600481018a905288151560248201529051600160a060020a0390921691829163e870ed919160448083019260209291908290030181600087803b158015611d9a57600080fd5b505af1158015611dae573d6000803e3d6000fd5b505050506040513d6020811015611dc457600080fd5b5051159150611ea2565b826001015483600301541015611ea257861515611e4557825460408051808201909152600f81527f616363657373206465636c696e65640000000000000000000000000000000000602082015260038501546001860154611e3f938c93600160a060020a039091169290910361220e565b50611ea2565b60025460408051808201909152600d81527f6f776e6572207061796d656e7400000000000000000000000000000000000000602082015260038501546001860154611ea0938c93600160a060020a039091169290910361220e565b505b60018215151415611fdd5760028301805460ff19166001908117909155604080518a8152602080820184905260809282018381528a519383019390935289517f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718948d9490938c938c93919291606084019160a08501919087019080838360005b83811015611f3a578181015183820152602001611f22565b50505050905090810190601f168015611f675780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611f9a578181015183820152602001611f82565b50505050905090810190601f168015611fc75780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a1612043565b60028301805460ff191660ff179055604080518981526000602082018190526080828401819052820181905260c06060830181905282015290517f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718918190036101000190a15b509695505050505050565b600254600160a060020a03163214806120715750600254600160a060020a031633145b151561207c57600080fd5b600f8190556040805182815290517fa288d3a2aba7d5dec44e93ff46d2f1129e251695be2046de105f9a9c6feefcaa9181900360200190a150565b600254600160a060020a03163214806120da5750600254600160a060020a031633145b15156120e557600080fd5b600160a060020a03811615156120fa57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600090600160a060020a031632148061214f5750600254600160a060020a031633145b151561215a57600080fd5b600a8290556040805183815290517f4114f8ef80b6de2161db580cbefa14e1892d15d3ebe2062c9914e4a5773114a39181900360200190a15050600a5490565b60006121a7600b54610a7f565b905090565b600254600160a060020a03163314806121c857506121c86119f8565b15156121d357600080fd5b60048190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b6000848152600d602052604081206001810154600382015484011161233757604051600160a060020a0386169084156108fc029085906000818181858888f19350505050158015612263573d6000803e3d6000fd5b508281600301540181600301819055507fad58d18ea7292f887da6f15bb4f0badddaa33d169713d09cf49710acc7c3a5b986858786604051808581526020018060200184600160a060020a0316600160a060020a03168152602001838152602001828103825285818151815260200191508051906020019080838360005b838110156122f95781810151838201526020016122e1565b50505050905090810190601f1680156123265780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a15b50949350505050565b600254600160a060020a03163214806123635750600254600160a060020a031633145b151561236e57600080fd5b600254600160a060020a0316ff5b604080516080810182526000808252602082018190529181018290526060810191909152905600a165627a7a723058202fb499b0671f272703d23a44987b8c3eea18538337d5e28f0d8dc8fca16109ae0029a165627a7a723058207dbf31d67589062030b4a2b7fe318b722e198afe2c9baf9ed2c099a1cc400bd5002960806040527f4f776e61626c6532303139303331353134313530304d4c0000000000000000006000557f42734163636573734374726c47727032303139303331353137323930304d4c0060035534801561005857600080fd5b506040516020806107a983398101604090815290516001805432600160a060020a0319918216811780845560028054841690921790915560048054909216600160a060020a03948516179091558216600090815260066020908152848220805460ff199081168517909155835490941682526005905292832080549092161790556106c09081906100e990396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100dc57806308ae4b0c1461010d5780630ae5e7391461014257806341c0e1b51461016357806342e7ba7b1461017857806354fd4d50146101995780636d2e4b1b146101c057806375861a95146101e157806385e68531146102025780638da5cb5b1461022357806395a078e814610238578063af570c0414610259578063cdb849b71461026e578063f2fde38b1461028f578063fdff9b4d146102b0575b005b3480156100e857600080fd5b506100f16102d1565b60408051600160a060020a039092168252519081900360200190f35b34801561011957600080fd5b5061012e600160a060020a03600435166102e0565b604080519115158252519081900360200190f35b34801561014e57600080fd5b506100da600160a060020a03600435166102f5565b34801561016f57600080fd5b506100da610371565b34801561018457600080fd5b5061012e600160a060020a03600435166103ad565b3480156101a557600080fd5b506101ae6103d0565b60408051918252519081900360200190f35b3480156101cc57600080fd5b506100da600160a060020a03600435166103d6565b3480156101ed57600080fd5b506100da600160a060020a0360043516610431565b34801561020e57600080fd5b506100da600160a060020a03600435166104ba565b34801561022f57600080fd5b506100f1610548565b34801561024457600080fd5b5061012e600160a060020a0360043516610557565b34801561026557600080fd5b506100f161057a565b34801561027a57600080fd5b506100da600160a060020a0360043516610589565b34801561029b57600080fd5b506100da600160a060020a036004351661060d565b3480156102bc57600080fd5b5061012e600160a060020a036004351661067f565b600154600160a060020a031681565b60056020526000908152604090205460ff1681565b3360009081526006602052604090205460ff16151560011461031657600080fd5b600160a060020a038116600081815260056020908152604091829020805460ff19166001179055815192835290517fb251eb052afc73ffd02ffe85ad79990a8b3fed60d76dbc2fa2fdd7123dffd9149281900390910190a150565b600254600160a060020a03163214806103945750600254600160a060020a031633145b151561039f57600080fd5b600254600160a060020a0316ff5b600160a060020a031660009081526006602052604090205460ff16151560011490565b60035481565b600154600160a060020a031632146103ed57600080fd5b600160a060020a038116151561040257600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a03163214806104545750600254600160a060020a031633145b151561045f57600080fd5b600160a060020a038116600081815260066020908152604091829020805460ff19166001179055815192835290517f93bcaab179551bde429187645251f8e1fb8ac85801fcb1cf91eb2c9043d611179281900390910190a150565b3360009081526006602052604090205460ff161515600114806104e5575033600160a060020a038216145b15156104f057600080fd5b600160a060020a038116600081815260056020908152604091829020805460ff19169055815192835290517f745cd29407db644ed93e3ceb61cbcab96d1dfb496989ac5d5bf514fc5a9fab9c9281900390910190a150565b600254600160a060020a031681565b600160a060020a031660009081526005602052604090205460ff16151560011490565b600454600160a060020a031681565b600254600160a060020a03163314806105aa575033600160a060020a038216145b15156105b557600080fd5b600160a060020a038116600081815260066020908152604091829020805460ff19169055815192835290517f2d6aa1a9629d125e23a0cf692cda7cd6795dff1652eedd4673b38ec31e387b959281900390910190a150565b600254600160a060020a03163214806106305750600254600160a060020a031633145b151561063b57600080fd5b600160a060020a038116151561065057600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60066020526000908152604090205460ff16815600a165627a7a72305820a3d364eca755c88b24e28f3029b5ff9767cb2358a596e6de29fab07635e9de520029608060408190527f41636365737369626c6532303139303232323133353930304d4c0000000000006000557f4f776e61626c6532303139303331353134313530304d4c0000000000000000006001557f4564697461626c6532303139303331353134313830304d4c00000000000000006004557f42617365436f6e74656e745479706532303139303331383130313230304d4c0060065560208061057883398101604052516002805432600160a060020a0319918216811790925560038054821690921790915560078054909116600160a060020a0390921691909117905561048b806100ed6000396000f3006080604052600436106100b95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100bb57806341c0e1b5146100ec57806354fd4d50146101015780636d2e4b1b146101285780638da5cb5b14610149578063af570c041461015e578063c287e0ed14610173578063cbcd446114610188578063e02dd9c2146101b1578063f14fcbc8146101c6578063f1551887146101de578063f2fde38b146101f3575b005b3480156100c757600080fd5b506100d0610214565b60408051600160a060020a039092168252519081900360200190f35b3480156100f857600080fd5b506100b9610223565b34801561010d57600080fd5b5061011661025f565b60408051918252519081900360200190f35b34801561013457600080fd5b506100b9600160a060020a0360043516610265565b34801561015557600080fd5b506100d06102c0565b34801561016a57600080fd5b506100d06102cf565b34801561017f57600080fd5b506100b96102de565b34801561019457600080fd5b5061019d610343565b604080519115158252519081900360200190f35b3480156101bd57600080fd5b50610116610354565b3480156101d257600080fd5b506100b960043561035a565b3480156101ea57600080fd5b5061019d6103bc565b3480156101ff57600080fd5b506100b9600160a060020a03600435166103ed565b600254600160a060020a031681565b600354600160a060020a03163214806102465750600354600160a060020a031633145b151561025157600080fd5b600354600160a060020a0316ff5b60065481565b600254600160a060020a0316321461027c57600080fd5b600160a060020a038116151561029157600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600354600160a060020a031681565b600754600160a060020a031681565b600354600160a060020a03163214806103015750600354600160a060020a031633145b151561030c57600080fd5b60055460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600354600160a060020a0316331490565b60055481565b600354600160a060020a03163314806103765750610376610343565b151561038157600080fd5b60058190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b6040516000907fed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88908290a150600190565b600354600160a060020a03163214806104105750600354600160a060020a031633145b151561041b57600080fd5b600160a060020a038116151561043057600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582041b0d335b9b4d89aa7e09dd15c22e519ef4e22dba7d5570bb82b50c0bb682de20029a165627a7a72305820f11fb3f913afcc66ed61e5d2ea0bfbf09602a396641a1e76d1cac8fa49d2bd520029`

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

// ActiveNodeAddresses is a free data retrieval call binding the contract method 0x52f82dd8.
//
// Solidity: function activeNodeAddresses(uint256 ) constant returns(address)
func (_BaseContentSpace *BaseContentSpaceCaller) ActiveNodeAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContentSpace.contract.Call(opts, out, "activeNodeAddresses", arg0)
	return *ret0, err
}

// ActiveNodeAddresses is a free data retrieval call binding the contract method 0x52f82dd8.
//
// Solidity: function activeNodeAddresses(uint256 ) constant returns(address)
func (_BaseContentSpace *BaseContentSpaceSession) ActiveNodeAddresses(arg0 *big.Int) (common.Address, error) {
	return _BaseContentSpace.Contract.ActiveNodeAddresses(&_BaseContentSpace.CallOpts, arg0)
}

// ActiveNodeAddresses is a free data retrieval call binding the contract method 0x52f82dd8.
//
// Solidity: function activeNodeAddresses(uint256 ) constant returns(address)
func (_BaseContentSpace *BaseContentSpaceCallerSession) ActiveNodeAddresses(arg0 *big.Int) (common.Address, error) {
	return _BaseContentSpace.Contract.ActiveNodeAddresses(&_BaseContentSpace.CallOpts, arg0)
}

// ActiveNodeLocators is a free data retrieval call binding the contract method 0x5272ae17.
//
// Solidity: function activeNodeLocators(uint256 ) constant returns(bytes)
func (_BaseContentSpace *BaseContentSpaceCaller) ActiveNodeLocators(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _BaseContentSpace.contract.Call(opts, out, "activeNodeLocators", arg0)
	return *ret0, err
}

// ActiveNodeLocators is a free data retrieval call binding the contract method 0x5272ae17.
//
// Solidity: function activeNodeLocators(uint256 ) constant returns(bytes)
func (_BaseContentSpace *BaseContentSpaceSession) ActiveNodeLocators(arg0 *big.Int) ([]byte, error) {
	return _BaseContentSpace.Contract.ActiveNodeLocators(&_BaseContentSpace.CallOpts, arg0)
}

// ActiveNodeLocators is a free data retrieval call binding the contract method 0x5272ae17.
//
// Solidity: function activeNodeLocators(uint256 ) constant returns(bytes)
func (_BaseContentSpace *BaseContentSpaceCallerSession) ActiveNodeLocators(arg0 *big.Int) ([]byte, error) {
	return _BaseContentSpace.Contract.ActiveNodeLocators(&_BaseContentSpace.CallOpts, arg0)
}

// CanNodePublish is a free data retrieval call binding the contract method 0x26683e14.
//
// Solidity: function canNodePublish(address candidate) constant returns(bool)
func (_BaseContentSpace *BaseContentSpaceCaller) CanNodePublish(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseContentSpace.contract.Call(opts, out, "canNodePublish", candidate)
	return *ret0, err
}

// CanNodePublish is a free data retrieval call binding the contract method 0x26683e14.
//
// Solidity: function canNodePublish(address candidate) constant returns(bool)
func (_BaseContentSpace *BaseContentSpaceSession) CanNodePublish(candidate common.Address) (bool, error) {
	return _BaseContentSpace.Contract.CanNodePublish(&_BaseContentSpace.CallOpts, candidate)
}

// CanNodePublish is a free data retrieval call binding the contract method 0x26683e14.
//
// Solidity: function canNodePublish(address candidate) constant returns(bool)
func (_BaseContentSpace *BaseContentSpaceCallerSession) CanNodePublish(candidate common.Address) (bool, error) {
	return _BaseContentSpace.Contract.CanNodePublish(&_BaseContentSpace.CallOpts, candidate)
}

// CanPublish is a free data retrieval call binding the contract method 0xcbcd4461.
//
// Solidity: function canPublish() constant returns(bool)
func (_BaseContentSpace *BaseContentSpaceCaller) CanPublish(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseContentSpace.contract.Call(opts, out, "canPublish")
	return *ret0, err
}

// CanPublish is a free data retrieval call binding the contract method 0xcbcd4461.
//
// Solidity: function canPublish() constant returns(bool)
func (_BaseContentSpace *BaseContentSpaceSession) CanPublish() (bool, error) {
	return _BaseContentSpace.Contract.CanPublish(&_BaseContentSpace.CallOpts)
}

// CanPublish is a free data retrieval call binding the contract method 0xcbcd4461.
//
// Solidity: function canPublish() constant returns(bool)
func (_BaseContentSpace *BaseContentSpaceCallerSession) CanPublish() (bool, error) {
	return _BaseContentSpace.Contract.CanPublish(&_BaseContentSpace.CallOpts)
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

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() constant returns(address)
func (_BaseContentSpace *BaseContentSpaceCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContentSpace.contract.Call(opts, out, "factory")
	return *ret0, err
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() constant returns(address)
func (_BaseContentSpace *BaseContentSpaceSession) Factory() (common.Address, error) {
	return _BaseContentSpace.Contract.Factory(&_BaseContentSpace.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() constant returns(address)
func (_BaseContentSpace *BaseContentSpaceCallerSession) Factory() (common.Address, error) {
	return _BaseContentSpace.Contract.Factory(&_BaseContentSpace.CallOpts)
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

// NumActiveNodes is a free data retrieval call binding the contract method 0x43f59ec7.
//
// Solidity: function numActiveNodes() constant returns(uint256)
func (_BaseContentSpace *BaseContentSpaceCaller) NumActiveNodes(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseContentSpace.contract.Call(opts, out, "numActiveNodes")
	return *ret0, err
}

// NumActiveNodes is a free data retrieval call binding the contract method 0x43f59ec7.
//
// Solidity: function numActiveNodes() constant returns(uint256)
func (_BaseContentSpace *BaseContentSpaceSession) NumActiveNodes() (*big.Int, error) {
	return _BaseContentSpace.Contract.NumActiveNodes(&_BaseContentSpace.CallOpts)
}

// NumActiveNodes is a free data retrieval call binding the contract method 0x43f59ec7.
//
// Solidity: function numActiveNodes() constant returns(uint256)
func (_BaseContentSpace *BaseContentSpaceCallerSession) NumActiveNodes() (*big.Int, error) {
	return _BaseContentSpace.Contract.NumActiveNodes(&_BaseContentSpace.CallOpts)
}

// NumPendingNodes is a free data retrieval call binding the contract method 0xf41a1587.
//
// Solidity: function numPendingNodes() constant returns(uint256)
func (_BaseContentSpace *BaseContentSpaceCaller) NumPendingNodes(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseContentSpace.contract.Call(opts, out, "numPendingNodes")
	return *ret0, err
}

// NumPendingNodes is a free data retrieval call binding the contract method 0xf41a1587.
//
// Solidity: function numPendingNodes() constant returns(uint256)
func (_BaseContentSpace *BaseContentSpaceSession) NumPendingNodes() (*big.Int, error) {
	return _BaseContentSpace.Contract.NumPendingNodes(&_BaseContentSpace.CallOpts)
}

// NumPendingNodes is a free data retrieval call binding the contract method 0xf41a1587.
//
// Solidity: function numPendingNodes() constant returns(uint256)
func (_BaseContentSpace *BaseContentSpaceCallerSession) NumPendingNodes() (*big.Int, error) {
	return _BaseContentSpace.Contract.NumPendingNodes(&_BaseContentSpace.CallOpts)
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

// PendingNodeAddresses is a free data retrieval call binding the contract method 0x6be9514c.
//
// Solidity: function pendingNodeAddresses(uint256 ) constant returns(address)
func (_BaseContentSpace *BaseContentSpaceCaller) PendingNodeAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseContentSpace.contract.Call(opts, out, "pendingNodeAddresses", arg0)
	return *ret0, err
}

// PendingNodeAddresses is a free data retrieval call binding the contract method 0x6be9514c.
//
// Solidity: function pendingNodeAddresses(uint256 ) constant returns(address)
func (_BaseContentSpace *BaseContentSpaceSession) PendingNodeAddresses(arg0 *big.Int) (common.Address, error) {
	return _BaseContentSpace.Contract.PendingNodeAddresses(&_BaseContentSpace.CallOpts, arg0)
}

// PendingNodeAddresses is a free data retrieval call binding the contract method 0x6be9514c.
//
// Solidity: function pendingNodeAddresses(uint256 ) constant returns(address)
func (_BaseContentSpace *BaseContentSpaceCallerSession) PendingNodeAddresses(arg0 *big.Int) (common.Address, error) {
	return _BaseContentSpace.Contract.PendingNodeAddresses(&_BaseContentSpace.CallOpts, arg0)
}

// PendingNodeLocators is a free data retrieval call binding the contract method 0x69e30ff8.
//
// Solidity: function pendingNodeLocators(uint256 ) constant returns(bytes)
func (_BaseContentSpace *BaseContentSpaceCaller) PendingNodeLocators(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _BaseContentSpace.contract.Call(opts, out, "pendingNodeLocators", arg0)
	return *ret0, err
}

// PendingNodeLocators is a free data retrieval call binding the contract method 0x69e30ff8.
//
// Solidity: function pendingNodeLocators(uint256 ) constant returns(bytes)
func (_BaseContentSpace *BaseContentSpaceSession) PendingNodeLocators(arg0 *big.Int) ([]byte, error) {
	return _BaseContentSpace.Contract.PendingNodeLocators(&_BaseContentSpace.CallOpts, arg0)
}

// PendingNodeLocators is a free data retrieval call binding the contract method 0x69e30ff8.
//
// Solidity: function pendingNodeLocators(uint256 ) constant returns(bytes)
func (_BaseContentSpace *BaseContentSpaceCallerSession) PendingNodeLocators(arg0 *big.Int) ([]byte, error) {
	return _BaseContentSpace.Contract.PendingNodeLocators(&_BaseContentSpace.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseContentSpace *BaseContentSpaceCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseContentSpace.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseContentSpace *BaseContentSpaceSession) Version() ([32]byte, error) {
	return _BaseContentSpace.Contract.Version(&_BaseContentSpace.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseContentSpace *BaseContentSpaceCallerSession) Version() ([32]byte, error) {
	return _BaseContentSpace.Contract.Version(&_BaseContentSpace.CallOpts)
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

// AddNode is a paid mutator transaction binding the contract method 0x64f0f050.
//
// Solidity: function addNode(address _nodeAddr, bytes _locator) returns()
func (_BaseContentSpace *BaseContentSpaceTransactor) AddNode(opts *bind.TransactOpts, _nodeAddr common.Address, _locator []byte) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "addNode", _nodeAddr, _locator)
}

// AddNode is a paid mutator transaction binding the contract method 0x64f0f050.
//
// Solidity: function addNode(address _nodeAddr, bytes _locator) returns()
func (_BaseContentSpace *BaseContentSpaceSession) AddNode(_nodeAddr common.Address, _locator []byte) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.AddNode(&_BaseContentSpace.TransactOpts, _nodeAddr, _locator)
}

// AddNode is a paid mutator transaction binding the contract method 0x64f0f050.
//
// Solidity: function addNode(address _nodeAddr, bytes _locator) returns()
func (_BaseContentSpace *BaseContentSpaceTransactorSession) AddNode(_nodeAddr common.Address, _locator []byte) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.AddNode(&_BaseContentSpace.TransactOpts, _nodeAddr, _locator)
}

// ApproveNode is a paid mutator transaction binding the contract method 0xdd4c97a0.
//
// Solidity: function approveNode(address _nodeAddr) returns()
func (_BaseContentSpace *BaseContentSpaceTransactor) ApproveNode(opts *bind.TransactOpts, _nodeAddr common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "approveNode", _nodeAddr)
}

// ApproveNode is a paid mutator transaction binding the contract method 0xdd4c97a0.
//
// Solidity: function approveNode(address _nodeAddr) returns()
func (_BaseContentSpace *BaseContentSpaceSession) ApproveNode(_nodeAddr common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.ApproveNode(&_BaseContentSpace.TransactOpts, _nodeAddr)
}

// ApproveNode is a paid mutator transaction binding the contract method 0xdd4c97a0.
//
// Solidity: function approveNode(address _nodeAddr) returns()
func (_BaseContentSpace *BaseContentSpaceTransactorSession) ApproveNode(_nodeAddr common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.ApproveNode(&_BaseContentSpace.TransactOpts, _nodeAddr)
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

// RemoveNode is a paid mutator transaction binding the contract method 0xb2b99ec9.
//
// Solidity: function removeNode(address _nodeAddr) returns()
func (_BaseContentSpace *BaseContentSpaceTransactor) RemoveNode(opts *bind.TransactOpts, _nodeAddr common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "removeNode", _nodeAddr)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xb2b99ec9.
//
// Solidity: function removeNode(address _nodeAddr) returns()
func (_BaseContentSpace *BaseContentSpaceSession) RemoveNode(_nodeAddr common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.RemoveNode(&_BaseContentSpace.TransactOpts, _nodeAddr)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xb2b99ec9.
//
// Solidity: function removeNode(address _nodeAddr) returns()
func (_BaseContentSpace *BaseContentSpaceTransactorSession) RemoveNode(_nodeAddr common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.RemoveNode(&_BaseContentSpace.TransactOpts, _nodeAddr)
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

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(address new_factory) returns()
func (_BaseContentSpace *BaseContentSpaceTransactor) SetFactory(opts *bind.TransactOpts, new_factory common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "setFactory", new_factory)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(address new_factory) returns()
func (_BaseContentSpace *BaseContentSpaceSession) SetFactory(new_factory common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.SetFactory(&_BaseContentSpace.TransactOpts, new_factory)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(address new_factory) returns()
func (_BaseContentSpace *BaseContentSpaceTransactorSession) SetFactory(new_factory common.Address) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.SetFactory(&_BaseContentSpace.TransactOpts, new_factory)
}

// SubmitNode is a paid mutator transaction binding the contract method 0x160eee74.
//
// Solidity: function submitNode(bytes _locator) returns()
func (_BaseContentSpace *BaseContentSpaceTransactor) SubmitNode(opts *bind.TransactOpts, _locator []byte) (*types.Transaction, error) {
	return _BaseContentSpace.contract.Transact(opts, "submitNode", _locator)
}

// SubmitNode is a paid mutator transaction binding the contract method 0x160eee74.
//
// Solidity: function submitNode(bytes _locator) returns()
func (_BaseContentSpace *BaseContentSpaceSession) SubmitNode(_locator []byte) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.SubmitNode(&_BaseContentSpace.TransactOpts, _locator)
}

// SubmitNode is a paid mutator transaction binding the contract method 0x160eee74.
//
// Solidity: function submitNode(bytes _locator) returns()
func (_BaseContentSpace *BaseContentSpaceTransactorSession) SubmitNode(_locator []byte) (*types.Transaction, error) {
	return _BaseContentSpace.Contract.SubmitNode(&_BaseContentSpace.TransactOpts, _locator)
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

// BaseContentSpaceNodeApprovedIterator is returned from FilterNodeApproved and is used to iterate over the raw logs and unpacked data for NodeApproved events raised by the BaseContentSpace contract.
type BaseContentSpaceNodeApprovedIterator struct {
	Event *BaseContentSpaceNodeApproved // Event containing the contract specifics and raw log

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
func (it *BaseContentSpaceNodeApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentSpaceNodeApproved)
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
		it.Event = new(BaseContentSpaceNodeApproved)
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
func (it *BaseContentSpaceNodeApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentSpaceNodeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentSpaceNodeApproved represents a NodeApproved event raised by the BaseContentSpace contract.
type BaseContentSpaceNodeApproved struct {
	Addr    common.Address
	Locator []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeApproved is a free log retrieval operation binding the contract event 0xd644c8164f225d3b7fdbcc404f279bb1e823ef0d93f88dd4b24e85d0e7bc6a54.
//
// Solidity: event NodeApproved(address addr, bytes locator)
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterNodeApproved(opts *bind.FilterOpts) (*BaseContentSpaceNodeApprovedIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "NodeApproved")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceNodeApprovedIterator{contract: _BaseContentSpace.contract, event: "NodeApproved", logs: logs, sub: sub}, nil
}

// WatchNodeApproved is a free log subscription operation binding the contract event 0xd644c8164f225d3b7fdbcc404f279bb1e823ef0d93f88dd4b24e85d0e7bc6a54.
//
// Solidity: event NodeApproved(address addr, bytes locator)
func (_BaseContentSpace *BaseContentSpaceFilterer) WatchNodeApproved(opts *bind.WatchOpts, sink chan<- *BaseContentSpaceNodeApproved) (event.Subscription, error) {

	logs, sub, err := _BaseContentSpace.contract.WatchLogs(opts, "NodeApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentSpaceNodeApproved)
				if err := _BaseContentSpace.contract.UnpackLog(event, "NodeApproved", log); err != nil {
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

// BaseContentSpaceNodeSubmittedIterator is returned from FilterNodeSubmitted and is used to iterate over the raw logs and unpacked data for NodeSubmitted events raised by the BaseContentSpace contract.
type BaseContentSpaceNodeSubmittedIterator struct {
	Event *BaseContentSpaceNodeSubmitted // Event containing the contract specifics and raw log

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
func (it *BaseContentSpaceNodeSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentSpaceNodeSubmitted)
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
		it.Event = new(BaseContentSpaceNodeSubmitted)
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
func (it *BaseContentSpaceNodeSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentSpaceNodeSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentSpaceNodeSubmitted represents a NodeSubmitted event raised by the BaseContentSpace contract.
type BaseContentSpaceNodeSubmitted struct {
	Addr    common.Address
	Locator []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeSubmitted is a free log retrieval operation binding the contract event 0xae5645569f32b946f7a747113c64094a29a6b84c5ddf55816ef4381ce8a3a46d.
//
// Solidity: event NodeSubmitted(address addr, bytes locator)
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterNodeSubmitted(opts *bind.FilterOpts) (*BaseContentSpaceNodeSubmittedIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "NodeSubmitted")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceNodeSubmittedIterator{contract: _BaseContentSpace.contract, event: "NodeSubmitted", logs: logs, sub: sub}, nil
}

// WatchNodeSubmitted is a free log subscription operation binding the contract event 0xae5645569f32b946f7a747113c64094a29a6b84c5ddf55816ef4381ce8a3a46d.
//
// Solidity: event NodeSubmitted(address addr, bytes locator)
func (_BaseContentSpace *BaseContentSpaceFilterer) WatchNodeSubmitted(opts *bind.WatchOpts, sink chan<- *BaseContentSpaceNodeSubmitted) (event.Subscription, error) {

	logs, sub, err := _BaseContentSpace.contract.WatchLogs(opts, "NodeSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentSpaceNodeSubmitted)
				if err := _BaseContentSpace.contract.UnpackLog(event, "NodeSubmitted", log); err != nil {
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

// BaseContentSpaceSetFactoryIterator is returned from FilterSetFactory and is used to iterate over the raw logs and unpacked data for SetFactory events raised by the BaseContentSpace contract.
type BaseContentSpaceSetFactoryIterator struct {
	Event *BaseContentSpaceSetFactory // Event containing the contract specifics and raw log

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
func (it *BaseContentSpaceSetFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseContentSpaceSetFactory)
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
		it.Event = new(BaseContentSpaceSetFactory)
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
func (it *BaseContentSpaceSetFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseContentSpaceSetFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseContentSpaceSetFactory represents a SetFactory event raised by the BaseContentSpace contract.
type BaseContentSpaceSetFactory struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetFactory is a free log retrieval operation binding the contract event 0x1c893ef9379093af30f458b9e74d2aba13c499660b68dec5e29af7b199c188b9.
//
// Solidity: event SetFactory(address factory)
func (_BaseContentSpace *BaseContentSpaceFilterer) FilterSetFactory(opts *bind.FilterOpts) (*BaseContentSpaceSetFactoryIterator, error) {

	logs, sub, err := _BaseContentSpace.contract.FilterLogs(opts, "SetFactory")
	if err != nil {
		return nil, err
	}
	return &BaseContentSpaceSetFactoryIterator{contract: _BaseContentSpace.contract, event: "SetFactory", logs: logs, sub: sub}, nil
}

// WatchSetFactory is a free log subscription operation binding the contract event 0x1c893ef9379093af30f458b9e74d2aba13c499660b68dec5e29af7b199c188b9.
//
// Solidity: event SetFactory(address factory)
func (_BaseContentSpace *BaseContentSpaceFilterer) WatchSetFactory(opts *bind.WatchOpts, sink chan<- *BaseContentSpaceSetFactory) (event.Subscription, error) {

	logs, sub, err := _BaseContentSpace.contract.WatchLogs(opts, "SetFactory")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseContentSpaceSetFactory)
				if err := _BaseContentSpace.contract.UnpackLog(event, "SetFactory", log); err != nil {
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
const BaseContentTypeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canPublish\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"objectHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"object_hash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accessRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"content_space\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"Commit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"UpdateRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AccessRequest\",\"type\":\"event\"}]"

// BaseContentTypeBin is the compiled bytecode used for deploying new contracts.
const BaseContentTypeBin = `0x608060408190527f41636365737369626c6532303139303232323133353930304d4c0000000000006000557f4f776e61626c6532303139303331353134313530304d4c0000000000000000006001557f4564697461626c6532303139303331353134313830304d4c00000000000000006004557f42617365436f6e74656e745479706532303139303331383130313230304d4c0060065560208061057883398101604052516002805432600160a060020a0319918216811790925560038054821690921790915560078054909116600160a060020a0390921691909117905561048b806100ed6000396000f3006080604052600436106100b95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100bb57806341c0e1b5146100ec57806354fd4d50146101015780636d2e4b1b146101285780638da5cb5b14610149578063af570c041461015e578063c287e0ed14610173578063cbcd446114610188578063e02dd9c2146101b1578063f14fcbc8146101c6578063f1551887146101de578063f2fde38b146101f3575b005b3480156100c757600080fd5b506100d0610214565b60408051600160a060020a039092168252519081900360200190f35b3480156100f857600080fd5b506100b9610223565b34801561010d57600080fd5b5061011661025f565b60408051918252519081900360200190f35b34801561013457600080fd5b506100b9600160a060020a0360043516610265565b34801561015557600080fd5b506100d06102c0565b34801561016a57600080fd5b506100d06102cf565b34801561017f57600080fd5b506100b96102de565b34801561019457600080fd5b5061019d610343565b604080519115158252519081900360200190f35b3480156101bd57600080fd5b50610116610354565b3480156101d257600080fd5b506100b960043561035a565b3480156101ea57600080fd5b5061019d6103bc565b3480156101ff57600080fd5b506100b9600160a060020a03600435166103ed565b600254600160a060020a031681565b600354600160a060020a03163214806102465750600354600160a060020a031633145b151561025157600080fd5b600354600160a060020a0316ff5b60065481565b600254600160a060020a0316321461027c57600080fd5b600160a060020a038116151561029157600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600354600160a060020a031681565b600754600160a060020a031681565b600354600160a060020a03163214806103015750600354600160a060020a031633145b151561030c57600080fd5b60055460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600354600160a060020a0316331490565b60055481565b600354600160a060020a03163314806103765750610376610343565b151561038157600080fd5b60058190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b6040516000907fed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88908290a150600190565b600354600160a060020a03163214806104105750600354600160a060020a031633145b151561041b57600080fd5b600160a060020a038116151561043057600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582041b0d335b9b4d89aa7e09dd15c22e519ef4e22dba7d5570bb82b50c0bb682de20029`

// DeployBaseContentType deploys a new Ethereum contract, binding an instance of BaseContentType to it.
func DeployBaseContentType(auth *bind.TransactOpts, backend bind.ContractBackend, content_space common.Address) (common.Address, *types.Transaction, *BaseContentType, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseContentTypeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaseContentTypeBin), backend, content_space)
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

// CanPublish is a free data retrieval call binding the contract method 0xcbcd4461.
//
// Solidity: function canPublish() constant returns(bool)
func (_BaseContentType *BaseContentTypeCaller) CanPublish(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseContentType.contract.Call(opts, out, "canPublish")
	return *ret0, err
}

// CanPublish is a free data retrieval call binding the contract method 0xcbcd4461.
//
// Solidity: function canPublish() constant returns(bool)
func (_BaseContentType *BaseContentTypeSession) CanPublish() (bool, error) {
	return _BaseContentType.Contract.CanPublish(&_BaseContentType.CallOpts)
}

// CanPublish is a free data retrieval call binding the contract method 0xcbcd4461.
//
// Solidity: function canPublish() constant returns(bool)
func (_BaseContentType *BaseContentTypeCallerSession) CanPublish() (bool, error) {
	return _BaseContentType.Contract.CanPublish(&_BaseContentType.CallOpts)
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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseContentType *BaseContentTypeCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseContentType.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseContentType *BaseContentTypeSession) Version() ([32]byte, error) {
	return _BaseContentType.Contract.Version(&_BaseContentType.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseContentType *BaseContentTypeCallerSession) Version() ([32]byte, error) {
	return _BaseContentType.Contract.Version(&_BaseContentType.CallOpts)
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

// BaseFactoryABI is the input ABI used to generate the binding from.
const BaseFactoryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"address_KMS\",\"type\":\"address\"}],\"name\":\"createLibrary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createContentType\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// BaseFactoryBin is the compiled bytecode used for deploying new contracts.
const BaseFactoryBin = `0x60806040527f4f776e61626c6532303139303331353134313530304d4c0000000000000000006000557f42617365466163746f727932303139303331393139353030304d4c000000000060035560018054600160a060020a0319908116329081179092556002805490911690911790556157c68061007e6000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f811461009a57806340b89f06146100cb57806341c0e1b5146100ec57806354fd4d5014610101578063575185ed146101285780636d2e4b1b1461013d5780638da5cb5b1461015e578063b8cfaf0514610173578063f2fde38b14610188575b005b3480156100a657600080fd5b506100af6101a9565b60408051600160a060020a039092168252519081900360200190f35b3480156100d757600080fd5b506100af600160a060020a03600435166101b8565b3480156100f857600080fd5b506100986101fe565b34801561010d57600080fd5b5061011661023a565b60408051918252519081900360200190f35b34801561013457600080fd5b506100af610240565b34801561014957600080fd5b50610098600160a060020a036004351661027e565b34801561016a57600080fd5b506100af6102d9565b34801561017f57600080fd5b506100af6102e8565b34801561019457600080fd5b50610098600160a060020a03600435166102f3565b600154600160a060020a031681565b600081336101c4610365565b600160a060020a03928316815291166020820152604080519182900301906000f0801580156101f7573d6000803e3d6000fd5b5092915050565b600254600160a060020a03163214806102215750600254600160a060020a031633145b151561022c57600080fd5b600254600160a060020a0316ff5b60035481565b60003361024b610375565b600160a060020a03909116815260405190819003602001906000f080158015610278573d6000803e3d6000fd5b50905090565b600154600160a060020a0316321461029557600080fd5b600160a060020a03811615156102aa57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a031681565b60003361024b610385565b600254600160a060020a03163214806103165750600254600160a060020a031633145b151561032157600080fd5b600160a060020a038116151561033657600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6040516146e48061039683390190565b6040516107a980614a7a83390190565b60405161057880615223833901905600608060408181527f41636365737369626c6532303139303232323133353930304d4c00000000000060009081557f4f776e61626c6532303139303331353134313530304d4c0000000000000000006001557f4564697461626c6532303139303331353134313830304d4c00000000000000006004557f426173654c69627261727932303139303331383130313330304d4c0000000000600655600c819055600d819055600e819055600f819055601355806146e48339810160405280516020909101516002805432600160a060020a03199182168117909255600380548216909217909155600780548216600160a060020a0393841617905560128054909116919092161790556145cf806101156000396000f3006080604052600436106101ea5763ffffffff60e060020a60003504166302d05d3f81146101ec5780630eaec2c51461021d5780630f58a7861461025257806316308394146102795780631b969895146102a05780631cdbee5a146102c15780631d0f4351146102e25780631e35d8fa1461030357806321770a84146103245780632393553b1461033957806326683e141461035157806329d002191461037257806329dedde5146103935780632c11f392146103b457806332eaf21b146103cc578063386493e0146103e157806341c0e1b514610402578063470750bb1461041757806349102e611461042c57806354fd4d501461044157806363dab9d414610456578063679a9a3c1461046e5780636d2e4b1b1461048f57806387e86b2c146104b05780638cb13c2e1461051b5780638da5cb5b14610533578063952e464b1461054857806395a078e814610560578063991a3a7c14610581578063af570c0414610599578063c287e0ed146105ae578063c65bcbe2146105c3578063c9e8e72d146105d8578063cbcd4461146105f9578063dc3c29c01461060e578063e02dd9c21461062f578063e5538fd214610644578063e8de515f14610659578063f14fcbc81461067a578063f155188714610692578063f2fde38b146106a7578063fd089196146106c8575b005b3480156101f857600080fd5b506102016106e9565b60408051600160a060020a039092168252519081900360200190f35b34801561022957600080fd5b5061023e600160a060020a03600435166106f8565b604080519115158252519081900360200190f35b34801561025e57600080fd5b506101ea600160a060020a0360043581169060243516610803565b34801561028557600080fd5b5061028e610978565b60408051918252519081900360200190f35b3480156102ac57600080fd5b5061023e600160a060020a036004351661097e565b3480156102cd57600080fd5b50610201600160a060020a0360043516610b1f565b3480156102ee57600080fd5b506101ea600160a060020a0360043516610b3a565b34801561030f57600080fd5b50610201600160a060020a0360043516610c4d565b34801561033057600080fd5b5061028e610e0e565b34801561034557600080fd5b50610201600435610e14565b34801561035d57600080fd5b5061023e600160a060020a0360043516610e3c565b34801561037e57600080fd5b5061023e600160a060020a0360043516610edb565b34801561039f57600080fd5b5061023e600160a060020a0360043516610fe6565b3480156103c057600080fd5b50610201600435611041565b3480156103d857600080fd5b5061020161104f565b3480156103ed57600080fd5b5061023e600160a060020a036004351661105e565b34801561040e57600080fd5b506101ea6111f4565b34801561042357600080fd5b5061028e611230565b34801561043857600080fd5b5061023e611236565b34801561044d57600080fd5b5061028e61144d565b34801561046257600080fd5b50610201600435611453565b34801561047a57600080fd5b506101ea600160a060020a03600435166114a0565b34801561049b57600080fd5b506101ea600160a060020a03600435166115b3565b3480156104bc57600080fd5b50604080516020600460443581810135601f810184900484028501840190955284845261023e948235600160a060020a031694602480351515953695946064949201919081908401838280828437509497506116019650505050505050565b34801561052757600080fd5b50610201600435611921565b34801561053f57600080fd5b5061020161192f565b34801561055457600080fd5b5061020160043561193e565b34801561056c57600080fd5b5061023e600160a060020a036004351661194c565b34801561058d57600080fd5b50610201600435611a49565b3480156105a557600080fd5b50610201611a57565b3480156105ba57600080fd5b506101ea611a66565b3480156105cf57600080fd5b5061028e611acb565b3480156105e457600080fd5b506101ea600160a060020a0360043516611ad1565b34801561060557600080fd5b5061023e611b21565b34801561061a57600080fd5b506101ea600160a060020a0360043516611b26565b34801561063b57600080fd5b5061028e611c39565b34801561065057600080fd5b5061028e611c3f565b34801561066557600080fd5b5061023e600160a060020a0360043516611c45565b34801561068657600080fd5b506101ea600435611ddb565b34801561069e57600080fd5b5061023e611e3d565b3480156106b357600080fd5b506101ea600160a060020a0360043516611ea0565b3480156106d457600080fd5b5061023e600160a060020a0360043516611f05565b600254600160a060020a031681565b6000806000806000600c546000141561071457600194506107fa565b5060005b600c548110156107f557600880548290811061073057fe5b600091825260209091200154600160a060020a0316935083156107ed5783915081600160a060020a03166395a078e8876040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b1580156107ab57600080fd5b505af11580156107bf573d6000803e3d6000fd5b505050506040513d60208110156107d557600080fd5b50519250600183151514156107ed57600194506107fa565b600101610718565b600094505b50505050919050565b600354600160a060020a03163214806108265750600354600160a060020a031633145b151561083157600080fd5b600160a060020a038083166000908152601060205260409020541615801561085f575061085d82610fe6565b155b1561090b57600b54600f5410156108b55781600b600f5481548110151561088257fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550610901565b600b80546001810182556000919091527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db9018054600160a060020a031916600160a060020a0384161790555b600f805460010190555b600160a060020a038281166000818152601060209081526040918290208054600160a060020a0319169486169485179055815192835282019290925281517f280016f7418306a55542432120fd1a239ef9fcc1a92694d8d44ca76be0249ea7929181900390910190a15050565b60135481565b6003546000908190600160a060020a03163214806109a65750600354600160a060020a031633145b15156109b157600080fd5b5060005b600d54811015610b145782600160a060020a03166009828154811015156109d857fe5b600091825260209091200154600160a060020a03161415610b0c576009805482908110610a0157fe5b60009182526020909120018054600160a060020a0319169055600d54600019018114610abd5760096001600d5403815481101515610a3b57fe5b60009182526020909120015460098054600160a060020a039092169183908110610a6157fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a0316021790555060096001600d5403815481101515610aa357fe5b60009182526020909120018054600160a060020a03191690555b600d805460001901905560408051600160a060020a038516815290517fdf9d78c5635b72b709c85300a786eb7238acbe5bffe01c60c16464e45c6eb6eb9181900360200190a160019150610b19565b6001016109b5565b600091505b50919050565b601060205260009081526040902054600160a060020a031681565b600354600160a060020a0316321480610b5d5750600354600160a060020a031633145b1515610b6857600080fd5b600a54600e541015610bb95780600a600e54815481101515610b8657fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550610c05565b600a80546001810182556000919091527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8018054600160a060020a031916600160a060020a0383161790555b600e8054600101905560408051600160a060020a038316815290517f3a94857e4393737f73edb175a7d0c195c7f635d9ae995e12740616ec55c9d4119181900360200190a150565b600080610c59326106f8565b1515610c6457600080fd5b600f5415610c8057610c7583610fe6565b1515610c8057600080fd5b82610c896120b2565b600160a060020a03909116815260405190819003602001906000f080158015610cb6573d6000803e3d6000fd5b50601254604080517fc9e8e72d000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015290519293509083169163c9e8e72d9160248082019260009290919082900301818387803b158015610d2157600080fd5b505af1158015610d35573d6000803e3d6000fd5b50505050600160a060020a038381166000908152601060205260408082205481517fe5385303000000000000000000000000000000000000000000000000000000008152908416600482015290519284169263e53853039260248084019391929182900301818387803b158015610dab57600080fd5b505af1158015610dbf573d6000803e3d6000fd5b505060408051600160a060020a0380861682528716602082015281517f3981e74ab81857b375ec391a4f7c31ee89462cd927de6d8fbdb98f77da009c569450908190039091019150a192915050565b600d5481565b6008805482908110610e2257fe5b600091825260209091200154600160a060020a0316905081565b600754604080517f26683e14000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151600093929092169182916326683e1491602480830192602092919082900301818887803b158015610ea857600080fd5b505af1158015610ebc573d6000803e3d6000fd5b505050506040513d6020811015610ed257600080fd5b50519392505050565b6003546000908190819081908190600160a060020a0387811691161415610f0557600194506107fa565b5060005b600d548110156107f5576009805482908110610f2157fe5b600091825260209091200154600160a060020a031693508315610fde5783915081600160a060020a03166395a078e8876040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b158015610f9c57600080fd5b505af1158015610fb0573d6000803e3d6000fd5b505050506040513d6020811015610fc657600080fd5b5051925060018315151415610fde57600194506107fa565b600101610f09565b600080805b600f548110156110365783600160a060020a0316600b8281548110151561100e57fe5b600091825260209091200154600160a060020a0316141561102e57600191505b600101610feb565b8192505b5050919050565b600a805482908110610e2257fe5b601254600160a060020a031681565b6003546000908190600160a060020a03163214806110865750600354600160a060020a031633145b151561109157600080fd5b5060005b600c54811015610b145782600160a060020a03166008828154811015156110b857fe5b600091825260209091200154600160a060020a031614156111ec5760088054829081106110e157fe5b60009182526020909120018054600160a060020a0319169055600c5460001901811461119d5760086001600c540381548110151561111b57fe5b60009182526020909120015460088054600160a060020a03909216918390811061114157fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a0316021790555060086001600c540381548110151561118357fe5b60009182526020909120018054600160a060020a03191690555b600c805460001901905560408051600160a060020a038516815290517fbbd97daa1862eb12f77ed128a557406737cee07b131b1e2d7140dff2005e197c9181900360200190a160019150610b19565b600101611095565b600354600160a060020a03163214806112175750600354600160a060020a031633145b151561122257600080fd5b600354600160a060020a0316ff5b600c5481565b600d546000903390819015156113165780600160a060020a0316638280dd8f60006040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b15801561129057600080fd5b505af11580156112a4573d6000803e3d6000fd5b505050506040513d60208110156112ba57600080fd5b505060408051600160a060020a038416815260016020820152606081830181905260009082015290517f70234ce475fee4ab40e5e55cf533f67f12b47ef4c860e62dd7affa84ead4b4429181900360a00190a160019250611448565b600160a060020a0382166000908152601460205260409020541561133d5760009250611448565b601154601354101561138e5781601160135481548110151561135b57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055506113da565b601180546001810182556000919091527f31ecc21a745e3968a04e9570e4425bc18fa8019c68028196b546d1669c200c68018054600160a060020a031916600160a060020a0384161790555b60138054600160a060020a0384166000818152601460209081526040918290206001948501905584549093019093558251908152329181019190915281517f0588a34cf0de4e025d359c89ca4bacbcbf175440909952d91c814412d9da996a929181900390910190a1600192505b505090565b60065481565b60006013546000148061146857508160135411155b156114755750600061149b565b601180548390811061148357fe5b600091825260209091200154600160a060020a031690505b919050565b600354600160a060020a03163214806114c35750600354600160a060020a031633145b15156114ce57600080fd5b600854600c54101561151f57806008600c548154811015156114ec57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a0316021790555061156b565b600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3018054600160a060020a031916600160a060020a0383161790555b600c8054600101905560408051600160a060020a038316815290517f218673669018c25b89bfbf1b58d0075e37c8847ef16e707b92355b7833e97d619181900360200190a150565b600254600160a060020a031632146115ca57600080fd5b600160a060020a03811615156115df57600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b60008060008060008061161332610edb565b151560011461162157600080fd5b600160a060020a0389166000908152601460205260409020546011805460001990920196508a9550908690811061165457fe5b600091825260208083209091018054600160a060020a031916905560138054600019018155600160a060020a038c16835260149091526040822091909155548510156117495760116013548154811015156116ab57fe5b60009182526020909120015460118054600160a060020a0390921694508491879081106116d457fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550601160135481548110151561171357fe5b600091825260208083209091018054600160a060020a0319169055600160a060020a038516825260149052604090206001860190555b83600160a060020a03166327c1c21d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561178757600080fd5b505af115801561179b573d6000803e3d6000fd5b505050506040513d60208110156117b157600080fd5b50519150600082131561191057600188151514156117d1575060006117d6565b506000195b83600160a060020a0316638280dd8f826040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b15801561181f57600080fd5b505af1158015611833573d6000803e3d6000fd5b505050506040513d602081101561184957600080fd5b505060408051600160a060020a038b16815289151560208281019190915260609282018381528a519383019390935289517f70234ce475fee4ab40e5e55cf533f67f12b47ef4c860e62dd7affa84ead4b442938d938d938d93919291608084019185019080838360005b838110156118cb5781810151838201526020016118b3565b50505050905090810190601f1680156118f85780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a160019550611915565b600095505b50505050509392505050565b6011805482908110610e2257fe5b600354600160a060020a031681565b6009805482908110610e2257fe5b6000806000806000600e546000141561196857600194506107fa565b5060005b600e548110156107f557600a80548290811061198457fe5b600091825260209091200154600160a060020a031693508315611a415783915081600160a060020a03166395a078e8876040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b1580156119ff57600080fd5b505af1158015611a13573d6000803e3d6000fd5b505050506040513d6020811015611a2957600080fd5b5051925060018315151415611a4157600194506107fa565b60010161196c565b600b805482908110610e2257fe5b600754600160a060020a031681565b600354600160a060020a0316321480611a895750600354600160a060020a031633145b1515611a9457600080fd5b60055460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600f5481565b600354600160a060020a0316321480611af45750600354600160a060020a031633145b1515611aff57600080fd5b60128054600160a060020a031916600160a060020a0392909216919091179055565b600090565b600354600160a060020a0316321480611b495750600354600160a060020a031633145b1515611b5457600080fd5b600954600d541015611ba557806009600d54815481101515611b7257fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550611bf1565b600980546001810182556000919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af018054600160a060020a031916600160a060020a0383161790555b600d8054600101905560408051600160a060020a038316815290517f1b88a571cc8ac2e87512f05648e79d184f5cc0cbb2889bc487c41f8b9a3202eb9181900360200190a150565b60055481565b600e5481565b6003546000908190600160a060020a0316321480611c6d5750600354600160a060020a031633145b1515611c7857600080fd5b5060005b600e54811015610b145782600160a060020a0316600a82815481101515611c9f57fe5b600091825260209091200154600160a060020a03161415611dd357600a805482908110611cc857fe5b60009182526020909120018054600160a060020a0319169055600e54600019018114611d8457600a6001600e5403815481101515611d0257fe5b600091825260209091200154600a8054600160a060020a039092169183908110611d2857fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550600a6001600e5403815481101515611d6a57fe5b60009182526020909120018054600160a060020a03191690555b600e805460001901905560408051600160a060020a038516815290517fc5224c4118417a068eeac7d714e6d8af6f99ec3fb611bc965185460b0e38f0819181900360200190a160019150610b19565b600101611c7c565b600354600160a060020a0316331480611df75750611df7611b21565b1515611e0257600080fd5b60058190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b6000611e483261194c565b80611e575750611e57326106f8565b80611e665750611e6632610edb565b1515611e7157600080fd5b6040517fed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e8890600090a150600190565b600354600160a060020a0316321480611ec35750600354600160a060020a031633145b1515611ece57600080fd5b600160a060020a0381161515611ee357600080fd5b60038054600160a060020a031916600160a060020a0392909216919091179055565b60035460009081908190600160a060020a0316321480611f2f5750600354600160a060020a031633145b1515611f3a57600080fd5b5050600f546000190160005b600f548110156120a85783600160a060020a0316600b82815481101515611f6957fe5b600091825260209091200154600160a060020a031614156120a057600b805482908110611f9257fe5b60009182526020909120018054600160a060020a031916905580821461203757600b805483908110611fc057fe5b600091825260209091200154600b8054600160a060020a039092169183908110611fe657fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600b80548390811061201d57fe5b60009182526020909120018054600160a060020a03191690555b600f829055600160a060020a0384166000818152601060209081526040918290208054600160a060020a0319169055815192835290517fd41375b9d347dfe722f90a780731abd23b7855f9cf14ea7063c4cab5f9ae58e29281900390910190a16001925061103a565b600101611f46565b5060009392505050565b6040516124e1806120c3833901905600608060408190527f4f776e61626c6532303139303331353134313530304d4c00000000000000000060009081557f4564697461626c6532303139303331353134313830304d4c00000000000000006003557f42617365436f6e74656e7432303139303331353137353130304d4c0000000000600555600c556020806124e1833981016040818152915160018054600160a060020a031990811632908117909255600280548216909217909155600980548216331790819055600019600b5560068054600160a060020a038086169190941617905516825291517fc3decc188980e855666b70498ca85e8fa284d97d30483d828fa126f7303d7d199181900360200190a1506123cf806101126000396000f3006080604052600436106101a75763ffffffff60e060020a60003504166217de9881146101a957806302d05d3f146101d0578063075d4782146102015780630c6d3f931461021d5780631a735f18146102885780632310167f146102d557806327c1c21d146102ea57806332eaf21b146102ff57806336ebffca14610314578063388642841461032957806338d0f5041461034157806341c0e1b5146103fa5780634dd707881461040f5780635267db441461042457806354fd4d501461043c5780635cc4aa9b1461045157806364ade32b146104625780636d2e4b1b146104775780637ca8f618146104985780638280dd8f146104b0578063879fe48f146104c85780638da5cb5b1461055e5780638f77920114610573578063a1ff106e14610588578063b816f5131461068c578063c287e0ed146106a1578063c9e8e72d146106b6578063cbcd4461146106d7578063d810f8c8146106ec578063e02dd9c214610701578063e538530314610716578063ee56d76714610737578063f14fcbc8146107d7578063f2fde38b146107ef578063f4d9bae814610810578063f81ab0ae14610828575b005b3480156101b557600080fd5b506101be61083d565b60408051918252519081900360200190f35b3480156101dc57600080fd5b506101e5610861565b60408051600160a060020a039092168252519081900360200190f35b610209610870565b604080519115158252519081900360200190f35b34801561022957600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526102099482359460248035600160a060020a03169536959460649492019190819084018382808284375094975050933594506109d19350505050565b34801561029457600080fd5b506102a0600435610a18565b60408051600160a060020a0390951685526020850193909352600091820b90910b838301526060830152519081900360800190f35b3480156102e157600080fd5b506101e5610a4c565b3480156102f657600080fd5b506101be610a5b565b34801561030b57600080fd5b506101e5610a61565b34801561032057600080fd5b506101e5610a70565b34801561033557600080fd5b506101be600435610a7f565b34801561034d57600080fd5b506040805160206004602480358281013584810280870186019097528086526103d796843560ff1696369660449591949091019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750949750610bd19650505050505050565b604051808360000b60000b81526020018281526020019250505060405180910390f35b34801561040657600080fd5b506101a7610d5b565b34801561041b57600080fd5b506101be610e40565b34801561043057600080fd5b506101be600435610e64565b34801561044857600080fd5b506101be610ef9565b610209600435602435604435610eff565b34801561046e57600080fd5b506101be6111c0565b34801561048357600080fd5b506101a7600160a060020a03600435166111c6565b3480156104a457600080fd5b506101be600435611221565b3480156104bc57600080fd5b506101be600435611240565b3480156104d457600080fd5b506040805160206004602480358281013584810280870186019097528086526101be96843560ff1696369660449591949091019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506113b79650505050505050565b34801561056a57600080fd5b506101e5611421565b34801561057f57600080fd5b506101be611430565b60408051602060046024803582810135601f810185900485028601850190965285855261020995833560ff1695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506114369650505050505050565b34801561069857600080fd5b506101e5611927565b3480156106ad57600080fd5b506101a7611936565b3480156106c257600080fd5b506101a7600160a060020a036004351661199b565b3480156106e357600080fd5b506102096119f8565b3480156106f857600080fd5b506101be611ac9565b34801561070d57600080fd5b506101be611aed565b34801561072257600080fd5b506101a7600160a060020a0360043516611af3565b34801561074357600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610209948235946024803515159536959460649492019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750611cb79650505050505050565b3480156107e357600080fd5b506101a760043561204e565b3480156107fb57600080fd5b506101a7600160a060020a03600435166120b7565b34801561081c57600080fd5b506101be600435612129565b34801561083457600080fd5b506101be61219a565b7f5075626c6973686564000000000000000000000000000000000000000000000081565b600154600160a060020a031681565b600080600061087d6119f8565b151561088857600080fd5b600454156108c757600454600e80546001810182556000919091527fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fd01555b6108d2600f546121ac565b6000600f556108e16001611240565b50600091506000600b5413156109825750600954604080517f49102e610000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169182916349102e619160048083019260209291908290030181600087803b15801561095357600080fd5b505af1158015610967573d6000803e3d6000fd5b505050506040513d602081101561097d57600080fd5b505191505b600b54600454604080518515158152602081019390935282810191909152517f901e6f3cdc4c61620d5d424116934b9af6e31ba79cdeaa349336d93ecfe846d49181900360600190a150919050565b600854600090600160a060020a0316158015906109f85750600854600160a060020a031633145b1515610a0357600080fd5b610a0f8585858561220e565b95945050505050565b600d602052600090815260408120805460018201546002830154600390930154600160a060020a03909216939092900b9084565b600854600160a060020a031681565b600b5481565b600754600160a060020a031681565b600654600160a060020a031681565b60085460009081908190600160a060020a031615610b2f5750600854604080517f45080442000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a0390921691829163450804429160248083019260209291908290030181600087803b158015610b0057600080fd5b505af1158015610b14573d6000803e3d6000fd5b505050506040513d6020811015610b2a57600080fd5b505191505b8115610b3d57819250610bca565b831515610b6c577f5075626c697368656400000000000000000000000000000000000000000000009250610bca565b6000841215610b9d577f44726166740000000000000000000000000000000000000000000000000000009250610bca565b6000841315610bca577f447261667420696e20726576696577000000000000000000000000000000000092505b5050919050565b60085460009081908190600019908290600160a060020a031615610d2e57506008546040517f0f82c16f00000000000000000000000000000000000000000000000000000000815260ff891660048201908152606060248301908152895160648401528951600160a060020a03909416938493630f82c16f938d938d938d9360448101916084909101906020808801910280838360005b83811015610c80578181015183820152602001610c68565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610cbf578181015183820152602001610ca7565b50505050905001955050505050506040805180830381600087803b158015610ce657600080fd5b505af1158015610cfa573d6000803e3d6000fd5b505050506040513d6040811015610d1057600080fd5b508051602090910151600a54909450909250831115610d2e57606491505b8160000b6000191415610d4957600a54600095509350610d50565b8183945094505b505050935093915050565b600254600090600160a060020a0316321480610d815750600254600160a060020a031633145b1515610d8c57600080fd5b600854600160a060020a031615610e355750600854604080517f9e99bbea0000000000000000000000000000000000000000000000000000000081529051600160a060020a03909216918291639e99bbea9160048083019260209291908290030181600087803b158015610dff57600080fd5b505af1158015610e13573d6000803e3d6000fd5b505050506040513d6020811015610e2957600080fd5b505115610e3557600080fd5b610e3d612340565b50565b7f447261667400000000000000000000000000000000000000000000000000000081565b600254600090600160a060020a031632148015610e9857506000821280610e985750600082138015610e9857506000600b54125b15610ea357600b8290555b600954600160a060020a0316331415610ebc57600b8290555b600b5460408051918252517fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a29181900360200190a15050600b5490565b60055481565b6000838152600d602052604081208054829081908190600160a060020a031615801590610f4957508354600160a060020a0316331480610f495750600254600160a060020a031633145b1515610f5457600080fd5b6008548715159350600160a060020a03161561100e57600854604080517f17685953000000000000000000000000000000000000000000000000000000008152600481018b9052602481018a90529051600160a060020a03909216935083916317685953916044808201926020929091908290030181600087803b158015610fdb57600080fd5b505af1158015610fef573d6000803e3d6000fd5b505050506040513d602081101561100557600080fd5b50518015935090505b8354600160a060020a031633141561104b57821561103b576002848101805460ff1916909117905561104b565b60028401805460ff191660fe1790555b836001015484600301541015611129576002840154600090810b810b136110cc57835460408051808201909152600681527f726566756e6400000000000000000000000000000000000000000000000000006020820152600386015460018701546110c6938c93600160a060020a039091169290910361220e565b50611129565b60025460408051808201909152600e81527f72656c6561736520657363726f77000000000000000000000000000000000000602082015260038601546001870154611127938c93600160a060020a039091169290910361220e565b505b6000888152600d60209081526040808320805473ffffffffffffffffffffffffffffffffffffffff191681556001810184905560028101805460ff191690556003019290925581518a8152908101899052808201889052841515606082015290517f7f1f4b28434ce7beab4983e64a8b5bb96e195a67029fdaff925028aec57fbc6b9181900360800190a150909695505050505050565b600a5481565b600154600160a060020a031632146111dd57600080fd5b600160a060020a03811615156111f257600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600e80548290811061122f57fe5b600091825260209091200154905081565b600080600061124d6119f8565b151561125857600080fd5b600854600160a060020a031615156112dc57600254600160a060020a031632148061128d5750600254600160a060020a031633145b80156112a457508360001914806112a45750836001145b156112b1578391506112d7565b600954600160a060020a0316331480156112ce57506000600b5412155b156112d7578391505b611374565b50600854604080517f3513a805000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a03909216918291633513a8059160248083019260209291908290030181600087803b15801561134557600080fd5b505af1158015611359573d6000803e3d6000fd5b505050506040513d602081101561136f57600080fd5b505191505b600b8290556040805183815290517fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a29181900360200190a15050600b5492915050565b60008060006113c7868686610bd1565b92509050600081900b156113da57600080fd5b6040805160ff881681526020810184905281517fa58326ee5bb617cb8b4f0d0f5f557c469d2d05d7a738f777037deda9c724b370929181900390910190a150949350505050565b600254600160a060020a031681565b600c5481565b600080600061144361237c565b600c80546001019055600954604080517f95a078e80000000000000000000000000000000000000000000000000000000081523260048201529051600160a060020a0390921694506000918291829187916395a078e891602480830192602092919082900301818787803b1580156114ba57600080fd5b505af11580156114ce573d6000803e3d6000fd5b505050506040513d60208110156114e457600080fd5b505115156114f157600080fd5b600254600160a060020a0316321461151d5761150e8c8a8a6113b7565b94503485111561151d57600080fd5b60408051608081018252338152346020808301918252600083850181815260608501828152600c548352600d90935294812084518154600160a060020a0391821673ffffffffffffffffffffffffffffffffffffffff199091161782559351600182015594516002860180549190920b60ff1660ff199091161790555160039093019290925560085490955016156116e557600860009054906101000a9004600160a060020a0316925082600160a060020a031663123e0e80600c548e8c8c6040518563ffffffff1660e060020a028152600401808581526020018460ff1660ff1681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561164457818101518382015260200161162c565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561168357818101518382015260200161166b565b505050509050019650505050505050602060405180830381600087803b1580156116ac57600080fd5b505af11580156116c0573d6000803e3d6000fd5b505050506040513d60208110156116d657600080fd5b5051915081156116e557600080fd5b7f089a6f1788a3c353423e1be4ba12533bdde7d908bb41abeee185af0acb3df562600c548d6004548e8e604051808681526020018560ff1660ff16815260200184600019166000191681526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611770578181015183820152602001611758565b50505050905090810190601f16801561179d5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156117d05781810151838201526020016117b8565b50505050905090810190601f1680156117fd5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15060005b885181101561188d57888181518110151561182a57fe5b6020908102909101015115611885577f515e0a48b385fce2a8e4d9f169a97c4f6ea669a752358f5e6ab37cc3c2e84c38898281518110151561186857fe5b602090810290910181015160408051918252519081900390910190a15b600101611813565b5060005b87518110156119155788818151811015156118a857fe5b602090810290910101511561190d577fb6e3239e521a6c66920ae634f8e921a37e6991d520ac44d52f8516397f41b68488828151811015156118e657fe5b602090810290910181015160408051600160a060020a039092168252519081900390910190a15b600101611891565b5060019b9a5050505050505050505050565b600954600160a060020a031681565b600254600160a060020a03163214806119595750600254600160a060020a031633145b151561196457600080fd5b60045460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600254600160a060020a03163214806119be5750600254600160a060020a031633145b15156119c957600080fd5b6007805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6002546000908190600160a060020a0316331480611a205750600954600160a060020a031633145b15611a2e5760019150611ac5565b50600954604080517f26683e140000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169182916326683e149160248083019260209291908290030181600087803b158015611a9657600080fd5b505af1158015611aaa573d6000803e3d6000fd5b505050506040513d6020811015611ac057600080fd5b505191505b5090565b7f447261667420696e20726576696577000000000000000000000000000000000081565b60045481565b60025460009081908190600160a060020a0316321480611b1d5750600254600160a060020a031633145b1515611b2857600080fd5b600854600160a060020a031615611bc857600860009054906101000a9004600160a060020a0316925082600160a060020a0316639e99bbea6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611b8f57600080fd5b505af1158015611ba3573d6000803e3d6000fd5b505050506040513d6020811015611bb957600080fd5b505191508115611bc857600080fd5b6008805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03861690811790915515611c725783925082600160a060020a0316637b1cdb3e6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611c3957600080fd5b505af1158015611c4d573d6000803e3d6000fd5b505050506040513d6020811015611c6357600080fd5b505190508015611c7257600080fd5b60085460408051600160a060020a039092168252517fa6f2e38f0cfebf27212317fced3ac40bc62e00bd33f38d69603710740c69acb79181900360200190a150505050565b600254600090819081908190600160a060020a0316331480611ce35750600754600160a060020a031633145b1515611cee57600080fd5b6000888152600d602052604090208054909350600160a060020a03161515611d1557600080fd5b600854879250600160a060020a031615611dce5750600854604080517fe870ed91000000000000000000000000000000000000000000000000000000008152600481018a905288151560248201529051600160a060020a0390921691829163e870ed919160448083019260209291908290030181600087803b158015611d9a57600080fd5b505af1158015611dae573d6000803e3d6000fd5b505050506040513d6020811015611dc457600080fd5b5051159150611ea2565b826001015483600301541015611ea257861515611e4557825460408051808201909152600f81527f616363657373206465636c696e65640000000000000000000000000000000000602082015260038501546001860154611e3f938c93600160a060020a039091169290910361220e565b50611ea2565b60025460408051808201909152600d81527f6f776e6572207061796d656e7400000000000000000000000000000000000000602082015260038501546001860154611ea0938c93600160a060020a039091169290910361220e565b505b60018215151415611fdd5760028301805460ff19166001908117909155604080518a8152602080820184905260809282018381528a519383019390935289517f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718948d9490938c938c93919291606084019160a08501919087019080838360005b83811015611f3a578181015183820152602001611f22565b50505050905090810190601f168015611f675780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611f9a578181015183820152602001611f82565b50505050905090810190601f168015611fc75780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a1612043565b60028301805460ff191660ff179055604080518981526000602082018190526080828401819052820181905260c06060830181905282015290517f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718918190036101000190a15b509695505050505050565b600254600160a060020a03163214806120715750600254600160a060020a031633145b151561207c57600080fd5b600f8190556040805182815290517fa288d3a2aba7d5dec44e93ff46d2f1129e251695be2046de105f9a9c6feefcaa9181900360200190a150565b600254600160a060020a03163214806120da5750600254600160a060020a031633145b15156120e557600080fd5b600160a060020a03811615156120fa57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600090600160a060020a031632148061214f5750600254600160a060020a031633145b151561215a57600080fd5b600a8290556040805183815290517f4114f8ef80b6de2161db580cbefa14e1892d15d3ebe2062c9914e4a5773114a39181900360200190a15050600a5490565b60006121a7600b54610a7f565b905090565b600254600160a060020a03163314806121c857506121c86119f8565b15156121d357600080fd5b60048190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b6000848152600d602052604081206001810154600382015484011161233757604051600160a060020a0386169084156108fc029085906000818181858888f19350505050158015612263573d6000803e3d6000fd5b508281600301540181600301819055507fad58d18ea7292f887da6f15bb4f0badddaa33d169713d09cf49710acc7c3a5b986858786604051808581526020018060200184600160a060020a0316600160a060020a03168152602001838152602001828103825285818151815260200191508051906020019080838360005b838110156122f95781810151838201526020016122e1565b50505050905090810190601f1680156123265780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a15b50949350505050565b600254600160a060020a03163214806123635750600254600160a060020a031633145b151561236e57600080fd5b600254600160a060020a0316ff5b604080516080810182526000808252602082018190529181018290526060810191909152905600a165627a7a723058202fb499b0671f272703d23a44987b8c3eea18538337d5e28f0d8dc8fca16109ae0029a165627a7a723058207dbf31d67589062030b4a2b7fe318b722e198afe2c9baf9ed2c099a1cc400bd5002960806040527f4f776e61626c6532303139303331353134313530304d4c0000000000000000006000557f42734163636573734374726c47727032303139303331353137323930304d4c0060035534801561005857600080fd5b506040516020806107a983398101604090815290516001805432600160a060020a0319918216811780845560028054841690921790915560048054909216600160a060020a03948516179091558216600090815260066020908152848220805460ff199081168517909155835490941682526005905292832080549092161790556106c09081906100e990396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100dc57806308ae4b0c1461010d5780630ae5e7391461014257806341c0e1b51461016357806342e7ba7b1461017857806354fd4d50146101995780636d2e4b1b146101c057806375861a95146101e157806385e68531146102025780638da5cb5b1461022357806395a078e814610238578063af570c0414610259578063cdb849b71461026e578063f2fde38b1461028f578063fdff9b4d146102b0575b005b3480156100e857600080fd5b506100f16102d1565b60408051600160a060020a039092168252519081900360200190f35b34801561011957600080fd5b5061012e600160a060020a03600435166102e0565b604080519115158252519081900360200190f35b34801561014e57600080fd5b506100da600160a060020a03600435166102f5565b34801561016f57600080fd5b506100da610371565b34801561018457600080fd5b5061012e600160a060020a03600435166103ad565b3480156101a557600080fd5b506101ae6103d0565b60408051918252519081900360200190f35b3480156101cc57600080fd5b506100da600160a060020a03600435166103d6565b3480156101ed57600080fd5b506100da600160a060020a0360043516610431565b34801561020e57600080fd5b506100da600160a060020a03600435166104ba565b34801561022f57600080fd5b506100f1610548565b34801561024457600080fd5b5061012e600160a060020a0360043516610557565b34801561026557600080fd5b506100f161057a565b34801561027a57600080fd5b506100da600160a060020a0360043516610589565b34801561029b57600080fd5b506100da600160a060020a036004351661060d565b3480156102bc57600080fd5b5061012e600160a060020a036004351661067f565b600154600160a060020a031681565b60056020526000908152604090205460ff1681565b3360009081526006602052604090205460ff16151560011461031657600080fd5b600160a060020a038116600081815260056020908152604091829020805460ff19166001179055815192835290517fb251eb052afc73ffd02ffe85ad79990a8b3fed60d76dbc2fa2fdd7123dffd9149281900390910190a150565b600254600160a060020a03163214806103945750600254600160a060020a031633145b151561039f57600080fd5b600254600160a060020a0316ff5b600160a060020a031660009081526006602052604090205460ff16151560011490565b60035481565b600154600160a060020a031632146103ed57600080fd5b600160a060020a038116151561040257600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a03163214806104545750600254600160a060020a031633145b151561045f57600080fd5b600160a060020a038116600081815260066020908152604091829020805460ff19166001179055815192835290517f93bcaab179551bde429187645251f8e1fb8ac85801fcb1cf91eb2c9043d611179281900390910190a150565b3360009081526006602052604090205460ff161515600114806104e5575033600160a060020a038216145b15156104f057600080fd5b600160a060020a038116600081815260056020908152604091829020805460ff19169055815192835290517f745cd29407db644ed93e3ceb61cbcab96d1dfb496989ac5d5bf514fc5a9fab9c9281900390910190a150565b600254600160a060020a031681565b600160a060020a031660009081526005602052604090205460ff16151560011490565b600454600160a060020a031681565b600254600160a060020a03163314806105aa575033600160a060020a038216145b15156105b557600080fd5b600160a060020a038116600081815260066020908152604091829020805460ff19169055815192835290517f2d6aa1a9629d125e23a0cf692cda7cd6795dff1652eedd4673b38ec31e387b959281900390910190a150565b600254600160a060020a03163214806106305750600254600160a060020a031633145b151561063b57600080fd5b600160a060020a038116151561065057600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60066020526000908152604090205460ff16815600a165627a7a72305820a3d364eca755c88b24e28f3029b5ff9767cb2358a596e6de29fab07635e9de520029608060408190527f41636365737369626c6532303139303232323133353930304d4c0000000000006000557f4f776e61626c6532303139303331353134313530304d4c0000000000000000006001557f4564697461626c6532303139303331353134313830304d4c00000000000000006004557f42617365436f6e74656e745479706532303139303331383130313230304d4c0060065560208061057883398101604052516002805432600160a060020a0319918216811790925560038054821690921790915560078054909116600160a060020a0390921691909117905561048b806100ed6000396000f3006080604052600436106100b95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100bb57806341c0e1b5146100ec57806354fd4d50146101015780636d2e4b1b146101285780638da5cb5b14610149578063af570c041461015e578063c287e0ed14610173578063cbcd446114610188578063e02dd9c2146101b1578063f14fcbc8146101c6578063f1551887146101de578063f2fde38b146101f3575b005b3480156100c757600080fd5b506100d0610214565b60408051600160a060020a039092168252519081900360200190f35b3480156100f857600080fd5b506100b9610223565b34801561010d57600080fd5b5061011661025f565b60408051918252519081900360200190f35b34801561013457600080fd5b506100b9600160a060020a0360043516610265565b34801561015557600080fd5b506100d06102c0565b34801561016a57600080fd5b506100d06102cf565b34801561017f57600080fd5b506100b96102de565b34801561019457600080fd5b5061019d610343565b604080519115158252519081900360200190f35b3480156101bd57600080fd5b50610116610354565b3480156101d257600080fd5b506100b960043561035a565b3480156101ea57600080fd5b5061019d6103bc565b3480156101ff57600080fd5b506100b9600160a060020a03600435166103ed565b600254600160a060020a031681565b600354600160a060020a03163214806102465750600354600160a060020a031633145b151561025157600080fd5b600354600160a060020a0316ff5b60065481565b600254600160a060020a0316321461027c57600080fd5b600160a060020a038116151561029157600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600354600160a060020a031681565b600754600160a060020a031681565b600354600160a060020a03163214806103015750600354600160a060020a031633145b151561030c57600080fd5b60055460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600354600160a060020a0316331490565b60055481565b600354600160a060020a03163314806103765750610376610343565b151561038157600080fd5b60058190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b6040516000907fed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e88908290a150600190565b600354600160a060020a03163214806104105750600354600160a060020a031633145b151561041b57600080fd5b600160a060020a038116151561043057600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582041b0d335b9b4d89aa7e09dd15c22e519ef4e22dba7d5570bb82b50c0bb682de20029a165627a7a72305820f11fb3f913afcc66ed61e5d2ea0bfbf09602a396641a1e76d1cac8fa49d2bd520029`

// DeployBaseFactory deploys a new Ethereum contract, binding an instance of BaseFactory to it.
func DeployBaseFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BaseFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaseFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseFactory{BaseFactoryCaller: BaseFactoryCaller{contract: contract}, BaseFactoryTransactor: BaseFactoryTransactor{contract: contract}, BaseFactoryFilterer: BaseFactoryFilterer{contract: contract}}, nil
}

// BaseFactory is an auto generated Go binding around an Ethereum contract.
type BaseFactory struct {
	BaseFactoryCaller     // Read-only binding to the contract
	BaseFactoryTransactor // Write-only binding to the contract
	BaseFactoryFilterer   // Log filterer for contract events
}

// BaseFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseFactorySession struct {
	Contract     *BaseFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseFactoryCallerSession struct {
	Contract *BaseFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BaseFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseFactoryTransactorSession struct {
	Contract     *BaseFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BaseFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseFactoryRaw struct {
	Contract *BaseFactory // Generic contract binding to access the raw methods on
}

// BaseFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseFactoryCallerRaw struct {
	Contract *BaseFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// BaseFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseFactoryTransactorRaw struct {
	Contract *BaseFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseFactory creates a new instance of BaseFactory, bound to a specific deployed contract.
func NewBaseFactory(address common.Address, backend bind.ContractBackend) (*BaseFactory, error) {
	contract, err := bindBaseFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseFactory{BaseFactoryCaller: BaseFactoryCaller{contract: contract}, BaseFactoryTransactor: BaseFactoryTransactor{contract: contract}, BaseFactoryFilterer: BaseFactoryFilterer{contract: contract}}, nil
}

// NewBaseFactoryCaller creates a new read-only instance of BaseFactory, bound to a specific deployed contract.
func NewBaseFactoryCaller(address common.Address, caller bind.ContractCaller) (*BaseFactoryCaller, error) {
	contract, err := bindBaseFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseFactoryCaller{contract: contract}, nil
}

// NewBaseFactoryTransactor creates a new write-only instance of BaseFactory, bound to a specific deployed contract.
func NewBaseFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseFactoryTransactor, error) {
	contract, err := bindBaseFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseFactoryTransactor{contract: contract}, nil
}

// NewBaseFactoryFilterer creates a new log filterer instance of BaseFactory, bound to a specific deployed contract.
func NewBaseFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseFactoryFilterer, error) {
	contract, err := bindBaseFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseFactoryFilterer{contract: contract}, nil
}

// bindBaseFactory binds a generic wrapper to an already deployed contract.
func bindBaseFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseFactory *BaseFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseFactory.Contract.BaseFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseFactory *BaseFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseFactory.Contract.BaseFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseFactory *BaseFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseFactory.Contract.BaseFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseFactory *BaseFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseFactory *BaseFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseFactory *BaseFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseFactory.Contract.contract.Transact(opts, method, params...)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseFactory *BaseFactoryCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseFactory.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseFactory *BaseFactorySession) Creator() (common.Address, error) {
	return _BaseFactory.Contract.Creator(&_BaseFactory.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_BaseFactory *BaseFactoryCallerSession) Creator() (common.Address, error) {
	return _BaseFactory.Contract.Creator(&_BaseFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseFactory *BaseFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseFactory.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseFactory *BaseFactorySession) Owner() (common.Address, error) {
	return _BaseFactory.Contract.Owner(&_BaseFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseFactory *BaseFactoryCallerSession) Owner() (common.Address, error) {
	return _BaseFactory.Contract.Owner(&_BaseFactory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseFactory *BaseFactoryCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseFactory.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseFactory *BaseFactorySession) Version() ([32]byte, error) {
	return _BaseFactory.Contract.Version(&_BaseFactory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseFactory *BaseFactoryCallerSession) Version() ([32]byte, error) {
	return _BaseFactory.Contract.Version(&_BaseFactory.CallOpts)
}

// CreateContentType is a paid mutator transaction binding the contract method 0xb8cfaf05.
//
// Solidity: function createContentType() returns(address)
func (_BaseFactory *BaseFactoryTransactor) CreateContentType(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseFactory.contract.Transact(opts, "createContentType")
}

// CreateContentType is a paid mutator transaction binding the contract method 0xb8cfaf05.
//
// Solidity: function createContentType() returns(address)
func (_BaseFactory *BaseFactorySession) CreateContentType() (*types.Transaction, error) {
	return _BaseFactory.Contract.CreateContentType(&_BaseFactory.TransactOpts)
}

// CreateContentType is a paid mutator transaction binding the contract method 0xb8cfaf05.
//
// Solidity: function createContentType() returns(address)
func (_BaseFactory *BaseFactoryTransactorSession) CreateContentType() (*types.Transaction, error) {
	return _BaseFactory.Contract.CreateContentType(&_BaseFactory.TransactOpts)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x575185ed.
//
// Solidity: function createGroup() returns(address)
func (_BaseFactory *BaseFactoryTransactor) CreateGroup(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseFactory.contract.Transact(opts, "createGroup")
}

// CreateGroup is a paid mutator transaction binding the contract method 0x575185ed.
//
// Solidity: function createGroup() returns(address)
func (_BaseFactory *BaseFactorySession) CreateGroup() (*types.Transaction, error) {
	return _BaseFactory.Contract.CreateGroup(&_BaseFactory.TransactOpts)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x575185ed.
//
// Solidity: function createGroup() returns(address)
func (_BaseFactory *BaseFactoryTransactorSession) CreateGroup() (*types.Transaction, error) {
	return _BaseFactory.Contract.CreateGroup(&_BaseFactory.TransactOpts)
}

// CreateLibrary is a paid mutator transaction binding the contract method 0x40b89f06.
//
// Solidity: function createLibrary(address address_KMS) returns(address)
func (_BaseFactory *BaseFactoryTransactor) CreateLibrary(opts *bind.TransactOpts, address_KMS common.Address) (*types.Transaction, error) {
	return _BaseFactory.contract.Transact(opts, "createLibrary", address_KMS)
}

// CreateLibrary is a paid mutator transaction binding the contract method 0x40b89f06.
//
// Solidity: function createLibrary(address address_KMS) returns(address)
func (_BaseFactory *BaseFactorySession) CreateLibrary(address_KMS common.Address) (*types.Transaction, error) {
	return _BaseFactory.Contract.CreateLibrary(&_BaseFactory.TransactOpts, address_KMS)
}

// CreateLibrary is a paid mutator transaction binding the contract method 0x40b89f06.
//
// Solidity: function createLibrary(address address_KMS) returns(address)
func (_BaseFactory *BaseFactoryTransactorSession) CreateLibrary(address_KMS common.Address) (*types.Transaction, error) {
	return _BaseFactory.Contract.CreateLibrary(&_BaseFactory.TransactOpts, address_KMS)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseFactory *BaseFactoryTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseFactory.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseFactory *BaseFactorySession) Kill() (*types.Transaction, error) {
	return _BaseFactory.Contract.Kill(&_BaseFactory.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BaseFactory *BaseFactoryTransactorSession) Kill() (*types.Transaction, error) {
	return _BaseFactory.Contract.Kill(&_BaseFactory.TransactOpts)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseFactory *BaseFactoryTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _BaseFactory.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseFactory *BaseFactorySession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseFactory.Contract.TransferCreatorship(&_BaseFactory.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_BaseFactory *BaseFactoryTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _BaseFactory.Contract.TransferCreatorship(&_BaseFactory.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseFactory *BaseFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseFactory *BaseFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseFactory.Contract.TransferOwnership(&_BaseFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseFactory *BaseFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseFactory.Contract.TransferOwnership(&_BaseFactory.TransactOpts, newOwner)
}

// BaseLibraryABI is the input ABI used to generate the binding from.
const BaseLibraryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"canContribute\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content_type\",\"type\":\"address\"},{\"name\":\"content_contract\",\"type\":\"address\"}],\"name\":\"addContentType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"approvalRequestsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"removeReviewerGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contentTypeContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"addAccessorGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content_type\",\"type\":\"address\"}],\"name\":\"createContent\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reviewerGroupsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contributorGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"canNodePublish\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"canReview\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"content_type\",\"type\":\"address\"}],\"name\":\"validType\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accessorGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressKMS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"removeContributorGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contributorGroupsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"submitApprovalRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingApprovalRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"addContributorGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content_contract\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"},{\"name\":\"note\",\"type\":\"string\"}],\"name\":\"approveContent\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approvalRequests\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reviewerGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"hasAccess\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contentTypes\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentTypesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"address_KMS\",\"type\":\"address\"}],\"name\":\"setAddressKMS\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canPublish\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"addReviewerGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"objectHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accessorGroupsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"removeAccessorGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"object_hash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accessRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content_type\",\"type\":\"address\"}],\"name\":\"removeContentType\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"address_KMS\",\"type\":\"address\"},{\"name\":\"content_space\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"content_type\",\"type\":\"address\"}],\"name\":\"ContentObjectCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ContributorGroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ContributorGroupRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ReviewerGroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ReviewerGroupRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"AccessorGroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"AccessorGroupRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentType\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"contentContract\",\"type\":\"address\"}],\"name\":\"ContentTypeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentType\",\"type\":\"address\"}],\"name\":\"ContentTypeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"operationCode\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"UnauthorizedOperation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"ApproveContentRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contentAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"note\",\"type\":\"string\"}],\"name\":\"ApproveContent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"Commit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"UpdateRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AccessRequest\",\"type\":\"event\"}]"

// BaseLibraryBin is the compiled bytecode used for deploying new contracts.
const BaseLibraryBin = `0x608060408181527f41636365737369626c6532303139303232323133353930304d4c00000000000060009081557f4f776e61626c6532303139303331353134313530304d4c0000000000000000006001557f4564697461626c6532303139303331353134313830304d4c00000000000000006004557f426173654c69627261727932303139303331383130313330304d4c0000000000600655600c819055600d819055600e819055600f819055601355806146e48339810160405280516020909101516002805432600160a060020a03199182168117909255600380548216909217909155600780548216600160a060020a0393841617905560128054909116919092161790556145cf806101156000396000f3006080604052600436106101ea5763ffffffff60e060020a60003504166302d05d3f81146101ec5780630eaec2c51461021d5780630f58a7861461025257806316308394146102795780631b969895146102a05780631cdbee5a146102c15780631d0f4351146102e25780631e35d8fa1461030357806321770a84146103245780632393553b1461033957806326683e141461035157806329d002191461037257806329dedde5146103935780632c11f392146103b457806332eaf21b146103cc578063386493e0146103e157806341c0e1b514610402578063470750bb1461041757806349102e611461042c57806354fd4d501461044157806363dab9d414610456578063679a9a3c1461046e5780636d2e4b1b1461048f57806387e86b2c146104b05780638cb13c2e1461051b5780638da5cb5b14610533578063952e464b1461054857806395a078e814610560578063991a3a7c14610581578063af570c0414610599578063c287e0ed146105ae578063c65bcbe2146105c3578063c9e8e72d146105d8578063cbcd4461146105f9578063dc3c29c01461060e578063e02dd9c21461062f578063e5538fd214610644578063e8de515f14610659578063f14fcbc81461067a578063f155188714610692578063f2fde38b146106a7578063fd089196146106c8575b005b3480156101f857600080fd5b506102016106e9565b60408051600160a060020a039092168252519081900360200190f35b34801561022957600080fd5b5061023e600160a060020a03600435166106f8565b604080519115158252519081900360200190f35b34801561025e57600080fd5b506101ea600160a060020a0360043581169060243516610803565b34801561028557600080fd5b5061028e610978565b60408051918252519081900360200190f35b3480156102ac57600080fd5b5061023e600160a060020a036004351661097e565b3480156102cd57600080fd5b50610201600160a060020a0360043516610b1f565b3480156102ee57600080fd5b506101ea600160a060020a0360043516610b3a565b34801561030f57600080fd5b50610201600160a060020a0360043516610c4d565b34801561033057600080fd5b5061028e610e0e565b34801561034557600080fd5b50610201600435610e14565b34801561035d57600080fd5b5061023e600160a060020a0360043516610e3c565b34801561037e57600080fd5b5061023e600160a060020a0360043516610edb565b34801561039f57600080fd5b5061023e600160a060020a0360043516610fe6565b3480156103c057600080fd5b50610201600435611041565b3480156103d857600080fd5b5061020161104f565b3480156103ed57600080fd5b5061023e600160a060020a036004351661105e565b34801561040e57600080fd5b506101ea6111f4565b34801561042357600080fd5b5061028e611230565b34801561043857600080fd5b5061023e611236565b34801561044d57600080fd5b5061028e61144d565b34801561046257600080fd5b50610201600435611453565b34801561047a57600080fd5b506101ea600160a060020a03600435166114a0565b34801561049b57600080fd5b506101ea600160a060020a03600435166115b3565b3480156104bc57600080fd5b50604080516020600460443581810135601f810184900484028501840190955284845261023e948235600160a060020a031694602480351515953695946064949201919081908401838280828437509497506116019650505050505050565b34801561052757600080fd5b50610201600435611921565b34801561053f57600080fd5b5061020161192f565b34801561055457600080fd5b5061020160043561193e565b34801561056c57600080fd5b5061023e600160a060020a036004351661194c565b34801561058d57600080fd5b50610201600435611a49565b3480156105a557600080fd5b50610201611a57565b3480156105ba57600080fd5b506101ea611a66565b3480156105cf57600080fd5b5061028e611acb565b3480156105e457600080fd5b506101ea600160a060020a0360043516611ad1565b34801561060557600080fd5b5061023e611b21565b34801561061a57600080fd5b506101ea600160a060020a0360043516611b26565b34801561063b57600080fd5b5061028e611c39565b34801561065057600080fd5b5061028e611c3f565b34801561066557600080fd5b5061023e600160a060020a0360043516611c45565b34801561068657600080fd5b506101ea600435611ddb565b34801561069e57600080fd5b5061023e611e3d565b3480156106b357600080fd5b506101ea600160a060020a0360043516611ea0565b3480156106d457600080fd5b5061023e600160a060020a0360043516611f05565b600254600160a060020a031681565b6000806000806000600c546000141561071457600194506107fa565b5060005b600c548110156107f557600880548290811061073057fe5b600091825260209091200154600160a060020a0316935083156107ed5783915081600160a060020a03166395a078e8876040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b1580156107ab57600080fd5b505af11580156107bf573d6000803e3d6000fd5b505050506040513d60208110156107d557600080fd5b50519250600183151514156107ed57600194506107fa565b600101610718565b600094505b50505050919050565b600354600160a060020a03163214806108265750600354600160a060020a031633145b151561083157600080fd5b600160a060020a038083166000908152601060205260409020541615801561085f575061085d82610fe6565b155b1561090b57600b54600f5410156108b55781600b600f5481548110151561088257fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550610901565b600b80546001810182556000919091527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db9018054600160a060020a031916600160a060020a0384161790555b600f805460010190555b600160a060020a038281166000818152601060209081526040918290208054600160a060020a0319169486169485179055815192835282019290925281517f280016f7418306a55542432120fd1a239ef9fcc1a92694d8d44ca76be0249ea7929181900390910190a15050565b60135481565b6003546000908190600160a060020a03163214806109a65750600354600160a060020a031633145b15156109b157600080fd5b5060005b600d54811015610b145782600160a060020a03166009828154811015156109d857fe5b600091825260209091200154600160a060020a03161415610b0c576009805482908110610a0157fe5b60009182526020909120018054600160a060020a0319169055600d54600019018114610abd5760096001600d5403815481101515610a3b57fe5b60009182526020909120015460098054600160a060020a039092169183908110610a6157fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a0316021790555060096001600d5403815481101515610aa357fe5b60009182526020909120018054600160a060020a03191690555b600d805460001901905560408051600160a060020a038516815290517fdf9d78c5635b72b709c85300a786eb7238acbe5bffe01c60c16464e45c6eb6eb9181900360200190a160019150610b19565b6001016109b5565b600091505b50919050565b601060205260009081526040902054600160a060020a031681565b600354600160a060020a0316321480610b5d5750600354600160a060020a031633145b1515610b6857600080fd5b600a54600e541015610bb95780600a600e54815481101515610b8657fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550610c05565b600a80546001810182556000919091527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8018054600160a060020a031916600160a060020a0383161790555b600e8054600101905560408051600160a060020a038316815290517f3a94857e4393737f73edb175a7d0c195c7f635d9ae995e12740616ec55c9d4119181900360200190a150565b600080610c59326106f8565b1515610c6457600080fd5b600f5415610c8057610c7583610fe6565b1515610c8057600080fd5b82610c896120b2565b600160a060020a03909116815260405190819003602001906000f080158015610cb6573d6000803e3d6000fd5b50601254604080517fc9e8e72d000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015290519293509083169163c9e8e72d9160248082019260009290919082900301818387803b158015610d2157600080fd5b505af1158015610d35573d6000803e3d6000fd5b50505050600160a060020a038381166000908152601060205260408082205481517fe5385303000000000000000000000000000000000000000000000000000000008152908416600482015290519284169263e53853039260248084019391929182900301818387803b158015610dab57600080fd5b505af1158015610dbf573d6000803e3d6000fd5b505060408051600160a060020a0380861682528716602082015281517f3981e74ab81857b375ec391a4f7c31ee89462cd927de6d8fbdb98f77da009c569450908190039091019150a192915050565b600d5481565b6008805482908110610e2257fe5b600091825260209091200154600160a060020a0316905081565b600754604080517f26683e14000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151600093929092169182916326683e1491602480830192602092919082900301818887803b158015610ea857600080fd5b505af1158015610ebc573d6000803e3d6000fd5b505050506040513d6020811015610ed257600080fd5b50519392505050565b6003546000908190819081908190600160a060020a0387811691161415610f0557600194506107fa565b5060005b600d548110156107f5576009805482908110610f2157fe5b600091825260209091200154600160a060020a031693508315610fde5783915081600160a060020a03166395a078e8876040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b158015610f9c57600080fd5b505af1158015610fb0573d6000803e3d6000fd5b505050506040513d6020811015610fc657600080fd5b5051925060018315151415610fde57600194506107fa565b600101610f09565b600080805b600f548110156110365783600160a060020a0316600b8281548110151561100e57fe5b600091825260209091200154600160a060020a0316141561102e57600191505b600101610feb565b8192505b5050919050565b600a805482908110610e2257fe5b601254600160a060020a031681565b6003546000908190600160a060020a03163214806110865750600354600160a060020a031633145b151561109157600080fd5b5060005b600c54811015610b145782600160a060020a03166008828154811015156110b857fe5b600091825260209091200154600160a060020a031614156111ec5760088054829081106110e157fe5b60009182526020909120018054600160a060020a0319169055600c5460001901811461119d5760086001600c540381548110151561111b57fe5b60009182526020909120015460088054600160a060020a03909216918390811061114157fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a0316021790555060086001600c540381548110151561118357fe5b60009182526020909120018054600160a060020a03191690555b600c805460001901905560408051600160a060020a038516815290517fbbd97daa1862eb12f77ed128a557406737cee07b131b1e2d7140dff2005e197c9181900360200190a160019150610b19565b600101611095565b600354600160a060020a03163214806112175750600354600160a060020a031633145b151561122257600080fd5b600354600160a060020a0316ff5b600c5481565b600d546000903390819015156113165780600160a060020a0316638280dd8f60006040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b15801561129057600080fd5b505af11580156112a4573d6000803e3d6000fd5b505050506040513d60208110156112ba57600080fd5b505060408051600160a060020a038416815260016020820152606081830181905260009082015290517f70234ce475fee4ab40e5e55cf533f67f12b47ef4c860e62dd7affa84ead4b4429181900360a00190a160019250611448565b600160a060020a0382166000908152601460205260409020541561133d5760009250611448565b601154601354101561138e5781601160135481548110151561135b57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055506113da565b601180546001810182556000919091527f31ecc21a745e3968a04e9570e4425bc18fa8019c68028196b546d1669c200c68018054600160a060020a031916600160a060020a0384161790555b60138054600160a060020a0384166000818152601460209081526040918290206001948501905584549093019093558251908152329181019190915281517f0588a34cf0de4e025d359c89ca4bacbcbf175440909952d91c814412d9da996a929181900390910190a1600192505b505090565b60065481565b60006013546000148061146857508160135411155b156114755750600061149b565b601180548390811061148357fe5b600091825260209091200154600160a060020a031690505b919050565b600354600160a060020a03163214806114c35750600354600160a060020a031633145b15156114ce57600080fd5b600854600c54101561151f57806008600c548154811015156114ec57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a0316021790555061156b565b600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3018054600160a060020a031916600160a060020a0383161790555b600c8054600101905560408051600160a060020a038316815290517f218673669018c25b89bfbf1b58d0075e37c8847ef16e707b92355b7833e97d619181900360200190a150565b600254600160a060020a031632146115ca57600080fd5b600160a060020a03811615156115df57600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b60008060008060008061161332610edb565b151560011461162157600080fd5b600160a060020a0389166000908152601460205260409020546011805460001990920196508a9550908690811061165457fe5b600091825260208083209091018054600160a060020a031916905560138054600019018155600160a060020a038c16835260149091526040822091909155548510156117495760116013548154811015156116ab57fe5b60009182526020909120015460118054600160a060020a0390921694508491879081106116d457fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550601160135481548110151561171357fe5b600091825260208083209091018054600160a060020a0319169055600160a060020a038516825260149052604090206001860190555b83600160a060020a03166327c1c21d6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561178757600080fd5b505af115801561179b573d6000803e3d6000fd5b505050506040513d60208110156117b157600080fd5b50519150600082131561191057600188151514156117d1575060006117d6565b506000195b83600160a060020a0316638280dd8f826040518263ffffffff1660e060020a02815260040180828152602001915050602060405180830381600087803b15801561181f57600080fd5b505af1158015611833573d6000803e3d6000fd5b505050506040513d602081101561184957600080fd5b505060408051600160a060020a038b16815289151560208281019190915260609282018381528a519383019390935289517f70234ce475fee4ab40e5e55cf533f67f12b47ef4c860e62dd7affa84ead4b442938d938d938d93919291608084019185019080838360005b838110156118cb5781810151838201526020016118b3565b50505050905090810190601f1680156118f85780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a160019550611915565b600095505b50505050509392505050565b6011805482908110610e2257fe5b600354600160a060020a031681565b6009805482908110610e2257fe5b6000806000806000600e546000141561196857600194506107fa565b5060005b600e548110156107f557600a80548290811061198457fe5b600091825260209091200154600160a060020a031693508315611a415783915081600160a060020a03166395a078e8876040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b1580156119ff57600080fd5b505af1158015611a13573d6000803e3d6000fd5b505050506040513d6020811015611a2957600080fd5b5051925060018315151415611a4157600194506107fa565b60010161196c565b600b805482908110610e2257fe5b600754600160a060020a031681565b600354600160a060020a0316321480611a895750600354600160a060020a031633145b1515611a9457600080fd5b60055460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600f5481565b600354600160a060020a0316321480611af45750600354600160a060020a031633145b1515611aff57600080fd5b60128054600160a060020a031916600160a060020a0392909216919091179055565b600090565b600354600160a060020a0316321480611b495750600354600160a060020a031633145b1515611b5457600080fd5b600954600d541015611ba557806009600d54815481101515611b7257fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550611bf1565b600980546001810182556000919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af018054600160a060020a031916600160a060020a0383161790555b600d8054600101905560408051600160a060020a038316815290517f1b88a571cc8ac2e87512f05648e79d184f5cc0cbb2889bc487c41f8b9a3202eb9181900360200190a150565b60055481565b600e5481565b6003546000908190600160a060020a0316321480611c6d5750600354600160a060020a031633145b1515611c7857600080fd5b5060005b600e54811015610b145782600160a060020a0316600a82815481101515611c9f57fe5b600091825260209091200154600160a060020a03161415611dd357600a805482908110611cc857fe5b60009182526020909120018054600160a060020a0319169055600e54600019018114611d8457600a6001600e5403815481101515611d0257fe5b600091825260209091200154600a8054600160a060020a039092169183908110611d2857fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550600a6001600e5403815481101515611d6a57fe5b60009182526020909120018054600160a060020a03191690555b600e805460001901905560408051600160a060020a038516815290517fc5224c4118417a068eeac7d714e6d8af6f99ec3fb611bc965185460b0e38f0819181900360200190a160019150610b19565b600101611c7c565b600354600160a060020a0316331480611df75750611df7611b21565b1515611e0257600080fd5b60058190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b6000611e483261194c565b80611e575750611e57326106f8565b80611e665750611e6632610edb565b1515611e7157600080fd5b6040517fed78a9defa7412748c9513ba9cf680f57703a46dd7e0fb0b1e94063423c73e8890600090a150600190565b600354600160a060020a0316321480611ec35750600354600160a060020a031633145b1515611ece57600080fd5b600160a060020a0381161515611ee357600080fd5b60038054600160a060020a031916600160a060020a0392909216919091179055565b60035460009081908190600160a060020a0316321480611f2f5750600354600160a060020a031633145b1515611f3a57600080fd5b5050600f546000190160005b600f548110156120a85783600160a060020a0316600b82815481101515611f6957fe5b600091825260209091200154600160a060020a031614156120a057600b805482908110611f9257fe5b60009182526020909120018054600160a060020a031916905580821461203757600b805483908110611fc057fe5b600091825260209091200154600b8054600160a060020a039092169183908110611fe657fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600b80548390811061201d57fe5b60009182526020909120018054600160a060020a03191690555b600f829055600160a060020a0384166000818152601060209081526040918290208054600160a060020a0319169055815192835290517fd41375b9d347dfe722f90a780731abd23b7855f9cf14ea7063c4cab5f9ae58e29281900390910190a16001925061103a565b600101611f46565b5060009392505050565b6040516124e1806120c3833901905600608060408190527f4f776e61626c6532303139303331353134313530304d4c00000000000000000060009081557f4564697461626c6532303139303331353134313830304d4c00000000000000006003557f42617365436f6e74656e7432303139303331353137353130304d4c0000000000600555600c556020806124e1833981016040818152915160018054600160a060020a031990811632908117909255600280548216909217909155600980548216331790819055600019600b5560068054600160a060020a038086169190941617905516825291517fc3decc188980e855666b70498ca85e8fa284d97d30483d828fa126f7303d7d199181900360200190a1506123cf806101126000396000f3006080604052600436106101a75763ffffffff60e060020a60003504166217de9881146101a957806302d05d3f146101d0578063075d4782146102015780630c6d3f931461021d5780631a735f18146102885780632310167f146102d557806327c1c21d146102ea57806332eaf21b146102ff57806336ebffca14610314578063388642841461032957806338d0f5041461034157806341c0e1b5146103fa5780634dd707881461040f5780635267db441461042457806354fd4d501461043c5780635cc4aa9b1461045157806364ade32b146104625780636d2e4b1b146104775780637ca8f618146104985780638280dd8f146104b0578063879fe48f146104c85780638da5cb5b1461055e5780638f77920114610573578063a1ff106e14610588578063b816f5131461068c578063c287e0ed146106a1578063c9e8e72d146106b6578063cbcd4461146106d7578063d810f8c8146106ec578063e02dd9c214610701578063e538530314610716578063ee56d76714610737578063f14fcbc8146107d7578063f2fde38b146107ef578063f4d9bae814610810578063f81ab0ae14610828575b005b3480156101b557600080fd5b506101be61083d565b60408051918252519081900360200190f35b3480156101dc57600080fd5b506101e5610861565b60408051600160a060020a039092168252519081900360200190f35b610209610870565b604080519115158252519081900360200190f35b34801561022957600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526102099482359460248035600160a060020a03169536959460649492019190819084018382808284375094975050933594506109d19350505050565b34801561029457600080fd5b506102a0600435610a18565b60408051600160a060020a0390951685526020850193909352600091820b90910b838301526060830152519081900360800190f35b3480156102e157600080fd5b506101e5610a4c565b3480156102f657600080fd5b506101be610a5b565b34801561030b57600080fd5b506101e5610a61565b34801561032057600080fd5b506101e5610a70565b34801561033557600080fd5b506101be600435610a7f565b34801561034d57600080fd5b506040805160206004602480358281013584810280870186019097528086526103d796843560ff1696369660449591949091019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750949750610bd19650505050505050565b604051808360000b60000b81526020018281526020019250505060405180910390f35b34801561040657600080fd5b506101a7610d5b565b34801561041b57600080fd5b506101be610e40565b34801561043057600080fd5b506101be600435610e64565b34801561044857600080fd5b506101be610ef9565b610209600435602435604435610eff565b34801561046e57600080fd5b506101be6111c0565b34801561048357600080fd5b506101a7600160a060020a03600435166111c6565b3480156104a457600080fd5b506101be600435611221565b3480156104bc57600080fd5b506101be600435611240565b3480156104d457600080fd5b506040805160206004602480358281013584810280870186019097528086526101be96843560ff1696369660449591949091019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506113b79650505050505050565b34801561056a57600080fd5b506101e5611421565b34801561057f57600080fd5b506101be611430565b60408051602060046024803582810135601f810185900485028601850190965285855261020995833560ff1695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506114369650505050505050565b34801561069857600080fd5b506101e5611927565b3480156106ad57600080fd5b506101a7611936565b3480156106c257600080fd5b506101a7600160a060020a036004351661199b565b3480156106e357600080fd5b506102096119f8565b3480156106f857600080fd5b506101be611ac9565b34801561070d57600080fd5b506101be611aed565b34801561072257600080fd5b506101a7600160a060020a0360043516611af3565b34801561074357600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610209948235946024803515159536959460649492019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750611cb79650505050505050565b3480156107e357600080fd5b506101a760043561204e565b3480156107fb57600080fd5b506101a7600160a060020a03600435166120b7565b34801561081c57600080fd5b506101be600435612129565b34801561083457600080fd5b506101be61219a565b7f5075626c6973686564000000000000000000000000000000000000000000000081565b600154600160a060020a031681565b600080600061087d6119f8565b151561088857600080fd5b600454156108c757600454600e80546001810182556000919091527fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fd01555b6108d2600f546121ac565b6000600f556108e16001611240565b50600091506000600b5413156109825750600954604080517f49102e610000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169182916349102e619160048083019260209291908290030181600087803b15801561095357600080fd5b505af1158015610967573d6000803e3d6000fd5b505050506040513d602081101561097d57600080fd5b505191505b600b54600454604080518515158152602081019390935282810191909152517f901e6f3cdc4c61620d5d424116934b9af6e31ba79cdeaa349336d93ecfe846d49181900360600190a150919050565b600854600090600160a060020a0316158015906109f85750600854600160a060020a031633145b1515610a0357600080fd5b610a0f8585858561220e565b95945050505050565b600d602052600090815260408120805460018201546002830154600390930154600160a060020a03909216939092900b9084565b600854600160a060020a031681565b600b5481565b600754600160a060020a031681565b600654600160a060020a031681565b60085460009081908190600160a060020a031615610b2f5750600854604080517f45080442000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a0390921691829163450804429160248083019260209291908290030181600087803b158015610b0057600080fd5b505af1158015610b14573d6000803e3d6000fd5b505050506040513d6020811015610b2a57600080fd5b505191505b8115610b3d57819250610bca565b831515610b6c577f5075626c697368656400000000000000000000000000000000000000000000009250610bca565b6000841215610b9d577f44726166740000000000000000000000000000000000000000000000000000009250610bca565b6000841315610bca577f447261667420696e20726576696577000000000000000000000000000000000092505b5050919050565b60085460009081908190600019908290600160a060020a031615610d2e57506008546040517f0f82c16f00000000000000000000000000000000000000000000000000000000815260ff891660048201908152606060248301908152895160648401528951600160a060020a03909416938493630f82c16f938d938d938d9360448101916084909101906020808801910280838360005b83811015610c80578181015183820152602001610c68565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610cbf578181015183820152602001610ca7565b50505050905001955050505050506040805180830381600087803b158015610ce657600080fd5b505af1158015610cfa573d6000803e3d6000fd5b505050506040513d6040811015610d1057600080fd5b508051602090910151600a54909450909250831115610d2e57606491505b8160000b6000191415610d4957600a54600095509350610d50565b8183945094505b505050935093915050565b600254600090600160a060020a0316321480610d815750600254600160a060020a031633145b1515610d8c57600080fd5b600854600160a060020a031615610e355750600854604080517f9e99bbea0000000000000000000000000000000000000000000000000000000081529051600160a060020a03909216918291639e99bbea9160048083019260209291908290030181600087803b158015610dff57600080fd5b505af1158015610e13573d6000803e3d6000fd5b505050506040513d6020811015610e2957600080fd5b505115610e3557600080fd5b610e3d612340565b50565b7f447261667400000000000000000000000000000000000000000000000000000081565b600254600090600160a060020a031632148015610e9857506000821280610e985750600082138015610e9857506000600b54125b15610ea357600b8290555b600954600160a060020a0316331415610ebc57600b8290555b600b5460408051918252517fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a29181900360200190a15050600b5490565b60055481565b6000838152600d602052604081208054829081908190600160a060020a031615801590610f4957508354600160a060020a0316331480610f495750600254600160a060020a031633145b1515610f5457600080fd5b6008548715159350600160a060020a03161561100e57600854604080517f17685953000000000000000000000000000000000000000000000000000000008152600481018b9052602481018a90529051600160a060020a03909216935083916317685953916044808201926020929091908290030181600087803b158015610fdb57600080fd5b505af1158015610fef573d6000803e3d6000fd5b505050506040513d602081101561100557600080fd5b50518015935090505b8354600160a060020a031633141561104b57821561103b576002848101805460ff1916909117905561104b565b60028401805460ff191660fe1790555b836001015484600301541015611129576002840154600090810b810b136110cc57835460408051808201909152600681527f726566756e6400000000000000000000000000000000000000000000000000006020820152600386015460018701546110c6938c93600160a060020a039091169290910361220e565b50611129565b60025460408051808201909152600e81527f72656c6561736520657363726f77000000000000000000000000000000000000602082015260038601546001870154611127938c93600160a060020a039091169290910361220e565b505b6000888152600d60209081526040808320805473ffffffffffffffffffffffffffffffffffffffff191681556001810184905560028101805460ff191690556003019290925581518a8152908101899052808201889052841515606082015290517f7f1f4b28434ce7beab4983e64a8b5bb96e195a67029fdaff925028aec57fbc6b9181900360800190a150909695505050505050565b600a5481565b600154600160a060020a031632146111dd57600080fd5b600160a060020a03811615156111f257600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600e80548290811061122f57fe5b600091825260209091200154905081565b600080600061124d6119f8565b151561125857600080fd5b600854600160a060020a031615156112dc57600254600160a060020a031632148061128d5750600254600160a060020a031633145b80156112a457508360001914806112a45750836001145b156112b1578391506112d7565b600954600160a060020a0316331480156112ce57506000600b5412155b156112d7578391505b611374565b50600854604080517f3513a805000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a03909216918291633513a8059160248083019260209291908290030181600087803b15801561134557600080fd5b505af1158015611359573d6000803e3d6000fd5b505050506040513d602081101561136f57600080fd5b505191505b600b8290556040805183815290517fda4f34b30fa0ba8a73fedb922f4d28e2a10a5d68e53cf8e942abce3ac09158a29181900360200190a15050600b5492915050565b60008060006113c7868686610bd1565b92509050600081900b156113da57600080fd5b6040805160ff881681526020810184905281517fa58326ee5bb617cb8b4f0d0f5f557c469d2d05d7a738f777037deda9c724b370929181900390910190a150949350505050565b600254600160a060020a031681565b600c5481565b600080600061144361237c565b600c80546001019055600954604080517f95a078e80000000000000000000000000000000000000000000000000000000081523260048201529051600160a060020a0390921694506000918291829187916395a078e891602480830192602092919082900301818787803b1580156114ba57600080fd5b505af11580156114ce573d6000803e3d6000fd5b505050506040513d60208110156114e457600080fd5b505115156114f157600080fd5b600254600160a060020a0316321461151d5761150e8c8a8a6113b7565b94503485111561151d57600080fd5b60408051608081018252338152346020808301918252600083850181815260608501828152600c548352600d90935294812084518154600160a060020a0391821673ffffffffffffffffffffffffffffffffffffffff199091161782559351600182015594516002860180549190920b60ff1660ff199091161790555160039093019290925560085490955016156116e557600860009054906101000a9004600160a060020a0316925082600160a060020a031663123e0e80600c548e8c8c6040518563ffffffff1660e060020a028152600401808581526020018460ff1660ff1681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561164457818101518382015260200161162c565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561168357818101518382015260200161166b565b505050509050019650505050505050602060405180830381600087803b1580156116ac57600080fd5b505af11580156116c0573d6000803e3d6000fd5b505050506040513d60208110156116d657600080fd5b5051915081156116e557600080fd5b7f089a6f1788a3c353423e1be4ba12533bdde7d908bb41abeee185af0acb3df562600c548d6004548e8e604051808681526020018560ff1660ff16815260200184600019166000191681526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611770578181015183820152602001611758565b50505050905090810190601f16801561179d5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156117d05781810151838201526020016117b8565b50505050905090810190601f1680156117fd5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15060005b885181101561188d57888181518110151561182a57fe5b6020908102909101015115611885577f515e0a48b385fce2a8e4d9f169a97c4f6ea669a752358f5e6ab37cc3c2e84c38898281518110151561186857fe5b602090810290910181015160408051918252519081900390910190a15b600101611813565b5060005b87518110156119155788818151811015156118a857fe5b602090810290910101511561190d577fb6e3239e521a6c66920ae634f8e921a37e6991d520ac44d52f8516397f41b68488828151811015156118e657fe5b602090810290910181015160408051600160a060020a039092168252519081900390910190a15b600101611891565b5060019b9a5050505050505050505050565b600954600160a060020a031681565b600254600160a060020a03163214806119595750600254600160a060020a031633145b151561196457600080fd5b60045460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600254600160a060020a03163214806119be5750600254600160a060020a031633145b15156119c957600080fd5b6007805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6002546000908190600160a060020a0316331480611a205750600954600160a060020a031633145b15611a2e5760019150611ac5565b50600954604080517f26683e140000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169182916326683e149160248083019260209291908290030181600087803b158015611a9657600080fd5b505af1158015611aaa573d6000803e3d6000fd5b505050506040513d6020811015611ac057600080fd5b505191505b5090565b7f447261667420696e20726576696577000000000000000000000000000000000081565b60045481565b60025460009081908190600160a060020a0316321480611b1d5750600254600160a060020a031633145b1515611b2857600080fd5b600854600160a060020a031615611bc857600860009054906101000a9004600160a060020a0316925082600160a060020a0316639e99bbea6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611b8f57600080fd5b505af1158015611ba3573d6000803e3d6000fd5b505050506040513d6020811015611bb957600080fd5b505191508115611bc857600080fd5b6008805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03861690811790915515611c725783925082600160a060020a0316637b1cdb3e6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611c3957600080fd5b505af1158015611c4d573d6000803e3d6000fd5b505050506040513d6020811015611c6357600080fd5b505190508015611c7257600080fd5b60085460408051600160a060020a039092168252517fa6f2e38f0cfebf27212317fced3ac40bc62e00bd33f38d69603710740c69acb79181900360200190a150505050565b600254600090819081908190600160a060020a0316331480611ce35750600754600160a060020a031633145b1515611cee57600080fd5b6000888152600d602052604090208054909350600160a060020a03161515611d1557600080fd5b600854879250600160a060020a031615611dce5750600854604080517fe870ed91000000000000000000000000000000000000000000000000000000008152600481018a905288151560248201529051600160a060020a0390921691829163e870ed919160448083019260209291908290030181600087803b158015611d9a57600080fd5b505af1158015611dae573d6000803e3d6000fd5b505050506040513d6020811015611dc457600080fd5b5051159150611ea2565b826001015483600301541015611ea257861515611e4557825460408051808201909152600f81527f616363657373206465636c696e65640000000000000000000000000000000000602082015260038501546001860154611e3f938c93600160a060020a039091169290910361220e565b50611ea2565b60025460408051808201909152600d81527f6f776e6572207061796d656e7400000000000000000000000000000000000000602082015260038501546001860154611ea0938c93600160a060020a039091169290910361220e565b505b60018215151415611fdd5760028301805460ff19166001908117909155604080518a8152602080820184905260809282018381528a519383019390935289517f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718948d9490938c938c93919291606084019160a08501919087019080838360005b83811015611f3a578181015183820152602001611f22565b50505050905090810190601f168015611f675780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611f9a578181015183820152602001611f82565b50505050905090810190601f168015611fc75780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a1612043565b60028301805460ff191660ff179055604080518981526000602082018190526080828401819052820181905260c06060830181905282015290517f475e9d68ca61f129cebee5af694af00ed0e3b3b0d4b74071fbb81d0e2b912718918190036101000190a15b509695505050505050565b600254600160a060020a03163214806120715750600254600160a060020a031633145b151561207c57600080fd5b600f8190556040805182815290517fa288d3a2aba7d5dec44e93ff46d2f1129e251695be2046de105f9a9c6feefcaa9181900360200190a150565b600254600160a060020a03163214806120da5750600254600160a060020a031633145b15156120e557600080fd5b600160a060020a03811615156120fa57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600090600160a060020a031632148061214f5750600254600160a060020a031633145b151561215a57600080fd5b600a8290556040805183815290517f4114f8ef80b6de2161db580cbefa14e1892d15d3ebe2062c9914e4a5773114a39181900360200190a15050600a5490565b60006121a7600b54610a7f565b905090565b600254600160a060020a03163314806121c857506121c86119f8565b15156121d357600080fd5b60048190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b6000848152600d602052604081206001810154600382015484011161233757604051600160a060020a0386169084156108fc029085906000818181858888f19350505050158015612263573d6000803e3d6000fd5b508281600301540181600301819055507fad58d18ea7292f887da6f15bb4f0badddaa33d169713d09cf49710acc7c3a5b986858786604051808581526020018060200184600160a060020a0316600160a060020a03168152602001838152602001828103825285818151815260200191508051906020019080838360005b838110156122f95781810151838201526020016122e1565b50505050905090810190601f1680156123265780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a15b50949350505050565b600254600160a060020a03163214806123635750600254600160a060020a031633145b151561236e57600080fd5b600254600160a060020a0316ff5b604080516080810182526000808252602082018190529181018290526060810191909152905600a165627a7a723058202fb499b0671f272703d23a44987b8c3eea18538337d5e28f0d8dc8fca16109ae0029a165627a7a723058207dbf31d67589062030b4a2b7fe318b722e198afe2c9baf9ed2c099a1cc400bd50029`

// DeployBaseLibrary deploys a new Ethereum contract, binding an instance of BaseLibrary to it.
func DeployBaseLibrary(auth *bind.TransactOpts, backend bind.ContractBackend, address_KMS common.Address, content_space common.Address) (common.Address, *types.Transaction, *BaseLibrary, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseLibraryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaseLibraryBin), backend, address_KMS, content_space)
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

// CanNodePublish is a free data retrieval call binding the contract method 0x26683e14.
//
// Solidity: function canNodePublish(address candidate) constant returns(bool)
func (_BaseLibrary *BaseLibraryCaller) CanNodePublish(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "canNodePublish", candidate)
	return *ret0, err
}

// CanNodePublish is a free data retrieval call binding the contract method 0x26683e14.
//
// Solidity: function canNodePublish(address candidate) constant returns(bool)
func (_BaseLibrary *BaseLibrarySession) CanNodePublish(candidate common.Address) (bool, error) {
	return _BaseLibrary.Contract.CanNodePublish(&_BaseLibrary.CallOpts, candidate)
}

// CanNodePublish is a free data retrieval call binding the contract method 0x26683e14.
//
// Solidity: function canNodePublish(address candidate) constant returns(bool)
func (_BaseLibrary *BaseLibraryCallerSession) CanNodePublish(candidate common.Address) (bool, error) {
	return _BaseLibrary.Contract.CanNodePublish(&_BaseLibrary.CallOpts, candidate)
}

// CanPublish is a free data retrieval call binding the contract method 0xcbcd4461.
//
// Solidity: function canPublish() constant returns(bool)
func (_BaseLibrary *BaseLibraryCaller) CanPublish(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "canPublish")
	return *ret0, err
}

// CanPublish is a free data retrieval call binding the contract method 0xcbcd4461.
//
// Solidity: function canPublish() constant returns(bool)
func (_BaseLibrary *BaseLibrarySession) CanPublish() (bool, error) {
	return _BaseLibrary.Contract.CanPublish(&_BaseLibrary.CallOpts)
}

// CanPublish is a free data retrieval call binding the contract method 0xcbcd4461.
//
// Solidity: function canPublish() constant returns(bool)
func (_BaseLibrary *BaseLibraryCallerSession) CanPublish() (bool, error) {
	return _BaseLibrary.Contract.CanPublish(&_BaseLibrary.CallOpts)
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

// ValidType is a free data retrieval call binding the contract method 0x29dedde5.
//
// Solidity: function validType(address content_type) constant returns(bool)
func (_BaseLibrary *BaseLibraryCaller) ValidType(opts *bind.CallOpts, content_type common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "validType", content_type)
	return *ret0, err
}

// ValidType is a free data retrieval call binding the contract method 0x29dedde5.
//
// Solidity: function validType(address content_type) constant returns(bool)
func (_BaseLibrary *BaseLibrarySession) ValidType(content_type common.Address) (bool, error) {
	return _BaseLibrary.Contract.ValidType(&_BaseLibrary.CallOpts, content_type)
}

// ValidType is a free data retrieval call binding the contract method 0x29dedde5.
//
// Solidity: function validType(address content_type) constant returns(bool)
func (_BaseLibrary *BaseLibraryCallerSession) ValidType(content_type common.Address) (bool, error) {
	return _BaseLibrary.Contract.ValidType(&_BaseLibrary.CallOpts, content_type)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseLibrary *BaseLibraryCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BaseLibrary.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseLibrary *BaseLibrarySession) Version() ([32]byte, error) {
	return _BaseLibrary.Contract.Version(&_BaseLibrary.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_BaseLibrary *BaseLibraryCallerSession) Version() ([32]byte, error) {
	return _BaseLibrary.Contract.Version(&_BaseLibrary.CallOpts)
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
// Solidity: function removeAccessorGroup(address group) returns(bool)
func (_BaseLibrary *BaseLibraryTransactor) RemoveAccessorGroup(opts *bind.TransactOpts, group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "removeAccessorGroup", group)
}

// RemoveAccessorGroup is a paid mutator transaction binding the contract method 0xe8de515f.
//
// Solidity: function removeAccessorGroup(address group) returns(bool)
func (_BaseLibrary *BaseLibrarySession) RemoveAccessorGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.RemoveAccessorGroup(&_BaseLibrary.TransactOpts, group)
}

// RemoveAccessorGroup is a paid mutator transaction binding the contract method 0xe8de515f.
//
// Solidity: function removeAccessorGroup(address group) returns(bool)
func (_BaseLibrary *BaseLibraryTransactorSession) RemoveAccessorGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.RemoveAccessorGroup(&_BaseLibrary.TransactOpts, group)
}

// RemoveContentType is a paid mutator transaction binding the contract method 0xfd089196.
//
// Solidity: function removeContentType(address content_type) returns(bool)
func (_BaseLibrary *BaseLibraryTransactor) RemoveContentType(opts *bind.TransactOpts, content_type common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "removeContentType", content_type)
}

// RemoveContentType is a paid mutator transaction binding the contract method 0xfd089196.
//
// Solidity: function removeContentType(address content_type) returns(bool)
func (_BaseLibrary *BaseLibrarySession) RemoveContentType(content_type common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.RemoveContentType(&_BaseLibrary.TransactOpts, content_type)
}

// RemoveContentType is a paid mutator transaction binding the contract method 0xfd089196.
//
// Solidity: function removeContentType(address content_type) returns(bool)
func (_BaseLibrary *BaseLibraryTransactorSession) RemoveContentType(content_type common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.RemoveContentType(&_BaseLibrary.TransactOpts, content_type)
}

// RemoveContributorGroup is a paid mutator transaction binding the contract method 0x386493e0.
//
// Solidity: function removeContributorGroup(address group) returns(bool)
func (_BaseLibrary *BaseLibraryTransactor) RemoveContributorGroup(opts *bind.TransactOpts, group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "removeContributorGroup", group)
}

// RemoveContributorGroup is a paid mutator transaction binding the contract method 0x386493e0.
//
// Solidity: function removeContributorGroup(address group) returns(bool)
func (_BaseLibrary *BaseLibrarySession) RemoveContributorGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.RemoveContributorGroup(&_BaseLibrary.TransactOpts, group)
}

// RemoveContributorGroup is a paid mutator transaction binding the contract method 0x386493e0.
//
// Solidity: function removeContributorGroup(address group) returns(bool)
func (_BaseLibrary *BaseLibraryTransactorSession) RemoveContributorGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.RemoveContributorGroup(&_BaseLibrary.TransactOpts, group)
}

// RemoveReviewerGroup is a paid mutator transaction binding the contract method 0x1b969895.
//
// Solidity: function removeReviewerGroup(address group) returns(bool)
func (_BaseLibrary *BaseLibraryTransactor) RemoveReviewerGroup(opts *bind.TransactOpts, group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.contract.Transact(opts, "removeReviewerGroup", group)
}

// RemoveReviewerGroup is a paid mutator transaction binding the contract method 0x1b969895.
//
// Solidity: function removeReviewerGroup(address group) returns(bool)
func (_BaseLibrary *BaseLibrarySession) RemoveReviewerGroup(group common.Address) (*types.Transaction, error) {
	return _BaseLibrary.Contract.RemoveReviewerGroup(&_BaseLibrary.TransactOpts, group)
}

// RemoveReviewerGroup is a paid mutator transaction binding the contract method 0x1b969895.
//
// Solidity: function removeReviewerGroup(address group) returns(bool)
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

// BaseLibraryContentTypeRemovedIterator is returned from FilterContentTypeRemoved and is used to iterate over the raw logs and unpacked data for ContentTypeRemoved events raised by the BaseLibrary contract.
type BaseLibraryContentTypeRemovedIterator struct {
	Event *BaseLibraryContentTypeRemoved // Event containing the contract specifics and raw log

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
func (it *BaseLibraryContentTypeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseLibraryContentTypeRemoved)
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
		it.Event = new(BaseLibraryContentTypeRemoved)
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
func (it *BaseLibraryContentTypeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseLibraryContentTypeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseLibraryContentTypeRemoved represents a ContentTypeRemoved event raised by the BaseLibrary contract.
type BaseLibraryContentTypeRemoved struct {
	ContentType common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContentTypeRemoved is a free log retrieval operation binding the contract event 0xd41375b9d347dfe722f90a780731abd23b7855f9cf14ea7063c4cab5f9ae58e2.
//
// Solidity: event ContentTypeRemoved(address contentType)
func (_BaseLibrary *BaseLibraryFilterer) FilterContentTypeRemoved(opts *bind.FilterOpts) (*BaseLibraryContentTypeRemovedIterator, error) {

	logs, sub, err := _BaseLibrary.contract.FilterLogs(opts, "ContentTypeRemoved")
	if err != nil {
		return nil, err
	}
	return &BaseLibraryContentTypeRemovedIterator{contract: _BaseLibrary.contract, event: "ContentTypeRemoved", logs: logs, sub: sub}, nil
}

// WatchContentTypeRemoved is a free log subscription operation binding the contract event 0xd41375b9d347dfe722f90a780731abd23b7855f9cf14ea7063c4cab5f9ae58e2.
//
// Solidity: event ContentTypeRemoved(address contentType)
func (_BaseLibrary *BaseLibraryFilterer) WatchContentTypeRemoved(opts *bind.WatchOpts, sink chan<- *BaseLibraryContentTypeRemoved) (event.Subscription, error) {

	logs, sub, err := _BaseLibrary.contract.WatchLogs(opts, "ContentTypeRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseLibraryContentTypeRemoved)
				if err := _BaseLibrary.contract.UnpackLog(event, "ContentTypeRemoved", log); err != nil {
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
const ContentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"address[]\"}],\"name\":\"runAccessInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"address[]\"}],\"name\":\"runAccess\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"runFinalize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proposed_status_code\",\"type\":\"int256\"}],\"name\":\"runStatusChange\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"name\":\"runDescribeStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"runCreate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"runKill\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"name\":\"runGrant\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"}],\"name\":\"Log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"LogBool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"a\",\"type\":\"address\"}],\"name\":\"LogAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"u\",\"type\":\"uint256\"}],\"name\":\"LogUint256\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"u\",\"type\":\"int256\"}],\"name\":\"LogInt256\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"b\",\"type\":\"bytes32\"}],\"name\":\"LogBytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"RunCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"RunKill\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"proposedStatusCode\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"returnStatusCode\",\"type\":\"int256\"}],\"name\":\"RunStatusChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"calculateAccessCharge\",\"type\":\"int256\"}],\"name\":\"RunAccessCharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"RunAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"RunFinalize\",\"type\":\"event\"}]"

// ContentBin is the compiled bytecode used for deploying new contracts.
const ContentBin = `0x60806040527f4f776e61626c6532303139303331353134313530304d4c0000000000000000006000557f436f6e74656e7432303139303331353137313530304d4c00000000000000000060035560018054600160a060020a0319908116329081179092556002805490911690911790556104aa8061007e6000396000f3006080604052600436106100cf5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100d15780630f82c16f14610102578063123e0e80146101bb578063176859531461025a5780633513a8051461026857806341c0e1b514610273578063450804421461028857806354fd4d50146102a05780636d2e4b1b146102b55780637b1cdb3e146102d65780638da5cb5b146102de5780639e99bbea146102d6578063e870ed91146102f3578063f2fde38b14610303575b005b3480156100dd57600080fd5b506100e6610324565b60408051600160a060020a039092168252519081900360200190f35b34801561010e57600080fd5b5060408051602060046024803582810135848102808701860190975280865261019896843560ff1696369660449591949091019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506103339650505050505050565b604051808360000b60000b81526020018281526020019250505060405180910390f35b6040805160206004604435818101358381028086018501909652808552610248958335956024803560ff1696369695606495939492019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506103409650505050505050565b60408051918252519081900360200190f35b61024860043560243561034a565b610248600435610352565b34801561027f57600080fd5b506100cf610355565b34801561029457600080fd5b50610248600435610391565b3480156102ac57600080fd5b50610248610397565b3480156102c157600080fd5b506100cf600160a060020a036004351661039d565b6102486103f8565b3480156102ea57600080fd5b506100e66103fd565b610248600435602435151561034a565b34801561030f57600080fd5b506100cf600160a060020a036004351661040c565b600154600160a060020a031681565b6000196000935093915050565b6000949350505050565b600092915050565b90565b600254600160a060020a03163214806103785750600254600160a060020a031633145b151561038357600080fd5b600254600160a060020a0316ff5b50600090565b60035481565b600154600160a060020a031632146103b457600080fd5b600160a060020a03811615156103c957600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600090565b600254600160a060020a031681565b600254600160a060020a031632148061042f5750600254600160a060020a031633145b151561043a57600080fd5b600160a060020a038116151561044f57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582023d99c277c7e28ddc38e9570b5809816080187824411fd857150a1dc485528900029`

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

// RunAccessInfo is a free data retrieval call binding the contract method 0x0f82c16f.
//
// Solidity: function runAccessInfo(uint8 , bytes32[] , address[] ) constant returns(int8, uint256)
func (_Content *ContentCaller) RunAccessInfo(opts *bind.CallOpts, arg0 uint8, arg1 [][32]byte, arg2 []common.Address) (int8, *big.Int, error) {
	var (
		ret0 = new(int8)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Content.contract.Call(opts, out, "runAccessInfo", arg0, arg1, arg2)
	return *ret0, *ret1, err
}

// RunAccessInfo is a free data retrieval call binding the contract method 0x0f82c16f.
//
// Solidity: function runAccessInfo(uint8 , bytes32[] , address[] ) constant returns(int8, uint256)
func (_Content *ContentSession) RunAccessInfo(arg0 uint8, arg1 [][32]byte, arg2 []common.Address) (int8, *big.Int, error) {
	return _Content.Contract.RunAccessInfo(&_Content.CallOpts, arg0, arg1, arg2)
}

// RunAccessInfo is a free data retrieval call binding the contract method 0x0f82c16f.
//
// Solidity: function runAccessInfo(uint8 , bytes32[] , address[] ) constant returns(int8, uint256)
func (_Content *ContentCallerSession) RunAccessInfo(arg0 uint8, arg1 [][32]byte, arg2 []common.Address) (int8, *big.Int, error) {
	return _Content.Contract.RunAccessInfo(&_Content.CallOpts, arg0, arg1, arg2)
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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Content *ContentCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Content.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Content *ContentSession) Version() ([32]byte, error) {
	return _Content.Contract.Version(&_Content.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Content *ContentCallerSession) Version() ([32]byte, error) {
	return _Content.Contract.Version(&_Content.CallOpts)
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

// RunFinalize is a paid mutator transaction binding the contract method 0x17685953.
//
// Solidity: function runFinalize(uint256 , uint256 ) returns(uint256)
func (_Content *ContentTransactor) RunFinalize(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "runFinalize", arg0, arg1)
}

// RunFinalize is a paid mutator transaction binding the contract method 0x17685953.
//
// Solidity: function runFinalize(uint256 , uint256 ) returns(uint256)
func (_Content *ContentSession) RunFinalize(arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _Content.Contract.RunFinalize(&_Content.TransactOpts, arg0, arg1)
}

// RunFinalize is a paid mutator transaction binding the contract method 0x17685953.
//
// Solidity: function runFinalize(uint256 , uint256 ) returns(uint256)
func (_Content *ContentTransactorSession) RunFinalize(arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _Content.Contract.RunFinalize(&_Content.TransactOpts, arg0, arg1)
}

// RunGrant is a paid mutator transaction binding the contract method 0xe870ed91.
//
// Solidity: function runGrant(uint256 , bool ) returns(uint256)
func (_Content *ContentTransactor) RunGrant(opts *bind.TransactOpts, arg0 *big.Int, arg1 bool) (*types.Transaction, error) {
	return _Content.contract.Transact(opts, "runGrant", arg0, arg1)
}

// RunGrant is a paid mutator transaction binding the contract method 0xe870ed91.
//
// Solidity: function runGrant(uint256 , bool ) returns(uint256)
func (_Content *ContentSession) RunGrant(arg0 *big.Int, arg1 bool) (*types.Transaction, error) {
	return _Content.Contract.RunGrant(&_Content.TransactOpts, arg0, arg1)
}

// RunGrant is a paid mutator transaction binding the contract method 0xe870ed91.
//
// Solidity: function runGrant(uint256 , bool ) returns(uint256)
func (_Content *ContentTransactorSession) RunGrant(arg0 *big.Int, arg1 bool) (*types.Transaction, error) {
	return _Content.Contract.RunGrant(&_Content.TransactOpts, arg0, arg1)
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

// ContentLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the Content contract.
type ContentLogIterator struct {
	Event *ContentLog // Event containing the contract specifics and raw log

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
func (it *ContentLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentLog)
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
		it.Event = new(ContentLog)
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
func (it *ContentLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentLog represents a Log event raised by the Content contract.
type ContentLog struct {
	Label string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0xcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab.
//
// Solidity: event Log(string label)
func (_Content *ContentFilterer) FilterLog(opts *bind.FilterOpts) (*ContentLogIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return &ContentLogIterator{contract: _Content.contract, event: "Log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0xcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab.
//
// Solidity: event Log(string label)
func (_Content *ContentFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *ContentLog) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentLog)
				if err := _Content.contract.UnpackLog(event, "Log", log); err != nil {
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

// ContentLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the Content contract.
type ContentLogAddressIterator struct {
	Event *ContentLogAddress // Event containing the contract specifics and raw log

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
func (it *ContentLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentLogAddress)
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
		it.Event = new(ContentLogAddress)
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
func (it *ContentLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentLogAddress represents a LogAddress event raised by the Content contract.
type ContentLogAddress struct {
	Label string
	A     common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x62ddffe5b5108385f7a590f100e1ee414ad9551a31f089e64e82998440785e1e.
//
// Solidity: event LogAddress(string label, address a)
func (_Content *ContentFilterer) FilterLogAddress(opts *bind.FilterOpts) (*ContentLogAddressIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "LogAddress")
	if err != nil {
		return nil, err
	}
	return &ContentLogAddressIterator{contract: _Content.contract, event: "LogAddress", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x62ddffe5b5108385f7a590f100e1ee414ad9551a31f089e64e82998440785e1e.
//
// Solidity: event LogAddress(string label, address a)
func (_Content *ContentFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *ContentLogAddress) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "LogAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentLogAddress)
				if err := _Content.contract.UnpackLog(event, "LogAddress", log); err != nil {
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

// ContentLogBoolIterator is returned from FilterLogBool and is used to iterate over the raw logs and unpacked data for LogBool events raised by the Content contract.
type ContentLogBoolIterator struct {
	Event *ContentLogBool // Event containing the contract specifics and raw log

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
func (it *ContentLogBoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentLogBool)
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
		it.Event = new(ContentLogBool)
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
func (it *ContentLogBoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentLogBoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentLogBool represents a LogBool event raised by the Content contract.
type ContentLogBool struct {
	Label string
	B     bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogBool is a free log retrieval operation binding the contract event 0x4c34c2f9a78632f29fa59aaed5514cb742fd9fbcfd7ccc2c03c85f2bbc621c47.
//
// Solidity: event LogBool(string label, bool b)
func (_Content *ContentFilterer) FilterLogBool(opts *bind.FilterOpts) (*ContentLogBoolIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "LogBool")
	if err != nil {
		return nil, err
	}
	return &ContentLogBoolIterator{contract: _Content.contract, event: "LogBool", logs: logs, sub: sub}, nil
}

// WatchLogBool is a free log subscription operation binding the contract event 0x4c34c2f9a78632f29fa59aaed5514cb742fd9fbcfd7ccc2c03c85f2bbc621c47.
//
// Solidity: event LogBool(string label, bool b)
func (_Content *ContentFilterer) WatchLogBool(opts *bind.WatchOpts, sink chan<- *ContentLogBool) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "LogBool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentLogBool)
				if err := _Content.contract.UnpackLog(event, "LogBool", log); err != nil {
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

// ContentLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the Content contract.
type ContentLogBytes32Iterator struct {
	Event *ContentLogBytes32 // Event containing the contract specifics and raw log

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
func (it *ContentLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentLogBytes32)
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
		it.Event = new(ContentLogBytes32)
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
func (it *ContentLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentLogBytes32 represents a LogBytes32 event raised by the Content contract.
type ContentLogBytes32 struct {
	Label string
	B     [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0x02d93529bba9d141e5e06733c52c7e6fbcb1149586adb5c24064b522ab26f1d7.
//
// Solidity: event LogBytes32(string label, bytes32 b)
func (_Content *ContentFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*ContentLogBytes32Iterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "LogBytes32")
	if err != nil {
		return nil, err
	}
	return &ContentLogBytes32Iterator{contract: _Content.contract, event: "LogBytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0x02d93529bba9d141e5e06733c52c7e6fbcb1149586adb5c24064b522ab26f1d7.
//
// Solidity: event LogBytes32(string label, bytes32 b)
func (_Content *ContentFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *ContentLogBytes32) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "LogBytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentLogBytes32)
				if err := _Content.contract.UnpackLog(event, "LogBytes32", log); err != nil {
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

// ContentLogInt256Iterator is returned from FilterLogInt256 and is used to iterate over the raw logs and unpacked data for LogInt256 events raised by the Content contract.
type ContentLogInt256Iterator struct {
	Event *ContentLogInt256 // Event containing the contract specifics and raw log

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
func (it *ContentLogInt256Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentLogInt256)
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
		it.Event = new(ContentLogInt256)
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
func (it *ContentLogInt256Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentLogInt256Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentLogInt256 represents a LogInt256 event raised by the Content contract.
type ContentLogInt256 struct {
	Label string
	U     *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogInt256 is a free log retrieval operation binding the contract event 0x3d9b341774178bb033613e3a7a1cadb2244b3bcbb1372905d2ba24dca38aeb22.
//
// Solidity: event LogInt256(string label, int256 u)
func (_Content *ContentFilterer) FilterLogInt256(opts *bind.FilterOpts) (*ContentLogInt256Iterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "LogInt256")
	if err != nil {
		return nil, err
	}
	return &ContentLogInt256Iterator{contract: _Content.contract, event: "LogInt256", logs: logs, sub: sub}, nil
}

// WatchLogInt256 is a free log subscription operation binding the contract event 0x3d9b341774178bb033613e3a7a1cadb2244b3bcbb1372905d2ba24dca38aeb22.
//
// Solidity: event LogInt256(string label, int256 u)
func (_Content *ContentFilterer) WatchLogInt256(opts *bind.WatchOpts, sink chan<- *ContentLogInt256) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "LogInt256")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentLogInt256)
				if err := _Content.contract.UnpackLog(event, "LogInt256", log); err != nil {
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

// ContentLogPaymentIterator is returned from FilterLogPayment and is used to iterate over the raw logs and unpacked data for LogPayment events raised by the Content contract.
type ContentLogPaymentIterator struct {
	Event *ContentLogPayment // Event containing the contract specifics and raw log

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
func (it *ContentLogPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentLogPayment)
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
		it.Event = new(ContentLogPayment)
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
func (it *ContentLogPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentLogPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentLogPayment represents a LogPayment event raised by the Content contract.
type ContentLogPayment struct {
	Label  string
	Payee  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogPayment is a free log retrieval operation binding the contract event 0x6a0f12fe24f7c34df8acc096f076bac6e3f1ff225f4fccc7b0f389657bde895a.
//
// Solidity: event LogPayment(string label, address payee, uint256 amount)
func (_Content *ContentFilterer) FilterLogPayment(opts *bind.FilterOpts) (*ContentLogPaymentIterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "LogPayment")
	if err != nil {
		return nil, err
	}
	return &ContentLogPaymentIterator{contract: _Content.contract, event: "LogPayment", logs: logs, sub: sub}, nil
}

// WatchLogPayment is a free log subscription operation binding the contract event 0x6a0f12fe24f7c34df8acc096f076bac6e3f1ff225f4fccc7b0f389657bde895a.
//
// Solidity: event LogPayment(string label, address payee, uint256 amount)
func (_Content *ContentFilterer) WatchLogPayment(opts *bind.WatchOpts, sink chan<- *ContentLogPayment) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "LogPayment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentLogPayment)
				if err := _Content.contract.UnpackLog(event, "LogPayment", log); err != nil {
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

// ContentLogUint256Iterator is returned from FilterLogUint256 and is used to iterate over the raw logs and unpacked data for LogUint256 events raised by the Content contract.
type ContentLogUint256Iterator struct {
	Event *ContentLogUint256 // Event containing the contract specifics and raw log

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
func (it *ContentLogUint256Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentLogUint256)
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
		it.Event = new(ContentLogUint256)
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
func (it *ContentLogUint256Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentLogUint256Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentLogUint256 represents a LogUint256 event raised by the Content contract.
type ContentLogUint256 struct {
	Label string
	U     *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogUint256 is a free log retrieval operation binding the contract event 0x31c369d7029afba34b21369bcf9a6ac132fb2621c34558b914859b768d05232d.
//
// Solidity: event LogUint256(string label, uint256 u)
func (_Content *ContentFilterer) FilterLogUint256(opts *bind.FilterOpts) (*ContentLogUint256Iterator, error) {

	logs, sub, err := _Content.contract.FilterLogs(opts, "LogUint256")
	if err != nil {
		return nil, err
	}
	return &ContentLogUint256Iterator{contract: _Content.contract, event: "LogUint256", logs: logs, sub: sub}, nil
}

// WatchLogUint256 is a free log subscription operation binding the contract event 0x31c369d7029afba34b21369bcf9a6ac132fb2621c34558b914859b768d05232d.
//
// Solidity: event LogUint256(string label, uint256 u)
func (_Content *ContentFilterer) WatchLogUint256(opts *bind.WatchOpts, sink chan<- *ContentLogUint256) (event.Subscription, error) {

	logs, sub, err := _Content.contract.WatchLogs(opts, "LogUint256")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentLogUint256)
				if err := _Content.contract.UnpackLog(event, "LogUint256", log); err != nil {
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
const EditableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canPublish\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"objectHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"object_hash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"Commit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"objectHash\",\"type\":\"bytes32\"}],\"name\":\"UpdateRequest\",\"type\":\"event\"}]"

// EditableBin is the compiled bytecode used for deploying new contracts.
const EditableBin = `0x60806040527f4f776e61626c6532303139303331353134313530304d4c0000000000000000006000557f4564697461626c6532303139303331353134313830304d4c000000000000000060035560018054600160a060020a0319908116329081179092556002805490911690911790556103ff8061007e6000396000f3006080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100a557806341c0e1b5146100d657806354fd4d50146100eb5780636d2e4b1b146101125780638da5cb5b14610133578063c287e0ed14610148578063cbcd44611461015d578063e02dd9c214610186578063f14fcbc81461019b578063f2fde38b146101b3575b005b3480156100b157600080fd5b506100ba6101d4565b60408051600160a060020a039092168252519081900360200190f35b3480156100e257600080fd5b506100a36101e3565b3480156100f757600080fd5b5061010061021f565b60408051918252519081900360200190f35b34801561011e57600080fd5b506100a3600160a060020a0360043516610225565b34801561013f57600080fd5b506100ba610280565b34801561015457600080fd5b506100a361028f565b34801561016957600080fd5b506101726102f4565b604080519115158252519081900360200190f35b34801561019257600080fd5b506101006102f9565b3480156101a757600080fd5b506100a36004356102ff565b3480156101bf57600080fd5b506100a3600160a060020a0360043516610361565b600154600160a060020a031681565b600254600160a060020a03163214806102065750600254600160a060020a031633145b151561021157600080fd5b600254600160a060020a0316ff5b60035481565b600154600160a060020a0316321461023c57600080fd5b600160a060020a038116151561025157600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a031681565b600254600160a060020a03163214806102b25750600254600160a060020a031633145b15156102bd57600080fd5b60045460408051918252517f1b6452b35fd3ee7f1fad8558a9d3e79233f94d15fe657df3871f8efc2f97ef199181900360200190a1565b600090565b60045481565b600254600160a060020a031633148061031b575061031b6102f4565b151561032657600080fd5b60048190556040805182815290517f9e8a51bb6b34b9d5d18c14fd753ee3bf44e2256512665a4577281ffcc91943ff9181900360200190a150565b600254600160a060020a03163214806103845750600254600160a060020a031633145b151561038f57600080fd5b600160a060020a03811615156103a457600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582020004a11734af66820d40ffd272686feb6e417c67f0074e05d967acf7f56e3210029`

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

// CanPublish is a free data retrieval call binding the contract method 0xcbcd4461.
//
// Solidity: function canPublish() constant returns(bool)
func (_Editable *EditableCaller) CanPublish(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Editable.contract.Call(opts, out, "canPublish")
	return *ret0, err
}

// CanPublish is a free data retrieval call binding the contract method 0xcbcd4461.
//
// Solidity: function canPublish() constant returns(bool)
func (_Editable *EditableSession) CanPublish() (bool, error) {
	return _Editable.Contract.CanPublish(&_Editable.CallOpts)
}

// CanPublish is a free data retrieval call binding the contract method 0xcbcd4461.
//
// Solidity: function canPublish() constant returns(bool)
func (_Editable *EditableCallerSession) CanPublish() (bool, error) {
	return _Editable.Contract.CanPublish(&_Editable.CallOpts)
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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Editable *EditableCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Editable.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Editable *EditableSession) Version() ([32]byte, error) {
	return _Editable.Contract.Version(&_Editable.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Editable *EditableCallerSession) Version() ([32]byte, error) {
	return _Editable.Contract.Version(&_Editable.CallOpts)
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
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x60806040527f4f776e61626c6532303139303331353134313530304d4c00000000000000000060005560018054600160a060020a0319908116329081179092556002805490911690911790556102968061005a6000396000f3006080604052600436106100775763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f811461007957806341c0e1b5146100aa57806354fd4d50146100bf5780636d2e4b1b146100e65780638da5cb5b14610107578063f2fde38b1461011c575b005b34801561008557600080fd5b5061008e61013d565b60408051600160a060020a039092168252519081900360200190f35b3480156100b657600080fd5b5061007761014c565b3480156100cb57600080fd5b506100d4610188565b60408051918252519081900360200190f35b3480156100f257600080fd5b50610077600160a060020a036004351661018e565b34801561011357600080fd5b5061008e6101e9565b34801561012857600080fd5b50610077600160a060020a03600435166101f8565b600154600160a060020a031681565b600254600160a060020a031632148061016f5750600254600160a060020a031633145b151561017a57600080fd5b600254600160a060020a0316ff5b60005481565b600154600160a060020a031632146101a557600080fd5b600160a060020a03811615156101ba57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a031681565b600254600160a060020a031632148061021b5750600254600160a060020a031633145b151561022657600080fd5b600160a060020a038116151561023b57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820216d6fef52e5de55f7c4a9af24e2dad13bb90c6437f243f6a238a6e629d416100029`

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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Ownable *OwnableCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Ownable *OwnableSession) Version() ([32]byte, error) {
	return _Ownable.Contract.Version(&_Ownable.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Ownable *OwnableCallerSession) Version() ([32]byte, error) {
	return _Ownable.Contract.Version(&_Ownable.CallOpts)
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
