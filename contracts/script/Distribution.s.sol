// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {IDistr, DISTR_PRECOMPILE_ADDRESS} from "../src/precompiles/IDistribution.sol";

contract SetAddressScript is Script {
    address ownerAddress;
    uint64 currentNonce;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        address contractAddress = DISTR_PRECOMPILE_ADDRESS;
        // Execute the transaction
        (bool success, ) = contractAddress.call(
            abi.encodeWithSignature(
                "setWithdrawAddress(string)",
                "6x1kch0sdjr5tuvjh0h3a55c6l5sr6m0phjeag9f2"
            )
        );
        require(success, "Transaction failed");

        // Log the success message
        console.log("Transfer success!");
        vm.stopBroadcast();
    }
}

contract WithDrawScript is Script {
    address ownerAddress;
    uint64 currentNonce;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        address contractAddress = DISTR_PRECOMPILE_ADDRESS;
        // Execute the transaction
        (bool success, ) = contractAddress.call(
            abi.encodeWithSignature(
                "withdrawRewards(string)",
                "6xvaloper1t3p2vzd7w036ahxf4kefsc9sn24pvlqpmk79jh"
            )
        );
        require(success, "Transaction failed");

        // Log the success message
        console.log("Transfer success!");
        vm.stopBroadcast();
    }
}
