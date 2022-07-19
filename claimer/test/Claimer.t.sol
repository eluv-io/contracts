// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../src/Claimer.sol";

error ErrOwner();
error ErrAllocated();
error ErrClaim();
error ErrBurn();
error ErrAuth();

contract ClaimerTest is Test {
    Claimer claimer;
    Claimer claimerTest;

    function setUp() public {
        claimer = new Claimer();
    }

    function testOwner() public {
        vm.prank(msg.sender);
        claimerTest = new Claimer();
        if(msg.sender != claimerTest.owner()){
            revert ErrOwner(); 
        }
    }

    function testMultipleAllocations() public {

        address currAdr1 = 0x7109709ECfa91a80626fF3989D68f67F5b1DD12D;
        address currAdr2 = 0x7109709ecFA91A80626Ff3989D68f67f5b1Dd12E;
        uint expirDate = 1000;
        claimer.addAuthorizedAdr(currAdr1);
        claimer.addAuthorizedAdr(currAdr2);

        claimer.allocate(currAdr1, 5, expirDate);
        claimer.allocate(currAdr1, 10, expirDate);
        if(claimer.getAmount(currAdr1, 0) != 5 || claimer.getAmount(currAdr1, 1) != 10) revert ErrAllocated();
        vm.prank(currAdr1); 
        claimer.burn(10);
        if((claimer.getAmount(currAdr1, 0) + claimer.getAmount(currAdr1, 1)) != 5 || claimer.burns(currAdr1) != 10 || claimer.claims(currAdr1) != 0) revert ErrBurn();
        vm.prank(currAdr1); 
        claimer.burn(5);
        if((claimer.getAmount(currAdr1, 0) + claimer.getAmount(currAdr1, 1)) != 0 || claimer.burns(currAdr1) != 15 || claimer.claims(currAdr1) != 0) revert ErrBurn();
        
        claimer.allocate(currAdr2, 15, expirDate);

        vm.prank(currAdr2); 
        claimer.claim(10); vm.prank(currAdr2); 
        claimer.burn(8); vm.prank(currAdr2);
        claimer.claim(8);
        if(claimer.getAmount(currAdr2, 0) != 0 || claimer.burns(currAdr2) != 5 || claimer.claims(currAdr2) != 10) revert ErrClaim();


    }

    function testAllocation() public {
        address currAdrr = 0x7109709ECfa91a80626fF3989D68f67F5b1DD12D;
        claimer.addAuthorizedAdr(currAdrr);
        if(!claimer.allocationAuthorities(currAdrr)){
            revert ErrAuth();
        }

        claimer.allocate(currAdrr, 10, 10);

        if(claimer.getAmount(currAdrr, 0) != 10){
            revert ErrAllocated();
        }

        vm.prank(currAdrr);
        claimer.claim(10);
        if(claimer.getAmount(currAdrr, 0) != 0 || claimer.claims(currAdrr) != 10){
            revert ErrClaim();
        }
    }


    function testBurn() public {
        address currAdrr = 0x7109709ecFA91A80626Ff3989D68f67f5b1Dd12E;
        claimer.addAuthorizedAdr(currAdrr);
        claimer.allocate(currAdrr, 20, 20);
        vm.prank(currAdrr); 
        claimer.burn(10);
        if(claimer.claims(currAdrr) != 0 || claimer.burns(currAdrr) != 10){
            revert ErrBurn();
        }

    }

    function testAddRmAuthorized() public {
        address currAdrr = 0x7109709eCfa91A80626fF3989d68f67F5b1dd12F;
        claimer.addAuthorizedAdr(currAdrr);
        if(!claimer.allocationAuthorities(currAdrr)){
            revert ErrAuth();
        }

        claimer.rmAuthorizedAdr(currAdrr);
        if(claimer.allocationAuthorities(currAdrr)){
            revert ErrAuth();
        }
    }
}
