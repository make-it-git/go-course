package main

import (
	"fmt"
)

func echo(arg string) {
	fmt.Println(arg)
}

func main() {
	go echo("Hello World")
	// time.Sleep(time.Millisecond * 10)
}
