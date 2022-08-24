// SPDX-License-Identifier: MIT
pragma solidity >=0.5.4 <0.9.0;

import "./oz_token/access/roles/MinterRole.sol";

/**
 * @title Redeemable
 * @dev Base contract that allows children to implement redeemable offers.
 *
 * Manage redeemable offers, typically associated with an NFT contract.
 * Redeemable offers can be added and removed by the MinterRole.
 * There can be 256 total offers per "Redeemable" instance (stored in a 256 bit bitmap).
 * Offers can get removed but they will be still keeping their "slot" so once
 * 256 offers are allocated, no more offers can be created regardless of how many
 * have already been removed.
 * Redemption state is stored in one 256 bit bitmap per tokenId.
 */
contract Redeemable is MinterRole {

    event RedeemableAdded(uint8 offerId);
    event RedeemableRemoved(uint8 offerId);
    event Redeem(uint256 tokenId, uint offerId);

    // Offers
    uint256 private offers;  // Bitmap of current offers - "1" means active
    uint16 private numOffers; // Number of offers created so far (16 bit to store '256')

    // Redemption state - store an offers bitmap for each tokenId
    mapping(uint256 => uint256) private redemptions;

    constructor() public {
        offers = 0;
        numOffers = 0;
    }

    /**
     * @dev Add an offer
     */
    function addRedeemableOffer() public onlyMinter returns (uint8) {
        require(numOffers < 256, "exceeded max number of offers");

        uint8 offerId = uint8(numOffers);
        uint256 mask = 1 << uint256(offerId);
        offers |= mask;
        numOffers ++;
        emit RedeemableAdded(offerId);
        return offerId;
    }

    /**
     * @dev Remove an active offer (simply deactivates it and retains it's 'slot')
     */
    function removeRedeemableOffer(uint8 offerId) public onlyMinter {
        require(offerId < numOffers, "bad offer id");

        uint256 mask = 1 << uint256(offerId);
        require(offers & mask > 0, "offer not active");

        offers = offers ^ mask;
        emit RedeemableRemoved(offerId);
    }

    /**
     * @dev Redeem an offer, as a token owner.
     * Calling contract must override and require caller is owner of tokenId.
     */
    function redeemOffer(uint256 tokenId, uint8 offerId) public payable {

        uint256 mask = 1 << uint256(offerId);
        require(offers & mask > 0, "offer not active");
        require(redemptions[tokenId] & mask == 0, "offer already redeemed");

        redemptions[tokenId] |= mask;
        emit Redeem(tokenId, offerId);
    }

    /**
     * @dev Check if an offer is already redeemed for a given token ID
     */
    function isOfferRedeemed(uint256 tokenId, uint8 offerId) public view returns (bool) {

        uint256 mask = 1 << uint256(offerId);
        if (offers & mask == 0) {
            return false; // This offer is not active
        }
        if (redemptions[tokenId] & mask > 0) {
            return true;  // This offer is redeemed for tokenId
        }
        return false;
    }

    /**
     * @dev Check if an offer is active
     */
    function isOfferActive(uint8 offerId) public view returns (bool) {
        if (offerId >= numOffers) {
            return false; // This offer has not been created
        }
        uint256 mask = 1 << uint256(offerId);
        return offers & mask > 0;
    }

    /**
     * @dev Return the offers bitmap and number of offers created so far
     */
    function getOffers() public view returns (uint256, uint16) {
        return (offers, numOffers);
    }

}
