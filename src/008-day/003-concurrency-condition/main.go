package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type Queue struct {
	buffer         []int
	capacity       int
	fullCondition  *sync.Cond
	emptyCondition *sync.Cond
	writeIndex     int
	readIndex      int
	totalElement   int
}

func (q *Queue) init() {
	var lock sync.Mutex
	q.fullCondition = sync.NewCond(&lock)
	q.emptyCondition = sync.NewCond(&lock)
}

func (q *Queue) enqueue(message int) {

	locker := q.fullCondition.L
	locker.Lock()

	for q.isFull() {
		q.fullCondition.Wait()
	}

	offset := q.writeIndex % q.capacity
	q.buffer[offset] = message

	q.writeIndex++
	q.totalElement++

	fmt.Printf("Message Added %v \t size %v \n", message, q.totalElement)
	q.emptyCondition.Signal()
	locker.Unlock()
}

func (q *Queue) isFull() bool {
	return q.totalElement >= q.capacity
}

func (q *Queue) dequeue() int {

	locker := q.emptyCondition.L
	locker.Lock()

	for q.isEmpty() {
		q.emptyCondition.Wait()
	}

	index := q.readIndex % q.capacity
	value := q.buffer[index]
	q.buffer[index] = 0

	q.readIndex++
	q.totalElement--

	q.fullCondition.Signal()
	locker.Unlock()
	return value
}

func (q *Queue) isEmpty() bool {
	return q.totalElement == 0
}

func main() {

	fmt.Println("NumGoroutine", runtime.NumGoroutine())
	fmt.Println("CPUs", runtime.NumCPU())

	const maxElement = 10
	queue := Queue{buffer: make([]int, maxElement), capacity: maxElement}
	queue.init()

	const maxMessage = 100
	go func() {
		//Consumer
		for {
			value := queue.dequeue()
			fmt.Println("C", "Message Received", value)
			//time.Sleep(time.Millisecond * 10)
		}

	}()

	go func() {
		//Consumer
		for {
			value := queue.dequeue()
			fmt.Println("A", "Message Received", value)
			//time.Sleep(time.Millisecond * 10)
		}

	}()

	go func() {
		//Consumer
		for {
			value := queue.dequeue()
			fmt.Println("B", "Message Received", value)
			//time.Sleep(time.Millisecond * 10)
		}

	}()

	for mcount := 0; mcount < maxMessage; mcount++ {
		message := int(rand.Int31n(maxMessage))
		queue.enqueue(message)
		//time.Sleep(time.Millisecond * 5)
	}
	fmt.Println("Producer and consumer started")

	fmt.Println("All done")
}
