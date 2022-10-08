// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "forge-std/console2.sol";
import "src/tradable/Redeemable.sol";

contract RedeemableTest is Test {
    Redeemable redeemable;
    address creator = 0x7109709ECfa91a80626fF3989D68f67F5b1DD12D;
    address minter = 0x2233709ecfa91A80626Ff3989d68f67F5B1dD12d;
    address user = 0x5372642648d93315f121dEa8B5b0E3568A894d94;

    event RedeemableAdded(uint256 tokenId);

    /**
     * Set up function, it is the first function executed when we run 'forge test'
     */
    function setUp() public {
        vm.prank(creator, creator); //adr is the creator and the owner of this smart contract
        redeemable = new Redeemable();
    }

    function testAdminRedeemableBasics() public {
        vm.prank(creator, creator);
        redeemable.grantRole(ElvCmnRoles.MINTER_ROLE, minter);

        vm.startPrank(minter, minter);
        uint256 res = redeemable.addRedeemableOffer();
        assertEq(res, 0, "result != 0");

        res = redeemable.addRedeemableOffer();
        assertEq(res, 1, "result != 1");

        redeemable.removeRedeemableOffer(0);

        vm.expectRevert(bytes("offer not active"));
        redeemable.removeRedeemableOffer(0);

        vm.stopPrank();
    }

    function testRedeemBasics() public {
        vm.expectRevert(bytes("offer not active"));
        redeemable.redeemOffer(user, 1000, 0);

        vm.prank(creator, creator);
        redeemable.grantRole(ElvCmnRoles.MINTER_ROLE, minter);

        vm.startPrank(minter, minter);
        uint256 res = redeemable.addRedeemableOffer();
        assertEq(res, 0, "result != 0");

        redeemable.redeemOffer(user, 1000, 0);

        vm.expectRevert(bytes("offer already redeemed"));
        redeemable.redeemOffer(user, 1000, 0);

        vm.expectRevert(bytes("offer already redeemed"));
        redeemable.redeemOffer(user, 1000, 0);

        redeemable.redeemOffer(user, 1001, 0);

        vm.stopPrank();
    }

    function testOfferOps() public {
        vm.prank(creator, creator);
        redeemable.grantRole(ElvCmnRoles.MINTER_ROLE, minter);

        vm.startPrank(minter, minter);

        uint256 expectedBitmap = 0;
        uint256 offers;
        uint16 num;

        // Create all offers
        for (uint256 i = 0; i < 256; i++) {
            uint256 res = redeemable.addRedeemableOffer();
            assertEq(res, i, "result != i");

            (offers, num) = redeemable.getOffers();
            assertEq(num, i + 1, "num != i + 1");
            expectedBitmap += 1 << i;
            assertEq(offers, expectedBitmap, "offer != expectedBitmap");
            // log_uint(offers);
        }

        vm.expectRevert(bytes("exceeded max number of offers"));
        redeemable.addRedeemableOffer();

        (offers, num) = redeemable.getOffers();
        //log_uint(offers);

        bool act = redeemable.isOfferActive(100);

        redeemable.removeRedeemableOffer(100);
        act = redeemable.isOfferActive(100);
        assertEq(act, false, "act not false");

        vm.expectRevert(bytes("exceeded max number of offers"));
        redeemable.addRedeemableOffer();

        // Boundary 'removes'
        redeemable.removeRedeemableOffer(255);
        act = redeemable.isOfferActive(255);
        assertEq(act, false, "act not false");

        redeemable.removeRedeemableOffer(0);
        act = redeemable.isOfferActive(0);
        assertEq(act, false, "act not false");

        vm.stopPrank();
    }

    function testRedeemOps() public {
        vm.prank(creator, creator);
        redeemable.grantRole(ElvCmnRoles.MINTER_ROLE, minter);

        vm.startPrank(minter, minter);

        uint256 expectedBitmap = 0;
        uint256 offers;
        uint16 num;

        // Create all offers
        for (uint256 i = 0; i < 256; i++) {
            uint256 res = redeemable.addRedeemableOffer();
            assertEq(res, i, "result != i");

            (offers, num) = redeemable.getOffers();
            assertEq(num, i + 1, "num != i + 1");
            expectedBitmap += 1 << i;
            assertEq(offers, expectedBitmap, "offer != expectedBitmap");
            // log_uint(offers);
        }

        // Redeem all offers
        uint256 tokenId = 131071;

        for (uint8 i = 0; i <= 255; i++) {
            bool res = redeemable.isOfferRedeemed(tokenId, i);
            assertEq(res, false, "isOfferRedeemed should be false");

            res = redeemable.isOfferRedeemed(tokenId, 255);
            assertEq(res, false, "isOfferRedeemed 255 should be false");

            redeemable.redeemOffer(user, tokenId, i);

            res = redeemable.isOfferRedeemed(tokenId, i);
            assertEq(res, true, "isOfferRedeemed should be true");

            if (i == 255) {
                break; // Avoid wrapping back to 0
            }
        }

        vm.stopPrank();
    }

    /*
     * Redeem a large batch for gas testing
     */
    function testRedeemMany() public {
        vm.prank(creator, creator);
        redeemable.grantRole(ElvCmnRoles.MINTER_ROLE, minter);

        vm.startPrank(minter, minter);
        uint256 res = redeemable.addRedeemableOffer();
        assertEq(res, 0, "result != 0");

        uint256 n = 800;

        for (uint256 i = 0; i < n; i++) {
            redeemable.redeemOffer(user, i, 0);
        }

        vm.stopPrank();
    }
}
