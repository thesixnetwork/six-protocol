// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/ERC20Impl.sol";

contract DeployERC20 is Script {
    address ownerAddress;
    uint64 currentNonce;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        ERC20Impl erc20 = new ERC20Impl("TestTokenD", "TTD", 1000000000000000000000000000, ownerAddress);
        console.log(address(erc20));
        vm.stopBroadcast();
    }

    function nonceUp() public {
        vm.setNonce(ownerAddress, currentNonce + uint64(1));
        currentNonce++;
    }
}
