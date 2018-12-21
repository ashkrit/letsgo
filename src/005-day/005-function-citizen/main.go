package main

import (
	"fmt"
	"time"
)

func main() {

	func() {
		fmt.Println("Hello")
	}()

	fn := func() {
		fmt.Println("Today is", time.Now())
	}
	fn()

	sgdInrFn := sgdToInr()
	value := sgdInrFn(1999)
	fmt.Printf("%f \n", value)
}

func sgdToInr() func(value float64) float64 {
	return func(value float64) float64 {
		return value * 52.6
	}
}
