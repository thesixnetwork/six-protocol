// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {SimpleTokenFactory} from "../src/SimpleTokenFactory.sol";

contract SimpleTokenFactoryDeployer is Script {
    SimpleTokenFactory public factory;
    address ownerAddress;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
    }

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        factory = new SimpleTokenFactory(
            ownerAddress
        );
        console.log(address(factory));

        vm.stopBroadcast();
    }
}

contract CreateTokenScript is Script {
    SimpleTokenFactory public factory;
    address factoryAddress;

    address ownerAddress;
    uint64 currentNonce;

    function setUp() public {
        string
            memory factoryContractInfoPath = "./broadcast/SimpleTokenFactoryDeployer.s.sol/666/run-latest.json";
        string memory factoryContractInfo = vm.readFile(
            factoryContractInfoPath
        );
        bytes memory routerJsonParsed = vm.parseJson(
            factoryContractInfo,
            ".transactions[0].contractAddress"
        );
        factoryAddress = abi.decode(routerJsonParsed, (address));

        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
    }

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        factory = SimpleTokenFactory(factoryAddress);
        console.log(address(factory));

        factory.grantRole(
            0x60f84e22dcea217c9a7721aaea62b0971fe29d0da3d3c250600e83e005d17763,
            ownerAddress
        );

        factory.createToken(
            "TestTokenD",
            "TTD",
            1000000000000000000000000000,
            ownerAddress
        );

        vm.stopBroadcast();
    }

    function nonceUp() public {
        vm.setNonce(ownerAddress, currentNonce + uint64(1));
        currentNonce++;
    }
}
