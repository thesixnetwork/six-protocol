// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/ActionRouter.sol";

contract ActionOne is Script {
    address ownerAddress;
    uint64 currentNonce;
    string virtualNftSchemaCode;
    address routerContractAddress;
    address membershipNftContractAddress;
    address divineEliteNftContractAddress;

    // action env
    string actionName;
    TokenIdMap[] tokenIdMapArray;
    string refId;
    string jsonParams;

    function setUp() public {
        string
            memory nftContractInfoPath = "./broadcast/ERC721.s.sol/666/run-latest.json";
        string memory nftContractInfo = vm.readFile(nftContractInfoPath);
        bytes memory membershipNftJsonParsed = vm.parseJson(
            nftContractInfo,
            ".transactions[0].contractAddress"
        );
        bytes memory divineNftJsonParsed = vm.parseJson(
            nftContractInfo,
            ".transactions[7].contractAddress"
        );
        membershipNftContractAddress = abi.decode(membershipNftJsonParsed, (address));
        divineEliteNftContractAddress = abi.decode(divineNftJsonParsed, (address));

        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
        virtualNftSchemaCode = "divineXmembership";

        tokenIdMapArray.push(
            TokenIdMap({
                nftSchemaName: "sixprotocol.membership",
                nftContractAddress: membershipNftContractAddress,
                tokenId: "5"
            })
        );

        tokenIdMapArray.push(
            TokenIdMap({
                nftSchemaName: "sixprotocol.divine_elite",
                nftContractAddress: divineEliteNftContractAddress,
                tokenId: "1"
            })
        );

        actionName = "bridge_4_to_2";
        refId = "";
        jsonParams = '[{"name":"amount","value": "10"}]';

        string
            memory routerContractInfoPath = "./broadcast/ActionRouter.s.sol/666/run-latest.json";
        string memory routerContractInfo = vm.readFile(routerContractInfoPath);
        bytes memory routerJsonParsed = vm.parseJson(
            routerContractInfo,
            ".transactions[0].contractAddress"
        );
        routerContractAddress = abi.decode(routerJsonParsed, (address));
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        Router actionRouter = Router(routerContractAddress);

        TokenIdMap[] memory _tokenIdMap = new TokenIdMap[](
            tokenIdMapArray.length
        );
        for (uint i = 0; i < tokenIdMapArray.length; i++) {
            _tokenIdMap[i] = tokenIdMapArray[i];
        }

        bool success = actionRouter.virtualAction(
            virtualNftSchemaCode,
            _tokenIdMap,
            actionName,
            refId,
            jsonParams
        );

        require(success, "Transaction failed. Check error message below:");

        // Log the success message
        console.log("Action executed successfully!");
        vm.stopBroadcast();
    }
}

