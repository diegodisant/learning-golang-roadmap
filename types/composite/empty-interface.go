package main

import "fmt"

func main() {
	test("stijiogjaoi")
	test(2.2010291)
	test(true)
}

func test(anyType interface{}) {
	fmt.Printf("(%v, %T)\n", anyType, anyType)
}
