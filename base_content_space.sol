//SPDX-License-Identifier: MIT
pragma solidity ^0.7.1;

import {Ownable} from "./ownable.sol";
import {Editable} from "./editable.sol";
import "./base_space_interfaces.sol";
import {BaseAccessControlGroup} from "./base_access_control_group.sol";
import {BaseContentType} from "./base_content_type.sol";
import {BaseLibrary} from "./base_library.sol";
import {BaseContent} from "./base_content.sol";
import {BaseAccessWalletFactory} from "./base_access_wallet.sol";
import {BaseAccessWallet} from "./base_access_wallet.sol";
import {AccessIndexor} from "./access_indexor.sol";
import {Container} from "./container.sol";
import "./user_space.sol";
import "./node_space.sol";
import "./node.sol";
import "./meta_object.sol";
import "./transactable.sol";
import "./lib_precompile.sol";
import "./base_tenant_space.sol";

/* -- Revision history --
BaseContentSpace20190221114100ML: First versioned released
BaseContentSpace20190319194900ML: Requires 0.4.24
BaseContentSpace20190320114200ML: Adding support for user-wallet
BaseContentSpace20190506153400ML: Moves dependant creation to factories, requires factory to be set after instantiation
BaseContentSpace20190510150900ML: Moves content creation from library to a dedicated content space factory
BaseContentSpace20190528193500ML: Moves node management to a parent class (INodeSpace)
BaseContentSpace20190605144600ML: Implements canConfirm to overloads default from Editable
BaseContentSpace20190801140400ML: Breaks AccessGroup creation to its own factory
BaseContentSpace20200309155700ML: Removes import of recording custom contract. To get all events, media_platform.sol should be used
BaseContentSpace20200316120600ML: Defaults visibility to ensure access is open
*/

