package card

import "testing"

func TestDeck(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"OK"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			deck := NewDeck()

			if l := len(deck.cards); l != totalCount {
				t.Errorf("Deck Length: expected %d got %d", totalCount, l)
			}

			if s := deck.String(); s == "" {
				t.Error("String: got empty string")
			}
		})
	}
}
