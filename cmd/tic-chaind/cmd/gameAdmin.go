package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	ticTacToeTypes "tic-chain/x/tictactoe/types"
)

// AddGenesisGameAdminCmd returns add-genesis-game-admin cobra Command.
func AddGenesisGameAdminCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-genesis-game-admin [address]",
		Short: "Add a genesis game admin to genesis.json",
		Long:  `Add a genesis game admin to genesis.json. The provided game admin must specify the account address. `,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := client.GetClientContextFromCmd(cmd)
			depCdc := ctx.Codec
			cdc := depCdc

			config := server.GetServerContextFromCmd(cmd).Config
			config.SetRoot(ctx.HomeDir)

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return fmt.Errorf("failed to parse address")
			}

			genFile := config.GenesisFile()
			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}
			GameGenState := ticTacToeTypes.GetGenesisStateFromAppState(appState)

			GameGenState.AdminAddress = addr.String()
			GameGenStateBz := cdc.MustMarshalJSON(&GameGenState)

			appState[ticTacToeTypes.ModuleName] = GameGenStateBz
			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}
			genDoc.AppState = appStateJSON
			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	return cmd
}
