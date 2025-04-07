// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {IStaking, STAKING_PRECOMPILE_ADDRESS} from "../src/precompiles/IStaking.sol";

contract StakeScript is Script {
    address ownerAddress;
    uint64 currentNonce;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        address contractAddress = STAKING_PRECOMPILE_ADDRESS;
        // Using the interface approach (recommended)
        IStaking stakeContract = IStaking(contractAddress);
        IStaking.Delegation memory delegationResult = stakeContract.delegation(
            "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv",
            "6xvaloper1t3p2vzd7w036ahxf4kefsc9sn24pvlqpmk79jh"
        );

        // Display results using the interface call
        console.log("=== Delegation Information ===");
        console.log("Amount:", delegationResult.balance.amount);
        console.log("Denom:", delegationResult.balance.denom);
        console.log(
            "Delegator:",
            delegationResult.delegation.delegator_address
        );
        console.log(
            "Validator:",
            delegationResult.delegation.validator_address
        );
        console.log("Shares:", delegationResult.delegation.shares);
        console.log("Decimals:", delegationResult.delegation.decimals);

        vm.stopBroadcast();
    }
}

contract SendScript is Script {
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
                "send(address,address,string,uint256)",
                ownerAddress,
                0xd907f36f7D83344057a619b6D83A45B3288c3c21,
                "asix",
                2 * 1e18
            )
        );
        require(success, "Transaction failed");

        // Log the success message
        console.log("Transfer success!");
        console.log(string(result));
        vm.stopBroadcast();
    }
}
