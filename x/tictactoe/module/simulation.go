package tictactoe

import (
	"math/rand"
	"tic-chain/testutil/sample"
	"tic-chain/x/tictactoe/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	tictactoesimulation "tic-chain/x/tictactoe/simulation"
)

// avoid unused import issue
var (
	_ = tictactoesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateGame = "op_weight_msg_create_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateGame int = 100

	opWeightMsgJoinGame = "op_weight_msg_join_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgJoinGame int = 100

	opWeightMsgPlayMove = "op_weight_msg_play_move"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPlayMove int = 100

	opWeightMsgDeleteGame = "op_weight_msg_delete_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteGame int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	tictactoeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&tictactoeGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateGame int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateGame, &weightMsgCreateGame, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGame = defaultWeightMsgCreateGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateGame,
		tictactoesimulation.SimulateMsgCreateGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgJoinGame int
	simState.AppParams.GetOrGenerate(opWeightMsgJoinGame, &weightMsgJoinGame, nil,
		func(_ *rand.Rand) {
			weightMsgJoinGame = defaultWeightMsgJoinGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgJoinGame,
		tictactoesimulation.SimulateMsgJoinGameX(am.accountKeeper, am.bankKeeper, am.keeper),
	))
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgJoinGame,
		tictactoesimulation.SimulateMsgJoinGameO(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPlayMove int
	simState.AppParams.GetOrGenerate(opWeightMsgPlayMove, &weightMsgPlayMove, nil,
		func(_ *rand.Rand) {
			weightMsgPlayMove = defaultWeightMsgPlayMove
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPlayMove,
		tictactoesimulation.SimulateMsgPlayMove(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteGame int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteGame, &weightMsgDeleteGame, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteGame = defaultWeightMsgDeleteGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteGame,
		tictactoesimulation.SimulateMsgDeleteGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateGame,
			defaultWeightMsgCreateGame,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tictactoesimulation.SimulateMsgCreateGame(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgJoinGame,
			defaultWeightMsgJoinGame,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tictactoesimulation.SimulateMsgJoinGameX(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgJoinGame,
			defaultWeightMsgJoinGame,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tictactoesimulation.SimulateMsgJoinGameO(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgPlayMove,
			defaultWeightMsgPlayMove,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tictactoesimulation.SimulateMsgPlayMove(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteGame,
			defaultWeightMsgDeleteGame,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tictactoesimulation.SimulateMsgDeleteGame(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
