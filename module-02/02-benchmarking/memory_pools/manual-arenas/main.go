package main

import (
	"fmt"
	"runtime"
	"sync"
)

func example() {
	printMemStats()
	blockSize := 1024
	arena := NewArena(1024*1024*100, blockSize) // 100MB, block 1kb
	printMemStats()

	var wg sync.WaitGroup
	numWorkers := 5
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			defer wg.Done()

			block := arena.Alloc()
			fmt.Printf("Worker %d allocated a block: %p, size %d bytes\n", id, block, len(block))
			// Address diff between memory blocks equals blockSize
			for i := range block {
				block[i] = byte(id + 10) // Fill block with sample data
			}
		}(i)
	}

	wg.Wait()

	arena.Reset()
	block := arena.Alloc()
	fmt.Printf("Value %v\n", block[0]) // Not always the same worker, memory is not cleared
	fmt.Printf("Value %v\n", block[blockSize-1])
	// fmt.Printf("Value %v\n", block[blockSize]) // panic

	printMemStats()
}

func main() {
	example()

	fmt.Println("example() completed")

	printMemStats()
	runtime.GC()
	fmt.Println("GC completed")
	printMemStats()

	fmt.Println("Reallocate arena")
	arena := NewArena(1024*1024*100, 1024)
	_ = arena
	printMemStats()
	runtime.GC()
	printMemStats()
	runtime.KeepAlive(arena)
}
