package main

import (
	"errors"
	"fmt"
	"reflect"
)

func SumInt8(a, b int8) int8 {
	return a + b
}

func SumInt32(a, b int32) int32 {
	return a + b
}

func SumInt(a, b interface{}) (interface{}, error) {
	if reflect.TypeOf(a).Kind() != reflect.Int8 &&
		reflect.TypeOf(a).Kind() != reflect.Int32 &&
		reflect.TypeOf(a).Kind() != reflect.Int64 {
		return nil, errors.New("unsupported type for a")
	}
	if reflect.TypeOf(b).Kind() != reflect.Int8 &&
		reflect.TypeOf(b).Kind() != reflect.Int32 &&
		reflect.TypeOf(b).Kind() != reflect.Int64 {
		return nil, errors.New("unsupported type for b")
	}

	valA := reflect.ValueOf(a).Int()
	valB := reflect.ValueOf(b).Int()

	return valA + valB, nil
}

func noGenerics() {
	var a32, b32 int32 = 10, 20
	c32, _ := SumInt(a32, b32)
	fmt.Printf("Type of c32 %T, value=%d\n", c32, c32)

	var anotherInt32 int64 = c32.(int64)
	_ = anotherInt32
}
