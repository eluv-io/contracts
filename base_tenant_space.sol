pragma solidity 0.4.24;

import {Ownable} from "./ownable.sol";
import {Accessible} from "./accessible.sol";
import {Editable} from "./editable.sol";
import "./base_space_interfaces.sol";
import {BaseAccessWalletFactory} from "./base_access_wallet.sol";
import {BaseAccessWallet} from "./base_access_wallet.sol";
import {BaseAccessControlGroup} from "./base_access_control_group.sol";
import {Container} from "./container.sol";
import "./user_space.sol";
import "./node_space.sol";
import "./meta_object.sol";
import "./lib_precompile.sol";
import "./base_content_space.sol";
import "./lib_enctoken.sol";

// CAREFUL: no storage! - otherwise it will conflict with the calling contract
contract TenantFuncTokenTransfer is MetaObject {

    event TenantTransfer(bytes32 ident, address to, uint256 amount);

    function checkTransferToken(bytes _encAuthToken, uint256 _amount, address _to) public view returns (bool) {

        address maybeAddr = EncToken.getAddress("iid", _encAuthToken);
        if (maybeAddr != address(this)) { // TODO: what is the address when delegatecall ...? maybe just always check this on the call side ...?
            return false;
        }

        uint256 maxAmount = EncToken.getUint("max", _encAuthToken);
        if (_amount > maxAmount) {
            return false;
        }

        bytes32 segIdent = bytes32(EncToken.getUint("_SEGIDENT_", _encAuthToken));
        uint256 otpOrd = EncToken.getUint("ord", _encAuthToken);
        uint8 segBit = uint8(otpOrd % 256);
        bool isSet = getBit(segIdent, segBit);

        return !isSet;
    }

    function transferToken(bytes _encAuthToken, uint256 _amount, address _to) public {

        address maybeAddr = EncToken.getAddress("iid", _encAuthToken);
        require(maybeAddr == address(this));

        uint256 maxAmount = EncToken.getUint("max", _encAuthToken);
        require(_amount <= maxAmount);

        string memory otpId = EncToken.getString("id", _encAuthToken);
        uint256 otpOrd = EncToken.getUint("ord", _encAuthToken);

        bytes32 segIdent = bytes32(EncToken.getUint("_SEGIDENT_", _encAuthToken));

        uint8 segBit = uint8(otpOrd % 256);
        bool wasSet = setAndGetBitInternal(segIdent, segBit);
        require(!wasSet);

        _to.transfer(_amount);

        emit TenantTransfer(segIdent, _to, _amount);
    }

}

