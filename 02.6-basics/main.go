package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var arr [4]int
	fmt.Println(arr[0]) // element at index 0
	fmt.Println(arr[3])
	// fmt.Println(arr[4]) // invalid index

	// python tuple example
	// arr = (100500, "i'am string")

	fmt.Println(arr) // [0 0 0 0]

	var arr2 [4]int = [4]int{10, 20, 30, 40}
	fmt.Println(arr2) // [10 20 30 40]

	arr3 := [4]int{10, 20, 30, 40}
	fmt.Println(arr3) // [10 20 30 40]

	// C style for loop
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("Index: %d, Element: %d\n", i, arr3[i])
	}

	// python style enumerate
	// for index, value in enumerate(...)
	for i, value := range arr3 {
		fmt.Printf("Index: %d, Element: %d\n", i, value)
	}

	for _, value := range arr3 {
		fmt.Printf("Element: %d\n", value)
	}

	for i := range arr3 {
		fmt.Printf("Index: %d\n", i)
	}

	for range arr3 {
		fmt.Println("?")
	}

	multiDimArr := [...][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	for i, el := range multiDimArr {
		// Index: 0, Element: [1 2 3 4]
		// Index: 1, Element: [5 6 7 8]
		fmt.Printf("Index: %d, Element: %d\n", i, el)
	}

	a := [4]int{1, 2, 3, 4}
	fmt.Println(a)
	// var b [2]int = a // cannot use a (variable of type [4]int) as [2]int value in variable declaration

	a1 := [3]string{"One", "Two", "Three"}
	var b1 = a1 // Copy of a1

	b1[0] = "One!!!"

	fmt.Println(a1) // [One Two Three]
	fmt.Println(b1) // [One!!! Two Three]

	// slices

	// https://go.dev/src/runtime/slice.go
	// type slice struct {
	//	array unsafe.Pointer
	//	len   int
	//	cap   int
	//}

	var zeroSlice []int
	fmt.Println(zeroSlice, zeroSlice == nil, len(zeroSlice)) // [] true 0

	var oneElementSlice = make([]int, 1, 10)
	fmt.Println(oneElementSlice, len(oneElementSlice), cap(oneElementSlice)) // [0] 1 10

	sliceWithLenEqCap := make([]int, 5)
	fmt.Println(sliceWithLenEqCap, len(sliceWithLenEqCap), cap(sliceWithLenEqCap)) // [0 0 0 0 0] 5 5

	myFavSlice := []string{"I", "like", "learning", "Go"}
	fmt.Println(myFavSlice) // [I like learning Go]

	fmt.Println(myFavSlice[1:2]) // like
	fmt.Println(myFavSlice[:3])  // I like learning

	myArr := [5]int{20, 15, 5, 30, 25}

	mySlice := myArr[1:4]

	// Array: [20 15 5 30 25], Length: 5, Capacity: 5
	fmt.Printf("Array: %v, Length: %d, Capacity: %d\n", myArr, len(myArr), cap(myArr))

	// Slice [15 5 30], Length: 3, Capacity: 4
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", mySlice, len(mySlice), cap(mySlice))

	for _, v := range myFavSlice {
		fmt.Println(v)
		// I
		// like
		// learning
		// Go
	}

	// func copy(dst, src []T) int
	s1 := []string{"a", "b", "c", "d"}
	s2 := make([]string, len(s1))

	e := copy(s2, s1)

	fmt.Println("Src:", s1)     // Src: [a b c d]
	fmt.Println("Dst:", s2)     // Dst: [a b c d]
	fmt.Println("Elements:", e) // Elements: 4

	s3 := make([]string, 2)
	e2 := copy(s3, s1)
	fmt.Println(s3, e2) // [a b] 2

	s3 = append(s3, "c")
	fmt.Println(s3) // [a b c]

	s3 = append(s3, "d", "e")
	fmt.Println(s3) // [a b c d e]
	// func append(slice []Type, elems ...Type) []Type

	// maps
	var m map[string]int
	fmt.Println(m, m == nil) // map[] true
	// m["test"] = 1 // panic: assignment to entry in nil map

	var m2 = make(map[string]int)
	m2["test"] = 1
	fmt.Println(m2) // map[test:1]

	// map literal
	var m3 = map[string]int{
		"hi": 100,
	}
	fmt.Println(m3) // map[hi:100]

	var m4 = map[string][]int{
		"a": []int{10, 20},
		"b": []int{30, 40},
	}
	fmt.Println(m4) // map[a:[10 20] b:[30 40]]

	m4["c"] = []int{50, 60, 70}

	fmt.Println(m4["c"])                 // [50 60 70]
	fmt.Println(m4["d"], m4["d"] == nil) // [] true

	value, ok := m4["d"]
	if !ok {
		fmt.Println("does not exist")
	} else {
		fmt.Println("key found", value)
	}

	delete(m4, "d")

	for k, v := range m4 {
		fmt.Printf("Key: %s, Value: %v\n", k, v)
	}

	m5 := m2
	m5["new"] = 10
	fmt.Println(m5, m2) // map[new:10 test:1] map[new:10 test:1]

	// runes

	word := `Straße`                          // 6 chars?
	fmt.Println(len(word))                    // 7
	fmt.Println(word[:2])                     // St
	fmt.Println(word[4:])                     // ße
	fmt.Println(word[4:6])                    // ß
	fmt.Println(word[5:])                     // �e
	fmt.Println(utf8.RuneCountInString(word)) // 6

	for pos, char := range word {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
		// character S starts at byte position 0
		// character t starts at byte position 1
		// character r starts at byte position 2
		// character a starts at byte position 3
		// character ß starts at byte position 4
		// character e starts at byte position 6
	}
}
