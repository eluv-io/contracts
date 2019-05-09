pragma solidity ^0.4.24;

import "./Transactable.sol";
import "../ownable.sol";

// abigen --sol channel/BatchTransaction.sol --pkg=contracts --out build/batch_trans.go

contract BatchTransaction is Ownable {

    event TransactionStatus(bool success, address from, address dest);

    constructor() public payable {}

    function () public payable {}

    // abi.encodePacked(address(this), _dest, _value, _ts)
    function executeBatch(uint8[] _v, bytes32[] _r, bytes32[] _s, address[] _from, address[] _dest, uint256[] _value, uint256[] _ts) public onlyOwner {
        // TODO: not sure if this is worth it - will just crash if the parameters are passed in incorrectly, which is the same as a revert...?
        require(_v.length == _r.length);
        require(_r.length == _s.length);
        require(_s.length == _from.length);
        require(_from.length == _dest.length);
        require(_dest.length == _value.length);
        require(_value.length == _ts.length);

        for (uint i = 0; i < _v.length; i++) {
            Transactable t = Transactable(_from[i]);
            bool success = t.execute(_v[i], _r[i], _s[i], _dest[i], _value[i], _ts[i]);

            if (!success) {
                // we failed to get the target wallet to pay so - in this scenario - we have to pay!
                // _dest[i].send(_value[i]); // TODO: transfer? error handling?
            }
            emit TransactionStatus(success, _from[i], _dest[i]);
        }
    }
}
