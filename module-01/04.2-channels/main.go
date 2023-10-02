package main

import (
	"fmt"
)

func echo(arg string, ch chan string) {
	ch <- arg // Send
}

func main() {
	var ch chan string

	fmt.Println(ch) // <nil>

	chInit := make(chan string)
	fmt.Println(chInit) // memory address

	go echo("Hello World", chInit)

	data := <-chInit  // Receive
	fmt.Println(data) // Hello World

	// buffered channels
	ch1 := make(chan string, 2)
	go echo("value 1", ch1)
	go echo("value 2", ch1)

	val1 := <-ch1
	val2 := <-ch1
	fmt.Println(val1, val2) // value 1 value 2 OR value 2 value 1

	close(ch1)

	// send to nil channel blocks forever
	// var c chan string
	// c <- "Hello, World!" // Panic: all goroutines are asleep - deadlock!

	// receive from nil channel blocks forever
	// var c chan string
	// fmt.Println(<-c) // Panic: all goroutines are asleep - deadlock!

	// send to closed channel -> panic
	// var c = make(chan string, 1)
	// c <- "Hello, World!"
	// close(c)
	// c <- "Hello, Panic!" // Panic: send on closed channel

	// A receive from a closed channel returns the zero value immediately.
	var c = make(chan int, 3)
	c <- 20
	c <- 10
	c <- 0
	close(c)
	for i := 0; i < 5; i++ {
		v, ok := <-c
		fmt.Printf("closed?: %v, value %d\n", !ok, v)
		// closed?: false, value 20
		// closed?: false, value 10
		// closed?: false, value 0
		// closed?: true, value 0
		// closed?: true, value 0
	}

	var c2 = make(chan int, 3)
	c2 <- 20
	c2 <- 10
	c2 <- 0
	close(c2)
	for v := range c2 {
		fmt.Printf("value %d\n", v)
		// value 20
		// value 10
		// value 0
	}
}

func directedEcho(arg string, ch chan<- string) {
	ch <- arg // Send Only
}

func directedReceive(ch <-chan string) {
	// ch <- "test" //  invalid operation: cannot send to receive-only channel ch (variable of type <-chan string)
	<-ch
}

func directedReceiveIfClosed(ch <-chan int) {
	value, ok := <-ch
	if ok {
		fmt.Println("process value", value)
	}
}
