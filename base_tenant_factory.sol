pragma solidity 0.5.4;

import {Ownable} from "./ownable.sol";
import {Utils} from "./utils.sol";
import {AccessIndexor} from "./access_indexor.sol";
import "./base_tenant_space.sol";
import "./base_access_control_group.sol";
import "./strings_utils.sol";
import "./address_utils.sol";

/* -- Revision history --
BaseTenantFactory20230404000000PM: First versioned released
*/

contract BaseTenantFactory is Ownable {
    bytes32 public version ="BsTenantFactory20230404120000PM"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    event CreateTenant(address tenantAddress);

    constructor(address payable _spaceAddr) public {
        require((msg.sender == _spaceAddr || factoryHelper.isValidTenantCreator(_spaceAddr)), "tenant_factory_creator(msg.sender) invalid");
        if (msg.sender == _spaceAddr) {
            require(factoryHelper.isBaseContentSpaceContract(_spaceAddr), "msg.sender != content_space_address");
        }
        contentSpace = _spaceAddr;
    }


    /*
     * 1. Checks if the msg.sender is content_space or
     *    is part of the group address provided in space
     *    metaObject '_ELV_GROUP_TENANT_AUTHORITIES'
     * 2. create baseTenantSpace contract ( has indexCaterory set to CATEGORY_CONTRACT())
     * 3. set rights to tx.origin for baseTenantSpace contract
    */
    function createTenant(string memory _tenantName, address _kmsAddr) public returns (address) {
        require(msg.sender == contentSpace || factoryHelper.isValidTenantCreator(contentSpace), "tenant_creator(msg.sender) invalid");
        address payable newTenant = address(new BaseTenantSpace(contentSpace, _tenantName, _kmsAddr));
        BaseTenantSpace theTenant = BaseTenantSpace(newTenant);

        IUserSpace userSpaceObj = IUserSpace(contentSpace);
        address payable userWallet = address(uint160(userSpaceObj.userWallets(tx.origin)));

        // due to a (intentional?) limitation of the EVM, this stanza *needs* to be duplicated as-is. in particular,
        //  factoring this code into a shared subroutine will cause it to fail.
        bool isV3Contract = Utils.checkV3Contract(userWallet);
        if (!isV3Contract) {
            AccessIndexor index = AccessIndexor(userWallet);
            theTenant.transferOwnership(tx.origin);
            index.setContractRights(address(newTenant), 0, 2);
        } else {
            // v3+ path ...
            theTenant.setRights(tx.origin, 0 /*TYPE_SEE*/, 2 /*ACCESS_CONFIRMED*/); // register tenant in user wallet
            theTenant.transferOwnership(tx.origin);
        }

        emit CreateTenant(newTenant);
        return newTenant;
    }
}

library factoryHelper {
    using strings for *;

    bytes public constant tenantCreatorMetaKey = bytes("_ELV_GROUP_TENANT_AUTHORITIES");
    bytes16 private constant expectedVersionSubString ="BaseContentSpace";

    function isValidTenantCreator(address payable _spaceAddr) internal view returns (bool) {

        BaseContentSpace space = BaseContentSpace(_spaceAddr);
        bytes memory value = space.getMeta(tenantCreatorMetaKey);

        // key not set
        if (value.length == 0) {
            return false;
        }

        strings.slice memory s = string(value).toSlice();
        strings.slice memory delim = ",".toSlice();
        string[] memory parts = new string[](s.count(delim) + 1);

        for (uint i = 0; i < parts.length; i++) {
            parts[i] = s.split(delim).toString();
            address gaddr = Address.bytesToAddress(bytes(parts[i]));
            address payable grpAddr = address(uint160(gaddr));

            // check if msg.sender is member of the group
            // TODO: should we check checkAccessRights...
            BaseAccessControlGroup theGroup = BaseAccessControlGroup(grpAddr);
            bool isMember = theGroup.membersMap(msg.sender);
            if (isMember) {
                return true;
            }
        }
        return false;
    }

    function isBaseContentSpaceContract(address payable _spaceAddr) internal view returns (bool){
        // reverts if no bytecode is found
        BaseContentSpace space = BaseContentSpace(_spaceAddr);
        bytes32 version = space.version();

        if (version != "") {
            // cast it to get first 16 bytes
            bytes16 substr = bytes16(version);
            if (substr == expectedVersionSubString) {
                return true;
            }
        }

        return false;
    }
}

