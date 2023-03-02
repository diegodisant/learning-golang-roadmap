package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var pointX float32 = 0.32121

	fmt.Printf("%d bytes \n", unsafe.Sizeof(pointX))
	fmt.Printf("pointX is %s\n", reflect.TypeOf(pointX))

	var pointY float64 = 0.28192819281

	fmt.Printf("%d bytes \n", unsafe.Sizeof(pointY))
	fmt.Printf("pointX is %s\n", reflect.TypeOf(pointY))

	pointZ := 0.281928191234

	fmt.Printf("%d bytes \n", unsafe.Sizeof(pointZ))
	fmt.Printf("pointZ is %s\n", reflect.TypeOf(pointZ))
}
