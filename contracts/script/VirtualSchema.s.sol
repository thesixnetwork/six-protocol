// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {INFTMNGR, NFTMNGR_PRECOMPILE_ADDRESS} from "../src/INFTManager.sol";
import {TransactionBatcher} from "../src/TransactionBatcher.sol";

contract VoteScript is Script {
    error TransactionFailed();
    address ownerAddress;
    uint64 currentNonce;
    string nftSchema;
    uint32 option;
    address batchContractAddress;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
        nftSchema = vm.envString("NFT_SCHEMA");

        string
            memory txBatcherPath = "./broadcast/BatchTx.s.sol/666/run-latest.json";
        string memory txBatcherContractInfo = vm.readFile(txBatcherPath);
        bytes memory txBatcherJsonParsed = vm.parseJson(
            txBatcherContractInfo,
            ".transactions[0].contractAddress"
        );
        batchContractAddress = abi.decode(txBatcherJsonParsed, (address));
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        option = 2;

        bytes memory data = abi.encodeWithSignature(
            "voteVirtualSchema(string,string,uint32)",
            "1",
            "sixprotocol.membership",
            option
        );

        bool success = callKeeeper(data);
        require(success, "Transaction failed. Check error message below:");

        console.log("Action executed successfully!");
        vm.stopBroadcast();
    }

    function callKeeeper(bytes memory datas) public payable returns (bool) {
        (bool success, ) = NFTMNGR_PRECOMPILE_ADDRESS.call{value: 0, gas: 3_000_000}(datas);
        if (!success) revert("transaction failed");
        return success;
    }
}
