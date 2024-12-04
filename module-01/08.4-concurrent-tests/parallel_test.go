package example

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func add(a int, b int) int {
	time.Sleep(time.Second)
	return a + b
}

func TestAdd(t *testing.T) {
	testCases := map[string]struct {
		a        int
		b        int
		expected int
	}{
		"0+0": {
			a:        0,
			b:        0,
			expected: 0,
		},
		"0+1": {
			a:        0,
			b:        1,
			expected: 1,
		},
		"100+500": {
			a:        100,
			b:        500,
			expected: 600,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			result := add(tc.a, tc.b)
			require.Equal(t, tc.expected, result)
		})
	}
}
