pragma solidity 0.4.24;

/* -- Revision history --
Accessible20190222135900ML: First versioned released
Accessible20190315141600ML: Migrated to 0.4.24
Accessible20200107160900ML: Moved Visibility filter from BaseContentObject to Accessible
*/

contract Accessible {

    bytes32 public version ="Accessible20200107160900ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    uint8 public visibility = 0;
    uint8 public constant CAN_SEE = 1;
    uint8 public constant CAN_ACCESS = 10;
    uint8 public constant CAN_EDIT = 100;

    event AccessRequest();

    function accessRequest() public returns (bool) {
        emit AccessRequest();
        return true;
    }

}
