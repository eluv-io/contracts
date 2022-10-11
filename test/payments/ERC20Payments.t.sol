// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {console} from "forge-std/console.sol";
import {Utilities} from "../utils/Utilities.sol";
import {stdStorage, StdStorage, Test} from "forge-std/Test.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {ERC20Payments} from "../../src/payments/ERC20Payments.sol";
import "src/tradable/ElvToken.sol";

contract ERC20PaymentsTest is Test {
    Utilities internal utils;
    address payable[] internal users;
    ElvToken internal token;
    ERC20Payments internal erc20Payments;

    address internal alice;
    address internal bob;
    address internal carol;

    function setUp() public virtual {
        utils = new Utilities();
        users = utils.createUsers(3);

        alice = users[0];
        vm.label(alice, "Alice");
        bob = users[1];
        vm.label(bob, "Bob");
        carol = users[2];
        vm.label(carol, "Carol");

        token = new ElvToken("ElvToken","ELV",20000000);
        erc20Payments = new ERC20Payments();
    }

    function aliceCreateERC20Payments(
        bytes16 paymentId,
        address[] memory receivers,
        uint256[] memory amounts,
        address oracle,
        uint256 initBal
    )
        private
    {
        token.transfer(alice, initBal * 10 ** token.decimals());
        uint256 total = erc20Payments.calculateTotal(amounts);
        vm.prank(alice);
        token.approve(address(erc20Payments), total);
        vm.prank(alice);
        erc20Payments.createPayment(receivers, ERC20Payments.Payment(paymentId, oracle), address(token), amounts);
    }

    function preloadedInit() private returns (bytes16 paymentId) {
        paymentId = "stringLiteral";
        address[] memory receivers = new address[](1);
        uint256[] memory amounts = new uint256[](1);
        receivers[0] = bob;
        amounts[0] = 1000 * 10 ** token.decimals();
        uint256 initBal = 10000;
        aliceCreateERC20Payments(paymentId, receivers, amounts, carol, initBal);
    }

    function testCreateERC20PaymentsState() public {
        console.log("Testing Correct ERC20Payments Creation");
        token.transfer(alice, 10000 * 10 ** 18);
        uint256 initialBalance = token.balanceOf(alice);
        bytes16 paymentId = "stringLiteral";
        address[] memory receivers = new address[](1);
        uint256[] memory amounts = new uint256[](1);
        receivers[0] = bob;
        amounts[0] = 1000 * 10 ** token.decimals();
        vm.prank(alice);
        token.approve(address(erc20Payments), amounts[0]);
        vm.prank(alice);
        erc20Payments.createPayment(receivers, ERC20Payments.Payment(paymentId, alice), address(token), amounts);
        (address s, address[] memory r, address tc, uint256[] memory a, address o, ERC20Payments.PaymentState st) =
            erc20Payments.getPayment(paymentId);
        require(
            initialBalance - token.balanceOf(alice) == 1000 * 10 ** token.decimals(),
            "Balance should be decreased by 1000"
        );
        require(token.balanceOf(alice) == 9000 * 10 ** token.decimals(), "Alice should have 9k");
        require(s == alice, "Sender is wrong");
        require(r[0] == bob, "Receiver is not Bob");
        require(tc == address(token), "Token contract is not token");
        require(a[0] == 1000 * 10 ** token.decimals(), "Amount is not 1000");
        require(o == alice, "Oracle is not alice");
        require(st == ERC20Payments.PaymentState.Created, "State is not Created");
    }

    function testCancelERC20Payments() public {
        console.log("Testing Correct ERC20Payments Cancellation");
        bytes16 paymentId = "stringLiteral";
        address[] memory receivers = new address[](1);
        uint256[] memory amounts = new uint256[](1);
        receivers[0] = bob;
        amounts[0] = 1000 * 10 ** token.decimals();
        uint256 init_bal = 10000;
        aliceCreateERC20Payments(paymentId, receivers, amounts, carol, init_bal);
        vm.prank(carol);
        erc20Payments.cancelPayment(paymentId);
        (address s, address[] memory r, address tc, uint256[] memory a, address o, ERC20Payments.PaymentState st) =
            erc20Payments.getPayment(paymentId);
        require(token.balanceOf(alice) == 10000 * 10 ** token.decimals(), "Alice should have 10k");
        require((token.balanceOf(address(erc20Payments)) == 0), "ERC20Payments should have 0 balance");
        require(s == alice, "Sender is wrong");
        require(r[0] == bob, "Receiver is not Bob");
        require(tc == address(token), "Token contract is not token");
        require(a[0] == 1000 * 10 ** token.decimals(), "Amount is not 1000");
        require(o == carol, "Oracle is not carol");
        require(st == ERC20Payments.PaymentState.Canceled, "State is not Canceled");
    }

    function testClaimERC20Payments() public {
        console.log("Testing Correct ERC20Payments Claimable");
        bytes16 paymentId = preloadedInit();
        vm.startPrank(carol);
        erc20Payments.claimPayment(paymentId);
        vm.stopPrank();
        (address s, address[] memory r, address tc, uint256[] memory a, address o, ERC20Payments.PaymentState st) =
            erc20Payments.getPayment(paymentId);
        require(token.balanceOf(alice) == 9000 * 10 ** token.decimals(), "Alice should have 9k");
        require(token.balanceOf(bob) == 1000 * 10 ** token.decimals(), "Bob should have 1k");
        require(token.balanceOf(carol) == 0, "Carol should have 0");
        require(token.balanceOf(address(erc20Payments)) == 0, "ERC20Payments should have 0 balance");
        require(s == alice, "Sender is wrong");
        require(r[0] == bob, "Receiver is not Bob");
        require(tc == address(token), "Token contract is not token");
        require(a[0] == 1000 * 10 ** token.decimals(), "Amount is not 1000");
        require(o == carol, "Oracle is not carol");
        require(st == ERC20Payments.PaymentState.Claimed, "State is not claimed");
    }

    function testFailRepeatedCreate() public {
        console.log("Testing Repeated ERC20Payments Creation");
        bytes16 paymentId = "stringLiteral";
        address[] memory receivers = new address[](1);
        uint256[] memory amounts = new uint256[](1);
        receivers[0] = bob;
        amounts[0] = 1000 * 10 ** token.decimals();
        uint256 init_bal = 10000;
        aliceCreateERC20Payments(paymentId, receivers, amounts, carol, init_bal);
        aliceCreateERC20Payments(paymentId, receivers, amounts, carol, init_bal);
    }

    function testFailClaim() public {
        console.log("Testing ERC20Payments Claim");
        bytes16 paymentId = preloadedInit();
        vm.prank(alice);
        erc20Payments.claimPayment(paymentId);
    }

    function testMultipleReceivers(uint256 len) public {
        //len = 1;
        vm.assume(len > 0 && len < 500);
        uint256[] memory amounts = new uint256[](len);
        for (uint256 i = 0; i < len; i++) {
            amounts[i] = (i + 1) * 100 * 10 ** token.decimals();
        }
        uint256 initBal = erc20Payments.calculateTotal(amounts);
        token.transfer(alice, initBal);

        console.log("Testing Multiple receivers");
        bytes16 paymentId = "stringLiteral";
        address[] memory receivers = new address[](amounts.length);
        address payable[] memory payables = utils.createUsers(amounts.length);
        for (uint256 i = 0; i < len; i++) {
            receivers[i] = payables[i];
        }

        vm.startPrank(alice);
        token.approve(address(erc20Payments), initBal);
        erc20Payments.createPayment(receivers, ERC20Payments.Payment(paymentId, carol), address(token), amounts);
        vm.stopPrank();
        vm.startPrank(carol);
        erc20Payments.claimPayment(paymentId);
        vm.stopPrank();
        (address s, address[] memory r, address tc,, address o, ERC20Payments.PaymentState st) =
            erc20Payments.getPayment(paymentId);
        require(token.balanceOf(alice) == 0, "Alice should have 0");
        for (uint256 i = 0; i < amounts.length; i++) {
            require(
                token.balanceOf(payables[i]) == amounts[i],
                string(abi.encodePacked("Recipient ", i, " should have ", amounts[i]))
            );
        }
        require(token.balanceOf(address(erc20Payments)) == 0, "ERC20Payments should have 0 balance");
        require(s == alice, "Sender is wrong");
        require(r[0] == payables[0], "Receiver is not Bob");
        require(tc == address(token), "Token contract is not token");
        require(o == carol, "Oracle is not carol");
        require(st == ERC20Payments.PaymentState.Claimed, "State is not Confirmed");
    }

    function testFailBobCancels() public {
        console.log("Testing Bob Cancelling");
        bytes16 paymentId = preloadedInit();
        vm.prank(bob);
        erc20Payments.cancelPayment(paymentId);
    }
}
