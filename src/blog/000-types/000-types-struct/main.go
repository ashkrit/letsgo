package main

import (
	"fmt"
)

type address struct {
	address1 string
	address2 string
	city     string
}

type contact struct {
	landLine int
	mobile   int
}

type person struct {
	firstName      string
	lastName       string
	age            int
	add            address
	contactDetails contact
}

func main() {

	var p1 person
	fmt.Println(p1)

	p2 := person{firstName: "James", lastName: "Bond", age: 35}
	fmt.Println(p2)

	p3 := person{firstName: "James", lastName: "Bond", age: 35,
		add:            address{address1: "30 Wellington Square", address2: "Street 81"},
		contactDetails: contact{mobile: 11119999}}

	fmt.Println(p3)

}
