package card

import "math/rand"

// Deck represents a standard 52-card deck of French playing cards.
type Deck struct {
	cards [totalCount]Card
}

// NewDeck creates a new deck of playing card.
// The cards in the deck are shuffled.
func NewDeck() *Deck {
	deck := &Deck{}

	for s := 0; s < suitCount; s++ {
		for r := 0; r < rankCount; r++ {
			deck.cards[s*rankCount+r].suit = Suit(s)
			deck.cards[s*rankCount+r].rank = Rank(r)
		}
	}

	deck.Shuffle()

	return deck
}

// Shuffle shuffles the cards in the Deck in O(n) time.
func (d *Deck) Shuffle() {
	n := len(d.cards)
	for i := 0; i < n; i++ {
		j := i + rand.Intn(n-i)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// String Impelements fmt.Stringer
func (d *Deck) String() string {
	str := ""
	for s := 0; s < suitCount; s++ {
		for r := 0; r < rankCount; r++ {
			str += d.cards[s*rankCount+r].String() + " "
		}
		str += "\n"
	}

	return str
}
