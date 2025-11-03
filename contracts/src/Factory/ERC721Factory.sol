// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "../MyNFT.sol";

contract ERC721Factory {
    event Deployed(address addr, uint256 salt, address initialOwner);

    function deployERC721(uint256 salt, string memory cname, string memory csymbol) public returns (address) {
        bytes32 saltBytes = bytes32(salt);
        MyNFT token = new MyNFT{salt: saltBytes}(cname, csymbol, msg.sender);
        emit Deployed(address(token), salt, msg.sender);
        return address(token);
    }

    function getDeploymentAddress(uint256 salt, string memory cname, string memory csymbol) public view returns (address addr) {
        bytes memory bytecode = abi.encodePacked(type(MyNFT).creationCode, abi.encode(cname, csymbol, msg.sender));
        bytes32 saltBytes = bytes32(salt);
        bytes32 bytecodeHash;
        assembly {
            bytecodeHash := keccak256(add(bytecode, 32), mload(bytecode))
        }

        assembly {
            let ptr := mload(0x40)
            mstore(ptr, 0xff)
            mstore(add(ptr, 0x01), address())
            mstore(add(ptr, 0x15), saltBytes)
            mstore(add(ptr, 0x35), bytecodeHash)
            let hash := keccak256(ptr, 0x55)
            addr := and(hash, 0x00FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF)
        }
    }
}
