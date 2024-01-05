package card

import (
	"fmt"
	"math/rand"
	"strings"
)

// Deck represents a standard 52-card deck of French playing cards.
type Deck struct {
	cards [cardCount]Card
}

// NewDeck creates a new deck of playing card.
// The cards in the deck are shuffled.
func NewDeck() *Deck {
	deck := new(Deck)
	for s := 0; s < suitCount; s++ {
		for r := 0; r < rankCount; r++ {
			deck.cards[s*rankCount+r].suit = Suit(s)
			deck.cards[s*rankCount+r].rank = Rank(r + 1) // Card ranks start from 1 (Ace)
		}
	}

	return deck
}

// verify verfies a deck of card is valid.
func (d *Deck) verify() error {
	if l := len(d.cards); l != cardCount {
		return fmt.Errorf("invalid number of cards: %d", l)
	}

	for s := 0; s < suitCount; s++ {
		for r := 1; r <= rankCount; r++ { // Card ranks start from 1 (Ace)
			c := New(Suit(s), Rank(r))
			found := false

			// Search through the deck
			for i := 0; i < len(d.cards); i++ {
				if d.cards[i].Equals(c) {
					found = true
					break
				}
			}

			if !found {
				return fmt.Errorf("card %s not found", c)
			}
		}
	}

	return nil
}

// Shuffle shuffles the cards in the Deck in O(n) time.
func (d *Deck) Shuffle() {
	// Fisher-Yates shuffle algorithm
	for i := len(d.cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// String Impelements the fmt.Stringer interface.
func (d *Deck) String() string {
	lines := make([]string, suitCount)
	for s := 0; s < suitCount; s++ {
		line := make([]string, rankCount)
		for r := 0; r < rankCount; r++ {
			line[r] = d.cards[s*rankCount+r].String()
		}
		lines[s] = strings.Join(line, " ")
	}

	return strings.Join(lines, "\n")
}
