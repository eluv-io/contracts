// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract ERC20Payments {
    struct PaymentID {
        bytes16 ref_id;
        address oracle_id;
    }

    enum PaymentState {
        Created,
        Withdrawn,
        Canceled
    }
    // Can be optimized by storing bytes instead of variable length arrays
    struct LockPayment {
        address sender;
        address[] receivers;
        address tokenContract;
        uint256[] amounts;
        address oracle;
        PaymentState state;
    }

    // Maybe this can be simplified to just be the contractId
    event Created(
        bytes16 indexed contractId,
        address indexed sender,
        address[] indexed receivers,
        address tokenContract,
        uint256[] amounts
    );

    event Withdrawn(bytes16 indexed contractId);
    event Canceled(bytes16 indexed contractId);

    modifier tokensTransferable(address _token, uint256 _amount) {
        require(_amount > 0, "token amount must be > 0");
        require(
            ERC20(_token).allowance(msg.sender, address(this)) >= _amount,
            "allowance must be >= amount"
        );
        _;
    }

    modifier contractExists(bytes16 _paymentID) {
        require(havePayment(_paymentID), "paymentID does not exist");
        _;
    }

    modifier paymentNotExists(bytes16 _paymentID) {
        require(!havePayment(_paymentID), "paymentID already exists");
        _;
    }

    modifier withdrawable(bytes16 _paymentID) {
        // This can be simplified, but for error clarity, we keep it as is
        require(paymentTransactions[_paymentID].state != PaymentState.Withdrawn, "withdrawable: already withdrawn");
        require(paymentTransactions[_paymentID].state != PaymentState.Canceled, "withdrawable: already refunded");
        _;
    }

    modifier refundable(bytes16 _paymentID) {
        require(paymentTransactions[_paymentID].oracle == msg.sender, "refundable: not oracle");
        // This can be simplified, but for error clarity, we keep it as is
        require(paymentTransactions[_paymentID].state != PaymentState.Canceled, "refundable: already refunded");
        require(paymentTransactions[_paymentID].state != PaymentState.Withdrawn, "refundable: already withdrawn");
        _;
    }

    modifier onlyOracle(bytes16 _paymentID) {
        require(paymentTransactions[_paymentID].oracle == msg.sender, "onlyOracle: message sender must be oracle");
        _;
    }

    mapping (bytes16 => LockPayment) paymentTransactions; // Change to paymentTransactions

    /**
     * @dev Create a new payment contract between the sender and the reciever arrays
     *
     * NOTE: _receiver must first call approve() on the token contract.
     *       See allowance check in tokensTransferable modifier.
     * @param _receivers Receivers of the tokens.
     * @param _paymentId PaymentID of the payment.
     * @param _tokenContract ERC20 Token contract address.
     * @param _amounts Amounts of the token to lock up.
     * @return refId Id of the payement.
     */
    function createPayment(
        address[] calldata _receivers,
        PaymentID calldata _paymentId,
        address _tokenContract,
        uint256[] calldata _amounts
    )
    external
    paymentNotExists(_paymentId.ref_id)
    returns (bytes16 refId)
    {
        refId = _paymentId.ref_id;

        uint256 total = verifyAmounts(_amounts, _tokenContract);
        // Reject if a contract already exists with the same ID.
        if (havePayment(refId))
            revert("Contract already exists");

        // This contract becomes the temporary owner of the tokens
        if (!ERC20(_tokenContract).transferFrom(msg.sender, address(this), total))
            revert("transferFrom sender to this failed");

        paymentTransactions[refId] =  LockPayment(
            msg.sender,
            _receivers,
            _tokenContract,
            _amounts,
            _paymentId.oracle_id,
            PaymentState.Created
        );

        emit Created(refId, msg.sender, _receivers, _tokenContract, _amounts);
    }

    /**
    * @dev Called by the oracle after the NFT has been transfered to the initiater's account
    *
    * @param _refId reference ID of the contract
    * @return bool true on success
     */
    function claimPayment(bytes16 _refId)
    external
    contractExists(_refId)
    withdrawable(_refId)
    onlyOracle(_refId)
    returns (bool)
    {
        LockPayment storage c = paymentTransactions[_refId];
        c.state = PaymentState.Withdrawn;
        // Iterate over all receivers and transfer the tokens to them
        for (uint i = 0; i < c.receivers.length; i++) {
            if (!ERC20(c.tokenContract).transfer(c.receivers[i], c.amounts[i]))
                revert("transfer failed");
        }
        emit Withdrawn(_refId);
        return true;
    }

    /**
     * @dev Called by the oracle if there was no withdraw to refund the payment sender.
     *
     * @param _refId reference ID of the contract
     * @return bool true on success
     */
    function cancelPayment(bytes16 _refId)
    external
    contractExists(_refId)
    refundable(_refId)
    returns (bool)
    {
        LockPayment storage c = paymentTransactions[_refId];
        c.state = PaymentState.Canceled;
        uint256 total = calculateTotal(c.amounts);
        ERC20(c.tokenContract).transfer(c.sender, total);
        emit Canceled(_refId);
        return true;
    }

    /**
     * @dev Get contract details.
     * @param _refId reference ID of the contract
     */
    function getContract(bytes16 _refId)
    public
    view
    returns (
        address sender,
        address[] memory receivers,
        address tokenContract,
        uint256[] memory amounts,
        address oracle,
        PaymentState state
    )
    {
        if (havePayment(_refId) == false)
            revert("Contract does not exist");
        LockPayment storage c = paymentTransactions[_refId];
        return (
        c.sender,
        c.receivers,
        c.tokenContract,
        c.amounts,
        c.oracle,
        c.state
        );
    }

    /**
     * @dev Is there a contract with id _refId.
     * @param _refId Id into paymentTransactions mapping.
     */
    function havePayment(bytes16 _refId)
    internal
    view
    returns (bool exists)
    {
        exists = (paymentTransactions[_refId].sender != address(0));
    }

    /**
     * @dev Calculate the total amount of the tokens and confirm that the sender has allocated the correct
     * number of tokens to the contract.
     * @param _amounts Amounts of the tokens.
     * @return total Total amount of the tokens.
     */
    function verifyAmounts(uint256[] memory _amounts, address _tokenContract)
    internal
    view
    returns (uint256 total)
    {
        require(_amounts.length > 0, "no amounts provided");
        total = 0;
        for (uint i=0; i < _amounts.length; i++) {
            require(_amounts[i] > 0, "amount must be greater than 0");
            total += _amounts[i];
        }
        require(
            ERC20(_tokenContract).allowance(msg.sender, address(this)) >= total,
            "allowance must be >= sum(amounts)"
        );
    }

    /**
    * @dev Convenience method to calculate the total amount of the tokens.
    * @param _amounts Amounts of the tokens.
    * @return total Total amount of the tokens.
    */

    function calculateTotal(uint256[] memory _amounts) public pure returns (uint256 total)
    {
        total = 0;
        for (uint i=0; i < _amounts.length; i++) {
            total += _amounts[i];
        }
    }
}