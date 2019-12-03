pragma solidity 0.4.24;


/**
 * UserSpace
 * Seperated from content_space to avoid circular dependencies.
 */

/* -- Revision history --
UserSpace20190506155300ML: First versioned released
*/

interface UserSpace {
    function getUserWallet(address _userAddr) external view returns (address);
}

contract UserSpaceImpl is UserSpace {

    bytes32 public version ="UserSpace20190506155300ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    mapping(address => address) userWallets;

    function getUserWallet(address _userAddr) public view returns (address) {
        return userWallets[_userAddr];
    }

}
