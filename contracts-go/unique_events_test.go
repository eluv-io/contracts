package contracts_go

import (
	"testing"

	contracts "github.com/eluv-io/contracts/contracts-go/contracts"
	c2019003 "github.com/eluv-io/contracts/contracts-go/contracts_20190331"
	c202002 "github.com/eluv-io/contracts/contracts-go/contracts_20200206"
	c202008 "github.com/eluv-io/contracts/contracts-go/contracts_20200803"
)

func TestUniqueEvents(t *testing.T) {
	ev201903 := c2019003.UniqueEvents
	ev202002 := c202002.UniqueEvents
	ev202008 := c202008.UniqueEvents

	evLast := contracts.UniqueEvents

	getId := func(version, eventName string) (string, bool) {
		switch version {
		case "201903":
			if e, ok := ev201903[eventName]; ok {
				return e.ID.String(), true
			}
		case "202002":
			if e, ok := ev202002[eventName]; ok {
				return e.ID.String(), true
			}
		case "202008":
			if e, ok := ev202008[eventName]; ok {
				return e.ID.String(), true
			}
		}
		return "", false
	}

	for name, ev := range evLast {
		for _, version := range []string{"201903", "202002", "202008"} {
			if id, ok := getId(version, name); ok && id != ev.ID.String() {
				switch name {
				case "RunAccessCharge", "VersionConfirm":
					// RunAccessCharge and VersionConfirm are known to have changed
				default:
					t.Errorf("id mismatch - version %s, event %s", version, name)
				}
			}
		}
	}

}
