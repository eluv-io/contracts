pragma solidity ^0.4.24;

// abigen --sol channel/TestElvPrecompile.sol --pkg=contracts --out build/elv_precompile.go

contract TestElvPrecompile {

    address public elv_precomp_addr = 255;

    event GetElvData(bytes32 data);

    function testPrecompile() public {

        int ord = 0;
        bool done = false;

        while (!done) {
            bytes32 ret = getTrans(ord, 0);
            if (ret[0] == 0) {
                done = true;
            } else {
                emit GetElvData(ret);
                ord++;
            }
        }
    }

    bytes4 precompSig = bytes4(keccak256("f()")); // TODO: signature is not important for pre-compile

    // input will be:
    // . uint256 identifier for transaction 'batch'
    // . int ordinal of transaction from batch
    // output will be:
    // . bytes of transaction? can we use zero length to indicate end of batch?
    function getTrans(int transOrd, uint256 batchId) constant returns (bytes32 ret) {

        // these have to be locals (?) so create copy
        bytes4 sig = precompSig; // TODO: pretty sure we don't care about the signature ...?
        bytes4 ordBytes = bytes4(transOrd);
        address callAddr = elv_precomp_addr;

        assembly {
            let x := mload(0x40)            // Find empty storage location using "free memory pointer"
            mstore(x, sig)                  // Place signature at beginning of empty storage
            mstore(add(x, 0x04), ordBytes)  // Place first argument directly next to signature
            mstore(add(x, 0x08), batchId)   // Place second argument next to first, (NOT padded to 32 bytes?)
            // input is now [4 bytes][4 bytes][32 bytes] = 40 bytes?

            let res := call(
                0,          // no gas (?)
                callAddr,   // To addr
                0,          // No value - i.e. tokens
                x,          // Inputs are stored at location x
                0x2C,       // Inputs are 44 bytes long
                x,          // Store output over input (saves space)
                0x20)       // Outputs are 32 bytes long

            if eq(res, 0) {
                revert(0, 0)
            }

            ret := mload(x)             // Assign output value to ret
            mstore(0x40, add(x, 0x20))  // Set storage pointer to empty space - i.e. after returned storage
        }
    }
}