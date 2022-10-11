// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/access/AccessControlEnumerable.sol";
import "./cmn/Roles.sol";

contract ElvToken is ERC20, ERC20Burnable, Pausable, AccessControlEnumerable {
    constructor(string memory name, string memory symbol, uint256 amount) ERC20(name, symbol) {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ElvCmnRoles.PAUSER_ROLE, msg.sender);
        _mint(msg.sender, amount * 10 ** decimals());
        _grantRole(ElvCmnRoles.MINTER_ROLE, msg.sender);
    }

    function pause() public onlyRole(ElvCmnRoles.PAUSER_ROLE) {
        _pause();
    }

    function unpause() public onlyRole(ElvCmnRoles.PAUSER_ROLE) {
        _unpause();
    }

    function mint(address to, uint256 amount) public onlyRole(ElvCmnRoles.MINTER_ROLE) {
        _mint(to, amount);
    }

    //    function _mint(address account, uint256 amount) internal override(ERC20,ERC20Capped) {
    //        super._mint(account, amount);
    //    }

    function _beforeTokenTransfer(address from, address to, uint256 amount) internal override whenNotPaused {
        super._beforeTokenTransfer(from, to, amount);
    }
}