contract ActionTwo is Script {
    address ownerAddress;
    uint64 currentNonce;
    string virtualNftSchemaCode;
    address routerContractAddress;
    address membershipNftContractAddress;
    address divineEliteNftContractAddress;

    // action env
    string actionName;
    TokenIdMap[] tokenIdMapArray;
    string refId;
    string jsonParams;

    function setUp() public {
        string
            memory nftContractInfoPath = "./broadcast/ERC721.s.sol/666/run-latest.json";
        string memory nftContractInfo = vm.readFile(nftContractInfoPath);
        bytes memory membershipNftJsonParsed = vm.parseJson(
            nftContractInfo,
            ".transactions[0].contractAddress"
        );
        bytes memory divineNftJsonParsed = vm.parseJson(
            nftContractInfo,
            ".transactions[7].contractAddress"
        );
        membershipNftContractAddress = abi.decode(membershipNftJsonParsed, (address));
        divineEliteNftContractAddress = abi.decode(divineNftJsonParsed, (address));

        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
        virtualNftSchemaCode = "divineXmembership";

        tokenIdMapArray.push(
            TokenIdMap({
                nftSchemaName: "sixprotocol.membership",
                nftContractAddress: membershipNftContractAddress,
                tokenId: "5"
            })
        );

        tokenIdMapArray.push(
            TokenIdMap({
                nftSchemaName: "sixprotocol.divine_elite",
                nftContractAddress: divineEliteNftContractAddress,
                tokenId: "1"
            })
        );

        actionName = "bridge_3_to_1";
        refId = "";
        jsonParams = '[]';

        string
            memory routerContractInfoPath = "./broadcast/ActionRouter.s.sol/666/run-latest.json";
        string memory routerContractInfo = vm.readFile(routerContractInfoPath);
        bytes memory routerJsonParsed = vm.parseJson(
            routerContractInfo,
            ".transactions[0].contractAddress"
        );
        routerContractAddress = abi.decode(routerJsonParsed, (address));
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        Router actionRouter = Router(routerContractAddress);

        TokenIdMap[] memory _tokenIdMap = new TokenIdMap[](
            tokenIdMapArray.length
        );
        for (uint i = 0; i < tokenIdMapArray.length; i++) {
            _tokenIdMap[i] = tokenIdMapArray[i];
        }

        bool success = actionRouter.virtualAction(
            virtualNftSchemaCode,
            _tokenIdMap,
            actionName,
            refId,
            jsonParams
        );

        require(success, "Transaction failed. Check error message below:");

        // Log the success message
        console.log("Action executed successfully!");
        vm.stopBroadcast();
    }
}

contract ActionThree is Script {
    address ownerAddress;
    uint64 currentNonce;
    string virtualNftSchemaCode;
    address routerContractAddress;
    address membershipNftContractAddress;
    address divineEliteNftContractAddress;

    // action env
    string actionName;
    TokenIdMap[] tokenIdMapArray;
    string refId;
    string jsonParams;

    function setUp() public {
        string
            memory nftContractInfoPath = "./broadcast/ERC721.s.sol/666/run-latest.json";
        string memory nftContractInfo = vm.readFile(nftContractInfoPath);
        bytes memory membershipNftJsonParsed = vm.parseJson(
            nftContractInfo,
            ".transactions[0].contractAddress"
        );
        bytes memory divineNftJsonParsed = vm.parseJson(
            nftContractInfo,
            ".transactions[7].contractAddress"
        );
        membershipNftContractAddress = abi.decode(membershipNftJsonParsed, (address));
        divineEliteNftContractAddress = abi.decode(divineNftJsonParsed, (address));

        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
        virtualNftSchemaCode = "divineXmembership";

        tokenIdMapArray.push(
            TokenIdMap({
                nftSchemaName: "sixprotocol.membership",
                nftContractAddress: membershipNftContractAddress,
                tokenId: "5"
            })
        );

        tokenIdMapArray.push(
            TokenIdMap({
                nftSchemaName: "sixprotocol.divine_elite",
                nftContractAddress: divineEliteNftContractAddress,
                tokenId: "1"
            })
        );

        actionName = "native_bridge";
        refId = "";
        jsonParams = '[{"name":"amount","value": "10"}]';

        string
            memory routerContractInfoPath = "./broadcast/ActionRouter.s.sol/666/run-latest.json";
        string memory routerContractInfo = vm.readFile(routerContractInfoPath);
        bytes memory routerJsonParsed = vm.parseJson(
            routerContractInfo,
            ".transactions[0].contractAddress"
        );
        routerContractAddress = abi.decode(routerJsonParsed, (address));
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        Router actionRouter = Router(routerContractAddress);

        TokenIdMap[] memory _tokenIdMap = new TokenIdMap[](
            tokenIdMapArray.length
        );
        for (uint i = 0; i < tokenIdMapArray.length; i++) {
            _tokenIdMap[i] = tokenIdMapArray[i];
        }

        bool success = actionRouter.virtualAction(
            virtualNftSchemaCode,
            _tokenIdMap,
            actionName,
            refId,
            jsonParams
        );

        require(success, "Transaction failed. Check error message below:");

        // Log the success message
        console.log("Action executed successfully!");
        vm.stopBroadcast();
    }
}
