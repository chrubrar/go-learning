package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// UnsafeCounter demonstrates a counter without proper synchronization.
// This type is intentionally unsafe to show race conditions.
type UnsafeCounter struct {
	Value int
}

// Increment adds 1 to the counter value without synchronization.
func (c *UnsafeCounter) Increment() {
	c.Value++ // This operation is not atomic!
}

// SafeCounter demonstrates a thread-safe counter using a mutex.
type SafeCounter struct {
	value int
	mutex sync.Mutex
}

// Increment adds 1 to the counter value in a thread-safe manner.
func (c *SafeCounter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

// GetValue returns the current value of the counter.
func (c *SafeCounter) GetValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

// DemonstrateMutex shows the difference between safe and unsafe counters
// when accessed concurrently.
func DemonstrateMutex() {
	// Unsafe counter demo
	unsafe := &UnsafeCounter{}
	for i := 0; i < 1000; i++ {
		go unsafe.Increment()
	}
	time.Sleep(time.Second)
	fmt.Printf("Unsafe counter (should be 1000): %d\n", unsafe.Value)

	// Safe counter demo
	safe := &SafeCounter{}
	for i := 0; i < 1000; i++ {
		go safe.Increment()
	}
	time.Sleep(time.Second)
	fmt.Printf("Safe counter (will be 1000): %d\n", safe.GetValue())
}
