package main

import "fmt"

func main() {

	x := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}
	fmt.Println(x)
	for _, value := range x {
		fmt.Println(value)
	}

	fmt.Printf("%T \n", x)
}
