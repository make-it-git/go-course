package main

import (
	"fmt"
	"log"
)

// type iface struct {
//	 tab  *itab
//	 data unsafe.Pointer
// }
//}

type CustomError struct{}

func (e *CustomError) Error() string {
	return "custom error"
}

func nonNilErrorInterface() {
	var typed *CustomError = nil
	var err error = typed
	fmt.Println(typed == nil) // true
	if err != nil {
		log.Println(err) // custom error
	}
}
