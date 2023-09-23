package main

import (
	"testing"
)

func TestTLog(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		value int
	}{
		{name: "test 1", value: 1},
		{name: "test 2", value: 2},
		{name: "test 3", value: 3},
		{name: "test 4", value: 4},
	}
	for _, tc := range tests {
		// tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()
			t.Log(tc.value)
		})
	}
}

// go test -v ./table_driven_test.go
//=== RUN   TestTLog
//=== PAUSE TestTLog
//=== CONT  TestTLog
//=== RUN   TestTLog/test_1
//    table_driven_test.go:22: 1
//=== RUN   TestTLog/test_2
//    table_driven_test.go:22: 2
//=== RUN   TestTLog/test_3
//    table_driven_test.go:22: 3
//=== RUN   TestTLog/test_4
//    table_driven_test.go:22: 4
//--- PASS: TestTLog (0.00s)
//    --- PASS: TestTLog/test_1 (0.00s)
//    --- PASS: TestTLog/test_2 (0.00s)
//    --- PASS: TestTLog/test_3 (0.00s)
//    --- PASS: TestTLog/test_4 (0.00s)
//PASS
//ok      command-line-arguments  0.233s

//go test -v -run TestTLog/test_1 ./table_driven_test.go
//=== RUN   TestTLog
//=== PAUSE TestTLog
//=== CONT  TestTLog
//=== RUN   TestTLog/test_1
//    table_driven_test.go:22: 1
//--- PASS: TestTLog (0.00s)
//    --- PASS: TestTLog/test_1 (0.00s)
//PASS
//ok      command-line-arguments  0.110s

//go test -v ./table_driven_test.go
//=== RUN   TestTLog
//=== PAUSE TestTLog
//=== CONT  TestTLog
//=== RUN   TestTLog/test_1
//=== PAUSE TestTLog/test_1
//=== RUN   TestTLog/test_2
//=== PAUSE TestTLog/test_2
//=== RUN   TestTLog/test_3
//=== PAUSE TestTLog/test_3
//=== RUN   TestTLog/test_4
//=== PAUSE TestTLog/test_4
//=== CONT  TestTLog/test_1
//=== CONT  TestTLog/test_3
//=== CONT  TestTLog/test_2
//=== NAME  TestTLog/test_3
//    table_driven_test.go:21: 4
//=== NAME  TestTLog/test_1
//    table_driven_test.go:21: 4
//=== CONT  TestTLog/test_4
//    table_driven_test.go:21: 4
//=== NAME  TestTLog/test_2
//    table_driven_test.go:21: 4
//--- PASS: TestTLog (0.00s)
//    --- PASS: TestTLog/test_3 (0.00s)
//    --- PASS: TestTLog/test_1 (0.00s)
//    --- PASS: TestTLog/test_4 (0.00s)
//    --- PASS: TestTLog/test_2 (0.00s)
//PASS
//ok      command-line-arguments  0.227s
