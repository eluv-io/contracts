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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820f158e127b3c082c5a724385bfbe08571f44364833b8d3d59033cc567878899fb0029`

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

// SimplePaymentChannelABI is the input ABI used to generate the binding from.
const SimplePaymentChannelABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_expiration\",\"type\":\"uint256\"}],\"name\":\"extendExpiration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimTimeout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"checkVRS\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"expiration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"recipient\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sender\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"closeChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"checkPackedAndHashed\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"splitSignature\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"checkEcrecover\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"expiration\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"ChannelOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"senderBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"recipientBalance\",\"type\":\"uint256\"}],\"name\":\"ChannelClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"checkAddr\",\"type\":\"address\"}],\"name\":\"CheckEcRecover\",\"type\":\"event\"}]"

// SimplePaymentChannelBin is the compiled bytecode used for deploying new contracts.
const SimplePaymentChannelBin = `0x60806040818152806108818339810160405280516020909101516000341161002657600080fd5b33600160a060020a038316141561003c57600080fd5b60008054600160a060020a031990811633179182905560018054600160a060020a03868116919093161790819055428401600281905560408051948416855291909216602084015282810191909152346060830152517f9957eac353f4e194b7faa928965f2dfdedd22f864e067b5b78f98d50157c2df59181900360800190a150506107b4806100cd6000396000f3006080604052600436106100ae5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630a114a7081146100b35780630e1da6c3146100cd5780634394c120146100e25780634665096d1461017857806366d003ac1461019f57806367e404ce146101d057806375be62fc146101e55780637af9999d1461020c578063a7bb580314610227578063ac9fc588146102a2578063bd190a2814610300575b600080fd5b3480156100bf57600080fd5b506100cb60043561033b565b005b3480156100d957600080fd5b506100cb610365565b3480156100ee57600080fd5b5061010360ff60043516602435604435610382565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561013d578181015183820152602001610125565b50505050905090810190601f16801561016a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561018457600080fd5b5061018d6103d8565b60408051918252519081900360200190f35b3480156101ab57600080fd5b506101b46103de565b60408051600160a060020a039092168252519081900360200190f35b3480156101dc57600080fd5b506101b46103ed565b3480156101f157600080fd5b506100cb60043560243560ff604435166064356084356103fc565b34801561021857600080fd5b5061018d6004356024356104e0565b34801561023357600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526102809436949293602493928401919081908401838280828437509497506105819650505050505050565b6040805160ff9094168452602084019290925282820152519081900360600190f35b3480156102ae57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101b49583359536956044949193909101919081908401838280828437509497506105ba9650505050505050565b34801561030c57600080fd5b5061032760043560243560ff60443516606435608435610644565b604080519115158252519081900360200190f35b600054600160a060020a0316331461035257600080fd5b600254811161036057600080fd5b600255565b60025442101561037457600080fd5b600054600160a060020a0316ff5b6040805160ff949094167f01000000000000000000000000000000000000000000000000000000000000000260208501526021840192909252604180840191909152815180840390910181526061909201905290565b60025481565b600154600160a060020a031681565b600054600160a060020a031681565b6001546000908190600160a060020a0316331461041857600080fd5b6104258787878787610644565b151561043057600080fd5b5085905060003031821115610448573031915061045c565b61045930318363ffffffff61077116565b90505b600154604051600160a060020a039091169083156108fc029084906000818181858888f19350505050158015610496573d6000803e3d6000fd5b50604080518281526020810184905281517f1273b19bd37db7541ec10aebcb7899427619b2685565bd3c7b184c2732f8118e929181900390910190a1600054600160a060020a0316ff5b604080516c010000000000000000000000003002602080830191909152603482018590526054808301859052835180840390910181526074909201928390528151600093918291908401908083835b6020831061054e5780518252601f19909201916020918201910161052f565b5181516020939093036101000a600019018019909116921691909117905260405192018290039091209695505050505050565b6000806000806000808651604114151561059a57600080fd5b505050506020830151604084015160609094015160001a94909392509050565b60008060008060006105cb86610581565b60408051600080825260208083018085528e905260ff8716838501526060830186905260808301859052925195995093975091955060019360a0808401949293601f19830193908390039091019190865af115801561062e573d6000803e3d6000fd5b5050604051601f19015198975050505050505050565b60008060003088886040516020018084600160a060020a0316600160a060020a03166c0100000000000000000000000002815260140183815260200182815260200193505050506040516020818303038152906040526040518082805190602001908083835b602083106106c95780518252601f1990920191602091820191016106aa565b51815160209384036101000a600019018019909216911617905260408051929094018290038220600080845283830180875282905260ff8e1684870152606084018d9052608084018c905294519098506001965060a080840196509194601f19820194509281900390910191865af1158015610749573d6000803e3d6000fd5b5050604051601f190151600054600160a060020a039081169116149998505050505050505050565b6000808383111561078157600080fd5b50509003905600a165627a7a72305820f358b6660036ecc1f1d822d5d75082f0593de27c409ec6e96d3981cab1b6e00d0029`

// DeploySimplePaymentChannel deploys a new Ethereum contract, binding an instance of SimplePaymentChannel to it.
func DeploySimplePaymentChannel(auth *bind.TransactOpts, backend bind.ContractBackend, _recipient common.Address, _duration *big.Int) (common.Address, *types.Transaction, *SimplePaymentChannel, error) {
	parsed, err := abi.JSON(strings.NewReader(SimplePaymentChannelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimplePaymentChannelBin), backend, _recipient, _duration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimplePaymentChannel{SimplePaymentChannelCaller: SimplePaymentChannelCaller{contract: contract}, SimplePaymentChannelTransactor: SimplePaymentChannelTransactor{contract: contract}, SimplePaymentChannelFilterer: SimplePaymentChannelFilterer{contract: contract}}, nil
}

// SimplePaymentChannel is an auto generated Go binding around an Ethereum contract.
type SimplePaymentChannel struct {
	SimplePaymentChannelCaller     // Read-only binding to the contract
	SimplePaymentChannelTransactor // Write-only binding to the contract
	SimplePaymentChannelFilterer   // Log filterer for contract events
}

// SimplePaymentChannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimplePaymentChannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplePaymentChannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimplePaymentChannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplePaymentChannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimplePaymentChannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplePaymentChannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimplePaymentChannelSession struct {
	Contract     *SimplePaymentChannel // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SimplePaymentChannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimplePaymentChannelCallerSession struct {
	Contract *SimplePaymentChannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SimplePaymentChannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimplePaymentChannelTransactorSession struct {
	Contract     *SimplePaymentChannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SimplePaymentChannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimplePaymentChannelRaw struct {
	Contract *SimplePaymentChannel // Generic contract binding to access the raw methods on
}

// SimplePaymentChannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimplePaymentChannelCallerRaw struct {
	Contract *SimplePaymentChannelCaller // Generic read-only contract binding to access the raw methods on
}

// SimplePaymentChannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimplePaymentChannelTransactorRaw struct {
	Contract *SimplePaymentChannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimplePaymentChannel creates a new instance of SimplePaymentChannel, bound to a specific deployed contract.
func NewSimplePaymentChannel(address common.Address, backend bind.ContractBackend) (*SimplePaymentChannel, error) {
	contract, err := bindSimplePaymentChannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimplePaymentChannel{SimplePaymentChannelCaller: SimplePaymentChannelCaller{contract: contract}, SimplePaymentChannelTransactor: SimplePaymentChannelTransactor{contract: contract}, SimplePaymentChannelFilterer: SimplePaymentChannelFilterer{contract: contract}}, nil
}

// NewSimplePaymentChannelCaller creates a new read-only instance of SimplePaymentChannel, bound to a specific deployed contract.
func NewSimplePaymentChannelCaller(address common.Address, caller bind.ContractCaller) (*SimplePaymentChannelCaller, error) {
	contract, err := bindSimplePaymentChannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimplePaymentChannelCaller{contract: contract}, nil
}

// NewSimplePaymentChannelTransactor creates a new write-only instance of SimplePaymentChannel, bound to a specific deployed contract.
func NewSimplePaymentChannelTransactor(address common.Address, transactor bind.ContractTransactor) (*SimplePaymentChannelTransactor, error) {
	contract, err := bindSimplePaymentChannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimplePaymentChannelTransactor{contract: contract}, nil
}

// NewSimplePaymentChannelFilterer creates a new log filterer instance of SimplePaymentChannel, bound to a specific deployed contract.
func NewSimplePaymentChannelFilterer(address common.Address, filterer bind.ContractFilterer) (*SimplePaymentChannelFilterer, error) {
	contract, err := bindSimplePaymentChannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimplePaymentChannelFilterer{contract: contract}, nil
}

// bindSimplePaymentChannel binds a generic wrapper to an already deployed contract.
func bindSimplePaymentChannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimplePaymentChannelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimplePaymentChannel *SimplePaymentChannelRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimplePaymentChannel.Contract.SimplePaymentChannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimplePaymentChannel *SimplePaymentChannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimplePaymentChannel.Contract.SimplePaymentChannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimplePaymentChannel *SimplePaymentChannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimplePaymentChannel.Contract.SimplePaymentChannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimplePaymentChannel *SimplePaymentChannelCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimplePaymentChannel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimplePaymentChannel *SimplePaymentChannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimplePaymentChannel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimplePaymentChannel *SimplePaymentChannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimplePaymentChannel.Contract.contract.Transact(opts, method, params...)
}

// CheckEcrecover is a free data retrieval call binding the contract method 0xac9fc588.
//
// Solidity: function checkEcrecover(bytes32 _hash, bytes _sig) constant returns(address)
func (_SimplePaymentChannel *SimplePaymentChannelCaller) CheckEcrecover(opts *bind.CallOpts, _hash [32]byte, _sig []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SimplePaymentChannel.contract.Call(opts, out, "checkEcrecover", _hash, _sig)
	return *ret0, err
}

// CheckEcrecover is a free data retrieval call binding the contract method 0xac9fc588.
//
// Solidity: function checkEcrecover(bytes32 _hash, bytes _sig) constant returns(address)
func (_SimplePaymentChannel *SimplePaymentChannelSession) CheckEcrecover(_hash [32]byte, _sig []byte) (common.Address, error) {
	return _SimplePaymentChannel.Contract.CheckEcrecover(&_SimplePaymentChannel.CallOpts, _hash, _sig)
}

// CheckEcrecover is a free data retrieval call binding the contract method 0xac9fc588.
//
// Solidity: function checkEcrecover(bytes32 _hash, bytes _sig) constant returns(address)
func (_SimplePaymentChannel *SimplePaymentChannelCallerSession) CheckEcrecover(_hash [32]byte, _sig []byte) (common.Address, error) {
	return _SimplePaymentChannel.Contract.CheckEcrecover(&_SimplePaymentChannel.CallOpts, _hash, _sig)
}

// CheckPackedAndHashed is a free data retrieval call binding the contract method 0x7af9999d.
//
// Solidity: function checkPackedAndHashed(uint256 _balance, uint256 _nonce) constant returns(bytes32)
func (_SimplePaymentChannel *SimplePaymentChannelCaller) CheckPackedAndHashed(opts *bind.CallOpts, _balance *big.Int, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SimplePaymentChannel.contract.Call(opts, out, "checkPackedAndHashed", _balance, _nonce)
	return *ret0, err
}

// CheckPackedAndHashed is a free data retrieval call binding the contract method 0x7af9999d.
//
// Solidity: function checkPackedAndHashed(uint256 _balance, uint256 _nonce) constant returns(bytes32)
func (_SimplePaymentChannel *SimplePaymentChannelSession) CheckPackedAndHashed(_balance *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _SimplePaymentChannel.Contract.CheckPackedAndHashed(&_SimplePaymentChannel.CallOpts, _balance, _nonce)
}

// CheckPackedAndHashed is a free data retrieval call binding the contract method 0x7af9999d.
//
// Solidity: function checkPackedAndHashed(uint256 _balance, uint256 _nonce) constant returns(bytes32)
func (_SimplePaymentChannel *SimplePaymentChannelCallerSession) CheckPackedAndHashed(_balance *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _SimplePaymentChannel.Contract.CheckPackedAndHashed(&_SimplePaymentChannel.CallOpts, _balance, _nonce)
}

// CheckVRS is a free data retrieval call binding the contract method 0x4394c120.
//
// Solidity: function checkVRS(uint8 _v, bytes32 _r, bytes32 _s) constant returns(bytes)
func (_SimplePaymentChannel *SimplePaymentChannelCaller) CheckVRS(opts *bind.CallOpts, _v uint8, _r [32]byte, _s [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _SimplePaymentChannel.contract.Call(opts, out, "checkVRS", _v, _r, _s)
	return *ret0, err
}

// CheckVRS is a free data retrieval call binding the contract method 0x4394c120.
//
// Solidity: function checkVRS(uint8 _v, bytes32 _r, bytes32 _s) constant returns(bytes)
func (_SimplePaymentChannel *SimplePaymentChannelSession) CheckVRS(_v uint8, _r [32]byte, _s [32]byte) ([]byte, error) {
	return _SimplePaymentChannel.Contract.CheckVRS(&_SimplePaymentChannel.CallOpts, _v, _r, _s)
}

// CheckVRS is a free data retrieval call binding the contract method 0x4394c120.
//
// Solidity: function checkVRS(uint8 _v, bytes32 _r, bytes32 _s) constant returns(bytes)
func (_SimplePaymentChannel *SimplePaymentChannelCallerSession) CheckVRS(_v uint8, _r [32]byte, _s [32]byte) ([]byte, error) {
	return _SimplePaymentChannel.Contract.CheckVRS(&_SimplePaymentChannel.CallOpts, _v, _r, _s)
}

// Expiration is a free data retrieval call binding the contract method 0x4665096d.
//
// Solidity: function expiration() constant returns(uint256)
func (_SimplePaymentChannel *SimplePaymentChannelCaller) Expiration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimplePaymentChannel.contract.Call(opts, out, "expiration")
	return *ret0, err
}

// Expiration is a free data retrieval call binding the contract method 0x4665096d.
//
// Solidity: function expiration() constant returns(uint256)
func (_SimplePaymentChannel *SimplePaymentChannelSession) Expiration() (*big.Int, error) {
	return _SimplePaymentChannel.Contract.Expiration(&_SimplePaymentChannel.CallOpts)
}

// Expiration is a free data retrieval call binding the contract method 0x4665096d.
//
// Solidity: function expiration() constant returns(uint256)
func (_SimplePaymentChannel *SimplePaymentChannelCallerSession) Expiration() (*big.Int, error) {
	return _SimplePaymentChannel.Contract.Expiration(&_SimplePaymentChannel.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0xbd190a28.
//
// Solidity: function isValidSignature(uint256 _balance, uint256 _nonce, uint8 _v, bytes32 _r, bytes32 _s) constant returns(bool)
func (_SimplePaymentChannel *SimplePaymentChannelCaller) IsValidSignature(opts *bind.CallOpts, _balance *big.Int, _nonce *big.Int, _v uint8, _r [32]byte, _s [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SimplePaymentChannel.contract.Call(opts, out, "isValidSignature", _balance, _nonce, _v, _r, _s)
	return *ret0, err
}

// IsValidSignature is a free data retrieval call binding the contract method 0xbd190a28.
//
// Solidity: function isValidSignature(uint256 _balance, uint256 _nonce, uint8 _v, bytes32 _r, bytes32 _s) constant returns(bool)
func (_SimplePaymentChannel *SimplePaymentChannelSession) IsValidSignature(_balance *big.Int, _nonce *big.Int, _v uint8, _r [32]byte, _s [32]byte) (bool, error) {
	return _SimplePaymentChannel.Contract.IsValidSignature(&_SimplePaymentChannel.CallOpts, _balance, _nonce, _v, _r, _s)
}

// IsValidSignature is a free data retrieval call binding the contract method 0xbd190a28.
//
// Solidity: function isValidSignature(uint256 _balance, uint256 _nonce, uint8 _v, bytes32 _r, bytes32 _s) constant returns(bool)
func (_SimplePaymentChannel *SimplePaymentChannelCallerSession) IsValidSignature(_balance *big.Int, _nonce *big.Int, _v uint8, _r [32]byte, _s [32]byte) (bool, error) {
	return _SimplePaymentChannel.Contract.IsValidSignature(&_SimplePaymentChannel.CallOpts, _balance, _nonce, _v, _r, _s)
}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() constant returns(address)
func (_SimplePaymentChannel *SimplePaymentChannelCaller) Recipient(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SimplePaymentChannel.contract.Call(opts, out, "recipient")
	return *ret0, err
}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() constant returns(address)
func (_SimplePaymentChannel *SimplePaymentChannelSession) Recipient() (common.Address, error) {
	return _SimplePaymentChannel.Contract.Recipient(&_SimplePaymentChannel.CallOpts)
}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() constant returns(address)
func (_SimplePaymentChannel *SimplePaymentChannelCallerSession) Recipient() (common.Address, error) {
	return _SimplePaymentChannel.Contract.Recipient(&_SimplePaymentChannel.CallOpts)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() constant returns(address)
func (_SimplePaymentChannel *SimplePaymentChannelCaller) Sender(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SimplePaymentChannel.contract.Call(opts, out, "sender")
	return *ret0, err
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() constant returns(address)
func (_SimplePaymentChannel *SimplePaymentChannelSession) Sender() (common.Address, error) {
	return _SimplePaymentChannel.Contract.Sender(&_SimplePaymentChannel.CallOpts)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() constant returns(address)
func (_SimplePaymentChannel *SimplePaymentChannelCallerSession) Sender() (common.Address, error) {
	return _SimplePaymentChannel.Contract.Sender(&_SimplePaymentChannel.CallOpts)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes _sig) constant returns(uint8, bytes32, bytes32)
func (_SimplePaymentChannel *SimplePaymentChannelCaller) SplitSignature(opts *bind.CallOpts, _sig []byte) (uint8, [32]byte, [32]byte, error) {
	var (
		ret0 = new(uint8)
		ret1 = new([32]byte)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _SimplePaymentChannel.contract.Call(opts, out, "splitSignature", _sig)
	return *ret0, *ret1, *ret2, err
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes _sig) constant returns(uint8, bytes32, bytes32)
func (_SimplePaymentChannel *SimplePaymentChannelSession) SplitSignature(_sig []byte) (uint8, [32]byte, [32]byte, error) {
	return _SimplePaymentChannel.Contract.SplitSignature(&_SimplePaymentChannel.CallOpts, _sig)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes _sig) constant returns(uint8, bytes32, bytes32)
func (_SimplePaymentChannel *SimplePaymentChannelCallerSession) SplitSignature(_sig []byte) (uint8, [32]byte, [32]byte, error) {
	return _SimplePaymentChannel.Contract.SplitSignature(&_SimplePaymentChannel.CallOpts, _sig)
}

// ClaimTimeout is a paid mutator transaction binding the contract method 0x0e1da6c3.
//
// Solidity: function claimTimeout() returns()
func (_SimplePaymentChannel *SimplePaymentChannelTransactor) ClaimTimeout(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimplePaymentChannel.contract.Transact(opts, "claimTimeout")
}

// ClaimTimeout is a paid mutator transaction binding the contract method 0x0e1da6c3.
//
// Solidity: function claimTimeout() returns()
func (_SimplePaymentChannel *SimplePaymentChannelSession) ClaimTimeout() (*types.Transaction, error) {
	return _SimplePaymentChannel.Contract.ClaimTimeout(&_SimplePaymentChannel.TransactOpts)
}

// ClaimTimeout is a paid mutator transaction binding the contract method 0x0e1da6c3.
//
// Solidity: function claimTimeout() returns()
func (_SimplePaymentChannel *SimplePaymentChannelTransactorSession) ClaimTimeout() (*types.Transaction, error) {
	return _SimplePaymentChannel.Contract.ClaimTimeout(&_SimplePaymentChannel.TransactOpts)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x75be62fc.
//
// Solidity: function closeChannel(uint256 _balance, uint256 _nonce, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_SimplePaymentChannel *SimplePaymentChannelTransactor) CloseChannel(opts *bind.TransactOpts, _balance *big.Int, _nonce *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SimplePaymentChannel.contract.Transact(opts, "closeChannel", _balance, _nonce, _v, _r, _s)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x75be62fc.
//
// Solidity: function closeChannel(uint256 _balance, uint256 _nonce, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_SimplePaymentChannel *SimplePaymentChannelSession) CloseChannel(_balance *big.Int, _nonce *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SimplePaymentChannel.Contract.CloseChannel(&_SimplePaymentChannel.TransactOpts, _balance, _nonce, _v, _r, _s)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x75be62fc.
//
// Solidity: function closeChannel(uint256 _balance, uint256 _nonce, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_SimplePaymentChannel *SimplePaymentChannelTransactorSession) CloseChannel(_balance *big.Int, _nonce *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SimplePaymentChannel.Contract.CloseChannel(&_SimplePaymentChannel.TransactOpts, _balance, _nonce, _v, _r, _s)
}

// ExtendExpiration is a paid mutator transaction binding the contract method 0x0a114a70.
//
// Solidity: function extendExpiration(uint256 _expiration) returns()
func (_SimplePaymentChannel *SimplePaymentChannelTransactor) ExtendExpiration(opts *bind.TransactOpts, _expiration *big.Int) (*types.Transaction, error) {
	return _SimplePaymentChannel.contract.Transact(opts, "extendExpiration", _expiration)
}

// ExtendExpiration is a paid mutator transaction binding the contract method 0x0a114a70.
//
// Solidity: function extendExpiration(uint256 _expiration) returns()
func (_SimplePaymentChannel *SimplePaymentChannelSession) ExtendExpiration(_expiration *big.Int) (*types.Transaction, error) {
	return _SimplePaymentChannel.Contract.ExtendExpiration(&_SimplePaymentChannel.TransactOpts, _expiration)
}

// ExtendExpiration is a paid mutator transaction binding the contract method 0x0a114a70.
//
// Solidity: function extendExpiration(uint256 _expiration) returns()
func (_SimplePaymentChannel *SimplePaymentChannelTransactorSession) ExtendExpiration(_expiration *big.Int) (*types.Transaction, error) {
	return _SimplePaymentChannel.Contract.ExtendExpiration(&_SimplePaymentChannel.TransactOpts, _expiration)
}

// SimplePaymentChannelChannelClosedIterator is returned from FilterChannelClosed and is used to iterate over the raw logs and unpacked data for ChannelClosed events raised by the SimplePaymentChannel contract.
type SimplePaymentChannelChannelClosedIterator struct {
	Event *SimplePaymentChannelChannelClosed // Event containing the contract specifics and raw log

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
func (it *SimplePaymentChannelChannelClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimplePaymentChannelChannelClosed)
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
		it.Event = new(SimplePaymentChannelChannelClosed)
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
func (it *SimplePaymentChannelChannelClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimplePaymentChannelChannelClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimplePaymentChannelChannelClosed represents a ChannelClosed event raised by the SimplePaymentChannel contract.
type SimplePaymentChannelChannelClosed struct {
	SenderBalance    *big.Int
	RecipientBalance *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterChannelClosed is a free log retrieval operation binding the contract event 0x1273b19bd37db7541ec10aebcb7899427619b2685565bd3c7b184c2732f8118e.
//
// Solidity: event ChannelClosed(uint256 senderBalance, uint256 recipientBalance)
func (_SimplePaymentChannel *SimplePaymentChannelFilterer) FilterChannelClosed(opts *bind.FilterOpts) (*SimplePaymentChannelChannelClosedIterator, error) {

	logs, sub, err := _SimplePaymentChannel.contract.FilterLogs(opts, "ChannelClosed")
	if err != nil {
		return nil, err
	}
	return &SimplePaymentChannelChannelClosedIterator{contract: _SimplePaymentChannel.contract, event: "ChannelClosed", logs: logs, sub: sub}, nil
}

// WatchChannelClosed is a free log subscription operation binding the contract event 0x1273b19bd37db7541ec10aebcb7899427619b2685565bd3c7b184c2732f8118e.
//
// Solidity: event ChannelClosed(uint256 senderBalance, uint256 recipientBalance)
func (_SimplePaymentChannel *SimplePaymentChannelFilterer) WatchChannelClosed(opts *bind.WatchOpts, sink chan<- *SimplePaymentChannelChannelClosed) (event.Subscription, error) {

	logs, sub, err := _SimplePaymentChannel.contract.WatchLogs(opts, "ChannelClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimplePaymentChannelChannelClosed)
				if err := _SimplePaymentChannel.contract.UnpackLog(event, "ChannelClosed", log); err != nil {
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

// SimplePaymentChannelChannelOpenedIterator is returned from FilterChannelOpened and is used to iterate over the raw logs and unpacked data for ChannelOpened events raised by the SimplePaymentChannel contract.
type SimplePaymentChannelChannelOpenedIterator struct {
	Event *SimplePaymentChannelChannelOpened // Event containing the contract specifics and raw log

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
func (it *SimplePaymentChannelChannelOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimplePaymentChannelChannelOpened)
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
		it.Event = new(SimplePaymentChannelChannelOpened)
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
func (it *SimplePaymentChannelChannelOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimplePaymentChannelChannelOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimplePaymentChannelChannelOpened represents a ChannelOpened event raised by the SimplePaymentChannel contract.
type SimplePaymentChannelChannelOpened struct {
	Sender     common.Address
	Recipient  common.Address
	Expiration *big.Int
	Deposit    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChannelOpened is a free log retrieval operation binding the contract event 0x9957eac353f4e194b7faa928965f2dfdedd22f864e067b5b78f98d50157c2df5.
//
// Solidity: event ChannelOpened(address sender, address recipient, uint256 expiration, uint256 deposit)
func (_SimplePaymentChannel *SimplePaymentChannelFilterer) FilterChannelOpened(opts *bind.FilterOpts) (*SimplePaymentChannelChannelOpenedIterator, error) {

	logs, sub, err := _SimplePaymentChannel.contract.FilterLogs(opts, "ChannelOpened")
	if err != nil {
		return nil, err
	}
	return &SimplePaymentChannelChannelOpenedIterator{contract: _SimplePaymentChannel.contract, event: "ChannelOpened", logs: logs, sub: sub}, nil
}

// WatchChannelOpened is a free log subscription operation binding the contract event 0x9957eac353f4e194b7faa928965f2dfdedd22f864e067b5b78f98d50157c2df5.
//
// Solidity: event ChannelOpened(address sender, address recipient, uint256 expiration, uint256 deposit)
func (_SimplePaymentChannel *SimplePaymentChannelFilterer) WatchChannelOpened(opts *bind.WatchOpts, sink chan<- *SimplePaymentChannelChannelOpened) (event.Subscription, error) {

	logs, sub, err := _SimplePaymentChannel.contract.WatchLogs(opts, "ChannelOpened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimplePaymentChannelChannelOpened)
				if err := _SimplePaymentChannel.contract.UnpackLog(event, "ChannelOpened", log); err != nil {
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

// SimplePaymentChannelCheckEcRecoverIterator is returned from FilterCheckEcRecover and is used to iterate over the raw logs and unpacked data for CheckEcRecover events raised by the SimplePaymentChannel contract.
type SimplePaymentChannelCheckEcRecoverIterator struct {
	Event *SimplePaymentChannelCheckEcRecover // Event containing the contract specifics and raw log

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
func (it *SimplePaymentChannelCheckEcRecoverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimplePaymentChannelCheckEcRecover)
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
		it.Event = new(SimplePaymentChannelCheckEcRecover)
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
func (it *SimplePaymentChannelCheckEcRecoverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimplePaymentChannelCheckEcRecoverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimplePaymentChannelCheckEcRecover represents a CheckEcRecover event raised by the SimplePaymentChannel contract.
type SimplePaymentChannelCheckEcRecover struct {
	CheckAddr common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCheckEcRecover is a free log retrieval operation binding the contract event 0x53d8579cf8cf21640b4171b2bbfdf0344f8397acd23543c571e3449f18964bf7.
//
// Solidity: event CheckEcRecover(address checkAddr)
func (_SimplePaymentChannel *SimplePaymentChannelFilterer) FilterCheckEcRecover(opts *bind.FilterOpts) (*SimplePaymentChannelCheckEcRecoverIterator, error) {

	logs, sub, err := _SimplePaymentChannel.contract.FilterLogs(opts, "CheckEcRecover")
	if err != nil {
		return nil, err
	}
	return &SimplePaymentChannelCheckEcRecoverIterator{contract: _SimplePaymentChannel.contract, event: "CheckEcRecover", logs: logs, sub: sub}, nil
}

// WatchCheckEcRecover is a free log subscription operation binding the contract event 0x53d8579cf8cf21640b4171b2bbfdf0344f8397acd23543c571e3449f18964bf7.
//
// Solidity: event CheckEcRecover(address checkAddr)
func (_SimplePaymentChannel *SimplePaymentChannelFilterer) WatchCheckEcRecover(opts *bind.WatchOpts, sink chan<- *SimplePaymentChannelCheckEcRecover) (event.Subscription, error) {

	logs, sub, err := _SimplePaymentChannel.contract.WatchLogs(opts, "CheckEcRecover")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimplePaymentChannelCheckEcRecover)
				if err := _SimplePaymentChannel.contract.UnpackLog(event, "CheckEcRecover", log); err != nil {
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
