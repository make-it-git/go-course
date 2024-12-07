package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

// go vet
// go vet module-02/03-linters/main.go

// golangci-lint
// curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.62.2
// golangci-lint run module-02/03-linters/main.go

// https://github.com/mgechev/revive
// go install github.com/mgechev/revive@latest
// ~/go/bin/revive module-02/03-linters/main.go

// https://github.com/securego/gosec
// curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.21.4
// ~/go/bin/gosec module-02/03-linters

// https://staticcheck.dev/
// go install honnef.co/go/tools/cmd/staticcheck@latest
// ~/go/bin/staticcheck ./module-02/03-linters/main.go

var G = 123 // Issue: global variable naming (revive, golangci-lint)

func main() {
	go memoryLeakExample() // Issue: Potential goroutine leak (staticcheck)
	unprotectedConcurrentAccess()
	defer fmt.Println("This defer does nothing useful") // Issue: useless defer (staticcheck)
	fmt.Println(unreachableCode())                      // Issue: always returns the same value (staticcheck)
	readSensitiveFile()                                 // Issue: hardcoded sensitive file path (gosec)

	// Issue: Unused variable
	unusedVar := "This is never used"
	_ = unusedVar

	// Issue: Inefficient string concatenation in a loop (staticcheck)
	concat := ""
	for i := 0; i < 1000; i++ {
		concat += "text"
	}

	fmt.Println(concat)

	// Issue: Shadowing outer variable (golangci-lint, go vet)
	x := 10
	if x := 5; x < 10 {
		fmt.Println(x)
	}
	_ = x

	// Issue: Log.Printf without proper format specifier (revive, staticcheck)
	log.Printf("Value: %s", 42)

	// Issue: Mutex lock copied by value (go vet)
	var mu sync.Mutex
	copyMutex := mu
	copyMutex.Lock()
}

func memoryLeakExample() {
	// Issue: Goroutine without termination (staticcheck)
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
}

func unprotectedConcurrentAccess() {
	// Issue: Data race due to lack of sync mechanisms (go vet, staticcheck)
	var counter int
	for i := 0; i < 10; i++ {
		go func() {
			counter++
		}()
	}
}

func unreachableCode() int {
	// Issue: Unreachable code (golangci-lint)
	if true {
		return 42
	}
	return 0
}

func readSensitiveFile() {
	// Issue: Hardcoded file path for sensitive data (gosec)
	data, err := ioutil.ReadFile("/etc/passwd")
	if err != nil {
		fmt.Println("Error reading file")
	}
	fmt.Println(string(data))
}

func unusedFunction() {
	// Issue: Unused function (golangci-lint)
	fmt.Println("This function is not used anywhere")
}

type emptyStruct struct{}

func (e emptyStruct) String() string {
	// Issue: Useless receiver method (revive)
	return "This serves no purpose"
}
