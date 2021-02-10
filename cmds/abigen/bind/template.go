// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package bind

import "github.com/ethereum/go-ethereum/accounts/abi"

// tmplData is the data structure required to fill the binding template.
type tmplData struct {
	Package   string                   // Name of the package to place the generated file in
	Contracts map[string]*tmplContract // List of contracts to generate into this file
	Libraries map[string]string        // Map the bytecode's link pattern to the library name
	Structs   map[string]*tmplStruct   // Contract struct type definitions
	Events    []*tmplEvent             // list of events with unique name and ID
}

// tmplContract contains the data needed to generate an individual contract binding.
type tmplContract struct {
	Type        string                 // Type name of the main contract binding
	InputABI    string                 // JSON ABI used as the input to generate the binding from
	InputBin    string                 // Optional EVM bytecode used to denetare deploy code from
	FuncSigs    map[string]string      // Optional map: string signature -> 4-byte signature
	Constructor abi.Method             // Contract constructor for deploy parametrization
	Calls       map[string]*tmplMethod // Contract calls that only read state data
	Transacts   map[string]*tmplMethod // Contract calls that write state data
	Events      map[string]*tmplEvent  // Contract events accessors
	Libraries   map[string]string      // Same as tmplData, but filtered to only keep what the contract needs
	Library     bool                   // Indicator whether the contract is a library
}

// tmplMethod is a wrapper around an abi.Method that contains a few preprocessed
// and cached data fields.
type tmplMethod struct {
	Original   abi.Method // Original method as parsed by the abi package
	Normalized abi.Method // Normalized version of the parsed method (capitalized names, non-anonymous args/returns)
	Structured bool       // Whether the returns should be accumulated into a struct
}

// tmplEvent is a wrapper around an abi.Event
type tmplEvent struct {
	KType      string    // Type name of a contract binding where the event was found
	Original   abi.Event // Original event as parsed by the abi package
	Normalized abi.Event // Normalized version of the parsed fields
}

type tmplEventList []*tmplEvent

func (l tmplEventList) Len() int {
	return len(l)
}

func (l tmplEventList) Less(i, j int) bool {
	return l[i].Normalized.Name < l[j].Normalized.Name
}

