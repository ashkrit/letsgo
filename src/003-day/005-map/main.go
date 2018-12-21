package main

import "fmt"

func main() {

	favItems := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(favItems)
	fmt.Println(favItems["bond_james"])

	for k, v := range favItems {

		fmt.Printf("%v likes are \n", k)
		for index, value := range v {
			fmt.Printf("\t %v %v \n", index, value)
		}
	}

	//Add elements
	favItems["Ashkrit"] = []string{"Google"}

	for k, v := range favItems {

		fmt.Printf("%v likes are \n", k)
		for index, value := range v {
			fmt.Printf("\t %v %v \n", index, value)
		}
	}

	//Delete

	delete(favItems, "Ashkrit")

	for k, v := range favItems {

		fmt.Printf("%v likes are \n", k)
		for index, value := range v {
			fmt.Printf("\t %v %v \n", index, value)
		}
	}

	//Check if key exists

	value, ok := favItems["Ashkrit"]
	fmt.Println("Value", value)
	fmt.Println("Key exists", ok)
}
