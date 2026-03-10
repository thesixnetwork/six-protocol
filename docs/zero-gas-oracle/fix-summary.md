# Fix Summary: Oracle Zero-Gas Voting Feature

**Date:** 2026-03-10  
**Branch:** (current working branch)  
**Files Changed:** 7

---

## Overview

Eight issues were identified and resolved across the oracle zero-gas voting feature. The most critical prevented the node from starting at all. The others ranged from incorrect state-mutation lifecycle to dead code and duplicate error constants.

---

## Bug #1 — CRITICAL: Missing Keeper Injection in `setAnteHandler`

**File:** `app/app.go`  
**Severity:** 🔴 Critical — Node panics on startup

### Problem
`HandlerOptions.Validate()` asserts that `NftOracleKeeper` and `NftAdminKeeper` are non-nil. The `setAnteHandler` function never assigned them, causing an immediate panic every time the node tried to start.

### Fix
```go
// Added to HandlerOptions struct literal inside setAnteHandler():
NftOracleKeeper: &app.NftoracleKeeper,
NftAdminKeeper:  &app.NftadminKeeper,
```

---

## Bug #2 — MEDIUM: Block-Level Spam Prevention Too Aggressive

**File:** `app/ante/gasless.go`  
**Severity:** 🟡 Medium — Legitimate oracle votes silently rejected

### Problem
`checkAndSetSpamPreventionCounter` blocked an oracle from submitting **any** gasless tx in a block once they had voted on **any** request. With multiple pending requests in the same block, oracles could only vote on one of them, breaking quorum for all others.

### Fix
Removed `checkAndSetSpamPreventionCounter` and all its callsites entirely. The existing per-request duplicate-vote check (`hasOracleAlreadyVoted`) is the correct and sufficient guard.

---

## Bug #3 — MEDIUM: State Write Inside Ante Handler

**File:** `app/ante/gasless.go`  
**Severity:** 🟡 Medium — Incorrect state mutation before tx execution

### Problem
`checkAndSetSpamPreventionCounter` called `SetOracleLastVoteHeight` (a KV-store write) during the ante handler phase. If any later decorator rejected the transaction, the oracle's vote-height would already be persisted, silently blocking that oracle from gasless voting for the rest of the block even though the tx never executed.

### Fix
Eliminated along with Bug #2 (same function removed).

---

## Bug #4 — MEDIUM: `CleanupOldVoteHeights` Never Called

**File:** `x/nftoracle/module/module.go`  
**Severity:** 🟡 Medium — Store grows unboundedly

### Problem
`CleanupOldVoteHeights` was implemented in `keeper/gasless_voting.go` but never wired into the module's lifecycle. Vote-height records accumulated in the KV store indefinitely.

### Fix
Wired into `EndBlock`:
```go
func (am AppModule) EndBlock(ctx context.Context) error {
    am.keeper.CleanupOldVoteHeights(ctx, 100)
    return nil
}
```

---

## Bug #5 — MEDIUM: `NftmngrKeeper` Initialized Before `NftadminKeeper`

**File:** `app/app.go`  
**Severity:** 🟡 Medium — `NftmngrKeeper` receives zero-value `NftadminKeeper`

### Problem
`app.NftmngrKeeper` was constructed at line 763 and passed `app.NftadminKeeper` as a dependency. `app.NftadminKeeper` was not initialized until line 775. The `NftmngrKeeper` therefore held a zero-value (uninitialized) admin keeper.

### Fix
Swapped initialization order: `NftadminKeeper` is now initialized first, then `NftmngrKeeper` receives the fully-initialized keeper.

---

## Issue #6 — LOW: `IsOracleGaslessEnabled()` Dead Code

**File:** `x/nftoracle/keeper/gasless_voting.go`  
**Severity:** 🟢 Low — Misleading dead code

### Problem
`IsOracleGaslessEnabled()` was hardcoded to `return true` and never called anywhere. Its presence implied a feature flag that did not exist.

### Fix
Removed the function entirely.

---

## Issue #7 — LOW: Duplicate Error Codes

**Files:** `app/ante/gasless.go`, `x/nftoracle/types/errors.go`  
**Severity:** 🟢 Low — Confusing duplicated constants

### Problem
Error codes 602 (`ErrCollectionOwnerRequestNotFound`) and 603 (`ErrCollectionOwnerRequestNotPending`) duplicated the semantics of the already-existing codes 300 (`ErrVerifyRequestNotFound`) and 301 (`ErrVerifyRequestNotPending`). Error code 601 (`ErrOracleSpamPrevention`) became unused after Bug #2/#3 was fixed.

### Fix
- Replaced `ErrCollectionOwnerRequestNotFound` → `ErrVerifyRequestNotFound` in `gasless.go`
- Replaced `ErrCollectionOwnerRequestNotPending` → `ErrVerifyRequestNotPending` in `gasless.go`
- Removed codes 601, 602, 603 from `errors.go`

---

## Issue #8 — LOW: Shallow Test Coverage

**Files:** `app/ante/gasless_test.go`, `testutil/keeper/combined.go`  
**Severity:** 🟢 Low — Missing test coverage

### Problem
The existing tests only verified basic cases (empty tx, multi-msg rejection) without exercising the real oracle permission system or any of the KV-store interactions.

### Fix
- Added `testutil/keeper/combined.go`: a shared `CombinedKeepers` factory that mounts both `nftoracle` and `nftadmin` stores on the same `CommitMultiStore`, enabling cross-keeper tests.
- Rewrote `gasless_test.go` with **35 test cases** covering:
  - Happy paths for all three message types (MintResponse, ActionResponse, CollectionVerify)
  - Permission rejection for non-oracle addresses
  - `RequestNotFound` and `NotPending` error paths
  - Duplicate-vote detection
  - Multi-oracle voting (other oracle already voted, but current oracle may still vote)
  - Infinite gas meter in both CheckTx and DeliverTx
  - Oracle priority assignment
  - Wrapped decorator skip (gasless) vs execute (non-gasless)
  - `SetOracleLastVoteHeight` / `GetOracleLastVoteHeight` / `DeleteOracleLastVoteHeight`
  - `CleanupOldVoteHeights` pruning logic
  - `VoteAloneDecorator` for all oracle message types

**All 35 tests pass.**

---

## Files Changed

| File | Change |
|------|--------|
| `app/app.go` | Bug #1: Add `NftOracleKeeper`/`NftAdminKeeper` to `HandlerOptions`; Bug #5: Reorder keeper init |
| `app/ante/gasless.go` | Bug #2+#3: Remove `checkAndSetSpamPreventionCounter`; Issue #7: Replace duplicate error refs |
| `x/nftoracle/module/module.go` | Bug #4: Wire `CleanupOldVoteHeights` in `EndBlock` |
| `x/nftoracle/keeper/gasless_voting.go` | Issue #6: Remove `IsOracleGaslessEnabled` dead code |
| `x/nftoracle/types/errors.go` | Issue #7: Remove codes 601–603 |
| `app/ante/gasless_test.go` | Issue #8: Comprehensive test suite (35 tests) |
| `testutil/keeper/combined.go` | Issue #8: New shared keeper factory for cross-module tests |
