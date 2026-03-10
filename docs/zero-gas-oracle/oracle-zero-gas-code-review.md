# Oracle Zero-Gas Voting — Code Review Report

**Branch:** Current development branch  
**Review Date:** 2026-03-10  
**Reviewed By:** GitHub Copilot  
**Files Reviewed:**
- `app/ante/gasless.go`
- `app/ante/cosmos.go`
- `app/ante/handler_options.go`
- `app/ante/gasless_test.go`
- `app/app.go` (`setAnteHandler`)
- `x/nftoracle/keeper/gasless_voting.go`
- `x/nftoracle/types/errors.go`

---

## Executive Summary

| Question | Answer |
|---|---|
| **Does this feature work?** | ⚠️ **Partially** — The logic is correct, but a critical wiring bug prevents the app from starting. |
| **Are there mistakes?** | 🔴 **Yes** — One critical bug + several medium/low issues listed below. |
| **Is it production-ready?** | ❌ **No** — Must fix the critical bug, add integration tests, and address the issues below first. |

---

## 1. Does the Feature Work?

### What is implemented correctly ✅

- **`GaslessDecorator`** (`app/ante/gasless.go`) — correctly wraps fee deduction. When a transaction is gasless, the wrapped `DeductFeeDecorator` is skipped and an infinite gas meter is applied.
- **`IsTxGasless()`** — correctly dispatches on the 3 oracle message types (`MsgSubmitMintResponse`, `MsgSubmitActionResponse`, `MsgSubmitVerifyCollectionOwner`).
- **`VoteAloneDecorator`** — correctly rejects oracle messages bundled with other messages.
- **Validation pipeline** — each `oracleXxx IsGasless()` function correctly checks oracle permission → request exists → request PENDING → no duplicate vote → spam prevention.
- **`gasless_voting.go`** — KV-store based vote-height tracking is implemented correctly.
- **Error codes** — new errors (600-604) are registered without conflicts.
- **`HandlerOptions.Validate()`** — correctly enforces that `NftOracleKeeper` and `NftAdminKeeper` are non-nil.

### What is broken 🔴

See **Critical Bug #1** below. The application will `panic` on startup because the two required keepers are never assigned into `HandlerOptions`.

---

## 2. Bugs & Mistakes

### 🔴 [CRITICAL] Bug #1 — Missing Keeper Injection in `setAnteHandler`

**File:** `app/app.go`, function `setAnteHandler`

The `HandlerOptions` struct has two required fields:

```go
// app/ante/handler_options.go
NftOracleKeeper *nftoraclekeeper.Keeper
NftAdminKeeper  *nftadminkeeper.Keeper
```

`Validate()` explicitly panics if they are `nil`:

```go
if options.NftOracleKeeper == nil {
    return errorsmod.Wrap(errortypes.ErrLogic, "nft oracle keeper is required for AnteHandler")
}
if options.NftAdminKeeper == nil {
    return errorsmod.Wrap(errortypes.ErrLogic, "nft admin keeper is required for AnteHandler")
}
```

However, in `setAnteHandler()` in `app/app.go`, **neither keeper is assigned**:

```go
// app/app.go — current (broken) code
options := ante.HandlerOptions{
    Cdc:                    app.appCodec,
    AccountKeeper:          app.AccountKeeper,
    BankKeeper:             app.BankKeeper,
    // ... other fields ...
    CircuitKeeper:          &app.CircuitBreakerKeeper,
    AllowUnorderedTx:       unsafeUnorderedTx,
    // ❌ NftOracleKeeper and NftAdminKeeper are MISSING here
}
if err := options.Validate(); err != nil {
    panic(err)  // <-- This WILL panic at startup
}
```

**Fix:**

```go
options := ante.HandlerOptions{
    // ...existing fields...
    NftOracleKeeper: &app.NftoracleKeeper,   // ← ADD THIS
    NftAdminKeeper:  &app.NftadminKeeper,    // ← ADD THIS
}
```

> ⚠️ **This bug prevents the node from starting.** The application will panic immediately on boot.

---

### 🟡 [MEDIUM] Bug #2 — Spam Prevention Blocks Legitimate Multi-Request Voting

**File:** `app/ante/gasless.go`, function `checkAndSetSpamPreventionCounter`

The current logic allows **only one oracle vote per block** across ALL request types:

