pragma solidity ^0.4.24;

import "./base_content_space.sol";

contract BaseContentFactoryExt is BaseContentFactory {

    // TODO: naming this the same as the event in BaseContentObject ...?
    event AccessRequest(
        address libraryAddress,
        address contentAddress,
        address userAddress,
        bytes32 contextHash,
        uint64 timestamp
    );

    function executeAccessBatch(address[] _contentAddrs, address[] _userAddrs, bytes32[] _ctxHashes, uint256[] _ts) public {

//        BaseContentSpace ourSpace = BaseContentSpace(contentSpace);
//        require(msg.sender == owner || ourSpace.checkKMSAddr(msg.sender) > 0);

        uint paramsLen = _contentAddrs.length;

        require(_userAddrs.length == paramsLen);
        require(_ctxHashes.length == paramsLen);
        require(_ts.length == paramsLen);

        for (uint i = 0; i < paramsLen; i++) {
            BaseContent cobj = BaseContent(_contentAddrs[i]);
//            require(msg.sender == owner || cobj.addressKMS() == msg.sender);
            emit AccessRequest(cobj.libraryAddress(), _contentAddrs[i], _userAddrs[i], _ctxHashes[i], uint64(_ts[i]));
        }
    }
}
