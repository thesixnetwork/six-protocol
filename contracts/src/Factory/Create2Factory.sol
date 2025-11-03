// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import {Router} from "../ActionRouter.sol";

contract RouterFactory {
    event Deployed(address addr, uint256 salt);

    function deploy(uint256 salt) public payable returns (address) {
        bytes32 saltBytes = bytes32(salt);
        return address(new Router{salt: saltBytes}());
    }

    function getAddress(uint256 salt) public view returns (address addr) {
        bytes memory bytecode = abi.encodePacked(type(Router).creationCode, abi.encode());
        bytes32 saltBytes = bytes32(salt);
        bytes32 bytecodeHash;
        assembly {
            bytecodeHash := keccak256(add(bytecode, 32), mload(bytecode))
        }

        // Inline assembly for CREATE2 address calculation:
        assembly {
            let ptr := mload(0x40)
            mstore(ptr, 0xff)                          // 1 byte
            mstore(add(ptr, 0x01), address())          // 20 bytes
            mstore(add(ptr, 0x15), saltBytes)          // 32 bytes
            mstore(add(ptr, 0x35), bytecodeHash)       // 32 bytes
            let hash := keccak256(ptr, 0x55)           // 1+20+32+32 = 85 bytes (0x55)
            addr := and(hash, 0x00FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF) // last 20 bytes
        }
    }
}
