// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

address constant NFTMNGR_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000055;

INFTMNGR constant NFTMNGR_CONTRACT = INFTMNGR(
    NFTMNGR_PRECOMPILE_ADDRESS
);

interface INFTMNGR {
    // Transactions
    function actionByAdmin(
        string memory nftSchemaName,
        string memory tokenId,
        string memory actionName,
        string memory refId,
        string memory jsonParam
    ) external returns (bool success);
}