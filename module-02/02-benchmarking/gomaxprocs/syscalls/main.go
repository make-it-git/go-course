package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// docker build -t syscall-example . && docker run --rm -it syscall-example
// strace -c -f ./app 100
// strace -c -f ./app 9999
// https://man7.org/linux/man-pages/man2/nanosleep.2.html
// https://github.com/golang/go/issues/25471
func main() {
	max_, _ := strconv.Atoi(os.Args[1])

	if max_ == 9999 {
		fmt.Println("Sleeping for 15 seconds")
		time.Sleep(time.Second * 15)
		return
	}

	n := 0
	for {
		time.Sleep(time.Second / 100)
		n += 1
		if n >= max_ {
			return
		}
	}
}
