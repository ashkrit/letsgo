package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and", p.age, "years old")
}

func main() {

	p1 := person{first: "James", last: "Bond", age: 30}
	p1.speak()

	p2 := person{first: "Ian", last: "Fleming", age: 55}
	p2.speak()
}
