package abidiff

import (
	"bytes"
	"fmt"
	"path/filepath"

	contracts "github.com/eluv-io/contracts/build"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/log"

	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
)

type ContractInfo struct {
	Name string
	ABI  string
}

// if someone knows how to find all the constants in a package, let me know...
var checkContracts = []ContractInfo{
	{"BaseContentSpace", contracts.BaseContentSpaceABI},
	{"BaseLibrary", contracts.BaseLibraryABI},
	{"BaseContent", contracts.BaseContentABI},
	{"BaseAccessControlGroup", contracts.BaseAccessControlGroupABI},
	{"BaseContentType", contracts.BaseContentABI},
	{"BaseAccessWallet", contracts.BaseAccessWalletABI},
}

type DiffItem struct {
	Breaking bool
	Report   string
}

func mergeCompareKeyLists(lkeys, rkeys []string, cmpFunc func(string), diffFunc func(key string, leftDiff bool)) {
	lidx := 0
	ridx := 0
	for {
		var cmp int
		ldone := lidx >= len(lkeys)
		rdone := ridx >= len(rkeys)
		if ldone && rdone {
			break
		} else if rdone {
			cmp = -1
		} else if ldone {
			cmp = 1
		} else {
			cmp = strings.Compare(lkeys[lidx], rkeys[ridx])
		}

		if cmp == 0 {
			cmpFunc(lkeys[lidx])
			lidx++
			ridx++
		} else if cmp < 0 {
			diffFunc(lkeys[lidx], true)
			lidx++
		} else {
			diffFunc(rkeys[ridx], false)
			ridx++
		}
	}
}

func VerifyAllABI(overwrite bool, storeDir string) ([]DiffItem, error) {

	allDiffs := make([]DiffItem, 0)

	for _, ci := range checkContracts {
		diffs, err := diffABIStore(ci.Name, ci.ABI, overwrite, storeDir)
		if err != nil {
			return nil, err
		}
		allDiffs = append(allDiffs, diffs...)
	}

	return allDiffs, nil
}

func diffABIStore(contractName, jsonCurrentABI string, overwrite bool, storeDir string) ([]DiffItem, error) {

	storeDirPath, err := filepath.Abs(storeDir)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(storeDirPath); err != nil {
		if os.IsNotExist(err) {
			log.Warn(fmt.Sprintf("store dir is not present, creating at %s", storeDirPath))
			os.MkdirAll(storeDirPath, os.ModePerm)
		}
	}

	storeFileName := path.Join(storeDirPath, fmt.Sprintf("%v.json", contractName))
	if _, err := os.Stat(storeFileName); err != nil {
		if os.IsNotExist(err) {
			return nil, ioutil.WriteFile(storeFileName, []byte(jsonCurrentABI), os.ModePerm)
		}
		return nil, err
	}

	// if overwrite is enabled, replace the existing files in store dir with new abi
	log.Debug("overwrite", "value", overwrite)
	if overwrite {
		return nil, ioutil.WriteFile(storeFileName, []byte(jsonCurrentABI), os.ModePerm)
	}

	storeABIStr, err := ioutil.ReadFile(storeFileName)
	if err != nil {
		return nil, err
	}

	storeABI, err := abi.JSON(strings.NewReader(string(storeABIStr)))
	if err != nil {
		return nil, err
	}

	currentABI, err := abi.JSON(strings.NewReader(jsonCurrentABI))
	if err != nil {
		return nil, err
	}

	return diffABIs(contractName, currentABI, storeABI)
}

func diffABIs(contractName string, currentABI, storeABI abi.ABI) ([]DiffItem, error) {

	diffs := make([]DiffItem, 0)

	getMethodKeys := func(theMap map[string]abi.Method) []string {
		keys := make([]string, 0)
		for k, _ := range theMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		return keys
	}

	currentMethodKeys := getMethodKeys(currentABI.Methods)
	storedMethodKeys := getMethodKeys(storeABI.Methods)

	cmpMethodFunc := func(methodKey string) {
		if bytes.Compare(currentABI.Methods[methodKey].ID(), storeABI.Methods[methodKey].ID()) != 0 {
			diffs = append(diffs, DiffItem{true, fmt.Sprintf("%v, current and stored METHOD signatures differ: %v", contractName, methodKey)})
		}
	}

	diffMethodFunc := func(methodKey string, currentDiff bool) {
		if currentDiff {
			diffs = append(diffs, DiffItem{false, fmt.Sprintf("%v, CURRENT contract has METHOD that is missing from stored: %v", contractName, methodKey)})
		} else {
			diffs = append(diffs, DiffItem{true, fmt.Sprintf("%v, STORED contract has METHOD that is missing from current: %v", contractName, methodKey)})
		}
	}

	mergeCompareKeyLists(currentMethodKeys, storedMethodKeys, cmpMethodFunc, diffMethodFunc)

	getEventKeys := func(theMap map[string]abi.Event) []string {
		keys := make([]string, 0)
		for k, _ := range theMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		return keys
	}

	currentEventKeys := getEventKeys(currentABI.Events)
	storedEventKeys := getEventKeys(storeABI.Events)

	cmpEventFunc := func(eventKey string) {
		if bytes.Compare(currentABI.Events[eventKey].ID().Bytes(), storeABI.Events[eventKey].ID().Bytes()) != 0 {
			diffs = append(diffs, DiffItem{true, fmt.Sprintf("%v, current and stored EVENT signatures differ: %v", contractName, eventKey)})
		}
	}

	diffEventFunc := func(eventKey string, currentDiff bool) {
		if currentDiff {
			diffs = append(diffs, DiffItem{false, fmt.Sprintf("%v, CURRENT contract has EVENT that is missing from stored: %v", contractName, eventKey)})
		} else {
			diffs = append(diffs, DiffItem{true, fmt.Sprintf("%v, STORED contract has EVENT that is missing from current: %v", contractName, eventKey)})
		}
	}

	mergeCompareKeyLists(currentEventKeys, storedEventKeys, cmpEventFunc, diffEventFunc)

	return diffs, nil
}
