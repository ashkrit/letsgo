package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	veh       vehicle
	fourWheel bool
}

type sedan struct {
	veh    vehicle
	luxury bool
}

func main() {

	t1 := truck{veh: vehicle{doors: 2, color: "Black"}, fourWheel: true}
	s1 := sedan{veh: vehicle{doors: 4, color: "grey"}, luxury: true}

	fmt.Println(t1)
	fmt.Println(s1)

	fmt.Println(t1.veh.color, t1.veh.doors, t1.fourWheel)
	fmt.Println(s1.veh.color, s1.veh.doors, s1.luxury)

	address := struct {
		zipCode string
		street  string
	}{"390019", "Near Airport"}

	fmt.Println(address)
}
