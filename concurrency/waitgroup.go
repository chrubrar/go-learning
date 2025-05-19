// Package concurrency provides examples of Go's concurrency primitives and patterns.
package concurrency

import (
	"fmt"
	"sync"
)

// AddWithWaitGroup demonstrates a simple counter protected by a WaitGroup.
// It increments a number using multiple goroutines and waits for all of them to complete.
func AddWithWaitGroup(numGoroutines int) int32 {
	var num int32 = 0
	wg := new(sync.WaitGroup)
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(w *sync.WaitGroup, n *int32) {
			defer w.Done()
			*n++
		}(wg, &num)
	}

	wg.Wait()
	return num
}

// DemonstrateWaitGroup shows how to use the WaitGroup example
func DemonstrateWaitGroup() {
	result := AddWithWaitGroup(1000)
	fmt.Printf("Final counter value: %d\n", result)
}
