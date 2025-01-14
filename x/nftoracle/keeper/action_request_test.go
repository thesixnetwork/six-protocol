package keeper_test

import (
	"encoding/base64"
	"strconv"
	"testing"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"

	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func createNActionRequest(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ActionOracleRequest {
	items := make([]types.ActionOracleRequest, n)
	for i := range items {
		items[i].Id = keeper.AppendActionRequest(ctx, items[i])
	}
	return items
}

func TestActionRequestGet(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionRequest(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetActionRequest(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestActionRequestRemove(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionRequest(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveActionRequest(ctx, item.Id)
		_, found := keeper.GetActionRequest(ctx, item.Id)
		require.False(t, found)
	}
}

func TestActionRequestGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionRequest(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllActionRequest(ctx)),
	)
}

func TestActionRequestCount(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionRequest(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetActionRequestCount(ctx))
}

func TestCreateActionRequest(t *testing.T) {
	k, _ := keepertest.NftoracleKeeper(t)
	actionSig := types.ActionSignature{
		Message:   "eyAibmZ0X3NjaGVtYV9jb2RlIjogIm1ocnMubWhycnMxIiwgInRva2VuX2lkIjogIjEiLCAiYWN0aW9uIjogInVzZUJlZXJDb3Vwb24iLCAiZXhwaXJlZF9hdCI6ICIyMDIyLTEwLTMxVDAwOjAwOjAwLjAwMFoiIH0=",
		Signature: "0xf3ea77a85d9b105dc8e46d92c129df743aa3b479fb30993657ddba20e6913f2d1003cd0841ae0f4a045d80747f3970ad068f19072855cf0593bab5fb6082fbe901",
	}
	_, _, err := ValidateActionSignature(k, actionSig)
	require.NoError(t, err)
}

func ValidateActionSignature(k *keeper.Keeper, actionSig types.ActionSignature) (*types.ActionOracleParam, *string, error) {
	sign_msg := "\x19Ethereum Signed Message:\n" + strconv.FormatInt(int64(len(actionSig.Message)), 10) + actionSig.Message

	data := []byte(sign_msg)
	hash := crypto.Keccak256Hash(data)
	hash_bytes := hash.Bytes()

	actionParam := &types.ActionOracleParam{}
	actionParamBz, err := base64.StdEncoding.DecodeString(actionSig.Message)
	if err != nil {
		return nil, nil, err
	}
	err = k.CDC.(*codec.ProtoCodec).UnmarshalJSON(actionParamBz, actionParam)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(types.ErrParsingActionParam, err.Error())
	}

	// validate signature format
	decode_signature, err := hexutil.Decode(actionSig.Signature)
	if err != nil {
		// log.Fatalf("Failed to decode signature: %v", msg.Signature)
		return nil, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid signature")
	}
	signature_with_revocery_id := decode_signature
	// remove last byte coz is is recovery id
	// decode_signature[64] -= 27 // this on should be checked whether it can be a weak point later // remove recovery id

	// get pulic key from signature
	sigPublicKey, err := crypto.Ecrecover(hash_bytes, decode_signature) // recover publickey from signature and hash
	if err != nil {
		return nil, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid signature or message")
	}

	// get address from public key
	pubEDCA, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "faild to unmarshal public key")
	}
	eth_address_from_pubkey := crypto.PubkeyToAddress(*pubEDCA)

	signatureNoRecoverID := signature_with_revocery_id[:len(signature_with_revocery_id)-1] // remove recovery id
	if verified := crypto.VerifySignature(sigPublicKey, hash.Bytes(), signatureNoRecoverID); !verified {
		return nil, nil, sdkerrors.Wrap(types.ErrVerifyingSignature, "invalid signature")
	}
	signer := eth_address_from_pubkey.Hex()
	return actionParam, &signer, nil
}
