package main

import (
	"fmt"
	"sync"
	"time"
)

// Without mutex - UNSAFE
type UnsafeCounter struct {
	value int
}

func (c *UnsafeCounter) Increment() {
	c.value++ // This operation is not atomic!
}

// With mutex - SAFE
type SafeCounter struct {
	value int
	mutex sync.Mutex
}

func (c *SafeCounter) Increment() {
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
}

func main() {
	// Unsafe counter demo
	unsafe := &UnsafeCounter{value: 0}
	for i := 0; i < 1000; i++ {
		go unsafe.Increment()
	}
	time.Sleep(time.Second)
	fmt.Printf("Unsafe counter (should be 1000): %d\n", unsafe.value)

	// Safe counter demo
	safe := &SafeCounter{value: 0}
	for i := 0; i < 1000; i++ {
		go safe.Increment()
	}
	time.Sleep(time.Second)
	fmt.Printf("Safe counter (will be 1000): %d\n", safe.value)
}
