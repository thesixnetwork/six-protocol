// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {BridgePrecompile, IBridge, BRIDGE_PRECOMPILE_ADDRESS} from "../src/IBridge.sol";

contract SendToCosmosScript is Script {
    address ownerAddress;
    uint64 currentNonce;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        // Initialize variables
        address contractAddress = BRIDGE_PRECOMPILE_ADDRESS;
        string memory destinationAddress = "6x1kch0sdjr5tuvjh0h3a55c6l5sr6m0phjeag9f2";
        uint256 amount = 5 * 1e18;

        // Execute the transaction
        (bool success, bytes memory result) = contractAddress.call(
            abi.encodeWithSignature(
                "transferToCosmos(string,uint256)",
                destinationAddress,
                amount
            )
        );

        require(success, "Transaction failed");

        // Log the success message
        console.log("Transfer success!");
        console.log(string(result));

        vm.stopBroadcast();
    }
}
