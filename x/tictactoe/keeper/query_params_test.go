package keeper_test

import (
	"testing"
	"tic-chain/x/tictactoe/types"

	"github.com/stretchr/testify/require"

	keepertest "tic-chain/testutil/keeper"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.TictactoeKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
