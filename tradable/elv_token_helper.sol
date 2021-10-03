pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./oz_token/ownership/Ownable.sol";
import "./elv_tradable_full.sol";

contract ElvTokenHelper is Ownable {
    function mintWithTokenURIMany(address[] memory tokAddrs, address[] memory to, uint256[] memory tokenIds, string[] memory tokenURIs)
    public onlyOwner returns (bool) {
        for (uint i = 0; i < tokAddrs.length; i++) {
            ElvTradable et = ElvTradable(tokAddrs[i]);
            if (!et.exists(tokenIds[i])) {
                et.mintWithTokenURI(to[i], tokenIds[i], tokenURIs[i]);
            }
        }
        return true;
    }

    // for this to work the 'from' address needs to be *this* contract
    function burnSignedAndMintMany(address burnAddr, address from, uint256 burnTokenId, uint8 v, bytes32 r, bytes32 s,
        address[] memory tokAddrs, address[] memory to, uint256[] memory tokenIds, string[] memory tokenURIs) public onlyOwner returns (bool) {
        ElvTradable etBurn = ElvTradable(burnAddr);
        etBurn.burnSigned(from, burnTokenId, v, r, s);
        return mintWithTokenURIMany(tokAddrs, to, tokenIds, tokenURIs);
    }
}
