package main

import "fmt"

func main() {

	favSport := "Tennis"
	switch favSport {
	case "Tennis":
		fmt.Println("Tennis lover")
	case "Cricket":
		fmt.Println("Cricket lover")
	default:
		fmt.Println("My fav sport is missing", favSport)
	}
}
