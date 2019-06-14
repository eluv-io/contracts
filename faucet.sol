pragma solidity 0.4.24;

import {Ownable} from "./ownable.sol";

contract Faucet is Ownable {

    event Withdrawal(address to, uint amount);
    event OneKTokenSent(address receiver);
    event TwoKTokenSent(address receiver);
    event FiveKTokenSent(address receiver);
    event FaucetOn(bool status);
    event FaucetOff(bool status);

    uint256 constant oneKToken = 1000000000000000000000;
    uint256 constant twoKToken = 2000000000000000000000;
    uint256 constant fiveKToken = 5000000000000000000000;
    uint256 constant oneHours = 1 minutes;
    uint256 constant twoHours = 2 minutes;
    uint256 constant fiveHours = 5 minutes;
    
    string public faucetName;
    bool public faucetStatus;
    mapping(address => uint256) status;

    modifier faucetOn() {
        require(faucetStatus);
        _;
    }

    modifier faucetOff() {
        require(!faucetStatus);
        _;
    }

    constructor(string _faucetName) public payable{
        faucetName = _faucetName;
        faucetStatus = true;

        emit FaucetOn(faucetStatus);
    }

    // withdrawal - sends "_withdraw_amount" wei to caller of this function
    // checks if the faucet has sufficient balance.
    function withdrawal(address _to, uint _withdraw_amount) internal {
        require(address(this).balance >= _withdraw_amount,
            "Insufficient balance in faucet for withdrawal request");
        _to.transfer(_withdraw_amount);
        emit Withdrawal(msg.sender, _withdraw_amount);
    }

// TODO as of now any one can give the _to address, should it be restricted to onlyOwner??

    // drip1000Token - sends 1000 ether to the caller 
    // with time lock of 1 hour
    function drip1000Token(address _to) public faucetOn() {
        require(checkStatus(_to));
        withdrawal(_to, oneKToken);
        updateStatus(_to, oneHours);

        emit OneKTokenSent(_to);
    }

    // drip2000Token - sends 2000 ether to the caller 
    // with time lock of 2 hours
    function drip2000Token(address _to) public faucetOn(){
        require(checkStatus(_to));
        withdrawal(_to, twoKToken);
        updateStatus(_to, twoHours);

        emit TwoKTokenSent(_to);
    }

    // drip5000Token - sends 5000 ether to the caller 
    // with time lock of 5 hours
    function drip5000Token(address _to) public faucetOn(){
        require(checkStatus(_to));
        withdrawal(_to, fiveKToken);
        updateStatus(_to, fiveHours);

        emit FiveKTokenSent(_to);
    }

    // get balance of the faucet
    function getBalance() public view returns (uint256) {
        return address(this).balance;
    }

    // turn faucet on
    function turnFaucetOn() public onlyOwner faucetOff() {
        faucetStatus = true;
        emit FaucetOn(faucetStatus);
    }

    // turn faucet off
    function turnFaucetOff() public onlyOwner faucetOn() {
        faucetStatus = false;
        emit FaucetOff(faucetStatus);
    }

    // Internal functions
    // checkStatus - verifies the time lock duration for the address provided
    function checkStatus(address _address) internal view returns (bool) {
        //check if first time address is requesting
        if(status[_address] == 0) {
            return true;
        }
        //if not first time check the timeLock
        else {
            if(block.timestamp >= status[_address]) {
                return true;
            }
            else {
                return false;
            }
        }
    }

    // Internal functions
    // updateStatus - updates the time lock duration for the address provided
    function updateStatus(address _address, uint256 _timelock) internal {
        status[_address] = block.timestamp + _timelock;
    }

}
