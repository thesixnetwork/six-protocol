package sixclient

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	nftmngrmoduletypes "github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// NFTSchemaRequest represents a request to create an NFT schema
type NFTSchemaRequest struct {
	Code                string                                    `json:"code"`
	Name                string                                    `json:"name"`
	Description         string                                    `json:"description"`
	SystemActioners     []string                                  `json:"system_actioners,omitempty"`
	OnchainData         *nftmngrmoduletypes.OnChainData           `json:"onchain_data,omitempty"`
	OriginData          *nftmngrmoduletypes.OriginData            `json:"origin_data,omitempty"`
	AttributeDefinition []*nftmngrmoduletypes.AttributeDefinition `json:"attribute_definition,omitempty"`
}

// CreateNFTSchema creates a new NFT schema
func (c *SixClient) CreateNFTSchema(req NFTSchemaRequest) (*TxResponse, error) {
	msg := &nftmngrmoduletypes.MsgCreateNFTSchema{
		Creator:              c.address.String(),
		NftSchemaCode:        req.Code,
		Name:                 req.Name,
		Description:          req.Description,
		SystemActioners:      req.SystemActioners,
		OnchainData:          req.OnchainData,
		OriginData:           req.OriginData,
		AttributeDefinitions: req.AttributeDefinition,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("create NFT schema: %s", req.Code))
}

// TransferSchemaOwnership transfers ownership of an NFT schema
func (c *SixClient) TransferSchemaOwnership(schemaCode, newOwner string) (*TxResponse, error) {
	msg := &nftmngrmoduletypes.MsgChangeSchemaOwner{
		Creator:       c.address.String(),
		NftSchemaCode: schemaCode,
		NewOwner:      newOwner,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("transfer schema ownership: %s", schemaCode))
}

// AddSystemActioner adds a system actioner to an NFT schema
func (c *SixClient) AddSystemActioner(schemaCode, actioner string) (*TxResponse, error) {
	msg := &nftmngrmoduletypes.MsgAddSystemActioner{
		Creator:        c.address.String(),
		NftSchemaCode:  schemaCode,
		SystemActioner: actioner,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("add system actioner to schema: %s", schemaCode))
}

// RemoveSystemActioner removes a system actioner from an NFT schema
func (c *SixClient) RemoveSystemActioner(schemaCode, actioner string) (*TxResponse, error) {
	msg := &nftmngrmoduletypes.MsgRemoveSystemActioner{
		Creator:        c.address.String(),
		NftSchemaCode:  schemaCode,
		SystemActioner: actioner,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("remove system actioner from schema: %s", schemaCode))
}

// CreateNFTData creates NFT data within a schema
func (c *SixClient) CreateNFTData(schemaCode, tokenId string, onchainImage *nftmngrmoduletypes.OnChainData, tokenAttributes []*nftmngrmoduletypes.NftAttributeValue) (*TxResponse, error) {
	msg := &nftmngrmoduletypes.MsgCreateNftData{
		Creator:         c.address.String(),
		NftSchemaCode:   schemaCode,
		TokenId:         tokenId,
		OnchainImage:    onchainImage,
		TokenAttributes: tokenAttributes,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("create NFT data: %s/%s", schemaCode, tokenId))
}

// QueryNFTSchema retrieves information about an NFT schema
func (c *SixClient) QueryNFTSchema(schemaCode string) (*NFTSchema, error) {
	path := fmt.Sprintf("/sixprotocol/nftmngr/nft_schema/%s", schemaCode)
	body, err := c.httpGet(path)
	if err != nil {
		return nil, err
	}

	var resp struct {
		NFTSchema nftmngrmoduletypes.NFTSchema `json:"nftSchema"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal NFT schema response: %w", err)
	}

	schema := resp.NFTSchema
	return &NFTSchema{
		Code:        schema.Code,
		Name:        schema.Name,
		Description: schema.Description,
		Owner:       schema.Owner,
		IsVerified:  schema.IsVerified,
	}, nil
}

// ListNFTSchemas retrieves all NFT schemas
func (c *SixClient) ListNFTSchemas() ([]NFTSchema, error) {
	path := "/sixprotocol/nftmngr/nft_schema"
	body, err := c.httpGet(path)
	if err != nil {
		return nil, err
	}

	var resp struct {
		NFTSchemas []nftmngrmoduletypes.NFTSchema `json:"nftSchema"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal NFT schemas response: %w", err)
	}

	schemas := make([]NFTSchema, len(resp.NFTSchemas))
	for i, schema := range resp.NFTSchemas {
		schemas[i] = NFTSchema{
			Code:        schema.Code,
			Name:        schema.Name,
			Description: schema.Description,
			Owner:       schema.Owner,
			IsVerified:  schema.IsVerified,
		}
	}

	return schemas, nil
}

