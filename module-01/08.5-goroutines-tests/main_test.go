package main

import (
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
	"time"
)

func TestQuorumWrite(t *testing.T) {
	qw := NewQuorumWriter(3)

	var wg sync.WaitGroup
	resultChan := make(chan bool, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go qw.Write(&wg, resultChan, i)
	}

	require.Eventually(t, func() bool {
		v := qw.IsQuorumAchieved()
		t.Logf("Is quorum achieved? %v", v)
		return v // проверяем, достигнут ли кворум
	}, 2*time.Second, 200*time.Millisecond, "Expected quorum to be achieved")

	wg.Wait()

	require.Equal(t, 3, qw.count, "Expected 3 successful writes")
}

func TestQuorumNotAchieved(t *testing.T) {
	qw := NewQuorumWriter(4)

	var wg sync.WaitGroup
	resultChan := make(chan bool, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go qw.Write(&wg, resultChan, i)
	}

	require.Eventually(t, func() bool {
		return !qw.IsQuorumAchieved() // проверяем, что кворум не достигнут
	}, 2*time.Second, 100*time.Millisecond, "Expected quorum not to be achieved")

	wg.Wait()

	require.False(t, qw.IsQuorumAchieved(), "Expected quorum not to be achieved")
}
