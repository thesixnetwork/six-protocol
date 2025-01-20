// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

address constant NFTMNGR_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000001055;

INFTMNGR constant NFTMNGR_CONTRACT = INFTMNGR(NFTMNGR_PRECOMPILE_ADDRESS);

struct TokenIdMap {
    string NftSchemaName;
    string tokenId;
}

struct ActionParameter {
    string Name;
    string Value;
}

interface INFTMNGR {
    // Transactions
    function addAction(
        string memory nftSchemaCode,
        string memory base64NewAction
    ) external returns (bool success);

    function addActionExecutor(
        string memory nftSchemaCode,
        address newExecutor
    ) external returns (bool success);

    function removeActionExecutor(
        string memory nftSchemaCode,
        address rmExecutor
    ) external returns (bool success);

    function addAttribute(
        string memory nftSchemaCode,
        uint32 location,
        string memory base64NewAttribute
    ) external returns (bool success);

    function changeOrgOwner(
        string memory orgName,
        address newOrgAddress
    ) external returns (bool success);

    function changeSchemaOwner(
        string memory nftSchemaCode,
        address newOrgAddress
    ) external returns (bool success);

    function createMetadata(
        string memory nftSchemaCode,
        string memory tokenId,
        string memory base64NewMetadata
    ) external returns (bool success);

    function createSchema(
        string memory base64NewSchema
    ) external returns (bool success);

    function actionByAdmin(
        string memory nftSchemaCode,
        string memory tokenId,
        string memory actionName,
        string memory refId,
        string memory jsonParam
    ) external returns (bool success);

    function resyncAttribute(
        string memory nftSchemaCode,
        string memory tokenId
    ) external returns (bool success);

    function updateSchemaAttribute(
        string memory nftSchemaCode,
        string memory base64NewAttribute
    ) external returns (bool success);

    function attributeOveride(
        string memory nftSchemaCode,
        uint32 overidingType
    ) external returns (bool success);

    function setBaseURI(
        string memory nftSchemaCode,
        string memory newBaseUri
    ) external returns (bool success);

    function setMetadataFormat(
        string memory nftSchemaCode,
        string memory newFormat
    ) external returns (bool success);

    function setMintAuth(
        string memory nftSchemaCode,
        uint32 authorizeTo
    ) external returns (bool success);

    function setOriginChain(
        string memory nftSchemaCode,
        string memory newOriginChain
    ) external returns (bool success);

    // new contract string incase that origin on non-evm
    function setOriginContract(
        string memory nftSchemaCode,
        string memory newOriginContract
    ) external returns (bool success);

    function setUriRetreival(
        string memory nftSchemaCode,
        string memory newUri
    ) external returns (bool success);

    function showAttribute(
        string memory nftSchemaCode,
        bool toShow,
        string[] memory attirbuteNames
    ) external returns (bool success);

    function toggleAction(
        string memory nftSchemaCode,
        string memory actionName,
        bool disable
    ) external returns (bool success);

    function updateAction(
        string memory nftSchemaCode,
        string memory base64UpdateAction
    ) external returns (bool success);

    function virtualAction(
        string memory vitualSchemaName,
        TokenIdMap memory tokenMap,
        string memory actionName,
        string memory refId,
        ActionParameter memory parameters
    ) external returns (bool success);

    function voteVirtualSchema(
      string memory proposalId,
      string memory nftSchemaCode,
      uint32 option
    ) external returns (bool success);

    // QUERY
    function isActionExecutor(
        string memory nftSchemaCode,
        address executor
    ) external view returns (bool isExecutor);

    function isSchemaOwner(
        string memory nftSchemaCode,
        address ownerAddress
    ) external view returns (bool isSchemaOwner);

    function getAttributeValue(
        string memory nftSchemaCode,
        string memory tokenId,
        string memory attributeName
    ) external view returns (string memory value);
}
