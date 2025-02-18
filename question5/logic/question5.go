package logic

import "slices"

// ReverseString returns the reverse of the input string
func ReverseString(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}
