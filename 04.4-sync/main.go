package main

import (
	"fmt"
	"sync"
	"time"
)

func process() {
	time.Sleep(time.Millisecond)
	fmt.Println("process")
}

func processWithWg(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond)
	fmt.Println("process with wg")
}

func main() {
	waitGroupExample()
	waitGroupExampleCopy()
	mutexExample()
	rwMutexExample()
	syncMapExample()
}

func waitGroupExample() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		process()
	}()

	wg.Wait()
}

func waitGroupExampleCopy() {
	var wg sync.WaitGroup

	wg.Add(1)
	go processWithWg(&wg)

	wg.Wait()
}

type Counter struct {
	m     sync.Mutex
	value int
}

func (c *Counter) Update(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	c.m.Lock()
	defer c.m.Unlock()
	c.value += n
}

func mutexExample() {
	var wg sync.WaitGroup

	c := Counter{}

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go c.Update(10, &wg)
	}

	wg.Wait()
	fmt.Printf("Result is %d\n", c.value) // 1000
}

type RWCounter struct {
	m     sync.RWMutex
	value int
}

func (c *RWCounter) Update(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	c.m.Lock()
	c.value += n
	c.m.Unlock()
}

func (c *RWCounter) GetValue() int {
	c.m.RLock()
	v := c.value
	defer c.m.RUnlock()

	return v
}

func rwMutexExample() {
	var wg sync.WaitGroup

	c := RWCounter{}

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go c.Update(10, &wg)
	}

	wg.Wait()
	fmt.Printf("Result is %d\n", c.GetValue()) // 1000
}

func syncMapExample() {
	var wg sync.WaitGroup
	var m sync.Map
	// m2 := m // assignment copies lock value to m2: sync.Map contains sync.Mutex

	wg.Add(10)

	for i := 1; i <= 5; i++ {
		go func(k int) {
			v := fmt.Sprintf("value %v", k)

			fmt.Println("Writing:", v)
			m.Store(k, v)
			wg.Done()
		}(i)
	}

	for i := 1; i <= 5; i++ {
		go func(k int) {
			v, _ := m.Load(k)
			fmt.Println("Reading: ", v)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
