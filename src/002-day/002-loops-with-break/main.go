package main

import "fmt"

func main() {

	years := 2000
	for {
		fmt.Println(years)
		years++

		if years > 2018 {
			break
		}
	}
}
