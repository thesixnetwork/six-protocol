// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

address constant NFTMNGR_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000001055;

INFTMNGR constant NFTMNGR_CONTRACT = INFTMNGR(NFTMNGR_PRECOMPILE_ADDRESS);

interface INFTMNGR {
    // Transactions
    function addAction(
        string memory nftSchemaName,
        string memory base64NewAction
    ) external returns (bool success);

    function addActionExecutor(
        string memory nftSchemaName,
        address newExecutor
    ) external returns (bool success);

    function removeActionExecutor(
        string memory nftSchemaName,
        address rmExecutor
    ) external returns (bool success);

    function addAttribute(
        string memory nftSchemaName,
        uint32 location,
        string memory base64NewAttribute
    ) external returns (bool success);

    function changeOrgOwner(
        string memory orgName,
        address newOrgAddress
    ) external returns (bool success);

    function changeSchemaOwner(
        string memory nftSchemaName,
        address newOrgAddress
    ) external returns (bool success);

    function createMetadata(
        string memory nftSchemaName,
        string memory tokenId,
        string memory base64NewMetadata
    ) external returns (bool success);

    function createSchema(
        string memory base64NewSchema
    ) external returns (bool success);

    function actionByAdmin(
        string memory nftSchemaName,
        string memory tokenId,
        string memory actionName,
        string memory refId,
        string memory jsonParam
    ) external returns (bool success);

    function resyncAttribute(
        string memory nftSchemaName,
        string memory tokenId
    ) external returns (bool success);

    function updateSchemaAttribute(
        string memory nftSchemaName,
        string memory base64NewAttribute
    ) external returns (bool success);

    function attributeOveride(
        string memory nftSchemaName,
        uint32 overidingType
    ) external returns (bool success);

    function setBaseURI(
        string memory nftSchemaName,
        string memory newBaseUri
    ) external returns (bool success);

    function setMetadataFormat(
        string memory nftSchemaName,
        string memory newFormat
    ) external returns (bool success);

    function setMintAuth(
        string memory nftSchemaName,
        uint32 authorizeTo
    ) external returns (bool success);

    function setOriginChain(
        string memory nftSchemaName,
        string memory newOriginChain
    ) external returns (bool success);

    // new contract string incase that origin on non-evm
    function setOriginContract(
        string memory nftSchemaName,
        string memory newOriginContract
    ) external returns (bool success);

    function setUriRetreival(
        string memory nftSchemaName,
        string memory newUri
    ) external returns (bool success);

    function showAttribute(
        string memory nftSchemaName,
        bool toShow,
        string[] memory attirbuteNames
    ) external returns (bool success);

    function toggleAction(
        string memory nftSchemaName,
        string memory actionName,
        bool disable
    ) external returns (bool success);

    function updateAction(
        string memory nftSchemaName,
        string memory base64UpdateAction
    ) external returns (bool success);

    // QUERY
    function isActionExecutor(
        string memory nftSchemaName,
        address executor
    ) external view returns (bool isExecutor);

    function isSchemaOwner(
        string memory nftSchemaName,
        address ownerAddress
    ) external view returns (bool isSchemaOwner);

    function getAttributeValue(
        string memory nftSchemaName,
        string memory tokenId,
        string memory attributeName
    ) external view returns (string memory value);
}
