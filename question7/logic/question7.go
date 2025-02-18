package logic

// Factorial returns the factorial of a number recursively
func Factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * Factorial(num-1)
}
