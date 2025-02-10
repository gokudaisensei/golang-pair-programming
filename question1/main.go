// question-1/main.go
package main

import (
	"fmt"
	"question1/logic"
)

func main() {
	var a, b, c int
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)
	fmt.Print("Enter third number: ")
	fmt.Scan(&c)

	largest := logic.FindLargest(a, b, c)
	fmt.Printf("The largest number is: %d\n", largest)
}
