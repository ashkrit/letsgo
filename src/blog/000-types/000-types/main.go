package main

import "fmt"

func main() {

	//Basic Types - https://golang.org/ref/spec#Types

	var value int
	var f float64
	var b bool
	var by byte
	var name string
	var x rune

	//Variable are only declared by compiler initialized to ZERO VALUE of its type
	//https://golang.org/ref/spec#The_zero_value

	fmt.Printf("value %T -> (%v) \n", value, value)
	fmt.Printf("f %T -> (%v) \n", f, f)
	fmt.Printf("b %T -> (%v) \n", b, b)
	fmt.Printf("by %T -> (%v) \n", by, by)
	fmt.Printf("name %T -> (%v) \n", name, name)
	fmt.Printf("x %T -> (%v) \n", x, x)

	fmt.Println("*******************")

	value1 := 10
	f1 := 3.14
	b1 := true
	name1 := "Say Hello"
	x1 := 100
	fmt.Printf("value %T -> (%v) \n", value1, value1)
	fmt.Printf("f %T -> (%v) \n", f1, f1)
	fmt.Printf("b %T -> (%v) \n", b1, b1)
	fmt.Printf("name %T -> (%v) \n", name1, name1)
	fmt.Printf("x %T -> (%v) \n", x1, x1)

}
