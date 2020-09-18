//SPDX-License-Identifier: MIT
pragma solidity ^0.7.1;

import {Ownable} from "./ownable.sol";

/* -- Revision history --
Node20190315105100ML: First versioned released
*/

contract Node is Ownable {

    event Log(string label);
    event LogBool(string label, bool b);
    event LogAddress(string label, address a);
    event LogUint256(string label, uint256 u);
    event LogInt256(string label, int256 u);
    event LogBytes32(string label, bytes32 b);
    
    constructor() {
        version = "Node20190315105100ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX
    }

    function log(string memory label) public onlyOwner {
        emit Log(label);
    }
}
