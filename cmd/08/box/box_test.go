package box_test

import (
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/08/box"
)

func TestFrom(t *testing.T) {
	tests := []struct {
		input    string
		expected box.Box
	}{
		{"1,2,3", box.From("1,2,3")},
		{"0,0,0", box.From("0,0,0")},
		{"-1,-2,-3", box.From("-1,-2,-3")},
	}

	for _, test := range tests {
		result := box.From(test.input)

		if result != test.expected {
			t.Errorf("From(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestDistanceFrom(t *testing.T) {
	tests := []struct {
		box1     box.Box
		box2     box.Box
		expected float64
	}{
		{box.From("0,0,0"), box.From("1,1,1"), 1.7320508075688772},
		{box.From("1,2,3"), box.From("4,5,6"), 5.196152422706632},
		{box.From("-1,-2,-3"), box.From("1,2,3"), 7.483314773547883},
	}

	for _, test := range tests {
		result := test.box1.DistanceFrom(test.box2)

		if result != test.expected {
			t.Errorf("DistanceFrom(%v, %v) = %v; want %v", test.box1, test.box2, result, test.expected)
		}
	}
}
