package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var letter byte = 'z'

	fmt.Printf("Size: %d\n", unsafe.Sizeof(letter))
	fmt.Printf("Type: %s\n", reflect.TypeOf(letter))
	fmt.Printf("Character: %c\n", letter)

	var commonString string = "abc"

	fmt.Println([]byte(commonString))
}
