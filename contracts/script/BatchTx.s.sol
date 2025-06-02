// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {TransactionBatcher} from "../src/TransactionBatcher.sol";

contract DeployBatch is Script {
    address ownerAddress;
    address nftContractAddress;
    address batchContractAddress;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");

        string
            memory txBatcherPath = "./broadcast/ActionRouter.s.sol/666/run-latest.json";
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

        TransactionBatcher batchEntry = new TransactionBatcher();

        batchContractAddress = address(batchEntry);

        console.log(address(batchContractAddress));
        vm.stopBroadcast();
    }
}

contract BatchScript is Script {
    address ownerAddress;
    address nftContractAddress;
    address batchContractAddress;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");

        string
            memory txBatcherPath = "./broadcast/ActionRouter.s.sol/666/run-latest.json";
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

        TransactionBatcher batchEntry = TransactionBatcher(
            payable(batchContractAddress)
        );

        // Prepare data for batch call
        address[] memory targets = new address[](5);
        uint[] memory values = new uint[](5);
        bytes[] memory datas = new bytes[](5);

        address[] memory recipients = new address[](5);
        for (uint160 i = 0; i < 5; i++) {
            recipients[i] = address(i + 2);
        }

        //recipients[0] = address(1);
        //recipients[1] = address(2);
        //recipients[2] = address(3);
        //recipients[3] = address(4);
        //recipients[4] = address(5);

        for (uint i = 0; i < 5; i++) {
            targets[i] = nftContractAddress;
            values[i] = 0;
            datas[i] = abi.encodeWithSignature(
                "safeMint(address,uint256)",
                recipients[i],
                i + 2
            );
        }

        // Call batchSend for minting
        batchEntry.batchSend(targets, values, datas);

        vm.stopBroadcast();
    }
}
