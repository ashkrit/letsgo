package main

import "fmt"

func main() {

	x := 42
	y := 20

	fmt.Println("==", x == y)
	fmt.Printf("%d <= %d %v \n", x, y, x <= y)
	fmt.Printf("%d >= %d %v \n", x, y, x >= y)
	fmt.Printf("%d != %d %v \n", x, y, x != y)
	fmt.Printf("%d > %d %v \n", x, y, x > y)
	fmt.Printf("%d < %d %v \n", x, y, x < y)

}
