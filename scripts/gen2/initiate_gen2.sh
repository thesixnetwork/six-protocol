# Create new chain
ignite scaffold chain sixnft --no-module --address-prefix 6nft

# Create module nftmngr
ignite scaffold module nftmngr

# Create NFT Data type fot NFT Schema
ignite scaffold type OpenseaDisplayOption display_type trait_type max_value \
    --no-message \
    --no-simulation \
    --module nftmngr

# maxValue is a string, so we modify the type to be a uint64 and also change format of the fields to be under_score_case
# message OpenseaDisplayOption {
  
#   string display_type = 1; 
#   string trait_type = 2; 
#   uint64 max_value = 3; 
# }
ignite scaffold type DisplayOption opensea \
    --no-message \
    --no-simulation \
    --module nftmngr
# modify proto file to be like this

# syntax = "proto3";
# package thesixnetwork.sixnft.nftmngr;
# import "nftmngr/opensea_display_option.proto";
# option go_package = "github.com/thesixnetwork/sixnft/x/nftmngr/types";

# message DisplayOption {
  
#   OpenseaDisplayOption opensea = 1; 
# }    

# Attribute Definition
ignite scaffold type AttributeDefinition name data_type display_value_field display_option default_mint_value  \
    --no-message \
    --no-simulation \
    --module nftmngr
# Modify proto file to be like this
# syntax = "proto3";
# package thesixnetwork.sixnft.nftmngr;
# import "nftmngr/display_option.proto";
# option go_package = "github.com/thesixnetwork/sixnft/x/nftmngr/types";

# message AttributeDefinition {
  
#   string name = 1; 
#   string data_type = 2; 
#   string display_value_field = 3; 
#   DisplayOption display_option = 4; 
#   string default_mint_value = 5; 
# }

ignite scaffold type OriginData origin_chain origin_contract_address origin_base_uri attribute_overriding metadata_format origin_attributes \
    --no-message \
    --no-simulation \
    --module nftmngr
# Modify proto file to be like this
# syntax = "proto3";
# package thesixnetwork.sixnft.nftmngr;
# import "nftmngr/attribute_definition.proto";
# option go_package = "github.com/thesixnetwork/sixnft/x/nftmngr/types";

# enum AttributeOverriding {
#   ORIGIN = 0;
#   CHAIN = 1;
# }

# message OriginData {
  
#   string origin_chain = 1; 
#   string origin_contract_address = 2; 
#   string origin_base_uri = 3; 
#   AttributeOverriding  attribute_overriding = 4; 
#   string metadata_format = 5; 
#   repeated AttributeDefinition origin_attributes = 6; 
# }

ignite scaffold type Action name desc when then \
    --no-message \
    --no-simulation \
    --module nftmngr
# Modify proto file to be like this
# syntax = "proto3";
# package thesixnetwork.sixnft.nftmngr;

# option go_package = "github.com/thesixnetwork/sixnft/x/nftmngr/types";

# message Action {
  
#   string name = 1; 
#   string desc = 2; 
#   string when = 3; 
#   repeated string then = 4; 
# }
ignite scaffold type OnChainData nft_attributes token_attributes actions then \
    --no-message \
    --no-simulation \
    --module nftmngr
# Modify proto file to be like this
# syntax = "proto3";
# package thesixnetwork.sixnft.nftmngr;
# import "nftmngr/attribute_definition.proto";
# import "nftmngr/action.proto";
# option go_package = "github.com/thesixnetwork/sixnft/x/nftmngr/types";

# message OnChainData {
  
#   repeated AttributeDefinition nft_attributes = 1; 
#   repeated AttributeDefinition token_attributes = 2; 
#   repeated Action actions = 3; 
# }
ignite scaffold map NFTSchema name owner origin_data onchain_data \
    --index code \
    --no-message \
    --no-simulation \
    --module nftmngr
# Modify proto file to be like this
# syntax = "proto3";
# package thesixnetwork.sixnft.nftmngr;


# import "nftmngr/on_chain_data.proto";
# import "nftmngr/origin_data.proto";

# option go_package = "github.com/thesixnetwork/sixnft/x/nftmngr/types";

# message NFTSchema {
#   string code = 1; 
#   string name = 2; 
#   string owner = 3; 
#   OriginData origin_data = 4; 
#   OnChainData onchain_data = 5; 
  
# }
ignite scaffold message createNFTSchema nft_schema_base64 \
    --desc "To create NFT schema from JSON in Base64 format" \
    --response code \
    --no-simulation \
    --module nftmngr

