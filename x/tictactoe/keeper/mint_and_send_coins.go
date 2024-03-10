package keeper

import (
	"context"
	"fmt"
	"tic-chain/x/tictactoe/types"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MintAndSendCoins mints and sends coins to players with => 15 points in the game leaderboard
// A player which manages to score 15 points based on this leaderboard gets 5 newly minted WIN tokens
// and their points tally is set back to 0.
func (k Keeper) MintAndSendCoins(ctx context.Context) error {
	// Get all players from the leaderboard
	leaderboard, err := k.GetAllLeaderboard(ctx)
	if err != nil {
		return fmt.Errorf("failed to get leaderboard: %v", err)
	}

	// Iterate through all players
	for _, player := range leaderboard.Players {
		// Check if the player has 15 or more points
		if player.Points >= 15 {
			// Convert player's address to sdk.AccAddress
			playerAddress, err := sdk.AccAddressFromBech32(player.PlayerId)
			if err != nil {
				return fmt.Errorf("failed to convert player address: %v", err)
			}

			// Send the minted coins to the player's address
			moduleAcct := k.accountKeeper.GetModuleAddress(types.ModuleName)

			var mintCoins sdk.Coins

			mintCoins = mintCoins.Add(sdk.NewCoin("WIN", math.NewInt(5)))
			if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
				return nil
			}

			if err := k.bankKeeper.SendCoins(ctx, moduleAcct, playerAddress, mintCoins); err != nil {
				return nil
			}
			// Reset the player's points tally to 0
			if err := k.ResetPlayerPoints(ctx, player.PlayerId); err != nil {
				return fmt.Errorf("failed to reset player's point: %v", err)
			}
		}
	}

	return nil
}
