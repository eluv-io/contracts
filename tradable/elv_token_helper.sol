pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./oz_token/ownership/Ownable.sol";
import "./elv_tradable_full.sol";

/**
 * @title Minting helper
 * @dev Contains functions that package multiple token function calls into a
 * single transaction. This contract must be added as a minter for the
 * associated token contracts.
 */
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

    uint256 public overrideHoldSecs;

    function setOverrideHoldSecs(uint256 _overrideHoldSecs) public onlyOwner {
        overrideHoldSecs = _overrideHoldSecs;
    }

    /**
     * @dev Burns the specified token for the owner. Useful when the owner does
     * not have gas, e.g. custodial wallet.
     * The message signed by the token owner is a concatenation (packing) of the following:
     *   token contract address | mint helper address | token ID
     *
     * @param token contract of the token to burn
     * @param tokenId ID of the token to burn
     * @param v part of the owner's signature giving permission to burn the token
     * @param r part of the owner's signature
     * @param s part of the owner's signature
     */
    function burnSigned(address token, uint256 tokenId, uint8 v, bytes32 r, bytes32 s) public onlyOwner returns (bool) {
        ElvTradable t = ElvTradable(token);
        return t.burnSigned(address(this), tokenId, v, r, s);
    }

    function burnSignedEIP191(address token, uint256 tokenId, uint8 v, bytes32 r, bytes32 s) public onlyOwner returns (bool) {
        ElvTradable t = ElvTradable(token);
        return t.burnSignedEIP191(address(this), tokenId, v, r, s);
    }

    // Deprecated - superceded by burnSignedAndMint
    // for this to work the 'from' address needs to be *this* contract
    function burnSignedAndMintMany(address burnAddr, address from, uint256 burnTokenId, uint8 v, bytes32 r, bytes32 s,
        address[] memory tokAddrs, address[] memory to, uint256[] memory tokenIds, string[] memory tokenURIs) public onlyOwner returns (bool) {
        ElvTradable etBurn = ElvTradable(burnAddr);
        etBurn.burnSigned(from, burnTokenId, v, r, s);
        if (overrideHoldSecs != 0) {
            for (uint i = 0; i < tokAddrs.length; i++) {
                ElvTradableLocal et = ElvTradableLocal(tokAddrs[i]);
                if (!et.exists(tokenIds[i])) {
                    et.mintHoldWithTokenURI(to[i], tokenIds[i], tokenURIs[i], overrideHoldSecs);
                }
            }
            return true;
        } else {
            return mintWithTokenURIMany(tokAddrs, to, tokenIds, tokenURIs);
        }
    }

    /**
     * @dev Redeems (mints) the specified tokenIds in exchange for (burning)
     * the burnTokenIds.
     * @param burnTokens contracts of the tokens to burn
     * @param burnTokenIds IDs of the tokens to burn
     * @param v part of the owner's signature giving permission to burn the token
     * @param r part of the owner's signature
     * @param s part of the owner's signature
     * @param to beneficiaries that will own the minted tokens
     * @param mintTokens contracts of the tokens to mint
     * @param mintTokenIds IDs to set when minting
     * @param mintTokenURIs URIs to set when minting
     */
    function burnSignedAndMint(
        address[] memory burnTokens,
        uint256[] memory burnTokenIds,
        uint8[] memory v,
        bytes32[] memory r,
        bytes32[] memory s,
        address[] memory to,
        address[] memory mintTokens,
        uint256[] memory mintTokenIds,
        string[] memory mintTokenURIs
    ) public onlyOwner returns (bool) {
        for (uint i = 0; i < burnTokens.length; i++) {
            ElvTradable etBurn = ElvTradable(burnTokens[i]);
            etBurn.burnSigned(address(this), burnTokenIds[i], v[i], r[i], s[i]);
        }

        if (overrideHoldSecs != 0) {
            for (uint i = 0; i < mintTokens.length; i++) {
                ElvTradableLocal et = ElvTradableLocal(mintTokens[i]);
                if (!et.exists(mintTokenIds[i])) {
                    et.mintHoldWithTokenURI(to[i], mintTokenIds[i], mintTokenURIs[i], overrideHoldSecs);
                }
            }
            return true;
        } else {
            return mintWithTokenURIMany(mintTokens, to, mintTokenIds, mintTokenURIs);
        }
    }

    function burnSignedEIP191AndMint(
        address[] memory burnTokens,
        uint256[] memory burnTokenIds,
        uint8[] memory v,
        bytes32[] memory r,
        bytes32[] memory s,
        address[] memory to,
        address[] memory mintTokens,
        uint256[] memory mintTokenIds,
        string[] memory mintTokenURIs
    ) public onlyOwner returns (bool) {
        for (uint i = 0; i < burnTokens.length; i++) {
            ElvTradable etBurn = ElvTradable(burnTokens[i]);
            etBurn.burnSignedEIP191(address(this), burnTokenIds[i], v[i], r[i], s[i]);
        }

        if (overrideHoldSecs != 0) {
            for (uint i = 0; i < mintTokens.length; i++) {
                ElvTradableLocal et = ElvTradableLocal(mintTokens[i]);
                if (!et.exists(mintTokenIds[i])) {
                    et.mintHoldWithTokenURI(to[i], mintTokenIds[i], mintTokenURIs[i], overrideHoldSecs);
                }
            }
            return true;
        } else {
            return mintWithTokenURIMany(mintTokens, to, mintTokenIds, mintTokenURIs);
        }
    }

    function redeemOfferSigned(address token, uint256 tokenId, uint8 offerId, uint8 v, bytes32 r, bytes32 s) public onlyOwner {
        ElvTradable t = ElvTradable(token);
        t.redeemOfferSigned(address(this), tokenId, offerId, v, r, s);
    }

    function redeemOfferSignedEIP191(address token, uint256 tokenId, uint8 offerId, uint8 v, bytes32 r, bytes32 s) public onlyOwner{
        ElvTradable t = ElvTradable(token);
        t.redeemOfferSignedEIP191(address(this), tokenId, offerId, v, r, s);
    }

}
