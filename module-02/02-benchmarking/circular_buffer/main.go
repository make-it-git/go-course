package main

import (
	"fmt"
	"sync"
	"time"
)

type CircularBuffer struct {
	data       []interface{}
	readIndex  int
	writeIndex int
	size       int
	mutex      sync.Mutex
	cond       *sync.Cond
}

func NewCircularBuffer(size int) *CircularBuffer {
	cb := &CircularBuffer{
		data: make([]interface{}, size),
		size: size,
	}
	cb.cond = sync.NewCond(&cb.mutex)
	return cb
}

// Write method - blocks when buffer is full
func (cb *CircularBuffer) Write(value interface{}) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	// Wait until there is space to write
	for (cb.writeIndex+1)%cb.size == cb.readIndex {
		cb.cond.Wait()
	}

	cb.data[cb.writeIndex] = value
	fmt.Printf("%v, writeIndex=%d written\n", value, cb.writeIndex)
	cb.writeIndex = (cb.writeIndex + 1) % cb.size

	cb.cond.Signal()
}

// Read method - blocks when buffer is empty
func (cb *CircularBuffer) Read() interface{} {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	// Wait until there is data to read
	for cb.readIndex == cb.writeIndex {
		cb.cond.Wait()
	}

	value := cb.data[cb.readIndex]
	fmt.Printf("%v, readIndex=%d read\n", value, cb.readIndex)
	cb.readIndex = (cb.readIndex + 1) % cb.size

	// Signal that there's space available for writing
	cb.cond.Signal()
	return value
}

func producer(cb *CircularBuffer, id int) {
	for {
		data := fmt.Sprintf("producer %d", id)
		cb.Write(data)
		time.Sleep(time.Second)
	}
}

func consumer(cb *CircularBuffer, id int) {
	for {
		// Simulate consuming data
		data := cb.Read()
		fmt.Printf("consumer %d: read \"%s\"\n", id, data)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	cb := NewCircularBuffer(5)

	for i := 1; i <= 2; i++ {
		go producer(cb, i)
		go consumer(cb, i)
	}

	// Keep the main function running indefinitely
	select {}
}
