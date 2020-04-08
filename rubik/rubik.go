// Package rubik models the standard 3x3 Rubik's Cube.
package rubik

// Color represents one of the six colors in the standard 3x3 Rubik's Cube.
type Color int

// The standard six colors in the standard 3x3 Rubik's Cube.
const (
	White Color = iota
	Yellow
	Blue
	Green
	Orange
	Red
)

func (c Color) String() string {
	switch c {
	case White:
		return "White"
	case Yellow:
		return "Yellow"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Orange:
		return "Orange"
	case Red:
		return "Red"
	default:
		return ""
	}
}

// Rotation represents a face rotation.
type Rotation int

const (
	// C90 is a clockwise 90-degrees rotation.
	C90 Rotation = iota
	// CC90 is a counterclockwise 90-degrees rotation.
	CC90
	// C180 is a clockwise 180-degrees rotation.
	C180
)

// Cube represents a standard 3x3 Rubik's Cube.
type Cube interface {
	F(Rotation)
	B(Rotation)
	R(Rotation)
	L(Rotation)
	U(Rotation)
	D(Rotation)
	Solved() bool
}

// corner cubelet
//      +--------+
//     /        /|
//    /   3    / |
//   /        /  |
//  +--------+ 2 +
//  |        |  /
//  |   1    | /
//  |        |/
//  +--------+

// edge cubelet
//      +--------+
//     //////////|
//    ////////// |
//   //////////  |
//  +--------+ 2 +
//  |        |  /
//  |   1    | /
//  |        |/
//  +--------+

// center cubelet
//      +--------+
//     //////////|
//    ///////////|
//   ////////////|
//  +--------+///+
//  |        |///
//  |   1    |//
//  |        |/
//  +--------+

// cube is the implementation of a standard 3x3 Rubik's Cube.
type cube struct {
}

// NewCube creates a new standard 3x3 Rubik's Cube.
// The created cube is in solved state and the white face is in the front.
func NewCube() Cube {
	return &cube{}
}

// F rotates the front face.
func (c *cube) F(Rotation) {

}

// B rotates the back face.
func (c *cube) B(Rotation) {

}

// R rotates the right face.
func (c *cube) R(Rotation) {

}

// L rotates the left face.
func (c *cube) L(Rotation) {

}

// U rotates the up face.
func (c *cube) U(Rotation) {

}

// D rotates the down face.
func (c *cube) D(Rotation) {

}

// Solved determines if the cube is solved.
func (c *cube) Solved() bool {
	return false
}
