package main

import "fmt"

func main() {
	eventsChannel := make(chan string)

	go sendEvents(eventsChannel)

	for eventValue := range eventsChannel {
		fmt.Println(eventValue)
	}
}

func sendEvents(eventsChannel chan<- string) {
	eventsChannel <- "a"
	eventsChannel <- "b"
	eventsChannel <- "c"
	eventsChannel <- "d"

	close(eventsChannel)
}
