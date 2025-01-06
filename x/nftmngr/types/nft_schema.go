package types

func (m *NFTSchema) ResultWithEmptyVirtualAction() NFTSchemaQueryResult {
	result := NFTSchemaQueryResult{
		Code:              m.Code,
		Name:              m.Name,
		Owner:             m.Owner,
		Description:       m.Description,
		OriginData:        m.OriginData,
		OnchainData:       &OnChainDataResult{
			NftAttributes:   m.OnchainData.NftAttributes,
			TokenAttributes: m.OnchainData.TokenAttributes,
			Actions:         m.OnchainData.Actions,
			VirtualActions:  []*VirtualAction{},
			Status:          m.OnchainData.Status,
		},
		IsVerified:        m.IsVerified,
		MintAuthorization: m.MintAuthorization,
	}

	return result
}


func (m *NFTSchemaQueryResult) AppendVirtualAction(newVirtualAction *VirtualAction) *NFTSchemaQueryResult {
	result := m
	result.OnchainData.VirtualActions = append(result.OnchainData.VirtualActions, newVirtualAction)
	return result
}