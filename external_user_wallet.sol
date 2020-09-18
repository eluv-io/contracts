//SPDX-License-Identifier: MIT
pragma solidity ^0.7.1;

import {Ownable} from "./ownable.sol";
import {Editable} from "./editable.sol";

contract ExternalUserWallet is Editable {

    address payable public addressTA; // trusted authority of the user - aka KMS
    address payable public addressExtUser;

    event CreateExtUserWallet(address contentSpace, address extUserAddr);

    constructor(address payable _content_space, address payable _taAddr, address payable _extUserAddr) payable {
        contentSpace = _content_space;
        addressTA = _taAddr;
        addressExtUser = _extUserAddr;
        emit CreateExtUserWallet(_content_space, _extUserAddr);
    }

    function claimOwnership() public {
        require(msg.sender == addressExtUser);
        owner = msg.sender;
    }

    // overloads ...
    function canEdit() public view override returns (bool) {
        return (msg.sender == owner || msg.sender == addressExtUser);
    }

    function canCommit() public view override returns (bool) {
        return (msg.sender == owner || msg.sender == addressExtUser);
    }

    function accessRequest() public returns (bool) {
        require(msg.sender == owner || msg.sender == addressExtUser);
        emit AccessRequestV3(uint256(keccak256(abi.encodePacked(address(this), block.timestamp))), address(0x0), 0x0, msg.sender, block.timestamp * 1000);
        return true;
    }
}