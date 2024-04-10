package card

import "fmt"

// Card represents a playing card with a suit and value.
type Card struct {
	Suit  Suit
	Value Value
}

// NewCard creates a new card with the given suit and value.
func NewCard(suit Suit, value Value) *Card {
	return &Card{Suit: suit, Value: value}
}

// Display returns a string representation of the card.
func (c Card) Display() string {
	return fmt.Sprintf("%s of %s", c.Value.String(), c.Suit)
}
