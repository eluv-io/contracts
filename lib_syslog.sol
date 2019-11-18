pragma solidity ^0.4.24;

library SystemLog {

    address constant elv_precomp_addr_sl = 253;

    bytes4 constant sigIdLog = bytes4(keccak256("log(address,string,string)"));

    function log(address _addr, string _level, string _msg) internal view returns (bool ret) {

        bytes4 sig = sigIdLog;
        address callAddr = elv_precomp_addr_sl;

        bytes32 truncLevel; // for simplicity we always truncate level to 32 bytes, since we expect it to be ERROR|WARN|INFO etc.
        uint msgLen;

        assembly {
            let x := mload(0x40)
            mstore(x, sig)

            mstore(add(x, 0x04), _addr)

            truncLevel := mload(add(_level, 32)) // truncate bytes at _level to 32 and store
            mstore(add(x, 0x24), truncLevel)

            let msgLen := mload(_msg) // load value at _msg - i.e. length of message?
            mstore(add(x, 0x44), msgLen) // hopefully this will be the length of _msg

            for { let i := 0 } lt(i, msgLen) { i := add(i, 32) } {
                mstore(add(x, add(0x44,i)), mload(add(add(_msg, 0x20), i)))
            }

            // input is now [4 bytes][32 bytes][32 bytes] = 68 bytes
            let totalLen := add(0x44, msgLen)

            let res := call(
            0,          // no gas
            callAddr,   // To addr
            0,          // No value - i.e. tokens
            x,          // Inputs are stored at location x
            totalLen,
            x,          // Store output over input
            0x01)       // Output is 1 byte bool

            if eq(res, 0) {
                revert(0, 0)
            }

            ret := x
            mstore(0x40, add(x, 0x01))  // Set storage pointer to empty space - i.e. after returned storage
        }
    }
}