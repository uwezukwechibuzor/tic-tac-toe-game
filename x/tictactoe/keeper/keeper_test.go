package keeper_test

import (
	"testing"
	"tic-chain/x/tictactoe/keeper"
	"tic-chain/x/tictactoe/types"

	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	TicApp "tic-chain/app"
)

const (
	initChain = true
)

var (
	adminAddress        = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	firstPlayerAddress  = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	secondPlayerAddress = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
)

type KeeperTestSuite struct {
	suite.Suite

	address     []sdk.AccAddress
	legacyAmino *codec.LegacyAmino
	ctx         sdk.Context
	app         *TicApp.App

	tictactoeKeeper keeper.Keeper
	msgServer       types.MsgServer
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app = TicApp.Setup(initChain)

	suite.legacyAmino = suite.app.LegacyAmino()
	suite.ctx = suite.app.BaseApp.NewContext(initChain)

	suite.tictactoeKeeper = suite.app.TictactoeKeeper
	suite.msgServer = keeper.NewMsgServerImpl(suite.tictactoeKeeper)

	suite.address = []sdk.AccAddress{adminAddress, firstPlayerAddress, secondPlayerAddress}
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
