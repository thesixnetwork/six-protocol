// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "openzeppelin-contracts/token/ERC20/extensions/ERC20Capped.sol";
import "openzeppelin-contracts/access/Ownable.sol";

contract ERC20Impl is ERC20Capped, Ownable {
    uint256 public totalAllocatedAmount;

    string[] public holdersList;

    address[] public assetUpdatersList;

    mapping(string => bool) public isHolders;

    mapping(string => uint256) public holderBalances;

    mapping(string => string) public holderExternalIds;

    mapping(address => bool) public isAssetUpdater;

    event AssetUpdaterAdded(address updater);
    event AssetUpdaterRemoved(address updater);
    event AllAssetUpdatersRevoked();
    event HolderAdded(string holder, uint256 amount, string externalId);
    event MultipleHoldersAdded(string[] holders, uint256[] amounts, string[] externalIds);
    event HolderUpdated(string holder, uint256 previousAmount, uint256 amount,string externalId);
    event MultipleHoldersUpdated(string[] holders, uint256[] previousAmounts, uint256[] amounts,string[] holdersExternalIds);

    error CallerNotAssetUpdater();
    error AssetUpdaterExists();
    error AssetUpdaterDoesNotExist();
    error HolderExists();
    error HolderDoesNotExist();
    error HoldersAndAmountLengthMismatch();
    error AllocationExceedsCap();

    constructor(
        string memory _name,
        string memory _symbol,
        uint256 _supply,
        address _owner
    ) ERC20(_name, _symbol) ERC20Capped(_supply) Ownable() {
        _mint(_owner, _supply);
        totalAllocatedAmount = 0;
    }

    modifier onlyAssetUpdater() {
        if(!isAssetUpdater[msg.sender]) {
            revert CallerNotAssetUpdater();
        }
        _;
    }

    function addAssetUpdater(address updater) external onlyOwner {
        if(!isAssetUpdater[updater]) {
            isAssetUpdater[updater] = true;
            assetUpdatersList.push(updater);

            emit AssetUpdaterAdded(updater);
        }
        else {
            revert AssetUpdaterExists();
        }
    }

    function removeAssetUpdater(address updater) external onlyOwner {
        bool found = false;
        for (uint256 i = 0; i < assetUpdatersList.length; ++i) {
            if (assetUpdatersList[i] == updater) {
                assetUpdatersList[i] = assetUpdatersList[
                    assetUpdatersList.length - 1
                ];
                assetUpdatersList.pop();
                found = true;
                break;
            }
        }
        if (found) {
            isAssetUpdater[updater] = false;
            emit AssetUpdaterRemoved(updater);
        }
        else {
            revert AssetUpdaterDoesNotExist();
        }

    }

    function revokeAllAssetUpdaters() external onlyOwner {
        for (uint256 i = 0; i < assetUpdatersList.length; ++i) {
            isAssetUpdater[assetUpdatersList[i]] = false;
        }
        delete assetUpdatersList;
        emit AllAssetUpdatersRevoked();
    }

    function burn(uint256 amount) external {
        _burn(msg.sender, amount);
    }

    function addHolder(
        string calldata holder,
        uint256 amount,
        string calldata externalId
    ) external onlyAssetUpdater {
        if(isHolders[holder]) {
            revert HolderExists();
        }
        isHolders[holder] = true;
        holderBalances[holder] = amount;
        holderExternalIds[holder] = externalId;
        totalAllocatedAmount += amount;
        holdersList.push(holder);

        if (totalAllocatedAmount > cap()) {
            revert AllocationExceedsCap();
        }

        emit HolderAdded(holder, amount, externalId);
    }

    function addMultipleHolders(
        string[] memory holders,
        uint256[] memory amounts,
        string[] memory externalIds
    ) external onlyAssetUpdater {
        if(holders.length != amounts.length) {
            revert HoldersAndAmountLengthMismatch();
        }

        uint256 localTotalAllocatedAmount = totalAllocatedAmount;
        for (uint256 i = 0; i < holders.length; ++i) {
            if(isHolders[holders[i]]) {
                revert HolderExists();
            }
            isHolders[holders[i]] = true;
            holderBalances[holders[i]] = amounts[i];
            holderExternalIds[holders[i]] = externalIds[i];
            localTotalAllocatedAmount += amounts[i];
            holdersList.push(holders[i]);
        }
        totalAllocatedAmount = localTotalAllocatedAmount;

        if (totalAllocatedAmount > cap()) {
            revert AllocationExceedsCap();
        }

        emit MultipleHoldersAdded(holders, amounts, externalIds);
    }

    function updateHolder(
        string calldata holder,
        uint256 amount,
        string calldata externalId
    ) external onlyAssetUpdater {
        if(!isHolders[holder]) {
            revert HolderDoesNotExist();
        }
        uint256 currentBalance = holderBalances[holder];
        holderBalances[holder] = amount;
        holderExternalIds[holder] = externalId;
        totalAllocatedAmount = totalAllocatedAmount - currentBalance + amount;

        if (totalAllocatedAmount > cap()) {
            revert AllocationExceedsCap();
        }

        emit HolderUpdated(holder, currentBalance, amount, externalId);
    }

    function updateMultipleHolders(
        string[] memory holders,
        uint256[] memory amounts,
        string[] memory externalIds
    ) external onlyAssetUpdater {
        if(holders.length != amounts.length) {
            revert HoldersAndAmountLengthMismatch();
        }

        uint256 localTotalAllocatedAmount = totalAllocatedAmount;

        uint256[] memory previousAmounts = new uint256[](holders.length);
        for (uint256 i = 0; i < holders.length; ++i) {
            if(!isHolders[holders[i]]) {
                revert HolderDoesNotExist();
            }
            uint256 currentBalance = holderBalances[holders[i]];
            previousAmounts[i] = currentBalance;
            holderBalances[holders[i]] = amounts[i];
            holderExternalIds[holders[i]] = externalIds[i];
            localTotalAllocatedAmount =
                localTotalAllocatedAmount -
                currentBalance +
                amounts[i];
        }
        totalAllocatedAmount = localTotalAllocatedAmount;

        if (totalAllocatedAmount > cap()) {
            revert AllocationExceedsCap();
        }
        
        emit MultipleHoldersUpdated(holders, previousAmounts, amounts, externalIds);
    }

    function getHolderBalance(
        string calldata holder
    ) external view returns (uint256) {
        return holderBalances[holder];
    }

    function getHolderExternalId(
        string calldata holder
    ) external view returns (string memory) {
        return holderExternalIds[holder];
    }

    function getTotalAllocatedAmount() external view returns (uint256) {
        return totalAllocatedAmount;
    }

    function getHoldersListPagination(
        uint256 offset,
        uint256 limit
    ) external view returns (string[] memory) {
        if(holdersList.length == 0) {
            return new string[](0);
        }

        if(offset >= holdersList.length || limit == 0) {
            return new string[](0);
        }

        uint256 end = offset + limit;
        if (end > holdersList.length) {
            end = holdersList.length;
        }

        string[] memory result = new string[](end - offset);
        for (uint256 i = offset; i < end; i++) {
            result[i - offset] = holdersList[i];
        }

        return result;
    }

    function getTotalHolders() external view returns (uint256) {
        return holdersList.length;
    }

    function getAssetUpdatersListPagination(
        uint256 offset,
        uint256 limit
    ) external view returns (address[] memory) {
        if(assetUpdatersList.length == 0) {
            return new address[](0);
        }

        if(offset >= assetUpdatersList.length || limit == 0) {
            return new address[](0);
        }

        uint256 end = offset + limit;
        if (end > assetUpdatersList.length) {
            end = assetUpdatersList.length;
        }

        address[] memory result = new address[](end - offset);
        for (uint256 i = offset; i < end; i++) {
            result[i - offset] = assetUpdatersList[i];
        }

        return result;
    }

    function getTotalAssetUpdaters() external view returns (uint256) {
        return assetUpdatersList.length;
    }
}
