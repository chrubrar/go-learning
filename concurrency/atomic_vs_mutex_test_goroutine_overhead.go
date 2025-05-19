package concurrency

import (
	"sync"
	"sync/atomic"
	"testing"
)

// This version shows the overhead of creating goroutines
// compared to the more efficient RunParallel approach

func BenchmarkAtomicWithGoroutines(b *testing.B) {
	var num int32
	var wg sync.WaitGroup
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&num, 1)
		}()
	}
	wg.Wait()
}

func BenchmarkMutexWithGoroutines(b *testing.B) {
	var num int32
	var mu sync.Mutex
	var wg sync.WaitGroup
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			num++
			mu.Unlock()
		}()
	}
	wg.Wait()
}

// Key differences from atomic_vs_mutex_test.go:
// 1. Creates a new goroutine for each operation (more overhead)
// 2. Uses WaitGroup instead of b.RunParallel
// 3. Less efficient as each operation requires goroutine creation
// 4. More memory usage due to goroutine allocation
