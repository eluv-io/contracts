pragma solidity ^0.4.24;

library AccessManager {

    address constant elv_precomp_addr_am = 253;

    bytes4 constant sigIdIsAllowed = bytes4(keccak256("isAllowed(address,string,string)"));

    function isAllowed(address _subject, string _resourceId, string _action) internal view returns (bool ret) {

        bytes4 sig = sigIdIsAllowed;
        address callAddr = elv_precomp_addr_am;
        bytes32 truncAction;
        bytes32 truncId;

        assembly {
            let x := mload(0x40)
            mstore(x, sig)
            mstore(add(x, 0x04), _subject)
            truncId := mload(add(_resourceId, 32))
            mstore(add(x, 0x24), truncId)
            truncAction := mload(add(_action, 32))
            mstore(add(x, 0x44), truncAction)

            let res := call(
            0,          // no gas
            callAddr,   // To addr
            0,          // No value - i.e. tokens
            x,          // Inputs are stored at location x
            0x64,       // Inputs are 100 bytes long
            x,          // Store output over input
            0x20)       // Output is 32 byte bool

            if eq(res, 0) {
                revert(0, 0)
            }

            ret := mload(x)
            mstore(0x40, add(x, 0x20))  // Set storage pointer to empty space - i.e. after returned storage
        }
    }
}