// QueryNFTData retrieves NFT data for a specific token
func (c *SixClient) QueryNFTData(schemaCode, tokenId string) (*nftmngrmoduletypes.NftData, error) {
	path := fmt.Sprintf("/sixprotocol/nftmngr/nft_data/%s/%s", schemaCode, tokenId)
	body, err := c.httpGet(path)
	if err != nil {
		return nil, err
	}

	var resp struct {
		NftData nftmngrmoduletypes.NftData `json:"nftData"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal NFT data response: %w", err)
	}

	return &resp.NftData, nil
}

// QueryNFTCollection retrieves NFT collection information
func (c *SixClient) QueryNFTCollection(schemaCode string) (*nftmngrmoduletypes.NftCollection, error) {
	path := fmt.Sprintf("/sixprotocol/nftmngr/nft_collection/%s", schemaCode)
	body, err := c.httpGet(path)
	if err != nil {
		return nil, err
	}

	var resp struct {
		NftCollection nftmngrmoduletypes.NftCollection `json:"nftCollection"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal NFT collection response: %w", err)
	}

	return &resp.NftCollection, nil
}

// SetNFTAttribute sets an attribute value for an NFT
func (c *SixClient) SetNFTAttribute(schemaCode, tokenId, location string, refId uint64, newValue *nftmngrmoduletypes.NftAttributeValue) (*TxResponse, error) {
	msg := &nftmngrmoduletypes.MsgSetAttributeOveriding{
		Creator:       c.address.String(),
		NftSchemaCode: schemaCode,
		TokenId:       tokenId,
		Location:      location,
		RefId:         refId,
		NewValue:      newValue,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("set NFT attribute: %s/%s", schemaCode, tokenId))
}

// PerformNFTAction performs an action on an NFT
func (c *SixClient) PerformNFTAction(schemaCode, tokenId, actionName string, refId uint64, parameters []*nftmngrmoduletypes.ActionParameter) (*TxResponse, error) {
	msg := &nftmngrmoduletypes.MsgPerformActionByNftAdmin{
		Creator:       c.address.String(),
		NftSchemaCode: schemaCode,
		TokenId:       tokenId,
		ActionName:    actionName,
		RefId:         refId,
		Parameters:    parameters,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("perform NFT action: %s/%s/%s", schemaCode, tokenId, actionName))
}

// ToggleNFTAction toggles an action's enabled/disabled state
func (c *SixClient) ToggleNFTAction(schemaCode, actionName string, disable bool) (*TxResponse, error) {
	msg := &nftmngrmoduletypes.MsgToggleAction{
		Creator:       c.address.String(),
		NftSchemaCode: schemaCode,
		ActionName:    actionName,
		Disable:       disable,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("toggle NFT action: %s/%s", schemaCode, actionName))
}

// SetBaseUri sets the base URI for an NFT schema
func (c *SixClient) SetBaseUri(schemaCode, newBaseUri string) (*TxResponse, error) {
	msg := &nftmngrmoduletypes.MsgSetBaseUri{
		Creator:       c.address.String(),
		NftSchemaCode: schemaCode,
		NewBaseUri:    newBaseUri,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("set base URI: %s", schemaCode))
}

// SetMetadataFormat sets the metadata format for an NFT schema
func (c *SixClient) SetMetadataFormat(schemaCode, newFormat string) (*TxResponse, error) {
	msg := &nftmngrmoduletypes.MsgSetMetadataFormat{
		Creator:       c.address.String(),
		NftSchemaCode: schemaCode,
		NewFormat:     newFormat,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("set metadata format: %s", schemaCode))
}

// SetOriginChain sets the origin chain for an NFT schema
func (c *SixClient) SetOriginChain(schemaCode, newOriginChain string) (*TxResponse, error) {
	msg := &nftmngrmoduletypes.MsgSetOriginChain{
		Creator:        c.address.String(),
		NftSchemaCode:  schemaCode,
		NewOriginChain: newOriginChain,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("set origin chain: %s", schemaCode))
}

// SetOriginContract sets the origin contract for an NFT schema
func (c *SixClient) SetOriginContract(schemaCode, newOriginContract string) (*TxResponse, error) {
	msg := &nftmngrmoduletypes.MsgSetOriginContract{
		Creator:           c.address.String(),
		NftSchemaCode:     schemaCode,
		NewOriginContract: newOriginContract,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("set origin contract: %s", schemaCode))
}

// SetUriRetrievalMethod sets the URI retrieval method for an NFT schema
func (c *SixClient) SetUriRetrievalMethod(schemaCode string, newMethod int32) (*TxResponse, error) {
	msg := &nftmngrmoduletypes.MsgSetUriRetrievalMethod{
		Creator:       c.address.String(),
		NftSchemaCode: schemaCode,
		NewMethod:     newMethod,
	}

	return c.BroadcastTx([]sdk.Msg{msg}, fmt.Sprintf("set URI retrieval method: %s", schemaCode))
}
