import { newMockEvent } from "matchstick-as"
import { ethereum, Address, BigInt } from "@graphprotocol/graph-ts"
import {
  AllAssetUpdatersRevoked,
  Approval,
  AssetUpdaterAdded,
  AssetUpdaterRemoved,
  HolderAdded,
  HolderUpdated,
  MultipleHoldersAdded,
  MultipleHoldersUpdated,
  OwnershipTransferred,
  Transfer
} from "../generated/ERC20Imp/ERC20Imp"

export function createAllAssetUpdatersRevokedEvent(): AllAssetUpdatersRevoked {
  let allAssetUpdatersRevokedEvent =
    changetype<AllAssetUpdatersRevoked>(newMockEvent())

  allAssetUpdatersRevokedEvent.parameters = new Array()

  return allAssetUpdatersRevokedEvent
}

export function createApprovalEvent(
  owner: Address,
  spender: Address,
  value: BigInt
): Approval {
  let approvalEvent = changetype<Approval>(newMockEvent())

  approvalEvent.parameters = new Array()

  approvalEvent.parameters.push(
    new ethereum.EventParam("owner", ethereum.Value.fromAddress(owner))
  )
  approvalEvent.parameters.push(
    new ethereum.EventParam("spender", ethereum.Value.fromAddress(spender))
  )
  approvalEvent.parameters.push(
    new ethereum.EventParam("value", ethereum.Value.fromUnsignedBigInt(value))
  )

  return approvalEvent
}

export function createAssetUpdaterAddedEvent(
  updater: Address
): AssetUpdaterAdded {
  let assetUpdaterAddedEvent = changetype<AssetUpdaterAdded>(newMockEvent())

  assetUpdaterAddedEvent.parameters = new Array()

  assetUpdaterAddedEvent.parameters.push(
    new ethereum.EventParam("updater", ethereum.Value.fromAddress(updater))
  )

  return assetUpdaterAddedEvent
}

export function createAssetUpdaterRemovedEvent(
  updater: Address
): AssetUpdaterRemoved {
  let assetUpdaterRemovedEvent = changetype<AssetUpdaterRemoved>(newMockEvent())

  assetUpdaterRemovedEvent.parameters = new Array()

  assetUpdaterRemovedEvent.parameters.push(
    new ethereum.EventParam("updater", ethereum.Value.fromAddress(updater))
  )

  return assetUpdaterRemovedEvent
}

export function createHolderAddedEvent(
  holder: string,
  amount: BigInt,
  externalId: string
): HolderAdded {
  let holderAddedEvent = changetype<HolderAdded>(newMockEvent())

  holderAddedEvent.parameters = new Array()

  holderAddedEvent.parameters.push(
    new ethereum.EventParam("holder", ethereum.Value.fromString(holder))
  )
  holderAddedEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )
  holderAddedEvent.parameters.push(
    new ethereum.EventParam("externalId", ethereum.Value.fromString(externalId))
  )

  return holderAddedEvent
}

export function createHolderUpdatedEvent(
  holder: string,
  previousAmount: BigInt,
  amount: BigInt,
  externalId: string
): HolderUpdated {
  let holderUpdatedEvent = changetype<HolderUpdated>(newMockEvent())

  holderUpdatedEvent.parameters = new Array()

  holderUpdatedEvent.parameters.push(
    new ethereum.EventParam("holder", ethereum.Value.fromString(holder))
  )
  holderUpdatedEvent.parameters.push(
    new ethereum.EventParam(
      "previousAmount",
      ethereum.Value.fromUnsignedBigInt(previousAmount)
    )
  )
  holderUpdatedEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )
  holderUpdatedEvent.parameters.push(
    new ethereum.EventParam("externalId", ethereum.Value.fromString(externalId))
  )

  return holderUpdatedEvent
}

export function createMultipleHoldersAddedEvent(
  holders: Array<string>,
  amounts: Array<BigInt>,
  externalIds: Array<string>
): MultipleHoldersAdded {
  let multipleHoldersAddedEvent =
    changetype<MultipleHoldersAdded>(newMockEvent())

  multipleHoldersAddedEvent.parameters = new Array()

  multipleHoldersAddedEvent.parameters.push(
    new ethereum.EventParam("holders", ethereum.Value.fromStringArray(holders))
  )
  multipleHoldersAddedEvent.parameters.push(
    new ethereum.EventParam(
      "amounts",
      ethereum.Value.fromUnsignedBigIntArray(amounts)
    )
  )
  multipleHoldersAddedEvent.parameters.push(
    new ethereum.EventParam(
      "externalIds",
      ethereum.Value.fromStringArray(externalIds)
    )
  )

  return multipleHoldersAddedEvent
}

export function createMultipleHoldersUpdatedEvent(
  holders: Array<string>,
  previousAmounts: Array<BigInt>,
  amounts: Array<BigInt>,
  holdersExternalIds: Array<string>
): MultipleHoldersUpdated {
  let multipleHoldersUpdatedEvent =
    changetype<MultipleHoldersUpdated>(newMockEvent())

  multipleHoldersUpdatedEvent.parameters = new Array()

  multipleHoldersUpdatedEvent.parameters.push(
    new ethereum.EventParam("holders", ethereum.Value.fromStringArray(holders))
  )
  multipleHoldersUpdatedEvent.parameters.push(
    new ethereum.EventParam(
      "previousAmounts",
      ethereum.Value.fromUnsignedBigIntArray(previousAmounts)
    )
  )
  multipleHoldersUpdatedEvent.parameters.push(
    new ethereum.EventParam(
      "amounts",
      ethereum.Value.fromUnsignedBigIntArray(amounts)
    )
  )
  multipleHoldersUpdatedEvent.parameters.push(
    new ethereum.EventParam(
      "holdersExternalIds",
      ethereum.Value.fromStringArray(holdersExternalIds)
    )
  )

  return multipleHoldersUpdatedEvent
}

export function createOwnershipTransferredEvent(
  previousOwner: Address,
  newOwner: Address
): OwnershipTransferred {
  let ownershipTransferredEvent =
    changetype<OwnershipTransferred>(newMockEvent())

  ownershipTransferredEvent.parameters = new Array()

  ownershipTransferredEvent.parameters.push(
    new ethereum.EventParam(
      "previousOwner",
      ethereum.Value.fromAddress(previousOwner)
    )
  )
  ownershipTransferredEvent.parameters.push(
    new ethereum.EventParam("newOwner", ethereum.Value.fromAddress(newOwner))
  )

  return ownershipTransferredEvent
}

export function createTransferEvent(
  from: Address,
  to: Address,
  value: BigInt
): Transfer {
  let transferEvent = changetype<Transfer>(newMockEvent())

  transferEvent.parameters = new Array()

  transferEvent.parameters.push(
    new ethereum.EventParam("from", ethereum.Value.fromAddress(from))
  )
  transferEvent.parameters.push(
    new ethereum.EventParam("to", ethereum.Value.fromAddress(to))
  )
  transferEvent.parameters.push(
    new ethereum.EventParam("value", ethereum.Value.fromUnsignedBigInt(value))
  )

  return transferEvent
}
