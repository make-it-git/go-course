package main

import (
	"fmt"
	"io"
	"os"
)

type NetworkSender struct {
}

func (s *NetworkSender) Write(p []byte) (n int, err error) {
	return fmt.Printf("Network send: %s", string(p))
}

type SampleInterface interface {
	Method() string
}

type OtherInterface interface {
	SampleInterface
}

type BidirectionalCommunication struct {
}

func (b *BidirectionalCommunication) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (b *BidirectionalCommunication) Write(p []byte) (n int, err error) {
	return 0, nil
}

type SampleStruct struct {
}

func (s *SampleStruct) Method() string {
	return ""
}

func main() {
	var w io.Writer
	w = os.Stdout
	_, _ = w.Write([]byte("Hello, world\n")) // Hello, world

	w = &NetworkSender{}
	_, _ = w.Write([]byte("Hello, world\n")) // Network send: Hello, world

	// differences
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s) // hello

	// number := i.(int) // panic: interface conversion: interface {} is string, not int
	number, ok := i.(int)
	if !ok {
		fmt.Printf("Type %T is not number, number=%d\n", i, number) // Type string is not number, number=0
	} else {
		fmt.Println(number)
	}

	// type conversion
	intVar := string(97)
	fmt.Println(intVar) // a

	var int32Var int32 = 100500
	var int64Var int64
	int64Var = int64(int32Var)
	fmt.Println(int64Var) // 100500

	// type switch
	var t interface{} = "hello"
	switch t := t.(type) {
	case string:
		fmt.Printf("string: %s\n", t) // string: hello
	case bool:
		fmt.Printf("boolean: %v\n", t)
	case int:
		fmt.Printf("integer: %d\n", t)
	default:
		fmt.Printf("unexpected: %T\n", t)
	}

	// zero value
	var sample SampleInterface
	fmt.Println(sample, sample == nil) // <nil> true

	var x io.ReadWriter = &BidirectionalCommunication{}
	_, _ = x.Read([]byte{})
	_, _ = x.Write([]byte{})

	var sampleInstance *SampleStruct = nil
	sample = sampleInstance
	fmt.Println(sample, sample == nil) // <nil> false
}
