pragma solidity ^0.4.21;

import {Ownable} from "./ownable.sol";


contract Editable is Ownable {


    event Commit(bytes32 objectHash);
    event UpdateRequest(bytes32 objectHash);

    bytes32 public objectHash;

    // intended to be overridden
    function canPublish() view returns (bool) {
        return false;
    }

    function commit(bytes32 object_hash) public {
        require(msg.sender == owner || canPublish());
        objectHash = object_hash;
        emit Commit(objectHash);
    }

    function updateRequest() public onlyOwner {
        emit UpdateRequest(objectHash);
    }

}
