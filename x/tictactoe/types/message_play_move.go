package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgPlayMove{}

func NewMsgPlayMove(playerAddress string, gameId uint64, move string) *MsgPlayMove {
	return &MsgPlayMove{
		PlayerAddress: playerAddress,
		GameId:        gameId,
		Move:          move,
	}
}

func (msg *MsgPlayMove) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.PlayerAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
