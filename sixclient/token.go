package sixclient

import (
	"encoding/json"
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tokenmngrmoduletypes "github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

// TokenRequest represents a request to create a custom token
type TokenRequest struct {
	Name         string   `json:"name"`
	Base         string   `json:"base"`
	MaxSupply    math.Int `json:"max_supply"`
	Mintee       string   `json:"mintee"`
	Description  string   `json:"description,omitempty"`
	TokenURI     string   `json:"token_uri,omitempty"`
	BurnRate     string   `json:"burn_rate,omitempty"`
	BurnFromPool bool     `json:"burn_from_pool,omitempty"`
}

// TokenBurnRequest represents a request to burn tokens
type TokenBurnRequest struct {
	Name   string   `json:"name"`
	Amount math.Int `json:"amount"`
}

// TokenMintRequest represents a request to mint tokens
type TokenMintRequest struct {
	Name   string   `json:"name"`
	Amount math.Int `json:"amount"`
	Mintee string   `json:"mintee"`
}

// CreateToken creates a new custom token
func (c *SixClient) CreateToken(req TokenRequest) (*TxResponse, error) {
	msg := &tokenmngrmoduletypes.MsgCreateToken{
		Creator:      c.address.String(),
		Name:         req.Name,
		Base:         req.Base,
		MaxSupply:    req.MaxSupply,
		Mintee:       req.Mintee,
		Description:  req.Description,
		TokenURI:     req.TokenURI,
		BurnRate:     req.BurnRate,
		BurnFromPool: req.BurnFromPool,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("create token: %s", req.Name))
}

// MintToken mints tokens to a specified address
func (c *SixClient) MintToken(req TokenMintRequest) (*TxResponse, error) {
	msg := &tokenmngrmoduletypes.MsgMint{
		Creator: c.address.String(),
		Name:    req.Name,
		Amount:  req.Amount,
		Mintee:  req.Mintee,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("mint token: %s", req.Name))
}

// BurnToken burns tokens from the creator's address
func (c *SixClient) BurnToken(req TokenBurnRequest) (*TxResponse, error) {
	msg := &tokenmngrmoduletypes.MsgBurn{
		Creator: c.address.String(),
		Name:    req.Name,
		Amount:  req.Amount,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("burn token: %s", req.Name))
}

// GrantBurnFrom grants burn permission to another address
func (c *SixClient) GrantBurnFrom(tokenName, grantee string) (*TxResponse, error) {
	msg := &tokenmngrmoduletypes.MsgGrantBurnFrom{
		Creator: c.address.String(),
		Name:    tokenName,
		Grantee: grantee,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("grant burn permission: %s to %s", tokenName, grantee))
}

// RevokeBurnFrom revokes burn permission from an address
func (c *SixClient) RevokeBurnFrom(tokenName, revokee string) (*TxResponse, error) {
	msg := &tokenmngrmoduletypes.MsgRevokeBurnFrom{
		Creator: c.address.String(),
		Name:    tokenName,
		Revokee: revokee,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("revoke burn permission: %s from %s", tokenName, revokee))
}

// GrantMintee grants minting permission to another address
func (c *SixClient) GrantMintee(tokenName, grantee string) (*TxResponse, error) {
	msg := &tokenmngrmoduletypes.MsgGrantMintee{
		Creator: c.address.String(),
		Name:    tokenName,
		Grantee: grantee,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("grant mint permission: %s to %s", tokenName, grantee))
}

// RevokeMintee revokes minting permission from an address
func (c *SixClient) RevokeMintee(tokenName, revokee string) (*TxResponse, error) {
	msg := &tokenmngrmoduletypes.MsgRevokeMintee{
		Creator: c.address.String(),
		Name:    tokenName,
		Revokee: revokee,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("revoke mint permission: %s from %s", tokenName, revokee))
}

// SetTokenURI updates the token URI
func (c *SixClient) SetTokenURI(tokenName, newURI string) (*TxResponse, error) {
	msg := &tokenmngrmoduletypes.MsgSetTokenURI{
		Creator:  c.address.String(),
		Name:     tokenName,
		TokenURI: newURI,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("set token URI: %s", tokenName))
}

// SetBurnRate updates the burn rate for a token
func (c *SixClient) SetBurnRate(tokenName, newRate string) (*TxResponse, error) {
	msg := &tokenmngrmoduletypes.MsgSetBurnRate{
		Creator:  c.address.String(),
		Name:     tokenName,
		BurnRate: newRate,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("set burn rate: %s", tokenName))
}

// QueryToken retrieves information about a custom token
func (c *SixClient) QueryToken(tokenName string) (*Token, error) {
	path := fmt.Sprintf("/sixprotocol/tokenmngr/token/%s", tokenName)
	body, err := c.httpGet(path)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Token tokenmngrmoduletypes.Token `json:"token"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal token response: %w", err)
	}

	token := resp.Token
	return &Token{
		Name:      token.Name,
		Base:      token.Base,
		Creator:   token.Creator,
		MaxSupply: token.MaxSupply,
	}, nil
}

// ListTokens retrieves all custom tokens (using fallback to cosmos bank supply)
func (c *SixClient) ListTokens() ([]Token, error) {
	// First try the sixprotocol endpoint
	path := "/sixprotocol/tokenmngr/token"
	body, err := c.httpGet(path)
	if err == nil {
		var resp struct {
			Tokens []tokenmngrmoduletypes.Token `json:"token"`
		}

		if err := json.Unmarshal(body, &resp); err == nil && len(resp.Tokens) > 0 {
			tokens := make([]Token, len(resp.Tokens))
			for i, token := range resp.Tokens {
				tokens[i] = Token{
					Name:      token.Name,
					Base:      token.Base,
					Creator:   token.Creator,
					MaxSupply: token.MaxSupply,
				}
			}
			return tokens, nil
		}
	}

	// Fallback to cosmos bank supply endpoint
	path = "/cosmos/bank/v1beta1/supply"
	body, err = c.httpGet(path)
	if err != nil {
		return nil, fmt.Errorf("failed to query tokens: %w", err)
	}

	var supplyResp struct {
		Supply []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"supply"`
	}

	if err := json.Unmarshal(body, &supplyResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal supply response: %w", err)
	}

	// Convert supply data to token format
	tokens := make([]Token, len(supplyResp.Supply))
	for i, coin := range supplyResp.Supply {
		amount, ok := math.NewIntFromString(coin.Amount)
		if !ok {
			amount = math.ZeroInt()
		}
		tokens[i] = Token{
			Name:      coin.Denom,
			Base:      coin.Denom,
			Creator:   "system",
			MaxSupply: amount, // Current supply as max supply fallback
		}
	}

	return tokens, nil
}

