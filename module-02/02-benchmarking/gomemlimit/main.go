package main

func allocate100Kb() []byte {
	s := make([]byte, 1024*100)
	return s
}

// GOMEMLIMIT=1MiB go run .
func main() {
	var x []byte
	printMemStats()
	for i := 0; i < 1000_000; i++ {
		if i > 0 && i%250_000 == 0 {
			printMemStats()
		}
		x = allocate100Kb()
	}
	_ = x
	printMemStats()
}
