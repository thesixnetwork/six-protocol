// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/opcode.sol";

contract OpcodeTest is Test {
    OpcodeCoverage oc;

    function setUp() public {
        oc = new OpcodeCoverage();
    }

    function test_ADD() public view {
        assertEq(oc.op_ADD(1, 2), 3);
    }
    function test_MUL() public view {
        assertEq(oc.op_MUL(2, 3), 6);
    }
    function test_SUB() public view {
        assertEq(oc.op_SUB(7, 4), 3);
    }
    function test_DIV() public view {
        assertEq(oc.op_DIV(6, 2), 3);
    }
    function test_SDIV() public view {
        assertEq(oc.op_SDIV(-6, 2), -3);
    }
    function test_MOD() public view {
        assertEq(oc.op_MOD(8, 3), 2);
    }
    function test_SMOD() public view{
        assertEq(oc.op_SMOD(-8, 3), -2);
    }
    function test_ADDMOD() public view {
        assertEq(oc.op_ADDMOD(5, 6, 7), addmod(5, 6, 7));
    }
    function test_MULMOD() public view {
        assertEq(oc.op_MULMOD(5, 6, 7), mulmod(5, 6, 7));
    }
    function test_EXP() public view {
        assertEq(oc.op_EXP(2, 5), 32);
    }
    function test_SIGNEXTEND() public view {
        assertEq(oc.op_SIGNEXTEND(1, int256(0x0000000000000080)), int256(-128));
    }

    function test_LT() public view {
        assertTrue(oc.op_LT(1, 2));
    }
    function test_GT() public view {
        assertTrue(oc.op_GT(2, 1));
    }
    function test_SLT() public view {
        assertTrue(oc.op_SLT(-2, 1));
    }
    function test_SGT() public view {
        assertTrue(oc.op_SGT(2, -1));
    }
    function test_EQ() public view {
        assertTrue(oc.op_EQ(3, 3));
    }
    function test_ISZERO() public view {
        assertTrue(oc.op_ISZERO(0));
    }
    function test_AND() public view {
        assertEq(oc.op_AND(0xFF, 0x0F), 0x0F);
    }
    function test_OR() public view {
        assertEq(oc.op_OR(0xF0, 0x0F), 0xFF);
    }
    function test_XOR() public view {
        assertEq(oc.op_XOR(0xF0, 0x0F), 0xFF);
    }
    function test_NOT() public view {
        assertEq(oc.op_NOT(0), type(uint256).max);
    }
    function test_BYTE() public view {
        assertEq(oc.op_BYTE(0, hex"123456"), 0x12);
    }
    function test_SHL() public view {
        assertEq(oc.op_SHL(1, 2), 4);
    }
    function test_SHR() public view {
        assertEq(oc.op_SHR(10, 2), 2);
    }
    function test_SAR() public view {
        assertEq(oc.op_SAR(-4, 1), -2);
    }

    function test_KECCAK256() public view {
        bytes memory data = abi.encodePacked(uint256(123));
        assertEq(oc.op_KECCAK256(data), keccak256(data));
    }

    function test_ADDRESS() public view {
        assertEq(oc.op_ADDRESS(), address(oc));
    }
    function test_BALANCE() public view {
        assertEq(oc.op_BALANCE(address(this)), address(this).balance);
    }
    function test_ORIGIN() public view {
        assertEq(oc.op_ORIGIN(), tx.origin);
    }
    function test_CALLER() public view {
        assertEq(oc.op_CALLER(), address(this));
    }
    function test_CHAINID() public view {
        assertEq(oc.op_CHAINID(), block.chainid);
    }

    function test_CODESIZE() public view {
        assertGt(oc.op_CODESIZE(), 0);
    }
    function test_GASPRICE() public view {
        assertEq(oc.op_GASPRICE(), tx.gasprice);
    }
    function test_SELFBALANCE() public view {
        assertEq(oc.op_SELFBALANCE(), address(oc).balance);
    }

    function test_BLOCKHASH() public view {
        // blockhash(0) returns 0, but recent block's hash should not be 0.
        assertEq(oc.op_BLOCKHASH(0), bytes32(0));
    }
    function test_COINBASE() public view {
        assertEq(oc.op_COINBASE(), block.coinbase);
    }
    function test_TIMESTAMP() public view {
        assertEq(oc.op_TIMESTAMP(), block.timestamp);
    }
    function test_NUMBER() public view {
        assertEq(oc.op_NUMBER(), block.number);
    }
    function test_DIFFICULTY() public view {
        assertEq(oc.op_DIFFICULTY(), block.prevrandao);
    }
    function test_GASLIMIT() public view {
        assertEq(oc.op_GASLIMIT(), block.gaslimit);
    }

    function test_SSTORE_SLOAD() public {
        oc.op_SSTORE(keccak256("key"), bytes32(uint256(1234)));
        assertEq(oc.op_SLOAD(keccak256("key")), bytes32(uint256(1234)));
    }
    function test_MSTORE_MLOAD() public view {
        assertEq(oc.op_MSTORE(0x80, 0x1234), 0x1234);
    }
    function test_MSTORE8() public view {
        assertEq(oc.op_MSTORE8(0x80, 0x12), 0x12);
    }

    function test_LOG0() public {
        oc.op_LOG0(0);
        // You can check logs in the forge trace, but not easily in Solidity.
    }

    function test_GAS() public view {
        assertGt(oc.op_GAS(), 0);
    }
    // function test_PC() public {
    //     assertGt(oc.op_PC(), 0);
    // }
    // function test_MSIZE() public {
    //     assertGt(oc.op_MSIZE(), 0);
    // }
}