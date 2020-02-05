pragma solidity 0.4.24;

import {Content} from "./content.sol";
import {BaseContent} from "./base_content.sol";
import {BaseContentSpace} from "./base_content_space.sol";
import {BaseAccessControlGroup} from "./base_access_control_group.sol";
import {AccessIndexor} from "./access_indexor.sol";


/* -- Revision history --
SmplScreener20191202201500ML: First versioned released
SmplScreener20191204144100ML: Fixes the logic of access to check first for direct access and allows non editor access
SmplScreener20200129174600ML: Allows editor from screener manager objec to modify avails

*/

contract SampleScreener is Content {

    bytes32 public version ="SmplScreener20200129174600ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    event SetAvailability(address content, uint256 start, uint256 end, uint8 status);

    struct Avail {
        uint256 start; //0 indicates no limit
        uint256 end; //0 indicates no limit
        uint8 status; //0 indicates disabled
    }
    mapping(address => Avail) public availability;

    address public screeners; //screener group

    address public screenerManager;

    address[] public assets;
    uint256 public assetsLength = 0;

    function setScreeners(address group) public onlyEditor {
        screeners = group;
    }

    modifier onlyEditor() {
        BaseContent content  = BaseContent(screenerManager);
        require((tx.origin == owner) || content.canEdit());
        _;
    }

    function setAvailability(address content, uint256 start, uint256 end) public onlyEditor {
        //Could be changed to only allow content owner instead of screener contract owner to set availability
        availability[content].start = start;
        availability[content].end = end;
        availability[content].status = 1;
        emit SetAvailability(content, start, end, 1);
    }

    function makeUnavailable(address content) public onlyEditor {
         availability[content].status = 0;
        emit SetAvailability(content, availability[content].start, availability[content].end, 0);
    }

    function makeAvailable(address content) public onlyEditor {
        availability[content].status = 1;
        emit SetAvailability(content, availability[content].start, availability[content].end, 1);
    }

    function runCreate() public payable returns (uint) {
        if (screenerManager == 0x0) {
            screenerManager = msg.sender;
        } else {
            availability[msg.sender] = Avail(0, 0, 0);
            assets.push(msg.sender);
            assetsLength = assetsLength+1;
        }
        return 0;
    }


    function runKill() public payable returns (uint) {
        delete availability[msg.sender];
        return 0;
    }


    function runAccessInfo( //Old API
        uint8, /*level*/
        bytes32[], /*customValues*/
        address[] /*stakeholders*/
    )
    public view returns (uint8, uint8, uint8, uint256) //Mask, visibilityCode, accessCode, accessCharge
    {
        BaseContent content  = BaseContent(msg.sender);
        //First check if user has been granted explicit right to the object
        address walletAddress =  BaseContentSpace(content.contentSpace()).userWallets(tx.origin);
        AccessIndexor wallet = AccessIndexor(walletAddress);
        if (wallet.checkContentObjectRights(msg.sender, wallet.TYPE_ACCESS())) {
            return (0, 0, 0 ,0); // Object is accessible
        }

        BaseAccessControlGroup group = BaseAccessControlGroup(screeners);
        if (group.hasAccessRight(tx.origin, false)) {
            if ((availability[msg.sender].status != 0) && (availability[msg.sender].start <= now) && ((availability[msg.sender].end == 0) || (availability[msg.sender].end >= now))) {
                return (0, 0, 0 ,0); // no charges
            } else {
                return (0, 2, 2, 0); //outside of availability
            }
        }
        return (0, 1, 1, 0); //Bar access to unauthorized user
    }

    /*  // not needed for now, might be required if special reporting is added
    //0 indicates that access request can proceed.
    // Other numbers can be used as error codes and would stop the processing.
    function runAccess(
        uint256 request_ID,
        uint8 level,
        bytes32[] custom_values,
        address[] stake_holders
    )
    public payable returns(uint)
    {
        return 0;
    }
    */


    /*
    function runAccessInfo( //new API
        address accessor,
        uint8 level,
        bytes32[] customValues,
        address[] stakeholders
    )
    public view returns (uint256[4]) //Mask, visibilityCode, accessCode, accessCharge
    {
        return [uint256(7), 0, 0, 0]; //7 is DEFAULT_SEE + DEFAULT_ACCESS + DEFAULT_CHARGE, hence the 3 tailing values are ignored
    }
    */


    /*
    //0 indicates that access request can proceed.
    // Other numbers can be used as error codes and would stop the processing.
    function runAccess( //new API
        bytes32 requestNonce,
        address accessor,
        uint8 level,
        uint256 value,
        bytes32[] custom_values,
        address[] stake_holders
    )
    public payable returns(uint)
    {
        return 0;
    }
    */




}
