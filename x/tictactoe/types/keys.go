package types

const (
	// ModuleName defines the module name
	ModuleName = "tictactoe"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tictactoe"
)

var (
	ParamsKey    = []byte("p_tictactoe")
	GameAdminKey = []byte{0x00}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	GameKey      = "Game/value/"
	GameCountKey = "Game/count/"
)

// GetGameAdminKey gets the key for the Game admin.
func GetGameAdminKey() []byte {
	return GameAdminKey
}

const (
	LeaderboardKey      = "Leaderboard/value/"
	LeaderboardCountKey = "Leaderboard/count/"
)
