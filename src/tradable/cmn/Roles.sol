// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

library ElvCmnRoles {
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
}
