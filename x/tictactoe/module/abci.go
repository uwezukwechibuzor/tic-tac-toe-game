package tictactoe

import (
	"context"
	"tic-chain/x/tictactoe/keeper"
)

// EndBlocker processes premium payment at every block.
func EndBlocker(ctx context.Context, k keeper.Keeper) {
	// MintAndSendCoins to mint and send coins to players with => 15 points
	if err := k.MintAndSendCoins(ctx); err != nil {
		return
	}
}
