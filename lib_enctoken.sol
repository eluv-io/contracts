pragma solidity ^0.4.24;

library PrecompEncToken {

    address constant elv_precomp_enc_token_addr = 254;

    bytes4 constant sigIdString = bytes4(keccak256("getString(string,bytes)"));
    bytes4 constant sigIdUint = bytes4(keccak256("getUint(string,bytes)"));
    bytes4 constant sigIdAddress = bytes4(keccak256("getAddress(string,bytes)"));

    function getString(string _attrib, bytes _tok) internal constant returns (string ret) {
        return getValue(_attrib, _tok, sigIdString);
    }

    function getUint(string _attrib, bytes _tok) internal constant returns (uint256 ret) {
        getValue(_attrib, _tok, sigIdUint);
    }

    function getAddress(string _attrib, bytes _tok) internal constant returns (address ret) {
        bytes4 sig = sigIdAddress;
        address callAddr = elv_precomp_enc_token_addr;
        bytes32 truncAttrib;

        uint tokLen = _tok.length;
        uint allLen = 4 + 32 + tokLen; // 4bytes + truncated token + token bytes length

        assembly {
            let x := mload(0x40)        // Find empty storage location using "free memory pointer"
            mstore(x, sig)              // Place signature at beginning of empty storage

            truncAttrib := mload(add(_attrib, 32))
            mstore(add(x, 0x04), truncAttrib)

            for { let i := 0 } lt(i, tokLen) { i := add(i, 32) } {
                mstore(add(x, add(0x24,i)), mload(add(add(_tok, 0x20), i)))
            }

            let res := call(
            0,
            callAddr,
            0,
            x,              // Inputs are stored at location x
            allLen,
            x,              // Store output over input (saves space)
            0x20)

            switch res
            case 0 {
                revert(0, 0)
            }
            default {
                ret := mload(x)
                // mstore(0x40, add(x, 0x20))  // Set storage pointer to empty space - i.e. after returned storage
            }
        }
    }

//    function getBytes(string tok, string attrib) internal constant returns (bytes ret) {
//        return msg.data; // TODO: bytes constant ...?
//    }

//    function callByBytes(bytes4 _func, bytes _param) public {
//        address _tmpAddr = elv_precomp_enc_token_addr;
//        uint paramLen = _param.length;
//        uint allLen = 4 + paramLen;
//        assembly {
//            let p := mload(0x40)
//            mstore(p, _func)
//            for { let i := 0 } lt(i, paramLen) { i := add(i, 32) } {
//                mstore(add(p, add(4,i)), mload(add(add(_param, 0x20), i)))
//            }
//
//            let success := call(not(0), _tmpAddr, 0, p, allLen, 0, 0)
//
//            let size := returndatasize
//            returndatacopy(p, 0, size)
//
//            switch success
//            case 0 {
//                revert(p, size)
//            }
//            default {
//                return(p, size)
//            }
//        }
//    }

    function getValue(string _attrib, bytes _tok, bytes4 _sigBytes) internal constant returns (string ret) {
        bytes4 sig = _sigBytes;
        address callAddr = elv_precomp_enc_token_addr;
        bytes32 truncAttrib;

        uint tokLen = _tok.length;
        uint allLen = 4 + 32 + tokLen; // 4bytes + truncated token + token bytes length

        assembly {
            let x := mload(0x40)        // Find empty storage location using "free memory pointer"
            mstore(x, sig)              // Place signature at beginning of empty storage

            truncAttrib := mload(add(_attrib, 32))
            mstore(add(x, 0x04), truncAttrib)

            for { let i := 0 } lt(i, tokLen) { i := add(i, 32) } {
                mstore(add(x, add(0x24,i)), mload(add(add(_tok, 0x20), i)))
            }

            let res := call(
            0,
            callAddr,
            0,
            x,              // Inputs are stored at location x
            allLen,
            x,              // Store output over input (saves space)
            0x40)

            switch res
            case 0 {
                revert(0, 0)
            }
            default {
                ret := x
                mstore(0x40, add(x, 0x40))  // Set storage pointer to empty space - i.e. after returned storage
            }
        }
    }
}
