// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {ITokenFactory, TOKENFACTORY_PRECOMPILE_ADDRESS} from "../src/precompiles/ITokenFactory.sol";

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
    address contractAddress = TOKENFACTORY_PRECOMPILE_ADDRESS;
    string memory destinationAddress = "6x1kch0sdjr5tuvjh0h3a55c6l5sr6m0phjeag9f2";
    uint256 amount = 1999 * 1e18;

    // Execute the transaction
    (bool success, ) = contractAddress.call(
      abi.encodeWithSignature("transferToCosmos(string,uint256)", destinationAddress, amount)
    );

    require(success, "Transaction failed");

    // Log the success message
    console.log("Transfer success!");

    vm.stopBroadcast();
  }
}

contract UnwrapStakeScript is Script {
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
    address contractAddress = TOKENFACTORY_PRECOMPILE_ADDRESS;
    uint256 amount = 1999 * 1e6;

    // Execute the transaction
    (bool success, ) = contractAddress.call(abi.encodeWithSignature("unwrapStakeToken(uint256)", amount));

    require(success, "Transaction failed");

    // Log the success message
    console.log("Transfer success!");

    vm.stopBroadcast();
  }
}

contract SendToCrossChainScript is Script {
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
    address contractAddress = TOKENFACTORY_PRECOMPILE_ADDRESS;
    string memory bridgeInAddress = "6x1kch0sdjr5tuvjh0h3a55c6l5sr6m0phjeag9f2";
    uint256 amount = 1999 * 1e18;
    string memory memo = "six-eth-bridge-001";
    string memory chain = "eth";

    // Execute the transaction
    (bool success, ) = contractAddress.call(abi.encodeWithSignature("transferToCrossChain(string,uint256,string,string)", bridgeInAddress,amount, memo, chain));

    require(success, "Transaction failed");

    // Log the success message
    console.log("Transfer success!");

    vm.stopBroadcast();
  }
}