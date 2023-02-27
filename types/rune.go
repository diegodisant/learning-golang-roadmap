package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var currentChar rune = 'a'

	fmt.Printf("Size %d \n", unsafe.Sizeof(currentChar))

	fmt.Printf("Type: %s \n", reflect.TypeOf(currentChar))

	fmt.Printf("Unicode CodePoint: %U\n", currentChar)

	fmt.Printf("Character: %c\n", currentChar)

	newString := "0b£"

	fmt.Printf("%U\n", []rune(newString))
	fmt.Println([]rune(newString))
}
