package aggregator

import (
	"math/rand"
	"testing"
)

func FuzzFetchData(f *testing.F) {
	// Добавим фиксированный seed для воспроизводимости
	f.Add(1)

	f.Fuzz(func(t *testing.T, seed int) {
		r := rand.New(rand.NewSource(int64(seed)))

		results, err := FetchData(r)

		if err != nil && len(results) == 0 {
			t.Errorf("expected some data even with errors, got none, seed: %d", seed)
		}

		for key := range results {
			if key != "ServiceA" && key != "ServiceB" && key != "ServiceC" {
				t.Errorf("unexpected key in results: %s, seed: %d", key, seed)
			}

		}
	})
}

// go test -fuzz=FuzzFetchData -fuzztime=30s aggregator/*.go
//fuzz: elapsed: 0s, gathering baseline coverage: 0/23 completed
//fuzz: elapsed: 0s, gathering baseline coverage: 23/23 completed, now fuzzing with 14 workers
//fuzz: elapsed: 3s, execs: 530 (177/sec), new interesting: 0 (total: 23)
//fuzz: elapsed: 4s, execs: 785 (171/sec), new interesting: 0 (total: 23)
//--- FAIL: FuzzFetchData (4.49s)
//    --- FAIL: FuzzFetchData (0.09s)
//        aggregate_test.go:18: expected some data even with errors, got none, seed: -171
//
//    Failing input written to testdata/fuzz/FuzzFetchData/dc825facf128513c
//    To re-run:
//    go test -run=FuzzFetchData/dc825facf128513c
