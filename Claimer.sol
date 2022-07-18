pragma solidity ^0.8.13;

contract Claimer {

    address owner;
    mapping (address => bool) public allocationAuthorities;

    /*
        An Allocation contains always an amount and an expiration date.
    */
    struct Allocation {
        uint amount;
        uint expirationDate;
    }

    /*
        We certify that the owner is the creator of the contract 
    */
    constructor ()  {
       owner = msg.sender;
    }

    mapping (address => Allocation[]) public allocations;
    mapping (address => uint) public claims;

    event AllocationUnlock();
    event AllocationClaim();

    modifier ownerOnly {
        require(owner == msg.sender, "Only Owner is authorized.");
        _;
    }

    modifier authorizedAdr(address _addr) {
        require(allocationAuthorities[_addr], "This address is not authorized.");
        _;
    }

    /*
        We allocate an allocation to a user
        @param _addr : id of the user
        @param _amount : amount given
        @param _expirationDate : expiration date of the allocation
    */
    function allocate(address _addr, uint _amount, uint _expirationDate) public ownerOnly authorizedAdr(_addr) {
        require(block.timestamp < _expirationDate, "The expiration date is not valid.");

        Allocation memory currentAllocation;
        currentAllocation.amount = _amount;
        currentAllocation.expirationDate = _expirationDate;

        allocations[_addr].push(currentAllocation);
    }

    /*
        @param _addr : id of the user we want to add to our authorized address
    */
    function addAuthorizedAdr(address _addr) public ownerOnly {
        allocationAuthorities[_addr] = true;
    }

    /*
        @param _addr : id of the user we want to remove to our authorized address
    */
    function rmAuthorizedAdr(address _addr) public ownerOnly authorizedAdr(_addr) {
        allocationAuthorities[_addr] = false;
    }

    /*
        User claims an amount and we verify that we can give it to them.
        This is to avoid giving users things they don't want.
        @param _amount : the amount that claims the user
    */
    function claim (uint _amount) public authorizedAdr(msg.sender){
        for(uint i = 0; i < allocations[msg.sender].length;  i++) {
            Allocation memory currentAllocation = allocations[msg.sender][i];
            if(currentAllocation.amount >= _amount && currentAllocation.expirationDate > block.timestamp){
                currentAllocation.amount -= _amount;
                claims[msg.sender] += _amount;
                break;
            }
        }
    }

    /*
        A user can burn his units
        @param _amount : the amount that burns the user
    */
    function burn (uint _amount) public authorizedAdr(msg.sender){
        for(uint i = 0; i < allocations[msg.sender].length;  i++) {
            Allocation memory currentAllocation = allocations[msg.sender][i];
            if(currentAllocation.amount >= _amount && currentAllocation.expirationDate > block.timestamp){
                currentAllocation.amount -= _amount;
                break;
            }
        }
    }


}

