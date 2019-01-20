package main

import "fmt"

func main() {

	counter := 0

	counter++
	fmt.Println("In main", counter)
	inc(counter)
	fmt.Println("After inc", counter)

	fmt.Println("Before pointer inc ", counter)
	incByPointer(&counter)
	fmt.Println("After pointer inc ", counter)

}

func inc(value int) {
	value++
	fmt.Println("In Inc", value)
}

func incByPointer(value *int) {
	*value++
	fmt.Printf("In Pointer Inc address %v value %v \n", value, *value)
}
