// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

address constant NFTMNGR_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000055;

INftmngr constant NFTMNGR_CONTRACT = INftmngr(
    NFTMNGR_PRECOMPILE_ADDRESS
);

struct ActionParameter {
    uint256 name;
    string value;
}

interface INftmngr {
    // Transactions
    function actionByAdmin(
        string memory nftSchemaName,
        string memory tokenId,
        string memory actionName,
        string memory refId,
        ActionParameter[] memory parameters
    ) external returns (bool success);
}
