package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// go test .
// ok      09-tests-pitfalls       0.110s
func TestFindErrorsWithIO(t *testing.T) {
	t.Run("Test find errors", func(t *testing.T) {
		result, err := findErrorsIn("sample.txt")
		if err != nil {
			t.Errorf("findErrorsIn() error = %v", err)
		}

		require.Equal(t, []string{"some error", "another error"}, result)
	})
}

func TestFindErrorsWithoutIO(t *testing.T) {
	t.Run("Test find errors", func(t *testing.T) {
		reader := strings.NewReader("one\ntwo\nsome error\nthree")
		result, err := findErrorsInWithoutIO(reader)
		if err != nil {
			t.Errorf("findErrorsInWithoutIO() error = %v", err)
		}

		require.Equal(t, []string{"some error"}, result)
	})
}
