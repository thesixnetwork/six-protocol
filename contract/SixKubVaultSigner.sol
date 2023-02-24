// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SixKubVaultSigner {
    struct ActionSigner{
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

    function setActionSigner(address actor, uint256 expire_epoch) external {
        ActionSigner memory signer = ActionSigner({
            created_epoch: block.timestamp,
            expired_epoch: expire_epoch
        });
        action_signer[actor][msg.sender] = signer;

        
        // make stuct of SetSignerParams
        SetSignerParams memory binder = SetSignerParams({
            actor_address: actor,
            expired_at: expire_epoch
        });

        // add signer to binedSigner
        binedSigner[msg.sender].signers.push(binder);
        binedSigner[msg.sender].actorCount += 1;
    }

    function getBindedSignerList(address actor) external view returns (BindedSignerList memory) {
        return binedSigner[actor];
    }

    function getActionSigner(address actor, address owner) external view returns (address _actor, address _owner, ActionSigner memory) {
        return (actor, msg.sender, action_signer[actor][owner]);
    }

    function getMyActionSigner(address actor) external view returns (address _actor, address _owner, ActionSigner memory) {
        return (actor, msg.sender, action_signer[actor][msg.sender]);
    }

    function removeActionSigner(address actor) external {
        ActionSigner memory signer = ActionSigner({
            created_epoch: block.timestamp,
            expired_epoch: 0
        });
        action_signer[actor][msg.sender] = signer;

        // add signer to binedSigner
        // loop and find index of signer to remove
        for (uint256 i = 0; i < binedSigner[msg.sender].signers.length; i++) {
            if (binedSigner[msg.sender].signers[i].actor_address == actor) {
                delete binedSigner[msg.sender].signers[i];
            }
        }
        binedSigner[msg.sender].actorCount -= 1;
    }


    // function verify(address actor, uint256 epoch, bytes memory signature) external view returns (bool) {
    //     binedSigner[actor].signers[1].actor_address;
    //     for (uint256 i = 0; i < binedSigner[actor].signers.length; i++) {
    //         if (binedSigner[actor].signers[i].actor_address == msg.sender) {
    //             break;
    //         }
    //         if (i == binedSigner[actor].signers.length - 1) {
    //             return false;
    //         }
    //     }
    //     if (action_signer[actor][msg.sender].expired_epoch < epoch) {
    //         return false;
    //     }
    //     bytes32 message = keccak256(abi.encodePacked(actor, msg.sender, epoch));
    //     return recoverSigner(message, signature) == msg.sender;
    // }

    // function recoverSigner(bytes32 message, bytes memory signature) public pure returns (address) {
    //     require(signature.length == 65, "invalid signature length");

    //     bytes32 r;
    //     bytes32 s;
    //     uint8 v;

    //     assembly {
    //         r := mload(add(signature, 0x20))
    //         s := mload(add(signature, 0x40))
    //         v := byte(0, mload(add(signature, 0x60)))
    //     }

    //     if (v < 27) {
    //         v += 27;
    //     }

    //     require(v == 27 || v == 28, "invalid signature v value");

    //     return ecrecover(message, v, r, s);
    // }
}