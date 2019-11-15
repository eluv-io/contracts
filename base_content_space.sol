pragma solidity 0.4.24;

import {Ownable} from "./ownable.sol";
import {Accessible} from "./accessible.sol";
import {Editable} from "./editable.sol";
import {BaseAccessControlGroup} from "./base_access_control_group.sol";
import {BaseContentType} from "./base_content_type.sol";
import {BaseLibrary} from "./base_library.sol";
import {BaseContent} from "./base_content.sol";
import {BaseAccessWallet} from "./base_access_wallet.sol";
import {AccessIndexor} from "./access_indexor.sol";
import {Container} from "./container.sol";
import "./user_space.sol";
import "./node_space.sol";
import "./node.sol";
import "./meta_object.sol";
import {KMSSpace} from "./kms_space.sol";
import {BaseFactory} from "./base_factory.sol";


/* -- Revision history --
BaseContentSpace20190221114100ML: First versioned released
BaseContentSpace20190319194900ML: Requires 0.4.24
BaseContentSpace20190320114200ML: Adding support for user-wallet
BaseContentSpace20190506153400ML: Moves dependant creation to factories, requires factory to be set after instantiation
BaseContentSpace20190510150900ML: Moves content creation from library to a dedicated content space factory
BaseContentSpace20190528193500ML: Moves node management to a parent class (NodeSpace)
BaseContentSpace20190605144600ML: Implements canConfirm to overloads default from Editable
BaseContentSpace20190801140400ML: Breaks AccessGroup creation to its own factory
*/


contract BaseContentSpace is MetaObject, Accessible, Container, UserSpace, NodeSpace, KMSSpace {

    bytes32 public version ="BaseContentSpace20190801140400ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    string public name;
    string public description;
    address public factory;
    address public groupFactory;
    address public libraryFactory;
    address public contentFactory;

    mapping(address => address) public nodeMapping;


    event CreateContentType(address contentTypeAddress);
    event CreateLibrary(address libraryAddress);
    event CreateGroup(address groupAddress);
    event CreateContent(address contentAddress);


    event EngageAccountLibrary(address accountAddress);
    event SetFactory(address factory);

    event RegisterNode(address nodeObjAddr);
    event UnregisterNode(address nodeObjAddr);


    event CreateSpace(bytes32 version, address owner);

    constructor(string memory content_space_name) public {
        name = content_space_name;
        contentSpace = address(this);
        emit CreateSpace(version, owner);
    }

    function setFactory(address new_factory) public onlyOwner {
        factory = new_factory;
    }

    function setGroupFactory(address new_factory) public onlyOwner {
        groupFactory = new_factory;
    }

    function setWalletFactory(address new_factory) public onlyOwner {
        walletFactory = new_factory;
    }

    function setLibraryFactory(address new_factory) public onlyOwner {
        libraryFactory = new_factory;
    }

    function setContentFactory(address new_factory) public onlyOwner {
        contentFactory = new_factory;
    }

    function setDescription(string memory content_space_description) public onlyOwner {
        description = content_space_description;
    }

    function canConfirm() public view returns (bool) {
        return canNodePublish(msg.sender);
    }

    // used to create a node contract instance. should be called by the address of the node that wishes to register.
    function registerSpaceNode() public returns (address) {
        require(nodeMapping[msg.sender] == 0x0); // for now can't re-register (or replace) node instance
        uint i = 0;
        for (; i < activeNodeAddresses.length; i++) {
            if (activeNodeAddresses[i] == msg.sender) {
                break;
            }
        }
        require(i < activeNodeAddresses.length); // node should be in active list
        address nodeAddr = BaseFactory(factory).createNode(msg.sender);
        nodeMapping[msg.sender] = nodeAddr;
        emit RegisterNode(nodeAddr);
        return nodeAddr;
    }

    function unregisterSpaceNode() public returns (bool) {
        require(nodeMapping[msg.sender] != 0x0);
        address nodeAddr = nodeMapping[msg.sender];
        delete nodeMapping[msg.sender];
        Node(nodeAddr).kill();
        emit UnregisterNode(nodeAddr);
    }


    function createContentType() public returns (address) {
        address contentTypeAddress = BaseFactory(factory).createContentType();
        emit CreateContentType(contentTypeAddress);
        return contentTypeAddress;
    }

    function createLibrary(address address_KMS) public returns (address) {
        address libraryAddress = BaseLibraryFactory(libraryFactory).createLibrary(address_KMS);
        emit CreateLibrary(libraryAddress);
        return libraryAddress;
    }

    function createContent(address lib, address content_type) public returns (address) {
        address contentAddress = BaseContentFactory(contentFactory).createContent(lib, content_type);
        emit CreateContent(contentAddress);
        return contentAddress;
    }

    function createGroup() public returns (address) {
        address groupAddress = BaseGroupFactory(groupFactory).createGroup();
        emit CreateGroup(groupAddress);
        return groupAddress;
    }

    function engageAccountLibrary() public returns (address) {
        emit EngageAccountLibrary(tx.origin);
    }

}



