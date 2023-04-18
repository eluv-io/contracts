// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.5.4;

import "../../base_content_space.sol";
import "../../base_tenant_factory.sol";
import {AccessIndexor} from "../../access_indexor.sol";
import "./og-forge-std/OgTest.sol";
import "../../tradable/Strings.sol";

contract BaseContentSpaceTest is OgTest {
    BaseContentSpace space;
    BaseTypeFactory typeFactory;
    BaseGroupFactory groupFactory;
    BaseLibraryFactory libFactory;
    ContentFactoryHelper cfHelper;
    BaseContentFactory contentFactory;
    BaseAccessWalletFactory walletFactory;
    BaseTenantFactory tenantFactory;

    address payable spaceAddr;

    address payable spaceCreator = address(uint160(vm.addr(0x1)));
    address libraryCreator = vm.addr(0x2);
    address tenantFactoryCreator = vm.addr(0x3);
    address kmsAddr = vm.addr(0x4);

    address groupCreator = vm.addr(0x5);
    address payable groupUser1 = address(uint160(vm.addr(0x6)));
    address payable groupUser2 = address(uint160(vm.addr(0x7)));

    uint64 salt = 12345;

    function setUp() public {
        vm.startPrank(spaceCreator, spaceCreator);
        space = new BaseContentSpace("test_space");
        spaceAddr = address(uint160(address(space)));

        typeFactory = new BaseTypeFactory(spaceAddr);
        groupFactory = new BaseGroupFactory(spaceAddr);
        libFactory = new BaseLibraryFactory(spaceAddr);
        cfHelper = new ContentFactoryHelper();
        address payable cfHelperAddr = address(uint160(address(cfHelper)));
        contentFactory = new BaseContentFactory(spaceAddr, cfHelperAddr);
        walletFactory = new BaseAccessWalletFactory(spaceAddr);

        space.setFactory(address(typeFactory));
        space.setWalletFactory(address(walletFactory));
        space.setGroupFactory(address(groupFactory));
        space.setLibraryFactory(address(libFactory));
        space.setContentFactory(address(contentFactory));
        vm.stopPrank();

        // make access wallet for all the keys
        vm.prank(spaceCreator);
        makeAccessWallet();
        vm.prank(libraryCreator);
        makeAccessWallet();
        vm.prank(tenantFactoryCreator);
        makeAccessWallet();
        vm.prank(kmsAddr);
        makeAccessWallet();
        vm.prank(groupCreator);
        makeAccessWallet();
        vm.prank(groupUser1);
        makeAccessWallet();
        vm.prank(groupUser2);
        makeAccessWallet();
    }

    function testOwner() public {
        assertEq(space.owner(), spaceCreator);
    }

    function testCreateLibrary() public {
        vm.startPrank(libraryCreator, libraryCreator);
        address libraryAddr = space.createLibrary(kmsAddr);
        console.log("library Address:", libraryAddr);
        vm.stopPrank();
    }

    function testCreateTenantFactoryAndBaseTenantSpace() public {
        // create the groups
        vm.startPrank(groupCreator, groupCreator);
        address group1 = createGroupAndAddUsers();
        address group2 = createGroupAndAddUsers();
        vm.stopPrank();
        bytes memory value = bytes(Strings.strConcat(vm.toString(group1), ",", vm.toString(group2)));

        vm.startPrank(spaceCreator, spaceCreator);
        bytes memory key = bytes("_ELV_GROUP_TENANT_AUTHORITIES");
        space.putMeta(key, value);
        bytes memory meta = space.getMeta(key);
        console.log("meta", string(meta));
        vm.stopPrank();

        vm.startPrank(groupUser1, groupUser1);
        tenantFactory = new BaseTenantFactory(spaceAddr);
        vm.stopPrank();

        vm.startPrank(groupUser1, groupUser1);
        address tenantAddress = tenantFactory.createTenant("TST_TNT", kmsAddr, salt);
        console.log(tenantAddress);
        vm.stopPrank();

        salt = 1111;
        vm.startPrank(groupUser1, groupUser1);
        tenantAddress = tenantFactory.createTenant("TST_TNT", kmsAddr, salt);
        console.log(tenantAddress);
        vm.stopPrank();
    }

    function testRevert_CreateBaseTenantSpaceWithSameSalt() public {
        // create the groups
        vm.startPrank(groupCreator, groupCreator);
        address group1 = createGroupAndAddUsers();
        address group2 = createGroupAndAddUsers();
        vm.stopPrank();
        bytes memory value = bytes(Strings.strConcat(vm.toString(group1), ",", vm.toString(group2)));

        vm.startPrank(spaceCreator, spaceCreator);
        bytes memory key = bytes("_ELV_GROUP_TENANT_AUTHORITIES");
        space.putMeta(key, value);
        bytes memory meta = space.getMeta(key);
        console.log("meta", string(meta));
        vm.stopPrank();

        vm.startPrank(groupUser1, groupUser1);
        tenantFactory = new BaseTenantFactory(spaceAddr);
        vm.stopPrank();

        vm.startPrank(groupUser1, groupUser1);
        address tenantAddress = tenantFactory.createTenant("TST_TNT", kmsAddr, salt);
        console.log(tenantAddress);
        vm.stopPrank();

        vm.startPrank(groupUser1, groupUser1);
        vm.expectRevert();
        tenantAddress = tenantFactory.createTenant("TST_TNT", kmsAddr, salt);
        console.log(tenantAddress);
        vm.stopPrank();
    }

    function testRevert_CreateTenantFactoryWhenMetaNotSet() public {
        // meta not set
        vm.startPrank(groupUser1, groupUser1);
        vm.expectRevert("tenant_factory_creator(msg.sender) invalid");
        tenantFactory = new BaseTenantFactory(spaceAddr);
        vm.stopPrank();
    }

    function testRevert_CreateTenantFactoryWhenNotPartOfGroup() public {
        // create the groups
        vm.startPrank(groupCreator, groupCreator);
        address group1 = createGroupAndAddUsers();
        address group2 = createGroupAndAddUsers();
        vm.stopPrank();
        bytes memory value = bytes(Strings.strConcat(vm.toString(group1), ",", vm.toString(group2)));

        vm.startPrank(spaceCreator, spaceCreator);
        bytes memory key = bytes("_ELV_GROUP_TENANT_AUTHORITIES");
        space.putMeta(key, value);
        bytes memory meta = space.getMeta(key);
        console.log("meta", string(meta));
        vm.stopPrank();

        vm.startPrank(spaceCreator, spaceCreator);
        vm.expectRevert("tenant_factory_creator(msg.sender) invalid");
        tenantFactory = new BaseTenantFactory(spaceAddr);
        vm.stopPrank();
    }

    function testCreateTenantFactoryBySpace() public {
        vm.startPrank(spaceAddr, spaceAddr);
        tenantFactory = new BaseTenantFactory(spaceAddr);
        vm.stopPrank();
    }

    function testFail_CreateTenantFactoryWhenSpaceAddressInvalid() public {
        vm.startPrank(spaceCreator, spaceCreator);
        tenantFactory = new BaseTenantFactory(spaceCreator);
        vm.stopPrank();
    }

    // ==================
    // helper funcs

    function makeAccessWallet() private returns (address) {
        return space.getAccessWallet();
    }

    function createGroupAndAddUsers() private returns (address) {
        address payable groupAddr = address(uint160(space.createGroup()));
        console.log("group Address:", groupAddr);
        BaseAccessControlGroup theGroup = BaseAccessControlGroup(groupAddr);

        theGroup.grantAccess(groupUser1);
        theGroup.grantAccess(groupUser2);
        return groupAddr;
    }
}
