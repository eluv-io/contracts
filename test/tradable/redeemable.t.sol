// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import "forge-std/Test.sol";
import "forge-std/console2.sol";
import "../../tradable/redeemable.sol";


contract RedeemableTest is Test {
    Redeemable redeemable;
    address adr = 0x7109709ECfa91a80626fF3989D68f67F5b1DD12D;

    /**
        Set up function, it is the first function executed when we run 'forge test'
    */
    function setUp() public {
        vm.prank(adr, adr); //adr is the creator and the owner of this smart contract
        redeemable = new Redeemable();
    }

}
