pragma solidity ^0.5.0;

import "./oz_erc721/ERC721.sol";
import "./oz_erc721/ERC721Enumerable.sol";
import "./oz_erc721/ERC721Metadata.sol";
import "./oz_erc721/ownership/Ownable.sol";
import "./oz_erc721/access/roles/MinterRole.sol";
import "./Strings.sol";

contract OwnableDelegateProxy {}

contract ProxyRegistry {
    mapping(address => OwnableDelegateProxy) public proxies;
}

/**
 * @title ElvTradable
 * ElvTradable - ERC721 contract that whitelists a trading address, and has minting functionality.
 */
contract ElvTradable is ERC721, ERC721Enumerable, ERC721Metadata, MinterRole, Ownable {
    using Strings for string;

    address proxyRegistryAddress;

    constructor(
        string memory _name,
        string memory _symbol,
        string memory _contractURI,
        address _proxyRegistryAddress
    ) public ERC721Metadata(_name, _symbol) {
        proxyRegistryAddress = _proxyRegistryAddress;
        contractURI = _contractURI;
    }

    string public contractURI;

    function setContractURI(string memory _newContractURI) public onlyOwner {
        contractURI = _newContractURI;
    }

    // straight from ERC721MetadataMintable in OpenZeppelin
    /**
     * @dev Function to mint tokens
     * @param to The address that will receive the minted tokens.
     * @param tokenId The token id to mint.
     * @param tokenURI The token URI of the minted token.
     * @return A boolean that indicates if the operation was successful.
    */
    function mintWithTokenURI(address to, uint256 tokenId, string memory tokenURI) public onlyMinter returns (bool) {
        _mint(to, tokenId);
        _setTokenURI(tokenId, tokenURI);
        return true;
    }

    /**
     * @dev Burns a specific ERC721 token.
     * @param tokenId uint256 id of the ERC721 token to be burned.
    */
    function burn(uint256 tokenId) public {
        require(_isApprovedOrOwner(msg.sender, tokenId));
        _burn(tokenId);
    }

    event SetTokenURI(uint256 indexed tokenId, string prevURI, string newURI);

    // allows the owner of a token to reset the URI.
    function setTokenURI(uint256 tokenId, string memory uri) public {
        require(_isApprovedOrOwner(msg.sender, tokenId));
        emit SetTokenURI(tokenId, this.tokenURI(tokenId), uri);
        _setTokenURI(tokenId, uri);
    }

    event SetProxyAddress(address indexed prevAddr, address indexed newAddr);

    function setProxyRegistryAddress(address _newProxy) public onlyOwner {
        emit SetProxyAddress(proxyRegistryAddress, _newProxy);
        proxyRegistryAddress = _newProxy;
    }

    /**
     * Allow for proxy contract to manage approvals.
     * Can be used to whitelist user's OpenSea proxy accounts (and others) to enable gas-less listings.
     */
    function isApprovedForAll(address owner, address operator)
    public
    view
    returns (bool)
    {
        if (proxyRegistryAddress != address(0)) {
            ProxyRegistry proxyRegistry = ProxyRegistry(proxyRegistryAddress);
            if (address(proxyRegistry.proxies(owner)) == operator) {
                return true;
            }
        }

        return super.isApprovedForAll(owner, operator);
    }
}