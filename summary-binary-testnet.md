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
| Binary version | `sixd v4.0.3-35-gecd6185a` → `v4.0.3-36-g2538722e` (after fixes) |

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

**Result:** ✅ Built successfully

---

## 2. Local Testnet

### Setup

```bash
./init_testnet.sh mynode 0
```

- Chain ID (Cosmos): `testnet`
- Chain ID (EVM): `666` (`0x29a`)
- Key algorithm: `eth_secp256k1` *(fixed from `secp256k1`)* 
- Keyring backend: `test`
- Addresses: Derived dynamically after key import (Keccak256 derivation)

### Test Accounts

| Account | Cosmos Address | EVM Hex Address | Initial Balance |
|---|---|---|---|
| Alice | `6x15murggkqrgpv970jcuqmn0dc280wk3txfhuek4` | `0xA6f83422C01a02C2f9F2c701B9BDB851DEeB4566` | 1,000,000,000,000 usix |
| Bob | `6x1jwr9hhq7c5x5x23tw4l26xs4nu6f26ct68das8` | `0x93865Bdc1EC50d432A2B757ead1A159f34956b0B` | 10,000,000,000,000 usix |
| Super-admin | `6x1lq9ylung2jxtn32nsmndc6kr2vz8utmpvdcuhh` | — | 200,000,000,000,000 usix |

> **Note:** With `eth_secp256k1`, Cosmos and EVM addresses share the same Keccak256 derivation. The same private key produces the same address on both layers.

### Verification

Chain started and produced blocks successfully. Node ready in ~12 seconds.

Key type confirmed:
```
pubkey: '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"A7ORN3A8TuvNIq3uzfJZn/aEntUXfK9/GKBfm7lZArwE"}'
```

---

## 3. Test Results

### 3.1 Unit Tests

All 24 unit tests pass across 2 packages:

```
ok  github.com/thesixnetwork/six-protocol/v4/x/precisebank/types   0.585s
ok  github.com/thesixnetwork/six-protocol/v4/x/precisebank/keeper  0.607s
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
| Fractional Balance | `sixd query precisebank fractional-balance --address=6x15murggkqrgpv970jcuqmn0dc280wk3txfhuek4` | `{"fractional_balance":"0"}` | ✅ |
| Fractional Balances | `sixd query precisebank fractional-balances` | `{"pagination":{}}` | ✅ |

REST API also confirmed working:
```
GET http://localhost:1317/thesixnetwork/six-protocol/precisebank/remainder → {"remainder":"0"}
```

### 3.3 EVM Balance Query via Precisebank

`eth_getBalance` for Alice's EVM address correctly returns usix balance converted to 18-decimal asix:

```
Alice usix balance:  1,000,000,000,000
eth_getBalance:      0xd3c217a4a19177e7c000  (≈ 1,000,000,000,000 × 10^12 asix)
✅ MATCH
```

**Conclusion:** Precisebank `GetBalance()` correctly computes `(bank_balance_usix × 10^12) + fractional_balance` and is properly integrated with the EVM keeper for `eth_getBalance` queries.

### 3.4 Cosmos Bank Send (usix)

```bash
sixd tx bank send alice bob 100usix --fees 300000usix -y
# Result: ✅ code: 0
# Bob balance: 10,000,000,000,000 → 10,000,000,000,100 usix (+100)
```

Standard Cosmos bank transfers using `usix` work as expected.

### 3.5 Cosmos Bank Send (asix) — Expected Failure

```bash
sixd tx bank send alice bob 500000000000asix --fees 300000usix -y
# Result: ❌ code: 5 — "spendable balance 0asix is smaller than 500000000000asix: insufficient funds"
```

**This is by design.** The standard bank `MsgSend` handler calls `bankKeeper.SpendableCoin()` directly, not through PreciseBankKeeper. Since `asix` is not stored in `x/bank`, the balance is reported as 0. Precisebank is a keeper-only module without its own message server — it only works when modules explicitly use `PreciseBankKeeper` (e.g., the EVM keeper).

### 3.6 EVM Transaction (cast send) — ✅ PASS (after fix)

```bash
cast send --rpc-url http://localhost:8545 \
  --private-key 0x<alice_privkey> \
  0x93865Bdc1EC50d432A2B757ead1A159f34956b0B \
  --value 500000000000
