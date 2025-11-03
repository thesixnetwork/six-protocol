// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "openzeppelin-contracts/access/AccessControl.sol";
import "./ERC20Impl.sol";

contract SimpleTokenFactory is AccessControl {
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant TOKEN_CREATOR_ROLE = keccak256("TOKEN_CREATOR_ROLE");

    address public lastCreatedToken;

    // create a map of token addresses to validate that the token is createed by this factory
    mapping(address => bool) public isTokenCreated;

    constructor(address admin) {
        // _grantRole(DEFAULT_ADMIN_ROLE,admin);
        _setRoleAdmin(TOKEN_CREATOR_ROLE, ADMIN_ROLE);
        _grantRole(ADMIN_ROLE,admin);
    }

    function createToken(
        string memory _name,
        string memory _symbol,
        uint256 _supply,
        address _owner
    ) public onlyRole(TOKEN_CREATOR_ROLE) {
        ERC20Impl token = new ERC20Impl(_name, _symbol, _supply, _owner);
        lastCreatedToken = address(token);
        isTokenCreated[address(token)] = true;
    }
}
