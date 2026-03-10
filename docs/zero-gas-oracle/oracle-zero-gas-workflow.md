# Oracle Zero-Gas Voting — Complete Workflow

## What Is It?

Oracle zero-gas voting allows **permissioned oracle nodes** to submit vote transactions on NFT oracle requests without paying gas fees or holding any $SIX balance. This removes a major operational dependency — oracle operators no longer need to fund wallets and manage gas budgets.

---

## Architecture Overview

```
Client (oracle binary)
        │
        ▼
  Tendermint P2P / RPC
        │
        ▼
  ┌─────────────────────────────────────────────────────────────────┐
  │                     Ante Handler Chain                          │
  │                                                                 │
  │  VoteAloneDecorator      ← reject bundled oracle msgs           │
  │  SetUpContextDecorator   ← set block gas limit                  │
  │  ValidateBasicDecorator                                         │
  │  GaslessDecorator        ← skip fee deduction for oracle msgs   │
  │    └─ IsTxGasless()      ← 3 checks: permission, request, dup   │
  │  SigVerificationDecorator                                       │
  │  IncrementSequenceDecorator                                     │
  └─────────────────────────────────────────────────────────────────┘
        │
        ▼
  Message Server (nftoracle module)
        │
        ├── SubmitMintResponse
        ├── SubmitActionResponse
        └── SubmitVerifyCollectionOwner
```

---

## Key Components

| Component | Package | Role |
|-----------|---------|------|
| `GaslessDecorator` | `app/ante` | Wraps fee decorators; skips them for gasless txs; sets infinite gas meter and oracle priority |
| `VoteAloneDecorator` | `app/ante` | Rejects any tx that bundles oracle votes with other messages |
| `IsTxGasless()` | `app/ante` | Performs the three gasless eligibility checks |
| `HasPermission()` | `x/nftadmin/keeper` | Checks the `"oracle"` permission list in the nftadmin store |
| `GetMintRequest` / `GetActionRequest` / `GetCollectionOwnerRequest` | `x/nftoracle/keeper` | Fetches the request and its current state |
| `CleanupOldVoteHeights` | `x/nftoracle/keeper` | Prunes the vote-height KV store in EndBlock |

---

## Detailed Flow: `SubmitMintResponse` (Gasless Path)

```
Oracle Node
  │
  │  MsgSubmitMintResponse {
  │    Creator:       "six1oracle...",
  │    MintRequestID: 42,
  │    Status:        "success",
  │    MetadataURI:   "ipfs://...",
  │  }
  ▼

1. VoteAloneDecorator
   ├─ tx.GetMsgs() → [MsgSubmitMintResponse]  (len == 1)
   └─ ✅ PASS (single oracle msg, no bundling)

2. SetUpContextDecorator
   └─ sets block gas limit context

3. ValidateBasicDecorator
   └─ checks msg.ValidateBasic()

4. GaslessDecorator.AnteHandle()
   │
   ├─ IsTxGasless(tx, ctx, oracleKeeper, adminKeeper)
   │    │
   │    ├─ msg type switch → oracleVoteIsGasless()
   │    │
   │    ├─ Check 1: Oracle Permission
   │    │    adminKeeper.HasPermission(ctx, "oracle", oracleAddr)
   │    │    ├─ GetAuthorization() → auth.Permissions["oracle"]
   │    │    └─ addr in addresses list?
   │    │         ✅ YES → continue
   │    │         ❌ NO  → ErrNoOraclePermission (tx rejected, pays gas)
   │    │
   │    ├─ Check 2: Request Exists and is PENDING
   │    │    oracleKeeper.GetMintRequest(ctx, 42)
   │    │         ✅ found + PENDING → continue
   │    │         ❌ not found       → ErrMintRequestNotFound (tx rejected, pays gas)
   │    │         ❌ not PENDING     → ErrMintRequestNotPending (tx rejected, pays gas)
   │    │
   │    └─ Check 3: Duplicate Vote
   │         hasOracleAlreadyVoted(request.Confirmers, oracleAddr)
   │              ✅ not in list → return true (gasless)
   │              ❌ already in list → ErrOracleAlreadyVoted (tx rejected, pays gas)
   │
   ├─ isGasless == true:
   │    ├─ ctx.IsCheckTx():  set InfiniteGasMeter + Priority = MaxInt64-100
   │    ├─ !ctx.IsCheckTx(): set InfiniteGasMeter
   │    └─ SKIP wrapped fee decorators (fee deduction is skipped)
   │
   └─ isGasless == false:
        └─ RUN wrapped fee decorators (normal fee deduction)

5. SigVerificationDecorator
   └─ verifies oracle's signature

6. IncrementSequenceDecorator
   └─ increments oracle account sequence

7. Message Server: SubmitMintResponse
   ├─ Re-validates request exists and is PENDING
   ├─ Appends oracle to request.Confirmers
   ├─ Increments request.CurrentConfirm
   ├─ If CurrentConfirm >= RequiredConfirm:
   │    ├─ Execute NFT mint (via nftmngr keeper)
   │    └─ Set request.Status = SUCCESS_WITH_CONSENSUS
   └─ SetMintRequest(ctx, updatedRequest)
```

