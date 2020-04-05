package card

import "testing"

func TestSuit(t *testing.T) {
	tests := []struct {
		name           string
		suit           Suit
		expectedString string
	}{
		{"Club", Club, "Club"},
		{"Diamond", Diamond, "Diamond"},
		{"Heart", Heart, "Heart"},
		{"Spade", Spade, "Spade"},
		{"Invalid", Suit(4), ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			str := tc.suit.String()

			if str != tc.expectedString {
				t.Errorf("String: expected %s got %s", tc.expectedString, str)
			}
		})
	}
}

func TestRank(t *testing.T) {
	tests := []struct {
		name           string
		rank           Rank
		expectedString string
	}{
		{"Ace", Ace, "Ace"},
		{"Two", Two, "Two"},
		{"Three", Three, "Three"},
		{"Four", Four, "Four"},
		{"Five", Five, "Five"},
		{"Six", Six, "Six"},
		{"Seven", Seven, "Seven"},
		{"Eight", Eight, "Eight"},
		{"Nine", Nine, "Nine"},
		{"Ten", Ten, "Ten"},
		{"Jack", Jack, "Jack"},
		{"Queen", Queen, "Queen"},
		{"King", King, "King"},
		{"Invalid", Rank(13), ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			str := tc.rank.String()

			if str != tc.expectedString {
				t.Errorf("String: expected %s got %s", tc.expectedString, str)
			}
		})
	}
}

func TestCard(t *testing.T) {
	tests := []struct {
		name           string
		suit           Suit
		rank           Rank
		expectedString string
	}{
		{"ClubAce", Club, Ace, "A♣"},
		{"ClubTwo", Club, Two, "2♣"},
		{"ClubThree", Club, Three, "3♣"},
		{"ClubFour", Club, Four, "4♣"},
		{"ClubFive", Club, Five, "5♣"},
		{"ClubSix", Club, Six, "6♣"},
		{"ClubSeven", Club, Seven, "7♣"},
		{"ClubEight", Club, Eight, "8♣"},
		{"ClubNine", Club, Nine, "9♣"},
		{"ClubTen", Club, Ten, "10♣"},
		{"ClubJack", Club, Jack, "J♣"},
		{"ClubQueen", Club, Queen, "Q♣"},
		{"ClubKing", Club, King, "K♣"},
		{"DiamondAce", Diamond, Ace, "A♦"},
		{"DiamondTwo", Diamond, Two, "2♦"},
		{"DiamondThree", Diamond, Three, "3♦"},
		{"DiamondFour", Diamond, Four, "4♦"},
		{"DiamondFive", Diamond, Five, "5♦"},
		{"DiamondSix", Diamond, Six, "6♦"},
		{"DiamondSeven", Diamond, Seven, "7♦"},
		{"DiamondEight", Diamond, Eight, "8♦"},
		{"DiamondNine", Diamond, Nine, "9♦"},
		{"DiamondTen", Diamond, Ten, "10♦"},
		{"DiamondJack", Diamond, Jack, "J♦"},
		{"DiamondQueen", Diamond, Queen, "Q♦"},
		{"DiamondKing", Diamond, King, "K♦"},
		{"HeartAce", Heart, Ace, "A♥"},
		{"HeartTwo", Heart, Two, "2♥"},
		{"HeartThree", Heart, Three, "3♥"},
		{"HeartFour", Heart, Four, "4♥"},
		{"HeartFive", Heart, Five, "5♥"},
		{"HeartSix", Heart, Six, "6♥"},
		{"HeartSeven", Heart, Seven, "7♥"},
		{"HeartEight", Heart, Eight, "8♥"},
		{"HeartNine", Heart, Nine, "9♥"},
		{"HeartTen", Heart, Ten, "10♥"},
		{"HeartJack", Heart, Jack, "J♥"},
		{"HeartQueen", Heart, Queen, "Q♥"},
		{"HeartKing", Heart, King, "K♥"},
		{"SpadeAce", Spade, Ace, "A♠"},
		{"SpadeTwo", Spade, Two, "2♠"},
		{"SpadeThree", Spade, Three, "3♠"},
		{"SpadeFour", Spade, Four, "4♠"},
		{"SpadeFive", Spade, Five, "5♠"},
		{"SpadeSix", Spade, Six, "6♠"},
		{"SpadeSeven", Spade, Seven, "7♠"},
		{"SpadeEight", Spade, Eight, "8♠"},
		{"SpadeNine", Spade, Nine, "9♠"},
		{"SpadeTen", Spade, Ten, "10♠"},
		{"SpadeJack", Spade, Jack, "J♠"},
		{"SpadeQueen", Spade, Queen, "Q♠"},
		{"SpadeKing", Spade, King, "K♠"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			card := New(tc.suit, tc.rank)

			if s := card.Suit(); s != tc.suit {
				t.Errorf("Suit: expected %s got %s", tc.suit, s)
			}

			if r := card.Rank(); r != tc.rank {
				t.Errorf("Rank: expected %s got %s", tc.rank, r)
			}

			if s := card.String(); s != tc.expectedString {
				t.Errorf("String: expected %s got %s", tc.expectedString, s)
			}
		})
	}
}
