// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokens

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

// ElvFundABI is the input ABI used to generate the binding from.
const ElvFundABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"drip1000Token\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"faucetStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"faucetName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"drip2000Token\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"turnFaucetOff\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"drip5000Token\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"turnFaucetOn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_faucetName\",\"type\":\"string\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"OneKTokenSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"TwoKTokenSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"FiveKTokenSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"FaucetOn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"FaucetOff\",\"type\":\"event\"}]"

// ElvFundBin is the compiled bytecode used for deploying new contracts.
const ElvFundBin = `0x608060408190527f4f776e61626c6532303139303532383139333830304d4c000000000000000000600055610a8c388190039081908339810160405280516001805432600160a060020a031991821681179092556002805490911690911790550180516100739060049060208401906100c2565b506005805460ff1916600117908190556040805160ff9290921615158252517fd2a3a2d721c242e2bb5894f541c2a05972d28379a5ae20423a5419a42da3bf639181900360200190a15061015d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061010357805160ff1916838001178555610130565b82800160010185558215610130579182015b82811115610130578251825591602001919060010190610115565b5061013c929150610140565b5090565b61015a91905b8082111561013c5760008155600101610146565b90565b6109208061016c6000396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301b8eb6081146100dc57806302d05d3f146100fd5780630dfec7961461012e57806312065fe014610157578063400207071461017e57806341c0e1b51461020857806354fd4d501461021d57806359a24294146102325780636d2e4b1b146102535780638da5cb5b14610274578063af570c0414610289578063b412506b1461029e578063cc230322146102b3578063ee8734d0146102d4578063f2fde38b146102e9575b005b3480156100e857600080fd5b506100da600160a060020a036004351661030a565b34801561010957600080fd5b5061011261038c565b60408051600160a060020a039092168252519081900360200190f35b34801561013a57600080fd5b5061014361039b565b604080519115158252519081900360200190f35b34801561016357600080fd5b5061016c6103a4565b60408051918252519081900360200190f35b34801561018a57600080fd5b506101936103a9565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101cd5781810151838201526020016101b5565b50505050905090810190601f1680156101fa5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561021457600080fd5b506100da610437565b34801561022957600080fd5b5061016c610473565b34801561023e57600080fd5b506100da600160a060020a0360043516610479565b34801561025f57600080fd5b506100da600160a060020a03600435166104fa565b34801561028057600080fd5b50610112610555565b34801561029557600080fd5b50610112610564565b3480156102aa57600080fd5b506100da610573565b3480156102bf57600080fd5b506100da600160a060020a03600435166105f2565b3480156102e057600080fd5b506100da610676565b3480156102f557600080fd5b506100da600160a060020a03600435166106fe565b60055460ff16151561031b57600080fd5b61032481610770565b151561032f57600080fd5b61034281683635c9adc5dea000006107c7565b61034d81603c6108d3565b60408051600160a060020a038316815290517fb451b2fb4e1bc44234460580d26a4f71a2693d5c5a72e515b233f2e20235386c9181900360200190a150565b600154600160a060020a031681565b60055460ff1681565b303190565b6004805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561042f5780601f106104045761010080835404028352916020019161042f565b820191906000526020600020905b81548152906001019060200180831161041257829003601f168201915b505050505081565b600254600160a060020a031632148061045a5750600254600160a060020a031633145b151561046557600080fd5b600254600160a060020a0316ff5b60005481565b60055460ff16151561048a57600080fd5b61049381610770565b151561049e57600080fd5b6104b181686c6b935b8bbd4000006107c7565b6104bc8160786108d3565b60408051600160a060020a038316815290517eaacc7878c456f81effdd03a5982712003474ea96fceb9ded19748a0160a60c9181900360200190a150565b600154600160a060020a0316321461051157600080fd5b600160a060020a038116151561052657600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a031681565b600354600160a060020a031681565b600254600160a060020a03163214806105965750600254600160a060020a031633145b15156105a157600080fd5b60055460ff1615156105b257600080fd5b6005805460ff19169055604080516000815290517f1f2881e73120325370afe6da38205f10340d474ae88deced4da5b739ca650bcf9181900360200190a1565b60055460ff16151561060357600080fd5b61060c81610770565b151561061757600080fd5b61062b8169010f0cf064dd592000006107c7565b6106378161012c6108d3565b60408051600160a060020a038316815290517f0cecbc87a45437e3e1904e0a78549073436f9ce37badc9b8619e01328a0999929181900360200190a150565b600254600160a060020a03163214806106995750600254600160a060020a031633145b15156106a457600080fd5b60055460ff16156106b457600080fd5b6005805460ff1916600117908190556040805160ff9290921615158252517fd2a3a2d721c242e2bb5894f541c2a05972d28379a5ae20423a5419a42da3bf639181900360200190a1565b600254600160a060020a03163214806107215750600254600160a060020a031633145b151561072c57600080fd5b600160a060020a038116151561074157600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a0381166000908152600660205260408120541515610797575060016107c2565b600160a060020a03821660009081526006602052604090205442106107be575060016107c2565b5060005b919050565b303181111561085d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f496e73756666696369656e742062616c616e636520696e20666175636574206660448201527f6f72207769746864726177616c20726571756573740000000000000000000000606482015290519081900360840190fd5b604051600160a060020a0383169082156108fc029083906000818181858888f19350505050158015610893573d6000803e3d6000fd5b50604080513381526020810183905281517f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65929181900390910190a15050565b600160a060020a0390911660009081526006602052604090204290910190555600a165627a7a72305820d802fa507f4daa6321de684a894b74640fa7ec5c957e365237411428b73f82330029`

