package types

// IntegerCoinDenom is the denomination for integer coins managed by x/bank.
// This is the native cosmos-side denomination of the chain.
const IntegerCoinDenom = "usix"

// ExtendedCoinDenom is the denomination for the extended IntegerCoinDenom.
// This represents the total balance of integer + fractional balances,
// used by x/evm where 18 decimal points are expected.
const ExtendedCoinDenom = "asix"
