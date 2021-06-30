pragma solidity ^0.5.0;

import "./oz_token/ERC721.sol";
import "./oz_token/ERC721Enumerable.sol";
import "./oz_token/ERC721Metadata.sol";
import "./oz_token/ownership/Ownable.sol";
import "./oz_token/access/roles/MinterRole.sol";
import "./Strings.sol";
import "./elv_wrapped.sol";

contract OwnableDelegateProxy {}

contract ProxyRegistry {
    mapping(address => OwnableDelegateProxy) public proxies;
}

contract OwnerProxyRegistry is ProxyRegistry, Ownable {

    int public countDelegates;

    constructor (address[10] memory initDelegates) public {
        for (uint i = 0; i < initDelegates.length; i++) {
            if (initDelegates[i] != address(0)) {
                addDelegate(initDelegates[i]);
            }
        }
    }

    function addDelegate(address from) public onlyOwner {
        require(from != msg.sender);
        proxies[from] = OwnableDelegateProxy(msg.sender);
        countDelegates++;
    }

    function finalize() public onlyOwner {
        selfdestruct(msg.sender);
    }
}

interface TransferFeeProxy {
    function getTransferFee(uint256 _tokenId) external view returns (uint256);
}

/**
 * @title ElvTradable
 * ElvTradable - ERC721 contract that whitelists a trading address, and has minting functionality.
 */
contract ElvTradable is ERC721, ERC721Enumerable, ERC721Metadata, MinterRole, Ownable {
    using Strings for string;

    address public proxyRegistryAddress;

    constructor(
        string memory _name,
        string memory _symbol,
        string memory _contractURI,
        address _proxyRegistryAddress,
        uint256 _baseTransferFee
    ) public ERC721Metadata(_name, _symbol) {
        proxyRegistryAddress = _proxyRegistryAddress;
        contractURI = _contractURI;
        baseTransferFee = _baseTransferFee;
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

    int public constant PROXY_TYPE_REGISTRY = 1;
    int public constant PROXY_TYPE_TRANSFER_FEE = 2;

    event SetProxyAddress(int proxyType, address indexed prevAddr, address indexed newAddr);

    function setProxyRegistryAddress(address _newProxy) public onlyOwner {
        emit SetProxyAddress(PROXY_TYPE_REGISTRY, proxyRegistryAddress, _newProxy);
        proxyRegistryAddress = _newProxy;
    }

    uint256 public baseTransferFee;
    address public transferFeeProxyAddress;

    function setTransferFreeProxyAddress(address _newProxy) public onlyOwner {
        emit SetProxyAddress(PROXY_TYPE_TRANSFER_FEE, transferFeeProxyAddress, _newProxy);
        transferFeeProxyAddress = _newProxy;
    }

    event BaseTransferFeeSet(uint256 prevFee, uint256 newFee);

    function setBaseTransferFee(uint256 _newBaseFee) public onlyOwner {
        emit BaseTransferFeeSet(baseTransferFee, _newBaseFee);
        baseTransferFee = _newBaseFee;
    }

    // naming is intended to be in line with proposal: https://github.com/ethereum/EIPs/issues/2665
    function getTransferFee(uint256 _tokenId) public view returns (uint256) {
        if (transferFeeProxyAddress == address(0)) {
            return baseTransferFee;
        }
        return TransferFeeProxy(transferFeeProxyAddress).getTransferFee(_tokenId);
    }

    function withdraw(uint256 _amount) public onlyOwner {
        msg.sender.transfer(_amount);
    }

    function transferFrom(address from, address to, uint256 tokenId) public payable {
        if (msg.value < getTransferFee(tokenId)) {
            require(isProxyApprovedForAll(ownerOf(tokenId), msg.sender), "transfer w/o proxy requires fee");
        }
        super.transferFrom(from, to, tokenId);
    }

    function isProxyApprovedForAll(address owner, address operator) public view returns (bool) {
        if (proxyRegistryAddress != address(0)) {
            ProxyRegistry proxyRegistry = ProxyRegistry(proxyRegistryAddress);
            if (address(proxyRegistry.proxies(owner)) == operator) {
                return true;
            }
        }
        return false;
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
        if (isProxyApprovedForAll(owner, operator))
            return true;
        return super.isApprovedForAll(owner, operator);
    }
}