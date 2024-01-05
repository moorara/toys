package rubik

import "testing"

func TestColor(t *testing.T) {
	tests := []struct {
		name           string
		color          Color
		expectedString string
	}{
		{"White", White, "White"},
		{"Yellow", Yellow, "Yellow"},
		{"Blue", Blue, "Blue"},
		{"Green", Green, "Green"},
		{"Orange", Orange, "Orange"},
		{"Red", Red, "Red"},
		{"Invalid", Color(6), "Color(6)"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if s := tc.color.String(); s != tc.expectedString {
				t.Errorf("String: expected %s got %s", tc.expectedString, s)
			}
		})
	}
}

func TestNewCube(t *testing.T) {

}

func TestCubeF(t *testing.T) {

}

func TestCubeB(t *testing.T) {

}

func TestCubeR(t *testing.T) {

}

func TestCubeL(t *testing.T) {

}

func TestCubeU(t *testing.T) {

}

func TestCubeD(t *testing.T) {

}

func TestSolved(t *testing.T) {

}
