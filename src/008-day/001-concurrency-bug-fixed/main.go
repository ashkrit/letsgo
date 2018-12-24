package main

import (
	"fmt"
	"runtime"
	"sync"
)

//THis is buggy counter.. count will never go to 100
// run with  go run -race main.go to see data race
/*
	this code is fixing 2 issues
	 - Race condition when counter variable is incremented
	 - Not all the go routines are completed before program finished
*/

func main() {

	fmt.Println("NumGoroutine", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(100)

	var lock sync.Mutex

	counter := 0

	for times := 0; times < 100; times++ {

		go func() {
			lock.Lock()

			counter++
			wg.Done()

			lock.Unlock()
		}()

	}
	wg.Wait()
	fmt.Println("Counter value", counter)
	fmt.Println("Current NumGoroutine", runtime.NumGoroutine())

}
