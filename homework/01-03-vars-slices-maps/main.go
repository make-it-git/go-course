package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	shadowedVariables("a1")
	overflows()
	sliceProblem()
}

func shadowedVariables(a string) {
	// many
	// lines
	// of
	// code
	// here
	a, b := "a2", "b2" // do not change these variables names
	// many
	// lines
	// of
	// code
	// here
	_ = b
	_ = a
	// Do not remove if/else/fmt.Println
	if a == "a1" {
		fmt.Println("Hooray! a == 'a1'")
	} else {
		fmt.Println("Error: shadowedVariables, a != 'a1'")
	}
}

// May be better data type?
func overflows() {
	var counter int32 = getInt32() + 10
	if strconv.Itoa(int(counter)) == "2147483657" {
		fmt.Println("Hooray! counter = 2147483657")
	} else {
		fmt.Println("Error: overflows, counter != 2147483647")
	}
}

func getInt32() int32 {
	return math.MaxInt32
}

// experiment with capacity
func sliceProblem() {
	x := make([]int, 0, 3)
	x = append(x, 10)
	x = append(x, 20)

	y := x
	y = append(y, 30)
	y[0] = 11
	if x[0] == 11 {
		fmt.Println("Error: x[0] != 10")
	}

	y = append(y, 40)
	y[0] = 12

	fmt.Println(x, "expected", "[10 20]")
	fmt.Println(y, "expected", "[12 20 30 40]")
}
