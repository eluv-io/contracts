pragma solidity 0.4.24;

import {Factory} from "./factory.sol";
import {Ownable} from "./ownable.sol";
/**
 * UserSpace
 * Seperated from content_space to avoid circular dependencies.
 */

/* -- Revision history --
UserSpace20190506155300ML: First versioned released
*/


contract UserSpace {

    bytes32 public version ="UserSpace20190506155300ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    mapping(address => address) public userWallets;

    event CreateAccessWallet(address wallet);
    event GetAccessWallet(address walletAddress);
    address public walletFactory;


    function createAccessWallet() public returns (address) {
        return createUserWallet(tx.origin);
    }

    // TODO: TESTING
    //    function createUserGuarantorWallet(address _user) public returns (bool) {
    //        if (userWallets[_user] != 0x0) {
    //            return false;
    //        }
    //        address walletAddress = BaseAccessWalletFactory(walletFactory).createAccessWallet();
    //        BaseAccessWallet wallet = BaseAccessWallet(walletAddress);
    //        // wallet.setGuarantor(msg.sender);
    //        wallet.transferOwnership(_user);
    //        userWallets[_user] = walletAddress;
    //        emit CreateAccessWallet(walletAddress); // TODO: different event here?
    //        return true;
    //    }

    //This methods revert when attempting to transfer ownership, so for now we make it private
    // Hence it will be assumed, that user are responsible for creating their wallet.
    function createUserWallet(address user) private returns (address) {
        require(userWallets[user] == 0x0);
        address walletAddress = Factory(walletFactory).create();
        if (user != tx.origin) {
            Ownable wallet = Ownable(walletAddress);
            wallet.transferOwnership(user);
        }
        emit CreateAccessWallet(walletAddress);
        userWallets[user] = walletAddress;
        return walletAddress;
    }

    function getAccessWallet() public returns(address) {
        address walletAddress;
        if (userWallets[tx.origin] == 0x0) {
            walletAddress = createAccessWallet();
        } else {
            walletAddress = userWallets[tx.origin];
        }

        emit GetAccessWallet(walletAddress);
        return walletAddress;
    }

    /* removed as the createUserWallet does not work for creating wallet on behalf of a user
    // Not sure we want that, if so it might have to be restricted -- to be thought through
    function getUserWallet(address user) public returns(address) {
        if (userWallets[user] == 0x0) {
            return createUserWallet(user);
        } else {
            return userWallets[user];
        }
    }
    */

}
