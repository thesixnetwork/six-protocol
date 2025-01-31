// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;
pragma experimental ABIEncoderV2;

contract TransactionBatcher {
    function batchSend(
        address[] memory targets,
        uint256[] memory values,
        bytes[] memory datas
    ) public payable {
        for (uint256 i = 0; i < targets.length; i++) {
            (bool success, ) = targets[i].call{value: values[i]}(datas[i]);
            if (!success) revert("transaction failed");
        }
    }

    function singleSend(
        address target,
        uint256 value,
        bytes memory data
    ) public payable {
        (bool success, ) = target.call{value: value}(data);
        if (!success) revert("transaction failed");
    }
}