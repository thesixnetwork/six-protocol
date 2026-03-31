# Precisebank Module — Build, Testnet & Testing Summary

## 1. Build

### Environment

| Item | Value |
|---|---|
| OS | macOS |
| Go toolchain | go1.25.7 (auto-downloaded via `GOTOOLCHAIN=auto`) |
| Cosmos SDK | v0.50.13 |
| Evmos fork | `thesixnetwork/evmos/v20@v20.0.0-six` |
| Branch | `feature/precisebank` |

### Build Issue: Go Version Mismatch

**Symptom:**
```
compile: version "go1.24.7" does not match go tool version "go1.25.7"
```

**Root Cause:** `~/.zshrc` exports `GOROOT=/Users/sankanvicha/goroot` which is a symlink to `~/go-src/go1.24.7/`. When `GOTOOLCHAIN=auto` downloads `go1.25.7` to compile the project, the hard-coded `GOROOT` forces it to use go1.24.7's stdlib, causing a version mismatch.

**Fix:**
```bash
unset GOROOT && go clean -cache && make install
```

**Result:** ✅ Built successfully — `sixd version 4.0.3-34-g281d18f2`

---

## 2. Local Testnet

### Setup

```bash
./init_testnet.sh mynode 0
```

- Chain ID (Cosmos): `testnet`
- Chain ID (EVM): `666` (`0x29a`)
- Key algorithm: `secp256k1`
- Keyring backend: `test`

### Test Accounts

| Account | Cosmos Address | EVM Address (Cosmos-derived) | Initial Balance |
|---|---|---|---|
| Alice | `6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq` | `0xd907f36f7D83344057a619b6D83A45B3288c3c21` | 1,000,000,000,000 usix |
| Bob | `6x13g50hqdqsjk85fmgqz2h5xdxq49lsmjdwlemsp` | `0x8a28fb81A084Ac7A276800957a19a6054BF86E4D` | 10,000,000,000,000 usix |
| Super-admin | `6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv` | — | 200,000,000,000,000 usix |

### Verification

Chain started and produced blocks successfully.

---

## 3. Test Results

### 3.1 Unit Tests

All 24 unit tests pass across 2 packages:

```
ok  github.com/thesixnetwork/six-protocol/v4/x/precisebank/keeper  0.793s
ok  github.com/thesixnetwork/six-protocol/v4/x/precisebank/types   1.100s
```

<details>
<summary>Full test list (all PASS)</summary>

**keeper package (10 tests):**
- TestRemainder_NilRequest
- TestRemainder_Default
- TestRemainder_WithValue
- TestFractionalBalance_NilRequest
- TestFractionalBalance_InvalidAddress
- TestFractionalBalance_Zero
- TestFractionalBalance_WithValue
- TestFractionalBalances_NilRequest
- TestFractionalBalances_Empty
- TestFractionalBalances_Multiple

**types package (14 tests):**
- TestConversionFactor
- TestConversionFactor_Immutable
- TestNewFractionalBalance
- TestFractionalBalance_Validate (6 sub-tests)
- TestValidateFractionalAmount (6 sub-tests)
- TestFractionalBalances_Validate
- TestFractionalBalances_SumAmount
- TestGenesisState_Validate
- TestGenesisState_TotalAmountWithRemainder
- TestDefaultGenesis
- TestConstants
- TestModuleName
- TestKeyPrefixes

</details>

### 3.2 gRPC Query Endpoints

| Query | Command | Result | Status |
|---|---|---|---|
| Remainder | `sixd query precisebank remainder` | `{"remainder":"0"}` | ✅ |
| Fractional Balance | `sixd query precisebank fractional-balance --address=6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq` | `{"fractional_balance":"0"}` | ✅ |
| Fractional Balances | `sixd query precisebank fractional-balances` | `{"pagination":{}}` | ✅ |

