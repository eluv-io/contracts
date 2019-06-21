pragma solidity 0.4.24;

/**
 * UserSpace
 * Seperated from content_space to avoid circular dependencies.
 */

/* -- Revision history --
UserSpace20190506155300ML: First versioned released
*/
contract UserSpace {
    //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX
    bytes32 public version ="UserSpace20190506155300ML";

    mapping(address => address) public userWallets;

}
