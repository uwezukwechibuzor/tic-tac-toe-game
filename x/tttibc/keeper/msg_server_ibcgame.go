package keeper

import (
	"context"
	"tic-chain/x/tttibc/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
)

func (k msgServer) SendIbcgame(goCtx context.Context, msg *types.MsgSendIbcgame) (*types.MsgSendIbcgameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.IbcgamePacketData

	// Transmit the packet
	_, err := k.TransmitIbcgamePacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendIbcgameResponse{}, nil
}
