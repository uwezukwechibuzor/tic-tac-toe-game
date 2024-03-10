package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgJoinGameX{}

func NewMsgJoinGame(playerAddress string, gameId uint64) *MsgJoinGameX {
	return &MsgJoinGameX{
		PlayerAddress: playerAddress,
		GameId:        gameId,
	}
}

func (msg *MsgJoinGameX) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.PlayerAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid playerAddress string (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgJoinGameO{}

func NewMsgJoinGameO(playerAddress string, gameId uint64) *MsgJoinGameO {
	return &MsgJoinGameO{
		PlayerAddress: playerAddress,
		GameId:        gameId,
	}
}

func (msg *MsgJoinGameO) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.PlayerAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid playerAddress string (%s)", err)
	}
	return nil
}
