// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

address constant DISTR_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000001007;

IDistr constant DISTR_CONTRACT = IDistr(DISTR_PRECOMPILE_ADDRESS);

interface IDistr {
    function setWithdrawAddress(
        string memory withdrawAddress
    ) external view returns (bool success);

    function withdrawRewards(
        string memory validatorAddress
    ) external view returns (bool success);

    function rewards(
        string memory validatorAddress,
        string memory delegatorAddress
    ) external view returns (Reward calldata reward);

    function allRewards(
        string memory delegatorAddress
    ) external view returns (Rewards calldata rewards);

    struct Coin {
        uint256 amount;
        string denom;
    }

    struct Reward {
        Coin[] coins;
        string validatorAddress;
    }

    struct Rewards {
        Reward[] reward;
        Coin[] total;
    }
}
