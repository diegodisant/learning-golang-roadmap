package main

import "fmt"

func main() {
  add := func(x, y int) int {
    return x + y
  }

  fmt.Println(add(1, 2))

  fmt.Println(doOperation(add, 1, 2))
}

func doOperation(operation func(x, y int) int, x, y int) int {
  return operation(x, y)
}
