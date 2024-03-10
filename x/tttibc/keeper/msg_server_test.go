package keeper_test

import (
	"context"
	"testing"
	"tic-chain/x/tttibc/keeper"
	"tic-chain/x/tttibc/types"

	"github.com/stretchr/testify/require"

	keepertest "tic-chain/testutil/keeper"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.TttibcKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
