package main

import (
	"fmt"
	"runtime"
)

func printMemStats() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Printf("Alloc = %v MB", memStats.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MB", memStats.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MB", memStats.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", memStats.NumGC)
}
