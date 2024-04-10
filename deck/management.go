package deck

import (
	"github.com/forrest321/poker/card"
	"math/rand"
	"time"
)

// Shuffle randomizes the order of cards in the deck.
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

// DealOne returns one card from the deck.
func (d *Deck) DealOne() card.Card {
	var dealtCard card.Card
	dealtCard, d.Cards = d.Cards[0], d.Cards[1:]
	return dealtCard
}
