// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {INFTMNGR, NFTMNGR_PRECOMPILE_ADDRESS} from "../src/precompiles/INFTManager.sol";

contract ToggleAction is Script {
    address ownerAddress;
    string nftSchema;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        nftSchema = vm.envString("NFT_SCHEMA");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        string memory actionName = "claim_friend_reach_20";
        bool disable = true;

        bytes memory data = abi.encodeWithSignature(
            "toggleAction(string,string,bool)",
            nftSchema,
            actionName,
            disable
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
