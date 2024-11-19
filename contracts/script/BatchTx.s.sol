// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {TransactionBatcher} from "../src/TransactionBatcher.sol";

contract BatchScript is Script {
    address ownerAddress;
    address nftContractAddress;
    address batchContractAddress;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        batchContractAddress = vm.envAddress("TX_BATCHER_CONTRACT");
        //batchContractAddress = 0xFB2dfC82205dBB138Cb64b9F8d2414369a3bdd2c;
        nftContractAddress = vm.envAddress("PATIENT_CONTRACT");
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
