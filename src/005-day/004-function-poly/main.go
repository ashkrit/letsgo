package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type SHAPE interface {
	area() float64
}

func main() {

	sq := square{length: 10}
	cir := circle{radius: 30}

	info(sq)
	info(cir)
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s circle) area() float64 {
	return math.Pi * s.radius * 2
}

func info(shape SHAPE) {
	value := shape.area()
	fmt.Printf("%T area is %v \n", shape, value)
}
