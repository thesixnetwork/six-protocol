// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

address constant BRIDGE_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000001069;

IBridge constant BRIDGE_CONTRACT = IBridge(BRIDGE_PRECOMPILE_ADDRESS);

interface IBridge {
    function transferToCosmos(string memory dst, uint256 amount) external returns (bool success);
}

contract BridgePrecompile {
    function transferToCosmos(string memory dst, uint256 amount) public returns (bool success) {
        require(amount > 0, "Amount must be greater than 0");

        // Potentially break down large transfers to avoid high gas fees
        uint256 maxTransfer = 10 ** 18; // Example maximum amount per transaction
        uint256 remainingAmount = amount;

        while (remainingAmount > 0) {
            uint256 transferAmount = remainingAmount > maxTransfer ? maxTransfer : remainingAmount;
            success = BRIDGE_CONTRACT.transferToCosmos(dst, transferAmount);
            require(success, "Transfer to Cosmos failed");
            remainingAmount -= transferAmount;
        }
    }
}
