package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type anotherperson struct {
	firstName string
	lastName  string
}

func main() {

	p1 := person{firstName: "James", lastName: "Bond"}
	p2 := person{firstName: "James", lastName: "Bond"}
	anotherp1 := anotherperson{firstName: "James", lastName: "Bond"}

	if p1 == p2 {
		fmt.Println("Same person found!!!!", p1)
	} else {
		fmt.Println("They are different", p1, p2)
	}

	fmt.Println("Another person", anotherp1)

	//Simple conversion

	p1 = person(anotherp1)
	anotherp1.lastName = "Lee" // Will have not effect on p1

	fmt.Println(p1, anotherp1)
}
