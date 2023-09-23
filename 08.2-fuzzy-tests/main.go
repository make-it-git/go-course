package main

import (
	"fmt"

	"08-tests/math"
)

func main() {
	result := math.AddWithError(2, 2)
	fmt.Println(result) // 4
	result = math.AddWithError(100, 10)
	fmt.Println(result) // 0
}