contract BaseContentSpace is MetaObject, Container, UserSpace, NodeSpace, IKmsSpace, IFactorySpace {

    string public name;
    string public description;
    address payable public factory;
    address payable public groupFactory;
    address payable public walletFactory;
    address payable public libraryFactory;
    address payable public contentFactory;

    mapping(address => address) public nodeMapping;

    mapping(string => bytes[]) kmsMapping;
    mapping(string => string)  kmsPublicKeys;

    event CreateContentType(address contentTypeAddress);
    event CreateLibrary(address libraryAddress);
    event CreateGroup(address groupAddress);
    event CreateContent(address contentAddress);
    event CreateAccessWallet(address wallet);
    event BindUserWallet(address wallet, address userAddr);

    event EngageAccountLibrary(address accountAddress);
    event SetFactory(address factory);

    event RegisterNode(address nodeObjAddr);
    event UnregisterNode(address nodeObjAddr);

    event AddKMSLocator(address sender,uint status);
    event RemoveKMSLocator(address sender, uint status);

    event CreateSpace(bytes32 version, address owner);
    event GetAccessWallet(address walletAddress);

    constructor(string memory content_space_name) {
        version ="BaseContentSpace20200626120600PO"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials X
        
        name = content_space_name;
        contentSpace = address(this);
        emit CreateSpace(version, owner);
        visibility = 10;
    }

    // override
    function setVisibility(uint8 _visibility_code) public override onlyOwner {
        visibility = _visibility_code;
        emit VisibilityChanged(address(this), address(0x0), visibility);
    }

    function setFactory(address payable new_factory) public onlyOwner {
        factory = new_factory;
    }

    function setGroupFactory(address payable new_factory) public onlyOwner {
        groupFactory = new_factory;
    }

    function setWalletFactory(address payable new_factory) public onlyOwner {
        walletFactory = new_factory;
    }

    function setLibraryFactory(address payable new_factory) public onlyOwner {
        libraryFactory = new_factory;
    }

    function setContentFactory(address payable new_factory) public onlyOwner {
        contentFactory = new_factory;
    }

    function setDescription(string memory content_space_description) public onlyOwner {
        description = content_space_description;
    }

    function canConfirm() public view override returns (bool) {
        INodeSpace bcs = INodeSpace(address(this));
        return bcs.canNodePublish(msg.sender);
    }

    function createContentType() public override returns (address payable) {
        address payable contentTypeAddress = BaseFactory(factory).createContentType();
        emit CreateContentType(contentTypeAddress);
        return contentTypeAddress;
    }

    function createLibrary(address payable address_KMS) public override returns (address payable) {
        address payable libraryAddress = BaseLibraryFactory(libraryFactory).createLibrary(address_KMS);
        emit CreateLibrary(libraryAddress);
        return libraryAddress;
    }

    function createContent(address payable lib, address payable content_type) public override returns (address payable) {
        address payable contentAddress = BaseContentFactory(contentFactory).createContent(lib, content_type);
        emit CreateContent(contentAddress);
        return contentAddress;
    }

    function createGroup() public override returns (address payable) {
        address payable groupAddress = BaseGroupFactory(groupFactory).createGroup();
        emit CreateGroup(groupAddress);
        return groupAddress;
    }
    
    function createUserWallet(address payable _user) external override returns (address payable) {
        return createUserWalletInternal(_user);
    }

    function createAccessWallet() public returns (address payable) {
        return createUserWalletInternal(msg.sender);
    }

    // This methods revert when attempting to transfer ownership, so for now we make it private
    // Hence it will be assumed, that user are responsible for creating their wallet.
    function createUserWalletInternal(address payable _user) public returns (address payable) {
        require(userWalletsMap[_user] == address(0x0));
        address payable walletAddress = BaseAccessWalletFactory(walletFactory).createAccessWallet();
        if (_user != msg.sender) {
            BaseAccessWallet wallet = BaseAccessWallet(walletAddress);
            wallet.transferOwnership(_user);
        }
        emit CreateAccessWallet(walletAddress);
        emit BindUserWallet(walletAddress, _user);
        userWalletsMap[_user] = walletAddress;
        return walletAddress;
    }

    function getAccessWallet() public returns(address) {
        address walletAddress;
        if (userWalletsMap[msg.sender] == address(0x0)) {
            walletAddress = createAccessWallet();
        } else {
            walletAddress = userWalletsMap[msg.sender];
        }

        emit GetAccessWallet(walletAddress);
        return walletAddress;
    }

    function getKMSID(address payable _kmsAddr) public view override returns (string memory){
        return Precompile.makeIDString(Precompile.CodeKMS(), _kmsAddr);
    }

    function checkKMS(string memory _kmsIdStr) public view returns (uint) {
        return kmsMapping[_kmsIdStr].length;
    }

    function checkKMSAddr(address payable _kmsAddr) public view override returns (uint) {
        string memory kmsID = getKMSID(_kmsAddr);
        return kmsMapping[kmsID].length;
    }

    // can be used to add or remove - i.e. set to ""
    function setKMSPublicKey(string memory _kmsID, string memory _pubKey) public onlyOwner {
        kmsPublicKeys[_kmsID] = _pubKey;
    }

    function matchesPrefix(bytes memory input, bytes memory prefix) pure internal returns (bool) {
        uint len = prefix.length;
        if (len > input.length) len = input.length;
        for (uint x = 0; x < len; x++) {
            if (input[x] != prefix[x]) return false;
        }
        return true;
    }

    function filterPrefix(bytes[] memory input, bytes memory prefix) view internal returns (bytes[] memory) {
        uint countMatch = 0;
        for (uint i = 0; i < input.length; i++) {
            if (matchesPrefix(input[i], prefix)) {
                countMatch++;
            }
        }
        bytes[] memory output = new bytes[](countMatch);
        if (countMatch == 0) return output;
        countMatch = 0;
        for (uint i = 0; i < input.length; i++) {
            if (matchesPrefix(input[i], prefix)) {
                output[countMatch] = input[i];
                countMatch++;
            }
        }
        return output;
    }

    function getKMSInfo(string memory _kmsID, bytes memory prefix) external view override returns (string memory, string memory) {
        bytes[] memory locators = kmsMapping[_kmsID];
        string memory publicKey = kmsPublicKeys[_kmsID];

        if (locators.length == 0) return ("", publicKey);
        bytes[] memory filtered = filterPrefix(locators, prefix);

        string memory output;
        for (uint i = 0; i < filtered.length; i++) {
            if (i == filtered.length -1) {
                output = string(abi.encodePacked(output, string(filtered[i])));
            } else {
                output = string(abi.encodePacked(output, string(filtered[i]), ","));
            }
        }
        return (output, publicKey);
    }


    // KMS mappings
    // mapping(address => string[]) public kmsMapping;
    // status -> 0 added
    // status -> 1 not added
    function addKMSLocator(string memory _kmsID, bytes memory _locator) public onlyOwner returns (bool) {
        bytes[] memory kmsLocators = kmsMapping[_kmsID];

        for (uint i = 0; i < kmsLocators.length; i++) {
            if (keccak256(kmsLocators[i]) == keccak256(_locator)) {
                emit AddKMSLocator(msg.sender, 1);
                return false;
            }
        }
        kmsMapping[_kmsID].push(_locator);
        emit AddKMSLocator(msg.sender, 0);
        return true;
    }

    // status -> 0 removed
    // status -> 1 not removed
    function removeKMSLocator(string memory _kmsID, bytes memory _locator) public onlyOwner returns (bool) {
        bytes[] memory kmsLocators = kmsMapping[_kmsID];
        for (uint i = 0; i < kmsLocators.length; i++) {
            if (keccak256(kmsLocators[i]) == keccak256(_locator)) {
                if (i != kmsLocators.length - 1) {
                    kmsMapping[_kmsID][i] = kmsLocators[kmsLocators.length - 1];
                }
                delete kmsMapping[_kmsID][kmsLocators.length - 1];
                kmsMapping[_kmsID].pop();
                emit RemoveKMSLocator(msg.sender,0);
                return true;
            }
        }
        emit RemoveKMSLocator(msg.sender,1);
        return false;
    }
    
}



