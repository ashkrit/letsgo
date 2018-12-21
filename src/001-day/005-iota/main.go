package main

import "fmt"

const (
	now   = 2018 + iota
	plus1 = 2018 + iota
	plus2 = 2018 + iota
	plus3 = 2018 + iota
)

func main() {

	fmt.Println(now)
	fmt.Println(plus1)
	fmt.Println(plus2)
	fmt.Println(plus3)
}
