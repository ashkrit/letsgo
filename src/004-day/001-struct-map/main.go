package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	favIce    []string
}

func main() {

	p1 := person{firstName: "Ashkrit", lastName: "Sharma",
		favIce: []string{"chocolate", "vanilla"}}

	p2 := person{firstName: "James", lastName: "Bond",
		favIce: []string{"strawberry", "vanilla"}}

	persons := map[string]person{p2.lastName: p2,
		p1.lastName: p1}

	p := persons["Sharma"]
	fmt.Printf("%v %v likes \n", p.firstName, p.lastName)
	for _, name := range p.favIce {
		fmt.Printf("\t %v", name)
	}

}
