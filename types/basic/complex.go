package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a float32 = 3
	var b float32 = 7

	c := complex(a, b)

	var d complex64 = 4 + 5i

	fmt.Printf("c size is from %d bytes\n", unsafe.Sizeof(c))
	fmt.Printf("d size is from %d bytes\n", unsafe.Sizeof(d))

	fmt.Printf("c type is %s\n", reflect.TypeOf(c))
	fmt.Printf("d type is %s\n", reflect.TypeOf(d))

	fmt.Println(c+d, c-d, c*d, c/d)
}
