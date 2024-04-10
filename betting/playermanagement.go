package betting

// PlayerManagement handles player-related actions in a betting round
type PlayerManagement interface {
	AddPlayer(playerID int) error
	RemovePlayer(playerID int) error
	PlayerFoldStatus(playerID int) bool
	SetCurrentPlayer(playerID int)
	GetCurrentPlayer() int
}
