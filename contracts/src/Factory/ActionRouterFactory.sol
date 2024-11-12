// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import {Router} from "../ActionRouter.sol";

contract RouterFactory {
    event Deployed(address addr, uint256 salt);

    function deploy(uint256 salt) public payable returns (address) {
        bytes32 saltBytes = bytes32(salt);
        return address(new Router{salt: saltBytes}());
    }

    function getAddress(uint256 salt) public view returns (address) {
        bytes memory bytecode = abi.encodePacked(type(Router).creationCode, abi.encode());
        bytes32 saltBytes = bytes32(salt);
        bytes32 hash = keccak256(abi.encodePacked(bytes1(0xff), address(this), saltBytes, keccak256(bytecode)));
        return address(uint160(uint256(hash)));
    }
}
