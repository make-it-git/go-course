package main

import (
	"fmt"
	"sync"
	"time"
)

type QuorumWriter struct {
	required int
	mu       sync.Mutex
	count    int
}

func NewQuorumWriter(required int) *QuorumWriter {
	return &QuorumWriter{
		required: required,
	}
}

func (qw *QuorumWriter) Write(wg *sync.WaitGroup, resultChan chan<- bool, id int) {
	defer wg.Done()

	// Имитация записи, которая может быть успешной или неуспешной
	time.Sleep(time.Second)
	success := id%2 == 0 // Успех только для четных id

	// Если запись успешна, увеличиваем счётчик успешных записей
	if success {
		qw.mu.Lock()
		qw.count++
		qw.mu.Unlock()
	}

	resultChan <- success
}

func (qw *QuorumWriter) IsQuorumAchieved() bool {
	qw.mu.Lock()
	defer qw.mu.Unlock()
	return qw.count >= qw.required
}

func main() {
	qw := NewQuorumWriter(3) // Требуем 3 успешных записи из 5

	var wg sync.WaitGroup
	resultChan := make(chan bool, 5)

	// Запускаем несколько горутин для записи
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go qw.Write(&wg, resultChan, i)
	}

	wg.Wait()

	if qw.IsQuorumAchieved() {
		fmt.Println("Quorum achieved!")
	} else {
		fmt.Println("Quorum not achieved.")
	}
}
