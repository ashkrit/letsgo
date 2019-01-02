package main

import "fmt"

/*
	3 types of channel
	 - Bi directional channel  - Both send & receive is possible
	 - Send channel  - Only send is possible and compiler gives and error
	- Receive channel  - Only receive is possible
*/
func main() {

	message := make(chan int)          // Send & Receive
	sendChannel := make(chan<- int)    // Send only
	receiveChannel := make(<-chan int) // Receive only

	fmt.Printf("%T \n", message)
	fmt.Printf("%T \n", sendChannel)
	fmt.Printf("%T \n", receiveChannel)

	sendChannel = message

	fmt.Printf("%T \n", sendChannel)

	go pingMessage(message)

	pongMessage(message)

}

func pingMessage(c chan<- int) {

	c <- 100
}

func pongMessage(c <-chan int) {

	pMessage := <-c
	fmt.Println("Pong Message", pMessage)
}
