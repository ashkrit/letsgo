package main

import "fmt"

func main() {

	message := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			message <- x
		}
		close(message)
	}()

	for m := range message {
		fmt.Println(m)
	}

	fmt.Println("Processing done")

}
