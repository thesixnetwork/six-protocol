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
  address factoryAddress;

  function setUp() public {
    ownerAddress = vm.envAddress("OWNER");
    currentNonce = vm.getNonce(ownerAddress);
    string memory factoryContractInfoPath = "./broadcast/FactoryDeploy.s.sol/666/run-latest.json";
    string memory factoryContractInfo = vm.readFile(factoryContractInfoPath);
    bytes memory factoryJsonParsed = vm.parseJson(factoryContractInfo, ".transactions[0].contractAddress");

    factoryAddress = abi.decode(factoryJsonParsed, (address));
  }

  function run() external {
    uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
    vm.startBroadcast(deployerPrivateKey);

    ERC721Factory factory = ERC721Factory(factoryAddress);

    uint256 salt = 69;

    address membershipNFTAddress = factory.deployERC721(salt, "MEMBERSHIP", "MEM");
    nonceUp(ownerAddress);

    console.log("MEMBERSHIP contract is :", membershipNFTAddress);

    MyNFT membershipNFT = MyNFT(payable(membershipNFTAddress));

    membershipNFT.setPreMinteeAddress(ownerAddress);
    nonceUp(ownerAddress);

    membershipNFT.setLimitedEditionSize(nftNumber);
    nonceUp(ownerAddress);

    roundFloor = 50;
    minted = 0;
    nftNumber = 50;

    uint256 round = nftNumber / roundFloor;
    uint256 remain = nftNumber % roundFloor;

    if (nftNumber >= roundFloor) {
      for (uint256 i = 0; i < round; i++) {
        membershipNFT.preMint(roundFloor);
        minted += roundFloor;
        nonceUp(ownerAddress);
      }
    }

    if (remain > 0 && minted < nftNumber && remain < roundFloor) {
      membershipNFT.preMint(remain);
      minted += remain;
      nonceUp(ownerAddress);
    }

    salt = 96;

    address divineNFTAddress = factory.deployERC721(salt, "DIVINE_ELITE", "DIV");
    nonceUp(ownerAddress);

    console.log("DIVINE contract is :", divineNFTAddress);

    MyNFT divineNFT = MyNFT(payable(divineNFTAddress));

    divineNFT.setPreMinteeAddress(ownerAddress);
    nonceUp(ownerAddress);

    divineNFT.setLimitedEditionSize(nftNumber);
    nonceUp(ownerAddress);

    roundFloor = 50;
    minted = 0;
    nftNumber = 50;

    round = nftNumber / roundFloor;
    remain = nftNumber % roundFloor;

    if (nftNumber >= roundFloor) {
      for (uint256 i = 0; i < round; i++) {
        divineNFT.preMint(roundFloor);
        minted += roundFloor;
        nonceUp(ownerAddress);
      }
    }

    if (remain > 0 && minted < nftNumber && remain < roundFloor) {
      divineNFT.preMint(remain);
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

    // MEMBERSHIP
    MyNFT membershipNFT = new MyNFT("MEM", "MEMBERSHIP", ownerAddress);
    nonceUp(ownerAddress);
    address membershipNFTAddress = address(membershipNFT);
    console.log("membershipt address: ", membershipNFTAddress);

    membershipNFT.setPreMinteeAddress(ownerAddress);
    nonceUp(ownerAddress);

    membershipNFT.setLimitedEditionSize(nftNumber);
    nonceUp(ownerAddress);

    uint256 round = nftNumber / roundFloor;
    uint256 remain = nftNumber % roundFloor;

    if (nftNumber >= roundFloor) {
      for (uint256 i = 0; i < round; i++) {
        membershipNFT.preMint(roundFloor);
        minted += roundFloor;
        nonceUp(ownerAddress);
      }
    }

    if (remain > 0 && minted < nftNumber && remain < roundFloor) {
      membershipNFT.preMint(remain);
      minted += remain;
      nonceUp(ownerAddress);
    }

    // DIVINE
    MyNFT divineNFT = new MyNFT("DIV", "DIVINEELITE", ownerAddress);
    nonceUp(ownerAddress);
    address divineNFTAddress = address(divineNFT);
    console.log("devine elite address : ", divineNFTAddress);

    divineNFT.setPreMinteeAddress(ownerAddress);
    nonceUp(ownerAddress);

    divineNFT.setLimitedEditionSize(nftNumber);
    nonceUp(ownerAddress);

    uint256 divround = nftNumber / roundFloor;
    uint256 divremain = nftNumber % roundFloor;

    if (nftNumber >= roundFloor) {
      for (uint256 i = 0; i < divround; i++) {
        divineNFT.preMint(roundFloor);
        minted += roundFloor;
        nonceUp(ownerAddress);
      }
    }

    if (divremain > 0 && minted < nftNumber && remain < roundFloor) {
      divineNFT.preMint(remain);
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

contract TransferToken is Script {
  address contractAdrress;
  address ownerAddress;
  address membershipNftContractAddress;

  function setUp() public {
    ownerAddress = vm.envAddress("OWNER");
    string memory nftContractInfoPath = "./broadcast/ERC721.s.sol/666/run-latest.json";
    string memory nftContractInfo = vm.readFile(nftContractInfoPath);
    bytes memory membershipNftJsonParsed = vm.parseJson(nftContractInfo, ".transactions[0].contractAddress");

    membershipNftContractAddress = abi.decode(membershipNftJsonParsed, (address));
  }

  function run() external {
    uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
    MyNFT nft = MyNFT(payable(membershipNftContractAddress));
    vm.startBroadcast(deployerPrivateKey);

    nft.transferFrom(ownerAddress, 0x3753C81072A56072840990D3D02f354Efb7425A3, 5);

    vm.stopBroadcast();
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
