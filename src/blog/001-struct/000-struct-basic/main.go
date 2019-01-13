package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	var p1 person                                      // Zero value
	var p2 = person{}                                  //Zero value
	p3 := person{firstName: "James", lastName: "Bond"} //Proper initialization
	p4 := person{firstName: "James"}                   //Partial initialization

	p5 := new(person) // Using new operator , this returns pointer
	p5.firstName = "James"
	p5.lastName = "Bond"

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)
	fmt.Printf("type %T  value %v ", p5, p5)

}
