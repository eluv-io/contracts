pragma solidity 0.5.4;

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
import {Utils} from "./utils.sol";
import "./user_space.sol";
import "./node_space.sol";
import "./node.sol";
import "./meta_object.sol";
import "./transactable.sol";
import "./lib_precompile.sol";
import "./base_tenant_space.sol";
import "./base_tenant_consumer_group.sol";
import "./erc721_proxy.sol";

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
BaseContentSpace20200626120600PO: authv3 changes
BaseContentSpace20200928110000PO: Replace tx.origin with msg.sender in some cases
*/

contract KmsManager is Ownable, IKmsSpace {

    mapping(string => bytes[]) kmsMapping;
    mapping(string => string)  kmsPublicKeys;

    event AddKMSLocator(address sender, uint status);
    event RemoveKMSLocator(address sender, uint status);

    function getKMSID(address _kmsAddr) public view returns (string memory) {
        return Precompile.makeIDString(Precompile.CodeKMS(), _kmsAddr);
    }

    function checkKMS(string memory _kmsIdStr) public view returns (uint) {
        return kmsMapping[_kmsIdStr].length;
    }

    function checkKMSAddr(address _kmsAddr) public view returns (uint) {
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

    function filterPrefix(bytes[] memory input, bytes memory prefix) pure internal returns (bytes[] memory) {
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

    function getKMSInfo(string calldata _kmsID, bytes calldata prefix) external view returns (string memory, string memory) {
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
                kmsMapping[_kmsID].length -= 1;
                emit RemoveKMSLocator(msg.sender,0);
                return true;
            }
        }
        emit RemoveKMSLocator(msg.sender,1);
        return false;
    }
}

