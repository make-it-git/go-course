package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Data struct {
	Value int
}

const N = 500000

func main() {
	var counter int64
	atomic.AddInt64(&counter, 1)

	// CAS
	cas()
	// Pointers
	pointersAtomic()
	pointersNonAtomic()
}

func cas() {
	var value int32 = 42
	success := atomic.CompareAndSwapInt32(&value, 42, 100500)
	if success {
		fmt.Println("Value was 42, now it's 100500")
	} else {
		fmt.Println("Value wasn't 42, no change made")
	}
}

func pointersAtomic() {
	var ptr atomic.Pointer[Data]
	ptr.Store(&Data{Value: 0}) // Initialize pointer

	wg := sync.WaitGroup{}
	// Goroutine 1: Increment the value 10 times
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= N; i++ {
			newValue := &Data{Value: i}
			ptr.Store(newValue)
		}
	}()

	// Goroutine 2: Increment the value 10 times
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= N; i++ {
			newValue := &Data{Value: i}
			ptr.Store(newValue)
		}
	}()

	wg.Wait()

	finalValue := ptr.Load()
	fmt.Println("Final Value (with atomic):", finalValue.Value)
}

func pointersNonAtomic() {
	var ptr *Data = &Data{Value: 0}

	wg := sync.WaitGroup{}

	// Goroutine 1: Increment the value 10 times
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= N; i++ {
			newValue := &Data{Value: i}
			ptr = newValue // Not atomic!
		}
	}()

	// Goroutine 2: Increment the value 10 times
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= N; i++ {
			newValue := &Data{Value: i}
			ptr = newValue // Not atomic!
		}
	}()

	wg.Wait()

	// Final value of the pointer
	fmt.Println("Final Value (without atomic):", ptr.Value)
}