contract BaseTenantSpace is MetaObject, Accessible, Container, IUserSpace, INodeSpace, IKmsSpace, IFactorySpace {

    bytes32 public version ="BaseTenantSpace20200504120000PO"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    string public name;
    string public description;

    function setDescription(string memory _desc) public onlyAdmin {
        description = _desc;
    }

    event CreateTenant(bytes32 version, address owner);
    event GetAccessWallet(address walletAddress);

    address contentSpace;
    constructor(address _contentSpace, string _tenantName, address _owner) public payable {
        name = _tenantName;
        BaseContentSpace spc = BaseContentSpace(_contentSpace);
        // allow either the space owner or a trusted address to refer to the space
        require(msg.sender == spc.owner() || spc.checkKMSAddr(msg.sender) > 0);
        contentSpace = address(_contentSpace);
        owner = _owner;
        emit CreateTenant(version, owner);
    }

    bytes32 public constant GROUP_ID_ADMIN = "tenant_admin";

    mapping(bytes32 => address[]) public groupsMapping;
    bytes32[] public groupIds;

    function isAdmin(address _candidate) public view returns (bool) {
        if (_candidate == owner) {
            return true;
        }

        address[] memory maybeAddrs = groupsMapping[GROUP_ID_ADMIN];
        for (uint256 i = 0; i < maybeAddrs.length; i++) {
            if (BaseAccessControlGroup(maybeAddrs[i]).hasManagerAccess(_candidate)) {
                return true;
            }
        }
        return false;
    }

    event FunctionsAdded(bytes4[] func4Bytes, address funcAddr);

    mapping(bytes4 => address) public funcMapping;

    function addFuncs(bytes4[] _func4Bytes, address _funcAddr) public onlyAdmin {
        for (uint256 i = 0; i < _func4Bytes.length; i++) {
            funcMapping[_func4Bytes[i]] = _funcAddr;
        }
        emit FunctionsAdded(_func4Bytes, _funcAddr);
    }

    function callFuncUintAddr(bytes4 _func4Bytes, uint256 _p1, address _p2, bytes _encAuthToken) public {
        address maybeFuncAddr = funcMapping[_func4Bytes];
        require(maybeFuncAddr != 0x0);

        // check signature!

        bool success = maybeFuncAddr.delegatecall(abi.encodeWithSelector(_func4Bytes, _encAuthToken, _p1, _p2));
        require(success);
    }

    event TenantTransfer(bytes32 ident, address to, uint256 amount);

    function transfer(address _to, uint256 _amount) public onlyAdmin {
        _to.transfer(_amount);
    }

    function transferToken(address _to, uint256 _amount, bytes _encAuthToken) public {

        address maybeAddr = EncToken.getAddress("iid", _encAuthToken);
        require(maybeAddr == address(this));

        uint256 maxAmount = EncToken.getUint("max", _encAuthToken);
        require(_amount <= maxAmount);

        string memory otpId = EncToken.getString("id", _encAuthToken);
        uint256 otpOrd = EncToken.getUint("ord", _encAuthToken);

        bytes32 segIdent = bytes32(EncToken.getUint("_SEGIDENT_", _encAuthToken));

        uint8 segBit = uint8(otpOrd % 256);
        bool wasSet = setAndGetBitInternal(segIdent, segBit);
        require(!wasSet);

        _to.transfer(_amount);

        emit TenantTransfer(segIdent, _to, _amount);
    }

    modifier onlyAdmin() {
        require(isAdmin(msg.sender));
        _;
    }

    event AddTenantGroup(bytes32 groupId, address groupAddr);
    event RemoveTenantGroup(bytes32 groupId, address groupAddr);

    function addGroup(bytes32 _id, address _groupAddr) public onlyAdmin {
        require(_groupAddr != 0x0);
        for (uint256 i = 0; i < groupsMapping[_id].length; i++) {
            require(_id != GROUP_ID_ADMIN); // only one tenant admin group allowed
            if (groupsMapping[_id][i] == _groupAddr) {
                return;
            }
        }
        if (groupsMapping[_id].push(_groupAddr) == 1) {
            groupIds.push(_id);
        }
        emit AddTenantGroup(_id, _groupAddr);
    }

    function checkIdsRemove(bytes32 _id) {
        if (groupsMapping[_id].length == 0) {
            for (uint256 i = 0; i < groupIds.length; i++) {
                if (groupIds[i] == _id) {
                    if (i != groupIds.length - 1) {
                        groupIds[i] = groupIds[groupIds.length - 1];
                    }
                    delete groupIds[groupIds.length - 1];
                    groupIds.length -= 1;
                }
            }
        }
    }

    function removeGroup(bytes32 _id, address _groupAddr) public onlyAdmin {
        require(_id != GROUP_ID_ADMIN); // can't change the tenant admin group
        for (uint256 i = 0; i < groupsMapping[_id].length; i++) {
            if (groupsMapping[_id][i] == _groupAddr) {
                uint256 len = groupsMapping[_id].length;
                if (i != len - 1) {
                    groupsMapping[_id][i] = groupsMapping[_id][len - 1];
                }
                delete groupsMapping[_id][len - 1];
                groupsMapping[_id].length -= 1;
                checkIdsRemove(_id);
                emit RemoveTenantGroup(_id, _groupAddr);
                return;
            }
        }
    }

    // for 'historical' reasons the tenant ID is based on the address of the admin group - not this contract!
    function getTenantID() public view returns (string) {
        address adminGroup = groupsMapping[GROUP_ID_ADMIN][0];
        return Precompile.makeIDString(Precompile.CodeTEN(), adminGroup);
    }

    // implement IUserSpace

    address public userManager;
    event SetUserManager(address newManager, address prevManager);

    function setUserManager(address _userMan) public onlyAdmin {
        emit SetUserManager(_userMan, userManager);
        userManager = _userMan;
    }

    function userWallets(address _userAddr) external view returns (address) {
        if (userManager != 0x0) {
            return IUserSpace(userManager).userWallets(_userAddr);
        }
        return IUserSpace(contentSpace).userWallets(_userAddr);
    }

    function createUserWallet(address _userAddr) external returns (address) {
        if (userManager != 0x0) {
            return IUserSpace(userManager).createUserWallet(_userAddr);
        }
        // at this point in time, deployed spaces do not implement this method so we will revert - but this is likely what we want to do anyway
        return IUserSpace(contentSpace).createUserWallet(_userAddr);
    }

    // implement INodeSpace

    address public nodeManager;
    event SetNodeManager(address newManager, address prevManager);

    function setNodeManager(address _nodeMan) public onlyAdmin {
        emit SetNodeManager(_nodeMan, nodeManager);
        nodeManager = _nodeMan;
    }

    function canConfirm() public view returns (bool) {
        if (nodeManager != 0x0) {
            return INodeSpace(nodeManager).canNodePublish(msg.sender);
        }
        return INodeSpace(contentSpace).canNodePublish(msg.sender);
    }

    function canNodePublish(address _candidate) external view returns (bool) {
        if (nodeManager != 0x0) {
            return INodeSpace(nodeManager).canNodePublish(_candidate);
        }
        return INodeSpace(contentSpace).canNodePublish(_candidate);
    }

    // implement IFactorySpace

    address public factoryManager;
    event SetFactoryManager(address newManager, address prevManager);

    function setFactoryManager(address _factMan) public onlyAdmin {
        emit SetFactoryManager(_factMan, factoryManager);
        factoryManager = _factMan;
    }

    function createContentType() public returns (address) {
        if (factoryManager != 0x0) {
            return IFactorySpace(factoryManager).createContentType();
        }
        return IFactorySpace(contentSpace).createContentType();
    }

    function createLibrary(address _kmsAddress) public returns (address) {
        if (factoryManager != 0x0) {
            return IFactorySpace(factoryManager).createLibrary(_kmsAddress);
        }
        return IFactorySpace(contentSpace).createLibrary(_kmsAddress);
    }

    function createContent(address _lib, address _contentType) public returns (address) {
        if (factoryManager != 0x0) {
            return IFactorySpace(factoryManager).createContent(_lib, _contentType);
        }
        return IFactorySpace(contentSpace).createContent(_lib, _contentType);
    }

    function createGroup() public returns (address) {
        if (factoryManager != 0x0) {
            return IFactorySpace(factoryManager).createGroup();
        }
        return IFactorySpace(contentSpace).createGroup();
    }

    address public kmsManager;
    event SetKmsManager(address newManager, address prevManager);

    function setKmsManager(address _kmsMan) public onlyAdmin {
        emit SetKmsManager(_kmsMan, kmsManager);
        kmsManager = _kmsMan;
    }

    // implement IKmsSpace

    function getKMSID(address _kmsAddr) public view returns (string){
        if (kmsManager != 0x0) {
            return IKmsSpace(kmsManager).getKMSID(_kmsAddr);
        }
        return IKmsSpace(contentSpace).getKMSID(_kmsAddr);
    }

    function checkKMSAddr(address _kmsAddr) public view returns (uint) {
        if (kmsManager != 0x0) {
            return IKmsSpace(kmsManager).checkKMSAddr(_kmsAddr);
        }
        return IKmsSpace(contentSpace).checkKMSAddr(_kmsAddr);
    }

    function getKMSInfo(string _kmsID, bytes _prefix) external view returns (string, string) {
        if (kmsManager != 0x0) {
            return IKmsSpace(kmsManager).getKMSInfo(_kmsID, _prefix);
        }
        return IKmsSpace(contentSpace).getKMSInfo(_kmsID, _prefix);
    }
}