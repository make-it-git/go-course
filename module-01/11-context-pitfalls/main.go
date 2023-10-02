package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type oneParam string
type otherParam string

func main() {
	// context values overriding
	ctx := context.Background() // root of all derived contexts
	ctx = context.WithValue(ctx, "param", 10)
	fmt.Println(ctx.Value("param")) // 10
	ctx = context.WithValue(ctx, "param", 20)
	fmt.Println(ctx.Value("param")) // 20

	ctx2 := context.Background()
	var param1 oneParam = "param"
	ctx2 = context.WithValue(ctx2, param1, 10)
	var param2 otherParam = "param"
	ctx2 = context.WithValue(ctx2, param2, 20)
	fmt.Println(ctx2.Value(param1)) // 10
	fmt.Println(ctx2.Value(param2)) // 20

	// advisory cancellation
	ctx3, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctx3.Done():
			fmt.Println("Got cancel signal")
			time.Sleep(time.Second * 3)
			fmt.Println("cancelled")
		}
	}()
	cancel()

	// explicit better than implicit
	ctx4 := context.Background()
	ctx4 = context.WithValue(ctx4, "importantParam", 123)
	myMethodImplicit(ctx4)
	// better
	myMethodExplicit(ctx4, 123)

	// explicit context value usage
	// addRequestID()
	// useRequestID()

	// context cancel propagation
	ctx5, cancel5 := context.WithCancel(context.Background())
	go cancellableMethod(ctx5)
	cancel5()

	networkRequest()

	time.Sleep(time.Second * 5)
}

func addRequestID(rw http.ResponseWriter, req *http.Request, next http.Handler) {
	ctx := context.WithValue(req.Context(), "request-id", req.Header.Get("request-id"))
	req = req.WithContext(ctx)
	next.ServeHTTP(rw, req)
}

func useRequestID(rw http.ResponseWriter, req *http.Request, next http.Handler) {
	requestID := req.Context().Value("request-id")
	fmt.Println("send request to other service", requestID)
}

func myMethodImplicit(ctx context.Context) {
	val := ctx.Value("importantParam")
	fmt.Println("importantParam", val) // importantParam 123
}

func myMethodExplicit(ctx context.Context, importantParam int) {
	fmt.Println("importantParam", importantParam) // importantParam 123
}

func cancellableMethod(ctx context.Context) {
	ctx = context.WithValue(ctx, "param", "value")
	cancellableMethod2(ctx)
}

func cancellableMethod2(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("cancelled inner context too")
	}
}

func networkRequest() {
	// Get "https://example.com": dial tcp: lookup example.com: i/o timeout
	const timeout = 1 * time.Second

	client := http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: timeout,
			}).DialContext,
		},
	}

	// Get "https://example.com": context deadline exceeded
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	rsp, err := client.Do(req)
	if rsp != nil {
		defer rsp.Body.Close()
	}
	if e, ok := err.(net.Error); ok && e.Timeout() {
		fmt.Printf("Do request timeout: %s\n", err)
		return
	} else if err != nil {
		fmt.Printf("Cannot do request: %s\n", err)
		return
	}

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("body: %s\n", body)
}
