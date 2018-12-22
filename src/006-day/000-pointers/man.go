package main

import "fmt"

func main() {

	x := 20
	fmt.Println(x, &x)

	y := &x
	fmt.Println(*y, y)
	incByValue(x)
	fmt.Println(x, &x)

	incByPointer(y)
	fmt.Println(*y, y)
}

func incByValue(z int) {
	z++
	fmt.Println(z, &z)
}

func incByPointer(z *int) {
	*z++
	fmt.Println(*z, z)
}
