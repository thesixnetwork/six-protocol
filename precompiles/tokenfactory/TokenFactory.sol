// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

address constant TOKENFACTORY_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000001069;

ITokenFactory constant BRIDGE_CONTRACT = ITokenFactory(TOKENFACTORY_PRECOMPILE_ADDRESS);

interface ITokenFactory {
  function transferToCosmos(string memory dst, uint256 amount) external returns (bool success);

  function unwrapStakeToken(uint256 amount) external returns (bool success);

  function changeDelegatorAddress(address oldAddr, address newAddr) external returns (bool success);
}
