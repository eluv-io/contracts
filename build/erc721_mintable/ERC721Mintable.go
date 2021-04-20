// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc721_mintable

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

// Erc721MintableABI is the input ABI used to generate the binding from.
const Erc721MintableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"_approved\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementsERC721\",\"outputs\":[{\"name\":\"_implementsERC721\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"_totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"name\":\"_infoUrl\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_approvedAddress\",\"type\":\"address\"},{\"name\":\"_metadata\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numTokensTotal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getOwnerTokens\",\"outputs\":[{\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// Erc721Mintable is an auto generated Go binding around an Ethereum contract.
type Erc721Mintable struct {
	Erc721MintableCaller     // Read-only binding to the contract
	Erc721MintableTransactor // Write-only binding to the contract
	Erc721MintableFilterer   // Log filterer for contract events
}

// Erc721MintableCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc721MintableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721MintableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc721MintableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721MintableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc721MintableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721MintableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc721MintableSession struct {
	Contract     *Erc721Mintable   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc721MintableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc721MintableCallerSession struct {
	Contract *Erc721MintableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Erc721MintableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc721MintableTransactorSession struct {
	Contract     *Erc721MintableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Erc721MintableRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc721MintableRaw struct {
	Contract *Erc721Mintable // Generic contract binding to access the raw methods on
}

// Erc721MintableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc721MintableCallerRaw struct {
	Contract *Erc721MintableCaller // Generic read-only contract binding to access the raw methods on
}

// Erc721MintableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc721MintableTransactorRaw struct {
	Contract *Erc721MintableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc721Mintable creates a new instance of Erc721Mintable, bound to a specific deployed contract.
func NewErc721Mintable(address common.Address, backend bind.ContractBackend) (*Erc721Mintable, error) {
	contract, err := bindErc721Mintable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc721Mintable{Erc721MintableCaller: Erc721MintableCaller{contract: contract}, Erc721MintableTransactor: Erc721MintableTransactor{contract: contract}, Erc721MintableFilterer: Erc721MintableFilterer{contract: contract}}, nil
}

// NewErc721MintableCaller creates a new read-only instance of Erc721Mintable, bound to a specific deployed contract.
func NewErc721MintableCaller(address common.Address, caller bind.ContractCaller) (*Erc721MintableCaller, error) {
	contract, err := bindErc721Mintable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721MintableCaller{contract: contract}, nil
}

// NewErc721MintableTransactor creates a new write-only instance of Erc721Mintable, bound to a specific deployed contract.
func NewErc721MintableTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc721MintableTransactor, error) {
	contract, err := bindErc721Mintable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721MintableTransactor{contract: contract}, nil
}

// NewErc721MintableFilterer creates a new log filterer instance of Erc721Mintable, bound to a specific deployed contract.
func NewErc721MintableFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc721MintableFilterer, error) {
	contract, err := bindErc721Mintable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc721MintableFilterer{contract: contract}, nil
}