REST API also confirmed working:
```
GET http://localhost:1317/thesixnetwork/six-protocol/precisebank/remainder → {"remainder":"0"}
```

### 3.3 EVM Balance Query via Precisebank

`eth_getBalance` for Alice's Cosmos-mapped EVM address correctly returns usix balance converted to 18-decimal asix:

```
Alice usix balance: 999,999,700,000
eth_getBalance:     999,999,700,000,000,000,000,000 asix (= 999,999,700,000 × 10^12)
✅ MATCH
```

Bob's balance after receiving 100 usix:
```
Bob usix balance:   10,000,000,000,100
eth_getBalance:     10,000,000,000,100,000,000,000,000 asix (= 10,000,000,000,100 × 10^12)
✅ MATCH
```

**Conclusion:** Precisebank `GetBalance()` correctly computes `(bank_balance_usix × 10^12) + fractional_balance` and is properly integrated with the EVM keeper for `eth_getBalance` queries.

### 3.4 Cosmos Bank Send (usix)

```bash
sixd tx bank send alice bob 100usix --fees 300000usix -y
# Result: ✅ code: 0, height: 2486
```

Standard Cosmos bank transfers using `usix` work as expected.

### 3.5 Cosmos Bank Send (asix) — Expected Failure

```bash
sixd tx bank send alice bob 500000000000asix --fees 300000usix -y
# Result: ❌ code: 5 — "spendable balance 0asix is smaller than 500000000000asix: insufficient funds"
```

**This is by design.** The standard bank `MsgSend` handler calls `bankKeeper.SpendableCoin()` directly, not through PreciseBankKeeper. Since `asix` is not stored in `x/bank`, the balance is reported as 0. Precisebank is a keeper-only module without its own message server — it only works when modules explicitly use `PreciseBankKeeper` (e.g., the EVM keeper).

### 3.6 EVM Transaction (cast send) — BUG FOUND

```bash
cast send --rpc-url http://localhost:8545 \
  --private-key 0xd28b370f7ec4e8f688be18a7fc6aaa6c0daf77bce3e90b9cc6ff44cc8dd1bd9d \
  0x8a28fb81A084Ac7A276800957a19a6054BF86E4D \
  --value 500000000000
# Result: ❌ "sender balance < tx cost (0 < 144375500000000000)"
```

See **Bug #1** below for root cause analysis.

---

## 4. Bugs & Issues Found

### Bug #1 (Critical): EVM Transaction Fails — secp256k1 Address Derivation Mismatch

**Severity:** Critical — Blocks all EVM transactions from testnet accounts

**Symptom:** All EVM transactions fail with `sender balance < tx cost (0 < ...)` even though `eth_getBalance` returns the correct balance.

**Root Cause:**

The test accounts are created with `--algo secp256k1` (`init_testnet.sh` line 13: `KEYALGO="secp256k1"`). Cosmos (secp256k1) and Ethereum (eth_secp256k1) derive addresses from the **same private key** using **different hash algorithms**:

| Path | Algorithm | Resulting Address |
|---|---|---|
| Cosmos | RIPEMD160(SHA256(compressed_pubkey)) | `0xd907f36f7D83344057a619b6D83A45B3288c3c21` |
| EVM (Keccak256) | Keccak256(uncompressed_pubkey)[12:] | `0xA6f83422C01a02C2f9F2c701B9BDB851DEeB4566` |

When `cast send` signs and broadcasts an EVM transaction:
1. The node recovers the signer address using **Keccak256** → gets `0xA6f8...`
2. The ante handler calls `evmKeeper.GetAccount(ctx, 0xA6f8...)` to check balance
3. PreciseBankKeeper looks up the Cosmos address mapped to `0xA6f8...` → finds no balance (no account exists at that address)
4. Returns balance = 0 → transaction rejected

