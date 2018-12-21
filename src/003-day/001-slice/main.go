package main

import "fmt"

func main() {

	x := []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println(x)
	for index, value := range x {
		fmt.Println(index, value)
	}

	fmt.Printf("%T \n", x)

	fmt.Println(x[:4])
	fmt.Println(x[4:])
	fmt.Println(x[4:6])

	x = append(x, 80)

	fmt.Println(x)

	x = append(x, 90, 100, 110)

	fmt.Println(x)

	x = append(x, []int{200, 210}...)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
