package main

import "fmt"

func main() {
	sample := [3]string{"a", "b", "c"}
	print(sample)
}

func print(sample [3]string) {
	fmt.Println(sample)
}