/* -- Revision history --
BaseFactory20190227170400ML: First versioned released
BaseFactory20190301105700ML: No changes version bump to test
BaseFactory20190319195000ML: with  0.4.24 migration
BaseFactory20190506153000ML: Split createLibrary out, adds access indexing
BaseFactory20190722161600ML: No changes, updated to provide generation for BsAccessCtrlGrp20190722161600ML
BaseFactory20190801140700ML: Removed access group creation to its own factory
BaseFactory20200203112400ML: Only records SEE rights in wallet upon creation, to avoid interference with transfer of ownership
BaseFactory20200316120700ML: Uses content-type setRights instead of going straight to the wallet
*/

contract BaseFactory is Ownable {

    constructor(){
        version ="BaseFactory20200316120700ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX
    }

    // TODO: similar issue as tx.origin in Ownable - can't use msg.sender here because it's the space. But as with ownable,
    //  don't think there's a legit spoofing attack here because the spoofee ends up with the rights.
    function createContentType() public returns (address payable) {
        address payable newType = address(new BaseContentType(msg.sender));
        BaseContentType theType = BaseContentType(newType);
        theType.setRights(tx.origin, 0 /*TYPE_SEE*/, 2 /*ACCESS_CONFIRMED*/); // register library in user wallet
        theType.transferOwnership(tx.origin);
        return newType;
    }

    function createNode(address _owner) public returns (address) {
        Node n = new Node(); // this sets owner to tx.origin?
        require(n.owner() == _owner);
        return address(n);
    }
}




/* -- Revision history --
BaseGroupFactory20190729115200ML: First versioned released
BaseGroupFactory20200203112000ML: Only records SEE rights in wallet upon creation, to avoid interference with transfer of ownership
BaseGroupFactory20200316120800ML: Uses group setRights instead of going straight to the wallet
*/

contract BaseGroupFactory is Ownable {

    constructor(){
        version ="BaseGroupFactory20200316120800ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX
    }

    // see note on BaseFactory
    function createGroup() public returns (address payable) {
        address payable newGroup = address(new BaseAccessControlGroup(msg.sender));
        BaseAccessControlGroup theGroup = BaseAccessControlGroup(newGroup);
        theGroup.setRights(tx.origin, 0 /*TYPE_SEE*/, 2 /*ACCESS_CONFIRMED*/);
        theGroup.transferOwnership(tx.origin);
        return newGroup;
    }
}



