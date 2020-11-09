// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const distDir = "../../dist"

type EventStruct struct {
	Event string `json:"event"`
	Topic string `json:"topic"`
}

func main() {
	files, err := ioutil.ReadDir(distDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		eventsMap := make(map[string]EventStruct)
		if strings.HasSuffix(f.Name(), ".abi") {
			jsonBytes, err := ioutil.ReadFile(path.Join(distDir, f.Name()))
			if err != nil {
				log.Fatal(err)
			}
			parsedABI, err := abi.JSON(bytes.NewReader(jsonBytes))
			if err != nil {
				log.Fatal(err)
			}

			for e, info := range parsedABI.Events {
				eventsMap[e] = EventStruct{Event: info.String(), Topic: info.ID().String()}
			}

			if len(eventsMap) > 0 {
				out, err := json.Marshal(eventsMap)
				if err != nil {
					log.Fatal(err)
				}

				eventsFileName := fmt.Sprintf("%v_events.json", strings.TrimSuffix(f.Name(), ".abi"))
				err = ioutil.WriteFile(path.Join(distDir, eventsFileName), out, 0755)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
