// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {IBank, BANK_PRECOMPILE_ADDRESS} from "../src/precompiles/IBank.sol";

contract BankScript is Script {
    address ownerAddress;
    uint64 currentNonce;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        address contractAddress = BANK_PRECOMPILE_ADDRESS;
        (bool success, bytes memory result) = contractAddress.call(
            abi.encodeWithSignature(
                "balance(address,string)",
                ownerAddress,
                "asix"
            )
        );

        require(success, "Transaction failed");

        // Log the success message
        // console.log("Transfer success!");
        console.log(string(result));
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
        address contractAddress = BANK_PRECOMPILE_ADDRESS;
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
