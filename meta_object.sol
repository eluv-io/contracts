pragma solidity 0.4.24;

import "./ownable.sol";

contract MetaObject is Ownable, IAdmin {

    mapping(bytes32 => bytes) mapSmallKeys;
    mapping(bytes => bytes) mapBigKeys;

    event ObjectMetaChanged(bytes key);

    // meant to be overridden in base classes
    function isAdmin(address _candidate) public view returns (bool) {
        if (_candidate == owner) {
            return true;
        }
        return false;
    }

    modifier onlyAdmin() {
        require(isAdmin(msg.sender));
        _;
    }

    function putMeta(bytes key, bytes value) public onlyAdmin {
        if (key.length <= 32) {
            bytes32 smallKey;
            uint256 keyLen = key.length;
            assembly {
                smallKey := mload(add(key, keyLen))
            }
            delete mapSmallKeys[smallKey];
            if (value.length > 0)
                mapSmallKeys[smallKey] = value;
        } else {
            delete mapBigKeys[key];
            if (value.length > 0)
                mapBigKeys[key] = value;
        }
        emit ObjectMetaChanged(key);
    }

    function getMeta(bytes key) public constant returns (bytes) {
        if (key.length <= 32) {
            bytes32 smallKey;
            uint256 keyLen = key.length;
            assembly {
                smallKey := mload(add(key, keyLen))
            }
            return mapSmallKeys[smallKey];
        }
        return mapBigKeys[key];
    }
}