contract BaseContentSpace is MetaObject, Container, UserSpace, NodeSpace, IKmsSpace, IFactorySpace {

    bytes32 public version ="BaseContentSpace20200928110000PO"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    string public name;
    string public description;
    address payable public factory;
    address payable public groupFactory;
    address payable public walletFactory;
    address payable public libraryFactory;
    address payable public contentFactory;

    mapping(address => address) public nodeMapping;

    address payable public kmsManAddr;

    function setKmsManager(address payable _kmsManAddr) public onlyOwner {
        kmsManAddr = _kmsManAddr;
    }

    event CreateContentType(address contentTypeAddress);
    event CreateLibrary(address libraryAddress);
    event CreateGroup(address groupAddress);
    event CreateContent(address contentAddress);
    event CreateAccessWallet(address wallet);
    event BindUserWallet(address wallet, address userAddr);

    event EngageAccountLibrary(address accountAddress);
    event SetFactory(address factory);

    event CreateSpace(bytes32 version, address owner);
    event GetAccessWallet(address walletAddress);

    constructor(string memory content_space_name) public {
        name = content_space_name;
        contentSpace = address(this);
        emit CreateSpace(version, owner);
        visibility = 10;
    }

    // override
    function setVisibility(uint8 _visibility_code) public onlyOwner {
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

    function canConfirm() public view returns (bool) {
        INodeSpace bcs = INodeSpace(address(this));
        return bcs.canNodePublish(msg.sender);
    }

    function createContentType() public returns (address) {
        require(msg.sender == tx.origin);
        address contentTypeAddress = BaseTypeFactory(factory).createContentType();
        emit CreateContentType(contentTypeAddress);
        return contentTypeAddress;
    }

    function createLibrary(address address_KMS) public returns (address) {
        require(msg.sender == tx.origin);
        address libraryAddress = BaseLibraryFactory(libraryFactory).createLibrary(address_KMS);
        emit CreateLibrary(libraryAddress);
        return libraryAddress;
    }

    function createContent(address payable lib, address payable content_type) public returns (address) {
        require(msg.sender == tx.origin || msg.sender == lib);
        address contentAddress = BaseContentFactory(contentFactory).createContent2(lib, content_type);
        emit CreateContent(contentAddress);
        return contentAddress;
    }

    function createGroup() public returns (address) {
        require(msg.sender == tx.origin);
        address groupAddress = BaseGroupFactory(groupFactory).createGroup();
        emit CreateGroup(groupAddress);
        return groupAddress;
    }

    function createUserWallet(address payable _user) external returns (address) {
        return createUserWalletInternal(_user);
    }

    function createAccessWallet() public returns (address) {
        return createUserWalletInternal(msg.sender);
    }

    // This methods revert when attempting to transfer ownership, so for now we make it private
    // Hence it will be assumed, that user are responsible for creating their wallet.
    function createUserWalletInternal(address payable _user) private returns (address) {
        require(userWallets[_user] == address(0x0));
        address payable walletAddress = BaseAccessWalletFactory(walletFactory).createAccessWallet();
        if (_user != msg.sender) {
            BaseAccessWallet wallet = BaseAccessWallet(walletAddress);
            wallet.transferOwnership(_user);
        }
        emit CreateAccessWallet(walletAddress);
        emit BindUserWallet(walletAddress, _user);
        userWallets[_user] = walletAddress;
        return walletAddress;
    }

    function getAccessWallet() public returns(address) {
        address walletAddress;
        if (userWallets[msg.sender] == address(0x0)) {
            walletAddress = createAccessWallet();
        } else {
            walletAddress = userWallets[msg.sender];
        }

        emit GetAccessWallet(walletAddress);
        return walletAddress;
    }

    // implement IKmsSpace

    function checkKMS(string calldata _kmsIdStr) external view returns (uint) {
        return KmsManager(kmsManAddr).checkKMS(_kmsIdStr);
    }

    function checkKMSAddr(address _kmsAddr) external view returns (uint) {
        return KmsManager(kmsManAddr).checkKMSAddr(_kmsAddr);
    }

    function getKMSID(address _kmsAddr) external view returns (string memory) {
        return KmsManager(kmsManAddr).getKMSID(_kmsAddr);
    }

    function getKMSInfo(string calldata _kmsID, bytes calldata prefix) external view returns (string memory, string memory) {
        return KmsManager(kmsManAddr).getKMSInfo(_kmsID, prefix);
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
BaseFactory20200928110000PO: Replace tx.origin with msg.sender in some cases
BaseFactory20201129223200ML: Fix glitch in creation of library (registration to the wrong index)
*/

contract BaseTypeFactory is Ownable {

    bytes32 public version ="BaseFactory20201129223200ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    constructor(address payable _spaceAddr) public {
        contentSpace = _spaceAddr;
    }

    // TODO: similar issue as tx.origin in Ownable - can't use msg.sender here because it's the space. But as with ownable,
    //  don't think there's a legit spoofing attack here because the spoofee ends up with the rights.
    function createContentType() public returns (address) {
        require(msg.sender == contentSpace);
        address payable newType = address(new BaseContentType(msg.sender));
        BaseContentType theType = BaseContentType(newType);

        IUserSpace userSpaceObj = IUserSpace(contentSpace);
        address payable userWallet = address(uint160(userSpaceObj.userWallets(tx.origin)));

        // due to a (intentional?) limitation of the EVM, this stanza *needs* to be duplicated as-is. in particular,
        //  factoring this code into a shared subroutine will cause it to fail.
        bool isV3Contract = Utils.checkV3Contract(userWallet);
        if (!isV3Contract) {
            AccessIndexor index = AccessIndexor(userWallet);
            theType.transferOwnership(tx.origin);
            index.setContentTypeRights(address(newType), 0, 2);
        } else {
            // v3+ path ...
            theType.setRights(tx.origin, 0 /*TYPE_SEE*/, 2 /*ACCESS_CONFIRMED*/); // register library in user wallet
            theType.transferOwnership(tx.origin);
        }

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
BaseGroupFactory20200928110000PO: Replace tx.origin with msg.sender in some cases
*/

contract BaseGroupFactory is Ownable {

    bytes32 public version ="BaseGroupFactory20200928110000PO"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    constructor(address payable _spaceAddr) public {
        contentSpace = _spaceAddr;
    }

    // see note on BaseFactory
    function createGroup() public returns (address) {
        require(msg.sender == contentSpace);
        address payable newGroup = address(new BaseAccessControlGroup(msg.sender));
        BaseAccessControlGroup theGroup = BaseAccessControlGroup(newGroup);

        IUserSpace userSpaceObj = IUserSpace(contentSpace);
        address payable userWallet = address(uint160(userSpaceObj.userWallets(tx.origin)));

        // due to a (intentional?) limitation of the EVM, this stanza *needs* to be duplicated as-is. in particular,
        //  factoring this code into a shared subroutine will cause it to fail.
        bool isV3Contract = Utils.checkV3Contract(userWallet);
        if (!isV3Contract) {
            AccessIndexor index = AccessIndexor(userWallet);
            theGroup.transferOwnership(tx.origin);
            index.setAccessGroupRights(newGroup, 0, 2);
        } else {
            // v3+ path ...
            theGroup.setRights(tx.origin, 0 /*TYPE_SEE*/, 2 /*ACCESS_CONFIRMED*/);
            theGroup.transferOwnership(tx.origin);
        }

        return newGroup;
    }
}

/* -- Revision history --
BaseLibFactory20190506153200ML: Split out of BaseFactory, adds access indexing
BaseLibFactory20200203111800ML: Only records SEE rights in wallet upon creation, to avoid interference with transfer of ownership
BaseLibFactory20200316121000ML: Uses group setRights instead of going straight to the wallet
BaseLibFactory20200928110000PO: Replace tx.origin with msg.sender in some cases
*/

contract BaseLibraryFactory is Ownable {

    bytes32 public version ="BaseLibFactory20200928110000PO"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    constructor(address payable _spaceAddr) public {
        contentSpace = _spaceAddr;
    }

    // see note on BaseFactory
    function createLibrary(address address_KMS) public returns (address) {
        require(msg.sender == contentSpace);
        address payable newLib = address(new BaseLibrary(address_KMS, msg.sender));
        BaseLibrary theLib = BaseLibrary(newLib);

        IUserSpace userSpaceObj = IUserSpace(contentSpace);
        address payable userWallet = address(uint160(userSpaceObj.userWallets(tx.origin)));

        // due to a (intentional?) limitation of the EVM, this stanza *needs* to be duplicated as-is. in particular,
        //  factoring this code into a shared subroutine will cause it to fail.
        bool isV3Contract = Utils.checkV3Contract(userWallet);
        if (!isV3Contract) {
            AccessIndexor index = AccessIndexor(userWallet);
            theLib.transferOwnership(tx.origin);
            index.setLibraryRights(newLib, 0, 2);
        } else {
            // v3+ path ...
            theLib.setRights(tx.origin, 0 /*TYPE_SEE*/, 2 /*ACCESS_CONFIRMED*/);  // register library in user wallet
            theLib.transferOwnership(tx.origin);
        }

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
BaseCtFactory20200803130000PO: authv3 changes
BaseCtFactory20200928110000PO: Replace tx.origin with msg.sender in some cases

*/

contract ContentFactoryHelper {
    function getContentObjectBytecode() public pure returns (bytes memory) {
        return type(BaseContent).creationCode;
    }
}

contract BaseContentFactory is Ownable {

    bytes32 public version ="BaseCtFactory20200928110000PO"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX
    bytes4 constant createContentSig = bytes4(keccak256("createContentInstance(address,address,address)"));

    // TODO: *not* adding a setter here means that we need to create a new factory if we want to change the helper - i.e. change
    //  the base content code. I *think* this is what we want to do.
    address payable factoryHelper;

    constructor(address payable _spaceAddr, address payable _helperAddr) public {
        contentSpace = _spaceAddr;
        factoryHelper = _helperAddr;
    }

    // creationBytecode uses the creationCode of baseContent ..
    // but this does not work and the contracts reverts
    function creationBytecode(
        address _spcAddr,
        address _libAddr,
        address _typeAddr) private view returns (bytes memory) {

        bytes memory bytecode = ContentFactoryHelper(factoryHelper).getContentObjectBytecode();
        return abi.encodePacked(bytecode, abi.encode(_spcAddr, _libAddr, _typeAddr));
    }

    // saltFor computes a salt for the given address: 'address++block.number'
    // The function is public for testing purposes.
    function saltFor(address addr) public view returns (uint256) {
        uint256 ret;
        ret = uint256(uint160(addr));
        ret <<= 96;
        return uint256(ret) + block.number;
    }

    function deploy(bytes memory bytecode, uint _salt) private returns (address payable _addr) {
        assembly {
            _addr := create2(0, add(bytecode, 0x20), mload(bytecode), _salt)

            if iszero(extcodesize(_addr)) {
                revert(0, 0)
            }
        }
    }

    function createContent2(address payable lib, address payable content_type) public  returns (address) {
        require(msg.sender == contentSpace);
        Container libraryObj = Container(lib);

        // this looks suspicious because it *can* be spoofed, but the object owner and the rights holder ends up being tx.origin
        //  as well so there doesn't seem to be a legit spoofing attack.
        require(libraryObj.canContribute(tx.origin)); // check if sender has contributor access
        require(libraryObj.validType(content_type));

        // BaseContent content = new BaseContent(msg.sender, lib, content_type);
        // ideally we would use this but it does not work
        bytes memory encodeBytes = creationBytecode(msg.sender, lib, content_type);
        address payable contentAddr = deploy(encodeBytes, saltFor(tx.origin));

        BaseContent content = BaseContent(contentAddr);

        content.setAddressKMS(libraryObj.addressKMS());
        content.setContentContractAddress(libraryObj.contentTypeContracts(content_type));

        IUserSpace userSpaceObj = IUserSpace(contentSpace);
        address payable userWallet = address(uint160(userSpaceObj.userWallets(tx.origin)));

        // due to a (intentional?) limitation of the EVM, this stanza *needs* to be duplicated as-is. in particular,
        //  factoring this code into a shared subroutine will cause it to fail.
        bool isV3Contract = Utils.checkV3Contract(userWallet);
        if (!isV3Contract) {
            AccessIndexor index = AccessIndexor(userWallet);
            content.transferOwnership(tx.origin);
            index.setContentObjectRights(address(content), 0, 2);
        } else {
            // v3+ path ...
            BaseContent(content).setRights(tx.origin, 0 /*TYPE_SEE*/, 2 /*ACCESS_CONFIRMED*/);
            content.transferOwnership(tx.origin);
        }

        return address(content);
    }

    uint32 public constant OP_ACCESS_REQUEST = 1;
    uint32 public constant OP_ACCESS_COMPLETE = 2;

    function isContract(address addr) public view returns (bool) {
        uint size;
        assembly { size := extcodesize(addr) }
        return size > 0;
    }

    function executeAccessBatch(
        uint32[] memory _opCodes,
        address payable[] memory _contentAddrs,
        address[] memory _userAddrs,
        uint256[] memory _requestNonces,
        bytes32[] memory _ctxHashes,
        uint256[] memory _ts,
        uint256[] memory) public {

        BaseContentSpace ourSpace = BaseContentSpace(contentSpace);
        require(msg.sender == owner || ourSpace.checkKMSAddr(msg.sender) > 0);

        uint paramsLen = _opCodes.length;

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
