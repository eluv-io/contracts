pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./oz_token/ownership/Ownable.sol";
import "./elv_tradable_full.sol";

contract ElvTokenHelper is Ownable {
    function mintWithTokenURIMany(address[] memory tok, address[] memory to, uint256[] memory tokenId, string[] memory tokenURI)
    public onlyOwner returns (bool) {
        for (uint i = 0; i < tok.length; i++) {
            ElvTradable(tok[i]).mintWithTokenURI(to[i], tokenId[i], tokenURI[i]);
        }
        return true;
    }
}
