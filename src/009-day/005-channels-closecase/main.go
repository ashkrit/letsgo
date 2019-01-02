package main

import "fmt"

func main() {

	message := make(chan string)

	go func() {
		message <- "Message 1"
		message <- "Message 2"
		message <- "Message 3"
		close(message)
	}()

	for {
		v1, ok := <-message
		if ok {
			fmt.Println("Message Received", v1)
		} else {
			fmt.Println("Channel is now closed")
			break
		}
	}

}
