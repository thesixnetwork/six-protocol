// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import {Script, console} from "forge-std/Script.sol";
import {Factory} from "../src/Factory/Create2Factory.sol";

import {ERC721Factory} from "../src/Factory/ERC721Factory.sol";
import {RouterFactory} from "../src/Factory/ActionRouterFactory.sol";

contract DeployERC721Factory is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        new ERC721Factory();
        vm.stopBroadcast();
    }
}

contract DeployRouterFactory is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        new RouterFactory();
        vm.stopBroadcast();
    }
}

contract Create2FactoryScript is Script {
    Factory public factory;

    function setUp() public {}

    function run() public {
        vm.startBroadcast(vm.envUint("PRIVATE_KEY"));

        factory = new Factory();

        vm.stopBroadcast();
    }
}