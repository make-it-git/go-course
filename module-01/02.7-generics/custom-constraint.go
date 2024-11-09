package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type MyStruct1 struct {
	value string
	a     int
}

func (x MyStruct1) Value() string {
	return x.value
}

type MyStruct2 struct {
	value string
	b     float32
}

func (x MyStruct2) Value() string {
	return x.value
}

type MyConstraint interface {
	Value() string
}

func test[T MyConstraint](x T) {
	fmt.Println(x.Value())
}

type MyConstraintAllStrings interface {
	constraints.Integer | ~string
}

type MyConstraintStrictStrings interface {
	constraints.Integer | string
}

func allStrings[T MyConstraintAllStrings](x T) {
	fmt.Println(x)
}

func strictStrings[T MyConstraintStrictStrings](x T) {
	fmt.Println(x)
}

type MyString string

func withSlice[S interface{ ~[]E }, E interface{}](s S) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func withSlice2[S ~[]E, E interface{}](s S) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func withSlice3[S ~[]E, E any](s S) {
	for _, v := range s {
		fmt.Println(v)
	}
}

type Number interface {
	int | int8 | int16
}

type Vector[T Number] []T

func addVectors[T Number](vec1 Vector[T], vec2 Vector[T]) Vector[T] {
	var result Vector[T]
	for i := range vec1 {
		result = append(result, vec1[i]+vec2[i])
	}
	return result
}

func customConstraints() {
	test(MyStruct1{value: "test1"})
	test(MyStruct2{value: "test2"})

	allStrings("test")
	allStrings(MyString("test"))
	strictStrings("test")
	//strictStrings(MyString("test"))

	withSlice([]string{"a", "b", "c"})
	withSlice2([]int{10, 20, 30})
	withSlice3([]float32{10, 20, 30, 40.100500})

	v1 := Vector[int]{1, 2, 3}
	v2 := Vector[int]{3, 4, 5}
	fmt.Printf("addVectors(v1, v2)=%f\n", addVectors(v1, v2))
}
