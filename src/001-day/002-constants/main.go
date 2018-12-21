package main

import "fmt"

const x = 10
const name = "James Bond"
const y int64 = 20

func main() {

	fmt.Printf("%v \t %T \n ", x, x)
	fmt.Printf("%v \t %T \n ", name, name)
	fmt.Printf("%v \t %T \n ", y, y)
}
