// Package card models the standard 52-card deck of French playing cards.
package card

const (
	suitCount  = 4
	rankCount  = 13
	totalCount = 52
)

// Suit represents one of the four suit symbols in a deck of playing cards.
type Suit int

const (
	// Club suit ♣
	Club Suit = iota
	// Diamond suit ♦
	Diamond
	// Heart suit ♥
	Heart
	// Spade suit ♠
	Spade
)

// String Impelements fmt.Stringer
func (s Suit) String() string {
	switch s {
	case Club:
		return "Club"
	case Diamond:
		return "Diamond"
	case Heart:
		return "Heart"
	case Spade:
		return "Spade"
	default:
		return ""
	}
}

// Rank represents one of the thirteen ranks in a deck of playing cards.
type Rank int

// The thirteen ranks in each of the four suits.
const (
	Ace Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// String Impelements fmt.Stringer
func (r Rank) String() string {
	switch r {
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return ""
	}
}

// Card represents a playing card in a standard 52-card deck.
type Card struct {
	suit Suit
	rank Rank
}

// New creates a new playing card given a suit and a rank.
func New(s Suit, r Rank) Card {
	return Card{
		suit: s,
		rank: r,
	}
}

// Suit returns the suit of a playing card.
func (c Card) Suit() Suit {
	return c.suit
}

// Rank returns the rank of a playing card.
func (c Card) Rank() Rank {
	return c.rank
}

// String Impelements fmt.Stringer
func (c Card) String() string {
	suit := ""
	switch c.suit {
	case Club:
		suit = "♣"
	case Diamond:
		suit = "♦"
	case Heart:
		suit = "♥"
	case Spade:
		suit = "♠"
	}

	rank := ""
	switch c.rank {
	case Ace:
		rank = "A"
	case Two:
		rank = "2"
	case Three:
		rank = "3"
	case Four:
		rank = "4"
	case Five:
		rank = "5"
	case Six:
		rank = "6"
	case Seven:
		rank = "7"
	case Eight:
		rank = "8"
	case Nine:
		rank = "9"
	case Ten:
		rank = "10"
	case Jack:
		rank = "J"
	case Queen:
		rank = "Q"
	case King:
		rank = "K"
	}

	return rank + suit
}
