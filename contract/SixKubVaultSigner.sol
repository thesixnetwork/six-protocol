// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SixKubVaultSigner {
    struct ActionSigner {
        uint256 created_epoch;
        uint256 expired_epoch;
    }

    struct SetSignerParams {
        address actor_address;
        uint256 expired_at;
    }

    struct BindedSignerList {
        SetSignerParams[] signers;
        uint64 actorCount;
    }

    mapping(address => BindedSignerList) public binedSigner;

    mapping(address => mapping(address => ActionSigner)) public action_signer;

    event EventTypeSetActionSigner(address actor,address owner,uint256 expired_at);
    event EventTypeRemoveActionSigner(address actor, address owner);

    function setActionSigner(
        address actor,
        uint256 expire_epoch
    ) external returns (bool pass) {
        if (msg.sender == address(0)) {
            return false;
        }
        address owner = msg.sender;
        ActionSigner memory signer = ActionSigner({
            created_epoch: block.timestamp,
            expired_epoch: expire_epoch
        });
        action_signer[actor][owner] = signer;

        // make stuct of SetSignerParams
        SetSignerParams memory binder = SetSignerParams({
            actor_address: actor,
            expired_at: expire_epoch
        });

        // if bined signer already exist, update expired_at
        if (binedSigner[owner].actorCount > 0) {
            for (uint256 i = 0; i < binedSigner[owner].signers.length; i++) {
                if (binedSigner[owner].signers[i].actor_address == actor) {
                    binedSigner[owner].signers[i].expired_at = expire_epoch;
                    break;
                }
            }
        } else {
            // add signer to binedSigner
            binedSigner[owner].signers.push(binder);
            binedSigner[owner].actorCount += 1;
        }
        emit EventTypeSetActionSigner(actor, owner, expire_epoch);
        return true;
    }

    function getBindedSignerList(
        address owner
    ) external view returns (BindedSignerList memory) {
        return binedSigner[owner];
    }

    function getActionSigner(
        address actor,
        address owner
    )
        external
        view
        returns (address _actor, address _owner, ActionSigner memory)
    {
        return (actor, owner, action_signer[actor][owner]);
    }

    function getMyActionSigner(
        address actor
    )
        external
        view
        returns (address _actor, address _owner, ActionSigner memory)
    {
        return (actor, msg.sender, action_signer[actor][msg.sender]);
    }

    function removeActionSigner(address actor) external returns (bool pass) {
        if (msg.sender == address(0)) {
            return false;
        }
        address owner = msg.sender;

        ActionSigner memory signer = ActionSigner({
            created_epoch: block.timestamp,
            expired_epoch: 0
        });
        action_signer[actor][owner] = signer;

        // if bined signer already exist, update expired_at
        if (binedSigner[owner].actorCount > 0) {
            for (uint256 i = 0; i < binedSigner[owner].signers.length; i++) {
                if (binedSigner[owner].signers[i].actor_address == actor) {
                    delete binedSigner[owner].signers[i];
                }
            }
            binedSigner[owner].actorCount -= 1;
        }
        emit EventTypeRemoveActionSigner(actor, owner);
        return true;
    }
}
