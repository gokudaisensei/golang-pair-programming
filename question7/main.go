package main

import (
	"fmt"
	"question7/logic"
)

func main() {
	// Get user input
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	// Calculate the factorial and show the result
	result := logic.Factorial(num)
	fmt.Printf("The factorial of %d is %d\n", num, result)
}
