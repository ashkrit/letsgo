package main

import "fmt"

func main() {

	x := foo()
	y, message := bar()
	fmt.Println(x, y, message)
}

func foo() int {
	return 10
}

func bar() (int, string) {
	return 10, "Hello"
}
