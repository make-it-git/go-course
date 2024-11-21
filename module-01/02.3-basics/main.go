package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var pointer *int
	var pointer2 *int

	fmt.Println(pointer) // nil

	// fmt.Println(*pointer)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x100cb5768]

	// How pointers work in C: https://beej.us/guide/bgc/html/split/pointers.html#pointers
	// Pointer Arithmetic: https://beej.us/guide/bgc/html/split/pointers2.html#pointers2

	if pointer != nil {

	}

	value := 42
	pointer = &value
	fmt.Println(pointer)  // 0x1400010e010
	fmt.Println(*pointer) // 42

	*pointer += 10
	fmt.Println(*pointer) // 52

	var increment = func(x *int) {
		*x++
	}

	var increment2 = func(x *int) {
		(*x)++
	}

	increment(pointer)
	fmt.Println(*pointer) // 53
	increment(&value)
	fmt.Println(*pointer) // 54
	increment2(&value)
	fmt.Println(*pointer) // 55

	pointer2 = &value
	fmt.Println(pointer2 == pointer) // true

	var pointerToSlice *[]int
	slice := []int{100, 200, 300}
	pointerToSlice = &slice
	fmt.Println(pointerToSlice)       // &[100 200 300]
	fmt.Println(*pointerToSlice)      // [100 200 300]
	fmt.Println((*pointerToSlice)[1]) // 200

	// pointer Arithmetic
	unsafePointerToSlice := unsafe.Pointer(pointerToSlice)
	fmt.Println((*[]int)(unsafePointerToSlice))  // &[100 200 300]
	fmt.Println(*(*[]int)(unsafePointerToSlice)) // [100 200 300]

	newUnsafePointerToSlice := unsafe.Pointer(uintptr(unsafePointerToSlice) + unsafe.Sizeof((*pointerToSlice)[0]))

	fmt.Println(unsafePointerToSlice, newUnsafePointerToSlice)                    // 0x1400000c018 0x1400000c020
	fmt.Println(uintptr(newUnsafePointerToSlice) - uintptr(unsafePointerToSlice)) // 8
}
