pragma solidity 0.4.24;

import {Ownable} from "./ownable.sol";
import {UserSpace} from "./user_space.sol";
import "./node.sol";
import {AccessIndexor} from "./access_indexor.sol";
import {BaseContentType} from "./base_content_type.sol";



/* -- Revision history --
BaseFactory20190227170400ML: First versioned released
BaseFactory20190301105700ML: No changes version bump to test
BaseFactory20190319195000ML: with  0.4.24 migration
BaseFactory20190506153000ML: Split createLibrary out, adds access indexing
BaseFactory20190722161600ML: No changes, updated to provide generation for BsAccessCtrlGrp20190722161600ML
BaseFactory20190801140700ML: Removed access group creation to its own factory
*/

contract BaseFactory is Ownable {

    bytes32 public version ="BaseFactory20190801140700ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    function createContentType() public returns (address) {
        address newType = (new BaseContentType(msg.sender));
        // register library in user wallet
        UserSpace contentSpaceObj = UserSpace(msg.sender);
        address walletAddress = contentSpaceObj.getAccessWallet();
        AccessIndexor userWallet = AccessIndexor(walletAddress);
        userWallet.setContentTypeRights(newType, userWallet.TYPE_EDIT(), userWallet.ACCESS_CONFIRMED());
        return newType;
    }

    function createNode(address _owner) public returns (address) {
        Node n = new Node(); // this sets owner to tx.origin?
        require(n.owner() == _owner);
        return address(n);
    }
}