---

## Detailed Flow: `SubmitActionResponse`

Identical to MintResponse except:
- Uses `GetActionRequest(ctx, msg.ActionRequestID)`
- Returns `ErrActionRequestNotFound` / `ErrActionRequestNotPending`
- Message server executes the NFT action on success

---

## Detailed Flow: `SubmitVerifyCollectionOwner`

Identical to MintResponse except:
- Uses `GetCollectionOwnerRequest(ctx, msg.VerifyRequestID)`
- Returns `ErrVerifyRequestNotFound` / `ErrVerifyRequestNotPending`
- Message server marks the NFT schema collection as verified on success

---

## Gasless Eligibility Decision Table

| Condition | Result |
|-----------|--------|
| Empty tx | Not gasless |
| Multiple messages in tx | Not gasless |
| Non-oracle message type | Not gasless |
| Oracle address lacks `"oracle"` permission | Rejected with error |
| Request ID not found | Rejected with error |
| Request status is not `PENDING` | Rejected with error |
| Oracle already in `request.Confirmers` | Rejected with error |
| All checks pass | ✅ **Gasless** |

---

## Permission Setup

Oracle addresses must be granted the `"oracle"` permission in the `nftadmin` module. This is done via governance or the root admin:

```json
{
  "root_admin": "six1...",
  "permissions": [
    {
      "name": "oracle",
      "addresses": [
        "six1oracle1...",
        "six1oracle2...",
        "six1oracle3..."
      ]
    }
  ]
}
```

The ante handler calls `nftadmin.HasPermission(ctx, "oracle", oracleAddr)` on every oracle vote. If an oracle is removed from this list, all future votes from that address will be rejected (and will pay normal gas).

---

## EndBlock Cleanup

Every block, `CleanupOldVoteHeights` runs with `maxAge = 100`:

```
EndBlock(ctx)
  └─ CleanupOldVoteHeights(ctx, maxAge=100)
       ├─ cutoff = ctx.BlockHeight() - 100
       ├─ Iterate all OracleLastVoteHeightPrefix entries
       └─ Delete entries where height < cutoff
```

This prevents the `OracleLastVoteHeight` KV sub-store from growing without bound. The vote-height store is only used for the `SetOracleLastVoteHeight` / `GetOracleLastVoteHeight` / `DeleteOracleLastVoteHeight` API (available for future spam-prevention features if needed).

---

## Error Reference

| Error | Code | Meaning |
|-------|------|---------|
| `ErrNoOraclePermission` | 3 | Sender is not in the oracle permission list |
| `ErrMintRequestNotFound` | 4 | No MintRequest with this ID exists |
| `ErrMintRequestNotPending` | 5 | MintRequest is not in PENDING state |
| `ErrActionRequestNotFound` | 100 | No ActionOracleRequest with this ID exists |
| `ErrActionRequestNotPending` | 101 | ActionOracleRequest is not in PENDING state |
| `ErrVerifyRequestNotFound` | 300 | No CollectionOwnerRequest with this ID exists |
| `ErrVerifyRequestNotPending` | 301 | CollectionOwnerRequest is not in PENDING state |
| `ErrOracleAlreadyVoted` | 600 | This oracle address is already in `request.Confirmers` |
| `ErrInvalidGaslessTransaction` | 604 | Tx shape is invalid for a gasless oracle vote |

---

## Ante Handler Chain (Full Context)

The oracle gasless feature sits in the Cosmos ante handler chain (not the EVM chain):

```
newCosmosAnteHandler(options):
  NewVoteAloneDecorator()
  authante.NewSetUpContextDecorator()
  authante.NewValidateBasicDecorator(options.Cdc)
  authante.NewTxTimeoutHeightDecorator()
  ...
  NewGaslessDecorator(
      wrappedDecorators: [NewDeductFeeDecorator(...)],
      nftOracleKeeper:    options.NftOracleKeeper,
      nftAdminKeeper:     options.NftAdminKeeper,
  )
  authante.NewSigGasConsumeDecorator(...)
  authante.NewSigVerificationDecorator(...)
  authante.NewIncrementSequenceDecorator(...)
```

`GaslessDecorator` wraps only `DeductFeeDecorator`. All other decorators still run for oracle transactions (signature verification, sequence increment, etc.).

---

## Transaction Priority

Gasless oracle transactions receive priority `math.MaxInt64 - 100` in the mempool. This ensures they are included quickly and not displaced by high-fee non-oracle transactions during periods of high chain activity.

---

## What Is NOT Gasless

- Any tx with more than one message (even if both are oracle vote types)
- Any message type other than `MsgSubmitMintResponse`, `MsgSubmitActionResponse`, `MsgSubmitVerifyCollectionOwner`
- Oracle votes on already-completed or expired requests
- Votes from addresses without the `"oracle"` permission
- Duplicate votes (oracle already in `Confirmers`)
