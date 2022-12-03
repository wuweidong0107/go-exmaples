package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "Ping message from other goroutine"
	}()

	msg := <-messages
	fmt.Println(msg)
}
