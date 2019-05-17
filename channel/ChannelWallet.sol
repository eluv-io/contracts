pragma solidity ^0.4.24;

import "./Transactable.sol";
import "./BatchTransaction.sol";

// abigen --sol channel/ChannelWallet.sol --pkg=contracts --out build/channel_wallet.go

// TODO: need a way to get the locator of the guarantor ?

contract ChannelWallet is Transactable {

    // the most recent timestamp that has been recorded for this channel
    // this only defines a flow of valid timestamps for the channel since it is the last timestamp that was recorded
    //  to the blockchain.
    uint256 public currentTimestamp;

    address public owner;
    address public guarantor;

    constructor(address _guarantor) public payable {
        owner = msg.sender;
        guarantor = _guarantor;
        currentTimestamp = now;
    }

    // TODO: not sure if this is necessary - there should be a default accessor for currentTimestamp?
    function validateTimestamp(uint256 _ts) public view returns (bool) {
        if (_ts > currentTimestamp)
            return true;
        return false;
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

    // TODO: self destruct, etc.
    // there will need to be a method for the owner to reclaim their funds. there will need to be a mandatory
    //  delay of sufficient length that it allows the guarantor to finalize any outstanding chits.

    event ExecStatus(address guarantor, int code);

    int public constant execStatusOk = 0;
    int public constant execStatusNonceFail = 1;
    int public constant execStatusBalanceFail = 2;
    int public constant execStatusSigFail = 3;
    int public constant execStatusSendFail = 4;

    function execute(uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts)
    external
    returns (bool) {

        require(msg.sender == guarantor); // only the guarantor can resolve transactions / chits

        if (_ts <= currentTimestamp) {
            emit ExecStatus(guarantor, execStatusNonceFail);
            return false;
        }

        if (address(this).balance < _value) {
            emit ExecStatus(guarantor, execStatusBalanceFail);
            return false;
        }

        // this check should not really fail because the guarantor should have at least validated the transaction before
        //  accepting it into a batch.
        bool checkTrans = validateTransaction(_v, _r, _s, _dest, _value, _ts);
        if (!checkTrans) {
            emit ExecStatus(guarantor, execStatusSigFail);
            return false;
        }

        currentTimestamp = _ts;

        // TODO: for a content access request there might be other data we want to pass ...?
        bool sent = _dest.send(_value);
        if (!sent) {
            emit ExecStatus(guarantor, execStatusSendFail);
            return false;
        }

        emit ExecStatus(guarantor, execStatusOk);

        return true;
    }
}