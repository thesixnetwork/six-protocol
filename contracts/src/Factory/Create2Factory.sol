// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import {Example} from "../Example.sol";

contract Factory {
    event Deployed(address addr, uint256 salt);

    function deploy(
        string memory message,
        uint256 salt
    ) public payable returns (address) {
        bytes32 saltBytes = bytes32(salt);
        return address(new Example{salt: saltBytes}(message));
    }

    function getAddress(
        bytes memory code,
        uint256 salt
    ) public view returns (address) {
        bytes32 hash = keccak256(
            abi.encodePacked(bytes1(0xff), address(this), salt, keccak256(code))
        );
        return address(uint160(uint256(hash)));
    }
}
