package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
}

func main() {

	p1 := person{FirstName: "James", LastName: "Bond"}
	fmt.Println(p1)

	toJson(p1)
	fromJson()

}

func fromJson() {
	content := `
	{"FirstName":"James","LastName":"Bond"}
	`
	p2 := person{}
	error := json.Unmarshal([]byte(content), &p2)
	if error != nil {
		fmt.Println("Error in parsing...", error)
	}
	fmt.Println(p2)
}

func toJson(p1 person) {
	data, erro := json.Marshal(p1)
	if erro != nil {
		fmt.Println("Error ", erro)
	}
	fmt.Println(string(data))
}
