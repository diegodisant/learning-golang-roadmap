package main

import "fmt"

func main() {
  text := make([]string, 2, 3)

  fmt.Println(text)

  newText := []string{"a", "b", "c"}
  fmt.Println(newText)

  newText = append(newText, "d")
  fmt.Println(newText)

  for _, currentChar := range newText {
    fmt.Println(currentChar)
  }
}
