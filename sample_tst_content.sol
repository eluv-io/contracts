pragma solidity 0.4.24;

import {Ownable} from "./ownable.sol";
import {Content} from "./content.sol";

/* -- Revision history --
TstContent20200113172800ML: First versioned released
*/

contract TstContent is Content {

    bytes32 public version = "TstContent20200113172800ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    mapping (int => bytes32) public statusCodeDescription;
    uint public runCreateCode = 0;
    uint public runEditCode = 100;
    uint public runKillCode = 0;
    int public runStatusChangeCode=0;
    uint public runAccessCode = 0;
    uint public runGrantCode = 0;
    uint public runFinalizeCode = 0;

    uint8 public maskCode = 7;
    uint8 public visibilityCode=0;
    uint8 public accessCode=0;
    uint256 public accessCharge=0;

    function runDescribeStatus(int code) public view returns (bytes32) {
        return statusCodeDescription[code];
    }

    function setStatusCodeDescription(int code, bytes32 description) public onlyOwner {
	statusCodeDescription[code] = description;
    }

    function setMaskCode(uint8 code) public onlyOwner {
 	    maskCode = code;
    }

    function setVisibilityCode(uint8 code) public onlyOwner {
        visibilityCode = code;
    }

    function setAccessCode(uint8 code) public onlyOwner {
        accessCode = code;
    }

    function setAccessCharge(uint256 code) public onlyOwner {
        accessCharge = code;
    }

    //0 indicates that the creation can proceed.
    // Other numbers can be used as error codes and would stop the processing.
    function runCreate() public payable returns (uint) {
        return runCreateCode;
    }
 
    function setRunCreateCode(uint code) public onlyOwner {
	    runCreateCode = code;
    }


   //0 indicates edit can proceed
   //100 indicates that custom contract does not modify default behavior
   function runEdit() public returns (uint) {
       return runEditCode;
   }

    function setRunEditCode(uint code) public onlyOwner {
        runEditCode = code;
    }

    //0 indicates that normal behavior should apply
    //100 indicates that normal behavior should apply and the custom contract should be killed too
    //1000 indicates that the deletion/inactivation can proceed without further validations
    //1100 indicates that the deletion can proceed without further validations and the custom contract should be killed too
    // Other numbers can be used as error codes and would stop the processing.
    function runKill() public payable returns (uint) {
        return runKillCode;
    }

    function setRunKillCode(uint code) public onlyOwner { 
        runKillCode = code;
    }

    // a negative number returned indicates that the licending fee to be paid is the default
    function runStatusChange(int proposed_status_code) public payable returns (int) {
        return runStatusChangeCode;
    }

    function setRunStatusChangeCode(int code) public onlyOwner {
        runStatusChangeCode = code;
    }


    function runAccessInfo(
        uint8, /*level*/
        bytes32[], /*customValues*/
        address[] /*stakeholders*/
    )
    public view returns (uint8, uint8, uint8, uint256) //Mask, visibilityCode, accessCode, accessCharge
    {
        return (maskCode, visibilityCode, accessCode, accessCharge); //7 is DEFAULT_SEE + DEFAULT_ACCESS + DEFAULT_CHARGE, hence the 3 tailing values are ignored
    }


    //0 indicates that access request can proceed.
    // Other numbers can be used as error codes and would stop the processing.
    function runAccess(
        uint256, /*request_ID*/
        uint8, /*level*/
        bytes32[], /*custom_values*/
        address[] /*stake_holders*/
    )
        public payable returns(uint)
    {
            return runAccessCode;
    }

    function setRunAccessCode(uint code) public onlyOwner {
        runAccessCode = code;
    }

    //0 indicates that access grant can proceed.
    // Other numbers can be used as error codes and would stop the processing.
    function runGrant(
        uint256, /*request_ID */
        bool /*access_granted */
    )
    public payable returns (uint)
    {
        return runGrantCode;
    }

    function setRunGrantCode(uint code) public onlyOwner {
        runGrantCode = code;
    }

    //the status is logged in an event at the end of the accessComplete function
    // behavior is currently unchanged regardless of result.
    // 0 indicates that the finalization can proceed.
    // Other numbers can be used as error codes and would stop the processing.
    function runFinalize(uint256 /*request_ID*/, uint256 /*score_pct*/) public payable returns (uint) {
        return runFinalizeCode;
    }

    function runFinalizeExt(uint256 requestID, uint256 score_pct, address originator) public payable returns (uint) {
        return runFinalizeCode;
    }

    function setRunFinalizeCode(uint code) public onlyOwner {
        runFinalizeCode = code;
    }
}
  
contract TstContentFactory is Ownable {    
   
    event NewTstContent(address instanceAddress);

    bytes32 public version = "TstContentFctry20100113182100ML";

    function instantiate() public onlyOwner returns (address) {
	address instanceAddress = new TstContent();
        emit NewTstContent(instanceAddress);
        return instanceAddress;
    }

}

