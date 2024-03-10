package keeper

import (
	"tic-chain/x/tictactoe/types"
)

var _ types.QueryServer = Keeper{}
