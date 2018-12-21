package main

import "fmt"

func main() {

	x := [5]int{}
	fmt.Println(x)

	x[0] = 10
	x[1] = 20
	x[2] = 30
	x[3] = 40
	x[4] = 50

	fmt.Println(x)

	for index, value := range x {
		fmt.Println(index, value)
	}

	fmt.Printf("%T \n", x)
}
