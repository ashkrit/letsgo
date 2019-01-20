package main

import "fmt"

const bufferSize = 1000 * 1000 * 1

func main() {

	name := "James Bond"
	doSomething(0, &name, [bufferSize]int{})

}

func doSomething(times int, name *string, buffer [bufferSize]int) {

	fmt.Println(times, name, *name)
	if times > 20 {
		return
	}
	doSomething(times+1, name, buffer)
}
