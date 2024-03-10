package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "tic-tac-toe-cli",
		Short: "A CLI tool to query tic-tac-toe endpoints",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Please indicate a subcommand.")
		},
	}

	serverAddress string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&serverAddress, "server", "s", "http://localhost:1317", "Server address")

	rootCmd.AddCommand(newGameBoardsCmd())
	rootCmd.AddCommand(newGameBoardCmd())
}

func Execute() error {
	return rootCmd.Execute()
}

func newGameBoardsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "game-boards",
		Short: "Get all game boards",
		RunE: func(cmd *cobra.Command, args []string) error {
			endpoint := fmt.Sprintf("%s/tic-chain/tictactoe/game-boards", serverAddress)
			return getAndPrintBoard(endpoint)
		},
	}
}

func newGameBoardCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "game-board [gameID]",
		Short: "Get a specific game board using game ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			endpoint := fmt.Sprintf("%s/tic-chain/tictactoe/game-board/%s", serverAddress, args[0])
			return getAndPrintBoard(endpoint)
		},
	}
}

func getAndPrintBoard(endpoint string) error {
	resp, err := http.Get(endpoint)
	if err != nil {
		return fmt.Errorf("failed to fetch data from %s: %v", endpoint, err)
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return fmt.Errorf("failed to decode JSON response from %s: %v", endpoint, err)
	}

	fmt.Println("Board:")
	if board, ok := data["board"].(string); ok {
		printBoard(board)
	} else if boards, ok := data["board"].([]interface{}); ok {
		for _, b := range boards {
			if board, ok := b.(string); ok {
				printBoard(board)
			}
		}
	} else {
		return fmt.Errorf("invalid board data format")
	}
	return nil
}

func printBoard(board string) {
	rows, cols := 3, 3

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			idx := i*cols + j
			if board[idx] == ' ' {
				fmt.Print("  ")
			} else {
				fmt.Printf("%c ", board[idx])
			}
		}
		fmt.Println()
	}
}
