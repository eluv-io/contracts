pragma solidity 0.4.24;


library Precompile {

    address public constant ELV_PRECOMP_ADDR = 255;
    bytes4 public constant SIG_ID_STRING = bytes4(keccak256("makeIDString(int,address)"));
    uint public constant ID_STR_LEN = 32;
    int public constant KMS = 11;

    function codeKMS() internal constant returns (int) {
        return KMS;
    }

    function makeIDString(int _code, address _addr) internal constant returns (string ret) {

        bytes4 sig = SIG_ID_STRING;
        bytes4 codeBytes = bytes4(_code);
        address callAddr = ELV_PRECOMP_ADDR;

        assembly {
            let x := mload(0x40)
            mstore(x, sig)
            mstore(add(x, 0x04), codeBytes)
            mstore(add(x, 0x08), _addr)
            // input is now [4 bytes][4 bytes][32 bytes] = 40 bytes

            let res := call(
            0,          // no gas
            callAddr,   // To addr
            0,          // No value - i.e. tokens
            x,          // Inputs are stored at location x
            0x28,       // Inputs are 40 bytes long
            x,          // Store output over input
            0x60)       // Outputs are 96 bytes long

            if eq(res, 0) {
                revert(0, 0)
            }

            ret := x
            mstore(0x40, add(x, 0x60))  // Set storage pointer to empty space - i.e. after returned storage
        }
    }
}
