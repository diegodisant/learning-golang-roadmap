package main

import "fmt"

func main() {
	// multiple variable declaration
	var isProductUpdated, hasTransitableStatus bool

	// single variable declaration
	var updateCounter int

	fmt.Println(isProductUpdated, hasTransitableStatus, updateCounter)

	// single variable declaration with default value
	var stockSum int = 0

	fmt.Println(stockSum)

	// multiple variable declaration with default values
	var (
		hasRejectedStatus = false
		hasAcceptedStatus = true
	)

	fmt.Println(hasRejectedStatus, hasAcceptedStatus)
}
