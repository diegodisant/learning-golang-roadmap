package main

import "fmt"

func main() {
  firstString := "this\nthat"

  fmt.Printf("firstString is: %s\n", firstString);

  secondString := `this\nthat`

  fmt.Printf("secondString is: %s\n", secondString)

  newString := "ab£"

  fmt.Println([]byte(newString))

  fmt.Println(len(newString))

  for _, currentChar := range newString {
    fmt.Println(string(currentChar))
  }
}
