package tictactoe

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "tic-chain/api/ticchain/tictactoe"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "GameAll",
					Use:       "list-game",
					Short:     "List all game",
				},
				{
					RpcMethod:      "Game",
					Use:            "show-game [id]",
					Short:          "Shows a game by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "GameBoard",
					Use:            "show-game-board [id]",
					Short:          "Shows a game board by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "GameBoardAll",
					Use:       "list-game-boards",
					Short:     "List all game boards",
				},
				{
					RpcMethod: "LeaderboardAll",
					Use:       "list-leaderboard",
					Short:     "List all leaderboard",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod: "CreateGame",
					Use:       "create-game",
					Short:     "Send a create-game tx",
				},
				{
					RpcMethod:      "JoinGameX",
					Use:            "join-game-x [game-id]",
					Short:          "Send a join-game tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "gameId"}},
				},
				{
					RpcMethod:      "JoinGameO",
					Use:            "join-game-o [game-id]",
					Short:          "Send a join-game as opponent tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "gameId"}},
				},
				{
					RpcMethod:      "PlayMove",
					Use:            "play-move [game-id] [move]",
					Short:          "Send a play-move tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "gameId"}, {ProtoField: "move"}},
				},
				{
					RpcMethod:      "DeleteGame",
					Use:            "delete-game [game-id]",
					Short:          "Send a delete-game tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "gameId"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
