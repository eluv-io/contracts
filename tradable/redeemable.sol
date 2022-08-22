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

    function addRedeemable() public ownerOnly {
        emit RedeemableAdded();
    }

    function removeRedeemable() public ownerOnly {
        emit RedeemableRemoved();
    }

    function redeem() {
        emit Redeem();
    }

}
