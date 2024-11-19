// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {INFTMNGR, NFTMNGR_PRECOMPILE_ADDRESS} from "../src/INFTManager.sol";

contract ResyncAttribute is Script {
    address ownerAddress;
    address destAddress;
    uint64 currentNonce;
    string nftSchema;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        destAddress = vm.envAddress("DEST_ADDRESS");
        nftSchema = vm.envString("NFT_SCHEMA");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        string memory tokenId = "1";

        bytes memory data = abi.encodeWithSignature(
            "resyncAttribute(string,string)",
            nftSchema,
            tokenId
        );

        bool success = callKeeeper(data);

        require(success, "Transaction failed. Check error message below:");

        // Log the success message
        console.log("Action executed successfully!");
        vm.stopBroadcast();
    }

    function callKeeeper(bytes memory datas) public payable returns (bool) {
        (bool success, ) = NFTMNGR_PRECOMPILE_ADDRESS.call{value: 0}(datas);
        if (!success) revert("transaction failed");
        return success;
    }
}

contract AttributeOverride is Script {
    address ownerAddress;
    uint64 currentNonce;
    string nftSchema;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
        nftSchema = vm.envString("NFT_SCHEMA");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        uint32 overrideType = 1;

        bytes memory data = abi.encodeWithSignature(
            "attributeOveride(string,uint32)",
            nftSchema,
            overrideType
        );

        bool success = callKeeeper(data);

        require(success, "Transaction failed. Check error message below:");

        // Log the success message
        console.log("Action executed successfully!");
        vm.stopBroadcast();
    }

    function callKeeeper(bytes memory datas) public payable returns (bool) {
        (bool success, ) = NFTMNGR_PRECOMPILE_ADDRESS.call{value: 0}(datas);
        if (!success) revert("transaction failed");
        return success;
    }
}


contract ShowAttribute is Script {
    address ownerAddress;
    string nftSchema;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        nftSchema = vm.envString("NFT_SCHEMA");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        bool toShow = true;
        string[] memory attributes = new string[](2);

        // TODO:: Append attributes
        attributes[0] = "points_a";
        attributes[1] = "points_b";

        bytes memory data = abi.encodeWithSignature(
            "showAttribute(string,bool,string[])",
            nftSchema,
            toShow,
            attributes
        );

        bool success = callKeeeper(data);

        require(success, "Transaction failed. Check error message below:");

        // Log the success message
        console.log("Action executed successfully!");
        vm.stopBroadcast();
    }

    function callKeeeper(bytes memory datas) public payable returns (bool) {
        (bool success, ) = NFTMNGR_PRECOMPILE_ADDRESS.call{value: 0}(datas);
        if (!success) revert("transaction failed");
        return success;
    }
}
