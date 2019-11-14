pragma solidity 0.4.24;

import {BaseContentFactory}  from "./base_content_space.sol";
import {Content} from "./content.sol";
import {BaseContent} from "./base_content.sol";


//BaseCtFactoryXt20191031115100PO: adds support for custom contract
//BaseCtFactoryXt20191031153200ML: passes accessor to the runAccess via the addresses array
//BaseCtFactoryXt20191031170400ML: adds request timestamp to event
//BaseCtFactoryXt20191031203100ML: change initialization of array

contract BaseContentFactoryExt is BaseContentFactory {

    bytes32 public version ="BaseCtFactoryXt20191031203100ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    // TODO: naming this the same as the event in BaseContentObject ...?
    event AccessRequest(
        uint256 timestamp,
        address libraryAddress,
        address contentAddress,
        address userAddress,
        bytes32 contextHash,
        uint64 request_timestamp
    );

    event AccessComplete(
        uint256 timestamp,
        address libraryAddress,
        address contentAddress,
        address userAddress,
        bytes32 contextHash,
        uint64 request_timestamp
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
                emit AccessRequest(now, cobj.libraryAddress(), _contentAddrs[i], _userAddrs[i], _ctxHashes[i], uint64(_ts[i]));
                if (cobj.contentContractAddress() != 0x0) {
                    bytes32[] memory emptyVals;
                    address[] memory paramAddrs =  new address[](1);
                    paramAddrs[0] = _userAddrs[i];
                    c = Content(cobj.contentContractAddress());
                    c.runAccess(_ts[i], 100, emptyVals, paramAddrs); // TODO: level?
                }
            } else if (_opCodes[i] == OP_ACCESS_COMPLETE) {
                emit AccessComplete(now, cobj.libraryAddress(), _contentAddrs[i], _userAddrs[i], _ctxHashes[i], uint64(_ts[i]));
                if (cobj.contentContractAddress() != 0x0) {
                    c = Content(cobj.contentContractAddress());
                    c.runFinalizeExt(_ts[i], _amt[i], _userAddrs[i]);
                }
            } else {
                require(false);
            }
        }
    }
}

