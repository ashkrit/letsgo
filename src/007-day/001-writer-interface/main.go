package main

import (
	"fmt"
)

/*
 Create new writer that keeps message in memory
*/

type InMemory struct {
	buffer *[]byte
	count  *int
}

func (w InMemory) message() string {
	return string(*w.buffer)
}

func (w InMemory) messageCount() int {
	return *w.count
}

func (w InMemory) Write(p []byte) (n int, err error) {
	*w.buffer = append(*w.buffer, p...)
	*w.count++
	return len(p), nil
}

func main() {

	var message []byte
	var messageCount int
	inMemoryWriter := InMemory{buffer: &message, count: &messageCount}

	fmt.Fprintln(inMemoryWriter, "Message 1")
	fmt.Fprintln(inMemoryWriter, "Message 2")
	fmt.Fprintln(inMemoryWriter, "Message 3")

	fmt.Println("No of message", inMemoryWriter.messageCount())
	fmt.Println(inMemoryWriter.message())

}
