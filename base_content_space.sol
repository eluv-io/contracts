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

    address[] public pendingNodeAddresses;
    bytes[] public pendingNodeLocators;

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

    event NodeSubmitted(address addr, bytes locator);

    function numActiveNodes() public view returns (uint) {
        return activeNodeLocators.length;
    }

    function numPendingNodes() public view returns (uint) {
        return pendingNodeLocators.length;
    }

    // we assume that this call is made from the submitted node - that is, from their address
    function submitNode(bytes _locator) public {
        require(!checkRedundantEntry(pendingNodeAddresses, pendingNodeLocators, msg.sender, _locator));
        require(!checkRedundantEntry(activeNodeAddresses, activeNodeLocators, msg.sender, _locator));
        require(pendingNodeAddresses.length < 100); // don't allow *too* much abuse - TODO: what value?
        pendingNodeLocators.push(_locator);
        pendingNodeAddresses.push(msg.sender);
        emit NodeSubmitted(msg.sender, _locator);
    }

    function approveNode(address _nodeAddr) public onlyOwner {
        bool found = false;
        for (uint i = 0; i < pendingNodeAddresses.length; i++) {
            if (pendingNodeAddresses[i] == _nodeAddr) {
                activeNodeAddresses.push(pendingNodeAddresses[i]);
                activeNodeLocators.push(pendingNodeLocators[i]);
                if (i != pendingNodeAddresses.length - 1) {
                    pendingNodeLocators[i] = pendingNodeLocators[pendingNodeLocators.length - 1];
                    pendingNodeAddresses[i] = pendingNodeAddresses[pendingNodeAddresses.length - 1];
                }
                delete pendingNodeLocators[pendingNodeLocators.length - 1];
                delete pendingNodeAddresses[pendingNodeAddresses.length - 1];
                found = true;
            }
        }
        require(found);
    }

    // check whether an address - which should represent a content fabric node - can confirm (publish?) a content object
    function canNodePublish(address candidate) public view returns (bool) {
        for (uint i = 0; i < activeNodeAddresses.length; i++) {
            if (activeNodeAddresses[i] == candidate) {
                return true;
            }
        }
        return false;
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
        address libraryAddress = new BaseLibrary(this, address_KMS);
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

