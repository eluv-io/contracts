pragma solidity 0.5.4;

import {Container} from "./container.sol";
import {BaseAccessControlGroup} from "./base_access_control_group.sol";
import {BaseContent} from "./base_content.sol";
import {IFactorySpace, INodeSpace} from "./base_space_interfaces.sol";
//import "./access_indexor.sol";
import "./meta_object.sol";

/* -- Revision history --
BaseLibrary20190221101700ML: First versioned released
BaseLibrary20190318101300ML: Migrated to 0.4.24
BaseLibrary20190506153700ML: Adds access indexing
BaseLibrary20190510151800ML: Modified createContent to use contentspace factory
BaseLibrary20190515103800ML: Overloads canPublish to take into account EDIT privilege granted for update request and commit
BaseLibrary20190522154000SS: Changed hash bytes32 to string
BaseLibrary20190523121700ML: Fixes logic of add/remove of groups to revert to compact arrays
BaseLibrary20190528151200ML: Uses Container abstraction
BaseLibrary20190605150200ML: Splits out canConfirm from canPublish
BaseLibrary20191010140800ML: Content can be deleted by content owner or the library owner
BaseLibrary20200110162700ML: Adds support for visibility, differentiates rights to access and edit library object and content
BaseLibrary20200211164300ML: Modified to conform to authV3 API
BaseLibrary20200316135200ML: Leverages inherited hasAccess
BaseLibrary20200928110000PO: Replace tx.origin with msg.sender in some cases
*/


