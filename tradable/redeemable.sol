// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

/**
 * @title Redeemable
 *
 *
 */
contract Redeemable {

    event RedeemableAdded();
    event RedeemableRemoved();
    event Redeem();

    modifier ownerOnly(){
        _;
    }

    function addRedeemable() public ownerOnly returns (uint){
        emit RedeemableAdded();
        return 10;
    }

    function removeRedeemable() public ownerOnly {
        emit RedeemableRemoved();
    }

    function redeem() public {
        emit Redeem();
    }

}
