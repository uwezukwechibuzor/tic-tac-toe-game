package keeper

import (
	"context"
	"fmt"
	"tic-chain/x/tictactoe/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// JoinGameX handles the transactions of joining to play the Tic Tac Toe game on the chain.
func (k msgServer) JoinGameX(goCtx context.Context, msg *types.MsgJoinGameX) (*types.MsgJoinGameXResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if game exists
	getGame, found := k.GetGame(ctx, msg.GameId)
	if !found {
		return nil, fmt.Errorf("game does not exist: %v", msg.GameId)
	}

	// admin address is not allowed to participate in games
	// get the game admin address
	adminAddress := k.GetAdmin(ctx)

	// Convert msg.PlayerAddress to sdk.AccAddress
	playerAddress, err := sdk.AccAddressFromBech32(msg.PlayerAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to convert msg.PlayerAddress: %v", err)
	}

	// Compare adminAddress and playerAddress
	if adminAddress.Equals(playerAddress) {
		return nil, fmt.Errorf("this address is not allowed to play games")
	}

	// check if game is still open
	if getGame.Status != "open" {
		return nil, fmt.Errorf("game is already in progress")
	}

	// check for if this player's position is available
	if getGame.FirstPlayer.PlayerId != "" {
		return nil, fmt.Errorf("there is an existing active player")
	}

	// check that same address is not player X and player O
	if getGame.SecondPlayer != nil && getGame.SecondPlayer.PlayerId == msg.PlayerAddress {
		return nil, fmt.Errorf("same player can not be player X and player O")
	}

	// set game details to store
	game := types.Game{
		GameId:       getGame.GetGameId(),
		AdminAddress: getGame.AdminAddress,
		Status:       "open",
		Board:        "         ",
		NextMove:     "X",
		Winner:       "",
		FirstPlayer: &types.Player{
			PlayerId: msg.PlayerAddress,
			Symbol:   "X",
			Points:   0,
		},
		SecondPlayer: getGame.SecondPlayer,
	}

	// change the game status after both players are active
	if game.SecondPlayer.PlayerId != "" {
		game.Status = "Inprogress"
	}

	// Make payment for joining the game
	if err := k.makePayment(ctx, playerAddress); err != nil {
		return nil, fmt.Errorf("failed to make payment: %v", err)
	}

	// set updated game details to store
	k.SetGame(ctx, game)

	return &types.MsgJoinGameXResponse{}, nil
}

// JoinGameX handles the transactions of joining to play the Tic Tac Toe game on the chain as the Opponent
func (k msgServer) JoinGameO(goCtx context.Context, msg *types.MsgJoinGameO) (*types.MsgJoinGameOResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if game exists
	getGame, found := k.GetGame(ctx, msg.GameId)
	if !found {
		return nil, fmt.Errorf("game does not exist: %v", msg.GameId)
	}

	// admin address is not allowed to participate in games
	// get the game admin address
	adminAddress := k.GetAdmin(ctx)

	// Convert msg.PlayerAddress to sdk.AccAddress
	playerAddress, err := sdk.AccAddressFromBech32(msg.PlayerAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to convert msg.PlayerAddress: %v", err)
	}

	// Compare adminAddress and playerAddress
	if adminAddress.Equals(playerAddress) {
		return nil, fmt.Errorf("this address is not allowed to play games")
	}

	// check if game is still open
	if getGame.Status != "open" {
		return nil, fmt.Errorf("game is already in progress")
	}

	// check for if this player's position is available
	if getGame.SecondPlayer.PlayerId != "" {
		return nil, fmt.Errorf("there is an existing active player")
	}

	// check that same address is not player X and player O
	if getGame.FirstPlayer != nil && getGame.FirstPlayer.PlayerId == msg.PlayerAddress {
		return nil, fmt.Errorf("same player can not be player X and player O")
	}

	// set game details to store
	game := types.Game{
		GameId:       getGame.GetGameId(),
		AdminAddress: getGame.AdminAddress,
		Status:       "open",
		Board:        "         ",
		NextMove:     getGame.NextMove,
		Winner:       "",
		SecondPlayer: &types.Player{
			PlayerId: msg.PlayerAddress,
			Symbol:   "O",
			Points:   0,
		},
		FirstPlayer: getGame.FirstPlayer,
	}

	// change the game status after both players are active
	if game.FirstPlayer.PlayerId != "" {
		game.Status = "Inprogress"
	}

	// Make payment for joining the game
	if err := k.makePayment(ctx, playerAddress); err != nil {
		return nil, fmt.Errorf("failed to make payment: %v", err)
	}

	// set updated game details to store
	k.SetGame(ctx, game)

	return &types.MsgJoinGameOResponse{}, nil
}
