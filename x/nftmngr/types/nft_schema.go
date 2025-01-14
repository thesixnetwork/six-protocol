package types

import (
	"fmt"
)

// Validate checks if the NFTSchema is valid
func (m *NFTSchema) Validate() error {
	if m.Code == "" {
		return fmt.Errorf("empty schema code")
	}
	if m.Name == "" {
		return fmt.Errorf("empty schema name")
	}
	if m.Owner == "" {
		return fmt.Errorf("empty schema owner")
	}
	if m.OnchainData == nil {
		return fmt.Errorf("missing onchain data")
	}
	return nil
}

// ResultWithEmptyVirtualAction creates a query result with empty virtual actions
func (m *NFTSchema) ResultWithEmptyVirtualAction() NFTSchemaQueryResult {
	if err := m.Validate(); err != nil {
		panic(fmt.Sprintf("invalid schema: %v", err))
	}

	result := NFTSchemaQueryResult{
		Code:        m.Code,
		Name:        m.Name,
		Owner:       m.Owner,
		Description: m.Description,
		OriginData:  m.OriginData,
		OnchainData: &OnChainDataResult{
			NftAttributes:   m.OnchainData.NftAttributes,
			TokenAttributes: m.OnchainData.TokenAttributes,
			Actions:         m.OnchainData.Actions,
			VirtualActions:  make([]*VirtualAction, 0),
			Status:          m.OnchainData.Status,
		},
		IsVerified:        m.IsVerified,
		MintAuthorization: m.MintAuthorization,
	}

	return result
}

// AppendVirtualAction safely adds a virtual action to the result
func (m *NFTSchemaQueryResult) AppendVirtualAction(newAction *VirtualAction) *NFTSchemaQueryResult {
	if newAction == nil {
		return m
	}

	if m.OnchainData == nil {
		m.OnchainData = &OnChainDataResult{
			VirtualActions: make([]*VirtualAction, 0),
		}
	}

	m.OnchainData.VirtualActions = append(
		m.OnchainData.VirtualActions,
		newAction,
	)
	return m
}
