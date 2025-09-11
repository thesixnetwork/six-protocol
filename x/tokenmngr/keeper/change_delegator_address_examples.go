package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Example demonstrates how to use the ChangeDelegatorAddress function
func ExampleChangeDelegatorAddress() {
	// This is a demonstration of how to use the function
	// In practice, you would have a real keeper instance and context

	fmt.Println("ChangeDelegatorAddress Usage Example")
	fmt.Println("====================================")

	// Example 1: Basic usage
	fmt.Println("\n1. Basic Usage:")
	fmt.Println("```go")
	fmt.Println("// Convert string addresses to sdk.AccAddress")
	fmt.Println(`oldAddrStr := "cosmos1oldaddress..."`)
	fmt.Println(`newAddrStr := "cosmos1newaddress..."`)
	fmt.Println("")
	fmt.Println("oldAddr, err := sdk.AccAddressFromBech32(oldAddrStr)")
	fmt.Println("if err != nil { return err }")
	fmt.Println("")
	fmt.Println("newAddr, err := sdk.AccAddressFromBech32(newAddrStr)")
	fmt.Println("if err != nil { return err }")
	fmt.Println("")
	fmt.Println("// Execute the address change")
	fmt.Println("err = keeper.ChangeDelegatorAddress(ctx, oldAddr, newAddr)")
	fmt.Println("if err != nil { return err }")
	fmt.Println("```")

	// Example 2: With validation
	fmt.Println("\n2. With Pre-validation:")
	fmt.Println("```go")
	fmt.Println("// Validate addresses first")
	fmt.Println("if oldAddr.Empty() || newAddr.Empty() {")
	fmt.Println(`    return errors.New("addresses cannot be empty")`)
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("if oldAddr.Equals(newAddr) {")
	fmt.Println(`    return errors.New("old and new addresses cannot be the same")`)
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("// Get info before change (optional)")
	fmt.Println("infoBefore, err := keeper.GetDelegatorStakingInfo(ctx, oldAddr)")
	fmt.Println("if err != nil { return err }")
	fmt.Println("")
	fmt.Println("// Execute change")
	fmt.Println("err = keeper.ChangeDelegatorAddress(ctx, oldAddr, newAddr)")
	fmt.Println("if err != nil { return err }")
	fmt.Println("```")

	// Example 3: Error handling
	fmt.Println("\n3. Error Handling:")
	fmt.Println("```go")
	fmt.Println("err := keeper.ChangeDelegatorAddress(ctx, oldAddr, newAddr)")
	fmt.Println("switch {")
	fmt.Println("case errors.Is(err, stakingtypes.ErrEmptyDelegatorAddr):")
	fmt.Println(`    log.Error("Invalid address provided")`)
	fmt.Println("case errors.Is(err, stakingtypes.ErrBadDelegatorAddr):")
	fmt.Println(`    log.Error("Bad delegator address")`)
	fmt.Println("case err != nil:")
	fmt.Println(`    log.Error("Failed to change delegator address", "error", err)`)
	fmt.Println("default:")
	fmt.Println(`    log.Info("Successfully changed delegator address")`)
	fmt.Println("}")
	fmt.Println("```")

	// Example 4: Complete workflow
	fmt.Println("\n4. Complete Workflow:")
	fmt.Println("```go")
	fmt.Println("func TransferStakingRecords(keeper *Keeper, ctx context.Context, oldAddrStr, newAddrStr string) error {")
	fmt.Println("    // Parse addresses")
	fmt.Println("    oldAddr, err := sdk.AccAddressFromBech32(oldAddrStr)")
	fmt.Println("    if err != nil {")
	fmt.Println(`        return fmt.Errorf("invalid old address: %w", err)`)
	fmt.Println("    }")
	fmt.Println("")
	fmt.Println("    newAddr, err := sdk.AccAddressFromBech32(newAddrStr)")
	fmt.Println("    if err != nil {")
	fmt.Println(`        return fmt.Errorf("invalid new address: %w", err)`)
	fmt.Println("    }")
	fmt.Println("")
	fmt.Println("    // Get current state for logging")
	fmt.Println("    infoBefore, err := keeper.GetDelegatorStakingInfo(ctx, oldAddr)")
	fmt.Println("    if err != nil {")
	fmt.Println(`        return fmt.Errorf("failed to get initial staking info: %w", err)`)
	fmt.Println("    }")
	fmt.Println("")
	fmt.Println("    // Log what we're about to transfer")
	fmt.Println(`    log.Info("Starting delegator address change",`)
	fmt.Println(`        "old_address", oldAddrStr,`)
	fmt.Println(`        "new_address", newAddrStr,`)
	fmt.Println(`        "delegations", len(infoBefore.Delegations),`)
	fmt.Println(`        "unbonding_delegations", len(infoBefore.UnbondingDelegations),`)
	fmt.Println(`        "redelegations", len(infoBefore.Redelegations),`)
	fmt.Println(`        "total_bonded", infoBefore.TotalBonded,`)
	fmt.Println(`        "total_unbonding", infoBefore.TotalUnbonding)`)
	fmt.Println("")
	fmt.Println("    // Execute the change")
	fmt.Println("    err = keeper.ChangeDelegatorAddress(ctx, oldAddr, newAddr)")
	fmt.Println("    if err != nil {")
	fmt.Println(`        return fmt.Errorf("failed to change delegator address: %w", err)`)
	fmt.Println("    }")
	fmt.Println("")
	fmt.Println("    // Verify the change")
	fmt.Println("    infoAfter, err := keeper.GetDelegatorStakingInfo(ctx, newAddr)")
	fmt.Println("    if err != nil {")
	fmt.Println(`        return fmt.Errorf("failed to get final staking info: %w", err)`)
	fmt.Println("    }")
	fmt.Println("")
	fmt.Println("    // Verify totals match")
	fmt.Println("    if !infoBefore.TotalBonded.Equal(infoAfter.TotalBonded) {")
	fmt.Println(`        return fmt.Errorf("total bonded amount mismatch: before=%s, after=%s",`)
	fmt.Println("            infoBefore.TotalBonded, infoAfter.TotalBonded)")
	fmt.Println("    }")
	fmt.Println("")
	fmt.Println(`    log.Info("Successfully completed delegator address change")`)
	fmt.Println("    return nil")
	fmt.Println("}")
	fmt.Println("```")

	fmt.Println("\n5. Important Notes:")
	fmt.Println("- This function requires administrative privileges")
	fmt.Println("- Test thoroughly in development before production use")
	fmt.Println("- Consider gas costs for large numbers of delegations")
	fmt.Println("- Backup state before running in production")
	fmt.Println("- Monitor for any reward distribution side effects")
}

