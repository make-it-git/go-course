package main

import "fmt"

func Map[T any, U any](input []T, mapper func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = mapper(v)
	}
	return result
}

func Filter[T any](input []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range input {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T any, U any](input []T, initial U, reducer func(U, T) U) U {
	acc := initial
	for _, v := range input {
		acc = reducer(acc, v)
	}
	return acc
}

func first[T any](s []T) T {
	return s[0]
}

func functionalExamples() {
	nums := []int{1, 2, 3, 4, 5, 6}

	sumOfSquaresOfEvens := Reduce(
		Map(Filter(nums, func(x int) bool { return x%2 == 0 }),
			func(x int) int { return x * x }),
		0,
		func(acc, x int) int { return acc + x },
	)

	fmt.Println("Sum of squares of evens:", sumOfSquaresOfEvens)

	s := []int{200, 300, 400}
	fmt.Printf("First of %s = %d\n", s, first(s))
}