// DeployElvFund deploys a new Ethereum contract, binding an instance of ElvFund to it.
func DeployElvFund(auth *bind.TransactOpts, backend bind.ContractBackend, _faucetName string) (common.Address, *types.Transaction, *ElvFund, error) {
	parsed, err := abi.JSON(strings.NewReader(ElvFundABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ElvFundBin), backend, _faucetName)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ElvFund{ElvFundCaller: ElvFundCaller{contract: contract}, ElvFundTransactor: ElvFundTransactor{contract: contract}, ElvFundFilterer: ElvFundFilterer{contract: contract}}, nil
}

// ElvFund is an auto generated Go binding around an Ethereum contract.
type ElvFund struct {
	ElvFundCaller     // Read-only binding to the contract
	ElvFundTransactor // Write-only binding to the contract
	ElvFundFilterer   // Log filterer for contract events
}

// ElvFundCaller is an auto generated read-only Go binding around an Ethereum contract.
type ElvFundCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElvFundTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ElvFundTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElvFundFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ElvFundFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElvFundSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ElvFundSession struct {
	Contract     *ElvFund          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ElvFundCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ElvFundCallerSession struct {
	Contract *ElvFundCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ElvFundTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ElvFundTransactorSession struct {
	Contract     *ElvFundTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ElvFundRaw is an auto generated low-level Go binding around an Ethereum contract.
type ElvFundRaw struct {
	Contract *ElvFund // Generic contract binding to access the raw methods on
}

// ElvFundCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ElvFundCallerRaw struct {
	Contract *ElvFundCaller // Generic read-only contract binding to access the raw methods on
}

// ElvFundTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ElvFundTransactorRaw struct {
	Contract *ElvFundTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElvFund creates a new instance of ElvFund, bound to a specific deployed contract.
func NewElvFund(address common.Address, backend bind.ContractBackend) (*ElvFund, error) {
	contract, err := bindElvFund(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ElvFund{ElvFundCaller: ElvFundCaller{contract: contract}, ElvFundTransactor: ElvFundTransactor{contract: contract}, ElvFundFilterer: ElvFundFilterer{contract: contract}}, nil
}

// NewElvFundCaller creates a new read-only instance of ElvFund, bound to a specific deployed contract.
func NewElvFundCaller(address common.Address, caller bind.ContractCaller) (*ElvFundCaller, error) {
	contract, err := bindElvFund(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ElvFundCaller{contract: contract}, nil
}

// NewElvFundTransactor creates a new write-only instance of ElvFund, bound to a specific deployed contract.
func NewElvFundTransactor(address common.Address, transactor bind.ContractTransactor) (*ElvFundTransactor, error) {
	contract, err := bindElvFund(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ElvFundTransactor{contract: contract}, nil
}

// NewElvFundFilterer creates a new log filterer instance of ElvFund, bound to a specific deployed contract.
func NewElvFundFilterer(address common.Address, filterer bind.ContractFilterer) (*ElvFundFilterer, error) {
	contract, err := bindElvFund(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ElvFundFilterer{contract: contract}, nil
}

// bindElvFund binds a generic wrapper to an already deployed contract.
func bindElvFund(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ElvFundABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElvFund *ElvFundRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ElvFund.Contract.ElvFundCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElvFund *ElvFundRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElvFund.Contract.ElvFundTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElvFund *ElvFundRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElvFund.Contract.ElvFundTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElvFund *ElvFundCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ElvFund.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElvFund *ElvFundTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElvFund.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElvFund *ElvFundTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElvFund.Contract.contract.Transact(opts, method, params...)
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_ElvFund *ElvFundCaller) ContentSpace(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ElvFund.contract.Call(opts, out, "contentSpace")
	return *ret0, err
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_ElvFund *ElvFundSession) ContentSpace() (common.Address, error) {
	return _ElvFund.Contract.ContentSpace(&_ElvFund.CallOpts)
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_ElvFund *ElvFundCallerSession) ContentSpace() (common.Address, error) {
	return _ElvFund.Contract.ContentSpace(&_ElvFund.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_ElvFund *ElvFundCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ElvFund.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_ElvFund *ElvFundSession) Creator() (common.Address, error) {
	return _ElvFund.Contract.Creator(&_ElvFund.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_ElvFund *ElvFundCallerSession) Creator() (common.Address, error) {
	return _ElvFund.Contract.Creator(&_ElvFund.CallOpts)
}

// FaucetName is a free data retrieval call binding the contract method 0x40020707.
//
// Solidity: function faucetName() constant returns(string)
func (_ElvFund *ElvFundCaller) FaucetName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ElvFund.contract.Call(opts, out, "faucetName")
	return *ret0, err
}

// FaucetName is a free data retrieval call binding the contract method 0x40020707.
//
// Solidity: function faucetName() constant returns(string)
func (_ElvFund *ElvFundSession) FaucetName() (string, error) {
	return _ElvFund.Contract.FaucetName(&_ElvFund.CallOpts)
}

// FaucetName is a free data retrieval call binding the contract method 0x40020707.
//
// Solidity: function faucetName() constant returns(string)
func (_ElvFund *ElvFundCallerSession) FaucetName() (string, error) {
	return _ElvFund.Contract.FaucetName(&_ElvFund.CallOpts)
}

// FaucetStatus is a free data retrieval call binding the contract method 0x0dfec796.
//
// Solidity: function faucetStatus() constant returns(bool)
func (_ElvFund *ElvFundCaller) FaucetStatus(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ElvFund.contract.Call(opts, out, "faucetStatus")
	return *ret0, err
}

// FaucetStatus is a free data retrieval call binding the contract method 0x0dfec796.
//
// Solidity: function faucetStatus() constant returns(bool)
func (_ElvFund *ElvFundSession) FaucetStatus() (bool, error) {
	return _ElvFund.Contract.FaucetStatus(&_ElvFund.CallOpts)
}

// FaucetStatus is a free data retrieval call binding the contract method 0x0dfec796.
//
// Solidity: function faucetStatus() constant returns(bool)
func (_ElvFund *ElvFundCallerSession) FaucetStatus() (bool, error) {
	return _ElvFund.Contract.FaucetStatus(&_ElvFund.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_ElvFund *ElvFundCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ElvFund.contract.Call(opts, out, "getBalance")
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_ElvFund *ElvFundSession) GetBalance() (*big.Int, error) {
	return _ElvFund.Contract.GetBalance(&_ElvFund.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_ElvFund *ElvFundCallerSession) GetBalance() (*big.Int, error) {
	return _ElvFund.Contract.GetBalance(&_ElvFund.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ElvFund *ElvFundCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ElvFund.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ElvFund *ElvFundSession) Owner() (common.Address, error) {
	return _ElvFund.Contract.Owner(&_ElvFund.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ElvFund *ElvFundCallerSession) Owner() (common.Address, error) {
	return _ElvFund.Contract.Owner(&_ElvFund.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_ElvFund *ElvFundCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ElvFund.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_ElvFund *ElvFundSession) Version() ([32]byte, error) {
	return _ElvFund.Contract.Version(&_ElvFund.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_ElvFund *ElvFundCallerSession) Version() ([32]byte, error) {
	return _ElvFund.Contract.Version(&_ElvFund.CallOpts)
}

// Drip1000Token is a paid mutator transaction binding the contract method 0x01b8eb60.
//
// Solidity: function drip1000Token(address _to) returns()
func (_ElvFund *ElvFundTransactor) Drip1000Token(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ElvFund.contract.Transact(opts, "drip1000Token", _to)
}

// Drip1000Token is a paid mutator transaction binding the contract method 0x01b8eb60.
//
// Solidity: function drip1000Token(address _to) returns()
func (_ElvFund *ElvFundSession) Drip1000Token(_to common.Address) (*types.Transaction, error) {
	return _ElvFund.Contract.Drip1000Token(&_ElvFund.TransactOpts, _to)
}

// Drip1000Token is a paid mutator transaction binding the contract method 0x01b8eb60.
//
// Solidity: function drip1000Token(address _to) returns()
func (_ElvFund *ElvFundTransactorSession) Drip1000Token(_to common.Address) (*types.Transaction, error) {
	return _ElvFund.Contract.Drip1000Token(&_ElvFund.TransactOpts, _to)
}

// Drip2000Token is a paid mutator transaction binding the contract method 0x59a24294.
//
// Solidity: function drip2000Token(address _to) returns()
func (_ElvFund *ElvFundTransactor) Drip2000Token(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ElvFund.contract.Transact(opts, "drip2000Token", _to)
}

// Drip2000Token is a paid mutator transaction binding the contract method 0x59a24294.
//
// Solidity: function drip2000Token(address _to) returns()
func (_ElvFund *ElvFundSession) Drip2000Token(_to common.Address) (*types.Transaction, error) {
	return _ElvFund.Contract.Drip2000Token(&_ElvFund.TransactOpts, _to)
}

// Drip2000Token is a paid mutator transaction binding the contract method 0x59a24294.
//
// Solidity: function drip2000Token(address _to) returns()
func (_ElvFund *ElvFundTransactorSession) Drip2000Token(_to common.Address) (*types.Transaction, error) {
	return _ElvFund.Contract.Drip2000Token(&_ElvFund.TransactOpts, _to)
}

// Drip5000Token is a paid mutator transaction binding the contract method 0xcc230322.
//
// Solidity: function drip5000Token(address _to) returns()
func (_ElvFund *ElvFundTransactor) Drip5000Token(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ElvFund.contract.Transact(opts, "drip5000Token", _to)
}

// Drip5000Token is a paid mutator transaction binding the contract method 0xcc230322.
//
// Solidity: function drip5000Token(address _to) returns()
func (_ElvFund *ElvFundSession) Drip5000Token(_to common.Address) (*types.Transaction, error) {
	return _ElvFund.Contract.Drip5000Token(&_ElvFund.TransactOpts, _to)
}

// Drip5000Token is a paid mutator transaction binding the contract method 0xcc230322.
//
// Solidity: function drip5000Token(address _to) returns()
func (_ElvFund *ElvFundTransactorSession) Drip5000Token(_to common.Address) (*types.Transaction, error) {
	return _ElvFund.Contract.Drip5000Token(&_ElvFund.TransactOpts, _to)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ElvFund *ElvFundTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElvFund.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ElvFund *ElvFundSession) Kill() (*types.Transaction, error) {
	return _ElvFund.Contract.Kill(&_ElvFund.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ElvFund *ElvFundTransactorSession) Kill() (*types.Transaction, error) {
	return _ElvFund.Contract.Kill(&_ElvFund.TransactOpts)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_ElvFund *ElvFundTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _ElvFund.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_ElvFund *ElvFundSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _ElvFund.Contract.TransferCreatorship(&_ElvFund.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_ElvFund *ElvFundTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _ElvFund.Contract.TransferCreatorship(&_ElvFund.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ElvFund *ElvFundTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ElvFund.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ElvFund *ElvFundSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ElvFund.Contract.TransferOwnership(&_ElvFund.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ElvFund *ElvFundTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ElvFund.Contract.TransferOwnership(&_ElvFund.TransactOpts, newOwner)
}

// TurnFaucetOff is a paid mutator transaction binding the contract method 0xb412506b.
//
// Solidity: function turnFaucetOff() returns()
func (_ElvFund *ElvFundTransactor) TurnFaucetOff(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElvFund.contract.Transact(opts, "turnFaucetOff")
}

// TurnFaucetOff is a paid mutator transaction binding the contract method 0xb412506b.
//
// Solidity: function turnFaucetOff() returns()
func (_ElvFund *ElvFundSession) TurnFaucetOff() (*types.Transaction, error) {
	return _ElvFund.Contract.TurnFaucetOff(&_ElvFund.TransactOpts)
}

// TurnFaucetOff is a paid mutator transaction binding the contract method 0xb412506b.
//
// Solidity: function turnFaucetOff() returns()
func (_ElvFund *ElvFundTransactorSession) TurnFaucetOff() (*types.Transaction, error) {
	return _ElvFund.Contract.TurnFaucetOff(&_ElvFund.TransactOpts)
}

// TurnFaucetOn is a paid mutator transaction binding the contract method 0xee8734d0.
//
// Solidity: function turnFaucetOn() returns()
func (_ElvFund *ElvFundTransactor) TurnFaucetOn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElvFund.contract.Transact(opts, "turnFaucetOn")
}

// TurnFaucetOn is a paid mutator transaction binding the contract method 0xee8734d0.
//
// Solidity: function turnFaucetOn() returns()
func (_ElvFund *ElvFundSession) TurnFaucetOn() (*types.Transaction, error) {
	return _ElvFund.Contract.TurnFaucetOn(&_ElvFund.TransactOpts)
}

// TurnFaucetOn is a paid mutator transaction binding the contract method 0xee8734d0.
//
// Solidity: function turnFaucetOn() returns()
func (_ElvFund *ElvFundTransactorSession) TurnFaucetOn() (*types.Transaction, error) {
	return _ElvFund.Contract.TurnFaucetOn(&_ElvFund.TransactOpts)
}

// ElvFundFaucetOffIterator is returned from FilterFaucetOff and is used to iterate over the raw logs and unpacked data for FaucetOff events raised by the ElvFund contract.
type ElvFundFaucetOffIterator struct {
	Event *ElvFundFaucetOff // Event containing the contract specifics and raw log

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
func (it *ElvFundFaucetOffIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElvFundFaucetOff)
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
		it.Event = new(ElvFundFaucetOff)
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
func (it *ElvFundFaucetOffIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElvFundFaucetOffIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElvFundFaucetOff represents a FaucetOff event raised by the ElvFund contract.
type ElvFundFaucetOff struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFaucetOff is a free log retrieval operation binding the contract event 0x1f2881e73120325370afe6da38205f10340d474ae88deced4da5b739ca650bcf.
//
// Solidity: event FaucetOff(bool status)
func (_ElvFund *ElvFundFilterer) FilterFaucetOff(opts *bind.FilterOpts) (*ElvFundFaucetOffIterator, error) {

	logs, sub, err := _ElvFund.contract.FilterLogs(opts, "FaucetOff")
	if err != nil {
		return nil, err
	}
	return &ElvFundFaucetOffIterator{contract: _ElvFund.contract, event: "FaucetOff", logs: logs, sub: sub}, nil
}

// WatchFaucetOff is a free log subscription operation binding the contract event 0x1f2881e73120325370afe6da38205f10340d474ae88deced4da5b739ca650bcf.
//
// Solidity: event FaucetOff(bool status)
func (_ElvFund *ElvFundFilterer) WatchFaucetOff(opts *bind.WatchOpts, sink chan<- *ElvFundFaucetOff) (event.Subscription, error) {

	logs, sub, err := _ElvFund.contract.WatchLogs(opts, "FaucetOff")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElvFundFaucetOff)
				if err := _ElvFund.contract.UnpackLog(event, "FaucetOff", log); err != nil {
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

// ElvFundFaucetOnIterator is returned from FilterFaucetOn and is used to iterate over the raw logs and unpacked data for FaucetOn events raised by the ElvFund contract.
type ElvFundFaucetOnIterator struct {
	Event *ElvFundFaucetOn // Event containing the contract specifics and raw log

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
func (it *ElvFundFaucetOnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElvFundFaucetOn)
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
		it.Event = new(ElvFundFaucetOn)
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
func (it *ElvFundFaucetOnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElvFundFaucetOnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElvFundFaucetOn represents a FaucetOn event raised by the ElvFund contract.
type ElvFundFaucetOn struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFaucetOn is a free log retrieval operation binding the contract event 0xd2a3a2d721c242e2bb5894f541c2a05972d28379a5ae20423a5419a42da3bf63.
//
// Solidity: event FaucetOn(bool status)
func (_ElvFund *ElvFundFilterer) FilterFaucetOn(opts *bind.FilterOpts) (*ElvFundFaucetOnIterator, error) {

	logs, sub, err := _ElvFund.contract.FilterLogs(opts, "FaucetOn")
	if err != nil {
		return nil, err
	}
	return &ElvFundFaucetOnIterator{contract: _ElvFund.contract, event: "FaucetOn", logs: logs, sub: sub}, nil
}

// WatchFaucetOn is a free log subscription operation binding the contract event 0xd2a3a2d721c242e2bb5894f541c2a05972d28379a5ae20423a5419a42da3bf63.
//
// Solidity: event FaucetOn(bool status)
func (_ElvFund *ElvFundFilterer) WatchFaucetOn(opts *bind.WatchOpts, sink chan<- *ElvFundFaucetOn) (event.Subscription, error) {

	logs, sub, err := _ElvFund.contract.WatchLogs(opts, "FaucetOn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElvFundFaucetOn)
				if err := _ElvFund.contract.UnpackLog(event, "FaucetOn", log); err != nil {
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

// ElvFundFiveKTokenSentIterator is returned from FilterFiveKTokenSent and is used to iterate over the raw logs and unpacked data for FiveKTokenSent events raised by the ElvFund contract.
type ElvFundFiveKTokenSentIterator struct {
	Event *ElvFundFiveKTokenSent // Event containing the contract specifics and raw log

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
func (it *ElvFundFiveKTokenSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElvFundFiveKTokenSent)
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
		it.Event = new(ElvFundFiveKTokenSent)
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
func (it *ElvFundFiveKTokenSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElvFundFiveKTokenSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElvFundFiveKTokenSent represents a FiveKTokenSent event raised by the ElvFund contract.
type ElvFundFiveKTokenSent struct {
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFiveKTokenSent is a free log retrieval operation binding the contract event 0x0cecbc87a45437e3e1904e0a78549073436f9ce37badc9b8619e01328a099992.
//
// Solidity: event FiveKTokenSent(address receiver)
func (_ElvFund *ElvFundFilterer) FilterFiveKTokenSent(opts *bind.FilterOpts) (*ElvFundFiveKTokenSentIterator, error) {

	logs, sub, err := _ElvFund.contract.FilterLogs(opts, "FiveKTokenSent")
	if err != nil {
		return nil, err
	}
	return &ElvFundFiveKTokenSentIterator{contract: _ElvFund.contract, event: "FiveKTokenSent", logs: logs, sub: sub}, nil
}

// WatchFiveKTokenSent is a free log subscription operation binding the contract event 0x0cecbc87a45437e3e1904e0a78549073436f9ce37badc9b8619e01328a099992.
//
// Solidity: event FiveKTokenSent(address receiver)
func (_ElvFund *ElvFundFilterer) WatchFiveKTokenSent(opts *bind.WatchOpts, sink chan<- *ElvFundFiveKTokenSent) (event.Subscription, error) {

	logs, sub, err := _ElvFund.contract.WatchLogs(opts, "FiveKTokenSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElvFundFiveKTokenSent)
				if err := _ElvFund.contract.UnpackLog(event, "FiveKTokenSent", log); err != nil {
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

// ElvFundOneKTokenSentIterator is returned from FilterOneKTokenSent and is used to iterate over the raw logs and unpacked data for OneKTokenSent events raised by the ElvFund contract.
type ElvFundOneKTokenSentIterator struct {
	Event *ElvFundOneKTokenSent // Event containing the contract specifics and raw log

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
func (it *ElvFundOneKTokenSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElvFundOneKTokenSent)
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
		it.Event = new(ElvFundOneKTokenSent)
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
func (it *ElvFundOneKTokenSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElvFundOneKTokenSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElvFundOneKTokenSent represents a OneKTokenSent event raised by the ElvFund contract.
type ElvFundOneKTokenSent struct {
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOneKTokenSent is a free log retrieval operation binding the contract event 0xb451b2fb4e1bc44234460580d26a4f71a2693d5c5a72e515b233f2e20235386c.
//
// Solidity: event OneKTokenSent(address receiver)
func (_ElvFund *ElvFundFilterer) FilterOneKTokenSent(opts *bind.FilterOpts) (*ElvFundOneKTokenSentIterator, error) {

	logs, sub, err := _ElvFund.contract.FilterLogs(opts, "OneKTokenSent")
	if err != nil {
		return nil, err
	}
	return &ElvFundOneKTokenSentIterator{contract: _ElvFund.contract, event: "OneKTokenSent", logs: logs, sub: sub}, nil
}

// WatchOneKTokenSent is a free log subscription operation binding the contract event 0xb451b2fb4e1bc44234460580d26a4f71a2693d5c5a72e515b233f2e20235386c.
//
// Solidity: event OneKTokenSent(address receiver)
func (_ElvFund *ElvFundFilterer) WatchOneKTokenSent(opts *bind.WatchOpts, sink chan<- *ElvFundOneKTokenSent) (event.Subscription, error) {

	logs, sub, err := _ElvFund.contract.WatchLogs(opts, "OneKTokenSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElvFundOneKTokenSent)
				if err := _ElvFund.contract.UnpackLog(event, "OneKTokenSent", log); err != nil {
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

// ElvFundTwoKTokenSentIterator is returned from FilterTwoKTokenSent and is used to iterate over the raw logs and unpacked data for TwoKTokenSent events raised by the ElvFund contract.
type ElvFundTwoKTokenSentIterator struct {
	Event *ElvFundTwoKTokenSent // Event containing the contract specifics and raw log

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
func (it *ElvFundTwoKTokenSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElvFundTwoKTokenSent)
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
		it.Event = new(ElvFundTwoKTokenSent)
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
func (it *ElvFundTwoKTokenSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElvFundTwoKTokenSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElvFundTwoKTokenSent represents a TwoKTokenSent event raised by the ElvFund contract.
type ElvFundTwoKTokenSent struct {
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTwoKTokenSent is a free log retrieval operation binding the contract event 0x00aacc7878c456f81effdd03a5982712003474ea96fceb9ded19748a0160a60c.
//
// Solidity: event TwoKTokenSent(address receiver)
func (_ElvFund *ElvFundFilterer) FilterTwoKTokenSent(opts *bind.FilterOpts) (*ElvFundTwoKTokenSentIterator, error) {

	logs, sub, err := _ElvFund.contract.FilterLogs(opts, "TwoKTokenSent")
	if err != nil {
		return nil, err
	}
	return &ElvFundTwoKTokenSentIterator{contract: _ElvFund.contract, event: "TwoKTokenSent", logs: logs, sub: sub}, nil
}

// WatchTwoKTokenSent is a free log subscription operation binding the contract event 0x00aacc7878c456f81effdd03a5982712003474ea96fceb9ded19748a0160a60c.
//
// Solidity: event TwoKTokenSent(address receiver)
func (_ElvFund *ElvFundFilterer) WatchTwoKTokenSent(opts *bind.WatchOpts, sink chan<- *ElvFundTwoKTokenSent) (event.Subscription, error) {

	logs, sub, err := _ElvFund.contract.WatchLogs(opts, "TwoKTokenSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElvFundTwoKTokenSent)
				if err := _ElvFund.contract.UnpackLog(event, "TwoKTokenSent", log); err != nil {
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

// ElvFundWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the ElvFund contract.
type ElvFundWithdrawalIterator struct {
	Event *ElvFundWithdrawal // Event containing the contract specifics and raw log

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
func (it *ElvFundWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElvFundWithdrawal)
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
		it.Event = new(ElvFundWithdrawal)
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
func (it *ElvFundWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElvFundWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElvFundWithdrawal represents a Withdrawal event raised by the ElvFund contract.
type ElvFundWithdrawal struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address to, uint256 amount)
func (_ElvFund *ElvFundFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*ElvFundWithdrawalIterator, error) {

	logs, sub, err := _ElvFund.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &ElvFundWithdrawalIterator{contract: _ElvFund.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address to, uint256 amount)
func (_ElvFund *ElvFundFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ElvFundWithdrawal) (event.Subscription, error) {

	logs, sub, err := _ElvFund.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElvFundWithdrawal)
				if err := _ElvFund.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x60806040527f4f776e61626c6532303139303532383139333830304d4c00000000000000000060005560018054600160a060020a0319908116329081179092556002805490911690911790556102c58061005a6000396000f3006080604052600436106100825763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f811461008457806341c0e1b5146100b557806354fd4d50146100ca5780636d2e4b1b146100f15780638da5cb5b14610112578063af570c0414610127578063f2fde38b1461013c575b005b34801561009057600080fd5b5061009961015d565b60408051600160a060020a039092168252519081900360200190f35b3480156100c157600080fd5b5061008261016c565b3480156100d657600080fd5b506100df6101a8565b60408051918252519081900360200190f35b3480156100fd57600080fd5b50610082600160a060020a03600435166101ae565b34801561011e57600080fd5b50610099610209565b34801561013357600080fd5b50610099610218565b34801561014857600080fd5b50610082600160a060020a0360043516610227565b600154600160a060020a031681565b600254600160a060020a031632148061018f5750600254600160a060020a031633145b151561019a57600080fd5b600254600160a060020a0316ff5b60005481565b600154600160a060020a031632146101c557600080fd5b600160a060020a03811615156101da57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a031681565b600354600160a060020a031681565b600254600160a060020a031632148061024a5750600254600160a060020a031633145b151561025557600080fd5b600160a060020a038116151561026a57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820747d708664be33b1214aeab86aa309e2912acf3fa357d263ffd39f0d0b64fefb0029`

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

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_Ownable *OwnableCaller) ContentSpace(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "contentSpace")
	return *ret0, err
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_Ownable *OwnableSession) ContentSpace() (common.Address, error) {
	return _Ownable.Contract.ContentSpace(&_Ownable.CallOpts)
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_Ownable *OwnableCallerSession) ContentSpace() (common.Address, error) {
	return _Ownable.Contract.ContentSpace(&_Ownable.CallOpts)
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
