package keeper_test

import (
	"testing"
	"tic-chain/x/tictactoe/types"

	"github.com/stretchr/testify/require"

	keepertest "tic-chain/testutil/keeper"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TictactoeKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
