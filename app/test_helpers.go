package app

import (
	"cosmossdk.io/log"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"
	"github.com/rs/zerolog"
)

// InitNewApp initiate a new App object
func InitNewApp() *App {
	// db := dbm.NewMemDB()
	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = DefaultNodeHome
	appOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue
	nodeHome := ""
	appOptions[flags.FlagHome] = nodeHome // ensure unique folder
	appOptions[server.FlagInvCheckPeriod] = 1
	// encCdc := MakeEncodingConfig()
	app, _ := New(
		log.NewCustomLogger(zerolog.Logger{}),
		simtestutil.DefaultStartUpConfig().DB,
		nil,
		true,
		nil,
		nil,
		simtestutil.DefaultStartUpConfig().BaseAppOption,
		simtestutil.DefaultStartUpConfig().BaseAppOption,
	)

	return app
}

// Initializes a new App without IBC functionality
func Setup(initChain bool) *App {
	app := InitNewApp()
	if initChain {
		app.InitChain(abci.ToRequestCommit().GetInitChain())
		app.DefaultGenesis()
		// commit genesis changes
		app.Commit()
		app.BeginBlocker(sdk.Context{})
	}

	return app
}
