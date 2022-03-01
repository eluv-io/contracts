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

var UniqueEvents map[string]*events.EventInfo
var EventsByType = make(map[reflect.Type]*events.EventInfo)
var EventsByID = make(map[common.Hash]*events.EventInfo)

func init() {
	res := make(map[string]*events.EventInfo)
	resId := make(map[common.Hash]*events.EventInfo)

	addEvents := func(
		events map[string]*events.EventInfo,
		eventSuffix string) {

		for name, event := range events {
			ev, ok := res[name]
			if !ok {
				res[name] = event
				resId[event.ID] = event
				continue
			}
			if ev.ID == event.ID || resId[event.ID] != nil {
				continue
			}
			res[fmt.Sprintf("%s_%s", name, eventSuffix)] = event
			resId[event.ID] = event
		}
	}

	addEvents(contracts.UniqueEvents, "")
	addEvents(elv_tradable.UniqueEvents, "tradable")
	addEvents(c202008.UniqueEvents, "20200803")
	addEvents(c202002.UniqueEvents, "20200206")
	addEvents(c201903.UniqueEvents, "20190331")

	UniqueEvents = res
	EventsByID = resId
	for name, event := range res {
		ev, _ := EventsByType[event.Type]
		if ev != nil {
			panic(fmt.Sprintf("duplicate event: %s with type %v for %s and type %v",
				ev.Name, ev.Type, name, event.Type))
		}
		EventsByType[event.Type] = event
	}
	_ = EventsByID
}