// ValidationResult represents the result of address validation
type ValidationResult struct {
	Valid        bool
	ErrorMessage string
}

// ValidateAddressChange validates inputs before calling ChangeDelegatorAddress
func ValidateAddressChange(oldAddr, newAddr sdk.AccAddress) ValidationResult {
	if oldAddr.Empty() {
		return ValidationResult{
			Valid:        false,
			ErrorMessage: "old address cannot be empty",
		}
	}

	if newAddr.Empty() {
		return ValidationResult{
			Valid:        false,
			ErrorMessage: "new address cannot be empty",
		}
	}

	if oldAddr.Equals(newAddr) {
		return ValidationResult{
			Valid:        false,
			ErrorMessage: "old and new addresses cannot be the same",
		}
	}

	return ValidationResult{
		Valid:        true,
		ErrorMessage: "",
	}
}

// FormatStakingInfo provides a human-readable format of staking information
func FormatStakingInfo(info *DelegatorStakingInfo) string {
	if info == nil {
		return "No staking information"
	}

	return fmt.Sprintf(
		"Delegator: %s\n"+
			"Delegations: %d (Total Bonded: %s)\n"+
			"Unbonding Delegations: %d (Total Unbonding: %s)\n"+
			"Redelegations: %d",
		info.DelegatorAddress,
		len(info.Delegations), info.TotalBonded.String(),
		len(info.UnbondingDelegations), info.TotalUnbonding.String(),
		len(info.Redelegations),
	)
}