// QueryTokenBurns retrieves burn information for a token
func (c *SixClient) QueryTokenBurns(tokenName string) ([]tokenmngrmoduletypes.TokenBurn, error) {
	path := fmt.Sprintf("/sixprotocol/tokenmngr/token_burn/%s", tokenName)
	body, err := c.httpGet(path)
	if err != nil {
		return nil, err
	}

	var resp struct {
		TokenBurns []tokenmngrmoduletypes.TokenBurn `json:"tokenBurn"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal token burns response: %w", err)
	}

	return resp.TokenBurns, nil
}

// QueryMinteeGrants retrieves mintee grants for a token
func (c *SixClient) QueryMinteeGrants(tokenName string) ([]string, error) {
	path := fmt.Sprintf("/sixprotocol/tokenmngr/mintee/%s", tokenName)
	body, err := c.httpGet(path)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Mintees []string `json:"mintee"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal mintee response: %w", err)
	}

	return resp.Mintees, nil
}

// QueryBurnGrants retrieves burn grants for a token
func (c *SixClient) QueryBurnGrants(tokenName string) ([]string, error) {
	path := fmt.Sprintf("/sixprotocol/tokenmngr/burn_from/%s", tokenName)
	body, err := c.httpGet(path)
	if err != nil {
		return nil, err
	}

	var resp struct {
		BurnFroms []string `json:"burnFrom"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal burn grants response: %w", err)
	}

	return resp.BurnFroms, nil
}

// GetTokenSupply retrieves the current supply of a token
func (c *SixClient) GetTokenSupply(denom string) (math.Int, error) {
	path := fmt.Sprintf("/cosmos/bank/v1beta1/supply/by_denom?denom=%s", denom)
	body, err := c.httpGet(path)
	if err != nil {
		return math.ZeroInt(), err
	}

	var resp struct {
		Amount struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"amount"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return math.ZeroInt(), fmt.Errorf("failed to unmarshal supply response: %w", err)
	}

	amount, ok := math.NewIntFromString(resp.Amount.Amount)
	if !ok {
		return math.ZeroInt(), fmt.Errorf("invalid amount format: %s", resp.Amount.Amount)
	}

	return amount, nil
}

// GetTotalSupply retrieves the total supply of all tokens
func (c *SixClient) GetTotalSupply() ([]Balance, error) {
	path := "/cosmos/bank/v1beta1/supply"
	body, err := c.httpGet(path)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Supply []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"supply"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal total supply response: %w", err)
	}

	supplies := make([]Balance, len(resp.Supply))
	for i, coin := range resp.Supply {
		amount, ok := math.NewIntFromString(coin.Amount)
		if !ok {
			amount = math.ZeroInt()
		}
		supplies[i] = Balance{
			Denom:  coin.Denom,
			Amount: amount,
		}
	}

	return supplies, nil
}
