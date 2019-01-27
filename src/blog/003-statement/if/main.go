package main

import (
	"fmt"
	"time"
)

func main() {

	x := 10

	//Basic if
	if x > 10 {
		fmt.Println("I am gt ", x)
	} else {
		fmt.Println("I am lt ", x)
	}

	//Sample of if containing init & check
	fmt.Println("Today is", time.Now().Weekday())

	if value := time.Now().Weekday(); value == time.Sunday {
		fmt.Println("Yahoooo.. today is sunday")
	} else {
		fmt.Println("Lets get back to work. I hate", value)
	}
}
