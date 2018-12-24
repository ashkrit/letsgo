package main

import (
	"fmt"
	"runtime"
)

//THis is buggy counter.. count will never go to 100
// run with  go run -race main.go to see data race
/*
	this code as 2 issues
	 - Race condition when counter variable is incremented
	 - Not all the go routines are completed before program finished
*/

func main() {

	fmt.Println("NumGoroutine", runtime.NumGoroutine())
	counter := 0
	for times := 0; times < 100; times++ {

		go func() {
			counter++
		}()

	}

	fmt.Println("Counter value", counter)
	fmt.Println("Current NumGoroutine", runtime.NumGoroutine())

}
