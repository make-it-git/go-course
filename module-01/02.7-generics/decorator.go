package main

import "fmt"

func Decorator[T any](fn func(T), decorator func(T) T) func(T) {
	return func(input T) {
		fn((input))
	}
}

func decoratorExample() {
	print := func(n int) { fmt.Println("Number:", n) }
	logger := func(n int) int {
		fmt.Println("called with n", n)
		return n
	}

	decorated := Decorator(print, logger)
	decorated(5)
}
