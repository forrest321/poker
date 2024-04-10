package deck

import (
	"github.com/forrest321/poker/card"
)

// Deck represents a collection of playing cards.
type Deck struct {
	Cards []card.Card
}

// NewDeck creates and returns a new deck of 52 playing cards.
func NewDeck() *Deck {
	var cards []card.Card

	// Iterate over each suit and value to create all combinations
	for _, suit := range []card.Suit{card.Spades, card.Hearts, card.Diamonds, card.Clubs} {
		for _, value := range []card.Value{card.Two, card.Three, card.Four, card.Five, card.Six, card.Seven, card.Eight, card.Nine, card.Ten, card.Jack, card.Queen, card.King, card.Ace} {
			cards = append(cards, *card.NewCard(suit, value))
		}
	}

	return &Deck{Cards: cards}
}
