pragma solidity ^0.8.13;

/**
 * MOCK IMPLEMENTATION
 * Stand in for oz_tradable MinterRole.sol
 */
contract MinterRole {

    modifier onlyMinter() {
        _;
    }
}
