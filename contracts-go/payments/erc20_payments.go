// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package payments

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"strings"

	c "github.com/eluv-io/contracts/contracts-go/events"

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Map of ABI names to *abi.ABI
// ABI names are constants starting with K_
var ParsedABIS = map[string]*abi.ABI{}

// Map of ABI names to *bind.BoundContract for log parsing only
// ABI names are constants starting with K_
var BoundContracts = map[string]*bind.BoundContract{}

// Map of Unique events names to *EventInfo.
// Unique events names are constants starting with E_
var UniqueEvents = map[string]*EventInfo{}

// Map of Unique events types to *EventInfo
var EventsByType = map[reflect.Type]*EventInfo{}

// Map of Unique events IDs to *EventInfo
var EventsByID = map[common.Hash]*EventInfo{}

// JSON returns a parsed ABI interface and error if it failed.
func JSON(reader io.Reader) (*abi.ABI, error) {
	dec := json.NewDecoder(reader)

	var anAbi abi.ABI
	if err := dec.Decode(&anAbi); err != nil {
		return nil, err
	}

	return &anAbi, nil
}

func parseABI(name string) (*abi.ABI, error) {
	sabi := ABIS[name]
	if sabi == "" {
		return nil, fmt.Errorf("no such ABI %s", name)
	}
	return JSON(strings.NewReader(sabi))
}

func ParsedABI(name string) (*abi.ABI, error) {
	pabi, ok := ParsedABIS[name]
	if ok {
		return pabi, nil
	}
	return parseABI(name)
}

// ERC20PaymentsPayment is an auto generated low-level Go binding around an user-defined struct.
type ERC20PaymentsPayment struct {
	PaymentId [16]byte
	OracleId  common.Address
}

func BoundContract(name string) *bind.BoundContract {
	bc, ok := BoundContracts[name]
	if !ok {
		anABI, err := ParsedABI(name)
		if err != nil {
			panic(err)
		}
		bc = bind.NewBoundContract(common.Address{}, *anABI, nil, nil, nil)
		BoundContracts[name] = bc
	}
	return bc
}

// Type names of contract binding
const (
	K_Payments = "Payments"
)

var ABIS = map[string]string{

	K_Payments: PaymentsABI,
}

// Unique events names.
// Unique events are events whose ID and name are unique across contracts.
const (
	E_Canceled = "Canceled"
	E_Claimed  = "Claimed"
	E_Created  = "Created"
)

type EventInfo = c.EventInfo
type EventType = c.EventType

func init() {
	for name, _ := range ABIS {
		a, err := parseABI(name)
		if err == nil {
			ParsedABIS[name] = a
		}
	}
	var ev *EventInfo

	ev = &EventInfo{
		Name: "Canceled",
		ID:   common.HexToHash("0x25ccef20460325159f2f1853d4e2f53d50e508519fd950a1b57ca5f4f92ea264"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*Canceled)(nil)),
				BoundContract: BoundContract(K_Payments),
			},
		},
	}
	UniqueEvents[E_Canceled] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "Claimed",
		ID:   common.HexToHash("0x51630b38d4d2aa8f785cae61e1d06989553fa9b886c9f286106e331a8ff2e57a"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*Claimed)(nil)),
				BoundContract: BoundContract(K_Payments),
			},
		},
	}
	UniqueEvents[E_Claimed] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "Created",
		ID:   common.HexToHash("0x1a5cbdfaa174b8c077f5e32871877ef08cbf1cdba62a1ca568f500124c7a2411"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*Created)(nil)),
				BoundContract: BoundContract(K_Payments),
			},
		},
	}
	UniqueEvents[E_Created] = ev
	EventsByType[ev.Types[0].Type] = ev

}

// Unique events structs

