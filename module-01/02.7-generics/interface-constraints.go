package main

import "fmt"

type Stringable interface {
	String() string
}

func convertToString[T Stringable](s []T) []string {
	var ret []string
	for _, v := range s {
		ret = append(ret, v.String())
	}

	return ret
}

type MyCustomStruct struct {
	a int
}

func (x MyCustomStruct) String() string {
	return fmt.Sprintf("Hello, a=%d", x.a)
}

type Stringer[T any] interface {
	ToString() string
}

type Person struct {
	Name  string
	Value int
}

func (p Person) ToString() string {
	return "Person: " + p.Name
}

func PrintString[T Stringer[T]](s T) {
	fmt.Println(s.ToString())
}

func interfaceConstraints() {
	slice := convertToString([]MyCustomStruct{
		{100},
		{500},
	})
	fmt.Println(slice)

	PrintString(Person{Name: "John", Value: 100500})
}
