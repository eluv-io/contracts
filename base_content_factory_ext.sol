pragma solidity ^0.4.24;

import "./base_content_space.sol";
import "./content.sol";

// abigen --sol ./base_content_factory_ext.sol --pkg=contracts_ext --out build/ext/base_content_factory_ext.go

contract BaseContentFactoryExt is BaseContentFactory {

    // TODO: naming this the same as the event in BaseContentObject ...?
    event AccessRequest(
        address libraryAddress,
        address contentAddress,
        address userAddress,
        bytes32 contextHash,
        uint64 timestamp
    );

    event AccessComplete(
        address libraryAddress,
        address contentAddress,
        address userAddress,
        bytes32 contextHash,
        uint64 timestamp
    );

    uint32 public constant OP_ACCESS_REQUEST = 1;
    uint32 public constant OP_ACCESS_COMPLETE = 2;

    function executeAccessBatch(uint32[] _opCodes, address[] _contentAddrs, address[] _userAddrs, bytes32[] _ctxHashes, uint256[] _ts, uint256[] _amt) public {

//        BaseContentSpace ourSpace = BaseContentSpace(contentSpace);
//        require(msg.sender == owner || ourSpace.checkKMSAddr(msg.sender) > 0);

        uint paramsLen = _opCodes.length;

        require(_contentAddrs.length == paramsLen);
        require(_userAddrs.length == paramsLen);
        require(_ctxHashes.length == paramsLen);
        require(_ts.length == paramsLen);

        for (uint i = 0; i < paramsLen; i++) {
            BaseContent cobj = BaseContent(_contentAddrs[i]);
            Content c;
            // require(msg.sender == owner || cobj.addressKMS() == msg.sender);
            if (_opCodes[i] == OP_ACCESS_REQUEST) {
                emit AccessRequest(cobj.libraryAddress(), _contentAddrs[i], _userAddrs[i], _ctxHashes[i], uint64(_ts[i]));
                if (cobj.contentContractAddress() != 0x0) {
                    bytes32[] memory emptyVals;
                    address[] memory emptyAddrs;
                    c = Content(cobj.contentContractAddress());
                    c.runAccess(_ts[i], 100, emptyVals, emptyAddrs); // TODO: level?
                }
            } else if (_opCodes[i] == OP_ACCESS_REQUEST) {
                emit AccessComplete(cobj.libraryAddress(), _contentAddrs[i], _userAddrs[i], _ctxHashes[i], uint64(_ts[i]));
                if (cobj.contentContractAddress() != 0x0) {
                    c = Content(cobj.contentContractAddress());
                    c.runFinalize(_ts[i], _amt[i]);
                }
            } else {
                require(false);
            }
        }
    }
}
