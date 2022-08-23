pragma solidity ^0.6.12;

/**
 * MOCK IMPLEMENTATION
 * Stand in for oz_tradable MinterRole.sol
 */
contract MinterRole {

    modifier onlyMinter() {
        _;
    }
}
