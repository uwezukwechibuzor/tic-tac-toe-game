package keeper

import (
	"fmt"
	"tic-chain/x/tictactoe/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// makePayment handles the payment required for joining the game.
func (k msgServer) makePayment(ctx sdk.Context, playerAddress sdk.AccAddress) error {
	// Joining a game requires a payment of 1 TTT (the native token of the protocol).
	// Get the payment as sdk.Coins
	amount := "1ttt"
	payment, err := sdk.ParseCoinsNormalized(amount)
	if err != nil {
		return fmt.Errorf("failed to parse payment: %v", err)
	}

	// Use the module account
	// When any player joins the Tic Tac Toe game, it requires a payment of 1 TTT (the native token of the protocol)
	// To facilitate this payment, we use the module account as an escrow account.
	// This approach ensures that the entry fee payment are collected securely and managed within the context
	// of the Tic Tac Toe game module, allowing for easier tracking and management of funds.
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, playerAddress, types.ModuleName, payment); err != nil {
		return fmt.Errorf("failed to send payment: %v", err)
	}

	return nil
}
