package aggregator

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func shouldFail(r rand.Source) bool {
	min := int64(1)
	max := int64(100)
	val := min + (r.Int63() % (max - min + 1))
	return val < 10
}

func fetchFromServiceA(r rand.Source) (string, error) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	if shouldFail(r) {
		return "", errors.New("service A error")
	}
	return "Data from Service A", nil
}

func fetchFromServiceB(r rand.Source) (string, error) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	if shouldFail(r) {
		return "", errors.New("service B error")
	}
	return "Data from Service B", nil
}

func fetchFromServiceC(r rand.Source) (string, error) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	if shouldFail(r) {
		return "", errors.New("service C error")
	}
	return "Data from Service C", nil
}

func FetchData(r rand.Source) (map[string]string, error) {
	var mu sync.Mutex
	results := make(map[string]string)
	errorsOccurred := make([]error, 0)

	var wg sync.WaitGroup
	services := map[string]func(r rand.Source) (string, error){
		"ServiceA": func(r rand.Source) (string, error) {
			return fetchFromServiceA(r)
		},
		"ServiceB": func(r rand.Source) (string, error) {
			return fetchFromServiceB(r)
		},
		"ServiceC": func(r rand.Source) (string, error) {
			return fetchFromServiceC(r)
		},
	}

	wg.Add(len(services))
	for name, fetchFunc := range services {
		go func(name string, fetchFunc func(r rand.Source) (string, error)) {
			defer wg.Done()
			data, err := fetchFunc(r)
			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				errorsOccurred = append(errorsOccurred, err)
				return
			}
			results[name] = data
		}(name, fetchFunc)
	}

	wg.Wait()

	if len(errorsOccurred) > 0 {
		return results, fmt.Errorf("one or more services failed: %v", errorsOccurred)
	}
	return results, nil
}
