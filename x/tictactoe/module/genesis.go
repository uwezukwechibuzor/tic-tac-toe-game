package tictactoe

import (
	"strings"
	"tic-chain/x/tictactoe/keeper"
	"tic-chain/x/tictactoe/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the game
	for _, elem := range genState.GameList {
		k.SetGame(ctx, elem)
	}

	adminAddr := sdk.AccAddress{}
	var err error
	if len(strings.TrimSpace(genState.AdminAddress)) != 0 {
		adminAddr, err = sdk.AccAddressFromBech32(genState.AdminAddress)
		if err != nil {
			panic(err)
		}
	}

	k.SetAdmin(ctx, adminAddr)

	// Set game count
	k.SetGameCount(ctx, genState.GameCount)
	// Set all the leaderboard
	for _, elem := range genState.LeaderboardList {
		k.SetLeaderboard(ctx, elem)
	}

	// Set leaderboard count
	k.SetLeaderboardCount(ctx, genState.LeaderboardCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AdminAddress = string(k.GetAdmin(ctx))
	genesis.GameList = k.GetAllGame(ctx)
	genesis.GameCount = k.GetGameCount(ctx)
	// genesis.LeaderboardList = k.GetAllLeaderboard(ctx)
	genesis.LeaderboardCount = k.GetLeaderboardCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
