package main

import (
	"fmt"
	"sync"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("%d ", i) // 9 or 10
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("\n")

	var wg2 sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func(i int) {
			fmt.Printf("%d ", i)
			wg2.Done()
		}(i) // copy of i
	}

	wg2.Wait()
}
