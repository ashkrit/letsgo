package main

import "fmt"

func main() {

	for value := 10; value < 100; value++ {
		fmt.Printf("Mod for %v / %v is %v \n", value, 4, value%4)
	}
}
