package main

import (
	"fmt"
)

type Request struct {
	Payload string
}

func server(c <-chan Request) {
	for work := range c {
		go safelyDo(work)
	}
}

func safelyDo(r Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("request error:", err)
		}
	}()
	do(r)
}

func do(r Request) {
	fmt.Println("handle request", r.Payload)
	if r.Payload == "do-panic" {
		panic("failed to process this request")
	}
	fmt.Println("done handling request", r.Payload)
}

// Note: random output order

// handle request success 1
// done handling request success 1

// handle request do-panic
// request error: failed to process this request

// handle request success 2
// done handling request success 2

func safeRecover() {
	c := make(chan Request)
	go server(c)
	c <- Request{Payload: "success 1"}
	c <- Request{Payload: "do-panic"}
	c <- Request{Payload: "success 2"}
	defer close(c)
}
