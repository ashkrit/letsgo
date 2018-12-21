package main

import "fmt"

func main() {

	for x := 65; x < 97; x++ {
		fmt.Println(x)
		for times := 0; times < 3; times++ {
			fmt.Printf("\t%#U \n", x)
		}
	}
}
