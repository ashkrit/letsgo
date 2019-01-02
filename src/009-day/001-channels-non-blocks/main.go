package main

import "fmt"

func main() {

	/*
	Channels are blocking and default channel(i.e with out capacity) is of zero capacity and
	it need receiver to hand-off message.

		In this example go routine is started that will publish message
	*/

	c := make(chan int)

	go func() {
		fmt.Println("Adding message")
		c <- 10
		fmt.Println("Message added")

	}()

	fmt.Println("Now working.....", <-c)
}
