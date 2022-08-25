// SPDX-License-Identifier: MIT
pragma solidity ^0.6.12;

import "forge-std/Script.sol";
import "forge-std/console2.sol";
import "src/redeemable.sol";

/**
    To run this script on the testnet:
    export PRIVATE_KEY=your_private_key
    export RPC_URL=https://host-766.contentfabric.io/eth 
    forge script script/tradable/redeemable.s.sol:RedeemableScript --fork-url ${RPC_URL} --private-key ${PRIVATE_KEY} --broadcast --legacy
 */
contract RedeemableScript is Script {

    function run() external {
        vm.startBroadcast();
        Redeemable redeemable = new Redeemable();

        //test add redeeemable offer
        for(uint i=0; i<5; i ++){
            redeemable.addRedeemableOffer();
        }

        //test remove redeemable offer
        redeemable.removeRedeemableOffer(0);
        redeemable.removeRedeemableOffer(2);
        redeemable.removeRedeemableOffer(4);

        //test the current state
        console.log("Expected : false, Result : ", vm.toString(redeemable.isOfferActive(0)));
        console.log("Expected : true, Result : ", vm.toString(redeemable.isOfferActive(1)));
        console.log("Expected : false, Result : ", vm.toString(redeemable.isOfferActive(2)));
        console.log("Expected : true, Result : ", vm.toString(redeemable.isOfferActive(3)));
        console.log("Expected : false, Result : ", vm.toString(redeemable.isOfferActive(4)));

        
        //test the offers for a specific token
        uint tokenId = 1000;
        redeemable.redeemOffer(tokenId, 1);
        
        console.log("Expected : false, Result : ", vm.toString(redeemable.isOfferRedeemed(tokenId, 0)));
        console.log("Expected : true, Result : ", vm.toString(redeemable.isOfferRedeemed(tokenId, 1)));
        console.log("Expected : false, Result : ", vm.toString(redeemable.isOfferRedeemed(tokenId, 2)));
        console.log("Expected : false, Result : ", vm.toString(redeemable.isOfferRedeemed(tokenId, 3)));
        console.log("Expected : false, Result : ", vm.toString(redeemable.isOfferRedeemed(tokenId, 4)));

        vm.stopBroadcast();
    }
}


//maybe we can add a getter for redemptions with a parameter tokenId
