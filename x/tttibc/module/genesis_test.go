package tttibc_test

import (
	"testing"
	"tic-chain/testutil/nullify"
	"tic-chain/x/tttibc/types"

	keepertest "tic-chain/testutil/keeper"

	tttibc "tic-chain/x/tttibc/module"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		GameList: []types.Game{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		GameCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TttibcKeeper(t)
	tttibc.InitGenesis(ctx, k, genesisState)
	got := tttibc.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.GameList, got.GameList)
	require.Equal(t, genesisState.GameCount, got.GameCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
