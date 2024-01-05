package card

// Player represents a participant in a game.
type Player struct {
	ID   int
	Name string
}

// String Impelements the fmt.Stringer interface.
func (p Player) String() string {
	return p.Name
}

// Game is the abstraction for any game played with standard French playing cards.
type Game interface {
	// Name returns the name of the game.
	Name() string
	// Players returns the list of participants involved in the game.
	Players() []Player
}