```go
func checkAndSetSpamPreventionCounter(ctx sdk.Context, oracle sdk.AccAddress, oracleKeeper nftoraclekeeper.Keeper) error {
    lastVoteHeight := oracleKeeper.GetOracleLastVoteHeight(ctx, oracle)
    currentHeight := ctx.BlockHeight()

    if lastVoteHeight == currentHeight {
        return errorsmod.Wrap(nftoracletypes.ErrOracleSpamPrevention, "oracle already voted in this block")
    }
    oracleKeeper.SetOracleLastVoteHeight(ctx, oracle, currentHeight)
    return nil
}
```

**Problem:** An oracle cannot vote on two different requests (e.g., one mint request + one action request) in the same block, even though they are completely independent. This is unnecessarily restrictive and will cause valid oracle votes to be rejected, breaking oracle consensus when network activity is high.

**What Sei Protocol does:** Sei limits to one vote per **vote period** (multiple blocks) for the same oracle price feed, not across different request IDs.

**Fix options:**
- Track spam per `(oracle_address, request_id)` pair instead of per oracle address only, OR
- Remove the block-level spam prevention entirely and rely solely on the duplicate-vote check (`hasOracleAlreadyVoted`), which is already request-scoped and therefore sufficient.

---

### 🟡 [MEDIUM] Bug #3 — `SetOracleLastVoteHeight` Is Called During `CheckTx` (State-Modification Side Effect)

**File:** `app/ante/gasless.go`, function `checkAndSetSpamPreventionCounter`

The ante handler calls `oracleKeeper.SetOracleLastVoteHeight(ctx, oracle, currentHeight)` inside `IsTxGasless()`, which is called from `AnteHandle()`. In Cosmos SDK, `CheckTx` runs in a cached context that **is discarded**. But `SetOracleLastVoteHeight` writes to KV store even during `CheckTx`.

**Problem:**
1. During `CheckTx`, the state write is ultimately discarded — the spam-prevention counter in store is **not updated** for `CheckTx`.
2. During `DeliverTx` / `FinalizeBlock`, the write is committed. But the function is called during ante handling, **before** the transaction is known to succeed. If a later decorator rejects the transaction, the vote height is still written to state — **incorrectly blocking future valid votes**.

**Fix:** Move the spam-prevention counter write (`SetOracleLastVoteHeight`) to the **message server** (`SubmitMintResponse`, etc.), not the ante handler. The ante handler should only read the counter for validation. Alternatively, only write it on `DeliverTx` after all checks pass.

---

### 🟡 [MEDIUM] Bug #4 — `CleanupOldVoteHeights` Is Never Called

**File:** `x/nftoracle/keeper/gasless_voting.go`

The function `CleanupOldVoteHeights(ctx, maxAge)` is implemented but **never invoked** anywhere in the codebase — not in `BeginBlock`, `EndBlock`, or any other lifecycle hook.

```go
// Defined but never called
func (k Keeper) CleanupOldVoteHeights(ctx context.Context, maxAge int64) { ... }
```

The `nftoracle` module's `BeginBlock` and `EndBlock` are empty stubs:

```go
// x/nftoracle/module/module.go
func (am AppModule) BeginBlock(_ context.Context) error { return nil }
func (am AppModule) EndBlock(_ context.Context) error   { return nil }
```

**Impact:** The `OracleLastVoteHeightPrefix` store grows indefinitely — one entry per oracle, permanently. While not immediately critical (entries are small and get overwritten on the next vote), this wastes storage over time.

**Fix:** Call `CleanupOldVoteHeights` in `EndBlock` with an appropriate `maxAge` (e.g., 100 blocks):

```go
func (am AppModule) EndBlock(ctx context.Context) error {
    sdkCtx := sdk.UnwrapSDKContext(ctx)
    am.keeper.CleanupOldVoteHeights(ctx, 100)
    _ = sdkCtx
    return nil
}
```

---

### 🟡 [MEDIUM] Bug #5 — `NftmngrKeeper` Initialized with Zero-Value `NftadminKeeper`

**File:** `app/app.go`

```go
// Line ~760-790 in app.go
app.NftmngrKeeper = nftmngrmodulekeeper.NewKeeper(
    ...
    app.NftadminKeeper,   // ← This is a ZERO VALUE at this point!
    ...
)

// NftadminKeeper is only initialized AFTER NftmngrKeeper:
app.NftadminKeeper = nftadminmodulekeeper.NewKeeper(...)
app.NftoracleKeeper = nftoraclemodulekeeper.NewKeeper(
    ...
    app.NftadminKeeper,  // ← This is fine, initialized just above
    ...
)
```

