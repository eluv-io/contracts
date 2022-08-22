// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import "forge-std/Script.sol";
import "forge-std/console2.sol";
import "tradable/redeemable.sol";

contract RedeemableScript is Script {

    function run() external {
        vm.startBroadcast();
        Redeemable redeemabpe = new Redeemable();

//        console.log("Creator of the smart contract : ", vm.toString(redeemable.creator()));
//        console.log("Owner of the smart contract : ", vm.toString(redeemable.owner()));

        vm.stopBroadcast();


    }
}
