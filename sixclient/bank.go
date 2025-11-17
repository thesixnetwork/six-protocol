package sixclient

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// SendTokens sends tokens from the client's address to another address
func (c *SixClient) SendTokens(toAddress string, amount math.Int, denom string) (*TxResponse, error) {
	toAddr, err := sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		return nil, err
	}

	msg := banktypes.NewMsgSend(
		c.address,
		toAddr,
		sdk.NewCoins(sdk.NewCoin(denom, amount)),
	)

	return c.BroadcastTx([]sdk.Msg{msg}, "send tokens via SixClient")
}

// SendSIX sends SIX tokens (convenience method)
func (c *SixClient) SendSIX(toAddress string, amount math.Int) (*TxResponse, error) {
	return c.SendTokens(toAddress, amount, Denom)
}

// SendMultiple sends tokens to multiple recipients in a single transaction
func (c *SixClient) SendMultiple(recipients []SendRequest) (*TxResponse, error) {
	var inputs []banktypes.Input
	var outputs []banktypes.Output

	// Calculate total amount needed
	totalAmount := sdk.NewCoins()
	for _, req := range recipients {
		totalAmount = totalAmount.Add(sdk.NewCoin(req.Denom, req.Amount))
	}

	// Create input
	inputs = append(inputs, banktypes.Input{
		Address: c.address.String(),
		Coins:   totalAmount,
	})

	// Create outputs
	for _, req := range recipients {
		outputs = append(outputs, banktypes.Output{
			Address: req.ToAddress,
			Coins:   sdk.NewCoins(sdk.NewCoin(req.Denom, req.Amount)),
		})
	}

	msg := banktypes.NewMsgMultiSend(inputs, outputs)
	return c.BroadcastTx([]sdk.Msg{msg}, "multi-send via SixClient")
}

// SendRequest represents a send request for multi-send operations
type SendRequest struct {
	ToAddress string   `json:"to_address"`
	Amount    math.Int `json:"amount"`
	Denom     string   `json:"denom"`
}

// NewSendRequest creates a new send request
func NewSendRequest(toAddress string, amount math.Int, denom string) SendRequest {
	return SendRequest{
		ToAddress: toAddress,
		Amount:    amount,
		Denom:     denom,
	}
}

// NewSIXSendRequest creates a new SIX send request
func NewSIXSendRequest(toAddress string, amount math.Int) SendRequest {
	return SendRequest{
		ToAddress: toAddress,
		Amount:    amount,
		Denom:     Denom,
	}
}

// SIXToUsix converts SIX amount to usix (multiply by 1,000,000)
func SIXToUsix(sixAmount int64) math.Int {
	return math.NewInt(sixAmount).Mul(math.NewInt(1000000))
}

// UsixToSIX converts usix amount to SIX (divide by 1,000,000)
func UsixToSIX(usixAmount math.Int) math.Int {
	return usixAmount.Quo(math.NewInt(1000000))
}

// GetSIXBalance returns the SIX balance in SIX units (not usix)
func (c *SixClient) GetSIXBalance() (math.Int, error) {
	balance, err := c.GetBalance(Denom)
	if err != nil {
		return math.ZeroInt(), err
	}
	return UsixToSIX(balance.Amount), nil
}
