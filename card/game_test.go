package card

import "testing"

func TestPlayer(t *testing.T) {
	tests := []struct {
		name           string
		playerID       int
		playerName     string
		expectedString string
	}{
		{"Alice", 1, "Alice", "Alice"},
		{"Bob", 2, "Bob", "Bob"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := Player{
				ID:   tc.playerID,
				Name: tc.playerName,
			}

			if str := p.String(); str != tc.expectedString {
				t.Errorf("String(): expected %s got %s", tc.expectedString, str)
			}
		})
	}
}
