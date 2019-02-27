pragma solidity ^0.4.21;

import {Accessible} from "./accessible.sol";
import {Editable} from "./editable.sol";
import {BaseAccessControlGroup} from "./base_access_control_group.sol";
import {BaseContentType} from "./base_content_type.sol";
import {BaseLibrary} from "./base_library.sol";
import "./accessible.sol";


contract BaseContentSpace is Accessible, Editable {

    string public name;
    string public description;

    address[] public activeNodeAddresses;
    bytes[] public activeNodeLocators;

    address[] pendingNodeAddresses;
    bytes[] pendingNodeLocators;

    function checkRedundantEntry(address[] _addrs, bytes[] _locators, address _nodeAddr, bytes _nodeLocator) pure internal returns (bool) {
        require(_addrs.length == _locators.length);
        for (uint i = 0; i < _addrs.length; i++) {
            // right now we assume that neither the address or the locator can be used redundantly
            if (keccak256(_locators[i]) == keccak256(_nodeLocator) || _addrs[i] == _nodeAddr) {
                return true;
            }
        }
        return false;
    }

    // we assume that this call is made from the submitted node - that is, from their address
    function submitNode(bytes _locator) public {
        require(!checkRedundantEntry(pendingNodeAddresses, pendingNodeLocators, msg.sender, _locator));
        require(!checkRedundantEntry(activeNodeAddresses, activeNodeLocators, msg.sender, _locator));
        require(pendingNodeAddresses.length < 100); // don't allow abuse - TODO: what value?
        pendingNodeLocators.push(_locator);
        pendingNodeAddresses.push(msg.sender);
    }

    event CreateContentType(address contentTypeAddress);
    event CreateLibrary(address libraryAddress);
    event CreateGroup(address groupAddress);
    event EngageAccountLibrary(address accountAddress);

    constructor(string memory content_space_name) public {
        name = content_space_name;
    }

    function setDescription(string memory content_space_description) public onlyOwner {
        description = content_space_description;
    }

    function createContentType() public returns (address) {
        address contentTypeAddress = new BaseContentType();
        emit CreateContentType(contentTypeAddress);
        return contentTypeAddress;
    }

    function createLibrary(address address_KMS) public returns (address) {
        address libraryAddress = new BaseLibrary(address_KMS);
        emit CreateLibrary(libraryAddress);
        return libraryAddress;
    }

    function createGroup() public returns (address) {
        address groupAddress = new BaseAccessControlGroup();
        emit CreateGroup(groupAddress);
        return groupAddress;
    }

    function engageAccountLibrary() public returns (address) {
        emit EngageAccountLibrary(tx.origin);
    }


}

