package main

import "fmt"

func sliceReverse(s []int) {
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Printf("[%d]=%d\n", i, s[i])
	}
}

type Iter[E any] func(body func(index int, value E))

func Backward[S ~[]E, E any](s S) Iter[E] {
	return func(body func(int, E)) {
		for i := len(s) - 1; i >= 0; i-- {
			body(i, s[i])
		}
	}
}

func funcRangeExample() {
	s := []int{10, 20, 30}
	fmt.Printf("reverse of slice %s\n", s)
	sliceReverse(s)

	fmt.Printf("backward of slice %s\n", s)
	backwardIter := Backward(s)
	backwardIter(func(index int, value int) {
		fmt.Printf("[%d]=%d\n", index, value)
	})
}
