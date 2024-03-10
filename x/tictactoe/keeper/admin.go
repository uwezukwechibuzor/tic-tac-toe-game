package keeper

import (
	"context"
	"tic-chain/x/tictactoe/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAdmin sets the Game admin account address
func (k Keeper) SetAdmin(ctx context.Context, admin sdk.AccAddress) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.GetGameAdminKey())

	// Convert admin address to bytes
	adminBytes := admin.Bytes()

	// Set the admin address in the store
	store.Set(types.GetGameAdminKey(), adminBytes)
}

// GetAdmin gets the Game admin account address.
func (k Keeper) GetAdmin(ctx context.Context) sdk.AccAddress {
	// Create store using storeAdapter
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.GetGameAdminKey())

	// fetch the admin address bytes from the store
	adminBytes := store.Get(types.GetGameAdminKey())

	// Convert bytes to sdk.AccAddress
	var admin sdk.AccAddress = adminBytes
	admin = adminBytes

	return admin
}
