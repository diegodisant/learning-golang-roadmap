package main

import "fmt"

func main() {
	eventsChannel := make(chan string, 3)

	eventsChannel <- "a"
	eventsChannel <- "b"
	eventsChannel <- "c"

	close(eventsChannel)

	for eventValue := range eventsChannel {
		fmt.Println(eventValue)
	}
}
