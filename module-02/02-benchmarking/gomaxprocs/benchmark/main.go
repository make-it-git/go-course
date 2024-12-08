package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func simulateSyscall(t time.Duration) {
	time.Sleep(t)
}

func main() {
	numCpuCores := []int{1, 2, 4, 8, 16, 32, 64, 128}
	durations := []time.Duration{
		time.Second, time.Second * 3, time.Millisecond,
	}
	for _, duration := range durations {
		fmt.Println("----------------")
		fmt.Printf("DURATION=%v\n", duration)
		for _, procs := range numCpuCores {
			fmt.Printf("GOMAXPROCS=%d, ", procs)
			start := time.Now()
			run(procs, duration)
			end := time.Now()
			fmt.Println(end.Sub(start))
		}
	}

}

func run(procs int, t time.Duration) {
	runtime.GOMAXPROCS(procs)
	var wg sync.WaitGroup

	for i := 0; i < 300_000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			simulateSyscall(t)
		}()
	}
	wg.Wait()
}
