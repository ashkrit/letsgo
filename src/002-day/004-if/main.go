package main

import "fmt"

func main() {

	x := 110

	if x <= 10 {
		fmt.Println("Less than eq to 10")
	} else if x > 100 {
		fmt.Println("is gt than 100")
	} else {
		fmt.Println("No luck")
	}
}
