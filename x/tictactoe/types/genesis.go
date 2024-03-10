package types

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		GameList:        []Game{},
		LeaderboardList: []Leaderboard{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in game
	gameIdMap := make(map[uint64]bool)
	gameCount := gs.GetGameCount()
	for _, elem := range gs.GameList {
		if _, ok := gameIdMap[elem.GameId]; ok {
			return fmt.Errorf("duplicated id for game")
		}
		if elem.GameId >= gameCount {
			return fmt.Errorf("game id should be lower or equal than the last id")
		}
		gameIdMap[elem.GameId] = true
	}

	// Check for duplicated IDs in leaderboard
	leaderboardIdMap := make(map[string]bool)
	leaderboardCount := gs.GetLeaderboardCount()
	for _, leaderboard := range gs.LeaderboardList {
		for _, player := range leaderboard.Players {
			if _, ok := leaderboardIdMap[player.PlayerId]; ok {
				return fmt.Errorf("duplicate player ID in leaderboard")
			}
			playerID, err := strconv.ParseUint(player.PlayerId, 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse player ID: %v", err)
			}
			if playerID >= leaderboardCount {
				return fmt.Errorf("player ID exceeds leaderboard count")
			}
			leaderboardIdMap[player.PlayerId] = true
		}
	}

	return gs.Params.Validate()
}

// GetGenesisStateFromAppState returns GenesisState given raw application genesis state
func GetGenesisStateFromAppState(appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		if err := json.Unmarshal(appState[ModuleName], &genesisState); err != nil {
			// Handle the error during unmarshalling
			return GenesisState{}
		}
	}
	return genesisState
}
