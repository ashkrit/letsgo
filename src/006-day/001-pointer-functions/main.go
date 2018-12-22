package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	last  string
}

func (p person) changeMe() {
	//Swap first & last name
	f := p.first
	p.first = p.last
	p.last = f
	fmt.Println("Inside change me ", p, &p)
}

func (p *person) changeMeP() {
	//Swap first & last name
	f := p.first
	(*p).first = p.last
	(*p).last = f
	fmt.Println("Inside change me pointer ", p, &p)
}

type UpdatePerson interface {
	changeMeP()
}

func swapNames(p UpdatePerson) {
	p.changeMeP()
}

func changeMeAgain(p *person) {

	f := p.first
	(*p).first = p.last
	(*p).last = f
	fmt.Println("Inside change me pointer ", p, &p)
}

func main() {

	p := person{first: "James", last: "Bond"}
	fmt.Println(p, &p)

	//compiler allows to pass value or reference
	p = person{first: "James", last: "Bond"}
	p.changeMe()
	fmt.Println("Outside", p, &p)

	//Methods compiler allows to pass value or reference
	p = person{first: "James", last: "Bond"}
	(&p).changeMe()
	fmt.Println("Outside", p, &p)

	//Methods compiler allows to pass value or reference
	p = person{first: "James", last: "Bond"}
	p.changeMeP()
	fmt.Println("Outside", p, &p)

	//Methods compiler allows to pass value or reference
	p = person{first: "James", last: "Bond"}
	(&p).changeMeP()
	fmt.Println("Outside", p, &p)

	p = person{first: "James", last: "Bond"}
	changeMeAgain(&p)
	fmt.Println("Outside", p, &p)

	//compiler allows only by reference when method is exposed as Interface or using method set
	p = person{first: "James", last: "Bond"}
	swapNames(&p)
	fmt.Println("Outside", p, &p)

	sort.Ints()
}
