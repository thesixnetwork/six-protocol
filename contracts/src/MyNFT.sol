// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "./ERC721A/ERC721A.sol";
import {ERC721AQueryable} from "./extensions/ERC721AQueryable.sol";
import {Ownable} from "openzeppelin-contracts/access/Ownable.sol";
import {ReentrancyGuard} from "openzeppelin-contracts/security/ReentrancyGuard.sol";

contract MyNFT is ERC721A, Ownable, ERC721AQueryable, ReentrancyGuard {
    event WithdrawMoney(uint256 indexed blocktime, uint256 indexed amount, address indexed sender);

    address public preMinteeAddress;
    uint256 public limitedEditionSize;
    string private _baseTokenURI;

    constructor(
      string memory _name,
      string memory _symbol
    ) ERC721A(_name, _symbol) {}

    function _startTokenId() internal view virtual override returns (uint256) {
        return 1;
    }

    function setLimitedEditionSize(uint256 _limitedEditionSize) external onlyOwner {
        limitedEditionSize = _limitedEditionSize;
    }

    function setPreMinteeAddress(address _preMinteeAddress) external onlyOwner {
        require(_preMinteeAddress != address(0), "The address should not be 0.");
        preMinteeAddress = _preMinteeAddress;
    }

    function preMint(uint256 quantity) external onlyOwner {
        require(totalSupply() + quantity <= limitedEditionSize, "Too many already minted");
        _safeMint(preMinteeAddress, quantity);
    }

    function _baseURI() internal view virtual override returns (string memory) {
        return _baseTokenURI;
    }

    function setBaseURI(string calldata baseURI) external onlyOwner {
        _baseTokenURI = baseURI;
    }

    function withdrawMoney() external nonReentrant onlyOwner {
        require(address(this).balance > 0, "No balance to withdraw.");
        emit WithdrawMoney(block.timestamp, address(this).balance, msg.sender);
        payable(msg.sender).transfer(address(this).balance);
    }

    function numberMinted(address _owner) public view returns (uint256) {
        return _numberMinted(_owner);
    }

    function getOwnershipData(uint256 tokenId) external view returns (TokenOwnership memory) {
        return _ownershipOf(tokenId);
    }

    receive() external payable {}
}
