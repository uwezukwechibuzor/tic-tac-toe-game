package keeper

import (
	"context"
	"fmt"
	"tic-chain/x/tictactoe/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteGame(goCtx context.Context, msg *types.MsgDeleteGame) (*types.MsgDeleteGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the game exists
	game, found := k.GetGame(ctx, msg.GameId)
	if !found {
		return nil, fmt.Errorf("game id not found: %v", msg.GameId)
	}

	// Checks that the game creator which is the admin
	// can only delete the game
	// the game was created by the admin
	if msg.AdminAddress != game.AdminAddress {
		return nil, fmt.Errorf("this address can not delete a game: %v", msg.AdminAddress)
	}

	// the game cannot be deleted if it is in progress
	if game.Status == "Inprogress" {
		return nil, fmt.Errorf("game in progress, can not be deleted: %v", game.Status)
	}

	// the game cannot be deleted if it has finished
	if game.Status == "finished" {
		return nil, fmt.Errorf("game has ended, can not be deleted: %v", game.Status)
	}

	k.RemoveGame(ctx, msg.GameId)

	return &types.MsgDeleteGameResponse{}, nil
}
