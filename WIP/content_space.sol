pragma solidity ^0.4.21;

import './library.sol';

contract ContentSpace is Ownable {


    string public name;
    string public description;

    event CreateContentType(address contentTypeAddress);
    event CreateLibrary(address libraryAddress);


    function ContentSpace(string memory contentSpaceName) public {
        name = contentSpaceName;
    }


    function setDescription(string memory contentSpaceDescription) public  {
        description = contentSpaceDescription;
    }


    function createContentType() public returns (address) {
        address contentTypeAddress = new Content();
        emit CreateContentType(contentTypeAddress);
        return contentTypeAddress;
    }


    function createLibrary(address address_KMS) public returns (address) {
        address libraryAddress = new ContentLibrary(address_KMS);
        emit CreateLibrary(libraryAddress);
        return libraryAddress;
    }


}