pragma solidity 0.4.24;

import "./ownable.sol";
import {UserSpace} from "./user_space.sol";

/* -- Revision history --
AccessIndexor20190423111800ML: First versioned release
*/


contract AccessIndexor is Ownable {

    bytes32 public version = "AccessIndexor20190506155200ML";

    event RightsChanged(address principal, address entity, uint8 aggregate);
    //event dbUint8(string label, uint8 value);

    address public contentSpace;

    uint8 public CATEGORY_CONTENT_OBJECT = 1;
    uint8 public CATEGORY_GROUP = 2;
    uint8 public CATEGORY_LIBRARY = 3;
    uint8 public CATEGORY_CONTENT_TYPE = 4;
    uint8 public CATEGORY_CONTRACT = 5;

    uint8 public constant TYPE_SEE = 0;
    uint8 public constant TYPE_ACCESS = 1;
    uint8 public constant TYPE_EDIT = 2;
    uint8 public constant ACCESS_NONE = 0; // No access
    uint8 public constant ACCESS_TENTATIVE = 1; //access was granted, but not accepted by recipient
    uint8 public constant ACCESS_CONFIRMED = 2; //access was granted and accepted by recipient
    uint8[3] ranking = [1,10,100];

    struct AccessIndex {
        uint8 category;
        mapping (address => uint8) rights; // rights map for a given object address
        address[] list; // All items for which rights are known
        uint256 length; // Number of items listed
    }

    //mapping (address => uint8) libraryRights;
    //address[] libraryList;
    AccessIndex public libraries; // = AccessIndex(libraryRights, libraryList, 0);
    AccessIndex public contentObjects;
    AccessIndex public accessGroups;
    AccessIndex public contentTypes;
    AccessIndex public contracts;

/*
    function isLibraryVisible(address lib) public view returns (bool) {
        return (libraries[lib] >= TYPE_SEE);
    }

    function isLibraryAccessible(address lib) public view returns (bool) {
        return (libraries[lib] >= TYPE_ACCESS);
    }

    function isLibraryEditable(address lib) public view returns (bool) {
        return (libraries[lib] >= TYPE_EDIT);
    }
*/

    constructor() public payable {
        libraries.category = CATEGORY_LIBRARY;
        accessGroups.category = CATEGORY_GROUP;
        contentObjects.category = CATEGORY_CONTENT_OBJECT;
        contentTypes.category = CATEGORY_CONTENT_TYPE;
        contracts.category = CATEGORY_CONTRACT;
    }

    function setContentSpace(address content_space) public onlyOwner {
        contentSpace = content_space;
    }

    function setLibraryRights(address lib, uint8 access_type, uint8 access) public  {
        setRights(libraries, lib, access_type, access);
    }

    function setAccessGroupRights(address group, uint8 access_type, uint8 access) public  {
        setRights(accessGroups, group, access_type, access);
    }

    function setContentObjectRights(address obj, uint8 access_type, uint8 access) public  {
        setRights(contentObjects, obj, access_type, access);
    }

    function setContentTypeRights(address obj, uint8 access_type, uint8 access) public  {
        setRights(contentTypes, obj, access_type, access);
    }

    function setContractRights(address obj, uint8 access_type, uint8 access) public  {
        setRights(contracts, obj, access_type, access);
    }

    function checkLibraryRights(address lib, uint8 access_type) public view returns(bool) {
        return checkRights(CATEGORY_LIBRARY, lib, access_type);
    }

    function checkAccessGroupRights(address group, uint8 access_type) public view returns(bool) {
        return checkRights(CATEGORY_GROUP, group, access_type);
    }

    function checkContentObjectRights(address group, uint8 access_type) public view returns(bool) {
        return checkRights(CATEGORY_CONTENT_OBJECT, group, access_type);
    }
    /*
    event dbAddress(string label, address addr);
    function checkContentObjectRights(address obj, uint8 access_type) public view returns(bool) {
        emit dbAddress("start", obj);
        bool directRights = checkRights(contentObjects, obj, access_type);
        if (directRights == true) {
            emit dbAddress("direct", address(this));
            return true;
        } else {
            AccessIndexor groupObj;
            address group;
            for (uint i = 0; i < accessGroups.length; i++) {
                group = accessGroups.list[i];
                if (group != 0x0) {
                    groupObj = AccessIndexor(group);
                    if (groupObj.checkContentObjectRights(obj, access_type) == true) {
                        emit dbAddress("group", group);
                        return true;
                    }
                }
            }
        }
        emit dbAddress("end", obj);
        return false;
    }
*/

    function checkContentTypeRights(address obj, uint8 access_type) public view returns(bool) {
        return checkRights(CATEGORY_CONTENT_TYPE, obj, access_type);
    }

    function checkContractRights(address obj, uint8 access_type) public view returns(bool) {
        return checkRights(CATEGORY_CONTRACT, obj, access_type);
    }

    function getLibraryRights(address lib) public view returns(uint8) {
        return libraries.rights[lib];
    }

    function getLibrariesLength() public view returns(uint256) {
        return libraries.length;
    }

    function getLibrary(uint256 position) public view returns(address) {
        return libraries.list[position];
    }

    function getAccessGroupRights(address group) public view returns(uint8) {
        return accessGroups.rights[group];
    }

    function getAccessGroupsLength() public view returns(uint256) {
        return accessGroups.length;
    }

    function getAccessGroup(uint256 position) public view returns(address) {
        return accessGroups.list[position];
    }

    function getContentObjectRights(address obj) public view returns(uint8) {
        return contentObjects.rights[obj];
    }

    function getContentObjectsLength() public view returns(uint256) {
        return contentObjects.length;
    }

    function getContentObject(uint256 position) public view returns(address) {
        return contentObjects.list[position];
    }

    function getContentTypeRights(address obj) public view returns(uint8) {
        return contentTypes.rights[obj];
    }

    function getContentTypesLength() public view returns(uint256) {
        return contentTypes.length;
    }

    function getContentType(uint256 position) public view returns(address) {
        return contentTypes.list[position];
    }

    function getContractRights(address obj) public view returns(uint8) {
        return contracts.rights[obj];
    }

    function getContractsLength() public view returns(uint256) {
        return contracts.length;
    }

    function getContract(uint256 position) public view returns(address) {
        return contracts.list[position];
    }

    function hasManagerAccess(address candidate) public view returns (bool) {
        return (candidate == owner);
    }

    function checkDirectRights(uint8 index_type, address obj, uint8 access_type) public view returns(bool) {
        Ownable instance = Ownable(obj);
        if (instance.owner() == tx.origin){
            return true;
        }
        if (index_type == CATEGORY_CONTENT_OBJECT) {
            return checkRawRights(contentObjects, obj, access_type);
        }
        if (index_type == CATEGORY_GROUP) {
            return checkRawRights(accessGroups, obj, access_type);
        }
        if (index_type == CATEGORY_LIBRARY) {
            return checkRawRights(libraries, obj, access_type);
        }
        if (index_type == CATEGORY_CONTRACT) {
            return checkRawRights(contracts, obj, access_type);
        }
        if (index_type == CATEGORY_CONTENT_TYPE) {
            return checkRawRights(contentTypes, obj, access_type);
        }
    }

    function checkRawRights(AccessIndex storage index, address obj, uint8 access_type) private view returns(bool) {
        uint8 currentAggregate = index.rights[obj];
        return (currentAggregate >= ranking[access_type]);
    }

    function checkRights(uint8 index_type, address obj, uint8 access_type) public view returns(bool) {
        bool directRights = checkDirectRights(index_type, obj, access_type);
        if (directRights == true) {
            return true;
        } else {
            AccessIndexor groupObj;
            address group;
            for (uint i = 0; i < accessGroups.length; i++) {
                group = accessGroups.list[i];
                if (group != 0x0) {
                    groupObj = AccessIndexor(group);
                    if (groupObj.checkDirectRights(index_type, obj, access_type) == true) {
                        return true;
                    }
                }
            }
        }
        Ownable instance = Ownable(obj);
        return (owner == instance.owner());
    }

    event dbAddress(string label, address addr);
    function setRights(AccessIndex storage index, address obj, uint8 access_type, uint8 access) private  {
        UserSpace space = UserSpace(contentSpace);
        address walletAddress = space.userWallets(tx.origin);
        if  (walletAddress == 0x0) {
            Ownable instance = Ownable(obj);
            require(tx.origin == instance.owner());
        }
        if (walletAddress != 0x0) {
            AccessIndexor wallet = AccessIndexor(walletAddress);
            require(wallet.checkRights(index.category, obj, TYPE_EDIT));
        }

        uint8 currentAggregate = index.rights[obj];
        uint8[3] memory currentRights;
        currentRights[0] = currentAggregate % 10;
        currentRights[1] = currentAggregate % 100 - currentRights[0];
        currentRights[2] = currentAggregate - currentRights[1] - currentRights[0];

        if (hasManagerAccess(tx.origin) == true) {
            //manager can revoke
            if (access == ACCESS_NONE) {
                currentRights[access_type] = 0;
            }

            //manager can confirm or infirm something tentative
            if ((currentRights[access_type] == ACCESS_TENTATIVE) && (access != ACCESS_TENTATIVE)) {
                currentRights[access_type] = access * ranking[access_type];
            }

            //if access grant is tentative, but user is manager, then make it confirmed
            if (access >= ACCESS_TENTATIVE) {
                currentRights[access_type] = ACCESS_CONFIRMED * ranking[access_type];
            }

            //to be discussed, manager can also adds itself viewing privileges, acting as a bookmark.
            if ((access_type == TYPE_SEE)  && (access != ACCESS_NONE)) {
                currentRights[access_type] = ACCESS_CONFIRMED;
            }
        } else {

            //library can add tentative rights, unless rights are already confirmed
            if ((access >= ACCESS_TENTATIVE) && (currentRights[access_type] == 0)) {
                currentRights[access_type] = ACCESS_TENTATIVE * ranking[access_type];
            }

            //library can withdraw privilege
            if (access == ACCESS_NONE) {
                currentRights[access_type] = ACCESS_NONE;
            }
        }
        uint8 newAggregate = currentRights[0] + currentRights[1] + currentRights[2];
        index.rights[obj] = newAggregate;

        //add to array if newly added
        if ((newAggregate != 0) && (currentAggregate == 0)) {
            addToList(index, obj);
        }

        //remove from array if removed
        if (newAggregate == 0) {
            removeFromList(index, obj);
        }

        emit RightsChanged(address(this), obj, newAggregate);
    }

    function addToList(AccessIndex storage index, address obj) private {
        if (index.length < index.list.length) {
            index.list[index.length] = obj;
        } else {
            index.list.push(obj);
        }
        index.length++;
    }

    function removeFromList(AccessIndex storage index, address obj) private returns (bool) {
        for (uint i = 0; i < index.length; i++) {
            if (index.list[i] == obj) {
                delete index.list[i];
                if (i != (index.length - 1)) {
                    index.list[i] = index.list[index.length - 1];
                    delete index.list[index.length - 1];
                }
                index.length--;
                return true;
            }
        }
        return false;
    }


}
