package main

import "fmt"

func main() {

	/*
		Channels are blocking and default channel(i.e with out capacity) is of zero capacity and
		it need receiver to hand-off message. Communication will only happen when both sender and receiver are ready.
	*/
	c := make(chan int)

	c <- 10

	fmt.Println(<-c)
}
