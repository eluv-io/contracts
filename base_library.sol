pragma solidity ^0.4.21;

import {Accessible} from "./accessible.sol";
import {Editable} from "./editable.sol";
import {BaseAccessControlGroup} from "./base_access_control_group.sol";
import {BaseContent} from "./base_content.sol";
import "./accessible.sol";
import "./base_content_space.sol";


contract BaseLibrary is Accessible, Editable {

    bytes32 public version ="BaseLibrary20190221101700ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    address public contentSpace;
    address[] public contributorGroups;
    address[] public reviewerGroups;
    address[] public accessorGroups;
    address[] public contentTypes;
    uint256 public contributorGroupsLength;
    uint256 public reviewerGroupsLength;
    uint256 public accessorGroupsLength;
    uint256 public contentTypesLength = 0;

    mapping ( address => address ) public contentTypeContracts;  // custom contracts map

    address[] public approvalRequests;
    address public addressKMS;
    uint256 public approvalRequestsLength = 0;
    mapping (address => uint256) private approvalRequestsMap; //index offset by 1 to avoid confusing 0 for removed

    event ContentObjectCreated(address contentAddress, address content_type);
    event ContributorGroupAdded(address group);
    event ContributorGroupRemoved(address group);
    event ReviewerGroupAdded(address group);
    event ReviewerGroupRemoved(address group);
    event AccessorGroupAdded(address group);
    event AccessorGroupRemoved(address group);
    event ContentTypeAdded(address contentType, address contentContract);
    event ContentTypeRemoved(address contentType);
    event UnauthorizedOperation(uint operationCode, address candidate);
    event ApproveContentRequest(address contentAddress, address submitter);
    event ApproveContent(address contentAddress, bool approved, string note);

    constructor(address address_KMS, address _content_space) public payable {
        contentSpace = _content_space;
        contributorGroupsLength = 0;
        reviewerGroupsLength = 0;
        accessorGroupsLength = 0;
        addressKMS = address_KMS;
    }

    function setAddressKMS(address address_KMS) public onlyOwner {
        addressKMS = address_KMS;
    }

    function addContributorGroup(address group) public onlyOwner {
        if (contributorGroupsLength < contributorGroups.length) {
            contributorGroups[contributorGroupsLength] = group;
        } else {
            contributorGroups.push(group);
        }
        contributorGroupsLength = contributorGroupsLength + 1;
        emit ContributorGroupAdded(group);
    }

    function removeContributorGroup(address group) public onlyOwner returns (bool) {
        for (uint i = 0; i < contributorGroupsLength; i++) {
            if (contributorGroups[i] == group) {
                delete contributorGroups[i];
                if (i != (contributorGroupsLength - 1)) {
                    contributorGroups[i] = contributorGroups[contributorGroupsLength - 1];
                    delete contributorGroups[contributorGroupsLength - 1];
                }
                contributorGroupsLength--;
                emit ContributorGroupRemoved(group);
                return true;
            }
        }
        return false;
    }

    function addReviewerGroup(address group) public onlyOwner {
        if (reviewerGroupsLength < reviewerGroups.length) {
            reviewerGroups[reviewerGroupsLength] = group;
        } else {
            reviewerGroups.push(group);
        }
        reviewerGroupsLength = reviewerGroupsLength + 1;
        emit ReviewerGroupAdded(group);
    }

    function removeReviewerGroup(address group) public onlyOwner returns (bool) {
        for (uint i = 0; i < reviewerGroupsLength; i++) {
            if (reviewerGroups[i] == group) {
                delete reviewerGroups[i];
                if (i != (reviewerGroupsLength - 1)) {
                    reviewerGroups[i] = reviewerGroups[reviewerGroupsLength - 1];
                    delete reviewerGroups[reviewerGroupsLength - 1];
                }
                reviewerGroupsLength--;
                emit ReviewerGroupRemoved(group);
                return true;
            }
        }
        return false;
    }

    function addAccessorGroup(address group) public onlyOwner {
        if (accessorGroupsLength < accessorGroups.length) {
            accessorGroups[accessorGroupsLength] = group;
        } else {
            accessorGroups.push(group);
        }
        accessorGroupsLength = accessorGroupsLength + 1;
        emit AccessorGroupAdded(group);
    }

    function removeAccessorGroup(address group) public onlyOwner returns (bool) {
        for (uint i = 0; i < accessorGroupsLength; i++) {
            if (accessorGroups[i] == group) {
                delete accessorGroups[i];
                if (i != (accessorGroupsLength - 1)) {
                    accessorGroups[i] = accessorGroups[accessorGroupsLength - 1];
                    delete accessorGroups[accessorGroupsLength - 1];
                }
                accessorGroupsLength--;
                emit AccessorGroupRemoved(group);
                return true;
            }
        }
        return false;
    }

    function addContentType(address content_type, address content_contract) public onlyOwner {
        if ((contentTypeContracts[content_type] == 0x0) && (validType(content_type) == false)) {
            if (contentTypesLength < contentTypes.length) {
                contentTypes[contentTypesLength] = content_type;
            } else {
                contentTypes.push(content_type);
            }
            contentTypesLength = contentTypesLength + 1;
        }
        contentTypeContracts[content_type] = content_contract;
        emit ContentTypeAdded(content_type, content_contract);
    }

    function removeContentType(address content_type) public onlyOwner returns (bool) {
        uint256 latestIndex = contentTypesLength - 1;
        for (uint256 i = 0; i < contentTypesLength; i++) {
            if (contentTypes[i] == content_type) {
                delete contentTypes[i];
                if (i != latestIndex) {
                    contentTypes[i] = contentTypes[latestIndex];
                    delete contentTypes[latestIndex];
                }
                contentTypesLength = latestIndex; //decrease by 1
                delete contentTypeContracts[content_type];
                emit ContentTypeRemoved(content_type);
                return true;
            }
        }
        return false;
    }

    function hasAccess(address candidate) public constant returns (bool) {
        if (accessorGroupsLength == 0) {
            return true;
        }
        address group;
        bool groupAccess;
        BaseAccessControlGroup groupContract;
        for (uint i = 0; i < accessorGroupsLength; i++) {
            group = accessorGroups[i];
            if (group != 0x0) {
                groupContract = BaseAccessControlGroup(group);
                groupAccess = groupContract.hasAccess(candidate);
                if (groupAccess == true) {
                    return true;
                }
            }
        }
        return false;
    }

    function canContribute(address candidate) public constant returns (bool) {
        if (contributorGroupsLength == 0) {
            return true;
        }
        address group;
        bool groupAccess;
        BaseAccessControlGroup groupContract;
        for (uint i = 0; i < contributorGroupsLength; i++) {
            group = contributorGroups[i];
            if (group != 0x0) {
                groupContract = BaseAccessControlGroup(group);
                groupAccess = groupContract.hasAccess(candidate);
                if (groupAccess == true) {
                    return true;
                }
            }
        }
        return false;
    }

    function canReview(address candidate) public constant returns (bool) {
        if (candidate == owner) {
            return true;
        }
        address group;
        bool groupAccess;
        BaseAccessControlGroup groupContract;
        for (uint i = 0; i < reviewerGroupsLength; i++) {
            group = reviewerGroups[i];
            if (group != 0x0) {
                groupContract = BaseAccessControlGroup(group);
                groupAccess = groupContract.hasAccess(candidate);
                if (groupAccess == true) {
                    return true;
                }
            }
        }
        return false;
    }

    // check whether an address - which should represent a content fabric node - can confirm (publish?) a content object
    function canNodePublish(address candidate) public view returns (bool) {
        BaseContentSpace bcs = BaseContentSpace(contentSpace);
        return bcs.canNodePublish(candidate);
    }

    function submitApprovalRequest() public returns (bool) {
        address contentContract = msg.sender;
        BaseContent c = BaseContent(contentContract);
        //require((c.owner() == tx.origin)); the publish already check authorization

        if (reviewerGroupsLength == 0) { //No review required
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
        emit ApproveContentRequest(contentContract, tx.origin);
        return true;
    }

    function getPendingApprovalRequest(uint256 index) public constant returns (address) {
        // Read and process the first content in the approvalRequests list
        if ((approvalRequestsLength == 0) || (approvalRequestsLength <= index)) {
            return 0;
        }
        return approvalRequests[index];
    }

    function approveContent(address content_contract, bool approved, string note) public returns ( bool ) {
        require(canReview(tx.origin) == true);
        // credit the account based on the percent_complete
        uint256 index = approvalRequestsMap[content_contract] - 1;
        BaseContent c = BaseContent(content_contract);

        // remove the request from the list
        delete approvalRequests[index];
        approvalRequestsLength--;
        approvalRequestsMap[content_contract] = 0;
        if (approvalRequestsLength > index) {
            address lastRequest = approvalRequests[approvalRequestsLength];
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

    function validType(address content_type) public view returns (bool) {
        bool isValidType = false;
        for (uint i = 0; i < contentTypesLength; i++) {
            if (contentTypes[i] == content_type) {
                isValidType = true;
            }
        }
        return isValidType;
    }

    function createContent(address content_type) public  returns (address) {
        require(canContribute(tx.origin)); //check if sender has contributor access
        if (contentTypesLength != 0) {
            require(validType(content_type));
        }
        BaseContent content = new BaseContent(this, content_type);
        content.setAddressKMS(addressKMS);
        content.setContentContractAddress(contentTypeContracts[content_type]);
        emit ContentObjectCreated(address(content), content_type);
        return address(content);
    }

    function accessRequest() public returns (bool) {
        require(hasAccess(tx.origin) || canContribute(tx.origin) || canReview(tx.origin));
        emit AccessRequest();
        return true;
    }
}

