// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ERC20Payments {
    struct Payment {
        bytes16 paymentId;
        address oracleId;
    }

    enum PaymentState {
        Created,
        Claimed,
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

    event Created(bytes16 indexed paymentId);
    event Claimed(bytes16 indexed paymentId);
    event Canceled(bytes16 indexed paymentId);

    modifier tokensTransferable(address _token, uint256 _amount) {
        require(_amount > 0, "token amount must be > 0");
        require(IERC20(_token).allowance(msg.sender, address(this)) >= _amount, "allowance must be >= amount");
        _;
    }

    modifier paymentExists(bytes16 _paymentId) {
        require(havePayment(_paymentId), "paymentId does not exist");
        _;
    }

    modifier paymentNotExists(bytes16 _paymentId) {
        require(!havePayment(_paymentId), "paymentId already exists");
        _;
    }

    modifier claimable(bytes16 _paymentId) {
        // This can be simplified, but for error clarity, we keep it as is
        require(paymentTransactions[_paymentId].state != PaymentState.Claimed, "claimable: already claimed");
        require(paymentTransactions[_paymentId].state != PaymentState.Canceled, "claimable: already refunded");
        _;
    }

    modifier refundable(bytes16 _paymentId) {
        require(paymentTransactions[_paymentId].oracle == msg.sender, "refundable: not oracle");
        // This can be simplified, but for error clarity, we keep it as is
        require(paymentTransactions[_paymentId].state != PaymentState.Canceled, "refundable: already refunded");
        require(paymentTransactions[_paymentId].state != PaymentState.Claimed, "refundable: already claimed");
        _;
    }

    modifier onlyOracle(bytes16 _paymentId) {
        require(paymentTransactions[_paymentId].oracle == msg.sender, "onlyOracle: message sender must be oracle");
        _;
    }

    mapping(bytes16 => LockPayment) paymentTransactions; // Change to paymentTransactions

    /**
     * @dev Create a new payment contract between the sender and the receiver arrays
     *
     * NOTE: _receiver must first call approve() on the token contract.
     * See allowance check in tokensTransferable modifier.
     * @param _receivers Receivers of the tokens.
     * @param _payment payment struct.
     * @param _tokenContract ERC20 Token contract address.
     * @param _amounts Amounts of the token to lock up.
     * @return paymentId payment id of the contract.
     */
    function createPayment(
        address[] calldata _receivers,
        Payment calldata _payment,
        address _tokenContract,
        uint256[] calldata _amounts
    )
        external
        paymentNotExists(_payment.paymentId)
        returns (bytes16 paymentId)
    {
        paymentId = _payment.paymentId;

        uint256 total = verifyAmounts(_amounts, _tokenContract);
        // Reject if a contract already exists with the same ID.
        if (havePayment(paymentId)) {
            revert("paymentId already exists");
        }

        // This contract becomes the temporary owner of the tokens
        if (!IERC20(_tokenContract).transferFrom(msg.sender, address(this), total)) {
            revert("transferFrom sender to this failed");
        }

        paymentTransactions[paymentId] =
            LockPayment(msg.sender, _receivers, _tokenContract, _amounts, _payment.oracleId, PaymentState.Created);

        emit Created(paymentId);
    }

    /**
     * @dev Called by the oracle after the NFT has been transferred to the initiator's account
     *
     * @param _paymentId payment id of the contract
     * @return bool true on success
     */
    function claimPayment(bytes16 _paymentId)
        external
        paymentExists(_paymentId)
        claimable(_paymentId)
        onlyOracle(_paymentId)
        returns (bool)
    {
        LockPayment storage c = paymentTransactions[_paymentId];

        // Iterate over all receivers and transfer the tokens to them
        for (uint256 i = 0; i < c.receivers.length; i++) {
            if (!IERC20(c.tokenContract).transfer(c.receivers[i], c.amounts[i])) {
                revert("transfer failed");
            }
        }

        c.state = PaymentState.Claimed;
        emit Claimed(_paymentId);
        return true;
    }

    /**
     * @dev Called by the oracle if there was no claim to refund the payment sender.
     *
     * @param _paymentId payment id of the contract
     * @return bool true on success
     */
    function cancelPayment(bytes16 _paymentId)
        external
        paymentExists(_paymentId)
        refundable(_paymentId)
        returns (bool)
    {
        LockPayment storage c = paymentTransactions[_paymentId];
        c.state = PaymentState.Canceled;
        uint256 total = calculateTotal(c.amounts);
        IERC20(c.tokenContract).transfer(c.sender, total);
        emit Canceled(_paymentId);
        return true;
    }

    /**
     * @dev Get contract details.
     * @param _paymentId payment id of the contract
     */
    function getPayment(bytes16 _paymentId)
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
        if (havePayment(_paymentId) == false) {
            revert("contract does not exist");
        }
        LockPayment storage c = paymentTransactions[_paymentId];
        return (c.sender, c.receivers, c.tokenContract, c.amounts, c.oracle, c.state);
    }

    /**
     * @dev Is there a contract with id _paymentId.
     * @param _paymentId payment id in paymentTransactions mapping.
     */
    function havePayment(bytes16 _paymentId) internal view returns (bool exists) {
        exists = (paymentTransactions[_paymentId].sender != address(0));
    }

    /**
     * @dev Calculate the total amount of the tokens and confirm that the sender has allocated the correct
     * number of tokens to the contract.
     * @param _amounts Amounts of the tokens.
     * @return total Total amount of the tokens.
     */
    function verifyAmounts(uint256[] memory _amounts, address _tokenContract) internal view returns (uint256 total) {
        require(_amounts.length > 0, "no amounts provided");
        total = 0;
        for (uint256 i = 0; i < _amounts.length; i++) {
            require(_amounts[i] > 0, "amount must be greater than 0");
            total += _amounts[i];
        }
        require(
            IERC20(_tokenContract).allowance(msg.sender, address(this)) >= total, "allowance must be >= sum(amounts)"
        );
    }

    /**
     * @dev Convenience method to calculate the total amount of the tokens.
     * @param _amounts Amounts of the tokens.
     * @return total Total amount of the tokens.
     */

    function calculateTotal(uint256[] memory _amounts) public pure returns (uint256 total) {
        total = 0;
        for (uint256 i = 0; i < _amounts.length; i++) {
            total += _amounts[i];
        }
    }
}
