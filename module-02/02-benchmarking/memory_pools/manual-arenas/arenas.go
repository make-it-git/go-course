package main

import (
	"sync"
)

type Arena struct {
	buffer []byte // This is the memory buffer simulating the arena
	offset int    // Offset into the arena where the next allocation should occur
	size   int    // Size of each allocation block
	mu     sync.Mutex
}

func NewArena(size, blockSize int) *Arena {
	return &Arena{
		buffer: make([]byte, size),
		size:   blockSize,
	}
}

// Alloc allocates a block of memory from the arena
func (a *Arena) Alloc() []byte {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.offset+a.size > len(a.buffer) {
		panic("Arena out of memory")
	}

	block := a.buffer[a.offset : a.offset+a.size]
	a.offset += a.size
	return block
}

// Reset resets the arena, allowing it to be reused
func (a *Arena) Reset() {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.offset = 0
}
