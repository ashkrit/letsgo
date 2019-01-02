package main

import "fmt"

func main() {

	/*
		Channels are blocking and it need receiver to hand-off message.
		In this example channel is created with 1 capacity. This channel can hold 1 message without blocking.
		If one more message is added then execution will block

	*/

	c := make(chan int, 1)

	c <- 10

	fmt.Println("Now working.....", <-c)
}
