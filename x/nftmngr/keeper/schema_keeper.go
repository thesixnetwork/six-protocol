package keeper

import (
	"context"
	"strconv"
	"strings"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	errormod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) CreateNftSchemaKeeper(ctx sdk.Context, creator string, schemaInput types.NFTSchemaINPUT) error {
	schemaInput.Owner = creator

	creatorAddress, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return err
	}

	// Validate Schema Message and return error if not valid
	_, err = ValidateNFTSchema(&schemaInput)
	if err != nil {
		return errormod.Wrap(types.ErrValidatingNFTSchema, err.Error())
	}

	// if mint_authorization is empty then set system to default
	if len(schemaInput.MintAuthorization) == 0 || schemaInput.MintAuthorization != types.KeyMintPermissionOnlySystem && schemaInput.MintAuthorization != types.KeyMintPermissionAll {
		schemaInput.MintAuthorization = types.KeyMintPermissionOnlySystem
	}

	// Check if the schemaInput already exists
	_, found := k.GetNFTSchema(ctx, schemaInput.Code)
	if found {
		return errormod.Wrap(types.ErrSchemaAlreadyExists, schemaInput.Code)
	}

	// check if schemaCode match with some virtual schema proposal
	err = k.checkDuplicateSchemaCodeWithVirtualSchemaProposal(ctx, schemaInput.Code)
	if err != nil {
		return err
	}

	foundOrganization, organizationName := GetOrganizationFromSchemaCode(schemaInput.Code)

	// NOTE: If there is organization in schemaInput code, check if the organization exists
	// org name must unique
	if foundOrganization {
		storedOrganization, found := k.GetOrganization(ctx, organizationName)
		if found {
			// Check owner of organization
			if storedOrganization.Owner != creator {
				return errormod.Wrap(types.ErrOrganizationOwner, creator)
			}
		} else {
			// Store organization
			k.SetOrganization(ctx, types.Organization{
				Owner: creator,
				Name:  organizationName,
			})
		}
	}

	MergeAllAttributesAndAlterOrderIndex(schemaInput.OriginData.OriginAttributes, schemaInput.OnchainData.NftAttributes, schemaInput.OnchainData.TokenAttributes)

	// parse schemaInput to NFTSchema
	schema := types.NFTSchema{
		Code:        schemaInput.Code,
		Name:        schemaInput.Name,
		Owner:       schemaInput.Owner,
		Description: schemaInput.Description,
		OriginData:  schemaInput.OriginData,
		OnchainData: &types.OnChainData{
			TokenAttributes: schemaInput.OnchainData.TokenAttributes,
			NftAttributes:   schemaInput.OnchainData.NftAttributes,
			Actions:         schemaInput.OnchainData.Actions,
			Status:          schemaInput.OnchainData.Status,
		},
		IsVerified:        schemaInput.IsVerified,
		MintAuthorization: schemaInput.MintAuthorization,
	}

	// loop over SchemaAttribute and add to nftmngr/code/name
	// schema attributes is a Global attribute of the collection,
	// which mean every metadata will shared the same value
	for _, schemaDefaultMintAttribute := range schemaInput.OnchainData.NftAttributes {
		// parse DefaultMintValue to SchemaAttributeValue
		schmaAttributeValue, err := ConvertDefaultMintValueToSchemaAttributeValue(schemaDefaultMintAttribute.DefaultMintValue)
		if err != nil {
			return errormod.Wrap(types.ErrParsingMetadataMessage, err.Error())
		}

		k.SetSchemaAttribute(ctx, types.SchemaAttribute{
			NftSchemaCode: schemaInput.Code,
			Name:          schemaDefaultMintAttribute.Name,
			DataType:      schemaDefaultMintAttribute.DataType,
			CurrentValue:  schmaAttributeValue,
			Creator:       creator,
		})
	}

	// Add the schemaInput to the store
	k.SetNFTSchema(ctx, schema)

	/*

		|------------------------------------|
		|                                    |
		|          **** ACTION ****          |
		|                                    |
		|------------------------------------|

	*/

	for i, action := range schemaInput.OnchainData.Actions {
		// check if action already exists
		_, isFound := k.GetActionOfSchema(ctx, schemaInput.Code, action.Name)
		if isFound {
			continue
		}
		k.SetActionOfSchema(ctx, types.ActionOfSchema{
			Name:          action.Name,
			NftSchemaCode: schemaInput.Code,
			Index:         uint64(i),
		})
	}

	// set action executor
	for _, actionExecutor := range schemaInput.SystemActioners {

		// validate address
		_, err := sdk.AccAddressFromBech32(actionExecutor)
		if err != nil {
			return errormod.Wrap(sdkerrors.ErrInvalidAddress, actionExecutor)
		}

		// check if actionExecutor already exists
		_, isFound := k.GetActionExecutor(ctx, schemaInput.Code, actionExecutor)
		if isFound {
			continue
		}

		val, found := k.GetExecutorOfSchema(ctx, schemaInput.Code)
		if !found {
			val = types.ExecutorOfSchema{
				NftSchemaCode:   schemaInput.Code,
				ExecutorAddress: []string{},
			}
		}

		// set executorOfSchema
		val.ExecutorAddress = append(val.ExecutorAddress, actionExecutor)

		k.SetExecutorOfSchema(ctx, types.ExecutorOfSchema{
			NftSchemaCode:   val.NftSchemaCode,
			ExecutorAddress: val.ExecutorAddress,
		})

		// set actionExecutor
		k.SetActionExecutor(ctx, types.ActionExecutor{
			Creator:         creator,
			NftSchemaCode:   schemaInput.Code,
			ExecutorAddress: actionExecutor,
		})
	}

	/*

		|------------------------------------|
		|                                    |
		|     **** ENHANCEMENT ****          |
		|                                    |
		|------------------------------------|

	*/

	// Check if schemaInput by contract already exists
	nftSchemaByContract, found := k.GetNFTSchemaByContract(ctx, schemaInput.OriginData.OriginContractAddress)
	if !found {
		nftSchemaByContract = types.NFTSchemaByContract{
			OriginContractAddress: schemaInput.OriginData.OriginContractAddress,
			SchemaCodes:           []string{schemaInput.Code},
		}
	} else {
		nftSchemaByContract.SchemaCodes = append(nftSchemaByContract.SchemaCodes, schemaInput.Code)
	}

	// Add the schemaInput code to the list of schemaInput codes
	k.SetNFTSchemaByContract(ctx, nftSchemaByContract)

	// **** SCHEMA FEE ****
	feeConfig, found := k.GetNFTFeeConfig(ctx)
	if found {
		// Get Denom
		amount, err := sdk.ParseCoinNormalized(feeConfig.SchemaFee.FeeAmount)
		if err != nil {
			return errormod.Wrap(types.ErrInvalidFeeAmount, err.Error())
		}
		feeBalances, found := k.GetNFTFeeBalance(ctx)
		if !found {
			feeBalances = types.NFTFeeBalance{
				FeeBalances: []string{
					"0" + amount.Denom,
				},
			}
		}

		if len(feeBalances.FeeBalances) > 0 {
			feeBalances.FeeBalances[types.FeeSubject_CREATE_NFT_SCHEMA] = "0" + amount.Denom
		}
		err = k.ProcessFee(ctx, &feeConfig, &feeBalances, types.FeeSubject_CREATE_NFT_SCHEMA, creatorAddress)
		if err != nil {
			return errormod.Wrap(types.ErrProcessingFee, err.Error())
		}
		// Set Fee Balance
		k.SetNFTFeeBalance(ctx, feeBalances)
	}
	// **** END SCHEMA FEE ****
	return nil
}

