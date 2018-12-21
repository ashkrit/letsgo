package main

import "fmt"

type bigint int

var x bigint

func main() {

	fmt.Println(x)
	fmt.Printf("%T \n", x)

	x = 42

	fmt.Println(x)
}
