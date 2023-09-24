package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// preemptive/cooperative multitasking

func main() {
	ctxBackground := context.Background()
	//ctxTodo := context.TODO()

	ctx := context.WithValue(ctxBackground, "value", 1)
	logger := logrus.New()
	ctx = context.WithValue(ctx, "logger", logger)

	ctx, cancel := context.WithTimeout(ctx, time.Hour)
	defer cancel()

	ctx, cancel2 := context.WithCancel(ctx)
	defer cancel2()

	http.HandleFunc("/", handleRequest)

	fmt.Println("Server is running...")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		panic(err)
	}
}

// curl localhost:4000
func handleRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handler started")
	ctx := req.Context()
	requestID := ctx.Value("request-id")
	if requestID == nil {
		ctx = context.WithValue(ctx, "request-id", uuid.New().String())
	}

	select {
	case <-time.After(5 * time.Second):
		_, _ = fmt.Fprintf(w, "Response from the server")

	// Handling request cancellation
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("Error:", err)
	}

	fmt.Println("Handler complete")
}
