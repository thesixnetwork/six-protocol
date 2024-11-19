// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import {Script, console} from "forge-std/Script.sol";
// import {Counter} from "../src/Counter.sol";
import {Factory} from "../src/Factory/Create2Factory.sol";
import {Example} from "../src/Example.sol";

contract FactoryDeployScript is Script {
    Factory public factory;
    bytes public code = type(Example).creationCode;

    function setUp() public {}

    function run() public {

        address factoryAddress = address(0x0168D74E02ef5bC86EdA4Ff60c63a292CFF3223A);

        vm.startBroadcast(vm.envUint("PK"));

        factory = Factory(factoryAddress);
        factory.deploy("Hello", 3);   

        vm.stopBroadcast();
    }
}