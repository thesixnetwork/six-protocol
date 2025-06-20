// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../src/opcode.sol";

contract OpcodeScript is Script {

    address ownerAddress;
    uint64 currentNonce;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
    }

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        OpcodeCoverage oc = new OpcodeCoverage();

        // Arithmetic
        require(oc.op_ADD(1, 2) == 3, "ADD failed");
        require(oc.op_MUL(2, 3) == 6, "MUL failed");
        require(oc.op_SUB(7, 4) == 3, "SUB failed");
        require(oc.op_DIV(6, 2) == 3, "DIV failed");
        require(oc.op_SDIV(-6, 2) == -3, "SDIV failed");
        require(oc.op_MOD(8, 3) == 2, "MOD failed");
        require(oc.op_SMOD(-8, 3) == -2, "SMOD failed");
        require(oc.op_ADDMOD(5, 6, 7) == addmod(5, 6, 7), "ADDMOD failed");
        require(oc.op_MULMOD(5, 6, 7) == mulmod(5, 6, 7), "MULMOD failed");
        require(oc.op_EXP(2, 5) == 32, "EXP failed");

        int256 extended = oc.testSignExtend();

        console.log(extended);
        // require(oc.op_SIGNEXTEND(1, int256(0x80)) == int256(-128), "SIGNEXTEND failed"); // failed

        // Comparison/Bitwise
        require(oc.op_LT(1, 2), "LT failed");
        require(oc.op_GT(2, 1), "GT failed");
        require(oc.op_SLT(-2, 1), "SLT failed");
        require(oc.op_SGT(2, -1), "SGT failed");
        require(oc.op_EQ(3, 3), "EQ failed");
        require(oc.op_ISZERO(0), "ISZERO failed");
        require(oc.op_AND(0xFF, 0x0F) == 0x0F, "AND failed");
        require(oc.op_OR(0xF0, 0x0F) == 0xFF, "OR failed");
        require(oc.op_XOR(0xF0, 0x0F) == 0xFF, "XOR failed");
        require(oc.op_NOT(0) == type(uint256).max, "NOT failed");
        require(oc.op_BYTE(0, hex"1234560000000000000000000000000000000000000000000000000000000000") == 0x12, "BYTE failed");
        require(oc.op_SHL(1, 2) == 4, "SHL failed");
        // require(oc.op_SHR(4, 1) == 2, "SHR failed"); // failed
        // require(oc.op_SAR(-4, 1) == -2, "SAR failed"); // failed

        // Crypto
        bytes memory data = abi.encodePacked(uint256(123));
        require(oc.op_KECCAK256(data) == keccak256(data), "KECCAK256 failed");

        // Environmental/Info
        require(oc.op_ADDRESS() == address(oc), "ADDRESS failed");
        require(oc.op_BALANCE(address(this)) == address(this).balance, "BALANCE failed");
        // require(oc.op_ORIGIN() == tx.origin, "ORIGIN failed"); // failed
        // require(oc.op_CALLER() == address(this), "CALLER failed"); // failed
        require(oc.op_CHAINID() == block.chainid, "CHAINID failed");
        require(oc.op_CODESIZE() > 0, "CODESIZE failed");
        require(oc.op_GASPRICE() == tx.gasprice, "GASPRICE failed");
        require(oc.op_SELFBALANCE() == address(oc).balance, "SELFBALANCE failed");

        // Block Info
        oc.op_BLOCKHASH(block.number - 1); // cannot check value deterministically
        require(oc.op_COINBASE() == block.coinbase, "COINBASE failed");
        require(oc.op_TIMESTAMP() == block.timestamp, "TIMESTAMP failed");
        require(oc.op_NUMBER() == block.number, "NUMBER failed");
        // Use PREVRANDAO instead of DIFFICULTY in modern EVM
        require(oc.op_DIFFICULTY() == block.prevrandao, "PREVRANDAO failed");
        require(oc.op_GASLIMIT() == block.gaslimit, "GASLIMIT failed");

        // Storage/mem opcodes
        bytes32 key = keccak256("key");
        oc.op_SSTORE(key, bytes32(uint256(1234)));
        require(oc.op_SLOAD(key) == bytes32(uint256(1234)), "SSTORE/SLOAD failed");
        require(oc.op_MSTORE(0x80, 0x1234) == 0x1234, "MSTORE failed");
        // require(oc.op_MSTORE8(0x80, 0x12) == 0x12, "MSTORE8 failed"); // failed
        // require(oc.op_MLOAD(0x80) == 0x1234, "MLOAD failed"); // failed

        // Logging (cannot check in-script, see transaction log)
        oc.op_LOG0(0);

        require(oc.op_GAS() > 0, "GAS failed");
        // The following are not safe with optimizer, so are commented out
        // oc.op_PC();
        // oc.op_MSIZE();

        vm.stopBroadcast();
    }

    function nonceUp() public {
        vm.setNonce(ownerAddress, currentNonce + uint64(1));
        currentNonce++;
    }
}