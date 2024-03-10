package keeper

import (
	"context"
	"encoding/binary"
	"fmt"
	"tic-chain/x/tictactoe/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetLeaderboardCount get the total number of leaderboard
func (k Keeper) GetLeaderboardCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.LeaderboardCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetLeaderboardCount set the total number of leaderboard
func (k Keeper) SetLeaderboardCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.LeaderboardCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetLeaderboard set a specific leaderboard in the store
func (k Keeper) SetLeaderboard(ctx context.Context, leaderboard types.Leaderboard) error {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LeaderboardKey))

	// Retrieve existing leaderboard
	for _, newPlayer := range leaderboard.Players {
		existingLeaderboardBytes := store.Get([]byte(newPlayer.PlayerId))
		var existingPlayer types.Player
		if existingLeaderboardBytes != nil {
			if err := k.cdc.Unmarshal(existingLeaderboardBytes, &existingPlayer); err != nil {
				return fmt.Errorf("failed to unmarshal existing leaderboard player: %v", err)
			}
			existingPlayer.Points += newPlayer.Points
		} else {
			existingPlayer = *newPlayer
		}

		// Set the player in the store
		playerBytes, err := k.cdc.Marshal(&existingPlayer)
		if err != nil {
			return fmt.Errorf("failed to marshal player: %v", err)
		}
		store.Set([]byte(newPlayer.PlayerId), playerBytes)
	}

	return nil
}

// GetAllLeaderboard returns all leaderboard
func (k Keeper) GetAllLeaderboard(ctx context.Context) (types.Leaderboard, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LeaderboardKey))

	var leaderboard types.Leaderboard

	// Iterate over all items in the store and return each player's details
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var player types.Player
		if err := k.cdc.Unmarshal(iterator.Value(), &player); err != nil {
			return types.Leaderboard{}, fmt.Errorf("failed to unmarshal player: %v", err)
		}

		leaderboard.Players = append(leaderboard.Players, &player)
	}

	return leaderboard, nil
}

// ResetPlayerPoints resets a player's points to zero
func (k Keeper) ResetPlayerPoints(ctx context.Context, playerID string) error {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LeaderboardKey))

	// Retrieve the player from the store
	playerBytes := store.Get([]byte(playerID))
	if playerBytes == nil {
		return fmt.Errorf("player with ID %s not found", playerID)
	}

	var player types.Player
	if err := k.cdc.Unmarshal(playerBytes, &player); err != nil {
		return fmt.Errorf("failed to unmarshal player: %v", err)
	}

	// Reset player's points to zero
	player.Points = 0

	// Set the updated player in the store
	updatedPlayerBytes, err := k.cdc.Marshal(&player)
	if err != nil {
		return fmt.Errorf("failed to marshal updated player: %v", err)
	}
	store.Set([]byte(playerID), updatedPlayerBytes)

	return nil
}