# Unmarshal JSON in msg_server_xxxxx
# func (k msgServer) CreateNFTSchema(goCtx context.Context, msg *types.MsgCreateNFTSchema) (*types.MsgCreateNFTSchemaResponse, error) {
# 	ctx := sdk.UnwrapSDKContext(goCtx)

# 	jsonSchema, err := base64.StdEncoding.DecodeString(msg.NftSchemaBase64)
# 	if err != nil {
# 		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
# 	}
# 	schema := types.NFTSchema{}
# 	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(jsonSchema, &schema)
# 	if err != nil {
# 		return nil, sdkerrors.Wrap(types.ErrParsingSchemaMessage, err.Error())
# 	}

# 	fmt.Println("####################### ", schema.Code)

# 	// TODO: Handling the message
# 	_ = ctx

# 	return &types.MsgCreateNFTSchemaResponse{}, nil
# }

# Test create NFTSchema
sixnftd tx nftmngr create-nft-schema --from alice \
    <BASE64 schema>

sixnftd query nftmngr show-nft-schema \
    <NFT schema code>


ignite scaffold type Status first_mint_complete \
    --no-message \
    --no-simulation \
    --module nftmngr

ignite scaffold type OnOffSwitch active \
    --no-message \
    --no-simulation \
    --module nftmngr


# Create NFT Attributes Value
ignite scaffold type NftAttributeValue name value \
    --no-message \
    --no-simulation \
    --module nftmngr

# Create NFT Data
ignite scaffold map NftData token_owner origin_image origin_attributes onchain_attributes \
    --index nft_schema_code,token_id \
    --no-message \
    --no-simulation \
    --module nftmngr
# Modify proto file
# syntax = "proto3";
# package thesixnetwork.sixnft.nftmngr;

# import "nftmngr/nft_attribute_value.proto";
# option go_package = "github.com/thesixnetwork/sixnft/x/nftmngr/types";

# message NftData {
#   string nftSchemaCode = 1; 
#   uint64 tokenId = 2; 
#   string tokenOwner = 3; 
#   string originImage = 4; 
#   repeated NftAttributeValue originAttributes = 5; 
#   repeated NftAttributeValue onchainAttributes = 6; 
  
# }

ignite scaffold message createMetadata nft_schema_code token_id base64NFTData \
    --desc "To create NFT Metadata from Base64" \
    --response nft_schema_code,token_id \
    --no-simulation \
    --module nftmngr

#Test create metadata
sixnftd tx nftmngr create-metadata buakaw1 1 --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake \
    <BASE64 NFT Data>


# Create Action
ignite scaffold message performActionByAdmin nft_schema_code token_id action \
    --desc "To do action" \
    --response nft_schema_code,token_id \
    --no-simulation \
    --module nftmngr


# Test create action
sixnftd tx nftmngr perform-action-by-nftadmin buakaw1 1 use10percentdiscount --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake

# Create evm_support module

ignite scaffold module evm_support 

ignite scaffold map AddressBinding \
    --index eth_address:string,native_address:string \
    --no-message \
    --no-simulation \
    --module evm_support

ignite scaffold message bindAddress eth_address:string signature:string signed_message:string \
    --desc "To bind EVM address with native address" \
    --response eth_address,native_address \
    --no-simulation \
    --module evm_support

ignite scaffold message removeBinding eth_address:string signature:string signed_message:string \
    --desc "To remove binding address" \
    --response eth_address \
    --no-simulation \
    --module evm_support

# string creator = 1;
#   string nftSchemaCode = 2;
#   string tokenId = 3;
#   string action = 4;

ignite scaffold map ActionByRefId creator nft_schema_code token_id action \
    --index ref_id \
    --no-message \
    --no-simulation \
    --module nftmngr

ignite scaffold module nftoracle

ignite scaffold list MintRequest nft_schema_code:string token_id:string required_confirm:uint \
    --no-message \
    --no-simulation \
    --module nftoracle


ignite scaffold message createMintRequest nft_schema_code:string token_id:string required_confirm:uint \
    --desc "To create Mint Request" \
    --response nft_schema_code,token_id \
    --no-simulation \
    --module nftoracle

# To create token module
ignite scaffold module admin

ignite scaffold type permissions \
    --no-message \
    --no-simulation \
    --module admin

ignite scaffold single authorization root_admin:string permissions:Permissions \
    --no-message \
    --no-simulation \
    --module admin

