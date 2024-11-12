// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {INFTMNGR, NFTMNGR_PRECOMPILE_ADDRESS} from "../src/INFTManager.sol";

contract SetBaseURI is Script {
    address ownerAddress;
    string nftSchema;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        nftSchema = vm.envString("NFT_SCHEMA");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        string memory newBaseUri = "idk";

        bytes memory data = abi.encodeWithSignature(
            "setBaseURI(string,string)",
            nftSchema,
            newBaseUri
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

contract SetMetadataFormat is Script {
    address ownerAddress;
    string nftSchema;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        nftSchema = vm.envString("NFT_SCHEMA");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        string memory newFormat = "opensea";

        bytes memory data = abi.encodeWithSignature(
            "setMetadataFormat(string,string)",
            nftSchema,
            newFormat
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


contract SetMintAuth is Script {
    address ownerAddress;
    string nftSchema;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        nftSchema = vm.envString("NFT_SCHEMA");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        uint32 authorizeTo = 1;

        bytes memory data = abi.encodeWithSignature(
            "setMintAuth(string,uint32)",
            nftSchema,
            authorizeTo
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

contract SetOriginChain is Script {
    address ownerAddress;
    string nftSchema;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        nftSchema = vm.envString("NFT_SCHEMA");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        string memory newChain = "newchain";

        bytes memory data = abi.encodeWithSignature(
            "setOriginChain(string,string)",
            nftSchema,
            newChain
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

contract SetOriginConrtact is Script {
    address ownerAddress;
    string nftSchema;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        nftSchema = vm.envString("NFT_SCHEMA");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        
        string memory stringContractAddress = "0xSomething";

        bytes memory data = abi.encodeWithSignature(
            "setOriginContract(string,string)",
            nftSchema,
            stringContractAddress
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

contract SetURIRetreival is Script {
    address ownerAddress;
    string nftSchema;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        nftSchema = vm.envString("NFT_SCHEMA");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        uint32 newMethod = 1;

        bytes memory data = abi.encodeWithSignature(
            "setUriRetreival(string,uint32)",
            nftSchema,
            newMethod
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