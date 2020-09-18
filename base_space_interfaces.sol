//SPDX-License-Identifier: MIT
pragma solidity ^0.7.1;

interface IFactorySpace {
    // this is the only method that's not directly called on the space
    function createContent(address payable lib, address payable content_type) external returns (address payable);

    // current factory methods of space - not including wallet ...
    function createContentType() external returns (address payable);
    function createLibrary(address payable address_KMS) external returns (address payable);
    function createGroup() external returns (address payable);
}

interface IUserSpace {
    function userWallets(address payable _userAddr) external view returns (address payable);
    function createUserWallet(address payable _user) external returns (address payable);
}

interface IKmsSpace {
    function checkKMSAddr(address payable _kmsAddr) external view returns (uint);
    function getKMSID(address payable _kmsAddr) external view returns (string memory);
    function getKMSInfo(string memory _kmsID, bytes memory prefix) external view returns (string memory, string memory);
}

interface INodeSpace {
    function canNodePublish(address payable candidate) external view returns (bool);
}

