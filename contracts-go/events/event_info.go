package events

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// EventInfo gather information about a 'unique event'.
type EventInfo struct {
	Name   string                                    // name of the event as in abi.Event
	ID     common.Hash                               // ID of the event
	Type   reflect.Type                              // type of the struct event
	Unpack func(log types.Log, ev interface{}) error // unpack the given log into the given event
}

func (ev *EventInfo) Value(log types.Log) (reflect.Value, error) {
	event := reflect.New(ev.Type.Elem())
	err := ev.Unpack(log, event.Interface())
	if err != nil {
		return reflect.Value{}, err
	}
	f := event.Elem().FieldByName("Raw")
	if f.IsValid() && f.CanSet() {
		f.Set(reflect.ValueOf(log))
	}
	return event, nil
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
		"type": fmt.Sprintf("%T", ev.Type),
	})
}
