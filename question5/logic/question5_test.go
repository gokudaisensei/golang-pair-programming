package logic

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"two characters", "ab", "ba"},
		{"three characters", "abc", "cba"},
		{"four characters", "abcd", "dcba"},
		{"five characters", "abcde", "edcba"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ReverseString(tt.input)
			if actual != tt.expected {
				t.Errorf("ReverseString(%s): expected %s, actual %s", tt.input, tt.expected, actual)
			}
		})
	}
}
