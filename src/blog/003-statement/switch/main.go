package main

import (
	"fmt"
)

func main() {

	//Simple switch case
	value := 10
	switch value {
	case 10:
		fmt.Println("Value is 10")
	default:
		fmt.Println("Some other value than 10")
	}

	//Switch with no expression and dynamic condition
	switch {
	case value >= 10:
		fmt.Println("Value is gt 10")
	case value >= 20:
		fmt.Println("Value is gt 20")
	}

	//Switch with multiple condition in single case
	specialValue := '@'
	switch specialValue {
	case '@', '!', '#':
		fmt.Println("This is special value")
	default:
		fmt.Println("This is normal value")
	}

	//Switch for type assertion
	var t interface{}
	t = "James"
	switch t.(type) {
	case int:
		fmt.Println("Int value", t)
	case string:
		fmt.Println("String value", t)
	}

}
