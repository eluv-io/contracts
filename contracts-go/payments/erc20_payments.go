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
		ID:   common.HexToHash("0xf5b628ae04a2e801c4184fb9b99c711c49643bbf8e1c6a0a70c10e3f5c32c0e2"),
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

// Created event with ID 0xf5b628ae04a2e801c4184fb9b99c711c49643bbf8e1c6a0a70c10e3f5c32c0e2
type Created struct {
	PaymentId     [16]byte
	Sender        common.Address
	Receivers     []common.Address
	TokenContract common.Address
	Amounts       []*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// PaymentsABI is the input ABI used to generate the binding from.
const PaymentsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"paymentId\",\"type\":\"bytes16\"}],\"name\":\"Canceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"paymentId\",\"type\":\"bytes16\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"paymentId\",\"type\":\"bytes16\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Created\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"calculateTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"_paymentId\",\"type\":\"bytes16\"}],\"name\":\"cancelPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"_paymentId\",\"type\":\"bytes16\"}],\"name\":\"claimPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_receivers\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"paymentId\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"oracleId\",\"type\":\"address\"}],\"internalType\":\"structERC20Payments.Payment\",\"name\":\"_payment\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"createPayment\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"paymentId\",\"type\":\"bytes16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"_paymentId\",\"type\":\"bytes16\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"enumERC20Payments.PaymentState\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PaymentsBin is the compiled bytecode used for deploying new contracts.
var PaymentsBin = "0x608060405234801561001057600080fd5b5061147d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063732760a11461005c578063820b71ca1461008d5780639c6173be146100b0578063e2f5f998146100d5578063f2c9f603146100f6575b600080fd5b61006f61006a366004611071565b610109565b6040516001600160801b031990911681526020015b60405180910390f35b6100a061009b366004611111565b610503565b6040519015158152602001610084565b6100c36100be366004611111565b610892565b6040516100849695949392919061117a565b6100e86100e3366004611249565b610a06565b604051908152602001610084565b6100a0610104366004611111565b610a52565b60006101186020860186611111565b61012181610d66565b1561016e5760405162461bcd60e51b81526020600482015260186024820152777061796d656e74496420616c72656164792065786973747360401b60448201526064015b60405180910390fd5b61017b6020870187611111565b915060006101bd8585808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508a9250610d8d915050565b90506101c883610d66565b156102105760405162461bcd60e51b81526020600482015260186024820152777061796d656e74496420616c72656164792065786973747360401b6044820152606401610165565b6040516323b872dd60e01b8152336004820152306024820152604481018290526001600160a01b038716906323b872dd906064016020604051808303816000875af1158015610263573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102879190611307565b6102de5760405162461bcd60e51b815260206004820152602260248201527f7472616e7366657246726f6d2073656e64657220746f2074686973206661696c604482015261195960f21b6064820152608401610165565b6040518060c00160405280336001600160a01b031681526020018a8a808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152505050908252506001600160a01b03881660208083019190915260408051888302818101840183528982529190930192918991899182919085019084908082843760009201919091525050509082525060209081019061038d9060408b01908b01611329565b6001600160a01b03168152602001600090526001600160801b03198416600090815260208181526040909120825181546001600160a01b0319166001600160a01b0390911617815582820151805191926103ef92600185019290910190610f54565b5060408201516002820180546001600160a01b0319166001600160a01b0390921691909117905560608201518051610431916003840191602090910190610fb9565b5060808201516004820180546001600160a01b039092166001600160a01b031983168117825560a0850151926001600160a81b03191617600160a01b83600281111561047f5761047f611142565b02179055505060405161049691508a908a90611344565b6040518091039020336001600160a01b0316846fffffffffffffffffffffffffffffffff19167ff5b628ae04a2e801c4184fb9b99c711c49643bbf8e1c6a0a70c10e3f5c32c0e28989896040516104ef93929190611384565b60405180910390a450509695505050505050565b60008161050f81610d66565b6105565760405162461bcd60e51b81526020600482015260186024820152771c185e5b595b9d125908191bd95cc81b9bdd08195e1a5cdd60421b6044820152606401610165565b8260016001600160801b03198216600090815260208190526040902060040154600160a01b900460ff16600281111561059157610591611142565b036105de5760405162461bcd60e51b815260206004820152601a60248201527f636c61696d61626c653a20616c726561647920636c61696d65640000000000006044820152606401610165565b60026001600160801b03198216600090815260208190526040902060040154600160a01b900460ff16600281111561061857610618611142565b036106655760405162461bcd60e51b815260206004820152601b60248201527f636c61696d61626c653a20616c726561647920726566756e64656400000000006044820152606401610165565b6001600160801b0319841660009081526020819052604090206004015484906001600160a01b031633146106ed5760405162461bcd60e51b815260206004820152602960248201527f6f6e6c794f7261636c653a206d6573736167652073656e646572206d757374206044820152686265206f7261636c6560b81b6064820152608401610165565b6001600160801b031985166000908152602081905260408120905b600182015481101561083b5760028201546001830180546001600160a01b039092169163a9059cbb919084908110610742576107426113d1565b6000918252602090912001546003850180546001600160a01b039092169185908110610770576107706113d1565b6000918252602090912001546040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af11580156107c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107eb9190611307565b6108295760405162461bcd60e51b815260206004820152600f60248201526e1d1c985b9cd9995c8819985a5b1959608a1b6044820152606401610165565b80610833816113fd565b915050610708565b5060048101805460ff60a01b1916600160a01b1790556040516001600160801b03198716907f51630b38d4d2aa8f785cae61e1d06989553fa9b886c9f286106e331a8ff2e57a90600090a250600195945050505050565b60006060600060606000806108a687610d66565b15156000036108f75760405162461bcd60e51b815260206004820152601760248201527f436f6e747261637420646f6573206e6f742065786973740000000000000000006044820152606401610165565b6001600160801b031987166000908152602081815260409182902080546002820154600483015460018401805487518188028101880190985280885294966001600160a01b0394851696919593851694600389019490841693600160a01b900460ff169287919083018282801561099757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610979575b50505050509450828054806020026020016040519081016040528092919081815260200182805480156109e957602002820191906000526020600020905b8154815260200190600101908083116109d5575b505050505092509650965096509650965096505091939550919395565b6000805b8251811015610a4c57828181518110610a2557610a256113d1565b602002602001015182610a389190611416565b915080610a44816113fd565b915050610a0a565b50919050565b600081610a5e81610d66565b610aa55760405162461bcd60e51b81526020600482015260186024820152771c185e5b595b9d125908191bd95cc81b9bdd08195e1a5cdd60421b6044820152606401610165565b6001600160801b0319831660009081526020819052604090206004015483906001600160a01b03163314610b145760405162461bcd60e51b8152602060048201526016602482015275726566756e6461626c653a206e6f74206f7261636c6560501b6044820152606401610165565b60026001600160801b03198216600090815260208190526040902060040154600160a01b900460ff166002811115610b4e57610b4e611142565b03610b9b5760405162461bcd60e51b815260206004820152601c60248201527f726566756e6461626c653a20616c726561647920726566756e646564000000006044820152606401610165565b60016001600160801b03198216600090815260208190526040902060040154600160a01b900460ff166002811115610bd557610bd5611142565b03610c225760405162461bcd60e51b815260206004820152601b60248201527f726566756e6461626c653a20616c726561647920636c61696d656400000000006044820152606401610165565b6001600160801b0319841660009081526020818152604080832060048101805460ff60a01b1916600160a11b179055600381018054835181860281018601909452808452919493610ca8939290830182828015610c9e57602002820191906000526020600020905b815481526020019060010190808311610c8a575b5050505050610a06565b6002830154835460405163a9059cbb60e01b81526001600160a01b03918216600482015260248101849052929350169063a9059cbb906044016020604051808303816000875af1158015610d00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d249190611307565b506040516001600160801b03198716907f25ccef20460325159f2f1853d4e2f53d50e508519fd950a1b57ca5f4f92ea26490600090a250600195945050505050565b6001600160801b0319166000908152602081905260409020546001600160a01b0316151590565b600080835111610dd55760405162461bcd60e51b81526020600482015260136024820152721b9bc8185b5bdd5b9d1cc81c1c9bdd9a591959606a1b6044820152606401610165565b506000805b8351811015610e85576000848281518110610df757610df76113d1565b602002602001015111610e4c5760405162461bcd60e51b815260206004820152601d60248201527f616d6f756e74206d7573742062652067726561746572207468616e20300000006044820152606401610165565b838181518110610e5e57610e5e6113d1565b602002602001015182610e719190611416565b915080610e7d816113fd565b915050610dda565b50604051636eb1769f60e11b815233600482015230602482015281906001600160a01b0384169063dd62ed3e90604401602060405180830381865afa158015610ed2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ef6919061142e565b1015610f4e5760405162461bcd60e51b815260206004820152602160248201527f616c6c6f77616e6365206d757374206265203e3d2073756d28616d6f756e74736044820152602960f81b6064820152608401610165565b92915050565b828054828255906000526020600020908101928215610fa9579160200282015b82811115610fa957825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610f74565b50610fb5929150610ff4565b5090565b828054828255906000526020600020908101928215610fa9579160200282015b82811115610fa9578251825591602001919060010190610fd9565b5b80821115610fb55760008155600101610ff5565b60008083601f84011261101b57600080fd5b50813567ffffffffffffffff81111561103357600080fd5b6020830191508360208260051b850101111561104e57600080fd5b9250929050565b80356001600160a01b038116811461106c57600080fd5b919050565b60008060008060008086880360a081121561108b57600080fd5b873567ffffffffffffffff808211156110a357600080fd5b6110af8b838c01611009565b90995097508791506040601f19840112156110c957600080fd5b60208a0196506110db60608b01611055565b955060808a01359250808311156110f157600080fd5b50506110ff89828a01611009565b979a9699509497509295939492505050565b60006020828403121561112357600080fd5b81356001600160801b03198116811461113b57600080fd5b9392505050565b634e487b7160e01b600052602160045260246000fd5b6003811061117657634e487b7160e01b600052602160045260246000fd5b9052565b6001600160a01b03878116825260c0602080840182905288519184018290526000928982019290919060e0860190855b818110156111c85785518516835294830194918301916001016111aa565b5050918916604086015284820360608601528751808352918101925080880160005b83811015611206578151855293820193908201906001016111ea565b5050506001600160a01b038616608085015250905061122860a0830184611158565b979650505050505050565b634e487b7160e01b600052604160045260246000fd5b6000602080838503121561125c57600080fd5b823567ffffffffffffffff8082111561127457600080fd5b818501915085601f83011261128857600080fd5b81358181111561129a5761129a611233565b8060051b604051601f19603f830116810181811085821117156112bf576112bf611233565b6040529182528482019250838101850191888311156112dd57600080fd5b938501935b828510156112fb578435845293850193928501926112e2565b98975050505050505050565b60006020828403121561131957600080fd5b8151801515811461113b57600080fd5b60006020828403121561133b57600080fd5b61113b82611055565b60008184825b85811015611379576001600160a01b0361136383611055565b168352602092830192919091019060010161134a565b509095945050505050565b6001600160a01b0384168152604060208201819052810182905260006001600160fb1b038311156113b457600080fd5b8260051b8085606085013760009201606001918252509392505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820161140f5761140f6113e7565b5060010190565b60008219821115611429576114296113e7565b500190565b60006020828403121561144057600080fd5b505191905056fea26469706673582212204664c8d9c7e85290d2d53d16687acb407fab88c5d3d8a07a3476f78d5246dcc264736f6c634300080d0033"

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

// GetContract is a free data retrieval call binding the contract method 0x9c6173be.
//
// Solidity: function getContract(bytes16 _paymentId) view returns(address sender, address[] receivers, address tokenContract, uint256[] amounts, address oracle, uint8 state)
func (_Payments *PaymentsCaller) GetContract(opts *bind.CallOpts, _paymentId [16]byte) (struct {
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
	err := _Payments.contract.Call(opts, out, "getContract", _paymentId)
	return *ret, err
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
	PaymentId     [16]byte
	Sender        common.Address
	Receivers     []common.Address
	TokenContract common.Address
	Amounts       []*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0xf5b628ae04a2e801c4184fb9b99c711c49643bbf8e1c6a0a70c10e3f5c32c0e2.
//
// Solidity: event Created(bytes16 indexed paymentId, address indexed sender, address[] indexed receivers, address tokenContract, uint256[] amounts)
func (_Payments *PaymentsFilterer) FilterCreated(opts *bind.FilterOpts, paymentId [][16]byte, sender []common.Address, receivers [][]common.Address) (*PaymentsCreatedIterator, error) {

	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiversRule []interface{}
	for _, receiversItem := range receivers {
		receiversRule = append(receiversRule, receiversItem)
	}

	logs, sub, err := _Payments.contract.FilterLogs(opts, "Created", paymentIdRule, senderRule, receiversRule)
	if err != nil {
		return nil, err
	}
	return &PaymentsCreatedIterator{contract: _Payments.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0xf5b628ae04a2e801c4184fb9b99c711c49643bbf8e1c6a0a70c10e3f5c32c0e2.
//
// Solidity: event Created(bytes16 indexed paymentId, address indexed sender, address[] indexed receivers, address tokenContract, uint256[] amounts)
func (_Payments *PaymentsFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *PaymentsCreated, paymentId [][16]byte, sender []common.Address, receivers [][]common.Address) (event.Subscription, error) {

	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiversRule []interface{}
	for _, receiversItem := range receivers {
		receiversRule = append(receiversRule, receiversItem)
	}

	logs, sub, err := _Payments.contract.WatchLogs(opts, "Created", paymentIdRule, senderRule, receiversRule)
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

// ParseCreated is a log parse operation binding the contract event 0xf5b628ae04a2e801c4184fb9b99c711c49643bbf8e1c6a0a70c10e3f5c32c0e2.
//
// Solidity: event Created(bytes16 indexed paymentId, address indexed sender, address[] indexed receivers, address tokenContract, uint256[] amounts)
func (_Payments *PaymentsFilterer) ParseCreated(log types.Log) (*PaymentsCreated, error) {
	event := new(PaymentsCreated)
	if err := _Payments.contract.UnpackLog(event, "Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
