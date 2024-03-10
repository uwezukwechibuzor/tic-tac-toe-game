package keeper

import (
	"context"
	"fmt"
	"tic-chain/x/tictactoe/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
func (k Keeper) LeaderboardAll(ctx context.Context, req *types.QueryAllLeaderboardRequest) (*types.QueryAllLeaderboardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var leaderboards []types.Leaderboard

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	leaderboardStore := prefix.NewStore(store, types.KeyPrefix(types.LeaderboardKey))

	pageRes, err := query.Paginate(leaderboardStore, req.Pagination, func(key []byte, value []byte) error {
		var leaderboard types.Leaderboard
		if err := k.cdc.Unmarshal(value, &leaderboard); err != nil {
			return err
		}

		leaderboards = append(leaderboards, leaderboard)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLeaderboardResponse{Leaderboard: leaderboards, Pagination: pageRes}, nil
}*/

func (k Keeper) LeaderboardAll(ctx context.Context, req *types.QueryAllLeaderboardRequest) (*types.QueryAllLeaderboardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	leaderboard, err := k.GetAllLeaderboard(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get leaderboard: %v", err)
	}

	return &types.QueryAllLeaderboardResponse{Leaderboard: leaderboard}, nil
}
