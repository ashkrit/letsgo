package main

import "fmt"

func main() {

	pingChannel := make(chan string)
	pongChannel := make(chan string)
	taskDoneSignal := make(chan bool)

	go func() {
		for x := 0; x < 10; x++ {
			pingChannel <- fmt.Sprintf("Ping(%v)", x)
			pongChannel <- fmt.Sprintf("Pong(%v)", x)
		}
		close(pingChannel)
		close(pongChannel)
		taskDoneSignal <- true
	}()

	for {
		select {
		case pingMessage := <-pingChannel:
			fmt.Println("PING Received ->", pingMessage)
		case pongMessage := <-pongChannel:
			fmt.Println("PONG Received ->", pongMessage)
		case tDone := <-taskDoneSignal:
			fmt.Println("Done .. last message", tDone)
			return
		default:
		}
	}

	fmt.Println("Processing done")

}