// Canceled event with ID 0x25ccef20460325159f2f1853d4e2f53d50e508519fd950a1b57ca5f4f92ea264
type Canceled struct {
	PaymentId [16]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// Claimed event with ID 0x51630b38d4d2aa8f785cae61e1d06989553fa9b886c9f286106e331a8ff2e57a
type Claimed struct {
	PaymentId [16]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// Created event with ID 0x1a5cbdfaa174b8c077f5e32871877ef08cbf1cdba62a1ca568f500124c7a2411
type Created struct {
	PaymentId [16]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// PaymentsABI is the input ABI used to generate the binding from.
const PaymentsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"paymentId\",\"type\":\"bytes16\"}],\"name\":\"Canceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"paymentId\",\"type\":\"bytes16\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"paymentId\",\"type\":\"bytes16\"}],\"name\":\"Created\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"calculateTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"_paymentId\",\"type\":\"bytes16\"}],\"name\":\"cancelPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"_paymentId\",\"type\":\"bytes16\"}],\"name\":\"claimPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_receivers\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"paymentId\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"oracleId\",\"type\":\"address\"}],\"internalType\":\"structERC20Payments.Payment\",\"name\":\"_payment\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"createPayment\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"paymentId\",\"type\":\"bytes16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"_paymentId\",\"type\":\"bytes16\"}],\"name\":\"getPayment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"enumERC20Payments.PaymentState\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPaymentIdAt\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPaymentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PaymentsBin is the compiled bytecode used for deploying new contracts.
var PaymentsBin = "0x608060405234801561001057600080fd5b50611465806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638894c9fb1161005b5780638894c9fb146100e95780638e2aa1f3146100fb578063e2f5f99814610120578063f2c9f6031461013357600080fd5b806333c53d8b14610082578063732760a1146100b3578063820b71ca146100c6575b600080fd5b610095610090366004611065565b610146565b6040516001600160801b031990911681526020015b60405180910390f35b6100956100c13660046110e6565b610185565b6100d96100d4366004611186565b61055f565b60405190151581526020016100aa565b6001545b6040519081526020016100aa565b61010e610109366004611186565b6108ee565b6040516100aa969594939291906111ef565b6100ed61012e3660046112be565b610a62565b6100d9610141366004611186565b610aae565b60006001828154811061015b5761015b61137c565b90600052602060002090600291828204019190066010029054906101000a900460801b9050919050565b60006101946020860186611186565b61019d81610dc2565b156101ef5760405162461bcd60e51b815260206004820152601860248201527f7061796d656e74496420616c726561647920657869737473000000000000000060448201526064015b60405180910390fd5b6101fc6020870187611186565b9150600061023e8585808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508a9250610de9915050565b6040516323b872dd60e01b8152336004820152306024820152604481018290529091506001600160a01b038716906323b872dd906064016020604051808303816000875af1158015610294573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102b89190611392565b61030f5760405162461bcd60e51b815260206004820152602260248201527f7472616e7366657246726f6d2073656e64657220746f2074686973206661696c604482015261195960f21b60648201526084016101e6565b6040518060c00160405280336001600160a01b031681526020018a8a808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152505050908252506001600160a01b0388166020808301919091526040805188830281810184018352898252919093019291899189918291908501908490808284376000920191909152505050908252506020908101906103be9060408b01908b016113b4565b6001600160a01b03168152602001600090526001600160801b03198416600090815260208181526040909120825181546001600160a01b0319166001600160a01b03909116178155828201518051919261042092600185019290910190610fb0565b5060408201516002820180546001600160a01b0319166001600160a01b0390921691909117905560608201518051610462916003840191602090910190611015565b5060808201516004820180546001600160a01b039092166001600160a01b031983168117825560a0850151926001600160a81b03191617600160a01b8360028111156104b0576104b06111b7565b02179055505060018054808201825560008281527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6600283040180546fffffffffffffffffffffffffffffffff939094166010026101000a92830219909316608088901c92909202919091179091556040516001600160801b0319861692507f1a5cbdfaa174b8c077f5e32871877ef08cbf1cdba62a1ca568f500124c7a24119190a250509695505050505050565b60008161056b81610dc2565b6105b25760405162461bcd60e51b81526020600482015260186024820152771c185e5b595b9d125908191bd95cc81b9bdd08195e1a5cdd60421b60448201526064016101e6565b8260016001600160801b03198216600090815260208190526040902060040154600160a01b900460ff1660028111156105ed576105ed6111b7565b0361063a5760405162461bcd60e51b815260206004820152601a60248201527f636c61696d61626c653a20616c726561647920636c61696d656400000000000060448201526064016101e6565b60026001600160801b03198216600090815260208190526040902060040154600160a01b900460ff166002811115610674576106746111b7565b036106c15760405162461bcd60e51b815260206004820152601b60248201527f636c61696d61626c653a20616c726561647920726566756e646564000000000060448201526064016101e6565b6001600160801b0319841660009081526020819052604090206004015484906001600160a01b031633146107495760405162461bcd60e51b815260206004820152602960248201527f6f6e6c794f7261636c653a206d6573736167652073656e646572206d757374206044820152686265206f7261636c6560b81b60648201526084016101e6565b6001600160801b031985166000908152602081905260408120905b60018201548110156108975760028201546001830180546001600160a01b039092169163a9059cbb91908490811061079e5761079e61137c565b6000918252602090912001546003850180546001600160a01b0390921691859081106107cc576107cc61137c565b6000918252602090912001546040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af1158015610823573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108479190611392565b6108855760405162461bcd60e51b815260206004820152600f60248201526e1d1c985b9cd9995c8819985a5b1959608a1b60448201526064016101e6565b8061088f816113e5565b915050610764565b5060048101805460ff60a01b1916600160a01b1790556040516001600160801b03198716907f51630b38d4d2aa8f785cae61e1d06989553fa9b886c9f286106e331a8ff2e57a90600090a250600195945050505050565b600060606000606060008061090287610dc2565b15156000036109535760405162461bcd60e51b815260206004820152601760248201527f636f6e747261637420646f6573206e6f7420657869737400000000000000000060448201526064016101e6565b6001600160801b031987166000908152602081815260409182902080546002820154600483015460018401805487518188028101880190985280885294966001600160a01b0394851696919593851694600389019490841693600160a01b900460ff16928791908301828280156109f357602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116109d5575b5050505050945082805480602002602001604051908101604052809291908181526020018280548015610a4557602002820191906000526020600020905b815481526020019060010190808311610a31575b505050505092509650965096509650965096505091939550919395565b6000805b8251811015610aa857828181518110610a8157610a8161137c565b602002602001015182610a9491906113fe565b915080610aa0816113e5565b915050610a66565b50919050565b600081610aba81610dc2565b610b015760405162461bcd60e51b81526020600482015260186024820152771c185e5b595b9d125908191bd95cc81b9bdd08195e1a5cdd60421b60448201526064016101e6565b6001600160801b0319831660009081526020819052604090206004015483906001600160a01b03163314610b705760405162461bcd60e51b8152602060048201526016602482015275726566756e6461626c653a206e6f74206f7261636c6560501b60448201526064016101e6565b60026001600160801b03198216600090815260208190526040902060040154600160a01b900460ff166002811115610baa57610baa6111b7565b03610bf75760405162461bcd60e51b815260206004820152601c60248201527f726566756e6461626c653a20616c726561647920726566756e6465640000000060448201526064016101e6565b60016001600160801b03198216600090815260208190526040902060040154600160a01b900460ff166002811115610c3157610c316111b7565b03610c7e5760405162461bcd60e51b815260206004820152601b60248201527f726566756e6461626c653a20616c726561647920636c61696d6564000000000060448201526064016101e6565b6001600160801b0319841660009081526020818152604080832060048101805460ff60a01b1916600160a11b179055600381018054835181860281018601909452808452919493610d04939290830182828015610cfa57602002820191906000526020600020905b815481526020019060010190808311610ce6575b5050505050610a62565b6002830154835460405163a9059cbb60e01b81526001600160a01b03918216600482015260248101849052929350169063a9059cbb906044016020604051808303816000875af1158015610d5c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d809190611392565b506040516001600160801b03198716907f25ccef20460325159f2f1853d4e2f53d50e508519fd950a1b57ca5f4f92ea26490600090a250600195945050505050565b6001600160801b0319166000908152602081905260409020546001600160a01b0316151590565b600080835111610e315760405162461bcd60e51b81526020600482015260136024820152721b9bc8185b5bdd5b9d1cc81c1c9bdd9a591959606a1b60448201526064016101e6565b506000805b8351811015610ee1576000848281518110610e5357610e5361137c565b602002602001015111610ea85760405162461bcd60e51b815260206004820152601d60248201527f616d6f756e74206d7573742062652067726561746572207468616e203000000060448201526064016101e6565b838181518110610eba57610eba61137c565b602002602001015182610ecd91906113fe565b915080610ed9816113e5565b915050610e36565b50604051636eb1769f60e11b815233600482015230602482015281906001600160a01b0384169063dd62ed3e90604401602060405180830381865afa158015610f2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f529190611416565b1015610faa5760405162461bcd60e51b815260206004820152602160248201527f616c6c6f77616e6365206d757374206265203e3d2073756d28616d6f756e74736044820152602960f81b60648201526084016101e6565b92915050565b828054828255906000526020600020908101928215611005579160200282015b8281111561100557825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610fd0565b50611011929150611050565b5090565b828054828255906000526020600020908101928215611005579160200282015b82811115611005578251825591602001919060010190611035565b5b808211156110115760008155600101611051565b60006020828403121561107757600080fd5b5035919050565b60008083601f84011261109057600080fd5b50813567ffffffffffffffff8111156110a857600080fd5b6020830191508360208260051b85010111156110c357600080fd5b9250929050565b80356001600160a01b03811681146110e157600080fd5b919050565b60008060008060008086880360a081121561110057600080fd5b873567ffffffffffffffff8082111561111857600080fd5b6111248b838c0161107e565b90995097508791506040601f198401121561113e57600080fd5b60208a01965061115060608b016110ca565b955060808a013592508083111561116657600080fd5b505061117489828a0161107e565b979a9699509497509295939492505050565b60006020828403121561119857600080fd5b81356001600160801b0319811681146111b057600080fd5b9392505050565b634e487b7160e01b600052602160045260246000fd5b600381106111eb57634e487b7160e01b600052602160045260246000fd5b9052565b6001600160a01b03878116825260c0602080840182905288519184018290526000928982019290919060e0860190855b8181101561123d57855185168352948301949183019160010161121f565b5050918916604086015284820360608601528751808352918101925080880160005b8381101561127b5781518552938201939082019060010161125f565b5050506001600160a01b038616608085015250905061129d60a08301846111cd565b979650505050505050565b634e487b7160e01b600052604160045260246000fd5b600060208083850312156112d157600080fd5b823567ffffffffffffffff808211156112e957600080fd5b818501915085601f8301126112fd57600080fd5b81358181111561130f5761130f6112a8565b8060051b604051601f19603f83011681018181108582111715611334576113346112a8565b60405291825284820192508381018501918883111561135257600080fd5b938501935b8285101561137057843584529385019392850192611357565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156113a457600080fd5b815180151581146111b057600080fd5b6000602082840312156113c657600080fd5b6111b0826110ca565b634e487b7160e01b600052601160045260246000fd5b6000600182016113f7576113f76113cf565b5060010190565b60008219821115611411576114116113cf565b500190565b60006020828403121561142857600080fd5b505191905056fea264697066735822122074e53b79e9983cd69bce362323011d32fb89769b3fb8984ee4c0d7fbb8ee549664736f6c634300080d0033"

// DeployPayments deploys a new Ethereum contract, binding an instance of Payments to it.
func DeployPayments(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Payments, error) {
	parsed, err := ParsedABI(K_Payments)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PaymentsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Payments{PaymentsCaller: PaymentsCaller{contract: contract}, PaymentsTransactor: PaymentsTransactor{contract: contract}, PaymentsFilterer: PaymentsFilterer{contract: contract}}, nil
}

// Payments is an auto generated Go binding around an Ethereum contract.
type Payments struct {
	PaymentsCaller     // Read-only binding to the contract
	PaymentsTransactor // Write-only binding to the contract
	PaymentsFilterer   // Log filterer for contract events
}

// PaymentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewPayments creates a new instance of Payments, bound to a specific deployed contract.
func NewPayments(address common.Address, backend bind.ContractBackend) (*Payments, error) {
	contract, err := bindPayments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Payments{PaymentsCaller: PaymentsCaller{contract: contract}, PaymentsTransactor: PaymentsTransactor{contract: contract}, PaymentsFilterer: PaymentsFilterer{contract: contract}}, nil
}

// NewPaymentsCaller creates a new read-only instance of Payments, bound to a specific deployed contract.
func NewPaymentsCaller(address common.Address, caller bind.ContractCaller) (*PaymentsCaller, error) {
	contract, err := bindPayments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentsCaller{contract: contract}, nil
}

// NewPaymentsTransactor creates a new write-only instance of Payments, bound to a specific deployed contract.
func NewPaymentsTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentsTransactor, error) {
	contract, err := bindPayments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentsTransactor{contract: contract}, nil
}

// NewPaymentsFilterer creates a new log filterer instance of Payments, bound to a specific deployed contract.
func NewPaymentsFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentsFilterer, error) {
	contract, err := bindPayments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentsFilterer{contract: contract}, nil
}

// bindPayments binds a generic wrapper to an already deployed contract.
func bindPayments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_Payments)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// CalculateTotal is a free data retrieval call binding the contract method 0xe2f5f998.
//
// Solidity: function calculateTotal(uint256[] _amounts) pure returns(uint256 total)
func (_Payments *PaymentsCaller) CalculateTotal(opts *bind.CallOpts, _amounts []*big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Payments.contract.Call(opts, out, "calculateTotal", _amounts)
	return *ret0, err
}

