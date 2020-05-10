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

// main difference from BaseContentSpace it tenant is *not* a NodeSpace *can* control publishing through INodeSpace
contract BaseTenantSpace is MetaObject, Accessible, Container, UserSpace, INodeSpace, IKmsSpace, IFactorySpace {

    bytes32 public version ="BaseTenantSpace20200504120000PO"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    string public name;
    string public description;

    // TODO: are all factories valid? most likely but likely want mechanism that can defer to base content space
    address public factory;
    address public groupFactory;
    address public walletFactory;
    address public libraryFactory;
    address public contentFactory;

    mapping(string => bytes[]) kmsMapping;
    mapping(string => string)  kmsPublicKeys;

    event CreateContentType(address contentTypeAddress);
    event CreateLibrary(address libraryAddress);
    event CreateGroup(address groupAddress);
    event CreateContent(address contentAddress);
    event CreateAccessWallet(address wallet);
    event BindUserWallet(address wallet, address userAddr);

    event AddKMSLocator(address sender,uint status);
    event RemoveKMSLocator(address sender, uint status);

    // TODO: create tenant?
    event CreateTenant(bytes32 version, address owner);
    event GetAccessWallet(address walletAddress);

    address contentSpace;
    constructor(address _contentSpace, string _tenantName) public payable {
        name = _tenantName;
        BaseContentSpace spc = BaseContentSpace(_contentSpace);
        // allow either the space owner or a trusted address to refer to the space
        require(msg.sender == spc.owner() || spc.checkKMSAddr(msg.sender) > 0);
        contentSpace = address(_contentSpace);
        emit CreateTenant(version, owner);
    }

    address public adminGroup;
    address public defaultUserGroup; // TODO: should this always just point to a group that contains *all* users in the tenant?

    event SetTenantGroups(address adminGroup, address defaultUserGroup);

    function isAdmin(address _candidate) public view returns (bool) {
        if (_candidate == owner) {
            return true;
        }
        if (adminGroup == 0x0) // TODO: don't know if it should be *valid* to not have an admin group but ...
            return false;
        return BaseAccessControlGroup(adminGroup).hasManagerAccess(_candidate);
    }

    modifier onlyAdmin() {
        require(isAdmin(msg.sender));
        _;
    }

    function setGroups(address _adminGroup, address _defaultUserGroup) onlyAdmin {
        if (_adminGroup != 0x0) {
            adminGroup = _adminGroup;
        }
        if (_defaultUserGroup != 0x0) {
            defaultUserGroup = _defaultUserGroup;
        }
        emit SetTenantGroups(adminGroup, defaultUserGroup);
    }

    // for 'historical' reasons the tenant ID is based on the address of the admin group - not this contract!
    function getTenantID() public view returns (string){
        return Precompile.makeIDString(Precompile.CodeTEN(), adminGroup);
    }

    function setFactory(address new_factory) public onlyAdmin {
        factory = new_factory;
    }

    function setGroupFactory(address new_factory) public onlyAdmin {
        groupFactory = new_factory;
    }

    function setWalletFactory(address new_factory) public onlyAdmin {
        walletFactory = new_factory;
    }

    function setLibraryFactory(address new_factory) public onlyAdmin {
        libraryFactory = new_factory;
    }

    function setContentFactory(address new_factory) public onlyAdmin {
        contentFactory = new_factory;
    }

    function setDescription(string memory content_space_description) public onlyAdmin {
        description = content_space_description;
    }

    // TODO: just point at space method?
    function canConfirm() public view returns (bool) {
        INodeSpace bcs = INodeSpace(address(this));
        return bcs.canNodePublish(msg.sender);
    }

    // TODO: ???
    function canNodePublish(address _candidate) external view returns (bool) {
        INodeSpace spc = INodeSpace(contentSpace);
        return spc.canNodePublish(_candidate);
    }

    // TODO: auth?
    function createContentType() public returns (address) {
        address contentTypeAddress = BaseFactory(factory).createContentType();
        emit CreateContentType(contentTypeAddress);
        return contentTypeAddress;
    }

    // TODO: auth?
    function createLibrary(address address_KMS) public returns (address) {
        address libraryAddress = BaseLibraryFactory(libraryFactory).createLibrary(address_KMS);
        emit CreateLibrary(libraryAddress);
        return libraryAddress;
    }

    // TODO: auth?
    function createContent(address lib, address content_type) public returns (address) {
        address contentAddress = BaseContentFactory(contentFactory).createContent(lib, content_type);
        emit CreateContent(contentAddress);
        return contentAddress;
    }

    // TODO: auth?
    function createGroup() public returns (address) {
        address groupAddress = BaseGroupFactory(groupFactory).createGroup();
        emit CreateGroup(groupAddress);
        return groupAddress;
    }

    // TODO: generally we likely want to allow the tenant to add KMS info but also fall back to the space ...?
    //  for simplicity (?) we might want to make it so that if the tenant provides *any* KMS info we *don't* fall back
    //  to the space. otherwise, there would be not mechanism for the tenant to 'override' the default KMS info.
    function getKMSID(address _kmsAddr) public view returns (string){
        return Precompile.makeIDString(Precompile.CodeKMS(), _kmsAddr);
    }

    function checkKMS(string _kmsIdStr) public view returns (uint) {
        return kmsMapping[_kmsIdStr].length;
    }

    function checkKMSAddr(address _kmsAddr) public view returns (uint) {
        string memory kmsID = getKMSID(_kmsAddr);
        return kmsMapping[kmsID].length;
    }

    // can be used to add or remove - i.e. set to ""
    function setKMSPublicKey(string _kmsID, string _pubKey) public onlyAdmin {
        kmsPublicKeys[_kmsID] = _pubKey;
    }

    function matchesPrefix(bytes input, bytes prefix) pure internal returns (bool) {
        uint len = prefix.length;
        if (len > input.length) len = input.length;
        for (uint x = 0; x < len; x++) {
            if (input[x] != prefix[x]) return false;
        }
        return true;
    }

    function filterPrefix(bytes[] input, bytes prefix) view internal returns (bytes[]) {
        uint countMatch = 0;
        for (uint i = 0; i < input.length; i++) {
            if (matchesPrefix(input[i], prefix)) {
                countMatch++;
            }
        }
        bytes[] memory output = new bytes[](countMatch);
        if (countMatch == 0) return output;
        countMatch = 0;
        for (i = 0; i < input.length; i++) {
            if (matchesPrefix(input[i], prefix)) {
                output[countMatch] = input[i];
                countMatch++;
            }
        }
        return output;
    }

    function getKMSInfo(string _kmsID, bytes prefix) external view returns (string, string) {
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
    function addKMSLocator(string _kmsID, bytes _locator) public onlyAdmin returns (bool) {
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
    function removeKMSLocator(string _kmsID, bytes _locator) public onlyAdmin returns (bool) {
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