package player

import "github.com/forrest321/poker/card"

// Player represents a participant in a poker game.
type Player struct {
	ID        int         // Unique identifier
	Hand      []card.Card // Current hand
	ChipCount int         // Current chip count
	InGame    bool        // Active in current game
	Folded    bool        // Whether the player has folded
	// Add other attributes as needed
}

// NewPlayer creates a new player with a specified ID and chip count.
func NewPlayer(id, chipCount int) *Player {
	return &Player{
		ID:        id,
		ChipCount: chipCount,
		InGame:    true,
		Folded:    false,
	}
}