// GetPayment is a free data retrieval call binding the contract method 0x8e2aa1f3.
//
// Solidity: function getPayment(bytes16 _paymentId) view returns(address sender, address[] receivers, address tokenContract, uint256[] amounts, address oracle, uint8 state)
func (_Payments *PaymentsCaller) GetPayment(opts *bind.CallOpts, _paymentId [16]byte) (struct {
	Sender        common.Address
	Receivers     []common.Address
	TokenContract common.Address
	Amounts       []*big.Int
	Oracle        common.Address
	State         uint8
}, error) {
	ret := new(struct {
		Sender        common.Address
		Receivers     []common.Address
		TokenContract common.Address
		Amounts       []*big.Int
		Oracle        common.Address
		State         uint8
	})
	out := ret
	err := _Payments.contract.Call(opts, out, "getPayment", _paymentId)
	return *ret, err
}

// GetPaymentIdAt is a free data retrieval call binding the contract method 0x33c53d8b.
//
// Solidity: function getPaymentIdAt(uint256 index) view returns(bytes16)
func (_Payments *PaymentsCaller) GetPaymentIdAt(opts *bind.CallOpts, index *big.Int) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _Payments.contract.Call(opts, out, "getPaymentIdAt", index)
	return *ret0, err
}

// TotalPaymentId is a free data retrieval call binding the contract method 0x8894c9fb.
//
// Solidity: function totalPaymentId() view returns(uint256)
func (_Payments *PaymentsCaller) TotalPaymentId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Payments.contract.Call(opts, out, "totalPaymentId")
	return *ret0, err
}

