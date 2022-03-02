package contracts_go

import (
	"fmt"
	"reflect"

	"github.com/eluv-io/contracts/contracts-go/contracts"
	c201903 "github.com/eluv-io/contracts/contracts-go/contracts_20190331"
	c202002 "github.com/eluv-io/contracts/contracts-go/contracts_20200206"
	c202008 "github.com/eluv-io/contracts/contracts-go/contracts_20200803"
	"github.com/eluv-io/contracts/contracts-go/events"
	"github.com/eluv-io/contracts/contracts-go/tradable"
	"github.com/ethereum/go-ethereum/common"
)

var (
	// UniqueEvents contains all events keyed by their name (e.g. "AccessRequest"). Events that have different
	// signatures in different contract versions (sub-packages) are treated as separate events. Their most recent
	// version will get the event's original name, while older versions are suffixed with the version date of the
	// contract package, e.g. VersionConfirm_20200206.
	UniqueEvents = make(map[string]*events.EventInfo)
	// EventsByName is an alias of UniqueEvents
	EventsByName = UniqueEvents
	// EventsByType contains all events keyed by their go type.
	EventsByType = make(map[reflect.Type]*events.EventInfo)
	// EventsByID contains all events keyed by their id (topic hash)
	EventsByID = make(map[common.Hash]*events.EventInfo)
	// EventNamesByID contains all event names keyed by their id (topic hash)
	EventNamesByID = make(map[common.Hash]string)
)

func init() {
	addEvents := func(
		events map[string]*events.EventInfo,
		eventSuffix string) {

		for name, event := range events {
			ev, ok := UniqueEvents[name]
			if !ok {
				UniqueEvents[name] = event
				EventsByID[event.ID] = event
				continue
			}
			if ev.ID == event.ID || EventsByID[event.ID] != nil {
				continue
			}
			UniqueEvents[fmt.Sprintf("%s_%s", name, eventSuffix)] = event
			EventsByID[event.ID] = event
		}
	}

	addEvents(contracts.UniqueEvents, "")
	addEvents(elv_tradable.UniqueEvents, "tradable")
	addEvents(c202008.UniqueEvents, "20200803")
	addEvents(c202002.UniqueEvents, "20200206")
	addEvents(c201903.UniqueEvents, "20190331")

	for name, event := range UniqueEvents {
		ev, _ := EventsByType[event.Type]
		if ev != nil {
			panic(fmt.Sprintf("duplicate event: %s with type %v for %s and type %v",
				ev.Name, ev.Type, name, event.Type))
		}
		EventsByType[event.Type] = event
		EventNamesByID[event.ID] = event.Name
	}
}
