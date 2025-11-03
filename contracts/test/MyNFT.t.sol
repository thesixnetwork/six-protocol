// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Test.sol";

import "../src/MyNFT.sol";

contract MyNFTTest is Test {
    MyNFT public myNFT;

    function setUp() public {
        myNFT = new MyNFT("TEST","TEST", address(0xDEE));
        myNFT.setPreMinteeAddress(address(0xDEE));
        myNFT.setLimitedEditionSize(260);
        myNFT.preMint(260);
    }

    function testMint() public {
        myNFT.setLimitedEditionSize(myNFT.limitedEditionSize() + 40);
        myNFT.preMint(40);
        assertEq(myNFT.balanceOf(address(0xDEE)), 300);
        assertEq(myNFT.ownerOf(1), address(0xDEE));
    }

    function testRevertSetPreMintFromNotOwner() public {
        vm.startPrank(address(1));
        vm.expectRevert("Ownable: caller is not the owner");
        myNFT.setPreMinteeAddress(address(0xDEE));
        vm.stopPrank();
    }

    function testRevertPreMintOverLimit() public {
        vm.expectRevert("Too many already minted");
        myNFT.preMint(200);
    }
}
