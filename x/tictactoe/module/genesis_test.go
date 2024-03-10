package tictactoe_test

import (
	"testing"
	"tic-chain/testutil/nullify"
	"tic-chain/x/tictactoe/types"

	keepertest "tic-chain/testutil/keeper"

	tictactoe "tic-chain/x/tictactoe/module"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		LeaderboardList: []types.Leaderboard{
			{},
			{},
		},
		LeaderboardCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TictactoeKeeper(t)
	tictactoe.InitGenesis(ctx, k, genesisState)
	got := tictactoe.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LeaderboardList, got.LeaderboardList)
	require.Equal(t, genesisState.LeaderboardCount, got.LeaderboardCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
