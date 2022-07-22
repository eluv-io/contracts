pragma solidity ^0.8.13;

contract Claimer {

    /*
        An Allocation contains always an amount and an expiration date.
    */
    struct Allocation {
        uint amount;
        uint expirationDate;
    }

    address public immutable owner;
    mapping (address => bool) public allocationAuthorities;
    mapping (address => Allocation[]) public allocations;
    mapping (address => uint) public claims;
    mapping (address => uint) public burns;

    event ClaimerAllocate(address _addr, uint _amount);
    event ClaimerClaim(uint _amount);
    event ClaimerBurn(uint _amount);


    /*
        We certify that the owner is the creator of the contract 
    */
    constructor () {
       owner = msg.sender;
       addAuthorizedAdr(msg.sender);
    }


    modifier ownerOnly {
        require(owner == msg.sender, "Only Owner is authorized.");
        _;
    }

    modifier authorizedAdr(address _addr) {
        require(allocationAuthorities[_addr] == true, "This address is not authorized.");
        _;
    }

    /*
        We allocate an allocation to a user
        @param _addr : id of the user
        @param _amount : amount given
        @param _expirationDate : expiration date of the allocation
    */
    function allocate(address _addr, uint _amount, uint _expirationDate) external ownerOnly authorizedAdr(_addr) {
        require(block.timestamp < _expirationDate, "The expiration date is not valid.");

        Allocation memory currentAllocation;
        currentAllocation.amount = _amount;
        currentAllocation.expirationDate = _expirationDate;

        allocations[_addr].push(currentAllocation);
        emit ClaimerAllocate(_addr, _amount);
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
    function rmAuthorizedAdr(address _addr) external ownerOnly authorizedAdr(_addr) {
        allocationAuthorities[_addr] = false;
    }

    /*
        User claims an amount and we verify that we can give it to them.
        We give the maximum to the user: for example if a user wants 200, but only 100 is allocated, we give him the 100.
        @param _amount : the amount that claims the user
    */
    function claim (uint _amount) external authorizedAdr(msg.sender){
        claimAndBurn(_amount, true);
    }

    /*
        A user can burn his units.
        As for claim, we burn the maximum possible.
        @param _amount : the amount that burns the user
    */
    function burn (uint _amount) external authorizedAdr(msg.sender){
        claimAndBurn(_amount, false);
    }

    /*
        Function that can do claiming or burning by using a boolean specifier
        @param _amount : the amount that claims/burns the user
        @param isClaiming : the boolean that specify the behavior of the function
    */
    function claimAndBurn(uint _amount, bool isClaiming) private {
        uint tot = 0;
        uint i = 0;
        bool increment = true;
        while(i < allocations[msg.sender].length){
            
            Allocation memory currentAllocation = allocations[msg.sender][i];
            if(currentAllocation.expirationDate > block.timestamp){ //we verify that the allocation is not expired
                if(currentAllocation.amount + tot >= _amount){
                    allocations[msg.sender][i].amount -= _amount - tot;
                    if(isClaiming){
                        claims[msg.sender] += _amount;
                        emit ClaimerClaim(_amount);
                    } else {
                        burns[msg.sender] += _amount;
                        emit ClaimerBurn(_amount);
                    }
                    return;
                } else {
                    tot += currentAllocation.amount;
                    allocations[msg.sender][i].amount = 0;
                }
            } else { //if expired, we remove it from the array
                uint lastIdx = allocations[msg.sender].length - 1;
                if(i != lastIdx){
                    allocations[msg.sender][i] = allocations[msg.sender][lastIdx];
                    increment = false; //we do not increment bc we want to verify the current i
                }
                allocations[msg.sender].pop();
            }

            if(increment) i++;  
            else increment = true;
            
        }

        if(isClaiming){
            claims[msg.sender] += tot;
            emit ClaimerClaim(tot);
        } else {
            burns[msg.sender] += tot;
            emit ClaimerBurn(tot);
        }

    }
    
  

    
    //============UTILS FUNCTIONS FOR TESTING=============
    function getAmount(address _addr, uint idx) external view returns(uint){
        return allocations[_addr][idx].amount;
    } 



}

