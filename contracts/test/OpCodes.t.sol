// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import "../src/OpCodes.sol";

contract OpcodeTest is Test {
    OpcodeCoverage opcodes;
    
    function setUp() public {
        opcodes = new OpcodeCoverage();
    }
    
    // Arithmetic tests
    function test_OpAdd() public view {
        assertEq(opcodes.opAdd(2, 3), 5);
    }
    
    function test_OpMul() public view {
        assertEq(opcodes.opMul(2, 3), 6);
    }
    
    function test_OpSub() public view {
        assertEq(opcodes.opSub(5, 3), 2);
    }
    
    function test_OpDiv() public view {
        assertEq(opcodes.opDiv(6, 3), 2);
    }
    
    function test_OpSdiv() public view {
        assertEq(opcodes.opSdiv(-6, 3), -2);
    }
    
    function test_OpMod() public view {
        assertEq(opcodes.opMod(5, 3), 2);
    }
    
    function test_OpSmod() public view {
        assertEq(opcodes.opSmod(-5, 3), -2);
    }
    
    function test_OpAddmod() public view {
        assertEq(opcodes.opAddmod(5, 3, 7), 1);
    }
    
    function test_OpMulmod() public view {
        assertEq(opcodes.opMulmod(5, 3, 7), 1);
    }
    
    function test_OpExp() public view {
        assertEq(opcodes.opExp(2, 3), 8);
    }
    
    function test_OpSignextend() public view {
        assertEq(opcodes.opSignextend(0, int256(0xff)), -1);
    }
    
    // Comparison/Bitwise tests
    function test_OpLt() public view {
        assertTrue(opcodes.opLt(2, 3));
        assertFalse(opcodes.opLt(3, 2));
    }
    
    function test_OpGt() public view {
        assertTrue(opcodes.opGt(3, 2));
        assertFalse(opcodes.opGt(2, 3));
    }
    
    function test_OpSlt() public view {
        assertTrue(opcodes.opSlt(-2, 3));
        assertFalse(opcodes.opSlt(3, -2));
    }
    
    function test_OpSgt() public view {
        assertTrue(opcodes.opSgt(3, -2));
        assertFalse(opcodes.opSgt(-2, 3));
    }
    
    function test_OpEq() public view {
        assertTrue(opcodes.opEq(3, 3));
        assertFalse(opcodes.opEq(3, 2));
    }
    
    function test_OpIszero() public view {
        assertTrue(opcodes.opIszero(0));
        assertFalse(opcodes.opIszero(1));
    }
    
    function test_OpAnd() public view {
        assertEq(opcodes.opAnd(3, 5), 1);
    }
    
    function test_OpOr() public view {
        assertEq(opcodes.opOr(3, 5), 7);
    }
    
    function test_OpXor() public view {
        assertEq(opcodes.opXor(3, 5), 6);
    }
    
    function test_OpNot() public view {
        // Testing on a small value for clearer results
        assertEq(opcodes.opNot(0), type(uint256).max);
    }
    
    function test_OpByte() public view {
        assertEq(opcodes.opByte(31, bytes32(uint256(0xff))), 0xff);
    }
    
    function test_OpShl() public view {
        assertEq(opcodes.opShl(1, 1), 2);
    }
    
    function test_OpShr() public view {
        assertEq(opcodes.opShr(1, 2), 1);
    }
    
    function test_OpSar() public view {
        assertEq(opcodes.opSar(1, -4), -2);
    }
    
    // Crypto tests
    function test_OpKeccak256() public view {
        bytes memory data = abi.encodePacked("test");
        bytes32 expectedHash = keccak256("test");
        assertEq(opcodes.opKeccak256(data), expectedHash);
    }
    
    // Environmental/Info tests
    function test_OpAddress() public view {
        assertEq(opcodes.opAddress(), address(opcodes));
    }
    
    function test_OpBalance() public view {
        // Test with current contract
        assertEq(opcodes.opBalance(address(this)), address(this).balance);
    }
    
    function test_OpOrigin() public view {
        assertEq(opcodes.opOrigin(), tx.origin);
    }
    
    function test_OpCaller() public view {
        assertEq(opcodes.opCaller(), address(this));
    }
    
    function test_OpCallvalue() public {
        assertEq(opcodes.opCallvalue(), 0);
    }
    
    function test_OpCalldataload() public view {
        // This is harder to test directly, we'll just ensure it executes
        opcodes.opCalldataload(0);
    }
    
    function test_OpCalldatasize() public view {
        // The size will depend on the function call itself
        opcodes.opCalldatasize();
    }
    
    function test_OpCodesize() public view {
        assertTrue(opcodes.opCodesize() > 0);
    }
    
    function test_OpGasprice() public view {
        assertEq(opcodes.opGasprice(), tx.gasprice);
    }
    
    function test_OpChainid() public view {
        assertEq(opcodes.opChainid(), block.chainid);
    }
    
    function test_OpSelfbalance() public view {
        assertEq(opcodes.opSelfbalance(), address(opcodes).balance);
    }
    
    function test_OpBasefee() public view {
        assertEq(opcodes.opBasefee(), block.basefee);
    }
    
    // Block Info tests
    function test_OpBlockhash() public view {
        // This depends on the blockchain state
        opcodes.opBlockhash(block.number - 1);
    }
    
    function test_OpCoinbase() public view {
        assertEq(opcodes.opCoinbase(), block.coinbase);
    }
    
    function test_OpTimestamp() public view {
        assertEq(opcodes.opTimestamp(), block.timestamp);
    }
    
    function test_OpNumber() public view {
        assertEq(opcodes.opNumber(), block.number);
    }
    
    function test_OpDifficulty() public view {
        assertEq(opcodes.opDifficulty(), block.prevrandao);
    }
    
    function test_OpGaslimit() public view {
        assertEq(opcodes.opGaslimit(), block.gaslimit);
    }
    
    // Storage/Memory tests
    function test_OpSstore() public {
        bytes32 key = bytes32(uint256(1));
        bytes32 value = bytes32(uint256(42));
        opcodes.opSstore(key, value);
        assertEq(opcodes.opSload(key), value);
    }
    
    function test_OpSload() public {
        bytes32 key = bytes32(uint256(2));
        bytes32 value = bytes32(uint256(43));
        opcodes.opSstore(key, value);
        assertEq(opcodes.opSload(key), value);
    }
    
    function test_OpMstore() public view {
        assertEq(opcodes.opMstore(0, 123), 123);
    }
    
    function test_OpMstore8() public view {
        assertEq(opcodes.opMstore8(0, 255), 255);
    }
    
    function test_OpMload() public view {
        // This test depends on memory state which is challenging to verify directly
        opcodes.opMload(0);
    }
    
    // Logging test
    function test_OpLog0() public {
        // Just verify it doesn't revert
        opcodes.opLog0(bytes32(uint256(123)));
    }
    
    // Misc tests
    function test_OpGas() public view {
        assertTrue(opcodes.opGas() > 0);
    }
    
    function testSignExtend() public view {
        int256 result = opcodes.testSignExtend();
        // Result will depend on the specific implementation
        assertEq(result, int256(int16(0x0080)));
    }
}
