package keeper

import (
	"context"
	"fmt"
	"tic-chain/x/tictactoe/types"
)

// processGameResult updates player scores and the leaderboard based on the result of the game.
// It takes the context and the game state as input arguments.
// processGameResult is responsible for updating player scores and the leaderboard based on the outcome of a game.
// It adjusts the points of the players according to whether the game was won by 'X', won by 'O', or resulted in a draw.
func (k msgServer) processGameResult(ctx context.Context, game *types.Game) error {
	switch game.Winner {
	case "X":
		game.FirstPlayer.Points += 3
	case "O":
		game.SecondPlayer.Points += 3
	case "draw":
		game.FirstPlayer.Points++
		game.SecondPlayer.Points++
	}

	// Update the leaderboard
	var leaderboard types.Leaderboard
	leaderboard.Players = append(leaderboard.Players, game.FirstPlayer, game.SecondPlayer)

	// Set leaderboard details to store
	if err := k.SetLeaderboard(ctx, leaderboard); err != nil {
		return fmt.Errorf("failed to update leaderboard: %v", err)
	}

	return nil
}
