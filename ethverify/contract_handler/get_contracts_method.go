package getContract

import (
	"fmt"
	"math/big"
	"strings"
	"time"
	"unicode"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	contracts "github.com/eluv-io/contracts/build"
)

type getContractConfig struct {
	ethclient    *ethclient.Client
	contractAddr common.Address
	contractType string
	co           bind.CallOpts
}

func NewGetContractConfig(contractAddrStr, contractType, elvmasterdRPCUrl string) (*getContractConfig, error) {

	isAddr := common.IsHexAddress(contractAddrStr)
	if contractAddrStr == "" || contractAddrStr == "0x" || !isAddr {
		return nil, fmt.Errorf("contract_address provided is invalid")
	}

	ec, err := ethclient.Dial(elvmasterdRPCUrl)
	if err != nil {
		return nil, err
	}

	config := getContractConfig{}

	config.contractAddr = common.HexToAddress(contractAddrStr)
	config.ethclient = ec
	config.contractType = contractType
	config.co = bind.CallOpts{}

	return &config, nil
}

func (gc *getContractConfig) GetContract() (map[string]interface{}, error) {

	out := make(map[string]interface{})
	var err error
	switch strings.ToLower(gc.contractType) {
	case "ownable":
		out, err = gc.getOwnable()
	case "editable":
		out, err = gc.getEditable()
	case "container":
		out, err = gc.getContainer()
	case "nodespace":
		out, err = gc.getNodeSpace()
	case "userspace":
		out, err = gc.getUserSpace()
	case "accessindexor":
		out, err = gc.getAccessIndexor()
	case "accessible":
		out, err = gc.getAccessible()
	case "basecontentspace":
		out, err = gc.getBaseContentSpace()
	case "basecontenttype":
		out, err = gc.getBaseContentType()
	case "basecontent":
		out, err = gc.getBaseContent()
	case "baselibrary":
		out, err = gc.getBaseLibrary()
	case "baseaccesswallet":
	case "baseaccesscontrolgroup":
	default:
		return nil, fmt.Errorf("contract_type provided is INVALID : %v", gc.contractType)
	}
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (gc *getContractConfig) getOwnable() (map[string]interface{}, error) {

	ownableInst, err := contracts.NewOwnable(gc.contractAddr, gc.ethclient)
	if err != nil {
		return nil, err
	}

	owner, err := ownableInst.Owner(&gc.co)
	if err != nil {
		return nil, err
	}

	creator, err := ownableInst.Creator(&gc.co)
	if err != nil {
		return nil, err
	}

	versionInBytes32, err := ownableInst.Version(&gc.co)
	if err != nil {
		return nil, err
	}
	version := strings.TrimFunc(string(versionInBytes32[:]), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	spc, err := ownableInst.ContentSpace(&gc.co)
	if err != nil {
		return nil, err
	}

	ownable := map[string]interface{}{
		"owner":        owner,
		"creator":      creator,
		"version":      version,
		"contentSpace": spc,
	}

	return ownable, nil
}

func (gc *getContractConfig) getEditable() (map[string]interface{}, error) {

	editableInst, err := contracts.NewEditable(gc.contractAddr, gc.ethclient)
	if err != nil {
		return nil, err
	}

	parentAddr, err := editableInst.ParentAddress(&gc.co)
	if err != nil {
		return nil, err
	}

	pendingHash, err := editableInst.PendingHash(&gc.co)
	if err != nil {
		return nil, err
	}

	objHash, err := editableInst.ObjectHash(&gc.co)
	if err != nil {
		return nil, err
	}

	totalVersionHashes, err := editableInst.CountVersionHashes(&gc.co)
	if err != nil {
		return nil, err
	}

	versionHashes := make(map[string]time.Time)
	for i := int64(0); i < totalVersionHashes.Int64(); i++ {
		timestamp, err := editableInst.VersionTimestamp(&gc.co, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		unixTimeUTC := time.Unix(timestamp.Int64(), 0)
		version, err := editableInst.VersionHashes(&gc.co, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		versionHashes[version] = unixTimeUTC
	}

	commitPending, err := editableInst.CommitPending(&gc.co)
	if err != nil {
		return nil, err
	}

	canCommit, err := editableInst.CanCommit(&gc.co)
	if err != nil {
		return nil, err
	}

	canConfirm, err := editableInst.CanConfirm(&gc.co)
	if err != nil {
		return nil, err
	}

	canEdit, err := editableInst.CanEdit(&gc.co)
	if err != nil {
		return nil, err
	}

	ownable, err := gc.getOwnable()
	if err != nil {
		return nil, err
	}

	editable := map[string]interface{}{
		"ownable":            ownable,
		"parentAddress":      parentAddr,
		"objectHash":         objHash,
		"pendingHash":        pendingHash,
		"totalVersionhashes": totalVersionHashes,
		"versionHashes":      versionHashes,
		"commitPending":      commitPending,
		"canCommit":          canCommit,
		"canConfirm":         canConfirm,
		"canEdit":            canEdit,
	}

	return editable, nil
}

func (gc *getContractConfig) getContainer() (map[string]interface{}, error) {

	containerInst, err := contracts.NewContainer(gc.contractAddr, gc.ethclient)
	if err != nil {
		return nil, err
	}

	addressKMS, err := containerInst.AddressKMS(&gc.co)
	if err != nil {
		return nil, err
	}

	requiresReview, err := containerInst.RequiresReview(&gc.co)
	if err != nil {
		return nil, err
	}

	contentTypeLength, err := containerInst.ContentTypesLength(&gc.co)
	if err != nil {
		return nil, err
	}

	var contentTypes []string
	contentTypesContracts := make(map[string]string)
	for i := int64(0); i < contentTypeLength.Int64(); i++ {
		ct, err := containerInst.ContentTypes(&gc.co, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		contentTypes = append(contentTypes, ct.String())

		ctContract, err := containerInst.ContentTypeContracts(&gc.co, ct)
		if err != nil {
			return nil, err
		}

		contentTypesContracts[ct.String()] = ctContract.String()
	}

	ownable, err := gc.getOwnable()
	if err != nil {
		return nil, err
	}

	editable, err := gc.getEditable()
	if err != nil {
		return nil, err
	}

	container := map[string]interface{}{
		"addressKMS":            addressKMS.String(),
		"requiresReview":        requiresReview,
		"contentTypeLength":     contentTypeLength.String(),
		"contentTypes":          contentTypes,
		"contentTypesContracts": contentTypesContracts,
		"ownable":               ownable,
		"editable":              editable,
	}

	return container, nil
}

func (gc *getContractConfig) getNodeSpace() (map[string]interface{}, error) {
	ownable, err := gc.getOwnable()
	if err != nil {
		return nil, err
	}

	nodeSpaceInst, err := contracts.NewNodeSpace(gc.contractAddr, gc.ethclient)
	if err != nil {
		return nil, err
	}

	numActiveNodes, err := nodeSpaceInst.NumActiveNodes(&gc.co)
	if err != nil {
		return nil, err
	}
	numPendingNodes, err := nodeSpaceInst.NumPendingNodes(&gc.co)
	if err != nil {
		return nil, err
	}

	activeNodes := make(map[string]interface{})
	for i := int64(0); i < numActiveNodes.Int64(); i++ {
		activeNodeAddr, err := nodeSpaceInst.ActiveNodeAddresses(&gc.co, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		activeNodeLocatorsbytes, err := nodeSpaceInst.ActiveNodeLocators(&gc.co, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		activeNodes[activeNodeAddr.String()] = activeNodeLocatorsbytes
	}

	pendingNodes := make(map[string]interface{})
	for i := int64(0); i < numPendingNodes.Int64(); i++ {
		pendingNodeAddr, err := nodeSpaceInst.PendingNodeAddresses(&gc.co, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		pendingNodeLocatorsBytes, err := nodeSpaceInst.PendingNodeLocators(&gc.co, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		pendingNodes[pendingNodeAddr.String()] = pendingNodeLocatorsBytes
	}

	nodeSpace := map[string]interface{}{
		"ownable":         ownable,
		"numActiveNodes":  numActiveNodes,
		"numPendingNodes": numPendingNodes,
		"activeNodes":     activeNodes,
		"pendingNodes":    pendingNodes,
	}

	return nodeSpace, nil
}

func (gc *getContractConfig) getUserSpace() (map[string]interface{}, error) {

	userSpaceInst, err := contracts.NewUserSpace(gc.contractAddr, gc.ethclient)
	if err != nil {
		return nil, err
	}

	versionInBytes32, err := userSpaceInst.Version(&gc.co)
	if err != nil {
		return nil, err
	}
	version := strings.TrimFunc(string(versionInBytes32[:]), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	userWallet := map[string]interface{}{
		"version": version,
	}

	return userWallet, nil
}

func (gc *getContractConfig) getAccessIndexor() (map[string]interface{}, error) {

	accessType := map[string]uint8{
		"view":   0,
		"access": 1,
		"edit":   2,
	}

	categoryType := map[uint8]string{
		1: "content_object",
		2: "group",
		3: "library",
		4: "content_type",
		5: "contract",
	}

	ownable, err := gc.getOwnable()
	if err != nil {
		return nil, err
	}

	AcsIndexorInst, err := contracts.NewAccessIndexor(gc.contractAddr, gc.ethclient)
	if err != nil {
		return nil, err
	}

	getAccessIndexor := func(objLength int64, getObj func(int64) (common.Address, error), getObjRights func(common.Address, uint8) (bool, error)) (map[string]map[string]bool, error) {
		objLists := make(map[string]map[string]bool)
		for i := int64(0); i < objLength; i++ {
			obj, err := getObj(i)
			if err != nil {
				return nil, err
			}

			objLists[obj.String()] = make(map[string]bool)
			for k, v := range accessType {
				rights, err := getObjRights(obj, v)
				if err != nil {
					return nil, err
				}
				objLists[obj.String()][k] = rights
			}
		}

		return objLists, nil
	}

	// ----------------

	libLength, err := AcsIndexorInst.GetLibrariesLength(&gc.co)
	if err != nil {
		return nil, err
	}

	libStruct, err := AcsIndexorInst.Libraries(&gc.co)
	if err != nil {
		return nil, err
	}

	getLibrary := func(index int64) (common.Address, error) {
		lib, err := AcsIndexorInst.GetLibrary(&gc.co, big.NewInt(index))
		if err != nil {
			return common.Address{}, err
		}
		return lib, nil
	}

	getLibraryRigths := func(addr common.Address, v uint8) (bool, error) {
		rights, err := AcsIndexorInst.CheckLibraryRights(&gc.co, addr, v)
		if err != nil {
			return false, err
		}
		return rights, nil
	}

	librariesList, err := getAccessIndexor(libLength.Int64(), getLibrary, getLibraryRigths)
	if err != nil {
		return nil, err
	}

	librariesInfo := map[string]interface{}{
		"librariesLength":   libLength,
		"librariesCategory": categoryType[libStruct.Category],
		"librariesList":     librariesList,
	}

	// -------------

	contentObjLength, err := AcsIndexorInst.GetContentObjectsLength(&gc.co)
	if err != nil {
		return nil, err
	}

	contentObjStruct, err := AcsIndexorInst.ContentObjects(&gc.co)
	if err != nil {
		return nil, err
	}

	getContentObject := func(index int64) (common.Address, error) {
		cObj, err := AcsIndexorInst.GetContentObject(&gc.co, big.NewInt(index))
		if err != nil {
			return common.Address{}, err
		}
		return cObj, nil
	}

	getContentObjectRights := func(addr common.Address, v uint8) (bool, error) {
		rights, err := AcsIndexorInst.CheckContentObjectRights(&gc.co, addr, v)
		if err != nil {
			return false, err
		}
		return rights, nil
	}

	contentObjectList, err := getAccessIndexor(contentObjLength.Int64(), getContentObject, getContentObjectRights)
	if err != nil {
		return nil, err
	}

	contentObjectsInfo := map[string]interface{}{
		"contentObjectsLength":   contentObjLength,
		"contentObjectsCategory": categoryType[contentObjStruct.Category],
		"contentObjectsList":     contentObjectList,
	}

	// --------------

	accessGrpLength, err := AcsIndexorInst.GetAccessGroupsLength(&gc.co)
	if err != nil {
		return nil, err
	}

	accessGrpStruct, err := AcsIndexorInst.AccessGroups(&gc.co)
	if err != nil {
		return nil, err
	}

	getAccessGroup := func(index int64) (common.Address, error) {
		acsGrp, err := AcsIndexorInst.GetAccessGroup(&gc.co, big.NewInt(index))
		if err != nil {
			return common.Address{}, err
		}
		return acsGrp, nil
	}

	getAccessGroupsRights := func(addr common.Address, v uint8) (bool, error) {
		rights, err := AcsIndexorInst.CheckAccessGroupRights(&gc.co, addr, v)
		if err != nil {
			return false, err
		}
		return rights, nil
	}

	accessGroupsList, err := getAccessIndexor(accessGrpLength.Int64(), getAccessGroup, getAccessGroupsRights)
	if err != nil {
		return nil, err
	}

	accessGroupsInfo := map[string]interface{}{
		"accessGroupsLength":   accessGrpLength,
		"accessGroupsCategory": categoryType[accessGrpStruct.Category],
		"accessGroupsList":     accessGroupsList,
	}

	// --------------------------------

	contentTypesLength, err := AcsIndexorInst.GetContentTypesLength(&gc.co)
	if err != nil {
		return nil, err
	}

	contentTypesStruct, err := AcsIndexorInst.ContentTypes(&gc.co)
	if err != nil {
		return nil, err
	}

	getContentType := func(index int64) (common.Address, error) {
		conType, err := AcsIndexorInst.GetContentType(&gc.co, big.NewInt(index))
		if err != nil {
			return common.Address{}, nil
		}
		return conType, nil
	}

	getContentTypeRights := func(addr common.Address, v uint8) (bool, error) {
		rights, err := AcsIndexorInst.CheckAccessGroupRights(&gc.co, addr, v)
		if err != nil {
			return false, err
		}
		return rights, nil
	}

	contentTypesList, err := getAccessIndexor(contentTypesLength.Int64(), getContentType, getContentTypeRights)
	if err != nil {
		return nil, err
	}

	contentTypesInfo := map[string]interface{}{
		"contentTypesLength":   contentTypesLength,
		"contentTypesCategory": categoryType[contentTypesStruct.Category],
		"contentTypesList":     contentTypesList,
	}

	// ------------------------------

	contractsLength, err := AcsIndexorInst.GetContractsLength(&gc.co)
	if err != nil {
		return nil, err
	}

	contractsStruct, err := AcsIndexorInst.Contracts(&gc.co)
	if err != nil {
		return nil, err
	}

	getContract := func(index int64) (common.Address, error) {
		contract, err := AcsIndexorInst.GetContract(&gc.co, big.NewInt(index))
		if err != nil {
			return common.Address{}, nil
		}
		return contract, nil
	}

	getContractRights := func(addr common.Address, v uint8) (bool, error) {
		rights, err := AcsIndexorInst.CheckContractRights(&gc.co, addr, v)
		if err != nil {
			return false, err
		}
		return rights, nil
	}

	contractsList, err := getAccessIndexor(contractsLength.Int64(), getContract, getContractRights)
	if err != nil {
		return nil, err
	}

	contractsInfo := map[string]interface{}{
		"contractsLength":   contractsLength,
		"contractsCategory": categoryType[contractsStruct.Category],
		"contractsList":     contractsList,
	}

	// -------------------------------------

	accessIndexor := map[string]interface{}{
		"ownable":            ownable,
		"librariesInfo":      librariesInfo,
		"contentObjectsInfo": contentObjectsInfo,
		"accessGroupsInfo":   accessGroupsInfo,
		"contentTypesInfo":   contentTypesInfo,
		"contractsInfo":      contractsInfo,
	}

	return accessIndexor, nil

}

func (gc *getContractConfig) getAccessible() (map[string]interface{}, error) {

	accessibleInst, err := contracts.NewAccessible(gc.contractAddr, gc.ethclient)
	if err != nil {
		return nil, err
	}

	versionInBytes32, err := accessibleInst.Version(&gc.co)
	if err != nil {
		return nil, err
	}
	version := strings.TrimFunc(string(versionInBytes32[:]), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	accessible := map[string]interface{}{
		"version": version,
	}

	return accessible, nil
}

func (gc *getContractConfig) getBaseContentSpace() (map[string]interface{}, error) {
	spaceInst, err := contracts.NewBaseContentSpace(gc.contractAddr, gc.ethclient)
	if err != nil {
		return nil, err
	}

	name, err := spaceInst.Name(&gc.co)
	if err != nil {
		return nil, err
	}
	description, err := spaceInst.Description(&gc.co)
	if err != nil {
		return nil, err
	}

	factory, err := spaceInst.Factory(&gc.co)
	if err != nil {
		return nil, err
	}
	groupFactory, err := spaceInst.GroupFactory(&gc.co)
	if err != nil {
		return nil, err
	}
	walletFactory, err := spaceInst.WalletFactory(&gc.co)
	if err != nil {
		return nil, err
	}
	libraryFactory, err := spaceInst.LibraryFactory(&gc.co)
	if err != nil {
		return nil, err
	}
	contentFactory, err := spaceInst.ContentFactory(&gc.co)
	if err != nil {
		return nil, err
	}

	accessible, err := gc.getAccessible()
	if err != nil {
		return nil, err
	}

	container, err := gc.getContainer()
	if err != nil {
		return nil, err
	}

	userSpace, err := gc.getUserSpace()
	if err != nil {
		return nil, err
	}
	nodeSpace, err := gc.getNodeSpace()
	if err != nil {
		return nil, err
	}

	baseContentSpace := map[string]interface{}{
		"data": map[string]interface{}{
			"name":           name,
			"description":    description,
			"factory":        factory,
			"groupFactory":   groupFactory,
			"walletFactory":  walletFactory,
			"libraryFactory": libraryFactory,
			"contentFactory": contentFactory,
		},
		"accessible": accessible,
		"container":  container,
		"xSpace": map[string]interface{}{
			"userSpace": userSpace,
			"nodeSpace": nodeSpace,
		},
	}

	return baseContentSpace, nil
}

func (gc *getContractConfig) getBaseContentType() (map[string]interface{}, error) {
	accessible, err := gc.getAccessible()
	if err != nil {
		return nil, err
	}

	editable, err := gc.getEditable()
	if err != nil {
		return nil, err
	}

	baseContentType := map[string]interface{}{
		"accessible": accessible,
		"editable":   editable,
	}

	return baseContentType, nil
}

func (gc *getContractConfig) getBaseContent() (map[string]interface{}, error) {
	bsContentInst, err := contracts.NewBaseContent(gc.contractAddr, gc.ethclient)
	if err != nil {
		return nil, err
	}

	requestID, err := bsContentInst.RequestID(&gc.co)
	if err != nil {
		return nil, err
	}

	contentType, err := bsContentInst.ContentType(&gc.co)
	if err != nil {
		return nil, err
	}

	addressKMS, err := bsContentInst.AddressKMS(&gc.co)
	if err != nil {
		return nil, err
	}

	contentContractAddress, err := bsContentInst.ContentContractAddress(&gc.co)
	if err != nil {
		return nil, err
	}

	libraryAddress, err := bsContentInst.LibraryAddress(&gc.co)
	if err != nil {
		return nil, err
	}

	accessCharge, err := bsContentInst.AccessCharge(&gc.co)
	if err != nil {
		return nil, err
	}

	statusCode, err := bsContentInst.StatusCode(&gc.co)
	if err != nil {
		return nil, err
	}

	visibility, err := bsContentInst.Visibility(&gc.co)
	if err != nil {
		return nil, err
	}

	editable, err := gc.getEditable()
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"contentType":            contentType,
		"addressKMS":             addressKMS,
		"contentContractAddress": contentContractAddress,
		"libraryAddress":         libraryAddress,
		"accessCharge":           accessCharge,
		"statusCode":             statusCode,
		"visibility":             visibility,
		"requestID":              requestID,
	}

	baseContent := map[string]interface{}{
		"data":     data,
		"editable": editable,
	}

	return baseContent, nil
}

func (gc *getContractConfig) getBaseLibrary() (map[string]interface{}, error) {
	libInst, err := contracts.NewBaseLibrary(gc.contractAddr, gc.ethclient)
	if err != nil {
		return nil, err
	}

	contributorGroupsLength, err := libInst.ContributorGroupsLength(&gc.co)
	if err != nil {
		return nil, err
	}
	var cGroupArr []string
	for i := int64(0); i < contributorGroupsLength.Int64(); i++ {
		cGroup, err := libInst.ContributorGroups(&gc.co, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		cGroupArr = append(cGroupArr, cGroup.String())
	}

	reviewerGroupsLength, err := libInst.ReviewerGroupsLength(&gc.co)
	if err != nil {
		return nil, err
	}
	var rGroupArr []string
	for i := int64(0); i < reviewerGroupsLength.Int64(); i++ {
		rGroup, err := libInst.ReviewerGroups(&gc.co, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		rGroupArr = append(rGroupArr, rGroup.String())
	}

	accessorGroupsLength, err := libInst.AccessorGroupsLength(&gc.co)
	if err != nil {
		return nil, err
	}
	var aGroupArr []string
	for i := int64(0); i < accessorGroupsLength.Int64(); i++ {
		aGroup, err := libInst.AccessorGroups(&gc.co, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		aGroupArr = append(aGroupArr, aGroup.String())
	}

	approvalRequestsLength, err := libInst.ApprovalRequestsLength(&gc.co)
	if err != nil {
		return nil, err
	}
	var approvalReqArr []string
	for i := int64(0); i < approvalRequestsLength.Int64(); i++ {
		approvalRequest, err := libInst.ApprovalRequests(&gc.co, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		approvalReqArr = append(approvalReqArr, approvalRequest.String())
	}

	requiresReview, err := libInst.RequiresReview(&gc.co)
	if err != nil {
		return nil, err
	}

	accessible, err := gc.getAccessible()
	if err != nil {
		return nil, err
	}
	container, err := gc.getContainer()
	if err != nil {
		return nil, err
	}

	baseLibrary := map[string]interface{}{
		"data": map[string]interface{}{
			"contributorGroupsLength": contributorGroupsLength,
			"contributorGroups":       cGroupArr,
			"reviewerGroupsLength":    reviewerGroupsLength,
			"reviewerGroups":          rGroupArr,
			"accessorGroupsLength":    accessorGroupsLength,
			"accessorGroups":          aGroupArr,
			"approvalRequestsLength":  approvalRequestsLength,
			"approvalRequests":        approvalReqArr,
			"requiresReview":          requiresReview,
		},
		"accessible": accessible,
		"container":  container,
	}

	return baseLibrary, nil
}

// ----- for space node locators -------------
// TODO : to convert cbor encoded node locator (requires github.com/qluvio/elv-common...)
// type apiEndpoint struct {
// 	API    string `json:"api"`
// 	Scheme string `json:"scheme"`
// 	Host   string `json:"host"`
// 	Port   string `json:"port"`
// 	Path   string `json:"path"`
// }
//
// func (e *apiEndpoint) ToURL() string {
// 	h := e.Host
// 	if e.Port != "" {
// 		h += ":" + e.Port
// 	}
// 	u := url.URL{Scheme: e.Scheme, Host: h, Path: e.Path}
// 	return u.String()
// }
//
// type nodeRegistration struct {
// 	Private     bool           `json:"private"`
// 	FabLocators []*apiEndpoint `json:"fab"`
// 	EthLocators []*apiEndpoint `json:"eth"`
// }
//
// type EthNode struct {
// 	ID          id.ID
// 	Private     bool
// 	FabLocators []string
// 	EthLocators []string
// }
//
// func (e *EthNode) String() string {
// 	fl := strings.Join(e.FabLocators, ",")
// 	el := strings.Join(e.EthLocators, ",")
// 	return fmt.Sprintf("id:%v,private:%v,fabLocators:%v,ethLocators:%v", e.ID.String(), e.Private, fl, el)
// }
//
// func cborToNodeV1(locBytes []byte) (*EthNode, error) {
// 	buf := bytes.NewBuffer(locBytes)
// 	var endpoints []*apiEndpoint
// 	err := codecs.NewCborCodec().Decoder(buf).Decode(&endpoints)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	node := &EthNode{}
//
// 	for _, e := range endpoints {
// 		h := e.Host
// 		if e.Port != "" {
// 			h += ":" + e.Port
// 		}
// 		u := url.URL{Scheme: e.Scheme, Host: h, Path: e.Path}
// 		l := u.String()
// 		switch e.API {
// 		case "fabric":
// 			node.FabLocators = append(node.FabLocators, l)
// 		case "eth":
// 			node.EthLocators = append(node.EthLocators, l)
// 		default:
// 			return nil, errors.E("encountered invalid api_endpoint type", errors.K.Invalid, "type", e.API)
// 		}
// 	}
//
// 	return node, nil
// }
//
// func cborToNode(locBytes []byte) (*EthNode, error) {
// 	buf := bytes.NewBuffer(locBytes)
// 	var nreg nodeRegistration
// 	err := codecs.NewCborCodec().Decoder(buf).Decode(&nreg)
// 	if err != nil {
// 		return cborToNodeV1(locBytes)
// 	}
//
// 	node := &EthNode{Private: nreg.Private}
//
// 	for _, e := range nreg.FabLocators {
// 		node.FabLocators = append(node.FabLocators, e.ToURL())
// 	}
// 	for _, e := range nreg.EthLocators {
// 		node.EthLocators = append(node.EthLocators, e.ToURL())
// 	}
//
// 	return node, nil
// }
