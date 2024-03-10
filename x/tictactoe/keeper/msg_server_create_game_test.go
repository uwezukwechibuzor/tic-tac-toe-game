package keeper_test

import (
	//"strings"

	"tic-chain/x/tictactoe/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestCreateGame() {
	type args struct {
		msgCreateGame types.MsgCreateGame
	}

	type errArgs struct {
		shouldPass bool
		contains   string
	}

	tests := []struct {
		name    string
		args    args
		errArgs errArgs
	}{
		{
			name: "Valid Game Created",
			args: args{
				msgCreateGame: *types.NewMsgCreateGame(suite.address[0].String(), 1),
			},
			errArgs: errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			name: "Invalid Admin Address",
			args: args{
				msgCreateGame: *types.NewMsgCreateGame(suite.address[1].String(), 1),
			},
			errArgs: errArgs{
				shouldPass: false,
				contains:   "invalid player address",
			},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			// set admin address
			suite.tictactoeKeeper.SetAdmin(suite.ctx, sdk.AccAddress(tc.args.msgCreateGame.Creator))

			// check that the game creator is the admin
			adminAddress := suite.tictactoeKeeper.GetAdmin(suite.ctx)

			suite.msgServer.CreateGame(sdk.WrapSDKContext(suite.ctx), &tc.args.msgCreateGame)

			// Get Games
			getGame, found := suite.tictactoeKeeper.GetGame(suite.ctx, tc.args.msgCreateGame.GameId)

			if tc.errArgs.shouldPass {
				suite.Require().True(found)
				suite.Require().NotEmpty(getGame)
				suite.Require().NotEmpty(adminAddress)
			} else {
				suite.Require().Empty(getGame)
				suite.Require().False(found)
			}
		})
	}
}