ignite scaffold message grantPermission name grantee \
    --desc "To grant permission" \
    --response grantee \
    --no-simulation \
    --module admin

ignite scaffold message revokePermission name revokee \
    --desc "To revoke permission" \
    --response revokee \
    --no-simulation \
    --module admin

ignite scaffold message mint amount:uint token:string \
    --desc "To mint token" \
    --response amount,token \
    --no-simulation \
    --module admin

ignite scaffold message burn amount:uint token:string \
    --desc "To burn token" \
    --response amount,token \
    --no-simulation \
    --module admin


ignite scaffold map Organization owner \
    --index name \
    --no-message \
    --no-simulation \
    --module nftmngr

ignite scaffold message submitMintResponse mintRequestID:uint base64NftData:string \
    --desc "To submit mint response" \
    --response mintRequestID \
    --no-simulation \
    --module nftoracle
ignite scaffold message setNFTAttribute nft_schema_code attribute_value:NftAttributeValue \
    --desc "To set NFT attribute" \
    --response nft_schema_code,attribute_name,attribute_value \
    --no-simulation \
    --module nftmngr


ignite scaffold list ActionRequest nft_schema_code:string token_id:string required_confirm:uint \
    --no-message \
    --no-simulation \
    --module nftoracle

ignite scaffold message createActionRequest vm:string base64_action_signature:string \
    --desc "To create Action Request" \
    --response request_id \
    --no-simulation \
    --module nftoracle

ignite scaffold message submitActionResponse actionRequestID:uint base64NftData:string \
    --desc "To submit action response" \
    --response actionRequestID \
    --no-simulation \
    --module nftoracle


ignite scaffold message changeSchemaOwner nft_schema_code:string new_owner:string \
    --desc "To update schema owner" \
    --response nft_schema_code,new_owner \
    --no-simulation \
    --module nftmngr

ignite scaffold list CollectionOwnerRequest nft_schema_code:string base64_owner_signature:string \
    --no-message \
    --no-simulation \
    --module nftoracle

ignite scaffold message createVerifyCollectionOwnerRequest nft_schema_code:string required_confirm:uint \          
    --desc "To create Verify Collection Owner Request" \
    --response id:uint,nft_schema_code:string,ownerAddress:string \
    --no-simulation \
    --module nftoracle

ignite scaffold message SubmitVerifyCollectionOwner verifyRequestID:uint schema_code:string contractOwnerAddress:string \
    --desc "To Sumint Verify Collection Owner" \
    --response verifyRequestID:uint \
    --no-simulation \
    --module nftoracle


ignite scaffold message resyncAttributes nft_schema_code:string token_id:string \
    --desc "To resync onchain attributes" \
    --response nft_schema_code \
    --no-simulation \
    --module nftmngr

ignite scaffold single OracleConfig \
    --no-message \
    --no-simulation \
    --module nftoracle


ignite scaffold message setMinimumConfirmation new_confirmation \
    --desc "To set minimum confirmation" \
    --response new_confirmation \
    --no-simulation \
    --module nftoracle

ignite scaffold map NFTSchemaByContract schema_codes:string \
    --index origin_contract_address,chain \
    --no-message \
    --no-simulation \
    --module nftmngr

ignite scaffold single NFTFeeConfig \
    --no-message \
    --no-simulation \
    --module nftmngr

ignite scaffold message setFeeConfig new_fee_config_base64:string \
    --desc "To set fee config" \
    --response new_fee_config \
    --no-simulation \
    --module nftmngr

ignite scaffold single NFTFeeBalance \
    --no-message \
    --no-simulation \
    --module nftmngr

ignite scaffold map actionActor \
    create_at:string expire_at:string \
    --index actorAddress:string,ownerAddress:string  \
    --module nftoracle --no-simulation

# create multiple metadatas in one message
ignite scaffold message createMultiMetadata nft_schema_code token_id base64NFTData \                          
    --desc "To create Multiple NFT Metadata(s) from Base64 required flag in base64" \
    --response nft_schema_code,token_id \
    --no-simulation \
    --module nftmngr

## modify proto/nftmngr/tx.proto
## from
# message MsgCreateMultiMetadata {
#   string creator = 1;
#   string nftSchemaCode = 2;
#   string tokenId = 3;
#   string base64NFTData = 4;
# }

## to
# message MsgCreateMultiMetadata {
#   string creator = 1;
#   string nftSchemaCode = 2;
#   repeated strin tokenId = 3;
#   string base64NFTData = 4;
# }

# run ignite g proto-go -y