package math_test

import (
	"testing"

	"08-tests/math"
)

func TestAdd(t *testing.T) {
	got := math.Add(1, 1)
	expected := 2

	if got != expected {
		t.Fail()
	}
}

func TestAddTableDriven(t *testing.T) {
	testCases := map[string]struct {
		a      int
		b      int
		result int
	}{
		"sum equal": {
			a:      10,
			b:      10,
			result: 20,
		},
		"sum with zero": {
			a:      0,
			b:      15,
			result: 15,
		},
	}

	for _, tc := range testCases {
		got := math.Add(tc.a, tc.b)

		if got != tc.result {
			t.Errorf("Expected %d but got %d", tc.result, got)
		}
	}
}