Meanwhile, `eth_getBalance` takes the address as a **parameter** (doesn't recover from a signature), so querying `0xd907...` directly works correctly.

**Fix:** Change `init_testnet.sh` line 13 from:
```bash
KEYALGO="secp256k1"
```
to:
```bash
KEYALGO="eth_secp256k1"
```

With `eth_secp256k1`, both Cosmos and EVM derive addresses using Keccak256, producing the **same address** for both layers. The chain already supports `eth_secp256k1` (see `client/keys/add.go` and `app/ante/sigverify.go`).

**Note:** After changing the key algorithm, the test account addresses will change. The mnemonics will derive different Cosmos addresses since `eth_secp256k1` uses a different address derivation path.

### Bug #2 (Minor): Ante Handler Vesting Check Uses Native BankKeeper

**Severity:** Minor — Affects only vesting accounts

**Location:** `app/app.go` line 1131 → `app/ante/evm/08_vesting.go` line 141

**Description:** The EVM ante handler's `CheckVesting()` function receives `app.BankKeeper` (native bank) and calls `bankKeeper.GetBalance(ctx, address, denom)` with `denom = "asix"`. Since native bank doesn't track `asix` balances (only precisebank does), this returns 0 for all accounts. For non-vesting accounts, this code path is skipped. For vesting accounts, this would incorrectly report insufficient spendable balance.

**Fix:** In `app/app.go` `setAnteHandler()`, change:
```go
BankKeeper: app.BankKeeper,
```
to:
```go
BankKeeper: app.PreciseBankKeeper,
```

This ensures the vesting check gets the correct extended balance via PreciseBankKeeper.

### Issue #3 (Design Limitation): No Cosmos-Level asix Transfers

**Severity:** Low — By design, but worth noting

**Description:** Users cannot send `asix` via standard Cosmos `MsgSend` (e.g., `sixd tx bank send alice bob 500asix`). The precisebank module has no message server and doesn't intercept bank `MsgSend`. It only works when other modules (like the EVM keeper) use `PreciseBankKeeper` directly.

**Impact:** Users who want to transfer sub-usix amounts can only do so through EVM transactions, not through Cosmos CLI.

**Possible Enhancement:** Implement a `SendRestriction` on the bank module to route `asix` sends through precisebank, or add a dedicated `MsgSend` to the precisebank module.

---

## 5. Test Coverage Gaps

The following areas lack test coverage and would benefit from additional tests:

| Area | What to Test |
|---|---|
| **Integration with EVM keeper** | End-to-end EVM transaction that triggers precisebank send/mint/burn |
| **Fractional balance operations** | Live chain test: send amounts that create non-zero fractional balances (e.g., send 500,000,000,000 asix = 0.5 usix worth of fractional) |
| **Mint/Burn via EVM** | Contract deployment and token operations that trigger mint/burn through EVM |
| **Invariant checks** | Run `sixd query precisebank` invariants after various operations |
| **Ante handler** | Verify fee deduction works correctly through PreciseBankKeeper |
| **Vesting account** | EVM transaction from a vesting account (blocked by Bug #2) |

---

## 6. Summary Table

| Test | Status | Notes |
|---|---|---|
| Build sixd | ✅ Pass | Requires `unset GOROOT` |
| Start local testnet | ✅ Pass | `init_testnet.sh mynode 0` |
| Unit tests (24 total) | ✅ All Pass | keeper: 10, types: 14 |
| gRPC queries (CLI) | ✅ Pass | remainder, fractional-balance, fractional-balances |
| REST API queries | ✅ Pass | `localhost:1317` |
| EVM `eth_getBalance` | ✅ Pass | Correctly returns usix × 10^12 |
| Cosmos bank send (usix) | ✅ Pass | Standard transfer works |
| Cosmos bank send (asix) | ❌ Expected Fail | By design — no precisebank message server |
| EVM transaction (cast send) | ❌ Bug #1 | secp256k1 → eth_secp256k1 address mismatch |
| EVM vesting check | ❌ Bug #2 | Ante handler uses native BankKeeper for asix |
