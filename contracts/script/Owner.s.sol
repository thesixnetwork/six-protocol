// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {INFTMNGR, NFTMNGR_PRECOMPILE_ADDRESS} from "../src/precompiles/INFTManager.sol";

contract ChangeOrgOwner is Script {
    address ownerAddress;
    address destAddress;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        destAddress = vm.envAddress("DEST_ADDRESS");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        string memory orgName = "TechSauceV1";

        bytes memory data = abi.encodeWithSignature(
            "changeOrgOwner(string,address)",
            orgName,
            destAddress
        );

        bool success = callKeeper(data);

        require(success, "Transaction failed. Check error message below:");

        // Log the success message
        console.log("Action executed successfully!");
        vm.stopBroadcast();
    }

    function callKeeper(bytes memory datas) public payable returns (bool) {
        (bool success, ) = NFTMNGR_PRECOMPILE_ADDRESS.call{value: 0}(datas);
        if (!success) revert("transaction failed");
        return success;
    }
}

contract ChangeSchemaOwner is Script {
    address ownerAddress;
    address destAddress;
    string nftSchema;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        nftSchema = vm.envString("NFT_SCHEMA");
        destAddress = vm.envAddress("DEST_ADDRESS");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        bytes memory data = abi.encodeWithSignature(
            "changeSchemaOwner(string,address)",
            nftSchema,
            destAddress
        );

        bool success = callKeeper(data);

        require(success, "Transaction failed. Check error message below:");

        // Log the success message
        console.log("Action executed successfully!");
        vm.stopBroadcast();
    }

    function callKeeper(bytes memory datas) public payable returns (bool) {
        (bool success, ) = NFTMNGR_PRECOMPILE_ADDRESS.call{value: 0}(datas);
        if (!success) revert("transaction failed");
        return success;
    }
}
