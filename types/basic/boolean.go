package main

import "fmt"

func main() {
	var isAsserted bool

	fmt.Printf("isAsserted value is %t\n", isAsserted)

	andOperation := 1 < 2 && 1 > 3
	fmt.Printf("1 < 2 && 1 > 3 is %t\n", andOperation)

	orOperation := 1 < 2 || 1 > 3
	fmt.Printf("1 < 2 || 1 > 3 is %t\n", orOperation)

	notOperation := !(1 > 2)
	fmt.Printf("!(1 > 2) is %t\n", notOperation)
}