// Total amount of fee collected from schema_input for each block then distribute to validators // ** In the begin block it will set to 0 again
func (k Keeper) ProcessFee(ctx sdk.Context, feeConfig *types.NFTFeeConfig, feeBalances *types.NFTFeeBalance, feeSubject types.FeeSubject, source sdk.AccAddress) error {
	currentFeeBalance, _ := sdk.ParseCoinNormalized(feeBalances.FeeBalances[int32(feeSubject)])
	feeAmount, _ := sdk.ParseCoinNormalized(feeConfig.SchemaFee.FeeAmount)
	// Plus fee amount to fee balance
	currentFeeBalance = currentFeeBalance.Add(feeAmount)
	feeBalances.FeeBalances[int32(feeSubject)] = strconv.FormatInt(currentFeeBalance.Amount.Int64(), 10) + feeAmount.Denom

	err := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, source, types.ModuleName, sdk.NewCoins(feeAmount),
	)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) SetBaseURIKeeper(ctx sdk.Context, creator, nftSchemaName, baseURI string) error {
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if creator != schema.Owner {
		return errormod.Wrap(types.ErrCreatorDoesNotMatch, creator)
	}

	schema.OriginData.OriginBaseUri = baseURI
	k.SetNFTSchema(ctx, schema)

	return nil
}

