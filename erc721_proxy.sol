pragma solidity 0.5.4;

import "./lib_external_call.sol";

contract ERC721Proxy {

    bytes4 constant sigBalanceOf = bytes4(keccak256("balanceOf(address)"));

    bytes32 public version = "TODO";

    address public externalAddress;
    string public configName;

    constructor(address payable _extAddr, string memory _configName)  public payable {
        externalAddress = _extAddr;
        configName = _configName;
    }

    // erc721 standard
    function balanceOf(address owner) public view returns (uint256 balance) {
        // block.number needs to be an ETH (external!) block num?!?!?!?!
        return ExternalCall.callUint(configName, externalAddress, block.number,
            abi.encodeWithSelector(sigBalanceOf, owner));
    }

}
