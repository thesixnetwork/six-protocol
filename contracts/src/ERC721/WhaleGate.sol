// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {ERC721} from "openzeppelin-contracts/token/ERC721/ERC721.sol";
import {ERC721Enumerable} from "openzeppelin-contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import {Ownable} from "openzeppelin-contracts/access/Ownable.sol";

contract WhaleGate is ERC721, ERC721Enumerable, Ownable {
  string public baseURI = "";
  uint256 private _nextTokenId;

  constructor(address initialOwner) ERC721("WHALEGATE", "WG") Ownable(initialOwner) {
    _nextTokenId = 1;
  }

  function setBaseURI(string memory _newUri) external {
    baseURI = _newUri;
  }

  function _baseURI() internal view override returns (string memory) {
    return baseURI;
  }

  function safeMint(address to) public onlyOwner returns (uint256) {
    uint256 tokenId = _nextTokenId++;
    _safeMint(to, tokenId);
    return tokenId;
  }

  // The following functions are overrides required by Solidity.

  function _update(
    address to,
    uint256 tokenId,
    address auth
  ) internal override(ERC721, ERC721Enumerable) returns (address) {
    return super._update(to, tokenId, auth);
  }

  function _increaseBalance(address account, uint128 value) internal override(ERC721, ERC721Enumerable) {
    super._increaseBalance(account, value);
  }

  function supportsInterface(bytes4 interfaceId) public view override(ERC721, ERC721Enumerable) returns (bool) {
    return super.supportsInterface(interfaceId);
  }
}
