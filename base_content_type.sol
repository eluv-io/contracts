pragma solidity 0.4.24;

import {Accessible} from "./accessible.sol";
import {Editable} from "./editable.sol";
import "./accessible.sol";
import "./base_content_space.sol";
import "./access_indexor.sol";


/* -- Revision history --
BaseContentType20190222145700ML: First versioned released
BaseContentType20190318101200ML: Migrated to 0.4.24
BaseContentType20190506153900ML: Adds access indexing
*/


contract BaseContentType is Accessible, Editable {

    bytes32 public version ="BaseContentType20190506153900ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    address public contentSpace;

    constructor(address content_space) public payable {
        contentSpace = content_space;
    }

    function canPublish() public view returns (bool) {
        return msg.sender == owner;
    }


    function setRights(address stakeholder, uint8 access_type, uint8 access) public {
        BaseContentSpace contentSpaceObj = BaseContentSpace(contentSpace);
        address walletAddress = contentSpaceObj.userWallets(stakeholder);
        if (walletAddress == 0x0){
            //stakeholder is not a user (hence group or wallet)
            setGroupRights(stakeholder, access_type, access);
        } else {
            setGroupRights(walletAddress, access_type, access);
        }
    }

    function setGroupRights(address group, uint8 access_type, uint8 access) public {
        AccessIndexor indexor = AccessIndexor(group);
        indexor.setContentObjectRights(address(this), access_type, access);
    }
}