`NftmngrKeeper` is constructed before `NftadminKeeper` is initialized, so it receives a zero-value keeper. This is a pre-existing bug but is relevant context: the gasless feature depends on `NftadminKeeper` being correctly set up.

**Fix:** Reorder keeper initialization so `NftadminKeeper` is created before `NftmngrKeeper`.

---

### 🟢 [LOW] Issue #6 — `IsOracleGaslessEnabled()` Is Always `true` (Not Configurable)

**File:** `x/nftoracle/keeper/gasless_voting.go`

```go
func (k Keeper) IsOracleGaslessEnabled(ctx context.Context) bool {
    return true  // Hardcoded, never reads from params
}
```

This function exists but is not used in the gasless logic at all. If it were wired in, it could serve as an on-chain toggle. As-is, it is dead code.

**Fix:** Either wire it into `IsTxGasless()` as a guard, or remove it.

---

### 🟢 [LOW] Issue #7 — `CollectionOwnerRequest` Uses Wrong Error Type

**File:** `app/ante/gasless.go`

```go
collectionRequest, found := oracleKeeper.GetCollectionOwnerRequest(ctx, msg.VerifyRequestID)
if !found {
    return false, errorsmod.Wrap(nftoracletypes.ErrCollectionOwnerRequestNotFound, ...)
}
if collectionRequest.Status != nftoracletypes.RequestStatus_PENDING {
    return false, errorsmod.Wrap(nftoracletypes.ErrCollectionOwnerRequestNotPending, ...)
}
```

The existing codebase already has `ErrVerifyRequestNotFound` (error code 300) and `ErrVerifyRequestNotPending` (error code 301) for collection-owner requests. The new errors `ErrCollectionOwnerRequestNotFound` (code 602) and `ErrCollectionOwnerRequestNotPending` (code 603) are duplicates with different codes, which can confuse clients.

**Fix:** Reuse the existing `ErrVerifyRequestNotFound` and `ErrVerifyRequestNotPending` error types.

---

### 🟢 [LOW] Issue #8 — Test Coverage Is Shallow

**File:** `app/ante/gasless_test.go`

The existing tests only cover:
- Empty transactions
- Multi-message rejection
- Non-oracle messages
- Decorator construction

**What is NOT tested:**
- A full end-to-end test where oracle permission is granted and a valid mint request is in PENDING state → `IsTxGasless()` returns `true`
- Spam prevention counter (the test manually calls `SetOracleLastVoteHeight` but never tests that the ante handler actually rejects a second vote in the same block)
- `oracleActionResponseIsGasless` and `oracleCollectionVerifyIsGasless` — zero coverage
- The `GaslessDecorator` setting an infinite gas meter for a truly gasless tx

---

## 3. Is It Production-Ready?

**No.** Below is a checklist:

