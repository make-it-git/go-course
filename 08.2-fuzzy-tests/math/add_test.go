package math_test

import (
	"testing"

	"08-tests/math"
)

func FuzzTestErrorAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int) {
		res := math.AddWithError(a, b)
		if res != a+b {
			t.Errorf("a=%d, b=%d, result=%d", a, b, res)
		}
	})
}

//  go test -fuzz=FuzzTestErrorAdd 08-tests/math
//fuzz: elapsed: 0s, gathering baseline coverage: 0/1 completed
//failure while testing seed corpus entry: FuzzTestErrorAdd/34084481b2afdc97
//fuzz: elapsed: 0s, gathering baseline coverage: 0/1 completed
//--- FAIL: FuzzTestErrorAdd (0.01s)
//    --- FAIL: FuzzTestErrorAdd (0.00s)
//        add_test.go:15: a=0, b=10, result=0
//
//FAIL
//exit status 1
//FAIL    08-tests/math   0.248s
