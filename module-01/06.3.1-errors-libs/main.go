package main

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

var Err1 = errors.New("error 1")
var Err2 = errors.New("error 2")

type CustomErr struct {
}

func (c CustomErr) Error() string {
	return "custom error"
}

func step1() error {
	return Err1
}

func step2() error {
	return Err2
}

func step3() error {
	return &CustomErr{}
}

func main() {
	var result error

	if err := step1(); err != nil {
		result = multierror.Append(result, err)
	}
	if err := step2(); err != nil {
		result = multierror.Append(result, err)
	}
	if err := step3(); err != nil {
		result = multierror.Append(result, err)
	}

	fmt.Println(result)
	// 3 errors occurred:
	//        * error 1
	//        * error 2
	//        * custom error

	if mutliErr, ok := result.(*multierror.Error); ok {
		for _, e := range mutliErr.Errors {
			if errors.Is(e, Err2) {
				fmt.Println("retry step2 may be?")
			}
		}
	}

	var customErr *CustomErr
	if errors.As(result, &customErr) {
		fmt.Println(customErr) // custom error
	}

	if errors.Is(result, Err2) {
		fmt.Println("We have err 2 in list")
	}

	var multiErr *multierror.Error
	multiErr = multierror.Append(multiErr, Err1)
	fmt.Println(multiErr.ErrorOrNil())
}
