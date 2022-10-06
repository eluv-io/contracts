// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {console} from "forge-std/console.sol";
import {Utilities} from "../utils/Utilities.sol";
import {stdStorage, StdStorage, Test} from "forge-std/Test.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {ERC20Payments} from "payments/ERC20Payments.sol";

// Deployed at 0x60aa8e13feb9301ec17cdf07ff6a45b9cd472279
contract TestToken is ERC20 {
    address public admin;

    constructor() ERC20("dust", "TestToken") {
        _mint(msg.sender, 20000000 * 10 ** decimals());
        admin = msg.sender;
    }

    function name() override public pure returns (string memory) {
        return "Test Coin";
    }
}

contract ERC20PaymentsTest is Test {

    Utilities internal utils;
    address payable[] internal users;
    ERC20 internal token;
    ERC20Payments internal escrow;


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
        token = new TestToken();
        escrow = new ERC20Payments();
    }

    function aliceCreateEscrow(bytes16 refId, address[] memory recievers, uint256[] memory amounts, address oracle, uint init_bal) private {
        token.transfer(alice, init_bal*10**token.decimals());
        uint256 total = escrow.calculateTotal(amounts);
        vm.prank(alice);
        token.approve(address(escrow), total);
        vm.prank(alice);
        escrow.createPayment(
            recievers,
            ERC20Payments.PaymentID(refId, oracle),
            address(token),
            amounts
        );
    }

    function preloadedInit() private returns (bytes16 refId) {
        refId = "stringliteral";
        address[] memory recievers = new address[](1);
        uint256[] memory amounts = new uint256[](1);
        recievers[0] = bob;
        amounts[0] = 1000*10**token.decimals();
        uint initBal = 10000;
        aliceCreateEscrow(refId, recievers, amounts, carol, initBal);
    }

    function testCreateEscrowState() public {
        console.log("Testing Correct Escrow Creation");
        token.transfer(alice, 10000*10**18);
        uint256 initialBalance = token.balanceOf(alice);
        bytes16 refId = "stringliteral";
        address[] memory recievers = new address[](1);
        uint256[] memory amounts = new uint256[](1);
        recievers[0] = bob;
        amounts[0] = 1000*10**token.decimals();
        vm.prank(alice);
        token.approve(address(escrow), amounts[0]);
        vm.prank(alice);
        escrow.createPayment(
            recievers, ERC20Payments.PaymentID(refId, alice),
            address(token),
            amounts
        );
        (address s, address[] memory r, address tc, uint256[] memory a, address o, ERC20Payments.PaymentState st) = escrow.getContract(refId);
        require(initialBalance-token.balanceOf(alice) == 1000*10**token.decimals(), "Balance should be decreased by 1000");
        require(token.balanceOf(alice) == 9000*10**token.decimals(), "Alice should have 9k");
        require(s == alice, "sender is wrong");
        require(r[0] == bob, "receiver is not Bob");
        require(tc == address(token), "token contract is not token");
        require(a[0] == 1000*10**token.decimals(), "amount is not 1000");
        require(o == alice, "oracle is not alice");
        require(st == ERC20Payments.PaymentState.Created, "state is not Created");
    }

    function testCancelEscrow() public {
        console.log("Testing Correct Escrow Cancelation");
        bytes16 refId = "stringliteral";
        address[] memory recievers = new address[](1);
        uint256[] memory amounts = new uint256[](1);
        recievers[0] = bob;
        amounts[0] = 1000*10**token.decimals();
        uint init_bal = 10000;
        aliceCreateEscrow(refId, recievers, amounts, carol, init_bal);
        vm.prank(carol);
        escrow.cancelPayment(refId);
        (address s, address[] memory r, address tc, uint256[] memory a, address o, ERC20Payments.PaymentState st) = escrow.getContract(refId);
        require(token.balanceOf(alice) == 10000*10**token.decimals(), "Alice should have 10k");
        require((token.balanceOf(address(escrow)) == 0), "Escrow should have 0 balance");
        require(s == alice, "sender is wrong");
        require(r[0] == bob, "receiver is not Bob");
        require(tc == address(token), "token contract is not token");
        require(a[0] == 1000*10**token.decimals(), "amount is not 1000");
        require(o == carol, "oracle is not carol");
        require(st == ERC20Payments.PaymentState.Canceled, "state is not Canceled");
    }

    function testWithdrawEscrow() public {
        console.log("Testing Correct Escrow Withdrawal");
        bytes16 refId = preloadedInit();
        vm.startPrank(carol);
        escrow.claimPayment(refId);
        vm.stopPrank();
        (address s, address[] memory r, address tc, uint256[] memory a, address o, ERC20Payments.PaymentState st) = escrow.getContract(refId);
        require(token.balanceOf(alice) == 9000*10**token.decimals(), "Alice should have 9k");
        require(token.balanceOf(bob) == 1000*10**token.decimals(), "Bob should have 1k");
        require(token.balanceOf(carol) == 0, "Carol should have 0");
        require(token.balanceOf(address(escrow)) == 0, "Escrow should have 0 balance");
        require(s == alice, "sender is wrong");
        require(r[0] == bob, "receiver is not Bob");
        require(tc == address(token), "token contract is not token");
        require(a[0] == 1000*10**token.decimals(), "amount is not 1000");
        require(o == carol, "oracle is not carol");
        require(st == ERC20Payments.PaymentState.Withdrawn, "state is not Withdrawn");
    }

    function testFailRepeatedCreate() public {
        console.log("Testing Repeated Escrow Creation");
        bytes16 refId = "stringliteral";
        address[] memory recievers = new address[](1);
        uint256[] memory amounts = new uint256[](1);
        recievers[0] = bob;
        amounts[0] = 1000*10**token.decimals();
        uint init_bal = 10000;
        aliceCreateEscrow(refId, recievers, amounts, carol, init_bal);
        // Idk why I can't get expect revert to work
        // vm.expectRevert(bytes("contractId already exists"));
        aliceCreateEscrow(refId, recievers, amounts, carol, init_bal);
    }

    function testFailWithdraw() public {
        console.log("Testing Escrow Withdrawal");
        bytes16 refId = preloadedInit();
        vm.prank(carol);
        // Idk why I can't get expect revert to work
        // vm.expectRevert(bytes("contractId does not exist"));
        vm.prank(alice);
        escrow.claimPayment(refId);
    }

    function testMultipleRecievers(uint len) public {
        vm.assume(len > 0 && len < 500);
        uint256[] memory amounts = new uint256[](len);
        for (uint i = 0; i < len; i++) {
            amounts[i] = (i+1)*100*10**token.decimals();
        }
        uint initBal = escrow.calculateTotal(amounts);
        token.transfer(alice, initBal);

        console.log("Testing Multiple Recievers");
        bytes16 refId = "stringliteral";
        address[] memory recievers = new address[](amounts.length);
        address payable[] memory payables = utils.createUsers(amounts.length);
        for (uint i = 0; i < len; i++) {
            recievers[i] = payables[i];
        }

        vm.startPrank(alice);
        token.approve(address(escrow), initBal);
        escrow.createPayment(recievers, ERC20Payments.PaymentID(refId, carol), address(token), amounts);
        vm.stopPrank();
        vm.startPrank(carol);
        escrow.claimPayment(refId);
        vm.stopPrank();
        (address s, address[] memory r, address tc, , address o, ERC20Payments.PaymentState st) = escrow.getContract(refId);
        require(token.balanceOf(alice) == 0, "Alice should have 0");
        for (uint i = 0; i < amounts.length; i++) {
            require(token.balanceOf(payables[i]) == amounts[i], string(abi.encodePacked("Recipient ", i, " should have ", amounts[i])));
        }
        require(token.balanceOf(address(escrow)) == 0, "Escrow should have 0 balance");
        require(s == alice, "sender is wrong");
        require(r[0] == payables[0], "receiver is not Bob");
        require(tc == address(token), "token contract is not token");
        require(o == carol, "oracle is not carol");
        require(st == ERC20Payments.PaymentState.Withdrawn, "state is not Confirmed");
    }

    function testFailBobCancels() public {
        console.log("Testing Bob Cancelling");
        bytes16 refId = preloadedInit();
        vm.prank(bob);
        escrow.cancelPayment(refId);
    }
}



