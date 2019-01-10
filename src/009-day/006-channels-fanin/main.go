package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	appleBasePrice  float64 = 157
	googleBasePrice float64 = 1000
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

	go listenForPrice(googleTicks, googleBasePrice)
	go listenForPrice(appleTicks, appleBasePrice)
	go fanInAllTicks(appleTicks, googleTicks, allTicks)

	for ticks := range allTicks {
		fmt.Println(ticks.when.Format("2006-01-02T15:04:05.999999-07:00"), ticks.symbol, ticks.price)
	}
}

func listenForPrice(stockRicks chan<- float64, basePrice float64) {
	func() {
		for ticks := 0; ticks < 500; ticks++ {
			stockRicks <- math.Round(rand.Float64() * basePrice)
			//time.Sleep(time.Duration(rand.Int31n(100)) * time.Millisecond)
		}
		close(stockRicks)
	}()
}

func fanInAllTicks(appleTicks <-chan float64, googleTicks <-chan float64, allTicks chan<- StockTicks) {
	var countDown sync.WaitGroup
	countDown.Add(2) // 2 stream are available

	go func() {
		for v := range appleTicks {
			allTicks <- StockTicks{symbol: "AAPL", price: v, when: time.Now()}
		}
		countDown.Done() // Mark one stream as complete
	}()

	go func() {
		for v := range googleTicks {
			allTicks <- StockTicks{symbol: "GOOGL", price: v, when: time.Now()}
		}
		countDown.Done() // Mark one stream as completed
	}()

	countDown.Wait()
	close(allTicks)
}
