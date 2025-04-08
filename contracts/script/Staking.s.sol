// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {IStaking, STAKING_PRECOMPILE_ADDRESS} from "../src/precompiles/IStaking.sol";

contract delegateScript is Script {
    address ownerAddress;
    uint64 currentNonce;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        address contractAddress = STAKING_PRECOMPILE_ADDRESS;
        // Execute the transaction
        (bool success, bytes memory result) = contractAddress.call(
            abi.encodeWithSignature(
                "delegate(string,uint256)",
                "6xvaloper1t3p2vzd7w036ahxf4kefsc9sn24pvlqpmk79jh",
                20000 * 1e18
            )
        );
        require(success, "Transaction failed");

        // Log the success message
        console.log("Transfer success!");
        console.log(string(result));
        vm.stopBroadcast();
    }
}

contract undelegateScript is Script {
    address ownerAddress;
    uint64 currentNonce;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        address contractAddress = STAKING_PRECOMPILE_ADDRESS;
        // Execute the transaction
        (bool success, bytes memory result) = contractAddress.call(
            abi.encodeWithSignature(
                "undelegate(string,uint256)",
                "6xvaloper1t3p2vzd7w036ahxf4kefsc9sn24pvlqpmk79jh",
                20000 * 1e18
            )
        );
        require(success, "Transaction failed");

        // Log the success message
        console.log("Transfer success!");
        console.log(string(result));
        vm.stopBroadcast();
    }
}
