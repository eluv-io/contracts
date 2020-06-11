pragma solidity ^0.4.24;

library AccessManager {

    string public constant CREATE_LIBRARY = "create-library";

    address constant elv_precomp_addr_am = 254;

    bytes4 constant sigIdIsAllowed = bytes4(keccak256("isAllowed(address,address,string)"));

    function constCreateLibrary() internal constant returns (string) {
        return CREATE_LIBRARY;
    }

    function isAllowed(address _subject, address _resource, string _action) internal view returns (bool ret) {
        return false;
    }

    // TODO: not yet implemented internally ...
    function isAllowedToDo(address _subject, address _resource, string _action) internal view returns (bool ret) {

        bytes4 sig = sigIdIsAllowed;
        address callAddr = elv_precomp_addr_am;
        bytes32 truncAction;

        assembly {
            let x := mload(0x40)
            mstore(x, sig)
            mstore(add(x, 0x04), _subject)
            mstore(add(x, 0x24), _resource)
            truncAction := mload(add(_action, 32))
            mstore(add(x, 0x44), truncAction)
        // input is now [4 bytes][32 bytes][32 bytes][32 bytes] = 100 bytes

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