contract BaseLibrary is MetaObject, Container {

    bytes32 public version ="BaseLibrary20200928110000PO"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    address payable[] public contributorGroups;
    address payable[] public reviewerGroups;
    address payable[] public accessorGroups;
    uint256 public contributorGroupsLength = 0;
    uint256 public reviewerGroupsLength = 0;
    uint256 public accessorGroupsLength = 0;

    address payable[] public approvalRequests;
    uint256 public approvalRequestsLength = 0;
    mapping (address => uint256) private approvalRequestsMap; //index offset by 1 to avoid confusing 0 for removed

    event ContentObjectCreated(address contentAddress, address content_type, address spaceAddress);
    event ContentObjectDeleted(address contentAddress, address spaceAddress);
    event ContributorGroupAdded(address group);
    event ContributorGroupRemoved(address group);
    event ReviewerGroupAdded(address group);
    event ReviewerGroupRemoved(address group);
    event AccessorGroupAdded(address group);
    event AccessorGroupRemoved(address group);
    event UnauthorizedOperation(uint operationCode, address candidate);
    event ApproveContentRequest(address contentAddress, address submitter);
    event ApproveContent(address contentAddress, bool approved, string note);
    event UpdateKmsAddress(address addressKms);

    constructor(address address_KMS, address payable _content_space) public payable {
        contentSpace = _content_space;
        addressKMS = address_KMS;
        visibility = 0;
        indexCategory = 3; // AccessIndexor CATEGORY_LIBRARY
    }

    function canConfirm() public view returns (bool) {
        INodeSpace bcs = INodeSpace(contentSpace);
        return bcs.canNodePublish(msg.sender);
    }

    function canPublish() public view returns (bool) {
        if ((msg.sender == owner) || canEdit()) {
            return true;
        }
        return false;
    }

    function canCommit() public view returns (bool) {
        return canEdit();
    }

    function addToGroupList(address payable _addGroup, address payable[] storage _groupList, uint256 _groupListLength) internal returns (uint256) {
        for (uint256 i = 0; i < _groupListLength; i++) {
            if (_addGroup == _groupList[i]) {
                return _groupListLength;
            }
        }
        if (_groupListLength < _groupList.length) {
            _groupList[_groupListLength] = _addGroup;
        } else {
            _groupList.push(_addGroup);
        }
        return (_groupListLength + 1);
    }

    function removeFromGroupList(address payable _removeGroup, address payable[] storage _groupList, uint256 _groupListLength) internal returns (uint256) {
        for (uint256 i = 0; i < _groupListLength; i++) {
            if (_removeGroup == _groupList[i]) {
                delete _groupList[i];
                if (i != _groupListLength - 1) {
                    _groupList[i] = _groupList[_groupListLength - 1];
                    delete _groupList[_groupListLength - 1];
                }
                return (_groupListLength - 1);
            }
        }
        return _groupListLength;
    }

    function addContributorGroup(address payable group) public onlyEditor {
        uint256 prevLen = contributorGroupsLength;
        contributorGroupsLength = addToGroupList(group, contributorGroups, prevLen);
        if (contributorGroupsLength > prevLen) {
            emit ContributorGroupAdded(group);
            /*
            AccessIndexor accessIndex = AccessIndexor(group);
            accessIndex.setLibraryRights(address(this), accessIndex.TYPE_SEE(), accessIndex.ACCESS_TENTATIVE());
            */
            setRights(group, 0 /*AccessIndexor TYPE_SEE*/, 1 /*AccessIndexor ACCESS_TENTATIVE*/);
        }
    }

    function groupIsListed(address payable group, address payable[] memory groups) internal pure returns (bool) {
       for (uint256 i = 0; i < groups.length; i++) {
           if (group == groups[i]) {
               return true;
           }
       }
       return false;
    }

    function removeContributorGroup(address payable group) public onlyEditor returns (bool) {
        uint256 prevLen = contributorGroupsLength;
        contributorGroupsLength = removeFromGroupList(group, contributorGroups, prevLen);
        if (contributorGroupsLength < prevLen) {
            emit ContributorGroupRemoved(group);
            /*
            AccessIndexor accessIndex = AccessIndexor(group);
            accessIndex.setLibraryRights(address(this), accessIndex.TYPE_SEE(), accessIndex.ACCESS_NONE());
            */
            setRights(group, 0 /*AccessIndexor TYPE_SEE*/, 0 /*AccessIndexor ACCESS_TENTATIVE*/);
            return true;
        }
        return false;
    }

    function addReviewerGroup(address payable group) public onlyEditor {
        uint256 prevLen = reviewerGroupsLength;
        reviewerGroupsLength = addToGroupList(group, reviewerGroups, prevLen);
        if (reviewerGroupsLength > prevLen) {
            emit ReviewerGroupAdded(group);
            /*
            AccessIndexor accessIndex = AccessIndexor(group);
            accessIndex.setLibraryRights(address(this), accessIndex.TYPE_ACCESS(), accessIndex.ACCESS_TENTATIVE());
            */
            setRights(group, 1 /*AccessIndexor TYPE_ACCESS*/, 1 /*AccessIndexor ACCESS_TENTATIVE*/);
        }
    }

    function removeReviewerGroup(address payable group) public onlyEditor returns (bool) {
        uint256 prevLen = reviewerGroupsLength;
        reviewerGroupsLength = removeFromGroupList(group, reviewerGroups, prevLen);
        if (reviewerGroupsLength < prevLen) {
            emit ReviewerGroupRemoved(group);
            if (!groupIsListed(group, accessorGroups)) {
                /*
	              AccessIndexor accessIndex = AccessIndexor(group);
                accessIndex.setLibraryRights(address(this), accessIndex.TYPE_ACCESS(), accessIndex.ACCESS_NONE());
                */
                setRights(group, 1 /*AccessIndexor TYPE_ACCESS*/, 0 /*AccessIndexor ACCESS_NONE*/);
            }
            return true;
        }
        return false;

    }

    function addAccessorGroup(address payable group) public onlyEditor {
        uint256 prevLen = accessorGroupsLength;
        accessorGroupsLength = addToGroupList(group, accessorGroups, prevLen);
        if (accessorGroupsLength > prevLen) {
            emit AccessorGroupAdded(group);
            /*
            AccessIndexor accessIndex = AccessIndexor(group);
            accessIndex.setLibraryRights(address(this), accessIndex.TYPE_ACCESS(), accessIndex.ACCESS_TENTATIVE());
            */
            setRights(group, 1 /*AccessIndexor TYPE_ACCESS*/,1 /*AccessIndexor ACCESS_TENTATIVE*/);
        }
    }

    function removeAccessorGroup(address payable group) public onlyEditor returns (bool) {
        uint256 prevLen = accessorGroupsLength;
        accessorGroupsLength = removeFromGroupList(group, accessorGroups, prevLen);
        if (accessorGroupsLength < prevLen) {
            emit AccessorGroupRemoved(group);
            if (!groupIsListed(group, reviewerGroups)) {
                /*
                AccessIndexor accessIndex = AccessIndexor(group);
                accessIndex.setLibraryRights(address(this), accessIndex.TYPE_ACCESS(), accessIndex.ACCESS_NONE());
                */
                setRights(group, 1 /*AccessIndexor TYPE_ACCESS*/, 0 /*AccessIndexor ACCESS_NONE*/);
            }
        }
        return false;
    }

    function hasGroupAccess(address _candidate, address payable[] memory _groupList) internal view returns (bool) {
        for (uint i = 0; i < _groupList.length; i++) {
            if (_groupList[i] != address(0x0)) {
                BaseAccessControlGroup groupContract = BaseAccessControlGroup(_groupList[i]);
                if (groupContract.hasAccess(_candidate)) {
                    return true;
                }
            }
        }
        return false;
    }
/*
    function hasAccess(address _candidate) public view returns (bool) {
        if ((visibility >= 10) || (_candidate == owner)) {
            return true;
        }
        address userWallet = IUserSpace(contentSpace).userWallets(_candidate);
        AccessIndexor wallet = AccessIndexor(userWallet);
        return wallet.checkLibraryRights(address(this), wallet.TYPE_ACCESS());
    }
    */

    function canContribute(address payable _candidate) public view returns (bool) {
        return hasEditorRight(_candidate) || hasGroupAccess(_candidate, contributorGroups);
    }

    function canReview(address _candidate) public view returns (bool) {
        if (_candidate == owner) {
            return true;
        }
        return hasGroupAccess(_candidate, reviewerGroups);
    }

    function requiresReview() public view returns (bool) {
        return (reviewerGroupsLength > 0);
    }

    function submitApprovalRequest() public returns (bool) {
        address payable contentContract = msg.sender;
        BaseContent c = BaseContent(contentContract);

        if (requiresReview() == false) { //No review required
            // 0 indicates approval, custom contract might overwrite that decision
            c.updateStatus(0);

            // Log event
            emit ApproveContent(contentContract, true, "");
            return true;
        }
        if (approvalRequestsMap[contentContract] != 0) {
            return false;
        }
        // Create a new approval request and add to pending list
        if (approvalRequestsLength < approvalRequests.length) {
            approvalRequests[approvalRequestsLength] = contentContract;
        } else {
            approvalRequests.push(contentContract);
        }
        approvalRequestsMap[contentContract] = approvalRequestsLength + 1;
        approvalRequestsLength++;

        // Log event
        emit ApproveContentRequest(contentContract, msg.sender);
        return true;
    }

    function getPendingApprovalRequest(uint256 index) public view returns (address) {
        // Read and process the first content in the approvalRequests list
        if ((approvalRequestsLength == 0) || (approvalRequestsLength <= index)) {
            return address(0x0);
        }
        return approvalRequests[index];
    }

    function approveContent(address payable content_contract, bool approved, string memory note) public returns ( bool ) {
        require(canReview(msg.sender) == true);
        // credit the account based on the percent_complete
        uint256 index = approvalRequestsMap[content_contract] - 1;
        BaseContent c = BaseContent(content_contract);

        // remove the request from the list
        delete approvalRequests[index];
        approvalRequestsLength--;
        approvalRequestsMap[content_contract] = 0;
        if (approvalRequestsLength > index) {
            address payable lastRequest = approvalRequests[approvalRequestsLength];
            approvalRequests[index] = lastRequest;
            delete approvalRequests[approvalRequestsLength];
            approvalRequestsMap[lastRequest] = index + 1;
        }

        int currentStatus = c.statusCode();
        if (currentStatus > 0) {
            int newStatusCode;
            if (approved == true) {
                newStatusCode = 0; // indicates approval, custom contract might overwrite that decision
            } else {
                newStatusCode = -1; // alternatively we could use more complex logic like (currentStatus + 1) * -1
            }
            c.updateStatus(newStatusCode);

            // Log event
            emit ApproveContent(content_contract, approved, note);
            return true;
        } else {
            return false;
        }
    }

    function createContent(address payable content_type) public  returns (address) {
        require(msg.sender == tx.origin, "only direct calls allowed");
        address content = IFactorySpace(contentSpace).createContent(address(this), content_type);
        emit ContentObjectCreated(content, content_type, contentSpace);
        return content;
    }

    // content can be deleted by content owner or the library owner - enforced inside the kill
    function deleteContent(address payable _contentAddr) public onlyEditor {
        BaseContent content = BaseContent(_contentAddr);
        content.kill();
        emit ContentObjectDeleted(_contentAddr, contentSpace);
    }

    function publish(address payable contentObj) public returns (bool) {
        require(msg.sender == contentObj);
        BaseContent content = BaseContent(contentObj);
        // Update the content contract to reflect the approval process
        content.updateStatus(1); //update status to in-review
        // mark with statusCode 1, which is the default for in-review - NOTE: could be change to be (currentStatus * -1)
        bool submitStatus = false;
        if (content.statusCode() > 0) {
            submitStatus = submitApprovalRequest();
        }
        return submitStatus;
    }

    function updateAddressKMS(address address_KMS) public onlyEditor {
        addressKMS = address_KMS;
        emit UpdateKmsAddress(addressKMS);
    }

}
