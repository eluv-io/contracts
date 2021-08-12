pragma solidity 0.5.4;

import {Editable} from "./editable.sol";
import {MetaObject, CounterObject} from "./meta_object.sol";
import {BaseTenantSpace} from "./base_tenant_space.sol";

contract BaseTenantConsumerGroup is MetaObject, CounterObject, Editable {

    bytes32 public version ="BsTenantConsGrp20210809150000PO";

    mapping(address => bool) public membersMap;
    uint256 public membersNum;

    event MemberAdded(address candidate);
    event MemberRevoked(address candidate);

    bool public isConsumerGroup;

    address payable public tenant;

    constructor(address payable _tenantSpace) public {
        tenant = _tenantSpace;
        isConsumerGroup = true;
        membersNum = 0;
    }

    function isAdmin(address payable _candidate) public view returns (bool) {
        if (_candidate == owner || _candidate == tenant) {
            return true;
        }
        return BaseTenantSpace(tenant).isAdmin(_candidate);
    }

    function hasAccess(address candidate) public view returns (bool) {
        return membersMap[candidate];
    }

    function grantAccessMany(address payable[] memory candidates) public {
        for (uint i = 0; i < candidates.length; i++) {
            grantAccess(candidates[i]);
        }
    }

    function grantAccess(address payable candidate) public {
        require(isAdmin(msg.sender) || (msg.sender == candidate));

        if (!membersMap[candidate]) {
            membersMap[candidate] = true;
            membersNum++;
        }

        emit MemberAdded(candidate);
    }

    function revokeAccess(address payable candidate) public {
        require(isAdmin(msg.sender) || (msg.sender == candidate));

        if (membersMap[candidate]) {
            delete membersMap[candidate];
            membersNum--;
        }

        emit MemberRevoked(candidate);
    }
}
