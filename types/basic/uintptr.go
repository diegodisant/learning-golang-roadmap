package main

import (
	"fmt"
	"unsafe"
)

type sample struct {
	a int
	b string
}

func main() {
	s := &sample{a: 1, b: "test"}

	// s has its own address
	fmt.Println(&s)

	// s starts from a address
	// and s.b you get at what location jump to get value from s.b indirectly
	p := unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Offsetof(s.b))

	fmt.Println(*(*string)(p))

	var location uintptr = 0xc82000c290

	fmt.Println("Value of location:", location)
	fmt.Printf("Type of location is %T", location)
}
