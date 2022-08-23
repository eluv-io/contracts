// SPDX-License-Identifier: MIT
pragma solidity ^0.6.12;

import "forge-std/Test.sol";
import "forge-std/console2.sol";
import "src/redeemable.sol";


contract RedeemableTest is Test {
    Redeemable redeemable;
    address creator = 0x7109709ECfa91a80626fF3989D68f67F5b1DD12D;
    address minter = 0x2233709ecfa91A80626Ff3989d68f67F5B1dD12d;

    event RedeemableAdded(uint tokenId);

    /**
        Set up function, it is the first function executed when we run 'forge test'
    */
    function setUp() public {
        vm.prank(creator, creator); //adr is the creator and the owner of this smart contract
        redeemable = new Redeemable();
    }

    function testAdminRedeemable() public{

        vm.prank(minter);
        uint res = redeemable.addRedeemableOffer();
        assertEq(res, 0, "result != 0");

        res = redeemable.addRedeemableOffer();
        assertEq(res, 1, "result != 1");

        redeemable.removeRedeemableOffer(0);

        vm.expectRevert(bytes("offer not active"));
        redeemable.removeRedeemableOffer(0);

    }

    function testRedeem() public{

        vm.expectRevert(bytes("offer not active"));
        redeemable.redeemOffer(1000, 0);

        vm.prank(minter);
        uint res = redeemable.addRedeemableOffer();
        assertEq(res, 0, "result != 0");

        redeemable.redeemOffer(1000, 0);

        vm.expectRevert(bytes("offer already redeemed"));
        redeemable.redeemOffer(1000, 0);

        vm.expectRevert(bytes("offer already redeemed"));
        redeemable.redeemOffer(1000, 0);

        redeemable.redeemOffer(1001, 0);

    }

}