func (k Keeper) SetMetadataFormatKeeper(ctx sdk.Context, creator, nftSchemaName, format string) error {
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if creator != schema.Owner {
		return errormod.Wrap(types.ErrCreatorDoesNotMatch, creator)
	}

	schema.OriginData.MetadataFormat = format

	k.SetNFTSchema(ctx, schema)

	return nil
}

func (k Keeper) SetMintAuthKeeper(ctx sdk.Context, creator, nftSchemaName string, authTo types.AuthorizeTo) error {
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if creator != schema.Owner {
		return errormod.Wrap(types.ErrCreatorDoesNotMatch, creator)
	}

	// Swith location of attribute
	switch authTo {
	case types.AuthorizeTo_SYSTEM:
		// append new nft_attributes to array of OnchainData.NftAttributes
		schema.MintAuthorization = types.KeyMintPermissionOnlySystem
	case types.AuthorizeTo_ALL:
		// append new token_attributes to array of OnchainData.TokenAttributes
		schema.MintAuthorization = types.KeyMintPermissionAll
		// end the case
	}

	k.SetNFTSchema(ctx, schema)

	return nil
}

func (k Keeper) SetOriginChainKeeper(ctx sdk.Context, creator, nftSchemaName, originChain string) error {
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if creator != schema.Owner {
		return errormod.Wrap(types.ErrCreatorDoesNotMatch, creator)
	}

	// to uppercase
	chainUpperCase := strings.ToUpper(originChain)
	schema.OriginData.OriginChain = chainUpperCase

	k.SetNFTSchema(ctx, schema)

	return nil
}

func (k Keeper) SetOriginContractKeeper(ctx sdk.Context, creator, nftSchemaName, contract string) error {
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if creator != schema.Owner {
		return errormod.Wrap(types.ErrCreatorDoesNotMatch, creator)
	}

	schema.OriginData.OriginContractAddress = contract
	k.SetNFTSchema(ctx, schema)

	return nil
}

func (k Keeper) SetURIRetrievalKeeper(ctx sdk.Context, creator, nftSchemaName string, method int32) error {
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if creator != schema.Owner {
		return errormod.Wrap(types.ErrCreatorDoesNotMatch, creator)
	}

	switch method {
	case 0:
		schema.OriginData.UriRetrievalMethod = types.URIRetrievalMethod_BASE
	case 1:
		schema.OriginData.UriRetrievalMethod = types.URIRetrievalMethod_TOKEN
	}

	k.SetNFTSchema(ctx, schema)

	return nil
}

func (k Keeper) checkDuplicateSchemaCodeWithVirtualSchemaProposal(ctx context.Context, schemaCode string) error {
	allAciveProposal := k.GetAllActiveVirtualSchemaProposal(ctx)
	for _, proposal := range allAciveProposal {
		virtualSchemaProposal, _ := k.GetVirtualSchemaProposal(ctx, proposal.Id)
		if virtualSchemaProposal.VirtualSchema.VirtualNftSchemaCode == schemaCode {
			return errormod.Wrap(types.ErrSchemaAlreadyExists, "schema code already exists in proposal: "+proposal.Id)
		}
	}
	return nil
}
