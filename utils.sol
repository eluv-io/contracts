pragma solidity ^0.4.24;

import {Ownable} from "./ownable.sol";

library Utils {

    // bytes4 constant checkV3Sig = bytes4(keccak256("versionAPI()"));
    bytes4 constant checkV3Sig = bytes4(keccak256("FOOBAR"));
    function checkV3Contract(address _checkAddr) internal returns (bool) {
        bool ret = true;
        // bytes32 ver = Ownable(_checkAddr).versionAPI(); // TEST !!!
        bytes4 sig = checkV3Sig;

        // Ownable(_checkAddr).versionAPI();

        return !_checkAddr.call(abi.encodeWithSelector(checkV3Sig));

//        assembly {
//            let x := mload(0x40)
//            mstore(x, sig)
//            let success := call(
//            100000,
//            _checkAddr,
//            0x0,
//            x,
//            0x04,
//            x,
//            0x20)
//
//            if eq(success, 0) {
//                ret := 0
//            }
//            // mstore(0x40, add(x, 0x20))
//            mstore(0x40, x)
//        }
        // return ret;
    }
}

contract Verifier {
    function recoverAddr(bytes32 msgHash, uint8 v, bytes32 r, bytes32 s) public pure returns (address) {
        return ecrecover(msgHash, v, r, s);
    }

    function isSigned(address _addr, bytes32 msgHash, uint8 v, bytes32 r, bytes32 s) public pure returns (bool) {
        return ecrecover(msgHash, v, r, s) == _addr;
    }
}
