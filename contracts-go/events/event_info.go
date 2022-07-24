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

func (ev *EventInfo) HasType(typ EventType) bool {
	for _, t := range ev.Types {
		if t == typ {
			return true
		}
	}
	return false
}

// Value returns a value constructed from the type of the event and filled via
// the bound contract's UnpackLog function.
// The optional typ parameter allows to disambiguate the type to select. This is
// useful in cases where the ID of an EventInfo refers to multiples EventType
// where each of the bound contract would be able to unpack the log with no error.
func (ev *EventInfo) Value(log types.Log, typ ...reflect.Type) (res reflect.Value, err error) {
	var ty reflect.Type
	if len(typ) > 0 {
		ty = typ[0]
	}
	for _, u := range ev.Types {
		if ty != nil && ty != u.Type {
			continue
		}
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
