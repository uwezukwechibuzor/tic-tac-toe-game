package keeper

import (
	"tic-chain/x/tttibc/types"
)

var _ types.QueryServer = Keeper{}
