package main

import (
	"fmt"
	"question5/logic"
)

func main() {
	// Get the user input string
	fmt.Print("Enter the string to reverse: ")
	var input string
	fmt.Scanln(&input)
	// Determine if it is a palindrome
	if input == logic.ReverseString(input) {
		fmt.Printf("%s is a palindrome.\n", input)
	} else {
		fmt.Printf("%s is not a palindrome.\n", input)
	}
}
