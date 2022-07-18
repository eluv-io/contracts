package events

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// EventInfo gather information about a 'unique event'.
type EventInfo struct {
	Name  string      // name of the event as in abi.Event
	ID    common.Hash // ID of the event
	Types []EventType // Normally a single item unless otherwise identical events differ in terms of "indexed" arguments.
}

func (ev *EventInfo) Value(log types.Log) (res reflect.Value, err error) {
	for _, u := range ev.Types {
		event := reflect.New(u.Type.Elem())
		err = u.BoundContract.UnpackLog(event.Interface(), ev.Name, log)
		if err != nil {
			continue
		}
		f := event.Elem().FieldByName("Raw")
		if f.IsValid() && f.CanSet() {
			f.Set(reflect.ValueOf(log))
		}
		return event, nil
	}
	return
}

func (ev *EventInfo) Event(log types.Log) (interface{}, error) {
	val, err := ev.Value(log)
	if err != nil {
		return nil, err
	}
	return val.Interface(), nil
}

func (ev *EventInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"name": ev.Name,
		"id":   ev.ID.String(),
		"type": ev.Types,
	})
}

// EventType contains the Go type of an event and a minimal bound contract used to unpack event logs.
type EventType struct {
	Type          reflect.Type // type of the struct event
	BoundContract *bind.BoundContract
}

func (ev *EventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type": fmt.Sprintf("%v", ev.Type),
	})
}