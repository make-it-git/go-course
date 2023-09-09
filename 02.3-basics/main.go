package main

import (
	"fmt"
)

func main() {
	var pointer *int

	fmt.Println(pointer) // nil
	if pointer != nil {

	}

	value := 42
	pointer = &value
	fmt.Println(pointer)  // 0x1400010e010
	fmt.Println(*pointer) // 42

	*pointer += 10
	fmt.Println(*pointer) // 52
}
