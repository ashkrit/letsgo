package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("NumGoroutine", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(100)

	var counter int32

	for times := 0; times < 100; times++ {

		go func() {

			atomic.AddInt32(&counter, 1)
			wg.Done()

		}()

	}
	wg.Wait()
	fmt.Println("Counter value", atomic.LoadInt32(&counter))
	fmt.Println("Current NumGoroutine", runtime.NumGoroutine())

}
