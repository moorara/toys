package card

import (
	"strings"
	"testing"
)

func TestDeck(t *testing.T) {
	deck := NewDeck()

	if err := deck.verify(); err != nil {
		t.Errorf("verify(): %s", err)
	}

	deck.Shuffle()

	if err := deck.verify(); err != nil {
		t.Errorf("verify(): %s", err)
	}

	str := deck.String()

	for s := 0; s < suitCount; s++ {
		for r := 1; r <= rankCount; r++ { // Card ranks start from 1 (Ace)
			c := New(Suit(s), Rank(r))
			if !strings.Contains(str, c.String()) {
				t.Errorf("String(): card %s not found", c.String())
			}
		}
	}
}
