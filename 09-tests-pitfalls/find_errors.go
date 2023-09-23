package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func findErrorsIn(file string) ([]string, error) {
	handle, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer handle.Close()

	result := make([]string, 0)

	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "error") {
			result = append(result, text)
		}
	}
	return result, nil
}

func findErrorsInWithIO(file string) ([]string, error) {
	handle, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer handle.Close()

	return findErrorsInWithoutIO(handle)
}

func findErrorsInWithoutIO(reader io.Reader) ([]string, error) {
	result := make([]string, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "error") {
			result = append(result, text)
		}
	}
	return result, nil
}
