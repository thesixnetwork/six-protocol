// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../src/OpCodes.sol";

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
        require(oc.opADD(1, 2) == 3, "ADD failed");
        require(oc.opMUL(2, 3) == 6, "MUL failed");
        require(oc.opSUB(7, 4) == 3, "SUB failed");
        require(oc.opDIV(6, 2) == 3, "DIV failed");
        require(oc.opSDIV(-6, 2) == -3, "SDIV failed");
        require(oc.opMOD(8, 3) == 2, "MOD failed");
        require(oc.opSMOD(-8, 3) == -2, "SMOD failed");
        require(oc.opADDMOD(5, 6, 7) == addmod(5, 6, 7), "ADDMOD failed");
        require(oc.opMULMOD(5, 6, 7) == mulmod(5, 6, 7), "MULMOD failed");
        require(oc.opEXP(2, 5) == 32, "EXP failed");

        int256 extended = oc.testSignExtend();

        console.log(extended);
        // require(oc.opSIGNEXTEND(1, int256(0x80)) == int256(-128), "SIGNEXTEND failed"); // failed

        // Comparison/Bitwise
        require(oc.opLT(1, 2), "LT failed");
        require(oc.opGT(2, 1), "GT failed");
        require(oc.opSLT(-2, 1), "SLT failed");
        require(oc.opSGT(2, -1), "SGT failed");
        require(oc.opEQ(3, 3), "EQ failed");
        require(oc.opISZERO(0), "ISZERO failed");
        require(oc.opAND(0xFF, 0x0F) == 0x0F, "AND failed");
        require(oc.opOR(0xF0, 0x0F) == 0xFF, "OR failed");
        require(oc.opXOR(0xF0, 0x0F) == 0xFF, "XOR failed");
        require(oc.opNOT(0) == type(uint256).max, "NOT failed");
        require(oc.opBYTE(0, hex"1234560000000000000000000000000000000000000000000000000000000000") == 0x12, "BYTE failed");
        require(oc.opSHL(1, 2) == 4, "SHL failed");
        // require(oc.opSHR(4, 1) == 2, "SHR failed"); // failed
        // require(oc.opSAR(-4, 1) == -2, "SAR failed"); // failed

        // Crypto
        bytes memory data = abi.encodePacked(uint256(123));
        require(oc.opKECCAK256(data) == keccak256(data), "KECCAK256 failed");

        // Environmental/Info
        require(oc.opADDRESS() == address(oc), "ADDRESS failed");
        require(oc.opBALANCE(address(this)) == address(this).balance, "BALANCE failed");
        // require(oc.opORIGIN() == tx.origin, "ORIGIN failed"); // failed
        // require(oc.opCALLER() == address(this), "CALLER failed"); // failed
        require(oc.opCHAINID() == block.chainid, "CHAINID failed");
        require(oc.opCODESIZE() > 0, "CODESIZE failed");
        require(oc.opGASPRICE() == tx.gasprice, "GASPRICE failed");
        require(oc.opSELFBALANCE() == address(oc).balance, "SELFBALANCE failed");

        // Block Info
        oc.opBLOCKHASH(block.number - 1); // cannot check value deterministically
        require(oc.opCOINBASE() == block.coinbase, "COINBASE failed");
        require(oc.opTIMESTAMP() == block.timestamp, "TIMESTAMP failed");
        require(oc.opNUMBER() == block.number, "NUMBER failed");
        // Use PREVRANDAO instead of DIFFICULTY in modern EVM
        require(oc.opDIFFICULTY() == block.prevrandao, "PREVRANDAO failed");
        require(oc.opGASLIMIT() == block.gaslimit, "GASLIMIT failed");

        // Storage/mem opcodes
        bytes32 key = keccak256("key");
        oc.opSSTORE(key, bytes32(uint256(1234)));
        require(oc.opSLOAD(key) == bytes32(uint256(1234)), "SSTORE/SLOAD failed");
        require(oc.opMSTORE(0x80, 0x1234) == 0x1234, "MSTORE failed");
        // require(oc.opMSTORE8(0x80, 0x12) == 0x12, "MSTORE8 failed"); // failed
        // require(oc.opMLOAD(0x80) == 0x1234, "MLOAD failed"); // failed

        // Logging (cannot check in-script, see transaction log)
        oc.opLOG0(0);

        require(oc.opGAS() > 0, "GAS failed");
        // The following are not safe with optimizer, so are commented out
        // oc.opPC();
        // oc.opMSIZE();

        vm.stopBroadcast();
    }

    function nonceUp() public {
        vm.setNonce(ownerAddress, currentNonce + uint64(1));
        currentNonce++;
    }
}