pragma solidity ^0.5.0;

import "./oz_token/ERC20Capped.sol";
import "./oz_token/ERC20Pausable.sol";
import "./oz_token/ERC20Detailed.sol";

contract ElvToken is ERC20Capped, ERC20Detailed, ERC20Pausable {
    constructor (uint256 cap, string memory name, string memory symbol, uint8 decimals)
    ERC20Capped(cap)
    ERC20Detailed(name, symbol, decimals)
    public {}
}
