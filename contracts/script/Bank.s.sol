// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {IBank} from "../src/IBank.sol";

contract BankScript is Script {
    address ownerAddress;
    uint64 currentNonce;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        IBank bank = IBank(payable(0x0000000000000000000000000000000000001001));
        uint256 myBalance = bank.balance(ownerAddress, "asix");
        console.log(myBalance);
        vm.stopBroadcast();
    }

    function nonceUp() public {
        vm.setNonce(ownerAddress, currentNonce + uint64(1));
        currentNonce++;
    }
}
