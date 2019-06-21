pragma solidity 0.4.24;

import {Ownable} from "./ownable.sol";
import {Accessible} from "./accessible.sol";
import {Container} from "./container.sol";
import {BaseContent} from "./base_content.sol";
import {BaseContentSpace} from "./base_content_space.sol";
import "./access_indexor.sol";
import "./transactable.sol";

/* -- Revision history --
BaseAccessWallet20190320114000ML: First versioned released
BaseAccessWallet20190506154400ML: Adds instantiation via factory, adds access indexing
BaseAccessWallet20190510151100ML: Supports modified getAccessInfo API
BaseAccessWallet20190528124400ML: Change base to be Accessible and Editable (thru Container)
*/

// abigen --sol base_access_wallet.sol --pkg=contracts --out build/base_access_wallet.go
contract BaseAccessWallet is Accessible, Container, AccessIndexor, Transactable {

    int public constant EXEC_STATUS_OK = 0;
    int public constant EXEC_STATUS_NONCE_FAIL = 1;
    int public constant EXEC_STATUS_BALANCE_FAIL = 2;
    int public constant EXEC_STATUS_SIG_FAIL = 3;
    int public constant EXEC_STATUS_SENDFAIL = 4;

    //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX
    bytes32 public version = "BaseAccessWallet20190611120000PO";

    // WIP - state channel support
    // the most recent timestamp that has been recorded for this channel
    // this only defines a flow of valid timestamps for the channel since it is the last timestamp that was recorded
    //  to the blockchain.
    uint256 public currentTimestamp;

    event ExecStatus(address guarantor, int code);

    constructor(address content_space)  public payable {
        contentSpace = content_space;
    }

    function execute(address _guarantor, uint8 _v, bytes32 _r, bytes32 _s, address _dest, uint256 _value, uint256 _ts)
    external
    returns (bool) {

        BaseContentSpace spc = BaseContentSpace(contentSpace);
        require(msg.sender == contentSpace || spc.checkKMSAddr(msg.sender) > 0);
        require(spc.checkKMSAddr(_guarantor) > 0);

        if (_ts <= currentTimestamp) {
            emit ExecStatus(_guarantor, EXEC_STATUS_NONCE_FAIL);
            return false;
        }

        if (address(this).balance < _value) {
            emit ExecStatus(_guarantor, EXEC_STATUS_BALANCE_FAIL);
            return false;
        }

        // this check should not really fail because the _guarantor should have at
        // least validated the transaction before accepting it into a batch.
        bool checkTrans = validateTransaction(_v, _r, _s, _dest, _value, _ts);
        if (!checkTrans) {
            emit ExecStatus(_guarantor, EXEC_STATUS_SIG_FAIL);
            return false;
        }

        currentTimestamp = _ts;

        // TODO: for a content access request there might be other data we want to pass ...?
        bool sent = _dest.send(_value);
        if (!sent) {
            emit ExecStatus(_guarantor, EXEC_STATUS_SENDFAIL);
            return false;
        }

        emit ExecStatus(_guarantor, EXEC_STATUS_OK);

        return true;
    }

    function accessRequestMsg(
        address content_address,
        uint8 level,
        string pke_requestor,
        string pke_AFGH,
        bytes32[] custom_values,
        address[] stakeholders
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(content_address, level, pke_requestor, pke_AFGH,
        custom_values, stakeholders));
    }

    function accessCompleteMsg(
        address content_address,
        uint256 request_ID,
        uint256 score_pct,
        bytes32 ml_out_hash
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(content_address, request_ID, score_pct, ml_out_hash));
    }

    function contentAccessRequest(
        address content_address,
        bytes, /*signature*/
        uint8 level,
        string pke_requestor,
        string pke_AFGH,
        bytes32[] custom_values,
        address[] stakeholders
    ) public returns (uint256) {
        /*
        //Signature should be valid for requested operation
        bytes32 message = accessRequestMsg(content_address,  level, pke_requestor, pke_AFGH,
        custom_values, stakeholders);
        require(owner == recoverSignerFromMessage(message,signature));
        */

        BaseContent content = BaseContent(content_address);
        uint256 requiredFund;
        uint8 visibilityCode;
        uint8 accessCode;
        (visibilityCode, accessCode, requiredFund) = content.getAccessInfo(level, custom_values, stakeholders);
        require(visibilityCode == 0);
        uint256 requestID = content.accessRequest.value((accessCode != 0) ? requiredFund : 0)
        (level, pke_requestor, pke_AFGH, custom_values, stakeholders);
        return requestID;

    }

    function contentAccessComplete(
        address content_address,
        bytes,/*signature*/
        uint256 request_ID,
        uint256 score_pct,
        bytes32 ml_out_hash
    ) public payable returns (bool) {
        /*
        //Signature should be valid for requested operation
        bytes32  message = accessCompleteMsg(content_address, request_ID, score_pct, ml_out_hash);
        require(owner == recoverSignerFromMessage(message,signature));
        */

        BaseContent content = BaseContent(content_address);
        return content.accessComplete(request_ID, score_pct, ml_out_hash);

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

}

/* -- Revision history --
BsAccWltFactory20190506154200ML: First versioned released
*/

contract BaseAccessWalletFactory is Ownable {

    //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX
    bytes32 public version ="BsAccWltFactory20190506154200ML";

    function createAccessWallet() public returns (address) {
        return  (new BaseAccessWallet(msg.sender));
    }
}