# Result: ✅ status: 1 (success), block: 364, gasUsed: 21000
```

**Transaction details:**
```
txHash:    0x52ededb5f61bded21c3c22a104d385e1bd338052c444d49a112bb8d876e2090c
from:      0xA6f83422C01a02C2f9F2c701B9BDB851DEeB4566 (Alice)
to:        0x93865Bdc1EC50d432A2B757ead1A159f34956b0B (Bob)
value:     500,000,000,000 asix (sub-usix amount)
gasUsed:   21,000
status:    1 (success)
```

### 3.7 Fractional Balance Creation via EVM — ✅ PASS

After the EVM send of 500,000,000,000 asix (< 1 usix), fractional balances are correctly created:

```json
{
  "fractional_balances": [
    {"address": "6x1jwr9hhq7c5x5x23tw4l26xs4nu6f26ct68das8", "fractional_balance": "500000000000"},
    {"address": "6x15murggkqrgpv970jcuqmn0dc280wk3txfhuek4", "fractional_balance": "499999979000"},
    {"address": "6x17xpfvakm2amg962yls6f84z3kell8c5llh54l0", "fractional_balance": "21000"}
  ]
}
```

**Verification:**
| Account | Fractional Balance (asix) | Explanation |
|---|---|---|
| Bob (recipient) | 500,000,000,000 | Received sub-usix transfer |
| Alice (sender) | 499,999,979,000 | Change after sent amount + gas |
| Fee collector | 21,000 | Gas fee (21000 gas × 1 wei) |
| **Total** | **1,000,000,000,000** | **= exactly 1 usix (10^12)** |

**Remainder:** `0` — perfectly balanced. The total fractional amounts sum to exactly 1 usix, confirming the precisebank invariant holds.

### 3.8 Export Eth Key — ✅ PASS

```bash
sixd keys unsafe-export-eth-key alice --home ~/.six --keyring-backend test
# Result: ✅ Successfully exports private key (D28B370F7E...)
```

---

## 4. Bugs Found & Fixed

### Bug #1 (Critical): EVM Transaction Fails — secp256k1 Address Derivation Mismatch

**Status:** ✅ **FIXED** in commit `2538722e`

**Symptom:** All EVM transactions fail with `sender balance < tx cost (0 < ...)` even though `eth_getBalance` returns the correct balance.

**Root Cause:** The test accounts were created with `--algo secp256k1` (`init_testnet.sh` line 13: `KEYALGO="secp256k1"`). Cosmos (secp256k1) and Ethereum (eth_secp256k1) derive addresses using different hash algorithms, producing different addresses from the same key.

The fix required changes at **three levels**:
1. `init_testnet.sh`: `KEYALGO="eth_secp256k1"`, `--algo` → `--key-type`, dynamic address derivation
2. `cmd/sixd/cmd/root.go`: Uncommented `WithKeyringOptions(sixkey.Option())` to register eth_secp256k1
3. `client/keys.go`: Enabled custom `addCmd` with eth_secp256k1 as default key type

### Bug #2 (Minor): Ante Handler Vesting Check Uses Native BankKeeper

**Status:** ✅ **FIXED** in commit `2538722e`

**Location:** `app/app.go` line 1131

**Fix:** Changed `BankKeeper: app.BankKeeper` → `BankKeeper: app.PreciseBankKeeper` in `setAnteHandler()`.

### Issue #3 (Design Limitation): No Cosmos-Level asix Transfers

**Status:** Not changed — by design

Users cannot send `asix` via standard Cosmos `MsgSend`. Precisebank only works when other modules use `PreciseBankKeeper` directly (e.g., EVM keeper).

---

## 5. Code Review Fixes Applied

In addition to the bugs above, the following code review issues were fixed in commit `2538722e`:

| # | Severity | File | Issue | Fix |
|---|---|---|---|---|
| 1 | Critical | `types/fractional_balance.go` | `ConversionFactor()` used `NewIntFromBigIntMut` which wraps internal `*big.Int` without copy — callers could mutate the global | Use `new(big.Int).Set()` for safe deep copy |
| 2 | Critical | `types/extended_balance.go` | Direct use of private `conversionFactor` var instead of safe `ConversionFactor()` function | Changed to `ConversionFactor()` |
| 3 | Medium | `keeper/fractional_balance.go` | Mixed pointer/value receivers (`*Keeper` vs `Keeper`) | Standardized all 8 methods to value receivers (`Keeper`) |
| 4 | Medium | `keeper/send.go` | `SendCoinsFromModuleToAccount` used `GetModuleAddress` which only returns address without verifying account exists | Changed to `GetModuleAccount` which verifies the account |

---

## 6. Files Changed

| File | Changes |
|---|---|
| `app/app.go` | `BankKeeper: app.BankKeeper` → `app.PreciseBankKeeper` in ante handler |
| `client/keys.go` | Enabled eth_secp256k1 key support (uncommented addCmd, runAddCmd, imports) |
| `cmd/sixd/cmd/root.go` | Enabled `WithKeyringOptions(sixkey.Option())` for eth_secp256k1 keyring |
| `init_testnet.sh` | `KEYALGO="eth_secp256k1"`, `--key-type` flag, dynamic address derivation |
| `x/precisebank/keeper/fractional_balance.go` | Standardized all methods to value receivers |
| `x/precisebank/keeper/send.go` | `GetModuleAddress` → `GetModuleAccount` |
| `x/precisebank/types/extended_balance.go` | Use `ConversionFactor()` instead of private var |
| `x/precisebank/types/fractional_balance.go` | Safe deep copy in `ConversionFactor()` |

---

## 7. Summary Table

| Test | Status | Notes |
|---|---|---|
| Build sixd | ✅ Pass | Requires `unset GOROOT` |
| Start local testnet | ✅ Pass | `init_testnet.sh mynode 0`, node ready in ~12s |
| Unit tests (24 total) | ✅ All Pass | keeper: 0.607s, types: 0.585s |
| gRPC queries (CLI) | ✅ Pass | remainder, fractional-balance, fractional-balances |
| REST API queries | ✅ Pass | `localhost:1317` |
| EVM `eth_getBalance` | ✅ Pass | Correctly returns usix × 10^12 |
| Cosmos bank send (usix) | ✅ Pass | Standard transfer works |
| Cosmos bank send (asix) | ❌ Expected Fail | By design — no precisebank message server |
| EVM transaction (cast send) | ✅ Pass | 500B asix transfer, status=1, block=364 |
| Fractional balance creation | ✅ Pass | 3 fractional balances created, sum = exactly 1 usix |
| Remainder invariant | ✅ Pass | Remainder = 0 (perfectly balanced) |
| Export eth key | ✅ Pass | `sixd keys unsafe-export-eth-key` works |
| Key algorithm | ✅ Pass | Confirmed `ethermint.crypto.v1.ethsecp256k1.PubKey` |
