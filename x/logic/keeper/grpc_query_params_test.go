package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/okp4/okp4d/testutil/keeper"
	"github.com/okp4/okp4d/x/logic/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.LogicKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryServiceParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryServiceParamsResponse{Params: params}, response)
}
