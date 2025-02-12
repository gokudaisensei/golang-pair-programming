package main

import (
	"fmt"
	"question3/logic"
)

func main() {
	// Get the year from the user
	var year int
	fmt.Print("Enter a year: ")
	_, err := fmt.Scanln(&year)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Run the IsLeapYearFunction from the logic package
	if logic.IsLeapYear(year) {
		fmt.Println(year, "is a leap year.")
	} else {
		fmt.Println(year, "is not a leap year.")
	}

}
