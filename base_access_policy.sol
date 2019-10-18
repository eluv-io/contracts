pragma solidity 0.4.24;

import "./ownable.sol";
import "./meta_object.sol";
import "./lib_accessmanager.sol";

contract BaseAccessPolicy is Ownable, MetaObject {

    event AccessPolicyChanged(address space);

    string public description;

    address[] public userSubjects;
    address[] public groupSubjects;
    // or just ...?
    address[] public subjects;

    string[] public actions;

    // TODO: content types and maybe groups as resources ???
    address[] public libraryResources;
    address[] public contentResources;
    // or just ...?
    address[] public resources;
    // conditions can be stored as part of metadata - i.e. in MetaObject

    string public constant EFFECT_ALLOW = "allow";
    string public constant EFFECT_DENY = "deny";
    string public effect;

    constructor(address _contentSpace)  public payable {
        contentSpace = _contentSpace; // TODO: why isn't content space in constructor of Ownable?
        effect = EFFECT_ALLOW; // default is allow
    }

    function addBasicPolicy(address _subject, address _resource, string _action) public returns (bool) {
        require(msg.sender == owner || AccessManager.isAllowed(msg.sender, address(this), "edit"));

        // TODO: add _subject, _resource, _action

        emit AccessPolicyChanged(contentSpace);
        return true;
    }

    function addSubject(address _subject) public returns (bool) {
        require(msg.sender == owner || AccessManager.isAllowed(msg.sender, address(this), "edit"));

        // TODO: add _subject

        emit AccessPolicyChanged(contentSpace);
        return true;
    }

    function addResource(address _resource) public returns (bool) {
        require(msg.sender == owner || AccessManager.isAllowed(msg.sender, address(this), "edit"));

        // TODO: add _resource

        emit AccessPolicyChanged(contentSpace);
        return true;
    }

    function addAction(string _action) public returns (bool) {
        require(msg.sender == owner || AccessManager.isAllowed(msg.sender, address(this), "edit"));

        // TODO: add _action

        emit AccessPolicyChanged(contentSpace);
        return true;
    }
}
