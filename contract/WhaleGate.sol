// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts@4.4.0/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts@4.4.0/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts@4.4.0/access/Ownable.sol";
import "@openzeppelin/contracts@4.4.0/utils/Counters.sol";

contract WhaleGate is ERC721, ERC721Enumerable, Ownable {
    using Counters for Counters.Counter;
    string public baseURI = "";

    Counters.Counter private _tokenIdCounter;

    constructor() ERC721("WHALEGATE", "WG") {
        _tokenIdCounter._value = 1;
    }

    function setBaseURI(string memory _newUri)external {
        baseURI = _newUri;
    }
    function _baseURI() internal view override returns (string memory) {
        return baseURI;
    }

    function safeMint(address to) public {
        uint256 tokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment();
        _safeMint(to, tokenId);
    }

    // The following functions are overrides required by Solidity.

    function _beforeTokenTransfer(address from, address to, uint256 tokenId)
        internal
        override(ERC721, ERC721Enumerable)
    {
        super._beforeTokenTransfer(from, to, tokenId);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, ERC721Enumerable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}