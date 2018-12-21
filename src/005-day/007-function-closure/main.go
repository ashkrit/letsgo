package main

import "fmt"

func main() {

	systemLogger := logWithContext("system")
	webLogger := logWithContext("webuser")

	systemLogger("This is first message")
	webLogger("Loading home page")

	systemLogger("This is second message")
	webLogger("Loading Products page")

	tracker1 := tracker()
	tracker2 := tracker()

	t1Value := tracker1()
	t1Value = tracker1()
	t1Value = tracker1()

	t2Value := tracker2()

	fmt.Println(t1Value)
	fmt.Println(t2Value)
}

func logWithContext(context string) func(string) {

	return func(message string) {
		fmt.Println("["+context+"]", message)
	}

}

func tracker() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
