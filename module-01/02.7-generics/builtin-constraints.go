package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type customString string

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func genericMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func builtinConstraints() {
	fmt.Printf("min(10, 20)=%d\n", min(10, 20))
	fmt.Printf("genericMin(30, 40)=%d\n", genericMin(30, 40))
	fmt.Printf("genericMin(100, 50)=%f\n", genericMin[float64](100.01, 100.02))

	a := customString("abbb")
	b := customString("aaaa")
	fmt.Printf("genericMin[customString](a, b)=%s\n", genericMin[customString](a, b))
}