// bindErc721Mintable binds a generic wrapper to an already deployed contract.
func bindErc721Mintable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc721MintableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721Mintable *Erc721MintableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Erc721Mintable.Contract.Erc721MintableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721Mintable *Erc721MintableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721Mintable.Contract.Erc721MintableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721Mintable *Erc721MintableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721Mintable.Contract.Erc721MintableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721Mintable *Erc721MintableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Erc721Mintable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721Mintable *Erc721MintableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721Mintable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721Mintable *Erc721MintableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721Mintable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 _balance)
func (_Erc721Mintable *Erc721MintableCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Erc721Mintable.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 _balance)
func (_Erc721Mintable *Erc721MintableSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Erc721Mintable.Contract.BalanceOf(&_Erc721Mintable.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 _balance)
func (_Erc721Mintable *Erc721MintableCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Erc721Mintable.Contract.BalanceOf(&_Erc721Mintable.CallOpts, _owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address _approved)
func (_Erc721Mintable *Erc721MintableCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Erc721Mintable.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address _approved)
func (_Erc721Mintable *Erc721MintableSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Erc721Mintable.Contract.GetApproved(&_Erc721Mintable.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address _approved)
func (_Erc721Mintable *Erc721MintableCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Erc721Mintable.Contract.GetApproved(&_Erc721Mintable.CallOpts, _tokenId)
}

// GetOwnerTokens is a free data retrieval call binding the contract method 0xd63d4af0.
//
// Solidity: function getOwnerTokens(address _owner) constant returns(uint256[] _tokenIds)
func (_Erc721Mintable *Erc721MintableCaller) GetOwnerTokens(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Erc721Mintable.contract.Call(opts, out, "getOwnerTokens", _owner)
	return *ret0, err
}

// GetOwnerTokens is a free data retrieval call binding the contract method 0xd63d4af0.
//
// Solidity: function getOwnerTokens(address _owner) constant returns(uint256[] _tokenIds)
func (_Erc721Mintable *Erc721MintableSession) GetOwnerTokens(_owner common.Address) ([]*big.Int, error) {
	return _Erc721Mintable.Contract.GetOwnerTokens(&_Erc721Mintable.CallOpts, _owner)
}

// GetOwnerTokens is a free data retrieval call binding the contract method 0xd63d4af0.
//
// Solidity: function getOwnerTokens(address _owner) constant returns(uint256[] _tokenIds)
func (_Erc721Mintable *Erc721MintableCallerSession) GetOwnerTokens(_owner common.Address) ([]*big.Int, error) {
	return _Erc721Mintable.Contract.GetOwnerTokens(&_Erc721Mintable.CallOpts, _owner)
}

// ImplementsERC721 is a free data retrieval call binding the contract method 0x1051db34.
//
// Solidity: function implementsERC721() constant returns(bool _implementsERC721)
func (_Erc721Mintable *Erc721MintableCaller) ImplementsERC721(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Erc721Mintable.contract.Call(opts, out, "implementsERC721")
	return *ret0, err
}

// ImplementsERC721 is a free data retrieval call binding the contract method 0x1051db34.
//
// Solidity: function implementsERC721() constant returns(bool _implementsERC721)
func (_Erc721Mintable *Erc721MintableSession) ImplementsERC721() (bool, error) {
	return _Erc721Mintable.Contract.ImplementsERC721(&_Erc721Mintable.CallOpts)
}

// ImplementsERC721 is a free data retrieval call binding the contract method 0x1051db34.
//
// Solidity: function implementsERC721() constant returns(bool _implementsERC721)
func (_Erc721Mintable *Erc721MintableCallerSession) ImplementsERC721() (bool, error) {
	return _Erc721Mintable.Contract.ImplementsERC721(&_Erc721Mintable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string _name)
func (_Erc721Mintable *Erc721MintableCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Erc721Mintable.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string _name)
func (_Erc721Mintable *Erc721MintableSession) Name() (string, error) {
	return _Erc721Mintable.Contract.Name(&_Erc721Mintable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string _name)
func (_Erc721Mintable *Erc721MintableCallerSession) Name() (string, error) {
	return _Erc721Mintable.Contract.Name(&_Erc721Mintable.CallOpts)
}

// NumTokensTotal is a free data retrieval call binding the contract method 0xaf129dc2.
//
// Solidity: function numTokensTotal() constant returns(uint256)
func (_Erc721Mintable *Erc721MintableCaller) NumTokensTotal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Erc721Mintable.contract.Call(opts, out, "numTokensTotal")
	return *ret0, err
}

// NumTokensTotal is a free data retrieval call binding the contract method 0xaf129dc2.
//
// Solidity: function numTokensTotal() constant returns(uint256)
func (_Erc721Mintable *Erc721MintableSession) NumTokensTotal() (*big.Int, error) {
	return _Erc721Mintable.Contract.NumTokensTotal(&_Erc721Mintable.CallOpts)
}

// NumTokensTotal is a free data retrieval call binding the contract method 0xaf129dc2.
//
// Solidity: function numTokensTotal() constant returns(uint256)
func (_Erc721Mintable *Erc721MintableCallerSession) NumTokensTotal() (*big.Int, error) {
	return _Erc721Mintable.Contract.NumTokensTotal(&_Erc721Mintable.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address _owner)
func (_Erc721Mintable *Erc721MintableCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Erc721Mintable.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address _owner)
func (_Erc721Mintable *Erc721MintableSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Erc721Mintable.Contract.OwnerOf(&_Erc721Mintable.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address _owner)
func (_Erc721Mintable *Erc721MintableCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Erc721Mintable.Contract.OwnerOf(&_Erc721Mintable.CallOpts, _tokenId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string _symbol)
func (_Erc721Mintable *Erc721MintableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Erc721Mintable.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string _symbol)
func (_Erc721Mintable *Erc721MintableSession) Symbol() (string, error) {
	return _Erc721Mintable.Contract.Symbol(&_Erc721Mintable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string _symbol)
func (_Erc721Mintable *Erc721MintableCallerSession) Symbol() (string, error) {
	return _Erc721Mintable.Contract.Symbol(&_Erc721Mintable.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x6914db60.
//
// Solidity: function tokenMetadata(uint256 _tokenId) constant returns(string _infoUrl)
func (_Erc721Mintable *Erc721MintableCaller) TokenMetadata(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Erc721Mintable.contract.Call(opts, out, "tokenMetadata", _tokenId)
	return *ret0, err
}

// TokenMetadata is a free data retrieval call binding the contract method 0x6914db60.
//
// Solidity: function tokenMetadata(uint256 _tokenId) constant returns(string _infoUrl)
func (_Erc721Mintable *Erc721MintableSession) TokenMetadata(_tokenId *big.Int) (string, error) {
	return _Erc721Mintable.Contract.TokenMetadata(&_Erc721Mintable.CallOpts, _tokenId)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x6914db60.
//
// Solidity: function tokenMetadata(uint256 _tokenId) constant returns(string _infoUrl)
func (_Erc721Mintable *Erc721MintableCallerSession) TokenMetadata(_tokenId *big.Int) (string, error) {
	return _Erc721Mintable.Contract.TokenMetadata(&_Erc721Mintable.CallOpts, _tokenId)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256 _tokenId)
func (_Erc721Mintable *Erc721MintableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Erc721Mintable.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256 _tokenId)
func (_Erc721Mintable *Erc721MintableSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Erc721Mintable.Contract.TokenOfOwnerByIndex(&_Erc721Mintable.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) constant returns(uint256 _tokenId)
func (_Erc721Mintable *Erc721MintableCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Erc721Mintable.Contract.TokenOfOwnerByIndex(&_Erc721Mintable.CallOpts, _owner, _index)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 _totalSupply)
func (_Erc721Mintable *Erc721MintableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Erc721Mintable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 _totalSupply)
func (_Erc721Mintable *Erc721MintableSession) TotalSupply() (*big.Int, error) {
	return _Erc721Mintable.Contract.TotalSupply(&_Erc721Mintable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 _totalSupply)
func (_Erc721Mintable *Erc721MintableCallerSession) TotalSupply() (*big.Int, error) {
	return _Erc721Mintable.Contract.TotalSupply(&_Erc721Mintable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_Erc721Mintable *Erc721MintableTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Mintable.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_Erc721Mintable *Erc721MintableSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Mintable.Contract.Approve(&_Erc721Mintable.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_Erc721Mintable *Erc721MintableTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Mintable.Contract.Approve(&_Erc721Mintable.TransactOpts, _to, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x84ba2642.
//
// Solidity: function mint(address _owner, uint256 _tokenId, address _approvedAddress, string _metadata) returns()
func (_Erc721Mintable *Erc721MintableTransactor) Mint(opts *bind.TransactOpts, _owner common.Address, _tokenId *big.Int, _approvedAddress common.Address, _metadata string) (*types.Transaction, error) {
	return _Erc721Mintable.contract.Transact(opts, "mint", _owner, _tokenId, _approvedAddress, _metadata)
}

// Mint is a paid mutator transaction binding the contract method 0x84ba2642.
//
// Solidity: function mint(address _owner, uint256 _tokenId, address _approvedAddress, string _metadata) returns()
func (_Erc721Mintable *Erc721MintableSession) Mint(_owner common.Address, _tokenId *big.Int, _approvedAddress common.Address, _metadata string) (*types.Transaction, error) {
	return _Erc721Mintable.Contract.Mint(&_Erc721Mintable.TransactOpts, _owner, _tokenId, _approvedAddress, _metadata)
}

// Mint is a paid mutator transaction binding the contract method 0x84ba2642.
//
// Solidity: function mint(address _owner, uint256 _tokenId, address _approvedAddress, string _metadata) returns()
func (_Erc721Mintable *Erc721MintableTransactorSession) Mint(_owner common.Address, _tokenId *big.Int, _approvedAddress common.Address, _metadata string) (*types.Transaction, error) {
	return _Erc721Mintable.Contract.Mint(&_Erc721Mintable.TransactOpts, _owner, _tokenId, _approvedAddress, _metadata)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_Erc721Mintable *Erc721MintableTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Mintable.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_Erc721Mintable *Erc721MintableSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Mintable.Contract.Transfer(&_Erc721Mintable.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_Erc721Mintable *Erc721MintableTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Mintable.Contract.Transfer(&_Erc721Mintable.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Erc721Mintable *Erc721MintableTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Mintable.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Erc721Mintable *Erc721MintableSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Mintable.Contract.TransferFrom(&_Erc721Mintable.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Erc721Mintable *Erc721MintableTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Mintable.Contract.TransferFrom(&_Erc721Mintable.TransactOpts, _from, _to, _tokenId)
}

// Erc721MintableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc721Mintable contract.
type Erc721MintableApprovalIterator struct {
	Event *Erc721MintableApproval // Event containing the contract specifics and raw log

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
func (it *Erc721MintableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721MintableApproval)
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
		it.Event = new(Erc721MintableApproval)
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
func (it *Erc721MintableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721MintableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721MintableApproval represents a Approval event raised by the Erc721Mintable contract.
type Erc721MintableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 _tokenId)
func (_Erc721Mintable *Erc721MintableFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*Erc721MintableApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _Erc721Mintable.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &Erc721MintableApprovalIterator{contract: _Erc721Mintable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 _tokenId)
func (_Erc721Mintable *Erc721MintableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc721MintableApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _Erc721Mintable.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721MintableApproval)
				if err := _Erc721Mintable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 _tokenId)
func (_Erc721Mintable *Erc721MintableFilterer) ParseApproval(log types.Log) (*Erc721MintableApproval, error) {
	event := new(Erc721MintableApproval)
	if err := _Erc721Mintable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Erc721MintableMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Erc721Mintable contract.
type Erc721MintableMintIterator struct {
	Event *Erc721MintableMint // Event containing the contract specifics and raw log

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
func (it *Erc721MintableMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721MintableMint)
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
		it.Event = new(Erc721MintableMint)
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
func (it *Erc721MintableMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721MintableMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721MintableMint represents a Mint event raised by the Erc721Mintable contract.
type Erc721MintableMint struct {
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId)
func (_Erc721Mintable *Erc721MintableFilterer) FilterMint(opts *bind.FilterOpts, _to []common.Address, _tokenId []*big.Int) (*Erc721MintableMintIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Erc721Mintable.contract.FilterLogs(opts, "Mint", _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Erc721MintableMintIterator{contract: _Erc721Mintable.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId)
func (_Erc721Mintable *Erc721MintableFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *Erc721MintableMint, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Erc721Mintable.contract.WatchLogs(opts, "Mint", _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721MintableMint)
				if err := _Erc721Mintable.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId)
func (_Erc721Mintable *Erc721MintableFilterer) ParseMint(log types.Log) (*Erc721MintableMint, error) {
	event := new(Erc721MintableMint)
	if err := _Erc721Mintable.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Erc721MintableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc721Mintable contract.
type Erc721MintableTransferIterator struct {
	Event *Erc721MintableTransfer // Event containing the contract specifics and raw log

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
func (it *Erc721MintableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721MintableTransfer)
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
		it.Event = new(Erc721MintableTransfer)
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
func (it *Erc721MintableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721MintableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721MintableTransfer represents a Transfer event raised by the Erc721Mintable contract.
type Erc721MintableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _tokenId)
func (_Erc721Mintable *Erc721MintableFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*Erc721MintableTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Erc721Mintable.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &Erc721MintableTransferIterator{contract: _Erc721Mintable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _tokenId)
func (_Erc721Mintable *Erc721MintableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc721MintableTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Erc721Mintable.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721MintableTransfer)
				if err := _Erc721Mintable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _tokenId)
func (_Erc721Mintable *Erc721MintableFilterer) ParseTransfer(log types.Log) (*Erc721MintableTransfer, error) {
	event := new(Erc721MintableTransfer)
	if err := _Erc721Mintable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
