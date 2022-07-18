pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./oz_token/ERC721.sol";
import "./oz_token/ownership/Ownable.sol";

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

interface ISettableTokenURI {
    function setTokenURI(uint256 tokenId, string calldata uri) external;
}

contract TransferProxyRegistry is ProxyRegistry, Ownable {

    int public countDelegates;

    // assumes that this contract is currently the transfer proxy of the target
    function proxyTransferFrom(address target, address from, address to, uint256 tokenId) public onlyOwner payable {
        if (proxies[from] == OwnableDelegateProxy(0)) {
            proxies[from] = OwnableDelegateProxy(address(this));
            countDelegates++;
        }
        IERC721(target).transferFrom(from, to, tokenId);
    }

    // assumes that this contract is currently the transfer proxy of the target
    function proxySetTokenURI(address target, uint256 tokenId, string memory uri) public onlyOwner payable {
        address from = IERC721(target).ownerOf(tokenId);
        if (proxies[from] == OwnableDelegateProxy(0)) {
            proxies[from] = OwnableDelegateProxy(address(this));
            countDelegates++;
        }
        ISettableTokenURI(target).setTokenURI(tokenId, uri);
    }

    function proxySetTokenURIMany(address target, uint256[] memory tokenIds, string[] memory uris) public onlyOwner payable {
        for (uint i = 0; i < tokenIds.length; i++) {
            proxySetTokenURI(target, tokenIds[i], uris[i]);
        }
    }

    function finalize() public onlyOwner {
        selfdestruct(msg.sender);
    }
}

interface TransferFeeProxy {
    function getTransferFee(uint256 _tokenId) external view returns (uint256);
}
