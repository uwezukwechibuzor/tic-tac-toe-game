package keeper

import (
	"tic-chain/x/tictactoe/types"
)

// processGameStatus processes the game status based on the new board state.
// It checks for winning combinations and updates the game's winner accordingly.
// If no winning combination is found and the board is full, the game is declared as a draw.
// Returns true if the game is over (either a win or draw), otherwise returns false.
func processGameStatus(newBoard string, game *types.Game) bool {
	// Possible winning combinations
	possibleWins := [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}

	// Flag to check if the board is full
	boardFull := true

	// Check for winning combinations
	for _, win := range possibleWins {
		if newBoard[win[0]] != ' ' && newBoard[win[0]] == newBoard[win[1]] && newBoard[win[0]] == newBoard[win[2]] {
			game.Winner = string(newBoard[win[0]])
			return true
		}
	}

	// Check if the board is full
	for _, cell := range newBoard {
		if cell == ' ' {
			boardFull = false
			break
		}
	}

	// Update game status based on board state
	if boardFull {
		if game.Winner != "X" && game.Winner != "O" {
			game.Winner = "draw"
		}
		return true
	}

	return false
}
