// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../MyNFT.sol";

contract ERC721Factory {
    event Deployed(address addr, uint256 salt);

    function deployERC721(uint256 salt) public returns (address) {
        bytes32 saltBytes = bytes32(salt);
        MyNFT token = new MyNFT{salt: saltBytes}("MyNFT", "NFT");
        emit Deployed(address(token), salt);
        return address(token);
    }

    function getDeploymentAddress(uint256 salt) public view returns (address) {
        bytes memory bytecode = abi.encodePacked(type(MyNFT).creationCode, abi.encode());
        bytes32 saltBytes = bytes32(salt);
        bytes32 hash = keccak256(abi.encodePacked(bytes1(0xff), address(this), saltBytes, keccak256(bytecode)));
        return address(uint160(uint256(hash)));
    }
}
