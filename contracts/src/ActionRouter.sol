// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IERC721} from "openzeppelin-contracts/token/ERC721/IERC721.sol";
import {INFTMNGR, NFTMNGR_PRECOMPILE_ADDRESS} from "./INFTManager.sol";

struct TokenIdMap {
    string nftSchemaName;
    string tokenId;
}

contract Router {
    error NotNFTOwner();
    error ModuleRejected();
    error TransactionFailed();

    function actionByNftOwner(
        address nftContractAddress,
        string memory nftSchemaName,
        string memory tokenId,
        string memory actionName,
        string memory refId,
        string memory jsonParam
    ) public returns (bool success) {
        uint256 tokenIdNumeric = stringToUint(tokenId);
        address owner = IERC721(nftContractAddress).ownerOf(tokenIdNumeric);
        if (msg.sender != owner) {
            revert NotNFTOwner();
        }

        bytes memory data = abi.encodeWithSignature(
            "actionByAdmin(string,string,string,string,string)", nftSchemaName, tokenId, actionName, refId, jsonParam
        );

        bool done = actionSend(data);

        if (!done) {
            revert ModuleRejected();
        }

        // Emit an event or perform other logic here
        emit ActionPerformed(nftSchemaName, tokenId, actionName, refId, jsonParam);

        return true;
    }

    function virtualAction(
        address nftContractAddress,
        string memory nftSchemaName,
        TokenIdMap[] memory tokenIdMap,
        string memory actionName,
        string memory refId,
        string memory jsonParam
    ) public returns (bool success) {
        uint256 tokenIdNumeric = stringToUint(tokenIdMap[0].tokenId);
        address owner = IERC721(nftContractAddress).ownerOf(tokenIdNumeric);
        if (msg.sender != owner) {
            revert NotNFTOwner();
        }

        string memory tokenIdMapJsonString = "[";
        for (uint256 i = 0; i < tokenIdMap.length; i++) {
            tokenIdMapJsonString = string(abi.encodePacked(
                tokenIdMapJsonString, '{"nftSchemaName":"', tokenIdMap[i].nftSchemaName, '","tokenId":"', tokenIdMap[i].tokenId, '"}'
            ));
            if (i < tokenIdMap.length - 1) {
                tokenIdMapJsonString = string(abi.encodePacked(tokenIdMapJsonString, ","));
            }
        }

        bytes memory data = abi.encodeWithSignature(
            "virtualAction(string,string,string,string,string)", nftSchemaName, tokenIdMapJsonString, actionName, refId, jsonParam
        );

        bool done = actionSend(data);

        if (!done) {
            revert ModuleRejected();
        }

        // Emit an event or perform other logic here
        emit ActionPerformed(nftSchemaName, tokenIdMapJsonString, actionName, refId, jsonParam);

        return true;
    }

    function stringToUint(string memory s) internal pure returns (uint256 result) {
        bytes memory b = bytes(s);
        uint256 i;
        result = 0;
        for (i = 0; i < b.length; i++) {
            uint256 c = uint256(uint8(b[i]));
            if (c >= 48 && c <= 57) {
                result = result * 10 + (c - 48);
            }
        }
    }

    event ActionPerformed(
        string indexed nftSchemaName, string indexed tokenId, string actionName, string refId, string jsonParam
    );

    function actionSend(bytes memory datas) public payable returns (bool) {
        (bool success,) = NFTMNGR_PRECOMPILE_ADDRESS.call{value: 0}(datas);
        if (!success) {
            revert TransactionFailed();
        }

        return success;
    }
}
