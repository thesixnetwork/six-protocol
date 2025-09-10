import {
  AllAssetUpdatersRevoked as AllAssetUpdatersRevokedEvent,
  Approval as ApprovalEvent,
  AssetUpdaterAdded as AssetUpdaterAddedEvent,
  AssetUpdaterRemoved as AssetUpdaterRemovedEvent,
  HolderAdded as HolderAddedEvent,
  HolderUpdated as HolderUpdatedEvent,
  MultipleHoldersAdded as MultipleHoldersAddedEvent,
  MultipleHoldersUpdated as MultipleHoldersUpdatedEvent,
  OwnershipTransferred as OwnershipTransferredEvent,
  Transfer as TransferEvent
} from "../generated/ERC20Imp/ERC20Imp"
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
} from "../generated/schema"

export function handleAllAssetUpdatersRevoked(
  event: AllAssetUpdatersRevokedEvent
): void {
  let entity = new AllAssetUpdatersRevoked(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleApproval(event: ApprovalEvent): void {
  let entity = new Approval(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.owner = event.params.owner
  entity.spender = event.params.spender
  entity.value = event.params.value

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleAssetUpdaterAdded(event: AssetUpdaterAddedEvent): void {
  let entity = new AssetUpdaterAdded(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.updater = event.params.updater

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleAssetUpdaterRemoved(
  event: AssetUpdaterRemovedEvent
): void {
  let entity = new AssetUpdaterRemoved(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.updater = event.params.updater

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleHolderAdded(event: HolderAddedEvent): void {
  let entity = new HolderAdded(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.holder = event.params.holder
  entity.amount = event.params.amount
  entity.externalId = event.params.externalId

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleHolderUpdated(event: HolderUpdatedEvent): void {
  let entity = new HolderUpdated(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.holder = event.params.holder
  entity.previousAmount = event.params.previousAmount
  entity.amount = event.params.amount
  entity.externalId = event.params.externalId

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleMultipleHoldersAdded(
  event: MultipleHoldersAddedEvent
): void {
  let entity = new MultipleHoldersAdded(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.holders = event.params.holders
  entity.amounts = event.params.amounts
  entity.externalIds = event.params.externalIds

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleMultipleHoldersUpdated(
  event: MultipleHoldersUpdatedEvent
): void {
  let entity = new MultipleHoldersUpdated(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.holders = event.params.holders
  entity.previousAmounts = event.params.previousAmounts
  entity.amounts = event.params.amounts
  entity.holdersExternalIds = event.params.holdersExternalIds

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleOwnershipTransferred(
  event: OwnershipTransferredEvent
): void {
  let entity = new OwnershipTransferred(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.previousOwner = event.params.previousOwner
  entity.newOwner = event.params.newOwner

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleTransfer(event: TransferEvent): void {
  let entity = new Transfer(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.from = event.params.from
  entity.to = event.params.to
  entity.value = event.params.value

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}
