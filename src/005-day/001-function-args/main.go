package main

import "fmt"

func main() {

	values := []int{10, 20, 30}
	total := foo(values...)
	fmt.Println(total)

	total = bar(values)
	fmt.Println(total)
}

func foo(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}

func bar(values []int) int {
	total := 0

	for _, value := range values {
		total += value
	}
	return total
}
