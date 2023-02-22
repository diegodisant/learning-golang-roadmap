package main

import "fmt"

func main() {
	// single line declaration
	const defaultDivisor = 2.5

	fmt.Println(defaultDivisor)

	// multiple constant declaration
	const (
		defaultValue   = 5.5
		defaultInitial = 2.1
	)

	fmt.Println(defaultValue, defaultInitial)
}
