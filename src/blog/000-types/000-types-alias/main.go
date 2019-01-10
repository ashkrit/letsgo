package main

import "fmt"

//Create Alias for int type
type RichInt int

func (v RichInt) toBinary() string {
	return fmt.Sprintf("%b", v)
}

func main() {

	var ri RichInt
	ri = 100

	fmt.Println("Value of rich int", ri)
	fmt.Println("Convert to Int", int(ri))
	fmt.Println("From int to Rich Int", RichInt(100))
	fmt.Println("Binary Value", ri.toBinary())
}
