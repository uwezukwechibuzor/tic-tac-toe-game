package keeper

import (
	"context"
	"fmt"
	"strconv"
	"tic-chain/x/tictactoe/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PlayMove(goCtx context.Context, msg *types.MsgPlayMove) (*types.MsgPlayMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currMove := ""
	nextMove := ""
	newBoard := ""
	currentPlayer := ""

	// check if game exists
	getGame, found := k.GetGame(ctx, msg.GameId)
	if !found {
		return nil, fmt.Errorf("game does not exist: %v", msg.GameId)
	}

	// check that playerAddress making a move signaled to play the game
	// this prevents other players from participating in this game
	if getGame.FirstPlayer.PlayerId != msg.PlayerAddress && getGame.SecondPlayer.PlayerId != msg.PlayerAddress {
		return nil, fmt.Errorf("player did not join this game: %v", msg.PlayerAddress)
	}

	// check that gaame status is still in progress
	// TODO: check the status is not still in the open stage
	if getGame.Status != "Inprogress" {
		return nil, fmt.Errorf("game play time has ended: %v", getGame.Status)
	}

	// get player to play
	if getGame.NextMove == "X" {
		currentPlayer = getGame.FirstPlayer.PlayerId
		currMove = getGame.FirstPlayer.Symbol
		nextMove = getGame.SecondPlayer.Symbol
	} else {
		currentPlayer = getGame.SecondPlayer.PlayerId
		currMove = getGame.SecondPlayer.Symbol
		nextMove = getGame.FirstPlayer.Symbol
	}

	// check player out of turn
	if msg.PlayerAddress != currentPlayer {
		return nil, fmt.Errorf("not player's turn to make a move")
	}

	// check move format
	position, err := strconv.Atoi(msg.Move)
	if err != nil {
		return nil, err
	}
	if (position < 1) || (position > 9) {
		return nil, fmt.Errorf("incorrect move format: numbers must be between 0 and 9: %v", msg.Move)
	}

	// check move validity and update
	if string(getGame.Board[position-1]) != " " {
		return nil, fmt.Errorf("incorrect move, positionition already taken: %v", getGame.Board)
	} else {
		newBoard = getGame.Board[:position-1] + currMove + getGame.Board[position:]
	}

	// process the player's game status
	finished := processGameStatus(newBoard, &getGame)
	if finished {
		getGame.Status = "finished"
	}

	// Update points and leaderboard based on game outcome
	if err := k.processGameResult(ctx, &getGame); err != nil {
		return nil, err
	}
	// set game data to store
	game := types.Game{
		GameId:       getGame.GetGameId(),
		AdminAddress: getGame.AdminAddress,
		Status:       getGame.Status,
		Board:        newBoard,
		Winner:       getGame.Winner,
		NextMove:     nextMove,
		FirstPlayer:  getGame.FirstPlayer,
		SecondPlayer: getGame.SecondPlayer,
	}

	k.SetGame(ctx, game)

	return &types.MsgPlayMoveResponse{}, nil
}
