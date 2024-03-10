package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteGame{}

func NewMsgDeleteGame(adminAddress string, gameId uint64) *MsgDeleteGame {
	return &MsgDeleteGame{
		AdminAddress: adminAddress,
		GameId:       gameId,
	}
}

func (msg *MsgDeleteGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
