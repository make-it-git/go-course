package main

import (
	"errors"
	"fmt"
)

var (
	ErrRepository = errors.New("db query error")
	ErrService    = errors.New("service error")
	ErrHandler    = errors.New("handler error")
)

func httpHandler(queryParams map[string]string) error {
	value, ok := queryParams["value"]
	if !ok {
		return fmt.Errorf("provide query param: %w", ErrHandler)
	}
	err := serviceMethod(value)
	if err != nil {
		return fmt.Errorf("service returned error: %w", err)
	}
	return nil
}

func serviceMethod(value string) error {
	if value == "service-error" {
		return ErrService
	}
	err := repositoryMethod(value)
	if err != nil {
		return fmt.Errorf("repo returned error: %w", err)
	}
	return nil
}

func repositoryMethod(value string) error {
	if value == "repo-error" {
		return ErrRepository
	}
	return nil
}

func main() {
	m := make(map[string]string)
	err := httpHandler(m)
	fmt.Println(errors.Is(err, ErrHandler)) // true

	m["value"] = "service-error"
	err = httpHandler(m)
	fmt.Println(errors.Is(err, ErrHandler), errors.Is(err, ErrService)) // false true
	fmt.Println(err)                                                    // service returned error: service error

	m["value"] = "repo-error"
	err = httpHandler(m)
	fmt.Println(errors.Is(err, ErrHandler), errors.Is(err, ErrService), errors.Is(err, ErrRepository)) // false false true
	fmt.Println(err)                                                                                   // service returned error: repo returned error: db query error

	err = fmt.Errorf("wrapped err: %w", ErrRepository)
	fmt.Println(err)                                 // wrapped err: db query error
	fmt.Println(errors.Unwrap(err))                  // db query error
	fmt.Println(errors.Unwrap(err) == ErrRepository) // true
}
