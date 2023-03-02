package main

import "fmt"

func main() {
	var currentValue int = 3

	var currentNumber *int

	currentNumber = &currentValue

	fmt.Println(currentNumber)
	fmt.Println(*currentNumber)

	*currentNumber = 5

	fmt.Println(*currentNumber)
	fmt.Println(currentValue)
}
