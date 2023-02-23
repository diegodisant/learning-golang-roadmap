package main

import (
	"fmt"
	"math/bits"
	"reflect"
	"unsafe"
)

func main() {
	var myNumber int = 3891

	fmt.Printf("int platform bit size is: %d\n", bits.UintSize)
	printIntegerProperties("myNumber", myNumber)

	var i8 int8 = 127
	printIntegerProperties("i8", i8)

	var i16 int16 = 32000
	printIntegerProperties("i16", i16)

	var i32 int32 = 1000000
	printIntegerProperties("i32", i32)

	var i64 int64 = 271928101
	printIntegerProperties("i64", i64)

	var u8 uint8 = 255
	printIntegerProperties("u8", u8)

	var u16 uint16 = 65535
	printIntegerProperties("u16", u16)

	var u32 uint32 = 29102911
	printIntegerProperties("u32", u32)

	var u64 uint64 = 27192719
	printIntegerProperties("u64", u64)
}

func printIntegerProperties(name string, number interface{}) {
	fmt.Println()
	fmt.Printf("%s type is %s\n", name, reflect.TypeOf(number))
	fmt.Printf("%s has a size of %d bytes\n", name, unsafe.Sizeof(number))
}
