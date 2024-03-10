package keeper

import (
	"context"
	"encoding/binary"
	"tic-chain/x/tttibc/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetGameCount get the total number of game
func (k Keeper) GetGameCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.GameCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetGameCount set the total number of game
func (k Keeper) SetGameCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.GameCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendGame appends a game in the store with a new id and update the count
func (k Keeper) AppendGame(
	ctx context.Context,
	game types.Game,
) uint64 {
	// Create the game
	count := k.GetGameCount(ctx)

	// Set the ID of the appended value
	game.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GameKey))
	appendedValue := k.cdc.MustMarshal(&game)
	store.Set(GetGameIDBytes(game.Id), appendedValue)

	// Update game count
	k.SetGameCount(ctx, count+1)

	return count
}

// SetGame set a specific game in the store
func (k Keeper) SetGame(ctx context.Context, game types.Game) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GameKey))
	b := k.cdc.MustMarshal(&game)
	store.Set(GetGameIDBytes(game.Id), b)
}

// GetGame returns a game from its id
func (k Keeper) GetGame(ctx context.Context, id uint64) (val types.Game, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GameKey))
	b := store.Get(GetGameIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGame removes a game from the store
func (k Keeper) RemoveGame(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GameKey))
	store.Delete(GetGameIDBytes(id))
}

// GetAllGame returns all game
func (k Keeper) GetAllGame(ctx context.Context) (list []types.Game) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GameKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Game
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetGameIDBytes returns the byte representation of the ID
func GetGameIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.GameKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
