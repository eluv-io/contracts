pragma solidity ^0.5.0;

import "./oz_token/ERC721.sol";
import "./oz_token/ERC721Enumerable.sol";
import "./oz_token/ERC721Metadata.sol";
import "./oz_token/ownership/Ownable.sol";
import "./oz_token/access/roles/MinterRole.sol";
import "./Strings.sol";
import "./elv_wrapped.sol";
import "./elv_token.sol";
import "./elv_token_helper.sol";
import "./elv_proxy.sol";

/**
 * @title ElvTradable
 * ElvTradable - ERC721 contract that whitelists a trading address, and has minting functionality.
 */
contract ElvTradable is ERC721, ERC721Enumerable, ERC721Metadata, ISettableTokenURI, MinterRole, Ownable {
    using Strings for string;

    address public proxyRegistryAddress;

    constructor(
        string memory _name,
        string memory _symbol,
        string memory _contractURI,
        address _proxyRegistryAddress,
        uint256 _baseTransferFee,
        uint256 _cap
    ) public ERC721Metadata(_name, _symbol) {
        proxyRegistryAddress = _proxyRegistryAddress;
        contractURI = _contractURI;
        baseTransferFee = _baseTransferFee;
        cap = _cap;
    }

    string public contractURI;

    function setContractURI(string memory _newContractURI) public onlyOwner {
        contractURI = _newContractURI;
    }

    uint256 public cap;
    uint256 public minted;

    // mapping from token id to ordinal position in which it was minted
    // this is different from _allTokensIndex in ERC721Enumerable because the indexes there can change with burns
    mapping(uint256 => uint256) private _allTokensOrds;

    function ordinalOfToken(uint256 tokenId) public view returns (uint256) {
        require(_exists(tokenId));
        return _allTokensOrds[tokenId];
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
        _allTokensOrds[tokenId] = minted;
        minted = minted.add(1);
        require(cap == 0 || minted <= cap);
        _mint(to, tokenId);
        _setTokenURI(tokenId, tokenURI);
        return true;
    }

    function isMinterSigned(address to, uint256 tokenId, string memory tokenURI, uint8 v, bytes32 r, bytes32 s) public view returns (bool) {
        return isMinter(ecrecover(keccak256(abi.encodePacked(address(this), to, tokenId, tokenURI)), v, r, s));
    }

    // _mint emits event Transfer(address(0), to, tokenId) - so will be indistinguishable
    function mintSignedWithTokenURI(address to, uint256 tokenId, string memory tokenURI, uint8 v, bytes32 r, bytes32 s) public returns (bool) {
        _allTokensOrds[tokenId] = minted;
        minted = minted.add(1);
        require(cap == 0 || minted <= cap);
        require(!_exists(tokenId));
        require(isMinterSigned(to, tokenId, tokenURI, v, r, s));
        _mint(to, tokenId);
        _setTokenURI(tokenId, tokenURI);
        return true;
    }

    function isOwnerSigned(address from, uint256 tokenId, uint8 v, bytes32 r, bytes32 s) public view returns (bool) {
        return _isApprovedOrOwner(ecrecover(keccak256(abi.encodePacked(address(this), from, tokenId)), v, r, s), tokenId);
    }

    function burnSigned(address from, uint256 tokenId, uint8 v, bytes32 r, bytes32 s) public returns (bool) {
        require(isOwnerSigned(from, tokenId, v, r, s));
        require(msg.sender == from && isMinter(from));
        _burn(tokenId);
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

    function exists(uint256 tokenId) public view returns (bool) {
        return _exists(tokenId);
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

    function setTransferFeeProxyAddress(address _newProxy) public onlyOwner {
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

    // override
    function transferFrom(address from, address to, uint256 tokenId) public payable {
        if (msg.value < getTransferFee(tokenId)) {
            require(isProxyApprovedForAll(ownerOf(tokenId), msg.sender), "transfer w/o proxy requires fee");
        }
        super.transferFrom(from, to, tokenId);
    }

    // override
    function safeTransferFrom(address from, address to, uint256 tokenId) public payable {
        if (msg.value < getTransferFee(tokenId)) {
            require(isProxyApprovedForAll(ownerOf(tokenId), msg.sender), "transfer w/o proxy requires fee");
        }
        super.safeTransferFrom(from, to, tokenId);
    }

    // override
    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory _data) public payable {
        if (msg.value < getTransferFee(tokenId)) {
            require(isProxyApprovedForAll(ownerOf(tokenId), msg.sender), "transfer w/o proxy requires fee");
        }
        super.safeTransferFrom(from, to, tokenId, _data);
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

// ElvTradableLocal

contract ElvTradableLocal is ElvTradable {
    using Strings for string;
    using SafeMath for uint256;

    constructor(
        string memory _name,
        string memory _symbol,
        string memory _contractURI,
        address _proxyRegistryAddress,
        uint256 _baseTransferFee,
        uint256 _cap,
        uint256 _defHoldSecs
    ) public ElvTradable(_name, _symbol, _contractURI, _proxyRegistryAddress, _baseTransferFee, _cap) {
        defHoldSecs = _defHoldSecs;
    }

    uint256 public defHoldSecs;

    mapping(uint256 => uint256) public _allTokensHolds;

    function mintHoldWithTokenURI(address to, uint256 tokenId, string memory tokenURI, uint256 holdSecs) public onlyMinter returns (bool) {
        _allTokensHolds[tokenId] = block.timestamp.add(holdSecs);
        return super.mintWithTokenURI(to, tokenId, tokenURI);
    }

    // override
    function mintWithTokenURI(address to, uint256 tokenId, string memory tokenURI) public onlyMinter returns (bool) {
        _allTokensHolds[tokenId] = block.timestamp.add(defHoldSecs);
        return super.mintWithTokenURI(to, tokenId, tokenURI);
    }

    // override
    function mintSignedWithTokenURI(address to, uint256 tokenId, string memory tokenURI, uint8 v, bytes32 r, bytes32 s) public returns (bool) {
        _allTokensHolds[tokenId] = block.timestamp.add(defHoldSecs);
        return super.mintSignedWithTokenURI(to, tokenId, tokenURI, v, r, s);
    }

    // override
    function transferFrom(address from, address to, uint256 tokenId) public payable {
        require(block.timestamp >= _allTokensHolds[tokenId]);
        super.transferFrom(from, to, tokenId);
    }

    // override
    function safeTransferFrom(address from, address to, uint256 tokenId) public payable {
        require(block.timestamp >= _allTokensHolds[tokenId]);
        super.safeTransferFrom(from, to, tokenId);
    }

    // override
    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory _data) public payable {
        require(block.timestamp >= _allTokensHolds[tokenId]);
        super.safeTransferFrom(from, to, tokenId, _data);
    }
}