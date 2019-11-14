pragma solidity 0.4.24;

import "./lib_precompile.sol";
import {Ownable} from "./ownable.sol";
import "./transactable.sol";

/**
 * UserSpace
 * Seperated from content_space to avoid circular dependencies.
 */

/* -- Revision history --
KMSSpace20191107155800ML: First versioned released
*/


contract KMSSpace is Ownable {

    bytes32 public version ="KMSSpace20191107155800ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX


    mapping(string => bytes[]) kmsMapping;
    mapping(string => string)  kmsPublicKeys;


    event AddKMSLocator(address sender,uint status);
    event RemoveKMSLocator(address sender, uint status);



    // TODO kmsAddr => kmsID
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
    function setKMSPublicKey(string _kmsID, string _pubKey) public onlyOwner {
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

    function getKMSInfo(string _kmsID, bytes prefix) public view returns (string, string) {
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
    function addKMSLocator(string _kmsID, bytes _locator) public onlyOwner returns (bool) {
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
    function removeKMSLocator(string _kmsID, bytes _locator) public onlyOwner returns (bool) {
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

    function executeBatch(uint8[] _v, bytes32[] _r, bytes32[] _s, address[] _from, address[] _dest, uint256[] _value, uint256[] _ts) public {

        require(msg.sender == owner || checkKMSAddr(msg.sender) > 0);

        // TODO: not sure if this is worth it - will just crash if the parameters are passed in incorrectly, which is the same as a revert...?
        require(_v.length == _r.length);
        require(_r.length == _s.length);
        require(_s.length == _from.length);
        require(_from.length == _dest.length);
        require(_dest.length == _value.length);
        require(_value.length == _ts.length);

        for (uint i = 0; i < _v.length; i++) {
            Transactable t = Transactable(_from[i]);
            bool success = t.execute(msg.sender, _v[i], _r[i], _s[i], _dest[i], _value[i], _ts[i]);

            if (!success) {
                // we failed to get the target wallet to pay so - in this scenario - we have to pay!
                // _dest[i].send(_value[i]); // TODO: transfer? error handling?
            }
        }
    }

}
