package main

import (
	"fmt"
	"math/rand"
	"time"

	"08-tests/aggregator"
	"08-tests/math"
)

func main() {
	result := math.AddWithError(2, 2)
	fmt.Println(result) // 4
	result = math.AddWithError(100, 10)
	fmt.Println(result) // 0

	data, errors := aggregator.FetchData(rand.New(rand.NewSource(time.Now().Unix())))
	fmt.Println(data)
	fmt.Println(errors)
	// map[ServiceA:Data from Service A ServiceB:Data from Service B]
	// one or more services failed: [service C error]
}
