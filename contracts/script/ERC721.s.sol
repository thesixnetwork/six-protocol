// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {MyNFT} from "../src/MyNFT.sol";
import {ERC721Factory} from "../src/Factory/ERC721Factory.sol";

contract DeployERC721WithFactory is Script {
    uint64 roundFloor = 50;
    uint256 minted = 0;
    uint64 nftNumber = 50;
    address ownerAddress;
    uint64 currentNonce;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        address factoryAddress = 0x3753C81072A56072840990D3D02f354Efb7425A3;
        ERC721Factory factory = ERC721Factory(factoryAddress);

        uint256 salt = 69;

        address erc721Address = factory.deployERC721(salt);
        nonceUp(ownerAddress);

        console.log("ERC721 deployed at:", erc721Address);

        MyNFT nft = MyNFT(payable(erc721Address));

        nft.setPreMinteeAddress(ownerAddress);
        nonceUp(ownerAddress);

        nft.setLimitedEditionSize(nftNumber);
        nonceUp(ownerAddress);

        uint256 round = nftNumber / roundFloor;
        uint256 remain = nftNumber % roundFloor;

        if (nftNumber >= roundFloor) {
            for (uint256 i = 0; i < round; i++) {
                nft.preMint(roundFloor);
                minted += roundFloor;
                nonceUp(ownerAddress);
            }
        }

        if (remain > 0 && minted < nftNumber && remain < roundFloor) {
            nft.preMint(remain);
            minted += remain;
            nonceUp(ownerAddress);
        }


        vm.stopBroadcast();
    }

        function nonceUp(address signer) public {
        vm.setNonce(signer, currentNonce + uint64(1));
        currentNonce++;
    }
}

contract DeployScript is Script {
    address ownerAddress;
    uint64 currentNonce;
    uint64 totalToken;
    uint64 roundFloor = 50;
    uint256 minted = 0;
    uint64 nftNumber = 50;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
        currentNonce = vm.getNonce(ownerAddress);
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        // MAIN
        MyNFT nft = new MyNFT();
        nonceUp(ownerAddress);
        address nftContractAddress = address(nft);

        nft.setPreMinteeAddress(ownerAddress);
        nonceUp(ownerAddress);

        nft.setLimitedEditionSize(nftNumber);
        nonceUp(ownerAddress);

        uint256 round = nftNumber / roundFloor;
        uint256 remain = nftNumber % roundFloor;

        if (nftNumber >= roundFloor) {
            for (uint256 i = 0; i < round; i++) {
                nft.preMint(roundFloor);
                minted += roundFloor;
                nonceUp(ownerAddress);
            }
        }

        if (remain > 0 && minted < nftNumber && remain < roundFloor) {
            nft.preMint(remain);
            minted += remain;
            nonceUp(ownerAddress);
        }
        console.log(address(nftContractAddress));
        vm.stopBroadcast();
    }

    function nonceUp(address signer) public {
        vm.setNonce(signer, currentNonce + uint64(1));
        currentNonce++;
    }
}

contract QueryTokenOwner is Script {
    address contractAddress;
    address ownerAddress;

    function setUp() public {
        ownerAddress = vm.envAddress("OWNER");
    }

    function run() external view {
        MyNFT nft = MyNFT(payable(0x3753C81072A56072840990D3D02f354Efb7425A3));

        console.log("Ownerof token 5", nft.totalSupply());
    }
}