/* -- Revision history --
BaseLibFactory20190506153200ML: Split out of BaseFactory, adds access indexing
BaseLibFactory20200203111800ML: Only records SEE rights in wallet upon creation, to avoid interference with transfer of ownership
BaseLibFactory20200316121000ML: Uses group setRights instead of going straight to the wallet
*/

contract BaseLibraryFactory is Ownable {

    constructor(){
        version ="BaseLibFactory20200316121000ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX
    }

    // see note on BaseFactory
    function createLibrary(address payable address_KMS) public returns (address payable) {
        address payable newLib = address(new BaseLibrary(address_KMS, msg.sender));
        BaseLibrary theLib = BaseLibrary(newLib);
        theLib.setRights(tx.origin, 0 /*TYPE_SEE*/, 2 /*ACCESS_CONFIRMED*/);  // register library in user wallet
        theLib.transferOwnership(tx.origin);
        return newLib;
    }
}


/* -- Revision history --
BaseCtFactory20190509171900ML: Split out of BaseLibraryFactory
BaseCtFactory20191017165200ML: Updated to reflect change in BaseContent20190801141600ML
BaseCtFactory20191219182100ML: Updated to reflect change in BaseContent20191219135200ML
BaseCtFactory20200203112500ML: Set rights to SEE upon creation to avoid interfering with ownership transfer
BaseCtFactory20200316121100ML: Uses content setRights instead of going straight to the wallet
BaseCtFactory20200422180700ML: Updated to reflect fix of deletion of content objects
*/

contract BaseContentFactory is Ownable {

    constructor(){
        version ="BaseCtFactory20200803130000PO"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX
    }

    // see note on BaseFactory re tx.origin
    function createContent(address payable lib, address payable content_type) public  returns (address payable) {
        Container libraryObj = Container(lib);

        // this looks suspicious because it *can* be spoofed, but the object owner and the rights holder ends up being tx.origin
        //  as well so there doesn't seem to be a legit spoofing attack.
        require(libraryObj.canContribute(tx.origin)); // check if sender has contributor access
        require(libraryObj.validType(content_type));

        BaseContent content = new BaseContent(msg.sender, lib, content_type);
        content.setAddressKMS(libraryObj.addressKMS());
        content.setContentContractAddress(payable(libraryObj.contentTypeContracts(content_type)));

        // register object in user wallet
        BaseContent(content).setRights(tx.origin, 0 /*TYPE_SEE*/, 2 /*ACCESS_CONFIRMED*/);
        content.transferOwnership(tx.origin);

        return address(content);
    }

    uint32 public constant OP_ACCESS_REQUEST = 1;
    uint32 public constant OP_ACCESS_COMPLETE = 2;

    function isContract(address addr) public returns (bool) {
        uint size;
        assembly { size := extcodesize(addr) }
        return size > 0;
    }

    function executeAccessBatch(uint32[] memory _opCodes, address payable[] memory _contentAddrs, address payable[] memory _userAddrs, uint256[] memory _requestNonces, bytes32[] memory _ctxHashes, uint256[] memory _ts, uint256[] memory _amt) public {

        //        BaseContentSpace ourSpace = BaseContentSpace(contentSpace);
        //        require(msg.sender == owner || ourSpace.checkKMSAddr(msg.sender) > 0);

        uint paramsLen = _opCodes.length;

        require(_contentAddrs.length == paramsLen);
        require(_requestNonces.length == paramsLen);
        require(_userAddrs.length == paramsLen);
        require(_ctxHashes.length == paramsLen);
        require(_ts.length == paramsLen);

        for (uint i = 0; i < paramsLen; i++) {
            BaseContent cobj = BaseContent(_contentAddrs[i]);
            // guard against race condition where content object is deleted before batch is executed.
            if (!isContract(_contentAddrs[i]))
                continue;
            // require(msg.sender == owner || cobj.addressKMS() == msg.sender);
            if (_opCodes[i] == OP_ACCESS_REQUEST) {
                cobj.accessRequestContext(_requestNonces[i], _ctxHashes[i], _userAddrs[i], _ts[i]);
            } else if (_opCodes[i] == OP_ACCESS_COMPLETE) {
                cobj.accessCompleteContext(_requestNonces[i], _ctxHashes[i], _userAddrs[i], _ts[i]);
            } else {
                require(false);
            }
        }
    }
}
