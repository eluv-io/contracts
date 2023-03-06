// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {console} from "forge-std/console.sol";
import {Test} from "forge-std/Test.sol";
import "src/tradable/ElvToken.sol";


contract ERC20PaymentsTest is Test {
    ElvToken internal token;
    address public deployer;

    function setUp() public virtual {
        vm.prank(msg.sender, msg.sender);
        deployer = msg.sender;
        token = new ElvToken("ElvToken","ELV",6,20000);
    }

    function testDecimals() public{
        assertEq(token.decimals(), 6);
    }

    function testTokenMinted() public {
        assertEq(token.balanceOf(deployer), 20000 * 10 ** token.decimals());
    }
}