| Item | Status |
|---|---|
| Critical keeper wiring bug fixed | ❌ Not fixed |
| Spam prevention semantics correct | ❌ Too aggressive (Bug #2) |
| State side-effect in ante handler | ❌ Wrong location (Bug #3) |
| Store cleanup wired | ❌ Never called (Bug #4) |
| Keeper init order correct | ❌ Pre-existing bug (Bug #5) |
| Integration tests with real keeper state | ❌ Missing |
| End-to-end test (node starts + oracle votes gaslessly) | ❌ Missing |
| `IsOracleGaslessEnabled()` wired or removed | ❌ Dead code |
| Duplicate error codes cleaned up | ❌ |

---

## 4. Recommendations & Advice

### Must Fix Before Merging

1. **Add the two missing keepers in `setAnteHandler`** (`NftOracleKeeper`, `NftAdminKeeper`). Without this the node cannot start. See Bug #1.

2. **Fix spam prevention to be per `(oracle, request_id)`** or remove the block-level rate limit entirely. The existing duplicate-vote check is sufficient. See Bug #2.

3. **Move `SetOracleLastVoteHeight` write out of the ante handler.** Ante handlers should not commit state for validation-only checks. Call it from the message server instead. See Bug #3.

### Should Fix Before Merging

4. **Wire `CleanupOldVoteHeights` into `EndBlock`** to prevent unbounded store growth. See Bug #4.

5. **Fix `NftmngrKeeper` initialization order** so it does not receive a zero-value `NftadminKeeper`. See Bug #5.

6. **Write integration tests** that cover a complete happy-path gasless vote (grant permission → create pending request → submit vote → verify no fee deducted).

### Nice to Have

7. Wire `IsOracleGaslessEnabled()` to module params so governance can toggle gasless voting on/off without a code upgrade.

8. Add `MsgSubmitSyncActionSigner` (if it also needs gasless treatment) — it is currently not handled.

9. Emit an event `EventTypeGaslessOracleVote` from the ante handler or message server for observability and monitoring.

10. Consider using a **feegrant allowance** approach as an alternative architecture — this is more native to Cosmos SDK and avoids the ante handler state-write problem entirely.

---

## 5. Flow Summary (How the Feature Is Supposed to Work)

```
Oracle submits MsgSubmitMintResponse
        │
        ▼
[VoteAloneDecorator]
  - tx has 1 msg? ✅ proceed
  - tx has >1 msg with oracle vote? ❌ reject
        │
        ▼
[NewSetUpContextDecorator, ValidateBasic, ...]
        │
        ▼
[GaslessDecorator]
  IsTxGasless()?
    → HasOraclePermission?         ✅ / ❌
    → MintRequest exists?          ✅ / ❌
    → MintRequest PENDING?         ✅ / ❌
    → Oracle not already voted?    ✅ / ❌
    → Not spam (block height)?     ✅ / ❌
  
  If gasless=true:
    → Skip DeductFeeDecorator
    → Apply InfiniteGasMeter
    → Set Priority = MaxInt64-100
  
  If gasless=false:
    → Run DeductFeeDecorator normally
        │
        ▼
[SigVerification, IncrementSequence, ...]
        │
        ▼
[Message Server: SubmitMintResponse]
  → Performs same checks again
  → Appends oracle to Confirmers[]
  → If Confirmers >= RequiredConfirm → finalize
```

> Note: The validation in the ante handler and the message server is **duplicated** by design (defense in depth), but this means any state change in the ante handler (like setting vote height) can get out of sync with the message server execution.

---

## 6. Diff of Critical Fix (Bug #1)

Apply this change to `app/app.go` in `setAnteHandler`:

```go
// BEFORE (broken):
options := ante.HandlerOptions{
    Cdc:                    app.appCodec,
    AccountKeeper:          app.AccountKeeper,
    BankKeeper:             app.BankKeeper,
    DistributionKeeper:     app.DistrKeeper,
    IBCKeeper:              app.IBCKeeper,
    StakingKeeper:          app.StakingKeeper,
    FeeMarketKeeper:        app.FeeMarketKeeper,
    EvmKeeper:              app.EVMKeeper,
    FeegrantKeeper:         app.FeeGrantKeeper,
    ExtensionOptionChecker: evmostypes.HasDynamicFeeExtensionOption,
    SignModeHandler:        txConfig.SignModeHandler(),
    SigGasConsumer:         ante.SigVerificationGasConsumer,
    MaxTxGasWanted:         maxGasWanted,
    TxFeeChecker:           ethante.NewDynamicFeeChecker(app.EVMKeeper),
    CircuitKeeper:          &app.CircuitBreakerKeeper,
    AllowUnorderedTx:       unsafeUnorderedTx,
    // ❌ Missing NftOracleKeeper and NftAdminKeeper
}

// AFTER (fixed):
options := ante.HandlerOptions{
    Cdc:                    app.appCodec,
    AccountKeeper:          app.AccountKeeper,
    BankKeeper:             app.BankKeeper,
    DistributionKeeper:     app.DistrKeeper,
    IBCKeeper:              app.IBCKeeper,
    StakingKeeper:          app.StakingKeeper,
    FeeMarketKeeper:        app.FeeMarketKeeper,
    EvmKeeper:              app.EVMKeeper,
    FeegrantKeeper:         app.FeeGrantKeeper,
    ExtensionOptionChecker: evmostypes.HasDynamicFeeExtensionOption,
    SignModeHandler:        txConfig.SignModeHandler(),
    SigGasConsumer:         ante.SigVerificationGasConsumer,
    MaxTxGasWanted:         maxGasWanted,
    TxFeeChecker:           ethante.NewDynamicFeeChecker(app.EVMKeeper),
    CircuitKeeper:          &app.CircuitBreakerKeeper,
    AllowUnorderedTx:       unsafeUnorderedTx,
    NftOracleKeeper:        &app.NftoracleKeeper,  // ✅ ADD
    NftAdminKeeper:         &app.NftadminKeeper,   // ✅ ADD
}
```

---

*End of review.*
