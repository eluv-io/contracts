pragma solidity ^0.5.2;

import "./openzeppelin-solidity/SafeMath.sol";
// import "zeppelin-solidity/contracts/ECRecovery.sol";

// abigen --sol channel/SimpleStateChannel.sol --pkg=contracts --out build/state_channel.go

/**
 * @title SimplePaymentChannel
 * @author Eric Olszewski <eolszewski@gmail.com>
 *
 * @dev Ethereum payment channels allow for off-chain transactions with an on-chain
 * settlement. In this implementation, a party (sender) can open a channel with a
 * deposit, expiration, and recipient. The sender can then sign transactions off-chain
 * and send them to the recipient, who can submit one of these signed transactions to
 * chain to close and settle the channel.
 */
contract SimplePaymentChannel {
    using SafeMath for uint256;
    // using ECRecovery for bytes32;

    //============================================================================
    // EVENTS
    //============================================================================

    event ChannelOpened(address sender, address recipient, uint expiration, uint256 deposit);
    event ChannelClosed(uint256 senderBalance, uint256 recipientBalance);

    // Creator of channel
    address payable public sender;
    // Recipient of channel
    address payable public recipient;
    // When sender can close the channel
    uint256 public expiration;

    //============================================================================
    // PUBLIC FUNCTIONS
    //============================================================================

    /**
    * @dev Constructor
    * @param _recipient address Address of the receiving party
    * @param _duration uint256 Time period past creation that defines expiration
    */
    constructor(address payable _recipient, uint256 _duration) public payable {
        // sanity checks
        require(msg.value > 0);
        require(msg.sender != _recipient);

        sender = msg.sender;
        recipient = _recipient;
        expiration = now + _duration;

        emit ChannelOpened(sender, recipient, expiration, msg.value);
    }

    /**
    * @dev Close a channel with a signed message from the sender
    * @param _balance uint256 The balance agreed to by the sender in their signed message
    */
    function closeChannel(uint256 _balance, uint256 _nonce,
        uint8 _v, bytes32 _r, bytes32 _s) public {
        // only recipient can call closeChannel
        require(msg.sender == recipient);
        require(isValidSignature(_balance, _nonce, _v, _r, _s));

        // temp vars for calculating sender and recipient remittances
        uint256 balance = _balance;
        uint256 remainder = 0;

        // if _balance > address(this).balance, send address(this).balance to recipient
        if (_balance > address(this).balance) {
            balance = address(this).balance;
        } else {
            remainder = address(this).balance.sub(balance);
        }

        // remit payment to recipient
        recipient.transfer(balance);

        emit ChannelClosed(remainder, balance);

        // remit remainder to sender
        selfdestruct(sender);
    }

    /**
    * @dev Extend the expiration date of the channel
    * @param _expiration uint256 Updated expiration date
    */
    function extendExpiration(uint256 _expiration) public {
        // only sender can call extendExpiration
        require(msg.sender == sender);
        require(_expiration > expiration);

        // update expiration
        expiration = _expiration;
    }

    /**
    * @dev Allows sender to claim channel balance if expired
    */
    function claimTimeout() public {
        // must be beyond expiration date
        require(now >= expiration);

        // remit payment to sender
        selfdestruct(sender);
    }

    //============================================================================
    // INTERNAL FUNCTIONS
    //============================================================================

    /*
        function recoverAddr(bytes32 msgHash, uint8 v, bytes32 r, bytes32 s) public pure returns (address) {
        return ecrecover(msgHash, v, r, s);
    }

    /**
    * @dev Derive and verify sender address from signed message and message hash
    * @param _balance uint256 The balance agreed to by the sender in their signed message
    */
    function isValidSignature(uint256 _balance, uint256 _nonce, uint8 _v, bytes32 _r, bytes32 _s)
    public
    // internal - make visible for testing
    // view
    returns (bool)
    {
        // signature is on [contract address | balance | nonce]
        // bytes32 checkHash = prefixedHash(keccak256(abi.encodePacked(address(this), _balance, _nonce)));
        bytes32 checkHash = keccak256(abi.encodePacked(address(this), _balance, _nonce));

        // check that the signature is from the payment sender
        address checkAddr = ecrecover(checkHash, _v, _r, _s);
        emit CheckEcRecover(checkAddr);
        return checkAddr == sender;
    }

    function splitSignature(bytes memory _sig) pure public returns (uint8, bytes32, bytes32) {
        require(_sig.length == 65);

        bytes32 r;
        bytes32 s;
        uint8 v;

        assembly {
            // first 32 bytes, after the length prefix
            r := mload(add(_sig, 32))
            // second 32 bytes
            s := mload(add(_sig, 64))
            // final byte (first byte of the next 32 bytes)
            v := byte(0, mload(add(_sig, 96)))
        }
        return (v, r, s);
    }

    event CheckEcRecover(address checkAddr);
    function checkEcrecover(bytes32 _hash, bytes memory _sig)
    public
    // view
    returns (address)
    {
        uint8 v;
        bytes32 r;
        bytes32 s;

        (v, r, s) = splitSignature(_sig);
        address checkAddr = ecrecover(_hash, v, r, s);
        emit CheckEcRecover(checkAddr);
        return checkAddr;
    }

    function checkVRS(uint8 _v, bytes32 _r, bytes32 _s)
    public
    view
    returns (bytes memory)
    {
        return abi.encodePacked(_v, _r, _s);
    }

    function checkPackedAndHashed(uint256 _balance, uint256 _nonce)
    public
    view
    returns (bytes32)
    {
        return keccak256(abi.encodePacked(address(this), _balance, _nonce));
    }

    /**
    * @dev Builds a prefixed hash to mimic the behavior of eth_sign
    * @param _hash bytes32 Message hash to be prefixed
    */
    function prefixedHash(bytes32 _hash)
    internal
    pure
    returns (bytes32)
    {
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", _hash));
    }
}