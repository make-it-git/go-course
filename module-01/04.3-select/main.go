package main

import (
	"fmt"
	"time"
)

func main() {
	selectExample()
	selectExampleDefault()
}

func selectExample() {
	one := make(chan string)
	two := make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 100)
		one <- "One"
	}()

	go func() {
		time.Sleep(time.Millisecond * 200)
		two <- "Two"
	}()

	select {
	case result := <-one:
		fmt.Println("Received:", result)
	case result := <-two:
		fmt.Println("Received:", result)
	}

	close(one)
	close(two)
}

func selectExampleDefault() {
	one := make(chan string)
	two := make(chan string)

	go func() {
		one <- "One"
	}()

	go func() {
		two <- "Two"
	}()

	select {
	case result := <-one:
		fmt.Println("Received:", result)
	case result := <-two:
		fmt.Println("Received:", result)
	default:
		fmt.Println("Default...")
	}

	close(one)
	close(two)
	// time.Sleep(time.Millisecond)
}
