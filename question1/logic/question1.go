package logic

// FindLargest returns the largest of three numbers.
func FindLargest(a, b, c int) int {
	switch {
	case a > b && a > c:
		return a
	case b > a && b > c:
		return b
	default:
		return c
	}
}
