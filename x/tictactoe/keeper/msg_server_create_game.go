package keeper

import (
	"context"
	"fmt"
	"tic-chain/x/tictactoe/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreateGame handles the transactions of creating new Tic Tac Toe game on the chain.
// It verifies whether the account has the necessary permission to create games
func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get the game admin address set at genesis
	adminAddress := k.GetAdmin(ctx)

	// Convert msg.Creator to sdk.AccAddress
	gameCreatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("failed to convert msg.Creator: %v", err)
	}

	// Compare adminAddress and gameCreatorAddress
	if !adminAddress.Equals(gameCreatorAddress) {
		return nil, fmt.Errorf("address has no permission to create games")
	}

	// get the game count
	gameID := k.GetGameCount(ctx)

	// Create game instance
	game := types.Game{
		GameId:       gameID,
		AdminAddress: msg.Creator,
		Status:       "open",
		Board:        "",
		NextMove:     "",
		Winner:       "null",
		FirstPlayer:  &types.Player{},
		SecondPlayer: &types.Player{},
	}

	// Increment and set the updated game count
	k.SetGameCount(ctx, gameID+1)

	// Set the newly created game to store
	k.SetGame(ctx, game)

	return &types.MsgCreateGameResponse{}, nil
}
