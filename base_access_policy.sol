pragma solidity 0.4.24;

import "./ownable.sol";
import "./meta_object.sol";
import "./lib_accessmanager.sol";
import "./editable.sol";
import "./lib_precompile.sol";

contract BaseAccessPolicy is MetaObject, AccessIndexor, Editable {

    string public description;

    address public tenant;

    function isAdmin(address _candidate) public view returns (bool) {

        return (msg.sender == owner || hasEditorRight(msg.sender)
            || AccessManager.isAllowed(msg.sender, Precompile.makePolicyId(address(this)), "edit"));
    }

    function setTenant(address _tenantAddr) public {
        require(isAdmin(msg.sender));
        tenant = _tenantAddr;
    }

    // TODO: actually, maybe this goes in content space ...?
    event AccessPolicyCreated(address space);
    event AccessPolicyDeleted(address space); // TODO: kill method?

    constructor(address _contentSpace)  public payable {
        contentSpace = _contentSpace; // TODO: why isn't content space in constructor of Ownable?
        emit AccessPolicyCreated(contentSpace);
    }

    int public constant OP_ADD_SUBJECT = 1;
    int public constant OP_REMOVE_SUBJECT = 2;
    int public constant OP_ADD_RESOURCE = 3;
    int public constant OP_REMOVE_RESOURCE = 4;
    int public constant OP_ADD_ACTION = 5;
    int public constant OP_REMOVE_ACTION = 6;
    int public constant OP_SET_EFFECT = 7;
    int public constant OP_CHANGE_DEF = 8;

    event AccessPolicyChanged(address space, address tenant, int op, string relTarget);

    // override from AccessIndexor...
    function setEntityRights(uint8 indexType, address obj, uint8 access_type, uint8 access) public  {
        require(Editable(obj).hasEditorRight(tx.origin)); // TODO: is this necessary or does setEntityRights suffice...?
        require(isAdmin(tx.origin));
        super.setEntityRights(indexType, obj, access_type, access);
        if (indexType == CATEGORY_CONTENT_OBJECT) {
            emit AccessPolicyChanged(contentSpace, tenant, OP_ADD_RESOURCE, Precompile.makeObjId(obj));
        } else if (indexType == CATEGORY_GROUP) {
            emit AccessPolicyChanged(contentSpace, tenant, OP_ADD_SUBJECT, Precompile.makeGroupId(obj));
        } else if (indexType == CATEGORY_LIBRARY) {
            emit AccessPolicyChanged(contentSpace, tenant, OP_ADD_RESOURCE, Precompile.makeLibId(obj));
        }
        // TODO: anything different (revert?) if it's not one of the handled types ...?
    }

    // override from Editable...
    function confirmCommit() public payable returns (bool) {
        return super.confirmCommit();
        emit AccessPolicyChanged(contentSpace, tenant, OP_CHANGE_DEF, "");
    }

    function addSubject(string _id) public returns (bool) {
        // require(isAdmin(msg.sender)); TODO: TEST !!!
        AccessManager.isAllowed(msg.sender, _id, "admin"); // edit?
        emit AccessPolicyChanged(contentSpace, tenant, OP_ADD_SUBJECT, _id);
        return true;
    }

    function removeSubject(string _id) public returns (bool) {
        require(isAdmin(msg.sender));
        AccessManager.isAllowed(msg.sender, _id, "admin"); // edit?
        emit AccessPolicyChanged(contentSpace, tenant, OP_REMOVE_SUBJECT, _id);
        return true;
    }

    function addResource(string _id) public returns (bool) {
        require(isAdmin(msg.sender));
        AccessManager.isAllowed(msg.sender, _id, "admin"); // edit?
        emit AccessPolicyChanged(contentSpace, tenant, OP_ADD_RESOURCE, _id);
        return true;
    }

    function removeResource(string _id) public returns (bool) {
        require(isAdmin(msg.sender));
        AccessManager.isAllowed(msg.sender, _id, "admin"); // edit?
        emit AccessPolicyChanged(contentSpace, tenant, OP_REMOVE_RESOURCE, _id);
        return true;
    }

    string public constant EFFECT_ALLOW = "allow";
    string public constant EFFECT_DENY = "deny";
    string public effect;

    function setEffect(string _newEffect) public returns (bool) {
        require(isAdmin(msg.sender));
        require(keccak256(abi.encodePacked(_newEffect)) == keccak256(abi.encodePacked(EFFECT_ALLOW)) ||
        keccak256(abi.encodePacked(_newEffect)) == keccak256(abi.encodePacked(EFFECT_DENY)));
        effect = _newEffect;
        emit AccessPolicyChanged(contentSpace, tenant, OP_SET_EFFECT, _newEffect);
        return true;
    }

    function addAction(string _action) public returns (bool) {
        require(isAdmin(msg.sender));
        emit AccessPolicyChanged(contentSpace, tenant, OP_ADD_ACTION, _action);
        return true;
    }

    function removeAction(string _action) public returns (bool) {
        require(isAdmin(msg.sender));
        emit AccessPolicyChanged(contentSpace, tenant, OP_REMOVE_ACTION, _action);
        return true;
    }
}