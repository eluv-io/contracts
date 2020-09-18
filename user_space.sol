//SPDX-License-Identifier: MIT
pragma solidity ^0.7.1;

import "./base_space_interfaces.sol";

/**
 * UserSpace
 * Seperated from content_space to avoid circular dependencies.
 */

/* -- Revision history --
UserSpace20190506155300ML: First versioned released
*/

contract UserSpace is IUserSpace {

    bytes32 public version ="UserSpace20190506155300ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX


    mapping(address => address) public userWalletsMap;

    function userWallets(address payable _userAddr) external view override returns (address payable) {
        return payable(userWalletsMap[_userAddr]);
    }

    // STUB impl - meant to be overridden
    function createUserWallet(address payable _user) external override returns (address payable) {
        require(false);
        return address(0x0);
    }
}