// CancelPayment is a paid mutator transaction binding the contract method 0xf2c9f603.
//
// Solidity: function cancelPayment(bytes16 _paymentId) returns(bool)
func (_Payments *PaymentsTransactor) CancelPayment(opts *bind.TransactOpts, _paymentId [16]byte) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "cancelPayment", _paymentId)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0x820b71ca.
//
// Solidity: function claimPayment(bytes16 _paymentId) returns(bool)
func (_Payments *PaymentsTransactor) ClaimPayment(opts *bind.TransactOpts, _paymentId [16]byte) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "claimPayment", _paymentId)
}

// CreatePayment is a paid mutator transaction binding the contract method 0x732760a1.
//
// Solidity: function createPayment(address[] _receivers, (bytes16,address) _payment, address _tokenContract, uint256[] _amounts) returns(bytes16 paymentId)
func (_Payments *PaymentsTransactor) CreatePayment(opts *bind.TransactOpts, _receivers []common.Address, _payment ERC20PaymentsPayment, _tokenContract common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "createPayment", _receivers, _payment, _tokenContract, _amounts)
}

// PaymentsCanceledIterator is returned from FilterCanceled and is used to iterate over the raw logs and unpacked data for Canceled events raised by the Payments contract.
type PaymentsCanceledIterator struct {
	Event *PaymentsCanceled // Event containing the contract specifics and raw log

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
func (it *PaymentsCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentsCanceled)
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
		it.Event = new(PaymentsCanceled)
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
func (it *PaymentsCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentsCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentsCanceled represents a Canceled event raised by the Payments contract.
type PaymentsCanceled struct {
	PaymentId [16]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCanceled is a free log retrieval operation binding the contract event 0x25ccef20460325159f2f1853d4e2f53d50e508519fd950a1b57ca5f4f92ea264.
//
// Solidity: event Canceled(bytes16 indexed paymentId)
func (_Payments *PaymentsFilterer) FilterCanceled(opts *bind.FilterOpts, paymentId [][16]byte) (*PaymentsCanceledIterator, error) {

	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _Payments.contract.FilterLogs(opts, "Canceled", paymentIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentsCanceledIterator{contract: _Payments.contract, event: "Canceled", logs: logs, sub: sub}, nil
}

// WatchCanceled is a free log subscription operation binding the contract event 0x25ccef20460325159f2f1853d4e2f53d50e508519fd950a1b57ca5f4f92ea264.
//
// Solidity: event Canceled(bytes16 indexed paymentId)
func (_Payments *PaymentsFilterer) WatchCanceled(opts *bind.WatchOpts, sink chan<- *PaymentsCanceled, paymentId [][16]byte) (event.Subscription, error) {

	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _Payments.contract.WatchLogs(opts, "Canceled", paymentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentsCanceled)
				if err := _Payments.contract.UnpackLog(event, "Canceled", log); err != nil {
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

// ParseCanceled is a log parse operation binding the contract event 0x25ccef20460325159f2f1853d4e2f53d50e508519fd950a1b57ca5f4f92ea264.
//
// Solidity: event Canceled(bytes16 indexed paymentId)
func (_Payments *PaymentsFilterer) ParseCanceled(log types.Log) (*PaymentsCanceled, error) {
	event := new(PaymentsCanceled)
	if err := _Payments.contract.UnpackLog(event, "Canceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentsClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Payments contract.
type PaymentsClaimedIterator struct {
	Event *PaymentsClaimed // Event containing the contract specifics and raw log

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
func (it *PaymentsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentsClaimed)
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
		it.Event = new(PaymentsClaimed)
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
func (it *PaymentsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentsClaimed represents a Claimed event raised by the Payments contract.
type PaymentsClaimed struct {
	PaymentId [16]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x51630b38d4d2aa8f785cae61e1d06989553fa9b886c9f286106e331a8ff2e57a.
//
// Solidity: event Claimed(bytes16 indexed paymentId)
func (_Payments *PaymentsFilterer) FilterClaimed(opts *bind.FilterOpts, paymentId [][16]byte) (*PaymentsClaimedIterator, error) {

	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _Payments.contract.FilterLogs(opts, "Claimed", paymentIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentsClaimedIterator{contract: _Payments.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x51630b38d4d2aa8f785cae61e1d06989553fa9b886c9f286106e331a8ff2e57a.
//
// Solidity: event Claimed(bytes16 indexed paymentId)
func (_Payments *PaymentsFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *PaymentsClaimed, paymentId [][16]byte) (event.Subscription, error) {

	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _Payments.contract.WatchLogs(opts, "Claimed", paymentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentsClaimed)
				if err := _Payments.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x51630b38d4d2aa8f785cae61e1d06989553fa9b886c9f286106e331a8ff2e57a.
//
// Solidity: event Claimed(bytes16 indexed paymentId)
func (_Payments *PaymentsFilterer) ParseClaimed(log types.Log) (*PaymentsClaimed, error) {
	event := new(PaymentsClaimed)
	if err := _Payments.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentsCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the Payments contract.
type PaymentsCreatedIterator struct {
	Event *PaymentsCreated // Event containing the contract specifics and raw log

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
func (it *PaymentsCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentsCreated)
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
		it.Event = new(PaymentsCreated)
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
func (it *PaymentsCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentsCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentsCreated represents a Created event raised by the Payments contract.
type PaymentsCreated struct {
	PaymentId [16]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0x1a5cbdfaa174b8c077f5e32871877ef08cbf1cdba62a1ca568f500124c7a2411.
//
// Solidity: event Created(bytes16 indexed paymentId)
func (_Payments *PaymentsFilterer) FilterCreated(opts *bind.FilterOpts, paymentId [][16]byte) (*PaymentsCreatedIterator, error) {

	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _Payments.contract.FilterLogs(opts, "Created", paymentIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentsCreatedIterator{contract: _Payments.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0x1a5cbdfaa174b8c077f5e32871877ef08cbf1cdba62a1ca568f500124c7a2411.
//
// Solidity: event Created(bytes16 indexed paymentId)
func (_Payments *PaymentsFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *PaymentsCreated, paymentId [][16]byte) (event.Subscription, error) {

	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _Payments.contract.WatchLogs(opts, "Created", paymentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentsCreated)
				if err := _Payments.contract.UnpackLog(event, "Created", log); err != nil {
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

// ParseCreated is a log parse operation binding the contract event 0x1a5cbdfaa174b8c077f5e32871877ef08cbf1cdba62a1ca568f500124c7a2411.
//
// Solidity: event Created(bytes16 indexed paymentId)
func (_Payments *PaymentsFilterer) ParseCreated(log types.Log) (*PaymentsCreated, error) {
	event := new(PaymentsCreated)
	if err := _Payments.contract.UnpackLog(event, "Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
