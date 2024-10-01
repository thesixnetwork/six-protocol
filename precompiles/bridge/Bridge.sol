// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

address constant BRIDGE_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000001069;

IBridge constant BRIDGE_CONTRACT = IBridge(BRIDGE_PRECOMPILE_ADDRESS);

interface IBridge {
    function transferToCosmos(string memory dst, uint256 amount) external returns (bool success);
}