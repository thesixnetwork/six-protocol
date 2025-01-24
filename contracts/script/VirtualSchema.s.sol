// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {INFTMNGR, NFTMNGR_PRECOMPILE_ADDRESS} from "../src/INFTManager.sol";

contract VoteScript is Script {
    address ownerAddress;
    uint64 currentNonce;
    string nftSchema;
    uint32 option;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
        nftSchema = vm.envString("NFT_SCHEMA");
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        option = 2;

        // bytes memory dataDivine = abi.encodeWithSignature(
        //     "voteVirtualSchema(string,string,uint32)",
        //     "1",
        //     "sixprotocol.divine_elite",
        //     option
        // );

        bytes memory dataMembership = abi.encodeWithSignature(
            "voteVirtualSchema(string,string,uint32)",
            "1",
            "sixprotocol.membership",
            option
        );

        bool success = callKeeper(dataMembership);
        require(success, "Transaction failed. Check error message below:");

        // callKeeper(dataMembership);
        // nonceUp(ownerAddress);

        // Log the success message
        console.log("Action executed successfully!");
        vm.stopBroadcast();
    }

    function callKeeper(bytes memory datas) public payable returns (bool) {
        (bool success, ) = NFTMNGR_PRECOMPILE_ADDRESS.call{value: 0}(datas);
        if (!success) revert("transaction failed");
        return success;
    }

    // function nonceUp(address signer) public {
    //     vm.setNonce(signer, currentNonce + uint64(1));
    //     currentNonce++;
    // }
}
