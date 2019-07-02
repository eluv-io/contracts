pragma solidity 0.4.24;

import {Ownable} from "./ownable.sol";


/* -- Revision history --
Editable20190222140100ML: First versioned released
Editable20190315141800ML: Migrated to 0.4.24
Editable20190515103600ML: Modified rights restriction on update to match the one on commit
Editable20190522154000SS: Changed hash bytes32 to string
Editable20190605144500ML: Renamed publish to confirm to avoid confusion in the case of the content-objects
*/


contract Editable is Ownable {

    bytes32 public version ="Editable20190607105600PO"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    event CommitPending(address spaceAddress, address parentAddress, string objectHash);
    event UpdateRequest(string objectHash);
    event VersionConfirm(string objectHash);
    event VersionDelete(string versionHash, int256 index);

    string public objectHash;
    string[] public versionHashes;
    string pendingHash;

    function countVersionHashes() public view returns (uint256) {
        return versionHashes.length;
    }

    // if _versionHash is present returns the index of the versionHash else length(versionHashes)+1
    function getVersionIndex(string _versionHash) public view returns (int256) {
       
        bytes32 vh = keccak256(abi.encode(_versionHash));
        for(uint256 i=0; i < versionHashes.length; i++){
            if (keccak256(bytes(versionHashes[i])) == vh) {
                return int256(i);
            }
        }
        return -1;
    }

    // intended to be overridden
    function canConfirm() public view returns (bool) {
        return false;
    }

    function canCommit() public view returns (bool) {
        return (tx.origin == owner);
    }

    // overridden in BaseContent to return library
    function parentAddress() returns (address) {
        return contentSpace;
    }

    function commit(string _objectHash) public {
        require(canCommit());
	    require(bytes(_objectHash).length < 128);
        pendingHash = _objectHash;
        emit CommitPending(contentSpace, parentAddress(), pendingHash);
    }

    function confirmCommit() public payable returns (bool) {
        require(canConfirm());

        if (bytes(objectHash).length > 0) {
            versionHashes.push(objectHash); // save existing version info
        }
        objectHash = pendingHash;
        pendingHash = "";
        emit VersionConfirm(objectHash);
    }

    function updateRequest() public {
        require(msg.sender == owner || canConfirm());
        emit UpdateRequest(objectHash);
    }

    function deleteVersion(string _versionHash) public returns (int256) {
        require(canCommit());
        
        int256 index = getVersionIndex(_versionHash);

        require(index != -1);

        // reset index value to default
        delete versionHashes[uint256(index)];
        emit VersionDelete(_versionHash, index);
        return index;
    }

}
