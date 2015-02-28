package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println("Got: ", msg)

	msgs := make(chan string, 2)

	msgs <- "buffered"
	msgs <- "channel"

	fmt.Println(<-msgs)
	fmt.Println(<-msgs)

}
