package main

import (
	"fmt"
)

func echo(arg string) {
	fmt.Println(arg)
}

// Don't communicate by sharing memory, share memory by communicating
func main() {
	go echo("Hello World")
	// time.Sleep(time.Millisecond)
}
