package main

import "fmt"

type stock struct {
	symbol string
	price  float64
}

func main() {

	s1 := allocateOnStack()
	fmt.Printf("Type =  %T value= %v \n", s1, s1)
	s2 := allocateOnHeap()
	fmt.Printf("Type = %T value= %v \n", s2, s2)
}

func allocateOnStack() stock {

	google := stock{symbol: "GOOG", price: 1109}
	return google
}

func allocateOnHeap() *stock {

	google := stock{symbol: "GOOG", price: 1109}
	return &google
}
