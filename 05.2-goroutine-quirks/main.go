package main

import (
	"fmt"
	"time"
)

func main() {
	s := First("test", Search, Search)
	fmt.Println(s)

	s2 := FirstWithSelect("test", Search, Search)
	fmt.Println(s2)

	s3 := FirstWithCancel("test", Search, Search)
	fmt.Println(s3)
}

func Search(query string) string {
	time.Sleep(time.Second)
	return "Result of " + query
}

func First(query string, replicas ...func(string) string) string {
	c := make(chan string)
	// c := make(chan string, len(replicas))
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func FirstWithSelect(query string, replicas ...func(string) string) string {
	c := make(chan string, 1)
	searchReplica := func(i int) {
		select {
		case c <- replicas[i](query):
		default:
		}
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func FirstWithCancel(query string, replicas ...func(string) string) string {
	c := make(chan string)
	done := make(chan struct{})
	defer close(done)
	searchReplica := func(i int) {
		select {
		case c <- replicas[i](query):
		case <-done:
		}
	}
	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}
