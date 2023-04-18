pragma solidity 0.5.4;

import {Ownable} from "./ownable.sol";
import {Utils} from "./utils.sol";
import {AccessIndexor} from "./access_indexor.sol";
import "./base_content_space.sol";
import "./base_tenant_space.sol";
import "./user_space.sol";
import "./base_access_control_group.sol";
import "./strings_utils.sol";
import "./address_utils.sol";

/* -- Revision history --
BaseTenantFactory20230404000000PM: First versioned released
*/

contract BaseTenantFactory is Ownable {
    bytes32 public version = "BsTenantFactory20230404120000PM"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    event CreateTenant(address tenantAddress);

    constructor(address payable _spaceAddr) public {
        require(
            (msg.sender == _spaceAddr || tenantFactoryHelper.isValidTenantCreator(_spaceAddr)),
            "tenant_factory_creator(msg.sender) invalid"
        );
        if (msg.sender == _spaceAddr) {
            require(tenantFactoryHelper.isBaseContentSpaceContract(_spaceAddr), "msg.sender != content_space_address");
        }
        contentSpace = _spaceAddr;
    }

    /*
     * 1. Checks if the msg.sender is content_space or
     *    is part of the group address provided in space
     *    metaObject '_ELV_GROUP_TENANT_AUTHORITIES'
     * 2. create baseTenantSpace contract using CREATE2 EVM opcode ( has indexCaterory set to CATEGORY_CONTRACT())
     * 3. set rights to tx.origin for baseTenantSpace contract
     * 4. if _salt > 0, tenant_address is calcualted as 'address++zero++nonce' else 'address++block.number++zero'
     *    for more info, check tenantFactoryHelper.saltFor() library method
    */
    function createTenant(string memory _tenantName, address _kmsAddr, uint64 _salt) public returns (address) {
        //address payable newTenant = address(new BaseTenantSpace(contentSpace, _tenantName, _kmsAddr));
        address payable newTenant = deployTenant(_tenantName, _kmsAddr, _salt);
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
            theTenant.setRights(tx.origin, 0, /*TYPE_SEE*/ 2 /*ACCESS_CONFIRMED*/ ); // register tenant in user wallet
            theTenant.transferOwnership(tx.origin);
        }

        emit CreateTenant(newTenant);
        return newTenant;
    }

    function deployTenant(string memory _tenantName, address _kmsAddr, uint64 _salt)
        private
        returns (address payable)
    {
        address payable addr;
        bytes memory baseTenantSpaceByteCode = type(BaseTenantSpace).creationCode;
        bytes memory constructorData = abi.encode(contentSpace, _tenantName, _kmsAddr);
        bytes memory encodeBytes = abi.encodePacked(baseTenantSpaceByteCode, constructorData);
        address payable tenantAddr = tenantFactoryHelper.deploy(encodeBytes, tenantFactoryHelper.saltFor(tx.origin, _salt));
        return tenantAddr;
    }
}

library tenantFactoryHelper {
    using strings for *;

    bytes public constant tenantCreatorMetaKey = bytes("_ELV_GROUP_TENANT_AUTHORITIES");
    bytes16 private constant expectedVersionSubString = "BaseContentSpace";

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

        for (uint256 i = 0; i < parts.length; i++) {
            parts[i] = s.split(delim).toString();
            address gaddr = Address.bytesToAddress(bytes(parts[i]));
            address payable grpAddr = address(uint160(gaddr));

            // check if msg.sender is member of the group
            // TODO: should we check checkAccessRights...
            BaseAccessControlGroup theGroup = BaseAccessControlGroup(grpAddr);
            bool isMember = theGroup.membersMap(tx.origin);
            if (isMember) {
                return true;
            }
        }
        return false;
    }

    function isBaseContentSpaceContract(address payable _spaceAddr) internal view returns (bool) {
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

    // saltFor computes a salt for the given address as: 'address++zero++nonce'.
    // If the provided nonce is zero, block.number is used and the salt becomes:
    // 'address++block.number++zero'. Hence the layout is 20 bytes for address,
    // 8 bytes for block number and 4 bytes for nonce.
    // The function is public for testing purposes.
    function saltFor(address addr, uint64 nonce) internal view returns (uint256) {
        uint256 ret = uint256(uint160(addr));
        ret <<= 96;
        uint256 blockNumber;
        if (nonce == 0) {
            // use block number but shift it to avoid collision on young chain
            // note that block.number has to fit into a uint64
            // see go-ethereum@v1.10.19/core/types/block.go in Header.SanityCheck
            blockNumber = block.number;
            blockNumber <<= 32;
        }
        return uint256(ret) + blockNumber + nonce;
    }

    function deploy(bytes memory bytecode, uint256 _salt) internal returns (address payable _addr) {
        /*
      NOTE: How to call create2
      create2(v, p, n, s)
      create new contract with code at memory p to p + n
      and send v wei
      and return the new address
      where new address = first 20 bytes of keccak256(0xff + address(this) + s + keccak256(mem[p…(p+n)))
            s = big-endian 256-bit value
      */
        assembly {
            _addr := create2(0, add(bytecode, 0x20), mload(bytecode), _salt)

            if iszero(extcodesize(_addr)) {
                revert(0, 0)
            }
        }
    }
}
