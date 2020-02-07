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

    event AccessRequestV3(
        bytes32 requestNonce,
        string contentHash,
        address libraryAddress, // likely will need for tenancy - but could be something else ...?
        bytes32 contextHash,
        address accessor,           // I've called this 'userAddress' - don't care too much but ideally it's the same name wherever it means the same thing!
        uint256 request_timestamp   // always millis - either set from context or with blockchain now()
    );

    function accessRequestV3(
        bytes32[] custom_values,
        address[] stakeholders
    ) public returns (bool) { //to be overloaded
        emit AccessRequestV3(0, "", 0x0, 0x0, msg.sender, now * 1000);
        return true;
    }
}
