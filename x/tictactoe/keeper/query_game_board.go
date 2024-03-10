package keeper

import (
	"context"
	"tic-chain/x/tictactoe/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GameBoard returns the game board for players in a given game id.
func (k Keeper) GameBoard(ctx context.Context, req *types.QueryGetGameBoardRequest) (*types.QueryGetGameBoardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	game, found := k.GetGame(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	// Return the game board with a given game id
	return &types.QueryGetGameBoardResponse{Board: game.Board}, nil
}

// GameBoardAll returns all existing game board for all players
func (k Keeper) GameBoardAll(ctx context.Context, req *types.QueryAllGameBoardRequest) (*types.QueryAllGameBoardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var boards []string

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	gameStore := prefix.NewStore(store, types.KeyPrefix(types.GameKey))

	pageRes, err := query.Paginate(gameStore, req.Pagination, func(key []byte, value []byte) error {
		var game types.Game
		if err := k.cdc.Unmarshal(value, &game); err != nil {
			return err
		}

		// append game.Boards to boards
		boards = append(boards, game.Board)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGameBoardResponse{Board: boards, Pagination: pageRes}, nil
}