func (l tmplEventList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// tmplField is a wrapper around a struct field with binding language
// struct type definition and relative filed name.
type tmplField struct {
	Type    string   // Field type representation depends on target binding language
	Name    string   // Field name converted from the raw user-defined field name
	SolKind abi.Type // Raw abi type information
}

// tmplStruct is a wrapper around an abi.tuple contains a auto-generated
// struct name.
type tmplStruct struct {
	Name   string       // Auto-generated struct name(before solidity v0.5.11) or raw name.
	Fields []*tmplField // Struct fields definition depends on the binding language.
}

// tmplSource is language to template mapping containing all the supported
// programming languages the package can generate to.
var tmplSource = map[Lang]string{
	LangGo:   tmplSourceGo,
	LangJava: tmplSourceJava,
}

// tmplSourceGo is the Go source template use to generate the contract binding
// based on.
const tmplSourceGo = `
// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package {{.Package}}

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"strings"
	"reflect"

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

// Map of ABI names to *abi.ABI
// ABI names are constants starting with K_
var ParsedABIS map[string]*abi.ABI

// Map of Unique events names to *EventInfo. 
// Unique events names are constants starting with E_ 
var UniqueEvents = map[string]*EventInfo{}

// Map of Unique events types to *EventInfo 
var EventsByType = map[reflect.Type]*EventInfo{}

// Map of Unique events IDs to *EventInfo 
var EventsByID   = map[common.Hash]*EventInfo{}


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

{{$structs := .Structs}}
{{range $structs}}
	// {{.Name}} is an auto generated low-level Go binding around an user-defined struct.
	type {{.Name}} struct {
	{{range $field := .Fields}}
	{{$field.Name}} {{$field.Type}}{{end}}
	}
{{end}}

// Type names of contract binding
const (
{{range $contract := .Contracts}}
    K_{{.Type}} = "{{.Type}}"{{end}}
)

var ABIS = map[string]string{
	{{range $contract := .Contracts}}
	K_{{.Type}}: {{.Type}}ABI,{{end}}
}

// Unique events names.
// Unique events are events whose ID and name are unique across contracts. 
const (
	{{range $event := .Events}}
	E_{{.Normalized.Name}} = "{{.Normalized.Name}}"{{end}} 
)

// EventInfo gather information about a 'unique event'.
type EventInfo struct {
	ID     common.Hash                                // ID of the event 
	Type   reflect.Type                               // type of the struct event 
	Unpack func(log *types.Log, ev interface{}) error // unpack the given log into the given event
}

func init() {
	ParsedABIS = make(map[string]*abi.ABI)
	for name, _ := range ABIS {
		a, err := parseABI(name)
		if err == nil {
			ParsedABIS[name] = a
		}
	}
	var ev *EventInfo
	{{range $event := .Events}}
	ev = &EventInfo{
		ID:     common.HexToHash("{{eventId .Normalized}}"), 
		Type:   reflect.TypeOf((*{{.Normalized.Name}})(nil)),
		Unpack: func(log *types.Log, ev interface{}) error{
			anABI, _ := ParsedABI(K_{{.KType}})
			if err := anABI.Unpack(ev, E_{{.Normalized.Name}}, log.Data); err != nil {
				return err
			}
			return nil
		},
	}
	UniqueEvents[E_{{.Normalized.Name}}] = ev
	EventsByType[ev.Type] = ev
	EventsByID[ev.ID] = ev
	{{end}}
}

// Unique events structs
{{range $event := .Events}}
// {{.Normalized.Name}} represents a {{.Normalized.Name}} event.
type {{.Normalized.Name}} struct { {{range .Normalized.Inputs}}
	{{capitalise .Name}} {{if .Indexed}}{{bindtopictype .Type $structs}}{{else}}{{bindtype .Type $structs}}{{end}}; {{end}}
	Raw types.Log // Blockchain specific contextual infos
}
{{end}}

{{range $contract := .Contracts}}
	// {{.Type}}ABI is the input ABI used to generate the binding from.
	const {{.Type}}ABI = "{{.InputABI}}"

	{{if $contract.FuncSigs}}
		// {{.Type}}FuncSigs maps the 4-byte function signature to its string representation.
		var {{.Type}}FuncSigs = map[string]string{
			{{range $strsig, $binsig := .FuncSigs}}"{{$binsig}}": "{{$strsig}}",
			{{end}}
		}
	{{end}}

	{{if .InputBin}}
		// {{.Type}}Bin is the compiled bytecode used for deploying new contracts.
		var {{.Type}}Bin = "0x{{.InputBin}}"

		// Deploy{{.Type}} deploys a new Ethereum contract, binding an instance of {{.Type}} to it.
		func Deploy{{.Type}}(auth *bind.TransactOpts, backend bind.ContractBackend {{range .Constructor.Inputs}}, {{.Name}} {{bindtype .Type $structs}}{{end}}) (common.Address, *types.Transaction, *{{.Type}}, error) {
	      parsed, err := ParsedABI(K_{{.Type}})
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  {{range $pattern, $name := .Libraries}}
			{{decapitalise $name}}Addr, _, _, _ := Deploy{{capitalise $name}}(auth, backend)
			{{$contract.Type}}Bin = strings.Replace({{$contract.Type}}Bin, "__${{$pattern}}$__", {{decapitalise $name}}Addr.String()[2:], -1)
		  {{end}}
		  address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex({{.Type}}Bin), backend {{range .Constructor.Inputs}}, {{.Name}}{{end}})
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  return address, tx, &{{.Type}}{ {{.Type}}Caller: {{.Type}}Caller{contract: contract}, {{.Type}}Transactor: {{.Type}}Transactor{contract: contract}, {{.Type}}Filterer: {{.Type}}Filterer{contract: contract} }, nil
		}
	{{end}}

	// {{.Type}} is an auto generated Go binding around an Ethereum contract.
	type {{.Type}} struct {
	  {{.Type}}Caller     // Read-only binding to the contract
	  {{.Type}}Transactor // Write-only binding to the contract
	  {{.Type}}Filterer   // Log filterer for contract events
	}

	// {{.Type}}Caller is an auto generated read-only Go binding around an Ethereum contract.
	type {{.Type}}Caller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// {{.Type}}Transactor is an auto generated write-only Go binding around an Ethereum contract.
	type {{.Type}}Transactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// {{.Type}}Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type {{.Type}}Filterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// New{{.Type}} creates a new instance of {{.Type}}, bound to a specific deployed contract.
	func New{{.Type}}(address common.Address, backend bind.ContractBackend) (*{{.Type}}, error) {
	  contract, err := bind{{.Type}}(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &{{.Type}}{ {{.Type}}Caller: {{.Type}}Caller{contract: contract}, {{.Type}}Transactor: {{.Type}}Transactor{contract: contract}, {{.Type}}Filterer: {{.Type}}Filterer{contract: contract} }, nil
	}

	// New{{.Type}}Caller creates a new read-only instance of {{.Type}}, bound to a specific deployed contract.
	func New{{.Type}}Caller(address common.Address, caller bind.ContractCaller) (*{{.Type}}Caller, error) {
	  contract, err := bind{{.Type}}(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &{{.Type}}Caller{contract: contract}, nil
	}

	// New{{.Type}}Transactor creates a new write-only instance of {{.Type}}, bound to a specific deployed contract.
	func New{{.Type}}Transactor(address common.Address, transactor bind.ContractTransactor) (*{{.Type}}Transactor, error) {
	  contract, err := bind{{.Type}}(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &{{.Type}}Transactor{contract: contract}, nil
	}

	// New{{.Type}}Filterer creates a new log filterer instance of {{.Type}}, bound to a specific deployed contract.
 	func New{{.Type}}Filterer(address common.Address, filterer bind.ContractFilterer) (*{{.Type}}Filterer, error) {
 	  contract, err := bind{{.Type}}(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &{{.Type}}Filterer{contract: contract}, nil
 	}

	// bind{{.Type}} binds a generic wrapper to an already deployed contract.
	func bind{{.Type}}(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := ParsedABI(K_{{.Type}})
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
	}

	{{range .Calls}}
		// {{.Normalized.Name}} is a free data retrieval call binding the contract method 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{formatmethod .Original $structs}}
		func (_{{$contract.Type}} *{{$contract.Type}}Caller) {{.Normalized.Name}}(opts *bind.CallOpts {{range .Normalized.Inputs}}, {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
			{{if .Structured}}ret := new(struct{
				{{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}}
				{{end}}
			}){{else}}var (
				{{range $i, $_ := .Normalized.Outputs}}ret{{$i}} = new({{bindtype .Type $structs}})
				{{end}}
			){{end}}
			out := {{if .Structured}}ret{{else}}{{if eq (len .Normalized.Outputs) 1}}ret0{{else}}&[]interface{}{
				{{range $i, $_ := .Normalized.Outputs}}ret{{$i}},
				{{end}}
			}{{end}}{{end}}
			err := _{{$contract.Type}}.contract.Call(opts, out, "{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
			return {{if .Structured}}*ret,{{else}}{{range $i, $_ := .Normalized.Outputs}}*ret{{$i}},{{end}}{{end}} err
		}

	{{end}}

	{{range .Transacts}}
		// {{.Normalized.Name}} is a paid mutator transaction binding the contract method 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{formatmethod .Original $structs}}
		func (_{{$contract.Type}} *{{$contract.Type}}Transactor) {{.Normalized.Name}}(opts *bind.TransactOpts {{range .Normalized.Inputs}}, {{.Name}} {{bindtype .Type $structs}} {{end}}) (*types.Transaction, error) {
			return _{{$contract.Type}}.contract.Transact(opts, "{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
		}
	{{end}}

	{{range .Events}}
		// {{$contract.Type}}{{.Normalized.Name}}Iterator is returned from Filter{{.Normalized.Name}} and is used to iterate over the raw logs and unpacked data for {{.Normalized.Name}} events raised by the {{$contract.Type}} contract.
		type {{$contract.Type}}{{.Normalized.Name}}Iterator struct {
			Event *{{$contract.Type}}{{.Normalized.Name}} // Event containing the contract specifics and raw log

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
		func (it *{{$contract.Type}}{{.Normalized.Name}}Iterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new({{$contract.Type}}{{.Normalized.Name}})
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
				it.Event = new({{$contract.Type}}{{.Normalized.Name}})
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
		func (it *{{$contract.Type}}{{.Normalized.Name}}Iterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *{{$contract.Type}}{{.Normalized.Name}}Iterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// {{$contract.Type}}{{.Normalized.Name}} represents a {{.Normalized.Name}} event raised by the {{$contract.Type}} contract.
		type {{$contract.Type}}{{.Normalized.Name}} struct { {{range .Normalized.Inputs}}
			{{capitalise .Name}} {{if .Indexed}}{{bindtopictype .Type $structs}}{{else}}{{bindtype .Type $structs}}{{end}}; {{end}}
			Raw types.Log // Blockchain specific contextual infos
		}

		// Filter{{.Normalized.Name}} is a free log retrieval operation binding the contract event 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{formatevent .Original $structs}}
 		func (_{{$contract.Type}} *{{$contract.Type}}Filterer) Filter{{.Normalized.Name}}(opts *bind.FilterOpts{{range .Normalized.Inputs}}{{if .Indexed}}, {{.Name}} []{{bindtype .Type $structs}}{{end}}{{end}}) (*{{$contract.Type}}{{.Normalized.Name}}Iterator, error) {
			{{range .Normalized.Inputs}}
			{{if .Indexed}}var {{.Name}}Rule []interface{}
			for _, {{.Name}}Item := range {{.Name}} {
				{{.Name}}Rule = append({{.Name}}Rule, {{.Name}}Item)
			}{{end}}{{end}}

			logs, sub, err := _{{$contract.Type}}.contract.FilterLogs(opts, "{{.Original.Name}}"{{range .Normalized.Inputs}}{{if .Indexed}}, {{.Name}}Rule{{end}}{{end}})
			if err != nil {
				return nil, err
			}
			return &{{$contract.Type}}{{.Normalized.Name}}Iterator{contract: _{{$contract.Type}}.contract, event: "{{.Original.Name}}", logs: logs, sub: sub}, nil
 		}

		// Watch{{.Normalized.Name}} is a free log subscription operation binding the contract event 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{formatevent .Original $structs}}
		func (_{{$contract.Type}} *{{$contract.Type}}Filterer) Watch{{.Normalized.Name}}(opts *bind.WatchOpts, sink chan<- *{{$contract.Type}}{{.Normalized.Name}}{{range .Normalized.Inputs}}{{if .Indexed}}, {{.Name}} []{{bindtype .Type $structs}}{{end}}{{end}}) (event.Subscription, error) {
			{{range .Normalized.Inputs}}
			{{if .Indexed}}var {{.Name}}Rule []interface{}
			for _, {{.Name}}Item := range {{.Name}} {
				{{.Name}}Rule = append({{.Name}}Rule, {{.Name}}Item)
			}{{end}}{{end}}

			logs, sub, err := _{{$contract.Type}}.contract.WatchLogs(opts, "{{.Original.Name}}"{{range .Normalized.Inputs}}{{if .Indexed}}, {{.Name}}Rule{{end}}{{end}})
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new({{$contract.Type}}{{.Normalized.Name}})
						if err := _{{$contract.Type}}.contract.UnpackLog(event, "{{.Original.Name}}", log); err != nil {
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

		// Parse{{.Normalized.Name}} is a log parse operation binding the contract event 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{formatevent .Original $structs}}
		func (_{{$contract.Type}} *{{$contract.Type}}Filterer) Parse{{.Normalized.Name}}(log types.Log) (*{{$contract.Type}}{{.Normalized.Name}}, error) {
			event := new({{$contract.Type}}{{.Normalized.Name}})
			if err := _{{$contract.Type}}.contract.UnpackLog(event, "{{.Original.Name}}", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	{{end}}
{{end}}
`

// tmplSourceJava is the Java source template use to generate the contract binding
// based on.
const tmplSourceJava = `
// This file is an automatically generated Java binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package {{.Package}};

import org.ethereum.geth.*;
import java.util.*;

{{$structs := .Structs}}
{{range $contract := .Contracts}}
{{if not .Library}}public {{end}}class {{.Type}} {
	// ABI is the input ABI used to generate the binding from.
	public final static String ABI = "{{.InputABI}}";
	{{if $contract.FuncSigs}}
		// {{.Type}}FuncSigs maps the 4-byte function signature to its string representation.
		public final static Map<String, String> {{.Type}}FuncSigs;
		static {
			Hashtable<String, String> temp = new Hashtable<String, String>();
			{{range $strsig, $binsig := .FuncSigs}}temp.put("{{$binsig}}", "{{$strsig}}");
			{{end}}
			{{.Type}}FuncSigs = Collections.unmodifiableMap(temp);
		}
	{{end}}
	{{if .InputBin}}
	// BYTECODE is the compiled bytecode used for deploying new contracts.
	public final static String BYTECODE = "0x{{.InputBin}}";

	// deploy deploys a new Ethereum contract, binding an instance of {{.Type}} to it.
	public static {{.Type}} deploy(TransactOpts auth, EthereumClient client{{range .Constructor.Inputs}}, {{bindtype .Type $structs}} {{.Name}}{{end}}) throws Exception {
		Interfaces args = Geth.newInterfaces({{(len .Constructor.Inputs)}});
		String bytecode = BYTECODE;
		{{if .Libraries}}

		// "link" contract to dependent libraries by deploying them first.
		{{range $pattern, $name := .Libraries}}
		{{capitalise $name}} {{decapitalise $name}}Inst = {{capitalise $name}}.deploy(auth, client);
		bytecode = bytecode.replace("__${{$pattern}}$__", {{decapitalise $name}}Inst.Address.getHex().substring(2));
		{{end}}
		{{end}}
		{{range $index, $element := .Constructor.Inputs}}Interface arg{{$index}} = Geth.newInterface();arg{{$index}}.set{{namedtype (bindtype .Type $structs) .Type}}({{.Name}});args.set({{$index}},arg{{$index}});
		{{end}}
		return new {{.Type}}(Geth.deployContract(auth, ABI, Geth.decodeFromHex(bytecode), client, args));
	}

	// Internal constructor used by contract deployment.
	private {{.Type}}(BoundContract deployment) {
		this.Address  = deployment.getAddress();
		this.Deployer = deployment.getDeployer();
		this.Contract = deployment;
	}
	{{end}}

	// Ethereum address where this contract is located at.
	public final Address Address;

	// Ethereum transaction in which this contract was deployed (if known!).
	public final Transaction Deployer;

	// Contract instance bound to a blockchain address.
	private final BoundContract Contract;

	// Creates a new instance of {{.Type}}, bound to a specific deployed contract.
	public {{.Type}}(Address address, EthereumClient client) throws Exception {
		this(Geth.bindContract(address, ABI, client));
	}

	{{range .Calls}}
	{{if gt (len .Normalized.Outputs) 1}}
	// {{capitalise .Normalized.Name}}Results is the output of a call to {{.Normalized.Name}}.
	public class {{capitalise .Normalized.Name}}Results {
		{{range $index, $item := .Normalized.Outputs}}public {{bindtype .Type $structs}} {{if ne .Name ""}}{{.Name}}{{else}}Return{{$index}}{{end}};
		{{end}}
	}
	{{end}}

	// {{.Normalized.Name}} is a free data retrieval call binding the contract method 0x{{printf "%x" .Original.ID}}.
	//
	// Solidity: {{.Original.String}}
	public {{if gt (len .Normalized.Outputs) 1}}{{capitalise .Normalized.Name}}Results{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}}{{end}}{{end}} {{.Normalized.Name}}(CallOpts opts{{range .Normalized.Inputs}}, {{bindtype .Type $structs}} {{.Name}}{{end}}) throws Exception {
		Interfaces args = Geth.newInterfaces({{(len .Normalized.Inputs)}});
		{{range $index, $item := .Normalized.Inputs}}Interface arg{{$index}} = Geth.newInterface();arg{{$index}}.set{{namedtype (bindtype .Type $structs) .Type}}({{.Name}});args.set({{$index}},arg{{$index}});
		{{end}}

		Interfaces results = Geth.newInterfaces({{(len .Normalized.Outputs)}});
		{{range $index, $item := .Normalized.Outputs}}Interface result{{$index}} = Geth.newInterface(); result{{$index}}.setDefault{{namedtype (bindtype .Type $structs) .Type}}(); results.set({{$index}}, result{{$index}});
		{{end}}

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "{{.Original.Name}}", args);
		{{if gt (len .Normalized.Outputs) 1}}
			{{capitalise .Normalized.Name}}Results result = new {{capitalise .Normalized.Name}}Results();
			{{range $index, $item := .Normalized.Outputs}}result.{{if ne .Name ""}}{{.Name}}{{else}}Return{{$index}}{{end}} = results.get({{$index}}).get{{namedtype (bindtype .Type $structs) .Type}}();
			{{end}}
			return result;
		{{else}}{{range .Normalized.Outputs}}return results.get(0).get{{namedtype (bindtype .Type $structs) .Type}}();{{end}}
		{{end}}
	}
	{{end}}

	{{range .Transacts}}
	// {{.Normalized.Name}} is a paid mutator transaction binding the contract method 0x{{printf "%x" .Original.ID}}.
	//
	// Solidity: {{.Original.String}}
	public Transaction {{.Normalized.Name}}(TransactOpts opts{{range .Normalized.Inputs}}, {{bindtype .Type $structs}} {{.Name}}{{end}}) throws Exception {
		Interfaces args = Geth.newInterfaces({{(len .Normalized.Inputs)}});
		{{range $index, $item := .Normalized.Inputs}}Interface arg{{$index}} = Geth.newInterface();arg{{$index}}.set{{namedtype (bindtype .Type $structs) .Type}}({{.Name}});args.set({{$index}},arg{{$index}});
		{{end}}
		return this.Contract.transact(opts, "{{.Original.Name}}"	, args);
	}
	{{end}}
}
{{end}}
`
