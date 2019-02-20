pragma solidity ^0.4.24;

import "./openzeppelin-solidity/Arrays.sol";
import "./Transactable.sol";

// abigen --sol channel/ElvWallet.sol --pkg=contracts --out build/elv_wallet.go

contract ElvWallet is Transactable {

    uint constant max_timestamps = 10;

    using Arrays for uint256[];

    uint256[] private timestamps;
    address public owner;

    constructor(address _owner) public payable {
        owner = _owner;
    }

    // should be called when transaction is originally submitted as a sanity check for validity
    function validateTimestamp(uint256 _ts) public view returns (bool, uint256) {
        uint256 idx = timestamps.findUpperBound(_ts);
        if (idx < timestamps.length) {
            // timestamp is 'within' range of array - check to see if it already exists or is less than range
            if (timestamps[idx] >= _ts) {
                return (false, 0); // new timestamp either already exists of is outside (below) range we've recorded
            }
        }
        return (true, idx);
    }

    function validateTransaction(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts)
    public
    view
    returns (bool) {
        bytes32 checkHash = keccak256(abi.encodePacked(address(this), _dest, _value, _ts));
        address checkAddr = ecrecover(checkHash, _v, _r, _s);
        if (checkAddr != owner) return false;
        return true;
    }

    function shiftTimestampsDown(uint256 fromIdx) private {
        for (uint256 i = 0; i < fromIdx; i++) {
            timestamps[i] = timestamps[i + 1];
        }
    }

    function shiftTimestampsUp(uint256 fromIdx) private {
        for (uint256 i = timestamps.length - 1; i < fromIdx; i--) {
            timestamps[i] = timestamps[i - 1];
        }
    }

    // TODO: self destruct, owner, creator, etc.
    // I think we want to let both the owner and creator destroy the contract - but then send any funds back to the owner

    event ExecStatus(int code);

    // Note that address recovered from signatures must be strictly increasing
    function execute(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts)
    external
    returns (bool) {

        if (address(this).balance < _value) {
            emit ExecStatus(1);
            return false;
        }

        // this check should not really fail because the guarantor should have at least validated the transaction before
        //  accepting it into a batch.
        bool checkTrans = validateTransaction(_v, _r, _s, _dest, _value, _ts);
        if (!checkTrans) {
            emit ExecStatus(2);
            return false;
        }

//        (bool valid, uint256 idx) = validateTimestamp(_ts);
//        if (!valid) {
//            emit ExecStatus(3);
//            return false;
//        }

        // TODO: for a content access request there might be other data we want to pass ...?
        bool sent = _dest.send(_value);
        if (!sent) {
            emit ExecStatus(4);
            return false;
        }

//        if (timestamps.length < max_timestamps) {
//            timestamps.length++;
//            shiftTimestampsUp(idx);
//        } else {
//            shiftTimestampsDown(idx);
//        }
//        timestamps[idx] = _ts;

        // emit ExecStatus(0);
        return true;
    }
}