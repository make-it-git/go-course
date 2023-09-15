package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var i = rand.Int() % 10
	var a [8]int

	// panic: runtime error: index out of range [9] with length 8
	fmt.Println(a[i])

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	defer handlePanic() // recovered from error stop
	unrecoverableError()

	// 3
	// 2
	// 1
}

func unrecoverableError() {
	panic("stop")
}

func handlePanic() {
	err := recover()
	if err != nil {
		fmt.Printf("recovered from error %s\n", err)
	}
}
