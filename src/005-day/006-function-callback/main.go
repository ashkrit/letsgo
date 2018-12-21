package main

import "fmt"

func main() {

	process(callbackHandler)

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	gt5 := func(i int) bool {
		return i > 5
	}

	by2 := func(i int) int {
		return i * 2
	}

	filterValues := RichIntSlice(values).
		filter(gt5).
		mapValue(by2)

	fmt.Println(filterValues)

}

func callbackHandler(message string) {
	fmt.Println("Message received ->", message)
}

func process(f func(message string)) {
	fmt.Println("Processing started")
	f("Started")

	f("In progress")

	fmt.Println("Completed")
	f("Completed")
}

type RichIntSlice []int

func (r RichIntSlice) filter(f func(int) bool) RichIntSlice {
	var filterValues []int
	for _, value := range []int(r) {
		if f(value) {
			filterValues = append(filterValues, value)
		}
	}
	return RichIntSlice(filterValues)
}

func (r RichIntSlice) mapValue(f func(int) int) RichIntSlice {
	var filterValues []int
	for _, value := range []int(r) {
		filterValues = append(filterValues, f(value))
	}
	return RichIntSlice(filterValues)
}
