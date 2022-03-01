package contracts_go

import (
	"fmt"
	"sort"
	"testing"

	"github.com/eluv-io/contracts/contracts-go/contracts"
	c2019003 "github.com/eluv-io/contracts/contracts-go/contracts_20190331"
	c202002 "github.com/eluv-io/contracts/contracts-go/contracts_20200206"
	c202008 "github.com/eluv-io/contracts/contracts-go/contracts_20200803"
)

// TestDuplicateEvents does not look at tradable
func TestDuplicateEvents(t *testing.T) {
	ev201903 := c2019003.UniqueEvents
	ev202002 := c202002.UniqueEvents
	ev202008 := c202008.UniqueEvents
	evLast := contracts.UniqueEvents

	//trLast := elv_tradable.UniqueEvents
	//fmt.Println("",
	//	"201903", len(ev201903),
	//	"202002", len(ev202002),
	//	"202008", len(ev202008),
	//	"latest", len(evLast),
	//	"tradable", len(trLast),
	//	"all", len(UniqueEvents))
	/*
		201903: 26 only in this version
		202002:  3 duplicates of other versions
		202008:  1 only in this version (CreateExtUserWallet)
		latest: 81
		trad:   15
		all:   126
	*/

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

func TestUniqueEvents(t *testing.T) {
	names := []string{}
	for name, _ := range UniqueEvents {
		names = append(names, name)
	}
	sort.Strings(names)

	fm := func(s string, maxLen int) string {
		for len(s) < maxLen {
			s = s + " "
		}
		return s
	}

	fmt.Println(fmt.Sprintf("all events (%d)", len(UniqueEvents)))
	for _, name := range names {
		fmt.Println(fm(name, 35),
			fm(fmt.Sprintf("%v", UniqueEvents[name].Type), 50),
			UniqueEvents[name].ID.String())
	}

}
