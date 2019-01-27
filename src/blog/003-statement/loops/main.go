package main

import (
	"fmt"
)

func main() {

	//C/C++/Java like
	value := 0
	for counter := 0; counter < 10; counter++ {
		value++
	}
	fmt.Println(value)

	//Just condition
	value = 0
	for value < 10 {
		value++
	}
	fmt.Println(value)

	//Infinite (With no condition)
	for {
		value++
		if value > 10 {
			break
		}
	}

	//for each loop for array/collections
	days := []string{"Sunday", "Monday"}
	for index, value := range days {
		fmt.Println("Index ", index, "Value ", value)
	}
}
