pragma solidity 0.5.4;

library ExternalCall {

    address constant elv_precomp_extern_call = address(253);

    bytes4 constant sigIdString = bytes4(keccak256("callString(string,address,uint256,bytes)"));
    bytes4 constant sigIdUint = bytes4(keccak256("callUint(string,address,uint256,bytes)"));
    bytes4 constant sigIdAddress = bytes4(keccak256("callAddress(string,address,uint256,bytes)"));

    function callUint(string memory _config, address _extAddr, uint256 _blockNum, bytes memory _params) internal view returns (uint256 ret) {
        bytes4 sig = sigIdUint;
        address callAddr = elv_precomp_extern_call;
        bytes32 truncConfig;

        uint paramsLen = _params.length;
        uint allLen = 4 + 96 + paramsLen; // 4bytes + truncated config + address + block num + params bytes length

        assembly {
            let x := mload(0x40)        // Find empty storage location using "free memory pointer"
            mstore(x, sig)              // Place signature at beginning of empty storage

            truncConfig := mload(add(_config, 32))
            mstore(add(x, 4), truncConfig)
            mstore(add(x, 36), _extAddr)
            mstore(add(x, 68), _blockNum)

            for { let i := 0 } lt(i, paramsLen) { i := add(i, 32) } {
                mstore(add(x, add(100,i)), mload(add(add(_params, 32), i)))
            }

            let res := staticcall(0, callAddr, x, allLen, x, 32)

            switch res
            case 0 {
                revert(0, 0)
            }
            default {
                ret := mload(x)
            }
        }
    }

}