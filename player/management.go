package player

import "github.com/forrest321/poker/card"

// AddCard adds a card to the player's hand.
func (p *Player) AddCard(c card.Card) {
	p.Hand = append(p.Hand, c)
}

// ResetHand clears the player's hand.
func (p *Player) ResetHand() {
	p.Hand = nil
}

// Bet decreases the player's chip count by the bet amount.
func (p *Player) Bet(amount int) {
	p.ChipCount -= amount // Ensure amount is not greater than ChipCount
}

// Fold sets the player's status to folded.
func (p *Player) Fold() {
	p.Folded = true
}

// UpdateChipCount changes the player's chip count.
func (p *Player) UpdateChipCount(amount int) {
	p.ChipCount += amount // Use negative amount to subtract chips
}

// IsActive returns whether the player is active in the game.
func (p *Player) IsActive() bool {
	return p.InGame && !p.Folded
}

// ... Any other management methods needed
