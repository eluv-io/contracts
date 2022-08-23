// SPDX-License-Identifier: MIT
pragma solidity ^0.6.12;

import "forge-std/Script.sol";
import "forge-std/console2.sol";
import "src/redeemable.sol";

contract RedeemableScript is Script {

    function run() external {
        vm.startBroadcast();
        Redeemable redeemable = new Redeemable();
        vm.stopBroadcast();
    }
}