/* -- Revision history --
BaseGroupFactory20190729115200ML: First versioned released
*/

contract BaseGroupFactory is Ownable {

    bytes32 public version ="BaseGroupFactory20190729115200ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    function createGroup() public returns (address) {
        address newGroup = (new BaseAccessControlGroup(msg.sender));
        // register library in user wallet
        BaseContentSpace contentSpaceObj = BaseContentSpace(msg.sender);
        address walletAddress = contentSpaceObj.getAccessWallet();
        BaseAccessWallet userWallet = BaseAccessWallet(walletAddress);
        userWallet.setAccessGroupRights(newGroup, userWallet.TYPE_EDIT(), userWallet.ACCESS_CONFIRMED());
        return newGroup;
    }
}
/* -- Revision history --
BaseFactory20190506153100ML: Split out of BaseFactory, adds access indexing
*/

contract BaseLibraryFactory is Ownable {

    bytes32 public version ="BaseLibFactory20190506153200ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    function createLibrary(address address_KMS) public returns (address) {
        address newLib = (new BaseLibrary(address_KMS, msg.sender));
        // register library in user wallet
        BaseContentSpace contentSpaceObj = BaseContentSpace(msg.sender);
        address walletAddress = contentSpaceObj.getAccessWallet();
        BaseAccessWallet userWallet = BaseAccessWallet(walletAddress);
        userWallet.setLibraryRights(newLib, userWallet.TYPE_EDIT(), userWallet.ACCESS_CONFIRMED());
        return newLib;
    }
}


/* -- Revision history --
BaseCtFactory20190509171900ML: Split out of BaseLibraryFactory
BaseCtFactory20191017165200ML: Updated to reflect change in BaseContent20190801141600ML
*/

contract BaseContentFactory is Ownable {

    bytes32 public version ="BaseCtFactory20191017165200ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    function createContent(address lib, address content_type) public  returns (address) {
        Container libraryObj = Container(lib);
        require(libraryObj.canContribute(tx.origin)); //check if sender has contributor access
        require(libraryObj.validType(content_type));

        BaseContent content = new BaseContent(msg.sender, lib, content_type);
        content.setAddressKMS(libraryObj.addressKMS());
        content.setContentContractAddress(libraryObj.contentTypeContracts(content_type));

        // register object in user wallet
        BaseContentSpace contentSpaceObj = BaseContentSpace(msg.sender);
        address walletAddress = contentSpaceObj.getAccessWallet();
        AccessIndexor userWallet = AccessIndexor(walletAddress);
        userWallet.setContentObjectRights(address(content), userWallet.TYPE_EDIT(), userWallet.ACCESS_CONFIRMED());

        return address(content);
    }
}

//BaseCtFactoryXt20191031115100PO: adds support for custom contract
//BaseCtFactoryXt20191031153200ML: passes accessor to the runAccess via the addresses array
//BaseCtFactoryXt20191031170400ML: adds request timestamp to event
//BaseCtFactoryXt20191031203100ML: change initialization of array

contract BaseContentFactoryExt is BaseContentFactory {

    bytes32 public version ="BaseCtFactoryXt20191031203100ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    uint32 public constant OP_ACCESS_REQUEST = 1;
    uint32 public constant OP_ACCESS_COMPLETE = 2;

    function isContract(address addr) returns (bool) {
        uint size;
        assembly { size := extcodesize(addr) }
        return size > 0;
    }

    function executeAccessBatch(uint32[] _opCodes, address[] _contentAddrs, address[] _userAddrs, bytes32[] _requestNonces, bytes32[] _ctxHashes, uint256[] _ts, uint256[] _amt) public {

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
                cobj.accessComplete(_requestNonces[i], _amt[i], 0x0);
            } else {
                require(false);
            }
        }
    }
}