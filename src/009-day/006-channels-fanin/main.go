package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type StockTicks struct {
	when   time.Time
	symbol string
	price  float64
}

func main() {

	appleTicks := make(chan float64)
	googleTicks := make(chan float64)
	allTicks := make(chan StockTicks)

	go listenForGooglePrice(googleTicks)
	go listenForApplePrice(appleTicks)
	go fanInAllTicks(appleTicks, googleTicks, allTicks)

	for ticks := range allTicks {
		fmt.Println(ticks.when.Format("2006-01-02T15:04:05.999999-07:00"), ticks.symbol, ticks.price)
	}
}

func listenForApplePrice(appleTicks chan<- float64) {
	func() {
		for ticks := 0; ticks < 5; ticks++ {
			appleTicks <- math.Round(rand.Float64() * 157)
			time.Sleep(time.Duration(rand.Int31n(100)) * time.Millisecond)
		}
		close(appleTicks)
	}()
}

func listenForGooglePrice(googleTicks chan<- float64) {
	func() {
		for ticks := 0; ticks < 5; ticks++ {
			googleTicks <- math.Round(rand.Float64() * 1000)
			time.Sleep(time.Duration(rand.Int31n(100)) * time.Millisecond)
		}
		close(googleTicks)
	}()
}

func fanInAllTicks(appleTicks <-chan float64, googleTicks <-chan float64, allTicks chan<- StockTicks) {
	var countDown sync.WaitGroup
	countDown.Add(2)

	go func() {
		for v := range appleTicks {
			allTicks <- StockTicks{symbol: "AAPL", price: v, when: time.Now()}
		}
		countDown.Done()
	}()

	go func() {
		for v := range googleTicks {
			allTicks <- StockTicks{symbol: "GOOGL", price: v, when: time.Now()}
		}
		countDown.Done()
	}()

	countDown.Wait()
	close(allTicks)
}
