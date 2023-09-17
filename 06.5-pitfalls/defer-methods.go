package main

import (
	"fmt"
)

type Data struct {
	value int
}

func (d Data) PrintValue() {
	fmt.Println(d.value)
}

func (d *Data) PrintValuePtr() {
	fmt.Println(d.value)
}

func deferMethods() {
	c := Data{value: 123}
	defer c.PrintValue() // 123
	c.value = 456

	c2 := Data{value: 123}
	defer c2.PrintValuePtr() // 456
	c2.value = 456
}
