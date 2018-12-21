package main

import "fmt"

func main() {

	x := []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println(x)

	//Dropping 30
	x = append(x[0:2], x[3:]...)

	fmt.Println(x)
}
