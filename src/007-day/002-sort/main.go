package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	last  string
	age   int
}

type personSorter struct {
	data    []person
	compare func(x, y int) bool
}

func (ps personSorter) Len() int {
	return len(ps.data)
}

func (ps personSorter) Less(x, y int) bool {
	return ps.compare(x, y)
}

func (ps personSorter) Swap(x, y int) {
	ps.data[x], ps.data[y] = ps.data[y], ps.data[x]
}

func main() {

	numbers := []int{10, 40, 1, 60, 9}
	fmt.Println(numbers)
	sort.Ints(numbers)
	fmt.Println(numbers)

	names := []string{"James", "Q", "B"}
	fmt.Println(names)
	sort.Strings(names)
	fmt.Println(names)

	p1 := person{first: "James", last: "BOND", age: 35}
	p2 := person{first: "Q", age: 60}
	p3 := person{first: "B", age: 40}

	persons := []person{p1, p2, p3}

	fmt.Println(persons)

	sortableData := personSorter{data: persons, compare: func(x, y int) bool {
		return persons[x].age < persons[y].age
	}}
	sort.Sort(sortableData)

	fmt.Println("By Age", sortableData.data)

	sortableData = personSorter{data: persons, compare: func(x, y int) bool {
		return persons[x].first < persons[y].first
	}}
	sort.Sort(sortableData)
	fmt.Println("By First Name", sortableData.data)

}
