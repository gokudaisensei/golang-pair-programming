package logic

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		num      int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
		{8, 40320},
		{9, 362880},
		{10, 3628800},
	}

	for _, test := range tests {
		if result := Factorial(test.num); result != test.expected {
			t.Errorf("SolveProblem(%d) = %d, want %d", test.num, result, test.expected)
		}
	}
}
