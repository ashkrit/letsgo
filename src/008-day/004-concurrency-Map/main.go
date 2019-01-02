package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Current NumGoroutine", runtime.NumGoroutine())
	var cache sync.Map

	var ele []string
	const maxElement = 10000
	const rangeValue = 999
	for v := 0; v < maxElement; v++ {
		ele = append(ele, fmt.Sprintf("%v", rand.Intn(rangeValue)))
	}

	var wq sync.WaitGroup
	wq.Add(len(ele))

	for v := range ele {
		go func(item int) {
			cache.Store(item, item)
			wq.Done()
		}(v)
	}

	wq.Wait()

	foundKeys := 0
	missingKeys := 0
	for v := range ele {
		_, e := cache.Load(v)
		if e {
			foundKeys++
		} else {
			missingKeys++
		}
	}

	fmt.Println("Found Key", foundKeys)
	fmt.Println("Missing Key", missingKeys)

}
