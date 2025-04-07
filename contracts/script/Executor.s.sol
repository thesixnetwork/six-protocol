// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {INFTMNGR, NFTMNGR_PRECOMPILE_ADDRESS} from "../src/precompiles/INFTManager.sol";

contract AddRouterExecutor is Script {
    address ownerAddress;
    uint64 currentNonce;
    address routerContractAddress;
    string nftSchema;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
        nftSchema = vm.envString("NFT_SCHEMA");

        string
            memory routerContractInfoPath = "./broadcast/ActionRouter.s.sol/666/run-latest.json";
        string memory routerContractInfo = vm.readFile(routerContractInfoPath);
        bytes memory routerJsonParsed = vm.parseJson(
            routerContractInfo,
            ".transactions[0].contractAddress"
        );
        routerContractAddress = abi.decode(routerJsonParsed, (address));
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        bytes memory data = abi.encodeWithSignature(
            "addActionExecutor(string,address)",
            nftSchema,
            routerContractAddress
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

contract RemoveRouterExecutor is Script {
    address ownerAddress;
    uint64 currentNonce;
    address routerContractAddress;
    string nftSchema;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
        nftSchema = vm.envString("NFT_SCHEMA");

        string
            memory routerContractInfoPath = "./broadcast/ActionRouter.s.sol/666/run-latest.json";
        string memory routerContractInfo = vm.readFile(routerContractInfoPath);
        bytes memory routerJsonParsed = vm.parseJson(
            routerContractInfo,
            ".transactions[0].contractAddress"
        );
        routerContractAddress = abi.decode(routerJsonParsed, (address));
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        bytes memory data = abi.encodeWithSignature(
            "removeActionExecutor(string,address)",
            nftSchema,
            routerContractAddress
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


contract AddExecutor is Script {
    address ownerAddress;
    uint64 currentNonce;
    address executorAddress;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
        executorAddress = vm.envAddress("EXECUTOR");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        string memory nftSchema = "TechSauceV1.GlobalSummit2024";

        bytes memory data = abi.encodeWithSignature(
            "addActionExecutor(string,address)",
            nftSchema,
            executorAddress
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


contract RemoveExecutor is Script {
    address ownerAddress;
    uint64 currentNonce;
    address executorAddress;
    string nftSchema;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
        nftSchema = vm.envString("NFT_SCHEMA");
        executorAddress = vm.envAddress("EXECUTOR");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        bytes memory data = abi.encodeWithSignature(
            "removeActionExecutor(string,address)",
            nftSchema,
            executorAddress
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
